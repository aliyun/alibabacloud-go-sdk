// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a new job plan that supports task types such as distillation.
//
// Description:
//
// ## Operation description
//
// - This API operation creates a new node plan. You can specify the workspace ID, node plan type, name, and steps.
//
// - If you use a scenario-specific distillation template, provide the `TemplateId` parameter and make sure that `JobPlanSteps` contains distillation configurations that match the template.
//
// - The `Tag` parameter follows the Alibaba Cloud label system specification and is used to add additional identity information to the node plan.
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
// Deletes a specified task plan by task plan ID.
//
// Description:
//
// ## Operation description.
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
// Retrieves the details of a specified distillation template, including preset configurations and default YAML.
//
// Description:
//
// ## Operation description
//
// - This API operation queries the full details of a specific distillation template to facilitate rendering the creation form.
//
// - The template ID is a required parameter, obtained from the `ListDistillationTemplates` operation.
//
// - The `DefaultConfig` field provides the complete EasyDistill configuration YAML (with comments). All paths in the YAML are relative paths, and callers do not need to perform absolute path conversion or string replacement.
//
// - All translatable fields are automatically parsed into the corresponding language version based on the `x-acs-accept-language` request header.
//
// - If the `TrainingOptions` field is missing, the template supports only the first stage of processing. Attempts to use such a template to create a task that includes the second stage will fail.
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
// Queries the details of a specified job plan by job plan ID.
//
// Description:
//
// ## Description
//
// This API operation retrieves the details of a job plan based on the specified job plan ID (`JobPlanId`), including but not limited to the name, type, and current step status of the job plan. Make sure the `JobPlanId` provided in the request is valid and belongs to your workspace.
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
// Retrieves summary information of distillation templates for rendering the template card list.
//
// Description:
//
// ## Operation description
//
// - This operation supports filtered queries by using the TemplateId, Category, and Keyword parameters.
//
// - Pagination is controlled by the PageNumber and PageSize parameters, consistent with other paginated operations of the same service.
//
// - Templates are public resources that do not belong to any workspace. Therefore, you do not need to specify WorkspaceId.
//
// - All translatable fields such as TemplateName and Description are automatically parsed into the corresponding language version based on the x-acs-accept-language request header.
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
// Retrieves the list of job plans in a specified workspace.
//
// Description:
//
// ## Operation description
//
// - `WorkspaceId` is a required parameter that specifies the workspace to which the job plans belong.
//
// - The `Tag` parameter must be encoded by using `EncodeURI` before being passed.
//
// - If both `TemplateId` and `HasTemplate` are specified, the value of `TemplateId` takes precedence for filtering.
//
// - `JobPlanName` supports exact match. Enclose the specific name in quotation marks.
//
// - By default, results are sorted in descending order by creation time (`GmtCreateTime`). Set the `Order` parameter to `ASC` to change the sort order.
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
// Retrieves the list of models from ModelGallery (deprecated, use the 2026-06-03 version instead).
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
// Updates the current step and tag information of a specified task plan.
//
// Description:
//
// ## Request description
//
// You can use this API operation to update the current execution step (`JobPlanCurrentStep`) and related tags (`Tag`) of a specific task plan identified by `JobPlanId`. If the request contains tag information, tags are updated or added based on the provided key-value pairs.
//
// - **JobPlanId*	- is a path parameter. You must provide a valid task plan ID.
//
// - **JobPlanCurrentStep*	- is an optional parameter that specifies the new current step of the task.
//
// - **Tag*	- is an optional parameter that specifies a list of key-value pairs used to label the task plan. Each tag consists of a `Key` and a `Value`.
//
// Note: Ensure that the `JobPlanId` you provide exists and that you have the permissions to modify it.
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
