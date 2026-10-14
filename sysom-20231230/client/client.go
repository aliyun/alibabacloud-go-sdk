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
	client.Endpoint, _err = client.GetEndpoint(dara.String("sysom"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Authorizes SysOM to diagnose ECS instances under the current account. You must call this operation to authorize diagnostics on a specific ECS instance before you can call the InvokeDiagnosis operation to initiate diagnostics on it.
//
// Description:
//
//	Notice: The diagnostics feature requires a service-linked role to be created under the Resource Access Management (RAM) user. Invoking this operation automatically checks whether the service-linked role exists and creates it if it does not. The RAM user that invokes this operation must have the ram:CreateServiceLinkedRole permission.</notice>
//
// Note the following when you invoke this operation to authorize SysOM to diagnose ECS instances:
//
// - Each authorization is valid for 7 days. After 7 days, the authorization expires and you must invoke this operation again to re-authorize.
//
// - If the SysOM service-linked role (AliyunServiceRoleForSysom) does not exist when you invoke this operation, automatic creation is performed. The Resource Access Management (RAM) user that invokes this operation must have the `ram:CreateServiceLinkedRole` permission.
//
// - When you invoke this operation to authorize diagnostics on a specific instance, the label `sysom:diagnosis` is automatically associated with the target ECS instance. SysOM only allows diagnostics on instances that have this label.
//
// @param request - AuthDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthDiagnosisResponse
func (client *Client) AuthDiagnosisWithOptions(request *AuthDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuthDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateRole) {
		body["autoCreateRole"] = request.AutoCreateRole
	}

	if !dara.IsNil(request.AutoInstallAgent) {
		body["autoInstallAgent"] = request.AutoInstallAgent
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthDiagnosis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/diagnosis/auth"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes SysOM to diagnose ECS instances under the current account. You must call this operation to authorize diagnostics on a specific ECS instance before you can call the InvokeDiagnosis operation to initiate diagnostics on it.
//
// Description:
//
//	Notice: The diagnostics feature requires a service-linked role to be created under the Resource Access Management (RAM) user. Invoking this operation automatically checks whether the service-linked role exists and creates it if it does not. The RAM user that invokes this operation must have the ram:CreateServiceLinkedRole permission.</notice>
//
// Note the following when you invoke this operation to authorize SysOM to diagnose ECS instances:
//
// - Each authorization is valid for 7 days. After 7 days, the authorization expires and you must invoke this operation again to re-authorize.
//
// - If the SysOM service-linked role (AliyunServiceRoleForSysom) does not exist when you invoke this operation, automatic creation is performed. The Resource Access Management (RAM) user that invokes this operation must have the `ram:CreateServiceLinkedRole` permission.
//
// - When you invoke this operation to authorize diagnostics on a specific instance, the label `sysom:diagnosis` is automatically associated with the target ECS instance. SysOM only allows diagnostics on instances that have this label.
//
// @param request - AuthDiagnosisRequest
//
// @return AuthDiagnosisResponse
func (client *Client) AuthDiagnosis(request *AuthDiagnosisRequest) (_result *AuthDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.AuthDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether resources such as clusters and nodes support the current configuration.
//
// Description:
//
// The instance list retrieved by this operation contains only machines managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - CheckAgentConfigSupportedRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAgentConfigSupportedResponse
func (client *Client) CheckAgentConfigSupportedWithOptions(request *CheckAgentConfigSupportedRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckAgentConfigSupportedResponse, _err error) {
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

	if !dara.IsNil(request.CheckType) {
		query["checkType"] = request.CheckType
	}

	if !dara.IsNil(request.Clusters) {
		query["clusters"] = request.Clusters
	}

	if !dara.IsNil(request.ConfigId) {
		query["configId"] = request.ConfigId
	}

	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	if !dara.IsNil(request.Instances) {
		query["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAgentConfigSupported"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/checkAgentConfigSupported"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAgentConfigSupportedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether resources such as clusters and nodes support the current configuration.
//
// Description:
//
// The instance list retrieved by this operation contains only machines managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - CheckAgentConfigSupportedRequest
//
// @return CheckAgentConfigSupportedResponse
func (client *Client) CheckAgentConfigSupported(request *CheckAgentConfigSupportedRequest) (_result *CheckAgentConfigSupportedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckAgentConfigSupportedResponse{}
	_body, _err := client.CheckAgentConfigSupportedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether target instances are supported by SysOM.
//
// Description:
//
// The instance list returned by this operation contains only instances that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - CheckInstanceSupportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceSupportResponse
func (client *Client) CheckInstanceSupportWithOptions(request *CheckInstanceSupportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceSupportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInstanceSupport"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/supportInstanceList/checkInstanceSupport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceSupportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether target instances are supported by SysOM.
//
// Description:
//
// The instance list returned by this operation contains only instances that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - CheckInstanceSupportRequest
//
// @return CheckInstanceSupportResponse
func (client *Client) CheckInstanceSupport(request *CheckInstanceSupportRequest) (_result *CheckInstanceSupportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstanceSupportResponse{}
	_body, _err := client.CheckInstanceSupportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CPU high agent streaming SSE interface.
//
// @param request - CpuHighAgentStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CpuHighAgentStreamResponseResponse
func (client *Client) CpuHighAgentStreamResponseWithSSE(request *CpuHighAgentStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *CpuHighAgentStreamResponseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.cpuHighAgentStreamResponseWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// Invokes the CPU high agent streaming SSE interface.
//
// @param request - CpuHighAgentStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CpuHighAgentStreamResponseResponse
func (client *Client) CpuHighAgentStreamResponseWithOptions(request *CpuHighAgentStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CpuHighAgentStreamResponseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LlmParamString) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CpuHighAgentStreamResponse"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/highCpuAgent/streamResponse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CpuHighAgentStreamResponseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CPU high agent streaming SSE interface.
//
// @param request - CpuHighAgentStreamResponseRequest
//
// @return CpuHighAgentStreamResponseResponse
func (client *Client) CpuHighAgentStreamResponse(request *CpuHighAgentStreamResponseRequest) (_result *CpuHighAgentStreamResponseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CpuHighAgentStreamResponseResponse{}
	_body, _err := client.CpuHighAgentStreamResponseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a component configuration.
//
// @param request - CreateAgentConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentConfigResponse
func (client *Client) CreateAgentConfigWithOptions(request *CreateAgentConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.ConfigName) {
		body["configName"] = request.ConfigName
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentConfig"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/createAgentConfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a component configuration.
//
// @param request - CreateAgentConfigRequest
//
// @return CreateAgentConfigResponse
func (client *Client) CreateAgentConfig(request *CreateAgentConfigRequest) (_result *CreateAgentConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentConfigResponse{}
	_body, _err := client.CreateAgentConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a contact for alert notifications.
//
// @param request - CreateAlertDestinationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertDestinationResponse
func (client *Client) CreateAlertDestinationWithOptions(request *CreateAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlertDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["app_id"] = request.AppId
	}

	if !dara.IsNil(request.AppSecret) {
		body["app_secret"] = request.AppSecret
	}

	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.Imbot) {
		body["imbot"] = request.Imbot
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.Target) {
		body["target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlertDestination"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/createDestination"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlertDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a contact for alert notifications.
//
// @param request - CreateAlertDestinationRequest
//
// @return CreateAlertDestinationResponse
func (client *Client) CreateAlertDestination(request *CreateAlertDestinationRequest) (_result *CreateAlertDestinationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlertDestinationResponse{}
	_body, _err := client.CreateAlertDestinationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an alert policy for push notifications.
//
// @param request - CreateAlertStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertStrategyResponse
func (client *Client) CreateAlertStrategyWithOptions(request *CreateAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlertStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.K8sLabel) {
		body["k8sLabel"] = request.K8sLabel
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlertStrategy"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/createStrategy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlertStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert policy for push notifications.
//
// @param request - CreateAlertStrategyRequest
//
// @return CreateAlertStrategyResponse
func (client *Client) CreateAlertStrategy(request *CreateAlertStrategyRequest) (_result *CreateAlertStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlertStrategyResponse{}
	_body, _err := client.CreateAlertStrategyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a VPC endpoint connection for a cluster.
//
// Description:
//
// - Use this operation together with the call_sseapi operation in the aliyun-tea-openapi-inner package.
//
// - Populate the parameters based on the input parameters of the general LLM service, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string to a dictionary before use. Refer to the response format of the general LLM service.
//
// @param request - CreateClusterVpcEndpointConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterVpcEndpointConnectionResponse
func (client *Client) CreateClusterVpcEndpointConnectionWithOptions(request *CreateClusterVpcEndpointConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterVpcEndpointConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["clusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClusterVpcEndpointConnection"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/k8sProxy/access/createClusterVpcEndpointConnection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterVpcEndpointConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a VPC endpoint connection for a cluster.
//
// Description:
//
// - Use this operation together with the call_sseapi operation in the aliyun-tea-openapi-inner package.
//
// - Populate the parameters based on the input parameters of the general LLM service, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string to a dictionary before use. Refer to the response format of the general LLM service.
//
// @param request - CreateClusterVpcEndpointConnectionRequest
//
// @return CreateClusterVpcEndpointConnectionResponse
func (client *Client) CreateClusterVpcEndpointConnection(request *CreateClusterVpcEndpointConnectionRequest) (_result *CreateClusterVpcEndpointConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterVpcEndpointConnectionResponse{}
	_body, _err := client.CreateClusterVpcEndpointConnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a SysOM instance inspection.
//
// @param request - CreateInstanceInspectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceInspectionResponse
func (client *Client) CreateInstanceInspectionWithOptions(request *CreateInstanceInspectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceInspectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Instance) {
		body["instance"] = request.Instance
	}

	if !dara.IsNil(request.Items) {
		body["items"] = request.Items
	}

	if !dara.IsNil(request.MetricSource) {
		body["metricSource"] = request.MetricSource
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceInspection"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/inspection/createInstanceInspection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceInspectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a SysOM instance inspection.
//
// @param request - CreateInstanceInspectionRequest
//
// @return CreateInstanceInspectionResponse
func (client *Client) CreateInstanceInspection(request *CreateInstanceInspectionRequest) (_result *CreateInstanceInspectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceInspectionResponse{}
	_body, _err := client.CreateInstanceInspectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an intelligent breakdown diagnostic node to diagnose the vmcore or dmesg log file specified in the input parameters.
//
// @param request - CreateVmcoreDiagnosisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVmcoreDiagnosisTaskResponse
func (client *Client) CreateVmcoreDiagnosisTaskWithOptions(request *CreateVmcoreDiagnosisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVmcoreDiagnosisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DebuginfoCommonUrl) {
		body["debuginfoCommonUrl"] = request.DebuginfoCommonUrl
	}

	if !dara.IsNil(request.DebuginfoUrl) {
		body["debuginfoUrl"] = request.DebuginfoUrl
	}

	if !dara.IsNil(request.DmesgUrl) {
		body["dmesgUrl"] = request.DmesgUrl
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	if !dara.IsNil(request.VmcoreUrl) {
		body["vmcoreUrl"] = request.VmcoreUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVmcoreDiagnosisTask"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crashAgent/diagnosis/createDiagnosisTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVmcoreDiagnosisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an intelligent breakdown diagnostic node to diagnose the vmcore or dmesg log file specified in the input parameters.
//
// @param request - CreateVmcoreDiagnosisTaskRequest
//
// @return CreateVmcoreDiagnosisTaskResponse
func (client *Client) CreateVmcoreDiagnosisTask(request *CreateVmcoreDiagnosisTaskRequest) (_result *CreateVmcoreDiagnosisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVmcoreDiagnosisTaskResponse{}
	_body, _err := client.CreateVmcoreDiagnosisTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a component configuration.
//
// @param request - DeleteAgentConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentConfigResponse
func (client *Client) DeleteAgentConfigWithOptions(request *DeleteAgentConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentConfigResponse, _err error) {
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

	if !dara.IsNil(request.ConfigId) {
		query["configId"] = request.ConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentConfig"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/deleteAgentConfig"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a component configuration.
//
// @param request - DeleteAgentConfigRequest
//
// @return DeleteAgentConfigResponse
func (client *Client) DeleteAgentConfig(request *DeleteAgentConfigRequest) (_result *DeleteAgentConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentConfigResponse{}
	_body, _err := client.DeleteAgentConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert contact.
//
// @param request - DeleteAlertDestinationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertDestinationResponse
func (client *Client) DeleteAlertDestinationWithOptions(request *DeleteAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertDestination"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/deleteDestination"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert contact.
//
// @param request - DeleteAlertDestinationRequest
//
// @return DeleteAlertDestinationResponse
func (client *Client) DeleteAlertDestination(request *DeleteAlertDestinationRequest) (_result *DeleteAlertDestinationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlertDestinationResponse{}
	_body, _err := client.DeleteAlertDestinationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert policy for push alerting.
//
// @param request - DeleteAlertStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertStrategyResponse
func (client *Client) DeleteAlertStrategyWithOptions(request *DeleteAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertStrategy"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/deleteStrategy"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert policy for push alerting.
//
// @param request - DeleteAlertStrategyRequest
//
// @return DeleteAlertStrategyResponse
func (client *Client) DeleteAlertStrategy(request *DeleteAlertStrategyRequest) (_result *DeleteAlertStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlertStrategyResponse{}
	_body, _err := client.DeleteAlertStrategyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metric data.
//
// Description:
//
// The instance list returned by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - DescribeMetricListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricListWithOptions(request *DescribeMetricListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeMetricListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.MetricName) {
		query["metricName"] = request.MetricName
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricList"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/get/describeMetricList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metric data.
//
// Description:
//
// The instance list returned by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - DescribeMetricListRequest
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricList(request *DescribeMetricListRequest) (_result *DescribeMetricListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeMetricListResponse{}
	_body, _err := client.DescribeMetricListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the response result from the Copilot service.
//
// Description:
//
// - Populate parameters according to the standard LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned data from a string to a dict before use. Refer to the standard LLM service response format.
//
// @param request - GenerateCopilotResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotResponseResponse
func (client *Client) GenerateCopilotResponseWithOptions(request *GenerateCopilotResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateCopilotResponseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LlmParamString) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCopilotResponse"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/copilot/generate_copilot_response"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the response result from the Copilot service.
//
// Description:
//
// - Populate parameters according to the standard LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned data from a string to a dict before use. Refer to the standard LLM service response format.
//
// @param request - GenerateCopilotResponseRequest
//
// @return GenerateCopilotResponseResponse
func (client *Client) GenerateCopilotResponse(request *GenerateCopilotResponseRequest) (_result *GenerateCopilotResponseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.GenerateCopilotResponseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides a streaming SSE interface for the OS Copilot service.
//
// Description:
//
// - Use this operation together with the call_sseapi interface in the aliyun-tea-openapi-inner package.
//
// - Populate the parameters based on the standard LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string data to a dict before use. Refer to the standard LLM service response format.
//
// @param request - GenerateCopilotStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotStreamResponseResponse
func (client *Client) GenerateCopilotStreamResponseWithSSE(request *GenerateCopilotStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *GenerateCopilotStreamResponseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.generateCopilotStreamResponseWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// Provides a streaming SSE interface for the OS Copilot service.
//
// Description:
//
// - Use this operation together with the call_sseapi interface in the aliyun-tea-openapi-inner package.
//
// - Populate the parameters based on the standard LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string data to a dict before use. Refer to the standard LLM service response format.
//
// @param request - GenerateCopilotStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotStreamResponseResponse
func (client *Client) GenerateCopilotStreamResponseWithOptions(request *GenerateCopilotStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateCopilotStreamResponseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LlmParamString) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCopilotStreamResponse"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/copilot/generate_copilot_stream_response"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateCopilotStreamResponseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a streaming SSE interface for the OS Copilot service.
//
// Description:
//
// - Use this operation together with the call_sseapi interface in the aliyun-tea-openapi-inner package.
//
// - Populate the parameters based on the standard LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string data to a dict before use. Refer to the standard LLM service response format.
//
// @param request - GenerateCopilotStreamResponseRequest
//
// @return GenerateCopilotStreamResponseResponse
func (client *Client) GenerateCopilotStreamResponse(request *GenerateCopilotStreamResponseRequest) (_result *GenerateCopilotStreamResponseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateCopilotStreamResponseResponse{}
	_body, _err := client.GenerateCopilotStreamResponseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries AI Infra analysis results.
//
// @param request - GetAIQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIQueryResultResponse
func (client *Client) GetAIQueryResultWithOptions(request *GetAIQueryResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAIQueryResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisId) {
		body["analysisId"] = request.AnalysisId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAIQueryResult"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/app_observ/aiAnalysis/query_result"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAIQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries AI Infra analysis results.
//
// @param request - GetAIQueryResultRequest
//
// @return GetAIQueryResultResponse
func (client *Client) GetAIQueryResult(request *GetAIQueryResultRequest) (_result *GetAIQueryResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAIQueryResultResponse{}
	_body, _err := client.GetAIQueryResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the number of unprocessed (undiagnosed) abnormal events at different severity levels for a node or pod.
//
// @param request - GetAbnormalEventsCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAbnormalEventsCountResponse
func (client *Client) GetAbnormalEventsCountWithOptions(request *GetAbnormalEventsCountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAbnormalEventsCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.Level) {
		query["level"] = request.Level
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Pod) {
		query["pod"] = request.Pod
	}

	if !dara.IsNil(request.ShowPod) {
		query["showPod"] = request.ShowPod
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAbnormalEventsCount"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/range/abnormaly_events_count"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAbnormalEventsCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the number of unprocessed (undiagnosed) abnormal events at different severity levels for a node or pod.
//
// @param request - GetAbnormalEventsCountRequest
//
// @return GetAbnormalEventsCountResponse
func (client *Client) GetAbnormalEventsCount(request *GetAbnormalEventsCountRequest) (_result *GetAbnormalEventsCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAbnormalEventsCountResponse{}
	_body, _err := client.GetAbnormalEventsCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific component.
//
// @param request - GetAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithOptions(request *GetAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["agent_id"] = request.AgentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/get_agent"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific component.
//
// @param request - GetAgentRequest
//
// @return GetAgentResponse
func (client *Client) GetAgent(request *GetAgentRequest) (_result *GetAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves component configurations.
//
// @param request - GetAgentConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentConfigResponse
func (client *Client) GetAgentConfigWithOptions(request *GetAgentConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentConfigResponse, _err error) {
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

	if !dara.IsNil(request.ConfigId) {
		query["configId"] = request.ConfigId
	}

	if !dara.IsNil(request.ConfigName) {
		query["configName"] = request.ConfigName
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentConfig"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/getAgentConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves component configurations.
//
// @param request - GetAgentConfigRequest
//
// @return GetAgentConfigResponse
func (client *Client) GetAgentConfig(request *GetAgentConfigRequest) (_result *GetAgentConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentConfigResponse{}
	_body, _err := client.GetAgentConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the execution status of an Agent installation task.
//
// @param request - GetAgentTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentTaskResponse
func (client *Client) GetAgentTaskWithOptions(request *GetAgentTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentTask"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/get_agent_task"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the execution status of an Agent installation task.
//
// @param request - GetAgentTaskRequest
//
// @return GetAgentTaskResponse
func (client *Client) GetAgentTask(request *GetAgentTaskRequest) (_result *GetAgentTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentTaskResponse{}
	_body, _err := client.GetAgentTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves diagnostic results.
//
// Description:
//
// The diagnostic process is asynchronous. When you call this operation, the diagnosis may still be in progress. Check the `data.status` field in the response to determine the current state. When `data.status == Success`, the diagnosis is complete and you can read the diagnostic results from `data.result`.
//
// @param request - GetAifgResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAifgResultResponse
func (client *Client) GetAifgResultWithOptions(request *GetAifgResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAifgResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAifgResult"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aifg/getAifgResult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAifgResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves diagnostic results.
//
// Description:
//
// The diagnostic process is asynchronous. When you call this operation, the diagnosis may still be in progress. Check the `data.status` field in the response to determine the current state. When `data.status == Success`, the diagnosis is complete and you can read the diagnostic results from `data.result`.
//
// @param request - GetAifgResultRequest
//
// @return GetAifgResultResponse
func (client *Client) GetAifgResult(request *GetAifgResultRequest) (_result *GetAifgResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAifgResultResponse{}
	_body, _err := client.GetAifgResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the URL link of a diagnostic result.
//
// Description:
//
// The diagnostic process is asynchronous. When you call this operation, the diagnosis may still be in progress. You can check the `data.status` field in the response data to determine the current state. When `data.status == Success`, the diagnosis is complete and you can read the diagnostic result from `data.result`.
//
// @param request - GetAifgUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAifgUrlResponse
func (client *Client) GetAifgUrlWithOptions(request *GetAifgUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAifgUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAifgUrl"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aifg/getAifgUrl"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAifgUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the URL link of a diagnostic result.
//
// Description:
//
// The diagnostic process is asynchronous. When you call this operation, the diagnosis may still be in progress. You can check the `data.status` field in the response data to determine the current state. When `data.status == Success`, the diagnosis is complete and you can read the diagnostic result from `data.result`.
//
// @param request - GetAifgUrlRequest
//
// @return GetAifgUrlResponse
func (client *Client) GetAifgUrl(request *GetAifgUrlRequest) (_result *GetAifgUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAifgUrlResponse{}
	_body, _err := client.GetAifgUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified alert contact.
//
// @param request - GetAlertDestinationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertDestinationResponse
func (client *Client) GetAlertDestinationWithOptions(request *GetAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlertDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertDestination"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/getDestination"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified alert contact.
//
// @param request - GetAlertDestinationRequest
//
// @return GetAlertDestinationResponse
func (client *Client) GetAlertDestination(request *GetAlertDestinationRequest) (_result *GetAlertDestinationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlertDestinationResponse{}
	_body, _err := client.GetAlertDestinationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an alert rule of a user by policy ID.
//
// @param request - GetAlertStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertStrategyResponse
func (client *Client) GetAlertStrategyWithOptions(request *GetAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlertStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertStrategy"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/getStrategy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an alert rule of a user by policy ID.
//
// @param request - GetAlertStrategyRequest
//
// @return GetAlertStrategyResponse
func (client *Client) GetAlertStrategy(request *GetAlertStrategyRequest) (_result *GetAlertStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlertStrategyResponse{}
	_body, _err := client.GetAlertStrategyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves configuration information by configuration name.
//
// @param request - GetConfigByNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigByNameResponse
func (client *Client) GetConfigByNameWithOptions(request *GetConfigByNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConfigByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["configName"] = request.ConfigName
	}

	if !dara.IsNil(request.ConfigType) {
		query["configType"] = request.ConfigType
	}

	if !dara.IsNil(request.EntityId) {
		query["entityId"] = request.EntityId
	}

	if !dara.IsNil(request.UseGlobalUid) {
		query["useGlobalUid"] = request.UseGlobalUid
	}

	if !dara.IsNil(request.VersionId) {
		query["versionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConfigByName"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/configManage/config/getConfigByName"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves configuration information by configuration name.
//
// @param request - GetConfigByNameRequest
//
// @return GetConfigByNameResponse
func (client *Client) GetConfigByName(request *GetConfigByNameRequest) (_result *GetConfigByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConfigByNameResponse{}
	_body, _err := client.GetConfigByNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the historical chat records of Copilot.
//
// @param request - GetCopilotHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCopilotHistoryResponse
func (client *Client) GetCopilotHistoryWithOptions(request *GetCopilotHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCopilotHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		body["count"] = request.Count
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCopilotHistory"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/copilot/get_copilot_history"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCopilotHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the historical chat records of Copilot.
//
// @param request - GetCopilotHistoryRequest
//
// @return GetCopilotHistoryResponse
func (client *Client) GetCopilotHistory(request *GetCopilotHistoryRequest) (_result *GetCopilotHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCopilotHistoryResponse{}
	_body, _err := client.GetCopilotHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves diagnostic results.
//
// Description:
//
// The diagnostic process is asynchronous. When you call this operation, the diagnosis may still be running. You can check the `data.status` field in the response to determine the status. When `data.status == Success`, the diagnosis is complete and you can read the diagnostic results from `data.result`.
//
// @param request - GetDiagnosisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiagnosisResultResponse
func (client *Client) GetDiagnosisResultWithOptions(request *GetDiagnosisResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDiagnosisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDiagnosisResult"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/diagnosis/get_diagnosis_results"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves diagnostic results.
//
// Description:
//
// The diagnostic process is asynchronous. When you call this operation, the diagnosis may still be running. You can check the `data.status` field in the response to determine the status. When `data.status == Success`, the diagnosis is complete and you can read the diagnostic results from `data.result`.
//
// @param request - GetDiagnosisResultRequest
//
// @return GetDiagnosisResultResponse
func (client *Client) GetDiagnosisResult(request *GetDiagnosisResultRequest) (_result *GetDiagnosisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.GetDiagnosisResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the GPU cluster parameter list.
//
// Description:
//
// After installing the component for the target ACK cluster:
//
// 1. First, when a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster has more than 50 nodes, only 50 instances are included in the first batch.
//
// 2. Then, the operating system console periodically checks the scaling status of the managed cluster. When a new ECS instance is added to the cluster, the operating system console automatically installs the component without user intervention.
//
// @param request - GetGpuClusterArgsOptionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGpuClusterArgsOptionsResponse
func (client *Client) GetGpuClusterArgsOptionsWithOptions(request *GetGpuClusterArgsOptionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGpuClusterArgsOptionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterType) {
		query["filterType"] = request.FilterType
	}

	if !dara.IsNil(request.FilterValues) {
		query["filterValues"] = request.FilterValues
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGpuClusterArgsOptions"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiTopology/topology/listClusterTopologyOptions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGpuClusterArgsOptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the GPU cluster parameter list.
//
// Description:
//
// After installing the component for the target ACK cluster:
//
// 1. First, when a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster has more than 50 nodes, only 50 instances are included in the first batch.
//
// 2. Then, the operating system console periodically checks the scaling status of the managed cluster. When a new ECS instance is added to the cluster, the operating system console automatically installs the component without user intervention.
//
// @param request - GetGpuClusterArgsOptionsRequest
//
// @return GetGpuClusterArgsOptionsResponse
func (client *Client) GetGpuClusterArgsOptions(request *GetGpuClusterArgsOptionsRequest) (_result *GetGpuClusterArgsOptionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGpuClusterArgsOptionsResponse{}
	_body, _err := client.GetGpuClusterArgsOptionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of GPU cluster alerts.
//
// Description:
//
// After installing the component on the target ACK cluster:
//
// 1. When a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster has more than 50 nodes, only the first 50 nodes are processed in the first batch.
//
// 2. The operating system console periodically checks the scaling status of managed clusters. When a new ECS instance is added to a cluster, the console automatically installs the component on it without user intervention.
//
// @param request - GetGpuClusterListAlarmRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGpuClusterListAlarmRecordResponse
func (client *Client) GetGpuClusterListAlarmRecordWithOptions(request *GetGpuClusterListAlarmRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGpuClusterListAlarmRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmCluster) {
		query["alarmCluster"] = request.AlarmCluster
	}

	if !dara.IsNil(request.AlarmExtInfo) {
		query["alarmExtInfo"] = request.AlarmExtInfo
	}

	if !dara.IsNil(request.AlarmTask) {
		query["alarmTask"] = request.AlarmTask
	}

	if !dara.IsNil(request.AlarmTime) {
		query["alarmTime"] = request.AlarmTime
	}

	if !dara.IsNil(request.AlarmType) {
		query["alarmType"] = request.AlarmType
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGpuClusterListAlarmRecord"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiTopology/topology/listAlarmRecord"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGpuClusterListAlarmRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of GPU cluster alerts.
//
// Description:
//
// After installing the component on the target ACK cluster:
//
// 1. When a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster has more than 50 nodes, only the first 50 nodes are processed in the first batch.
//
// 2. The operating system console periodically checks the scaling status of managed clusters. When a new ECS instance is added to a cluster, the console automatically installs the component on it without user intervention.
//
// @param request - GetGpuClusterListAlarmRecordRequest
//
// @return GetGpuClusterListAlarmRecordResponse
func (client *Client) GetGpuClusterListAlarmRecord(request *GetGpuClusterListAlarmRecordRequest) (_result *GetGpuClusterListAlarmRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGpuClusterListAlarmRecordResponse{}
	_body, _err := client.GetGpuClusterListAlarmRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the GPU cluster topology.
//
// Description:
//
// After you install the component on the target ACK cluster:
//
// 1. When a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster has more than 50 nodes, only the first 50 nodes are processed in the first batch.
//
// 2. The operating system console periodically checks the scaling status of managed clusters. When a new ECS instance is added to a cluster, the console automatically installs the component on it without requiring manual intervention.
//
// @param request - GetGpuClusterTopologyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGpuClusterTopologyResponse
func (client *Client) GetGpuClusterTopologyWithOptions(request *GetGpuClusterTopologyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGpuClusterTopologyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["channelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClusterId) {
		query["clusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CommunicationId) {
		query["communicationId"] = request.CommunicationId
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["instanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.PodList) {
		query["podList"] = request.PodList
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskIdList) {
		query["taskIdList"] = request.TaskIdList
	}

	if !dara.IsNil(request.TopoType) {
		query["topoType"] = request.TopoType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGpuClusterTopology"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiTopology/topology/clusterTopology"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGpuClusterTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the GPU cluster topology.
//
// Description:
//
// After you install the component on the target ACK cluster:
//
// 1. When a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster has more than 50 nodes, only the first 50 nodes are processed in the first batch.
//
// 2. The operating system console periodically checks the scaling status of managed clusters. When a new ECS instance is added to a cluster, the console automatically installs the component on it without requiring manual intervention.
//
// @param request - GetGpuClusterTopologyRequest
//
// @return GetGpuClusterTopologyResponse
func (client *Client) GetGpuClusterTopology(request *GetGpuClusterTopologyRequest) (_result *GetGpuClusterTopologyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGpuClusterTopologyResponse{}
	_body, _err := client.GetGpuClusterTopologyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the health status distribution of nodes or pods over a specified time period.
//
// @param request - GetHealthPercentageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHealthPercentageResponse
func (client *Client) GetHealthPercentageWithOptions(request *GetHealthPercentageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHealthPercentageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHealthPercentage"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/range/health_percentage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHealthPercentageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the health status distribution of nodes or pods over a specified time period.
//
// @param request - GetHealthPercentageRequest
//
// @return GetHealthPercentageResponse
func (client *Client) GetHealthPercentage(request *GetHealthPercentageRequest) (_result *GetHealthPercentageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHealthPercentageResponse{}
	_body, _err := client.GetHealthPercentageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the number of nodes or node Pods in a cluster.
//
// @param request - GetHostCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostCountResponse
func (client *Client) GetHostCountWithOptions(request *GetHostCountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHostCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostCount"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/range/host_count"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the number of nodes or node Pods in a cluster.
//
// @param request - GetHostCountRequest
//
// @return GetHostCountResponse
func (client *Client) GetHostCount(request *GetHostCountRequest) (_result *GetHostCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHostCountResponse{}
	_body, _err := client.GetHostCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of a specific field under an instance.
//
// @param request - GetHotSpotUniqListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotSpotUniqListResponse
func (client *Client) GetHotSpotUniqListWithOptions(request *GetHotSpotUniqListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotSpotUniqListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BegEnd) {
		body["beg_end"] = request.BegEnd
	}

	if !dara.IsNil(request.BegStart) {
		body["beg_start"] = request.BegStart
	}

	if !dara.IsNil(request.Instance) {
		body["instance"] = request.Instance
	}

	if !dara.IsNil(request.Pid) {
		body["pid"] = request.Pid
	}

	if !dara.IsNil(request.Table) {
		body["table"] = request.Table
	}

	if !dara.IsNil(request.Uniq) {
		body["uniq"] = request.Uniq
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotSpotUniqList"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/livetrace_proxy/hotspot_uniq_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotSpotUniqListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of a specific field under an instance.
//
// @param request - GetHotSpotUniqListRequest
//
// @return GetHotSpotUniqListResponse
func (client *Client) GetHotSpotUniqList(request *GetHotSpotUniqListRequest) (_result *GetHotSpotUniqListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotSpotUniqListResponse{}
	_body, _err := client.GetHotSpotUniqListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of an asynchronous task for CVE hotfixes on a node.
//
// Description:
//
// The hotfix task process is asynchronous. When you call this operation, the task may still be running. You can check the `data.status` field in the response to determine the task status. If `data.status` is `Success`, the task succeeded, and you can obtain the task result from the `data.result` field.
//
// @param request - GetHotfixTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotfixTaskResponse
func (client *Client) GetHotfixTaskWithOptions(request *GetHotfixTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotfixTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotfixTask"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/hotfix_manager/task/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotfixTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an asynchronous task for CVE hotfixes on a node.
//
// Description:
//
// The hotfix task process is asynchronous. When you call this operation, the task may still be running. You can check the `data.status` field in the response to determine the task status. If `data.status` is `Success`, the task succeeded, and you can obtain the task result from the `data.result` field.
//
// @param request - GetHotfixTaskRequest
//
// @return GetHotfixTaskResponse
func (client *Client) GetHotfixTask(request *GetHotfixTaskRequest) (_result *GetHotfixTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotfixTaskResponse{}
	_body, _err := client.GetHotfixTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves hot spot analysis results.
//
// @param request - GetHotspotAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotAnalysisResponse
func (client *Client) GetHotspotAnalysisWithOptions(request *GetHotspotAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["appType"] = request.AppType
	}

	if !dara.IsNil(request.BegEnd) {
		body["beg_end"] = request.BegEnd
	}

	if !dara.IsNil(request.BegStart) {
		body["beg_start"] = request.BegStart
	}

	if !dara.IsNil(request.Instance) {
		body["instance"] = request.Instance
	}

	if !dara.IsNil(request.Pid) {
		body["pid"] = request.Pid
	}

	if !dara.IsNil(request.Table) {
		body["table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotspotAnalysis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/post/livetrace_hotspot_analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotspotAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves hot spot analysis results.
//
// @param request - GetHotspotAnalysisRequest
//
// @return GetHotspotAnalysisResponse
func (client *Client) GetHotspotAnalysis(request *GetHotspotAnalysisRequest) (_result *GetHotspotAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotAnalysisResponse{}
	_body, _err := client.GetHotspotAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the results of a hot spot comparison trace.
//
// @param request - GetHotspotCompareRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotCompareResponse
func (client *Client) GetHotspotCompareWithOptions(request *GetHotspotCompareRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotCompareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Beg1End) {
		body["beg1_end"] = request.Beg1End
	}

	if !dara.IsNil(request.Beg1Start) {
		body["beg1_start"] = request.Beg1Start
	}

	if !dara.IsNil(request.Beg2End) {
		body["beg2_end"] = request.Beg2End
	}

	if !dara.IsNil(request.Beg2Start) {
		body["beg2_start"] = request.Beg2Start
	}

	if !dara.IsNil(request.HotType) {
		body["hot_type"] = request.HotType
	}

	if !dara.IsNil(request.Instance1) {
		body["instance1"] = request.Instance1
	}

	if !dara.IsNil(request.Instance2) {
		body["instance2"] = request.Instance2
	}

	if !dara.IsNil(request.Pid1) {
		body["pid1"] = request.Pid1
	}

	if !dara.IsNil(request.Pid2) {
		body["pid2"] = request.Pid2
	}

	if !dara.IsNil(request.Table) {
		body["table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotspotCompare"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/post/livetrace_hotspot_compare"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotspotCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the results of a hot spot comparison trace.
//
// @param request - GetHotspotCompareRequest
//
// @return GetHotspotCompareResponse
func (client *Client) GetHotspotCompare(request *GetHotspotCompareRequest) (_result *GetHotspotCompareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotCompareResponse{}
	_body, _err := client.GetHotspotCompareWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of hot spot instances.
//
// @param request - GetHotspotInstanceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotInstanceListResponse
func (client *Client) GetHotspotInstanceListWithOptions(request *GetHotspotInstanceListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotInstanceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BegEnd) {
		body["beg_end"] = request.BegEnd
	}

	if !dara.IsNil(request.BegStart) {
		body["beg_start"] = request.BegStart
	}

	if !dara.IsNil(request.Table) {
		body["table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotspotInstanceList"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/post/livetrace_hotspot_instance_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotspotInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of hot spot instances.
//
// @param request - GetHotspotInstanceListRequest
//
// @return GetHotspotInstanceListResponse
func (client *Client) GetHotspotInstanceList(request *GetHotspotInstanceListRequest) (_result *GetHotspotInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotInstanceListResponse{}
	_body, _err := client.GetHotspotInstanceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the PID list of a specified instance.
//
// @param request - GetHotspotPidListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotPidListResponse
func (client *Client) GetHotspotPidListWithOptions(request *GetHotspotPidListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotPidListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BegEnd) {
		body["beg_end"] = request.BegEnd
	}

	if !dara.IsNil(request.BegStart) {
		body["beg_start"] = request.BegStart
	}

	if !dara.IsNil(request.Instance) {
		body["instance"] = request.Instance
	}

	if !dara.IsNil(request.Table) {
		body["table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotspotPidList"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/post/livetrace_hotspot_pid_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotspotPidListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the PID list of a specified instance.
//
// @param request - GetHotspotPidListRequest
//
// @return GetHotspotPidListResponse
func (client *Client) GetHotspotPidList(request *GetHotspotPidListRequest) (_result *GetHotspotPidListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotPidListResponse{}
	_body, _err := client.GetHotspotPidListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves hot spot tracking results.
//
// @param request - GetHotspotTrackingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotTrackingResponse
func (client *Client) GetHotspotTrackingWithOptions(request *GetHotspotTrackingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotTrackingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BegEnd) {
		body["beg_end"] = request.BegEnd
	}

	if !dara.IsNil(request.BegStart) {
		body["beg_start"] = request.BegStart
	}

	if !dara.IsNil(request.HotType) {
		body["hot_type"] = request.HotType
	}

	if !dara.IsNil(request.Instance) {
		body["instance"] = request.Instance
	}

	if !dara.IsNil(request.Pid) {
		body["pid"] = request.Pid
	}

	if !dara.IsNil(request.Table) {
		body["table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotspotTracking"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/post/livetrace_hotspot_tracking"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotspotTrackingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves hot spot tracking results.
//
// @param request - GetHotspotTrackingRequest
//
// @return GetHotspotTrackingResponse
func (client *Client) GetHotspotTracking(request *GetHotspotTrackingRequest) (_result *GetHotspotTrackingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotTrackingResponse{}
	_body, _err := client.GetHotspotTrackingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a SysOM inspection report.
//
// @param request - GetInspectionReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInspectionReportResponse
func (client *Client) GetInspectionReportWithOptions(request *GetInspectionReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInspectionReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["reportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInspectionReport"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/inspection/getInspectionReport"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInspectionReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a SysOM inspection report.
//
// @param request - GetInspectionReportRequest
//
// @return GetInspectionReportResponse
func (client *Client) GetInspectionReport(request *GetInspectionReportRequest) (_result *GetInspectionReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInspectionReportResponse{}
	_body, _err := client.GetInspectionReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the real-time health score of a cluster or node.
//
// @param request - GetInstantScoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstantScoreResponse
func (client *Client) GetInstantScoreWithOptions(request *GetInstantScoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstantScoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstantScore"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/instant/score"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstantScoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the real-time health score of a cluster or node.
//
// @param request - GetInstantScoreRequest
//
// @return GetInstantScoreResponse
func (client *Client) GetInstantScore(request *GetInstantScoreRequest) (_result *GetInstantScoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstantScoreResponse{}
	_body, _err := client.GetInstantScoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of GPU clusters.
//
// Description:
//
// After installing the component for the target ACK cluster:
//
// 1. First, when a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster contains more than 50 nodes, only 50 instances are included in the first batch.
//
// 2. Then, the operating system console periodically checks the scaling status of the managed cluster. When a new ECS instance is added to the cluster, the operating system console automatically installs the component on it without user intervention.
//
// @param request - GetListGpuClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListGpuClusterResponse
func (client *Client) GetListGpuClusterWithOptions(request *GetListGpuClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetListGpuClusterResponse, _err error) {
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

	if !dara.IsNil(request.TimeRegion) {
		query["timeRegion"] = request.TimeRegion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetListGpuCluster"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiTopology/topology/listCluster"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetListGpuClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of GPU clusters.
//
// Description:
//
// After installing the component for the target ACK cluster:
//
// 1. First, when a cluster is managed for the first time, the component is installed on all ECS instances currently in the cluster. If the cluster contains more than 50 nodes, only 50 instances are included in the first batch.
//
// 2. Then, the operating system console periodically checks the scaling status of the managed cluster. When a new ECS instance is added to the cluster, the operating system console automatically installs the component on it without user intervention.
//
// @param request - GetListGpuClusterRequest
//
// @return GetListGpuClusterResponse
func (client *Client) GetListGpuCluster(request *GetListGpuClusterRequest) (_result *GetListGpuClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetListGpuClusterResponse{}
	_body, _err := client.GetListGpuClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of AI Infra analysis records.
//
// @param request - GetListRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListRecordResponse
func (client *Client) GetListRecordWithOptions(request *GetListRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetListRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisId) {
		query["analysisId"] = request.AnalysisId
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.CustomId) {
		query["customId"] = request.CustomId
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetListRecord"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/app_observ/aiAnalysis/list_record"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetListRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of AI Infra analysis records.
//
// @param request - GetListRecordRequest
//
// @return GetListRecordResponse
func (client *Client) GetListRecord(request *GetListRecordRequest) (_result *GetListRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetListRecordResponse{}
	_body, _err := client.GetListRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the proportion of anomaly issues for nodes in a cluster or pods in a node within a specified time range.
//
// @param request - GetProblemPercentageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProblemPercentageResponse
func (client *Client) GetProblemPercentageWithOptions(request *GetProblemPercentageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProblemPercentageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProblemPercentage"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/range/problem_percentage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProblemPercentageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the proportion of anomaly issues for nodes in a cluster or pods in a node within a specified time range.
//
// @param request - GetProblemPercentageRequest
//
// @return GetProblemPercentageResponse
func (client *Client) GetProblemPercentage(request *GetProblemPercentageRequest) (_result *GetProblemPercentageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProblemPercentageResponse{}
	_body, _err := client.GetProblemPercentageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the health score trend.
//
// @param request - GetRangeScoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRangeScoreResponse
func (client *Client) GetRangeScoreWithOptions(request *GetRangeScoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRangeScoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRangeScore"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/range/score"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRangeScoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the health score trend.
//
// @param request - GetRangeScoreRequest
//
// @return GetRangeScoreResponse
func (client *Client) GetRangeScore(request *GetRangeScoreRequest) (_result *GetRangeScoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRangeScoreResponse{}
	_body, _err := client.GetRangeScoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the real-time resource usage of a cluster or node.
//
// @param request - GetResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcesResponse
func (client *Client) GetResourcesWithOptions(request *GetResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResources"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/instant/resource"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the real-time resource usage of a cluster or node.
//
// @param request - GetResourcesRequest
//
// @return GetResourcesResponse
func (client *Client) GetResources(request *GetResourcesRequest) (_result *GetResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourcesResponse{}
	_body, _err := client.GetResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a feature module.
//
// Description:
//
// Retrieves the service configuration status.
//
// @param tmpReq - GetServiceFuncStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceFuncStatusResponse
func (client *Client) GetServiceFuncStatusWithOptions(tmpReq *GetServiceFuncStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceFuncStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetServiceFuncStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("params"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		query["channel"] = request.Channel
	}

	if !dara.IsNil(request.ParamsShrink) {
		query["params"] = request.ParamsShrink
	}

	if !dara.IsNil(request.ServiceName) {
		query["service_name"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceFuncStatus"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/func-switch/get-service-func-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceFuncStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a feature module.
//
// Description:
//
// Retrieves the service configuration status.
//
// @param request - GetServiceFuncStatusRequest
//
// @return GetServiceFuncStatusResponse
func (client *Client) GetServiceFuncStatus(request *GetServiceFuncStatusRequest) (_result *GetServiceFuncStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceFuncStatusResponse{}
	_body, _err := client.GetServiceFuncStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution status and diagnostic result of a diagnostic task by task ID.
//
// @param request - GetVmcoreDiagnosisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVmcoreDiagnosisTaskResponse
func (client *Client) GetVmcoreDiagnosisTaskWithOptions(request *GetVmcoreDiagnosisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVmcoreDiagnosisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVmcoreDiagnosisTask"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crashAgent/diagnosis/queryTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVmcoreDiagnosisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status and diagnostic result of a diagnostic task by task ID.
//
// @param request - GetVmcoreDiagnosisTaskRequest
//
// @return GetVmcoreDiagnosisTaskResponse
func (client *Client) GetVmcoreDiagnosisTask(request *GetVmcoreDiagnosisTaskRequest) (_result *GetVmcoreDiagnosisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVmcoreDiagnosisTaskResponse{}
	_body, _err := client.GetVmcoreDiagnosisTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes SysOM to ensure that the service-linked role exists.
//
// Description:
//
// Some SysOM API operations require role assumption based on the `AliyunServiceRoleForSysom` service-linked role. Before using SysOM features, invoke this operation to perform initialization and ensure that the service-linked role has been created.
//
// - `check_only`: If this parameter is set to True, only a check is performed to determine whether the service-linked role exists. The service-linked role is not created. If this parameter is set to False or left empty, automatic creation of the service-linked role is performed if it does not exist.
//
// >
//
// > Note: When you call this operation to initialize the role, you agree to the user agreement of the operating system console by default. For more information, see [Operating system console overview](https://www.alibabacloud.com/help/en/alinux/product-overview/os-console-overview) and [Alibaba Cloud Service Trial Terms](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud202001091714_51956.html).
//
// @param request - InitialSysomRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitialSysomResponse
func (client *Client) InitialSysomWithOptions(request *InitialSysomRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InitialSysomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckOnly) {
		body["check_only"] = request.CheckOnly
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitialSysom"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/initial"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InitialSysomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes SysOM to ensure that the service-linked role exists.
//
// Description:
//
// Some SysOM API operations require role assumption based on the `AliyunServiceRoleForSysom` service-linked role. Before using SysOM features, invoke this operation to perform initialization and ensure that the service-linked role has been created.
//
// - `check_only`: If this parameter is set to True, only a check is performed to determine whether the service-linked role exists. The service-linked role is not created. If this parameter is set to False or left empty, automatic creation of the service-linked role is performed if it does not exist.
//
// >
//
// > Note: When you call this operation to initialize the role, you agree to the user agreement of the operating system console by default. For more information, see [Operating system console overview](https://www.alibabacloud.com/help/en/alinux/product-overview/os-console-overview) and [Alibaba Cloud Service Trial Terms](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud202001091714_51956.html).
//
// @param request - InitialSysomRequest
//
// @return InitialSysomResponse
func (client *Client) InitialSysom(request *InitialSysomRequest) (_result *InitialSysomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitialSysomResponse{}
	_body, _err := client.InitialSysomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an agent on a specified instance.
//
// Description:
//
// Calling this operation to install an agent is asynchronous. After you call this operation, a task_id is returned. You can use this ID to call the GetAgentTask operation to query the task execution status.
//
// @param request - InstallAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentResponse
func (client *Client) InstallAgentWithOptions(request *InstallAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agent_version"] = request.AgentVersion
	}

	if !dara.IsNil(request.InstallType) {
		body["install_type"] = request.InstallType
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallAgent"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/install_agent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an agent on a specified instance.
//
// Description:
//
// Calling this operation to install an agent is asynchronous. After you call this operation, a task_id is returned. You can use this ID to call the GetAgentTask operation to query the task execution status.
//
// @param request - InstallAgentRequest
//
// @return InstallAgentResponse
func (client *Client) InstallAgent(request *InstallAgentRequest) (_result *InstallAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAgentResponse{}
	_body, _err := client.InstallAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs a component on an ACK cluster.
//
// Description:
//
// After you install a component on the target ACK cluster:
//
// 1. First, when the cluster is managed for the first time, the component is installed on all existing ECS instances in the cluster. If the cluster contains more than 50 nodes, only 50 instances are included in the first batch.
//
// 2. Then, the operating system console periodically checks the scaling status of the managed cluster. When a new ECS instance is added to the cluster, the operating system console automatically installs the component on the instance without user intervention.
//
// @param request - InstallAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentForClusterResponse
func (client *Client) InstallAgentForClusterWithOptions(request *InstallAgentForClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAgentForClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agent_version"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClusterId) {
		body["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigId) {
		body["config_id"] = request.ConfigId
	}

	if !dara.IsNil(request.GrayscaleConfig) {
		body["grayscale_config"] = request.GrayscaleConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallAgentForCluster"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/install_agent_by_cluster"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAgentForClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs a component on an ACK cluster.
//
// Description:
//
// After you install a component on the target ACK cluster:
//
// 1. First, when the cluster is managed for the first time, the component is installed on all existing ECS instances in the cluster. If the cluster contains more than 50 nodes, only 50 instances are included in the first batch.
//
// 2. Then, the operating system console periodically checks the scaling status of the managed cluster. When a new ECS instance is added to the cluster, the operating system console automatically installs the component on the instance without user intervention.
//
// @param request - InstallAgentForClusterRequest
//
// @return InstallAgentForClusterResponse
func (client *Client) InstallAgentForCluster(request *InstallAgentForClusterRequest) (_result *InstallAgentForClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAgentForClusterResponse{}
	_body, _err := client.InstallAgentForClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an Agent on a specified instance.
//
// Description:
//
// Calling this operation to install an Agent is asynchronous. After the call, a task_id is returned. You can use this ID to call the GetAgentTask operation to check the task execution status.
//
// @param request - InstallAgentWithTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentWithTypeResponse
func (client *Client) InstallAgentWithTypeWithOptions(request *InstallAgentWithTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAgentWithTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ConfigId) {
		body["configId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceType) {
		body["instanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallAgentWithType"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/installAgent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAgentWithTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an Agent on a specified instance.
//
// Description:
//
// Calling this operation to install an Agent is asynchronous. After the call, a task_id is returned. You can use this ID to call the GetAgentTask operation to check the task execution status.
//
// @param request - InstallAgentWithTypeRequest
//
// @return InstallAgentWithTypeResponse
func (client *Client) InstallAgentWithType(request *InstallAgentWithTypeRequest) (_result *InstallAgentWithTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAgentWithTypeResponse{}
	_body, _err := client.InstallAgentWithTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs CVE hotfixes asynchronously on specified instances.
//
// Description:
//
// This operation asynchronously installs the agent. After the operation is called, a task ID is returned. You can use this ID to call the GetHotfixTask operation to query the execution status of the task.
//
// @param request - InstallHotfixRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallHotfixResponse
func (client *Client) InstallHotfixWithOptions(request *InstallHotfixRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallHotfixResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CveId) {
		body["cveId"] = request.CveId
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallHotfix"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/hotfix_manager/hotfix/install"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallHotfixResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs CVE hotfixes asynchronously on specified instances.
//
// Description:
//
// This operation asynchronously installs the agent. After the operation is called, a task ID is returned. You can use this ID to call the GetHotfixTask operation to query the execution status of the task.
//
// @param request - InstallHotfixRequest
//
// @return InstallHotfixResponse
func (client *Client) InstallHotfix(request *InstallHotfixRequest) (_result *InstallHotfixResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallHotfixResponse{}
	_body, _err := client.InstallHotfixWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an anomaly diagnosis.
//
// @param request - InvokeAnomalyDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeAnomalyDiagnosisResponse
func (client *Client) InvokeAnomalyDiagnosisWithOptions(request *InvokeAnomalyDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokeAnomalyDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Uuid) {
		query["uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAnomalyDiagnosis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/diagnosis/invoke_anomaly_diagnose"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeAnomalyDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an anomaly diagnosis.
//
// @param request - InvokeAnomalyDiagnosisRequest
//
// @return InvokeAnomalyDiagnosisResponse
func (client *Client) InvokeAnomalyDiagnosis(request *InvokeAnomalyDiagnosisRequest) (_result *InvokeAnomalyDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokeAnomalyDiagnosisResponse{}
	_body, _err := client.InvokeAnomalyDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a diagnostic task.
//
// Description:
//
// The following requirements must be met to diagnose a target ECS instance:
//
// - The target ECS instance status must be Running.
//
// - Cloud Assistant must be installed on the target ECS instance. If it is not installed, refer to [Install the Cloud Assistant Agent](https://www.alibabacloud.com/help/en/ecs/user-guide/install-the-cloud-assistant-agent) for installation.
//
// - You must invoke the AuthDiagnosis operation to authorize SysOM to diagnose the target ECS instance. If authorization is not granted, this operation directly fails.
//
// - This operation depends on the SysOM service-linked role (AliyunServiceRoleForSysom) being created. This operation does not automatically create the service-linked role. If the service-linked role does not exist, first invoke AuthDiagnosis to associate the authorization. That operation creates the aforementioned service-linked role.
//
// @param request - InvokeDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeDiagnosisResponse
func (client *Client) InvokeDiagnosisWithOptions(request *InvokeDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokeDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["channel"] = request.Channel
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.ServiceName) {
		body["service_name"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeDiagnosis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/diagnosis/invoke_diagnosis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a diagnostic task.
//
// Description:
//
// The following requirements must be met to diagnose a target ECS instance:
//
// - The target ECS instance status must be Running.
//
// - Cloud Assistant must be installed on the target ECS instance. If it is not installed, refer to [Install the Cloud Assistant Agent](https://www.alibabacloud.com/help/en/ecs/user-guide/install-the-cloud-assistant-agent) for installation.
//
// - You must invoke the AuthDiagnosis operation to authorize SysOM to diagnose the target ECS instance. If authorization is not granted, this operation directly fails.
//
// - This operation depends on the SysOM service-linked role (AliyunServiceRoleForSysom) being created. This operation does not automatically create the service-linked role. If the service-linked role does not exist, first invoke AuthDiagnosis to associate the authorization. That operation creates the aforementioned service-linked role.
//
// @param request - InvokeDiagnosisRequest
//
// @return InvokeDiagnosisResponse
func (client *Client) InvokeDiagnosis(request *InvokeDiagnosisRequest) (_result *InvokeDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.InvokeDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves anomaly event information for a cluster, node, or pod within a specified time range.
//
// @param request - ListAbnormalyEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAbnormalyEventsResponse
func (client *Client) ListAbnormalyEventsWithOptions(request *ListAbnormalyEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAbnormalyEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Event) {
		query["event"] = request.Event
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.Level) {
		query["level"] = request.Level
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pod) {
		query["pod"] = request.Pod
	}

	if !dara.IsNil(request.ShowPod) {
		query["showPod"] = request.ShowPod
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAbnormalyEvents"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/range/abnormaly_events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAbnormalyEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves anomaly event information for a cluster, node, or pod within a specified time range.
//
// @param request - ListAbnormalyEventsRequest
//
// @return ListAbnormalyEventsResponse
func (client *Client) ListAbnormalyEvents(request *ListAbnormalyEventsRequest) (_result *ListAbnormalyEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAbnormalyEventsResponse{}
	_body, _err := client.ListAbnormalyEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the installation records of an Agent.
//
// @param request - ListAgentInstallRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentInstallRecordsResponse
func (client *Client) ListAgentInstallRecordsWithOptions(request *ListAgentInstallRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentInstallRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.InstanceId) {
		query["instance_id"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["plugin_id"] = request.PluginId
	}

	if !dara.IsNil(request.PluginVersion) {
		query["plugin_version"] = request.PluginVersion
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentInstallRecords"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/list_agent_install_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentInstallRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the installation records of an Agent.
//
// @param request - ListAgentInstallRecordsRequest
//
// @return ListAgentInstallRecordsResponse
func (client *Client) ListAgentInstallRecords(request *ListAgentInstallRecordsRequest) (_result *ListAgentInstallRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentInstallRecordsResponse{}
	_body, _err := client.ListAgentInstallRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of Agents.
//
// @param request - ListAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithOptions(request *ListAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
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
		Action:      dara.String("ListAgents"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/list_agents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of Agents.
//
// @param request - ListAgentsRequest
//
// @return ListAgentsResponse
func (client *Client) ListAgents(request *ListAgentsRequest) (_result *ListAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentsResponse{}
	_body, _err := client.ListAgentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of diagnostic tasks.
//
// Description:
//
// This operation retrieves the list of regions where the current user has machines managed by SysOM. If the user has ECS instances in a region but those instances are not managed by SysOM, that region does not appear in the results.
//
// @param request - ListAifgJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAifgJobsResponse
func (client *Client) ListAifgJobsWithOptions(request *ListAifgJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAifgJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAifgJobs"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aifg/listAifgJobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAifgJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of diagnostic tasks.
//
// Description:
//
// This operation retrieves the list of regions where the current user has machines managed by SysOM. If the user has ECS instances in a region but those instances are not managed by SysOM, that region does not appear in the results.
//
// @param request - ListAifgJobsRequest
//
// @return ListAifgJobsResponse
func (client *Client) ListAifgJobs(request *ListAifgJobsRequest) (_result *ListAifgJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAifgJobsResponse{}
	_body, _err := client.ListAifgJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of alert contacts.
//
// @param request - ListAlertDestinationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertDestinationsResponse
func (client *Client) ListAlertDestinationsWithOptions(request *ListAlertDestinationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertDestinationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
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

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertDestinations"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/listDestinations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertDestinationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of alert contacts.
//
// @param request - ListAlertDestinationsRequest
//
// @return ListAlertDestinationsResponse
func (client *Client) ListAlertDestinations(request *ListAlertDestinationsRequest) (_result *ListAlertDestinationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertDestinationsResponse{}
	_body, _err := client.ListAlertDestinationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all alert metrics.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertItemsResponse
func (client *Client) ListAlertItemsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertItemsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertItems"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/listItems"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all alert metrics.
//
// @return ListAlertItemsResponse
func (client *Client) ListAlertItems() (_result *ListAlertItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertItemsResponse{}
	_body, _err := client.ListAlertItemsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all alert policies configured by the user for push notifications.
//
// @param request - ListAlertStrategiesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertStrategiesResponse
func (client *Client) ListAlertStrategiesWithOptions(request *ListAlertStrategiesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
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

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertStrategies"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/listStrategies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all alert policies configured by the user for push notifications.
//
// @param request - ListAlertStrategiesRequest
//
// @return ListAlertStrategiesResponse
func (client *Client) ListAlertStrategies(request *ListAlertStrategiesRequest) (_result *ListAlertStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlertStrategiesResponse{}
	_body, _err := client.ListAlertStrategiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of managed and unmanaged instances along with their instance information.
//
// @param request - ListAllInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllInstancesResponse
func (client *Client) ListAllInstancesWithOptions(request *ListAllInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.Filters) {
		query["filters"] = request.Filters
	}

	if !dara.IsNil(request.InstanceType) {
		query["instanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.ManagedType) {
		query["managedType"] = request.ManagedType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["pluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllInstances"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/listAllInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of managed and unmanaged instances along with their instance information.
//
// @param request - ListAllInstancesRequest
//
// @return ListAllInstancesResponse
func (client *Client) ListAllInstances(request *ListAllInstancesRequest) (_result *ListAllInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAllInstancesResponse{}
	_body, _err := client.ListAllInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the installation records of cluster components.
//
// @param request - ListClusterAgentInstallRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterAgentInstallRecordsResponse
func (client *Client) ListClusterAgentInstallRecordsWithOptions(request *ListClusterAgentInstallRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterAgentInstallRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentConfigId) {
		query["agent_config_id"] = request.AgentConfigId
	}

	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["plugin_id"] = request.PluginId
	}

	if !dara.IsNil(request.PluginVersion) {
		query["plugin_version"] = request.PluginVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterAgentInstallRecords"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/list_cluster_agent_install_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterAgentInstallRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the installation records of cluster components.
//
// @param request - ListClusterAgentInstallRecordsRequest
//
// @return ListClusterAgentInstallRecordsResponse
func (client *Client) ListClusterAgentInstallRecords(request *ListClusterAgentInstallRecordsRequest) (_result *ListClusterAgentInstallRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterAgentInstallRecordsResponse{}
	_body, _err := client.ListClusterAgentInstallRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all managed clusters for the current user.
//
// @param request - ListClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterStatus) {
		query["cluster_status"] = request.ClusterStatus
	}

	if !dara.IsNil(request.ClusterType) {
		query["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/cluster/list_clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all managed clusters for the current user.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of diagnostic history records.
//
// @param request - ListDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosisResponse
func (client *Client) ListDiagnosisWithOptions(request *ListDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Params) {
		query["params"] = request.Params
	}

	if !dara.IsNil(request.ServiceName) {
		query["service_name"] = request.ServiceName
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnosis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/diagnosis/list_diagnosis"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of diagnostic history records.
//
// @param request - ListDiagnosisRequest
//
// @return ListDiagnosisResponse
func (client *Client) ListDiagnosis(request *ListDiagnosisRequest) (_result *ListDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnosisResponse{}
	_body, _err := client.ListDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the health status list of cluster nodes or pods within a specified time range.
//
// @param request - ListInstanceHealthRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceHealthResponse
func (client *Client) ListInstanceHealthWithOptions(request *ListInstanceHealthRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceHealthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceHealth"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/cluster_health/range/instance_health_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceHealthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the health status list of cluster nodes or pods within a specified time range.
//
// @param request - ListInstanceHealthRequest
//
// @return ListInstanceHealthResponse
func (client *Client) ListInstanceHealth(request *ListInstanceHealthRequest) (_result *ListInstanceHealthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceHealthResponse{}
	_body, _err := client.ListInstanceHealthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all values of a specific type of instance information.
//
// @param request - ListInstanceInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceInfoResponse
func (client *Client) ListInstanceInfoWithOptions(request *ListInstanceInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InfoType) {
		query["infoType"] = request.InfoType
	}

	if !dara.IsNil(request.InstanceType) {
		query["instanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.ManagedType) {
		query["managedType"] = request.ManagedType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PluginId) {
		query["pluginId"] = request.PluginId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceInfo"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/listInstanceInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all values of a specific type of instance information.
//
// @param request - ListInstanceInfoRequest
//
// @return ListInstanceInfoResponse
func (client *Client) ListInstanceInfo(request *ListInstanceInfoRequest) (_result *ListInstanceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceInfoResponse{}
	_body, _err := client.ListInstanceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the instance status.
//
// Description:
//
// Retrieves the list of machines managed by SysOM.
//
// @param request - ListInstanceStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceStatusResponse
func (client *Client) ListInstanceStatusWithOptions(request *ListInstanceStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceStatus"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/list_instance_status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the instance status.
//
// Description:
//
// Retrieves the list of machines managed by SysOM.
//
// @param request - ListInstanceStatusRequest
//
// @return ListInstanceStatusResponse
func (client *Client) ListInstanceStatus(request *ListInstanceStatusRequest) (_result *ListInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceStatusResponse{}
	_body, _err := client.ListInstanceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of managed instances.
//
// Description:
//
// This operation retrieves the list of instances that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/list_instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of managed instances.
//
// Description:
//
// This operation retrieves the list of instances that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves ECS information lists for instances, such as tag lists and public IP address lists.
//
// Description:
//
// The instance list retrieved by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - ListInstancesEcsInfoListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesEcsInfoListResponse
func (client *Client) ListInstancesEcsInfoListWithOptions(request *ListInstancesEcsInfoListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesEcsInfoListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InfoType) {
		query["info_type"] = request.InfoType
	}

	if !dara.IsNil(request.InstanceId) {
		query["instance_id"] = request.InstanceId
	}

	if !dara.IsNil(request.ManagedType) {
		query["managed_type"] = request.ManagedType
	}

	if !dara.IsNil(request.PluginId) {
		query["plugin_id"] = request.PluginId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstancesEcsInfoList"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/listInstancesEcsInfoList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesEcsInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves ECS information lists for instances, such as tag lists and public IP address lists.
//
// Description:
//
// The instance list retrieved by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - ListInstancesEcsInfoListRequest
//
// @return ListInstancesEcsInfoListResponse
func (client *Client) ListInstancesEcsInfoList(request *ListInstancesEcsInfoListRequest) (_result *ListInstancesEcsInfoListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesEcsInfoListResponse{}
	_body, _err := client.ListInstancesEcsInfoListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about managed and unmanaged instances, including ECS information.
//
// Description:
//
// The instance list returned by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param tmpReq - ListInstancesWithEcsInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesWithEcsInfoResponse
func (client *Client) ListInstancesWithEcsInfoWithOptions(tmpReq *ListInstancesWithEcsInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesWithEcsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListInstancesWithEcsInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceTag) {
		request.InstanceTagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceTag, dara.String("instance_tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.HealthStatus) {
		query["health_status"] = request.HealthStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["instance_id"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIdName) {
		query["instance_id_name"] = request.InstanceIdName
	}

	if !dara.IsNil(request.InstanceName) {
		query["instance_name"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceTagShrink) {
		query["instance_tag"] = request.InstanceTagShrink
	}

	if !dara.IsNil(request.IsManaged) {
		query["is_managed"] = request.IsManaged
	}

	if !dara.IsNil(request.OsName) {
		query["os_name"] = request.OsName
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateIp) {
		query["private_ip"] = request.PrivateIp
	}

	if !dara.IsNil(request.PublicIp) {
		query["public_ip"] = request.PublicIp
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resource_group_id"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupIdName) {
		query["resource_group_id_name"] = request.ResourceGroupIdName
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["resource_group_name"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstancesWithEcsInfo"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/listInstancesWithEcsInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesWithEcsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about managed and unmanaged instances, including ECS information.
//
// Description:
//
// The instance list returned by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - ListInstancesWithEcsInfoRequest
//
// @return ListInstancesWithEcsInfoResponse
func (client *Client) ListInstancesWithEcsInfo(request *ListInstancesWithEcsInfoRequest) (_result *ListInstancesWithEcsInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesWithEcsInfoResponse{}
	_body, _err := client.ListInstancesWithEcsInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of instances for plug-in installation, update, or uninstallation.
//
// Description:
//
// The instance list retrieved by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - ListPluginsInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginsInstancesResponse
func (client *Client) ListPluginsInstancesWithOptions(request *ListPluginsInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginsInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.InstanceIdName) {
		query["instance_id_name"] = request.InstanceIdName
	}

	if !dara.IsNil(request.InstanceTag) {
		query["instance_tag"] = request.InstanceTag
	}

	if !dara.IsNil(request.OperationType) {
		query["operation_type"] = request.OperationType
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["plugin_id"] = request.PluginId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPluginsInstances"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/listPluginsInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPluginsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of instances for plug-in installation, update, or uninstallation.
//
// Description:
//
// The instance list retrieved by this operation contains only machines that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - ListPluginsInstancesRequest
//
// @return ListPluginsInstancesResponse
func (client *Client) ListPluginsInstances(request *ListPluginsInstancesRequest) (_result *ListPluginsInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPluginsInstancesResponse{}
	_body, _err := client.ListPluginsInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of pods in a cluster or instance.
//
// @param request - ListPodsOfInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPodsOfInstanceResponse
func (client *Client) ListPodsOfInstanceWithOptions(request *ListPodsOfInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPodsOfInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.Current) {
		query["current"] = request.Current
	}

	if !dara.IsNil(request.Instance) {
		query["instance"] = request.Instance
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPodsOfInstance"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/list_pod_of_instance"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPodsOfInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of pods in a cluster or instance.
//
// @param request - ListPodsOfInstanceRequest
//
// @return ListPodsOfInstanceResponse
func (client *Client) ListPodsOfInstance(request *ListPodsOfInstanceRequest) (_result *ListPodsOfInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPodsOfInstanceResponse{}
	_body, _err := client.ListPodsOfInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all regions that contain managed instances.
//
// Description:
//
// This operation retrieves the list of regions where the current user has instances managed by SysOM. If a user has ECS instances in a region but none of them are managed by SysOM, that region is not included in the response.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/instance/list_regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all regions that contain managed instances.
//
// Description:
//
// This operation retrieves the list of regions where the current user has instances managed by SysOM. If a user has ECS instances in a region but none of them are managed by SysOM, that region is not included in the response.
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of historical down diagnosis tasks.
//
// @param request - ListVmcoreDiagnosisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVmcoreDiagnosisTaskResponse
func (client *Client) ListVmcoreDiagnosisTaskWithOptions(request *ListVmcoreDiagnosisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVmcoreDiagnosisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Days) {
		query["days"] = request.Days
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVmcoreDiagnosisTask"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crashAgent/diagnosis/queryTaskList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVmcoreDiagnosisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of historical down diagnosis tasks.
//
// @param request - ListVmcoreDiagnosisTaskRequest
//
// @return ListVmcoreDiagnosisTaskResponse
func (client *Client) ListVmcoreDiagnosisTask(request *ListVmcoreDiagnosisTaskRequest) (_result *ListVmcoreDiagnosisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVmcoreDiagnosisTaskResponse{}
	_body, _err := client.ListVmcoreDiagnosisTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously queries the CVE hotfix status on specified instances.
//
// Description:
//
// The call to this operation to install the agent is asynchronous. After you call this operation, a task_id is returned. You can use this ID to call the GetHotfixTask operation to obtain the execution status of the task.
//
// @param request - QueryHotfixStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHotfixStatusResponse
func (client *Client) QueryHotfixStatusWithOptions(request *QueryHotfixStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryHotfixStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CveId) {
		body["cveId"] = request.CveId
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryHotfixStatus"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/hotfix_manager/hotfix/query_status"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryHotfixStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously queries the CVE hotfix status on specified instances.
//
// Description:
//
// The call to this operation to install the agent is asynchronous. After you call this operation, a task_id is returned. You can use this ID to call the GetHotfixTask operation to obtain the execution status of the task.
//
// @param request - QueryHotfixStatusRequest
//
// @return QueryHotfixStatusResponse
func (client *Client) QueryHotfixStatus(request *QueryHotfixStatusRequest) (_result *QueryHotfixStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryHotfixStatusResponse{}
	_body, _err := client.QueryHotfixStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts AI job analysis.
//
// @param request - StartAIAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAIAnalysisResponse
func (client *Client) StartAIAnalysisWithOptions(request *StartAIAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAIAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisTool) {
		body["analysisTool"] = request.AnalysisTool
	}

	if !dara.IsNil(request.AnalysisParams) {
		body["analysis_params"] = request.AnalysisParams
	}

	if !dara.IsNil(request.Channel) {
		body["channel"] = request.Channel
	}

	if !dara.IsNil(request.Comms) {
		body["comms"] = request.Comms
	}

	if !dara.IsNil(request.CreatedBy) {
		body["created_by"] = request.CreatedBy
	}

	if !dara.IsNil(request.Instance) {
		body["instance"] = request.Instance
	}

	if !dara.IsNil(request.InstanceType) {
		body["instance_type"] = request.InstanceType
	}

	if !dara.IsNil(request.IterationFunc) {
		body["iteration_func"] = request.IterationFunc
	}

	if !dara.IsNil(request.IterationMod) {
		body["iteration_mod"] = request.IterationMod
	}

	if !dara.IsNil(request.IterationRange) {
		body["iteration_range"] = request.IterationRange
	}

	if !dara.IsNil(request.Pids) {
		body["pids"] = request.Pids
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.Timeout) {
		body["timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Uid) {
		body["uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAIAnalysis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/post/start_ai_analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAIAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts AI job analysis.
//
// @param request - StartAIAnalysisRequest
//
// @return StartAIAnalysisResponse
func (client *Client) StartAIAnalysis(request *StartAIAnalysisRequest) (_result *StartAIAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAIAnalysisResponse{}
	_body, _err := client.StartAIAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an AI Infra differential analysis.
//
// Description:
//
// Currently, only comparative analysis of the same PID across different steps within the same AI Infra analysis record is supported.
//
// @param request - StartAIDiffAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAIDiffAnalysisResponse
func (client *Client) StartAIDiffAnalysisWithOptions(request *StartAIDiffAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAIDiffAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Task1) {
		body["task1"] = request.Task1
	}

	if !dara.IsNil(request.Task2) {
		body["task2"] = request.Task2
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAIDiffAnalysis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/appObserv/aiAnalysis/diffAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAIDiffAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an AI Infra differential analysis.
//
// Description:
//
// Currently, only comparative analysis of the same PID across different steps within the same AI Infra analysis record is supported.
//
// @param request - StartAIDiffAnalysisRequest
//
// @return StartAIDiffAnalysisResponse
func (client *Client) StartAIDiffAnalysis(request *StartAIDiffAnalysisRequest) (_result *StartAIDiffAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAIDiffAnalysisResponse{}
	_body, _err := client.StartAIDiffAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers AI application diagnostics.
//
// Description:
//
// Currently, only comparative analysis of the same PID across different steps within the same AI Infra analysis record is supported.
//
// @param request - StartAifgAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAifgAnalysisResponse
func (client *Client) StartAifgAnalysisWithOptions(request *StartAifgAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAifgAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["channel"] = request.Channel
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.ServiceName) {
		body["serviceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAifgAnalysis"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aifg/startAifgAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAifgAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers AI application diagnostics.
//
// Description:
//
// Currently, only comparative analysis of the same PID across different steps within the same AI Infra analysis record is supported.
//
// @param request - StartAifgAnalysisRequest
//
// @return StartAifgAnalysisResponse
func (client *Client) StartAifgAnalysis(request *StartAifgAnalysisRequest) (_result *StartAifgAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAifgAnalysisResponse{}
	_body, _err := client.StartAifgAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls a specified version of a component.
//
// Description:
//
// Calling this operation to uninstall an Agent is asynchronous. After the call, a task_id is returned. You can use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UninstallAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentResponse
func (client *Client) UninstallAgentWithOptions(request *UninstallAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agent_version"] = request.AgentVersion
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallAgent"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/uninstall_agent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls a specified version of a component.
//
// Description:
//
// Calling this operation to uninstall an Agent is asynchronous. After the call, a task_id is returned. You can use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UninstallAgentRequest
//
// @return UninstallAgentResponse
func (client *Client) UninstallAgent(request *UninstallAgentRequest) (_result *UninstallAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallAgentResponse{}
	_body, _err := client.UninstallAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls a component from a cluster.
//
// @param request - UninstallAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentForClusterResponse
func (client *Client) UninstallAgentForClusterWithOptions(request *UninstallAgentForClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallAgentForClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agent_version"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClusterId) {
		body["cluster_id"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallAgentForCluster"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/uninstall_agent_by_cluster"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallAgentForClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls a component from a cluster.
//
// @param request - UninstallAgentForClusterRequest
//
// @return UninstallAgentForClusterResponse
func (client *Client) UninstallAgentForCluster(request *UninstallAgentForClusterRequest) (_result *UninstallAgentForClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallAgentForClusterResponse{}
	_body, _err := client.UninstallAgentForClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls a specified version of a component.
//
// Description:
//
// Calling this operation to uninstall an Agent is asynchronous. After the call, a task_id is returned. Use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UninstallAgentWithTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentWithTypeResponse
func (client *Client) UninstallAgentWithTypeWithOptions(request *UninstallAgentWithTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallAgentWithTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.InstanceType) {
		body["instanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallAgentWithType"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/uninstallAgent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallAgentWithTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls a specified version of a component.
//
// Description:
//
// Calling this operation to uninstall an Agent is asynchronous. After the call, a task_id is returned. Use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UninstallAgentWithTypeRequest
//
// @return UninstallAgentWithTypeResponse
func (client *Client) UninstallAgentWithType(request *UninstallAgentWithTypeRequest) (_result *UninstallAgentWithTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallAgentWithTypeResponse{}
	_body, _err := client.UninstallAgentWithTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously uninstalls CVE hotfixes from the specified instances.
//
// Description:
//
// This operation is asynchronous. After you call this operation, a task_id is returned. You can use this ID to call the GetHotfixTask operation to query the execution status of the task.
//
// @param request - UninstallHotfixRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallHotfixResponse
func (client *Client) UninstallHotfixWithOptions(request *UninstallHotfixRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallHotfixResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CveId) {
		body["cveId"] = request.CveId
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallHotfix"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/hotfix_manager/hotfix/uninstall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallHotfixResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously uninstalls CVE hotfixes from the specified instances.
//
// Description:
//
// This operation is asynchronous. After you call this operation, a task_id is returned. You can use this ID to call the GetHotfixTask operation to query the execution status of the task.
//
// @param request - UninstallHotfixRequest
//
// @return UninstallHotfixResponse
func (client *Client) UninstallHotfix(request *UninstallHotfixRequest) (_result *UninstallHotfixResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallHotfixResponse{}
	_body, _err := client.UninstallHotfixWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the component configuration.
//
// Description:
//
// After the configuration is updated, the new configuration is delivered to all resources that use it.
//
// @param request - UpdateAgentConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentConfigResponse
func (client *Client) UpdateAgentConfigWithOptions(request *UpdateAgentConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.ConfigId) {
		body["configId"] = request.ConfigId
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgentConfig"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/updateAgentConfig"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the component configuration.
//
// Description:
//
// After the configuration is updated, the new configuration is delivered to all resources that use it.
//
// @param request - UpdateAgentConfigRequest
//
// @return UpdateAgentConfigResponse
func (client *Client) UpdateAgentConfig(request *UpdateAgentConfigRequest) (_result *UpdateAgentConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentConfigResponse{}
	_body, _err := client.UpdateAgentConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an alert contact.
//
// Description:
//
// 、
//
// @param request - UpdateAlertDestinationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertDestinationResponse
func (client *Client) UpdateAlertDestinationWithOptions(request *UpdateAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["app_id"] = request.AppId
	}

	if !dara.IsNil(request.AppSecret) {
		body["app_secret"] = request.AppSecret
	}

	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.Imbot) {
		body["imbot"] = request.Imbot
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.Target) {
		body["target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertDestination"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/updateDestination"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alert contact.
//
// Description:
//
// 、
//
// @param request - UpdateAlertDestinationRequest
//
// @return UpdateAlertDestinationResponse
func (client *Client) UpdateAlertDestination(request *UpdateAlertDestinationRequest) (_result *UpdateAlertDestinationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlertDestinationResponse{}
	_body, _err := client.UpdateAlertDestinationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of a push alert policy.
//
// @param request - UpdateAlertEnabledRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertEnabledResponse
func (client *Client) UpdateAlertEnabledWithOptions(request *UpdateAlertEnabledRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertEnabledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertEnabled"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/updateEnabled"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of a push alert policy.
//
// @param request - UpdateAlertEnabledRequest
//
// @return UpdateAlertEnabledResponse
func (client *Client) UpdateAlertEnabled(request *UpdateAlertEnabledRequest) (_result *UpdateAlertEnabledResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlertEnabledResponse{}
	_body, _err := client.UpdateAlertEnabledWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a push alert policy.
//
// @param request - UpdateAlertStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertStrategyResponse
func (client *Client) UpdateAlertStrategyWithOptions(request *UpdateAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.K8sLabel) {
		body["k8sLabel"] = request.K8sLabel
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertStrategy"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/alertPusher/alert/updateStrategy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a push alert policy.
//
// @param request - UpdateAlertStrategyRequest
//
// @return UpdateAlertStrategyResponse
func (client *Client) UpdateAlertStrategy(request *UpdateAlertStrategyRequest) (_result *UpdateAlertStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlertStrategyResponse{}
	_body, _err := client.UpdateAlertStrategyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a cluster component.
//
// @param request - UpdateClusterAgentConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterAgentConfigResponse
func (client *Client) UpdateClusterAgentConfigWithOptions(request *UpdateClusterAgentConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateClusterAgentConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.ClusterId) {
		body["clusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigId) {
		body["configId"] = request.ConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClusterAgentConfig"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/updateClusterAgentConfig"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterAgentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a cluster component.
//
// @param request - UpdateClusterAgentConfigRequest
//
// @return UpdateClusterAgentConfigResponse
func (client *Client) UpdateClusterAgentConfig(request *UpdateClusterAgentConfigRequest) (_result *UpdateClusterAgentConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateClusterAgentConfigResponse{}
	_body, _err := client.UpdateClusterAgentConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attention level of an anomaly item. Adjusting the attention level affects the sensitivity of the anomaly detection algorithm.
//
// @param request - UpdateEventsAttentionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventsAttentionResponse
func (client *Client) UpdateEventsAttentionWithOptions(request *UpdateEventsAttentionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEventsAttentionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mode) {
		body["mode"] = request.Mode
	}

	if !dara.IsNil(request.Range) {
		body["range"] = request.Range
	}

	if !dara.IsNil(request.Uuid) {
		body["uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventsAttention"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/openapi/proxy/post/cluster_update_events_attention"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventsAttentionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attention level of an anomaly item. Adjusting the attention level affects the sensitivity of the anomaly detection algorithm.
//
// @param request - UpdateEventsAttentionRequest
//
// @return UpdateEventsAttentionResponse
func (client *Client) UpdateEventsAttention(request *UpdateEventsAttentionRequest) (_result *UpdateEventsAttentionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEventsAttentionResponse{}
	_body, _err := client.UpdateEventsAttentionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a service feature module.
//
// Description:
//
// - Populate parameters according to the general LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned data from a string to a dict before use. Refer to the general LLM service response format.
//
// @param tmpReq - UpdateFuncSwitchRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFuncSwitchRecordResponse
func (client *Client) UpdateFuncSwitchRecordWithOptions(tmpReq *UpdateFuncSwitchRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFuncSwitchRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFuncSwitchRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("params"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		query["channel"] = request.Channel
	}

	if !dara.IsNil(request.ParamsShrink) {
		query["params"] = request.ParamsShrink
	}

	if !dara.IsNil(request.ServiceName) {
		query["service_name"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFuncSwitchRecord"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/func-switch/update-service-func-switch"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFuncSwitchRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a service feature module.
//
// Description:
//
// - Populate parameters according to the general LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned data from a string to a dict before use. Refer to the general LLM service response format.
//
// @param request - UpdateFuncSwitchRecordRequest
//
// @return UpdateFuncSwitchRecordResponse
func (client *Client) UpdateFuncSwitchRecord(request *UpdateFuncSwitchRecordRequest) (_result *UpdateFuncSwitchRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFuncSwitchRecordResponse{}
	_body, _err := client.UpdateFuncSwitchRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the component configuration of an instance.
//
// @param request - UpdateInstanceAgentConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceAgentConfigResponse
func (client *Client) UpdateInstanceAgentConfigWithOptions(request *UpdateInstanceAgentConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceAgentConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.ConfigId) {
		body["configId"] = request.ConfigId
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceAgentConfig"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/updateInstanceAgentConfig"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceAgentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the component configuration of an instance.
//
// @param request - UpdateInstanceAgentConfigRequest
//
// @return UpdateInstanceAgentConfigResponse
func (client *Client) UpdateInstanceAgentConfig(request *UpdateInstanceAgentConfigRequest) (_result *UpdateInstanceAgentConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceAgentConfigResponse{}
	_body, _err := client.UpdateInstanceAgentConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an installed component to a specified version.
//
// Description:
//
// Calling this operation to update an Agent is asynchronous. After the call, a task_id is returned. Use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UpgradeAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentResponse
func (client *Client) UpgradeAgentWithOptions(request *UpgradeAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agent_version"] = request.AgentVersion
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeAgent"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/upgrade_agent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an installed component to a specified version.
//
// Description:
//
// Calling this operation to update an Agent is asynchronous. After the call, a task_id is returned. Use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UpgradeAgentRequest
//
// @return UpgradeAgentResponse
func (client *Client) UpgradeAgent(request *UpgradeAgentRequest) (_result *UpgradeAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeAgentResponse{}
	_body, _err := client.UpgradeAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a component for a cluster.
//
// @param request - UpgradeAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentForClusterResponse
func (client *Client) UpgradeAgentForClusterWithOptions(request *UpgradeAgentForClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeAgentForClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agent_version"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClusterId) {
		body["cluster_id"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeAgentForCluster"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/upgrade_agent_by_cluster"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeAgentForClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a component for a cluster.
//
// @param request - UpgradeAgentForClusterRequest
//
// @return UpgradeAgentForClusterResponse
func (client *Client) UpgradeAgentForCluster(request *UpgradeAgentForClusterRequest) (_result *UpgradeAgentForClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeAgentForClusterResponse{}
	_body, _err := client.UpgradeAgentForClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the version of an installed component to a specified version.
//
// Description:
//
// Calling this operation to update an Agent is asynchronous. After the call, a task_id is returned. Use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UpgradeAgentWithTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentWithTypeResponse
func (client *Client) UpgradeAgentWithTypeWithOptions(request *UpgradeAgentWithTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeAgentWithTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		body["agentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.InstanceType) {
		body["instanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeAgentWithType"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/am/agent/upgradeAgent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeAgentWithTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version of an installed component to a specified version.
//
// Description:
//
// Calling this operation to update an Agent is asynchronous. After the call, a task_id is returned. Use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UpgradeAgentWithTypeRequest
//
// @return UpgradeAgentWithTypeResponse
func (client *Client) UpgradeAgentWithType(request *UpgradeAgentWithTypeRequest) (_result *UpgradeAgentWithTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeAgentWithTypeResponse{}
	_body, _err := client.UpgradeAgentWithTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) cpuHighAgentStreamResponseWithSSE_opYieldFunc(_yield chan *CpuHighAgentStreamResponseResponse, _yieldErr chan error, request *CpuHighAgentStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LlmParamString) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CpuHighAgentStreamResponse"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/highCpuAgent/streamResponse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

func (client *Client) generateCopilotStreamResponseWithSSE_opYieldFunc(_yield chan *GenerateCopilotStreamResponseResponse, _yieldErr chan error, request *GenerateCopilotStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LlmParamString) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCopilotStreamResponse"),
		Version:     dara.String("2023-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/copilot/generate_copilot_stream_response"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
