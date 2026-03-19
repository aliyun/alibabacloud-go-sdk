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
// 拓展练问答对生成
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
// 同步基础练问答对生成
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
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
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
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
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
// 阿里云控制台/列出购买过的资源列表
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
// 智能批改/口语评测/统计/调用量
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
// 智能批改/口语评测/统计/并发数
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
// 智能批改/口语评测/统计/调用错误
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
// 阿里云控制台/授权凭证创建
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
// 阿里云控制台/创建项目
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
// 中文作文辅导
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
// 中文作文辅导
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
// 英语作文辅导
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
// 英语作文辅导
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
// 英文释义
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
// 英文释义
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
// 进行拓展练对话
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
// 拓展练根据上下文进行润色
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
// 拓展练语境翻译
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
// 语法检测
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
// 进行同步练对话
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
// 同步练语境翻译
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
// 进行AI对话
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
// 调整难度
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
// 语法检测
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
// 句子润色
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
// 对话重试
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
// 进行对话-流式输出
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
// 进行对话-流式输出
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
// 开启自由对话
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
// 获取鉴权参数
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
// 翻译消息内容
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
// 拓展练小助手
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
// 同步练小助手
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
// 获取请求鉴权参数
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
// 批量获取文章详情
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
// 获取文章列表
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
// 获取书本下的目录信息
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
// 获取包含年级下的书本列表
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
// 获取有资源的年级信息
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
// 获取文章内容详情
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
// 聊天/聊天接口
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
// 聊天/聊天接口
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
// API密钥管理/复制API密钥
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
// API密钥管理/创建API密钥
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
// 客户管理/创建客户
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

	if !dara.IsNil(request.AllowedModels) {
		body["allowedModels"] = request.AllowedModels
	}

	if !dara.IsNil(request.Contact) {
		body["contact"] = request.Contact
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
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
// 对话管理/新建对话
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
// 模型管理/创建模型
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
// API密钥管理/删除API密钥
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
// 客户管理/删除客户
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
// 对话管理/删除对话
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
// 模型管理/删除模型
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
// API密钥管理/获取API密钥详情
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
// API密钥管理/获取API密钥列表
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
// 客户管理/获取客户列表
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
// 对话管理/获取对话详情
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
// 对话管理/获取对话列表
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
// 模型管理/获取模型详情
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
// 模型管理/获取模型列表
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
// Nacos配置/获取Nacos服务提供者列表
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
// Nacos配置/获取Nacos标签列表
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
// 模型观测/获取观测图表数据
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

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
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
// 模型观测/获取观测日志列表
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

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
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
// 模型观测/获取观测指标数据
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

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupBy) {
		query["groupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
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
// 客户管理/更新客户
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

	if !dara.IsNil(request.AllowedModels) {
		body["allowedModels"] = request.AllowedModels
	}

	if !dara.IsNil(request.Contact) {
		body["contact"] = request.Contact
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
// 对话管理/更新对话
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
// 模型管理/更新模型
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
// 个性化文生图/基于一个预训练模型创建图片推理任务
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
// 个性化文生图/通过唯一的图片编号获取图片内容
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
// 个性化文生图/查询预制模型推理任务的状态
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
// 个性化文生图/基于一个模型创建图片推理任务
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
// 个性化文生图/创建一个模型训练任务
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
// 个性化文生图/图片二进制内容获取
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
// 个性化文生图/查询模型推理任务的状态和结果信息
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
// 个性化文生图/查询模型训练任务列表
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
// 个性化文生图/模型训练状态查询
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
// 阿里云控制台/获取应用访问识别码(appkey)信息
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
// 阿里云控制台/获取项目列表
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
// 阿里云控制台/获取项目列表
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
// 阿里云控制台/已经购买过的服务项目
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
// 阿里云控制台/更新项目信息
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
