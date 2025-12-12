// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AddBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBusinessCategoryResponse
func (client *Client) AddBusinessCategoryWithContext(ctx context.Context, request *AddBusinessCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddBusinessCategoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleCategoryResponse
func (client *Client) AddRuleCategoryWithContext(ctx context.Context, request *AddRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddRuleCategoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleV4Response
func (client *Client) AddRuleV4WithContext(ctx context.Context, request *AddRuleV4Request, runtime *dara.RuntimeOptions) (_result *AddRuleV4Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyWsTokenResponse
func (client *Client) ApplyWsTokenWithContext(ctx context.Context, request *ApplyWsTokenRequest, runtime *dara.RuntimeOptions) (_result *ApplyWsTokenResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssignReviewerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerResponse
func (client *Client) AssignReviewerWithContext(ctx context.Context, request *AssignReviewerRequest, runtime *dara.RuntimeOptions) (_result *AssignReviewerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerBySessionGroupResponse
func (client *Client) AssignReviewerBySessionGroupWithContext(ctx context.Context, request *AssignReviewerBySessionGroupRequest, runtime *dara.RuntimeOptions) (_result *AssignReviewerBySessionGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSubmitReviewInfoResponse
func (client *Client) BatchSubmitReviewInfoWithContext(ctx context.Context, request *BatchSubmitReviewInfoRequest, runtime *dara.RuntimeOptions) (_result *BatchSubmitReviewInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsrVocabResponse
func (client *Client) CreateAsrVocabWithContext(ctx context.Context, request *CreateAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *CreateAsrVocabResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCheckTypeToSchemeResponse
func (client *Client) CreateCheckTypeToSchemeWithContext(ctx context.Context, request *CreateCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateCheckTypeToSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMiningTaskResponse
func (client *Client) CreateMiningTaskWithContext(ctx context.Context, request *CreateMiningTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMiningTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityCheckSchemeResponse
func (client *Client) CreateQualityCheckSchemeWithContext(ctx context.Context, request *CreateQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityCheckSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemeTaskConfigResponse
func (client *Client) CreateSchemeTaskConfigWithContext(ctx context.Context, request *CreateSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSchemeTaskConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillGroupConfigResponse
func (client *Client) CreateSkillGroupConfigWithContext(ctx context.Context, request *CreateSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillGroupConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskAssignRuleResponse
func (client *Client) CreateTaskAssignRuleWithContext(ctx context.Context, request *CreateTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateTaskAssignRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningConfigResponse
func (client *Client) CreateWarningConfigWithContext(ctx context.Context, request *CreateWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateWarningConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningStrategyConfigResponse
func (client *Client) CreateWarningStrategyConfigWithContext(ctx context.Context, request *CreateWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateWarningStrategyConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DelRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelRuleCategoryResponse
func (client *Client) DelRuleCategoryWithContext(ctx context.Context, request *DelRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *DelRuleCategoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsrVocabResponse
func (client *Client) DeleteAsrVocabWithContext(ctx context.Context, request *DeleteAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *DeleteAsrVocabResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessCategoryResponse
func (client *Client) DeleteBusinessCategoryWithContext(ctx context.Context, request *DeleteBusinessCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteBusinessCategoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCheckTypeToSchemeResponse
func (client *Client) DeleteCheckTypeToSchemeWithContext(ctx context.Context, request *DeleteCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteCheckTypeToSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomizationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizationConfigResponse
func (client *Client) DeleteCustomizationConfigWithContext(ctx context.Context, request *DeleteCustomizationConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomizationConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSetResponse
func (client *Client) DeleteDataSetWithContext(ctx context.Context, request *DeleteDataSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeletePrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrecisionTaskResponse
func (client *Client) DeletePrecisionTaskWithContext(ctx context.Context, request *DeletePrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *DeletePrecisionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityCheckSchemeResponse
func (client *Client) DeleteQualityCheckSchemeWithContext(ctx context.Context, request *DeleteQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityCheckSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleV4Response
func (client *Client) DeleteRuleV4WithContext(ctx context.Context, request *DeleteRuleV4Request, runtime *dara.RuntimeOptions) (_result *DeleteRuleV4Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchemeTaskConfigResponse
func (client *Client) DeleteSchemeTaskConfigWithContext(ctx context.Context, request *DeleteSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSchemeTaskConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillGroupConfigResponse
func (client *Client) DeleteSkillGroupConfigWithContext(ctx context.Context, request *DeleteSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillGroupConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskAssignRuleResponse
func (client *Client) DeleteTaskAssignRuleWithContext(ctx context.Context, request *DeleteTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskAssignRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningConfigResponse
func (client *Client) DeleteWarningConfigWithContext(ctx context.Context, request *DeleteWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarningConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningStrategyConfigResponse
func (client *Client) DeleteWarningStrategyConfigWithContext(ctx context.Context, request *DeleteWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarningStrategyConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsrVocabResponse
func (client *Client) GetAsrVocabWithContext(ctx context.Context, request *GetAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *GetAsrVocabResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetBusinessCategoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusinessCategoryListResponse
func (client *Client) GetBusinessCategoryListWithContext(ctx context.Context, request *GetBusinessCategoryListRequest, runtime *dara.RuntimeOptions) (_result *GetBusinessCategoryListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomizationConfigListResponse
func (client *Client) GetCustomizationConfigListWithContext(ctx context.Context, request *GetCustomizationConfigListRequest, runtime *dara.RuntimeOptions) (_result *GetCustomizationConfigListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMiningTaskResultResponse
func (client *Client) GetMiningTaskResultWithContext(ctx context.Context, request *GetMiningTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetMiningTaskResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNextResultToVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNextResultToVerifyResponse
func (client *Client) GetNextResultToVerifyWithContext(ctx context.Context, request *GetNextResultToVerifyRequest, runtime *dara.RuntimeOptions) (_result *GetNextResultToVerifyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrecisionTaskResponse
func (client *Client) GetPrecisionTaskWithContext(ctx context.Context, request *GetPrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *GetPrecisionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityCheckSchemeResponse
func (client *Client) GetQualityCheckSchemeWithContext(ctx context.Context, request *GetQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *GetQualityCheckSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultResponse
func (client *Client) GetResultWithContext(ctx context.Context, request *GetResultRequest, runtime *dara.RuntimeOptions) (_result *GetResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetResultToReviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultToReviewResponse
func (client *Client) GetResultToReviewWithContext(ctx context.Context, request *GetResultToReviewRequest, runtime *dara.RuntimeOptions) (_result *GetResultToReviewResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleResponse
func (client *Client) GetRuleWithContext(ctx context.Context, request *GetRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleByIdResponse
func (client *Client) GetRuleByIdWithContext(ctx context.Context, request *GetRuleByIdRequest, runtime *dara.RuntimeOptions) (_result *GetRuleByIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleCategoryResponse
func (client *Client) GetRuleCategoryWithContext(ctx context.Context, request *GetRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetRuleCategoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleDetailResponse
func (client *Client) GetRuleDetailWithContext(ctx context.Context, request *GetRuleDetailRequest, runtime *dara.RuntimeOptions) (_result *GetRuleDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleV4Response
func (client *Client) GetRuleV4WithContext(ctx context.Context, request *GetRuleV4Request, runtime *dara.RuntimeOptions) (_result *GetRuleV4Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRulesCountListResponse
func (client *Client) GetRulesCountListWithContext(ctx context.Context, request *GetRulesCountListRequest, runtime *dara.RuntimeOptions) (_result *GetRulesCountListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemeTaskConfigResponse
func (client *Client) GetSchemeTaskConfigWithContext(ctx context.Context, request *GetSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSchemeTaskConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScoreInfoResponse
func (client *Client) GetScoreInfoWithContext(ctx context.Context, request *GetScoreInfoRequest, runtime *dara.RuntimeOptions) (_result *GetScoreInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupConfigResponse
func (client *Client) GetSkillGroupConfigWithContext(ctx context.Context, request *GetSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyncResultResponse
func (client *Client) GetSyncResultWithContext(ctx context.Context, request *GetSyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetSyncResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarningStrategyConfigResponse
func (client *Client) GetWarningStrategyConfigWithContext(ctx context.Context, request *GetWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *GetWarningStrategyConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - HandleComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleComplaintResponse
func (client *Client) HandleComplaintWithContext(ctx context.Context, request *HandleComplaintRequest, runtime *dara.RuntimeOptions) (_result *HandleComplaintResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidRuleResponse
func (client *Client) InvalidRuleWithContext(ctx context.Context, request *InvalidRuleRequest, runtime *dara.RuntimeOptions) (_result *InvalidRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsrVocabResponse
func (client *Client) ListAsrVocabWithContext(ctx context.Context, request *ListAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *ListAsrVocabResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSetResponse
func (client *Client) ListDataSetWithContext(ctx context.Context, request *ListDataSetRequest, runtime *dara.RuntimeOptions) (_result *ListDataSetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrecisionTaskResponse
func (client *Client) ListPrecisionTaskWithContext(ctx context.Context, request *ListPrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *ListPrecisionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityCheckSchemeResponse
func (client *Client) ListQualityCheckSchemeWithContext(ctx context.Context, request *ListQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *ListQualityCheckSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
func (client *Client) ListRulesWithContext(ctx context.Context, request *ListRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesV4Response
func (client *Client) ListRulesV4WithContext(ctx context.Context, request *ListRulesV4Request, runtime *dara.RuntimeOptions) (_result *ListRulesV4Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemeTaskConfigResponse
func (client *Client) ListSchemeTaskConfigWithContext(ctx context.Context, request *ListSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *ListSchemeTaskConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionGroupResponse
func (client *Client) ListSessionGroupWithContext(ctx context.Context, request *ListSessionGroupRequest, runtime *dara.RuntimeOptions) (_result *ListSessionGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupConfigResponse
func (client *Client) ListSkillGroupConfigWithContext(ctx context.Context, request *ListSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *ListSkillGroupConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTaskAssignRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskAssignRulesResponse
func (client *Client) ListTaskAssignRulesWithContext(ctx context.Context, request *ListTaskAssignRulesRequest, runtime *dara.RuntimeOptions) (_result *ListTaskAssignRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningConfigResponse
func (client *Client) ListWarningConfigWithContext(ctx context.Context, request *ListWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *ListWarningConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningStrategyConfigResponse
func (client *Client) ListWarningStrategyConfigWithContext(ctx context.Context, request *ListWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *ListWarningStrategyConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionResponse
func (client *Client) RevertAssignedSessionWithContext(ctx context.Context, request *RevertAssignedSessionRequest, runtime *dara.RuntimeOptions) (_result *RevertAssignedSessionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionGroupResponse
func (client *Client) RevertAssignedSessionGroupWithContext(ctx context.Context, request *RevertAssignedSessionGroupRequest, runtime *dara.RuntimeOptions) (_result *RevertAssignedSessionGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveConfigDataSetResponse
func (client *Client) SaveConfigDataSetWithContext(ctx context.Context, request *SaveConfigDataSetRequest, runtime *dara.RuntimeOptions) (_result *SaveConfigDataSetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitComplaintResponse
func (client *Client) SubmitComplaintWithContext(ctx context.Context, request *SubmitComplaintRequest, runtime *dara.RuntimeOptions) (_result *SubmitComplaintResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitPrecisionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPrecisionTaskResponse
func (client *Client) SubmitPrecisionTaskWithContext(ctx context.Context, request *SubmitPrecisionTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitPrecisionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitQualityCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityCheckTaskResponse
func (client *Client) SubmitQualityCheckTaskWithContext(ctx context.Context, request *SubmitQualityCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitQualityCheckTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitReviewInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitReviewInfoResponse
func (client *Client) SubmitReviewInfoWithContext(ctx context.Context, request *SubmitReviewInfoRequest, runtime *dara.RuntimeOptions) (_result *SubmitReviewInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SyncQualityCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncQualityCheckResponse
func (client *Client) SyncQualityCheckWithContext(ctx context.Context, request *SyncQualityCheckRequest, runtime *dara.RuntimeOptions) (_result *SyncQualityCheckResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestRuleV4Response
func (client *Client) TestRuleV4WithContext(ctx context.Context, request *TestRuleV4Request, runtime *dara.RuntimeOptions) (_result *TestRuleV4Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAsrVocabResponse
func (client *Client) UpdateAsrVocabWithContext(ctx context.Context, request *UpdateAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *UpdateAsrVocabResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCheckTypeToSchemeResponse
func (client *Client) UpdateCheckTypeToSchemeWithContext(ctx context.Context, request *UpdateCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCheckTypeToSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckDataResponse
func (client *Client) UpdateQualityCheckDataWithContext(ctx context.Context, request *UpdateQualityCheckDataRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityCheckDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckSchemeResponse
func (client *Client) UpdateQualityCheckSchemeWithContext(ctx context.Context, request *UpdateQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityCheckSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleResponse
func (client *Client) UpdateRuleWithContext(ctx context.Context, request *UpdateRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleByIdResponse
func (client *Client) UpdateRuleByIdWithContext(ctx context.Context, request *UpdateRuleByIdRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleByIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleToSchemeResponse
func (client *Client) UpdateRuleToSchemeWithContext(ctx context.Context, request *UpdateRuleToSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleToSchemeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleV4Response
func (client *Client) UpdateRuleV4WithContext(ctx context.Context, request *UpdateRuleV4Request, runtime *dara.RuntimeOptions) (_result *UpdateRuleV4Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSchemeTaskConfigResponse
func (client *Client) UpdateSchemeTaskConfigWithContext(ctx context.Context, request *UpdateSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSchemeTaskConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillGroupConfigResponse
func (client *Client) UpdateSkillGroupConfigWithContext(ctx context.Context, request *UpdateSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillGroupConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSyncQualityCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSyncQualityCheckDataResponse
func (client *Client) UpdateSyncQualityCheckDataWithContext(ctx context.Context, request *UpdateSyncQualityCheckDataRequest, runtime *dara.RuntimeOptions) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskAssignRuleResponse
func (client *Client) UpdateTaskAssignRuleWithContext(ctx context.Context, request *UpdateTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskAssignRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningConfigResponse
func (client *Client) UpdateWarningConfigWithContext(ctx context.Context, request *UpdateWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateWarningConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningStrategyConfigResponse
func (client *Client) UpdateWarningStrategyConfigWithContext(ctx context.Context, request *UpdateWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateWarningStrategyConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadAudioDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAudioDataResponse
func (client *Client) UploadAudioDataWithContext(ctx context.Context, request *UploadAudioDataRequest, runtime *dara.RuntimeOptions) (_result *UploadAudioDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataResponse
func (client *Client) UploadDataWithContext(ctx context.Context, request *UploadDataRequest, runtime *dara.RuntimeOptions) (_result *UploadDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataSyncResponse
func (client *Client) UploadDataSyncWithContext(ctx context.Context, request *UploadDataSyncRequest, runtime *dara.RuntimeOptions) (_result *UploadDataSyncResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataSyncForLLMResponse
func (client *Client) UploadDataSyncForLLMWithContext(ctx context.Context, request *UploadDataSyncForLLMRequest, runtime *dara.RuntimeOptions) (_result *UploadDataSyncForLLMResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataV4Response
func (client *Client) UploadDataV4WithContext(ctx context.Context, request *UploadDataV4Request, runtime *dara.RuntimeOptions) (_result *UploadDataV4Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRuleResponse
func (client *Client) UploadRuleWithContext(ctx context.Context, request *UploadRuleRequest, runtime *dara.RuntimeOptions) (_result *UploadRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifyFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyFileResponse
func (client *Client) VerifyFileWithContext(ctx context.Context, request *VerifyFileRequest, runtime *dara.RuntimeOptions) (_result *VerifyFileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - VerifySentenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySentenceResponse
func (client *Client) VerifySentenceWithContext(ctx context.Context, request *VerifySentenceRequest, runtime *dara.RuntimeOptions) (_result *VerifySentenceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
