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
	client.Endpoint, _err = client.GetEndpoint(dara.String("qualitycheck"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBusinessCategoryResponse
func (client *Client) AddBusinessCategoryWithOptions(request *AddBusinessCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddBusinessCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBusinessCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddBusinessCategoryRequest
//
// @return AddBusinessCategoryResponse
func (client *Client) AddBusinessCategory(request *AddBusinessCategoryRequest) (_result *AddBusinessCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.AddBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleCategoryResponse
func (client *Client) AddRuleCategoryWithOptions(request *AddRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRuleCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddRuleCategoryRequest
//
// @return AddRuleCategoryResponse
func (client *Client) AddRuleCategory(request *AddRuleCategoryRequest) (_result *AddRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.AddRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # V4创建规则
//
// @param request - AddRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleV4Response
func (client *Client) AddRuleV4WithOptions(request *AddRuleV4Request, runtime *dara.RuntimeOptions) (_result *AddRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IsCopy) {
		body["IsCopy"] = request.IsCopy
	}

	if !dara.IsNil(request.JsonStrForRule) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # V4创建规则
//
// @param request - AddRuleV4Request
//
// @return AddRuleV4Response
func (client *Client) AddRuleV4(request *AddRuleV4Request) (_result *AddRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRuleV4Response{}
	_body, _err := client.AddRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申领实时语音所需token
//
// @param request - ApplyWsTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyWsTokenResponse
func (client *Client) ApplyWsTokenWithOptions(request *ApplyWsTokenRequest, runtime *dara.RuntimeOptions) (_result *ApplyWsTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyWsToken"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyWsTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申领实时语音所需token
//
// @param request - ApplyWsTokenRequest
//
// @return ApplyWsTokenResponse
func (client *Client) ApplyWsToken(request *ApplyWsTokenRequest) (_result *ApplyWsTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyWsTokenResponse{}
	_body, _err := client.ApplyWsTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AssignReviewerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerResponse
func (client *Client) AssignReviewerWithOptions(request *AssignReviewerRequest, runtime *dara.RuntimeOptions) (_result *AssignReviewerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignReviewer"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignReviewerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssignReviewerRequest
//
// @return AssignReviewerResponse
func (client *Client) AssignReviewer(request *AssignReviewerRequest) (_result *AssignReviewerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssignReviewerResponse{}
	_body, _err := client.AssignReviewerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI AssignReviewerBySessionGroup is deprecated
//
// Summary:
//
// 会话组批量分配
//
// @param request - AssignReviewerBySessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerBySessionGroupResponse
func (client *Client) AssignReviewerBySessionGroupWithOptions(request *AssignReviewerBySessionGroupRequest, runtime *dara.RuntimeOptions) (_result *AssignReviewerBySessionGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignReviewerBySessionGroup"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignReviewerBySessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AssignReviewerBySessionGroup is deprecated
//
// Summary:
//
// 会话组批量分配
//
// @param request - AssignReviewerBySessionGroupRequest
//
// @return AssignReviewerBySessionGroupResponse
// Deprecated
func (client *Client) AssignReviewerBySessionGroup(request *AssignReviewerBySessionGroupRequest) (_result *AssignReviewerBySessionGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssignReviewerBySessionGroupResponse{}
	_body, _err := client.AssignReviewerBySessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量复核
//
// @param request - BatchSubmitReviewInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSubmitReviewInfoResponse
func (client *Client) BatchSubmitReviewInfoWithOptions(request *BatchSubmitReviewInfoRequest, runtime *dara.RuntimeOptions) (_result *BatchSubmitReviewInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSubmitReviewInfo"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSubmitReviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量复核
//
// @param request - BatchSubmitReviewInfoRequest
//
// @return BatchSubmitReviewInfoResponse
func (client *Client) BatchSubmitReviewInfo(request *BatchSubmitReviewInfoRequest) (_result *BatchSubmitReviewInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSubmitReviewInfoResponse{}
	_body, _err := client.BatchSubmitReviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建热词模型
//
// @param request - CreateAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsrVocabResponse
func (client *Client) CreateAsrVocabWithOptions(request *CreateAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *CreateAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热词模型
//
// @param request - CreateAsrVocabRequest
//
// @return CreateAsrVocabResponse
func (client *Client) CreateAsrVocab(request *CreateAsrVocabRequest) (_result *CreateAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CreateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建质检方案中的质检维度
//
// @param request - CreateCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCheckTypeToSchemeResponse
func (client *Client) CreateCheckTypeToSchemeWithOptions(request *CreateCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateCheckTypeToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCheckTypeToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建质检方案中的质检维度
//
// @param request - CreateCheckTypeToSchemeRequest
//
// @return CreateCheckTypeToSchemeResponse
func (client *Client) CreateCheckTypeToScheme(request *CreateCheckTypeToSchemeRequest) (_result *CreateCheckTypeToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCheckTypeToSchemeResponse{}
	_body, _err := client.CreateCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建标签挖掘任务
//
// @param request - CreateMiningTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMiningTaskResponse
func (client *Client) CreateMiningTaskWithOptions(request *CreateMiningTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMiningTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackUrl) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.FilePath) {
		body["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Param) {
		body["Param"] = request.Param
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMiningTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMiningTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标签挖掘任务
//
// @param request - CreateMiningTaskRequest
//
// @return CreateMiningTaskResponse
func (client *Client) CreateMiningTask(request *CreateMiningTaskRequest) (_result *CreateMiningTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMiningTaskResponse{}
	_body, _err := client.CreateMiningTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增质检方案
//
// @param request - CreateQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityCheckSchemeResponse
func (client *Client) CreateQualityCheckSchemeWithOptions(request *CreateQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增质检方案
//
// @param request - CreateQualityCheckSchemeRequest
//
// @return CreateQualityCheckSchemeResponse
func (client *Client) CreateQualityCheckScheme(request *CreateQualityCheckSchemeRequest) (_result *CreateQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQualityCheckSchemeResponse{}
	_body, _err := client.CreateQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建质检任务
//
// @param request - CreateSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemeTaskConfigResponse
func (client *Client) CreateSchemeTaskConfigWithOptions(request *CreateSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建质检任务
//
// @param request - CreateSchemeTaskConfigRequest
//
// @return CreateSchemeTaskConfigResponse
func (client *Client) CreateSchemeTaskConfig(request *CreateSchemeTaskConfigRequest) (_result *CreateSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSchemeTaskConfigResponse{}
	_body, _err := client.CreateSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateSkillGroupConfig is deprecated
//
// @param request - CreateSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillGroupConfigResponse
func (client *Client) CreateSkillGroupConfigWithOptions(request *CreateSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateSkillGroupConfig is deprecated
//
// @param request - CreateSkillGroupConfigRequest
//
// @return CreateSkillGroupConfigResponse
// Deprecated
func (client *Client) CreateSkillGroupConfig(request *CreateSkillGroupConfigRequest) (_result *CreateSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CreateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskAssignRuleResponse
func (client *Client) CreateTaskAssignRuleWithOptions(request *CreateTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateTaskAssignRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTaskAssignRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateTaskAssignRuleRequest
//
// @return CreateTaskAssignRuleResponse
func (client *Client) CreateTaskAssignRule(request *CreateTaskAssignRuleRequest) (_result *CreateTaskAssignRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CreateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用户
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningConfigResponse
func (client *Client) CreateWarningConfigWithOptions(request *CreateWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateWarningConfigRequest
//
// @return CreateWarningConfigResponse
func (client *Client) CreateWarningConfig(request *CreateWarningConfigRequest) (_result *CreateWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CreateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-新增
//
// @param request - CreateWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningStrategyConfigResponse
func (client *Client) CreateWarningStrategyConfigWithOptions(request *CreateWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-新增
//
// @param request - CreateWarningStrategyConfigRequest
//
// @return CreateWarningStrategyConfigResponse
func (client *Client) CreateWarningStrategyConfig(request *CreateWarningStrategyConfigRequest) (_result *CreateWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CreateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DelRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelRuleCategoryResponse
func (client *Client) DelRuleCategoryWithOptions(request *DelRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *DelRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelRuleCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DelRuleCategoryRequest
//
// @return DelRuleCategoryResponse
func (client *Client) DelRuleCategory(request *DelRuleCategoryRequest) (_result *DelRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.DelRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsrVocabResponse
func (client *Client) DeleteAsrVocabWithOptions(request *DeleteAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *DeleteAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAsrVocabRequest
//
// @return DeleteAsrVocabResponse
func (client *Client) DeleteAsrVocab(request *DeleteAsrVocabRequest) (_result *DeleteAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.DeleteAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessCategoryResponse
func (client *Client) DeleteBusinessCategoryWithOptions(request *DeleteBusinessCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteBusinessCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBusinessCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteBusinessCategoryRequest
//
// @return DeleteBusinessCategoryResponse
func (client *Client) DeleteBusinessCategory(request *DeleteBusinessCategoryRequest) (_result *DeleteBusinessCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.DeleteBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除质检唯独
//
// @param request - DeleteCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCheckTypeToSchemeResponse
func (client *Client) DeleteCheckTypeToSchemeWithOptions(request *DeleteCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteCheckTypeToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCheckTypeToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除质检唯独
//
// @param request - DeleteCheckTypeToSchemeRequest
//
// @return DeleteCheckTypeToSchemeResponse
func (client *Client) DeleteCheckTypeToScheme(request *DeleteCheckTypeToSchemeRequest) (_result *DeleteCheckTypeToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCheckTypeToSchemeResponse{}
	_body, _err := client.DeleteCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteCustomizationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizationConfigResponse
func (client *Client) DeleteCustomizationConfigWithOptions(request *DeleteCustomizationConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomizationConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomizationConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomizationConfigRequest
//
// @return DeleteCustomizationConfigResponse
func (client *Client) DeleteCustomizationConfig(request *DeleteCustomizationConfigRequest) (_result *DeleteCustomizationConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.DeleteCustomizationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteDataSet is deprecated
//
// @param request - DeleteDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSetResponse
func (client *Client) DeleteDataSetWithOptions(request *DeleteDataSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSet"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteDataSet is deprecated
//
// @param request - DeleteDataSetRequest
//
// @return DeleteDataSetResponse
// Deprecated
func (client *Client) DeleteDataSet(request *DeleteDataSetRequest) (_result *DeleteDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.DeleteDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeletePrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrecisionTaskResponse
func (client *Client) DeletePrecisionTaskWithOptions(request *DeletePrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *DeletePrecisionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrecisionTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeletePrecisionTaskRequest
//
// @return DeletePrecisionTaskResponse
func (client *Client) DeletePrecisionTask(request *DeletePrecisionTaskRequest) (_result *DeletePrecisionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrecisionTaskResponse{}
	_body, _err := client.DeletePrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除质检方案
//
// @param request - DeleteQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityCheckSchemeResponse
func (client *Client) DeleteQualityCheckSchemeWithOptions(request *DeleteQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除质检方案
//
// @param request - DeleteQualityCheckSchemeRequest
//
// @return DeleteQualityCheckSchemeResponse
func (client *Client) DeleteQualityCheckScheme(request *DeleteQualityCheckSchemeRequest) (_result *DeleteQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityCheckSchemeResponse{}
	_body, _err := client.DeleteQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// 删除规则
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelete) {
		body["ForceDelete"] = request.ForceDelete
	}

	if !dara.IsNil(request.IsSchemeData) {
		body["IsSchemeData"] = request.IsSchemeData
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// 删除规则
//
// @param request - DeleteRuleRequest
//
// @return DeleteRuleResponse
// Deprecated
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # V4删除规则
//
// @param request - DeleteRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleV4Response
func (client *Client) DeleteRuleV4WithOptions(request *DeleteRuleV4Request, runtime *dara.RuntimeOptions) (_result *DeleteRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelete) {
		body["ForceDelete"] = request.ForceDelete
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # V4删除规则
//
// @param request - DeleteRuleV4Request
//
// @return DeleteRuleV4Response
func (client *Client) DeleteRuleV4(request *DeleteRuleV4Request) (_result *DeleteRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRuleV4Response{}
	_body, _err := client.DeleteRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除质检任务
//
// @param request - DeleteSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchemeTaskConfigResponse
func (client *Client) DeleteSchemeTaskConfigWithOptions(request *DeleteSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除质检任务
//
// @param request - DeleteSchemeTaskConfigRequest
//
// @return DeleteSchemeTaskConfigResponse
func (client *Client) DeleteSchemeTaskConfig(request *DeleteSchemeTaskConfigRequest) (_result *DeleteSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSchemeTaskConfigResponse{}
	_body, _err := client.DeleteSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteSkillGroupConfig is deprecated
//
// @param request - DeleteSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillGroupConfigResponse
func (client *Client) DeleteSkillGroupConfigWithOptions(request *DeleteSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteSkillGroupConfig is deprecated
//
// @param request - DeleteSkillGroupConfigRequest
//
// @return DeleteSkillGroupConfigResponse
// Deprecated
func (client *Client) DeleteSkillGroupConfig(request *DeleteSkillGroupConfigRequest) (_result *DeleteSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.DeleteSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskAssignRuleResponse
func (client *Client) DeleteTaskAssignRuleWithOptions(request *DeleteTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskAssignRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTaskAssignRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTaskAssignRuleRequest
//
// @return DeleteTaskAssignRuleResponse
func (client *Client) DeleteTaskAssignRule(request *DeleteTaskAssignRuleRequest) (_result *DeleteTaskAssignRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.DeleteTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningConfigResponse
func (client *Client) DeleteWarningConfigWithOptions(request *DeleteWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteWarningConfigRequest
//
// @return DeleteWarningConfigResponse
func (client *Client) DeleteWarningConfig(request *DeleteWarningConfigRequest) (_result *DeleteWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.DeleteWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-删除
//
// @param request - DeleteWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningStrategyConfigResponse
func (client *Client) DeleteWarningStrategyConfigWithOptions(request *DeleteWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-删除
//
// @param request - DeleteWarningStrategyConfigRequest
//
// @return DeleteWarningStrategyConfigResponse
func (client *Client) DeleteWarningStrategyConfig(request *DeleteWarningStrategyConfigRequest) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.DeleteWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsrVocabResponse
func (client *Client) GetAsrVocabWithOptions(request *GetAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *GetAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAsrVocabRequest
//
// @return GetAsrVocabResponse
func (client *Client) GetAsrVocab(request *GetAsrVocabRequest) (_result *GetAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.GetAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetBusinessCategoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusinessCategoryListResponse
func (client *Client) GetBusinessCategoryListWithOptions(request *GetBusinessCategoryListRequest, runtime *dara.RuntimeOptions) (_result *GetBusinessCategoryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBusinessCategoryList"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetBusinessCategoryListRequest
//
// @return GetBusinessCategoryListResponse
func (client *Client) GetBusinessCategoryList(request *GetBusinessCategoryListRequest) (_result *GetBusinessCategoryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.GetBusinessCategoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取语音模型列表
//
// @param request - GetCustomizationConfigListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomizationConfigListResponse
func (client *Client) GetCustomizationConfigListWithOptions(request *GetCustomizationConfigListRequest, runtime *dara.RuntimeOptions) (_result *GetCustomizationConfigListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomizationConfigList"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取语音模型列表
//
// @param request - GetCustomizationConfigListRequest
//
// @return GetCustomizationConfigListResponse
func (client *Client) GetCustomizationConfigList(request *GetCustomizationConfigListRequest) (_result *GetCustomizationConfigListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.GetCustomizationConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标签挖掘任务结果
//
// @param request - GetMiningTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMiningTaskResultResponse
func (client *Client) GetMiningTaskResultWithOptions(request *GetMiningTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetMiningTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMiningTaskResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMiningTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标签挖掘任务结果
//
// @param request - GetMiningTaskResultRequest
//
// @return GetMiningTaskResultResponse
func (client *Client) GetMiningTaskResult(request *GetMiningTaskResultRequest) (_result *GetMiningTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMiningTaskResultResponse{}
	_body, _err := client.GetMiningTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNextResultToVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNextResultToVerifyResponse
func (client *Client) GetNextResultToVerifyWithOptions(request *GetNextResultToVerifyRequest, runtime *dara.RuntimeOptions) (_result *GetNextResultToVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNextResultToVerify"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNextResultToVerifyRequest
//
// @return GetNextResultToVerifyResponse
func (client *Client) GetNextResultToVerify(request *GetNextResultToVerifyRequest) (_result *GetNextResultToVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.GetNextResultToVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrecisionTaskResponse
func (client *Client) GetPrecisionTaskWithOptions(request *GetPrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *GetPrecisionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrecisionTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetPrecisionTaskRequest
//
// @return GetPrecisionTaskResponse
func (client *Client) GetPrecisionTask(request *GetPrecisionTaskRequest) (_result *GetPrecisionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPrecisionTaskResponse{}
	_body, _err := client.GetPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检方案
//
// @param request - GetQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityCheckSchemeResponse
func (client *Client) GetQualityCheckSchemeWithOptions(request *GetQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *GetQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检方案
//
// @param request - GetQualityCheckSchemeRequest
//
// @return GetQualityCheckSchemeResponse
func (client *Client) GetQualityCheckScheme(request *GetQualityCheckSchemeRequest) (_result *GetQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityCheckSchemeResponse{}
	_body, _err := client.GetQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检结果
//
// @param request - GetResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultResponse
func (client *Client) GetResultWithOptions(request *GetResultRequest, runtime *dara.RuntimeOptions) (_result *GetResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检结果
//
// @param request - GetResultRequest
//
// @return GetResultResponse
func (client *Client) GetResult(request *GetResultRequest) (_result *GetResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResultResponse{}
	_body, _err := client.GetResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检结果详情用于复核
//
// @param request - GetResultToReviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultToReviewResponse
func (client *Client) GetResultToReviewWithOptions(request *GetResultToReviewRequest, runtime *dara.RuntimeOptions) (_result *GetResultToReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResultToReview"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检结果详情用于复核
//
// @param request - GetResultToReviewRequest
//
// @return GetResultToReviewResponse
func (client *Client) GetResultToReview(request *GetResultToReviewRequest) (_result *GetResultToReviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.GetResultToReviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRule is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleResponse
func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRule is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleRequest
//
// @return GetRuleResponse
// Deprecated
func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRuleById is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// 获取规则
//
// @param request - GetRuleByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleByIdResponse
func (client *Client) GetRuleByIdWithOptions(request *GetRuleByIdRequest, runtime *dara.RuntimeOptions) (_result *GetRuleByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleById"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRuleById is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// 获取规则
//
// @param request - GetRuleByIdRequest
//
// @return GetRuleByIdResponse
// Deprecated
func (client *Client) GetRuleById(request *GetRuleByIdRequest) (_result *GetRuleByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleByIdResponse{}
	_body, _err := client.GetRuleByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleCategoryResponse
func (client *Client) GetRuleCategoryWithOptions(request *GetRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRuleCategoryRequest
//
// @return GetRuleCategoryResponse
func (client *Client) GetRuleCategory(request *GetRuleCategoryRequest) (_result *GetRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.GetRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRuleDetail is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleDetailResponse
func (client *Client) GetRuleDetailWithOptions(request *GetRuleDetailRequest, runtime *dara.RuntimeOptions) (_result *GetRuleDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleDetail"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRuleDetail is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// @param request - GetRuleDetailRequest
//
// @return GetRuleDetailResponse
// Deprecated
func (client *Client) GetRuleDetail(request *GetRuleDetailRequest) (_result *GetRuleDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.GetRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # V4获取规则
//
// @param request - GetRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleV4Response
func (client *Client) GetRuleV4WithOptions(request *GetRuleV4Request, runtime *dara.RuntimeOptions) (_result *GetRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # V4获取规则
//
// @param request - GetRuleV4Request
//
// @return GetRuleV4Response
func (client *Client) GetRuleV4(request *GetRuleV4Request) (_result *GetRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleV4Response{}
	_body, _err := client.GetRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得规则列表
//
// @param request - GetRulesCountListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRulesCountListResponse
func (client *Client) GetRulesCountListWithOptions(request *GetRulesCountListRequest, runtime *dara.RuntimeOptions) (_result *GetRulesCountListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.BusinessRange) {
		body["BusinessRange"] = request.BusinessRange
	}

	if !dara.IsNil(request.CategoryName) {
		body["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.CountTotal) {
		body["CountTotal"] = request.CountTotal
	}

	if !dara.IsNil(request.CreateEmpid) {
		body["CreateEmpid"] = request.CreateEmpid
	}

	if !dara.IsNil(request.CreateUserId) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LastUpdateEmpid) {
		body["LastUpdateEmpid"] = request.LastUpdateEmpid
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequireInfos) {
		body["RequireInfos"] = request.RequireInfos
	}

	if !dara.IsNil(request.Rid) {
		body["Rid"] = request.Rid
	}

	if !dara.IsNil(request.RuleIdOrRuleName) {
		body["RuleIdOrRuleName"] = request.RuleIdOrRuleName
	}

	if !dara.IsNil(request.RuleScoreSingleType) {
		body["RuleScoreSingleType"] = request.RuleScoreSingleType
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SchemeId) {
		body["SchemeId"] = request.SchemeId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.TypeName) {
		body["TypeName"] = request.TypeName
	}

	if !dara.IsNil(request.UpdateEndTime) {
		body["UpdateEndTime"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		body["UpdateStartTime"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UpdateUserId) {
		body["UpdateUserId"] = request.UpdateUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRulesCountList"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRulesCountListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得规则列表
//
// @param request - GetRulesCountListRequest
//
// @return GetRulesCountListResponse
func (client *Client) GetRulesCountList(request *GetRulesCountListRequest) (_result *GetRulesCountListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRulesCountListResponse{}
	_body, _err := client.GetRulesCountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检任务配置详情
//
// @param request - GetSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemeTaskConfigResponse
func (client *Client) GetSchemeTaskConfigWithOptions(request *GetSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检任务配置详情
//
// @param request - GetSchemeTaskConfigRequest
//
// @return GetSchemeTaskConfigResponse
func (client *Client) GetSchemeTaskConfig(request *GetSchemeTaskConfigRequest) (_result *GetSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSchemeTaskConfigResponse{}
	_body, _err := client.GetSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetScoreInfo is deprecated
//
// @param request - GetScoreInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScoreInfoResponse
func (client *Client) GetScoreInfoWithOptions(request *GetScoreInfoRequest, runtime *dara.RuntimeOptions) (_result *GetScoreInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScoreInfo"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScoreInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetScoreInfo is deprecated
//
// @param request - GetScoreInfoRequest
//
// @return GetScoreInfoResponse
// Deprecated
func (client *Client) GetScoreInfo(request *GetScoreInfoRequest) (_result *GetScoreInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScoreInfoResponse{}
	_body, _err := client.GetScoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetSkillGroupConfig is deprecated
//
// @param request - GetSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupConfigResponse
func (client *Client) GetSkillGroupConfigWithOptions(request *GetSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetSkillGroupConfig is deprecated
//
// @param request - GetSkillGroupConfigRequest
//
// @return GetSkillGroupConfigResponse
// Deprecated
func (client *Client) GetSkillGroupConfig(request *GetSkillGroupConfigRequest) (_result *GetSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.GetSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetSyncResult is deprecated, please use Qualitycheck::2019-01-15::GetResult instead.
//
// @param request - GetSyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyncResultResponse
func (client *Client) GetSyncResultWithOptions(request *GetSyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetSyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSyncResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetSyncResult is deprecated, please use Qualitycheck::2019-01-15::GetResult instead.
//
// @param request - GetSyncResultRequest
//
// @return GetSyncResultResponse
// Deprecated
func (client *Client) GetSyncResult(request *GetSyncResultRequest) (_result *GetSyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSyncResultResponse{}
	_body, _err := client.GetSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-详情
//
// @param request - GetWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarningStrategyConfigResponse
func (client *Client) GetWarningStrategyConfigWithOptions(request *GetWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *GetWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-详情
//
// @param request - GetWarningStrategyConfigRequest
//
// @return GetWarningStrategyConfigResponse
func (client *Client) GetWarningStrategyConfig(request *GetWarningStrategyConfigRequest) (_result *GetWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.GetWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - HandleComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleComplaintResponse
func (client *Client) HandleComplaintWithOptions(request *HandleComplaintRequest, runtime *dara.RuntimeOptions) (_result *HandleComplaintResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HandleComplaint"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HandleComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - HandleComplaintRequest
//
// @return HandleComplaintResponse
func (client *Client) HandleComplaint(request *HandleComplaintRequest) (_result *HandleComplaintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &HandleComplaintResponse{}
	_body, _err := client.HandleComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI InvalidRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// @param request - InvalidRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidRuleResponse
func (client *Client) InvalidRuleWithOptions(request *InvalidRuleRequest, runtime *dara.RuntimeOptions) (_result *InvalidRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvalidRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvalidRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI InvalidRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// @param request - InvalidRuleRequest
//
// @return InvalidRuleResponse
// Deprecated
func (client *Client) InvalidRule(request *InvalidRuleRequest) (_result *InvalidRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InvalidRuleResponse{}
	_body, _err := client.InvalidRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取热词模型列表
//
// @param request - ListAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsrVocabResponse
func (client *Client) ListAsrVocabWithOptions(request *ListAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *ListAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热词模型列表
//
// @param request - ListAsrVocabRequest
//
// @return ListAsrVocabResponse
func (client *Client) ListAsrVocab(request *ListAsrVocabRequest) (_result *ListAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.ListAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListDataSet is deprecated
//
// Summary:
//
// 获取数据集列表
//
// @param request - ListDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSetResponse
func (client *Client) ListDataSetWithOptions(request *ListDataSetRequest, runtime *dara.RuntimeOptions) (_result *ListDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSet"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDataSet is deprecated
//
// Summary:
//
// 获取数据集列表
//
// @param request - ListDataSetRequest
//
// @return ListDataSetResponse
// Deprecated
func (client *Client) ListDataSet(request *ListDataSetRequest) (_result *ListDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSetResponse{}
	_body, _err := client.ListDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrecisionTaskResponse
func (client *Client) ListPrecisionTaskWithOptions(request *ListPrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *ListPrecisionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrecisionTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPrecisionTaskRequest
//
// @return ListPrecisionTaskResponse
func (client *Client) ListPrecisionTask(request *ListPrecisionTaskRequest) (_result *ListPrecisionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrecisionTaskResponse{}
	_body, _err := client.ListPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 质检方案列表
//
// @param request - ListQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityCheckSchemeResponse
func (client *Client) ListQualityCheckSchemeWithOptions(request *ListQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *ListQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 质检方案列表
//
// @param request - ListQualityCheckSchemeRequest
//
// @return ListQualityCheckSchemeResponse
func (client *Client) ListQualityCheckScheme(request *ListQualityCheckSchemeRequest) (_result *ListQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.ListQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListRules is deprecated, please use Qualitycheck::2019-01-15::ListRulesV4 instead.
//
// @param request - ListRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRules"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListRules is deprecated, please use Qualitycheck::2019-01-15::ListRulesV4 instead.
//
// @param request - ListRulesRequest
//
// @return ListRulesResponse
// Deprecated
func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # V4获得规则列表
//
// @param request - ListRulesV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesV4Response
func (client *Client) ListRulesV4WithOptions(request *ListRulesV4Request, runtime *dara.RuntimeOptions) (_result *ListRulesV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.BusinessRange) {
		body["BusinessRange"] = request.BusinessRange
	}

	if !dara.IsNil(request.CategoryName) {
		body["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.CountTotal) {
		body["CountTotal"] = request.CountTotal
	}

	if !dara.IsNil(request.CreateEmpid) {
		body["CreateEmpid"] = request.CreateEmpid
	}

	if !dara.IsNil(request.CreateUserId) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LastUpdateEmpid) {
		body["LastUpdateEmpid"] = request.LastUpdateEmpid
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequireInfos) {
		body["RequireInfos"] = request.RequireInfos
	}

	if !dara.IsNil(request.Rid) {
		body["Rid"] = request.Rid
	}

	if !dara.IsNil(request.RuleIdOrRuleName) {
		body["RuleIdOrRuleName"] = request.RuleIdOrRuleName
	}

	if !dara.IsNil(request.RuleScoreSingleType) {
		body["RuleScoreSingleType"] = request.RuleScoreSingleType
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SchemeId) {
		body["SchemeId"] = request.SchemeId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.TypeName) {
		body["TypeName"] = request.TypeName
	}

	if !dara.IsNil(request.UpdateEndTime) {
		body["UpdateEndTime"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		body["UpdateStartTime"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UpdateUserId) {
		body["UpdateUserId"] = request.UpdateUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRulesV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # V4获得规则列表
//
// @param request - ListRulesV4Request
//
// @return ListRulesV4Response
func (client *Client) ListRulesV4(request *ListRulesV4Request) (_result *ListRulesV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRulesV4Response{}
	_body, _err := client.ListRulesV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检任务列表
//
// @param request - ListSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemeTaskConfigResponse
func (client *Client) ListSchemeTaskConfigWithOptions(request *ListSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *ListSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检任务列表
//
// @param request - ListSchemeTaskConfigRequest
//
// @return ListSchemeTaskConfigResponse
func (client *Client) ListSchemeTaskConfig(request *ListSchemeTaskConfigRequest) (_result *ListSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSchemeTaskConfigResponse{}
	_body, _err := client.ListSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListSessionGroup is deprecated
//
// Summary:
//
// 获取会话组列表
//
// @param request - ListSessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionGroupResponse
func (client *Client) ListSessionGroupWithOptions(request *ListSessionGroupRequest, runtime *dara.RuntimeOptions) (_result *ListSessionGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessionGroup"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListSessionGroup is deprecated
//
// Summary:
//
// 获取会话组列表
//
// @param request - ListSessionGroupRequest
//
// @return ListSessionGroupResponse
// Deprecated
func (client *Client) ListSessionGroup(request *ListSessionGroupRequest) (_result *ListSessionGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSessionGroupResponse{}
	_body, _err := client.ListSessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupConfig is deprecated
//
// @param request - ListSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupConfigResponse
func (client *Client) ListSkillGroupConfigWithOptions(request *ListSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *ListSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupConfig is deprecated
//
// @param request - ListSkillGroupConfigRequest
//
// @return ListSkillGroupConfigResponse
// Deprecated
func (client *Client) ListSkillGroupConfig(request *ListSkillGroupConfigRequest) (_result *ListSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.ListSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListTaskAssignRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskAssignRulesResponse
func (client *Client) ListTaskAssignRulesWithOptions(request *ListTaskAssignRulesRequest, runtime *dara.RuntimeOptions) (_result *ListTaskAssignRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskAssignRules"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTaskAssignRulesRequest
//
// @return ListTaskAssignRulesResponse
func (client *Client) ListTaskAssignRules(request *ListTaskAssignRulesRequest) (_result *ListTaskAssignRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.ListTaskAssignRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningConfigResponse
func (client *Client) ListWarningConfigWithOptions(request *ListWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *ListWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListWarningConfigRequest
//
// @return ListWarningConfigResponse
func (client *Client) ListWarningConfig(request *ListWarningConfigRequest) (_result *ListWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.ListWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-列表
//
// @param request - ListWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningStrategyConfigResponse
func (client *Client) ListWarningStrategyConfigWithOptions(request *ListWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *ListWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-列表
//
// @param request - ListWarningStrategyConfigRequest
//
// @return ListWarningStrategyConfigResponse
func (client *Client) ListWarningStrategyConfig(request *ListWarningStrategyConfigRequest) (_result *ListWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.ListWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量回收
//
// @param request - RevertAssignedSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionResponse
func (client *Client) RevertAssignedSessionWithOptions(request *RevertAssignedSessionRequest, runtime *dara.RuntimeOptions) (_result *RevertAssignedSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertAssignedSession"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertAssignedSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量回收
//
// @param request - RevertAssignedSessionRequest
//
// @return RevertAssignedSessionResponse
func (client *Client) RevertAssignedSession(request *RevertAssignedSessionRequest) (_result *RevertAssignedSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevertAssignedSessionResponse{}
	_body, _err := client.RevertAssignedSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI RevertAssignedSessionGroup is deprecated
//
// Summary:
//
// 会话组批量回收
//
// @param request - RevertAssignedSessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionGroupResponse
func (client *Client) RevertAssignedSessionGroupWithOptions(request *RevertAssignedSessionGroupRequest, runtime *dara.RuntimeOptions) (_result *RevertAssignedSessionGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertAssignedSessionGroup"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertAssignedSessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RevertAssignedSessionGroup is deprecated
//
// Summary:
//
// 会话组批量回收
//
// @param request - RevertAssignedSessionGroupRequest
//
// @return RevertAssignedSessionGroupResponse
// Deprecated
func (client *Client) RevertAssignedSessionGroup(request *RevertAssignedSessionGroupRequest) (_result *RevertAssignedSessionGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevertAssignedSessionGroupResponse{}
	_body, _err := client.RevertAssignedSessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI SaveConfigDataSet is deprecated
//
// @param request - SaveConfigDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveConfigDataSetResponse
func (client *Client) SaveConfigDataSetWithOptions(request *SaveConfigDataSetRequest, runtime *dara.RuntimeOptions) (_result *SaveConfigDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveConfigDataSet"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI SaveConfigDataSet is deprecated
//
// @param request - SaveConfigDataSetRequest
//
// @return SaveConfigDataSetResponse
// Deprecated
func (client *Client) SaveConfigDataSet(request *SaveConfigDataSetRequest) (_result *SaveConfigDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.SaveConfigDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitComplaintResponse
func (client *Client) SubmitComplaintWithOptions(request *SubmitComplaintRequest, runtime *dara.RuntimeOptions) (_result *SubmitComplaintResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitComplaint"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitComplaintRequest
//
// @return SubmitComplaintResponse
func (client *Client) SubmitComplaint(request *SubmitComplaintRequest) (_result *SubmitComplaintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.SubmitComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPrecisionTaskResponse
func (client *Client) SubmitPrecisionTaskWithOptions(request *SubmitPrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitPrecisionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitPrecisionTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitPrecisionTaskRequest
//
// @return SubmitPrecisionTaskResponse
func (client *Client) SubmitPrecisionTask(request *SubmitPrecisionTaskRequest) (_result *SubmitPrecisionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitPrecisionTaskResponse{}
	_body, _err := client.SubmitPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitQualityCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityCheckTaskResponse
func (client *Client) SubmitQualityCheckTaskWithOptions(request *SubmitQualityCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitQualityCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitQualityCheckTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitQualityCheckTaskRequest
//
// @return SubmitQualityCheckTaskResponse
func (client *Client) SubmitQualityCheckTask(request *SubmitQualityCheckTaskRequest) (_result *SubmitQualityCheckTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.SubmitQualityCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitReviewInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitReviewInfoResponse
func (client *Client) SubmitReviewInfoWithOptions(request *SubmitReviewInfoRequest, runtime *dara.RuntimeOptions) (_result *SubmitReviewInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitReviewInfo"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitReviewInfoRequest
//
// @return SubmitReviewInfoResponse
func (client *Client) SubmitReviewInfo(request *SubmitReviewInfoRequest) (_result *SubmitReviewInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.SubmitReviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交复核结果v4
//
// @param request - SubmitReviewInfoV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitReviewInfoV4Response
func (client *Client) SubmitReviewInfoV4WithOptions(request *SubmitReviewInfoV4Request, runtime *dara.RuntimeOptions) (_result *SubmitReviewInfoV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitReviewInfoV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitReviewInfoV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交复核结果v4
//
// @param request - SubmitReviewInfoV4Request
//
// @return SubmitReviewInfoV4Response
func (client *Client) SubmitReviewInfoV4(request *SubmitReviewInfoV4Request) (_result *SubmitReviewInfoV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitReviewInfoV4Response{}
	_body, _err := client.SubmitReviewInfoV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SyncQualityCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncQualityCheckResponse
func (client *Client) SyncQualityCheckWithOptions(request *SyncQualityCheckRequest, runtime *dara.RuntimeOptions) (_result *SyncQualityCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncQualityCheck"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SyncQualityCheckRequest
//
// @return SyncQualityCheckResponse
func (client *Client) SyncQualityCheck(request *SyncQualityCheckRequest) (_result *SyncQualityCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.SyncQualityCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 测试规则
//
// @param request - TestRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestRuleV4Response
func (client *Client) TestRuleV4WithOptions(request *TestRuleV4Request, runtime *dara.RuntimeOptions) (_result *TestRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsSchemeData) {
		body["IsSchemeData"] = request.IsSchemeData
	}

	if !dara.IsNil(request.TestJson) {
		body["TestJson"] = request.TestJson
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试规则
//
// @param request - TestRuleV4Request
//
// @return TestRuleV4Response
func (client *Client) TestRuleV4(request *TestRuleV4Request) (_result *TestRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestRuleV4Response{}
	_body, _err := client.TestRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAsrVocabResponse
func (client *Client) UpdateAsrVocabWithOptions(request *UpdateAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *UpdateAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAsrVocabRequest
//
// @return UpdateAsrVocabResponse
func (client *Client) UpdateAsrVocab(request *UpdateAsrVocabRequest) (_result *UpdateAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.UpdateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新质检方案中的质检维度
//
// @param request - UpdateCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCheckTypeToSchemeResponse
func (client *Client) UpdateCheckTypeToSchemeWithOptions(request *UpdateCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCheckTypeToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCheckTypeToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新质检方案中的质检维度
//
// @param request - UpdateCheckTypeToSchemeRequest
//
// @return UpdateCheckTypeToSchemeResponse
func (client *Client) UpdateCheckTypeToScheme(request *UpdateCheckTypeToSchemeRequest) (_result *UpdateCheckTypeToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCheckTypeToSchemeResponse{}
	_body, _err := client.UpdateCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新会话随录数据
//
// @param request - UpdateQualityCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckDataResponse
func (client *Client) UpdateQualityCheckDataWithOptions(request *UpdateQualityCheckDataRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityCheckDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityCheckData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新会话随录数据
//
// @param request - UpdateQualityCheckDataRequest
//
// @return UpdateQualityCheckDataResponse
func (client *Client) UpdateQualityCheckData(request *UpdateQualityCheckDataRequest) (_result *UpdateQualityCheckDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.UpdateQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新质检方案
//
// @param request - UpdateQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckSchemeResponse
func (client *Client) UpdateQualityCheckSchemeWithOptions(request *UpdateQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新质检方案
//
// @param request - UpdateQualityCheckSchemeRequest
//
// @return UpdateQualityCheckSchemeResponse
func (client *Client) UpdateQualityCheckScheme(request *UpdateQualityCheckSchemeRequest) (_result *UpdateQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityCheckSchemeResponse{}
	_body, _err := client.UpdateQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateRule is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// @param request - UpdateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleResponse
func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateRule is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// @param request - UpdateRuleRequest
//
// @return UpdateRuleResponse
// Deprecated
func (client *Client) UpdateRule(request *UpdateRuleRequest) (_result *UpdateRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleResponse{}
	_body, _err := client.UpdateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateRuleById is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// 更新规则
//
// @param request - UpdateRuleByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleByIdResponse
func (client *Client) UpdateRuleByIdWithOptions(request *UpdateRuleByIdRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IsCopy) {
		body["IsCopy"] = request.IsCopy
	}

	if !dara.IsNil(request.JsonStrForRule) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	if !dara.IsNil(request.ReturnRelatedSchemes) {
		body["ReturnRelatedSchemes"] = request.ReturnRelatedSchemes
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRuleById"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateRuleById is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// 更新规则
//
// @param request - UpdateRuleByIdRequest
//
// @return UpdateRuleByIdResponse
// Deprecated
func (client *Client) UpdateRuleById(request *UpdateRuleByIdRequest) (_result *UpdateRuleByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleByIdResponse{}
	_body, _err := client.UpdateRuleByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新质检方案的规则
//
// @param request - UpdateRuleToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleToSchemeResponse
func (client *Client) UpdateRuleToSchemeWithOptions(request *UpdateRuleToSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRuleToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新质检方案的规则
//
// @param request - UpdateRuleToSchemeRequest
//
// @return UpdateRuleToSchemeResponse
func (client *Client) UpdateRuleToScheme(request *UpdateRuleToSchemeRequest) (_result *UpdateRuleToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleToSchemeResponse{}
	_body, _err := client.UpdateRuleToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # V4更新规则
//
// @param request - UpdateRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleV4Response
func (client *Client) UpdateRuleV4WithOptions(request *UpdateRuleV4Request, runtime *dara.RuntimeOptions) (_result *UpdateRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JsonStrForRule) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # V4更新规则
//
// @param request - UpdateRuleV4Request
//
// @return UpdateRuleV4Response
func (client *Client) UpdateRuleV4(request *UpdateRuleV4Request) (_result *UpdateRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleV4Response{}
	_body, _err := client.UpdateRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新质检任务
//
// @param request - UpdateSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSchemeTaskConfigResponse
func (client *Client) UpdateSchemeTaskConfigWithOptions(request *UpdateSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新质检任务
//
// @param request - UpdateSchemeTaskConfigRequest
//
// @return UpdateSchemeTaskConfigResponse
func (client *Client) UpdateSchemeTaskConfig(request *UpdateSchemeTaskConfigRequest) (_result *UpdateSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSchemeTaskConfigResponse{}
	_body, _err := client.UpdateSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateSkillGroupConfig is deprecated
//
// @param request - UpdateSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillGroupConfigResponse
func (client *Client) UpdateSkillGroupConfigWithOptions(request *UpdateSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateSkillGroupConfig is deprecated
//
// @param request - UpdateSkillGroupConfigRequest
//
// @return UpdateSkillGroupConfigResponse
// Deprecated
func (client *Client) UpdateSkillGroupConfig(request *UpdateSkillGroupConfigRequest) (_result *UpdateSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.UpdateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateSyncQualityCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSyncQualityCheckDataResponse
func (client *Client) UpdateSyncQualityCheckDataWithOptions(request *UpdateSyncQualityCheckDataRequest, runtime *dara.RuntimeOptions) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSyncQualityCheckData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSyncQualityCheckDataRequest
//
// @return UpdateSyncQualityCheckDataResponse
func (client *Client) UpdateSyncQualityCheckData(request *UpdateSyncQualityCheckDataRequest) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.UpdateSyncQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskAssignRuleResponse
func (client *Client) UpdateTaskAssignRuleWithOptions(request *UpdateTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskAssignRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskAssignRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTaskAssignRuleRequest
//
// @return UpdateTaskAssignRuleResponse
func (client *Client) UpdateTaskAssignRule(request *UpdateTaskAssignRuleRequest) (_result *UpdateTaskAssignRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.UpdateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningConfigResponse
func (client *Client) UpdateWarningConfigWithOptions(request *UpdateWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateWarningConfigRequest
//
// @return UpdateWarningConfigResponse
func (client *Client) UpdateWarningConfig(request *UpdateWarningConfigRequest) (_result *UpdateWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.UpdateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-更新
//
// @param request - UpdateWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningStrategyConfigResponse
func (client *Client) UpdateWarningStrategyConfigWithOptions(request *UpdateWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-更新
//
// @param request - UpdateWarningStrategyConfigRequest
//
// @return UpdateWarningStrategyConfigResponse
func (client *Client) UpdateWarningStrategyConfig(request *UpdateWarningStrategyConfigRequest) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.UpdateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UploadAudioDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAudioDataResponse
func (client *Client) UploadAudioDataWithOptions(request *UploadAudioDataRequest, runtime *dara.RuntimeOptions) (_result *UploadAudioDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadAudioData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadAudioDataRequest
//
// @return UploadAudioDataResponse
func (client *Client) UploadAudioData(request *UploadAudioDataRequest) (_result *UploadAudioDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.UploadAudioDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UploadData is deprecated, please use Qualitycheck::2019-01-15::UploadDataV4 instead.
//
// Summary:
//
// 推荐使用UploadDataV4接口,支持更长的JsonStr,但仅支持POST方法.
//
// @param request - UploadDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataResponse
func (client *Client) UploadDataWithOptions(request *UploadDataRequest, runtime *dara.RuntimeOptions) (_result *UploadDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UploadData is deprecated, please use Qualitycheck::2019-01-15::UploadDataV4 instead.
//
// Summary:
//
// 推荐使用UploadDataV4接口,支持更长的JsonStr,但仅支持POST方法.
//
// @param request - UploadDataRequest
//
// @return UploadDataResponse
// Deprecated
func (client *Client) UploadData(request *UploadDataRequest) (_result *UploadDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataResponse{}
	_body, _err := client.UploadDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// http_hsf
//
// @param request - UploadDataSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataSyncResponse
func (client *Client) UploadDataSyncWithOptions(request *UploadDataSyncRequest, runtime *dara.RuntimeOptions) (_result *UploadDataSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDataSync"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// http_hsf
//
// @param request - UploadDataSyncRequest
//
// @return UploadDataSyncResponse
func (client *Client) UploadDataSync(request *UploadDataSyncRequest) (_result *UploadDataSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.UploadDataSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// http_hsf
//
// @param request - UploadDataSyncForLLMRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataSyncForLLMResponse
func (client *Client) UploadDataSyncForLLMWithOptions(request *UploadDataSyncForLLMRequest, runtime *dara.RuntimeOptions) (_result *UploadDataSyncForLLMResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDataSyncForLLM"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataSyncForLLMResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// http_hsf
//
// @param request - UploadDataSyncForLLMRequest
//
// @return UploadDataSyncForLLMResponse
func (client *Client) UploadDataSyncForLLM(request *UploadDataSyncForLLMRequest) (_result *UploadDataSyncForLLMResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataSyncForLLMResponse{}
	_body, _err := client.UploadDataSyncForLLMWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # UploadDataV4
//
// @param request - UploadDataV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataV4Response
func (client *Client) UploadDataV4WithOptions(request *UploadDataV4Request, runtime *dara.RuntimeOptions) (_result *UploadDataV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDataV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UploadDataV4
//
// @param request - UploadDataV4Request
//
// @return UploadDataV4Response
func (client *Client) UploadDataV4(request *UploadDataV4Request) (_result *UploadDataV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataV4Response{}
	_body, _err := client.UploadDataV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UploadRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRuleResponse
func (client *Client) UploadRuleWithOptions(request *UploadRuleRequest, runtime *dara.RuntimeOptions) (_result *UploadRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadRuleRequest
//
// @return UploadRuleResponse
func (client *Client) UploadRule(request *UploadRuleRequest) (_result *UploadRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadRuleResponse{}
	_body, _err := client.UploadRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VerifyFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyFileResponse
func (client *Client) VerifyFileWithOptions(request *VerifyFileRequest, runtime *dara.RuntimeOptions) (_result *VerifyFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyFile"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifyFileRequest
//
// @return VerifyFileResponse
func (client *Client) VerifyFile(request *VerifyFileRequest) (_result *VerifyFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyFileResponse{}
	_body, _err := client.VerifyFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - VerifySentenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySentenceResponse
func (client *Client) VerifySentenceWithOptions(request *VerifySentenceRequest, runtime *dara.RuntimeOptions) (_result *VerifySentenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifySentence"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifySentenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifySentenceRequest
//
// @return VerifySentenceResponse
func (client *Client) VerifySentence(request *VerifySentenceRequest) (_result *VerifySentenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifySentenceResponse{}
	_body, _err := client.VerifySentenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
