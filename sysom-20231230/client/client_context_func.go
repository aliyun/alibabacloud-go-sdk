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
// This API is used to authorize SysOM to diagnose ECS instances under your account. You can only invoke the InvokeDiagnosis API to initiate diagnosis on a specific ECS instance after authorizing it through this API.
//
// Description:
//
//	Notice: The diagnosis feature requires a service-linked role to be created under a Resource Access Management (RAM) user. When you call this API, it automatically checks whether the service-linked role exists. If the role does not exist, the API automatically creates it. This requires the RAM user invoking this API to have the ram:CreateServiceLinkedRole permission.</notice>
//
// When calling this API to authorize SysOM to diagnose ECS instances, note the following:
//
// - Each authorization is valid for 7 days. After 7 days, the authorization expires, and you must call this API again to re-authorize.
//
// - If the SysOM service-linked role (AliyunServiceRoleForSysom) does not exist when you call this API, it will be automatically created. This requires the RAM user invoking this API to have the `ram:CreateServiceLinkedRole` permission.
//
// - When you authorize a specific instance through this API, the system automatically adds the label `sysom:diagnosis` to the target ECS instance. SysOM can only diagnose instances that have this label.
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
// # Check whether the target instance is supported by SysOM
//
// Description:
//
// The instance list returned by this API includes only machines that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it will not appear in the list.
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
// # High-CPU agent streaming API
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
// # High-CPU agent streaming API
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
// This API is used to create an alert contact for push notifications.
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
// # Create an alert policy for push notifications
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
// 创建集群Vpc端点连接
//
// Description:
//
// - 需配合aliyun-tea-openapi-inner包的call_sseapi接口使用
//
// - 需要按通用LLM服务输入参数填充参数，转为string后赋给llmParamString
//
// - 返回数据需将string转为dict后使用，参考通用LLM服务返回格式
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
// 创建实例巡检
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
// This API creates an intelligent breakdown diagnosis task to diagnose the vmcore or dmesg log file provided in the parameters.
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
// This API is used to delete an alert contact.
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
// User deletes the alert policy for push notifications.
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
// # Query metrics
//
// Description:
//
// The instance list obtained by this API includes only the machines that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it will not appear in the list.
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
// # Obtain the Return Result of the copilot service
//
// Description:
//
// - You must fill in the input parameters according to the standard LLM service input parameters, convert them into a string, and assign the result to llmParamString.
//
// - The returned data must be converted from a string to a dict before use. Refer to the standard LLM service return format.
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
// # Stream Copilot service API
//
// Description:
//
// - Must be used together with the call_sseapi API of the aliyun-tea-openapi-inner package.
//
// - You must populate the input parameters according to the standard LLM service input parameters, convert them into a string, and assign the result to llmParamString.
//
// - The returned data is a string that you must convert into a dictionary for use, following the standard LLM service response format.
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
// # Stream Copilot service API
//
// Description:
//
// - Must be used together with the call_sseapi API of the aliyun-tea-openapi-inner package.
//
// - You must populate the input parameters according to the standard LLM service input parameters, convert them into a string, and assign the result to llmParamString.
//
// - The returned data is a string that you must convert into a dictionary for use, following the standard LLM service response format.
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
// # View AI Infra Analysis Result
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
// Obtain the quantity of unprocessed (undiagnosed) anomalous activity at different Levels for edge zones/pods.
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
// # Retrieve the details of a widget
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
// # Obtain the task execution status of Agent installation
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
// This API is used to obtain the specified alert contact information.
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
// Obtain an alert for a user by policy ID.
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
// # Retrieve copilot chat history
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
// Obtain the diagnosis result.
//
// Description:
//
// The diagnosis flow is asynchronous. Therefore, when you invoke this API, the diagnosis may still be executing and not yet ended. You can check the `data.status` field in the returned data to determine the status. When `data.status == "Success"`, it indicates that the diagnosis succeeded, and you can read the diagnosis result from `data.result`.
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
// Obtain the proportion of edge zone/pod health statuses over a period of time
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
// # Obtain the number of edge zones in a cluster or the number of pods in an edge zone
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
// Obtain the list of a specific field under an instance.
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
// # Obtain hot spot analysis results
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
// # Obtain hot spot comparison tracing results
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
// # Obtain the hot spot instance list
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
// # Obtain the PID list of a specific instance
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
// # Obtain hot spot tracing results
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
// 获取巡检报告
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
// Obtain real-time cluster/edge zone health degree score
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
// # AI Infra retrieves the list of analysis records
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
// Obtain the proportion of abnormal issues in pods within edge zones or in an edge zone within a cluster over a specified period of time.
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
// # Retrieve the health score trend
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
// # Obtain real-time resource usage of clusters or edge zones
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
// # Obtain Function Modules Configuration
//
// Description:
//
// This API is used to retrieve the service configuration status.
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
// This API queries the task execution status and diagnosis result based on the job ID.
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
// Initialize SysOM and ensure that the service role exists.
//
// Description:
//
// Some SysOM APIs require role assumption based on the `AliyunServiceRoleForSysom` service role. Therefore, before using SysOM features, you must invoke this API to perform initialization and ensure that the service role has been created.
//
// - `check_only`: If this parameter is set to True, the API only checks whether the service role exists and does not create it. If this parameter is set to False or omitted, the API automatically creates the service role if it does not exist.
//
// >
//
// > Note: When you invoke this API to initialize the role, you are deemed to have accepted the User Agreement of the operating system console by default. For more information, see [Overview of the Operating System Console](https://help.aliyun.com/zh/alinux/product-overview/os-console-overview?spm=a2c4g.11186623.help-menu-2632541.d_0_7.35a829ffLjQtgg) and [Alibaba Cloud Service Trial Terms](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud202001091714_51956.html).
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
// # Install an agent on the specified instance
//
// Description:
//
// The API call to install an agent is asynchronous. After invoking this API, a task_id is returned. You can use this ID to invoke the GetAgentTask API to retrieve the job execution status.
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
// # Install widgets on a cluster
//
// Description:
//
// After you install widgets on the specified ACK cluster:
//
// 1. When the cluster is first enrolled, widgets are installed on all ECS instances in the cluster (if the cluster contains more than 50 nodes, widgets are installed on only 50 nodes in the first batch).
//
// 2. The operating system console periodically checks for scale-in or scale-out events in the enrolled cluster. Whenever new ECS instances are added to the cluster, the operating system console automatically installs widgets on them without requiring user intervention.
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
// # Initiate diagnosis for anomalous activity
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
// Initiate a diagnosis.
//
// Description:
//
// Diagnosing the target ECS instance has the following requirements:
//
// - The instance status of the target ECS instance must be running.
//
// - The Cloud Assistant Agent must already be installed on the target ECS instance. If it is not installed, install it by referring to [Install the Cloud Assistant Agent](https://help.aliyun.com/zh/ecs/user-guide/install-the-cloud-assistant-agent).
//
// - You must invoke the AuthDiagnosis API to authorize SysOM to diagnose the target ECS instance. If this authorization is not granted, the API call will fail immediately.
//
// - This API depends on the existence of the SysOM service-linked role (AliyunServiceRoleForSysom). This API does not create the service-linked role automatically. If the service-linked role does not exist, you must first call AuthDiagnosis to perform authorization, which will create the aforementioned service-linked role.
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
// Obtain anomalous activity information for clusters, edge zones, or pods within a specified time period.
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
// # List installation records of the agent
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
// # Retrieve the Agent List
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
// This API is used to obtain the alert contact list.
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
// # Retrieve all alerting items
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
// # Used to obtain all alert policies for push notifications of a user
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
// This API is used to obtain a list of managed or unmanaged instances along with instance information.
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
// # Obtain cluster widget installation records
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
// # Retrieve all clusters managed by the current user
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
// Obtain the diagnosis history list.
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
// Obtain a list of cluster node or pod health scores within a specified time period.
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
// # Obtain instance status
//
// Description:
//
// This API is used to obtain the list of machines managed by SysOM.
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
// # Obtain the instance list
//
// Description:
//
// The instance list returned by this API includes only the machines that have been managed by SysOM. If an ECS instance exists but has not been managed by SysOM, it will not appear in the list.
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
// Obtain a list of ECS information, such as the tag list, public IP address list, and so on.
//
// Description:
//
// The instance list returned by this API includes only machines that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it will not appear in the list.
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
// Obtain information about managed or unmanaged instances, including ECS information.
//
// Description:
//
// The current API returns a list of instances that have already been managed by SysOM. If an ECS instance exists but has not been managed by SysOM, it will not appear in the list.
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
// # Obtain the list of instances for plugin installation, update, or uninstallation
//
// Description:
//
// The instance list returned by this API consists of machines that are already managed by SysOM. If an ECS instance exists but is not managed by SysOM, it will not appear in the list.
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
// # Retrieve the list of pods in a cluster or instance
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
// # List all areas where machines are managed
//
// Description:
//
// This API retrieves the list of areas where the current user has machines managed by SysOM. If the user has ECS instances in an area but those instances are not managed by SysOM, that area will not appear in the API response.
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
// Query the history list of breakdown diagnosis jobs.
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
// Start an AI job analysis.
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
// Start AI Infra differential analysis.
//
// Description:
//
// Currently, only comparative analysis between different steps under the same AI Infra analysis record and the same pid is supported.
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
// # Uninstall a specified version of the widget
//
// Description:
//
// The API call to uninstall an Agent is asynchronous. After invoking this API, a task_id is returned. You can use this ID to invoke the GetAgentTask API to retrieve the execution status of the job.
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
// # Uninstall a widget from a cluster
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
// This API is used to update an alert contact.
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
func (client *Client) UpdateAlertDestinationWithContext(ctx context.Context, request *UpdateAlertDestinationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlertDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
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
// # User updates the status of a push alert policy
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
// # Update push alert policy
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
// Update the follow level of an anomalous activity to adjust the sensitivity of the anomaly detection algorithm by modifying the follow level.
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
// Update the service function module configuration.
//
// Description:
//
// - You must fill in the parameters according to the input parameters of the general LLM service, convert them to a string, and assign the result to `llmParamString`.
//
// - To use the returned data, convert the string back to a dictionary, following the response format of the general LLM service.
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
// Update the version of the installed widget to the specified version.
//
// Description:
//
// The API call to update the Agent is asynchronous. After invoking this API, a task_id is returned. You can use this ID to invoke the GetAgentTask API to retrieve the execution status of the job.
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
// # Update widget for cluster
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
