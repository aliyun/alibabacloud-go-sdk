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
// 发送对话消息
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
// 发送对话消息
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
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
// 发送对话消息
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
// 停止对话
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

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
// 停止对话
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
// 创建应用服务实例
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
// 创建应用服务实例
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
// 创建自定义agent
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
	if !dara.IsNil(tmpReq.Tools) {
		request.ToolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tools, dara.String("Tools"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.EnableTools) {
		query["EnableTools"] = request.EnableTools
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
// 创建自定义agent
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
// 删除应用服务实例
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
// 删除应用服务实例
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
// 删除Custom Agent
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

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
// 删除Custom Agent
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
// 查询应用服务详情
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
// 查询应用服务详情
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
// 查询应用服务列表
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
// 查询应用服务列表
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
// 查询事件信息列表
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
// 查询事件信息列表
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
// 查看实例认证信息
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
// 查看实例认证信息
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
// 查看服务连接信息
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
// 查看服务连接信息
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
// 查询服务白名单
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
// 查询服务白名单
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
// 查看实例RAG配置
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
// 查看实例RAG配置
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
// 查看实例SSL配置
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
// 查看实例SSL配置
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
// 查看实例存储配置
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
// 查看实例存储配置
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
// 获取会话列表
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

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
// 获取会话列表
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
// 查询CustomAgent
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

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
// 查询CustomAgent
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
// 获取会话历史消息
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
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
// 获取会话历史消息
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
// 获取Custom Agent列表
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
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
// 获取Custom Agent列表
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
// 获取专属Agent可用工具
//
// @param request - ListCustomAgentToolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAgentToolsResponse
func (client *Client) ListCustomAgentToolsWithOptions(request *ListCustomAgentToolsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomAgentToolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
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
// 获取专属Agent可用工具
//
// @param request - ListCustomAgentToolsRequest
//
// @return ListCustomAgentToolsResponse
func (client *Client) ListCustomAgentTools(request *ListCustomAgentToolsRequest) (_result *ListCustomAgentToolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomAgentToolsResponse{}
	_body, _err := client.ListCustomAgentToolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Supabase Auth相关配置
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
// 修改Supabase Auth相关配置
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
// 修改实例RAG配置
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
// 修改实例RAG配置
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
// 修改服务白名单
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
// 修改服务白名单
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
// 修改实例RAG配置
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
// 修改实例RAG配置
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
// 修改实例SSL配置
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
// 修改实例SSL配置
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
// 修改实例存储配置
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
// 修改实例存储配置
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
// 消息终端用户反馈、点赞/点踩
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

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
// 消息终端用户反馈、点赞/点踩
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
// 重置实例密码
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
// 重置实例密码
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
// 重启实例
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
// 重启实例
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
// 启动实例
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
// 启动实例
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
// 暂停实例
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
// 暂停实例
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
// 更新Custom Agent
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
	if !dara.IsNil(tmpReq.Tools) {
		request.ToolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tools, dara.String("Tools"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	if !dara.IsNil(request.EnableTools) {
		query["EnableTools"] = request.EnableTools
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
// 更新Custom Agent
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
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
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
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
