// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Assign personnel to the worker nodes (annotation, quality inspection, and validation) of an annotation job.
//
// @param request - AddWorkNodeWorkforceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkNodeWorkforceResponse
func (client *Client) AddWorkNodeWorkforceWithContext(ctx context.Context, TenantId *string, TaskId *string, WorkNodeId *string, request *AddWorkNodeWorkforceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddWorkNodeWorkforceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIds) {
		body["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWorkNodeWorkforce"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/worknodes/" + dara.PercentEncode(dara.StringValue(WorkNodeId)) + "/workforce"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWorkNodeWorkforceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Append data to a job.
//
// @param request - AppendAllDataToTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppendAllDataToTaskResponse
func (client *Client) AppendAllDataToTaskWithContext(ctx context.Context, TenantId *string, TaskId *string, request *AppendAllDataToTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AppendAllDataToTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AppendAllDataToTask"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/appendAllDataToTask"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AppendAllDataToTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an annotation job for the current tenant.
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithContext(ctx context.Context, TenantId *string, request *CreateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTask"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can add a new template for the current tenant and customize the annotation template based on your business requirements.
//
// @param request - CreateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithContext(ctx context.Context, TenantId *string, request *CreateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/templates"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a member to the tenant.
//
// @param request - CreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithContext(ctx context.Context, TenantId *string, request *CreateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.AccountType) {
		body["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.Role) {
		body["Role"] = request.Role
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/users"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a job under the current tenant.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithContext(ctx context.Context, TenantId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTask"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the template under the current tenant.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithContext(ctx context.Context, TenantId *string, TemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/templates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a member within a tenant.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithContext(ctx context.Context, TenantId *string, UserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/users/" + dara.PercentEncode(dara.StringValue(UserId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Export the result data of an annotation job.
//
// @param request - ExportAnnotationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportAnnotationsResponse
func (client *Client) ExportAnnotationsWithContext(ctx context.Context, TenantId *string, TaskId *string, request *ExportAnnotationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportAnnotationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssPath) {
		query["OssPath"] = request.OssPath
	}

	if !dara.IsNil(request.RegisterDataset) {
		query["RegisterDataset"] = request.RegisterDataset
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportAnnotations"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/annotations/export"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportAnnotationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the information of a single annotation export result.
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithContext(ctx context.Context, TenantId *string, JobId *string, request *GetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(JobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the information of a single subtask package.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubtaskResponse
func (client *Client) GetSubtaskWithContext(ctx context.Context, TenantId *string, TaskID *string, SubtaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSubtaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubtask"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskID)) + "/subtasks/" + dara.PercentEncode(dara.StringValue(SubtaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubtaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query a single annotated data item in a subtask package.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubtaskItemResponse
func (client *Client) GetSubtaskItemWithContext(ctx context.Context, TenantId *string, TaskId *string, SubtaskId *string, ItemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSubtaskItemResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubtaskItem"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/subtasks/" + dara.PercentEncode(dara.StringValue(SubtaskId)) + "/items/" + dara.PercentEncode(dara.StringValue(ItemId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubtaskItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the information of a single annotation job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, TenantId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the current statistics information of a job.
//
// @param request - GetTaskStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatisticsResponse
func (client *Client) GetTaskStatisticsWithContext(ctx context.Context, TenantId *string, TaskId *string, request *GetTaskStatisticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.StatType) {
		query["StatType"] = request.StatType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskStatistics"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/statistics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the current status of a job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithContext(ctx context.Context, TenantId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskStatus"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the current template information of a job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskTemplateResponse
func (client *Client) GetTaskTemplateWithContext(ctx context.Context, TenantId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskTemplate"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/template"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query job template questions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskTemplateQuestionsResponse
func (client *Client) GetTaskTemplateQuestionsWithContext(ctx context.Context, TenantId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskTemplateQuestionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskTemplateQuestions"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/template/questions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskTemplateQuestionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the data display information in the job template.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskTemplateViewsResponse
func (client *Client) GetTaskTemplateViewsWithContext(ctx context.Context, TenantId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskTemplateViewsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskTemplateViews"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/template/views"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskTemplateViewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the personnel configuration information of each node in a job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskWorkforceResponse
func (client *Client) GetTaskWorkforceWithContext(ctx context.Context, TenantId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskWorkforceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskWorkforce"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/workforce"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskWorkforceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query statistics of each member in a job.
//
// @param request - GetTaskWorkforceStatisticRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskWorkforceStatisticResponse
func (client *Client) GetTaskWorkforceStatisticWithContext(ctx context.Context, TenantId *string, TaskId *string, request *GetTaskWorkforceStatisticRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskWorkforceStatisticResponse, _err error) {
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

	if !dara.IsNil(request.StatType) {
		query["StatType"] = request.StatType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskWorkforceStatistic"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/workforce/statistic"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskWorkforceStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query template information under a tenant.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithContext(ctx context.Context, TenantId *string, TemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/templates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query question information such as Radio and Multiple Choice in a template.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateQuestionsResponse
func (client *Client) GetTemplateQuestionsWithContext(ctx context.Context, TenantId *string, TemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateQuestionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateQuestions"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/templates/" + dara.PercentEncode(dara.StringValue(TemplateId)) + "/questions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateQuestionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the display information such as images, text, and audio in the template.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateViewResponse
func (client *Client) GetTemplateViewWithContext(ctx context.Context, TenantId *string, TemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateViewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateView"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/templates/" + dara.PercentEncode(dara.StringValue(TemplateId)) + "/views"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query information about the iTAG tenant.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTenantResponse
func (client *Client) GetTenantWithContext(ctx context.Context, TenantId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTenantResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTenant"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTenantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the details of a single member in a tenant.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithContext(ctx context.Context, TenantId *string, UserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/users/" + dara.PercentEncode(dara.StringValue(UserId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays a list of all exported annotation results.
//
// @param request - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, TenantId *string, request *ListJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Display the annotation data of a single subtask package.
//
// @param request - ListSubtaskItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubtaskItemsResponse
func (client *Client) ListSubtaskItemsWithContext(ctx context.Context, TenantId *string, TaskID *string, SubtaskId *string, request *ListSubtaskItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubtaskItemsResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubtaskItems"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskID)) + "/subtasks/" + dara.PercentEncode(dara.StringValue(SubtaskId)) + "/items"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubtaskItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the list of subtask packages.
//
// @param request - ListSubtasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubtasksResponse
func (client *Client) ListSubtasksWithContext(ctx context.Context, TenantId *string, TaskID *string, request *ListSubtasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubtasksResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubtasks"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskID)) + "/subtasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubtasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the list of annotation jobs for the current tenant.
//
// @param request - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, TenantId *string, request *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Display the template list of the current tenant.
//
// @param tmpReq - ListTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithContext(ctx context.Context, TenantId *string, tmpReq *ListTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.TypesShrink) {
		query["Types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplates"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query iTAG tenants under an Alibaba Cloud account.
//
// @param request - ListTenantsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantsResponse
func (client *Client) ListTenantsWithContext(ctx context.Context, request *ListTenantsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTenantsResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTenants"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTenantsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays all annotate members under the current tenant.
//
// @param request - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, TenantId *string, request *ListUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// Summary:
//
// # Delete Node Personnel
//
// @param request - RemoveWorkNodeWorkforceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveWorkNodeWorkforceResponse
func (client *Client) RemoveWorkNodeWorkforceWithContext(ctx context.Context, TenantId *string, TaskId *string, WorkNodeId *string, request *RemoveWorkNodeWorkforceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveWorkNodeWorkforceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIds) {
		body["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveWorkNodeWorkforce"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/worknodes/" + dara.PercentEncode(dara.StringValue(WorkNodeId)) + "/workforce"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveWorkNodeWorkforceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify a job under the current tenant.
//
// @param request - UpdateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskResponse
func (client *Client) UpdateTaskWithContext(ctx context.Context, TenantId *string, TaskId *string, request *UpdateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTask"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update job members.
//
// @param request - UpdateTaskWorkforceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskWorkforceResponse
func (client *Client) UpdateTaskWorkforceWithContext(ctx context.Context, TenantId *string, TaskId *string, request *UpdateTaskWorkforceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTaskWorkforceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Workforce) {
		body["Workforce"] = request.Workforce
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskWorkforce"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/workforce"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskWorkforceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the template under the current tenant.
//
// @param request - UpdateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithContext(ctx context.Context, TenantId *string, TemplateId *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/templates/" + dara.PercentEncode(dara.StringValue(TemplateId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the information of an iTAG tenant.
//
// @param request - UpdateTenantRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantResponse
func (client *Client) UpdateTenantWithContext(ctx context.Context, TenantId *string, request *UpdateTenantRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTenantResponse, _err error) {
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

	if !dara.IsNil(request.TenantName) {
		body["TenantName"] = request.TenantName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTenant"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTenantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update member information within a tenant.
//
// @param request - UpdateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithContext(ctx context.Context, TenantId *string, UserId *string, request *UpdateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Role) {
		body["Role"] = request.Role
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2022-06-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(TenantId)) + "/users/" + dara.PercentEncode(dara.StringValue(UserId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
