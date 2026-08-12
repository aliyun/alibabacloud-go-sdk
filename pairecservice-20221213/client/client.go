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
		"cn-shenzhen":    dara.String("pairecservice.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai":    dara.String("pairecservice.cn-shanghai.aliyuncs.com"),
		"cn-hongkong":    dara.String("pairecservice.cn-hongkong.aliyuncs.com"),
		"cn-hangzhou":    dara.String("pairecservice.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":     dara.String("pairecservice.cn-beijing.aliyuncs.com"),
		"ap-southeast-5": dara.String("pairecservice.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-1": dara.String("pairecservice.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":   dara.String("pairecservice.eu-central-1.aliyuncs.com"),
		"us-east-1":      dara.String("pairecservice.us-east-1.aliyuncs.com"),
		"us-west-1":      dara.String("pairecservice.us-west-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("pairecservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Applies an engine configuration.
//
// @param request - ApplyEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyEngineConfigResponse
func (client *Client) ApplyEngineConfigWithOptions(EngineConfigId *string, request *ApplyEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApplyEngineConfigResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId)) + "/action/apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyEngineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies an engine configuration.
//
// @param request - ApplyEngineConfigRequest
//
// @return ApplyEngineConfigResponse
func (client *Client) ApplyEngineConfig(EngineConfigId *string, request *ApplyEngineConfigRequest) (_result *ApplyEngineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyEngineConfigResponse{}
	_body, _err := client.ApplyEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs data backflow for the latest job of a specified feature consistency check job configuration.
//
// @param request - BackflowFeatureConsistencyCheckJobDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackflowFeatureConsistencyCheckJobDataResponse
func (client *Client) BackflowFeatureConsistencyCheckJobDataWithOptions(request *BackflowFeatureConsistencyCheckJobDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BackflowFeatureConsistencyCheckJobDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FeatureConsistencyCheckJobConfigId) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemFeatures) {
		body["ItemFeatures"] = request.ItemFeatures
	}

	if !dara.IsNil(request.LogItemId) {
		body["LogItemId"] = request.LogItemId
	}

	if !dara.IsNil(request.LogRequestId) {
		body["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.LogRequestTime) {
		body["LogRequestTime"] = request.LogRequestTime
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.SceneName) {
		body["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.Scores) {
		body["Scores"] = request.Scores
	}

	if !dara.IsNil(request.ServiceName) {
		body["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.UserFeatures) {
		body["UserFeatures"] = request.UserFeatures
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BackflowFeatureConsistencyCheckJobData"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/action/backflowdata"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BackflowFeatureConsistencyCheckJobDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs data backflow for the latest job of a specified feature consistency check job configuration.
//
// @param request - BackflowFeatureConsistencyCheckJobDataRequest
//
// @return BackflowFeatureConsistencyCheckJobDataResponse
func (client *Client) BackflowFeatureConsistencyCheckJobData(request *BackflowFeatureConsistencyCheckJobDataRequest) (_result *BackflowFeatureConsistencyCheckJobDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BackflowFeatureConsistencyCheckJobDataResponse{}
	_body, _err := client.BackflowFeatureConsistencyCheckJobDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API changes the version of a recall management service.
//
// Description:
//
// ## Request
//
// Use this API to change the version of a recall management service. Ensure that the provided `RecallManagementServiceId`, `InstanceId`, and target `RecallManagementServiceVersionId` are valid, and that you have the required permissions for these resources.
//
// - **RecallManagementServiceId**: The unique identifier of the recall management service.
//
// - **InstanceId**: The instance ID associated with the recall management service.
//
// - **RecallManagementServiceVersionId**: The target version ID to switch to.
//
// Note: Before changing the version, confirm that the new version is fully tested and ready for production.
//
// @param request - ChangeRecallManagementServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeRecallManagementServiceVersionResponse
func (client *Client) ChangeRecallManagementServiceVersionWithOptions(RecallManagementServiceId *string, request *ChangeRecallManagementServiceVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeRecallManagementServiceVersionResponse, _err error) {
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

	if !dara.IsNil(request.RecallManagementServiceVersionId) {
		body["RecallManagementServiceVersionId"] = request.RecallManagementServiceVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeRecallManagementServiceVersion"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/action/changeversion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeRecallManagementServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API changes the version of a recall management service.
//
// Description:
//
// ## Request
//
// Use this API to change the version of a recall management service. Ensure that the provided `RecallManagementServiceId`, `InstanceId`, and target `RecallManagementServiceVersionId` are valid, and that you have the required permissions for these resources.
//
// - **RecallManagementServiceId**: The unique identifier of the recall management service.
//
// - **InstanceId**: The instance ID associated with the recall management service.
//
// - **RecallManagementServiceVersionId**: The target version ID to switch to.
//
// Note: Before changing the version, confirm that the new version is fully tested and ready for production.
//
// @param request - ChangeRecallManagementServiceVersionRequest
//
// @return ChangeRecallManagementServiceVersionResponse
func (client *Client) ChangeRecallManagementServiceVersion(RecallManagementServiceId *string, request *ChangeRecallManagementServiceVersionRequest) (_result *ChangeRecallManagementServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeRecallManagementServiceVersionResponse{}
	_body, _err := client.ChangeRecallManagementServiceVersionWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a conversation message to an agent. Supports Server-Sent Events (SSE). Creates a new session or continues a conversation in an existing session.
//
// Description:
//
// ## Operation description
//
// - Call this API operation to send a conversation message to an agent. Server-Sent Events (SSE) is supported.
//
// - If the `ConversationId` parameter is specified, the conversation continues in the context of the specified existing session. If this parameter is not specified, automatic creation of a new session is performed.
//
// - The `Config` field allows you to pass additional information input. The value must be in JSON format.
//
// - If the request succeeds, the response includes the message ID, reply content, and other information for this conversation. If a fault occurs, the corresponding error code and error message are returned.
//
// @param request - ChatConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatConversationResponse
func (client *Client) ChatConversationWithSSE(request *ChatConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ChatConversationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatConversationWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// Sends a conversation message to an agent. Supports Server-Sent Events (SSE). Creates a new session or continues a conversation in an existing session.
//
// Description:
//
// ## Operation description
//
// - Call this API operation to send a conversation message to an agent. Server-Sent Events (SSE) is supported.
//
// - If the `ConversationId` parameter is specified, the conversation continues in the context of the specified existing session. If this parameter is not specified, automatic creation of a new session is performed.
//
// - The `Config` field allows you to pass additional information input. The value must be in JSON format.
//
// - If the request succeeds, the response includes the message ID, reply content, and other information for this conversation. If a fault occurs, the corresponding error code and error message are returned.
//
// @param request - ChatConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatConversationResponse
func (client *Client) ChatConversationWithOptions(request *ChatConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChatConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatConversation"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/conversations/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a conversation message to an agent. Supports Server-Sent Events (SSE). Creates a new session or continues a conversation in an existing session.
//
// Description:
//
// ## Operation description
//
// - Call this API operation to send a conversation message to an agent. Server-Sent Events (SSE) is supported.
//
// - If the `ConversationId` parameter is specified, the conversation continues in the context of the specified existing session. If this parameter is not specified, automatic creation of a new session is performed.
//
// - The `Config` field allows you to pass additional information input. The value must be in JSON format.
//
// - If the request succeeds, the response includes the message ID, reply content, and other information for this conversation. If a fault occurs, the corresponding error code and error message are returned.
//
// @param request - ChatConversationRequest
//
// @return ChatConversationResponse
func (client *Client) ChatConversation(request *ChatConversationRequest) (_result *ChatConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChatConversationResponse{}
	_body, _err := client.ChatConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies access to resources configured for an instance.
//
// @param request - CheckInstanceResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceResourcesResponse
func (client *Client) CheckInstanceResourcesWithOptions(InstanceId *string, request *CheckInstanceResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInstanceResources"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/action/checkresources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies access to resources configured for an instance.
//
// @param request - CheckInstanceResourcesRequest
//
// @return CheckInstanceResourcesResponse
func (client *Client) CheckInstanceResources(InstanceId *string, request *CheckInstanceResourcesRequest) (_result *CheckInstanceResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstanceResourcesResponse{}
	_body, _err := client.CheckInstanceResourcesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates a traffic control task expression.
//
// Description:
//
// ## Description
//
// This operation validates a traffic control task expression for a specific instance and table. Provide the correct `InstanceId`, `TableMetaId`, and `Expression` parameters.
//
// @param request - CheckTrafficControlTaskExpressionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTrafficControlTaskExpressionResponse
func (client *Client) CheckTrafficControlTaskExpressionWithOptions(request *CheckTrafficControlTaskExpressionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckTrafficControlTaskExpressionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TableMetaId) {
		query["TableMetaId"] = request.TableMetaId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckTrafficControlTaskExpression"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/action/checkexpression"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckTrafficControlTaskExpressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates a traffic control task expression.
//
// Description:
//
// ## Description
//
// This operation validates a traffic control task expression for a specific instance and table. Provide the correct `InstanceId`, `TableMetaId`, and `Expression` parameters.
//
// @param request - CheckTrafficControlTaskExpressionRequest
//
// @return CheckTrafficControlTaskExpressionResponse
func (client *Client) CheckTrafficControlTaskExpression(request *CheckTrafficControlTaskExpressionRequest) (_result *CheckTrafficControlTaskExpressionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckTrafficControlTaskExpressionResponse{}
	_body, _err := client.CheckTrafficControlTaskExpressionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones an engine configuration.
//
// @param request - CloneEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneEngineConfigResponse
func (client *Client) CloneEngineConfigWithOptions(EngineConfigId *string, request *CloneEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneEngineConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneEngineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones an engine configuration.
//
// @param request - CloneEngineConfigRequest
//
// @return CloneEngineConfigResponse
func (client *Client) CloneEngineConfig(EngineConfigId *string, request *CloneEngineConfigRequest) (_result *CloneEngineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneEngineConfigResponse{}
	_body, _err := client.CloneEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones a specified experiment.
//
// @param request - CloneExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneExperimentResponse
func (client *Client) CloneExperimentWithOptions(ExperimentId *string, request *CloneExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneExperimentResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones a specified experiment.
//
// @param request - CloneExperimentRequest
//
// @return CloneExperimentResponse
func (client *Client) CloneExperiment(ExperimentId *string, request *CloneExperimentRequest) (_result *CloneExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneExperimentResponse{}
	_body, _err := client.CloneExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones an experiment group to a specified environment.
//
// @param request - CloneExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneExperimentGroupResponse
func (client *Client) CloneExperimentGroupWithOptions(ExperimentGroupId *string, request *CloneExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneExperimentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LayerId) {
		body["LayerId"] = request.LayerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones an experiment group to a specified environment.
//
// @param request - CloneExperimentGroupRequest
//
// @return CloneExperimentGroupResponse
func (client *Client) CloneExperimentGroup(ExperimentGroupId *string, request *CloneExperimentGroupRequest) (_result *CloneExperimentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneExperimentGroupResponse{}
	_body, _err := client.CloneExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones a feature consistency check job configuration.
//
// @param request - CloneFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneFeatureConsistencyCheckJobConfigResponse
func (client *Client) CloneFeatureConsistencyCheckJobConfigWithOptions(SourceFeatureConsistencyCheckJobConfigId *string, request *CloneFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneFeatureConsistencyCheckJobConfigResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs/" + dara.PercentEncode(dara.StringValue(SourceFeatureConsistencyCheckJobConfigId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones a feature consistency check job configuration.
//
// @param request - CloneFeatureConsistencyCheckJobConfigRequest
//
// @return CloneFeatureConsistencyCheckJobConfigResponse
func (client *Client) CloneFeatureConsistencyCheckJobConfig(SourceFeatureConsistencyCheckJobConfigId *string, request *CloneFeatureConsistencyCheckJobConfigRequest) (_result *CloneFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CloneFeatureConsistencyCheckJobConfigWithOptions(SourceFeatureConsistencyCheckJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones a specified laboratory to a designated environment. You can specify whether to clone the experiment groups within the laboratory.
//
// @param request - CloneLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneLaboratoryResponse
func (client *Client) CloneLaboratoryWithOptions(LaboratoryId *string, request *CloneLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneLaboratoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloneExperimentGroup) {
		body["CloneExperimentGroup"] = request.CloneExperimentGroup
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones a specified laboratory to a designated environment. You can specify whether to clone the experiment groups within the laboratory.
//
// @param request - CloneLaboratoryRequest
//
// @return CloneLaboratoryResponse
func (client *Client) CloneLaboratory(LaboratoryId *string, request *CloneLaboratoryRequest) (_result *CloneLaboratoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneLaboratoryResponse{}
	_body, _err := client.CloneLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones a specified traffic control task to a new instance.
//
// Description:
//
// ## Request description
//
// This API clones an existing traffic control task to another specified instance. Ensure that the `InstanceId` you provide is valid and that you have the required permissions for the target instance.
//
// Note: The cloning process does not affect the status or configuration of the original task.
//
// @param request - CloneTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneTrafficControlTaskResponse
func (client *Client) CloneTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *CloneTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneTrafficControlTaskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones a specified traffic control task to a new instance.
//
// Description:
//
// ## Request description
//
// This API clones an existing traffic control task to another specified instance. Ensure that the `InstanceId` you provide is valid and that you have the required permissions for the target instance.
//
// Note: The cloning process does not affect the status or configuration of the original task.
//
// @param request - CloneTrafficControlTaskRequest
//
// @return CloneTrafficControlTaskResponse
func (client *Client) CloneTrafficControlTask(TrafficControlTaskId *string, request *CloneTrafficControlTaskRequest) (_result *CloneTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneTrafficControlTaskResponse{}
	_body, _err := client.CloneTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对比样本一致性任务
//
// @param request - CompareSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareSampleConsistencyJobResponse
func (client *Client) CompareSampleConsistencyJobWithOptions(SampleConsistencyJobId *string, request *CompareSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CompareSampleConsistencyJobResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/compare"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareSampleConsistencyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对比样本一致性任务
//
// @param request - CompareSampleConsistencyJobRequest
//
// @return CompareSampleConsistencyJobResponse
func (client *Client) CompareSampleConsistencyJob(SampleConsistencyJobId *string, request *CompareSampleConsistencyJobRequest) (_result *CompareSampleConsistencyJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CompareSampleConsistencyJobResponse{}
	_body, _err := client.CompareSampleConsistencyJobWithOptions(SampleConsistencyJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AB metric.
//
// @param request - CreateABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABMetricResponse
func (client *Client) CreateABMetricWithOptions(request *CreateABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateABMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregationByUser) {
		body["AggregationByUser"] = request.AggregationByUser
	}

	if !dara.IsNil(request.Definition) {
		body["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Denominator) {
		body["Denominator"] = request.Denominator
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsBinomialDistribution) {
		body["IsBinomialDistribution"] = request.IsBinomialDistribution
	}

	if !dara.IsNil(request.LeftMetricId) {
		body["LeftMetricId"] = request.LeftMetricId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedSignificance) {
		body["NeedSignificance"] = request.NeedSignificance
	}

	if !dara.IsNil(request.Numerator) {
		body["Numerator"] = request.Numerator
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.ResultResourceId) {
		body["ResultResourceId"] = request.ResultResourceId
	}

	if !dara.IsNil(request.RightMetricId) {
		body["RightMetricId"] = request.RightMetricId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StatisticsCycle) {
		body["StatisticsCycle"] = request.StatisticsCycle
	}

	if !dara.IsNil(request.TableMetaId) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AB metric.
//
// @param request - CreateABMetricRequest
//
// @return CreateABMetricResponse
func (client *Client) CreateABMetric(request *CreateABMetricRequest) (_result *CreateABMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABMetricResponse{}
	_body, _err := client.CreateABMetricWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an A/B metric group.
//
// @param request - CreateABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABMetricGroupResponse
func (client *Client) CreateABMetricGroupWithOptions(request *CreateABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateABMetricGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ABMetricIds) {
		body["ABMetricIds"] = request.ABMetricIds
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

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an A/B metric group.
//
// @param request - CreateABMetricGroupRequest
//
// @return CreateABMetricGroupResponse
func (client *Client) CreateABMetricGroup(request *CreateABMetricGroupRequest) (_result *CreateABMetricGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABMetricGroupResponse{}
	_body, _err := client.CreateABMetricGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates multiple calculation jobs.
//
// @param request - CreateCalculationJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCalculationJobsResponse
func (client *Client) CreateCalculationJobsWithOptions(request *CreateCalculationJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCalculationJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ABMetricIds) {
		body["ABMetricIds"] = request.ABMetricIds
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCalculationJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/batch/calculationjobs/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCalculationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple calculation jobs.
//
// @param request - CreateCalculationJobsRequest
//
// @return CreateCalculationJobsResponse
func (client *Client) CreateCalculationJobs(request *CreateCalculationJobsRequest) (_result *CreateCalculationJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCalculationJobsResponse{}
	_body, _err := client.CreateCalculationJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a crowd that represents a group of users.
//
// @param request - CreateCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCrowdResponse
func (client *Client) CreateCrowdWithOptions(request *CreateCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCrowdResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a crowd that represents a group of users.
//
// @param request - CreateCrowdRequest
//
// @return CreateCrowdResponse
func (client *Client) CreateCrowd(request *CreateCrowdRequest) (_result *CreateCrowdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCrowdResponse{}
	_body, _err := client.CreateCrowdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data diagnosis task. This API supports various analysis types.
//
// Description:
//
// ## Description
//
// - This API creates a data diagnosis task. It supports various analysis types, including item or user change rate analysis, user preference statistics cycle analysis, two-table join analysis, basic statistical analysis, and abnormal behavior analysis.
//
// - The content of the `Config` parameter depends on the value of the `Type` parameter. For more information, see the example configurations in this topic.
//
// - To run the task on a schedule, specify the `CycleTime` parameter. If this parameter is omitted, the task runs only once.
//
// - The optional `TopNQuantity` parameter specifies the number of top results to return.
//
// @param request - CreateDataDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataDiagnosisResponse
func (client *Client) CreateDataDiagnosisWithOptions(request *CreateDataDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.CycleTime) {
		body["CycleTime"] = request.CycleTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LeftTableMetaId) {
		body["LeftTableMetaId"] = request.LeftTableMetaId
	}

	if !dara.IsNil(request.LeftTablePartitionField) {
		body["LeftTablePartitionField"] = request.LeftTablePartitionField
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PartitionField) {
		body["PartitionField"] = request.PartitionField
	}

	if !dara.IsNil(request.RightTableMetaId) {
		body["RightTableMetaId"] = request.RightTableMetaId
	}

	if !dara.IsNil(request.RightTablePartitionField) {
		body["RightTablePartitionField"] = request.RightTablePartitionField
	}

	if !dara.IsNil(request.TableMetaId) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.TopNQuantity) {
		body["TopNQuantity"] = request.TopNQuantity
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataDiagnosis"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnoses"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data diagnosis task. This API supports various analysis types.
//
// Description:
//
// ## Description
//
// - This API creates a data diagnosis task. It supports various analysis types, including item or user change rate analysis, user preference statistics cycle analysis, two-table join analysis, basic statistical analysis, and abnormal behavior analysis.
//
// - The content of the `Config` parameter depends on the value of the `Type` parameter. For more information, see the example configurations in this topic.
//
// - To run the task on a schedule, specify the `CycleTime` parameter. If this parameter is omitted, the task runs only once.
//
// - The optional `TopNQuantity` parameter specifies the number of top results to return.
//
// @param request - CreateDataDiagnosisRequest
//
// @return CreateDataDiagnosisResponse
func (client *Client) CreateDataDiagnosis(request *CreateDataDiagnosisRequest) (_result *CreateDataDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataDiagnosisResponse{}
	_body, _err := client.CreateDataDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data diagnosis (rerun) job for a specified time period.
//
// Description:
//
// ## Description
//
// This operation creates a data diagnosis (rerun) job for a specific instance within a specified time frame. To ensure the job runs correctly, provide accurate values for the `DataDiagnosisId`, `InstanceId`, `StartDate`, and `EndDate` parameters.
//
// @param request - CreateDataDiagnosisJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataDiagnosisJobsResponse
func (client *Client) CreateDataDiagnosisJobsWithOptions(request *CreateDataDiagnosisJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataDiagnosisJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataDiagnosisId) {
		body["DataDiagnosisId"] = request.DataDiagnosisId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataDiagnosisJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/batch/datadiagnosisjobs/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataDiagnosisJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data diagnosis (rerun) job for a specified time period.
//
// Description:
//
// ## Description
//
// This operation creates a data diagnosis (rerun) job for a specific instance within a specified time frame. To ensure the job runs correctly, provide accurate values for the `DataDiagnosisId`, `InstanceId`, `StartDate`, and `EndDate` parameters.
//
// @param request - CreateDataDiagnosisJobsRequest
//
// @return CreateDataDiagnosisJobsResponse
func (client *Client) CreateDataDiagnosisJobs(request *CreateDataDiagnosisJobsRequest) (_result *CreateDataDiagnosisJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataDiagnosisJobsResponse{}
	_body, _err := client.CreateDataDiagnosisJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an engine configuration.
//
// @param request - CreateEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEngineConfigResponse
func (client *Client) CreateEngineConfigWithOptions(request *CreateEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEngineConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEngineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an engine configuration.
//
// @param request - CreateEngineConfigRequest
//
// @return CreateEngineConfigResponse
func (client *Client) CreateEngineConfig(request *CreateEngineConfigRequest) (_result *CreateEngineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEngineConfigResponse{}
	_body, _err := client.CreateEngineConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an experiment in a specified experiment group.
//
// @param request - CreateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentResponse
func (client *Client) CreateExperimentWithOptions(request *CreateExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExperimentGroupId) {
		body["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.FlowPercent) {
		body["FlowPercent"] = request.FlowPercent
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an experiment in a specified experiment group.
//
// @param request - CreateExperimentRequest
//
// @return CreateExperimentResponse
func (client *Client) CreateExperiment(request *CreateExperimentRequest) (_result *CreateExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentResponse{}
	_body, _err := client.CreateExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an experiment group in a specified layer. You can use experiment groups to categorize experiments and observe their overall performance.
//
// @param request - CreateExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentGroupResponse
func (client *Client) CreateExperimentGroupWithOptions(request *CreateExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.CrowdId) {
		body["CrowdId"] = request.CrowdId
	}

	if !dara.IsNil(request.CrowdTargetType) {
		body["CrowdTargetType"] = request.CrowdTargetType
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DistributionTimeDuration) {
		body["DistributionTimeDuration"] = request.DistributionTimeDuration
	}

	if !dara.IsNil(request.DistributionType) {
		body["DistributionType"] = request.DistributionType
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LayerId) {
		body["LayerId"] = request.LayerId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedAA) {
		body["NeedAA"] = request.NeedAA
	}

	if !dara.IsNil(request.RandomFlow) {
		body["RandomFlow"] = request.RandomFlow
	}

	if !dara.IsNil(request.ReservedBuckets) {
		body["ReservedBuckets"] = request.ReservedBuckets
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an experiment group in a specified layer. You can use experiment groups to categorize experiments and observe their overall performance.
//
// @param request - CreateExperimentGroupRequest
//
// @return CreateExperimentGroupResponse
func (client *Client) CreateExperimentGroup(request *CreateExperimentGroupRequest) (_result *CreateExperimentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentGroupResponse{}
	_body, _err := client.CreateExperimentGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a feature consistency check job.
//
// @param request - CreateFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureConsistencyCheckJobResponse
func (client *Client) CreateFeatureConsistencyCheckJobWithOptions(request *CreateFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureConsistencyCheckJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.FeatureConsistencyCheckJobConfigId) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SamplingDuration) {
		body["SamplingDuration"] = request.SamplingDuration
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFeatureConsistencyCheckJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a feature consistency check job.
//
// @param request - CreateFeatureConsistencyCheckJobRequest
//
// @return CreateFeatureConsistencyCheckJobResponse
func (client *Client) CreateFeatureConsistencyCheckJob(request *CreateFeatureConsistencyCheckJobRequest) (_result *CreateFeatureConsistencyCheckJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CreateFeatureConsistencyCheckJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configure a feature consistency check task.
//
// @param request - CreateFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureConsistencyCheckJobConfigResponse
func (client *Client) CreateFeatureConsistencyCheckJobConfigWithOptions(request *CreateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureConsistencyCheckJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CompareFeature) {
		body["CompareFeature"] = request.CompareFeature
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetMountPath) {
		body["DatasetMountPath"] = request.DatasetMountPath
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		body["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.DatasetUri) {
		body["DatasetUri"] = request.DatasetUri
	}

	if !dara.IsNil(request.DefaultRoute) {
		body["DefaultRoute"] = request.DefaultRoute
	}

	if !dara.IsNil(request.EasServiceName) {
		body["EasServiceName"] = request.EasServiceName
	}

	if !dara.IsNil(request.EasyRecPackagePath) {
		body["EasyRecPackagePath"] = request.EasyRecPackagePath
	}

	if !dara.IsNil(request.EasyRecVersion) {
		body["EasyRecVersion"] = request.EasyRecVersion
	}

	if !dara.IsNil(request.FeatureDisplayExclude) {
		body["FeatureDisplayExclude"] = request.FeatureDisplayExclude
	}

	if !dara.IsNil(request.FeatureLandingResourceId) {
		body["FeatureLandingResourceId"] = request.FeatureLandingResourceId
	}

	if !dara.IsNil(request.FeaturePriority) {
		body["FeaturePriority"] = request.FeaturePriority
	}

	if !dara.IsNil(request.FeatureStoreItemId) {
		body["FeatureStoreItemId"] = request.FeatureStoreItemId
	}

	if !dara.IsNil(request.FeatureStoreModelId) {
		body["FeatureStoreModelId"] = request.FeatureStoreModelId
	}

	if !dara.IsNil(request.FeatureStoreProjectId) {
		body["FeatureStoreProjectId"] = request.FeatureStoreProjectId
	}

	if !dara.IsNil(request.FeatureStoreProjectName) {
		body["FeatureStoreProjectName"] = request.FeatureStoreProjectName
	}

	if !dara.IsNil(request.FeatureStoreSeqFeatureView) {
		body["FeatureStoreSeqFeatureView"] = request.FeatureStoreSeqFeatureView
	}

	if !dara.IsNil(request.FeatureStoreUserId) {
		body["FeatureStoreUserId"] = request.FeatureStoreUserId
	}

	if !dara.IsNil(request.FgJarVersion) {
		body["FgJarVersion"] = request.FgJarVersion
	}

	if !dara.IsNil(request.FgJsonFileName) {
		body["FgJsonFileName"] = request.FgJsonFileName
	}

	if !dara.IsNil(request.GenerateZip) {
		body["GenerateZip"] = request.GenerateZip
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemIdField) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !dara.IsNil(request.ItemTable) {
		body["ItemTable"] = request.ItemTable
	}

	if !dara.IsNil(request.ItemTablePartitionField) {
		body["ItemTablePartitionField"] = request.ItemTablePartitionField
	}

	if !dara.IsNil(request.ItemTablePartitionFieldFormat) {
		body["ItemTablePartitionFieldFormat"] = request.ItemTablePartitionFieldFormat
	}

	if !dara.IsNil(request.MaxcomputeSchema) {
		body["MaxcomputeSchema"] = request.MaxcomputeSchema
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OssResourceId) {
		body["OssResourceId"] = request.OssResourceId
	}

	if !dara.IsNil(request.PredictWorkerCount) {
		body["PredictWorkerCount"] = request.PredictWorkerCount
	}

	if !dara.IsNil(request.PredictWorkerCpu) {
		body["PredictWorkerCpu"] = request.PredictWorkerCpu
	}

	if !dara.IsNil(request.PredictWorkerMemory) {
		body["PredictWorkerMemory"] = request.PredictWorkerMemory
	}

	if !dara.IsNil(request.ResourceConfig) {
		body["ResourceConfig"] = request.ResourceConfig
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SwitchId) {
		body["SwitchId"] = request.SwitchId
	}

	if !dara.IsNil(request.UseFeatureStore) {
		body["UseFeatureStore"] = request.UseFeatureStore
	}

	if !dara.IsNil(request.UserIdField) {
		body["UserIdField"] = request.UserIdField
	}

	if !dara.IsNil(request.UserTable) {
		body["UserTable"] = request.UserTable
	}

	if !dara.IsNil(request.UserTablePartitionField) {
		body["UserTablePartitionField"] = request.UserTablePartitionField
	}

	if !dara.IsNil(request.UserTablePartitionFieldFormat) {
		body["UserTablePartitionFieldFormat"] = request.UserTablePartitionFieldFormat
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkflowName) {
		body["WorkflowName"] = request.WorkflowName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configure a feature consistency check task.
//
// @param request - CreateFeatureConsistencyCheckJobConfigRequest
//
// @return CreateFeatureConsistencyCheckJobConfigResponse
func (client *Client) CreateFeatureConsistencyCheckJobConfig(request *CreateFeatureConsistencyCheckJobConfigRequest) (_result *CreateFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CreateFeatureConsistencyCheckJobConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a configuration resource for a specified instance.
//
// @param request - CreateInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResourceResponse
func (client *Client) CreateInstanceResourceWithOptions(InstanceId *string, request *CreateInstanceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.Group) {
		body["Group"] = request.Group
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration resource for a specified instance.
//
// @param request - CreateInstanceResourceRequest
//
// @return CreateInstanceResourceResponse
func (client *Client) CreateInstanceResource(InstanceId *string, request *CreateInstanceResourceRequest) (_result *CreateInstanceResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResourceResponse{}
	_body, _err := client.CreateInstanceResourceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a laboratory. A laboratory isolates a segment of traffic for running experiments.
//
// @param request - CreateLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLaboratoryResponse
func (client *Client) CreateLaboratoryWithOptions(request *CreateLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLaboratoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BucketCount) {
		body["BucketCount"] = request.BucketCount
	}

	if !dara.IsNil(request.BucketType) {
		body["BucketType"] = request.BucketType
	}

	if !dara.IsNil(request.Buckets) {
		body["Buckets"] = request.Buckets
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a laboratory. A laboratory isolates a segment of traffic for running experiments.
//
// @param request - CreateLaboratoryRequest
//
// @return CreateLaboratoryResponse
func (client *Client) CreateLaboratory(request *CreateLaboratoryRequest) (_result *CreateLaboratoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLaboratoryResponse{}
	_body, _err := client.CreateLaboratoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a layer in a specified laboratory for layered experiments. Layers are orthogonal to each other, allowing experiments to run independently and preventing traffic starvation.
//
// @param request - CreateLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLayerResponse
func (client *Client) CreateLayerWithOptions(request *CreateLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLayerResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LaboratoryId) {
		body["LaboratoryId"] = request.LaboratoryId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a layer in a specified laboratory for layered experiments. Layers are orthogonal to each other, allowing experiments to run independently and preventing traffic starvation.
//
// @param request - CreateLayerRequest
//
// @return CreateLayerResponse
func (client *Client) CreateLayer(request *CreateLayerRequest) (_result *CreateLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLayerResponse{}
	_body, _err := client.CreateLayerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ABTest parameter for a specific scene in a specified environment.
//
// Description:
//
// ## Operation description.
//
// @param request - CreateParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParamResponse
func (client *Client) CreateParamWithOptions(request *CreateParamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateParamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParam"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ABTest parameter for a specific scene in a specified environment.
//
// Description:
//
// ## Operation description.
//
// @param request - CreateParamRequest
//
// @return CreateParamResponse
func (client *Client) CreateParam(request *CreateParamRequest) (_result *CreateParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateParamResponse{}
	_body, _err := client.CreateParamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initializes a Recall Management configuration, including the instance ID, user information, and network configurations.
//
// Description:
//
// ## Request
//
// @param request - CreateRecallManagementConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecallManagementConfigResponse
func (client *Client) CreateRecallManagementConfigWithOptions(request *CreateRecallManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRecallManagementConfigResponse, _err error) {
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

	if !dara.IsNil(request.NetworkConfigs) {
		body["NetworkConfigs"] = request.NetworkConfigs
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecallManagementConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementconfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecallManagementConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes a Recall Management configuration, including the instance ID, user information, and network configurations.
//
// Description:
//
// ## Request
//
// @param request - CreateRecallManagementConfigRequest
//
// @return CreateRecallManagementConfigResponse
func (client *Client) CreateRecallManagementConfig(request *CreateRecallManagementConfigRequest) (_result *CreateRecallManagementConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRecallManagementConfigResponse{}
	_body, _err := client.CreateRecallManagementConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new recall management service.
//
// Description:
//
// ## Description
//
// To create a recall management service, call this API with a specified instance ID, service name, and service description. Ensure that the `InstanceId` parameter is valid.
//
// @param request - CreateRecallManagementServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecallManagementServiceResponse
func (client *Client) CreateRecallManagementServiceWithOptions(request *CreateRecallManagementServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRecallManagementServiceResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecallManagementService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecallManagementServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new recall management service.
//
// Description:
//
// ## Description
//
// To create a recall management service, call this API with a specified instance ID, service name, and service description. Ensure that the `InstanceId` parameter is valid.
//
// @param request - CreateRecallManagementServiceRequest
//
// @return CreateRecallManagementServiceResponse
func (client *Client) CreateRecallManagementService(request *CreateRecallManagementServiceRequest) (_result *CreateRecallManagementServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRecallManagementServiceResponse{}
	_body, _err := client.CreateRecallManagementServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new recall management service version that supports configuring multiple recall strategies.
//
// Description:
//
// ## Request
//
// - Use this API to create a new version of a recall management service.
//
// - To create a new version from an existing one, specify the source recall management service version ID.
//
// - You can configure detailed recall rules, such as the recall name, description, priority, and recall type.
//
// - Configure operators such as filter, trigger, feature extraction, and join.
//
// - The merge configuration specifies how to merge multiple recall results and supports two merge methods: weight-based and alternating.
//
// - All configuration items are optional.
//
// @param request - CreateRecallManagementServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecallManagementServiceVersionResponse
func (client *Client) CreateRecallManagementServiceVersionWithOptions(RecallManagementServiceId *string, request *CreateRecallManagementServiceVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRecallManagementServiceVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Configs) {
		body["Configs"] = request.Configs
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SourceRecallManagementServiceVersionId) {
		body["SourceRecallManagementServiceVersionId"] = request.SourceRecallManagementServiceVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecallManagementServiceVersion"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecallManagementServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new recall management service version that supports configuring multiple recall strategies.
//
// Description:
//
// ## Request
//
// - Use this API to create a new version of a recall management service.
//
// - To create a new version from an existing one, specify the source recall management service version ID.
//
// - You can configure detailed recall rules, such as the recall name, description, priority, and recall type.
//
// - Configure operators such as filter, trigger, feature extraction, and join.
//
// - The merge configuration specifies how to merge multiple recall results and supports two merge methods: weight-based and alternating.
//
// - All configuration items are optional.
//
// @param request - CreateRecallManagementServiceVersionRequest
//
// @return CreateRecallManagementServiceVersionResponse
func (client *Client) CreateRecallManagementServiceVersion(RecallManagementServiceId *string, request *CreateRecallManagementServiceVersionRequest) (_result *CreateRecallManagementServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRecallManagementServiceVersionResponse{}
	_body, _err := client.CreateRecallManagementServiceVersionWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a configuration for a specified version of the Recall Manager service, including its recall and merge settings.
//
// Description:
//
// ## Request
//
// - This API creates a configuration for a specific version of the Recall Management Service.
//
// - The `ConfigType` parameter specifies the configuration type, which can be either recall or merge.
//
// - Use the `RecallConfig` and `MergeConfig` parameters to provide the recall and merge configurations, respectively.
//
// - Required parameters must be provided in the specified data formats.
//
// - Optional parameter values must be consistent with your business logic.
//
// @param request - CreateRecallManagementServiceVersionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecallManagementServiceVersionConfigResponse
func (client *Client) CreateRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, request *CreateRecallManagementServiceVersionConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRecallManagementServiceVersionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigType) {
		body["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MergeConfig) {
		body["MergeConfig"] = request.MergeConfig
	}

	if !dara.IsNil(request.RecallConfig) {
		body["RecallConfig"] = request.RecallConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecallManagementServiceVersionConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionId)) + "/configs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration for a specified version of the Recall Manager service, including its recall and merge settings.
//
// Description:
//
// ## Request
//
// - This API creates a configuration for a specific version of the Recall Management Service.
//
// - The `ConfigType` parameter specifies the configuration type, which can be either recall or merge.
//
// - Use the `RecallConfig` and `MergeConfig` parameters to provide the recall and merge configurations, respectively.
//
// - Required parameters must be provided in the specified data formats.
//
// - Optional parameter values must be consistent with your business logic.
//
// @param request - CreateRecallManagementServiceVersionConfigRequest
//
// @return CreateRecallManagementServiceVersionConfigResponse
func (client *Client) CreateRecallManagementServiceVersionConfig(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, request *CreateRecallManagementServiceVersionConfigRequest) (_result *CreateRecallManagementServiceVersionConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.CreateRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId, RecallManagementServiceVersionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a recall management table that supports multiple data sources and configuration options.
//
// Description:
//
// ## Request
//
// - The **InstanceId**, **Name**, **Description**, **Type**, and **DataSource*	- parameters are required.
//
// - The **RecallType*	- parameter is optional. If provided, it must be a valid enum value.
//
// - For each field in the **Fields*	- parameter, you must define its name, type, and attributes. You must mark at least one field as Primary.
//
// - If you use MaxCompute as the data source, you must specify the **MaxcomputeProjectName*	- and **MaxcomputeTableName*	- parameters. The **MaxcomputeSchema*	- parameter is optional.
//
// - For vector fields, the values of the **VectorDimension*	- and **VectorMetricType*	- parameters must match the actual data.
//
// - Use the **Config*	- field to provide additional configuration as a JSON string.
//
// - Use fluctuation threshold parameters, such as **EnableRowCountFluctuationThreshold**, to monitor changes in row count or table size. Enable these parameters and set appropriate thresholds as needed.
//
// @param request - CreateRecallManagementTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecallManagementTableResponse
func (client *Client) CreateRecallManagementTableWithOptions(request *CreateRecallManagementTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRecallManagementTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.DataSource) {
		body["DataSource"] = request.DataSource
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableDataSizeFluctuationThreshold) {
		body["EnableDataSizeFluctuationThreshold"] = request.EnableDataSizeFluctuationThreshold
	}

	if !dara.IsNil(request.EnableRowCountFluctuationThreshold) {
		body["EnableRowCountFluctuationThreshold"] = request.EnableRowCountFluctuationThreshold
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxDataSizeFluctuationThreshold) {
		body["MaxDataSizeFluctuationThreshold"] = request.MaxDataSizeFluctuationThreshold
	}

	if !dara.IsNil(request.MaxRowCountFluctuationThreshold) {
		body["MaxRowCountFluctuationThreshold"] = request.MaxRowCountFluctuationThreshold
	}

	if !dara.IsNil(request.MaxcomputeProjectName) {
		body["MaxcomputeProjectName"] = request.MaxcomputeProjectName
	}

	if !dara.IsNil(request.MaxcomputeSchema) {
		body["MaxcomputeSchema"] = request.MaxcomputeSchema
	}

	if !dara.IsNil(request.MaxcomputeTableName) {
		body["MaxcomputeTableName"] = request.MaxcomputeTableName
	}

	if !dara.IsNil(request.MinDataSizeFluctuationThreshold) {
		body["MinDataSizeFluctuationThreshold"] = request.MinDataSizeFluctuationThreshold
	}

	if !dara.IsNil(request.MinRowCountFluctuationThreshold) {
		body["MinRowCountFluctuationThreshold"] = request.MinRowCountFluctuationThreshold
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RecallType) {
		body["RecallType"] = request.RecallType
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecallManagementTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecallManagementTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a recall management table that supports multiple data sources and configuration options.
//
// Description:
//
// ## Request
//
// - The **InstanceId**, **Name**, **Description**, **Type**, and **DataSource*	- parameters are required.
//
// - The **RecallType*	- parameter is optional. If provided, it must be a valid enum value.
//
// - For each field in the **Fields*	- parameter, you must define its name, type, and attributes. You must mark at least one field as Primary.
//
// - If you use MaxCompute as the data source, you must specify the **MaxcomputeProjectName*	- and **MaxcomputeTableName*	- parameters. The **MaxcomputeSchema*	- parameter is optional.
//
// - For vector fields, the values of the **VectorDimension*	- and **VectorMetricType*	- parameters must match the actual data.
//
// - Use the **Config*	- field to provide additional configuration as a JSON string.
//
// - Use fluctuation threshold parameters, such as **EnableRowCountFluctuationThreshold**, to monitor changes in row count or table size. Enable these parameters and set appropriate thresholds as needed.
//
// @param request - CreateRecallManagementTableRequest
//
// @return CreateRecallManagementTableResponse
func (client *Client) CreateRecallManagementTable(request *CreateRecallManagementTableRequest) (_result *CreateRecallManagementTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRecallManagementTableResponse{}
	_body, _err := client.CreateRecallManagementTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源规则
//
// @param request - CreateResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceRuleResponse
func (client *Client) CreateResourceRuleWithOptions(request *CreateResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceRuleResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricOperationType) {
		body["MetricOperationType"] = request.MetricOperationType
	}

	if !dara.IsNil(request.MetricPullInfo) {
		body["MetricPullInfo"] = request.MetricPullInfo
	}

	if !dara.IsNil(request.MetricPullPeriod) {
		body["MetricPullPeriod"] = request.MetricPullPeriod
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RuleComputingDefinition) {
		body["RuleComputingDefinition"] = request.RuleComputingDefinition
	}

	if !dara.IsNil(request.RuleItems) {
		body["RuleItems"] = request.RuleItems
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源规则
//
// @param request - CreateResourceRuleRequest
//
// @return CreateResourceRuleResponse
func (client *Client) CreateResourceRule(request *CreateResourceRuleRequest) (_result *CreateResourceRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceRuleResponse{}
	_body, _err := client.CreateResourceRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源规则条目
//
// @param request - CreateResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceRuleItemResponse
func (client *Client) CreateResourceRuleItemWithOptions(ResourceRuleId *string, request *CreateResourceRuleItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceRuleItemResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxValue) {
		body["MaxValue"] = request.MaxValue
	}

	if !dara.IsNil(request.MinValue) {
		body["MinValue"] = request.MinValue
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceRuleItem"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/items"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceRuleItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源规则条目
//
// @param request - CreateResourceRuleItemRequest
//
// @return CreateResourceRuleItemResponse
func (client *Client) CreateResourceRuleItem(ResourceRuleId *string, request *CreateResourceRuleItemRequest) (_result *CreateResourceRuleItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceRuleItemResponse{}
	_body, _err := client.CreateResourceRuleItemWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建样本一致性任务
//
// @param request - CreateSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSampleConsistencyJobResponse
func (client *Client) CreateSampleConsistencyJobWithOptions(request *CreateSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSampleConsistencyJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EasModelServiceName) {
		body["EasModelServiceName"] = request.EasModelServiceName
	}

	if !dara.IsNil(request.FeatureSaveResourceId) {
		body["FeatureSaveResourceId"] = request.FeatureSaveResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemIdField) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PartitionField) {
		body["PartitionField"] = request.PartitionField
	}

	if !dara.IsNil(request.PartitionFieldFormat) {
		body["PartitionFieldFormat"] = request.PartitionFieldFormat
	}

	if !dara.IsNil(request.RequestIdField) {
		body["RequestIdField"] = request.RequestIdField
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SampleTableName) {
		body["SampleTableName"] = request.SampleTableName
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.UserIdField) {
		body["UserIdField"] = request.UserIdField
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSampleConsistencyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建样本一致性任务
//
// @param request - CreateSampleConsistencyJobRequest
//
// @return CreateSampleConsistencyJobResponse
func (client *Client) CreateSampleConsistencyJob(request *CreateSampleConsistencyJobRequest) (_result *CreateSampleConsistencyJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSampleConsistencyJobResponse{}
	_body, _err := client.CreateSampleConsistencyJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scene for metric and experiment analysis.
//
// @param request - CreateSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSceneResponse
func (client *Client) CreateSceneWithOptions(request *CreateSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSceneResponse, _err error) {
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

	if !dara.IsNil(request.Flows) {
		body["Flows"] = request.Flows
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scene for metric and experiment analysis.
//
// @param request - CreateSceneRequest
//
// @return CreateSceneResponse
func (client *Client) CreateScene(request *CreateSceneRequest) (_result *CreateSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSceneResponse{}
	_body, _err := client.CreateSceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a sub-crowd for a specified crowd.
//
// @param request - CreateSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCrowdResponse
func (client *Client) CreateSubCrowdWithOptions(CrowdId *string, request *CreateSubCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSubCrowdResponse, _err error) {
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

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sub-crowd for a specified crowd.
//
// @param request - CreateSubCrowdRequest
//
// @return CreateSubCrowdResponse
func (client *Client) CreateSubCrowd(CrowdId *string, request *CreateSubCrowdRequest) (_result *CreateSubCrowdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSubCrowdResponse{}
	_body, _err := client.CreateSubCrowdWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data table.
//
// @param request - CreateTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableMetaResponse
func (client *Client) CreateTableMetaWithOptions(request *CreateTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTableMetaResponse, _err error) {
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

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Module) {
		body["Module"] = request.Module
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data table.
//
// @param request - CreateTableMetaRequest
//
// @return CreateTableMetaResponse
func (client *Client) CreateTableMeta(request *CreateTableMetaRequest) (_result *CreateTableMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTableMetaResponse{}
	_body, _err := client.CreateTableMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new traffic control target and sets its parameters, such as the item condition, event, and value.
//
// Description:
//
// ## Usage notes
//
// - The `ItemConditionType` parameter supports two formats: array format (Array) and expression format (Expression). Based on the format you select, you must provide either `ItemConditionArray` or `ItemConditionExpress`.
//
// - The `StatisPeriod` parameter defaults to daily. For hourly statistics, you must set this parameter explicitly.
//
// - The `ToleranceValue` and `NewProductRegulation` parameters are optional.
//
// - The `Status` parameter controls whether a new traffic control target takes effect immediately. By default, new targets are inactive.
//
// @param request - CreateTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlTargetResponse
func (client *Client) CreateTrafficControlTargetWithOptions(request *CreateTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrafficControlTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Event) {
		body["Event"] = request.Event
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NewProductRegulation) {
		body["NewProductRegulation"] = request.NewProductRegulation
	}

	if !dara.IsNil(request.RecallName) {
		body["RecallName"] = request.RecallName
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisPeriod) {
		body["StatisPeriod"] = request.StatisPeriod
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ToleranceValue) {
		body["ToleranceValue"] = request.ToleranceValue
	}

	if !dara.IsNil(request.TrafficControlTaskId) {
		body["TrafficControlTaskId"] = request.TrafficControlTaskId
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrafficControlTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new traffic control target and sets its parameters, such as the item condition, event, and value.
//
// Description:
//
// ## Usage notes
//
// - The `ItemConditionType` parameter supports two formats: array format (Array) and expression format (Expression). Based on the format you select, you must provide either `ItemConditionArray` or `ItemConditionExpress`.
//
// - The `StatisPeriod` parameter defaults to daily. For hourly statistics, you must set this parameter explicitly.
//
// - The `ToleranceValue` and `NewProductRegulation` parameters are optional.
//
// - The `Status` parameter controls whether a new traffic control target takes effect immediately. By default, new targets are inactive.
//
// @param request - CreateTrafficControlTargetRequest
//
// @return CreateTrafficControlTargetResponse
func (client *Client) CreateTrafficControlTarget(request *CreateTrafficControlTargetRequest) (_result *CreateTrafficControlTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrafficControlTargetResponse{}
	_body, _err := client.CreateTrafficControlTargetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new traffic control task with multiple conditions and target configurations.
//
// Description:
//
// ## Request
//
// - Use this API to create a new traffic control task. You can define a traffic control policy for different scenarios, time ranges, and conditions for users or items.
//
// - The `ExecutionTime` parameter specifies the execution time mode for the task. If you select the `TimeRange` mode, you must provide both the `StartTime` and `EndTime` parameters.
//
// - The `TrafficControlTargets` parameter is required. For each traffic control target, you must specify its name, time range, event type, and expected value.
//
// - You can use the `UserConditionType` and `ItemConditionType` parameters to define conditions for the target user group and items.
//
// - Set the `ControlLogic` parameter to `Guaranteed` for guaranteed control or to `Approach` for approach control.
//
// - To configure new product regulation, use the `NewProductRegulation` field.
//
// @param request - CreateTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlTaskResponse
func (client *Client) CreateTrafficControlTaskWithOptions(request *CreateTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrafficControlTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BehaviorTableMetaId) {
		body["BehaviorTableMetaId"] = request.BehaviorTableMetaId
	}

	if !dara.IsNil(request.ControlGranularity) {
		body["ControlGranularity"] = request.ControlGranularity
	}

	if !dara.IsNil(request.ControlLogic) {
		body["ControlLogic"] = request.ControlLogic
	}

	if !dara.IsNil(request.ControlType) {
		body["ControlType"] = request.ControlType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EffectiveSceneIds) {
		body["EffectiveSceneIds"] = request.EffectiveSceneIds
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecutionTime) {
		body["ExecutionTime"] = request.ExecutionTime
	}

	if !dara.IsNil(request.FlinkResourceId) {
		body["FlinkResourceId"] = request.FlinkResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.ItemTableMetaId) {
		body["ItemTableMetaId"] = request.ItemTableMetaId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PreExperimentIds) {
		body["PreExperimentIds"] = request.PreExperimentIds
	}

	if !dara.IsNil(request.ProdExperimentIds) {
		body["ProdExperimentIds"] = request.ProdExperimentIds
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceIds) {
		body["ServiceIds"] = request.ServiceIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisBehaviorConditionArray) {
		body["StatisBehaviorConditionArray"] = request.StatisBehaviorConditionArray
	}

	if !dara.IsNil(request.StatisBehaviorConditionExpress) {
		body["StatisBehaviorConditionExpress"] = request.StatisBehaviorConditionExpress
	}

	if !dara.IsNil(request.StatisBehaviorConditionType) {
		body["StatisBehaviorConditionType"] = request.StatisBehaviorConditionType
	}

	if !dara.IsNil(request.TrafficControlTargets) {
		body["TrafficControlTargets"] = request.TrafficControlTargets
	}

	if !dara.IsNil(request.UserConditionArray) {
		body["UserConditionArray"] = request.UserConditionArray
	}

	if !dara.IsNil(request.UserConditionExpress) {
		body["UserConditionExpress"] = request.UserConditionExpress
	}

	if !dara.IsNil(request.UserConditionType) {
		body["UserConditionType"] = request.UserConditionType
	}

	if !dara.IsNil(request.UserTableMetaId) {
		body["UserTableMetaId"] = request.UserTableMetaId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new traffic control task with multiple conditions and target configurations.
//
// Description:
//
// ## Request
//
// - Use this API to create a new traffic control task. You can define a traffic control policy for different scenarios, time ranges, and conditions for users or items.
//
// - The `ExecutionTime` parameter specifies the execution time mode for the task. If you select the `TimeRange` mode, you must provide both the `StartTime` and `EndTime` parameters.
//
// - The `TrafficControlTargets` parameter is required. For each traffic control target, you must specify its name, time range, event type, and expected value.
//
// - You can use the `UserConditionType` and `ItemConditionType` parameters to define conditions for the target user group and items.
//
// - Set the `ControlLogic` parameter to `Guaranteed` for guaranteed control or to `Approach` for approach control.
//
// - To configure new product regulation, use the `NewProductRegulation` field.
//
// @param request - CreateTrafficControlTaskRequest
//
// @return CreateTrafficControlTaskResponse
func (client *Client) CreateTrafficControlTask(request *CreateTrafficControlTaskRequest) (_result *CreateTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrafficControlTaskResponse{}
	_body, _err := client.CreateTrafficControlTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对指定资源规则中的计算公式进行调试
//
// @param tmpReq - DebugResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugResourceRuleResponse
func (client *Client) DebugResourceRuleWithOptions(ResourceRuleId *string, tmpReq *DebugResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DebugResourceRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DebugResourceRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetricInfo) {
		request.MetricInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricInfo, dara.String("MetricInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricInfoShrink) {
		query["MetricInfo"] = request.MetricInfoShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/action/debug"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugResourceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对指定资源规则中的计算公式进行调试
//
// @param request - DebugResourceRuleRequest
//
// @return DebugResourceRuleResponse
func (client *Client) DebugResourceRule(ResourceRuleId *string, request *DebugResourceRuleRequest) (_result *DebugResourceRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DebugResourceRuleResponse{}
	_body, _err := client.DebugResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified A/B test metric.
//
// @param request - DeleteABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABMetricResponse
func (client *Client) DeleteABMetricWithOptions(ABMetricId *string, request *DeleteABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteABMetricResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics/" + dara.PercentEncode(dara.StringValue(ABMetricId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified A/B test metric.
//
// @param request - DeleteABMetricRequest
//
// @return DeleteABMetricResponse
func (client *Client) DeleteABMetric(ABMetricId *string, request *DeleteABMetricRequest) (_result *DeleteABMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABMetricResponse{}
	_body, _err := client.DeleteABMetricWithOptions(ABMetricId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an A/B test metric group.
//
// @param request - DeleteABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABMetricGroupResponse
func (client *Client) DeleteABMetricGroupWithOptions(ABMetricGroupId *string, request *DeleteABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteABMetricGroupResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an A/B test metric group.
//
// @param request - DeleteABMetricGroupRequest
//
// @return DeleteABMetricGroupResponse
func (client *Client) DeleteABMetricGroup(ABMetricGroupId *string, request *DeleteABMetricGroupRequest) (_result *DeleteABMetricGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABMetricGroupResponse{}
	_body, _err := client.DeleteABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified audience.
//
// @param request - DeleteCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrowdResponse
func (client *Client) DeleteCrowdWithOptions(CrowdId *string, request *DeleteCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCrowdResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified audience.
//
// @param request - DeleteCrowdRequest
//
// @return DeleteCrowdResponse
func (client *Client) DeleteCrowd(CrowdId *string, request *DeleteCrowdRequest) (_result *DeleteCrowdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCrowdResponse{}
	_body, _err := client.DeleteCrowdWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data diagnosis configuration using the specified DataDiagnosisId and InstanceId.
//
// Description:
//
// ## Description
//
// Ensure you provide the correct `DataDiagnosisId` and `InstanceId` to avoid accidental deletion.
//
// @param request - DeleteDataDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataDiagnosisResponse
func (client *Client) DeleteDataDiagnosisWithOptions(DataDiagnosisId *string, request *DeleteDataDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataDiagnosisResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataDiagnosis"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnoses/" + dara.PercentEncode(dara.StringValue(DataDiagnosisId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data diagnosis configuration using the specified DataDiagnosisId and InstanceId.
//
// Description:
//
// ## Description
//
// Ensure you provide the correct `DataDiagnosisId` and `InstanceId` to avoid accidental deletion.
//
// @param request - DeleteDataDiagnosisRequest
//
// @return DeleteDataDiagnosisResponse
func (client *Client) DeleteDataDiagnosis(DataDiagnosisId *string, request *DeleteDataDiagnosisRequest) (_result *DeleteDataDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataDiagnosisResponse{}
	_body, _err := client.DeleteDataDiagnosisWithOptions(DataDiagnosisId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified engine configuration.
//
// Description:
//
// Deletes a specified engine configuration.
//
// @param request - DeleteEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEngineConfigResponse
func (client *Client) DeleteEngineConfigWithOptions(EngineConfigId *string, request *DeleteEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEngineConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteAll) {
		query["DeleteAll"] = request.DeleteAll
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEngineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified engine configuration.
//
// Description:
//
// Deletes a specified engine configuration.
//
// @param request - DeleteEngineConfigRequest
//
// @return DeleteEngineConfigResponse
func (client *Client) DeleteEngineConfig(EngineConfigId *string, request *DeleteEngineConfigRequest) (_result *DeleteEngineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEngineConfigResponse{}
	_body, _err := client.DeleteEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified experiment.
//
// @param request - DeleteExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithOptions(ExperimentId *string, request *DeleteExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified experiment.
//
// @param request - DeleteExperimentRequest
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperiment(ExperimentId *string, request *DeleteExperimentRequest) (_result *DeleteExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentResponse{}
	_body, _err := client.DeleteExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified experiment group.
//
// @param request - DeleteExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentGroupResponse
func (client *Client) DeleteExperimentGroupWithOptions(ExperimentGroupId *string, request *DeleteExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentGroupResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified experiment group.
//
// @param request - DeleteExperimentGroupRequest
//
// @return DeleteExperimentGroupResponse
func (client *Client) DeleteExperimentGroup(ExperimentGroupId *string, request *DeleteExperimentGroupRequest) (_result *DeleteExperimentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentGroupResponse{}
	_body, _err := client.DeleteExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a configuration resource from an instance.
//
// @param request - DeleteInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResourceResponse
func (client *Client) DeleteInstanceResourceWithOptions(InstanceId *string, ResourceId *string, request *DeleteInstanceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResourceResponse, _err error) {
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
		Action:      dara.String("DeleteInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a configuration resource from an instance.
//
// @param request - DeleteInstanceResourceRequest
//
// @return DeleteInstanceResourceResponse
func (client *Client) DeleteInstanceResource(InstanceId *string, ResourceId *string, request *DeleteInstanceResourceRequest) (_result *DeleteInstanceResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResourceResponse{}
	_body, _err := client.DeleteInstanceResourceWithOptions(InstanceId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified Lab.
//
// @param request - DeleteLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLaboratoryResponse
func (client *Client) DeleteLaboratoryWithOptions(LaboratoryId *string, request *DeleteLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLaboratoryResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified Lab.
//
// @param request - DeleteLaboratoryRequest
//
// @return DeleteLaboratoryResponse
func (client *Client) DeleteLaboratory(LaboratoryId *string, request *DeleteLaboratoryRequest) (_result *DeleteLaboratoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLaboratoryResponse{}
	_body, _err := client.DeleteLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified layer.
//
// @param request - DeleteLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayerResponse
func (client *Client) DeleteLayerWithOptions(LayerId *string, request *DeleteLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLayerResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers/" + dara.PercentEncode(dara.StringValue(LayerId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified layer.
//
// @param request - DeleteLayerRequest
//
// @return DeleteLayerResponse
func (client *Client) DeleteLayer(LayerId *string, request *DeleteLayerRequest) (_result *DeleteLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLayerResponse{}
	_body, _err := client.DeleteLayerWithOptions(LayerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified parameter.
//
// @param request - DeleteParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParamResponse
func (client *Client) DeleteParamWithOptions(ParamId *string, request *DeleteParamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteParamResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParam"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params/" + dara.PercentEncode(dara.StringValue(ParamId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified parameter.
//
// @param request - DeleteParamRequest
//
// @return DeleteParamResponse
func (client *Client) DeleteParam(ParamId *string, request *DeleteParamRequest) (_result *DeleteParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteParamResponse{}
	_body, _err := client.DeleteParamWithOptions(ParamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a recall management service.
//
// Description:
//
// ## Request description
//
// This operation deletes a recall management service based on the RecallManagementServiceId and InstanceId. Before you call this API, ensure you have the correct information for the service to be deleted.
//
// @param request - DeleteRecallManagementServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecallManagementServiceResponse
func (client *Client) DeleteRecallManagementServiceWithOptions(RecallManagementServiceId *string, request *DeleteRecallManagementServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRecallManagementServiceResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecallManagementService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecallManagementServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a recall management service.
//
// Description:
//
// ## Request description
//
// This operation deletes a recall management service based on the RecallManagementServiceId and InstanceId. Before you call this API, ensure you have the correct information for the service to be deleted.
//
// @param request - DeleteRecallManagementServiceRequest
//
// @return DeleteRecallManagementServiceResponse
func (client *Client) DeleteRecallManagementService(RecallManagementServiceId *string, request *DeleteRecallManagementServiceRequest) (_result *DeleteRecallManagementServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRecallManagementServiceResponse{}
	_body, _err := client.DeleteRecallManagementServiceWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified version of a recall management service.
//
// Description:
//
// ## Request
//
// This operation deletes a specific version of a recall management service. You must provide the recall management service ID, the recall management service version ID, and the instance ID. This operation is irreversible, so back up all critical data before proceeding.
//
// @param request - DeleteRecallManagementServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecallManagementServiceVersionResponse
func (client *Client) DeleteRecallManagementServiceVersionWithOptions(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, request *DeleteRecallManagementServiceVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRecallManagementServiceVersionResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecallManagementServiceVersion"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecallManagementServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified version of a recall management service.
//
// Description:
//
// ## Request
//
// This operation deletes a specific version of a recall management service. You must provide the recall management service ID, the recall management service version ID, and the instance ID. This operation is irreversible, so back up all critical data before proceeding.
//
// @param request - DeleteRecallManagementServiceVersionRequest
//
// @return DeleteRecallManagementServiceVersionResponse
func (client *Client) DeleteRecallManagementServiceVersion(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, request *DeleteRecallManagementServiceVersionRequest) (_result *DeleteRecallManagementServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRecallManagementServiceVersionResponse{}
	_body, _err := client.DeleteRecallManagementServiceVersionWithOptions(RecallManagementServiceId, RecallManagementServiceVersionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified recall management service version configuration.
//
// Description:
//
// ## Usage notes
//
// - Specify the recall management service ID, recall management version ID, and recall management configuration ID.
//
// - `InstanceId` is a required query parameter.
//
// - The request fails if any of the specified IDs are invalid.
//
// @param request - DeleteRecallManagementServiceVersionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecallManagementServiceVersionConfigResponse
func (client *Client) DeleteRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, RecallManagementServiceVersionConfigId *string, request *DeleteRecallManagementServiceVersionConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRecallManagementServiceVersionConfigResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecallManagementServiceVersionConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionId)) + "/configs/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionConfigId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified recall management service version configuration.
//
// Description:
//
// ## Usage notes
//
// - Specify the recall management service ID, recall management version ID, and recall management configuration ID.
//
// - `InstanceId` is a required query parameter.
//
// - The request fails if any of the specified IDs are invalid.
//
// @param request - DeleteRecallManagementServiceVersionConfigRequest
//
// @return DeleteRecallManagementServiceVersionConfigResponse
func (client *Client) DeleteRecallManagementServiceVersionConfig(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, RecallManagementServiceVersionConfigId *string, request *DeleteRecallManagementServiceVersionConfigRequest) (_result *DeleteRecallManagementServiceVersionConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.DeleteRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId, RecallManagementServiceVersionId, RecallManagementServiceVersionConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API deletes a recall management table using the specified recall management table ID and instance ID.
//
// Description:
//
// ## Request
//
// - The required **path parameter*	- `RecallManagementTableId` specifies the ID of the recall management table to delete.
//
// - The required **query parameter*	- `InstanceId` specifies the ID of the instance.
//
// - A successful operation returns a `RequestId` in the response body for request tracking.
//
// @param request - DeleteRecallManagementTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecallManagementTableResponse
func (client *Client) DeleteRecallManagementTableWithOptions(RecallManagementTableId *string, request *DeleteRecallManagementTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRecallManagementTableResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecallManagementTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables/" + dara.PercentEncode(dara.StringValue(RecallManagementTableId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecallManagementTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API deletes a recall management table using the specified recall management table ID and instance ID.
//
// Description:
//
// ## Request
//
// - The required **path parameter*	- `RecallManagementTableId` specifies the ID of the recall management table to delete.
//
// - The required **query parameter*	- `InstanceId` specifies the ID of the instance.
//
// - A successful operation returns a `RequestId` in the response body for request tracking.
//
// @param request - DeleteRecallManagementTableRequest
//
// @return DeleteRecallManagementTableResponse
func (client *Client) DeleteRecallManagementTable(RecallManagementTableId *string, request *DeleteRecallManagementTableRequest) (_result *DeleteRecallManagementTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRecallManagementTableResponse{}
	_body, _err := client.DeleteRecallManagementTableWithOptions(RecallManagementTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源规则
//
// @param request - DeleteResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceRuleResponse
func (client *Client) DeleteResourceRuleWithOptions(ResourceRuleId *string, request *DeleteResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceRuleResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源规则
//
// @param request - DeleteResourceRuleRequest
//
// @return DeleteResourceRuleResponse
func (client *Client) DeleteResourceRule(ResourceRuleId *string, request *DeleteResourceRuleRequest) (_result *DeleteResourceRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceRuleResponse{}
	_body, _err := client.DeleteResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源规则条目
//
// @param request - DeleteResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceRuleItemResponse
func (client *Client) DeleteResourceRuleItemWithOptions(ResourceRuleId *string, ResourceRuleItemId *string, request *DeleteResourceRuleItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceRuleItemResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceRuleItem"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/items/" + dara.PercentEncode(dara.StringValue(ResourceRuleItemId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceRuleItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源规则条目
//
// @param request - DeleteResourceRuleItemRequest
//
// @return DeleteResourceRuleItemResponse
func (client *Client) DeleteResourceRuleItem(ResourceRuleId *string, ResourceRuleItemId *string, request *DeleteResourceRuleItemRequest) (_result *DeleteResourceRuleItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceRuleItemResponse{}
	_body, _err := client.DeleteResourceRuleItemWithOptions(ResourceRuleId, ResourceRuleItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定的样本一致性任务
//
// @param request - DeleteSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSampleConsistencyJobResponse
func (client *Client) DeleteSampleConsistencyJobWithOptions(SampleConsistencyJobId *string, request *DeleteSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSampleConsistencyJobResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSampleConsistencyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定的样本一致性任务
//
// @param request - DeleteSampleConsistencyJobRequest
//
// @return DeleteSampleConsistencyJobResponse
func (client *Client) DeleteSampleConsistencyJob(SampleConsistencyJobId *string, request *DeleteSampleConsistencyJobRequest) (_result *DeleteSampleConsistencyJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSampleConsistencyJobResponse{}
	_body, _err := client.DeleteSampleConsistencyJobWithOptions(SampleConsistencyJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified scenario.
//
// @param request - DeleteSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneResponse
func (client *Client) DeleteSceneWithOptions(SceneId *string, request *DeleteSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSceneResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes/" + dara.PercentEncode(dara.StringValue(SceneId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified scenario.
//
// @param request - DeleteSceneRequest
//
// @return DeleteSceneResponse
func (client *Client) DeleteScene(SceneId *string, request *DeleteSceneRequest) (_result *DeleteSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSceneResponse{}
	_body, _err := client.DeleteSceneWithOptions(SceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified subcrowd.
//
// @param request - DeleteSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubCrowdResponse
func (client *Client) DeleteSubCrowdWithOptions(CrowdId *string, SubCrowdId *string, request *DeleteSubCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSubCrowdResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds/" + dara.PercentEncode(dara.StringValue(SubCrowdId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified subcrowd.
//
// @param request - DeleteSubCrowdRequest
//
// @return DeleteSubCrowdResponse
func (client *Client) DeleteSubCrowd(CrowdId *string, SubCrowdId *string, request *DeleteSubCrowdRequest) (_result *DeleteSubCrowdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSubCrowdResponse{}
	_body, _err := client.DeleteSubCrowdWithOptions(CrowdId, SubCrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data table.
//
// @param request - DeleteTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableMetaResponse
func (client *Client) DeleteTableMetaWithOptions(TableMetaId *string, request *DeleteTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTableMetaResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas/" + dara.PercentEncode(dara.StringValue(TableMetaId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data table.
//
// @param request - DeleteTableMetaRequest
//
// @return DeleteTableMetaResponse
func (client *Client) DeleteTableMeta(TableMetaId *string, request *DeleteTableMetaRequest) (_result *DeleteTableMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTableMetaResponse{}
	_body, _err := client.DeleteTableMetaWithOptions(TableMetaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified traffic control target.
//
// Description:
//
// ## Request
//
// - **TrafficControlTargetId*	- is a required path parameter that specifies the traffic control target to delete.
//
// - **InstanceId*	- is a required query parameter that specifies the instance ID for this operation.
//
// - A successful response includes a `RequestId` field to track the request.
//
// @param request - DeleteTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlTargetResponse
func (client *Client) DeleteTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *DeleteTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrafficControlTargetResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficControlTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified traffic control target.
//
// Description:
//
// ## Request
//
// - **TrafficControlTargetId*	- is a required path parameter that specifies the traffic control target to delete.
//
// - **InstanceId*	- is a required query parameter that specifies the instance ID for this operation.
//
// - A successful response includes a `RequestId` field to track the request.
//
// @param request - DeleteTrafficControlTargetRequest
//
// @return DeleteTrafficControlTargetResponse
func (client *Client) DeleteTrafficControlTarget(TrafficControlTargetId *string, request *DeleteTrafficControlTargetRequest) (_result *DeleteTrafficControlTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrafficControlTargetResponse{}
	_body, _err := client.DeleteTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified traffic control task.
//
// Description:
//
// ## Description
//
// - This API uses `TrafficControlTaskId` and `InstanceId` to delete a traffic control task.
//
// - Ensure the `TrafficControlTaskId` and `InstanceId` are correct, or the operation may fail.
//
// - This operation is irreversible. Proceed with caution.
//
// @param request - DeleteTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlTaskResponse
func (client *Client) DeleteTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *DeleteTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrafficControlTaskResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified traffic control task.
//
// Description:
//
// ## Description
//
// - This API uses `TrafficControlTaskId` and `InstanceId` to delete a traffic control task.
//
// - Ensure the `TrafficControlTaskId` and `InstanceId` are correct, or the operation may fail.
//
// - This operation is irreversible. Proceed with caution.
//
// @param request - DeleteTrafficControlTaskRequest
//
// @return DeleteTrafficControlTaskResponse
func (client *Client) DeleteTrafficControlTask(TrafficControlTaskId *string, request *DeleteTrafficControlTaskRequest) (_result *DeleteTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrafficControlTaskResponse{}
	_body, _err := client.DeleteTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys Flink code for a traffic control task in a specified environment.
//
// Description:
//
// ## Overview
//
// - This API deploys Flink code for a specific traffic control task.
//
// - `TrafficControlTaskId` is a path parameter and requires a valid task ID.
//
// - `InstanceId` and `Environment` are required request body parameters that specify the instance ID and the target deployment environment.
//
// - The optional `RetryDeploy` parameter specifies whether to automatically retry the deployment on failure. The default value is `false`.
//
// - The value for `Environment` must be one of the following: Daily, Pre, or Prod.
//
// @param request - DeployTrafficControlTaskCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployTrafficControlTaskCodeResponse
func (client *Client) DeployTrafficControlTaskCodeWithOptions(TrafficControlTaskId *string, request *DeployTrafficControlTaskCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployTrafficControlTaskCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RetryDeploy) {
		body["RetryDeploy"] = request.RetryDeploy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployTrafficControlTaskCode"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/deploycode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployTrafficControlTaskCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys Flink code for a traffic control task in a specified environment.
//
// Description:
//
// ## Overview
//
// - This API deploys Flink code for a specific traffic control task.
//
// - `TrafficControlTaskId` is a path parameter and requires a valid task ID.
//
// - `InstanceId` and `Environment` are required request body parameters that specify the instance ID and the target deployment environment.
//
// - The optional `RetryDeploy` parameter specifies whether to automatically retry the deployment on failure. The default value is `false`.
//
// - The value for `Environment` must be one of the following: Daily, Pre, or Prod.
//
// @param request - DeployTrafficControlTaskCodeRequest
//
// @return DeployTrafficControlTaskCodeResponse
func (client *Client) DeployTrafficControlTaskCode(TrafficControlTaskId *string, request *DeployTrafficControlTaskCodeRequest) (_result *DeployTrafficControlTaskCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployTrafficControlTaskCodeResponse{}
	_body, _err := client.DeployTrafficControlTaskCodeWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Exports a specified table from the recall engine to a MaxCompute project.
//
// Description:
//
// ## Description
//
// Use this API to export a specific table from the recall engine to Alibaba Cloud MaxCompute for further data processing or analysis. Ensure the provided MaxCompute project name, schema, and table name are valid and that you have the required permissions.
//
// ### Usage notes
//
// - The `Partitions` field must be a JSON object that specifies the table partitions to export.
//
// - The request may fail if any required parameters are missing or incorrect.
//
// - The export process is asynchronous and may take some time. You can use the returned job ID to track the status of the job.
//
// @param request - ExportRecallManagementTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportRecallManagementTableResponse
func (client *Client) ExportRecallManagementTableWithOptions(RecallManagementTableId *string, request *ExportRecallManagementTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportRecallManagementTableResponse, _err error) {
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

	if !dara.IsNil(request.MaxcomputeProjectName) {
		body["MaxcomputeProjectName"] = request.MaxcomputeProjectName
	}

	if !dara.IsNil(request.MaxcomputeSchema) {
		body["MaxcomputeSchema"] = request.MaxcomputeSchema
	}

	if !dara.IsNil(request.MaxcomputeTableName) {
		body["MaxcomputeTableName"] = request.MaxcomputeTableName
	}

	if !dara.IsNil(request.Partitions) {
		body["Partitions"] = request.Partitions
	}

	if !dara.IsNil(request.RecallManagementTableVersionId) {
		body["RecallManagementTableVersionId"] = request.RecallManagementTableVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportRecallManagementTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables/" + dara.PercentEncode(dara.StringValue(RecallManagementTableId)) + "/action/export"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportRecallManagementTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports a specified table from the recall engine to a MaxCompute project.
//
// Description:
//
// ## Description
//
// Use this API to export a specific table from the recall engine to Alibaba Cloud MaxCompute for further data processing or analysis. Ensure the provided MaxCompute project name, schema, and table name are valid and that you have the required permissions.
//
// ### Usage notes
//
// - The `Partitions` field must be a JSON object that specifies the table partitions to export.
//
// - The request may fail if any required parameters are missing or incorrect.
//
// - The export process is asynchronous and may take some time. You can use the returned job ID to track the status of the job.
//
// @param request - ExportRecallManagementTableRequest
//
// @return ExportRecallManagementTableResponse
func (client *Client) ExportRecallManagementTable(RecallManagementTableId *string, request *ExportRecallManagementTableRequest) (_result *ExportRecallManagementTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExportRecallManagementTableResponse{}
	_body, _err := client.ExportRecallManagementTableWithOptions(RecallManagementTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an algorithm customization script.
//
// @param request - GenerateAlgorithmCustomizationScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAlgorithmCustomizationScriptResponse
func (client *Client) GenerateAlgorithmCustomizationScriptWithOptions(AlgorithmCustomizationId *string, request *GenerateAlgorithmCustomizationScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateAlgorithmCustomizationScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeployMode) {
		body["DeployMode"] = request.DeployMode
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleFieldTypes) {
		body["ModuleFieldTypes"] = request.ModuleFieldTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAlgorithmCustomizationScript"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithmcustomizations/" + dara.PercentEncode(dara.StringValue(AlgorithmCustomizationId)) + "/action/generatescript"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAlgorithmCustomizationScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an algorithm customization script.
//
// @param request - GenerateAlgorithmCustomizationScriptRequest
//
// @return GenerateAlgorithmCustomizationScriptResponse
func (client *Client) GenerateAlgorithmCustomizationScript(AlgorithmCustomizationId *string, request *GenerateAlgorithmCustomizationScriptRequest) (_result *GenerateAlgorithmCustomizationScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateAlgorithmCustomizationScriptResponse{}
	_body, _err := client.GenerateAlgorithmCustomizationScriptWithOptions(AlgorithmCustomizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates Flink code for a specified traffic control task ID and instance information.
//
// Description:
//
// ## Description
//
// - This API generates Flink code for a specified traffic control task ID, instance ID, and environment type.
//
// - The `Environment` parameter accepts three values: `Daily` for the daily environment, `Pre` for the pre-release environment, and `Prod` for the production environment.
//
// - Check the `PreNeedConfig` field in the response. A `true` value indicates that necessary configuration information might be missing in the pre-release environment. If this occurs, add or adjust the required settings.
//
// @param request - GenerateTrafficControlTaskCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTrafficControlTaskCodeResponse
func (client *Client) GenerateTrafficControlTaskCodeWithOptions(TrafficControlTaskId *string, request *GenerateTrafficControlTaskCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateTrafficControlTaskCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateTrafficControlTaskCode"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/generatecode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTrafficControlTaskCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates Flink code for a specified traffic control task ID and instance information.
//
// Description:
//
// ## Description
//
// - This API generates Flink code for a specified traffic control task ID, instance ID, and environment type.
//
// - The `Environment` parameter accepts three values: `Daily` for the daily environment, `Pre` for the pre-release environment, and `Prod` for the production environment.
//
// - Check the `PreNeedConfig` field in the response. A `true` value indicates that necessary configuration information might be missing in the pre-release environment. If this occurs, add or adjust the required settings.
//
// @param request - GenerateTrafficControlTaskCodeRequest
//
// @return GenerateTrafficControlTaskCodeResponse
func (client *Client) GenerateTrafficControlTaskCode(TrafficControlTaskId *string, request *GenerateTrafficControlTaskCodeRequest) (_result *GenerateTrafficControlTaskCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateTrafficControlTaskCodeResponse{}
	_body, _err := client.GenerateTrafficControlTaskCodeWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 产生流量调控的相关召回配置
//
// @param request - GenerateTrafficControlTaskConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTrafficControlTaskConfigResponse
func (client *Client) GenerateTrafficControlTaskConfigWithOptions(TrafficControlTaskId *string, request *GenerateTrafficControlTaskConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateTrafficControlTaskConfigResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateTrafficControlTaskConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/generateconfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTrafficControlTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 产生流量调控的相关召回配置
//
// @param request - GenerateTrafficControlTaskConfigRequest
//
// @return GenerateTrafficControlTaskConfigResponse
func (client *Client) GenerateTrafficControlTaskConfig(TrafficControlTaskId *string, request *GenerateTrafficControlTaskConfigRequest) (_result *GenerateTrafficControlTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateTrafficControlTaskConfigResponse{}
	_body, _err := client.GenerateTrafficControlTaskConfigWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of an A/B metric.
//
// @param request - GetABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetABMetricResponse
func (client *Client) GetABMetricWithOptions(ABMetricId *string, request *GetABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetABMetricResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics/" + dara.PercentEncode(dara.StringValue(ABMetricId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of an A/B metric.
//
// @param request - GetABMetricRequest
//
// @return GetABMetricResponse
func (client *Client) GetABMetric(ABMetricId *string, request *GetABMetricRequest) (_result *GetABMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetABMetricResponse{}
	_body, _err := client.GetABMetricWithOptions(ABMetricId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an A/B testing metric group.
//
// @param request - GetABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetABMetricGroupResponse
func (client *Client) GetABMetricGroupWithOptions(ABMetricGroupId *string, request *GetABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetABMetricGroupResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an A/B testing metric group.
//
// @param request - GetABMetricGroupRequest
//
// @return GetABMetricGroupResponse
func (client *Client) GetABMetricGroup(ABMetricGroupId *string, request *GetABMetricGroupRequest) (_result *GetABMetricGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetABMetricGroupResponse{}
	_body, _err := client.GetABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of a specified calculation job.
//
// @param request - GetCalculationJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCalculationJobResponse
func (client *Client) GetCalculationJobWithOptions(CalculationJobId *string, request *GetCalculationJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCalculationJobResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCalculationJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/calculationjobs/" + dara.PercentEncode(dara.StringValue(CalculationJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCalculationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a specified calculation job.
//
// @param request - GetCalculationJobRequest
//
// @return GetCalculationJobResponse
func (client *Client) GetCalculationJob(CalculationJobId *string, request *GetCalculationJobRequest) (_result *GetCalculationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCalculationJobResponse{}
	_body, _err := client.GetCalculationJobWithOptions(CalculationJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves details of a data diagnosis task using its data diagnosis task ID and instance ID.
//
// Description:
//
// ## Request
//
// - This API retrieves the details of a specific data diagnosis task using the provided `DataDiagnosisId` (data diagnosis task configuration ID) and `InstanceId` (instance ID).
//
// - The `CycleTime` field specifies the time for periodic execution. If this field is empty, the task does not execute periodically.
//
// - The value of `Type` determines the content of the `Config` field. For details about the required configuration for each type, see the relevant documentation.
//
// - `GmtCreateTime` and `GmtModifiedTime` are timestamps for the record\\"s creation and modification times, respectively.
//
// @param request - GetDataDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataDiagnosisResponse
func (client *Client) GetDataDiagnosisWithOptions(DataDiagnosisId *string, request *GetDataDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDataDiagnosisResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataDiagnosis"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnoses/" + dara.PercentEncode(dara.StringValue(DataDiagnosisId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details of a data diagnosis task using its data diagnosis task ID and instance ID.
//
// Description:
//
// ## Request
//
// - This API retrieves the details of a specific data diagnosis task using the provided `DataDiagnosisId` (data diagnosis task configuration ID) and `InstanceId` (instance ID).
//
// - The `CycleTime` field specifies the time for periodic execution. If this field is empty, the task does not execute periodically.
//
// - The value of `Type` determines the content of the `Config` field. For details about the required configuration for each type, see the relevant documentation.
//
// - `GmtCreateTime` and `GmtModifiedTime` are timestamps for the record\\"s creation and modification times, respectively.
//
// @param request - GetDataDiagnosisRequest
//
// @return GetDataDiagnosisResponse
func (client *Client) GetDataDiagnosis(DataDiagnosisId *string, request *GetDataDiagnosisRequest) (_result *GetDataDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataDiagnosisResponse{}
	_body, _err := client.GetDataDiagnosisWithOptions(DataDiagnosisId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an engine configuration.
//
// @param request - GetEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEngineConfigResponse
func (client *Client) GetEngineConfigWithOptions(EngineConfigId *string, request *GetEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEngineConfigResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEngineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an engine configuration.
//
// @param request - GetEngineConfigRequest
//
// @return GetEngineConfigResponse
func (client *Client) GetEngineConfig(EngineConfigId *string, request *GetEngineConfigRequest) (_result *GetEngineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEngineConfigResponse{}
	_body, _err := client.GetEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified experiment.
//
// @param request - GetExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithOptions(ExperimentId *string, request *GetExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified experiment.
//
// @param request - GetExperimentRequest
//
// @return GetExperimentResponse
func (client *Client) GetExperiment(ExperimentId *string, request *GetExperimentRequest) (_result *GetExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentResponse{}
	_body, _err := client.GetExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves details for a specified experiment group.
//
// @param request - GetExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentGroupResponse
func (client *Client) GetExperimentGroupWithOptions(ExperimentGroupId *string, request *GetExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentGroupResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details for a specified experiment group.
//
// @param request - GetExperimentGroupRequest
//
// @return GetExperimentGroupResponse
func (client *Client) GetExperimentGroup(ExperimentGroupId *string, request *GetExperimentGroupRequest) (_result *GetExperimentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentGroupResponse{}
	_body, _err := client.GetExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of a feature consistency check job.
//
// @param request - GetFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConsistencyCheckJobResponse
func (client *Client) GetFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId *string, request *GetFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureConsistencyCheckJobResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureConsistencyCheckJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a feature consistency check job.
//
// @param request - GetFeatureConsistencyCheckJobRequest
//
// @return GetFeatureConsistencyCheckJobResponse
func (client *Client) GetFeatureConsistencyCheckJob(FeatureConsistencyCheckJobId *string, request *GetFeatureConsistencyCheckJobRequest) (_result *GetFeatureConsistencyCheckJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureConsistencyCheckJobResponse{}
	_body, _err := client.GetFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration details of a feature consistency check task.
//
// @param request - GetFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConsistencyCheckJobConfigResponse
func (client *Client) GetFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId *string, request *GetFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureConsistencyCheckJobConfigResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration details of a feature consistency check task.
//
// @param request - GetFeatureConsistencyCheckJobConfigRequest
//
// @return GetFeatureConsistencyCheckJobConfigResponse
func (client *Client) GetFeatureConsistencyCheckJobConfig(FeatureConsistencyCheckJobConfigId *string, request *GetFeatureConsistencyCheckJobConfigRequest) (_result *GetFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.GetFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of a specified PAI-REC instance.
//
// @param request - GetInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(InstanceId *string, request *GetInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a specified PAI-REC instance.
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(InstanceId *string, request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific resource in a specified instance.
//
// @param request - GetInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResourceResponse
func (client *Client) GetInstanceResourceWithOptions(InstanceId *string, ResourceId *string, request *GetInstanceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResourceResponse, _err error) {
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
		Action:      dara.String("GetInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific resource in a specified instance.
//
// @param request - GetInstanceResourceRequest
//
// @return GetInstanceResourceResponse
func (client *Client) GetInstanceResource(InstanceId *string, ResourceId *string, request *GetInstanceResourceRequest) (_result *GetInstanceResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResourceResponse{}
	_body, _err := client.GetInstanceResourceWithOptions(InstanceId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of data tables under a data source.
//
// @param request - GetInstanceResourceTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResourceTableResponse
func (client *Client) GetInstanceResourceTableWithOptions(InstanceId *string, ResourceId *string, TableName *string, request *GetInstanceResourceTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResourceTableResponse, _err error) {
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
		Action:      dara.String("GetInstanceResourceTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(TableName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResourceTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of data tables under a data source.
//
// @param request - GetInstanceResourceTableRequest
//
// @return GetInstanceResourceTableResponse
func (client *Client) GetInstanceResourceTable(InstanceId *string, ResourceId *string, TableName *string, request *GetInstanceResourceTableRequest) (_result *GetInstanceResourceTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResourceTableResponse{}
	_body, _err := client.GetInstanceResourceTableWithOptions(InstanceId, ResourceId, TableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified laboratory.
//
// @param request - GetLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLaboratoryResponse
func (client *Client) GetLaboratoryWithOptions(LaboratoryId *string, request *GetLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLaboratoryResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified laboratory.
//
// @param request - GetLaboratoryRequest
//
// @return GetLaboratoryResponse
func (client *Client) GetLaboratory(LaboratoryId *string, request *GetLaboratoryRequest) (_result *GetLaboratoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLaboratoryResponse{}
	_body, _err := client.GetLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified layer.
//
// @param request - GetLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerResponse
func (client *Client) GetLayerWithOptions(LayerId *string, request *GetLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLayerResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers/" + dara.PercentEncode(dara.StringValue(LayerId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified layer.
//
// @param request - GetLayerRequest
//
// @return GetLayerResponse
func (client *Client) GetLayer(LayerId *string, request *GetLayerRequest) (_result *GetLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLayerResponse{}
	_body, _err := client.GetLayerWithOptions(LayerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the recall management configuration.
//
// @param request - GetRecallManagementConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecallManagementConfigResponse
func (client *Client) GetRecallManagementConfigWithOptions(request *GetRecallManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRecallManagementConfigResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecallManagementConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementconfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecallManagementConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the recall management configuration.
//
// @param request - GetRecallManagementConfigRequest
//
// @return GetRecallManagementConfigResponse
func (client *Client) GetRecallManagementConfig(request *GetRecallManagementConfigRequest) (_result *GetRecallManagementConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRecallManagementConfigResponse{}
	_body, _err := client.GetRecallManagementConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific recall management job, including its status and log.
//
// Description:
//
// ## Description
//
// Retrieves the details of a specific recall management job using its `RecallManagementJobId` and `InstanceId`. The response includes the job\\"s status (such as Init, Running, Success, or Failed), start and end times, related table information, and operation log. To make a request, specify the `RecallManagementJobId` as a path parameter and the `InstanceId` as a query parameter.
//
// @param request - GetRecallManagementJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecallManagementJobResponse
func (client *Client) GetRecallManagementJobWithOptions(RecallManagementJobId *string, request *GetRecallManagementJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRecallManagementJobResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecallManagementJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementjobs/" + dara.PercentEncode(dara.StringValue(RecallManagementJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecallManagementJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific recall management job, including its status and log.
//
// Description:
//
// ## Description
//
// Retrieves the details of a specific recall management job using its `RecallManagementJobId` and `InstanceId`. The response includes the job\\"s status (such as Init, Running, Success, or Failed), start and end times, related table information, and operation log. To make a request, specify the `RecallManagementJobId` as a path parameter and the `InstanceId` as a query parameter.
//
// @param request - GetRecallManagementJobRequest
//
// @return GetRecallManagementJobResponse
func (client *Client) GetRecallManagementJob(RecallManagementJobId *string, request *GetRecallManagementJobRequest) (_result *GetRecallManagementJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRecallManagementJobResponse{}
	_body, _err := client.GetRecallManagementJobWithOptions(RecallManagementJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified recall management service, including its status and version.
//
// Description:
//
// ## Request
//
// @param request - GetRecallManagementServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecallManagementServiceResponse
func (client *Client) GetRecallManagementServiceWithOptions(RecallManagementServiceId *string, request *GetRecallManagementServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRecallManagementServiceResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecallManagementService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecallManagementServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified recall management service, including its status and version.
//
// Description:
//
// ## Request
//
// @param request - GetRecallManagementServiceRequest
//
// @return GetRecallManagementServiceResponse
func (client *Client) GetRecallManagementService(RecallManagementServiceId *string, request *GetRecallManagementServiceRequest) (_result *GetRecallManagementServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRecallManagementServiceResponse{}
	_body, _err := client.GetRecallManagementServiceWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the version details of the recall management service.
//
// @param request - GetRecallManagementServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecallManagementServiceVersionResponse
func (client *Client) GetRecallManagementServiceVersionWithOptions(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, request *GetRecallManagementServiceVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRecallManagementServiceVersionResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecallManagementServiceVersion"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecallManagementServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the version details of the recall management service.
//
// @param request - GetRecallManagementServiceVersionRequest
//
// @return GetRecallManagementServiceVersionResponse
func (client *Client) GetRecallManagementServiceVersion(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, request *GetRecallManagementServiceVersionRequest) (_result *GetRecallManagementServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRecallManagementServiceVersionResponse{}
	_body, _err := client.GetRecallManagementServiceVersionWithOptions(RecallManagementServiceId, RecallManagementServiceVersionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API retrieves the configuration details of a specific recall management service version.
//
// Description:
//
// ## Request
//
// Retrieves the configuration details for a specific version of a Recall Management Service by specifying its service, version, and configuration IDs. Ensure the parameter values are correct. The `InstanceId` is crucial for locating the correct instance.
//
// - **Note**: All path parameters (`RecallManagementServiceId`, `RecallManagementServiceVersionId`, and `RecallManagementServiceVersionConfigId`) are required and must reference an existing resource.
//
// - **Extended configuration**: The response includes the `ExtendedConfig` field, which is used for future extensions and custom settings. Parse this field as needed.
//
// @param request - GetRecallManagementServiceVersionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecallManagementServiceVersionConfigResponse
func (client *Client) GetRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, RecallManagementServiceVersionConfigId *string, request *GetRecallManagementServiceVersionConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRecallManagementServiceVersionConfigResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecallManagementServiceVersionConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionId)) + "/configs/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API retrieves the configuration details of a specific recall management service version.
//
// Description:
//
// ## Request
//
// Retrieves the configuration details for a specific version of a Recall Management Service by specifying its service, version, and configuration IDs. Ensure the parameter values are correct. The `InstanceId` is crucial for locating the correct instance.
//
// - **Note**: All path parameters (`RecallManagementServiceId`, `RecallManagementServiceVersionId`, and `RecallManagementServiceVersionConfigId`) are required and must reference an existing resource.
//
// - **Extended configuration**: The response includes the `ExtendedConfig` field, which is used for future extensions and custom settings. Parse this field as needed.
//
// @param request - GetRecallManagementServiceVersionConfigRequest
//
// @return GetRecallManagementServiceVersionConfigResponse
func (client *Client) GetRecallManagementServiceVersionConfig(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, RecallManagementServiceVersionConfigId *string, request *GetRecallManagementServiceVersionConfigRequest) (_result *GetRecallManagementServiceVersionConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.GetRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId, RecallManagementServiceVersionId, RecallManagementServiceVersionConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified recall management table, including its table structure and configuration.
//
// Description:
//
// ## Request
//
// This API retrieves details of a specific recall management table using the provided `RecallManagementTableId` and `InstanceId`. Ensure you provide the correct values for these parameters.
//
// - **Note**: The `CanDelete` field indicates whether the data table supports delete operations. Use this value to control delete functionality in your application.
//
// - The `Fields` list contains the definitions for each field in the data table, including their name, type, and properties.
//
// - To monitor data changes, you can configure or query the fluctuation thresholds for row count and size using the corresponding fields.
//
// @param request - GetRecallManagementTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecallManagementTableResponse
func (client *Client) GetRecallManagementTableWithOptions(RecallManagementTableId *string, request *GetRecallManagementTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRecallManagementTableResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecallManagementTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables/" + dara.PercentEncode(dara.StringValue(RecallManagementTableId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecallManagementTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified recall management table, including its table structure and configuration.
//
// Description:
//
// ## Request
//
// This API retrieves details of a specific recall management table using the provided `RecallManagementTableId` and `InstanceId`. Ensure you provide the correct values for these parameters.
//
// - **Note**: The `CanDelete` field indicates whether the data table supports delete operations. Use this value to control delete functionality in your application.
//
// - The `Fields` list contains the definitions for each field in the data table, including their name, type, and properties.
//
// - To monitor data changes, you can configure or query the fluctuation thresholds for row count and size using the corresponding fields.
//
// @param request - GetRecallManagementTableRequest
//
// @return GetRecallManagementTableResponse
func (client *Client) GetRecallManagementTable(RecallManagementTableId *string, request *GetRecallManagementTableRequest) (_result *GetRecallManagementTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRecallManagementTableResponse{}
	_body, _err := client.GetRecallManagementTableWithOptions(RecallManagementTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源规则详细信息
//
// @param request - GetResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceRuleResponse
func (client *Client) GetResourceRuleWithOptions(ResourceRuleId *string, request *GetResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceRuleResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源规则详细信息
//
// @param request - GetResourceRuleRequest
//
// @return GetResourceRuleResponse
func (client *Client) GetResourceRule(ResourceRuleId *string, request *GetResourceRuleRequest) (_result *GetResourceRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceRuleResponse{}
	_body, _err := client.GetResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取样本一致性任务详细信息
//
// @param request - GetSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSampleConsistencyJobResponse
func (client *Client) GetSampleConsistencyJobWithOptions(SampleConsistencyJobId *string, request *GetSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSampleConsistencyJobResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSampleConsistencyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取样本一致性任务详细信息
//
// @param request - GetSampleConsistencyJobRequest
//
// @return GetSampleConsistencyJobResponse
func (client *Client) GetSampleConsistencyJob(SampleConsistencyJobId *string, request *GetSampleConsistencyJobRequest) (_result *GetSampleConsistencyJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSampleConsistencyJobResponse{}
	_body, _err := client.GetSampleConsistencyJobWithOptions(SampleConsistencyJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified scene.
//
// @param request - GetSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneResponse
func (client *Client) GetSceneWithOptions(SceneId *string, request *GetSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSceneResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes/" + dara.PercentEncode(dara.StringValue(SceneId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified scene.
//
// @param request - GetSceneRequest
//
// @return GetSceneResponse
func (client *Client) GetScene(SceneId *string, request *GetSceneRequest) (_result *GetSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSceneResponse{}
	_body, _err := client.GetSceneWithOptions(SceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a service.
//
// @param request - GetServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(ServiceId *string, request *GetServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/services/" + dara.PercentEncode(dara.StringValue(ServiceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a service.
//
// @param request - GetServiceRequest
//
// @return GetServiceResponse
func (client *Client) GetService(ServiceId *string, request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(ServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a sub-crowd.
//
// @param request - GetSubCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubCrowdResponse
func (client *Client) GetSubCrowdWithOptions(CrowdId *string, SubCrowdId *string, request *GetSubCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSubCrowdResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds/" + dara.PercentEncode(dara.StringValue(SubCrowdId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a sub-crowd.
//
// @param request - GetSubCrowdRequest
//
// @return GetSubCrowdResponse
func (client *Client) GetSubCrowd(CrowdId *string, SubCrowdId *string, request *GetSubCrowdRequest) (_result *GetSubCrowdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSubCrowdResponse{}
	_body, _err := client.GetSubCrowdWithOptions(CrowdId, SubCrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns the details of a table.
//
// @param request - GetTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableMetaResponse
func (client *Client) GetTableMetaWithOptions(TableMetaId *string, request *GetTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableMetaResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas/" + dara.PercentEncode(dara.StringValue(TableMetaId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the details of a table.
//
// @param request - GetTableMetaRequest
//
// @return GetTableMetaResponse
func (client *Client) GetTableMeta(TableMetaId *string, request *GetTableMetaRequest) (_result *GetTableMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableMetaResponse{}
	_body, _err := client.GetTableMetaWithOptions(TableMetaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of a traffic throttling objective by its ID.
//
// Description:
//
// ## Request
//
// Retrieves the detailed configuration of a traffic control target. The configuration includes the target name, time range, and condition settings. This operation requires the `TrafficControlTargetId` and `InstanceId` parameters.
//
// @param request - GetTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTargetResponse
func (client *Client) GetTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *GetTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrafficControlTargetResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrafficControlTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a traffic throttling objective by its ID.
//
// Description:
//
// ## Request
//
// Retrieves the detailed configuration of a traffic control target. The configuration includes the target name, time range, and condition settings. This operation requires the `TrafficControlTargetId` and `InstanceId` parameters.
//
// @param request - GetTrafficControlTargetRequest
//
// @return GetTrafficControlTargetResponse
func (client *Client) GetTrafficControlTarget(TrafficControlTargetId *string, request *GetTrafficControlTargetRequest) (_result *GetTrafficControlTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrafficControlTargetResponse{}
	_body, _err := client.GetTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a traffic control task with a specified ID.
//
// Description:
//
// ## Description
//
// - This operation retrieves the details of a specific traffic control task, including but not limited to the task name, description, and status.
//
// - TrafficControlTaskId and InstanceId are required parameters that specify the task ID and instance ID to query.
//
// - Optional parameters such as Environment, Version, and ControlTargetFilter help refine the request to retrieve more specific task data or version information.
//
// - Check the returned data structure, especially the TrafficControlTargets section, which contains multiple control targets and their related properties.
//
// @param request - GetTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTaskResponse
func (client *Client) GetTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *GetTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrafficControlTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlTargetFilter) {
		query["ControlTargetFilter"] = request.ControlTargetFilter
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a traffic control task with a specified ID.
//
// Description:
//
// ## Description
//
// - This operation retrieves the details of a specific traffic control task, including but not limited to the task name, description, and status.
//
// - TrafficControlTaskId and InstanceId are required parameters that specify the task ID and instance ID to query.
//
// - Optional parameters such as Environment, Version, and ControlTargetFilter help refine the request to retrieve more specific task data or version information.
//
// - Check the returned data structure, especially the TrafficControlTargets section, which contains multiple control targets and their related properties.
//
// @param request - GetTrafficControlTaskRequest
//
// @return GetTrafficControlTaskResponse
func (client *Client) GetTrafficControlTask(TrafficControlTaskId *string, request *GetTrafficControlTaskRequest) (_result *GetTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrafficControlTaskResponse{}
	_body, _err := client.GetTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves traffic allocation details for a specific traffic control task.
//
// Description:
//
// ## Description
//
// This API retrieves the traffic details for a specific traffic control task. The request must include the `TrafficControlTaskId`, `InstanceId`, and `Environment`.
//
// - `TrafficControlTaskId`: The unique identifier for the traffic control task.
//
// - `InstanceId`: The instance ID.
//
// - `Environment`: The environment type, such as the production environment (Prod).
//
// The response includes the traffic allocation for each experiment and traffic control target. This data allows you to analyze and manage your traffic control strategies.
//
// @param request - GetTrafficControlTaskTrafficRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrafficControlTaskTrafficResponse
func (client *Client) GetTrafficControlTaskTrafficWithOptions(TrafficControlTaskId *string, request *GetTrafficControlTaskTrafficRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrafficControlTaskTrafficResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrafficControlTaskTraffic"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/trafficinfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrafficControlTaskTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves traffic allocation details for a specific traffic control task.
//
// Description:
//
// ## Description
//
// This API retrieves the traffic details for a specific traffic control task. The request must include the `TrafficControlTaskId`, `InstanceId`, and `Environment`.
//
// - `TrafficControlTaskId`: The unique identifier for the traffic control task.
//
// - `InstanceId`: The instance ID.
//
// - `Environment`: The environment type, such as the production environment (Prod).
//
// The response includes the traffic allocation for each experiment and traffic control target. This data allows you to analyze and manage your traffic control strategies.
//
// @param request - GetTrafficControlTaskTrafficRequest
//
// @return GetTrafficControlTaskTrafficResponse
func (client *Client) GetTrafficControlTaskTraffic(TrafficControlTaskId *string, request *GetTrafficControlTaskTrafficRequest) (_result *GetTrafficControlTaskTrafficResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrafficControlTaskTrafficResponse{}
	_body, _err := client.GetTrafficControlTaskTrafficWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of A/B metric groups.
//
// @param request - ListABMetricGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABMetricGroupsResponse
func (client *Client) ListABMetricGroupsWithOptions(request *ListABMetricGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABMetricGroupsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Realtime) {
		query["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABMetricGroups"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABMetricGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of A/B metric groups.
//
// @param request - ListABMetricGroupsRequest
//
// @return ListABMetricGroupsResponse
func (client *Client) ListABMetricGroups(request *ListABMetricGroupsRequest) (_result *ListABMetricGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABMetricGroupsResponse{}
	_body, _err := client.ListABMetricGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists A/B testing metrics.
//
// @param request - ListABMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABMetricsResponse
func (client *Client) ListABMetricsWithOptions(request *ListABMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABMetricsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Realtime) {
		query["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.TableMetaId) {
		query["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABMetrics"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists A/B testing metrics.
//
// @param request - ListABMetricsRequest
//
// @return ListABMetricsResponse
func (client *Client) ListABMetrics(request *ListABMetricsRequest) (_result *ListABMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABMetricsResponse{}
	_body, _err := client.ListABMetricsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists calculation jobs.
//
// @param request - ListCalculationJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalculationJobsResponse
func (client *Client) ListCalculationJobsWithOptions(request *ListCalculationJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCalculationJobsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCalculationJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/calculationjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCalculationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists calculation jobs.
//
// @param request - ListCalculationJobsRequest
//
// @return ListCalculationJobsResponse
func (client *Client) ListCalculationJobs(request *ListCalculationJobsRequest) (_result *ListCalculationJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCalculationJobsResponse{}
	_body, _err := client.ListCalculationJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all users within a specified crowd, including those in its sub-crowds.
//
// @param request - ListCrowdUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrowdUsersResponse
func (client *Client) ListCrowdUsersWithOptions(CrowdId *string, request *ListCrowdUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCrowdUsersResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrowdUsers"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrowdUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all users within a specified crowd, including those in its sub-crowds.
//
// @param request - ListCrowdUsersRequest
//
// @return ListCrowdUsersResponse
func (client *Client) ListCrowdUsers(CrowdId *string, request *ListCrowdUsersRequest) (_result *ListCrowdUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCrowdUsersResponse{}
	_body, _err := client.ListCrowdUsersWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the Crowds in a specified instance.
//
// @param request - ListCrowdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrowdsResponse
func (client *Client) ListCrowdsWithOptions(request *ListCrowdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCrowdsResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrowds"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrowdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the Crowds in a specified instance.
//
// @param request - ListCrowdsRequest
//
// @return ListCrowdsResponse
func (client *Client) ListCrowds(request *ListCrowdsRequest) (_result *ListCrowdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCrowdsResponse{}
	_body, _err := client.ListCrowdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the data diagnostic tasks for the specified instance.
//
// Description:
//
// ## Request
//
// This API retrieves a list of data diagnosis tasks. It requires the `InstanceId` parameter and accepts optional parameters—such as data diagnosis type, page number, and page size—for filtering and pagination.
//
// - The **Types*	- parameter accepts multiple data diagnosis types, allowing you to view reports for all selected types at once.
//
// - To paginate results, use the `PageNumber` and `PageSize` parameters.
//
// @param tmpReq - ListDataDiagnosesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDiagnosesResponse
func (client *Client) ListDataDiagnosesWithOptions(tmpReq *ListDataDiagnosesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataDiagnosesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataDiagnosesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TypesShrink) {
		query["Types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataDiagnoses"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnoses"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataDiagnosesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the data diagnostic tasks for the specified instance.
//
// Description:
//
// ## Request
//
// This API retrieves a list of data diagnosis tasks. It requires the `InstanceId` parameter and accepts optional parameters—such as data diagnosis type, page number, and page size—for filtering and pagination.
//
// - The **Types*	- parameter accepts multiple data diagnosis types, allowing you to view reports for all selected types at once.
//
// - To paginate results, use the `PageNumber` and `PageSize` parameters.
//
// @param request - ListDataDiagnosesRequest
//
// @return ListDataDiagnosesResponse
func (client *Client) ListDataDiagnoses(request *ListDataDiagnosesRequest) (_result *ListDataDiagnosesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataDiagnosesResponse{}
	_body, _err := client.ListDataDiagnosesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data diagnosis jobs for a specified instance.
//
// Description:
//
// ## Description
//
// This API retrieves data diagnosis jobs for a specific instance. You can filter the jobs by parameters such as `InstanceId` and `Status`, and use the `PageNumber` and `PageSize` parameters to control pagination.
//
// - The **InstanceId*	- parameter is required and specifies the instance to query.
//
// - Optional parameters include **Status**, **Types**, **PageNumber**, and **PageSize**.
//
// - Note: If you omit filter conditions, the operation returns all matching data diagnosis jobs.
//
// @param tmpReq - ListDataDiagnosisJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDiagnosisJobsResponse
func (client *Client) ListDataDiagnosisJobsWithOptions(tmpReq *ListDataDiagnosisJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataDiagnosisJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataDiagnosisJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TypesShrink) {
		query["Types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataDiagnosisJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnosisjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataDiagnosisJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data diagnosis jobs for a specified instance.
//
// Description:
//
// ## Description
//
// This API retrieves data diagnosis jobs for a specific instance. You can filter the jobs by parameters such as `InstanceId` and `Status`, and use the `PageNumber` and `PageSize` parameters to control pagination.
//
// - The **InstanceId*	- parameter is required and specifies the instance to query.
//
// - Optional parameters include **Status**, **Types**, **PageNumber**, and **PageSize**.
//
// - Note: If you omit filter conditions, the operation returns all matching data diagnosis jobs.
//
// @param request - ListDataDiagnosisJobsRequest
//
// @return ListDataDiagnosisJobsResponse
func (client *Client) ListDataDiagnosisJobs(request *ListDataDiagnosisJobsRequest) (_result *ListDataDiagnosisJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataDiagnosisJobsResponse{}
	_body, _err := client.ListDataDiagnosisJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specify parameters to retrieve data diagnosis reports for a specific time range.
//
// Description:
//
// ## Request
//
// - This API retrieves data diagnosis reports based on parameters such as the data diagnosis ID, instance ID, and a date range.
//
// - The `FeatureName` parameter filters reports by a specific feature, and the `TopN` parameter limits the number of results.
//
// - The `RemainRateType` parameter specifies the type of retention rate report, such as a periodic report.
//
// - Report content includes item and user change rate analysis, periodic user preference analysis, correlation analysis, basic statistical analysis, and anomaly detection.
//
// - Important: All date values must use the `YYYY-MM-DD` format.
//
// @param request - ListDataDiagnosisReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDiagnosisReportsResponse
func (client *Client) ListDataDiagnosisReportsWithOptions(DataDiagnosisId *string, request *ListDataDiagnosisReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataDiagnosisReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RemainRateType) {
		query["RemainRateType"] = request.RemainRateType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataDiagnosisReports"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnoses/" + dara.PercentEncode(dara.StringValue(DataDiagnosisId)) + "/reports"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataDiagnosisReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specify parameters to retrieve data diagnosis reports for a specific time range.
//
// Description:
//
// ## Request
//
// - This API retrieves data diagnosis reports based on parameters such as the data diagnosis ID, instance ID, and a date range.
//
// - The `FeatureName` parameter filters reports by a specific feature, and the `TopN` parameter limits the number of results.
//
// - The `RemainRateType` parameter specifies the type of retention rate report, such as a periodic report.
//
// - Report content includes item and user change rate analysis, periodic user preference analysis, correlation analysis, basic statistical analysis, and anomaly detection.
//
// - Important: All date values must use the `YYYY-MM-DD` format.
//
// @param request - ListDataDiagnosisReportsRequest
//
// @return ListDataDiagnosisReportsResponse
func (client *Client) ListDataDiagnosisReports(DataDiagnosisId *string, request *ListDataDiagnosisReportsRequest) (_result *ListDataDiagnosisReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataDiagnosisReportsResponse{}
	_body, _err := client.ListDataDiagnosisReportsWithOptions(DataDiagnosisId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of engine configurations.
//
// @param request - ListEngineConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEngineConfigsResponse
func (client *Client) ListEngineConfigsWithOptions(request *ListEngineConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEngineConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEngineConfigs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEngineConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of engine configurations.
//
// @param request - ListEngineConfigsRequest
//
// @return ListEngineConfigsResponse
func (client *Client) ListEngineConfigs(request *ListEngineConfigsRequest) (_result *ListEngineConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEngineConfigsResponse{}
	_body, _err := client.ListEngineConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns a list of experiment groups.
//
// @param request - ListExperimentGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentGroupsResponse
func (client *Client) ListExperimentGroupsWithOptions(request *ListExperimentGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExperimentGroupsResponse, _err error) {
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

	if !dara.IsNil(request.LayerId) {
		query["LayerId"] = request.LayerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeRangeEnd) {
		query["TimeRangeEnd"] = request.TimeRangeEnd
	}

	if !dara.IsNil(request.TimeRangeStart) {
		query["TimeRangeStart"] = request.TimeRangeStart
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExperimentGroups"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExperimentGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a list of experiment groups.
//
// @param request - ListExperimentGroupsRequest
//
// @return ListExperimentGroupsResponse
func (client *Client) ListExperimentGroups(request *ListExperimentGroupsRequest) (_result *ListExperimentGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentGroupsResponse{}
	_body, _err := client.ListExperimentGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the experiments in the specified experiment group.
//
// @param request - ListExperimentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentsResponse
func (client *Client) ListExperimentsWithOptions(request *ListExperimentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExperimentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExperimentGroupId) {
		query["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExperiments"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the experiments in the specified experiment group.
//
// @param request - ListExperimentsRequest
//
// @return ListExperimentsResponse
func (client *Client) ListExperiments(request *ListExperimentsRequest) (_result *ListExperimentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentsResponse{}
	_body, _err := client.ListExperimentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of feature consistency check task configurations.
//
// @param request - ListFeatureConsistencyCheckJobConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobConfigsResponse
func (client *Client) ListFeatureConsistencyCheckJobConfigsWithOptions(request *ListFeatureConsistencyCheckJobConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobConfigsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobConfigs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of feature consistency check task configurations.
//
// @param request - ListFeatureConsistencyCheckJobConfigsRequest
//
// @return ListFeatureConsistencyCheckJobConfigsResponse
func (client *Client) ListFeatureConsistencyCheckJobConfigs(request *ListFeatureConsistencyCheckJobConfigsRequest) (_result *ListFeatureConsistencyCheckJobConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobConfigsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists feature reports for a feature consistency check job.
//
// @param request - ListFeatureConsistencyCheckJobFeatureReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobFeatureReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobFeatureReportsWithOptions(FeatureConsistencyCheckJobId *string, request *ListFeatureConsistencyCheckJobFeatureReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobFeatureReportsResponse, _err error) {
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

	if !dara.IsNil(request.LogItemId) {
		query["LogItemId"] = request.LogItemId
	}

	if !dara.IsNil(request.LogRequestId) {
		query["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.LogUserId) {
		query["LogUserId"] = request.LogUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobFeatureReports"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId)) + "/featurereports"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobFeatureReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists feature reports for a feature consistency check job.
//
// @param request - ListFeatureConsistencyCheckJobFeatureReportsRequest
//
// @return ListFeatureConsistencyCheckJobFeatureReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobFeatureReports(FeatureConsistencyCheckJobId *string, request *ListFeatureConsistencyCheckJobFeatureReportsRequest) (_result *ListFeatureConsistencyCheckJobFeatureReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobFeatureReportsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobFeatureReportsWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the score difference reports for a feature consistency check job.
//
// @param tmpReq - ListFeatureConsistencyCheckJobScoreReportsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobScoreReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobScoreReportsWithOptions(FeatureConsistencyCheckJobId *string, tmpReq *ListFeatureConsistencyCheckJobScoreReportsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobScoreReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFeatureConsistencyCheckJobScoreReportsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExcludeRequestIds) {
		request.ExcludeRequestIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeRequestIds, dara.String("ExcludeRequestIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeRequestIdsShrink) {
		query["ExcludeRequestIds"] = request.ExcludeRequestIdsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobScoreReports"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId)) + "/scorereports"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobScoreReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the score difference reports for a feature consistency check job.
//
// @param request - ListFeatureConsistencyCheckJobScoreReportsRequest
//
// @return ListFeatureConsistencyCheckJobScoreReportsResponse
func (client *Client) ListFeatureConsistencyCheckJobScoreReports(FeatureConsistencyCheckJobId *string, request *ListFeatureConsistencyCheckJobScoreReportsRequest) (_result *ListFeatureConsistencyCheckJobScoreReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobScoreReportsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobScoreReportsWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of feature consistency check jobs.
//
// @param request - ListFeatureConsistencyCheckJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureConsistencyCheckJobsResponse
func (client *Client) ListFeatureConsistencyCheckJobsWithOptions(request *ListFeatureConsistencyCheckJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureConsistencyCheckJobsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureConsistencyCheckJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureConsistencyCheckJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of feature consistency check jobs.
//
// @param request - ListFeatureConsistencyCheckJobsRequest
//
// @return ListFeatureConsistencyCheckJobsResponse
func (client *Client) ListFeatureConsistencyCheckJobs(request *ListFeatureConsistencyCheckJobsRequest) (_result *ListFeatureConsistencyCheckJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureConsistencyCheckJobsResponse{}
	_body, _err := client.ListFeatureConsistencyCheckJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all schemas for a specified resource in an instance.
//
// Description:
//
// ## Description
//
// To retrieve a list of all schemas for a specified resource, provide the instance ID (InstanceId) and resource ID (ResourceId). Use the optional SchemaName parameter to filter the schemas by a name prefix.
//
// @param request - ListInstanceResourceSchemasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResourceSchemasResponse
func (client *Client) ListInstanceResourceSchemasWithOptions(InstanceId *string, ResourceId *string, request *ListInstanceResourceSchemasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceResourceSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceResourceSchemas"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/schemas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResourceSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all schemas for a specified resource in an instance.
//
// Description:
//
// ## Description
//
// To retrieve a list of all schemas for a specified resource, provide the instance ID (InstanceId) and resource ID (ResourceId). Use the optional SchemaName parameter to filter the schemas by a name prefix.
//
// @param request - ListInstanceResourceSchemasRequest
//
// @return ListInstanceResourceSchemasResponse
func (client *Client) ListInstanceResourceSchemas(InstanceId *string, ResourceId *string, request *ListInstanceResourceSchemasRequest) (_result *ListInstanceResourceSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResourceSchemasResponse{}
	_body, _err := client.ListInstanceResourceSchemasWithOptions(InstanceId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of data tables for a specified instance and data source.
//
// Description:
//
// ## Description
//
// Provide the instance ID (InstanceId) and data source ID (ResourceId) to retrieve a list of data tables from the specified data source. Use the optional `MaxcomputeSchema` parameter to filter the results by a MaxCompute schema.
//
// @param request - ListInstanceResourceTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResourceTablesResponse
func (client *Client) ListInstanceResourceTablesWithOptions(InstanceId *string, ResourceId *string, request *ListInstanceResourceTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceResourceTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxcomputeSchema) {
		query["MaxcomputeSchema"] = request.MaxcomputeSchema
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceResourceTables"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResourceTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data tables for a specified instance and data source.
//
// Description:
//
// ## Description
//
// Provide the instance ID (InstanceId) and data source ID (ResourceId) to retrieve a list of data tables from the specified data source. Use the optional `MaxcomputeSchema` parameter to filter the results by a MaxCompute schema.
//
// @param request - ListInstanceResourceTablesRequest
//
// @return ListInstanceResourceTablesResponse
func (client *Client) ListInstanceResourceTables(InstanceId *string, ResourceId *string, request *ListInstanceResourceTablesRequest) (_result *ListInstanceResourceTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResourceTablesResponse{}
	_body, _err := client.ListInstanceResourceTablesWithOptions(InstanceId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the resources configured for an instance.
//
// @param request - ListInstanceResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResourcesResponse
func (client *Client) ListInstanceResourcesWithOptions(InstanceId *string, request *ListInstanceResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceResources"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the resources configured for an instance.
//
// @param request - ListInstanceResourcesRequest
//
// @return ListInstanceResourcesResponse
func (client *Client) ListInstanceResources(InstanceId *string, request *ListInstanceResourcesRequest) (_result *ListInstanceResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResourcesResponse{}
	_body, _err := client.ListInstanceResourcesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets a list of PAIRec instances.
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances"),
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
// Gets a list of PAIRec instances.
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
// Retrieves the laboratories in a specified scene.
//
// @param request - ListLaboratoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLaboratoriesResponse
func (client *Client) ListLaboratoriesWithOptions(request *ListLaboratoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLaboratoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLaboratories"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLaboratoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the laboratories in a specified scene.
//
// @param request - ListLaboratoriesRequest
//
// @return ListLaboratoriesResponse
func (client *Client) ListLaboratories(request *ListLaboratoriesRequest) (_result *ListLaboratoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLaboratoriesResponse{}
	_body, _err := client.ListLaboratoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of layers in a specified laboratory.
//
// @param request - ListLayersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayersResponse
func (client *Client) ListLayersWithOptions(request *ListLayersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLayersResponse, _err error) {
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

	if !dara.IsNil(request.LaboratoryId) {
		query["LaboratoryId"] = request.LaboratoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLayers"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLayersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of layers in a specified laboratory.
//
// @param request - ListLayersRequest
//
// @return ListLayersResponse
func (client *Client) ListLayers(request *ListLayersRequest) (_result *ListLayersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLayersResponse{}
	_body, _err := client.ListLayersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists parameters.
//
// @param request - ListParamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParamsResponse
func (client *Client) ListParamsWithOptions(request *ListParamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListParamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Encrypted) {
		query["Encrypted"] = request.Encrypted
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListParams"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists parameters.
//
// @param request - ListParamsRequest
//
// @return ListParamsResponse
func (client *Client) ListParams(request *ListParamsRequest) (_result *ListParamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListParamsResponse{}
	_body, _err := client.ListParamsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of recall management tasks that match specified conditions.
//
// Description:
//
// ## Request
//
// - Use this API operation to retrieve a list of recall management tasks.
//
// - The `InstanceId` and `Type` parameters are required. All other parameters are optional.
//
// - Use the `Condition` parameter to set filter conditions on specific table types, such as filtering by `RecallManagementTableId`.
//
// - Use the `SortBy` and `Order` parameters to control the sort order of the results. The default sort order is ascending by creation time.
//
// - Use the `PageNumber` and `PageSize` parameters for pagination. The `PageNumber` parameter defaults to 1, and the `PageSize` parameter defaults to 10.
//
// - The response includes details about each recall management task, such as its basic information and status.
//
// @param request - ListRecallManagementJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecallManagementJobsResponse
func (client *Client) ListRecallManagementJobsWithOptions(request *ListRecallManagementJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecallManagementJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Condition) {
		query["Condition"] = request.Condition
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecallManagementJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecallManagementJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of recall management tasks that match specified conditions.
//
// Description:
//
// ## Request
//
// - Use this API operation to retrieve a list of recall management tasks.
//
// - The `InstanceId` and `Type` parameters are required. All other parameters are optional.
//
// - Use the `Condition` parameter to set filter conditions on specific table types, such as filtering by `RecallManagementTableId`.
//
// - Use the `SortBy` and `Order` parameters to control the sort order of the results. The default sort order is ascending by creation time.
//
// - Use the `PageNumber` and `PageSize` parameters for pagination. The `PageNumber` parameter defaults to 1, and the `PageSize` parameter defaults to 10.
//
// - The response includes details about each recall management task, such as its basic information and status.
//
// @param request - ListRecallManagementJobsRequest
//
// @return ListRecallManagementJobsResponse
func (client *Client) ListRecallManagementJobs(request *ListRecallManagementJobsRequest) (_result *ListRecallManagementJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecallManagementJobsResponse{}
	_body, _err := client.ListRecallManagementJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of all versions for a specified Recall Management Service.
//
// Description:
//
// ## Description
//
// This operation queries the details of all versions for a specific Recall Management Service, including the version ID, name, effective status, creation time, and modification time. For accurate results, provide the correct `RecallManagementServiceId` and `InstanceId`.
//
// - Use the `PageNumber` and `PageSize` parameters for pagination. By default, the query starts from the first page and returns 50 entries per page.
//
// - Use the `SortBy` parameter to sort the results by creation time or modification time. By default, the results are sorted by creation time in ascending order.
//
// - The `Order` parameter specifies the sort order. Valid values are `ASC` for ascending order and `DESC` for descending order.
//
// @param request - ListRecallManagementServiceVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecallManagementServiceVersionsResponse
func (client *Client) ListRecallManagementServiceVersionsWithOptions(RecallManagementServiceId *string, request *ListRecallManagementServiceVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecallManagementServiceVersionsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecallManagementServiceVersions"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecallManagementServiceVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of all versions for a specified Recall Management Service.
//
// Description:
//
// ## Description
//
// This operation queries the details of all versions for a specific Recall Management Service, including the version ID, name, effective status, creation time, and modification time. For accurate results, provide the correct `RecallManagementServiceId` and `InstanceId`.
//
// - Use the `PageNumber` and `PageSize` parameters for pagination. By default, the query starts from the first page and returns 50 entries per page.
//
// - Use the `SortBy` parameter to sort the results by creation time or modification time. By default, the results are sorted by creation time in ascending order.
//
// - The `Order` parameter specifies the sort order. Valid values are `ASC` for ascending order and `DESC` for descending order.
//
// @param request - ListRecallManagementServiceVersionsRequest
//
// @return ListRecallManagementServiceVersionsResponse
func (client *Client) ListRecallManagementServiceVersions(RecallManagementServiceId *string, request *ListRecallManagementServiceVersionsRequest) (_result *ListRecallManagementServiceVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecallManagementServiceVersionsResponse{}
	_body, _err := client.ListRecallManagementServiceVersionsWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of recall management services under a specified instance.
//
// Description:
//
// ## Operation description
//
// By calling the `ListRecallManagementServices` operation, you can retrieve the list of recall management services under a specific instance based on given parameters such as InstanceId, PageNumber, and PageSize. You can sort results by creation time or update time in ascending or descending order.
//
// - **InstanceId*	- is required and specifies the instance to which the services belong.
//
// - The pagination parameters **PageNumber*	- and **PageSize*	- allow you to control the amount of returned data and the page from which to start displaying results. By default, data from the first page is returned.
//
// - Use the **SortBy*	- and **Order*	- parameters to customize the sorting of the list.
//
// @param request - ListRecallManagementServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecallManagementServicesResponse
func (client *Client) ListRecallManagementServicesWithOptions(request *ListRecallManagementServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecallManagementServicesResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecallManagementServices"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecallManagementServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of recall management services under a specified instance.
//
// Description:
//
// ## Operation description
//
// By calling the `ListRecallManagementServices` operation, you can retrieve the list of recall management services under a specific instance based on given parameters such as InstanceId, PageNumber, and PageSize. You can sort results by creation time or update time in ascending or descending order.
//
// - **InstanceId*	- is required and specifies the instance to which the services belong.
//
// - The pagination parameters **PageNumber*	- and **PageSize*	- allow you to control the amount of returned data and the page from which to start displaying results. By default, data from the first page is returned.
//
// - Use the **SortBy*	- and **Order*	- parameters to customize the sorting of the list.
//
// @param request - ListRecallManagementServicesRequest
//
// @return ListRecallManagementServicesResponse
func (client *Client) ListRecallManagementServices(request *ListRecallManagementServicesRequest) (_result *ListRecallManagementServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecallManagementServicesResponse{}
	_body, _err := client.ListRecallManagementServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all versions of a specified RecallManagementTable.
//
// Description:
//
// ## Usage
//
// - To retrieve the version history of a specific RecallManagementTable, provide the `RecallManagementTableId` and `InstanceId`.
//
// - Use the `SortBy` parameter to sort the results by creation time or update time. By default, the results are sorted by creation time in ascending order.
//
// - The `PageNumber` and `PageSize` parameters enable pagination, which allows you to control the number of items to return and the page to display.
//
// - If the `Order` parameter is not specified, the results are sorted in ascending order by default.
//
// @param request - ListRecallManagementTableVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecallManagementTableVersionsResponse
func (client *Client) ListRecallManagementTableVersionsWithOptions(RecallManagementTableId *string, request *ListRecallManagementTableVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecallManagementTableVersionsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecallManagementTableVersions"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables/" + dara.PercentEncode(dara.StringValue(RecallManagementTableId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecallManagementTableVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all versions of a specified RecallManagementTable.
//
// Description:
//
// ## Usage
//
// - To retrieve the version history of a specific RecallManagementTable, provide the `RecallManagementTableId` and `InstanceId`.
//
// - Use the `SortBy` parameter to sort the results by creation time or update time. By default, the results are sorted by creation time in ascending order.
//
// - The `PageNumber` and `PageSize` parameters enable pagination, which allows you to control the number of items to return and the page to display.
//
// - If the `Order` parameter is not specified, the results are sorted in ascending order by default.
//
// @param request - ListRecallManagementTableVersionsRequest
//
// @return ListRecallManagementTableVersionsResponse
func (client *Client) ListRecallManagementTableVersions(RecallManagementTableId *string, request *ListRecallManagementTableVersionsRequest) (_result *ListRecallManagementTableVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecallManagementTableVersionsResponse{}
	_body, _err := client.ListRecallManagementTableVersionsWithOptions(RecallManagementTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the recall management tables for a specified instance. Pagination and sorting are supported.
//
// Description:
//
// ## Request
//
// - **InstanceId*	- is a required parameter specifying the instance to query.
//
// - The **Name*	- and **Type*	- parameters filter recall management tables by name or type.
//
// - The **PageNumber*	- and **PageSize*	- parameters control pagination. By default, the query returns the first 50 records.
//
// - You can sort results by creation time (GmtCreateTime) or modification time (GmtModifiedTime) in ascending (ASC) or descending (DESC) order.
//
// @param request - ListRecallManagementTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecallManagementTablesResponse
func (client *Client) ListRecallManagementTablesWithOptions(request *ListRecallManagementTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecallManagementTablesResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecallManagementTables"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecallManagementTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the recall management tables for a specified instance. Pagination and sorting are supported.
//
// Description:
//
// ## Request
//
// - **InstanceId*	- is a required parameter specifying the instance to query.
//
// - The **Name*	- and **Type*	- parameters filter recall management tables by name or type.
//
// - The **PageNumber*	- and **PageSize*	- parameters control pagination. By default, the query returns the first 50 records.
//
// - You can sort results by creation time (GmtCreateTime) or modification time (GmtModifiedTime) in ascending (ASC) or descending (DESC) order.
//
// @param request - ListRecallManagementTablesRequest
//
// @return ListRecallManagementTablesResponse
func (client *Client) ListRecallManagementTables(request *ListRecallManagementTablesRequest) (_result *ListRecallManagementTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecallManagementTablesResponse{}
	_body, _err := client.ListRecallManagementTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - ListResourceRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRulesResponse
func (client *Client) ListResourceRulesWithOptions(request *ListResourceRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceRuleId) {
		query["ResourceRuleId"] = request.ResourceRuleId
	}

	if !dara.IsNil(request.ResourceRuleName) {
		query["ResourceRuleName"] = request.ResourceRuleName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceRules"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - ListResourceRulesRequest
//
// @return ListResourceRulesResponse
func (client *Client) ListResourceRules(request *ListResourceRulesRequest) (_result *ListResourceRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceRulesResponse{}
	_body, _err := client.ListResourceRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取样本一致性任务列表
//
// @param request - ListSampleConsistencyJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSampleConsistencyJobsResponse
func (client *Client) ListSampleConsistencyJobsWithOptions(request *ListSampleConsistencyJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSampleConsistencyJobsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSampleConsistencyJobs"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSampleConsistencyJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取样本一致性任务列表
//
// @param request - ListSampleConsistencyJobsRequest
//
// @return ListSampleConsistencyJobsResponse
func (client *Client) ListSampleConsistencyJobs(request *ListSampleConsistencyJobsRequest) (_result *ListSampleConsistencyJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSampleConsistencyJobsResponse{}
	_body, _err := client.ListSampleConsistencyJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of scenes.
//
// @param request - ListScenesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScenesResponse
func (client *Client) ListScenesWithOptions(request *ListScenesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListScenesResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScenes"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of scenes.
//
// @param request - ListScenesRequest
//
// @return ListScenesResponse
func (client *Client) ListScenes(request *ListScenesRequest) (_result *ListScenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScenesResponse{}
	_body, _err := client.ListScenesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the subcrowds for a specified crowd.
//
// @param request - ListSubCrowdsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubCrowdsResponse
func (client *Client) ListSubCrowdsWithOptions(CrowdId *string, request *ListSubCrowdsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubCrowdsResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubCrowds"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId)) + "/subcrowds"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubCrowdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the subcrowds for a specified crowd.
//
// @param request - ListSubCrowdsRequest
//
// @return ListSubCrowdsResponse
func (client *Client) ListSubCrowds(CrowdId *string, request *ListSubCrowdsRequest) (_result *ListSubCrowdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubCrowdsResponse{}
	_body, _err := client.ListSubCrowdsWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of data tables.
//
// @param request - ListTableMetasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableMetasResponse
func (client *Client) ListTableMetasWithOptions(request *ListTableMetasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTableMetasResponse, _err error) {
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

	if !dara.IsNil(request.Module) {
		query["Module"] = request.Module
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableMetas"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data tables.
//
// @param request - ListTableMetasRequest
//
// @return ListTableMetasResponse
func (client *Client) ListTableMetas(request *ListTableMetasRequest) (_result *ListTableMetasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTableMetasResponse{}
	_body, _err := client.ListTableMetasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the historical traffic records for a specific traffic control target.
//
// Description:
//
// ## Usage notes
//
// - The `TrafficControlTargetId`, `InstanceId`, and `Environment` parameters are required.
//
// - You can use `StartTime` and `EndTime` to specify the time range.
//
// - The `Threshold` parameter is optional.
//
// - Use `ExperimentId` and `ExperimentGroupId` to filter data for a specific experiment or experiment group.
//
// - Use `ItemId` to filter traffic data for a specific item.
//
// - The supported environments are the Daily environment, pre-production environment (Pre), and production environment (Prod).
//
// @param request - ListTrafficControlTargetTrafficHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficControlTargetTrafficHistoryResponse
func (client *Client) ListTrafficControlTargetTrafficHistoryWithOptions(TrafficControlTargetId *string, request *ListTrafficControlTargetTrafficHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrafficControlTargetTrafficHistoryResponse, _err error) {
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

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.ExperimentGroupId) {
		query["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.ExperimentId) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemId) {
		query["ItemId"] = request.ItemId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrafficControlTargetTrafficHistory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/traffichistories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrafficControlTargetTrafficHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the historical traffic records for a specific traffic control target.
//
// Description:
//
// ## Usage notes
//
// - The `TrafficControlTargetId`, `InstanceId`, and `Environment` parameters are required.
//
// - You can use `StartTime` and `EndTime` to specify the time range.
//
// - The `Threshold` parameter is optional.
//
// - Use `ExperimentId` and `ExperimentGroupId` to filter data for a specific experiment or experiment group.
//
// - Use `ItemId` to filter traffic data for a specific item.
//
// - The supported environments are the Daily environment, pre-production environment (Pre), and production environment (Prod).
//
// @param request - ListTrafficControlTargetTrafficHistoryRequest
//
// @return ListTrafficControlTargetTrafficHistoryResponse
func (client *Client) ListTrafficControlTargetTrafficHistory(TrafficControlTargetId *string, request *ListTrafficControlTargetTrafficHistoryRequest) (_result *ListTrafficControlTargetTrafficHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrafficControlTargetTrafficHistoryResponse{}
	_body, _err := client.ListTrafficControlTargetTrafficHistoryWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists traffic control tasks that meet specified conditions.
//
// Description:
//
// ## Request
//
// - This API retrieves a list of traffic control tasks.
//
// - Use query parameters to filter and sort the results.
//
// - This operation supports pagination. You can also retrieve all results in a single response.
//
// - Note: The `InstanceId` is a required parameter. All other parameters are optional.
//
// @param request - ListTrafficControlTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficControlTasksResponse
func (client *Client) ListTrafficControlTasksWithOptions(request *ListTrafficControlTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrafficControlTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ControlTargetFilter) {
		query["ControlTargetFilter"] = request.ControlTargetFilter
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TrafficControlTaskId) {
		query["TrafficControlTaskId"] = request.TrafficControlTaskId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrafficControlTasks"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrafficControlTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists traffic control tasks that meet specified conditions.
//
// Description:
//
// ## Request
//
// - This API retrieves a list of traffic control tasks.
//
// - Use query parameters to filter and sort the results.
//
// - This operation supports pagination. You can also retrieve all results in a single response.
//
// - Note: The `InstanceId` is a required parameter. All other parameters are optional.
//
// @param request - ListTrafficControlTasksRequest
//
// @return ListTrafficControlTasksResponse
func (client *Client) ListTrafficControlTasks(request *ListTrafficControlTasksRequest) (_result *ListTrafficControlTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrafficControlTasksResponse{}
	_body, _err := client.ListTrafficControlTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Takes an experiment offline.
//
// @param request - OfflineExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineExperimentResponse
func (client *Client) OfflineExperimentWithOptions(ExperimentId *string, request *OfflineExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineExperimentResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Takes an experiment offline.
//
// @param request - OfflineExperimentRequest
//
// @return OfflineExperimentResponse
func (client *Client) OfflineExperiment(ExperimentId *string, request *OfflineExperimentRequest) (_result *OfflineExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineExperimentResponse{}
	_body, _err := client.OfflineExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Takes a specified experiment group offline.
//
// @param request - OfflineExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineExperimentGroupResponse
func (client *Client) OfflineExperimentGroupWithOptions(ExperimentGroupId *string, request *OfflineExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineExperimentGroupResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId)) + "/action/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Takes a specified experiment group offline.
//
// @param request - OfflineExperimentGroupRequest
//
// @return OfflineExperimentGroupResponse
func (client *Client) OfflineExperimentGroup(ExperimentGroupId *string, request *OfflineExperimentGroupRequest) (_result *OfflineExperimentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineExperimentGroupResponse{}
	_body, _err := client.OfflineExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Takes the specified laboratory offline.
//
// @param request - OfflineLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineLaboratoryResponse
func (client *Client) OfflineLaboratoryWithOptions(LaboratoryId *string, request *OfflineLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineLaboratoryResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId)) + "/action/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Takes the specified laboratory offline.
//
// @param request - OfflineLaboratoryRequest
//
// @return OfflineLaboratoryResponse
func (client *Client) OfflineLaboratory(LaboratoryId *string, request *OfflineLaboratoryRequest) (_result *OfflineLaboratoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineLaboratoryResponse{}
	_body, _err := client.OfflineLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Takes a specified recall management service offline.
//
// Description:
//
// ## Description
//
// Use this API to take a specific recall management service offline. Ensure that the provided `RecallManagementServiceId` and `InstanceId` are accurate to prevent unintended operations.
//
// - **Important**: Once a recall management service is taken offline, it stops processing new requests until you reactivate it.
//
// - Back up any required data or configurations before you perform this operation in case you need to restore the current state.
//
// @param request - OfflineRecallManagementServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineRecallManagementServiceResponse
func (client *Client) OfflineRecallManagementServiceWithOptions(RecallManagementServiceId *string, request *OfflineRecallManagementServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineRecallManagementServiceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineRecallManagementService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/action/offline"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineRecallManagementServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Takes a specified recall management service offline.
//
// Description:
//
// ## Description
//
// Use this API to take a specific recall management service offline. Ensure that the provided `RecallManagementServiceId` and `InstanceId` are accurate to prevent unintended operations.
//
// - **Important**: Once a recall management service is taken offline, it stops processing new requests until you reactivate it.
//
// - Back up any required data or configurations before you perform this operation in case you need to restore the current state.
//
// @param request - OfflineRecallManagementServiceRequest
//
// @return OfflineRecallManagementServiceResponse
func (client *Client) OfflineRecallManagementService(RecallManagementServiceId *string, request *OfflineRecallManagementServiceRequest) (_result *OfflineRecallManagementServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineRecallManagementServiceResponse{}
	_body, _err := client.OfflineRecallManagementServiceWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Brings a specified experiment online.
//
// @param request - OnlineExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineExperimentResponse
func (client *Client) OnlineExperimentWithOptions(ExperimentId *string, request *OnlineExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineExperimentResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Brings a specified experiment online.
//
// @param request - OnlineExperimentRequest
//
// @return OnlineExperimentResponse
func (client *Client) OnlineExperiment(ExperimentId *string, request *OnlineExperimentRequest) (_result *OnlineExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineExperimentResponse{}
	_body, _err := client.OnlineExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Brings a specified experiment group online.
//
// @param request - OnlineExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineExperimentGroupResponse
func (client *Client) OnlineExperimentGroupWithOptions(ExperimentGroupId *string, request *OnlineExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineExperimentGroupResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId)) + "/action/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Brings a specified experiment group online.
//
// @param request - OnlineExperimentGroupRequest
//
// @return OnlineExperimentGroupResponse
func (client *Client) OnlineExperimentGroup(ExperimentGroupId *string, request *OnlineExperimentGroupRequest) (_result *OnlineExperimentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineExperimentGroupResponse{}
	_body, _err := client.OnlineExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a specified laboratory for experimental analysis.
//
// @param request - OnlineLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineLaboratoryResponse
func (client *Client) OnlineLaboratoryWithOptions(LaboratoryId *string, request *OnlineLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineLaboratoryResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId)) + "/action/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a specified laboratory for experimental analysis.
//
// @param request - OnlineLaboratoryRequest
//
// @return OnlineLaboratoryResponse
func (client *Client) OnlineLaboratory(LaboratoryId *string, request *OnlineLaboratoryRequest) (_result *OnlineLaboratoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineLaboratoryResponse{}
	_body, _err := client.OnlineLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation brings a specified Recall Management Service online.
//
// Description:
//
// ## Request
//
// You can use this operation to bring a Recall Management Service online by specifying the Recall Management Service ID and the instance ID. Ensure that the `RecallManagementServiceId` and `InstanceId` are correct and that you have the required permissions.
//
// @param request - OnlineRecallManagementServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineRecallManagementServiceResponse
func (client *Client) OnlineRecallManagementServiceWithOptions(RecallManagementServiceId *string, request *OnlineRecallManagementServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OnlineRecallManagementServiceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineRecallManagementService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/action/online"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineRecallManagementServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation brings a specified Recall Management Service online.
//
// Description:
//
// ## Request
//
// You can use this operation to bring a Recall Management Service online by specifying the Recall Management Service ID and the instance ID. Ensure that the `RecallManagementServiceId` and `InstanceId` are correct and that you have the required permissions.
//
// @param request - OnlineRecallManagementServiceRequest
//
// @return OnlineRecallManagementServiceResponse
func (client *Client) OnlineRecallManagementService(RecallManagementServiceId *string, request *OnlineRecallManagementServiceRequest) (_result *OnlineRecallManagementServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OnlineRecallManagementServiceResponse{}
	_body, _err := client.OnlineRecallManagementServiceWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes a MaxCompute table with the recall engine. This operation allows you to publish specific partitions and select a synchronization mode.
//
// Description:
//
// ## Request details
//
// This API synchronizes a specified MaxCompute table with the recall engine. You must provide the correct `RecallManagementTableId` in the path parameter and the instance ID in the request body. You can also specify the table partitions to publish, whether to skip the threshold check, and the synchronization mode. To publish specific partitions, provide them as key-value pairs in the `Partitions` field.
//
// @param request - PublishRecallManagementTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRecallManagementTableResponse
func (client *Client) PublishRecallManagementTableWithOptions(RecallManagementTableId *string, request *PublishRecallManagementTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishRecallManagementTableResponse, _err error) {
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

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Partition) {
		body["Partition"] = request.Partition
	}

	if !dara.IsNil(request.Partitions) {
		body["Partitions"] = request.Partitions
	}

	if !dara.IsNil(request.SkipThresholdCheck) {
		body["SkipThresholdCheck"] = request.SkipThresholdCheck
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishRecallManagementTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables/" + dara.PercentEncode(dara.StringValue(RecallManagementTableId)) + "/action/publish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishRecallManagementTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes a MaxCompute table with the recall engine. This operation allows you to publish specific partitions and select a synchronization mode.
//
// Description:
//
// ## Request details
//
// This API synchronizes a specified MaxCompute table with the recall engine. You must provide the correct `RecallManagementTableId` in the path parameter and the instance ID in the request body. You can also specify the table partitions to publish, whether to skip the threshold check, and the synchronization mode. To publish specific partitions, provide them as key-value pairs in the `Partitions` field.
//
// @param request - PublishRecallManagementTableRequest
//
// @return PublishRecallManagementTableResponse
func (client *Client) PublishRecallManagementTable(RecallManagementTableId *string, request *PublishRecallManagementTableRequest) (_result *PublishRecallManagementTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishRecallManagementTableResponse{}
	_body, _err := client.PublishRecallManagementTableWithOptions(RecallManagementTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If an experiment is stable and performs well, you can push all traffic to it. This action retires the original experiment group and creates a new one that contains only this experiment. The new group receives 100% of the traffic.
//
// @param request - PushAllExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushAllExperimentResponse
func (client *Client) PushAllExperimentWithOptions(ExperimentId *string, request *PushAllExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushAllExperimentResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushAllExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/action/pushall"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushAllExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If an experiment is stable and performs well, you can push all traffic to it. This action retires the original experiment group and creates a new one that contains only this experiment. The new group receives 100% of the traffic.
//
// @param request - PushAllExperimentRequest
//
// @return PushAllExperimentResponse
func (client *Client) PushAllExperiment(ExperimentId *string, request *PushAllExperimentRequest) (_result *PushAllExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushAllExperimentResponse{}
	_body, _err := client.PushAllExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送指标到指定资源规则
//
// @param tmpReq - PushResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResourceRuleResponse
func (client *Client) PushResourceRuleWithOptions(ResourceRuleId *string, tmpReq *PushResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushResourceRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushResourceRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetricInfo) {
		request.MetricInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricInfo, dara.String("MetricInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricInfoShrink) {
		query["MetricInfo"] = request.MetricInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/action/push"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushResourceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送指标到指定资源规则
//
// @param request - PushResourceRuleRequest
//
// @return PushResourceRuleResponse
func (client *Client) PushResourceRule(ResourceRuleId *string, request *PushResourceRuleRequest) (_result *PushResourceRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushResourceRuleResponse{}
	_body, _err := client.PushResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves statistics for a specified data diagnosis task within a time range.
//
// Description:
//
// ## Request description
//
// - The `DataDiagnosisId` parameter is required and specifies the data diagnosis task.
//
// - The `InstanceId` parameter is also required and specifies the instance.
//
// - The `StartDate` and `EndDate` parameters specify the start and end dates of the time range. The format is YYYY-MM-DD.
//
// - The `RemainRateType` parameter is optional. It specifies the retention rate report type. The default value is \\"Period\\", which indicates a periodic report.
//
// - The response includes the request ID (`RequestId`) and a `Statistics` object. This object contains the dates of task failures (`FailedDates`) and dates with missing task data (`NoDataDates`).
//
// @param request - QueryDataDiagnosisStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataDiagnosisStatisticsResponse
func (client *Client) QueryDataDiagnosisStatisticsWithOptions(DataDiagnosisId *string, request *QueryDataDiagnosisStatisticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryDataDiagnosisStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RemainRateType) {
		query["RemainRateType"] = request.RemainRateType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDataDiagnosisStatistics"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnoses/" + dara.PercentEncode(dara.StringValue(DataDiagnosisId)) + "/statistics/action/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDataDiagnosisStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves statistics for a specified data diagnosis task within a time range.
//
// Description:
//
// ## Request description
//
// - The `DataDiagnosisId` parameter is required and specifies the data diagnosis task.
//
// - The `InstanceId` parameter is also required and specifies the instance.
//
// - The `StartDate` and `EndDate` parameters specify the start and end dates of the time range. The format is YYYY-MM-DD.
//
// - The `RemainRateType` parameter is optional. It specifies the retention rate report type. The default value is \\"Period\\", which indicates a periodic report.
//
// - The response includes the request ID (`RequestId`) and a `Statistics` object. This object contains the dates of task failures (`FailedDates`) and dates with missing task data (`NoDataDates`).
//
// @param request - QueryDataDiagnosisStatisticsRequest
//
// @return QueryDataDiagnosisStatisticsResponse
func (client *Client) QueryDataDiagnosisStatistics(DataDiagnosisId *string, request *QueryDataDiagnosisStatisticsRequest) (_result *QueryDataDiagnosisStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryDataDiagnosisStatisticsResponse{}
	_body, _err := client.QueryDataDiagnosisStatisticsWithOptions(DataDiagnosisId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves records from a specified recall management table.
//
// Description:
//
// ## Request
//
// This API retrieves records from a specific recall management table using the provided primary keys. You must provide a valid `InstanceId` and `RecallManagementTableId`, and a non-empty `PrimaryKeys` list. If you specify `RecallManagementTableVersionId`, the API returns records from that version; otherwise, it uses the currently published version.
//
// @param request - QueryRecallManagementTableRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecallManagementTableRecordsResponse
func (client *Client) QueryRecallManagementTableRecordsWithOptions(RecallManagementTableId *string, request *QueryRecallManagementTableRecordsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryRecallManagementTableRecordsResponse, _err error) {
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

	if !dara.IsNil(request.PrimaryKeys) {
		query["PrimaryKeys"] = request.PrimaryKeys
	}

	if !dara.IsNil(request.RecallManagementTableVersionId) {
		query["RecallManagementTableVersionId"] = request.RecallManagementTableVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecallManagementTableRecords"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables/" + dara.PercentEncode(dara.StringValue(RecallManagementTableId)) + "/queryrecords"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecallManagementTableRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves records from a specified recall management table.
//
// Description:
//
// ## Request
//
// This API retrieves records from a specific recall management table using the provided primary keys. You must provide a valid `InstanceId` and `RecallManagementTableId`, and a non-empty `PrimaryKeys` list. If you specify `RecallManagementTableVersionId`, the API returns records from that version; otherwise, it uses the currently published version.
//
// @param request - QueryRecallManagementTableRecordsRequest
//
// @return QueryRecallManagementTableRecordsResponse
func (client *Client) QueryRecallManagementTableRecords(RecallManagementTableId *string, request *QueryRecallManagementTableRecordsRequest) (_result *QueryRecallManagementTableRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRecallManagementTableRecordsResponse{}
	_body, _err := client.QueryRecallManagementTableRecordsWithOptions(RecallManagementTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看样本一致性任务差异的详情
//
// @param request - QuerySampleConsistencyJobDifferenceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySampleConsistencyJobDifferenceResponse
func (client *Client) QuerySampleConsistencyJobDifferenceWithOptions(SampleConsistencyJobId *string, request *QuerySampleConsistencyJobDifferenceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QuerySampleConsistencyJobDifferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySampleConsistencyJobDifference"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/querydifference"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySampleConsistencyJobDifferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看样本一致性任务差异的详情
//
// @param request - QuerySampleConsistencyJobDifferenceRequest
//
// @return QuerySampleConsistencyJobDifferenceResponse
func (client *Client) QuerySampleConsistencyJobDifference(SampleConsistencyJobId *string, request *QuerySampleConsistencyJobDifferenceRequest) (_result *QuerySampleConsistencyJobDifferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySampleConsistencyJobDifferenceResponse{}
	_body, _err := client.QuerySampleConsistencyJobDifferenceWithOptions(SampleConsistencyJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the traffic control details of a target item for a given environment and date range.
//
// Description:
//
// ## Request
//
// Use this API to query the details of single-item control for a given traffic control target on a specified date and for a specific instance ID and environment. The details include traffic data and feature information for the top 100 items before and after the control is applied. Ensure that the `TrafficControlTargetId`, `InstanceId`, and `Environment` parameters are accurate, and that the `Date` is in YYYY-MM-DD format. Although the `Date` parameter is optional, we recommend specifying a date for meaningful results.
//
// @param request - QueryTrafficControlTargetItemReportDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTrafficControlTargetItemReportDetailResponse
func (client *Client) QueryTrafficControlTargetItemReportDetailWithOptions(TrafficControlTargetId *string, request *QueryTrafficControlTargetItemReportDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTrafficControlTargetItemReportDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTrafficControlTargetItemReportDetail"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/itemcontrolreportdetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTrafficControlTargetItemReportDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the traffic control details of a target item for a given environment and date range.
//
// Description:
//
// ## Request
//
// Use this API to query the details of single-item control for a given traffic control target on a specified date and for a specific instance ID and environment. The details include traffic data and feature information for the top 100 items before and after the control is applied. Ensure that the `TrafficControlTargetId`, `InstanceId`, and `Environment` parameters are accurate, and that the `Date` is in YYYY-MM-DD format. Although the `Date` parameter is optional, we recommend specifying a date for meaningful results.
//
// @param request - QueryTrafficControlTargetItemReportDetailRequest
//
// @return QueryTrafficControlTargetItemReportDetailResponse
func (client *Client) QueryTrafficControlTargetItemReportDetail(TrafficControlTargetId *string, request *QueryTrafficControlTargetItemReportDetailRequest) (_result *QueryTrafficControlTargetItemReportDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTrafficControlTargetItemReportDetailResponse{}
	_body, _err := client.QueryTrafficControlTargetItemReportDetailWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the deployment status and related information of a specified traffic control task in a specific environment.
//
// Description:
//
// ## Operation description
//
// You can call this operation to query the deployment result of a traffic control task specified by TrafficControlTaskId for a given instance ID and environment. Make sure that the specified InstanceId is associated with your account and that the Environment parameter value is valid (Daily for daily environment, Pre for staging environment, Prod for production environment). All request parameters are required.
//
// @param request - QueryTrafficControlTaskDeployResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTrafficControlTaskDeployResultResponse
func (client *Client) QueryTrafficControlTaskDeployResultWithOptions(TrafficControlTaskId *string, request *QueryTrafficControlTaskDeployResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTrafficControlTaskDeployResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTrafficControlTaskDeployResult"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/queryresult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTrafficControlTaskDeployResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the deployment status and related information of a specified traffic control task in a specific environment.
//
// Description:
//
// ## Operation description
//
// You can call this operation to query the deployment result of a traffic control task specified by TrafficControlTaskId for a given instance ID and environment. Make sure that the specified InstanceId is associated with your account and that the Environment parameter value is valid (Daily for daily environment, Pre for staging environment, Prod for production environment). All request parameters are required.
//
// @param request - QueryTrafficControlTaskDeployResultRequest
//
// @return QueryTrafficControlTaskDeployResultResponse
func (client *Client) QueryTrafficControlTaskDeployResult(TrafficControlTaskId *string, request *QueryTrafficControlTaskDeployResultRequest) (_result *QueryTrafficControlTaskDeployResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTrafficControlTaskDeployResultResponse{}
	_body, _err := client.QueryTrafficControlTaskDeployResultWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a detailed report on item control for a specified traffic control task.
//
// Description:
//
// ## Description
//
// - This API retrieves the item control results for a specific traffic control task within a given time range.
//
// - `TrafficControlTaskId` is the unique identifier for a traffic control task.
//
// - `InstanceId` specifies the instance that runs the task.
//
// - The `Environment` parameter specifies the task\\"s execution environment. Valid values are Daily (development environment), Pre (staging environment), and Prod (production environment).
//
// - `StartTime` and `EndTime` specify the start and end of the time range for the report, respectively. The format is "YYYY-MM-DD HH:MM:SS".
//
// - The specified start and end times must be valid and span no more than two consecutive calendar days.
//
// @param request - QueryTrafficControlTaskItemReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTrafficControlTaskItemReportResponse
func (client *Client) QueryTrafficControlTaskItemReportWithOptions(TrafficControlTaskId *string, request *QueryTrafficControlTaskItemReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTrafficControlTaskItemReportResponse, _err error) {
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

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTrafficControlTaskItemReport"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/queryitemreport"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTrafficControlTaskItemReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a detailed report on item control for a specified traffic control task.
//
// Description:
//
// ## Description
//
// - This API retrieves the item control results for a specific traffic control task within a given time range.
//
// - `TrafficControlTaskId` is the unique identifier for a traffic control task.
//
// - `InstanceId` specifies the instance that runs the task.
//
// - The `Environment` parameter specifies the task\\"s execution environment. Valid values are Daily (development environment), Pre (staging environment), and Prod (production environment).
//
// - `StartTime` and `EndTime` specify the start and end of the time range for the report, respectively. The format is "YYYY-MM-DD HH:MM:SS".
//
// - The specified start and end times must be valid and span no more than two consecutive calendar days.
//
// @param request - QueryTrafficControlTaskItemReportRequest
//
// @return QueryTrafficControlTaskItemReportResponse
func (client *Client) QueryTrafficControlTaskItemReport(TrafficControlTaskId *string, request *QueryTrafficControlTaskItemReportRequest) (_result *QueryTrafficControlTaskItemReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTrafficControlTaskItemReportResponse{}
	_body, _err := client.QueryTrafficControlTaskItemReportWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a traffic control task for the specified instance and environment.
//
// Description:
//
// ## Request
//
// Use this API to release a traffic control task for a specific instance and environment (Daily, Pre, or Prod). Your request must include the `TrafficControlTaskId`, `InstanceId`, and `Environment` parameters.
//
// - `TrafficControlTaskId`: The unique ID of the traffic control task.
//
// - `InstanceId`: The ID of the target instance.
//
// - `Environment`: The execution environment for the traffic control task. Valid values: `Daily`, `Pre`, and `Prod`.
//
// The request succeeds only if all required parameters are provided correctly. A successful response includes a `RequestId`.
//
// @param request - ReleaseTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseTrafficControlTaskResponse
func (client *Client) ReleaseTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *ReleaseTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReleaseTrafficControlTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/release"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a traffic control task for the specified instance and environment.
//
// Description:
//
// ## Request
//
// Use this API to release a traffic control task for a specific instance and environment (Daily, Pre, or Prod). Your request must include the `TrafficControlTaskId`, `InstanceId`, and `Environment` parameters.
//
// - `TrafficControlTaskId`: The unique ID of the traffic control task.
//
// - `InstanceId`: The ID of the target instance.
//
// - `Environment`: The execution environment for the traffic control task. Valid values: `Daily`, `Pre`, and `Prod`.
//
// The request succeeds only if all required parameters are provided correctly. A successful response includes a `RequestId`.
//
// @param request - ReleaseTrafficControlTaskRequest
//
// @return ReleaseTrafficControlTaskResponse
func (client *Client) ReleaseTrafficControlTask(TrafficControlTaskId *string, request *ReleaseTrafficControlTaskRequest) (_result *ReleaseTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseTrafficControlTaskResponse{}
	_body, _err := client.ReleaseTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve a metric group\\"s report.
//
// @param request - ReportABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportABMetricGroupResponse
func (client *Client) ReportABMetricGroupWithOptions(ABMetricGroupId *string, request *ReportABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReportABMetricGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseExperimentId) {
		body["BaseExperimentId"] = request.BaseExperimentId
	}

	if !dara.IsNil(request.DimensionFields) {
		body["DimensionFields"] = request.DimensionFields
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ExperimentGroupId) {
		body["ExperimentGroupId"] = request.ExperimentGroupId
	}

	if !dara.IsNil(request.ExperimentIds) {
		body["ExperimentIds"] = request.ExperimentIds
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ReportType) {
		body["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TimeStatisticsMethod) {
		body["TimeStatisticsMethod"] = request.TimeStatisticsMethod
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId)) + "/action/report"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve a metric group\\"s report.
//
// @param request - ReportABMetricGroupRequest
//
// @return ReportABMetricGroupResponse
func (client *Client) ReportABMetricGroup(ABMetricGroupId *string, request *ReportABMetricGroupRequest) (_result *ReportABMetricGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReportABMetricGroupResponse{}
	_body, _err := client.ReportABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 样本一致性任务报表
//
// @param request - ReportSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportSampleConsistencyJobResponse
func (client *Client) ReportSampleConsistencyJobWithOptions(SampleConsistencyJobId *string, request *ReportSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReportSampleConsistencyJobResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/report"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportSampleConsistencyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 样本一致性任务报表
//
// @param request - ReportSampleConsistencyJobRequest
//
// @return ReportSampleConsistencyJobResponse
func (client *Client) ReportSampleConsistencyJob(SampleConsistencyJobId *string, request *ReportSampleConsistencyJobRequest) (_result *ReportSampleConsistencyJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReportSampleConsistencyJobResponse{}
	_body, _err := client.ReportSampleConsistencyJobWithOptions(SampleConsistencyJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Conducts conversations with users through an AI shopping guide to provide product recommendation services.
//
// Description:
//
// ## Operation description
//
// - This API is used to send conversation messages to the AI shopping guide and supports Server-Sent Events (SSE).
//
// - `InstanceId`, `SessionId`, `SceneId`, `ServiceId`, `Environment`, `Uid`, and `Language` are required parameters. Ensure the accuracy of these values to obtain optimal responses.
//
// - The `InputMessage` must contain at least one text-type message that describes the user\\"s request or question.
//
// - Based on the provided input, the system returns corresponding recommendation results or other relevant information.
//
// - Check the returned `StopReason` field to understand whether the session has ended and the reason.
//
// @param request - ShoppingAssistantRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShoppingAssistantResponse
func (client *Client) ShoppingAssistantWithSSE(request *ShoppingAssistantRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ShoppingAssistantResponse, _yieldErr chan error) {
	defer close(_yield)
	client.shoppingAssistantWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// Conducts conversations with users through an AI shopping guide to provide product recommendation services.
//
// Description:
//
// ## Operation description
//
// - This API is used to send conversation messages to the AI shopping guide and supports Server-Sent Events (SSE).
//
// - `InstanceId`, `SessionId`, `SceneId`, `ServiceId`, `Environment`, `Uid`, and `Language` are required parameters. Ensure the accuracy of these values to obtain optimal responses.
//
// - The `InputMessage` must contain at least one text-type message that describes the user\\"s request or question.
//
// - Based on the provided input, the system returns corresponding recommendation results or other relevant information.
//
// - Check the returned `StopReason` field to understand whether the session has ended and the reason.
//
// @param request - ShoppingAssistantRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShoppingAssistantResponse
func (client *Client) ShoppingAssistantWithOptions(request *ShoppingAssistantRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ShoppingAssistantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Contents) {
		body["Contents"] = request.Contents
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InputMessage) {
		body["InputMessage"] = request.InputMessage
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ShoppingAssistant"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/conversations/shopping_assistant/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ShoppingAssistantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Conducts conversations with users through an AI shopping guide to provide product recommendation services.
//
// Description:
//
// ## Operation description
//
// - This API is used to send conversation messages to the AI shopping guide and supports Server-Sent Events (SSE).
//
// - `InstanceId`, `SessionId`, `SceneId`, `ServiceId`, `Environment`, `Uid`, and `Language` are required parameters. Ensure the accuracy of these values to obtain optimal responses.
//
// - The `InputMessage` must contain at least one text-type message that describes the user\\"s request or question.
//
// - Based on the provided input, the system returns corresponding recommendation results or other relevant information.
//
// - Check the returned `StopReason` field to understand whether the session has ended and the reason.
//
// @param request - ShoppingAssistantRequest
//
// @return ShoppingAssistantResponse
func (client *Client) ShoppingAssistant(request *ShoppingAssistantRequest) (_result *ShoppingAssistantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ShoppingAssistantResponse{}
	_body, _err := client.ShoppingAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Splits the target values for a traffic control target into time intervals.
//
// Description:
//
// Splits the target values for a traffic control target into time intervals.
//
// @param request - SplitTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitTrafficControlTargetResponse
func (client *Client) SplitTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *SplitTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SplitTrafficControlTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SetPoints) {
		body["SetPoints"] = request.SetPoints
	}

	if !dara.IsNil(request.SetValues) {
		body["SetValues"] = request.SetValues
	}

	if !dara.IsNil(request.TimePoints) {
		body["TimePoints"] = request.TimePoints
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SplitTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/action/split"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SplitTrafficControlTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Splits the target values for a traffic control target into time intervals.
//
// Description:
//
// Splits the target values for a traffic control target into time intervals.
//
// @param request - SplitTrafficControlTargetRequest
//
// @return SplitTrafficControlTargetResponse
func (client *Client) SplitTrafficControlTarget(TrafficControlTargetId *string, request *SplitTrafficControlTargetRequest) (_result *SplitTrafficControlTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SplitTrafficControlTargetResponse{}
	_body, _err := client.SplitTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a traffic control task for a specific traffic control target.
//
// Description:
//
// ## Request
//
// Call this operation to start a traffic control task by providing the `TrafficControlTargetId` and `InstanceId`.
//
// @param request - StartTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTrafficControlTargetResponse
func (client *Client) StartTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *StartTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartTrafficControlTargetResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/action/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTrafficControlTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a traffic control task for a specific traffic control target.
//
// Description:
//
// ## Request
//
// Call this operation to start a traffic control task by providing the `TrafficControlTargetId` and `InstanceId`.
//
// @param request - StartTrafficControlTargetRequest
//
// @return StartTrafficControlTargetResponse
func (client *Client) StartTrafficControlTarget(TrafficControlTargetId *string, request *StartTrafficControlTargetRequest) (_result *StartTrafficControlTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTrafficControlTargetResponse{}
	_body, _err := client.StartTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a traffic control task with a specified ID for instances in different environments.
//
// Description:
//
// ## Request details
//
// - This operation starts the traffic control task identified by `TrafficControlTaskId`.
//
// - `InstanceId` specifies the target instance.
//
// - `Environment` specifies the target environment. Valid values: Daily, Pre, and Prod.
//
// - Ensure that all required parameters are set correctly before you call this operation. The specified `TrafficControlTaskId` must exist in the system.
//
// @param request - StartTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTrafficControlTaskResponse
func (client *Client) StartTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *StartTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartTrafficControlTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a traffic control task with a specified ID for instances in different environments.
//
// Description:
//
// ## Request details
//
// - This operation starts the traffic control task identified by `TrafficControlTaskId`.
//
// - `InstanceId` specifies the target instance.
//
// - `Environment` specifies the target environment. Valid values: Daily, Pre, and Prod.
//
// - Ensure that all required parameters are set correctly before you call this operation. The specified `TrafficControlTaskId` must exist in the system.
//
// @param request - StartTrafficControlTaskRequest
//
// @return StartTrafficControlTaskResponse
func (client *Client) StartTrafficControlTask(TrafficControlTaskId *string, request *StartTrafficControlTaskRequest) (_result *StartTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTrafficControlTaskResponse{}
	_body, _err := client.StartTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止样本一致性任务
//
// @param request - StopSampleConsistencyJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSampleConsistencyJobResponse
func (client *Client) StopSampleConsistencyJobWithOptions(SampleConsistencyJobId *string, request *StopSampleConsistencyJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopSampleConsistencyJobResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSampleConsistencyJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/sampleconsistencyjobs/" + dara.PercentEncode(dara.StringValue(SampleConsistencyJobId)) + "/action/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopSampleConsistencyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止样本一致性任务
//
// @param request - StopSampleConsistencyJobRequest
//
// @return StopSampleConsistencyJobResponse
func (client *Client) StopSampleConsistencyJob(SampleConsistencyJobId *string, request *StopSampleConsistencyJobRequest) (_result *StopSampleConsistencyJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopSampleConsistencyJobResponse{}
	_body, _err := client.StopSampleConsistencyJobWithOptions(SampleConsistencyJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a traffic control Flink task with the specified ID.
//
// Description:
//
// ## Request description
//
// You can call this operation to stop a specific traffic control Flink task based on the specified TrafficControlTaskId. Make sure that you have prepared the correct InstanceId and the environment to which the instance belongs (Daily for daily environment, Pre for staging environment, Prod for production environment). Include this information in the request body to ensure that the operation is correctly performed.
//
// @param request - StopTrafficControlFlinkTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrafficControlFlinkTaskResponse
func (client *Client) StopTrafficControlFlinkTaskWithOptions(TrafficControlTaskId *string, request *StopTrafficControlFlinkTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTrafficControlFlinkTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTrafficControlFlinkTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/stopflink"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTrafficControlFlinkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a traffic control Flink task with the specified ID.
//
// Description:
//
// ## Request description
//
// You can call this operation to stop a specific traffic control Flink task based on the specified TrafficControlTaskId. Make sure that you have prepared the correct InstanceId and the environment to which the instance belongs (Daily for daily environment, Pre for staging environment, Prod for production environment). Include this information in the request body to ensure that the operation is correctly performed.
//
// @param request - StopTrafficControlFlinkTaskRequest
//
// @return StopTrafficControlFlinkTaskResponse
func (client *Client) StopTrafficControlFlinkTask(TrafficControlTaskId *string, request *StopTrafficControlFlinkTaskRequest) (_result *StopTrafficControlFlinkTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTrafficControlFlinkTaskResponse{}
	_body, _err := client.StopTrafficControlFlinkTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a specific traffic control target.
//
// Description:
//
// ## Request
//
// This operation stops a traffic control task using the provided `TrafficControlTargetId` and `InstanceId`. Ensure that the parameter values are accurate to avoid stopping the wrong target or instance.
//
// @param request - StopTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrafficControlTargetResponse
func (client *Client) StopTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *StopTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTrafficControlTargetResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId)) + "/action/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTrafficControlTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a specific traffic control target.
//
// Description:
//
// ## Request
//
// This operation stops a traffic control task using the provided `TrafficControlTargetId` and `InstanceId`. Ensure that the parameter values are accurate to avoid stopping the wrong target or instance.
//
// @param request - StopTrafficControlTargetRequest
//
// @return StopTrafficControlTargetResponse
func (client *Client) StopTrafficControlTarget(TrafficControlTargetId *string, request *StopTrafficControlTargetRequest) (_result *StopTrafficControlTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTrafficControlTargetResponse{}
	_body, _err := client.StopTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a traffic control task for a specific instance and environment.
//
// Description:
//
// ## Usage notes
//
// - This API stops a traffic control task identified by a specific `TrafficControlTaskId`.
//
// - The `InstanceId` and `Environment` parameters are required to identify the target instance and its environment.
//
// - Ensure that you provide the correct `TrafficControlTaskId` to prevent the request from failing.
//
// @param request - StopTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrafficControlTaskResponse
func (client *Client) StopTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *StopTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTrafficControlTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a traffic control task for a specific instance and environment.
//
// Description:
//
// ## Usage notes
//
// - This API stops a traffic control task identified by a specific `TrafficControlTaskId`.
//
// - The `InstanceId` and `Environment` parameters are required to identify the target instance and its environment.
//
// - Ensure that you provide the correct `TrafficControlTaskId` to prevent the request from failing.
//
// @param request - StopTrafficControlTaskRequest
//
// @return StopTrafficControlTaskResponse
func (client *Client) StopTrafficControlTask(TrafficControlTaskId *string, request *StopTrafficControlTaskRequest) (_result *StopTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTrafficControlTaskResponse{}
	_body, _err := client.StopTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Syncs the replay log for a feature consistency check job.
//
// @param request - SyncFeatureConsistencyCheckJobReplayLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncFeatureConsistencyCheckJobReplayLogResponse
func (client *Client) SyncFeatureConsistencyCheckJobReplayLogWithOptions(request *SyncFeatureConsistencyCheckJobReplayLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncFeatureConsistencyCheckJobReplayLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextFeatures) {
		body["ContextFeatures"] = request.ContextFeatures
	}

	if !dara.IsNil(request.FeatureConsistencyCheckJobConfigId) {
		body["FeatureConsistencyCheckJobConfigId"] = request.FeatureConsistencyCheckJobConfigId
	}

	if !dara.IsNil(request.GeneratedFeatures) {
		body["GeneratedFeatures"] = request.GeneratedFeatures
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogItemId) {
		body["LogItemId"] = request.LogItemId
	}

	if !dara.IsNil(request.LogRequestId) {
		body["LogRequestId"] = request.LogRequestId
	}

	if !dara.IsNil(request.LogRequestTime) {
		body["LogRequestTime"] = request.LogRequestTime
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.RawFeatures) {
		body["RawFeatures"] = request.RawFeatures
	}

	if !dara.IsNil(request.SceneName) {
		body["SceneName"] = request.SceneName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncFeatureConsistencyCheckJobReplayLog"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/action/syncreplaylog"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncFeatureConsistencyCheckJobReplayLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Syncs the replay log for a feature consistency check job.
//
// @param request - SyncFeatureConsistencyCheckJobReplayLogRequest
//
// @return SyncFeatureConsistencyCheckJobReplayLogResponse
func (client *Client) SyncFeatureConsistencyCheckJobReplayLog(request *SyncFeatureConsistencyCheckJobReplayLogRequest) (_result *SyncFeatureConsistencyCheckJobReplayLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncFeatureConsistencyCheckJobReplayLogResponse{}
	_body, _err := client.SyncFeatureConsistencyCheckJobReplayLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates a feature consistency check job.
//
// @param request - TerminateFeatureConsistencyCheckJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateFeatureConsistencyCheckJobResponse
func (client *Client) TerminateFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId *string, request *TerminateFeatureConsistencyCheckJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TerminateFeatureConsistencyCheckJobResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateFeatureConsistencyCheckJob"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobId)) + "/action/terminate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a feature consistency check job.
//
// @param request - TerminateFeatureConsistencyCheckJobRequest
//
// @return TerminateFeatureConsistencyCheckJobResponse
func (client *Client) TerminateFeatureConsistencyCheckJob(FeatureConsistencyCheckJobId *string, request *TerminateFeatureConsistencyCheckJobRequest) (_result *TerminateFeatureConsistencyCheckJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TerminateFeatureConsistencyCheckJobResponse{}
	_body, _err := client.TerminateFeatureConsistencyCheckJobWithOptions(FeatureConsistencyCheckJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the metric configuration of an existing ABTest experiment.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to update the attributes of a specified ABTest metric, including whether to calculate significance and the aggregation method. Make sure that you have obtained the correct `ABMetricId` before calling this operation.
//
// - `NeedSignificance`: Specifies whether to perform significance analysis on the current metric. Default value: `false`.
//
// - `AggregationByUser`: When significance calculation is enabled, specifies whether to aggregate data by user or by sample. Default value: `false` (by sample).
//
// - `Numerator` and `Denominator`: The specific definitions of the numerator and denominator used in significance calculation.
//
// - `IsBinomialDistribution`: Valid only for derived metrics. Specifies whether the metric follows a binomial distribution, which affects subsequent data processing logic.
//
// Note: You do not need to provide all fields at the same time. Include only the parameters whose values you want to change in the request body.
//
// @param request - UpdateABMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABMetricResponse
func (client *Client) UpdateABMetricWithOptions(ABMetricId *string, request *UpdateABMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AggregationByUser) {
		body["AggregationByUser"] = request.AggregationByUser
	}

	if !dara.IsNil(request.Definition) {
		body["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Denominator) {
		body["Denominator"] = request.Denominator
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsBinomialDistribution) {
		body["IsBinomialDistribution"] = request.IsBinomialDistribution
	}

	if !dara.IsNil(request.LeftMetricId) {
		body["LeftMetricId"] = request.LeftMetricId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedSignificance) {
		body["NeedSignificance"] = request.NeedSignificance
	}

	if !dara.IsNil(request.Numerator) {
		body["Numerator"] = request.Numerator
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.ResultResourceId) {
		body["ResultResourceId"] = request.ResultResourceId
	}

	if !dara.IsNil(request.RightMetricId) {
		body["RightMetricId"] = request.RightMetricId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.StatisticsCycle) {
		body["StatisticsCycle"] = request.StatisticsCycle
	}

	if !dara.IsNil(request.TableMetaId) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateABMetric"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetrics/" + dara.PercentEncode(dara.StringValue(ABMetricId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the metric configuration of an existing ABTest experiment.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to update the attributes of a specified ABTest metric, including whether to calculate significance and the aggregation method. Make sure that you have obtained the correct `ABMetricId` before calling this operation.
//
// - `NeedSignificance`: Specifies whether to perform significance analysis on the current metric. Default value: `false`.
//
// - `AggregationByUser`: When significance calculation is enabled, specifies whether to aggregate data by user or by sample. Default value: `false` (by sample).
//
// - `Numerator` and `Denominator`: The specific definitions of the numerator and denominator used in significance calculation.
//
// - `IsBinomialDistribution`: Valid only for derived metrics. Specifies whether the metric follows a binomial distribution, which affects subsequent data processing logic.
//
// Note: You do not need to provide all fields at the same time. Include only the parameters whose values you want to change in the request body.
//
// @param request - UpdateABMetricRequest
//
// @return UpdateABMetricResponse
func (client *Client) UpdateABMetric(ABMetricId *string, request *UpdateABMetricRequest) (_result *UpdateABMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABMetricResponse{}
	_body, _err := client.UpdateABMetricWithOptions(ABMetricId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an A/B test metric group.
//
// @param request - UpdateABMetricGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABMetricGroupResponse
func (client *Client) UpdateABMetricGroupWithOptions(ABMetricGroupId *string, request *UpdateABMetricGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABMetricGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ABMetricIds) {
		body["ABMetricIds"] = request.ABMetricIds
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

	if !dara.IsNil(request.Realtime) {
		body["Realtime"] = request.Realtime
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateABMetricGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/abmetricgroups/" + dara.PercentEncode(dara.StringValue(ABMetricGroupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABMetricGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an A/B test metric group.
//
// @param request - UpdateABMetricGroupRequest
//
// @return UpdateABMetricGroupResponse
func (client *Client) UpdateABMetricGroup(ABMetricGroupId *string, request *UpdateABMetricGroupRequest) (_result *UpdateABMetricGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABMetricGroupResponse{}
	_body, _err := client.UpdateABMetricGroupWithOptions(ABMetricGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a crowd\\"s information, such as its name and description.
//
// @param request - UpdateCrowdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCrowdResponse
func (client *Client) UpdateCrowdWithOptions(CrowdId *string, request *UpdateCrowdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCrowdResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCrowd"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/crowds/" + dara.PercentEncode(dara.StringValue(CrowdId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a crowd\\"s information, such as its name and description.
//
// @param request - UpdateCrowdRequest
//
// @return UpdateCrowdResponse
func (client *Client) UpdateCrowd(CrowdId *string, request *UpdateCrowdRequest) (_result *UpdateCrowdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCrowdResponse{}
	_body, _err := client.UpdateCrowdWithOptions(CrowdId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specified data diagnosis task.
//
// Description:
//
// ## Request
//
// This API updates the configuration of an existing data diagnosis task, including the instance ID, task name, task type, and specific configuration content. Provide the `DataDiagnosisId` in the request path to identify the task to update. You must also specify the `Config` parameter based on the task `Type`. For periodic runs, set the execution time in the `CycleTime` field. If a periodic run is not required, omit this field.
//
// ## Usage notes
//
// - `DataDiagnosisId` is a required path parameter that uniquely identifies a data diagnosis task.
//
// - The structure of the `Config` field varies depending on the value of `Type`. Refer to the examples in this document for configuration details.
//
// - To disable periodic runs, omit the `CycleTime` field.
//
// - When updating a task for two-table join analysis (`JoinTables`), provide the information for the left and right tables, including `LeftTableMetaId` and `RightTableMetaId`.
//
// - The `InstanceId`, `Name`, and `Type` parameters are required for all types of data diagnosis tasks.
//
// @param request - UpdateDataDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataDiagnosisResponse
func (client *Client) UpdateDataDiagnosisWithOptions(DataDiagnosisId *string, request *UpdateDataDiagnosisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDataDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.CycleTime) {
		body["CycleTime"] = request.CycleTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LeftTableMetaId) {
		body["LeftTableMetaId"] = request.LeftTableMetaId
	}

	if !dara.IsNil(request.LeftTablePartitionField) {
		body["LeftTablePartitionField"] = request.LeftTablePartitionField
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PartitionField) {
		body["PartitionField"] = request.PartitionField
	}

	if !dara.IsNil(request.RightTableMetaId) {
		body["RightTableMetaId"] = request.RightTableMetaId
	}

	if !dara.IsNil(request.RightTablePartitionField) {
		body["RightTablePartitionField"] = request.RightTablePartitionField
	}

	if !dara.IsNil(request.TableMetaId) {
		body["TableMetaId"] = request.TableMetaId
	}

	if !dara.IsNil(request.TopNQuantity) {
		body["TopNQuantity"] = request.TopNQuantity
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataDiagnosis"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datadiagnoses/" + dara.PercentEncode(dara.StringValue(DataDiagnosisId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specified data diagnosis task.
//
// Description:
//
// ## Request
//
// This API updates the configuration of an existing data diagnosis task, including the instance ID, task name, task type, and specific configuration content. Provide the `DataDiagnosisId` in the request path to identify the task to update. You must also specify the `Config` parameter based on the task `Type`. For periodic runs, set the execution time in the `CycleTime` field. If a periodic run is not required, omit this field.
//
// ## Usage notes
//
// - `DataDiagnosisId` is a required path parameter that uniquely identifies a data diagnosis task.
//
// - The structure of the `Config` field varies depending on the value of `Type`. Refer to the examples in this document for configuration details.
//
// - To disable periodic runs, omit the `CycleTime` field.
//
// - When updating a task for two-table join analysis (`JoinTables`), provide the information for the left and right tables, including `LeftTableMetaId` and `RightTableMetaId`.
//
// - The `InstanceId`, `Name`, and `Type` parameters are required for all types of data diagnosis tasks.
//
// @param request - UpdateDataDiagnosisRequest
//
// @return UpdateDataDiagnosisResponse
func (client *Client) UpdateDataDiagnosis(DataDiagnosisId *string, request *UpdateDataDiagnosisRequest) (_result *UpdateDataDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDataDiagnosisResponse{}
	_body, _err := client.UpdateDataDiagnosisWithOptions(DataDiagnosisId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specific engine configuration.
//
// @param request - UpdateEngineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEngineConfigResponse
func (client *Client) UpdateEngineConfigWithOptions(EngineConfigId *string, request *UpdateEngineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEngineConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEngineConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/engineconfigs/" + dara.PercentEncode(dara.StringValue(EngineConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEngineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specific engine configuration.
//
// @param request - UpdateEngineConfigRequest
//
// @return UpdateEngineConfigResponse
func (client *Client) UpdateEngineConfig(EngineConfigId *string, request *UpdateEngineConfigRequest) (_result *UpdateEngineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEngineConfigResponse{}
	_body, _err := client.UpdateEngineConfigWithOptions(EngineConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the properties of a specified experiment, such as its name.
//
// @param request - UpdateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperimentWithOptions(ExperimentId *string, request *UpdateExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowPercent) {
		body["FlowPercent"] = request.FlowPercent
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperiment"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the properties of a specified experiment, such as its name.
//
// @param request - UpdateExperimentRequest
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperiment(ExperimentId *string, request *UpdateExperimentRequest) (_result *UpdateExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentResponse{}
	_body, _err := client.UpdateExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information for a specified experiment group, such as its name and description.
//
// @param request - UpdateExperimentGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentGroupResponse
func (client *Client) UpdateExperimentGroupWithOptions(ExperimentGroupId *string, request *UpdateExperimentGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.CrowdId) {
		body["CrowdId"] = request.CrowdId
	}

	if !dara.IsNil(request.CrowdTargetType) {
		body["CrowdTargetType"] = request.CrowdTargetType
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DistributionTimeDuration) {
		body["DistributionTimeDuration"] = request.DistributionTimeDuration
	}

	if !dara.IsNil(request.DistributionType) {
		body["DistributionType"] = request.DistributionType
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LayerId) {
		body["LayerId"] = request.LayerId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedAA) {
		body["NeedAA"] = request.NeedAA
	}

	if !dara.IsNil(request.RandomFlow) {
		body["RandomFlow"] = request.RandomFlow
	}

	if !dara.IsNil(request.ReservcedBuckets) {
		body["ReservcedBuckets"] = request.ReservcedBuckets
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperimentGroup"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experimentgroups/" + dara.PercentEncode(dara.StringValue(ExperimentGroupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information for a specified experiment group, such as its name and description.
//
// @param request - UpdateExperimentGroupRequest
//
// @return UpdateExperimentGroupResponse
func (client *Client) UpdateExperimentGroup(ExperimentGroupId *string, request *UpdateExperimentGroupRequest) (_result *UpdateExperimentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentGroupResponse{}
	_body, _err := client.UpdateExperimentGroupWithOptions(ExperimentGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration details of a feature consistency check task, such as the name.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to update the configuration of an existing feature consistency check task. By providing new configuration parameters, you can modify multiple properties including the instance ID, name, and scene ID. Ensure that all required parameters are included in the request, and provide optional parameters as needed.
//
// - **FeatureConsistencyCheckJobConfigId*	- is a path parameter that specifies the feature consistency check task to update.
//
// - All other parameters are in the request body. Some are required (such as InstanceId and Name), and the rest are optional.
//
// - The SampleRate value must be a floating-point number between 0 and 1, which indicates the sampling ratio.
//
// - If you use FeatureStore-related features, make sure that you correctly set the IsUseFeatureStore flag and the related FeatureStore	- fields.
//
// - For network configuration parameters (such as VpcId and SwitchId), make sure that the values match your Alibaba Cloud environment.
//
// @param request - UpdateFeatureConsistencyCheckJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFeatureConsistencyCheckJobConfigResponse
func (client *Client) UpdateFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId *string, request *UpdateFeatureConsistencyCheckJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFeatureConsistencyCheckJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CompareFeature) {
		body["CompareFeature"] = request.CompareFeature
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetMountPath) {
		body["DatasetMountPath"] = request.DatasetMountPath
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		body["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.DatasetUri) {
		body["DatasetUri"] = request.DatasetUri
	}

	if !dara.IsNil(request.DefaultRoute) {
		body["DefaultRoute"] = request.DefaultRoute
	}

	if !dara.IsNil(request.EasServiceName) {
		body["EasServiceName"] = request.EasServiceName
	}

	if !dara.IsNil(request.EasyRecPackagePath) {
		body["EasyRecPackagePath"] = request.EasyRecPackagePath
	}

	if !dara.IsNil(request.EasyRecVersion) {
		body["EasyRecVersion"] = request.EasyRecVersion
	}

	if !dara.IsNil(request.FeatureDisplayExclude) {
		body["FeatureDisplayExclude"] = request.FeatureDisplayExclude
	}

	if !dara.IsNil(request.FeatureLandingResourceId) {
		body["FeatureLandingResourceId"] = request.FeatureLandingResourceId
	}

	if !dara.IsNil(request.FeaturePriority) {
		body["FeaturePriority"] = request.FeaturePriority
	}

	if !dara.IsNil(request.FeatureStoreItemId) {
		body["FeatureStoreItemId"] = request.FeatureStoreItemId
	}

	if !dara.IsNil(request.FeatureStoreModelId) {
		body["FeatureStoreModelId"] = request.FeatureStoreModelId
	}

	if !dara.IsNil(request.FeatureStoreProjectId) {
		body["FeatureStoreProjectId"] = request.FeatureStoreProjectId
	}

	if !dara.IsNil(request.FeatureStoreProjectName) {
		body["FeatureStoreProjectName"] = request.FeatureStoreProjectName
	}

	if !dara.IsNil(request.FeatureStoreSeqFeatureView) {
		body["FeatureStoreSeqFeatureView"] = request.FeatureStoreSeqFeatureView
	}

	if !dara.IsNil(request.FeatureStoreUserId) {
		body["FeatureStoreUserId"] = request.FeatureStoreUserId
	}

	if !dara.IsNil(request.FgJarVersion) {
		body["FgJarVersion"] = request.FgJarVersion
	}

	if !dara.IsNil(request.FgJsonFileName) {
		body["FgJsonFileName"] = request.FgJsonFileName
	}

	if !dara.IsNil(request.GenerateZip) {
		body["GenerateZip"] = request.GenerateZip
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsUseFeatureStore) {
		body["IsUseFeatureStore"] = request.IsUseFeatureStore
	}

	if !dara.IsNil(request.ItemIdField) {
		body["ItemIdField"] = request.ItemIdField
	}

	if !dara.IsNil(request.ItemTable) {
		body["ItemTable"] = request.ItemTable
	}

	if !dara.IsNil(request.ItemTablePartitionField) {
		body["ItemTablePartitionField"] = request.ItemTablePartitionField
	}

	if !dara.IsNil(request.ItemTablePartitionFieldFormat) {
		body["ItemTablePartitionFieldFormat"] = request.ItemTablePartitionFieldFormat
	}

	if !dara.IsNil(request.MaxcomputeSchema) {
		body["MaxcomputeSchema"] = request.MaxcomputeSchema
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OssResourceId) {
		body["OssResourceId"] = request.OssResourceId
	}

	if !dara.IsNil(request.PredictWorkerCount) {
		body["PredictWorkerCount"] = request.PredictWorkerCount
	}

	if !dara.IsNil(request.PredictWorkerCpu) {
		body["PredictWorkerCpu"] = request.PredictWorkerCpu
	}

	if !dara.IsNil(request.PredictWorkerMemory) {
		body["PredictWorkerMemory"] = request.PredictWorkerMemory
	}

	if !dara.IsNil(request.ResourceConfig) {
		body["ResourceConfig"] = request.ResourceConfig
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SwitchId) {
		body["SwitchId"] = request.SwitchId
	}

	if !dara.IsNil(request.UserIdField) {
		body["UserIdField"] = request.UserIdField
	}

	if !dara.IsNil(request.UserTable) {
		body["UserTable"] = request.UserTable
	}

	if !dara.IsNil(request.UserTablePartitionField) {
		body["UserTablePartitionField"] = request.UserTablePartitionField
	}

	if !dara.IsNil(request.UserTablePartitionFieldFormat) {
		body["UserTablePartitionFieldFormat"] = request.UserTablePartitionFieldFormat
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkflowName) {
		body["WorkflowName"] = request.WorkflowName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFeatureConsistencyCheckJobConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/featureconsistencycheck/jobconfigs/" + dara.PercentEncode(dara.StringValue(FeatureConsistencyCheckJobConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration details of a feature consistency check task, such as the name.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to update the configuration of an existing feature consistency check task. By providing new configuration parameters, you can modify multiple properties including the instance ID, name, and scene ID. Ensure that all required parameters are included in the request, and provide optional parameters as needed.
//
// - **FeatureConsistencyCheckJobConfigId*	- is a path parameter that specifies the feature consistency check task to update.
//
// - All other parameters are in the request body. Some are required (such as InstanceId and Name), and the rest are optional.
//
// - The SampleRate value must be a floating-point number between 0 and 1, which indicates the sampling ratio.
//
// - If you use FeatureStore-related features, make sure that you correctly set the IsUseFeatureStore flag and the related FeatureStore	- fields.
//
// - For network configuration parameters (such as VpcId and SwitchId), make sure that the values match your Alibaba Cloud environment.
//
// @param request - UpdateFeatureConsistencyCheckJobConfigRequest
//
// @return UpdateFeatureConsistencyCheckJobConfigResponse
func (client *Client) UpdateFeatureConsistencyCheckJobConfig(FeatureConsistencyCheckJobConfigId *string, request *UpdateFeatureConsistencyCheckJobConfigRequest) (_result *UpdateFeatureConsistencyCheckJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFeatureConsistencyCheckJobConfigResponse{}
	_body, _err := client.UpdateFeatureConsistencyCheckJobConfigWithOptions(FeatureConsistencyCheckJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specified resource for a specified instance.
//
// @param request - UpdateInstanceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResourceResponse
func (client *Client) UpdateInstanceResourceWithOptions(InstanceId *string, ResourceId *string, request *UpdateInstanceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceResource"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/resources/" + dara.PercentEncode(dara.StringValue(ResourceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified resource for a specified instance.
//
// @param request - UpdateInstanceResourceRequest
//
// @return UpdateInstanceResourceResponse
func (client *Client) UpdateInstanceResource(InstanceId *string, ResourceId *string, request *UpdateInstanceResourceRequest) (_result *UpdateInstanceResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResourceResponse{}
	_body, _err := client.UpdateInstanceResourceWithOptions(InstanceId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a laboratory\\"s information, such as its name.
//
// @param request - UpdateLaboratoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLaboratoryResponse
func (client *Client) UpdateLaboratoryWithOptions(LaboratoryId *string, request *UpdateLaboratoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLaboratoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BucketCount) {
		body["BucketCount"] = request.BucketCount
	}

	if !dara.IsNil(request.BucketType) {
		body["BucketType"] = request.BucketType
	}

	if !dara.IsNil(request.Buckets) {
		body["Buckets"] = request.Buckets
	}

	if !dara.IsNil(request.DebugCrowdId) {
		body["DebugCrowdId"] = request.DebugCrowdId
	}

	if !dara.IsNil(request.DebugUsers) {
		body["DebugUsers"] = request.DebugUsers
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLaboratory"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/laboratories/" + dara.PercentEncode(dara.StringValue(LaboratoryId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLaboratoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a laboratory\\"s information, such as its name.
//
// @param request - UpdateLaboratoryRequest
//
// @return UpdateLaboratoryResponse
func (client *Client) UpdateLaboratory(LaboratoryId *string, request *UpdateLaboratoryRequest) (_result *UpdateLaboratoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLaboratoryResponse{}
	_body, _err := client.UpdateLaboratoryWithOptions(LaboratoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name and description of a specified layer.
//
// @param request - UpdateLayerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLayerResponse
func (client *Client) UpdateLayerWithOptions(LayerId *string, request *UpdateLayerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLayerResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLayer"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/layers/" + dara.PercentEncode(dara.StringValue(LayerId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name and description of a specified layer.
//
// @param request - UpdateLayerRequest
//
// @return UpdateLayerResponse
func (client *Client) UpdateLayer(LayerId *string, request *UpdateLayerRequest) (_result *UpdateLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLayerResponse{}
	_body, _err := client.UpdateLayerWithOptions(LayerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information for a specified parameter, such as its value.
//
// @param request - UpdateParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParamResponse
func (client *Client) UpdateParamWithOptions(ParamId *string, request *UpdateParamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateParamResponse, _err error) {
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

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateParam"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/params/" + dara.PercentEncode(dara.StringValue(ParamId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information for a specified parameter, such as its value.
//
// @param request - UpdateParamRequest
//
// @return UpdateParamResponse
func (client *Client) UpdateParam(ParamId *string, request *UpdateParamRequest) (_result *UpdateParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateParamResponse{}
	_body, _err := client.UpdateParamWithOptions(ParamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the recall management configuration, including the instance ID, password, and network configuration.
//
// Description:
//
// ## Request
//
// - `InstanceId` is required. It specifies the instance to update.
//
// - `Password` and `NetworkConfigs` are optional.
//
// - Use `NetworkConfigs` to define the network by specifying the Virtual Private Cloud (VPC) ID (`VpcId`) and mapping availability zones to VSwitch IDs (`VswitchIds`).
//
// - Note: Ensure that sensitive information, such as the password, is transmitted securely.
//
// @param request - UpdateRecallManagementConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecallManagementConfigResponse
func (client *Client) UpdateRecallManagementConfigWithOptions(request *UpdateRecallManagementConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRecallManagementConfigResponse, _err error) {
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

	if !dara.IsNil(request.NetworkConfigs) {
		body["NetworkConfigs"] = request.NetworkConfigs
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecallManagementConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementconfigs"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecallManagementConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the recall management configuration, including the instance ID, password, and network configuration.
//
// Description:
//
// ## Request
//
// - `InstanceId` is required. It specifies the instance to update.
//
// - `Password` and `NetworkConfigs` are optional.
//
// - Use `NetworkConfigs` to define the network by specifying the Virtual Private Cloud (VPC) ID (`VpcId`) and mapping availability zones to VSwitch IDs (`VswitchIds`).
//
// - Note: Ensure that sensitive information, such as the password, is transmitted securely.
//
// @param request - UpdateRecallManagementConfigRequest
//
// @return UpdateRecallManagementConfigResponse
func (client *Client) UpdateRecallManagementConfig(request *UpdateRecallManagementConfigRequest) (_result *UpdateRecallManagementConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRecallManagementConfigResponse{}
	_body, _err := client.UpdateRecallManagementConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the instance ID and description of a specified recall management service.
//
// Description:
//
// ## Request description
//
// This operation updates the instance ID and description of a specific recall management service. Make sure to specify the `InstanceId` and `Description` fields in the request body.
//
// - **RecallManagementServiceId**: The unique identifier of the recall management service.
//
// - **InstanceId**: The instance ID to associate with this recall management service.
//
// - **Description**: A new description for the recall management service.
//
// Note: You must provide all required parameters, or the update may fail.
//
// @param request - UpdateRecallManagementServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecallManagementServiceResponse
func (client *Client) UpdateRecallManagementServiceWithOptions(RecallManagementServiceId *string, request *UpdateRecallManagementServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRecallManagementServiceResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecallManagementService"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecallManagementServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the instance ID and description of a specified recall management service.
//
// Description:
//
// ## Request description
//
// This operation updates the instance ID and description of a specific recall management service. Make sure to specify the `InstanceId` and `Description` fields in the request body.
//
// - **RecallManagementServiceId**: The unique identifier of the recall management service.
//
// - **InstanceId**: The instance ID to associate with this recall management service.
//
// - **Description**: A new description for the recall management service.
//
// Note: You must provide all required parameters, or the update may fail.
//
// @param request - UpdateRecallManagementServiceRequest
//
// @return UpdateRecallManagementServiceResponse
func (client *Client) UpdateRecallManagementService(RecallManagementServiceId *string, request *UpdateRecallManagementServiceRequest) (_result *UpdateRecallManagementServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRecallManagementServiceResponse{}
	_body, _err := client.UpdateRecallManagementServiceWithOptions(RecallManagementServiceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specific Recall Management Service version.
//
// Description:
//
// ## Request
//
// This API updates the recall and merge configurations for a specific recall management service version. Your request must include the correct `InstanceId` and the configurations to update. Refer to the parameter descriptions for details on required parameters.
//
// @param request - UpdateRecallManagementServiceVersionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecallManagementServiceVersionConfigResponse
func (client *Client) UpdateRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, RecallManagementServiceVersionConfigId *string, request *UpdateRecallManagementServiceVersionConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRecallManagementServiceVersionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigType) {
		body["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MergeConfig) {
		body["MergeConfig"] = request.MergeConfig
	}

	if !dara.IsNil(request.RecallConfig) {
		body["RecallConfig"] = request.RecallConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecallManagementServiceVersionConfig"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementservices/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionId)) + "/configs/" + dara.PercentEncode(dara.StringValue(RecallManagementServiceVersionConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specific Recall Management Service version.
//
// Description:
//
// ## Request
//
// This API updates the recall and merge configurations for a specific recall management service version. Your request must include the correct `InstanceId` and the configurations to update. Refer to the parameter descriptions for details on required parameters.
//
// @param request - UpdateRecallManagementServiceVersionConfigRequest
//
// @return UpdateRecallManagementServiceVersionConfigResponse
func (client *Client) UpdateRecallManagementServiceVersionConfig(RecallManagementServiceId *string, RecallManagementServiceVersionId *string, RecallManagementServiceVersionConfigId *string, request *UpdateRecallManagementServiceVersionConfigRequest) (_result *UpdateRecallManagementServiceVersionConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRecallManagementServiceVersionConfigResponse{}
	_body, _err := client.UpdateRecallManagementServiceVersionConfigWithOptions(RecallManagementServiceId, RecallManagementServiceVersionId, RecallManagementServiceVersionConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a recall management table specified by its ID.
//
// Description:
//
// ## Request details
//
// - Updates the recall management table specified by `RecallManagementTableId`.
//
// - You can enable fluctuation thresholds for the row count or data size and define the specific ranges for these thresholds.
//
// - You can add or modify fields in the table, including their names, types, and attributes.
//
// - The `InstanceId` parameter is required and identifies the specific instance.
//
// - For vector-related fields, you can also specify the vector dimension and metric type.
//
// - Note: Optional parameters in the request body selectively update the target table.
//
// @param request - UpdateRecallManagementTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecallManagementTableResponse
func (client *Client) UpdateRecallManagementTableWithOptions(RecallManagementTableId *string, request *UpdateRecallManagementTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRecallManagementTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableDataSizeFluctuationThreshold) {
		body["EnableDataSizeFluctuationThreshold"] = request.EnableDataSizeFluctuationThreshold
	}

	if !dara.IsNil(request.EnableRowCountFluctuationThreshold) {
		body["EnableRowCountFluctuationThreshold"] = request.EnableRowCountFluctuationThreshold
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.IndexVersionId) {
		body["IndexVersionId"] = request.IndexVersionId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxDataSizeFluctuationThreshold) {
		body["MaxDataSizeFluctuationThreshold"] = request.MaxDataSizeFluctuationThreshold
	}

	if !dara.IsNil(request.MaxRowCountFluctuationThreshold) {
		body["MaxRowCountFluctuationThreshold"] = request.MaxRowCountFluctuationThreshold
	}

	if !dara.IsNil(request.MinDataSizeFluctuationThreshold) {
		body["MinDataSizeFluctuationThreshold"] = request.MinDataSizeFluctuationThreshold
	}

	if !dara.IsNil(request.MinRowCountFluctuationThreshold) {
		body["MinRowCountFluctuationThreshold"] = request.MinRowCountFluctuationThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecallManagementTable"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recallmanagementtables/" + dara.PercentEncode(dara.StringValue(RecallManagementTableId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecallManagementTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a recall management table specified by its ID.
//
// Description:
//
// ## Request details
//
// - Updates the recall management table specified by `RecallManagementTableId`.
//
// - You can enable fluctuation thresholds for the row count or data size and define the specific ranges for these thresholds.
//
// - You can add or modify fields in the table, including their names, types, and attributes.
//
// - The `InstanceId` parameter is required and identifies the specific instance.
//
// - For vector-related fields, you can also specify the vector dimension and metric type.
//
// - Note: Optional parameters in the request body selectively update the target table.
//
// @param request - UpdateRecallManagementTableRequest
//
// @return UpdateRecallManagementTableResponse
func (client *Client) UpdateRecallManagementTable(RecallManagementTableId *string, request *UpdateRecallManagementTableRequest) (_result *UpdateRecallManagementTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRecallManagementTableResponse{}
	_body, _err := client.UpdateRecallManagementTableWithOptions(RecallManagementTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - UpdateResourceRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceRuleResponse
func (client *Client) UpdateResourceRuleWithOptions(ResourceRuleId *string, request *UpdateResourceRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceRuleResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricOperationType) {
		body["MetricOperationType"] = request.MetricOperationType
	}

	if !dara.IsNil(request.MetricPullInfo) {
		body["MetricPullInfo"] = request.MetricPullInfo
	}

	if !dara.IsNil(request.MetricPullPeriod) {
		body["MetricPullPeriod"] = request.MetricPullPeriod
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RuleComputingDefinition) {
		body["RuleComputingDefinition"] = request.RuleComputingDefinition
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceRule"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源规则列表
//
// @param request - UpdateResourceRuleRequest
//
// @return UpdateResourceRuleResponse
func (client *Client) UpdateResourceRule(ResourceRuleId *string, request *UpdateResourceRuleRequest) (_result *UpdateResourceRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceRuleResponse{}
	_body, _err := client.UpdateResourceRuleWithOptions(ResourceRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源规则条目
//
// @param request - UpdateResourceRuleItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceRuleItemResponse
func (client *Client) UpdateResourceRuleItemWithOptions(ResourceRuleId *string, ResourceRuleItemId *string, request *UpdateResourceRuleItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceRuleItemResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxValue) {
		body["MaxValue"] = request.MaxValue
	}

	if !dara.IsNil(request.MinValue) {
		body["MinValue"] = request.MinValue
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceRuleItem"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resourcerules/" + dara.PercentEncode(dara.StringValue(ResourceRuleId)) + "/items/" + dara.PercentEncode(dara.StringValue(ResourceRuleItemId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceRuleItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源规则条目
//
// @param request - UpdateResourceRuleItemRequest
//
// @return UpdateResourceRuleItemResponse
func (client *Client) UpdateResourceRuleItem(ResourceRuleId *string, ResourceRuleItemId *string, request *UpdateResourceRuleItemRequest) (_result *UpdateResourceRuleItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceRuleItemResponse{}
	_body, _err := client.UpdateResourceRuleItemWithOptions(ResourceRuleId, ResourceRuleItemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information for a scene, such as its name and description.
//
// @param request - UpdateSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSceneResponse
func (client *Client) UpdateSceneWithOptions(SceneId *string, request *UpdateSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSceneResponse, _err error) {
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

	if !dara.IsNil(request.Flows) {
		body["Flows"] = request.Flows
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScene"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/scenes/" + dara.PercentEncode(dara.StringValue(SceneId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information for a scene, such as its name and description.
//
// @param request - UpdateSceneRequest
//
// @return UpdateSceneResponse
func (client *Client) UpdateScene(SceneId *string, request *UpdateSceneRequest) (_result *UpdateSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSceneResponse{}
	_body, _err := client.UpdateSceneWithOptions(SceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data table.
//
// @param request - UpdateTableMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableMetaResponse
func (client *Client) UpdateTableMetaWithOptions(TableMetaId *string, request *UpdateTableMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTableMetaResponse, _err error) {
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

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Module) {
		body["Module"] = request.Module
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableMeta"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tablemetas/" + dara.PercentEncode(dara.StringValue(TableMetaId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data table.
//
// @param request - UpdateTableMetaRequest
//
// @return UpdateTableMetaResponse
func (client *Client) UpdateTableMeta(TableMetaId *string, request *UpdateTableMetaRequest) (_result *UpdateTableMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTableMetaResponse{}
	_body, _err := client.UpdateTableMetaWithOptions(TableMetaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a traffic control target, including its control period, conditions, and value.
//
// Description:
//
// ## Description
//
// - Updates a traffic control target specified by its ID.
//
// - `TrafficControlTargetId` is a path parameter that specifies the ID of the traffic control target to update.
//
// - The `ItemConditionType` parameter specifies the format of the item condition, which can be either `Array` or `Expression`. Based on your selection, you must provide a value for either the `ItemConditionArray` or `ItemConditionExpress` parameter.
//
// - If `NewProductRegulation` is set to `true`, the control rule applies to a new product.
//
// - The `StatisPeriod` parameter specifies the statistics period. Valid values are `Daily` and `hourly`.
//
// - Ensure that the time interval between `StartTime` and `EndTime` is reasonable and meets your business requirements.
//
// @param request - UpdateTrafficControlTargetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTargetResponse
func (client *Client) UpdateTrafficControlTargetWithOptions(TrafficControlTargetId *string, request *UpdateTrafficControlTargetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrafficControlTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewParam3) {
		query["new-param-3"] = request.NewParam3
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Event) {
		body["Event"] = request.Event
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NewProductRegulation) {
		body["NewProductRegulation"] = request.NewProductRegulation
	}

	if !dara.IsNil(request.RecallName) {
		body["RecallName"] = request.RecallName
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisPeriod) {
		body["StatisPeriod"] = request.StatisPeriod
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.ToleranceValue) {
		body["ToleranceValue"] = request.ToleranceValue
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrafficControlTarget"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltargets/" + dara.PercentEncode(dara.StringValue(TrafficControlTargetId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrafficControlTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a traffic control target, including its control period, conditions, and value.
//
// Description:
//
// ## Description
//
// - Updates a traffic control target specified by its ID.
//
// - `TrafficControlTargetId` is a path parameter that specifies the ID of the traffic control target to update.
//
// - The `ItemConditionType` parameter specifies the format of the item condition, which can be either `Array` or `Expression`. Based on your selection, you must provide a value for either the `ItemConditionArray` or `ItemConditionExpress` parameter.
//
// - If `NewProductRegulation` is set to `true`, the control rule applies to a new product.
//
// - The `StatisPeriod` parameter specifies the statistics period. Valid values are `Daily` and `hourly`.
//
// - Ensure that the time interval between `StartTime` and `EndTime` is reasonable and meets your business requirements.
//
// @param request - UpdateTrafficControlTargetRequest
//
// @return UpdateTrafficControlTargetResponse
func (client *Client) UpdateTrafficControlTarget(TrafficControlTargetId *string, request *UpdateTrafficControlTargetRequest) (_result *UpdateTrafficControlTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrafficControlTargetResponse{}
	_body, _err := client.UpdateTrafficControlTargetWithOptions(TrafficControlTargetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration and target of a specified traffic control task.
//
// Description:
//
// ## Usage notes
//
// - Use this API to update an existing traffic control task.
//
// - When `ExecutionTime` is set to `TimeRange`, you must also provide `StartTime` and `EndTime`.
//
// - Each element in the `TrafficControlTargets` array is a traffic control target. Ensure each target\\"s time range, condition type, and other information are complete and valid.
//
// - If you set `UserConditionType` or `ItemConditionType` to `Expression`, you must specify the corresponding expression field (for example, `UserConditionExpress`).
//
// - `ServiceIds` and `EffectiveSceneIds` are optional parameters. If you include them, ensure the ID lists are correctly formatted.
//
// - Ensure you complete all required fields to avoid a failed request.
//
// @param request - UpdateTrafficControlTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTaskResponse
func (client *Client) UpdateTrafficControlTaskWithOptions(TrafficControlTaskId *string, request *UpdateTrafficControlTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrafficControlTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BehaviorTableMetaId) {
		body["BehaviorTableMetaId"] = request.BehaviorTableMetaId
	}

	if !dara.IsNil(request.ControlGranularity) {
		body["ControlGranularity"] = request.ControlGranularity
	}

	if !dara.IsNil(request.ControlLogic) {
		body["ControlLogic"] = request.ControlLogic
	}

	if !dara.IsNil(request.ControlType) {
		body["ControlType"] = request.ControlType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EffectiveSceneIds) {
		body["EffectiveSceneIds"] = request.EffectiveSceneIds
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecutionTime) {
		body["ExecutionTime"] = request.ExecutionTime
	}

	if !dara.IsNil(request.FlinkResourceId) {
		body["FlinkResourceId"] = request.FlinkResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemConditionArray) {
		body["ItemConditionArray"] = request.ItemConditionArray
	}

	if !dara.IsNil(request.ItemConditionExpress) {
		body["ItemConditionExpress"] = request.ItemConditionExpress
	}

	if !dara.IsNil(request.ItemConditionType) {
		body["ItemConditionType"] = request.ItemConditionType
	}

	if !dara.IsNil(request.ItemTableMetaId) {
		body["ItemTableMetaId"] = request.ItemTableMetaId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PreExperimentIds) {
		body["PreExperimentIds"] = request.PreExperimentIds
	}

	if !dara.IsNil(request.ProdExperimentIds) {
		body["ProdExperimentIds"] = request.ProdExperimentIds
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceIds) {
		body["ServiceIds"] = request.ServiceIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisBaeaviorConditionArray) {
		body["StatisBaeaviorConditionArray"] = request.StatisBaeaviorConditionArray
	}

	if !dara.IsNil(request.StatisBehaviorConditionArray) {
		body["StatisBehaviorConditionArray"] = request.StatisBehaviorConditionArray
	}

	if !dara.IsNil(request.StatisBehaviorConditionExpress) {
		body["StatisBehaviorConditionExpress"] = request.StatisBehaviorConditionExpress
	}

	if !dara.IsNil(request.StatisBehaviorConditionType) {
		body["StatisBehaviorConditionType"] = request.StatisBehaviorConditionType
	}

	if !dara.IsNil(request.TrafficControlTargets) {
		body["TrafficControlTargets"] = request.TrafficControlTargets
	}

	if !dara.IsNil(request.UserConditionArray) {
		body["UserConditionArray"] = request.UserConditionArray
	}

	if !dara.IsNil(request.UserConditionExpress) {
		body["UserConditionExpress"] = request.UserConditionExpress
	}

	if !dara.IsNil(request.UserConditionType) {
		body["UserConditionType"] = request.UserConditionType
	}

	if !dara.IsNil(request.UserTableMetaId) {
		body["UserTableMetaId"] = request.UserTableMetaId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrafficControlTask"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrafficControlTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration and target of a specified traffic control task.
//
// Description:
//
// ## Usage notes
//
// - Use this API to update an existing traffic control task.
//
// - When `ExecutionTime` is set to `TimeRange`, you must also provide `StartTime` and `EndTime`.
//
// - Each element in the `TrafficControlTargets` array is a traffic control target. Ensure each target\\"s time range, condition type, and other information are complete and valid.
//
// - If you set `UserConditionType` or `ItemConditionType` to `Expression`, you must specify the corresponding expression field (for example, `UserConditionExpress`).
//
// - `ServiceIds` and `EffectiveSceneIds` are optional parameters. If you include them, ensure the ID lists are correctly formatted.
//
// - Ensure you complete all required fields to avoid a failed request.
//
// @param request - UpdateTrafficControlTaskRequest
//
// @return UpdateTrafficControlTaskResponse
func (client *Client) UpdateTrafficControlTask(TrafficControlTaskId *string, request *UpdateTrafficControlTaskRequest) (_result *UpdateTrafficControlTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrafficControlTaskResponse{}
	_body, _err := client.UpdateTrafficControlTaskWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the traffic parameters for a specified traffic control task, including target traffic and actual traffic.
//
// Description:
//
// ## Request
//
// This API updates the traffic configuration for a specific traffic control task. The configuration includes the traffic control target ID, record time, target traffic, and overall traffic. Ensure that the provided`TrafficControlTaskId` is valid and within your permission scope. Additionally, each object in the`Traffics` array must contain the required fields.
//
// @param request - UpdateTrafficControlTaskTrafficRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficControlTaskTrafficResponse
func (client *Client) UpdateTrafficControlTaskTrafficWithOptions(TrafficControlTaskId *string, request *UpdateTrafficControlTaskTrafficRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrafficControlTaskTrafficResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewParam3) {
		query["new-param-3"] = request.NewParam3
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Traffics) {
		body["Traffics"] = request.Traffics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrafficControlTaskTraffic"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trafficcontroltasks/" + dara.PercentEncode(dara.StringValue(TrafficControlTaskId)) + "/action/traffic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrafficControlTaskTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the traffic parameters for a specified traffic control task, including target traffic and actual traffic.
//
// Description:
//
// ## Request
//
// This API updates the traffic configuration for a specific traffic control task. The configuration includes the traffic control target ID, record time, target traffic, and overall traffic. Ensure that the provided`TrafficControlTaskId` is valid and within your permission scope. Additionally, each object in the`Traffics` array must contain the required fields.
//
// @param request - UpdateTrafficControlTaskTrafficRequest
//
// @return UpdateTrafficControlTaskTrafficResponse
func (client *Client) UpdateTrafficControlTaskTraffic(TrafficControlTaskId *string, request *UpdateTrafficControlTaskTrafficRequest) (_result *UpdateTrafficControlTaskTrafficResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrafficControlTaskTrafficResponse{}
	_body, _err := client.UpdateTrafficControlTaskTrafficWithOptions(TrafficControlTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传数据
//
// @param request - UploadRecommendationDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRecommendationDataResponse
func (client *Client) UploadRecommendationDataWithOptions(request *UploadRecommendationDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadRecommendationDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadRecommendationData"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/recommendationdata/action/upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadRecommendationDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传数据
//
// @param request - UploadRecommendationDataRequest
//
// @return UploadRecommendationDataResponse
func (client *Client) UploadRecommendationData(request *UploadRecommendationDataRequest) (_result *UploadRecommendationDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadRecommendationDataResponse{}
	_body, _err := client.UploadRecommendationDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) chatConversationWithSSE_opYieldFunc(_yield chan *ChatConversationResponse, _yieldErr chan error, request *ChatConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatConversation"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/conversations/chat"),
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

func (client *Client) shoppingAssistantWithSSE_opYieldFunc(_yield chan *ShoppingAssistantResponse, _yieldErr chan error, request *ShoppingAssistantRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Contents) {
		body["Contents"] = request.Contents
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Environment) {
		body["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InputMessage) {
		body["InputMessage"] = request.InputMessage
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ShoppingAssistant"),
		Version:     dara.String("2022-12-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/conversations/shopping_assistant/chat"),
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
