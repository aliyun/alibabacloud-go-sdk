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
	client.SignatureAlgorithm = dara.String("v2")
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":  dara.String("cn-beijing.fnf.aliyuncs.com"),
		"cn-hangzhou": dara.String("cn-hangzhou.fnf.aliyuncs.com"),
		"cn-shanghai": dara.String("cn-shanghai.fnf.aliyuncs.com"),
		"cn-shenzhen": dara.String("cn-shenzhen.fnf.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("fnf"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// ## [](#)Usage notes
//
//   - The number of flows that each user can create is restricted by resources. For more information, see [Limits](https://help.aliyun.com/document_detail/122093.html). If you want to create more flows, submit a ticket.
//
//   - At the user level, flows are distinguished by name. The name of a flow within one account must be unique.
//
// @param tmpReq - CreateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowResponse
func (client *Client) CreateFlowWithOptions(tmpReq *CreateFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Environment) {
		request.EnvironmentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Environment, dara.String("Environment"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		body["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvironmentShrink) {
		body["Environment"] = request.EnvironmentShrink
	}

	if !dara.IsNil(request.ExecutionMode) {
		body["ExecutionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.ExternalStorageLocation) {
		body["ExternalStorageLocation"] = request.ExternalStorageLocation
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlow"),
		Version:     dara.String("2019-03-15"),
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
// ## [](#)Usage notes
//
//   - The number of flows that each user can create is restricted by resources. For more information, see [Limits](https://help.aliyun.com/document_detail/122093.html). If you want to create more flows, submit a ticket.
//
//   - At the user level, flows are distinguished by name. The name of a flow within one account must be unique.
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
// 创建流程版本别名
//
// @param tmpReq - CreateFlowAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowAliasResponse
func (client *Client) CreateFlowAliasWithOptions(tmpReq *CreateFlowAliasRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFlowAliasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoutingConfigurations) {
		request.RoutingConfigurationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoutingConfigurations, dara.String("RoutingConfigurations"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RoutingConfigurationsShrink) {
		body["RoutingConfigurations"] = request.RoutingConfigurationsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowAlias"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流程版本别名
//
// @param request - CreateFlowAliasRequest
//
// @return CreateFlowAliasResponse
func (client *Client) CreateFlowAlias(request *CreateFlowAliasRequest) (_result *CreateFlowAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowAliasResponse{}
	_body, _err := client.CreateFlowAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a time-based schedule.
//
// @param request - CreateScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduleResponse
func (client *Client) CreateScheduleWithOptions(request *CreateScheduleRequest, runtime *dara.RuntimeOptions) (_result *CreateScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SignatureVersion) {
		query["SignatureVersion"] = request.SignatureVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CronExpression) {
		body["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enable) {
		body["Enable"] = request.Enable
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.Payload) {
		body["Payload"] = request.Payload
	}

	if !dara.IsNil(request.ScheduleName) {
		body["ScheduleName"] = request.ScheduleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSchedule"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a time-based schedule.
//
// @param request - CreateScheduleRequest
//
// @return CreateScheduleResponse
func (client *Client) CreateSchedule(request *CreateScheduleRequest) (_result *CreateScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScheduleResponse{}
	_body, _err := client.CreateScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an existing flow.
//
// Description:
//
// ## [](#)Usage notes
//
// A delete operation is asynchronous. If this operation is successful, the system returns a successful response. If an existing flow is pending to be deleted, a new flow of the same name will not be affected by the existing one. After you delete a flow, you cannot query its historical executions. All executions in progress will stop after their most recent steps are complete.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlow"),
		Version:     dara.String("2019-03-15"),
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
// Deletes an existing flow.
//
// Description:
//
// ## [](#)Usage notes
//
// A delete operation is asynchronous. If this operation is successful, the system returns a successful response. If an existing flow is pending to be deleted, a new flow of the same name will not be affected by the existing one. After you delete a flow, you cannot query its historical executions. All executions in progress will stop after their most recent steps are complete.
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
// 删除流程别名
//
// @param request - DeleteFlowAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowAliasResponse
func (client *Client) DeleteFlowAliasWithOptions(request *DeleteFlowAliasRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowAlias"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流程别名
//
// @param request - DeleteFlowAliasRequest
//
// @return DeleteFlowAliasResponse
func (client *Client) DeleteFlowAlias(request *DeleteFlowAliasRequest) (_result *DeleteFlowAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowAliasResponse{}
	_body, _err := client.DeleteFlowAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除流程版本
//
// @param request - DeleteFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowVersionResponse
func (client *Client) DeleteFlowVersionWithOptions(request *DeleteFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowVersion) {
		body["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowVersion"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流程版本
//
// @param request - DeleteFlowVersionRequest
//
// @return DeleteFlowVersionResponse
func (client *Client) DeleteFlowVersion(request *DeleteFlowVersionRequest) (_result *DeleteFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowVersionResponse{}
	_body, _err := client.DeleteFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a time-based scheduling task.
//
// @param request - DeleteScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduleResponse
func (client *Client) DeleteScheduleWithOptions(request *DeleteScheduleRequest, runtime *dara.RuntimeOptions) (_result *DeleteScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.ScheduleName) {
		body["ScheduleName"] = request.ScheduleName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSchedule"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a time-based scheduling task.
//
// @param request - DeleteScheduleRequest
//
// @return DeleteScheduleResponse
func (client *Client) DeleteSchedule(request *DeleteScheduleRequest) (_result *DeleteScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScheduleResponse{}
	_body, _err := client.DeleteScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an execution in a flow. The long polling mode is supported. The maximum waiting period for long polling depends on the value of the WaitTimeSeconds parameter.
//
// @param request - DescribeExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExecutionResponse
func (client *Client) DescribeExecutionWithOptions(request *DescribeExecutionRequest, runtime *dara.RuntimeOptions) (_result *DescribeExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExecution"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an execution in a flow. The long polling mode is supported. The maximum waiting period for long polling depends on the value of the WaitTimeSeconds parameter.
//
// @param request - DescribeExecutionRequest
//
// @return DescribeExecutionResponse
func (client *Client) DescribeExecution(request *DescribeExecutionRequest) (_result *DescribeExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExecutionResponse{}
	_body, _err := client.DescribeExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a flow.
//
// @param request - DescribeFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowResponse
func (client *Client) DescribeFlowWithOptions(request *DescribeFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlow"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a flow.
//
// @param request - DescribeFlowRequest
//
// @return DescribeFlowResponse
func (client *Client) DescribeFlow(request *DescribeFlowRequest) (_result *DescribeFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowResponse{}
	_body, _err := client.DescribeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流程版本别名详情
//
// @param request - DescribeFlowAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowAliasResponse
func (client *Client) DescribeFlowAliasWithOptions(request *DescribeFlowAliasRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowAlias"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流程版本别名详情
//
// @param request - DescribeFlowAliasRequest
//
// @return DescribeFlowAliasResponse
func (client *Client) DescribeFlowAlias(request *DescribeFlowAliasRequest) (_result *DescribeFlowAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowAliasResponse{}
	_body, _err := client.DescribeFlowAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 MapRun 详情
//
// @param request - DescribeMapRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMapRunResponse
func (client *Client) DescribeMapRunWithOptions(request *DescribeMapRunRequest, runtime *dara.RuntimeOptions) (_result *DescribeMapRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMapRun"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMapRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 MapRun 详情
//
// @param request - DescribeMapRunRequest
//
// @return DescribeMapRunResponse
func (client *Client) DescribeMapRun(request *DescribeMapRunRequest) (_result *DescribeMapRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMapRunResponse{}
	_body, _err := client.DescribeMapRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions where CloudFlow is available.
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
		Version:     dara.String("2019-03-15"),
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
// Queries the regions where CloudFlow is available.
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
// Queries the detailed information about a time-based schedule.
//
// @param request - DescribeScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScheduleResponse
func (client *Client) DescribeScheduleWithOptions(request *DescribeScheduleRequest, runtime *dara.RuntimeOptions) (_result *DescribeScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSchedule"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a time-based schedule.
//
// @param request - DescribeScheduleRequest
//
// @return DescribeScheduleResponse
func (client *Client) DescribeSchedule(request *DescribeScheduleRequest) (_result *DescribeScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScheduleResponse{}
	_body, _err := client.DescribeScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about each step in an execution process.
//
// @param request - GetExecutionHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecutionHistoryResponse
func (client *Client) GetExecutionHistoryWithOptions(request *GetExecutionHistoryRequest, runtime *dara.RuntimeOptions) (_result *GetExecutionHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExecutionHistory"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExecutionHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about each step in an execution process.
//
// @param request - GetExecutionHistoryRequest
//
// @return GetExecutionHistoryResponse
func (client *Client) GetExecutionHistory(request *GetExecutionHistoryRequest) (_result *GetExecutionHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExecutionHistoryResponse{}
	_body, _err := client.GetExecutionHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all historical executions of a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// After you delete a flow, you cannot query its historical executions, even if you create a flow of the same name.
//
// @param request - ListExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionsResponse
func (client *Client) ListExecutionsWithOptions(request *ListExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExecutions"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all historical executions of a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// After you delete a flow, you cannot query its historical executions, even if you create a flow of the same name.
//
// @param request - ListExecutionsRequest
//
// @return ListExecutionsResponse
func (client *Client) ListExecutions(request *ListExecutionsRequest) (_result *ListExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExecutionsResponse{}
	_body, _err := client.ListExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流程版本别名列表
//
// @param request - ListFlowAliasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowAliasesResponse
func (client *Client) ListFlowAliasesWithOptions(request *ListFlowAliasesRequest, runtime *dara.RuntimeOptions) (_result *ListFlowAliasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlowAliases"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowAliasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流程版本别名列表
//
// @param request - ListFlowAliasesRequest
//
// @return ListFlowAliasesResponse
func (client *Client) ListFlowAliases(request *ListFlowAliasesRequest) (_result *ListFlowAliasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlowAliasesResponse{}
	_body, _err := client.ListFlowAliasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流程版本列表
//
// @param request - ListFlowVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowVersionsResponse
func (client *Client) ListFlowVersionsWithOptions(request *ListFlowVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListFlowVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlowVersions"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流程版本列表
//
// @param request - ListFlowVersionsRequest
//
// @return ListFlowVersionsResponse
func (client *Client) ListFlowVersions(request *ListFlowVersionsRequest) (_result *ListFlowVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlowVersionsResponse{}
	_body, _err := client.ListFlowVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of flows.
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
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlows"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Queries a list of flows.
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
// Queries time-based schedules in a flow.
//
// @param request - ListSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchedulesResponse
func (client *Client) ListSchedulesWithOptions(request *ListSchedulesRequest, runtime *dara.RuntimeOptions) (_result *ListSchedulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchedules"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries time-based schedules in a flow.
//
// @param request - ListSchedulesRequest
//
// @return ListSchedulesResponse
func (client *Client) ListSchedules(request *ListSchedulesRequest) (_result *ListSchedulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSchedulesResponse{}
	_body, _err := client.ListSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布流程版本
//
// @param request - PublishFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFlowVersionResponse
func (client *Client) PublishFlowVersionWithOptions(request *PublishFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *PublishFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishFlowVersion"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishFlowVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布流程版本
//
// @param request - PublishFlowVersionRequest
//
// @return PublishFlowVersionResponse
func (client *Client) PublishFlowVersion(request *PublishFlowVersionRequest) (_result *PublishFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishFlowVersionResponse{}
	_body, _err := client.PublishFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reports a failed task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskFailed is used to call back `pattern: waitForCallback` indicates that the current task fails to be executed.
//
// In the new service (CloudFlow), the task step that ReportTaskFailed is used to call back `TaskMode: WaitForCustomCallback` indicates that the current task fails to be executed.
//
// @param request - ReportTaskFailedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportTaskFailedResponse
func (client *Client) ReportTaskFailedWithOptions(request *ReportTaskFailedRequest, runtime *dara.RuntimeOptions) (_result *ReportTaskFailedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskToken) {
		query["TaskToken"] = request.TaskToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Cause) {
		body["Cause"] = request.Cause
	}

	if !dara.IsNil(request.Error) {
		body["Error"] = request.Error
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportTaskFailed"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportTaskFailedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports a failed task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskFailed is used to call back `pattern: waitForCallback` indicates that the current task fails to be executed.
//
// In the new service (CloudFlow), the task step that ReportTaskFailed is used to call back `TaskMode: WaitForCustomCallback` indicates that the current task fails to be executed.
//
// @param request - ReportTaskFailedRequest
//
// @return ReportTaskFailedResponse
func (client *Client) ReportTaskFailed(request *ReportTaskFailedRequest) (_result *ReportTaskFailedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportTaskFailedResponse{}
	_body, _err := client.ReportTaskFailedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reports a successful task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskSucceeded is used to call back pattern: waitForCallback indicates that the current task is successfully executed.
//
// In the new service (CloudFlow), the task step that ReportTaskSucceeded is used to call back TaskMode: WaitForCustomCallback indicates that the current task is successfully executed.
//
// @param request - ReportTaskSucceededRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportTaskSucceededResponse
func (client *Client) ReportTaskSucceededWithOptions(request *ReportTaskSucceededRequest, runtime *dara.RuntimeOptions) (_result *ReportTaskSucceededResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskToken) {
		query["TaskToken"] = request.TaskToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Output) {
		body["Output"] = request.Output
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportTaskSucceeded"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportTaskSucceededResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports a successful task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskSucceeded is used to call back pattern: waitForCallback indicates that the current task is successfully executed.
//
// In the new service (CloudFlow), the task step that ReportTaskSucceeded is used to call back TaskMode: WaitForCustomCallback indicates that the current task is successfully executed.
//
// @param request - ReportTaskSucceededRequest
//
// @return ReportTaskSucceededResponse
func (client *Client) ReportTaskSucceeded(request *ReportTaskSucceededRequest) (_result *ReportTaskSucceededResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportTaskSucceededResponse{}
	_body, _err := client.ReportTaskSucceededWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts the execution of a workflow.
//
// Description:
//
// ## [](#)Usage notes
//
//   - The flow is created. A flow only in standard mode is supported.
//
//   - If you do not specify an execution, the system automatically generates an execution and starts the execution.
//
//   - If an ongoing execution has the same name as that of the execution to be started, the system directly returns the ongoing execution.
//
//   - If the ongoing execution with the same name has ended (succeeded or failed), `ExecutionAlreadyExists` is returned.
//
//   - If no execution with the same name exists, the system starts a new execution.
//
// @param request - StartExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartExecutionResponse
func (client *Client) StartExecutionWithOptions(request *StartExecutionRequest, runtime *dara.RuntimeOptions) (_result *StartExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackFnFTaskToken) {
		body["CallbackFnFTaskToken"] = request.CallbackFnFTaskToken
	}

	if !dara.IsNil(request.ExecutionName) {
		body["ExecutionName"] = request.ExecutionName
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.Input) {
		body["Input"] = request.Input
	}

	if !dara.IsNil(request.Qualifier) {
		body["Qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartExecution"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the execution of a workflow.
//
// Description:
//
// ## [](#)Usage notes
//
//   - The flow is created. A flow only in standard mode is supported.
//
//   - If you do not specify an execution, the system automatically generates an execution and starts the execution.
//
//   - If an ongoing execution has the same name as that of the execution to be started, the system directly returns the ongoing execution.
//
//   - If the ongoing execution with the same name has ended (succeeded or failed), `ExecutionAlreadyExists` is returned.
//
//   - If no execution with the same name exists, the system starts a new execution.
//
// @param request - StartExecutionRequest
//
// @return StartExecutionResponse
func (client *Client) StartExecution(request *StartExecutionRequest) (_result *StartExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartExecutionResponse{}
	_body, _err := client.StartExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronously starts an execution in a flow.
//
// Description:
//
//	Only flows of the express execution mode are supported.
//
// @param request - StartSyncExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSyncExecutionResponse
func (client *Client) StartSyncExecutionWithOptions(request *StartSyncExecutionRequest, runtime *dara.RuntimeOptions) (_result *StartSyncExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionName) {
		body["ExecutionName"] = request.ExecutionName
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.Input) {
		body["Input"] = request.Input
	}

	if !dara.IsNil(request.Qualifier) {
		body["Qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSyncExecution"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSyncExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronously starts an execution in a flow.
//
// Description:
//
//	Only flows of the express execution mode are supported.
//
// @param request - StartSyncExecutionRequest
//
// @return StartSyncExecutionResponse
func (client *Client) StartSyncExecution(request *StartSyncExecutionRequest) (_result *StartSyncExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartSyncExecutionResponse{}
	_body, _err := client.StartSyncExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an execution that is in progress in a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// The flow must be in progress.
//
// @param request - StopExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopExecutionResponse
func (client *Client) StopExecutionWithOptions(request *StopExecutionRequest, runtime *dara.RuntimeOptions) (_result *StopExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cause) {
		body["Cause"] = request.Cause
	}

	if !dara.IsNil(request.Error) {
		body["Error"] = request.Error
	}

	if !dara.IsNil(request.ExecutionName) {
		body["ExecutionName"] = request.ExecutionName
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopExecution"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an execution that is in progress in a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// The flow must be in progress.
//
// @param request - StopExecutionRequest
//
// @return StopExecutionResponse
func (client *Client) StopExecution(request *StopExecutionRequest) (_result *StopExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopExecutionResponse{}
	_body, _err := client.StopExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a flow.
//
// @param tmpReq - UpdateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowResponse
func (client *Client) UpdateFlowWithOptions(tmpReq *UpdateFlowRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Environment) {
		request.EnvironmentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Environment, dara.String("Environment"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		body["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvironmentShrink) {
		body["Environment"] = request.EnvironmentShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlow"),
		Version:     dara.String("2019-03-15"),
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
// Updates a flow.
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
// 更新流程版本别名配置
//
// @param tmpReq - UpdateFlowAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowAliasResponse
func (client *Client) UpdateFlowAliasWithOptions(tmpReq *UpdateFlowAliasRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFlowAliasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoutingConfigurations) {
		request.RoutingConfigurationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoutingConfigurations, dara.String("RoutingConfigurations"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RoutingConfigurationsShrink) {
		body["RoutingConfigurations"] = request.RoutingConfigurationsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlowAlias"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流程版本别名配置
//
// @param request - UpdateFlowAliasRequest
//
// @return UpdateFlowAliasResponse
func (client *Client) UpdateFlowAlias(request *UpdateFlowAliasRequest) (_result *UpdateFlowAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFlowAliasResponse{}
	_body, _err := client.UpdateFlowAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 MapRun 配置
//
// @param request - UpdateMapRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMapRunResponse
func (client *Client) UpdateMapRunWithOptions(request *UpdateMapRunRequest, runtime *dara.RuntimeOptions) (_result *UpdateMapRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMapRun"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMapRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 MapRun 配置
//
// @param request - UpdateMapRunRequest
//
// @return UpdateMapRunResponse
func (client *Client) UpdateMapRun(request *UpdateMapRunRequest) (_result *UpdateMapRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMapRunResponse{}
	_body, _err := client.UpdateMapRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a time-based schedule.
//
// @param request - UpdateScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduleResponse
func (client *Client) UpdateScheduleWithOptions(request *UpdateScheduleRequest, runtime *dara.RuntimeOptions) (_result *UpdateScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CronExpression) {
		body["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enable) {
		body["Enable"] = request.Enable
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.Payload) {
		body["Payload"] = request.Payload
	}

	if !dara.IsNil(request.ScheduleName) {
		body["ScheduleName"] = request.ScheduleName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSchedule"),
		Version:     dara.String("2019-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a time-based schedule.
//
// @param request - UpdateScheduleRequest
//
// @return UpdateScheduleResponse
func (client *Client) UpdateSchedule(request *UpdateScheduleRequest) (_result *UpdateScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScheduleResponse{}
	_body, _err := client.UpdateScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
