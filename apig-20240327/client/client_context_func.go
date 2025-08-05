// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # Authorize the security group for gateway to access services
//
// @param request - AddGatewaySecurityGroupRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewaySecurityGroupRuleResponse
func (client *Client) AddGatewaySecurityGroupRuleWithContext(ctx context.Context, gatewayId *string, request *AddGatewaySecurityGroupRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddGatewaySecurityGroupRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.PortRanges) {
		body["portRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["securityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewaySecurityGroupRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/security-group-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewaySecurityGroupRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes consumer authentication rules.
//
// @param request - BatchDeleteConsumerAuthorizationRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteConsumerAuthorizationRuleResponse
func (client *Client) BatchDeleteConsumerAuthorizationRuleWithContext(ctx context.Context, request *BatchDeleteConsumerAuthorizationRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchDeleteConsumerAuthorizationRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerAuthorizationRuleIds) {
		query["consumerAuthorizationRuleIds"] = request.ConsumerAuthorizationRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Resource Group Transfer
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/move-resource-group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateAndAttachPolicy
//
// @param request - CreateAndAttachPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndAttachPolicyResponse
func (client *Client) CreateAndAttachPolicyWithContext(ctx context.Context, request *CreateAndAttachPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAndAttachPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.ClassName) {
		body["className"] = request.ClassName
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndAttachPolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndAttachPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消费者
//
// @param request - CreateConsumerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerResponse
func (client *Client) CreateConsumerWithContext(ctx context.Context, request *CreateConsumerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AkSkIdentityConfigs) {
		body["akSkIdentityConfigs"] = request.AkSkIdentityConfigs
	}

	if !dara.IsNil(request.ApikeyIdentityConfig) {
		body["apikeyIdentityConfig"] = request.ApikeyIdentityConfig
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.JwtIdentityConfig) {
		body["jwtIdentityConfig"] = request.JwtIdentityConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消费者授权规则
//
// @param request - CreateConsumerAuthorizationRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerAuthorizationRuleResponse
func (client *Client) CreateConsumerAuthorizationRuleWithContext(ctx context.Context, consumerId *string, request *CreateConsumerAuthorizationRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerAuthorizationRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceInfos) {
		body["authorizationResourceInfos"] = request.AuthorizationResourceInfos
	}

	if !dara.IsNil(request.ExpireMode) {
		body["expireMode"] = request.ExpireMode
	}

	if !dara.IsNil(request.ExpireTimestamp) {
		body["expireTimestamp"] = request.ExpireTimestamp
	}

	if !dara.IsNil(request.ParentResourceType) {
		body["parentResourceType"] = request.ParentResourceType
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a consumer authentication rule.
//
// @param request - CreateConsumerAuthorizationRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerAuthorizationRulesResponse
func (client *Client) CreateConsumerAuthorizationRulesWithContext(ctx context.Context, request *CreateConsumerAuthorizationRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerAuthorizationRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRules) {
		body["authorizationRules"] = request.AuthorizationRules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerAuthorizationRules"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerAuthorizationRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a domain name.
//
// Description:
//
// Create Domain.
//
// @param request - CreateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CaCertIdentifier) {
		body["caCertIdentifier"] = request.CaCertIdentifier
	}

	if !dara.IsNil(request.CertIdentifier) {
		body["certIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.ClientCACert) {
		body["clientCACert"] = request.ClientCACert
	}

	if !dara.IsNil(request.ForceHttps) {
		body["forceHttps"] = request.ForceHttps
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Http2Option) {
		body["http2Option"] = request.Http2Option
	}

	if !dara.IsNil(request.MTLSEnabled) {
		body["mTLSEnabled"] = request.MTLSEnabled
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TlsCipherSuitesConfig) {
		body["tlsCipherSuitesConfig"] = request.TlsCipherSuitesConfig
	}

	if !dara.IsNil(request.TlsMax) {
		body["tlsMax"] = request.TlsMax
	}

	if !dara.IsNil(request.TlsMin) {
		body["tlsMin"] = request.TlsMin
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateEnvironment is deprecated
//
// Summary:
//
// # CreateEnvironment
//
// Description:
//
// Create environment.
//
// @param request - CreateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["alias"] = request.Alias
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建云原生网关
//
// @param request - CreateGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayResponse
func (client *Client) CreateGatewayWithContext(ctx context.Context, request *CreateGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		body["chargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.LogConfig) {
		body["logConfig"] = request.LogConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NetworkAccessConfig) {
		body["networkAccessConfig"] = request.NetworkAccessConfig
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Spec) {
		body["spec"] = request.Spec
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		body["vpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneConfig) {
		body["zoneConfig"] = request.ZoneConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an HTTP API.
//
// @param request - CreateHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiResponse
func (client *Client) CreateHttpApiWithContext(ctx context.Context, request *CreateHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHttpApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProtocols) {
		body["agentProtocols"] = request.AgentProtocols
	}

	if !dara.IsNil(request.AiProtocols) {
		body["aiProtocols"] = request.AiProtocols
	}

	if !dara.IsNil(request.AuthConfig) {
		body["authConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.BasePath) {
		body["basePath"] = request.BasePath
	}

	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnableAuth) {
		body["enableAuth"] = request.EnableAuth
	}

	if !dara.IsNil(request.IngressConfig) {
		body["ingressConfig"] = request.IngressConfig
	}

	if !dara.IsNil(request.ModelCategory) {
		body["modelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Protocols) {
		body["protocols"] = request.Protocols
	}

	if !dara.IsNil(request.RemoveBasePathOnForward) {
		body["removeBasePathOnForward"] = request.RemoveBasePathOnForward
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.VersionConfig) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create an Operation for HTTP API
//
// @param request - CreateHttpApiOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiOperationResponse
func (client *Client) CreateHttpApiOperationWithContext(ctx context.Context, httpApiId *string, request *CreateHttpApiOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHttpApiOperationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Operations) {
		body["operations"] = request.Operations
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpApiOperationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route for an HTTP API.
//
// @param request - CreateHttpApiRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiRouteResponse
func (client *Client) CreateHttpApiRouteWithContext(ctx context.Context, httpApiId *string, request *CreateHttpApiRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHttpApiRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BackendConfig) {
		body["backendConfig"] = request.BackendConfig
	}

	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DomainIds) {
		body["domainIds"] = request.DomainIds
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Match) {
		body["match"] = request.Match
	}

	if !dara.IsNil(request.McpRouteConfig) {
		body["mcpRouteConfig"] = request.McpRouteConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpApiRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建API
//
// @param request - CreatePluginAttachmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePluginAttachmentResponse
func (client *Client) CreatePluginAttachmentWithContext(ctx context.Context, request *CreatePluginAttachmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePluginAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PluginConfig) {
		body["pluginConfig"] = request.PluginConfig
	}

	if !dara.IsNil(request.PluginId) {
		body["pluginId"] = request.PluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePluginAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Policy
//
// @param request - CreatePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["className"] = request.ClassName
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create policy resource mount
//
// @param request - CreatePolicyAttachmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyAttachmentResponse
func (client *Client) CreatePolicyAttachmentWithContext(ctx context.Context, request *CreatePolicyAttachmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePolicyAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		body["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PolicyId) {
		body["policyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicyAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-attachments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service.
//
// Description:
//
// You can call this operation to create multiple services at a time.
//
// @param request - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithContext(ctx context.Context, request *CreateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceConfigs) {
		body["serviceConfigs"] = request.ServiceConfigs
	}

	if !dara.IsNil(request.SourceType) {
		body["sourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateService"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消费者
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerResponse
func (client *Client) DeleteConsumerWithContext(ctx context.Context, consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消费者授权规则
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerAuthorizationRuleResponse
func (client *Client) DeleteConsumerAuthorizationRuleWithContext(ctx context.Context, consumerAuthorizationRuleId *string, consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerAuthorizationRuleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteDomain
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithContext(ctx context.Context, domainId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteEnvironment is deprecated
//
// Summary:
//
// # DeleteEnvironment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironmentWithContext(ctx context.Context, environmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments/" + dara.PercentEncode(dara.StringValue(environmentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Gateway
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithContext(ctx context.Context, gatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete the security group rule of a gateway
//
// @param request - DeleteGatewaySecurityGroupRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewaySecurityGroupRuleResponse
func (client *Client) DeleteGatewaySecurityGroupRuleWithContext(ctx context.Context, gatewayId *string, securityGroupRuleId *string, request *DeleteGatewaySecurityGroupRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewaySecurityGroupRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CascadingDelete) {
		query["cascadingDelete"] = request.CascadingDelete
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewaySecurityGroupRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/security-group-rules/" + dara.PercentEncode(dara.StringValue(securityGroupRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewaySecurityGroupRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an HTTP API.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiResponse
func (client *Client) DeleteHttpApiWithContext(ctx context.Context, httpApiId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHttpApiResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Operation
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiOperationResponse
func (client *Client) DeleteHttpApiOperationWithContext(ctx context.Context, httpApiId *string, operationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHttpApiOperationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations/" + dara.PercentEncode(dara.StringValue(operationId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpApiOperationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete the route of an HttpApi
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiRouteResponse
func (client *Client) DeleteHttpApiRouteWithContext(ctx context.Context, httpApiId *string, routeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHttpApiRouteResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes/" + dara.PercentEncode(dara.StringValue(routeId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpApiRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除挂载规则API
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePluginAttachmentResponse
func (client *Client) DeletePluginAttachmentWithContext(ctx context.Context, pluginAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePluginAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments/" + dara.PercentEncode(dara.StringValue(pluginAttachmentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePluginAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Policy
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithContext(ctx context.Context, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete policy resource attachment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyAttachmentResponse
func (client *Client) DeletePolicyAttachmentWithContext(ctx context.Context, policyAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePolicyAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-attachments/" + dara.PercentEncode(dara.StringValue(policyAttachmentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithContext(ctx context.Context, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteService"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Deploy HttpApi
//
// @param request - DeployHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployHttpApiResponse
func (client *Client) DeployHttpApiWithContext(ctx context.Context, httpApiId *string, request *DeployHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployHttpApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HttpApiConfig) {
		body["httpApiConfig"] = request.HttpApiConfig
	}

	if !dara.IsNil(request.RestApiConfig) {
		body["restApiConfig"] = request.RestApiConfig
	}

	if !dara.IsNil(request.RouteId) {
		body["routeId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/deploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Export HTTP API
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportHttpApiResponse
func (client *Client) ExportHttpApiWithContext(ctx context.Context, httpApiId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportHttpApiResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/export"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerResponse
func (client *Client) GetConsumerWithContext(ctx context.Context, consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者授权规则
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerAuthorizationRuleResponse
func (client *Client) GetConsumerAuthorizationRuleWithContext(ctx context.Context, consumerAuthorizationRuleId *string, consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerAuthorizationRuleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains data from dashboards.
//
// @param tmpReq - GetDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDashboardResponse
func (client *Client) GetDashboardWithContext(ctx context.Context, gatewayId *string, tmpReq *GetDashboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDashboardResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetDashboardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ApiId) {
		query["apiId"] = request.ApiId
	}

	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PluginClassId) {
		query["pluginClassId"] = request.PluginClassId
	}

	if !dara.IsNil(request.PluginId) {
		query["pluginId"] = request.PluginId
	}

	if !dara.IsNil(request.RouteId) {
		query["routeId"] = request.RouteId
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.UpstreamCluster) {
		query["upstreamCluster"] = request.UpstreamCluster
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDashboard"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/dashboards"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDashboardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a domain name.
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithContext(ctx context.Context, domainId *string, request *GetDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WithStatistics) {
		query["withStatistics"] = request.WithStatistics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetEnvironment is deprecated
//
// Summary:
//
// # GetEnvironment
//
// @param request - GetEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironmentWithContext(ctx context.Context, environmentId *string, request *GetEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WithStatistics) {
		query["withStatistics"] = request.WithStatistics
	}

	if !dara.IsNil(request.WithVpcInfo) {
		query["withVpcInfo"] = request.WithVpcInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments/" + dara.PercentEncode(dara.StringValue(environmentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about an instance, such as the virtual private cloud (VPC) and vSwitch to which the instance belongs and its ingress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayResponse
func (client *Client) GetGatewayWithContext(ctx context.Context, gatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Read HttpApi
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiResponse
func (client *Client) GetHttpApiWithContext(ctx context.Context, httpApiId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHttpApiResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Operation
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiOperationResponse
func (client *Client) GetHttpApiOperationWithContext(ctx context.Context, httpApiId *string, operationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHttpApiOperationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations/" + dara.PercentEncode(dara.StringValue(operationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpApiOperationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a route of an HTTP API.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiRouteResponse
func (client *Client) GetHttpApiRouteWithContext(ctx context.Context, httpApiId *string, routeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHttpApiRouteResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes/" + dara.PercentEncode(dara.StringValue(routeId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpApiRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetPluginAttachment。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPluginAttachmentResponse
func (client *Client) GetPluginAttachmentWithContext(ctx context.Context, pluginAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPluginAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments/" + dara.PercentEncode(dara.StringValue(pluginAttachmentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPluginAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithContext(ctx context.Context, policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Policy Resource Attachment
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyAttachmentResponse
func (client *Client) GetPolicyAttachmentWithContext(ctx context.Context, policyAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPolicyAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-attachments/" + dara.PercentEncode(dara.StringValue(policyAttachmentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get resource overview information
//
// @param request - GetResourceOverviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceOverviewResponse
func (client *Client) GetResourceOverviewWithContext(ctx context.Context, request *GetResourceOverviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceOverviewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceOverview"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/overview/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceOverviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithContext(ctx context.Context, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetService"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve Tracing Configuration
//
// @param request - GetTraceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceConfigResponse
func (client *Client) GetTraceConfigWithContext(ctx context.Context, gatewayId *string, request *GetTraceConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTraceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTraceConfig"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/trace"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTraceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports HTTP APIs. You can call this operation to import OpenAPI 2.0 and OpenAPI 3.0.x definition files to create REST APIs.
//
// @param request - ImportHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportHttpApiResponse
func (client *Client) ImportHttpApiWithContext(ctx context.Context, request *ImportHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ImportHttpApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.McpRouteId) {
		body["mcpRouteId"] = request.McpRouteId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SpecContentBase64) {
		body["specContentBase64"] = request.SpecContentBase64
	}

	if !dara.IsNil(request.SpecFileUrl) {
		body["specFileUrl"] = request.SpecFileUrl
	}

	if !dara.IsNil(request.SpecOssConfig) {
		body["specOssConfig"] = request.SpecOssConfig
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	if !dara.IsNil(request.TargetHttpApiId) {
		body["targetHttpApiId"] = request.TargetHttpApiId
	}

	if !dara.IsNil(request.VersionConfig) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/import"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者列表
//
// @param request - ListConsumersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumersResponse
func (client *Client) ListConsumersWithContext(ctx context.Context, request *ListConsumersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumers"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain names.
//
// @param request - ListDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithContext(ctx context.Context, request *ListDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomains"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListEnvironments is deprecated
//
// Summary:
//
// # ListEnvironments
//
// @param request - ListEnvironmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironmentsWithContext(ctx context.Context, request *ListEnvironmentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasLike) {
		query["aliasLike"] = request.AliasLike
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayNameLike) {
		query["gatewayNameLike"] = request.GatewayNameLike
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironments"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param tmpReq - ListGatewaysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewaysResponse
func (client *Client) ListGatewaysWithContext(ctx context.Context, tmpReq *ListGatewaysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewaysResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListGatewaysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGateways"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewaysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Operations
//
// @param request - ListHttpApiOperationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApiOperationsResponse
func (client *Client) ListHttpApiOperationsWithContext(ctx context.Context, httpApiId *string, request *ListHttpApiOperationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHttpApiOperationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerAuthorizationRuleId) {
		query["consumerAuthorizationRuleId"] = request.ConsumerAuthorizationRuleId
	}

	if !dara.IsNil(request.ForDeploy) {
		query["forDeploy"] = request.ForDeploy
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Method) {
		query["method"] = request.Method
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PathLike) {
		query["pathLike"] = request.PathLike
	}

	if !dara.IsNil(request.WithConsumerInEnvironmentId) {
		query["withConsumerInEnvironmentId"] = request.WithConsumerInEnvironmentId
	}

	if !dara.IsNil(request.WithConsumerInfoById) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !dara.IsNil(request.WithPluginAttachmentByPluginId) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHttpApiOperations"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpApiOperationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routes of an HTTP API.
//
// @param request - ListHttpApiRoutesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApiRoutesResponse
func (client *Client) ListHttpApiRoutesWithContext(ctx context.Context, httpApiId *string, request *ListHttpApiRoutesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHttpApiRoutesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerAuthorizationRuleId) {
		query["consumerAuthorizationRuleId"] = request.ConsumerAuthorizationRuleId
	}

	if !dara.IsNil(request.DeployStatuses) {
		query["deployStatuses"] = request.DeployStatuses
	}

	if !dara.IsNil(request.DomainId) {
		query["domainId"] = request.DomainId
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.ForDeploy) {
		query["forDeploy"] = request.ForDeploy
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PathLike) {
		query["pathLike"] = request.PathLike
	}

	if !dara.IsNil(request.WithAuthPolicyInfo) {
		query["withAuthPolicyInfo"] = request.WithAuthPolicyInfo
	}

	if !dara.IsNil(request.WithConsumerInfoById) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !dara.IsNil(request.WithPluginAttachmentByPluginId) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHttpApiRoutes"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpApiRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of HTTP APIs.
//
// @param request - ListHttpApisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApisResponse
func (client *Client) ListHttpApisWithContext(ctx context.Context, request *ListHttpApisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHttpApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Types) {
		query["types"] = request.Types
	}

	if !dara.IsNil(request.WithAPIsPublishedToEnvironment) {
		query["withAPIsPublishedToEnvironment"] = request.WithAPIsPublishedToEnvironment
	}

	if !dara.IsNil(request.WithAuthPolicyInEnvironmentId) {
		query["withAuthPolicyInEnvironmentId"] = request.WithAuthPolicyInEnvironmentId
	}

	if !dara.IsNil(request.WithAuthPolicyList) {
		query["withAuthPolicyList"] = request.WithAuthPolicyList
	}

	if !dara.IsNil(request.WithConsumerInfoById) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !dara.IsNil(request.WithEnvironmentInfo) {
		query["withEnvironmentInfo"] = request.WithEnvironmentInfo
	}

	if !dara.IsNil(request.WithEnvironmentInfoById) {
		query["withEnvironmentInfoById"] = request.WithEnvironmentInfoById
	}

	if !dara.IsNil(request.WithIngressInfo) {
		query["withIngressInfo"] = request.WithIngressInfo
	}

	if !dara.IsNil(request.WithPluginAttachmentByPluginId) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
	}

	if !dara.IsNil(request.WithPolicyConfigs) {
		query["withPolicyConfigs"] = request.WithPolicyConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHttpApis"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取挂载列表
//
// @param request - ListPluginAttachmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginAttachmentsResponse
func (client *Client) ListPluginAttachmentsWithContext(ctx context.Context, request *ListPluginAttachmentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginAttachmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.AttachResourceTypes) {
		query["attachResourceTypes"] = request.AttachResourceTypes
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["pluginId"] = request.PluginId
	}

	if !dara.IsNil(request.WithParentResource) {
		query["withParentResource"] = request.WithParentResource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPluginAttachments"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPluginAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListPlugins
//
// @param request - ListPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginsResponse
func (client *Client) ListPluginsWithContext(ctx context.Context, request *ListPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.IncludeBuiltinAiGateway) {
		query["includeBuiltinAiGateway"] = request.IncludeBuiltinAiGateway
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginClassId) {
		query["pluginClassId"] = request.PluginClassId
	}

	if !dara.IsNil(request.PluginClassName) {
		query["pluginClassName"] = request.PluginClassName
	}

	if !dara.IsNil(request.WithAttachmentInfo) {
		query["withAttachmentInfo"] = request.WithAttachmentInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlugins"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugins"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPluginsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListPolicies
//
// @param request - ListPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.WithAttachments) {
		query["withAttachments"] = request.WithAttachments
	}

	if !dara.IsNil(request.WithSystemPolicy) {
		query["withSystemPolicy"] = request.WithSystemPolicy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListPolicyClasses
//
// @param request - ListPolicyClassesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyClassesResponse
func (client *Client) ListPolicyClassesWithContext(ctx context.Context, request *ListPolicyClassesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPolicyClassesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicyClasses"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-classes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyClassesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of services.
//
// @param request - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithContext(ctx context.Context, request *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SourceTypes) {
		query["sourceTypes"] = request.SourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServices"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListSslCerts
//
// @param request - ListSslCertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSslCertsResponse
func (client *Client) ListSslCertsWithContext(ctx context.Context, request *ListSslCertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSslCertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNameLike) {
		query["certNameLike"] = request.CertNameLike
	}

	if !dara.IsNil(request.DomainName) {
		query["domainName"] = request.DomainName
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSslCerts"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ssl/certs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// # Retrieve the availability zones under a cloud-native API gateway region
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListZonesResponse
func (client *Client) ListZonesWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListZonesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListZones"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/zones"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of consumer authentication rules.
//
// @param request - QueryConsumerAuthorizationRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConsumerAuthorizationRulesResponse
func (client *Client) QueryConsumerAuthorizationRulesWithContext(ctx context.Context, request *QueryConsumerAuthorizationRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryConsumerAuthorizationRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiNameLike) {
		query["apiNameLike"] = request.ApiNameLike
	}

	if !dara.IsNil(request.ConsumerId) {
		query["consumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.ConsumerNameLike) {
		query["consumerNameLike"] = request.ConsumerNameLike
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GroupByApi) {
		query["groupByApi"] = request.GroupByApi
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentResourceId) {
		query["parentResourceId"] = request.ParentResourceId
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["resourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConsumerAuthorizationRules"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConsumerAuthorizationRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a consumer authorization rule.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveConsumerAuthorizationRuleResponse
func (client *Client) RemoveConsumerAuthorizationRuleWithContext(ctx context.Context, consumerAuthorizationRuleId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveConsumerAuthorizationRuleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Gateway Restart
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartGatewayResponse
func (client *Client) RestartGatewayWithContext(ctx context.Context, gatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublishes an HTTP API.
//
// @param request - UndeployHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UndeployHttpApiResponse
func (client *Client) UndeployHttpApiWithContext(ctx context.Context, httpApiId *string, request *UndeployHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UndeployHttpApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.OperationId) {
		body["operationId"] = request.OperationId
	}

	if !dara.IsNil(request.RouteId) {
		body["routeId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UndeployHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/undeploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UndeployHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateAndAttachPolicy
//
// @param request - UpdateAndAttachPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAndAttachPolicyResponse
func (client *Client) UpdateAndAttachPolicyWithContext(ctx context.Context, policyId *string, request *UpdateAndAttachPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAndAttachPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAndAttachPolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAndAttachPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新消费者
//
// @param request - UpdateConsumerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerResponse
func (client *Client) UpdateConsumerWithContext(ctx context.Context, consumerId *string, request *UpdateConsumerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConsumerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AkSkIdentityConfigs) {
		body["akSkIdentityConfigs"] = request.AkSkIdentityConfigs
	}

	if !dara.IsNil(request.ApikeyIdentityConfig) {
		body["apikeyIdentityConfig"] = request.ApikeyIdentityConfig
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.JwtIdentityConfig) {
		body["jwtIdentityConfig"] = request.JwtIdentityConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConsumerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新消费者授权规则
//
// @param request - UpdateConsumerAuthorizationRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerAuthorizationRuleResponse
func (client *Client) UpdateConsumerAuthorizationRuleWithContext(ctx context.Context, consumerId *string, consumerAuthorizationRuleId *string, request *UpdateConsumerAuthorizationRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConsumerAuthorizationRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceInfos) {
		body["authorizationResourceInfos"] = request.AuthorizationResourceInfos
	}

	if !dara.IsNil(request.ExpireMode) {
		body["expireMode"] = request.ExpireMode
	}

	if !dara.IsNil(request.ExpireTimestamp) {
		body["expireTimestamp"] = request.ExpireTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a domain name.
//
// @param request - UpdateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithContext(ctx context.Context, domainId *string, request *UpdateDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CaCertIdentifier) {
		body["caCertIdentifier"] = request.CaCertIdentifier
	}

	if !dara.IsNil(request.CertIdentifier) {
		body["certIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.ClientCACert) {
		body["clientCACert"] = request.ClientCACert
	}

	if !dara.IsNil(request.ForceHttps) {
		body["forceHttps"] = request.ForceHttps
	}

	if !dara.IsNil(request.Http2Option) {
		body["http2Option"] = request.Http2Option
	}

	if !dara.IsNil(request.MTLSEnabled) {
		body["mTLSEnabled"] = request.MTLSEnabled
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TlsCipherSuitesConfig) {
		body["tlsCipherSuitesConfig"] = request.TlsCipherSuitesConfig
	}

	if !dara.IsNil(request.TlsMax) {
		body["tlsMax"] = request.TlsMax
	}

	if !dara.IsNil(request.TlsMin) {
		body["tlsMin"] = request.TlsMin
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateEnvironment is deprecated
//
// Summary:
//
// # UpdateEnvironment
//
// @param request - UpdateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironmentWithContext(ctx context.Context, environmentId *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["alias"] = request.Alias
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments/" + dara.PercentEncode(dara.StringValue(environmentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the feature configuration of the gateway
//
// @param request - UpdateGatewayFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayFeatureResponse
func (client *Client) UpdateGatewayFeatureWithContext(ctx context.Context, gatewayId *string, name *string, request *UpdateGatewayFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Value) {
		body["value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayFeature"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/gateway-features/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Change the name of a gateway instance
//
// @param request - UpdateGatewayNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayNameResponse
func (client *Client) UpdateGatewayNameWithContext(ctx context.Context, gatewayId *string, request *UpdateGatewayNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayName"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/name"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an HTTP API.
//
// @param request - UpdateHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiResponse
func (client *Client) UpdateHttpApiWithContext(ctx context.Context, httpApiId *string, request *UpdateHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHttpApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProtocols) {
		body["agentProtocols"] = request.AgentProtocols
	}

	if !dara.IsNil(request.AiProtocols) {
		body["aiProtocols"] = request.AiProtocols
	}

	if !dara.IsNil(request.AuthConfig) {
		body["authConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.BasePath) {
		body["basePath"] = request.BasePath
	}

	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnableAuth) {
		body["enableAuth"] = request.EnableAuth
	}

	if !dara.IsNil(request.IngressConfig) {
		body["ingressConfig"] = request.IngressConfig
	}

	if !dara.IsNil(request.OnlyChangeConfig) {
		body["onlyChangeConfig"] = request.OnlyChangeConfig
	}

	if !dara.IsNil(request.Protocols) {
		body["protocols"] = request.Protocols
	}

	if !dara.IsNil(request.RemoveBasePathOnForward) {
		body["removeBasePathOnForward"] = request.RemoveBasePathOnForward
	}

	if !dara.IsNil(request.VersionConfig) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Operation
//
// @param request - UpdateHttpApiOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiOperationResponse
func (client *Client) UpdateHttpApiOperationWithContext(ctx context.Context, httpApiId *string, operationId *string, request *UpdateHttpApiOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHttpApiOperationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Operation) {
		body["operation"] = request.Operation
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations/" + dara.PercentEncode(dara.StringValue(operationId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpApiOperationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the route of an HTTP API.
//
// @param request - UpdateHttpApiRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiRouteResponse
func (client *Client) UpdateHttpApiRouteWithContext(ctx context.Context, httpApiId *string, routeId *string, request *UpdateHttpApiRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHttpApiRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BackendConfig) {
		body["backendConfig"] = request.BackendConfig
	}

	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DomainIds) {
		body["domainIds"] = request.DomainIds
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Match) {
		body["match"] = request.Match
	}

	if !dara.IsNil(request.McpRouteConfig) {
		body["mcpRouteConfig"] = request.McpRouteConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes/" + dara.PercentEncode(dara.StringValue(routeId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpApiRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新挂载规则API
//
// @param request - UpdatePluginAttachmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePluginAttachmentResponse
func (client *Client) UpdatePluginAttachmentWithContext(ctx context.Context, pluginAttachmentId *string, request *UpdatePluginAttachmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePluginAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.PluginConfig) {
		body["pluginConfig"] = request.PluginConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments/" + dara.PercentEncode(dara.StringValue(pluginAttachmentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePluginAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Policy
//
// @param request - UpdatePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicyWithContext(ctx context.Context, policyId *string, request *UpdatePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Upgrade the gateway version
//
// @param request - UpgradeGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeGatewayResponse
func (client *Client) UpgradeGatewayWithContext(ctx context.Context, gatewayId *string, request *UpgradeGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeGatewayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("UpgradeGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/upgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
