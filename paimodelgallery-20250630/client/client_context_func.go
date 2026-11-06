// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建任务计划。
//
// @param request - CreateJobPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobPlanResponse
func (client *Client) CreateJobPlanWithContext(ctx context.Context, request *CreateJobPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobPlanName) {
		body["JobPlanName"] = request.JobPlanName
	}

	if !dara.IsNil(request.JobPlanSteps) {
		body["JobPlanSteps"] = request.JobPlanSteps
	}

	if !dara.IsNil(request.JobPlanType) {
		body["JobPlanType"] = request.JobPlanType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJobPlan"),
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobplans"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务计划。
//
// @param request - DeleteJobPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobPlanResponse
func (client *Client) DeleteJobPlanWithContext(ctx context.Context, JobPlanId *string, request *DeleteJobPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteJobPlanResponse, _err error) {
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
		Action:      dara.String("DeleteJobPlan"),
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobplans/" + dara.PercentEncode(dara.StringValue(JobPlanId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询蒸馏模板详情。
//
// @param request - GetDistillationTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDistillationTemplateResponse
func (client *Client) GetDistillationTemplateWithContext(ctx context.Context, TemplateId *string, request *GetDistillationTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDistillationTemplateResponse, _err error) {
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
		Action:      dara.String("GetDistillationTemplate"),
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/distillationtemplates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDistillationTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务计划。
//
// @param request - GetJobPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobPlanResponse
func (client *Client) GetJobPlanWithContext(ctx context.Context, JobPlanId *string, request *GetJobPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobPlanResponse, _err error) {
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
		Action:      dara.String("GetJobPlan"),
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobplans/" + dara.PercentEncode(dara.StringValue(JobPlanId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询蒸馏模板列表。
//
// @param request - ListDistillationTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDistillationTemplatesResponse
func (client *Client) ListDistillationTemplatesWithContext(ctx context.Context, request *ListDistillationTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDistillationTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateKind) {
		query["TemplateKind"] = request.TemplateKind
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDistillationTemplates"),
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/distillationtemplates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDistillationTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务计划列表。
//
// @param tmpReq - ListJobPlansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobPlansResponse
func (client *Client) ListJobPlansWithContext(ctx context.Context, tmpReq *ListJobPlansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListJobPlansShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HasTemplate) {
		query["HasTemplate"] = request.HasTemplate
	}

	if !dara.IsNil(request.JobPlanName) {
		query["JobPlanName"] = request.JobPlanName
	}

	if !dara.IsNil(request.JobPlanType) {
		query["JobPlanType"] = request.JobPlanType
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

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobPlans"),
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobplans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobPlansResponse{}
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
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelgallery/models"),
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

// Summary:
//
// 更新任务计划。
//
// @param request - UpdateJobPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobPlanResponse
func (client *Client) UpdateJobPlanWithContext(ctx context.Context, JobPlanId *string, request *UpdateJobPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateJobPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobPlanCurrentStep) {
		body["JobPlanCurrentStep"] = request.JobPlanCurrentStep
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJobPlan"),
		Version:     dara.String("2025-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobplans/" + dara.PercentEncode(dara.StringValue(JobPlanId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJobPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
