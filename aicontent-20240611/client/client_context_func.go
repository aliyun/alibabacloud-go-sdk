// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Generate Q&A pairs to expand data.
//
// @param request - AITeacherExpansionPracticeTaskGenerateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AITeacherExpansionPracticeTaskGenerateResponse
func (client *Client) AITeacherExpansionPracticeTaskGenerateWithContext(ctx context.Context, request *AITeacherExpansionPracticeTaskGenerateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AITeacherExpansionPracticeTaskGenerateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.KeySentences) {
		body["keySentences"] = request.KeySentences
	}

	if !dara.IsNil(request.KeyWords) {
		body["keyWords"] = request.KeyWords
	}

	if !dara.IsNil(request.LearningObject) {
		body["learningObject"] = request.LearningObject
	}

	if !dara.IsNil(request.TextContent) {
		body["textContent"] = request.TextContent
	}

	if !dara.IsNil(request.Textbook) {
		body["textbook"] = request.Textbook
	}

	if !dara.IsNil(request.Topic) {
		body["topic"] = request.Topic
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AITeacherExpansionPracticeTaskGenerate"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/expansionPractice/generateTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AITeacherExpansionPracticeTaskGenerateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronous basic practice is primarily for dialogue tasks with a ground truth. Although this mode allows some deviation from the ground truth, the AI strictly requires users to follow it.
//
// @param request - AITeacherSyncPracticeTaskGenerateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AITeacherSyncPracticeTaskGenerateResponse
func (client *Client) AITeacherSyncPracticeTaskGenerateWithContext(ctx context.Context, request *AITeacherSyncPracticeTaskGenerateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AITeacherSyncPracticeTaskGenerateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.KeySentences) {
		body["keySentences"] = request.KeySentences
	}

	if !dara.IsNil(request.KeyWords) {
		body["keyWords"] = request.KeyWords
	}

	if !dara.IsNil(request.LearningObject) {
		body["learningObject"] = request.LearningObject
	}

	if !dara.IsNil(request.TextContent) {
		body["textContent"] = request.TextContent
	}

	if !dara.IsNil(request.Textbook) {
		body["textbook"] = request.Textbook
	}

	if !dara.IsNil(request.Topic) {
		body["topic"] = request.Topic
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AITeacherSyncPracticeTaskGenerate"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/syncPractice/generateTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AITeacherSyncPracticeTaskGenerateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the services available on the Alibaba Cloud Console.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("AliyunConsoleOpenApiQueryAliyunConsoleServcieList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunconsole/queryAliyunConsoleServcieList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the services available in the Alibaba Cloud console.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServiceListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("AliyunConsoleOpenApiQueryAliyunConsoleServiceList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/queryAliyunConsoleServiceList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud Management Console / List purchased resources
//
// @param request - AliyunConsoleOpenApiQueryPaidResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AliyunConsoleOpenApiQueryPaidResourceResponse
func (client *Client) AliyunConsoleOpenApiQueryPaidResourceWithContext(ctx context.Context, request *AliyunConsoleOpenApiQueryPaidResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AliyunConsoleOpenApiQueryPaidResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AliyunConsoleOpenApiQueryPaidResource"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/queryPaidResource"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AliyunConsoleOpenApiQueryPaidResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Intelligent Correction / Oral Evaluation / Statistics / call volume
//
// @param request - CountOralEvaluationStatisticsCallsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOralEvaluationStatisticsCallsResponse
func (client *Client) CountOralEvaluationStatisticsCallsWithContext(ctx context.Context, request *CountOralEvaluationStatisticsCallsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CountOralEvaluationStatisticsCallsResponse, _err error) {
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
		Action:      dara.String("CountOralEvaluationStatisticsCalls"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/countOralEvaluationStatisticsCalls"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CountOralEvaluationStatisticsCallsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Intelligent grading / oral evaluation / statistics / concurrency
//
// @param request - CountOralEvaluationStatisticsConcurrentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOralEvaluationStatisticsConcurrentResponse
func (client *Client) CountOralEvaluationStatisticsConcurrentWithContext(ctx context.Context, request *CountOralEvaluationStatisticsConcurrentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CountOralEvaluationStatisticsConcurrentResponse, _err error) {
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
		Action:      dara.String("CountOralEvaluationStatisticsConcurrent"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/countOralEvaluationStatisticsConcurrent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CountOralEvaluationStatisticsConcurrentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves statistics about API call errors for the oral evaluation service.
//
// @param request - CountOralEvaluationStatisticsErrorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOralEvaluationStatisticsErrorResponse
func (client *Client) CountOralEvaluationStatisticsErrorWithContext(ctx context.Context, request *CountOralEvaluationStatisticsErrorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CountOralEvaluationStatisticsErrorResponse, _err error) {
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
		Action:      dara.String("CountOralEvaluationStatisticsError"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/countOralEvaluationStatisticsError"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CountOralEvaluationStatisticsErrorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access warrant.
//
// @param request - CreateAccessWarrantRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessWarrantResponse
func (client *Client) CreateAccessWarrantWithContext(ctx context.Context, request *CreateAccessWarrantRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAccessWarrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.RequestSign) {
		body["requestSign"] = request.RequestSign
	}

	if !dara.IsNil(request.Timestamp) {
		body["timestamp"] = request.Timestamp
	}

	if !dara.IsNil(request.UserClientIp) {
		body["userClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	if !dara.IsNil(request.WarrantAvailable) {
		body["warrantAvailable"] = request.WarrantAvailable
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessWarrant"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/createAccessWarrant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessWarrantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud console > Create Project
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		body["projectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ProjectType) {
		body["projectType"] = request.ProjectType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/createProject"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a workflow for Chinese composition tutoring.
//
// @param request - ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherChineseCompositionTutoringWorkflowRunWithSSECtx(ctx context.Context, request *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse, _yieldErr chan error) {
	defer close(_yield)
	client.executeAITeacherChineseCompositionTutoringWorkflowRunWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Executes a workflow for Chinese composition tutoring.
//
// @param request - ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherChineseCompositionTutoringWorkflowRunWithContext(ctx context.Context, request *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EssayOutline) {
		body["essayOutline"] = request.EssayOutline
	}

	if !dara.IsNil(request.EssayRequirements) {
		body["essayRequirements"] = request.EssayRequirements
	}

	if !dara.IsNil(request.EssayTopic) {
		body["essayTopic"] = request.EssayTopic
	}

	if !dara.IsNil(request.EssayType) {
		body["essayType"] = request.EssayType
	}

	if !dara.IsNil(request.EssayWordCount) {
		body["essayWordCount"] = request.EssayWordCount
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.ResponseMode) {
		body["responseMode"] = request.ResponseMode
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherChineseCompositionTutoringWorkflowRun"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/intelligentAgent/chineseCompositionTutoring/workflowRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # English Composition Tutoring
//
// @param request - ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherEnglishCompositionTutoringWorkflowRunWithSSECtx(ctx context.Context, request *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse, _yieldErr chan error) {
	defer close(_yield)
	client.executeAITeacherEnglishCompositionTutoringWorkflowRunWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// # English Composition Tutoring
//
// @param request - ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherEnglishCompositionTutoringWorkflowRunWithContext(ctx context.Context, request *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EssayOutline) {
		body["essayOutline"] = request.EssayOutline
	}

	if !dara.IsNil(request.EssayRequirements) {
		body["essayRequirements"] = request.EssayRequirements
	}

	if !dara.IsNil(request.EssayTopic) {
		body["essayTopic"] = request.EssayTopic
	}

	if !dara.IsNil(request.EssayType) {
		body["essayType"] = request.EssayType
	}

	if !dara.IsNil(request.EssayWordCount) {
		body["essayWordCount"] = request.EssayWordCount
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.ResponseMode) {
		body["responseMode"] = request.ResponseMode
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherEnglishCompositionTutoringWorkflowRun"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/intelligentAgent/englishCompositionTutoring/workflowRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Answers English-related questions.
//
// @param request - ExecuteAITeacherEnglishParaphraseChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherEnglishParaphraseChatMessageResponse
func (client *Client) ExecuteAITeacherEnglishParaphraseChatMessageWithSSECtx(ctx context.Context, request *ExecuteAITeacherEnglishParaphraseChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ExecuteAITeacherEnglishParaphraseChatMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.executeAITeacherEnglishParaphraseChatMessageWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Answers English-related questions.
//
// @param request - ExecuteAITeacherEnglishParaphraseChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherEnglishParaphraseChatMessageResponse
func (client *Client) ExecuteAITeacherEnglishParaphraseChatMessageWithContext(ctx context.Context, request *ExecuteAITeacherEnglishParaphraseChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherEnglishParaphraseChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.QuestionId) {
		body["questionId"] = request.QuestionId
	}

	if !dara.IsNil(request.QuestionInfo) {
		body["questionInfo"] = request.QuestionInfo
	}

	if !dara.IsNil(request.ResponseMode) {
		body["responseMode"] = request.ResponseMode
	}

	if !dara.IsNil(request.UserAnswer) {
		body["userAnswer"] = request.UserAnswer
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherEnglishParaphraseChatMessage"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/intelligentAgent/englishParaphrase/chatMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherEnglishParaphraseChatMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Expansion dialogues are for open-ended conversations. In these conversations, the AI poses open-ended questions, but the user must stay on topic. If a user\\"s response is off-topic, the AI steers the conversation back on topic. If the user gives two consecutive off-topic responses, the AI moves on to the next topic.
//
// @param request - ExecuteAITeacherExpansionDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueResponse
func (client *Client) ExecuteAITeacherExpansionDialogueWithContext(ctx context.Context, request *ExecuteAITeacherExpansionDialogueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Background) {
		body["background"] = request.Background
	}

	if !dara.IsNil(request.DialogueTasks) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !dara.IsNil(request.LanguageCode) {
		body["languageCode"] = request.LanguageCode
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	if !dara.IsNil(request.RoleInfo) {
		body["roleInfo"] = request.RoleInfo
	}

	if !dara.IsNil(request.StartSentence) {
		body["startSentence"] = request.StartSentence
	}

	if !dara.IsNil(request.Topic) {
		body["topic"] = request.Topic
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherExpansionDialogue"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/expansionPractice/executeExpansionTraining"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses context to polish the expanded text.
//
// @param request - ExecuteAITeacherExpansionDialogueRefineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueRefineResponse
func (client *Client) ExecuteAITeacherExpansionDialogueRefineWithContext(ctx context.Context, request *ExecuteAITeacherExpansionDialogueRefineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueRefineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Background) {
		body["background"] = request.Background
	}

	if !dara.IsNil(request.DialogueTasks) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !dara.IsNil(request.LanguageCode) {
		body["languageCode"] = request.LanguageCode
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	if !dara.IsNil(request.RoleInfo) {
		body["roleInfo"] = request.RoleInfo
	}

	if !dara.IsNil(request.StartSentence) {
		body["startSentence"] = request.StartSentence
	}

	if !dara.IsNil(request.Topic) {
		body["topic"] = request.Topic
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherExpansionDialogueRefine"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/expansionPractice/refineByContext"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueRefineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Further Contextual Translation Practice.
//
// @param request - ExecuteAITeacherExpansionDialogueTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueTranslateResponse
func (client *Client) ExecuteAITeacherExpansionDialogueTranslateWithContext(ctx context.Context, request *ExecuteAITeacherExpansionDialogueTranslateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Background) {
		body["background"] = request.Background
	}

	if !dara.IsNil(request.DialogueTasks) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	if !dara.IsNil(request.RoleInfo) {
		body["roleInfo"] = request.RoleInfo
	}

	if !dara.IsNil(request.StartSentence) {
		body["startSentence"] = request.StartSentence
	}

	if !dara.IsNil(request.Topic) {
		body["topic"] = request.Topic
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherExpansionDialogueTranslate"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/expansionPractice/translate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a grammar check.
//
// @param request - ExecuteAITeacherGrammarCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherGrammarCheckResponse
func (client *Client) ExecuteAITeacherGrammarCheckWithContext(ctx context.Context, request *ExecuteAITeacherGrammarCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherGrammarCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherGrammarCheck"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/common/grammarChecking"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherGrammarCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Practice synchronous dialogue.
//
// @param request - ExecuteAITeacherSyncDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherSyncDialogueResponse
func (client *Client) ExecuteAITeacherSyncDialogueWithContext(ctx context.Context, request *ExecuteAITeacherSyncDialogueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherSyncDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DialogueTasks) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !dara.IsNil(request.LanguageCode) {
		body["languageCode"] = request.LanguageCode
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherSyncDialogue"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/syncPractice/executeSyncTraining"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherSyncDialogueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can practice contextual translation in real-time.
//
// @param request - ExecuteAITeacherSyncDialogueTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherSyncDialogueTranslateResponse
func (client *Client) ExecuteAITeacherSyncDialogueTranslateWithContext(ctx context.Context, request *ExecuteAITeacherSyncDialogueTranslateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAITeacherSyncDialogueTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DialogueTasks) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherSyncDialogueTranslate"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/syncPractice/translate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAITeacherSyncDialogueTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a dialogue turn with the Textbook Assistant.
//
// @param request - ExecuteTextbookAssistantDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantDialogueResponse
func (client *Client) ExecuteTextbookAssistantDialogueWithContext(ctx context.Context, request *ExecuteTextbookAssistantDialogueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	if !dara.IsNil(request.UserMessage) {
		body["userMessage"] = request.UserMessage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantDialogue"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/ExecuteDialogue"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantDialogueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adjusts the difficulty of the textbook assistant\\"s dialogue.
//
// @param request - ExecuteTextbookAssistantDifficultyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantDifficultyResponse
func (client *Client) ExecuteTextbookAssistantDifficultyWithContext(ctx context.Context, request *ExecuteTextbookAssistantDifficultyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantDifficultyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.Assistant) {
		body["assistant"] = request.Assistant
	}

	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantDifficulty"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/ExecuteDifficulty"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantDifficultyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a grammar check.
//
// @param request - ExecuteTextbookAssistantGrammarCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantGrammarCheckResponse
func (client *Client) ExecuteTextbookAssistantGrammarCheckWithContext(ctx context.Context, request *ExecuteTextbookAssistantGrammarCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantGrammarCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	if !dara.IsNil(request.User) {
		body["user"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantGrammarCheck"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/ExecuteGrammarCheck"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantGrammarCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refines a sentence based on the conversational context.
//
// @param request - ExecuteTextbookAssistantRefineByContextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantRefineByContextResponse
func (client *Client) ExecuteTextbookAssistantRefineByContextWithContext(ctx context.Context, request *ExecuteTextbookAssistantRefineByContextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantRefineByContextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	if !dara.IsNil(request.User) {
		body["user"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantRefineByContext"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/RefineByContext"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantRefineByContextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation retries a conversation.
//
// @param request - ExecuteTextbookAssistantRetryConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantRetryConversationResponse
func (client *Client) ExecuteTextbookAssistantRetryConversationWithContext(ctx context.Context, request *ExecuteTextbookAssistantRetryConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantRetryConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Assistant) {
		body["assistant"] = request.Assistant
	}

	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantRetryConversation"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/RetryConversation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantRetryConversationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a conversation and returns a streaming output.
//
// @param request - ExecuteTextbookAssistantSseDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantSseDialogueResponse
func (client *Client) ExecuteTextbookAssistantSseDialogueWithSSECtx(ctx context.Context, request *ExecuteTextbookAssistantSseDialogueRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ExecuteTextbookAssistantSseDialogueResponse, _yieldErr chan error) {
	defer close(_yield)
	client.executeTextbookAssistantSseDialogueWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Starts a conversation and returns a streaming output.
//
// @param request - ExecuteTextbookAssistantSseDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantSseDialogueResponse
func (client *Client) ExecuteTextbookAssistantSseDialogueWithContext(ctx context.Context, request *ExecuteTextbookAssistantSseDialogueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantSseDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	if !dara.IsNil(request.UserMessage) {
		body["userMessage"] = request.UserMessage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantSseDialogue"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/ExecuteSseDialogue"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantSseDialogueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a conversation with the AI teacher. The teacher then sends the initial message.
//
// @param request - ExecuteTextbookAssistantStartConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantStartConversationResponse
func (client *Client) ExecuteTextbookAssistantStartConversationWithContext(ctx context.Context, request *ExecuteTextbookAssistantStartConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantStartConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticleId) {
		body["articleId"] = request.ArticleId
	}

	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantStartConversation"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/StartConversation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantStartConversationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a suggested response from the textbook-based AI teacher.
//
// @param request - ExecuteTextbookAssistantSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantSuggestionResponse
func (client *Client) ExecuteTextbookAssistantSuggestionWithContext(ctx context.Context, request *ExecuteTextbookAssistantSuggestionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantSuggestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Assistant) {
		body["assistant"] = request.Assistant
	}

	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantSuggestion"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/Suggestion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantSuggestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Translates the content of a message.
//
// @param request - ExecuteTextbookAssistantTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantTranslateResponse
func (client *Client) ExecuteTextbookAssistantTranslateWithContext(ctx context.Context, request *ExecuteTextbookAssistantTranslateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteTextbookAssistantTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Assistant) {
		body["assistant"] = request.Assistant
	}

	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantTranslate"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/ExecuteTranslate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTextbookAssistantTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Supplemental Practice Assistant
//
// @param request - GetAITeacherExpansionDialogueSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITeacherExpansionDialogueSuggestionResponse
func (client *Client) GetAITeacherExpansionDialogueSuggestionWithContext(ctx context.Context, request *GetAITeacherExpansionDialogueSuggestionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAITeacherExpansionDialogueSuggestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Background) {
		body["background"] = request.Background
	}

	if !dara.IsNil(request.DialogueTasks) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !dara.IsNil(request.LanguageCode) {
		body["languageCode"] = request.LanguageCode
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	if !dara.IsNil(request.RoleInfo) {
		body["roleInfo"] = request.RoleInfo
	}

	if !dara.IsNil(request.StartSentence) {
		body["startSentence"] = request.StartSentence
	}

	if !dara.IsNil(request.Topic) {
		body["topic"] = request.Topic
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAITeacherExpansionDialogueSuggestion"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/expansionPractice/suggestion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAITeacherExpansionDialogueSuggestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Sync Practice Assistant
//
// @param request - GetAITeacherSyncDialogueSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITeacherSyncDialogueSuggestionResponse
func (client *Client) GetAITeacherSyncDialogueSuggestionWithContext(ctx context.Context, request *GetAITeacherSyncDialogueSuggestionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAITeacherSyncDialogueSuggestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DialogueTasks) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !dara.IsNil(request.LanguageCode) {
		body["languageCode"] = request.LanguageCode
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAITeacherSyncDialogueSuggestion"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aiteacher/syncPractice/suggestion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAITeacherSyncDialogueSuggestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an authorization token to make API calls.
//
// @param request - GetTextbookAssistantTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextbookAssistantTokenResponse
func (client *Client) GetTextbookAssistantTokenWithContext(ctx context.Context, request *GetTextbookAssistantTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextbookAssistantTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		body["deviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextbookAssistantToken"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/teachingResource/GetToken"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextbookAssistantTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch get article details
//
// @param request - ListTextbookAssistantArticleDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantArticleDetailsResponse
func (client *Client) ListTextbookAssistantArticleDetailsWithContext(ctx context.Context, request *ListTextbookAssistantArticleDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextbookAssistantArticleDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ArticleIdList) {
		body["articleIdList"] = request.ArticleIdList
	}

	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTextbookAssistantArticleDetails"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/teachingResource/ListArticleDetails"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextbookAssistantArticleDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a list of articles.
//
// @param request - ListTextbookAssistantArticlesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantArticlesResponse
func (client *Client) ListTextbookAssistantArticlesWithContext(ctx context.Context, request *ListTextbookAssistantArticlesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextbookAssistantArticlesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.DirectoryId) {
		body["directoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTextbookAssistantArticles"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/teachingResource/ListArticles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextbookAssistantArticlesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the table of contents of a book.
//
// @param request - ListTextbookAssistantBookDirectoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantBookDirectoriesResponse
func (client *Client) ListTextbookAssistantBookDirectoriesWithContext(ctx context.Context, request *ListTextbookAssistantBookDirectoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextbookAssistantBookDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.BookId) {
		body["bookId"] = request.BookId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTextbookAssistantBookDirectories"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/teachingResource/ListBookDirectories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextbookAssistantBookDirectoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of books for a specified grade.
//
// @param request - ListTextbookAssistantBooksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantBooksResponse
func (client *Client) ListTextbookAssistantBooksWithContext(ctx context.Context, request *ListTextbookAssistantBooksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextbookAssistantBooksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.BookId) {
		body["bookId"] = request.BookId
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	if !dara.IsNil(request.Volume) {
		body["volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTextbookAssistantBooks"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/teachingResource/ListBooks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextbookAssistantBooksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the available grades and volumes for the Textbook Assistant.
//
// @param request - ListTextbookAssistantGradeVolumesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantGradeVolumesResponse
func (client *Client) ListTextbookAssistantGradeVolumesWithContext(ctx context.Context, request *ListTextbookAssistantGradeVolumesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextbookAssistantGradeVolumesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTextbookAssistantGradeVolumes"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/teachingResource/ListGradeVolumes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextbookAssistantGradeVolumesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Article Details
//
// @param request - ListTextbookAssistantSceneDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantSceneDetailsResponse
func (client *Client) ListTextbookAssistantSceneDetailsWithContext(ctx context.Context, request *ListTextbookAssistantSceneDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextbookAssistantSceneDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.SceneIdList) {
		body["sceneIdList"] = request.SceneIdList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTextbookAssistantSceneDetails"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/teachingResource/ListSceneDetails"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextbookAssistantSceneDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds model groups to departments in batches.
//
// Description:
//
// Binds model groups to departments in batches.
//
// @param request - ModelRouterBatchBindModelGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterBatchBindModelGroupResponse
func (client *Client) ModelRouterBatchBindModelGroupWithContext(ctx context.Context, request *ModelRouterBatchBindModelGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterBatchBindModelGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedModelGroupConfig) {
		body["allowedModelGroupConfig"] = request.AllowedModelGroupConfig
	}

	if !dara.IsNil(request.ClientIdList) {
		body["clientIdList"] = request.ClientIdList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterBatchBindModelGroup"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/batch-bind-model-group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterBatchBindModelGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates API keys in batches for members under a department in organization management.
//
// @param request - ModelRouterBatchCreateMemberApiKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterBatchCreateMemberApiKeysResponse
func (client *Client) ModelRouterBatchCreateMemberApiKeysWithContext(ctx context.Context, id *string, request *ModelRouterBatchCreateMemberApiKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterBatchCreateMemberApiKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireAt) {
		body["expireAt"] = request.ExpireAt
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.UserIds) {
		body["userIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterBatchCreateMemberApiKeys"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/member-apikeys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterBatchCreateMemberApiKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages models by performing batch model creation.
//
// @param request - ModelRouterBatchCreateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterBatchCreateModelResponse
func (client *Client) ModelRouterBatchCreateModelWithContext(ctx context.Context, request *ModelRouterBatchCreateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterBatchCreateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BaseUrl) {
		body["baseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.Models) {
		body["models"] = request.Models
	}

	if !dara.IsNil(request.Symbol) {
		body["symbol"] = request.Symbol
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterBatchCreateModel"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/models/batch"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterBatchCreateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables API keys in batches for members under a department in organization management.
//
// @param request - ModelRouterBatchDisableMemberApiKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterBatchDisableMemberApiKeysResponse
func (client *Client) ModelRouterBatchDisableMemberApiKeysWithContext(ctx context.Context, id *string, request *ModelRouterBatchDisableMemberApiKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterBatchDisableMemberApiKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIds) {
		body["userIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterBatchDisableMemberApiKeys"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/member-apikeys/disable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterBatchDisableMemberApiKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch resets member authorizations to inherit under a department in organization management.
//
// @param request - ModelRouterBatchResetMemberAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterBatchResetMemberAuthorizationResponse
func (client *Client) ModelRouterBatchResetMemberAuthorizationWithContext(ctx context.Context, id *string, request *ModelRouterBatchResetMemberAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterBatchResetMemberAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIds) {
		body["userIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterBatchResetMemberAuthorization"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/member-authorizations/reset"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterBatchResetMemberAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch sets member authorization under a department.
//
// @param request - ModelRouterBatchSetMemberAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterBatchSetMemberAuthorizationResponse
func (client *Client) ModelRouterBatchSetMemberAuthorizationWithContext(ctx context.Context, id *string, request *ModelRouterBatchSetMemberAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterBatchSetMemberAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedModelGroupConfig) {
		body["allowedModelGroupConfig"] = request.AllowedModelGroupConfig
	}

	if !dara.IsNil(request.UserIdList) {
		body["userIdList"] = request.UserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterBatchSetMemberAuthorization"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/member-authorizations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterBatchSetMemberAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the usage monitoring tab configuration.
//
// @param request - ModelRouterBillingCostTabsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterBillingCostTabsResponse
func (client *Client) ModelRouterBillingCostTabsWithContext(ctx context.Context, request *ModelRouterBillingCostTabsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterBillingCostTabsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterBillingCostTabs"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/cost/tabs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterBillingCostTabsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a chat conversation.
//
// @param request - ModelRouterChatCompletionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterChatCompletionsResponse
func (client *Client) ModelRouterChatCompletionsWithSSECtx(ctx context.Context, request *ModelRouterChatCompletionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ModelRouterChatCompletionsResponse, _yieldErr chan error) {
	defer close(_yield)
	client.modelRouterChatCompletionsWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Initiates a chat conversation.
//
// @param request - ModelRouterChatCompletionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterChatCompletionsResponse
func (client *Client) ModelRouterChatCompletionsWithContext(ctx context.Context, request *ModelRouterChatCompletionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterChatCompletionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterChatCompletions"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterChatCompletionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables balance-based throttling for a department.
//
// @param request - ModelRouterConfigureClientBalanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterConfigureClientBalanceResponse
func (client *Client) ModelRouterConfigureClientBalanceWithContext(ctx context.Context, id *string, request *ModelRouterConfigureClientBalanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterConfigureClientBalanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.EnableBalance) {
		body["enableBalance"] = request.EnableBalance
	}

	if !dara.IsNil(request.InitialBalance) {
		body["initialBalance"] = request.InitialBalance
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterConfigureClientBalance"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterConfigureClientBalanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the sub-wallet balance of a member in an organization.
//
// @param request - ModelRouterConfigureMemberBalanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterConfigureMemberBalanceResponse
func (client *Client) ModelRouterConfigureMemberBalanceWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterConfigureMemberBalanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterConfigureMemberBalanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.EnableBalance) {
		body["enableBalance"] = request.EnableBalance
	}

	if !dara.IsNil(request.InitialBalance) {
		body["initialBalance"] = request.InitialBalance
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterConfigureMemberBalance"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterConfigureMemberBalanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies an API key.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCopyApiKeyResponse
func (client *Client) ModelRouterCopyApiKeyWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCopyApiKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCopyApiKey"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/apikeys/" + dara.PercentEncode(dara.StringValue(id)) + "/copy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCopyApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API key.
//
// @param request - ModelRouterCreateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateApiKeyResponse
func (client *Client) ModelRouterCreateApiKeyWithContext(ctx context.Context, request *ModelRouterCreateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateApiKey"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/apikeys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages customers or creates a balance transaction.
//
// @param request - ModelRouterCreateBalanceTransactionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateBalanceTransactionResponse
func (client *Client) ModelRouterCreateBalanceTransactionWithContext(ctx context.Context, id *string, request *ModelRouterCreateBalanceTransactionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateBalanceTransactionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["amount"] = request.Amount
	}

	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.IdempotencyKey) {
		body["idempotencyKey"] = request.IdempotencyKey
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateBalanceTransaction"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/transactions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateBalanceTransactionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a billing rule.
//
// @param request - ModelRouterCreateBillingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateBillingRuleResponse
func (client *Client) ModelRouterCreateBillingRuleWithContext(ctx context.Context, request *ModelRouterCreateBillingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateBillingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BillingType) {
		body["billingType"] = request.BillingType
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["effectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.ExpireTime) {
		body["expireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.PricingConfig) {
		body["pricingConfig"] = request.PricingConfig
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateBillingRule"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateBillingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a customer.
//
// @param request - ModelRouterCreateClientRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateClientResponse
func (client *Client) ModelRouterCreateClientWithContext(ctx context.Context, request *ModelRouterCreateClientRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		body["address"] = request.Address
	}

	if !dara.IsNil(request.AllowedModelGroupConfig) {
		body["allowedModelGroupConfig"] = request.AllowedModelGroupConfig
	}

	if !dara.IsNil(request.AllowedModels) {
		body["allowedModels"] = request.AllowedModels
	}

	if !dara.IsNil(request.Contact) {
		body["contact"] = request.Contact
	}

	if !dara.IsNil(request.Discount) {
		body["discount"] = request.Discount
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateClient"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a conversation.
//
// @param request - ModelRouterCreateConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateConversationResponse
func (client *Client) ModelRouterCreateConversationWithContext(ctx context.Context, request *ModelRouterCreateConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatData) {
		body["chatData"] = request.ChatData
	}

	if !dara.IsNil(request.ModelIds) {
		body["modelIds"] = request.ModelIds
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateConversation"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/conversations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateConversationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API key for a member in the organization.
//
// @param request - ModelRouterCreateMemberApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateMemberApiKeyResponse
func (client *Client) ModelRouterCreateMemberApiKeyWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterCreateMemberApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateMemberApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireAt) {
		body["expireAt"] = request.ExpireAt
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateMemberApiKey"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/apikeys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateMemberApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a balance transaction on a member sub-wallet in organization management.
//
// @param request - ModelRouterCreateMemberBalanceTransactionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateMemberBalanceTransactionResponse
func (client *Client) ModelRouterCreateMemberBalanceTransactionWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterCreateMemberBalanceTransactionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateMemberBalanceTransactionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["amount"] = request.Amount
	}

	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.IdempotencyKey) {
		body["idempotencyKey"] = request.IdempotencyKey
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateMemberBalanceTransaction"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/transactions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateMemberBalanceTransactionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a balance subscription for a member in an organization.
//
// @param request - ModelRouterCreateMemberSubscriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateMemberSubscriptionResponse
func (client *Client) ModelRouterCreateMemberSubscriptionWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterCreateMemberSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateMemberSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["amount"] = request.Amount
	}

	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["effectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.IdempotencyKey) {
		body["idempotencyKey"] = request.IdempotencyKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateMemberSubscription"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/subscription"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateMemberSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs model creation.
//
// @param request - ModelRouterCreateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateModelResponse
func (client *Client) ModelRouterCreateModelWithContext(ctx context.Context, request *ModelRouterCreateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BaseUrl) {
		body["baseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Extensions) {
		body["extensions"] = request.Extensions
	}

	if !dara.IsNil(request.InOut) {
		body["inOut"] = request.InOut
	}

	if !dara.IsNil(request.MaxInputLength) {
		body["maxInputLength"] = request.MaxInputLength
	}

	if !dara.IsNil(request.MaxOutputLength) {
		body["maxOutputLength"] = request.MaxOutputLength
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelType) {
		body["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Symbol) {
		body["symbol"] = request.Symbol
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateModel"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/models"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a manual model group.
//
// Description:
//
// Creates a manual model group.
//
// @param request - ModelRouterCreateModelGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateModelGroupResponse
func (client *Client) ModelRouterCreateModelGroupWithContext(ctx context.Context, request *ModelRouterCreateModelGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateModelGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelList) {
		body["modelList"] = request.ModelList
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateModelGroup"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateModelGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a periodic recharge subscription for customer management.
//
// @param request - ModelRouterCreateSubscriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateSubscriptionResponse
func (client *Client) ModelRouterCreateSubscriptionWithContext(ctx context.Context, id *string, request *ModelRouterCreateSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["effectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.IdempotencyKey) {
		body["idempotencyKey"] = request.IdempotencyKey
	}

	if !dara.IsNil(request.SubscriptionAmount) {
		body["subscriptionAmount"] = request.SubscriptionAmount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateSubscription"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/subscription"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user.
//
// Description:
//
// Creates a user.
//
// @param request - ModelRouterCreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterCreateUserResponse
func (client *Client) ModelRouterCreateUserWithContext(ctx context.Context, request *ModelRouterCreateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterCreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentRoles) {
		body["departmentRoles"] = request.DepartmentRoles
	}

	if !dara.IsNil(request.LoginName) {
		body["loginName"] = request.LoginName
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterCreateUser"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/users"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterCreateUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API key.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterDeleteApiKeyResponse
func (client *Client) ModelRouterDeleteApiKeyWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterDeleteApiKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterDeleteApiKey"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/apikeys/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterDeleteApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a customer.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterDeleteClientResponse
func (client *Client) ModelRouterDeleteClientWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterDeleteClientResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterDeleteClient"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterDeleteClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a conversation.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterDeleteConversationResponse
func (client *Client) ModelRouterDeleteConversationWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterDeleteConversationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterDeleteConversation"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/conversations/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterDeleteConversationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterDeleteModelResponse
func (client *Client) ModelRouterDeleteModelWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterDeleteModelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterDeleteModel"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/models/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterDeleteModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a manual group.
//
// Description:
//
// Deletes a manual group.
//
// @param request - ModelRouterDeleteModelGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterDeleteModelGroupResponse
func (client *Client) ModelRouterDeleteModelGroupWithContext(ctx context.Context, groupId *string, request *ModelRouterDeleteModelGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterDeleteModelGroupResponse, _err error) {
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
		Action:      dara.String("ModelRouterDeleteModelGroup"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterDeleteModelGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user.
//
// @param request - ModelRouterDeleteUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterDeleteUserResponse
func (client *Client) ModelRouterDeleteUserWithContext(ctx context.Context, id *string, request *ModelRouterDeleteUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterDeleteUserResponse, _err error) {
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
		Action:      dara.String("ModelRouterDeleteUser"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/users/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterDeleteUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports the balance change records of a member in the organization.
//
// @param request - ModelRouterExportMemberBalanceOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterExportMemberBalanceOrdersResponse
func (client *Client) ModelRouterExportMemberBalanceOrdersWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterExportMemberBalanceOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterExportMemberBalanceOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		query["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterExportMemberBalanceOrders"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/orders/export"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterExportMemberBalanceOrdersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Billing Center/Queries the total cost trend of bills.
//
// Description:
//
// Queries user role assignments.
//
// @param request - ModelRouterGetBillingBillSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetBillingBillSummaryResponse
func (client *Client) ModelRouterGetBillingBillSummaryWithContext(ctx context.Context, request *ModelRouterGetBillingBillSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetBillingBillSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelTypes) {
		query["modelTypes"] = request.ModelTypes
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterGetBillingBillSummary"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/bills/summary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetBillingBillSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Client Management/Get department balance
//
// @param request - ModelRouterGetClientBalanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetClientBalanceResponse
func (client *Client) ModelRouterGetClientBalanceWithContext(ctx context.Context, id *string, request *ModelRouterGetClientBalanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetClientBalanceResponse, _err error) {
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
		Action:      dara.String("ModelRouterGetClientBalance"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetClientBalanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the balance change logs of a department.
//
// @param request - ModelRouterGetClientBalanceLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetClientBalanceLogsResponse
func (client *Client) ModelRouterGetClientBalanceLogsWithContext(ctx context.Context, id *string, request *ModelRouterGetClientBalanceLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetClientBalanceLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeType) {
		query["changeType"] = request.ChangeType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterGetClientBalanceLogs"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetClientBalanceLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the balance overview of a department.
//
// @param request - ModelRouterGetDeptBalanceSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetDeptBalanceSummaryResponse
func (client *Client) ModelRouterGetDeptBalanceSummaryWithContext(ctx context.Context, id *string, request *ModelRouterGetDeptBalanceSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetDeptBalanceSummaryResponse, _err error) {
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
		Action:      dara.String("ModelRouterGetDeptBalanceSummary"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance-summary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetDeptBalanceSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of API keys for a member in the organization.
//
// @param request - ModelRouterGetMemberApiKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetMemberApiKeysResponse
func (client *Client) ModelRouterGetMemberApiKeysWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterGetMemberApiKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetMemberApiKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterGetMemberApiKeys"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/apikeys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetMemberApiKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the sub-wallet balance of a member in an organization.
//
// @param request - ModelRouterGetMemberBalanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetMemberBalanceResponse
func (client *Client) ModelRouterGetMemberBalanceWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterGetMemberBalanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetMemberBalanceResponse, _err error) {
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
		Action:      dara.String("ModelRouterGetMemberBalance"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetMemberBalanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the balance change logs of a member in an organization.
//
// @param request - ModelRouterGetMemberBalanceLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetMemberBalanceLogsResponse
func (client *Client) ModelRouterGetMemberBalanceLogsWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterGetMemberBalanceLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetMemberBalanceLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeType) {
		query["changeType"] = request.ChangeType
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.SkipTotal) {
		query["skipTotal"] = request.SkipTotal
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterGetMemberBalanceLogs"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetMemberBalanceLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cross-department role assignments of a user.
//
// Description:
//
// Queries the role assignments of a user.
//
// @param request - ModelRouterGetUserRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterGetUserRolesResponse
func (client *Client) ModelRouterGetUserRolesWithContext(ctx context.Context, id *string, request *ModelRouterGetUserRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterGetUserRolesResponse, _err error) {
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
		Action:      dara.String("ModelRouterGetUserRoles"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/users/" + dara.PercentEncode(dara.StringValue(id)) + "/roles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterGetUserRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries balance change records.
//
// Description:
//
// This API operation is deprecated. Do not use it.
//
// @param request - ModelRouterListBalanceOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterListBalanceOrdersResponse
func (client *Client) ModelRouterListBalanceOrdersWithContext(ctx context.Context, id *string, request *ModelRouterListBalanceOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterListBalanceOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		query["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderType) {
		query["orderType"] = request.OrderType
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterListBalanceOrders"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/orders"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterListBalanceOrdersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of members in a specified department.
//
// @param request - ModelRouterListDeptMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterListDeptMembersResponse
func (client *Client) ModelRouterListDeptMembersWithContext(ctx context.Context, id *string, request *ModelRouterListDeptMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterListDeptMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfig) {
		query["authConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.IncludeAuthorization) {
		query["includeAuthorization"] = request.IncludeAuthorization
	}

	if !dara.IsNil(request.IncludeBalance) {
		query["includeBalance"] = request.IncludeBalance
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Model) {
		query["model"] = request.Model
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterListDeptMembers"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/members"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterListDeptMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the balance change records of a member in the organization.
//
// @param request - ModelRouterListMemberBalanceOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterListMemberBalanceOrdersResponse
func (client *Client) ModelRouterListMemberBalanceOrdersWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterListMemberBalanceOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterListMemberBalanceOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		query["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.OrderType) {
		query["orderType"] = request.OrderType
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterListMemberBalanceOrders"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/orders"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterListMemberBalanceOrdersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the subscription list of a member in the organization.
//
// @param request - ModelRouterListMemberSubscriptionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterListMemberSubscriptionsResponse
func (client *Client) ModelRouterListMemberSubscriptionsWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterListMemberSubscriptionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterListMemberSubscriptionsResponse, _err error) {
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
		Action:      dara.String("ModelRouterListMemberSubscriptions"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/subscription"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterListMemberSubscriptionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of periodic recharge subscriptions.
//
// Description:
//
// This operation is deprecated. Do not use it.
//
// @param request - ModelRouterListSubscriptionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterListSubscriptionsResponse
func (client *Client) ModelRouterListSubscriptionsWithContext(ctx context.Context, id *string, request *ModelRouterListSubscriptionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterListSubscriptionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		query["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterListSubscriptions"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/subscription"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterListSubscriptionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a pre-signed URL for downloading a Migu source file.
//
// Description:
//
// Creates a user.
//
// @param request - ModelRouterMiguDownloadSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterMiguDownloadSourceResponse
func (client *Client) ModelRouterMiguDownloadSourceWithContext(ctx context.Context, request *ModelRouterMiguDownloadSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterMiguDownloadSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceId) {
		query["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterMiguDownloadSource"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/pipeline/api/aigc/source/download"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterMiguDownloadSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages Migu source files and retrieves a pre-signed URL for source file upload.
//
// Description:
//
// Updates a user.
//
// @param request - ModelRouterMiguUploadSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterMiguUploadSourceResponse
func (client *Client) ModelRouterMiguUploadSourceWithContext(ctx context.Context, request *ModelRouterMiguUploadSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterMiguUploadSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileType) {
		body["fileType"] = request.FileType
	}

	if !dara.IsNil(request.ServiceName) {
		body["serviceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterMiguUploadSource"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/pipeline/api/aigc/source/upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterMiguUploadSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an API key.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryApiKeyResponse
func (client *Client) ModelRouterQueryApiKeyWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryApiKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryApiKey"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/apikeys/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of API keys.
//
// @param request - ModelRouterQueryApiKeyListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryApiKeyListResponse
func (client *Client) ModelRouterQueryApiKeyListWithContext(ctx context.Context, request *ModelRouterQueryApiKeyListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryApiKeyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.IncludeMemberKeys) {
		query["includeMemberKeys"] = request.IncludeMemberKeys
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryApiKeyList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/apikeys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryApiKeyListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries billing details in batches.
//
// Description:
//
// Queries the user list.
//
// @param request - ModelRouterQueryBillingCostBreakdownRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryBillingCostBreakdownResponse
func (client *Client) ModelRouterQueryBillingCostBreakdownWithContext(ctx context.Context, request *ModelRouterQueryBillingCostBreakdownRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryBillingCostBreakdownResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["granularity"] = request.Granularity
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelTypes) {
		query["modelTypes"] = request.ModelTypes
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryBillingCostBreakdown"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/cost/breakdown"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryBillingCostBreakdownResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries request-granularity billing details from the Billing Center.
//
// Description:
//
// Queries the user list.
//
// @param request - ModelRouterQueryBillingDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryBillingDetailsResponse
func (client *Client) ModelRouterQueryBillingDetailsWithContext(ctx context.Context, request *ModelRouterQueryBillingDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryBillingDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.ModelCodes) {
		query["modelCodes"] = request.ModelCodes
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelTypes) {
		query["modelTypes"] = request.ModelTypes
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequestId) {
		query["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryBillingDetails"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/details"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryBillingDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Billing management / Query billing rule list
//
// @param request - ModelRouterQueryBillingRuleListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryBillingRuleListResponse
func (client *Client) ModelRouterQueryBillingRuleListWithContext(ctx context.Context, request *ModelRouterQueryBillingRuleListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryBillingRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveOnly) {
		query["activeOnly"] = request.ActiveOnly
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModelCode) {
		query["modelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryBillingRuleList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryBillingRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the discount modification history for a client.
//
// @param request - ModelRouterQueryClientDiscountLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryClientDiscountLogsResponse
func (client *Client) ModelRouterQueryClientDiscountLogsWithContext(ctx context.Context, id *string, request *ModelRouterQueryClientDiscountLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryClientDiscountLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryClientDiscountLogs"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/discount-logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryClientDiscountLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of clients.
//
// @param request - ModelRouterQueryClientListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryClientListResponse
func (client *Client) ModelRouterQueryClientListWithContext(ctx context.Context, request *ModelRouterQueryClientListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryClientListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryClientList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryClientListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the customer tree structure.
//
// @param request - ModelRouterQueryClientTreeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryClientTreeResponse
func (client *Client) ModelRouterQueryClientTreeWithContext(ctx context.Context, request *ModelRouterQueryClientTreeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryClientTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryClientTree"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/tree"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryClientTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a conversation.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryConversationResponse
func (client *Client) ModelRouterQueryConversationWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryConversationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryConversation"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/conversations/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryConversationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Conversation management/Conversation list
//
// @param request - ModelRouterQueryConversationListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryConversationListResponse
func (client *Client) ModelRouterQueryConversationListWithContext(ctx context.Context, request *ModelRouterQueryConversationListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryConversationListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryConversationList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/conversations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryConversationListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves model usage details for billing management.
//
// @param request - ModelRouterQueryCostModelDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryCostModelDetailResponse
func (client *Client) ModelRouterQueryCostModelDetailWithContext(ctx context.Context, request *ModelRouterQueryCostModelDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryCostModelDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryCostModelDetail"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/cost/model-detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryCostModelDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of models called for billing management.
//
// @param request - ModelRouterQueryCostModelListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryCostModelListResponse
func (client *Client) ModelRouterQueryCostModelListWithContext(ctx context.Context, request *ModelRouterQueryCostModelListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryCostModelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["granularity"] = request.Granularity
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelTypes) {
		query["modelTypes"] = request.ModelTypes
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryCostModelList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/cost/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryCostModelListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves cost overview metrics for billing management.
//
// @param request - ModelRouterQueryCostOverviewMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryCostOverviewMetricsResponse
func (client *Client) ModelRouterQueryCostOverviewMetricsWithContext(ctx context.Context, request *ModelRouterQueryCostOverviewMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryCostOverviewMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["granularity"] = request.Granularity
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelTypes) {
		query["modelTypes"] = request.ModelTypes
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryCostOverviewMetrics"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/cost/overview"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryCostOverviewMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves usage cost trends.
//
// @param request - ModelRouterQueryCostTrendMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryCostTrendMetricsResponse
func (client *Client) ModelRouterQueryCostTrendMetricsWithContext(ctx context.Context, request *ModelRouterQueryCostTrendMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryCostTrendMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["granularity"] = request.Granularity
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelTypes) {
		query["modelTypes"] = request.ModelTypes
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryCostTrendMetrics"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/cost/trend"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryCostTrendMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details for a specific model.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelResponse
func (client *Client) ModelRouterQueryModelWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryModel"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/models/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a model group.
//
// Description:
//
// Queries the details of a model group.
//
// @param request - ModelRouterQueryModelGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelGroupResponse
func (client *Client) ModelRouterQueryModelGroupWithContext(ctx context.Context, groupId *string, request *ModelRouterQueryModelGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelGroupResponse, _err error) {
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
		Action:      dara.String("ModelRouterQueryModelGroup"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the departments bound to a model group by paging.
//
// Description:
//
// Queries the departments bound to a model group by paging.
//
// @param request - ModelRouterQueryModelGroupClientsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelGroupClientsResponse
func (client *Client) ModelRouterQueryModelGroupClientsWithContext(ctx context.Context, groupId *string, request *ModelRouterQueryModelGroupClientsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelGroupClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryModelGroupClients"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/clients"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelGroupClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of model groups by paging.
//
// Description:
//
// Queries the list of model groups by paging.
//
// @param request - ModelRouterQueryModelGroupListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelGroupListResponse
func (client *Client) ModelRouterQueryModelGroupListWithContext(ctx context.Context, request *ModelRouterQueryModelGroupListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelGroupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryModelGroupList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query for models within a model group.
//
// Description:
//
// Queries models within a group with pagination.
//
// @param request - ModelRouterQueryModelGroupModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelGroupModelsResponse
func (client *Client) ModelRouterQueryModelGroupModelsWithContext(ctx context.Context, groupId *string, request *ModelRouterQueryModelGroupModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelGroupModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryModelGroupModels"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelGroupModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the users attached to a model group by paging.
//
// Description:
//
// Queries the users attached to a model group by paging.
//
// @param request - ModelRouterQueryModelGroupUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelGroupUsersResponse
func (client *Client) ModelRouterQueryModelGroupUsersWithContext(ctx context.Context, groupId *string, request *ModelRouterQueryModelGroupUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelGroupUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryModelGroupUsers"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelGroupUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bound groups and models by API key.
//
// Description:
//
// Queries the bound groups and models by API key.
//
// @param request - ModelRouterQueryModelGroupsByApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelGroupsByApiKeyResponse
func (client *Client) ModelRouterQueryModelGroupsByApiKeyWithContext(ctx context.Context, id *string, request *ModelRouterQueryModelGroupsByApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelGroupsByApiKeyResponse, _err error) {
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
		Action:      dara.String("ModelRouterQueryModelGroupsByApiKey"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/apikeys/" + dara.PercentEncode(dara.StringValue(id)) + "/model-groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelGroupsByApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Model management/Get model list
//
// @param request - ModelRouterQueryModelListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryModelListResponse
func (client *Client) ModelRouterQueryModelListWithContext(ctx context.Context, request *ModelRouterQueryModelListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryModelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryModelList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryModelListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Nacos service providers through Nacos configuration.
//
// @param request - ModelRouterQueryNacosProvidersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryNacosProvidersResponse
func (client *Client) ModelRouterQueryNacosProvidersWithContext(ctx context.Context, request *ModelRouterQueryNacosProvidersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryNacosProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryNacosProviders"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/nacos/providers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryNacosProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of tags from Nacos.
//
// @param request - ModelRouterQueryNacosTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryNacosTagsResponse
func (client *Client) ModelRouterQueryNacosTagsWithContext(ctx context.Context, request *ModelRouterQueryNacosTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryNacosTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigType) {
		query["configType"] = request.ConfigType
	}

	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryNacosTags"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/nacos/tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryNacosTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves monitoring chart data for model observation.
//
// Description:
//
// Queries a list of users.
//
// @param request - ModelRouterQueryObservationChartsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryObservationChartsResponse
func (client *Client) ModelRouterQueryObservationChartsWithContext(ctx context.Context, request *ModelRouterQueryObservationChartsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryObservationChartsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeRange) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryObservationCharts"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/observation/charts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryObservationChartsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of observation logs for model monitoring.
//
// @param request - ModelRouterQueryObservationLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryObservationLogsResponse
func (client *Client) ModelRouterQueryObservationLogsWithContext(ctx context.Context, request *ModelRouterQueryObservationLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryObservationLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeRange) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryObservationLogs"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/observation/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryObservationLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves observability metric data for model API calls.
//
// @param request - ModelRouterQueryObservationMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryObservationMetricsResponse
func (client *Client) ModelRouterQueryObservationMetricsWithContext(ctx context.Context, request *ModelRouterQueryObservationMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryObservationMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.NeedTotalCount) {
		query["needTotalCount"] = request.NeedTotalCount
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeRange) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryObservationMetrics"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/observation/metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryObservationMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries usage details in batches for usage management.
//
// @param request - ModelRouterQueryUsageBreakdownRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryUsageBreakdownResponse
func (client *Client) ModelRouterQueryUsageBreakdownWithContext(ctx context.Context, request *ModelRouterQueryUsageBreakdownRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryUsageBreakdownResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIds) {
		query["clientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["granularity"] = request.Granularity
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberUserIds) {
		query["memberUserIds"] = request.MemberUserIds
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryUsageBreakdown"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/usage/breakdown"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryUsageBreakdownResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of users.
//
// Description:
//
// Queries the list of users.
//
// @param request - ModelRouterQueryUserListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterQueryUserListResponse
func (client *Client) ModelRouterQueryUserListWithContext(ctx context.Context, request *ModelRouterQueryUserListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterQueryUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Phone) {
		query["phone"] = request.Phone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterQueryUserList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterQueryUserListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the authorization of a member to inherit from the organization.
//
// @param request - ModelRouterResetMemberAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterResetMemberAuthorizationResponse
func (client *Client) ModelRouterResetMemberAuthorizationWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterResetMemberAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterResetMemberAuthorizationResponse, _err error) {
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
		Action:      dara.String("ModelRouterResetMemberAuthorization"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/authorization"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterResetMemberAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Flow control management / Save flow control configuration
//
// @param request - ModelRouterSaveFlowConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterSaveFlowConfigResponse
func (client *Client) ModelRouterSaveFlowConfigWithContext(ctx context.Context, request *ModelRouterSaveFlowConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterSaveFlowConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Rpm) {
		body["rpm"] = request.Rpm
	}

	if !dara.IsNil(request.SmoothFlowEnabled) {
		body["smoothFlowEnabled"] = request.SmoothFlowEnabled
	}

	if !dara.IsNil(request.Tpm) {
		body["tpm"] = request.Tpm
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterSaveFlowConfig"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/flow-config"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterSaveFlowConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches and locates nodes in the department tree for organization management.
//
// @param request - ModelRouterSearchClientTreeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterSearchClientTreeResponse
func (client *Client) ModelRouterSearchClientTreeWithContext(ctx context.Context, request *ModelRouterSearchClientTreeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterSearchClientTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterSearchClientTree"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/tree/search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterSearchClientTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the member authorization model for an organization.
//
// @param request - ModelRouterSetMemberAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterSetMemberAuthorizationResponse
func (client *Client) ModelRouterSetMemberAuthorizationWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterSetMemberAuthorizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterSetMemberAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedModelGroupConfig) {
		body["allowedModelGroupConfig"] = request.AllowedModelGroupConfig
	}

	if !dara.IsNil(request.AllowedModels) {
		body["allowedModels"] = request.AllowedModels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterSetMemberAuthorization"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/authorization"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterSetMemberAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets user roles or changes the department to which a user belongs.
//
// Description:
//
// Sets user roles or changes the department to which a user belongs.
//
// @param request - ModelRouterSetUserRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterSetUserRolesResponse
func (client *Client) ModelRouterSetUserRolesWithContext(ctx context.Context, id *string, request *ModelRouterSetUserRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterSetUserRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentRoles) {
		body["departmentRoles"] = request.DepartmentRoles
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterSetUserRoles"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/users/" + dara.PercentEncode(dara.StringValue(id)) + "/roles"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterSetUserRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 组织管理/停止成员订阅
//
// @param request - ModelRouterStopMemberSubscriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterStopMemberSubscriptionResponse
func (client *Client) ModelRouterStopMemberSubscriptionWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterStopMemberSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterStopMemberSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterStopMemberSubscription"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/subscription/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterStopMemberSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a periodic recharge subscription for customer management.
//
// @param request - ModelRouterStopSubscriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterStopSubscriptionResponse
func (client *Client) ModelRouterStopSubscriptionWithContext(ctx context.Context, id *string, request *ModelRouterStopSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterStopSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterStopSubscription"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id)) + "/balance/subscription/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterStopSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers funds from a department to a member within an organization.
//
// @param request - ModelRouterTransferToMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterTransferToMemberResponse
func (client *Client) ModelRouterTransferToMemberWithContext(ctx context.Context, clientId *string, id *string, request *ModelRouterTransferToMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterTransferToMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["amount"] = request.Amount
	}

	if !dara.IsNil(request.BalanceType) {
		body["balanceType"] = request.BalanceType
	}

	if !dara.IsNil(request.IdempotencyKey) {
		body["idempotencyKey"] = request.IdempotencyKey
	}

	if !dara.IsNil(request.MonthlyQuota) {
		body["monthlyQuota"] = request.MonthlyQuota
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterTransferToMember"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(clientId)) + "/members/" + dara.PercentEncode(dara.StringValue(id)) + "/transfer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterTransferToMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of an API key.
//
// @param request - ModelRouterUpdateApiKeyStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterUpdateApiKeyStatusResponse
func (client *Client) ModelRouterUpdateApiKeyStatusWithContext(ctx context.Context, id *string, request *ModelRouterUpdateApiKeyStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterUpdateApiKeyStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterUpdateApiKeyStatus"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/apikeys/" + dara.PercentEncode(dara.StringValue(id)) + "/status"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterUpdateApiKeyStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Billing management/Update billing rules
//
// @param request - ModelRouterUpdateBillingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterUpdateBillingRuleResponse
func (client *Client) ModelRouterUpdateBillingRuleWithContext(ctx context.Context, id *string, request *ModelRouterUpdateBillingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterUpdateBillingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BillingType) {
		body["billingType"] = request.BillingType
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["effectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.ExpireTime) {
		body["expireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.PricingConfig) {
		body["pricingConfig"] = request.PricingConfig
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterUpdateBillingRule"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/billing/rules/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterUpdateBillingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates customer information.
//
// @param request - ModelRouterUpdateClientRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterUpdateClientResponse
func (client *Client) ModelRouterUpdateClientWithContext(ctx context.Context, id *string, request *ModelRouterUpdateClientRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterUpdateClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		body["address"] = request.Address
	}

	if !dara.IsNil(request.AllowedModelGroupConfig) {
		body["allowedModelGroupConfig"] = request.AllowedModelGroupConfig
	}

	if !dara.IsNil(request.AllowedModels) {
		body["allowedModels"] = request.AllowedModels
	}

	if !dara.IsNil(request.Contact) {
		body["contact"] = request.Contact
	}

	if !dara.IsNil(request.Discount) {
		body["discount"] = request.Discount
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterUpdateClient"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/clients/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterUpdateClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Conversation management / Update conversation
//
// @param request - ModelRouterUpdateConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterUpdateConversationResponse
func (client *Client) ModelRouterUpdateConversationWithContext(ctx context.Context, id *string, request *ModelRouterUpdateConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterUpdateConversationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatData) {
		body["chatData"] = request.ChatData
	}

	if !dara.IsNil(request.MessageCount) {
		body["messageCount"] = request.MessageCount
	}

	if !dara.IsNil(request.ModelIds) {
		body["modelIds"] = request.ModelIds
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterUpdateConversation"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/conversations/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterUpdateConversationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Model Management / Update Model
//
// @param request - ModelRouterUpdateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterUpdateModelResponse
func (client *Client) ModelRouterUpdateModelWithContext(ctx context.Context, id *string, request *ModelRouterUpdateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterUpdateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BaseUrl) {
		body["baseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.MaxInputLength) {
		body["maxInputLength"] = request.MaxInputLength
	}

	if !dara.IsNil(request.MaxOutputLength) {
		body["maxOutputLength"] = request.MaxOutputLength
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelType) {
		body["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Symbol) {
		body["symbol"] = request.Symbol
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterUpdateModel"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/models/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterUpdateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a manual model group.
//
// Description:
//
// Edits a manual group.
//
// @param request - ModelRouterUpdateModelGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterUpdateModelGroupResponse
func (client *Client) ModelRouterUpdateModelGroupWithContext(ctx context.Context, groupId *string, request *ModelRouterUpdateModelGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterUpdateModelGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelList) {
		body["modelList"] = request.ModelList
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterUpdateModelGroup"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/model-groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterUpdateModelGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates user information.
//
// Description:
//
// Updates user information.
//
// @param request - ModelRouterUpdateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelRouterUpdateUserResponse
func (client *Client) ModelRouterUpdateUserWithContext(ctx context.Context, id *string, request *ModelRouterUpdateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModelRouterUpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterUpdateUser"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/users/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelRouterUpdateUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Personalized text-to-image: Create image inference tasks using a pre-trained model.
//
// @param request - PersonalizedTextToImageAddInferenceJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageAddInferenceJobResponse
func (client *Client) PersonalizedTextToImageAddInferenceJobWithContext(ctx context.Context, request *PersonalizedTextToImageAddInferenceJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PersonalizedTextToImageAddInferenceJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageNumber) {
		body["imageNumber"] = request.ImageNumber
	}

	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.Seed) {
		body["seed"] = request.Seed
	}

	if !dara.IsNil(request.Strength) {
		body["strength"] = request.Strength
	}

	if !dara.IsNil(request.TrainSteps) {
		body["trainSteps"] = request.TrainSteps
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PersonalizedTextToImageAddInferenceJob"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/addPreModelInferenceJob"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PersonalizedTextToImageAddInferenceJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the content of an image from the personalized text-to-image service using its unique image ID.
//
// @param request - PersonalizedTextToImageQueryImageAssetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageQueryImageAssetResponse
func (client *Client) PersonalizedTextToImageQueryImageAssetWithContext(ctx context.Context, request *PersonalizedTextToImageQueryImageAssetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PersonalizedTextToImageQueryImageAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncodeFormat) {
		query["encodeFormat"] = request.EncodeFormat
	}

	if !dara.IsNil(request.ImageId) {
		query["imageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PersonalizedTextToImageQueryImageAsset"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/queryImageAssetFromImageId"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("any"),
	}
	_result = &PersonalizedTextToImageQueryImageAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a personalized text-to-image inference job.
//
// @param request - PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
func (client *Client) PersonalizedTextToImageQueryPreModelInferenceJobInfoWithContext(ctx context.Context, request *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InferenceJobId) {
		query["inferenceJobId"] = request.InferenceJobId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PersonalizedTextToImageQueryPreModelInferenceJobInfo"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/queryPreModelInferenceJobInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an inference job to generate images based on a personalized text-to-image model.
//
// @param request - Personalizedtxt2imgAddInferenceJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgAddInferenceJobResponse
func (client *Client) Personalizedtxt2imgAddInferenceJobWithContext(ctx context.Context, request *Personalizedtxt2imgAddInferenceJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *Personalizedtxt2imgAddInferenceJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageNumber) {
		body["imageNumber"] = request.ImageNumber
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.Seed) {
		body["seed"] = request.Seed
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Personalizedtxt2imgAddInferenceJob"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/addInferenceJob"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &Personalizedtxt2imgAddInferenceJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Personalized text-to-image: Create a model training task.
//
// @param request - Personalizedtxt2imgAddModelTrainJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgAddModelTrainJobResponse
func (client *Client) Personalizedtxt2imgAddModelTrainJobWithContext(ctx context.Context, request *Personalizedtxt2imgAddModelTrainJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *Personalizedtxt2imgAddModelTrainJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ObjectType) {
		body["objectType"] = request.ObjectType
	}

	if !dara.IsNil(request.TrainSteps) {
		body["trainSteps"] = request.TrainSteps
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Personalizedtxt2imgAddModelTrainJob"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/addModelTrainJob"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &Personalizedtxt2imgAddModelTrainJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the binary data of an image generated by the personalized text-to-image service.
//
// @param request - Personalizedtxt2imgQueryImageAssetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryImageAssetResponse
func (client *Client) Personalizedtxt2imgQueryImageAssetWithContext(ctx context.Context, request *Personalizedtxt2imgQueryImageAssetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *Personalizedtxt2imgQueryImageAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncodeFormat) {
		query["encodeFormat"] = request.EncodeFormat
	}

	if !dara.IsNil(request.ImageId) {
		query["imageId"] = request.ImageId
	}

	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.PromptId) {
		query["promptId"] = request.PromptId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Personalizedtxt2imgQueryImageAsset"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/queryImageAsset"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("any"),
	}
	_result = &Personalizedtxt2imgQueryImageAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the status and results of a Personalizedtxt2img inference job.
//
// @param request - Personalizedtxt2imgQueryInferenceJobInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryInferenceJobInfoResponse
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfoWithContext(ctx context.Context, request *Personalizedtxt2imgQueryInferenceJobInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *Personalizedtxt2imgQueryInferenceJobInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InferenceJobId) {
		query["inferenceJobId"] = request.InferenceJobId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Personalizedtxt2imgQueryInferenceJobInfo"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/queryInferenceJobInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &Personalizedtxt2imgQueryInferenceJobInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Your personalized model training tasks: image generation and query models.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryModelTrainJobListResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainJobListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *Personalizedtxt2imgQueryModelTrainJobListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("Personalizedtxt2imgQueryModelTrainJobList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/queryModelTrainJobList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &Personalizedtxt2imgQueryModelTrainJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the training status of a personalized text-to-image model.
//
// @param request - Personalizedtxt2imgQueryModelTrainStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryModelTrainStatusResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainStatusWithContext(ctx context.Context, request *Personalizedtxt2imgQueryModelTrainStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *Personalizedtxt2imgQueryModelTrainStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelId) {
		query["modelId"] = request.ModelId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Personalizedtxt2imgQueryModelTrainStatus"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/personalizedtxt2img/queryModelTrainStatus"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &Personalizedtxt2imgQueryModelTrainStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an application access ID (appkey).
//
// @param request - QueryApplicationAccessIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApplicationAccessIdResponse
func (client *Client) QueryApplicationAccessIdWithContext(ctx context.Context, request *QueryApplicationAccessIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryApplicationAccessIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationAccessId) {
		query["applicationAccessId"] = request.ApplicationAccessId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryApplicationAccessId"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/queryApplicationAccessId"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryApplicationAccessIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud console / Project list
//
// @param request - QueryProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectResponse
func (client *Client) QueryProjectWithContext(ctx context.Context, request *QueryProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["projectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProject"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/queryProject"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud console / Project List
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectListResponse
func (client *Client) QueryProjectListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryProjectListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProjectList"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/queryProjectList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProjectListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud Console / Purchased Services
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPurchasedServiceResponse
func (client *Client) QueryPurchasedServiceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryPurchasedServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPurchasedService"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/queryPurchasedService"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPurchasedServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud Console / Update project information
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithContext(ctx context.Context, request *UpdateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["projectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aliyunConsole/updateProject"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) executeAITeacherChineseCompositionTutoringWorkflowRunWithSSECtx_opYieldFunc(_yield chan *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse, _yieldErr chan error, ctx context.Context, request *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EssayOutline) {
		body["essayOutline"] = request.EssayOutline
	}

	if !dara.IsNil(request.EssayRequirements) {
		body["essayRequirements"] = request.EssayRequirements
	}

	if !dara.IsNil(request.EssayTopic) {
		body["essayTopic"] = request.EssayTopic
	}

	if !dara.IsNil(request.EssayType) {
		body["essayType"] = request.EssayType
	}

	if !dara.IsNil(request.EssayWordCount) {
		body["essayWordCount"] = request.EssayWordCount
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.ResponseMode) {
		body["responseMode"] = request.ResponseMode
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherChineseCompositionTutoringWorkflowRun"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/intelligentAgent/chineseCompositionTutoring/workflowRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) executeAITeacherEnglishCompositionTutoringWorkflowRunWithSSECtx_opYieldFunc(_yield chan *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse, _yieldErr chan error, ctx context.Context, request *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EssayOutline) {
		body["essayOutline"] = request.EssayOutline
	}

	if !dara.IsNil(request.EssayRequirements) {
		body["essayRequirements"] = request.EssayRequirements
	}

	if !dara.IsNil(request.EssayTopic) {
		body["essayTopic"] = request.EssayTopic
	}

	if !dara.IsNil(request.EssayType) {
		body["essayType"] = request.EssayType
	}

	if !dara.IsNil(request.EssayWordCount) {
		body["essayWordCount"] = request.EssayWordCount
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.ResponseMode) {
		body["responseMode"] = request.ResponseMode
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherEnglishCompositionTutoringWorkflowRun"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/intelligentAgent/englishCompositionTutoring/workflowRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) executeAITeacherEnglishParaphraseChatMessageWithSSECtx_opYieldFunc(_yield chan *ExecuteAITeacherEnglishParaphraseChatMessageResponse, _yieldErr chan error, ctx context.Context, request *ExecuteAITeacherEnglishParaphraseChatMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.QuestionId) {
		body["questionId"] = request.QuestionId
	}

	if !dara.IsNil(request.QuestionInfo) {
		body["questionInfo"] = request.QuestionInfo
	}

	if !dara.IsNil(request.ResponseMode) {
		body["responseMode"] = request.ResponseMode
	}

	if !dara.IsNil(request.UserAnswer) {
		body["userAnswer"] = request.UserAnswer
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAITeacherEnglishParaphraseChatMessage"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/api/v1/intelligentAgent/englishParaphrase/chatMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) executeTextbookAssistantSseDialogueWithSSECtx_opYieldFunc(_yield chan *ExecuteTextbookAssistantSseDialogueResponse, _yieldErr chan error, ctx context.Context, request *ExecuteTextbookAssistantSseDialogueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		body["authToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChatId) {
		body["chatId"] = request.ChatId
	}

	if !dara.IsNil(request.Scenario) {
		body["scenario"] = request.Scenario
	}

	if !dara.IsNil(request.UserMessage) {
		body["userMessage"] = request.UserMessage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTextbookAssistantSseDialogue"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/textbookAssistant/dialogue/ExecuteSseDialogue"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) modelRouterChatCompletionsWithSSECtx_opYieldFunc(_yield chan *ModelRouterChatCompletionsResponse, _yieldErr chan error, ctx context.Context, request *ModelRouterChatCompletionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelRouterChatCompletions"),
		Version:     dara.String("20240611"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/modelRouter/open/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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
