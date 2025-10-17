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
	client.Endpoint, _err = client.GetEndpoint(dara.String("chatbot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 申请流式网关AccessToken
//
// @param request - ApplyForStreamAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyForStreamAccessTokenResponse
func (client *Client) ApplyForStreamAccessTokenWithOptions(request *ApplyForStreamAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *ApplyForStreamAccessTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyForStreamAccessToken"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyForStreamAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请流式网关AccessToken
//
// @param request - ApplyForStreamAccessTokenRequest
//
// @return ApplyForStreamAccessTokenResponse
func (client *Client) ApplyForStreamAccessToken(request *ApplyForStreamAccessTokenRequest) (_result *ApplyForStreamAccessTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyForStreamAccessTokenResponse{}
	_body, _err := client.ApplyForStreamAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话-联想API
//
// @param tmpReq - AssociateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateResponse
func (client *Client) AssociateWithOptions(tmpReq *AssociateRequest, runtime *dara.RuntimeOptions) (_result *AssociateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AssociateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Perspective) {
		request.PerspectiveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Perspective, dara.String("Perspective"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PerspectiveShrink) {
		query["Perspective"] = request.PerspectiveShrink
	}

	if !dara.IsNil(request.RecommendNum) {
		query["RecommendNum"] = request.RecommendNum
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Associate"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话-联想API
//
// @param request - AssociateRequest
//
// @return AssociateResponse
func (client *Client) Associate(request *AssociateRequest) (_result *AssociateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateResponse{}
	_body, _err := client.AssociateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取欢迎语
//
// @param request - BeginSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginSessionResponse
func (client *Client) BeginSessionWithOptions(request *BeginSessionRequest, runtime *dara.RuntimeOptions) (_result *BeginSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SandBox) {
		body["SandBox"] = request.SandBox
	}

	if !dara.IsNil(request.VendorParam) {
		body["VendorParam"] = request.VendorParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BeginSession"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BeginSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取欢迎语
//
// @param request - BeginSessionRequest
//
// @return BeginSessionResponse
func (client *Client) BeginSession(request *BeginSessionRequest) (_result *BeginSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BeginSessionResponse{}
	_body, _err := client.BeginSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消机器人发布
//
// @param request - CancelInstancePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelInstancePublishTaskResponse
func (client *Client) CancelInstancePublishTaskWithOptions(request *CancelInstancePublishTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelInstancePublishTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelInstancePublishTask"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelInstancePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消机器人发布
//
// @param request - CancelInstancePublishTaskRequest
//
// @return CancelInstancePublishTaskResponse
func (client *Client) CancelInstancePublishTask(request *CancelInstancePublishTaskRequest) (_result *CancelInstancePublishTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelInstancePublishTaskResponse{}
	_body, _err := client.CancelInstancePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消发布任务
//
// @param request - CancelPublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPublishTaskResponse
func (client *Client) CancelPublishTaskWithOptions(request *CancelPublishTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelPublishTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPublishTask"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消发布任务
//
// @param request - CancelPublishTaskRequest
//
// @return CancelPublishTaskResponse
func (client *Client) CancelPublishTask(request *CancelPublishTaskRequest) (_result *CancelPublishTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelPublishTaskResponse{}
	_body, _err := client.CancelPublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话API
//
// @param tmpReq - ChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithOptions(tmpReq *ChatRequest, runtime *dara.RuntimeOptions) (_result *ChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Perspective) {
		request.PerspectiveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Perspective, dara.String("Perspective"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentName) {
		query["IntentName"] = request.IntentName
	}

	if !dara.IsNil(request.KnowledgeId) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.PerspectiveShrink) {
		query["Perspective"] = request.PerspectiveShrink
	}

	if !dara.IsNil(request.SandBox) {
		query["SandBox"] = request.SandBox
	}

	if !dara.IsNil(request.SenderId) {
		query["SenderId"] = request.SenderId
	}

	if !dara.IsNil(request.SenderNick) {
		query["SenderNick"] = request.SenderNick
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	if !dara.IsNil(request.VendorParam) {
		query["VendorParam"] = request.VendorParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Chat"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话API
//
// @param request - ChatRequest
//
// @return ChatResponse
func (client *Client) Chat(request *ChatRequest) (_result *ChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatResponse{}
	_body, _err := client.ChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 继续机器人发布
//
// @param request - ContinueInstancePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueInstancePublishTaskResponse
func (client *Client) ContinueInstancePublishTaskWithOptions(request *ContinueInstancePublishTaskRequest, runtime *dara.RuntimeOptions) (_result *ContinueInstancePublishTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinueInstancePublishTask"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinueInstancePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 继续机器人发布
//
// @param request - ContinueInstancePublishTaskRequest
//
// @return ContinueInstancePublishTaskResponse
func (client *Client) ContinueInstancePublishTask(request *ContinueInstancePublishTaskRequest) (_result *ContinueInstancePublishTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinueInstancePublishTaskResponse{}
	_body, _err := client.ContinueInstancePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增类目
//
// @param request - CreateCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCategoryResponse
func (client *Client) CreateCategoryWithOptions(request *CreateCategoryRequest, runtime *dara.RuntimeOptions) (_result *CreateCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		body["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.KnowledgeType) {
		body["KnowledgeType"] = request.KnowledgeType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentCategoryId) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCategory"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增类目
//
// @param request - CreateCategoryRequest
//
// @return CreateCategoryResponse
func (client *Client) CreateCategory(request *CreateCategoryRequest) (_result *CreateCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCategoryResponse{}
	_body, _err := client.CreateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建FAQ关联问
//
// @param request - CreateConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConnQuestionResponse
func (client *Client) CreateConnQuestionWithOptions(request *CreateConnQuestionRequest, runtime *dara.RuntimeOptions) (_result *CreateConnQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnQuestionId) {
		body["ConnQuestionId"] = request.ConnQuestionId
	}

	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConnQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ关联问
//
// @param request - CreateConnQuestionRequest
//
// @return CreateConnQuestionResponse
func (client *Client) CreateConnQuestion(request *CreateConnQuestionRequest) (_result *CreateConnQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConnQuestionResponse{}
	_body, _err := client.CreateConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-创建
//
// @param request - CreateDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDSEntityResponse
func (client *Client) CreateDSEntityWithOptions(request *CreateDSEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateDSEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EntityName) {
		query["EntityName"] = request.EntityName
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDSEntity"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-创建
//
// @param request - CreateDSEntityRequest
//
// @return CreateDSEntityResponse
func (client *Client) CreateDSEntity(request *CreateDSEntityRequest) (_result *CreateDSEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDSEntityResponse{}
	_body, _err := client.CreateDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-创建
//
// @param tmpReq - CreateDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDSEntityValueResponse
func (client *Client) CreateDSEntityValueWithOptions(tmpReq *CreateDSEntityValueRequest, runtime *dara.RuntimeOptions) (_result *CreateDSEntityValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDSEntityValueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Synonyms) {
		request.SynonymsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Synonyms, dara.String("Synonyms"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SynonymsShrink) {
		body["Synonyms"] = request.SynonymsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDSEntityValue"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-创建
//
// @param request - CreateDSEntityValueRequest
//
// @return CreateDSEntityValueResponse
func (client *Client) CreateDSEntityValue(request *CreateDSEntityValueRequest) (_result *CreateDSEntityValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDSEntityValueResponse{}
	_body, _err := client.CreateDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文档
//
// @param tmpReq - CreateDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocResponse
func (client *Client) CreateDocWithOptions(tmpReq *CreateDocRequest, runtime *dara.RuntimeOptions) (_result *CreateDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocMetadata) {
		request.DocMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocMetadata, dara.String("DocMetadata"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagIds) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, dara.String("TagIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DocMetadataShrink) {
		query["DocMetadata"] = request.DocMetadataShrink
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Meta) {
		query["Meta"] = request.Meta
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TagIdsShrink) {
		query["TagIds"] = request.TagIdsShrink
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDoc"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档
//
// @param request - CreateDocRequest
//
// @return CreateDocResponse
func (client *Client) CreateDoc(request *CreateDocRequest) (_result *CreateDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDocResponse{}
	_body, _err := client.CreateDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建FAQ
//
// @param request - CreateFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFaqResponse
func (client *Client) CreateFaqWithOptions(request *CreateFaqRequest, runtime *dara.RuntimeOptions) (_result *CreateFaqResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.SolutionContent) {
		body["SolutionContent"] = request.SolutionContent
	}

	if !dara.IsNil(request.SolutionType) {
		body["SolutionType"] = request.SolutionType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFaq"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ
//
// @param request - CreateFaqRequest
//
// @return CreateFaqResponse
func (client *Client) CreateFaq(request *CreateFaqRequest) (_result *CreateFaqResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFaqResponse{}
	_body, _err := client.CreateFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-创建
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Introduction) {
		query["Introduction"] = request.Introduction
	}

	if !dara.IsNil(request.LanguageCode) {
		query["LanguageCode"] = request.LanguageCode
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RobotType) {
		query["RobotType"] = request.RobotType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-创建
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建机器人发布任务
//
// @param request - CreateInstancePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstancePublishTaskResponse
func (client *Client) CreateInstancePublishTaskWithOptions(request *CreateInstancePublishTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateInstancePublishTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstancePublishTask"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstancePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建机器人发布任务
//
// @param request - CreateInstancePublishTaskRequest
//
// @return CreateInstancePublishTaskResponse
func (client *Client) CreateInstancePublishTask(request *CreateInstancePublishTaskRequest) (_result *CreateInstancePublishTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstancePublishTaskResponse{}
	_body, _err := client.CreateInstancePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-创建
//
// @param tmpReq - CreateIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntentResponse
func (client *Client) CreateIntentWithOptions(tmpReq *CreateIntentRequest, runtime *dara.RuntimeOptions) (_result *CreateIntentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntentDefinition) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, dara.String("IntentDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentDefinitionShrink) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIntent"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-创建
//
// @param request - CreateIntentRequest
//
// @return CreateIntentResponse
func (client *Client) CreateIntent(request *CreateIntentRequest) (_result *CreateIntentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIntentResponse{}
	_body, _err := client.CreateIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-LGF-创建
//
// @param tmpReq - CreateLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLgfResponse
func (client *Client) CreateLgfWithOptions(tmpReq *CreateLgfRequest, runtime *dara.RuntimeOptions) (_result *CreateLgfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLgfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LgfDefinition) {
		request.LgfDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LgfDefinition, dara.String("LgfDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LgfDefinitionShrink) {
		query["LgfDefinition"] = request.LgfDefinitionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLgf"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-创建
//
// @param request - CreateLgfRequest
//
// @return CreateLgfResponse
func (client *Client) CreateLgf(request *CreateLgfRequest) (_result *CreateLgfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLgfResponse{}
	_body, _err := client.CreateLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-创建
//
// @param request - CreatePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePerspectiveResponse
func (client *Client) CreatePerspectiveWithOptions(request *CreatePerspectiveRequest, runtime *dara.RuntimeOptions) (_result *CreatePerspectiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
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
		Action:      dara.String("CreatePerspective"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-创建
//
// @param request - CreatePerspectiveRequest
//
// @return CreatePerspectiveResponse
func (client *Client) CreatePerspective(request *CreatePerspectiveRequest) (_result *CreatePerspectiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePerspectiveResponse{}
	_body, _err := client.CreatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建发布任务
//
// @param tmpReq - CreatePublishTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublishTaskResponse
func (client *Client) CreatePublishTaskWithOptions(tmpReq *CreatePublishTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatePublishTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePublishTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataIdList) {
		request.DataIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataIdList, dara.String("DataIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DataIdListShrink) {
		query["DataIdList"] = request.DataIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePublishTask"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePublishTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建发布任务
//
// @param request - CreatePublishTaskRequest
//
// @return CreatePublishTaskResponse
func (client *Client) CreatePublishTask(request *CreatePublishTaskRequest) (_result *CreatePublishTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePublishTaskResponse{}
	_body, _err := client.CreatePublishTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建FAQ相似问
//
// @param request - CreateSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimQuestionResponse
func (client *Client) CreateSimQuestionWithOptions(request *CreateSimQuestionRequest, runtime *dara.RuntimeOptions) (_result *CreateSimQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSimQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ相似问
//
// @param request - CreateSimQuestionRequest
//
// @return CreateSimQuestionResponse
func (client *Client) CreateSimQuestion(request *CreateSimQuestionRequest) (_result *CreateSimQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSimQuestionResponse{}
	_body, _err := client.CreateSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建FAQ答案
//
// @param request - CreateSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSolutionResponse
func (client *Client) CreateSolutionWithOptions(request *CreateSolutionRequest, runtime *dara.RuntimeOptions) (_result *CreateSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.KnowledgeId) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.PerspectiveCodes) {
		query["PerspectiveCodes"] = request.PerspectiveCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSolution"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建FAQ答案
//
// @param request - CreateSolutionRequest
//
// @return CreateSolutionResponse
func (client *Client) CreateSolution(request *CreateSolutionRequest) (_result *CreateSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSolutionResponse{}
	_body, _err := client.CreateSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-话术-创建
//
// @param tmpReq - CreateUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserSayResponse
func (client *Client) CreateUserSayWithOptions(tmpReq *CreateUserSayRequest, runtime *dara.RuntimeOptions) (_result *CreateUserSayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserSayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserSayDefinition) {
		request.UserSayDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserSayDefinition, dara.String("UserSayDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserSayDefinitionShrink) {
		query["UserSayDefinition"] = request.UserSayDefinitionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserSay"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-话术-创建
//
// @param request - CreateUserSayRequest
//
// @return CreateUserSayResponse
func (client *Client) CreateUserSay(request *CreateUserSayRequest) (_result *CreateUserSayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserSayResponse{}
	_body, _err := client.CreateUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除类目
//
// @param request - DeleteCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithOptions(request *DeleteCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCategory"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除类目
//
// @param request - DeleteCategoryRequest
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategory(request *DeleteCategoryRequest) (_result *DeleteCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除FAQ关联问
//
// @param request - DeleteConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnQuestionResponse
func (client *Client) DeleteConnQuestionWithOptions(request *DeleteConnQuestionRequest, runtime *dara.RuntimeOptions) (_result *DeleteConnQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OutlineId) {
		body["OutlineId"] = request.OutlineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ关联问
//
// @param request - DeleteConnQuestionRequest
//
// @return DeleteConnQuestionResponse
func (client *Client) DeleteConnQuestion(request *DeleteConnQuestionRequest) (_result *DeleteConnQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConnQuestionResponse{}
	_body, _err := client.DeleteConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-删除
//
// @param request - DeleteDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDSEntityResponse
func (client *Client) DeleteDSEntityWithOptions(request *DeleteDSEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteDSEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDSEntity"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-删除
//
// @param request - DeleteDSEntityRequest
//
// @return DeleteDSEntityResponse
func (client *Client) DeleteDSEntity(request *DeleteDSEntityRequest) (_result *DeleteDSEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDSEntityResponse{}
	_body, _err := client.DeleteDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-删除
//
// @param request - DeleteDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDSEntityValueResponse
func (client *Client) DeleteDSEntityValueWithOptions(request *DeleteDSEntityValueRequest, runtime *dara.RuntimeOptions) (_result *DeleteDSEntityValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityValueId) {
		query["EntityValueId"] = request.EntityValueId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDSEntityValue"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-删除
//
// @param request - DeleteDSEntityValueRequest
//
// @return DeleteDSEntityValueResponse
func (client *Client) DeleteDSEntityValue(request *DeleteDSEntityValueRequest) (_result *DeleteDSEntityValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDSEntityValueResponse{}
	_body, _err := client.DeleteDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档删除
//
// @param request - DeleteDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocResponse
func (client *Client) DeleteDocWithOptions(request *DeleteDocRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.KnowledgeId) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDoc"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档删除
//
// @param request - DeleteDocRequest
//
// @return DeleteDocResponse
func (client *Client) DeleteDoc(request *DeleteDocRequest) (_result *DeleteDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDocResponse{}
	_body, _err := client.DeleteDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除FAQ，如果是已发布的知识，删除之后，变成已删除未发布，需要发布才能真正删除
//
// @param request - DeleteFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaqResponse
func (client *Client) DeleteFaqWithOptions(request *DeleteFaqRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaqResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaq"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ，如果是已发布的知识，删除之后，变成已删除未发布，需要发布才能真正删除
//
// @param request - DeleteFaqRequest
//
// @return DeleteFaqResponse
func (client *Client) DeleteFaq(request *DeleteFaqRequest) (_result *DeleteFaqResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaqResponse{}
	_body, _err := client.DeleteFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-删除
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-删除
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-删除
//
// @param request - DeleteIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntentResponse
func (client *Client) DeleteIntentWithOptions(request *DeleteIntentRequest, runtime *dara.RuntimeOptions) (_result *DeleteIntentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIntent"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-删除
//
// @param request - DeleteIntentRequest
//
// @return DeleteIntentResponse
func (client *Client) DeleteIntent(request *DeleteIntentRequest) (_result *DeleteIntentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIntentResponse{}
	_body, _err := client.DeleteIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-LGF-删除
//
// @param request - DeleteLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLgfResponse
func (client *Client) DeleteLgfWithOptions(request *DeleteLgfRequest, runtime *dara.RuntimeOptions) (_result *DeleteLgfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.LgfId) {
		query["LgfId"] = request.LgfId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLgf"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-删除
//
// @param request - DeleteLgfRequest
//
// @return DeleteLgfResponse
func (client *Client) DeleteLgf(request *DeleteLgfRequest) (_result *DeleteLgfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLgfResponse{}
	_body, _err := client.DeleteLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-删除
//
// @param request - DeletePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePerspectiveResponse
func (client *Client) DeletePerspectiveWithOptions(request *DeletePerspectiveRequest, runtime *dara.RuntimeOptions) (_result *DeletePerspectiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.PerspectiveId) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePerspective"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-删除
//
// @param request - DeletePerspectiveRequest
//
// @return DeletePerspectiveResponse
func (client *Client) DeletePerspective(request *DeletePerspectiveRequest) (_result *DeletePerspectiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePerspectiveResponse{}
	_body, _err := client.DeletePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除FAQ相似问
//
// @param request - DeleteSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSimQuestionResponse
func (client *Client) DeleteSimQuestionWithOptions(request *DeleteSimQuestionRequest, runtime *dara.RuntimeOptions) (_result *DeleteSimQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SimQuestionId) {
		body["SimQuestionId"] = request.SimQuestionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSimQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ相似问
//
// @param request - DeleteSimQuestionRequest
//
// @return DeleteSimQuestionResponse
func (client *Client) DeleteSimQuestion(request *DeleteSimQuestionRequest) (_result *DeleteSimQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSimQuestionResponse{}
	_body, _err := client.DeleteSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除FAQ答案
//
// @param request - DeleteSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSolutionResponse
func (client *Client) DeleteSolutionWithOptions(request *DeleteSolutionRequest, runtime *dara.RuntimeOptions) (_result *DeleteSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SolutionId) {
		body["SolutionId"] = request.SolutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSolution"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除FAQ答案
//
// @param request - DeleteSolutionRequest
//
// @return DeleteSolutionResponse
func (client *Client) DeleteSolution(request *DeleteSolutionRequest) (_result *DeleteSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSolutionResponse{}
	_body, _err := client.DeleteSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-用户话术-删除
//
// @param request - DeleteUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserSayResponse
func (client *Client) DeleteUserSayWithOptions(request *DeleteUserSayRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserSayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.UserSayId) {
		query["UserSayId"] = request.UserSayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserSay"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-用户话术-删除
//
// @param request - DeleteUserSayRequest
//
// @return DeleteUserSayResponse
func (client *Client) DeleteUserSay(request *DeleteUserSayRequest) (_result *DeleteUserSayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserSayResponse{}
	_body, _err := client.DeleteUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看单个类目信息
//
// @param request - DescribeCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCategoryResponse
func (client *Client) DescribeCategoryWithOptions(request *DescribeCategoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCategory"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看单个类目信息
//
// @param request - DescribeCategoryRequest
//
// @return DescribeCategoryResponse
func (client *Client) DescribeCategory(request *DescribeCategoryRequest) (_result *DescribeCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCategoryResponse{}
	_body, _err := client.DescribeCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-详情
//
// @param request - DescribeDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDSEntityResponse
func (client *Client) DescribeDSEntityWithOptions(request *DescribeDSEntityRequest, runtime *dara.RuntimeOptions) (_result *DescribeDSEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDSEntity"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-详情
//
// @param request - DescribeDSEntityRequest
//
// @return DescribeDSEntityResponse
func (client *Client) DescribeDSEntity(request *DescribeDSEntityRequest) (_result *DescribeDSEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDSEntityResponse{}
	_body, _err := client.DescribeDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档详情
//
// @param request - DescribeDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocResponse
func (client *Client) DescribeDocWithOptions(request *DescribeDocRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.KnowledgeId) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.ShowDetail) {
		query["ShowDetail"] = request.ShowDetail
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDoc"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档详情
//
// @param request - DescribeDocRequest
//
// @return DescribeDocResponse
func (client *Client) DescribeDoc(request *DescribeDocRequest) (_result *DescribeDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDocResponse{}
	_body, _err := client.DescribeDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 知识详情
//
// @param request - DescribeFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaqResponse
func (client *Client) DescribeFaqWithOptions(request *DescribeFaqRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaqResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaq"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 知识详情
//
// @param request - DescribeFaqRequest
//
// @return DescribeFaqResponse
func (client *Client) DescribeFaq(request *DescribeFaqRequest) (_result *DescribeFaqResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFaqResponse{}
	_body, _err := client.DescribeFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-详情
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-详情
//
// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-详情
//
// @param request - DescribeIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIntentResponse
func (client *Client) DescribeIntentWithOptions(request *DescribeIntentRequest, runtime *dara.RuntimeOptions) (_result *DescribeIntentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IntentId) {
		body["IntentId"] = request.IntentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIntent"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-详情
//
// @param request - DescribeIntentRequest
//
// @return DescribeIntentResponse
func (client *Client) DescribeIntent(request *DescribeIntentRequest) (_result *DescribeIntentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIntentResponse{}
	_body, _err := client.DescribeIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-详情
//
// @param request - DescribePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePerspectiveResponse
func (client *Client) DescribePerspectiveWithOptions(request *DescribePerspectiveRequest, runtime *dara.RuntimeOptions) (_result *DescribePerspectiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.PerspectiveId) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePerspective"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-详情
//
// @param request - DescribePerspectiveRequest
//
// @return DescribePerspectiveResponse
func (client *Client) DescribePerspective(request *DescribePerspectiveRequest) (_result *DescribePerspectiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePerspectiveResponse{}
	_body, _err := client.DescribePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 问答点赞、点踩API
//
// @param request - FeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FeedbackResponse
func (client *Client) FeedbackWithOptions(request *FeedbackRequest, runtime *dara.RuntimeOptions) (_result *FeedbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Feedback) {
		query["Feedback"] = request.Feedback
	}

	if !dara.IsNil(request.FeedbackContent) {
		query["FeedbackContent"] = request.FeedbackContent
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Feedback"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 问答点赞、点踩API
//
// @param request - FeedbackRequest
//
// @return FeedbackResponse
func (client *Client) Feedback(request *FeedbackRequest) (_result *FeedbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FeedbackResponse{}
	_body, _err := client.FeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成用户免登Token
//
// @param request - GenerateUserAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUserAccessTokenResponse
func (client *Client) GenerateUserAccessTokenWithOptions(request *GenerateUserAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateUserAccessTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.ForeignId) {
		query["ForeignId"] = request.ForeignId
	}

	if !dara.IsNil(request.Nick) {
		query["Nick"] = request.Nick
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUserAccessToken"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUserAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成用户免登Token
//
// @param request - GenerateUserAccessTokenRequest
//
// @return GenerateUserAccessTokenResponse
func (client *Client) GenerateUserAccessToken(request *GenerateUserAccessTokenRequest) (_result *GenerateUserAccessTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateUserAccessTokenResponse{}
	_body, _err := client.GenerateUserAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间信息
//
// @param request - GetAgentInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentInfoResponse
func (client *Client) GetAgentInfoWithOptions(request *GetAgentInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAgentInfoResponse, _err error) {
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
		Action:      dara.String("GetAgentInfo"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间信息
//
// @param request - GetAgentInfoRequest
//
// @return GetAgentInfoResponse
func (client *Client) GetAgentInfo(request *GetAgentInfoRequest) (_result *GetAgentInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAgentInfoResponse{}
	_body, _err := client.GetAgentInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步函数执行结果接口
//
// @param request - GetAsyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncResultResponse
func (client *Client) GetAsyncResultWithOptions(request *GetAsyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncResult"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步函数执行结果接口
//
// @param request - GetAsyncResultRequest
//
// @return GetAsyncResultResponse
func (client *Client) GetAsyncResult(request *GetAsyncResultRequest) (_result *GetAsyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.GetAsyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询机器人接待人次和对话轮次
//
// @param request - GetBotSessionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBotSessionDataResponse
func (client *Client) GetBotSessionDataWithOptions(request *GetBotSessionDataRequest, runtime *dara.RuntimeOptions) (_result *GetBotSessionDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RobotInstanceId) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBotSessionData"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBotSessionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机器人接待人次和对话轮次
//
// @param request - GetBotSessionDataRequest
//
// @return GetBotSessionDataResponse
func (client *Client) GetBotSessionData(request *GetBotSessionDataRequest) (_result *GetBotSessionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBotSessionDataResponse{}
	_body, _err := client.GetBotSessionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询机器人发布进度
//
// @param request - GetInstancePublishTaskStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancePublishTaskStateResponse
func (client *Client) GetInstancePublishTaskStateWithOptions(request *GetInstancePublishTaskStateRequest, runtime *dara.RuntimeOptions) (_result *GetInstancePublishTaskStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstancePublishTaskState"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstancePublishTaskStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询机器人发布进度
//
// @param request - GetInstancePublishTaskStateRequest
//
// @return GetInstancePublishTaskStateResponse
func (client *Client) GetInstancePublishTaskState(request *GetInstancePublishTaskStateRequest) (_result *GetInstancePublishTaskStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstancePublishTaskStateResponse{}
	_body, _err := client.GetInstancePublishTaskStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询发布进度
//
// @param request - GetPublishTaskStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublishTaskStateResponse
func (client *Client) GetPublishTaskStateWithOptions(request *GetPublishTaskStateRequest, runtime *dara.RuntimeOptions) (_result *GetPublishTaskStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPublishTaskState"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPublishTaskStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询发布进度
//
// @param request - GetPublishTaskStateRequest
//
// @return GetPublishTaskStateResponse
func (client *Client) GetPublishTaskState(request *GetPublishTaskStateRequest) (_result *GetPublishTaskStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPublishTaskStateResponse{}
	_body, _err := client.GetPublishTaskStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 初始化im连接信息
//
// @param request - InitIMConnectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitIMConnectResponse
func (client *Client) InitIMConnectWithOptions(request *InitIMConnectRequest, runtime *dara.RuntimeOptions) (_result *InitIMConnectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.UserAccessToken) {
		query["UserAccessToken"] = request.UserAccessToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitIMConnect"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitIMConnectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化im连接信息
//
// @param request - InitIMConnectRequest
//
// @return InitIMConnectResponse
func (client *Client) InitIMConnect(request *InitIMConnectRequest) (_result *InitIMConnectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitIMConnectResponse{}
	_body, _err := client.InitIMConnectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-绑定类目
//
// @param request - LinkInstanceCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LinkInstanceCategoryResponse
func (client *Client) LinkInstanceCategoryWithOptions(request *LinkInstanceCategoryRequest, runtime *dara.RuntimeOptions) (_result *LinkInstanceCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbilityType) {
		query["AbilityType"] = request.AbilityType
	}

	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryIds) {
		body["CategoryIds"] = request.CategoryIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LinkInstanceCategory"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LinkInstanceCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-绑定类目
//
// @param request - LinkInstanceCategoryRequest
//
// @return LinkInstanceCategoryResponse
func (client *Client) LinkInstanceCategory(request *LinkInstanceCategoryRequest) (_result *LinkInstanceCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LinkInstanceCategoryResponse{}
	_body, _err := client.LinkInstanceCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间列表
//
// @param request - ListAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentResponse
func (client *Client) ListAgentWithOptions(request *ListAgentRequest, runtime *dara.RuntimeOptions) (_result *ListAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.GoodsCodes) {
		query["GoodsCodes"] = request.GoodsCodes
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
		Action:      dara.String("ListAgent"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间列表
//
// @param request - ListAgentRequest
//
// @return ListAgentResponse
func (client *Client) ListAgent(request *ListAgentRequest) (_result *ListAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAgentResponse{}
	_body, _err := client.ListAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 类目列表
//
// @param request - ListCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoryResponse
func (client *Client) ListCategoryWithOptions(request *ListCategoryRequest, runtime *dara.RuntimeOptions) (_result *ListCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeType) {
		body["KnowledgeType"] = request.KnowledgeType
	}

	if !dara.IsNil(request.ParentCategoryId) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCategory"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 类目列表
//
// @param request - ListCategoryRequest
//
// @return ListCategoryResponse
func (client *Client) ListCategory(request *ListCategoryRequest) (_result *ListCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCategoryResponse{}
	_body, _err := client.ListCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询FAQ关联问列表
//
// @param request - ListConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnQuestionResponse
func (client *Client) ListConnQuestionWithOptions(request *ListConnQuestionRequest, runtime *dara.RuntimeOptions) (_result *ListConnQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询FAQ关联问列表
//
// @param request - ListConnQuestionRequest
//
// @return ListConnQuestionResponse
func (client *Client) ListConnQuestion(request *ListConnQuestionRequest) (_result *ListConnQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConnQuestionResponse{}
	_body, _err := client.ListConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-列表
//
// @param request - ListDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDSEntityResponse
func (client *Client) ListDSEntityWithOptions(request *ListDSEntityRequest, runtime *dara.RuntimeOptions) (_result *ListDSEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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
		Action:      dara.String("ListDSEntity"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-列表
//
// @param request - ListDSEntityRequest
//
// @return ListDSEntityResponse
func (client *Client) ListDSEntity(request *ListDSEntityRequest) (_result *ListDSEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDSEntityResponse{}
	_body, _err := client.ListDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-列表
//
// @param request - ListDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDSEntityValueResponse
func (client *Client) ListDSEntityValueWithOptions(request *ListDSEntityValueRequest, runtime *dara.RuntimeOptions) (_result *ListDSEntityValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityValueId) {
		body["EntityValueId"] = request.EntityValueId
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDSEntityValue"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-列表
//
// @param request - ListDSEntityValueRequest
//
// @return ListDSEntityValueResponse
func (client *Client) ListDSEntityValue(request *ListDSEntityValueRequest) (_result *ListDSEntityValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDSEntityValueResponse{}
	_body, _err := client.ListDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - ListInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
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

	if !dara.IsNil(request.RobotType) {
		query["RobotType"] = request.RobotType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstance"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - ListInstanceRequest
//
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-列表
//
// @param request - ListIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntentResponse
func (client *Client) ListIntentWithOptions(request *ListIntentRequest, runtime *dara.RuntimeOptions) (_result *ListIntentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentName) {
		query["IntentName"] = request.IntentName
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
		Action:      dara.String("ListIntent"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-列表
//
// @param request - ListIntentRequest
//
// @return ListIntentResponse
func (client *Client) ListIntent(request *ListIntentRequest) (_result *ListIntentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntentResponse{}
	_body, _err := client.ListIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-LGF-列表
//
// @param request - ListLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLgfResponse
func (client *Client) ListLgfWithOptions(request *ListLgfRequest, runtime *dara.RuntimeOptions) (_result *ListLgfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.LgfText) {
		query["LgfText"] = request.LgfText
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
		Action:      dara.String("ListLgf"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-列表
//
// @param request - ListLgfRequest
//
// @return ListLgfResponse
func (client *Client) ListLgf(request *ListLgfRequest) (_result *ListLgfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLgfResponse{}
	_body, _err := client.ListLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的SaaS信息列表
//
// @param request - ListSaasInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSaasInfoResponse
func (client *Client) ListSaasInfoWithOptions(request *ListSaasInfoRequest, runtime *dara.RuntimeOptions) (_result *ListSaasInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.SaasGroupCodes) {
		query["SaasGroupCodes"] = request.SaasGroupCodes
	}

	if !dara.IsNil(request.SaasName) {
		query["SaasName"] = request.SaasName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSaasInfo"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSaasInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的SaaS信息列表
//
// @param request - ListSaasInfoRequest
//
// @return ListSaasInfoResponse
func (client *Client) ListSaasInfo(request *ListSaasInfoRequest) (_result *ListSaasInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSaasInfoResponse{}
	_body, _err := client.ListSaasInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的权限组信息
//
// @param request - ListSaasPermissionGroupInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSaasPermissionGroupInfosResponse
func (client *Client) ListSaasPermissionGroupInfosWithOptions(request *ListSaasPermissionGroupInfosRequest, runtime *dara.RuntimeOptions) (_result *ListSaasPermissionGroupInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSaasPermissionGroupInfos"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSaasPermissionGroupInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务空间下可集成的权限组信息
//
// @param request - ListSaasPermissionGroupInfosRequest
//
// @return ListSaasPermissionGroupInfosResponse
func (client *Client) ListSaasPermissionGroupInfos(request *ListSaasPermissionGroupInfosRequest) (_result *ListSaasPermissionGroupInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSaasPermissionGroupInfosResponse{}
	_body, _err := client.ListSaasPermissionGroupInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # FAQ相似问列表
//
// @param request - ListSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSimQuestionResponse
func (client *Client) ListSimQuestionWithOptions(request *ListSimQuestionRequest, runtime *dara.RuntimeOptions) (_result *ListSimQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSimQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # FAQ相似问列表
//
// @param request - ListSimQuestionRequest
//
// @return ListSimQuestionResponse
func (client *Client) ListSimQuestion(request *ListSimQuestionRequest) (_result *ListSimQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSimQuestionResponse{}
	_body, _err := client.ListSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # FAQ答案列表
//
// @param request - ListSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSolutionResponse
func (client *Client) ListSolutionWithOptions(request *ListSolutionRequest, runtime *dara.RuntimeOptions) (_result *ListSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSolution"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # FAQ答案列表
//
// @param request - ListSolutionRequest
//
// @return ListSolutionResponse
func (client *Client) ListSolution(request *ListSolutionRequest) (_result *ListSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSolutionResponse{}
	_body, _err := client.ListSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tongyi对话明细查询接口
//
// @param request - ListTongyiChatHistorysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTongyiChatHistorysResponse
func (client *Client) ListTongyiChatHistorysWithOptions(request *ListTongyiChatHistorysRequest, runtime *dara.RuntimeOptions) (_result *ListTongyiChatHistorysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.RobotInstanceId) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTongyiChatHistorys"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTongyiChatHistorysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tongyi对话明细查询接口
//
// @param request - ListTongyiChatHistorysRequest
//
// @return ListTongyiChatHistorysResponse
func (client *Client) ListTongyiChatHistorys(request *ListTongyiChatHistorysRequest) (_result *ListTongyiChatHistorysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTongyiChatHistorysResponse{}
	_body, _err := client.ListTongyiChatHistorysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询通义晓蜜的单个会话对话记录
//
// @param request - ListTongyiConversationLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTongyiConversationLogsResponse
func (client *Client) ListTongyiConversationLogsWithOptions(request *ListTongyiConversationLogsRequest, runtime *dara.RuntimeOptions) (_result *ListTongyiConversationLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.RobotInstanceId) {
		query["RobotInstanceId"] = request.RobotInstanceId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTongyiConversationLogs"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTongyiConversationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询通义晓蜜的单个会话对话记录
//
// @param request - ListTongyiConversationLogsRequest
//
// @return ListTongyiConversationLogsResponse
func (client *Client) ListTongyiConversationLogs(request *ListTongyiConversationLogsRequest) (_result *ListTongyiConversationLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTongyiConversationLogsResponse{}
	_body, _err := client.ListTongyiConversationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 话术-列表
//
// @param request - ListUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserSayResponse
func (client *Client) ListUserSayWithOptions(request *ListUserSayRequest, runtime *dara.RuntimeOptions) (_result *ListUserSayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
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
		Action:      dara.String("ListUserSay"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 话术-列表
//
// @param request - ListUserSayRequest
//
// @return ListUserSayResponse
func (client *Client) ListUserSay(request *ListUserSayRequest) (_result *ListUserSayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserSayResponse{}
	_body, _err := client.ListUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 统一NLU接口
//
// @param request - NluRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NluResponse
func (client *Client) NluWithOptions(request *NluRequest, runtime *dara.RuntimeOptions) (_result *NluResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Nlu"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NluResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 统一NLU接口
//
// @param request - NluRequest
//
// @return NluResponse
func (client *Client) Nlu(request *NluRequest) (_result *NluResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &NluResponse{}
	_body, _err := client.NluWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-列表
//
// @param request - QueryPerspectivesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPerspectivesResponse
func (client *Client) QueryPerspectivesWithOptions(request *QueryPerspectivesRequest, runtime *dara.RuntimeOptions) (_result *QueryPerspectivesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPerspectives"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-列表
//
// @param request - QueryPerspectivesRequest
//
// @return QueryPerspectivesResponse
func (client *Client) QueryPerspectives(request *QueryPerspectivesRequest) (_result *QueryPerspectivesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPerspectivesResponse{}
	_body, _err := client.QueryPerspectivesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档重试
//
// @param request - RetryDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryDocResponse
func (client *Client) RetryDocWithOptions(request *RetryDocRequest, runtime *dara.RuntimeOptions) (_result *RetryDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.KnowledgeId) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryDoc"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档重试
//
// @param request - RetryDocRequest
//
// @return RetryDocResponse
func (client *Client) RetryDoc(request *RetryDocRequest) (_result *RetryDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryDocResponse{}
	_body, _err := client.RetryDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档搜索
//
// @param tmpReq - SearchDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDocResponse
func (client *Client) SearchDocWithOptions(tmpReq *SearchDocRequest, runtime *dara.RuntimeOptions) (_result *SearchDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryIds) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, dara.String("CategoryIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagIds) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, dara.String("TagIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.CategoryIdsShrink) {
		query["CategoryIds"] = request.CategoryIdsShrink
	}

	if !dara.IsNil(request.CreateTimeBegin) {
		query["CreateTimeBegin"] = request.CreateTimeBegin
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateUserName) {
		query["CreateUserName"] = request.CreateUserName
	}

	if !dara.IsNil(request.EndTimeBegin) {
		query["EndTimeBegin"] = request.EndTimeBegin
	}

	if !dara.IsNil(request.EndTimeEnd) {
		query["EndTimeEnd"] = request.EndTimeEnd
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ModifyTimeBegin) {
		query["ModifyTimeBegin"] = request.ModifyTimeBegin
	}

	if !dara.IsNil(request.ModifyTimeEnd) {
		query["ModifyTimeEnd"] = request.ModifyTimeEnd
	}

	if !dara.IsNil(request.ModifyUserName) {
		query["ModifyUserName"] = request.ModifyUserName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessStatus) {
		query["ProcessStatus"] = request.ProcessStatus
	}

	if !dara.IsNil(request.SearchScope) {
		query["SearchScope"] = request.SearchScope
	}

	if !dara.IsNil(request.StartTimeBegin) {
		query["StartTimeBegin"] = request.StartTimeBegin
	}

	if !dara.IsNil(request.StartTimeEnd) {
		query["StartTimeEnd"] = request.StartTimeEnd
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIdsShrink) {
		query["TagIds"] = request.TagIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchDoc"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档搜索
//
// @param request - SearchDocRequest
//
// @return SearchDocResponse
func (client *Client) SearchDoc(request *SearchDocRequest) (_result *SearchDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchDocResponse{}
	_body, _err := client.SearchDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 知识搜索
//
// @param tmpReq - SearchFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFaqResponse
func (client *Client) SearchFaqWithOptions(tmpReq *SearchFaqRequest, runtime *dara.RuntimeOptions) (_result *SearchFaqResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchFaqShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryIds) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, dara.String("CategoryIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryIdsShrink) {
		body["CategoryIds"] = request.CategoryIdsShrink
	}

	if !dara.IsNil(request.CreateTimeBegin) {
		body["CreateTimeBegin"] = request.CreateTimeBegin
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateUserName) {
		body["CreateUserName"] = request.CreateUserName
	}

	if !dara.IsNil(request.EndTimeBegin) {
		body["EndTimeBegin"] = request.EndTimeBegin
	}

	if !dara.IsNil(request.EndTimeEnd) {
		body["EndTimeEnd"] = request.EndTimeEnd
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ModifyTimeBegin) {
		body["ModifyTimeBegin"] = request.ModifyTimeBegin
	}

	if !dara.IsNil(request.ModifyTimeEnd) {
		body["ModifyTimeEnd"] = request.ModifyTimeEnd
	}

	if !dara.IsNil(request.ModifyUserName) {
		body["ModifyUserName"] = request.ModifyUserName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchScope) {
		body["SearchScope"] = request.SearchScope
	}

	if !dara.IsNil(request.StartTimeBegin) {
		body["StartTimeBegin"] = request.StartTimeBegin
	}

	if !dara.IsNil(request.StartTimeEnd) {
		body["StartTimeEnd"] = request.StartTimeEnd
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchFaq"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 知识搜索
//
// @param request - SearchFaqRequest
//
// @return SearchFaqResponse
func (client *Client) SearchFaq(request *SearchFaqRequest) (_result *SearchFaqResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchFaqResponse{}
	_body, _err := client.SearchFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 大模型问答调试信息
//
// @param request - TongyiChatDebugInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TongyiChatDebugInfoResponse
func (client *Client) TongyiChatDebugInfoWithOptions(request *TongyiChatDebugInfoRequest, runtime *dara.RuntimeOptions) (_result *TongyiChatDebugInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TongyiChatDebugInfo"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TongyiChatDebugInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大模型问答调试信息
//
// @param request - TongyiChatDebugInfoRequest
//
// @return TongyiChatDebugInfoResponse
func (client *Client) TongyiChatDebugInfo(request *TongyiChatDebugInfoRequest) (_result *TongyiChatDebugInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TongyiChatDebugInfoResponse{}
	_body, _err := client.TongyiChatDebugInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑类目
//
// @param request - UpdateCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCategoryResponse
func (client *Client) UpdateCategoryWithOptions(request *UpdateCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		body["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCategory"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑类目
//
// @param request - UpdateCategoryRequest
//
// @return UpdateCategoryResponse
func (client *Client) UpdateCategory(request *UpdateCategoryRequest) (_result *UpdateCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.UpdateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新FAQ关联问
//
// @param request - UpdateConnQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnQuestionResponse
func (client *Client) UpdateConnQuestionWithOptions(request *UpdateConnQuestionRequest, runtime *dara.RuntimeOptions) (_result *UpdateConnQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnQuestionId) {
		body["ConnQuestionId"] = request.ConnQuestionId
	}

	if !dara.IsNil(request.OutlineId) {
		body["OutlineId"] = request.OutlineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ关联问
//
// @param request - UpdateConnQuestionRequest
//
// @return UpdateConnQuestionResponse
func (client *Client) UpdateConnQuestion(request *UpdateConnQuestionRequest) (_result *UpdateConnQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConnQuestionResponse{}
	_body, _err := client.UpdateConnQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体-更新
//
// @param request - UpdateDSEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDSEntityResponse
func (client *Client) UpdateDSEntityWithOptions(request *UpdateDSEntityRequest, runtime *dara.RuntimeOptions) (_result *UpdateDSEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityName) {
		query["EntityName"] = request.EntityName
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDSEntity"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDSEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体-更新
//
// @param request - UpdateDSEntityRequest
//
// @return UpdateDSEntityResponse
func (client *Client) UpdateDSEntity(request *UpdateDSEntityRequest) (_result *UpdateDSEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDSEntityResponse{}
	_body, _err := client.UpdateDSEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实体成员-更新
//
// @param tmpReq - UpdateDSEntityValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDSEntityValueResponse
func (client *Client) UpdateDSEntityValueWithOptions(tmpReq *UpdateDSEntityValueRequest, runtime *dara.RuntimeOptions) (_result *UpdateDSEntityValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDSEntityValueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Synonyms) {
		request.SynonymsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Synonyms, dara.String("Synonyms"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityValueId) {
		query["EntityValueId"] = request.EntityValueId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SynonymsShrink) {
		body["Synonyms"] = request.SynonymsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDSEntityValue"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDSEntityValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实体成员-更新
//
// @param request - UpdateDSEntityValueRequest
//
// @return UpdateDSEntityValueResponse
func (client *Client) UpdateDSEntityValue(request *UpdateDSEntityValueRequest) (_result *UpdateDSEntityValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDSEntityValueResponse{}
	_body, _err := client.UpdateDSEntityValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档变更
//
// @param tmpReq - UpdateDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocResponse
func (client *Client) UpdateDocWithOptions(tmpReq *UpdateDocRequest, runtime *dara.RuntimeOptions) (_result *UpdateDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocMetadata) {
		request.DocMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocMetadata, dara.String("DocMetadata"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagIds) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, dara.String("TagIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DocMetadataShrink) {
		query["DocMetadata"] = request.DocMetadataShrink
	}

	if !dara.IsNil(request.DocName) {
		query["DocName"] = request.DocName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.KnowledgeId) {
		query["KnowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.Meta) {
		query["Meta"] = request.Meta
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TagIdsShrink) {
		query["TagIds"] = request.TagIdsShrink
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDoc"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档变更
//
// @param request - UpdateDocRequest
//
// @return UpdateDocResponse
func (client *Client) UpdateDoc(request *UpdateDocRequest) (_result *UpdateDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDocResponse{}
	_body, _err := client.UpdateDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新FAQ
//
// @param request - UpdateFaqRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFaqResponse
func (client *Client) UpdateFaqWithOptions(request *UpdateFaqRequest, runtime *dara.RuntimeOptions) (_result *UpdateFaqResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.KnowledgeId) {
		body["KnowledgeId"] = request.KnowledgeId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFaq"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFaqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ
//
// @param request - UpdateFaqRequest
//
// @return UpdateFaqResponse
func (client *Client) UpdateFaq(request *UpdateFaqRequest) (_result *UpdateFaqResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFaqResponse{}
	_body, _err := client.UpdateFaqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Introduction) {
		query["Introduction"] = request.Introduction
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人-修改
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-更新
//
// @param tmpReq - UpdateIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntentResponse
func (client *Client) UpdateIntentWithOptions(tmpReq *UpdateIntentRequest, runtime *dara.RuntimeOptions) (_result *UpdateIntentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntentDefinition) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, dara.String("IntentDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentDefinitionShrink) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIntent"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIntentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-更新
//
// @param request - UpdateIntentRequest
//
// @return UpdateIntentResponse
func (client *Client) UpdateIntent(request *UpdateIntentRequest) (_result *UpdateIntentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIntentResponse{}
	_body, _err := client.UpdateIntentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-LGF-更新
//
// @param tmpReq - UpdateLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLgfResponse
func (client *Client) UpdateLgfWithOptions(tmpReq *UpdateLgfRequest, runtime *dara.RuntimeOptions) (_result *UpdateLgfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLgfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LgfDefinition) {
		request.LgfDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LgfDefinition, dara.String("LgfDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LgfDefinitionShrink) {
		query["LgfDefinition"] = request.LgfDefinitionShrink
	}

	if !dara.IsNil(request.LgfId) {
		query["LgfId"] = request.LgfId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLgf"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLgfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-LGF-更新
//
// @param request - UpdateLgfRequest
//
// @return UpdateLgfResponse
func (client *Client) UpdateLgf(request *UpdateLgfRequest) (_result *UpdateLgfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLgfResponse{}
	_body, _err := client.UpdateLgfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角-修改
//
// @param request - UpdatePerspectiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePerspectiveResponse
func (client *Client) UpdatePerspectiveWithOptions(request *UpdatePerspectiveRequest, runtime *dara.RuntimeOptions) (_result *UpdatePerspectiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PerspectiveId) {
		query["PerspectiveId"] = request.PerspectiveId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePerspective"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePerspectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角-修改
//
// @param request - UpdatePerspectiveRequest
//
// @return UpdatePerspectiveResponse
func (client *Client) UpdatePerspective(request *UpdatePerspectiveRequest) (_result *UpdatePerspectiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePerspectiveResponse{}
	_body, _err := client.UpdatePerspectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新FAQ相似问
//
// @param request - UpdateSimQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSimQuestionResponse
func (client *Client) UpdateSimQuestionWithOptions(request *UpdateSimQuestionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSimQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SimQuestionId) {
		body["SimQuestionId"] = request.SimQuestionId
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSimQuestion"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSimQuestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ相似问
//
// @param request - UpdateSimQuestionRequest
//
// @return UpdateSimQuestionResponse
func (client *Client) UpdateSimQuestion(request *UpdateSimQuestionRequest) (_result *UpdateSimQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSimQuestionResponse{}
	_body, _err := client.UpdateSimQuestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新FAQ答案
//
// @param request - UpdateSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSolutionResponse
func (client *Client) UpdateSolutionWithOptions(request *UpdateSolutionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.PerspectiveCodes) {
		body["PerspectiveCodes"] = request.PerspectiveCodes
	}

	if !dara.IsNil(request.SolutionId) {
		body["SolutionId"] = request.SolutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSolution"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新FAQ答案
//
// @param request - UpdateSolutionRequest
//
// @return UpdateSolutionResponse
func (client *Client) UpdateSolution(request *UpdateSolutionRequest) (_result *UpdateSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSolutionResponse{}
	_body, _err := client.UpdateSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图-话术-更新
//
// @param tmpReq - UpdateUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserSayResponse
func (client *Client) UpdateUserSayWithOptions(tmpReq *UpdateUserSayRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserSayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateUserSayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserSayDefinition) {
		request.UserSayDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserSayDefinition, dara.String("UserSayDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserSayDefinitionShrink) {
		query["UserSayDefinition"] = request.UserSayDefinitionShrink
	}

	if !dara.IsNil(request.UserSayId) {
		query["UserSayId"] = request.UserSayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserSay"),
		Version:     dara.String("2022-04-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserSayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图-话术-更新
//
// @param request - UpdateUserSayRequest
//
// @return UpdateUserSayResponse
func (client *Client) UpdateUserSay(request *UpdateUserSayRequest) (_result *UpdateUserSayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserSayResponse{}
	_body, _err := client.UpdateUserSayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
