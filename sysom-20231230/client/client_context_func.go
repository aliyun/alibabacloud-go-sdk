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
// Authorizes SysOM to diagnose ECS instances under the current account. You must call this operation to authorize diagnostics for a specific ECS instance before you can call the InvokeDiagnosis operation to initiate diagnostics on it.
//
// Description:
//
//	Notice: The diagnostics feature requires a service-linked role to be created under the Resource Access Management (RAM) user. This operation automatically checks whether the service-linked role exists and creates it if it does not. The RAM user that invokes this operation must have the ram:CreateServiceLinkedRole permission.</notice>
//
// Note the following when you invoke this operation to authorize SysOM to diagnose ECS instances:
//
// - Each authorization is valid for 7 days. After the authorization expires, invoke this operation again to re-authorize.
//
// - If the SysOM service-linked role (AliyunServiceRoleForSysom) does not exist when you invoke this operation, automatic creation is performed. The RAM user that invokes this operation must have the `ram:CreateServiceLinkedRole` permission.
//
// - When you invoke this operation to authorize diagnostics for a specific instance, the label `sysom:diagnosis` is automatically associated with the target ECS instance. SysOM only allows diagnostics on instances that have this label.
//
// @param request - AuthDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthDiagnosisResponse
func (client *Client) AuthDiagnosisWithContext(ctx context.Context, request *AuthDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuthDiagnosisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a target instance is supported by SysOM.
//
// Description:
//
// This operation retrieves the list of instances that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - CheckInstanceSupportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceSupportResponse
func (client *Client) CheckInstanceSupportWithContext(ctx context.Context, request *CheckInstanceSupportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceSupportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CPU High Agent streaming SSE interface.
//
// @param request - CpuHighAgentStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CpuHighAgentStreamResponseResponse
func (client *Client) CpuHighAgentStreamResponseWithSSECtx(ctx context.Context, request *CpuHighAgentStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *CpuHighAgentStreamResponseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.cpuHighAgentStreamResponseWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Calls the CPU High Agent streaming SSE interface.
//
// @param request - CpuHighAgentStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CpuHighAgentStreamResponseResponse
func (client *Client) CpuHighAgentStreamResponseWithContext(ctx context.Context, request *CpuHighAgentStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CpuHighAgentStreamResponseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertDestinationResponse
func (client *Client) CreateAlertDestinationWithContext(ctx context.Context, request *CreateAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlertDestinationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert push strategy.
//
// @param request - CreateAlertStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertStrategyResponse
func (client *Client) CreateAlertStrategyWithContext(ctx context.Context, request *CreateAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlertStrategyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Use this operation with the call_sseapi interface of the aliyun-tea-openapi-inner package.
//
// - Populate parameters according to the general LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string to a dictionary before use. Refer to the general LLM service response format.
//
// @param request - CreateClusterVpcEndpointConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterVpcEndpointConnectionResponse
func (client *Client) CreateClusterVpcEndpointConnectionWithContext(ctx context.Context, request *CreateClusterVpcEndpointConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterVpcEndpointConnectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceInspectionResponse
func (client *Client) CreateInstanceInspectionWithContext(ctx context.Context, request *CreateInstanceInspectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceInspectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an intelligent breakdown diagnostic node that diagnoses the specified vmcore or dmesg log file based on the input parameters.
//
// @param request - CreateVmcoreDiagnosisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVmcoreDiagnosisTaskResponse
func (client *Client) CreateVmcoreDiagnosisTaskWithContext(ctx context.Context, request *CreateVmcoreDiagnosisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVmcoreDiagnosisTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertDestinationResponse
func (client *Client) DeleteAlertDestinationWithContext(ctx context.Context, request *DeleteAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertDestinationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert policy for push notifications.
//
// @param request - DeleteAlertStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertStrategyResponse
func (client *Client) DeleteAlertStrategyWithContext(ctx context.Context, request *DeleteAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlertStrategyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// The instance list returned by this operation contains only instances that are managed by SysOM. If an ECS instance exists but is not managed by SysOM, it does not appear in the list.
//
// @param request - DescribeMetricListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricListWithContext(ctx context.Context, request *DescribeMetricListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeMetricListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the response result of the copilot service
//
// Description:
//
// - Parameters need to be filled in according to the standard LLM service input parameters, converted to a string, and assigned to llmParamString
//
// - The returned data needs to be converted from string to dict before use. Refer to the standard LLM service response format
//
// @param request - GenerateCopilotResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotResponseResponse
func (client *Client) GenerateCopilotResponseWithContext(ctx context.Context, request *GenerateCopilotResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateCopilotResponseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the streaming SSE endpoint of the OS Copilot service.
//
// Description:
//
// - Use this operation together with the call_sseapi operation in the aliyun-tea-openapi-inner package.
//
// - Populate the parameters based on the standard LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string to a dictionary before use. Refer to the standard LLM service response format.
//
// @param request - GenerateCopilotStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotStreamResponseResponse
func (client *Client) GenerateCopilotStreamResponseWithSSECtx(ctx context.Context, request *GenerateCopilotStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *GenerateCopilotStreamResponseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.generateCopilotStreamResponseWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Calls the streaming SSE endpoint of the OS Copilot service.
//
// Description:
//
// - Use this operation together with the call_sseapi operation in the aliyun-tea-openapi-inner package.
//
// - Populate the parameters based on the standard LLM service input parameters, convert them to a string, and assign the string to llmParamString.
//
// - Convert the returned string to a dictionary before use. Refer to the standard LLM service response format.
//
// @param request - GenerateCopilotStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotStreamResponseResponse
func (client *Client) GenerateCopilotStreamResponseWithContext(ctx context.Context, request *GenerateCopilotStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateCopilotStreamResponseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AI Infra analysis results.
//
// @param request - GetAIQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIQueryResultResponse
func (client *Client) GetAIQueryResultWithContext(ctx context.Context, request *GetAIQueryResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAIQueryResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the count of unhandled (undiagnosed) abnormal events of different levels for nodes/Pods
//
// @param request - GetAbnormalEventsCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAbnormalEventsCountResponse
func (client *Client) GetAbnormalEventsCountWithContext(ctx context.Context, request *GetAbnormalEventsCountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAbnormalEventsCountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get details of a specific agent
//
// @param request - GetAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithContext(ctx context.Context, request *GetAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentTaskResponse
func (client *Client) GetAgentTaskWithContext(ctx context.Context, request *GetAgentTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertDestinationResponse
func (client *Client) GetAlertDestinationWithContext(ctx context.Context, request *GetAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlertDestinationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an alert for a user based on the policy ID.
//
// @param request - GetAlertStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertStrategyResponse
func (client *Client) GetAlertStrategyWithContext(ctx context.Context, request *GetAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlertStrategyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the chat history of Copilot.
//
// @param request - GetCopilotHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCopilotHistoryResponse
func (client *Client) GetCopilotHistoryWithContext(ctx context.Context, request *GetCopilotHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCopilotHistoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the diagnostic result.
//
// Description:
//
// The diagnostic process is asynchronous. When you call this operation, the diagnosis may still be in progress. You can check the `data.status` field in the response to determine the status. When `data.status == Success`, the diagnosis is complete and you can read the diagnostic result from `data.result`.
//
// @param request - GetDiagnosisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiagnosisResultResponse
func (client *Client) GetDiagnosisResultWithContext(ctx context.Context, request *GetDiagnosisResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDiagnosisResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHealthPercentageResponse
func (client *Client) GetHealthPercentageWithContext(ctx context.Context, request *GetHealthPercentageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHealthPercentageResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the number of nodes or the number of Pods on nodes in a cluster.
//
// @param request - GetHostCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostCountResponse
func (client *Client) GetHostCountWithContext(ctx context.Context, request *GetHostCountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHostCountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the list of a specific field under an instance.
//
// @param request - GetHotSpotUniqListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotSpotUniqListResponse
func (client *Client) GetHotSpotUniqListWithContext(ctx context.Context, request *GetHotSpotUniqListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotSpotUniqListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotAnalysisResponse
func (client *Client) GetHotspotAnalysisWithContext(ctx context.Context, request *GetHotspotAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotAnalysisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Hotspot Comparison Tracing Results
//
// @param request - GetHotspotCompareRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotCompareResponse
func (client *Client) GetHotspotCompareWithContext(ctx context.Context, request *GetHotspotCompareRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotCompareResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Hotspot Instance List
//
// @param request - GetHotspotInstanceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotInstanceListResponse
func (client *Client) GetHotspotInstanceListWithContext(ctx context.Context, request *GetHotspotInstanceListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotInstanceListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotPidListResponse
func (client *Client) GetHotspotPidListWithContext(ctx context.Context, request *GetHotspotPidListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotPidListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotTrackingResponse
func (client *Client) GetHotspotTrackingWithContext(ctx context.Context, request *GetHotspotTrackingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHotspotTrackingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInspectionReportResponse
func (client *Client) GetInspectionReportWithContext(ctx context.Context, request *GetInspectionReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInspectionReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get real-time cluster/node health score
//
// @param request - GetInstantScoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstantScoreResponse
func (client *Client) GetInstantScoreWithContext(ctx context.Context, request *GetInstantScoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstantScoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListRecordResponse
func (client *Client) GetListRecordWithContext(ctx context.Context, request *GetListRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetListRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the proportion of abnormal issues in cluster nodes/pods within a specified time range
//
// @param request - GetProblemPercentageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProblemPercentageResponse
func (client *Client) GetProblemPercentageWithContext(ctx context.Context, request *GetProblemPercentageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProblemPercentageResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRangeScoreResponse
func (client *Client) GetRangeScoreWithContext(ctx context.Context, request *GetRangeScoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRangeScoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcesResponse
func (client *Client) GetResourcesWithContext(ctx context.Context, request *GetResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetServiceFuncStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceFuncStatusResponse
func (client *Client) GetServiceFuncStatusWithContext(ctx context.Context, tmpReq *GetServiceFuncStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceFuncStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVmcoreDiagnosisTaskResponse
func (client *Client) GetVmcoreDiagnosisTaskWithContext(ctx context.Context, request *GetVmcoreDiagnosisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVmcoreDiagnosisTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - `check_only`: If this parameter is set to True, the operation only checks whether the service-linked role exists and does not create it. If this parameter is set to False or left empty, the operation performs automatic creation of the service-linked role if it does not exist.
//
// >
//
// > Note: When you call this operation to initialize the role through the API, you agree to the user agreement of the operating system console by default. For more information, see [Operating system console overview](https://www.alibabacloud.com/help/en/alinux/product-overview/os-console-overview) and [Alibaba Cloud Service Trial Terms](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud202001091714_51956.html).
//
// @param request - InitialSysomRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitialSysomResponse
func (client *Client) InitialSysomWithContext(ctx context.Context, request *InitialSysomRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InitialSysomResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Calling this operation to install an Agent is asynchronous. After the call, a task_id is returned. You can use this ID to call the GetAgentTask operation to retrieve the task execution status.
//
// @param request - InstallAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentResponse
func (client *Client) InstallAgentWithContext(ctx context.Context, request *InstallAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAgentResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Install component for cluster
//
// Description:
//
// After installing a component for the target ACK cluster:
//
// 1. First, when the cluster is managed for the first time, the component will be installed on all ECS instances currently in the cluster. If the cluster has more than 50 nodes, only 50 instances will be covered in the first batch.
//
// 2. Then, the SysOM console periodically checks the scaling status of the managed cluster. Once a new ECS instance is added to the cluster, the SysOM console automatically installs the component on it without user intervention.
//
// @param request - InstallAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentForClusterResponse
func (client *Client) InstallAgentForClusterWithContext(ctx context.Context, request *InstallAgentForClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAgentForClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Calling this operation to install an Agent is asynchronous. After the call, a task_id is returned. You can use this ID to call the GetAgentTask operation to retrieve the task execution status.
//
// @param request - InstallAgentWithTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentWithTypeResponse
func (client *Client) InstallAgentWithTypeWithContext(ctx context.Context, request *InstallAgentWithTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAgentWithTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an anomaly diagnostics task.
//
// @param request - InvokeAnomalyDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeAnomalyDiagnosisResponse
func (client *Client) InvokeAnomalyDiagnosisWithContext(ctx context.Context, request *InvokeAnomalyDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokeAnomalyDiagnosisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiate Diagnosis.
//
// Description:
//
// The following requirements must be met to diagnose a target ECS instance:
//
// - The target ECS instance must be in the Running state.
//
// - The Cloud Assistant must be installed on the target ECS instance. If it is not installed, refer to [Install the Cloud Assistant Agent](https://help.aliyun.com/zh/ecs/user-guide/install-the-cloud-assistant-agent) for installation.
//
// - You must call the AuthDiagnosis API to authorize SysOM to diagnose the target ECS instance. If authorization is not granted, this API will fail directly.
//
// - This API requires that the SysOM service-linked role (AliyunServiceRoleForSysom) has been created. This API does not automatically create the service role. If the service role does not exist, you must first call AuthDiagnosis for authorization, which will create the aforementioned service role.
//
// @param request - InvokeDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeDiagnosisResponse
func (client *Client) InvokeDiagnosisWithContext(ctx context.Context, request *InvokeDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokeDiagnosisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAbnormalyEventsResponse
func (client *Client) ListAbnormalyEventsWithContext(ctx context.Context, request *ListAbnormalyEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAbnormalyEventsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentInstallRecordsResponse
func (client *Client) ListAgentInstallRecordsWithContext(ctx context.Context, request *ListAgentInstallRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentInstallRecordsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of agents.
//
// @param request - ListAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithContext(ctx context.Context, request *ListAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # This API is used to get the list of alert contacts
//
// @param request - ListAlertDestinationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertDestinationsResponse
func (client *Client) ListAlertDestinationsWithContext(ctx context.Context, request *ListAlertDestinationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertDestinationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertItemsResponse
func (client *Client) ListAlertItemsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertItemsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all push alert policies for the current user.
//
// @param request - ListAlertStrategiesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertStrategiesResponse
func (client *Client) ListAlertStrategiesWithContext(ctx context.Context, request *ListAlertStrategiesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertStrategiesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is used to retrieve a list of managed/unmanaged instances along with their instance information.
//
// @param request - ListAllInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllInstancesResponse
func (client *Client) ListAllInstancesWithContext(ctx context.Context, request *ListAllInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get cluster component installation records
//
// @param request - ListClusterAgentInstallRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterAgentInstallRecordsResponse
func (client *Client) ListClusterAgentInstallRecordsWithContext(ctx context.Context, request *ListClusterAgentInstallRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterAgentInstallRecordsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve all managed clusters of the current user
//
// @param request - ListClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithContext(ctx context.Context, request *ListClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the list of diagnostic history.
//
// @param request - ListDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosisResponse
func (client *Client) ListDiagnosisWithContext(ctx context.Context, request *ListDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDiagnosisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the health status list of cluster nodes or Pods within a specified time range.
//
// @param request - ListInstanceHealthRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceHealthResponse
func (client *Client) ListInstanceHealthWithContext(ctx context.Context, request *ListInstanceHealthRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceHealthResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves instance statuses.
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
func (client *Client) ListInstanceStatusWithContext(ctx context.Context, request *ListInstanceStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of instances.
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
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves lists of ECS information for instances, such as tag lists and public IP address lists.
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
func (client *Client) ListInstancesEcsInfoListWithContext(ctx context.Context, request *ListInstancesEcsInfoListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesEcsInfoListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListInstancesWithEcsInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesWithEcsInfoResponse
func (client *Client) ListInstancesWithEcsInfoWithContext(ctx context.Context, tmpReq *ListInstancesWithEcsInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesWithEcsInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginsInstancesResponse
func (client *Client) ListPluginsInstancesWithContext(ctx context.Context, request *ListPluginsInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginsInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPodsOfInstanceResponse
func (client *Client) ListPodsOfInstanceWithContext(ctx context.Context, request *ListPodsOfInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPodsOfInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the historical crash diagnosis task list.
//
// @param request - ListVmcoreDiagnosisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVmcoreDiagnosisTaskResponse
func (client *Client) ListVmcoreDiagnosisTaskWithContext(ctx context.Context, request *ListVmcoreDiagnosisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVmcoreDiagnosisTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start AI job analysis.
//
// @param request - StartAIAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAIAnalysisResponse
func (client *Client) StartAIAnalysisWithContext(ctx context.Context, request *StartAIAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAIAnalysisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Currently, only comparative analysis of the same pid across different steps within the same AI Infra analysis record is supported.
//
// @param request - StartAIDiffAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAIDiffAnalysisResponse
func (client *Client) StartAIDiffAnalysisWithContext(ctx context.Context, request *StartAIDiffAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAIDiffAnalysisResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param request - UninstallAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentResponse
func (client *Client) UninstallAgentWithContext(ctx context.Context, request *UninstallAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallAgentResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentForClusterResponse
func (client *Client) UninstallAgentForClusterWithContext(ctx context.Context, request *UninstallAgentForClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallAgentForClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param request - UninstallAgentWithTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentWithTypeResponse
func (client *Client) UninstallAgentWithTypeWithContext(ctx context.Context, request *UninstallAgentWithTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallAgentWithTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// .
//
// @param request - UpdateAlertDestinationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertDestinationResponse
func (client *Client) UpdateAlertDestinationWithContext(ctx context.Context, request *UpdateAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertDestinationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertEnabledResponse
func (client *Client) UpdateAlertEnabledWithContext(ctx context.Context, request *UpdateAlertEnabledRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertEnabledResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertStrategyResponse
func (client *Client) UpdateAlertStrategyWithContext(ctx context.Context, request *UpdateAlertStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertStrategyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventsAttentionResponse
func (client *Client) UpdateEventsAttentionWithContext(ctx context.Context, request *UpdateEventsAttentionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEventsAttentionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateFuncSwitchRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFuncSwitchRecordResponse
func (client *Client) UpdateFuncSwitchRecordWithContext(ctx context.Context, tmpReq *UpdateFuncSwitchRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFuncSwitchRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Updating the Agent by calling this operation is asynchronous. After you call this operation, a task_id is returned. You can use this ID to call the GetAgentTask operation to query the execution status of the task.
//
// @param request - UpgradeAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentResponse
func (client *Client) UpgradeAgentWithContext(ctx context.Context, request *UpgradeAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeAgentResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates components for a cluster.
//
// @param request - UpgradeAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentForClusterResponse
func (client *Client) UpgradeAgentForClusterWithContext(ctx context.Context, request *UpgradeAgentForClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeAgentForClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Calling this operation to update the Agent is asynchronous. After the call, a task_id is returned. You can use this ID to call the GetAgentTask operation to retrieve the execution status of the task.
//
// @param request - UpgradeAgentWithTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentWithTypeResponse
func (client *Client) UpgradeAgentWithTypeWithContext(ctx context.Context, request *UpgradeAgentWithTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeAgentWithTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) cpuHighAgentStreamResponseWithSSECtx_opYieldFunc(_yield chan *CpuHighAgentStreamResponseResponse, _yieldErr chan error, ctx context.Context, request *CpuHighAgentStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
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

func (client *Client) generateCopilotStreamResponseWithSSECtx_opYieldFunc(_yield chan *GenerateCopilotStreamResponseResponse, _yieldErr chan error, ctx context.Context, request *GenerateCopilotStreamResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
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
