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
	client.Endpoint, _err = client.GetEndpoint(dara.String("epg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 向出口代理地址池添加IP地址
//
// @param request - AddIpAddressesToEgressProxyIpPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpAddressesToEgressProxyIpPoolResponse
func (client *Client) AddIpAddressesToEgressProxyIpPoolWithOptions(request *AddIpAddressesToEgressProxyIpPoolRequest, runtime *dara.RuntimeOptions) (_result *AddIpAddressesToEgressProxyIpPoolResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyIpPoolId) {
		body["EgressProxyIpPoolId"] = request.EgressProxyIpPoolId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ZoneMappings) {
		bodyFlat["ZoneMappings"] = request.ZoneMappings
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddIpAddressesToEgressProxyIpPool"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIpAddressesToEgressProxyIpPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向出口代理地址池添加IP地址
//
// @param request - AddIpAddressesToEgressProxyIpPoolRequest
//
// @return AddIpAddressesToEgressProxyIpPoolResponse
func (client *Client) AddIpAddressesToEgressProxyIpPool(request *AddIpAddressesToEgressProxyIpPoolRequest) (_result *AddIpAddressesToEgressProxyIpPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddIpAddressesToEgressProxyIpPoolResponse{}
	_body, _err := client.AddIpAddressesToEgressProxyIpPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches egress proxy rule groups to an egress proxy configuration.
//
// Description:
//
// When you call this operation to attach egress proxy rule groups to an egress proxy configuration, take note of the following items:
//
// - This operation attaches one or more egress proxy rule groups to a specified egress proxy configuration and specifies the priority of each egress proxy rule group in the egress proxy configuration.
//
// - The **AttachEgressProxyRuleGroupsToConfiguration*	- operation is asynchronous. After you send the request, the system returns a response, but the egress proxy rule groups are not yet attached. The attachment task is still running in the background. You can call GetEgressProxyConfigurationAttribute to query the status of the egress proxy configuration:
//
//   - If the egress proxy configuration is in the **Configuring*	- state, the configuration is being updated.
//
//   - If the egress proxy configuration is in the **Available*	- state, the attachment is successful.
//
// @param request - AttachEgressProxyRuleGroupsToConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachEgressProxyRuleGroupsToConfigurationResponse
func (client *Client) AttachEgressProxyRuleGroupsToConfigurationWithOptions(request *AttachEgressProxyRuleGroupsToConfigurationRequest, runtime *dara.RuntimeOptions) (_result *AttachEgressProxyRuleGroupsToConfigurationResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyConfigurationId) {
		body["EgressProxyConfigurationId"] = request.EgressProxyConfigurationId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRuleGroups) {
		bodyFlat["EgressProxyRuleGroups"] = request.EgressProxyRuleGroups
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachEgressProxyRuleGroupsToConfiguration"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachEgressProxyRuleGroupsToConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches egress proxy rule groups to an egress proxy configuration.
//
// Description:
//
// When you call this operation to attach egress proxy rule groups to an egress proxy configuration, take note of the following items:
//
// - This operation attaches one or more egress proxy rule groups to a specified egress proxy configuration and specifies the priority of each egress proxy rule group in the egress proxy configuration.
//
// - The **AttachEgressProxyRuleGroupsToConfiguration*	- operation is asynchronous. After you send the request, the system returns a response, but the egress proxy rule groups are not yet attached. The attachment task is still running in the background. You can call GetEgressProxyConfigurationAttribute to query the status of the egress proxy configuration:
//
//   - If the egress proxy configuration is in the **Configuring*	- state, the configuration is being updated.
//
//   - If the egress proxy configuration is in the **Available*	- state, the attachment is successful.
//
// @param request - AttachEgressProxyRuleGroupsToConfigurationRequest
//
// @return AttachEgressProxyRuleGroupsToConfigurationResponse
func (client *Client) AttachEgressProxyRuleGroupsToConfiguration(request *AttachEgressProxyRuleGroupsToConfigurationRequest) (_result *AttachEgressProxyRuleGroupsToConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachEgressProxyRuleGroupsToConfigurationResponse{}
	_body, _err := client.AttachEgressProxyRuleGroupsToConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an egress proxy configuration.
//
// Description:
//
// When you call this operation to create an egress proxy configuration, take note of the following items:
//
// - The **CreateEgressProxyConfiguration*	- operation is asynchronous. After you call this operation, the system returns the ID of the egress proxy configuration, but the configuration is not yet created. The creation task is still running in the background. You can call GetEgressProxyConfigurationAttribute to query the status of the egress proxy configuration:
//
//   - If the egress proxy configuration is in the **Provisioning*	- state, the configuration is being created.
//
//   - If the egress proxy configuration is in the **Available*	- state, the configuration is created.
//
// @param request - CreateEgressProxyConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEgressProxyConfigurationResponse
func (client *Client) CreateEgressProxyConfigurationWithOptions(request *CreateEgressProxyConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateEgressProxyConfigurationResponse, _err error) {
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

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DefaultRulePhaseActions) {
		bodyFlat["DefaultRulePhaseActions"] = request.DefaultRulePhaseActions
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyConfigurationName) {
		body["EgressProxyConfigurationName"] = request.EgressProxyConfigurationName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEgressProxyConfiguration"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEgressProxyConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an egress proxy configuration.
//
// Description:
//
// When you call this operation to create an egress proxy configuration, take note of the following items:
//
// - The **CreateEgressProxyConfiguration*	- operation is asynchronous. After you call this operation, the system returns the ID of the egress proxy configuration, but the configuration is not yet created. The creation task is still running in the background. You can call GetEgressProxyConfigurationAttribute to query the status of the egress proxy configuration:
//
//   - If the egress proxy configuration is in the **Provisioning*	- state, the configuration is being created.
//
//   - If the egress proxy configuration is in the **Available*	- state, the configuration is created.
//
// @param request - CreateEgressProxyConfigurationRequest
//
// @return CreateEgressProxyConfigurationResponse
func (client *Client) CreateEgressProxyConfiguration(request *CreateEgressProxyConfigurationRequest) (_result *CreateEgressProxyConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEgressProxyConfigurationResponse{}
	_body, _err := client.CreateEgressProxyConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an egress proxy gateway instance.
//
// Description:
//
// The **CreateEgressProxyGateway*	- operation is asynchronous. After you call this operation, the system returns an instance ID, but the egress proxy gateway instance is not yet created and the creation task is still running in the background. You can call GetEgressProxyGatewayAttribute to query the creation status of the egress proxy gateway instance:
//
// - When the egress proxy gateway instance is in the **Provisioning*	- state, the instance is being created.
//
// - When the egress proxy gateway instance is in the **Active*	- state, the instance is created.
//
// @param request - CreateEgressProxyGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEgressProxyGatewayResponse
func (client *Client) CreateEgressProxyGatewayWithOptions(request *CreateEgressProxyGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateEgressProxyGatewayResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyConfigurationId) {
		body["EgressProxyConfigurationId"] = request.EgressProxyConfigurationId
	}

	if !dara.IsNil(request.EgressProxyGatewayName) {
		body["EgressProxyGatewayName"] = request.EgressProxyGatewayName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ListenerProperties) {
		bodyFlat["ListenerProperties"] = request.ListenerProperties
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TlsInterceptProperties) {
		bodyFlat["TlsInterceptProperties"] = request.TlsInterceptProperties
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneMappings) {
		bodyFlat["ZoneMappings"] = request.ZoneMappings
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEgressProxyGateway"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEgressProxyGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an egress proxy gateway instance.
//
// Description:
//
// The **CreateEgressProxyGateway*	- operation is asynchronous. After you call this operation, the system returns an instance ID, but the egress proxy gateway instance is not yet created and the creation task is still running in the background. You can call GetEgressProxyGatewayAttribute to query the creation status of the egress proxy gateway instance:
//
// - When the egress proxy gateway instance is in the **Provisioning*	- state, the instance is being created.
//
// - When the egress proxy gateway instance is in the **Active*	- state, the instance is created.
//
// @param request - CreateEgressProxyGatewayRequest
//
// @return CreateEgressProxyGatewayResponse
func (client *Client) CreateEgressProxyGateway(request *CreateEgressProxyGatewayRequest) (_result *CreateEgressProxyGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEgressProxyGatewayResponse{}
	_body, _err := client.CreateEgressProxyGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建出口代理地址池
//
// @param request - CreateEgressProxyIpPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEgressProxyIpPoolResponse
func (client *Client) CreateEgressProxyIpPoolWithOptions(request *CreateEgressProxyIpPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateEgressProxyIpPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		body["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyGatewayId) {
		body["EgressProxyGatewayId"] = request.EgressProxyGatewayId
	}

	if !dara.IsNil(request.EgressProxyIpPoolName) {
		body["EgressProxyIpPoolName"] = request.EgressProxyIpPoolName
	}

	if !dara.IsNil(request.HashFactor) {
		body["HashFactor"] = request.HashFactor
	}

	if !dara.IsNil(request.HashHeaderName) {
		body["HashHeaderName"] = request.HashHeaderName
	}

	if !dara.IsNil(request.Scheduler) {
		body["Scheduler"] = request.Scheduler
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	if !dara.IsNil(request.ZoneMappings) {
		bodyFlat["ZoneMappings"] = request.ZoneMappings
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEgressProxyIpPool"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEgressProxyIpPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建出口代理地址池
//
// @param request - CreateEgressProxyIpPoolRequest
//
// @return CreateEgressProxyIpPoolResponse
func (client *Client) CreateEgressProxyIpPool(request *CreateEgressProxyIpPoolRequest) (_result *CreateEgressProxyIpPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEgressProxyIpPoolResponse{}
	_body, _err := client.CreateEgressProxyIpPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an egress rule group.
//
// Description:
//
// When you call this operation to create an egress rule group, take note of the following items:
//
// - You can create egress rules in batches by phase (PreDns, PreRequest, and PostResponse) when you create the rule group.
//
// - Within the same phase of a rule group, rule priorities (Priority) must be unique. Valid values: 1 to 1000. A smaller value indicates a higher priority.
//
// - This operation is asynchronous. After you call this operation, the system returns the ID of the egress rule group, but the rule group is not yet created. The creation task is still running in the background. You can call GetEgressProxyRuleGroupAttribute to query the status of the egress rule group:
//
//   - If the egress rule group is in the **Provisioning*	- state, the rule group is being created.
//
//   - If the egress rule group is in the **Available*	- state, the rule group is created.
//
// @param request - CreateEgressProxyRuleGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEgressProxyRuleGroupResponse
func (client *Client) CreateEgressProxyRuleGroupWithOptions(request *CreateEgressProxyRuleGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateEgressProxyRuleGroupResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyIpPoolId) {
		body["EgressProxyIpPoolId"] = request.EgressProxyIpPoolId
	}

	if !dara.IsNil(request.EgressProxyRuleGroupName) {
		body["EgressProxyRuleGroupName"] = request.EgressProxyRuleGroupName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRules) {
		bodyFlat["EgressProxyRules"] = request.EgressProxyRules
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEgressProxyRuleGroup"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEgressProxyRuleGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an egress rule group.
//
// Description:
//
// When you call this operation to create an egress rule group, take note of the following items:
//
// - You can create egress rules in batches by phase (PreDns, PreRequest, and PostResponse) when you create the rule group.
//
// - Within the same phase of a rule group, rule priorities (Priority) must be unique. Valid values: 1 to 1000. A smaller value indicates a higher priority.
//
// - This operation is asynchronous. After you call this operation, the system returns the ID of the egress rule group, but the rule group is not yet created. The creation task is still running in the background. You can call GetEgressProxyRuleGroupAttribute to query the status of the egress rule group:
//
//   - If the egress rule group is in the **Provisioning*	- state, the rule group is being created.
//
//   - If the egress rule group is in the **Available*	- state, the rule group is created.
//
// @param request - CreateEgressProxyRuleGroupRequest
//
// @return CreateEgressProxyRuleGroupResponse
func (client *Client) CreateEgressProxyRuleGroup(request *CreateEgressProxyRuleGroupRequest) (_result *CreateEgressProxyRuleGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEgressProxyRuleGroupResponse{}
	_body, _err := client.CreateEgressProxyRuleGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates multiple egress proxy rules.
//
// Description:
//
// When you call this operation to create multiple egress proxy rules in a specified egress rule group, note the following items:
//
// - Egress rules are created in groups based on the effective phase (PreDns, PreRequest, or PostResponse).
//
// - Within the same effective phase of an egress rule group, the priority of rules must be unique.
//
// - Valid values for the priority: 1 to 1000. A smaller value indicates a higher priority.
//
// - This operation is asynchronous. After you call this operation, the system returns the IDs of the created egress rules, but the rules are not yet created. The creation task is still running in the background. You can call the GetEgressProxyRuleGroupAttribute operation to query the status of the egress rule group:
//
//   - If the egress rule group is in the **Provisioning*	- state, the rules are being created.
//
//   - If the egress rule group is in the **Available*	- state, the rules are created.
//
// @param request - CreateEgressProxyRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEgressProxyRulesResponse
func (client *Client) CreateEgressProxyRulesWithOptions(request *CreateEgressProxyRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateEgressProxyRulesResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyRuleGroupId) {
		body["EgressProxyRuleGroupId"] = request.EgressProxyRuleGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRules) {
		bodyFlat["EgressProxyRules"] = request.EgressProxyRules
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEgressProxyRules"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEgressProxyRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple egress proxy rules.
//
// Description:
//
// When you call this operation to create multiple egress proxy rules in a specified egress rule group, note the following items:
//
// - Egress rules are created in groups based on the effective phase (PreDns, PreRequest, or PostResponse).
//
// - Within the same effective phase of an egress rule group, the priority of rules must be unique.
//
// - Valid values for the priority: 1 to 1000. A smaller value indicates a higher priority.
//
// - This operation is asynchronous. After you call this operation, the system returns the IDs of the created egress rules, but the rules are not yet created. The creation task is still running in the background. You can call the GetEgressProxyRuleGroupAttribute operation to query the status of the egress rule group:
//
//   - If the egress rule group is in the **Provisioning*	- state, the rules are being created.
//
//   - If the egress rule group is in the **Available*	- state, the rules are created.
//
// @param request - CreateEgressProxyRulesRequest
//
// @return CreateEgressProxyRulesResponse
func (client *Client) CreateEgressProxyRules(request *CreateEgressProxyRulesRequest) (_result *CreateEgressProxyRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEgressProxyRulesResponse{}
	_body, _err := client.CreateEgressProxyRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy configuration.
//
// @param request - DeleteEgressProxyConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEgressProxyConfigurationResponse
func (client *Client) DeleteEgressProxyConfigurationWithOptions(request *DeleteEgressProxyConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteEgressProxyConfigurationResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyConfigurationId) {
		body["EgressProxyConfigurationId"] = request.EgressProxyConfigurationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEgressProxyConfiguration"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEgressProxyConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy configuration.
//
// @param request - DeleteEgressProxyConfigurationRequest
//
// @return DeleteEgressProxyConfigurationResponse
func (client *Client) DeleteEgressProxyConfiguration(request *DeleteEgressProxyConfigurationRequest) (_result *DeleteEgressProxyConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEgressProxyConfigurationResponse{}
	_body, _err := client.DeleteEgressProxyConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy gateway instance.
//
// Description:
//
// ## Request description
//
// - `EgressProxyGatewayId` is a required parameter that specifies the egress proxy gateway instance ID to delete.
//
// - `ClientToken` is an optional parameter used to ensure the idempotence of the request.
//
// @param request - DeleteEgressProxyGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEgressProxyGatewayResponse
func (client *Client) DeleteEgressProxyGatewayWithOptions(request *DeleteEgressProxyGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteEgressProxyGatewayResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyGatewayId) {
		body["EgressProxyGatewayId"] = request.EgressProxyGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEgressProxyGateway"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEgressProxyGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy gateway instance.
//
// Description:
//
// ## Request description
//
// - `EgressProxyGatewayId` is a required parameter that specifies the egress proxy gateway instance ID to delete.
//
// - `ClientToken` is an optional parameter used to ensure the idempotence of the request.
//
// @param request - DeleteEgressProxyGatewayRequest
//
// @return DeleteEgressProxyGatewayResponse
func (client *Client) DeleteEgressProxyGateway(request *DeleteEgressProxyGatewayRequest) (_result *DeleteEgressProxyGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEgressProxyGatewayResponse{}
	_body, _err := client.DeleteEgressProxyGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy IP address pool.
//
// Description:
//
// ## Request description
//
// - **EgressProxyIpPoolId*	- is a required parameter that specifies the ID of the egress proxy IP address pool to delete.
//
// - **ClientToken*	- is an optional parameter used to ensure the idempotence of the request.
//
// - **DryRun*	- is an optional parameter. Default value: `false`. If you set this parameter to `true`, the system only prechecks the request and does not perform the delete operation.
//
// @param request - DeleteEgressProxyIpPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEgressProxyIpPoolResponse
func (client *Client) DeleteEgressProxyIpPoolWithOptions(request *DeleteEgressProxyIpPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteEgressProxyIpPoolResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyIpPoolId) {
		body["EgressProxyIpPoolId"] = request.EgressProxyIpPoolId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEgressProxyIpPool"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEgressProxyIpPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy IP address pool.
//
// Description:
//
// ## Request description
//
// - **EgressProxyIpPoolId*	- is a required parameter that specifies the ID of the egress proxy IP address pool to delete.
//
// - **ClientToken*	- is an optional parameter used to ensure the idempotence of the request.
//
// - **DryRun*	- is an optional parameter. Default value: `false`. If you set this parameter to `true`, the system only prechecks the request and does not perform the delete operation.
//
// @param request - DeleteEgressProxyIpPoolRequest
//
// @return DeleteEgressProxyIpPoolResponse
func (client *Client) DeleteEgressProxyIpPool(request *DeleteEgressProxyIpPoolRequest) (_result *DeleteEgressProxyIpPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEgressProxyIpPoolResponse{}
	_body, _err := client.DeleteEgressProxyIpPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy rule group.
//
// @param request - DeleteEgressProxyRuleGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEgressProxyRuleGroupResponse
func (client *Client) DeleteEgressProxyRuleGroupWithOptions(request *DeleteEgressProxyRuleGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteEgressProxyRuleGroupResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyRuleGroupId) {
		body["EgressProxyRuleGroupId"] = request.EgressProxyRuleGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEgressProxyRuleGroup"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEgressProxyRuleGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an egress proxy rule group.
//
// @param request - DeleteEgressProxyRuleGroupRequest
//
// @return DeleteEgressProxyRuleGroupResponse
func (client *Client) DeleteEgressProxyRuleGroup(request *DeleteEgressProxyRuleGroupRequest) (_result *DeleteEgressProxyRuleGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEgressProxyRuleGroupResponse{}
	_body, _err := client.DeleteEgressProxyRuleGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple egress proxy rules.
//
// @param request - DeleteEgressProxyRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEgressProxyRulesResponse
func (client *Client) DeleteEgressProxyRulesWithOptions(request *DeleteEgressProxyRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteEgressProxyRulesResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyRuleGroupId) {
		body["EgressProxyRuleGroupId"] = request.EgressProxyRuleGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRuleIds) {
		bodyFlat["EgressProxyRuleIds"] = request.EgressProxyRuleIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEgressProxyRules"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEgressProxyRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple egress proxy rules.
//
// @param request - DeleteEgressProxyRulesRequest
//
// @return DeleteEgressProxyRulesResponse
func (client *Client) DeleteEgressProxyRules(request *DeleteEgressProxyRulesRequest) (_result *DeleteEgressProxyRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEgressProxyRulesResponse{}
	_body, _err := client.DeleteEgressProxyRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries zones.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries zones.
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates egress proxy rule groups from an egress proxy configuration.
//
// Description:
//
// When you call this operation to disassociate egress proxy rule groups from an egress proxy configuration, take note of the following items:
//
// - This operation disassociates one or more egress proxy rule groups from a specified egress proxy configuration. After the disassociation, these egress proxy rule groups no longer take effect in the egress proxy configuration.
//
// - The **DetachEgressProxyRuleGroupsFromConfiguration*	- operation is asynchronous. After you send a request, the system returns a response, but the egress proxy rule groups are not yet disassociated. The disassociation task is still running in the background. You can call the GetEgressProxyConfigurationAttribute operation to query the status of the egress proxy configuration:
//
//   - If the egress proxy configuration is in the **Configuring*	- state, the configuration is being updated.
//
//   - If the egress proxy configuration is in the **Available*	- state, the disassociation is successful.
//
// @param request - DetachEgressProxyRuleGroupsFromConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachEgressProxyRuleGroupsFromConfigurationResponse
func (client *Client) DetachEgressProxyRuleGroupsFromConfigurationWithOptions(request *DetachEgressProxyRuleGroupsFromConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DetachEgressProxyRuleGroupsFromConfigurationResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyConfigurationId) {
		body["EgressProxyConfigurationId"] = request.EgressProxyConfigurationId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRuleGroupIds) {
		bodyFlat["EgressProxyRuleGroupIds"] = request.EgressProxyRuleGroupIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachEgressProxyRuleGroupsFromConfiguration"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachEgressProxyRuleGroupsFromConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates egress proxy rule groups from an egress proxy configuration.
//
// Description:
//
// When you call this operation to disassociate egress proxy rule groups from an egress proxy configuration, take note of the following items:
//
// - This operation disassociates one or more egress proxy rule groups from a specified egress proxy configuration. After the disassociation, these egress proxy rule groups no longer take effect in the egress proxy configuration.
//
// - The **DetachEgressProxyRuleGroupsFromConfiguration*	- operation is asynchronous. After you send a request, the system returns a response, but the egress proxy rule groups are not yet disassociated. The disassociation task is still running in the background. You can call the GetEgressProxyConfigurationAttribute operation to query the status of the egress proxy configuration:
//
//   - If the egress proxy configuration is in the **Configuring*	- state, the configuration is being updated.
//
//   - If the egress proxy configuration is in the **Available*	- state, the disassociation is successful.
//
// @param request - DetachEgressProxyRuleGroupsFromConfigurationRequest
//
// @return DetachEgressProxyRuleGroupsFromConfigurationResponse
func (client *Client) DetachEgressProxyRuleGroupsFromConfiguration(request *DetachEgressProxyRuleGroupsFromConfigurationRequest) (_result *DetachEgressProxyRuleGroupsFromConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachEgressProxyRuleGroupsFromConfigurationResponse{}
	_body, _err := client.DetachEgressProxyRuleGroupsFromConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an egress proxy configuration.
//
// Description:
//
// *GetEgressProxyConfigurationAttribute*	- is a synchronous API operation that returns the details of a specified egress proxy configuration in real time after the call. You can determine the current status of the egress proxy configuration based on the returned **Status**. For example, **Provisioning*	- indicates that the configuration is being created, and **Available*	- indicates that the configuration is created and available.
//
// @param request - GetEgressProxyConfigurationAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEgressProxyConfigurationAttributeResponse
func (client *Client) GetEgressProxyConfigurationAttributeWithOptions(request *GetEgressProxyConfigurationAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetEgressProxyConfigurationAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyConfigurationId) {
		body["EgressProxyConfigurationId"] = request.EgressProxyConfigurationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEgressProxyConfigurationAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEgressProxyConfigurationAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an egress proxy configuration.
//
// Description:
//
// *GetEgressProxyConfigurationAttribute*	- is a synchronous API operation that returns the details of a specified egress proxy configuration in real time after the call. You can determine the current status of the egress proxy configuration based on the returned **Status**. For example, **Provisioning*	- indicates that the configuration is being created, and **Available*	- indicates that the configuration is created and available.
//
// @param request - GetEgressProxyConfigurationAttributeRequest
//
// @return GetEgressProxyConfigurationAttributeResponse
func (client *Client) GetEgressProxyConfigurationAttribute(request *GetEgressProxyConfigurationAttributeRequest) (_result *GetEgressProxyConfigurationAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEgressProxyConfigurationAttributeResponse{}
	_body, _err := client.GetEgressProxyConfigurationAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an egress proxy gateway instance.
//
// Description:
//
// *GetEgressProxyGatewayAttribute*	- is a synchronous operation that returns the current attribute information of the egress proxy gateway in real time after it is called.
//
// You can determine the current status of the egress proxy gateway based on the returned **Status**. For example, **Provisioning*	- indicates that the instance is being created, **Active*	- indicates that the instance is running, **Configuring*	- indicates that the instance is being reconfigured, and **Inactive*	- indicates that the instance is stopped.
//
// @param request - GetEgressProxyGatewayAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEgressProxyGatewayAttributeResponse
func (client *Client) GetEgressProxyGatewayAttributeWithOptions(request *GetEgressProxyGatewayAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetEgressProxyGatewayAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyGatewayId) {
		body["EgressProxyGatewayId"] = request.EgressProxyGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEgressProxyGatewayAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEgressProxyGatewayAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an egress proxy gateway instance.
//
// Description:
//
// *GetEgressProxyGatewayAttribute*	- is a synchronous operation that returns the current attribute information of the egress proxy gateway in real time after it is called.
//
// You can determine the current status of the egress proxy gateway based on the returned **Status**. For example, **Provisioning*	- indicates that the instance is being created, **Active*	- indicates that the instance is running, **Configuring*	- indicates that the instance is being reconfigured, and **Inactive*	- indicates that the instance is stopped.
//
// @param request - GetEgressProxyGatewayAttributeRequest
//
// @return GetEgressProxyGatewayAttributeResponse
func (client *Client) GetEgressProxyGatewayAttribute(request *GetEgressProxyGatewayAttributeRequest) (_result *GetEgressProxyGatewayAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEgressProxyGatewayAttributeResponse{}
	_body, _err := client.GetEgressProxyGatewayAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询出口代理地址池详细信息
//
// Description:
//
// ## 请求说明
//
// 通过提供`EgressProxyIpPoolId`，可以获取到与该ID关联的出口代理IP池的所有相关信息，包括但不限于名称、描述、网络类型等。请确保提供的`EgressProxyIpPoolId`是有效的且属于您的账户。
//
// @param request - GetEgressProxyIpPoolAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEgressProxyIpPoolAttributeResponse
func (client *Client) GetEgressProxyIpPoolAttributeWithOptions(request *GetEgressProxyIpPoolAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetEgressProxyIpPoolAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyIpPoolId) {
		body["EgressProxyIpPoolId"] = request.EgressProxyIpPoolId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEgressProxyIpPoolAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEgressProxyIpPoolAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询出口代理地址池详细信息
//
// Description:
//
// ## 请求说明
//
// 通过提供`EgressProxyIpPoolId`，可以获取到与该ID关联的出口代理IP池的所有相关信息，包括但不限于名称、描述、网络类型等。请确保提供的`EgressProxyIpPoolId`是有效的且属于您的账户。
//
// @param request - GetEgressProxyIpPoolAttributeRequest
//
// @return GetEgressProxyIpPoolAttributeResponse
func (client *Client) GetEgressProxyIpPoolAttribute(request *GetEgressProxyIpPoolAttributeRequest) (_result *GetEgressProxyIpPoolAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEgressProxyIpPoolAttributeResponse{}
	_body, _err := client.GetEgressProxyIpPoolAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询出口代理规则组详细信息
//
// @param request - GetEgressProxyRuleGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEgressProxyRuleGroupAttributeResponse
func (client *Client) GetEgressProxyRuleGroupAttributeWithOptions(request *GetEgressProxyRuleGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetEgressProxyRuleGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRuleGroupId) {
		body["EgressProxyRuleGroupId"] = request.EgressProxyRuleGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEgressProxyRuleGroupAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEgressProxyRuleGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询出口代理规则组详细信息
//
// @param request - GetEgressProxyRuleGroupAttributeRequest
//
// @return GetEgressProxyRuleGroupAttributeResponse
func (client *Client) GetEgressProxyRuleGroupAttribute(request *GetEgressProxyRuleGroupAttributeRequest) (_result *GetEgressProxyRuleGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEgressProxyRuleGroupAttributeResponse{}
	_body, _err := client.GetEgressProxyRuleGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of egress proxy configurations.
//
// Description:
//
// *ListEgressProxyConfigurations*	- is a synchronous API operation that returns a list of egress proxy configurations that meet the specified conditions in real time after the call. The returned results are paginated by using **MaxResults*	- and **NextToken**:
//
// - If **NextToken*	- is empty, all data has been returned.
//
// - If **NextToken*	- has a value, you can use it as the **NextToken*	- for the next call to continue the query.
//
// @param request - ListEgressProxyConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEgressProxyConfigurationsResponse
func (client *Client) ListEgressProxyConfigurationsWithOptions(request *ListEgressProxyConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *ListEgressProxyConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyConfigurationIds) {
		bodyFlat["EgressProxyConfigurationIds"] = request.EgressProxyConfigurationIds
	}

	if !dara.IsNil(request.EgressProxyConfigurationNames) {
		bodyFlat["EgressProxyConfigurationNames"] = request.EgressProxyConfigurationNames
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEgressProxyConfigurations"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEgressProxyConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of egress proxy configurations.
//
// Description:
//
// *ListEgressProxyConfigurations*	- is a synchronous API operation that returns a list of egress proxy configurations that meet the specified conditions in real time after the call. The returned results are paginated by using **MaxResults*	- and **NextToken**:
//
// - If **NextToken*	- is empty, all data has been returned.
//
// - If **NextToken*	- has a value, you can use it as the **NextToken*	- for the next call to continue the query.
//
// @param request - ListEgressProxyConfigurationsRequest
//
// @return ListEgressProxyConfigurationsResponse
func (client *Client) ListEgressProxyConfigurations(request *ListEgressProxyConfigurationsRequest) (_result *ListEgressProxyConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEgressProxyConfigurationsResponse{}
	_body, _err := client.ListEgressProxyConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询出口代理网关列表
//
// Description:
//
// ## 请求说明
//
// - `MaxResults` 参数用于指定每页返回的最大结果数量，默认值为20，最大可设置为1000。
//
// - `NextToken` 参数用于分页查询时获取下一页的数据。
//
// - `Skip` 参数用于跳过指定数量的结果，默认值为0。
//
// - `VpcIds` 参数允许通过指定一个或多个VPC ID来过滤查询结果，一次最多支持20个ID。
//
// - `EgressProxyGatewayIds` 和 `EgressProxyGatewayNames` 分别允许通过实例ID和实例名称来过滤查询结果，每次请求中每个参数最多支持20个值。
//
// - `Status` 参数可用于根据实例状态（如Active, Inactive等）进行过滤。
//
// - `ResourceGroupId` 参数允许按资源组ID过滤查询结果。
//
// - `Tags` 参数支持基于标签键值对的过滤条件，可以用来更精确地定位特定资源。
//
// 请注意，所有数组类型的参数在实际使用时需要按照API的具体要求格式化输入数据。
//
// @param request - ListEgressProxyGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEgressProxyGatewaysResponse
func (client *Client) ListEgressProxyGatewaysWithOptions(request *ListEgressProxyGatewaysRequest, runtime *dara.RuntimeOptions) (_result *ListEgressProxyGatewaysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyConfigurationIds) {
		bodyFlat["EgressProxyConfigurationIds"] = request.EgressProxyConfigurationIds
	}

	if !dara.IsNil(request.EgressProxyGatewayIds) {
		bodyFlat["EgressProxyGatewayIds"] = request.EgressProxyGatewayIds
	}

	if !dara.IsNil(request.EgressProxyGatewayNames) {
		bodyFlat["EgressProxyGatewayNames"] = request.EgressProxyGatewayNames
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcIds) {
		bodyFlat["VpcIds"] = request.VpcIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEgressProxyGateways"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEgressProxyGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询出口代理网关列表
//
// Description:
//
// ## 请求说明
//
// - `MaxResults` 参数用于指定每页返回的最大结果数量，默认值为20，最大可设置为1000。
//
// - `NextToken` 参数用于分页查询时获取下一页的数据。
//
// - `Skip` 参数用于跳过指定数量的结果，默认值为0。
//
// - `VpcIds` 参数允许通过指定一个或多个VPC ID来过滤查询结果，一次最多支持20个ID。
//
// - `EgressProxyGatewayIds` 和 `EgressProxyGatewayNames` 分别允许通过实例ID和实例名称来过滤查询结果，每次请求中每个参数最多支持20个值。
//
// - `Status` 参数可用于根据实例状态（如Active, Inactive等）进行过滤。
//
// - `ResourceGroupId` 参数允许按资源组ID过滤查询结果。
//
// - `Tags` 参数支持基于标签键值对的过滤条件，可以用来更精确地定位特定资源。
//
// 请注意，所有数组类型的参数在实际使用时需要按照API的具体要求格式化输入数据。
//
// @param request - ListEgressProxyGatewaysRequest
//
// @return ListEgressProxyGatewaysResponse
func (client *Client) ListEgressProxyGateways(request *ListEgressProxyGatewaysRequest) (_result *ListEgressProxyGatewaysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEgressProxyGatewaysResponse{}
	_body, _err := client.ListEgressProxyGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询出口代理地址池列表
//
// Description:
//
// ## 请求说明
//
// - `MaxResults` 参数用于控制每页返回的结果数量，默认为20，最大值为1000。
//
// - `NextToken` 用于分页查询时获取下一页的数据。
//
// - `Skip` 参数用于跳过指定数量的记录，默认值为0。
//
// - 可以通过 `EgressProxyGatewayIds`, `EgressProxyIpPoolIds`, `EgressProxyIpPoolNames` 等参数进行过滤查询。
//
// - `Status` 参数支持按状态过滤结果。
//
// - `Tags` 参数可用于更细粒度的资源筛选。
//
// - 返回结果中包含 `NextToken` 以便于继续请求后续页面的数据。
//
// - `CreateTime` 字段表示的是UTC时间格式。
//
// @param request - ListEgressProxyIpPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEgressProxyIpPoolsResponse
func (client *Client) ListEgressProxyIpPoolsWithOptions(request *ListEgressProxyIpPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListEgressProxyIpPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyGatewayIds) {
		bodyFlat["EgressProxyGatewayIds"] = request.EgressProxyGatewayIds
	}

	if !dara.IsNil(request.EgressProxyIpPoolIds) {
		bodyFlat["EgressProxyIpPoolIds"] = request.EgressProxyIpPoolIds
	}

	if !dara.IsNil(request.EgressProxyIpPoolNames) {
		bodyFlat["EgressProxyIpPoolNames"] = request.EgressProxyIpPoolNames
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEgressProxyIpPools"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEgressProxyIpPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询出口代理地址池列表
//
// Description:
//
// ## 请求说明
//
// - `MaxResults` 参数用于控制每页返回的结果数量，默认为20，最大值为1000。
//
// - `NextToken` 用于分页查询时获取下一页的数据。
//
// - `Skip` 参数用于跳过指定数量的记录，默认值为0。
//
// - 可以通过 `EgressProxyGatewayIds`, `EgressProxyIpPoolIds`, `EgressProxyIpPoolNames` 等参数进行过滤查询。
//
// - `Status` 参数支持按状态过滤结果。
//
// - `Tags` 参数可用于更细粒度的资源筛选。
//
// - 返回结果中包含 `NextToken` 以便于继续请求后续页面的数据。
//
// - `CreateTime` 字段表示的是UTC时间格式。
//
// @param request - ListEgressProxyIpPoolsRequest
//
// @return ListEgressProxyIpPoolsResponse
func (client *Client) ListEgressProxyIpPools(request *ListEgressProxyIpPoolsRequest) (_result *ListEgressProxyIpPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEgressProxyIpPoolsResponse{}
	_body, _err := client.ListEgressProxyIpPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询出口代理规则组列表
//
// @param request - ListEgressProxyRuleGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEgressProxyRuleGroupsResponse
func (client *Client) ListEgressProxyRuleGroupsWithOptions(request *ListEgressProxyRuleGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListEgressProxyRuleGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRuleGroupIds) {
		bodyFlat["EgressProxyRuleGroupIds"] = request.EgressProxyRuleGroupIds
	}

	if !dara.IsNil(request.EgressProxyRuleGroupNames) {
		bodyFlat["EgressProxyRuleGroupNames"] = request.EgressProxyRuleGroupNames
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEgressProxyRuleGroups"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEgressProxyRuleGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询出口代理规则组列表
//
// @param request - ListEgressProxyRuleGroupsRequest
//
// @return ListEgressProxyRuleGroupsResponse
func (client *Client) ListEgressProxyRuleGroups(request *ListEgressProxyRuleGroupsRequest) (_result *ListEgressProxyRuleGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEgressProxyRuleGroupsResponse{}
	_body, _err := client.ListEgressProxyRuleGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags bound to the specified resources.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags bound to the specified resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a resource belongs.
//
// Description:
//
// Shifts a resource to another resource group. After the shift, the resource belongs to the new resource group. For valid values of the ResourceType parameter, see the metric description.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a resource belongs.
//
// Description:
//
// Shifts a resource to another resource group. After the shift, the resource belongs to the new resource group. For valid values of the ResourceType parameter, see the metric description.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从出口代理地址池移除IP地址
//
// @param request - RemoveIpAddressesFromEgressProxyIpPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveIpAddressesFromEgressProxyIpPoolResponse
func (client *Client) RemoveIpAddressesFromEgressProxyIpPoolWithOptions(request *RemoveIpAddressesFromEgressProxyIpPoolRequest, runtime *dara.RuntimeOptions) (_result *RemoveIpAddressesFromEgressProxyIpPoolResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyIpPoolId) {
		body["EgressProxyIpPoolId"] = request.EgressProxyIpPoolId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ZoneMappings) {
		bodyFlat["ZoneMappings"] = request.ZoneMappings
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveIpAddressesFromEgressProxyIpPool"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveIpAddressesFromEgressProxyIpPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从出口代理地址池移除IP地址
//
// @param request - RemoveIpAddressesFromEgressProxyIpPoolRequest
//
// @return RemoveIpAddressesFromEgressProxyIpPoolResponse
func (client *Client) RemoveIpAddressesFromEgressProxyIpPool(request *RemoveIpAddressesFromEgressProxyIpPoolRequest) (_result *RemoveIpAddressesFromEgressProxyIpPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveIpAddressesFromEgressProxyIpPoolResponse{}
	_body, _err := client.RemoveIpAddressesFromEgressProxyIpPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and attaches tags to a specified list of resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and attaches tags to a specified list of resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from the specified resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		body["All"] = request.All
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		bodyFlat["TagKey"] = request.TagKey
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from the specified resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新出口代理配置属性
//
// @param request - UpdateEgressProxyConfigurationAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEgressProxyConfigurationAttributeResponse
func (client *Client) UpdateEgressProxyConfigurationAttributeWithOptions(request *UpdateEgressProxyConfigurationAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEgressProxyConfigurationAttributeResponse, _err error) {
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

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DefaultRulePhaseActions) {
		bodyFlat["DefaultRulePhaseActions"] = request.DefaultRulePhaseActions
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyConfigurationId) {
		body["EgressProxyConfigurationId"] = request.EgressProxyConfigurationId
	}

	if !dara.IsNil(request.EgressProxyConfigurationName) {
		body["EgressProxyConfigurationName"] = request.EgressProxyConfigurationName
	}

	if !dara.IsNil(request.EgressProxyRuleGroups) {
		bodyFlat["EgressProxyRuleGroups"] = request.EgressProxyRuleGroups
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEgressProxyConfigurationAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEgressProxyConfigurationAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新出口代理配置属性
//
// @param request - UpdateEgressProxyConfigurationAttributeRequest
//
// @return UpdateEgressProxyConfigurationAttributeResponse
func (client *Client) UpdateEgressProxyConfigurationAttribute(request *UpdateEgressProxyConfigurationAttributeRequest) (_result *UpdateEgressProxyConfigurationAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEgressProxyConfigurationAttributeResponse{}
	_body, _err := client.UpdateEgressProxyConfigurationAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改出口代理网关属性
//
// @param request - UpdateEgressProxyGatewayAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEgressProxyGatewayAttributeResponse
func (client *Client) UpdateEgressProxyGatewayAttributeWithOptions(request *UpdateEgressProxyGatewayAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEgressProxyGatewayAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AddListenerProperties) {
		bodyFlat["AddListenerProperties"] = request.AddListenerProperties
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DefaultEgressProxyIpPoolId) {
		body["DefaultEgressProxyIpPoolId"] = request.DefaultEgressProxyIpPoolId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyGatewayId) {
		body["EgressProxyGatewayId"] = request.EgressProxyGatewayId
	}

	if !dara.IsNil(request.EgressProxyGatewayName) {
		body["EgressProxyGatewayName"] = request.EgressProxyGatewayName
	}

	if !dara.IsNil(request.RemoveListenerProperties) {
		bodyFlat["RemoveListenerProperties"] = request.RemoveListenerProperties
	}

	if !dara.IsNil(request.TlsInterceptProperties) {
		bodyFlat["TlsInterceptProperties"] = request.TlsInterceptProperties
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEgressProxyGatewayAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEgressProxyGatewayAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改出口代理网关属性
//
// @param request - UpdateEgressProxyGatewayAttributeRequest
//
// @return UpdateEgressProxyGatewayAttributeResponse
func (client *Client) UpdateEgressProxyGatewayAttribute(request *UpdateEgressProxyGatewayAttributeRequest) (_result *UpdateEgressProxyGatewayAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEgressProxyGatewayAttributeResponse{}
	_body, _err := client.UpdateEgressProxyGatewayAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新出口代理地址池属性
//
// Description:
//
// ## 请求说明
//
// - `EgressProxyIpPoolId` 是必填参数，用来指定要更新的出口地址池ID。
//
// - `DryRun` 参数可以设置为 `true` 来预检此次请求而不实际执行更改，默认值为 `false`。
//
// - `Scheduler` 和 `HashFactor` 参数用于定义出口IP的选择逻辑及相关的哈希因子，可根据具体需求进行设置。
//
// - 其他参数如 `EgressProxyIpPoolName`, `Description` 等均为可选项，提供给用户根据需要自定义出口地址池的名字和描述等信息。
//
// @param request - UpdateEgressProxyIpPoolAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEgressProxyIpPoolAttributeResponse
func (client *Client) UpdateEgressProxyIpPoolAttributeWithOptions(request *UpdateEgressProxyIpPoolAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEgressProxyIpPoolAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyIpPoolId) {
		body["EgressProxyIpPoolId"] = request.EgressProxyIpPoolId
	}

	if !dara.IsNil(request.EgressProxyIpPoolName) {
		body["EgressProxyIpPoolName"] = request.EgressProxyIpPoolName
	}

	if !dara.IsNil(request.HashFactor) {
		body["HashFactor"] = request.HashFactor
	}

	if !dara.IsNil(request.HashHeaderName) {
		body["HashHeaderName"] = request.HashHeaderName
	}

	if !dara.IsNil(request.Scheduler) {
		body["Scheduler"] = request.Scheduler
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEgressProxyIpPoolAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEgressProxyIpPoolAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新出口代理地址池属性
//
// Description:
//
// ## 请求说明
//
// - `EgressProxyIpPoolId` 是必填参数，用来指定要更新的出口地址池ID。
//
// - `DryRun` 参数可以设置为 `true` 来预检此次请求而不实际执行更改，默认值为 `false`。
//
// - `Scheduler` 和 `HashFactor` 参数用于定义出口IP的选择逻辑及相关的哈希因子，可根据具体需求进行设置。
//
// - 其他参数如 `EgressProxyIpPoolName`, `Description` 等均为可选项，提供给用户根据需要自定义出口地址池的名字和描述等信息。
//
// @param request - UpdateEgressProxyIpPoolAttributeRequest
//
// @return UpdateEgressProxyIpPoolAttributeResponse
func (client *Client) UpdateEgressProxyIpPoolAttribute(request *UpdateEgressProxyIpPoolAttributeRequest) (_result *UpdateEgressProxyIpPoolAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEgressProxyIpPoolAttributeResponse{}
	_body, _err := client.UpdateEgressProxyIpPoolAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新出口代理规则组属性
//
// @param request - UpdateEgressProxyRuleGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEgressProxyRuleGroupAttributeResponse
func (client *Client) UpdateEgressProxyRuleGroupAttributeWithOptions(request *UpdateEgressProxyRuleGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEgressProxyRuleGroupAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyIpPoolId) {
		body["EgressProxyIpPoolId"] = request.EgressProxyIpPoolId
	}

	if !dara.IsNil(request.EgressProxyRuleGroupId) {
		body["EgressProxyRuleGroupId"] = request.EgressProxyRuleGroupId
	}

	if !dara.IsNil(request.EgressProxyRuleGroupName) {
		body["EgressProxyRuleGroupName"] = request.EgressProxyRuleGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEgressProxyRuleGroupAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEgressProxyRuleGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新出口代理规则组属性
//
// @param request - UpdateEgressProxyRuleGroupAttributeRequest
//
// @return UpdateEgressProxyRuleGroupAttributeResponse
func (client *Client) UpdateEgressProxyRuleGroupAttribute(request *UpdateEgressProxyRuleGroupAttributeRequest) (_result *UpdateEgressProxyRuleGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEgressProxyRuleGroupAttributeResponse{}
	_body, _err := client.UpdateEgressProxyRuleGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新出口代理规则属性
//
// @param request - UpdateEgressProxyRulesAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEgressProxyRulesAttributeResponse
func (client *Client) UpdateEgressProxyRulesAttributeWithOptions(request *UpdateEgressProxyRulesAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEgressProxyRulesAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EgressProxyRuleGroupId) {
		body["EgressProxyRuleGroupId"] = request.EgressProxyRuleGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EgressProxyRules) {
		bodyFlat["EgressProxyRules"] = request.EgressProxyRules
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEgressProxyRulesAttribute"),
		Version:     dara.String("2026-07-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEgressProxyRulesAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新出口代理规则属性
//
// @param request - UpdateEgressProxyRulesAttributeRequest
//
// @return UpdateEgressProxyRulesAttributeResponse
func (client *Client) UpdateEgressProxyRulesAttribute(request *UpdateEgressProxyRulesAttributeRequest) (_result *UpdateEgressProxyRulesAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEgressProxyRulesAttributeResponse{}
	_body, _err := client.UpdateEgressProxyRulesAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
