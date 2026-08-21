// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 获取模型部署方案匹配资源
//
// @param request - GetModelDeploymentResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelDeploymentResourcesResponse
func (client *Client) GetModelDeploymentResourcesWithContext(ctx context.Context, ModelId *string, request *GetModelDeploymentResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelDeploymentResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizKey) {
		query["BizKey"] = request.BizKey
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.ProfileId) {
		query["ProfileId"] = request.ProfileId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelDeploymentResources"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/deployment-resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelDeploymentResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成模型部署配置
//
// @param request - GetModelDeploymentSpecRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelDeploymentSpecResponse
func (client *Client) GetModelDeploymentSpecWithContext(ctx context.Context, ModelId *string, request *GetModelDeploymentSpecRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelDeploymentSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizKey) {
		query["BizKey"] = request.BizKey
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.ProfileId) {
		query["ProfileId"] = request.ProfileId
	}

	if !dara.IsNil(request.ResourceSelections) {
		query["ResourceSelections"] = request.ResourceSelections
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelDeploymentSpec"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/deployment-spec"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelDeploymentSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型部署方案列表
//
// @param request - ListModelDeploymentProfilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelDeploymentProfilesResponse
func (client *Client) ListModelDeploymentProfilesWithContext(ctx context.Context, ModelId *string, request *ListModelDeploymentProfilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelDeploymentProfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizKey) {
		query["BizKey"] = request.BizKey
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelDeploymentProfiles"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/deployment-profiles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelDeploymentProfilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ModelGallery模型列表
//
// @param tmpReq - ListModelGalleryModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelGalleryModelsResponse
func (client *Client) ListModelGalleryModelsWithContext(ctx context.Context, tmpReq *ListModelGalleryModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelGalleryModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListModelGalleryModelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Conditions) {
		request.ConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Conditions, dara.String("Conditions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collections) {
		query["Collections"] = request.Collections
	}

	if !dara.IsNil(request.Compressible) {
		query["Compressible"] = request.Compressible
	}

	if !dara.IsNil(request.ConditionsShrink) {
		query["Conditions"] = request.ConditionsShrink
	}

	if !dara.IsNil(request.DeepThink) {
		query["DeepThink"] = request.DeepThink
	}

	if !dara.IsNil(request.Demonstrable) {
		query["Demonstrable"] = request.Demonstrable
	}

	if !dara.IsNil(request.Deployable) {
		query["Deployable"] = request.Deployable
	}

	if !dara.IsNil(request.Distillable) {
		query["Distillable"] = request.Distillable
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Evaluable) {
		query["Evaluable"] = request.Evaluable
	}

	if !dara.IsNil(request.FunctionCall) {
		query["FunctionCall"] = request.FunctionCall
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelSeries) {
		query["ModelSeries"] = request.ModelSeries
	}

	if !dara.IsNil(request.ModelType) {
		query["ModelType"] = request.ModelType
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SupportedCompressionResource) {
		query["SupportedCompressionResource"] = request.SupportedCompressionResource
	}

	if !dara.IsNil(request.SupportedDistillationResource) {
		query["SupportedDistillationResource"] = request.SupportedDistillationResource
	}

	if !dara.IsNil(request.SupportedEvaluationResource) {
		query["SupportedEvaluationResource"] = request.SupportedEvaluationResource
	}

	if !dara.IsNil(request.SupportedInferenceResource) {
		query["SupportedInferenceResource"] = request.SupportedInferenceResource
	}

	if !dara.IsNil(request.SupportedTrainingResource) {
		query["SupportedTrainingResource"] = request.SupportedTrainingResource
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.Task) {
		query["Task"] = request.Task
	}

	if !dara.IsNil(request.Trainable) {
		query["Trainable"] = request.Trainable
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelGalleryModels"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelGalleryModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
