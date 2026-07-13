// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 绑定上游身份提供商
//
// @param request - BindIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindIdentityProviderResponse
func (client *Client) BindIdentityProviderWithContext(ctx context.Context, request *BindIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *BindIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		query["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.IdpMetadata) {
		query["IdpMetadata"] = request.IdpMetadata
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoginEnabled) {
		query["LoginEnabled"] = request.LoginEnabled
	}

	if !dara.IsNil(request.SyncEnabled) {
		query["SyncEnabled"] = request.SyncEnabled
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindIdentityProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为指定AgentTeams实例异步开通并配置阿里云公网NAT网关。
//
// Description:
//
// ## 请求说明
//
// - 本接口用于为特定的AgentTeams实例创建公网NAT网关，并自动完成EIP申请、绑定以及SNAT规则的设置。
//
// - 接口调用后将返回一个异步任务ID，实际的NAT网关、EIP及SNAT资源ID会在异步任务完成后通过任务结果提供。
//
// - NAT网关名称由系统自动生成，格式为`magic-create-for-vpc-{vpcId}`。
//
// - 支持GET和POST方法进行请求。
//
// - `eipBandwidth`参数指定了自动申请EIP时的带宽大小，默认值为5Mbps，范围在1-200Mbps之间。
//
// - 如果`instanceId`为空或无效，或者提供的`eipBandwidth`不在允许范围内，API将返回错误响应。
//
// @param request - ConfigureNatGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureNatGatewayResponse
func (client *Client) ConfigureNatGatewayWithContext(ctx context.Context, request *ConfigureNatGatewayRequest, runtime *dara.RuntimeOptions) (_result *ConfigureNatGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EipAllocationId) {
		query["EipAllocationId"] = request.EipAllocationId
	}

	if !dara.IsNil(request.EipBandwidth) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NatGatewayInstanceId) {
		query["NatGatewayInstanceId"] = request.NatGatewayInstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureNatGateway"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureNatGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建凭证
//
// @param request - CreateCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCredentialResponse
func (client *Client) CreateCredentialWithContext(ctx context.Context, request *CreateCredentialRequest, runtime *dara.RuntimeOptions) (_result *CreateCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCredential"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建指定配置的集群实例。
//
// Description:
//
// ## 请求说明
//
// - 本接口支持通过表单参数或 query 参数传递请求信息。
//
// - `instanceSpec` 和 `networkType` 等部分参数有默认值，若未指定则使用默认值。
//
// - 必须提供 `instanceName`, `regionId`, `vpcId`, 和 `vSwitchId` 参数。
//
// - `networkType` 支持三种选项：`PRIVATE_PUBNET`, `PRIVATE_NET`, `PUB_NET`，默认为 `PRIVATE_NET`。
//
// - 如果指定了 `zoneId`，则会尝试在该可用区创建实例；否则将根据系统策略选择合适的可用区。
//
// @param tmpReq - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, tmpReq *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Zones) {
		request.ZonesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Zones, dara.String("Zones"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZonesShrink) {
		query["Zones"] = request.ZonesShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.PaymentType) {
		body["PaymentType"] = request.PaymentType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建MCP
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param tmpReq - CreateMcpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpResponse
func (client *Client) CreateMcpWithContext(ctx context.Context, tmpReq *CreateMcpRequest, runtime *dara.RuntimeOptions) (_result *CreateMcpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMcpShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Addresses) {
		request.AddressesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addresses, dara.String("Addresses"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddressesShrink) {
		body["Addresses"] = request.AddressesShrink
	}

	if !dara.IsNil(request.AuthConfig) {
		body["AuthConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.AuthEnabled) {
		body["AuthEnabled"] = request.AuthEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CreateType) {
		body["CreateType"] = request.CreateType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SwaggerConfig) {
		body["SwaggerConfig"] = request.SwaggerConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcp"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param tmpReq - CreateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
func (client *Client) CreateModelWithContext(ctx context.Context, tmpReq *CreateModelRequest, runtime *dara.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Protocols) {
		request.ProtocolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Protocols, dara.String("Protocols"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProtocolsShrink) {
		body["Protocols"] = request.ProtocolsShrink
	}

	if !dara.IsNil(request.Provider) {
		body["Provider"] = request.Provider
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	if !dara.IsNil(request.ProviderName) {
		body["ProviderName"] = request.ProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModel"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param tmpReq - CreateModelProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelProviderResponse
func (client *Client) CreateModelProviderWithContext(ctx context.Context, tmpReq *CreateModelProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateModelProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiKeys) {
		request.ApiKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiKeys, dara.String("ApiKeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Protocols) {
		request.ProtocolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Protocols, dara.String("Protocols"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		body["Address"] = request.Address
	}

	if !dara.IsNil(request.ApiKeysShrink) {
		body["ApiKeys"] = request.ApiKeysShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProtocolsShrink) {
		body["Protocols"] = request.ProtocolsShrink
	}

	if !dara.IsNil(request.Provider) {
		body["Provider"] = request.Provider
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建指定实例的Endpoint，支持多种组件和网关类型。
//
// Description:
//
// ## 请求说明
//
// - 当前controller使用的是普通参数绑定，不是`@RequestBody`，因此参数更适合按query/form方式传递。
//
// - `domain`字段会在服务端进行`trim + lowerCase`处理。
//
// - `query`和`headers`必须是JSON object字符串格式，不能为数组。
//
// - 创建操作仅将数据保存到数据库；只有在更新时，如果满足`ELEMENT/MATRIX + AI_GATEWAY + INTERNET`且域名或证书发生变化，才会触发AI Gateway域名同步逻辑。
//
// @param request - CreateServiceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceEndpointResponse
func (client *Client) CreateServiceEndpointWithContext(ctx context.Context, request *CreateServiceEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.Component) {
		query["Component"] = request.Component
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceEndpoint"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建团队
//
// @param tmpReq - CreateTeamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTeamResponse
func (client *Client) CreateTeamWithContext(ctx context.Context, tmpReq *CreateTeamRequest, runtime *dara.RuntimeOptions) (_result *CreateTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTeamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TeamMembers) {
		request.TeamMembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TeamMembers, dara.String("TeamMembers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminName) {
		query["AdminName"] = request.AdminName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TeamMembersShrink) {
		query["TeamMembers"] = request.TeamMembersShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTeam"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTeamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMethod) {
		query["AuthMethod"] = request.AuthMethod
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Worker
//
// @param tmpReq - CreateWorkerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkerResponse
func (client *Client) CreateWorkerWithContext(ctx context.Context, tmpReq *CreateWorkerRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Channels) {
		request.ChannelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Channels, dara.String("Channels"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Credentials) {
		request.CredentialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Credentials, dara.String("Credentials"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Groups) {
		request.GroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Groups, dara.String("Groups"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LimitConfig) {
		request.LimitConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LimitConfig, dara.String("LimitConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.McpServers) {
		request.McpServersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.McpServers, dara.String("McpServers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Model) {
		request.ModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Model, dara.String("Model"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Skills) {
		request.SkillsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Skills, dara.String("Skills"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Subagents) {
		request.SubagentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Subagents, dara.String("Subagents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.Agents) {
		query["Agents"] = request.Agents
	}

	if !dara.IsNil(request.ChannelsShrink) {
		query["Channels"] = request.ChannelsShrink
	}

	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.GroupsShrink) {
		query["Groups"] = request.GroupsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LimitConfigShrink) {
		query["LimitConfig"] = request.LimitConfigShrink
	}

	if !dara.IsNil(request.McpServersShrink) {
		query["McpServers"] = request.McpServersShrink
	}

	if !dara.IsNil(request.ModelShrink) {
		query["Model"] = request.ModelShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillsShrink) {
		query["Skills"] = request.SkillsShrink
	}

	if !dara.IsNil(request.Soul) {
		query["Soul"] = request.Soul
	}

	if !dara.IsNil(request.SubagentsShrink) {
		query["Subagents"] = request.SubagentsShrink
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialsShrink) {
		body["Credentials"] = request.CredentialsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorker"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Worker本地纳管启动Token
//
// @param request - CreateWorkerBootstrapTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkerBootstrapTokenResponse
func (client *Client) CreateWorkerBootstrapTokenWithContext(ctx context.Context, request *CreateWorkerBootstrapTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkerBootstrapTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkerBootstrapToken"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkerBootstrapTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除凭证
//
// @param request - DeleteCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredentialWithContext(ctx context.Context, request *DeleteCredentialRequest, runtime *dara.RuntimeOptions) (_result *DeleteCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCredential"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于释放指定的AgentTeams实例，并清理相关资源。
//
// Description:
//
// ## 请求说明
//
// - 本API支持`GET`和`POST`方法，两者语义相同。
//
// - 使用`POST`方法时，参数通过`application/x-www-form-urlencoded`格式提交。
//
// - 当前实例状态为`CREATING`、`DELETING`或`DELETED`时，请求将被拒绝。
//
// - 成功调用后，实例状态将首先更改为`DELETING`，实际的资源清理过程由后台异步执行。
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除MCP
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - DeleteMcpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpResponse
func (client *Client) DeleteMcpWithContext(ctx context.Context, request *DeleteMcpRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcp"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - DeleteModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithContext(ctx context.Context, request *DeleteModelRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模型供应商
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - DeleteModelProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelProviderResponse
func (client *Client) DeleteModelProviderWithContext(ctx context.Context, request *DeleteModelProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于删除指定AgentTeams实例下的endpoint，并清理相关资源。
//
// Description:
//
// ## 请求说明
//
// - 该接口支持通过GET或POST方法调用。
//
// - 如果目标endpoint是`WORKER`类型，系统将自动清理与之关联的APIG/AI Gateway云资源及KubeOne worker service配置。
//
// - 请求参数必须包含`instanceId`和`endpointId`，且不能为空。
//
// - 成功响应会返回HTTP状态码200以及成功标志；错误响应则根据具体情况返回相应的HTTP状态码（如400、404、409）及错误信息。
//
// @param request - DeleteServiceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceEndpointResponse
func (client *Client) DeleteServiceEndpointWithContext(ctx context.Context, request *DeleteServiceEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceEndpoint"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除团队
//
// @param request - DeleteTeamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTeamResponse
func (client *Client) DeleteTeamWithContext(ctx context.Context, request *DeleteTeamRequest, runtime *dara.RuntimeOptions) (_result *DeleteTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTeam"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTeamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Worker
//
// @param request - DeleteWorkerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkerResponse
func (client *Client) DeleteWorkerWithContext(ctx context.Context, request *DeleteWorkerRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorker"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询凭证详情
//
// @param request - GetCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialResponse
func (client *Client) GetCredentialWithContext(ctx context.Context, request *GetCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCredential"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定类型的上游身份提供商详情
//
// @param request - GetIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProviderWithContext(ctx context.Context, request *GetIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		query["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过实例ID查询指定实例的详细信息。
//
// Description:
//
// ## 请求说明
//
// - 该接口支持`GET`和`POST`方法。
//
// - 请求时必须在头部包含`X-User-Id`，用于校验实例归属。
//
// - `X-Acs-Request-Id`为可选项，如果提供，则响应中的`requestId`将优先使用此值。
//
// - 必须通过`instanceId`参数指定要查询的实例。
//
// - 成功响应会返回实例的详细配置信息及状态。
//
// - 如果请求失败，根据错误类型返回相应的HTTP状态码及错误消息。
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定AgentTeams实例关联的异步任务状态，支持分页。
//
// Description:
//
// ## 请求说明
//
// - 本接口用于查询特定AgentTeams实例下的异步任务执行状态。
//
// - 目前仅支持查询与实例生命周期相关的创建实例任务。
//
// - 可通过`taskCode`参数指定要查询的任务类型，默认为创建实例任务。
//
// - 支持使用`maxResults`和`nextToken`进行结果分页。
//
// - 当任务处于暂停(`PAUSED`)状态时，会返回用户需要采取行动的信息(`recoveryMessage`)。
//
// - 注意：当前不支持通过`taskId`直接查询任务状态。
//
// @param request - GetInstanceAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceAsyncTaskResponse
func (client *Client) GetInstanceAsyncTaskWithContext(ctx context.Context, request *GetInstanceAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskCode) {
		query["TaskCode"] = request.TaskCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceAsyncTask"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例 OSS 挂载到 ACS 所需的 RAM 授权链接。
//
// Description:
//
// ## 请求说明
//
// - 该接口支持`GET`和`POST`方法。
//
// - 请求时必须在头部包含`X-User-Id`，用于校验实例归属。
//
// - 必须通过`instanceId`参数指定实例，后端会根据实例信息生成授权链接。
//
// - 成功响应会返回 RAM 控制台授权链接，不会创建 RAM 角色或策略。
//
// @param request - GetInstanceOssMountRamAuthorizeUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceOssMountRamAuthorizeUrlResponse
func (client *Client) GetInstanceOssMountRamAuthorizeUrlWithContext(ctx context.Context, request *GetInstanceOssMountRamAuthorizeUrlRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceOssMountRamAuthorizeUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceOssMountRamAuthorizeUrl"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceOssMountRamAuthorizeUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询MCP详情
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - GetMcpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpResponse
func (client *Client) GetMcpWithContext(ctx context.Context, request *GetMcpRequest, runtime *dara.RuntimeOptions) (_result *GetMcpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcp"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模型调用摘要
//
// @param request - GetModelInvocationSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelInvocationSummaryResponse
func (client *Client) GetModelInvocationSummaryWithContext(ctx context.Context, request *GetModelInvocationSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetModelInvocationSummaryResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelInvocationSummary"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelInvocationSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型供应商详情
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - GetModelProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelProviderResponse
func (client *Client) GetModelProviderWithContext(ctx context.Context, request *GetModelProviderRequest, runtime *dara.RuntimeOptions) (_result *GetModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定实例的NAT网关及其SNAT规则的配置状态。
//
// Description:
//
// ## 请求说明
//
// 通过此API，您可以获取特定实例关联的NAT网关配置详情及SNAT规则的状态。该接口支持GET或POST方法调用，并需要提供`instanceId`作为请求参数来指定要查询的实例。
//
// ### 注意事项
//
// - 确保提供的`instanceId`是有效的且属于您的账户。
//
// - 根据返回的状态值（如`READY`, `NEED_CONFIGURE_NAT_GATEWAY`, `NEED_CONFIGURE_SNAT_RULE`），采取相应的操作以完成NAT网关或SNAT规则的配置。
//
// - 当状态为`NEED_CONFIGURE_NAT_GATEWAY`时，表示当前VPC下没有可用的NAT网关；而`NEED_CONFIGURE_SNAT_RULE`则意味着虽然存在NAT网关但某些子网CIDR未被SNAT规则覆盖。
//
// @param request - GetNatGatewayStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNatGatewayStatusResponse
func (client *Client) GetNatGatewayStatusWithContext(ctx context.Context, request *GetNatGatewayStatusRequest, runtime *dara.RuntimeOptions) (_result *GetNatGatewayStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNatGatewayStatus"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNatGatewayStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定ID的Endpoint详细信息，支持通过实例ID进行校验。
//
// Description:
//
// ## 请求说明
//
// - 该API用于根据`endpointId`查询单个Endpoint的具体配置与状态信息。
//
// - 可选参数`instanceId`用于验证Endpoint是否属于特定实例。
//
// - 请求方式支持`GET`和`POST`，其中`GET`使用query string传递参数，而`POST`则可以通过form参数提交。
//
// - 如果`endpointId`缺失或为空，则会返回`InvalidParameter`错误。
//
// - 当请求的Endpoint不存在、不属于提供的实例或者不属于当前用户时，将收到相应的资源不存在类错误响应。
//
// @param request - GetServiceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceEndpointResponse
func (client *Client) GetServiceEndpointWithContext(ctx context.Context, request *GetServiceEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetServiceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceEndpoint"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务统计摘要
//
// @param request - GetTaskStatsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatsSummaryResponse
func (client *Client) GetTaskStatsSummaryWithContext(ctx context.Context, request *GetTaskStatsSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetTaskStatsSummaryResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskStatsSummary"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatsSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询团队详情
//
// @param request - GetTeamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTeamResponse
func (client *Client) GetTeamWithContext(ctx context.Context, request *GetTeamRequest, runtime *dara.RuntimeOptions) (_result *GetTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTeam"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTeamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Token趋势统计
//
// @param request - GetTokenTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenTrendResponse
func (client *Client) GetTokenTrendWithContext(ctx context.Context, request *GetTokenTrendRequest, runtime *dara.RuntimeOptions) (_result *GetTokenTrendResponse, _err error) {
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

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTokenTrend"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工具调用分布
//
// @param request - GetToolCallDistributionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetToolCallDistributionResponse
func (client *Client) GetToolCallDistributionWithContext(ctx context.Context, request *GetToolCallDistributionRequest, runtime *dara.RuntimeOptions) (_result *GetToolCallDistributionResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToolCallDistribution"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetToolCallDistributionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户详情
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户初始密码
//
// @param request - GetUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPasswordResponse
func (client *Client) GetUserPasswordWithContext(ctx context.Context, request *GetUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *GetUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserPassword"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Worker详情
//
// @param request - GetWorkerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkerResponse
func (client *Client) GetWorkerWithContext(ctx context.Context, request *GetWorkerRequest, runtime *dara.RuntimeOptions) (_result *GetWorkerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorker"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Worker本地纳管启动选项
//
// @param request - GetWorkerBootstrapOptionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkerBootstrapOptionsResponse
func (client *Client) GetWorkerBootstrapOptionsWithContext(ctx context.Context, request *GetWorkerBootstrapOptionsRequest, runtime *dara.RuntimeOptions) (_result *GetWorkerBootstrapOptionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkerBootstrapOptions"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkerBootstrapOptionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Worker最大可升级版本
//
// @param request - GetWorkerMaxVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkerMaxVersionResponse
func (client *Client) GetWorkerMaxVersionWithContext(ctx context.Context, request *GetWorkerMaxVersionRequest, runtime *dara.RuntimeOptions) (_result *GetWorkerMaxVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkerMaxVersion"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkerMaxVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Worker统计摘要
//
// @param request - GetWorkerStatsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkerStatsSummaryResponse
func (client *Client) GetWorkerStatsSummaryWithContext(ctx context.Context, request *GetWorkerStatsSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetWorkerStatsSummaryResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkerStatsSummary"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkerStatsSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询凭证列表
//
// @param request - ListCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCredentialsResponse
func (client *Client) ListCredentialsWithContext(ctx context.Context, request *ListCredentialsRequest, runtime *dara.RuntimeOptions) (_result *ListCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NameLike) {
		query["NameLike"] = request.NameLike
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCredentials"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询上游身份提供商绑定列表
//
// @param request - ListIdentityProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProvidersWithContext(ctx context.Context, request *ListIdentityProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListIdentityProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
		Action:      dara.String("ListIdentityProviders"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于查询符合条件的实例列表，并支持分页和模糊匹配。
//
// Description:
//
// ## 请求说明
//
// - **分页规则**：
//
//   - 如果传了 `NextToken`，优先按 `NextToken` 解析 offset。
//
//   - 如果没传 `NextToken`，则使用 `skip`。
//
//   - `MaxResults` 的取值范围为 1 到 100，非法值会返回 `400` 错误。
//
//   - `NextToken` 必须是有效的整数，否则会返回 `400` 错误。
//
//   - `skip` 的值不能小于 0，否则会返回 `400` 错误。
//
// - **排序规则**：列表按创建时间倒序返回。
//
// - **请求参数**：
//
//   - `instanceName`：实例名称，支持模糊匹配。
//
//   - `status`：实例状态。
//
//   - `MaxResults`：分页大小，默认值为 20。
//
//   - `NextToken`：下一页游标。
//
//   - `skip`：跳过的记录数，默认值为 0。
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试模型供应商和模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有Magic实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - ListMcpToolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpToolsResponse
func (client *Client) ListMcpToolsWithContext(ctx context.Context, request *ListMcpToolsRequest, runtime *dara.RuntimeOptions) (_result *ListMcpToolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcpTools"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpToolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询MCP列表
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - ListMcpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpsResponse
func (client *Client) ListMcpsWithContext(ctx context.Context, request *ListMcpsRequest, runtime *dara.RuntimeOptions) (_result *ListMcpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
		Action:      dara.String("ListMcps"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型供应商列表
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - ListModelProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProvidersResponse
func (client *Client) ListModelProvidersWithContext(ctx context.Context, request *ListModelProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListModelProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
		Action:      dara.String("ListModelProviders"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型列表
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - ListModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelsResponse
func (client *Client) ListModelsWithContext(ctx context.Context, request *ListModelsRequest, runtime *dara.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProviderName) {
		query["ProviderName"] = request.ProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModels"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 此API用于查询指定实例下的AI网关端点列表。
//
// Description:
//
// ## 请求说明
//
// - `instanceId` 是必填参数，用来指定 AgentTeams 实例 ID。
//
// - 可选参数包括 `component`, `serviceName`, `networkType`, 和 `domainType`，用于进一步筛选返回的端点列表。
//
// - 不支持通过 `status` 参数进行筛选。
//
// @param request - ListServiceEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceEndpointsResponse
func (client *Client) ListServiceEndpointsWithContext(ctx context.Context, request *ListServiceEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Component) {
		query["Component"] = request.Component
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceEndpoints"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户侧APIG可用的SSL证书列表
//
// Description:
//
// ## 请求说明
//
// - 该API用于获取与指定AgentTeams实例相关的SSL证书列表。
//
// - 可通过`certNameLike`和`domainName`参数进行模糊搜索或精确匹配证书名称及绑定域名。
//
// - 分页参数`pageNumber`和`pageSize`允许客户端控制返回结果的数量和页码，默认每页显示10条记录。
//
// - 成功响应将包含请求ID、是否成功标志、错误代码（如果有的话）、HTTP状态码、本次请求的最大结果数、下一页标记（如果有更多数据的话）、总证书数量以及具体的证书详情列表。
//
// @param request - ListSslCertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSslCertsResponse
func (client *Client) ListSslCertsWithContext(ctx context.Context, request *ListSslCertsRequest, runtime *dara.RuntimeOptions) (_result *ListSslCertsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
		Action:      dara.String("ListSslCerts"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSslCertsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 团队详情列表
//
// @param request - ListTeamDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamDetailsResponse
func (client *Client) ListTeamDetailsWithContext(ctx context.Context, request *ListTeamDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListTeamDetailsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTeamDetails"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTeamDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Team任务列表
//
// @param request - ListTeamTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamTasksResponse
func (client *Client) ListTeamTasksWithContext(ctx context.Context, request *ListTeamTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTeamTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Team) {
		query["Team"] = request.Team
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTeamTasks"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTeamTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询团队列表
//
// @param request - ListTeamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamsResponse
func (client *Client) ListTeamsWithContext(ctx context.Context, request *ListTeamsRequest, runtime *dara.RuntimeOptions) (_result *ListTeamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NameLike) {
		query["NameLike"] = request.NameLike
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTeams"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTeamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户列表
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NameLike) {
		query["NameLike"] = request.NameLike
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Worker统计详情列表
//
// @param request - ListWorkerStatsDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkerStatsDetailsResponse
func (client *Client) ListWorkerStatsDetailsWithContext(ctx context.Context, request *ListWorkerStatsDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkerStatsDetailsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkerStatsDetails"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkerStatsDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Worker列表
//
// @param tmpReq - ListWorkersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkersResponse
func (client *Client) ListWorkersWithContext(ctx context.Context, tmpReq *ListWorkersRequest, runtime *dara.RuntimeOptions) (_result *ListWorkersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Group) {
		request.GroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Group, dara.String("Group"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.Credential) {
		query["Credential"] = request.Credential
	}

	if !dara.IsNil(request.GroupShrink) {
		query["Group"] = request.GroupShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Mcp) {
		query["Mcp"] = request.Mcp
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelProvider) {
		query["ModelProvider"] = request.ModelProvider
	}

	if !dara.IsNil(request.NameLike) {
		query["NameLike"] = request.NameLike
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkers"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新CMS工作空间
//
// @param request - PutCmsWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCmsWorkspaceResponse
func (client *Client) PutCmsWorkspaceWithContext(ctx context.Context, request *PutCmsWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *PutCmsWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCmsWorkspace"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCmsWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定实例、worker、团队或个人的功能特性状态。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询特定`instanceId`下的不同目标（如`INSTANCE`、`WORKER`、`TEAM`、`HUMAN`）的功能特性状态。
//
// - `targetScope`参数定义了查询的目标类型，根据不同的`targetScope`值，可能需要提供额外的`resourceName`参数来指定具体的资源名称。
//
// - 如果提供了`featureCodes`列表，则返回这些特定功能特性的状态；否则，将返回指定`targetScope`下所有功能特性的状态。
//
// - 当使用`WORKER`、`TEAM`或`HUMAN`作为`targetScope`时，请确保正确填写对应的`resourceName`。
//
// - 对于`INSTANCE`级别的查询，无需提供`resourceName`。
//
// - 特性支持情况受基础版本、工作器版本等因素影响，并通过`unsupportedReasonCode`和`unsupportedReason`字段给出不支持的具体原因。
//
// @param request - QueryFeaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFeaturesResponse
func (client *Client) QueryFeaturesWithContext(ctx context.Context, request *QueryFeaturesRequest, runtime *dara.RuntimeOptions) (_result *QueryFeaturesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.TargetScope) {
		query["TargetScope"] = request.TargetScope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFeatures"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFeaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前AgentTeams Resource Pool配置支持的所有可用区ID。
//
// Description:
//
// ## 请求说明
//
// @param request - QuerySupportedZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySupportedZonesResponse
func (client *Client) QuerySupportedZonesWithContext(ctx context.Context, request *QuerySupportedZonesRequest, runtime *dara.RuntimeOptions) (_result *QuerySupportedZonesResponse, _err error) {
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
		Action:      dara.String("QuerySupportedZones"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySupportedZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置用户密码
//
// @param request - ResetUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPasswordWithContext(ctx context.Context, request *ResetUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetUserPassword"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试模型供应商和模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - TestModelProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestModelProviderResponse
func (client *Client) TestModelProviderWithContext(ctx context.Context, request *TestModelProviderRequest, runtime *dara.RuntimeOptions) (_result *TestModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	if !dara.IsNil(request.ProviderName) {
		body["ProviderName"] = request.ProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestModelProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestModelProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑上游身份提供商
//
// @param request - UnbindIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindIdentityProviderResponse
func (client *Client) UnbindIdentityProviderWithContext(ctx context.Context, request *UnbindIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *UnbindIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		query["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindIdentityProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新凭证密钥
//
// Description:
//
// ## 请求说明
//
// - 该接口用于更新 AgentTeams 实例下已有 Credential 的密钥明文。
//
// - 仅更新 Agent Identity TokenVault 中同名 APIKeyCredentialProvider 的密钥值，不修改本地元数据（description、createTime、updateTime、status）。
//
// - 响应不包含 apiKey 明文；如需绑定 Worker 明细请调用 GetCredential。
//
// @param request - UpdateCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialResponse
func (client *Client) UpdateCredentialWithContext(ctx context.Context, request *UpdateCredentialRequest, runtime *dara.RuntimeOptions) (_result *UpdateCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCredential"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新上游身份提供商绑定
//
// @param request - UpdateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProviderWithContext(ctx context.Context, request *UpdateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		query["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.IdpMetadata) {
		query["IdpMetadata"] = request.IdpMetadata
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoginEnabled) {
		query["LoginEnabled"] = request.LoginEnabled
	}

	if !dara.IsNil(request.SyncEnabled) {
		query["SyncEnabled"] = request.SyncEnabled
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdentityProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于更改指定AgentTeams实例的名称，支持通过GET或POST方法调用。
//
// Description:
//
// ## 请求说明
//
// - 推荐使用`POST`方法，并以表单形式提交参数。
//
// - 当前实现不支持JSON格式的请求体，请勿尝试使用`@RequestBody`方式调用。
//
// - 必须提供有效的`instanceId`和非空的`instanceName`作为参数。
//
// - 该接口仅允许修改实例名称(`instanceName`)，不允许通过此接口变更命名空间(`namespace`)。
//
// @param tmpReq - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, tmpReq *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Zones) {
		request.ZonesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Zones, dara.String("Zones"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.ZonesShrink) {
		query["Zones"] = request.ZonesShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启暂停中的创建实例异步任务。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于重新启动一个处于暂停状态的创建实例任务。
//
// - 目前仅支持 `agentteams:pay-order:create` 类型的任务。
//
// - 确保提供的 `instanceId`、`taskCode` 和 `taskId` 参数准确无误，否则可能导致请求失败。
//
// - 如果任务不是暂停状态（PAUSED），则不允许调用此接口进行更新。
//
// @param request - UpdateInstanceAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceAsyncTaskResponse
func (client *Client) UpdateInstanceAsyncTaskWithContext(ctx context.Context, request *UpdateInstanceAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsResume) {
		query["IsResume"] = request.IsResume
	}

	if !dara.IsNil(request.TaskCode) {
		query["TaskCode"] = request.TaskCode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceAsyncTask"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新MCP
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param tmpReq - UpdateMcpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpResponse
func (client *Client) UpdateMcpWithContext(ctx context.Context, tmpReq *UpdateMcpRequest, runtime *dara.RuntimeOptions) (_result *UpdateMcpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMcpShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Addresses) {
		request.AddressesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addresses, dara.String("Addresses"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddressesShrink) {
		body["Addresses"] = request.AddressesShrink
	}

	if !dara.IsNil(request.AuthConfig) {
		body["AuthConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.AuthEnabled) {
		body["AuthEnabled"] = request.AuthEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CreateType) {
		body["CreateType"] = request.CreateType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SwaggerConfig) {
		body["SwaggerConfig"] = request.SwaggerConfig
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcp"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param request - UpdateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelResponse
func (client *Client) UpdateModelWithContext(ctx context.Context, request *UpdateModelRequest, runtime *dara.RuntimeOptions) (_result *UpdateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModel"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模型供应商
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询当前登录用户所拥有的所有AgentTeams实例。
//
// - 用户身份通过请求头`X-User-Id`传递，服务端会根据此ID自动过滤只返回属于该用户的实例数据。
//
// - 支持使用`instanceName`进行模糊匹配以及通过`status`参数精确匹配实例状态来过滤查询结果。
//
// - 默认情况下，结果将按照创建时间倒序排列，并且可以通过设置`limit`和`offset`参数来进行分页控制，默认每页显示20条记录。
//
// - 如果请求中缺少`X-User-Id`或者其值为空，则会返回400错误；如果指定的实例不存在或不属于当前用户，则返回404错误。
//
// @param tmpReq - UpdateModelProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelProviderResponse
func (client *Client) UpdateModelProviderWithContext(ctx context.Context, tmpReq *UpdateModelProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateModelProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiKeys) {
		request.ApiKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiKeys, dara.String("ApiKeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Protocols) {
		request.ProtocolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Protocols, dara.String("Protocols"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		body["Address"] = request.Address
	}

	if !dara.IsNil(request.ApiKeysShrink) {
		body["ApiKeys"] = request.ApiKeysShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProtocolsShrink) {
		body["Protocols"] = request.ProtocolsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelProvider"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于更新指定Endpoint的域名和SSL证书信息。
//
// Description:
//
// ## 请求说明
//
// - 本API支持更新`ELEMENT`、`MATRIX`类型的Endpoint。
//
// - 如果尝试更新其他类型的Endpoint，将返回400错误。
//
// - 当`endpointId`不存在或不属于当前用户实例时，将返回404错误。
//
// - 更新域名时，系统会创建或复用新的HTTPS domain，并将其绑定到原endpoint route上，同时解绑旧domain，但不会删除旧domain。
//
// - 若不提供`domain`或`certIdentifier`参数，则保持原有设置不变。
//
// - 其他如`component`、`gatewayType`等字段即使传入也不会被更新。
//
// @param request - UpdateServiceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceEndpointResponse
func (client *Client) UpdateServiceEndpointWithContext(ctx context.Context, request *UpdateServiceEndpointRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceEndpoint"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新团队
//
// @param tmpReq - UpdateTeamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTeamResponse
func (client *Client) UpdateTeamWithContext(ctx context.Context, tmpReq *UpdateTeamRequest, runtime *dara.RuntimeOptions) (_result *UpdateTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTeamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TeamMembers) {
		request.TeamMembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TeamMembers, dara.String("TeamMembers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TeamMembersShrink) {
		query["TeamMembers"] = request.TeamMembersShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTeam"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTeamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户信息
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMethod) {
		query["AuthMethod"] = request.AuthMethod
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Worker
//
// @param tmpReq - UpdateWorkerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkerResponse
func (client *Client) UpdateWorkerWithContext(ctx context.Context, tmpReq *UpdateWorkerRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Channels) {
		request.ChannelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Channels, dara.String("Channels"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Credentials) {
		request.CredentialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Credentials, dara.String("Credentials"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LimitConfig) {
		request.LimitConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LimitConfig, dara.String("LimitConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.McpServers) {
		request.McpServersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.McpServers, dara.String("McpServers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Model) {
		request.ModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Model, dara.String("Model"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Skills) {
		request.SkillsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Skills, dara.String("Skills"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Agents) {
		query["Agents"] = request.Agents
	}

	if !dara.IsNil(request.ChannelsShrink) {
		query["Channels"] = request.ChannelsShrink
	}

	if !dara.IsNil(request.CredentialsShrink) {
		query["Credentials"] = request.CredentialsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LimitConfigShrink) {
		query["LimitConfig"] = request.LimitConfigShrink
	}

	if !dara.IsNil(request.McpServersShrink) {
		query["McpServers"] = request.McpServersShrink
	}

	if !dara.IsNil(request.ModelShrink) {
		query["Model"] = request.ModelShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillsShrink) {
		query["Skills"] = request.SkillsShrink
	}

	if !dara.IsNil(request.Soul) {
		query["Soul"] = request.Soul
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorker"),
		Version:     dara.String("2026-06-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
