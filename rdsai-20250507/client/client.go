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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("rdsai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Sends chat messages.
//
// @param tmpReq - ChatMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatMessagesResponse
func (client *Client) ChatMessagesWithSSE(tmpReq *ChatMessagesRequest, runtime *dara.RuntimeOptions, _yield chan *ChatMessagesResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatMessagesWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
	return
}

// Summary:
//
// Sends chat messages.
//
// @param tmpReq - ChatMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatMessagesResponse
func (client *Client) ChatMessagesWithOptions(tmpReq *ChatMessagesRequest, runtime *dara.RuntimeOptions) (_result *ChatMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChatMessagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Inputs) {
		request.InputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Inputs, dara.String("Inputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.EventMode) {
		query["EventMode"] = request.EventMode
	}

	if !dara.IsNil(request.InputsShrink) {
		query["Inputs"] = request.InputsShrink
	}

	if !dara.IsNil(request.ParentMessageId) {
		query["ParentMessageId"] = request.ParentMessageId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatMessages"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends chat messages.
//
// @param request - ChatMessagesRequest
//
// @return ChatMessagesResponse
func (client *Client) ChatMessages(request *ChatMessagesRequest) (_result *ChatMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatMessagesResponse{}
	_body, _err := client.ChatMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a conversation.
//
// @param request - ChatMessagesTaskStopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatMessagesTaskStopResponse
func (client *Client) ChatMessagesTaskStopWithOptions(request *ChatMessagesTaskStopRequest, runtime *dara.RuntimeOptions) (_result *ChatMessagesTaskStopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatMessagesTaskStop"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatMessagesTaskStopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a conversation.
//
// @param request - ChatMessagesTaskStopRequest
//
// @return ChatMessagesTaskStopResponse
func (client *Client) ChatMessagesTaskStop(request *ChatMessagesTaskStopRequest) (_result *ChatMessagesTaskStopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatMessagesTaskStopResponse{}
	_body, _err := client.ChatMessagesTaskStopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the related topics.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param tmpReq - CreateAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppInstanceResponse
func (client *Client) CreateAppInstanceWithOptions(tmpReq *CreateAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBInstanceConfig) {
		request.DBInstanceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBInstanceConfig, dara.String("DBInstanceConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceConfigShrink) {
		query["DBInstanceConfig"] = request.DBInstanceConfigShrink
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DashboardPassword) {
		query["DashboardPassword"] = request.DashboardPassword
	}

	if !dara.IsNil(request.DashboardUsername) {
		query["DashboardUsername"] = request.DashboardUsername
	}

	if !dara.IsNil(request.DatabasePassword) {
		query["DatabasePassword"] = request.DatabasePassword
	}

	if !dara.IsNil(request.InitializeWithExistingData) {
		query["InitializeWithExistingData"] = request.InitializeWithExistingData
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.PublicEndpointEnabled) {
		query["PublicEndpointEnabled"] = request.PublicEndpointEnabled
	}

	if !dara.IsNil(request.PublicNetworkAccessEnabled) {
		query["PublicNetworkAccessEnabled"] = request.PublicNetworkAccessEnabled
	}

	if !dara.IsNil(request.RAGEnabled) {
		query["RAGEnabled"] = request.RAGEnabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppInstance"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the related topics.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - CreateAppInstanceRequest
//
// @return CreateAppInstanceResponse
func (client *Client) CreateAppInstance(request *CreateAppInstanceRequest) (_result *CreateAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppInstanceResponse{}
	_body, _err := client.CreateAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dedicated agent.
//
// @param tmpReq - CreateCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomAgentResponse
func (client *Client) CreateCustomAgentWithOptions(tmpReq *CreateCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCustomAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SkillIds) {
		request.SkillIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SkillIds, dara.String("SkillIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tools) {
		request.ToolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tools, dara.String("Tools"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableTools) {
		query["EnableTools"] = request.EnableTools
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillIdsShrink) {
		query["SkillIds"] = request.SkillIdsShrink
	}

	if !dara.IsNil(request.SystemPrompt) {
		query["SystemPrompt"] = request.SystemPrompt
	}

	if !dara.IsNil(request.ToolsShrink) {
		query["Tools"] = request.ToolsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomAgent"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dedicated agent.
//
// @param request - CreateCustomAgentRequest
//
// @return CreateCustomAgentResponse
func (client *Client) CreateCustomAgent(request *CreateCustomAgentRequest) (_result *CreateCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomAgentResponse{}
	_body, _err := client.CreateCustomAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建边缘函数
//
// @param tmpReq - CreateEdgeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEdgeFunctionResponse
func (client *Client) CreateEdgeFunctionWithOptions(tmpReq *CreateEdgeFunctionRequest, runtime *dara.RuntimeOptions) (_result *CreateEdgeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEdgeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Code) {
		request.CodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Code, dara.String("Code"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomConfig) {
		request.CustomConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomConfig, dara.String("CustomConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Envs) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, dara.String("Envs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CodeShrink) {
		query["Code"] = request.CodeShrink
	}

	if !dara.IsNil(request.CustomConfigShrink) {
		query["CustomConfig"] = request.CustomConfigShrink
	}

	if !dara.IsNil(request.EdgeFunctionName) {
		query["EdgeFunctionName"] = request.EdgeFunctionName
	}

	if !dara.IsNil(request.EnvsShrink) {
		query["Envs"] = request.EnvsShrink
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEdgeFunction"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEdgeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建边缘函数
//
// @param request - CreateEdgeFunctionRequest
//
// @return CreateEdgeFunctionResponse
func (client *Client) CreateEdgeFunction(request *CreateEdgeFunctionRequest) (_result *CreateEdgeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEdgeFunctionResponse{}
	_body, _err := client.CreateEdgeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an inspection task for multiple instances.
//
// @param request - CreateInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInspectionTaskResponse
func (client *Client) CreateInspectionTaskWithOptions(request *CreateInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateInspectionTaskResponse, _err error) {
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

	if !dara.IsNil(request.InspectionItems) {
		query["InspectionItems"] = request.InspectionItems
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReportLanguage) {
		query["ReportLanguage"] = request.ReportLanguage
	}

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInspectionTask"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInspectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an inspection task for multiple instances.
//
// @param request - CreateInspectionTaskRequest
//
// @return CreateInspectionTaskResponse
func (client *Client) CreateInspectionTask(request *CreateInspectionTaskRequest) (_result *CreateInspectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInspectionTaskResponse{}
	_body, _err := client.CreateInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new scheduled inspection configuration for multiple instances.
//
// @param request - CreateScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledTaskResponse
func (client *Client) CreateScheduledTaskWithOptions(request *CreateScheduledTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
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

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReportLanguage) {
		query["ReportLanguage"] = request.ReportLanguage
	}

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledTask"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new scheduled inspection configuration for multiple instances.
//
// @param request - CreateScheduledTaskRequest
//
// @return CreateScheduledTaskResponse
func (client *Client) CreateScheduledTask(request *CreateScheduledTaskRequest) (_result *CreateScheduledTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CreateScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a user-defined skill.
//
// @param tmpReq - CreateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillResponse
func (client *Client) CreateSkillWithOptions(tmpReq *CreateSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Dbtypes) {
		request.DbtypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dbtypes, dara.String("Dbtypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		query["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.DbtypesShrink) {
		query["Dbtypes"] = request.DbtypesShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkill"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a user-defined skill.
//
// @param request - CreateSkillRequest
//
// @return CreateSkillResponse
func (client *Client) CreateSkill(request *CreateSkillRequest) (_result *CreateSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillResponse{}
	_body, _err := client.CreateSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the related topics.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  If you delete an RDS Supabase instance, the created RDS for PostgreSQL instance and the created NAT gateway are not automatically deleted. You must manually release the instance and delete the Internet NAT gateway and EIP.
//
// @param request - DeleteAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInstanceResponse
func (client *Client) DeleteAppInstanceWithOptions(request *DeleteAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppInstance"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the related topics.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  If you delete an RDS Supabase instance, the created RDS for PostgreSQL instance and the created NAT gateway are not automatically deleted. You must manually release the instance and delete the Internet NAT gateway and EIP.
//
// @param request - DeleteAppInstanceRequest
//
// @return DeleteAppInstanceResponse
func (client *Client) DeleteAppInstance(request *DeleteAppInstanceRequest) (_result *DeleteAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppInstanceResponse{}
	_body, _err := client.DeleteAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the dedicated agent created by a user.
//
// @param request - DeleteCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomAgentResponse
func (client *Client) DeleteCustomAgentWithOptions(request *DeleteCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomAgent"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the dedicated agent created by a user.
//
// @param request - DeleteCustomAgentRequest
//
// @return DeleteCustomAgentResponse
func (client *Client) DeleteCustomAgent(request *DeleteCustomAgentRequest) (_result *DeleteCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomAgentResponse{}
	_body, _err := client.DeleteCustomAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除边缘函数
//
// @param request - DeleteEdgeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEdgeFunctionResponse
func (client *Client) DeleteEdgeFunctionWithOptions(request *DeleteEdgeFunctionRequest, runtime *dara.RuntimeOptions) (_result *DeleteEdgeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EdgeFunctionName) {
		query["EdgeFunctionName"] = request.EdgeFunctionName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEdgeFunction"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEdgeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除边缘函数
//
// @param request - DeleteEdgeFunctionRequest
//
// @return DeleteEdgeFunctionResponse
func (client *Client) DeleteEdgeFunction(request *DeleteEdgeFunctionRequest) (_result *DeleteEdgeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEdgeFunctionResponse{}
	_body, _err := client.DeleteEdgeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified inspection configuration.
//
// @param request - DeleteScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledTaskResponse
func (client *Client) DeleteScheduledTaskWithOptions(request *DeleteScheduledTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteScheduledTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScheduledId) {
		query["ScheduledId"] = request.ScheduledId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduledTask"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified inspection configuration.
//
// @param request - DeleteScheduledTaskRequest
//
// @return DeleteScheduledTaskResponse
func (client *Client) DeleteScheduledTask(request *DeleteScheduledTaskRequest) (_result *DeleteScheduledTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.DeleteScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified skill.
//
// @param request - DeleteSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkillWithOptions(request *DeleteSkillRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkill"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Deletes the specified skill.
//
// @param request - DeleteSkillRequest
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkill(request *DeleteSkillRequest) (_result *DeleteSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSkillResponse{}
	_body, _err := client.DeleteSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeAppInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppInstanceAttributeResponse
func (client *Client) DescribeAppInstanceAttributeWithOptions(request *DescribeAppInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppInstanceAttribute"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeAppInstanceAttributeRequest
//
// @return DescribeAppInstanceAttributeResponse
func (client *Client) DescribeAppInstanceAttribute(request *DescribeAppInstanceAttributeRequest) (_result *DescribeAppInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppInstanceAttributeResponse{}
	_body, _err := client.DescribeAppInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the RDS Supabase instances.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeAppInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppInstancesResponse
func (client *Client) DescribeAppInstancesWithOptions(request *DescribeAppInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppInstances"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the RDS Supabase instances.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeAppInstancesRequest
//
// @return DescribeAppInstancesResponse
func (client *Client) DescribeAppInstances(request *DescribeAppInstancesRequest) (_result *DescribeAppInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppInstancesResponse{}
	_body, _err := client.DescribeAppInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询边缘函数列表
//
// @param request - DescribeEdgeFunctionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEdgeFunctionsResponse
func (client *Client) DescribeEdgeFunctionsWithOptions(request *DescribeEdgeFunctionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEdgeFunctionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EdgeFunctionName) {
		query["EdgeFunctionName"] = request.EdgeFunctionName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEdgeFunctions"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEdgeFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询边缘函数列表
//
// @param request - DescribeEdgeFunctionsRequest
//
// @return DescribeEdgeFunctionsResponse
func (client *Client) DescribeEdgeFunctions(request *DescribeEdgeFunctionsRequest) (_result *DescribeEdgeFunctionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEdgeFunctionsResponse{}
	_body, _err := client.DescribeEdgeFunctionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the events.
//
// @param request - DescribeEventsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsListResponse
func (client *Client) DescribeEventsListWithOptions(request *DescribeEventsListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsListResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventsList"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the events.
//
// @param request - DescribeEventsListRequest
//
// @return DescribeEventsListResponse
func (client *Client) DescribeEventsList(request *DescribeEventsListRequest) (_result *DescribeEventsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventsListResponse{}
	_body, _err := client.DescribeEventsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authentication information about an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceAuthInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAuthInfoResponse
func (client *Client) DescribeInstanceAuthInfoWithOptions(request *DescribeInstanceAuthInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAuthInfoResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAuthInfo"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAuthInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authentication information about an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceAuthInfoRequest
//
// @return DescribeInstanceAuthInfoResponse
func (client *Client) DescribeInstanceAuthInfo(request *DescribeInstanceAuthInfoRequest) (_result *DescribeInstanceAuthInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAuthInfoResponse{}
	_body, _err := client.DescribeInstanceAuthInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the endpoint of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceEndpointsResponse
func (client *Client) DescribeInstanceEndpointsWithOptions(request *DescribeInstanceEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceEndpoints"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoint of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceEndpointsRequest
//
// @return DescribeInstanceEndpointsResponse
func (client *Client) DescribeInstanceEndpoints(request *DescribeInstanceEndpointsRequest) (_result *DescribeInstanceEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceEndpointsResponse{}
	_body, _err := client.DescribeInstanceEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceIpWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceIpWhitelistResponse
func (client *Client) DescribeInstanceIpWhitelistWithOptions(request *DescribeInstanceIpWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceIpWhitelist"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceIpWhitelistRequest
//
// @return DescribeInstanceIpWhitelistResponse
func (client *Client) DescribeInstanceIpWhitelist(request *DescribeInstanceIpWhitelistRequest) (_result *DescribeInstanceIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceIpWhitelistResponse{}
	_body, _err := client.DescribeInstanceIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the RAG agent configurations of an RDS Supabase instance.
//
// @param request - DescribeInstanceRAGConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceRAGConfigResponse
func (client *Client) DescribeInstanceRAGConfigWithOptions(request *DescribeInstanceRAGConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceRAGConfigResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceRAGConfig"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceRAGConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the RAG agent configurations of an RDS Supabase instance.
//
// @param request - DescribeInstanceRAGConfigRequest
//
// @return DescribeInstanceRAGConfigResponse
func (client *Client) DescribeInstanceRAGConfig(request *DescribeInstanceRAGConfigRequest) (_result *DescribeInstanceRAGConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceRAGConfigResponse{}
	_body, _err := client.DescribeInstanceRAGConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SSL settings of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSSLResponse
func (client *Client) DescribeInstanceSSLWithOptions(request *DescribeInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSSLResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSSL"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SSL settings of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeInstanceSSLRequest
//
// @return DescribeInstanceSSLResponse
func (client *Client) DescribeInstanceSSL(request *DescribeInstanceSSLRequest) (_result *DescribeInstanceSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceSSLResponse{}
	_body, _err := client.DescribeInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage configurations of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  Only Object Storage Service (OSS) is supported for the storage of RDS Supabase.
//
// @param request - DescribeInstanceStorageConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceStorageConfigResponse
func (client *Client) DescribeInstanceStorageConfigWithOptions(request *DescribeInstanceStorageConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceStorageConfigResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceStorageConfig"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceStorageConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage configurations of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  Only Object Storage Service (OSS) is supported for the storage of RDS Supabase.
//
// @param request - DescribeInstanceStorageConfigRequest
//
// @return DescribeInstanceStorageConfigResponse
func (client *Client) DescribeInstanceStorageConfig(request *DescribeInstanceStorageConfigRequest) (_result *DescribeInstanceStorageConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceStorageConfigResponse{}
	_body, _err := client.DescribeInstanceStorageConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the history conversations of a user.
//
// @param request - GetConversationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationsResponse
func (client *Client) GetConversationsWithOptions(request *GetConversationsRequest, runtime *dara.RuntimeOptions) (_result *GetConversationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LastId) {
		query["LastId"] = request.LastId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Pinned) {
		query["Pinned"] = request.Pinned
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConversations"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConversationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the history conversations of a user.
//
// @param request - GetConversationsRequest
//
// @return GetConversationsResponse
func (client *Client) GetConversations(request *GetConversationsRequest) (_result *GetConversationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConversationsResponse{}
	_body, _err := client.GetConversationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dedicated agents created by a user.
//
// @param request - GetCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomAgentResponse
func (client *Client) GetCustomAgentWithOptions(request *GetCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *GetCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomAgent"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dedicated agents created by a user.
//
// @param request - GetCustomAgentRequest
//
// @return GetCustomAgentResponse
func (client *Client) GetCustomAgent(request *GetCustomAgentRequest) (_result *GetCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomAgentResponse{}
	_body, _err := client.GetCustomAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the content of a specified inspection report.
//
// @param request - GetInspectionReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInspectionReportResponse
func (client *Client) GetInspectionReportWithOptions(request *GetInspectionReportRequest, runtime *dara.RuntimeOptions) (_result *GetInspectionReportResponse, _err error) {
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

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInspectionReport"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the content of a specified inspection report.
//
// @param request - GetInspectionReportRequest
//
// @return GetInspectionReportResponse
func (client *Client) GetInspectionReport(request *GetInspectionReportRequest) (_result *GetInspectionReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInspectionReportResponse{}
	_body, _err := client.GetInspectionReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries specific conversation messages.
//
// @param request - GetMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessagesResponse
func (client *Client) GetMessagesWithOptions(request *GetMessagesRequest, runtime *dara.RuntimeOptions) (_result *GetMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.EventMode) {
		query["EventMode"] = request.EventMode
	}

	if !dara.IsNil(request.FirstId) {
		query["FirstId"] = request.FirstId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessages"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries specific conversation messages.
//
// @param request - GetMessagesRequest
//
// @return GetMessagesResponse
func (client *Client) GetMessages(request *GetMessagesRequest) (_result *GetMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMessagesResponse{}
	_body, _err := client.GetMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IDs of all instances that are included by a specified scheduled inspection configuration.
//
// @param request - GetScheduledInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledInstancesResponse
func (client *Client) GetScheduledInstancesWithOptions(request *GetScheduledInstancesRequest, runtime *dara.RuntimeOptions) (_result *GetScheduledInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScheduledId) {
		query["ScheduledId"] = request.ScheduledId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduledInstances"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs of all instances that are included by a specified scheduled inspection configuration.
//
// @param request - GetScheduledInstancesRequest
//
// @return GetScheduledInstancesResponse
func (client *Client) GetScheduledInstances(request *GetScheduledInstancesRequest) (_result *GetScheduledInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScheduledInstancesResponse{}
	_body, _err := client.GetScheduledInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of all inspection reports for a specified scheduled task. You can filter and paginate inspection reports by time range.
//
// @param request - GetScheduledReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledReportsResponse
func (client *Client) GetScheduledReportsWithOptions(request *GetScheduledReportsRequest, runtime *dara.RuntimeOptions) (_result *GetScheduledReportsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScheduledId) {
		query["ScheduledId"] = request.ScheduledId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduledReports"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of all inspection reports for a specified scheduled task. You can filter and paginate inspection reports by time range.
//
// @param request - GetScheduledReportsRequest
//
// @return GetScheduledReportsResponse
func (client *Client) GetScheduledReports(request *GetScheduledReportsRequest) (_result *GetScheduledReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScheduledReportsResponse{}
	_body, _err := client.GetScheduledReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a specified skill. You can obtain the details of user-defined skills or the system preset skills.
//
// @param request - GetSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillResponse
func (client *Client) GetSkillWithOptions(request *GetSkillRequest, runtime *dara.RuntimeOptions) (_result *GetSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkill"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a specified skill. You can obtain the details of user-defined skills or the system preset skills.
//
// @param request - GetSkillRequest
//
// @return GetSkillResponse
func (client *Client) GetSkill(request *GetSkillRequest) (_result *GetSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillResponse{}
	_body, _err := client.GetSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the individual inspection reports of all non-scheduled tasks under a specified user. Pagination is supported.
//
// @param request - GetStandAloneReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandAloneReportsResponse
func (client *Client) GetStandAloneReportsWithOptions(request *GetStandAloneReportsRequest, runtime *dara.RuntimeOptions) (_result *GetStandAloneReportsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandAloneReports"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandAloneReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the individual inspection reports of all non-scheduled tasks under a specified user. Pagination is supported.
//
// @param request - GetStandAloneReportsRequest
//
// @return GetStandAloneReportsResponse
func (client *Client) GetStandAloneReports(request *GetStandAloneReportsRequest) (_result *GetStandAloneReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandAloneReportsResponse{}
	_body, _err := client.GetStandAloneReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dedicated agents created by a user.
//
// @param request - ListCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAgentResponse
func (client *Client) ListCustomAgentWithOptions(request *ListCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *ListCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomAgent"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dedicated agents created by a user.
//
// @param request - ListCustomAgentRequest
//
// @return ListCustomAgentResponse
func (client *Client) ListCustomAgent(request *ListCustomAgentRequest) (_result *ListCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomAgentResponse{}
	_body, _err := client.ListCustomAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom agent tools of the user.
//
// @param request - ListCustomAgentToolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAgentToolsResponse
func (client *Client) ListCustomAgentToolsWithOptions(runtime *dara.RuntimeOptions) (_result *ListCustomAgentToolsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomAgentTools"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomAgentToolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom agent tools of the user.
//
// @return ListCustomAgentToolsResponse
func (client *Client) ListCustomAgentTools() (_result *ListCustomAgentToolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomAgentToolsResponse{}
	_body, _err := client.ListCustomAgentToolsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information of all inspection configurations under a specified user.
//
// @param request - ListScheduledTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledTasksResponse
func (client *Client) ListScheduledTasksWithOptions(request *ListScheduledTasksRequest, runtime *dara.RuntimeOptions) (_result *ListScheduledTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScheduledId) {
		query["ScheduledId"] = request.ScheduledId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduledTasks"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information of all inspection configurations under a specified user.
//
// @param request - ListScheduledTasksRequest
//
// @return ListScheduledTasksResponse
func (client *Client) ListScheduledTasks(request *ListScheduledTasksRequest) (_result *ListScheduledTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.ListScheduledTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the user-defined skills and all system preset skills of the current user.
//
// @param request - ListSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillResponse
func (client *Client) ListSkillWithOptions(request *ListSkillRequest, runtime *dara.RuntimeOptions) (_result *ListSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkill"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the user-defined skills and all system preset skills of the current user.
//
// @param request - ListSkillRequest
//
// @return ListSkillResponse
func (client *Client) ListSkill(request *ListSkillRequest) (_result *ListSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillResponse{}
	_body, _err := client.ListSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the authentication configurations of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param tmpReq - ModifyInstanceAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAuthConfigResponse
func (client *Client) ModifyInstanceAuthConfigWithOptions(tmpReq *ModifyInstanceAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstanceAuthConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigList) {
		request.ConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigList, dara.String("ConfigList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigListShrink) {
		query["ConfigList"] = request.ConfigListShrink
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAuthConfig"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAuthConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the authentication configurations of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifyInstanceAuthConfigRequest
//
// @return ModifyInstanceAuthConfigResponse
func (client *Client) ModifyInstanceAuthConfig(request *ModifyInstanceAuthConfigRequest) (_result *ModifyInstanceAuthConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceAuthConfigResponse{}
	_body, _err := client.ModifyInstanceAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the general configurations of an instance, such as the EIP and NAT settings.
//
// @param request - ModifyInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfigWithOptions(request *ModifyInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.ConfigValue) {
		query["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceConfig"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the general configurations of an instance, such as the EIP and NAT settings.
//
// @param request - ModifyInstanceConfigRequest
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfig(request *ModifyInstanceConfigRequest) (_result *ModifyInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.ModifyInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifyInstanceIpWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceIpWhitelistResponse
func (client *Client) ModifyInstanceIpWhitelistWithOptions(request *ModifyInstanceIpWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IpWhitelist) {
		query["IpWhitelist"] = request.IpWhitelist
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceIpWhitelist"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifyInstanceIpWhitelistRequest
//
// @return ModifyInstanceIpWhitelistResponse
func (client *Client) ModifyInstanceIpWhitelist(request *ModifyInstanceIpWhitelistRequest) (_result *ModifyInstanceIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceIpWhitelistResponse{}
	_body, _err := client.ModifyInstanceIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the RAG agent configurations of an RDS Supabase instance.
//
// @param tmpReq - ModifyInstanceRAGConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceRAGConfigResponse
func (client *Client) ModifyInstanceRAGConfigWithOptions(tmpReq *ModifyInstanceRAGConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceRAGConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstanceRAGConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigList) {
		request.ConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigList, dara.String("ConfigList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigListShrink) {
		query["ConfigList"] = request.ConfigListShrink
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceRAGConfig"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceRAGConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the RAG agent configurations of an RDS Supabase instance.
//
// @param request - ModifyInstanceRAGConfigRequest
//
// @return ModifyInstanceRAGConfigResponse
func (client *Client) ModifyInstanceRAGConfig(request *ModifyInstanceRAGConfigRequest) (_result *ModifyInstanceRAGConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceRAGConfigResponse{}
	_body, _err := client.ModifyInstanceRAGConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the SSL settings of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifyInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceSSLResponse
func (client *Client) ModifyInstanceSSLWithOptions(request *ModifyInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CAType) {
		query["CAType"] = request.CAType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SSLEnabled) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	if !dara.IsNil(request.ServerCert) {
		query["ServerCert"] = request.ServerCert
	}

	if !dara.IsNil(request.ServerKey) {
		query["ServerKey"] = request.ServerKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceSSL"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the SSL settings of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifyInstanceSSLRequest
//
// @return ModifyInstanceSSLResponse
func (client *Client) ModifyInstanceSSL(request *ModifyInstanceSSLRequest) (_result *ModifyInstanceSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceSSLResponse{}
	_body, _err := client.ModifyInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the storage configurations of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  Only Object Storage Service (OSS) is supported for the storage of RDS Supabase.
//
// @param tmpReq - ModifyInstanceStorageConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceStorageConfigResponse
func (client *Client) ModifyInstanceStorageConfigWithOptions(tmpReq *ModifyInstanceStorageConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceStorageConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstanceStorageConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigList) {
		request.ConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigList, dara.String("ConfigList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigListShrink) {
		query["ConfigList"] = request.ConfigListShrink
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceStorageConfig"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceStorageConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the storage configurations of an RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  Only Object Storage Service (OSS) is supported for the storage of RDS Supabase.
//
// @param request - ModifyInstanceStorageConfigRequest
//
// @return ModifyInstanceStorageConfigResponse
func (client *Client) ModifyInstanceStorageConfig(request *ModifyInstanceStorageConfigRequest) (_result *ModifyInstanceStorageConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceStorageConfigResponse{}
	_body, _err := client.ModifyInstanceStorageConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the SSL settings of RDS Supabase instances in batches.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param tmpReq - ModifyInstancesSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstancesSSLResponse
func (client *Client) ModifyInstancesSSLWithOptions(tmpReq *ModifyInstancesSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstancesSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstancesSSLShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceNames) {
		request.InstanceNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceNames, dara.String("InstanceNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CAType) {
		query["CAType"] = request.CAType
	}

	if !dara.IsNil(request.InstanceNamesShrink) {
		query["InstanceNames"] = request.InstanceNamesShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SSLEnabled) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	if !dara.IsNil(request.ServerCert) {
		query["ServerCert"] = request.ServerCert
	}

	if !dara.IsNil(request.ServerKey) {
		query["ServerKey"] = request.ServerKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstancesSSL"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstancesSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the SSL settings of RDS Supabase instances in batches.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifyInstancesSSLRequest
//
// @return ModifyInstancesSSLResponse
func (client *Client) ModifyInstancesSSL(request *ModifyInstancesSSLRequest) (_result *ModifyInstancesSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstancesSSLResponse{}
	_body, _err := client.ModifyInstancesSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the returned messages.
//
// @param request - ModifyMessagesFeedbacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMessagesFeedbacksResponse
func (client *Client) ModifyMessagesFeedbacksWithOptions(request *ModifyMessagesFeedbacksRequest, runtime *dara.RuntimeOptions) (_result *ModifyMessagesFeedbacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.Rating) {
		query["Rating"] = request.Rating
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMessagesFeedbacks"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMessagesFeedbacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the returned messages.
//
// @param request - ModifyMessagesFeedbacksRequest
//
// @return ModifyMessagesFeedbacksResponse
func (client *Client) ModifyMessagesFeedbacks(request *ModifyMessagesFeedbacksRequest) (_result *ModifyMessagesFeedbacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMessagesFeedbacksResponse{}
	_body, _err := client.ModifyMessagesFeedbacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an existing inspection configuration.
//
// @param request - ModifyScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduledTaskResponse
func (client *Client) ModifyScheduledTaskWithOptions(request *ModifyScheduledTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyScheduledTaskResponse, _err error) {
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

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ReportLanguage) {
		query["ReportLanguage"] = request.ReportLanguage
	}

	if !dara.IsNil(request.ScheduledId) {
		query["ScheduledId"] = request.ScheduledId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScheduledTask"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an existing inspection configuration.
//
// @param request - ModifyScheduledTaskRequest
//
// @return ModifyScheduledTaskResponse
func (client *Client) ModifyScheduledTask(request *ModifyScheduledTaskRequest) (_result *ModifyScheduledTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.ModifyScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the logon password of the RDS Supabase instance and the access password of the database.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  You can only change the password of a RDS Supabase Dashboard user.
//
// @param request - ResetInstancePasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetInstancePasswordResponse
func (client *Client) ResetInstancePasswordWithOptions(request *ResetInstancePasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetInstancePasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DashboardPassword) {
		query["DashboardPassword"] = request.DashboardPassword
	}

	if !dara.IsNil(request.DatabasePassword) {
		query["DatabasePassword"] = request.DatabasePassword
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetInstancePassword"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetInstancePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the logon password of the RDS Supabase instance and the access password of the database.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// >  You can only change the password of a RDS Supabase Dashboard user.
//
// @param request - ResetInstancePasswordRequest
//
// @return ResetInstancePasswordResponse
func (client *Client) ResetInstancePassword(request *ResetInstancePasswordRequest) (_result *ResetInstancePasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetInstancePasswordResponse{}
	_body, _err := client.ResetInstancePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an RDS Supabase instance that is in the Running state.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - RestartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstance"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an RDS Supabase instance that is in the Running state.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - RestartInstanceRequest
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a stopped RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - StartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstance"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a stopped RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - StartInstanceRequest
//
// @return StartInstanceResponse
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a running RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - StopInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstance"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a running RDS Supabase instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - StopInstanceRequest
//
// @return StopInstanceResponse
func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the custom agent.
//
// @param tmpReq - UpdateCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomAgentResponse
func (client *Client) UpdateCustomAgentWithOptions(tmpReq *UpdateCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCustomAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SkillIds) {
		request.SkillIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SkillIds, dara.String("SkillIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tools) {
		request.ToolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tools, dara.String("Tools"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	if !dara.IsNil(request.EnableTools) {
		query["EnableTools"] = request.EnableTools
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillIdsShrink) {
		query["SkillIds"] = request.SkillIdsShrink
	}

	if !dara.IsNil(request.SystemPrompt) {
		query["SystemPrompt"] = request.SystemPrompt
	}

	if !dara.IsNil(request.ToolsShrink) {
		query["Tools"] = request.ToolsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomAgent"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the custom agent.
//
// @param request - UpdateCustomAgentRequest
//
// @return UpdateCustomAgentResponse
func (client *Client) UpdateCustomAgent(request *UpdateCustomAgentRequest) (_result *UpdateCustomAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomAgentResponse{}
	_body, _err := client.UpdateCustomAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新边缘函数
//
// @param tmpReq - UpdateEdgeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEdgeFunctionResponse
func (client *Client) UpdateEdgeFunctionWithOptions(tmpReq *UpdateEdgeFunctionRequest, runtime *dara.RuntimeOptions) (_result *UpdateEdgeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateEdgeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Code) {
		request.CodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Code, dara.String("Code"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomConfig) {
		request.CustomConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomConfig, dara.String("CustomConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Envs) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, dara.String("Envs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CodeShrink) {
		query["Code"] = request.CodeShrink
	}

	if !dara.IsNil(request.CustomConfigShrink) {
		query["CustomConfig"] = request.CustomConfigShrink
	}

	if !dara.IsNil(request.EdgeFunctionName) {
		query["EdgeFunctionName"] = request.EdgeFunctionName
	}

	if !dara.IsNil(request.EnvsShrink) {
		query["Envs"] = request.EnvsShrink
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEdgeFunction"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEdgeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新边缘函数
//
// @param request - UpdateEdgeFunctionRequest
//
// @return UpdateEdgeFunctionResponse
func (client *Client) UpdateEdgeFunction(request *UpdateEdgeFunctionRequest) (_result *UpdateEdgeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEdgeFunctionResponse{}
	_body, _err := client.UpdateEdgeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a specified skill.
//
// @param tmpReq - UpdateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillResponse
func (client *Client) UpdateSkillWithOptions(tmpReq *UpdateSkillRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Dbtypes) {
		request.DbtypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dbtypes, dara.String("Dbtypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		query["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.DbtypesShrink) {
		query["Dbtypes"] = request.DbtypesShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkill"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a specified skill.
//
// @param request - UpdateSkillRequest
//
// @return UpdateSkillResponse
func (client *Client) UpdateSkill(request *UpdateSkillRequest) (_result *UpdateSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillResponse{}
	_body, _err := client.UpdateSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) chatMessagesWithSSE_opYieldFunc(_yield chan *ChatMessagesResponse, _yieldErr chan error, tmpReq *ChatMessagesRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &ChatMessagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Inputs) {
		request.InputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Inputs, dara.String("Inputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.EventMode) {
		query["EventMode"] = request.EventMode
	}

	if !dara.IsNil(request.InputsShrink) {
		query["Inputs"] = request.InputsShrink
	}

	if !dara.IsNil(request.ParentMessageId) {
		query["ParentMessageId"] = request.ParentMessageId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatMessages"),
		Version:     dara.String("2025-05-07"),
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
