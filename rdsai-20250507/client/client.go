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
		"cn-wulanchabu": dara.String("rdsai.aliyuncs.com"),
		"cn-shenzhen":   dara.String("rdsai.aliyuncs.com"),
		"cn-beijing":    dara.String("rdsai.aliyuncs.com"),
		"cn-shanghai":   dara.String("rdsai.aliyuncs.com"),
		"cn-guangzhou":  dara.String("rdsai.aliyuncs.com"),
		"cn-hangzhou":   dara.String("rdsai.aliyuncs.com"),
	}
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
// Adds an exclusion item to an import task.
//
// Description:
//
// Adds a job exclusion item.
//
// @param request - AddContextDatabaseSyncJobExclusionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddContextDatabaseSyncJobExclusionResponse
func (client *Client) AddContextDatabaseSyncJobExclusionWithOptions(request *AddContextDatabaseSyncJobExclusionRequest, runtime *dara.RuntimeOptions) (_result *AddContextDatabaseSyncJobExclusionResponse, _err error) {
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

	if !dara.IsNil(request.DocumentId) {
		query["DocumentId"] = request.DocumentId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddContextDatabaseSyncJobExclusion"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddContextDatabaseSyncJobExclusionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an exclusion item to an import task.
//
// Description:
//
// Adds a job exclusion item.
//
// @param request - AddContextDatabaseSyncJobExclusionRequest
//
// @return AddContextDatabaseSyncJobExclusionResponse
func (client *Client) AddContextDatabaseSyncJobExclusion(request *AddContextDatabaseSyncJobExclusionRequest) (_result *AddContextDatabaseSyncJobExclusionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddContextDatabaseSyncJobExclusionResponse{}
	_body, _err := client.AddContextDatabaseSyncJobExclusionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an in-progress discovery task by task ID.
//
// Description:
//
// You can specify a ClientToken to ensure idempotence for write operations. If you do not specify this parameter, the backend automatically generates one.
//
// @param request - CancelContextDatabaseSyncDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelContextDatabaseSyncDiscoveryResponse
func (client *Client) CancelContextDatabaseSyncDiscoveryWithOptions(request *CancelContextDatabaseSyncDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *CancelContextDatabaseSyncDiscoveryResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelContextDatabaseSyncDiscovery"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelContextDatabaseSyncDiscoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an in-progress discovery task by task ID.
//
// Description:
//
// You can specify a ClientToken to ensure idempotence for write operations. If you do not specify this parameter, the backend automatically generates one.
//
// @param request - CancelContextDatabaseSyncDiscoveryRequest
//
// @return CancelContextDatabaseSyncDiscoveryResponse
func (client *Client) CancelContextDatabaseSyncDiscovery(request *CancelContextDatabaseSyncDiscoveryRequest) (_result *CancelContextDatabaseSyncDiscoveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelContextDatabaseSyncDiscoveryResponse{}
	_body, _err := client.CancelContextDatabaseSyncDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a run.
//
// Description:
//
// Write operations support idempotence through ClientToken. If ClientToken is not provided, the backend generates one automatically. Cancellation is best-effort. Completed RunItems are not affected.
//
// @param request - CancelContextDatabaseSyncRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelContextDatabaseSyncRunResponse
func (client *Client) CancelContextDatabaseSyncRunWithOptions(request *CancelContextDatabaseSyncRunRequest, runtime *dara.RuntimeOptions) (_result *CancelContextDatabaseSyncRunResponse, _err error) {
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

	if !dara.IsNil(request.RunId) {
		query["RunId"] = request.RunId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelContextDatabaseSyncRun"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelContextDatabaseSyncRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a run.
//
// Description:
//
// Write operations support idempotence through ClientToken. If ClientToken is not provided, the backend generates one automatically. Cancellation is best-effort. Completed RunItems are not affected.
//
// @param request - CancelContextDatabaseSyncRunRequest
//
// @return CancelContextDatabaseSyncRunResponse
func (client *Client) CancelContextDatabaseSyncRun(request *CancelContextDatabaseSyncRunRequest) (_result *CancelContextDatabaseSyncRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelContextDatabaseSyncRunResponse{}
	_body, _err := client.CancelContextDatabaseSyncRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a conversation message.
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
// Sends a conversation message.
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
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

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

	if !dara.IsNil(request.FilesShrink) {
		query["Files"] = request.FilesShrink
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
// Sends a conversation message.
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

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Creates a custom API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - CreateApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKeyWithOptions(request *CreateApiKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DailyTokenQuota) {
		query["DailyTokenQuota"] = request.DailyTokenQuota
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.LimitRate) {
		query["LimitRate"] = request.LimitRate
	}

	if !dara.IsNil(request.LimitType) {
		query["LimitType"] = request.LimitType
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.TokenQuota) {
		query["TokenQuota"] = request.TokenQuota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiKey"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - CreateApiKeyRequest
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKey(request *CreateApiKeyRequest) (_result *CreateApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CreateApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
//	Warning: This API operation incurs fees. Read the related documentation carefully before you perform this operation.
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
	if !dara.IsNil(tmpReq.Components) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, dara.String("Components"), dara.String("json"))
	}

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

	if !dara.IsNil(request.ComponentsShrink) {
		query["Components"] = request.ComponentsShrink
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

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
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
// Creates an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
//	Warning: This API operation incurs fees. Read the related documentation carefully before you perform this operation.
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
// Creates a context database API key.
//
// Description:
//
// Creates an API key and returns the plaintext apiKey.
//
// @param request - CreateContextDatabaseApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextDatabaseApiKeyResponse
func (client *Client) CreateContextDatabaseApiKeyWithOptions(request *CreateContextDatabaseApiKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateContextDatabaseApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextDatabaseApiKey"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextDatabaseApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a context database API key.
//
// Description:
//
// Creates an API key and returns the plaintext apiKey.
//
// @param request - CreateContextDatabaseApiKeyRequest
//
// @return CreateContextDatabaseApiKeyResponse
func (client *Client) CreateContextDatabaseApiKey(request *CreateContextDatabaseApiKeyRequest) (_result *CreateContextDatabaseApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContextDatabaseApiKeyResponse{}
	_body, _err := client.CreateContextDatabaseApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ContextDB member.
//
// Description:
//
// Creates a member. If GenerateInitialKey is set to true, the first API key is also issued, and the plaintext ApiKey is returned in the response. This is a sensitive field and is returned only once. Store it securely. After the member is created, you can call the List or Get operation to query the member and the metadata of the API keys under the member.
//
// @param request - CreateContextDatabaseMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextDatabaseMemberResponse
func (client *Client) CreateContextDatabaseMemberWithOptions(request *CreateContextDatabaseMemberRequest, runtime *dara.RuntimeOptions) (_result *CreateContextDatabaseMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GenerateInitialKey) {
		query["GenerateInitialKey"] = request.GenerateInitialKey
	}

	if !dara.IsNil(request.InitialKeyName) {
		query["InitialKeyName"] = request.InitialKeyName
	}

	if !dara.IsNil(request.MemberName) {
		query["MemberName"] = request.MemberName
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextDatabaseMember"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextDatabaseMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ContextDB member.
//
// Description:
//
// Creates a member. If GenerateInitialKey is set to true, the first API key is also issued, and the plaintext ApiKey is returned in the response. This is a sensitive field and is returned only once. Store it securely. After the member is created, you can call the List or Get operation to query the member and the metadata of the API keys under the member.
//
// @param request - CreateContextDatabaseMemberRequest
//
// @return CreateContextDatabaseMemberResponse
func (client *Client) CreateContextDatabaseMember(request *CreateContextDatabaseMemberRequest) (_result *CreateContextDatabaseMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContextDatabaseMemberResponse{}
	_body, _err := client.CreateContextDatabaseMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a synchronization connection.
//
// Description:
//
// For write operations, you can specify a ClientToken to ensure idempotency. If no ClientToken is provided, the backend automatically generates one. Upon successful creation, the complete information of the created object is returned, including the ID that can be used for subsequent queries or updates.
//
// @param request - CreateContextDatabaseSyncConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextDatabaseSyncConnectionResponse
func (client *Client) CreateContextDatabaseSyncConnectionWithOptions(request *CreateContextDatabaseSyncConnectionRequest, runtime *dara.RuntimeOptions) (_result *CreateContextDatabaseSyncConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthKind) {
		query["AuthKind"] = request.AuthKind
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextDatabaseSyncConnection"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextDatabaseSyncConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a synchronization connection.
//
// Description:
//
// For write operations, you can specify a ClientToken to ensure idempotency. If no ClientToken is provided, the backend automatically generates one. Upon successful creation, the complete information of the created object is returned, including the ID that can be used for subsequent queries or updates.
//
// @param request - CreateContextDatabaseSyncConnectionRequest
//
// @return CreateContextDatabaseSyncConnectionResponse
func (client *Client) CreateContextDatabaseSyncConnection(request *CreateContextDatabaseSyncConnectionRequest) (_result *CreateContextDatabaseSyncConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContextDatabaseSyncConnectionResponse{}
	_body, _err := client.CreateContextDatabaseSyncConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an import task.
//
// Description:
//
// Write operations support idempotence via ClientToken. If ClientToken is not provided, the backend generates one automatically. Target is an implicit concept — you only need to provide KnowledgeBaseId, and the backend automatically creates or matches the Target.
//
// On success, the response returns the complete information of the created object. The Id field is active and can be used for subsequent queries or updates.
//
// @param tmpReq - CreateContextDatabaseSyncJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextDatabaseSyncJobResponse
func (client *Client) CreateContextDatabaseSyncJobWithOptions(tmpReq *CreateContextDatabaseSyncJobRequest, runtime *dara.RuntimeOptions) (_result *CreateContextDatabaseSyncJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateContextDatabaseSyncJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filters) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, dara.String("Filters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FiltersShrink) {
		query["Filters"] = request.FiltersShrink
	}

	if !dara.IsNil(request.KnowledgeBaseId) {
		query["KnowledgeBaseId"] = request.KnowledgeBaseId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetKnowledgeBaseId) {
		query["TargetKnowledgeBaseId"] = request.TargetKnowledgeBaseId
	}

	if !dara.IsNil(request.TargetKnowledgeBaseName) {
		query["TargetKnowledgeBaseName"] = request.TargetKnowledgeBaseName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextDatabaseSyncJob"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextDatabaseSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an import task.
//
// Description:
//
// Write operations support idempotence via ClientToken. If ClientToken is not provided, the backend generates one automatically. Target is an implicit concept — you only need to provide KnowledgeBaseId, and the backend automatically creates or matches the Target.
//
// On success, the response returns the complete information of the created object. The Id field is active and can be used for subsequent queries or updates.
//
// @param request - CreateContextDatabaseSyncJobRequest
//
// @return CreateContextDatabaseSyncJobResponse
func (client *Client) CreateContextDatabaseSyncJob(request *CreateContextDatabaseSyncJobRequest) (_result *CreateContextDatabaseSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContextDatabaseSyncJobResponse{}
	_body, _err := client.CreateContextDatabaseSyncJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data selection scope.
//
// Description:
//
// Write operations support idempotence by passing a ClientToken. If no token is provided, the backend generates one automatically. After creating a data selection scope, you typically need to call StartContextDatabaseSyncSourceDiscovery to discover available nodes.
//
// On success, the response returns the complete information of the created object. The Id field can be used for subsequent queries or updates.
//
// @param tmpReq - CreateContextDatabaseSyncSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextDatabaseSyncSourceResponse
func (client *Client) CreateContextDatabaseSyncSourceWithOptions(tmpReq *CreateContextDatabaseSyncSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateContextDatabaseSyncSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateContextDatabaseSyncSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceConfig) {
		request.SourceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceConfig, dara.String("SourceConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SourceConfigShrink) {
		query["SourceConfig"] = request.SourceConfigShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextDatabaseSyncSource"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextDatabaseSyncSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data selection scope.
//
// Description:
//
// Write operations support idempotence by passing a ClientToken. If no token is provided, the backend generates one automatically. After creating a data selection scope, you typically need to call StartContextDatabaseSyncSourceDiscovery to discover available nodes.
//
// On success, the response returns the complete information of the created object. The Id field can be used for subsequent queries or updates.
//
// @param request - CreateContextDatabaseSyncSourceRequest
//
// @return CreateContextDatabaseSyncSourceResponse
func (client *Client) CreateContextDatabaseSyncSource(request *CreateContextDatabaseSyncSourceRequest) (_result *CreateContextDatabaseSyncSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContextDatabaseSyncSourceResponse{}
	_body, _err := client.CreateContextDatabaseSyncSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workspace, the first member, and the first API key in a one-time onboarding process.
//
// Description:
//
// Creates a workspace along with the first member and the first API key in a one-time onboarding process. The plaintext API key is returned in the response.
//
// @param request - CreateContextDatabaseWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextDatabaseWorkspaceResponse
func (client *Client) CreateContextDatabaseWorkspaceWithOptions(request *CreateContextDatabaseWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *CreateContextDatabaseWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberName) {
		query["MemberName"] = request.MemberName
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextDatabaseWorkspace"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextDatabaseWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace, the first member, and the first API key in a one-time onboarding process.
//
// Description:
//
// Creates a workspace along with the first member and the first API key in a one-time onboarding process. The plaintext API key is returned in the response.
//
// @param request - CreateContextDatabaseWorkspaceRequest
//
// @return CreateContextDatabaseWorkspaceResponse
func (client *Client) CreateContextDatabaseWorkspace(request *CreateContextDatabaseWorkspaceRequest) (_result *CreateContextDatabaseWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContextDatabaseWorkspaceResponse{}
	_body, _err := client.CreateContextDatabaseWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user-specific agent.
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
// Creates a user-specific agent.
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
// Creates a batch instance inspection task.
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

	if !dara.IsNil(request.ReportRegionId) {
		query["ReportRegionId"] = request.ReportRegionId
	}

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
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
// Creates a batch instance inspection task.
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
// Triggers an asynchronous export task for MO usage details in CSV format.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - CreateMOUsageDetailExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMOUsageDetailExportResponse
func (client *Client) CreateMOUsageDetailExportWithOptions(request *CreateMOUsageDetailExportRequest, runtime *dara.RuntimeOptions) (_result *CreateMOUsageDetailExportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UsageType) {
		query["UsageType"] = request.UsageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMOUsageDetailExport"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMOUsageDetailExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers an asynchronous export task for MO usage details in CSV format.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - CreateMOUsageDetailExportRequest
//
// @return CreateMOUsageDetailExportResponse
func (client *Client) CreateMOUsageDetailExport(request *CreateMOUsageDetailExportRequest) (_result *CreateMOUsageDetailExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMOUsageDetailExportResponse{}
	_body, _err := client.CreateMOUsageDetailExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a sandbox template.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param tmpReq - CreateSandboxTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSandboxTemplateResponse
func (client *Client) CreateSandboxTemplateWithOptions(tmpReq *CreateSandboxTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateSandboxTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSandboxTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultCpu) {
		query["DefaultCpu"] = request.DefaultCpu
	}

	if !dara.IsNil(request.DefaultMemory) {
		query["DefaultMemory"] = request.DefaultMemory
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Image) {
		query["Image"] = request.Image
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSandboxTemplate"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSandboxTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sandbox template.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - CreateSandboxTemplateRequest
//
// @return CreateSandboxTemplateResponse
func (client *Client) CreateSandboxTemplate(request *CreateSandboxTemplateRequest) (_result *CreateSandboxTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSandboxTemplateResponse{}
	_body, _err := client.CreateSandboxTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduled inspection configuration for batch instances.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
//
//	Warning: This API operation incurs fees. Read the related feature documentation carefully before you perform this operation.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
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

	if !dara.IsNil(request.InspectionItems) {
		query["InspectionItems"] = request.InspectionItems
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

	if !dara.IsNil(request.ReportRegionId) {
		query["ReportRegionId"] = request.ReportRegionId
	}

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
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
// Creates a scheduled inspection configuration for batch instances.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
//
//	Warning: This API operation incurs fees. Read the related feature documentation carefully before you perform this operation.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
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
// Creates a user-defined skill.
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

	if !dara.IsNil(request.UpdateNote) {
		query["UpdateNote"] = request.UpdateNote
	}

	if !dara.IsNil(request.UploadId) {
		query["UploadId"] = request.UploadId
	}

	if !dara.IsNil(request.UploadToken) {
		query["UploadToken"] = request.UploadToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Creates a user-defined skill.
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
// Deletes a custom API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DeleteApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKeyWithOptions(request *DeleteApiKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiKey"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DeleteApiKeyRequest
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKey(request *DeleteApiKeyRequest) (_result *DeleteApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.DeleteApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an RDS AI application instance.
//
// Description:
//
// ### Applicable DPI engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
//
//	Warning: This API operation incurs fees. Read the related feature documentation carefully before you perform this operation.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
//	Notice: Deleting an RDS Supabase project does not automatically delete the ApsaraDB RDS for PostgreSQL instance or the Internet NAT gateway that were created along with the project. You must [manual release the instance](https://help.aliyun.com/document_detail/96749.html) and delete the [Internet NAT gateway](https://help.aliyun.com/document_detail/121139.html) and [EIP](https://help.aliyun.com/document_detail/121527.html).
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

	if !dara.IsNil(request.DeleteDBInstance) {
		query["DeleteDBInstance"] = request.DeleteDBInstance
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
// Deletes an RDS AI application instance.
//
// Description:
//
// ### Applicable DPI engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
//
//	Warning: This API operation incurs fees. Read the related feature documentation carefully before you perform this operation.
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
//	Notice: Deleting an RDS Supabase project does not automatically delete the ApsaraDB RDS for PostgreSQL instance or the Internet NAT gateway that were created along with the project. You must [manual release the instance](https://help.aliyun.com/document_detail/96749.html) and delete the [Internet NAT gateway](https://help.aliyun.com/document_detail/121139.html) and [EIP](https://help.aliyun.com/document_detail/121527.html).
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
// Deletes a context database member.
//
// Description:
//
// Deletes a member. This is a hard delete and cannot be undone.
//
// @param request - DeleteContextDatabaseMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContextDatabaseMemberResponse
func (client *Client) DeleteContextDatabaseMemberWithOptions(request *DeleteContextDatabaseMemberRequest, runtime *dara.RuntimeOptions) (_result *DeleteContextDatabaseMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContextDatabaseMember"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContextDatabaseMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a context database member.
//
// Description:
//
// Deletes a member. This is a hard delete and cannot be undone.
//
// @param request - DeleteContextDatabaseMemberRequest
//
// @return DeleteContextDatabaseMemberResponse
func (client *Client) DeleteContextDatabaseMember(request *DeleteContextDatabaseMemberRequest) (_result *DeleteContextDatabaseMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContextDatabaseMemberResponse{}
	_body, _err := client.DeleteContextDatabaseMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data source.
//
// Description:
//
// For write operations, you can specify a ClientToken to ensure idempotency. If you do not specify a ClientToken, the backend automatically generates one.
//
// @param request - DeleteContextDatabaseSyncConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContextDatabaseSyncConnectionResponse
func (client *Client) DeleteContextDatabaseSyncConnectionWithOptions(request *DeleteContextDatabaseSyncConnectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteContextDatabaseSyncConnectionResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContextDatabaseSyncConnection"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContextDatabaseSyncConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data source.
//
// Description:
//
// For write operations, you can specify a ClientToken to ensure idempotency. If you do not specify a ClientToken, the backend automatically generates one.
//
// @param request - DeleteContextDatabaseSyncConnectionRequest
//
// @return DeleteContextDatabaseSyncConnectionResponse
func (client *Client) DeleteContextDatabaseSyncConnection(request *DeleteContextDatabaseSyncConnectionRequest) (_result *DeleteContextDatabaseSyncConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContextDatabaseSyncConnectionResponse{}
	_body, _err := client.DeleteContextDatabaseSyncConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an import task.
//
// Description:
//
// You can provide a ClientToken to ensure idempotency for write operations. If not provided, the backend automatically generates one.
//
// @param request - DeleteContextDatabaseSyncJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContextDatabaseSyncJobResponse
func (client *Client) DeleteContextDatabaseSyncJobWithOptions(request *DeleteContextDatabaseSyncJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteContextDatabaseSyncJobResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContextDatabaseSyncJob"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContextDatabaseSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an import task.
//
// Description:
//
// You can provide a ClientToken to ensure idempotency for write operations. If not provided, the backend automatically generates one.
//
// @param request - DeleteContextDatabaseSyncJobRequest
//
// @return DeleteContextDatabaseSyncJobResponse
func (client *Client) DeleteContextDatabaseSyncJob(request *DeleteContextDatabaseSyncJobRequest) (_result *DeleteContextDatabaseSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContextDatabaseSyncJobResponse{}
	_body, _err := client.DeleteContextDatabaseSyncJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data selection scope.
//
// Description:
//
// Write operations support idempotence through ClientToken. If ClientToken is not provided, the backend generates one automatically.
//
// @param request - DeleteContextDatabaseSyncSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContextDatabaseSyncSourceResponse
func (client *Client) DeleteContextDatabaseSyncSourceWithOptions(request *DeleteContextDatabaseSyncSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteContextDatabaseSyncSourceResponse, _err error) {
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

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContextDatabaseSyncSource"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContextDatabaseSyncSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data selection scope.
//
// Description:
//
// Write operations support idempotence through ClientToken. If ClientToken is not provided, the backend generates one automatically.
//
// @param request - DeleteContextDatabaseSyncSourceRequest
//
// @return DeleteContextDatabaseSyncSourceResponse
func (client *Client) DeleteContextDatabaseSyncSource(request *DeleteContextDatabaseSyncSourceRequest) (_result *DeleteContextDatabaseSyncSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContextDatabaseSyncSourceResponse{}
	_body, _err := client.DeleteContextDatabaseSyncSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a ContextDB workspace.
//
// Description:
//
// Deletes a workspace. This is a hard delete and cannot be recovered. After the workspace is deleted, the local metadata is soft-deleted synchronously. Deleted workspaces no longer count toward the quota.
//
// @param request - DeleteContextDatabaseWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContextDatabaseWorkspaceResponse
func (client *Client) DeleteContextDatabaseWorkspaceWithOptions(request *DeleteContextDatabaseWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteContextDatabaseWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContextDatabaseWorkspace"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContextDatabaseWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a ContextDB workspace.
//
// Description:
//
// Deletes a workspace. This is a hard delete and cannot be recovered. After the workspace is deleted, the local metadata is soft-deleted synchronously. Deleted workspaces no longer count toward the quota.
//
// @param request - DeleteContextDatabaseWorkspaceRequest
//
// @return DeleteContextDatabaseWorkspaceResponse
func (client *Client) DeleteContextDatabaseWorkspace(request *DeleteContextDatabaseWorkspaceRequest) (_result *DeleteContextDatabaseWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContextDatabaseWorkspaceResponse{}
	_body, _err := client.DeleteContextDatabaseWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dedicated agent created by the user.
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
// Deletes a dedicated agent created by the user.
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
// Deletes a sandbox template.
//
// Description:
//
// ### Supported engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DeleteSandboxTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSandboxTemplateResponse
func (client *Client) DeleteSandboxTemplateWithOptions(request *DeleteSandboxTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSandboxTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSandboxTemplate"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSandboxTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a sandbox template.
//
// Description:
//
// ### Supported engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DeleteSandboxTemplateRequest
//
// @return DeleteSandboxTemplateResponse
func (client *Client) DeleteSandboxTemplate(request *DeleteSandboxTemplateRequest) (_result *DeleteSandboxTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSandboxTemplateResponse{}
	_body, _err := client.DeleteSandboxTemplateWithOptions(request, runtime)
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
// Deletes a specified Skill.
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

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Deletes a specified Skill.
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
// Queries the details of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Queries the details of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Queries a list of RDS AI application instances.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Queries a list of RDS AI application instances.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Queries the list of supported sandbox templates.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeCommonSandboxTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommonSandboxTemplatesResponse
func (client *Client) DescribeCommonSandboxTemplatesWithOptions(request *DescribeCommonSandboxTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommonSandboxTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCommonSandboxTemplates"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCommonSandboxTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of supported sandbox templates.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeCommonSandboxTemplatesRequest
//
// @return DescribeCommonSandboxTemplatesResponse
func (client *Client) DescribeCommonSandboxTemplates(request *DescribeCommonSandboxTemplatesRequest) (_result *DescribeCommonSandboxTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCommonSandboxTemplatesResponse{}
	_body, _err := client.DescribeCommonSandboxTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of events.
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
// Queries a list of events.
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
// Queries the authentication information of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Queries the authentication information of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Queries the endpoint information of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Queries the endpoint information of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Queries the IP whitelist of an RDS AI application instance.
//
// Description:
//
// ### Supported engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
// Queries the IP whitelist of an RDS AI application instance.
//
// Description:
//
// ### Supported engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Queries the RAG Agent configuration of an RDS AI application instance.
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Queries the RAG Agent configuration of an RDS AI application instance.
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
// Queries the SSL configuration of an RDS AI application instance.
//
// Description:
//
// ### Supported engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Queries the SSL configuration of an RDS AI application instance.
//
// Description:
//
// ### Supported engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Queries the storage configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// > Currently, only Object Storage Service (OSS) is supported.
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Queries the storage configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// > Currently, only Object Storage Service (OSS) is supported.
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
// Queries the token usage records of RDS AI Assistant Ultimate Edition.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeMOTokenUsageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMOTokenUsageDetailResponse
func (client *Client) DescribeMOTokenUsageDetailWithOptions(request *DescribeMOTokenUsageDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeMOTokenUsageDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ConsumerName) {
		query["ConsumerName"] = request.ConsumerName
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UsageType) {
		query["UsageType"] = request.UsageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMOTokenUsageDetail"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMOTokenUsageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the token usage records of RDS AI Assistant Ultimate Edition.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeMOTokenUsageDetailRequest
//
// @return DescribeMOTokenUsageDetailResponse
func (client *Client) DescribeMOTokenUsageDetail(request *DescribeMOTokenUsageDetailRequest) (_result *DescribeMOTokenUsageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMOTokenUsageDetailResponse{}
	_body, _err := client.DescribeMOTokenUsageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the token usage summary for a model operator instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeMOTokenUsageSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMOTokenUsageSummaryResponse
func (client *Client) DescribeMOTokenUsageSummaryWithOptions(request *DescribeMOTokenUsageSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeMOTokenUsageSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UsageType) {
		query["UsageType"] = request.UsageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMOTokenUsageSummary"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMOTokenUsageSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the token usage summary for a model operator instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeMOTokenUsageSummaryRequest
//
// @return DescribeMOTokenUsageSummaryResponse
func (client *Client) DescribeMOTokenUsageSummary(request *DescribeMOTokenUsageSummaryRequest) (_result *DescribeMOTokenUsageSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMOTokenUsageSummaryResponse{}
	_body, _err := client.DescribeMOTokenUsageSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status or download link of an asynchronous CSV export task for MO usage details.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeMOUsageDetailExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMOUsageDetailExportResponse
func (client *Client) DescribeMOUsageDetailExportWithOptions(request *DescribeMOUsageDetailExportRequest, runtime *dara.RuntimeOptions) (_result *DescribeMOUsageDetailExportResponse, _err error) {
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
		Action:      dara.String("DescribeMOUsageDetailExport"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMOUsageDetailExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status or download link of an asynchronous CSV export task for MO usage details.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeMOUsageDetailExportRequest
//
// @return DescribeMOUsageDetailExportResponse
func (client *Client) DescribeMOUsageDetailExport(request *DescribeMOUsageDetailExportRequest) (_result *DescribeMOUsageDetailExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMOUsageDetailExportResponse{}
	_body, _err := client.DescribeMOUsageDetailExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information and usage of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeModelOperatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelOperatorResponse
func (client *Client) DescribeModelOperatorWithOptions(request *DescribeModelOperatorRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelOperatorResponse, _err error) {
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

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelOperator"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelOperatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information and usage of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeModelOperatorRequest
//
// @return DescribeModelOperatorResponse
func (client *Client) DescribeModelOperator(request *DescribeModelOperatorRequest) (_result *DescribeModelOperatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelOperatorResponse{}
	_body, _err := client.DescribeModelOperatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instance monitoring data of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param tmpReq - DescribeMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorDataResponse
func (client *Client) DescribeMonitorDataWithOptions(tmpReq *DescribeMonitorDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeMonitorDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiKeyName) {
		request.ApiKeyNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiKeyName, dara.String("ApiKeyName"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyNameShrink) {
		query["ApiKeyName"] = request.ApiKeyNameShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorData"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instance monitoring data of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeMonitorDataRequest
//
// @return DescribeMonitorDataResponse
func (client *Client) DescribeMonitorData(request *DescribeMonitorDataRequest) (_result *DescribeMonitorDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorDataResponse{}
	_body, _err := client.DescribeMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of sandbox templates used to create Supabase sandboxes.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeSandboxTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSandboxTemplatesResponse
func (client *Client) DescribeSandboxTemplatesWithOptions(request *DescribeSandboxTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSandboxTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSandboxTemplates"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSandboxTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of sandbox templates used to create Supabase sandboxes.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - DescribeSandboxTemplatesRequest
//
// @return DescribeSandboxTemplatesResponse
func (client *Client) DescribeSandboxTemplates(request *DescribeSandboxTemplatesRequest) (_result *DescribeSandboxTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSandboxTemplatesResponse{}
	_body, _err := client.DescribeSandboxTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelist of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeWhitelistIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWhitelistIpsResponse
func (client *Client) DescribeWhitelistIpsWithOptions(request *DescribeWhitelistIpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWhitelistIpsResponse, _err error) {
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
		Action:      dara.String("DescribeWhitelistIps"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWhitelistIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelist of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - DescribeWhitelistIpsRequest
//
// @return DescribeWhitelistIpsResponse
func (client *Client) DescribeWhitelistIps(request *DescribeWhitelistIpsRequest) (_result *DescribeWhitelistIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWhitelistIpsResponse{}
	_body, _err := client.DescribeWhitelistIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the sandbox and Edge Routine capabilities of a Supabase instance. Note: This operation deletes all sandboxes and Edge Routines of the instance. Fully assess the business risks before performing this operation.
//
// Description:
//
// Disables the sandbox and Edge Routine capabilities of a Supabase instance.
//
//	Notice: This operation deletes all sandboxes and Edge Routines of the Supabase instance. Fully assess the business risks before performing this operation.
//
// @param request - DisableAgentRuntimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAgentRuntimeResponse
func (client *Client) DisableAgentRuntimeWithOptions(request *DisableAgentRuntimeRequest, runtime *dara.RuntimeOptions) (_result *DisableAgentRuntimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
		Action:      dara.String("DisableAgentRuntime"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAgentRuntimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the sandbox and Edge Routine capabilities of a Supabase instance. Note: This operation deletes all sandboxes and Edge Routines of the instance. Fully assess the business risks before performing this operation.
//
// Description:
//
// Disables the sandbox and Edge Routine capabilities of a Supabase instance.
//
//	Notice: This operation deletes all sandboxes and Edge Routines of the Supabase instance. Fully assess the business risks before performing this operation.
//
// @param request - DisableAgentRuntimeRequest
//
// @return DisableAgentRuntimeResponse
func (client *Client) DisableAgentRuntime(request *DisableAgentRuntimeRequest) (_result *DisableAgentRuntimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAgentRuntimeResponse{}
	_body, _err := client.DisableAgentRuntimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the sandbox and Edge Routine capabilities for a Supabase instance. Read the operation description before you call this operation.
//
// Description:
//
// During the public preview, the sandbox and Edge Routine features are free of charge.
//
// ### Before you begin
//
// Before you use this feature, [authorize the service-linked role](https://api.aliyun.com/api/ResourceManager/2020-03-31/CreateServiceLinkedRole?spm=api-workbench.API%20Document.0.0.4ea75094rJgPzK&RegionId=cn-beijing&tab=DEBUG&params=%7B%2522ServiceName%2522:%2522supabase.rdsai.aliyuncs.com%2522%7D&sdkStyle=old). The service-linked role used is [AliyunServiceRoleForRDSAISupabase](https://www.alibabacloud.com/help/en/ram/product-overview/services-that-work-with-service-linked-roles).
//
// @param request - EnableAgentRuntimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAgentRuntimeResponse
func (client *Client) EnableAgentRuntimeWithOptions(request *EnableAgentRuntimeRequest, runtime *dara.RuntimeOptions) (_result *EnableAgentRuntimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAgentRuntime"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAgentRuntimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the sandbox and Edge Routine capabilities for a Supabase instance. Read the operation description before you call this operation.
//
// Description:
//
// During the public preview, the sandbox and Edge Routine features are free of charge.
//
// ### Before you begin
//
// Before you use this feature, [authorize the service-linked role](https://api.aliyun.com/api/ResourceManager/2020-03-31/CreateServiceLinkedRole?spm=api-workbench.API%20Document.0.0.4ea75094rJgPzK&RegionId=cn-beijing&tab=DEBUG&params=%7B%2522ServiceName%2522:%2522supabase.rdsai.aliyuncs.com%2522%7D&sdkStyle=old). The service-linked role used is [AliyunServiceRoleForRDSAISupabase](https://www.alibabacloud.com/help/en/ram/product-overview/services-that-work-with-service-linked-roles).
//
// @param request - EnableAgentRuntimeRequest
//
// @return EnableAgentRuntimeResponse
func (client *Client) EnableAgentRuntime(request *EnableAgentRuntimeRequest) (_result *EnableAgentRuntimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableAgentRuntimeResponse{}
	_body, _err := client.EnableAgentRuntimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of large language models supported by the current RDS Supabase AI gateway.
//
// Description:
//
// ### Applicable engine
//
// RDS Supabase.
//
// @param request - GetAvailableLLMModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAvailableLLMModelsResponse
func (client *Client) GetAvailableLLMModelsWithOptions(request *GetAvailableLLMModelsRequest, runtime *dara.RuntimeOptions) (_result *GetAvailableLLMModelsResponse, _err error) {
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
		Action:      dara.String("GetAvailableLLMModels"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAvailableLLMModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of large language models supported by the current RDS Supabase AI gateway.
//
// Description:
//
// ### Applicable engine
//
// RDS Supabase.
//
// @param request - GetAvailableLLMModelsRequest
//
// @return GetAvailableLLMModelsResponse
func (client *Client) GetAvailableLLMModels(request *GetAvailableLLMModelsRequest) (_result *GetAvailableLLMModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAvailableLLMModelsResponse{}
	_body, _err := client.GetAvailableLLMModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the available models for a session.
//
// @param request - GetChatModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatModelResponse
func (client *Client) GetChatModelWithOptions(request *GetChatModelRequest, runtime *dara.RuntimeOptions) (_result *GetChatModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatModel"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the available models for a session.
//
// @param request - GetChatModelRequest
//
// @return GetChatModelResponse
func (client *Client) GetChatModel(request *GetChatModelRequest) (_result *GetChatModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatModelResponse{}
	_body, _err := client.GetChatModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a datasource configuration.
//
// Description:
//
// Retrieves the details of a connection.
//
// @param request - GetContextDatabaseSyncConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncConnectionResponse
func (client *Client) GetContextDatabaseSyncConnectionWithOptions(request *GetContextDatabaseSyncConnectionRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncConnection"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a datasource configuration.
//
// Description:
//
// Retrieves the details of a connection.
//
// @param request - GetContextDatabaseSyncConnectionRequest
//
// @return GetContextDatabaseSyncConnectionResponse
func (client *Client) GetContextDatabaseSyncConnection(request *GetContextDatabaseSyncConnectionRequest) (_result *GetContextDatabaseSyncConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncConnectionResponse{}
	_body, _err := client.GetContextDatabaseSyncConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorization scope of a data source.
//
// Description:
//
// Queries the authorization scope of a connection.
//
// @param request - GetContextDatabaseSyncConnectionAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncConnectionAuthorizationResponse
func (client *Client) GetContextDatabaseSyncConnectionAuthorizationWithOptions(request *GetContextDatabaseSyncConnectionAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncConnectionAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncConnectionAuthorization"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncConnectionAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization scope of a data source.
//
// Description:
//
// Queries the authorization scope of a connection.
//
// @param request - GetContextDatabaseSyncConnectionAuthorizationRequest
//
// @return GetContextDatabaseSyncConnectionAuthorizationResponse
func (client *Client) GetContextDatabaseSyncConnectionAuthorization(request *GetContextDatabaseSyncConnectionAuthorizationRequest) (_result *GetContextDatabaseSyncConnectionAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncConnectionAuthorizationResponse{}
	_body, _err := client.GetContextDatabaseSyncConnectionAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a discovery task by task ID.
//
// Description:
//
// Queries the status of a discovery task by task ID.
//
// @param request - GetContextDatabaseSyncDiscoveryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncDiscoveryTaskResponse
func (client *Client) GetContextDatabaseSyncDiscoveryTaskWithOptions(request *GetContextDatabaseSyncDiscoveryTaskRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncDiscoveryTaskResponse, _err error) {
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

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncDiscoveryTask"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncDiscoveryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a discovery task by task ID.
//
// Description:
//
// Queries the status of a discovery task by task ID.
//
// @param request - GetContextDatabaseSyncDiscoveryTaskRequest
//
// @return GetContextDatabaseSyncDiscoveryTaskResponse
func (client *Client) GetContextDatabaseSyncDiscoveryTask(request *GetContextDatabaseSyncDiscoveryTaskRequest) (_result *GetContextDatabaseSyncDiscoveryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncDiscoveryTaskResponse{}
	_body, _err := client.GetContextDatabaseSyncDiscoveryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an import job.
//
// Description:
//
// Retrieves the details of a job.
//
// @param request - GetContextDatabaseSyncJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncJobResponse
func (client *Client) GetContextDatabaseSyncJobWithOptions(request *GetContextDatabaseSyncJobRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncJob"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an import job.
//
// Description:
//
// Retrieves the details of a job.
//
// @param request - GetContextDatabaseSyncJobRequest
//
// @return GetContextDatabaseSyncJobResponse
func (client *Client) GetContextDatabaseSyncJob(request *GetContextDatabaseSyncJobRequest) (_result *GetContextDatabaseSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncJobResponse{}
	_body, _err := client.GetContextDatabaseSyncJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether an import task can be started.
//
// Description:
//
// Queries whether a job can be started.
//
// @param request - GetContextDatabaseSyncJobReadinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncJobReadinessResponse
func (client *Client) GetContextDatabaseSyncJobReadinessWithOptions(request *GetContextDatabaseSyncJobReadinessRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncJobReadinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncJobReadiness"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncJobReadinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether an import task can be started.
//
// Description:
//
// Queries whether a job can be started.
//
// @param request - GetContextDatabaseSyncJobReadinessRequest
//
// @return GetContextDatabaseSyncJobReadinessResponse
func (client *Client) GetContextDatabaseSyncJobReadiness(request *GetContextDatabaseSyncJobReadinessRequest) (_result *GetContextDatabaseSyncJobReadinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncJobReadinessResponse{}
	_body, _err := client.GetContextDatabaseSyncJobReadinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scheduling configuration of an import task.
//
// Description:
//
// Queries the scheduling configuration of a job.
//
// @param request - GetContextDatabaseSyncJobScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncJobScheduleResponse
func (client *Client) GetContextDatabaseSyncJobScheduleWithOptions(request *GetContextDatabaseSyncJobScheduleRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncJobScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncJobSchedule"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncJobScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scheduling configuration of an import task.
//
// Description:
//
// Queries the scheduling configuration of a job.
//
// @param request - GetContextDatabaseSyncJobScheduleRequest
//
// @return GetContextDatabaseSyncJobScheduleResponse
func (client *Client) GetContextDatabaseSyncJobSchedule(request *GetContextDatabaseSyncJobScheduleRequest) (_result *GetContextDatabaseSyncJobScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncJobScheduleResponse{}
	_body, _err := client.GetContextDatabaseSyncJobScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a run.
//
// Description:
//
// Retrieves the details of a run.
//
// @param request - GetContextDatabaseSyncRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncRunResponse
func (client *Client) GetContextDatabaseSyncRunWithOptions(request *GetContextDatabaseSyncRunRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RunId) {
		query["RunId"] = request.RunId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncRun"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a run.
//
// Description:
//
// Retrieves the details of a run.
//
// @param request - GetContextDatabaseSyncRunRequest
//
// @return GetContextDatabaseSyncRunResponse
func (client *Client) GetContextDatabaseSyncRun(request *GetContextDatabaseSyncRunRequest) (_result *GetContextDatabaseSyncRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncRunResponse{}
	_body, _err := client.GetContextDatabaseSyncRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data selection scope.
//
// Description:
//
// Retrieves the details of a source.
//
// @param request - GetContextDatabaseSyncSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncSourceResponse
func (client *Client) GetContextDatabaseSyncSourceWithOptions(request *GetContextDatabaseSyncSourceRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncSource"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data selection scope.
//
// Description:
//
// Retrieves the details of a source.
//
// @param request - GetContextDatabaseSyncSourceRequest
//
// @return GetContextDatabaseSyncSourceResponse
func (client *Client) GetContextDatabaseSyncSource(request *GetContextDatabaseSyncSourceRequest) (_result *GetContextDatabaseSyncSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncSourceResponse{}
	_body, _err := client.GetContextDatabaseSyncSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current node selection in the data selection scope.
//
// Description:
//
// Queries the current node selection of the source.
//
// @param request - GetContextDatabaseSyncSourceSelectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextDatabaseSyncSourceSelectionResponse
func (client *Client) GetContextDatabaseSyncSourceSelectionWithOptions(request *GetContextDatabaseSyncSourceSelectionRequest, runtime *dara.RuntimeOptions) (_result *GetContextDatabaseSyncSourceSelectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextDatabaseSyncSourceSelection"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextDatabaseSyncSourceSelectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current node selection in the data selection scope.
//
// Description:
//
// Queries the current node selection of the source.
//
// @param request - GetContextDatabaseSyncSourceSelectionRequest
//
// @return GetContextDatabaseSyncSourceSelectionResponse
func (client *Client) GetContextDatabaseSyncSourceSelection(request *GetContextDatabaseSyncSourceSelectionRequest) (_result *GetContextDatabaseSyncSourceSelectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextDatabaseSyncSourceSelectionResponse{}
	_body, _err := client.GetContextDatabaseSyncSourceSelectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical conversations of a user.
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

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Queries the historical conversations of a user.
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
// Queries a dedicated agent created by the user.
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
// Queries a dedicated agent created by the user.
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
// Retrieves the details of an inspection report by report ID.
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
// Retrieves the details of an inspection report by report ID.
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
// Queries the details of specific conversation messages.
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

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Queries the details of specific conversation messages.
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
// Retrieves order information for RDS AI Assistant Ultimate Edition.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - GetModelOperatorOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelOperatorOrderResponse
func (client *Client) GetModelOperatorOrderWithOptions(request *GetModelOperatorOrderRequest, runtime *dara.RuntimeOptions) (_result *GetModelOperatorOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelOperatorOrder"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelOperatorOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves order information for RDS AI Assistant Ultimate Edition.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - GetModelOperatorOrderRequest
//
// @return GetModelOperatorOrderResponse
func (client *Client) GetModelOperatorOrder(request *GetModelOperatorOrderRequest) (_result *GetModelOperatorOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetModelOperatorOrderResponse{}
	_body, _err := client.GetModelOperatorOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of all instance IDs included in a specified scheduled inspection configuration.
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
// Queries the list of all instance IDs included in a specified scheduled inspection configuration.
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
// Queries all inspection reports under a specified scheduled task, with support for filtering by time range and pagination.
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
// Queries all inspection reports under a specified scheduled task, with support for filtering by time range and pagination.
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
// Retrieves the details of a specified skill. You can retrieve your own skills or system preset skills.
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
	if !dara.IsNil(request.Download) {
		query["Download"] = request.Download
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Retrieves the details of a specified skill. You can retrieve your own skills or system preset skills.
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
// Queries the list of individual inspection reports for all non-scheduled tasks under a specified user, with pagination support.
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
// Queries the list of individual inspection reports for all non-scheduled tasks under a specified user, with pagination support.
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
// Queries custom API keys.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - ListApiKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiKeysResponse
func (client *Client) ListApiKeysWithOptions(request *ListApiKeysRequest, runtime *dara.RuntimeOptions) (_result *ListApiKeysResponse, _err error) {
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

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiKeys"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom API keys.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - ListApiKeysRequest
//
// @return ListApiKeysResponse
func (client *Client) ListApiKeys(request *ListApiKeysRequest) (_result *ListApiKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApiKeysResponse{}
	_body, _err := client.ListApiKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the API keys for a context database.
//
// Description:
//
// Lists the API keys under a specified member. The plaintext key values are not returned.
//
// @param request - ListContextDatabaseApiKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseApiKeysResponse
func (client *Client) ListContextDatabaseApiKeysWithOptions(request *ListContextDatabaseApiKeysRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseApiKeysResponse, _err error) {
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

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseApiKeys"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseApiKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the API keys for a context database.
//
// Description:
//
// Lists the API keys under a specified member. The plaintext key values are not returned.
//
// @param request - ListContextDatabaseApiKeysRequest
//
// @return ListContextDatabaseApiKeysResponse
func (client *Client) ListContextDatabaseApiKeys(request *ListContextDatabaseApiKeysRequest) (_result *ListContextDatabaseApiKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseApiKeysResponse{}
	_body, _err := client.ListContextDatabaseApiKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists context database members.
//
// Description:
//
// Lists all members in a specified workspace. Each member includes a list of API keys (plaintext values are not returned).
//
// @param request - ListContextDatabaseMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseMembersResponse
func (client *Client) ListContextDatabaseMembersWithOptions(request *ListContextDatabaseMembersRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseMembersResponse, _err error) {
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

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseMembers"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists context database members.
//
// Description:
//
// Lists all members in a specified workspace. Each member includes a list of API keys (plaintext values are not returned).
//
// @param request - ListContextDatabaseMembersRequest
//
// @return ListContextDatabaseMembersResponse
func (client *Client) ListContextDatabaseMembers(request *ListContextDatabaseMembersRequest) (_result *ListContextDatabaseMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseMembersResponse{}
	_body, _err := client.ListContextDatabaseMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the datasource configurations in a workspace.
//
// Description:
//
// The list operation returns paginated results by default. Page and PageSize are optional parameters.
//
// The Data list field in the response is returned as an array. Total indicates the total number of records that match the specified conditions.
//
// @param request - ListContextDatabaseSyncConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncConnectionsResponse
func (client *Client) ListContextDatabaseSyncConnectionsWithOptions(request *ListContextDatabaseSyncConnectionsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncConnections"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the datasource configurations in a workspace.
//
// Description:
//
// The list operation returns paginated results by default. Page and PageSize are optional parameters.
//
// The Data list field in the response is returned as an array. Total indicates the total number of records that match the specified conditions.
//
// @param request - ListContextDatabaseSyncConnectionsRequest
//
// @return ListContextDatabaseSyncConnectionsResponse
func (client *Client) ListContextDatabaseSyncConnections(request *ListContextDatabaseSyncConnectionsRequest) (_result *ListContextDatabaseSyncConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncConnectionsResponse{}
	_body, _err := client.ListContextDatabaseSyncConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the documents associated with an import task.
//
// Description:
//
// By default, this operation returns paginated results. The Page and PageSize parameters are optional. The Data field in the response is returned as an array, and the Total field indicates the total number of matching entries.
//
// @param request - ListContextDatabaseSyncJobDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncJobDocumentsResponse
func (client *Client) ListContextDatabaseSyncJobDocumentsWithOptions(request *ListContextDatabaseSyncJobDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncJobDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncJobDocuments"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncJobDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the documents associated with an import task.
//
// Description:
//
// By default, this operation returns paginated results. The Page and PageSize parameters are optional. The Data field in the response is returned as an array, and the Total field indicates the total number of matching entries.
//
// @param request - ListContextDatabaseSyncJobDocumentsRequest
//
// @return ListContextDatabaseSyncJobDocumentsResponse
func (client *Client) ListContextDatabaseSyncJobDocuments(request *ListContextDatabaseSyncJobDocumentsRequest) (_result *ListContextDatabaseSyncJobDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncJobDocumentsResponse{}
	_body, _err := client.ListContextDatabaseSyncJobDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the exclusions of an import task.
//
// Description:
//
// This operation uses pagination by default. Page and PageSize are optional parameters.
//
// The Data field in the response returns an array. Total indicates the total number of records that match the specified conditions.
//
// @param request - ListContextDatabaseSyncJobExclusionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncJobExclusionsResponse
func (client *Client) ListContextDatabaseSyncJobExclusionsWithOptions(request *ListContextDatabaseSyncJobExclusionsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncJobExclusionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncJobExclusions"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncJobExclusionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the exclusions of an import task.
//
// Description:
//
// This operation uses pagination by default. Page and PageSize are optional parameters.
//
// The Data field in the response returns an array. Total indicates the total number of records that match the specified conditions.
//
// @param request - ListContextDatabaseSyncJobExclusionsRequest
//
// @return ListContextDatabaseSyncJobExclusionsResponse
func (client *Client) ListContextDatabaseSyncJobExclusions(request *ListContextDatabaseSyncJobExclusionsRequest) (_result *ListContextDatabaseSyncJobExclusionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncJobExclusionsResponse{}
	_body, _err := client.ListContextDatabaseSyncJobExclusionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists import tasks.
//
// Description:
//
// The list operation returns paginated results by default. Page and PageSize are optional parameters.
//
// The Data list field in the response is returned as an array. Total indicates the total number of records that match the specified conditions.
//
// @param request - ListContextDatabaseSyncJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncJobsResponse
func (client *Client) ListContextDatabaseSyncJobsWithOptions(request *ListContextDatabaseSyncJobsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeBaseId) {
		query["KnowledgeBaseId"] = request.KnowledgeBaseId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncJobs"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists import tasks.
//
// Description:
//
// The list operation returns paginated results by default. Page and PageSize are optional parameters.
//
// The Data list field in the response is returned as an array. Total indicates the total number of records that match the specified conditions.
//
// @param request - ListContextDatabaseSyncJobsRequest
//
// @return ListContextDatabaseSyncJobsResponse
func (client *Client) ListContextDatabaseSyncJobs(request *ListContextDatabaseSyncJobsRequest) (_result *ListContextDatabaseSyncJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncJobsResponse{}
	_body, _err := client.ListContextDatabaseSyncJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists available synchronization plugins.
//
// Description:
//
// This operation returns a fixed-length array of available plugins in the workspace and does not support pagination.
//
// The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncPluginsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncPluginsResponse
func (client *Client) ListContextDatabaseSyncPluginsWithOptions(request *ListContextDatabaseSyncPluginsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncPluginsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncPlugins"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists available synchronization plugins.
//
// Description:
//
// This operation returns a fixed-length array of available plugins in the workspace and does not support pagination.
//
// The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncPluginsRequest
//
// @return ListContextDatabaseSyncPluginsResponse
func (client *Client) ListContextDatabaseSyncPlugins(request *ListContextDatabaseSyncPluginsRequest) (_result *ListContextDatabaseSyncPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncPluginsResponse{}
	_body, _err := client.ListContextDatabaseSyncPluginsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the run records.
//
// Description:
//
// The list API returns paginated results by default. The Page and PageSize parameters are optional. Run records are a list of RunItems that show the processing results of each document in each import run.
//
// The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncRunItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncRunItemsResponse
func (client *Client) ListContextDatabaseSyncRunItemsWithOptions(request *ListContextDatabaseSyncRunItemsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncRunItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RunId) {
		query["RunId"] = request.RunId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncRunItems"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncRunItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the run records.
//
// Description:
//
// The list API returns paginated results by default. The Page and PageSize parameters are optional. Run records are a list of RunItems that show the processing results of each document in each import run.
//
// The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncRunItemsRequest
//
// @return ListContextDatabaseSyncRunItemsResponse
func (client *Client) ListContextDatabaseSyncRunItems(request *ListContextDatabaseSyncRunItemsRequest) (_result *ListContextDatabaseSyncRunItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncRunItemsResponse{}
	_body, _err := client.ListContextDatabaseSyncRunItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出 Run 日志
//
// Description:
//
// 列表接口默认分页返回，Page 与 PageSize 为可选参数。
//
// 响应中的 Data 列表字段使用数组返回，Total 表示符合条件的总条数。
//
// @param request - ListContextDatabaseSyncRunLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncRunLogsResponse
func (client *Client) ListContextDatabaseSyncRunLogsWithOptions(request *ListContextDatabaseSyncRunLogsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncRunLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RunId) {
		query["RunId"] = request.RunId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncRunLogs"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncRunLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出 Run 日志
//
// Description:
//
// 列表接口默认分页返回，Page 与 PageSize 为可选参数。
//
// 响应中的 Data 列表字段使用数组返回，Total 表示符合条件的总条数。
//
// @param request - ListContextDatabaseSyncRunLogsRequest
//
// @return ListContextDatabaseSyncRunLogsResponse
func (client *Client) ListContextDatabaseSyncRunLogs(request *ListContextDatabaseSyncRunLogsRequest) (_result *ListContextDatabaseSyncRunLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncRunLogsResponse{}
	_body, _err := client.ListContextDatabaseSyncRunLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the runs.
//
// Description:
//
// By default, this operation returns paginated results. The Page and PageSize parameters are optional.
//
// The Data field in the response is returned as an array, and the Total field indicates the total number of items that meet the query conditions.
//
// @param request - ListContextDatabaseSyncRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncRunsResponse
func (client *Client) ListContextDatabaseSyncRunsWithOptions(request *ListContextDatabaseSyncRunsRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncRuns"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the runs.
//
// Description:
//
// By default, this operation returns paginated results. The Page and PageSize parameters are optional.
//
// The Data field in the response is returned as an array, and the Total field indicates the total number of items that meet the query conditions.
//
// @param request - ListContextDatabaseSyncRunsRequest
//
// @return ListContextDatabaseSyncRunsResponse
func (client *Client) ListContextDatabaseSyncRuns(request *ListContextDatabaseSyncRunsRequest) (_result *ListContextDatabaseSyncRunsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncRunsResponse{}
	_body, _err := client.ListContextDatabaseSyncRunsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the discovered tree nodes for the data selection scope.
//
// Description:
//
// List operations return paginated results by default. Page and PageSize are optional parameters. The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncSourceTreeNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncSourceTreeNodesResponse
func (client *Client) ListContextDatabaseSyncSourceTreeNodesWithOptions(request *ListContextDatabaseSyncSourceTreeNodesRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncSourceTreeNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncSourceTreeNodes"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncSourceTreeNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the discovered tree nodes for the data selection scope.
//
// Description:
//
// List operations return paginated results by default. Page and PageSize are optional parameters. The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncSourceTreeNodesRequest
//
// @return ListContextDatabaseSyncSourceTreeNodesResponse
func (client *Client) ListContextDatabaseSyncSourceTreeNodes(request *ListContextDatabaseSyncSourceTreeNodesRequest) (_result *ListContextDatabaseSyncSourceTreeNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncSourceTreeNodesResponse{}
	_body, _err := client.ListContextDatabaseSyncSourceTreeNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the data selection scopes.
//
// Description:
//
// By default, the list operation returns paginated results. Page and PageSize are optional parameters. The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseSyncSourcesResponse
func (client *Client) ListContextDatabaseSyncSourcesWithOptions(request *ListContextDatabaseSyncSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseSyncSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseSyncSources"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseSyncSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the data selection scopes.
//
// Description:
//
// By default, the list operation returns paginated results. Page and PageSize are optional parameters. The Data list field in the response is returned as an array, and Total indicates the total number of entries that meet the conditions.
//
// @param request - ListContextDatabaseSyncSourcesRequest
//
// @return ListContextDatabaseSyncSourcesResponse
func (client *Client) ListContextDatabaseSyncSources(request *ListContextDatabaseSyncSourcesRequest) (_result *ListContextDatabaseSyncSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseSyncSourcesResponse{}
	_body, _err := client.ListContextDatabaseSyncSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists ContextDB workspaces.
//
// Description:
//
// Lists ContextDB workspaces.
//
// @param request - ListContextDatabaseWorkspacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextDatabaseWorkspacesResponse
func (client *Client) ListContextDatabaseWorkspacesWithOptions(request *ListContextDatabaseWorkspacesRequest, runtime *dara.RuntimeOptions) (_result *ListContextDatabaseWorkspacesResponse, _err error) {
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContextDatabaseWorkspaces"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextDatabaseWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists ContextDB workspaces.
//
// Description:
//
// Lists ContextDB workspaces.
//
// @param request - ListContextDatabaseWorkspacesRequest
//
// @return ListContextDatabaseWorkspacesResponse
func (client *Client) ListContextDatabaseWorkspaces(request *ListContextDatabaseWorkspacesRequest) (_result *ListContextDatabaseWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContextDatabaseWorkspacesResponse{}
	_body, _err := client.ListContextDatabaseWorkspacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of dedicated agents created by the user.
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
// Queries the list of dedicated agents created by the user.
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
// Queries the list of Agent tools for a user.
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
// Queries the list of Agent tools for a user.
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
// Retrieves the AI gateway token usage for an RDS Supabase instance.
//
// Description:
//
// ### Applicable engine
//
// RDS Supabase.
//
// @param request - ListLLMTokenUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLLMTokenUsageResponse
func (client *Client) ListLLMTokenUsageWithOptions(request *ListLLMTokenUsageRequest, runtime *dara.RuntimeOptions) (_result *ListLLMTokenUsageResponse, _err error) {
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

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
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
		Action:      dara.String("ListLLMTokenUsage"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLLMTokenUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the AI gateway token usage for an RDS Supabase instance.
//
// Description:
//
// ### Applicable engine
//
// RDS Supabase.
//
// @param request - ListLLMTokenUsageRequest
//
// @return ListLLMTokenUsageResponse
func (client *Client) ListLLMTokenUsage(request *ListLLMTokenUsageRequest) (_result *ListLLMTokenUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLLMTokenUsageResponse{}
	_body, _err := client.ListLLMTokenUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of basic information about all inspection configurations under a specified user UID.
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
// Queries the list of basic information about all inspection configurations under a specified user UID.
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
// Retrieves the custom skills of the current user and all system-preset skills.
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

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Retrieves the custom skills of the current user and all system-preset skills.
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
// Modifies the modules of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param tmpReq - ModifyAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppInstanceResponse
func (client *Client) ModifyAppInstanceWithOptions(tmpReq *ModifyAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Components) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, dara.String("Components"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ComponentsShrink) {
		query["Components"] = request.ComponentsShrink
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
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
		Action:      dara.String("ModifyAppInstance"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the modules of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifyAppInstanceRequest
//
// @return ModifyAppInstanceResponse
func (client *Client) ModifyAppInstance(request *ModifyAppInstanceRequest) (_result *ModifyAppInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppInstanceResponse{}
	_body, _err := client.ModifyAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the authentication configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Modifies the authentication configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Modifies the general configurations of an instance, such as network EIP and NAT configurations.
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
// Modifies the general configurations of an instance, such as network EIP and NAT configurations.
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
// Modifies the IP whitelist of an RDS AI application instance.
//
// Description:
//
// ### Supported engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
// Modifies the IP whitelist of an RDS AI application instance.
//
// Description:
//
// ### Supported engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Modifies the RAG Agent configuration of an RDS AI application instance.
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
// Modifies the RAG Agent configuration of an RDS AI application instance.
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
// Modifies the SSL configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
// Modifies the SSL configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Modifies the storage configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// > Currently, only Object Storage Service (OSS) is supported.
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
// Modifies the storage configuration of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// > Currently, only Object Storage Service (OSS) is supported.
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
// Batch modifies the SSL configurations of RDS AI application instances.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Batch modifies the SSL configurations of RDS AI application instances.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// Modifies message feedback.
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

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Modifies message feedback.
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
// Modifies a sandbox template.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param tmpReq - ModifySandboxTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySandboxTemplateResponse
func (client *Client) ModifySandboxTemplateWithOptions(tmpReq *ModifySandboxTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifySandboxTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifySandboxTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultCpu) {
		query["DefaultCpu"] = request.DefaultCpu
	}

	if !dara.IsNil(request.DefaultMemory) {
		query["DefaultMemory"] = request.DefaultMemory
	}

	if !dara.IsNil(request.Image) {
		query["Image"] = request.Image
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySandboxTemplate"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySandboxTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a sandbox template.
//
// Description:
//
// ### Applicable engine
//
// # RDS Supabase
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ModifySandboxTemplateRequest
//
// @return ModifySandboxTemplateResponse
func (client *Client) ModifySandboxTemplate(request *ModifySandboxTemplateRequest) (_result *ModifySandboxTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySandboxTemplateResponse{}
	_body, _err := client.ModifySandboxTemplateWithOptions(request, runtime)
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

	if !dara.IsNil(request.InspectionItems) {
		query["InspectionItems"] = request.InspectionItems
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
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
// Modifies the IP whitelist of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - ModifyWhitelistIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWhitelistIpsResponse
func (client *Client) ModifyWhitelistIpsWithOptions(request *ModifyWhitelistIpsRequest, runtime *dara.RuntimeOptions) (_result *ModifyWhitelistIpsResponse, _err error) {
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

	if !dara.IsNil(request.IpWhitelist) {
		query["IpWhitelist"] = request.IpWhitelist
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWhitelistIps"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWhitelistIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP whitelist of an RDS AI Assistant Ultimate Edition instance.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - ModifyWhitelistIpsRequest
//
// @return ModifyWhitelistIpsResponse
func (client *Client) ModifyWhitelistIps(request *ModifyWhitelistIpsRequest) (_result *ModifyWhitelistIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWhitelistIpsResponse{}
	_body, _err := client.ModifyWhitelistIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an exclusion item for an import job.
//
// Description:
//
// Deletes an exclusion item for a job.
//
// @param request - RemoveContextDatabaseSyncJobExclusionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveContextDatabaseSyncJobExclusionResponse
func (client *Client) RemoveContextDatabaseSyncJobExclusionWithOptions(request *RemoveContextDatabaseSyncJobExclusionRequest, runtime *dara.RuntimeOptions) (_result *RemoveContextDatabaseSyncJobExclusionResponse, _err error) {
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

	if !dara.IsNil(request.DocumentId) {
		query["DocumentId"] = request.DocumentId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveContextDatabaseSyncJobExclusion"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveContextDatabaseSyncJobExclusionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an exclusion item for an import job.
//
// Description:
//
// Deletes an exclusion item for a job.
//
// @param request - RemoveContextDatabaseSyncJobExclusionRequest
//
// @return RemoveContextDatabaseSyncJobExclusionResponse
func (client *Client) RemoveContextDatabaseSyncJobExclusion(request *RemoveContextDatabaseSyncJobExclusionRequest) (_result *RemoveContextDatabaseSyncJobExclusionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveContextDatabaseSyncJobExclusionResponse{}
	_body, _err := client.RemoveContextDatabaseSyncJobExclusionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a custom API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - RenameApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameApiKeyResponse
func (client *Client) RenameApiKeyWithOptions(request *RenameApiKeyRequest, runtime *dara.RuntimeOptions) (_result *RenameApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameApiKey"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a custom API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - RenameApiKeyRequest
//
// @return RenameApiKeyResponse
func (client *Client) RenameApiKey(request *RenameApiKeyRequest) (_result *RenameApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenameApiKeyResponse{}
	_body, _err := client.RenameApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - ResetApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetApiKeyResponse
func (client *Client) ResetApiKeyWithOptions(request *ResetApiKeyRequest, runtime *dara.RuntimeOptions) (_result *ResetApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetApiKey"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - ResetApiKeyRequest
//
// @return ResetApiKeyResponse
func (client *Client) ResetApiKey(request *ResetApiKeyRequest) (_result *ResetApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetApiKeyResponse{}
	_body, _err := client.ResetApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the keys of an RDS AI instance, including AnonKey, ServiceKey, and JwtSecret. After the reset, the old keys and secret become invalid immediately. Ensure that your applications are updated accordingly.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ResetInstanceKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetInstanceKeysResponse
func (client *Client) ResetInstanceKeysWithOptions(request *ResetInstanceKeysRequest, runtime *dara.RuntimeOptions) (_result *ResetInstanceKeysResponse, _err error) {
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
		Action:      dara.String("ResetInstanceKeys"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetInstanceKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the keys of an RDS AI instance, including AnonKey, ServiceKey, and JwtSecret. After the reset, the old keys and secret become invalid immediately. Ensure that your applications are updated accordingly.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// @param request - ResetInstanceKeysRequest
//
// @return ResetInstanceKeysResponse
func (client *Client) ResetInstanceKeys(request *ResetInstanceKeysRequest) (_result *ResetInstanceKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetInstanceKeysResponse{}
	_body, _err := client.ResetInstanceKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the logon password and database access password of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// > Currently, only the password of the RDS Supabase Dashboard user can be modified.
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

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
// Resets the logon password and database access password of an RDS AI application instance.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
//
// [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html)
//
// > Currently, only the password of the RDS Supabase Dashboard user can be modified.
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
// Restarts an ApsaraDB RDS AI application instance that is in the Running state.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Restarts an ApsaraDB RDS AI application instance that is in the Running state.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related documentation
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
// 基于快照重试失败明细
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。 重试会基于失败项生成新的 Run，原 Run 状态保持不变。
//
// @param request - RetryFailedContextDatabaseSyncRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryFailedContextDatabaseSyncRunResponse
func (client *Client) RetryFailedContextDatabaseSyncRunWithOptions(request *RetryFailedContextDatabaseSyncRunRequest, runtime *dara.RuntimeOptions) (_result *RetryFailedContextDatabaseSyncRunResponse, _err error) {
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

	if !dara.IsNil(request.RunId) {
		query["RunId"] = request.RunId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryFailedContextDatabaseSyncRun"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryFailedContextDatabaseSyncRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 基于快照重试失败明细
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。 重试会基于失败项生成新的 Run，原 Run 状态保持不变。
//
// @param request - RetryFailedContextDatabaseSyncRunRequest
//
// @return RetryFailedContextDatabaseSyncRunResponse
func (client *Client) RetryFailedContextDatabaseSyncRun(request *RetryFailedContextDatabaseSyncRunRequest) (_result *RetryFailedContextDatabaseSyncRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryFailedContextDatabaseSyncRunResponse{}
	_body, _err := client.RetryFailedContextDatabaseSyncRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes an API key for a context database.
//
// Description:
//
// Revokes an API key.
//
// @param request - RevokeContextDatabaseApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeContextDatabaseApiKeyResponse
func (client *Client) RevokeContextDatabaseApiKeyWithOptions(request *RevokeContextDatabaseApiKeyRequest, runtime *dara.RuntimeOptions) (_result *RevokeContextDatabaseApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeContextDatabaseApiKey"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeContextDatabaseApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes an API key for a context database.
//
// Description:
//
// Revokes an API key.
//
// @param request - RevokeContextDatabaseApiKeyRequest
//
// @return RevokeContextDatabaseApiKeyResponse
func (client *Client) RevokeContextDatabaseApiKey(request *RevokeContextDatabaseApiKeyRequest) (_result *RevokeContextDatabaseApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeContextDatabaseApiKeyResponse{}
	_body, _err := client.RevokeContextDatabaseApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually triggers a run.
//
// Description:
//
// Write operations support idempotence by passing a ClientToken. If no ClientToken is provided, the backend generates one automatically.
//
// @param request - StartContextDatabaseSyncRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartContextDatabaseSyncRunResponse
func (client *Client) StartContextDatabaseSyncRunWithOptions(request *StartContextDatabaseSyncRunRequest, runtime *dara.RuntimeOptions) (_result *StartContextDatabaseSyncRunResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartContextDatabaseSyncRun"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartContextDatabaseSyncRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually triggers a run.
//
// Description:
//
// Write operations support idempotence by passing a ClientToken. If no ClientToken is provided, the backend generates one automatically.
//
// @param request - StartContextDatabaseSyncRunRequest
//
// @return StartContextDatabaseSyncRunResponse
func (client *Client) StartContextDatabaseSyncRun(request *StartContextDatabaseSyncRunRequest) (_result *StartContextDatabaseSyncRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartContextDatabaseSyncRunResponse{}
	_body, _err := client.StartContextDatabaseSyncRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a snapshot discovery task for data selection scope configuration.
//
// Description:
//
// Write operations support idempotence by passing a ClientToken. If no ClientToken is provided, the backend generates one automatically. Discovery is an asynchronous operation. The operation returns a discovery task identifier. Poll for results to retrieve SelectedNodeKeys.
//
// @param request - StartContextDatabaseSyncSourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartContextDatabaseSyncSourceDiscoveryResponse
func (client *Client) StartContextDatabaseSyncSourceDiscoveryWithOptions(request *StartContextDatabaseSyncSourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *StartContextDatabaseSyncSourceDiscoveryResponse, _err error) {
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

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartContextDatabaseSyncSourceDiscovery"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartContextDatabaseSyncSourceDiscoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a snapshot discovery task for data selection scope configuration.
//
// Description:
//
// Write operations support idempotence by passing a ClientToken. If no ClientToken is provided, the backend generates one automatically. Discovery is an asynchronous operation. The operation returns a discovery task identifier. Poll for results to retrieve SelectedNodeKeys.
//
// @param request - StartContextDatabaseSyncSourceDiscoveryRequest
//
// @return StartContextDatabaseSyncSourceDiscoveryResponse
func (client *Client) StartContextDatabaseSyncSourceDiscovery(request *StartContextDatabaseSyncSourceDiscoveryRequest) (_result *StartContextDatabaseSyncSourceDiscoveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartContextDatabaseSyncSourceDiscoveryResponse{}
	_body, _err := client.StartContextDatabaseSyncSourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an RDS AI application instance that is in the Stopped state.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
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
// Starts an RDS AI application instance that is in the Stopped state.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
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
// Pauses an RDS AI application instance that is in the Running state.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
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
	if !dara.IsNil(request.BranchName) {
		query["BranchName"] = request.BranchName
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
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
// Pauses an RDS AI application instance that is in the Running state.
//
// Description:
//
// ### Applicable engine
//
// # RDS PostgreSQL
//
// ### Related feature documentation
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
// Modifies the quota of an API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param tmpReq - UpdateApiKeyQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiKeyQuotaResponse
func (client *Client) UpdateApiKeyQuotaWithOptions(tmpReq *UpdateApiKeyQuotaRequest, runtime *dara.RuntimeOptions) (_result *UpdateApiKeyQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateApiKeyQuotaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Keys) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, dara.String("Keys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KeysShrink) {
		query["Keys"] = request.KeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiKeyQuota"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiKeyQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the quota of an API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - UpdateApiKeyQuotaRequest
//
// @return UpdateApiKeyQuotaResponse
func (client *Client) UpdateApiKeyQuota(request *UpdateApiKeyQuotaRequest) (_result *UpdateApiKeyQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApiKeyQuotaResponse{}
	_body, _err := client.UpdateApiKeyQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the API key of a context database.
//
// Description:
//
// Updates the display metadata of an API key. At least one of Name or Description must be specified. The plaintext key is not reissued.
//
// @param request - UpdateContextDatabaseApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseApiKeyResponse
func (client *Client) UpdateContextDatabaseApiKeyWithOptions(request *UpdateContextDatabaseApiKeyRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseApiKeyResponse, _err error) {
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

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseApiKey"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the API key of a context database.
//
// Description:
//
// Updates the display metadata of an API key. At least one of Name or Description must be specified. The plaintext key is not reissued.
//
// @param request - UpdateContextDatabaseApiKeyRequest
//
// @return UpdateContextDatabaseApiKeyResponse
func (client *Client) UpdateContextDatabaseApiKey(request *UpdateContextDatabaseApiKeyRequest) (_result *UpdateContextDatabaseApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseApiKeyResponse{}
	_body, _err := client.UpdateContextDatabaseApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a context database member.
//
// Description:
//
// Updates the role or status of a member.
//
// @param request - UpdateContextDatabaseMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseMemberResponse
func (client *Client) UpdateContextDatabaseMemberWithOptions(request *UpdateContextDatabaseMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseMember"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a context database member.
//
// Description:
//
// Updates the role or status of a member.
//
// @param request - UpdateContextDatabaseMemberRequest
//
// @return UpdateContextDatabaseMemberResponse
func (client *Client) UpdateContextDatabaseMember(request *UpdateContextDatabaseMemberRequest) (_result *UpdateContextDatabaseMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseMemberResponse{}
	_body, _err := client.UpdateContextDatabaseMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新连接元数据
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。
//
// @param request - UpdateContextDatabaseSyncConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseSyncConnectionResponse
func (client *Client) UpdateContextDatabaseSyncConnectionWithOptions(request *UpdateContextDatabaseSyncConnectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseSyncConnectionResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseSyncConnection"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseSyncConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新连接元数据
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。
//
// @param request - UpdateContextDatabaseSyncConnectionRequest
//
// @return UpdateContextDatabaseSyncConnectionResponse
func (client *Client) UpdateContextDatabaseSyncConnection(request *UpdateContextDatabaseSyncConnectionRequest) (_result *UpdateContextDatabaseSyncConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseSyncConnectionResponse{}
	_body, _err := client.UpdateContextDatabaseSyncConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 单独轮换连接凭据
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。
//
// @param tmpReq - UpdateContextDatabaseSyncConnectionCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseSyncConnectionCredentialResponse
func (client *Client) UpdateContextDatabaseSyncConnectionCredentialWithOptions(tmpReq *UpdateContextDatabaseSyncConnectionCredentialRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseSyncConnectionCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateContextDatabaseSyncConnectionCredentialShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Credentials) {
		request.CredentialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Credentials, dara.String("Credentials"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.CredentialsShrink) {
		query["Credentials"] = request.CredentialsShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseSyncConnectionCredential"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseSyncConnectionCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单独轮换连接凭据
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。
//
// @param request - UpdateContextDatabaseSyncConnectionCredentialRequest
//
// @return UpdateContextDatabaseSyncConnectionCredentialResponse
func (client *Client) UpdateContextDatabaseSyncConnectionCredential(request *UpdateContextDatabaseSyncConnectionCredentialRequest) (_result *UpdateContextDatabaseSyncConnectionCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseSyncConnectionCredentialResponse{}
	_body, _err := client.UpdateContextDatabaseSyncConnectionCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Job 元数据
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。 Target 为隐含概念，调用方只需传入 KnowledgeBaseId，后端自动创建/匹配 Target。
//
// @param tmpReq - UpdateContextDatabaseSyncJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseSyncJobResponse
func (client *Client) UpdateContextDatabaseSyncJobWithOptions(tmpReq *UpdateContextDatabaseSyncJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseSyncJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateContextDatabaseSyncJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filters) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, dara.String("Filters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.FiltersShrink) {
		query["Filters"] = request.FiltersShrink
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TargetKnowledgeBaseId) {
		query["TargetKnowledgeBaseId"] = request.TargetKnowledgeBaseId
	}

	if !dara.IsNil(request.TargetKnowledgeBaseName) {
		query["TargetKnowledgeBaseName"] = request.TargetKnowledgeBaseName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseSyncJob"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Job 元数据
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。 Target 为隐含概念，调用方只需传入 KnowledgeBaseId，后端自动创建/匹配 Target。
//
// @param request - UpdateContextDatabaseSyncJobRequest
//
// @return UpdateContextDatabaseSyncJobResponse
func (client *Client) UpdateContextDatabaseSyncJob(request *UpdateContextDatabaseSyncJobRequest) (_result *UpdateContextDatabaseSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseSyncJobResponse{}
	_body, _err := client.UpdateContextDatabaseSyncJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Job 调度配置
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。
//
// @param tmpReq - UpdateContextDatabaseSyncJobScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseSyncJobScheduleResponse
func (client *Client) UpdateContextDatabaseSyncJobScheduleWithOptions(tmpReq *UpdateContextDatabaseSyncJobScheduleRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseSyncJobScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateContextDatabaseSyncJobScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DailyTimes) {
		request.DailyTimesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DailyTimes, dara.String("DailyTimes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DailyTimesShrink) {
		query["DailyTimes"] = request.DailyTimesShrink
	}

	if !dara.IsNil(request.IntervalMinutes) {
		query["IntervalMinutes"] = request.IntervalMinutes
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MissedPolicy) {
		query["MissedPolicy"] = request.MissedPolicy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseSyncJobSchedule"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseSyncJobScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Job 调度配置
//
// Description:
//
// 写操作可传入 ClientToken 实现幂等；未传时由后端自动生成。
//
// @param request - UpdateContextDatabaseSyncJobScheduleRequest
//
// @return UpdateContextDatabaseSyncJobScheduleResponse
func (client *Client) UpdateContextDatabaseSyncJobSchedule(request *UpdateContextDatabaseSyncJobScheduleRequest) (_result *UpdateContextDatabaseSyncJobScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseSyncJobScheduleResponse{}
	_body, _err := client.UpdateContextDatabaseSyncJobScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the metadata of a data selection scope.
//
// Description:
//
// Write operations support idempotence via ClientToken. If ClientToken is not provided, the backend generates one automatically.
//
// @param tmpReq - UpdateContextDatabaseSyncSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseSyncSourceResponse
func (client *Client) UpdateContextDatabaseSyncSourceWithOptions(tmpReq *UpdateContextDatabaseSyncSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseSyncSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateContextDatabaseSyncSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceConfig) {
		request.SourceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceConfig, dara.String("SourceConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SourceConfigShrink) {
		query["SourceConfig"] = request.SourceConfigShrink
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseSyncSource"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseSyncSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the metadata of a data selection scope.
//
// Description:
//
// Write operations support idempotence via ClientToken. If ClientToken is not provided, the backend generates one automatically.
//
// @param request - UpdateContextDatabaseSyncSourceRequest
//
// @return UpdateContextDatabaseSyncSourceResponse
func (client *Client) UpdateContextDatabaseSyncSource(request *UpdateContextDatabaseSyncSourceRequest) (_result *UpdateContextDatabaseSyncSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseSyncSourceResponse{}
	_body, _err := client.UpdateContextDatabaseSyncSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the explicit node selection for the data selection scope.
//
// Description:
//
// For write operations, you can pass the ClientToken parameter to ensure idempotency. If this parameter is not provided, the backend automatically generates one.
//
// @param tmpReq - UpdateContextDatabaseSyncSourceSelectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseSyncSourceSelectionResponse
func (client *Client) UpdateContextDatabaseSyncSourceSelectionWithOptions(tmpReq *UpdateContextDatabaseSyncSourceSelectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseSyncSourceSelectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateContextDatabaseSyncSourceSelectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SelectedNodeKeys) {
		request.SelectedNodeKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedNodeKeys, dara.String("SelectedNodeKeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.SelectedNodeKeysShrink) {
		query["SelectedNodeKeys"] = request.SelectedNodeKeysShrink
	}

	if !dara.IsNil(request.SelectionPolicy) {
		query["SelectionPolicy"] = request.SelectionPolicy
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseSyncSourceSelection"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseSyncSourceSelectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the explicit node selection for the data selection scope.
//
// Description:
//
// For write operations, you can pass the ClientToken parameter to ensure idempotency. If this parameter is not provided, the backend automatically generates one.
//
// @param request - UpdateContextDatabaseSyncSourceSelectionRequest
//
// @return UpdateContextDatabaseSyncSourceSelectionResponse
func (client *Client) UpdateContextDatabaseSyncSourceSelection(request *UpdateContextDatabaseSyncSourceSelectionRequest) (_result *UpdateContextDatabaseSyncSourceSelectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseSyncSourceSelectionResponse{}
	_body, _err := client.UpdateContextDatabaseSyncSourceSelectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a ContextDB workspace.
//
// Description:
//
// Modifies the name of a workspace.
//
// @param request - UpdateContextDatabaseWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextDatabaseWorkspaceResponse
func (client *Client) UpdateContextDatabaseWorkspaceWithOptions(request *UpdateContextDatabaseWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateContextDatabaseWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextDatabaseWorkspace"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextDatabaseWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a ContextDB workspace.
//
// Description:
//
// Modifies the name of a workspace.
//
// @param request - UpdateContextDatabaseWorkspaceRequest
//
// @return UpdateContextDatabaseWorkspaceResponse
func (client *Client) UpdateContextDatabaseWorkspace(request *UpdateContextDatabaseWorkspaceRequest) (_result *UpdateContextDatabaseWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContextDatabaseWorkspaceResponse{}
	_body, _err := client.UpdateContextDatabaseWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a Custom Agent.
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
// Updates a Custom Agent.
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
// Updates the alert threshold percentage for an API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param tmpReq - UpdateMOQuotaAlertThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMOQuotaAlertThresholdResponse
func (client *Client) UpdateMOQuotaAlertThresholdWithOptions(tmpReq *UpdateMOQuotaAlertThresholdRequest, runtime *dara.RuntimeOptions) (_result *UpdateMOQuotaAlertThresholdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMOQuotaAlertThresholdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Apikey) {
		request.ApikeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Apikey, dara.String("Apikey"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApikeyShrink) {
		query["Apikey"] = request.ApikeyShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMOQuotaAlertThreshold"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMOQuotaAlertThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the alert threshold percentage for an API key.
//
// Description:
//
// ### Applicable engine
//
// [RDS AI Assistant Ultimate Edition](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-mysql/rds-copilot-ultra)
//
// @param request - UpdateMOQuotaAlertThresholdRequest
//
// @return UpdateMOQuotaAlertThresholdResponse
func (client *Client) UpdateMOQuotaAlertThreshold(request *UpdateMOQuotaAlertThresholdRequest) (_result *UpdateMOQuotaAlertThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMOQuotaAlertThresholdResponse{}
	_body, _err := client.UpdateMOQuotaAlertThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of a specified Skill.
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

	if !dara.IsNil(request.UpdateNote) {
		query["UpdateNote"] = request.UpdateNote
	}

	if !dara.IsNil(request.UploadId) {
		query["UploadId"] = request.UploadId
	}

	if !dara.IsNil(request.UploadToken) {
		query["UploadToken"] = request.UploadToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
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
// Updates the information of a specified Skill.
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

// Summary:
//
// Validates datasource credentials and authorization.
//
// Description:
//
// Validates the connection credentials and authorization.
//
// @param request - VerifyContextDatabaseSyncConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyContextDatabaseSyncConnectionResponse
func (client *Client) VerifyContextDatabaseSyncConnectionWithOptions(request *VerifyContextDatabaseSyncConnectionRequest, runtime *dara.RuntimeOptions) (_result *VerifyContextDatabaseSyncConnectionResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyContextDatabaseSyncConnection"),
		Version:     dara.String("2025-05-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyContextDatabaseSyncConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates datasource credentials and authorization.
//
// Description:
//
// Validates the connection credentials and authorization.
//
// @param request - VerifyContextDatabaseSyncConnectionRequest
//
// @return VerifyContextDatabaseSyncConnectionResponse
func (client *Client) VerifyContextDatabaseSyncConnection(request *VerifyContextDatabaseSyncConnectionRequest) (_result *VerifyContextDatabaseSyncConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyContextDatabaseSyncConnectionResponse{}
	_body, _err := client.VerifyContextDatabaseSyncConnectionWithOptions(request, runtime)
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
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

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

	if !dara.IsNil(request.FilesShrink) {
		query["Files"] = request.FilesShrink
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
