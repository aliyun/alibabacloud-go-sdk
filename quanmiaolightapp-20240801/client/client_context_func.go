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
// 取消异步任务
//
// @param request - CancelAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAsyncTaskResponse
func (client *Client) CancelAsyncTaskWithContext(ctx context.Context, workspaceId *string, request *CancelAsyncTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAsyncTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/cancelAsyncTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出挖掘任务明细
//
// @param tmpReq - ExportAnalysisTagDetailByTaskIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportAnalysisTagDetailByTaskIdResponse
func (client *Client) ExportAnalysisTagDetailByTaskIdWithContext(ctx context.Context, workspaceId *string, tmpReq *ExportAnalysisTagDetailByTaskIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportAnalysisTagDetailByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportAnalysisTagDetailByTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("categories"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		body["categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportAnalysisTagDetailByTaskId"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/exportAnalysisTagDetailByTaskId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新闻播报-抽取分类获取播报热点
//
// @param request - GenerateBroadcastNewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateBroadcastNewsResponse
func (client *Client) GenerateBroadcastNewsWithContext(ctx context.Context, workspaceId *string, request *GenerateBroadcastNewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateBroadcastNewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateBroadcastNews"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/GenerateBroadcastNews"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateBroadcastNewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-标签挖掘-获取示例输出格式
//
// @param tmpReq - GenerateOutputFormatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateOutputFormatResponse
func (client *Client) GenerateOutputFormatWithContext(ctx context.Context, workspaceId *string, tmpReq *GenerateOutputFormatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateOutputFormatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateOutputFormatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["businessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateOutputFormat"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/generateOutputFormat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateOutputFormatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业VOC分析任务结果
//
// @param request - GetEnterpriseVocAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnterpriseVocAnalysisTaskResponse
func (client *Client) GetEnterpriseVocAnalysisTaskWithContext(ctx context.Context, workspaceId *string, request *GetEnterpriseVocAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnterpriseVocAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnterpriseVocAnalysisTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/getEnterpriseVocAnalysisTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取作业批改结果
//
// @param request - GetEssayCorrectionTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEssayCorrectionTaskResponse
func (client *Client) GetEssayCorrectionTaskWithContext(ctx context.Context, workspaceId *string, request *GetEssayCorrectionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEssayCorrectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEssayCorrectionTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/getEssayCorrectionTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEssayCorrectionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件内容
//
// @param request - GetFileContentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileContentResponse
func (client *Client) GetFileContentWithContext(ctx context.Context, workspaceId *string, request *GetFileContentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFileContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["fileKey"] = request.FileKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileContent"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/getFileContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取挖掘分析任务结果
//
// @param request - GetTagMiningAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTagMiningAnalysisTaskResponse
func (client *Client) GetTagMiningAnalysisTaskWithContext(ctx context.Context, workspaceId *string, request *GetTagMiningAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTagMiningAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTagMiningAnalysisTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/getTagMiningAnalysisTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTagMiningAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频理解-获取配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoAnalysisConfigResponse
func (client *Client) GetVideoAnalysisConfigWithContext(ctx context.Context, workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoAnalysisConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoAnalysisConfig"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/getVideoAnalysisConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoAnalysisConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-获取视频理解异步任务结果
//
// @param request - GetVideoAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoAnalysisTaskResponse
func (client *Client) GetVideoAnalysisTaskWithContext(ctx context.Context, workspaceId *string, request *GetVideoAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoAnalysisTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/getVideoAnalysisTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能拆条-获取配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoDetectShotConfigResponse
func (client *Client) GetVideoDetectShotConfigWithContext(ctx context.Context, workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoDetectShotConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoDetectShotConfig"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/getVideoDetectShotConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoDetectShotConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-获取视频拆条异步任务结果
//
// @param request - GetVideoDetectShotTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoDetectShotTaskResponse
func (client *Client) GetVideoDetectShotTaskWithContext(ctx context.Context, workspaceId *string, request *GetVideoDetectShotTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoDetectShotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoDetectShotTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/getVideoDetectShotTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoDetectShotTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 热点新闻推荐
//
// @param request - HotNewsRecommendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotNewsRecommendResponse
func (client *Client) HotNewsRecommendWithContext(ctx context.Context, workspaceId *string, request *HotNewsRecommendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *HotNewsRecommendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotNewsRecommend"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/hotNewsRecommend"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HotNewsRecommendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取挖掘分析结果明细列表
//
// @param request - ListAnalysisTagDetailByTaskIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnalysisTagDetailByTaskIdResponse
func (client *Client) ListAnalysisTagDetailByTaskIdWithContext(ctx context.Context, workspaceId *string, request *ListAnalysisTagDetailByTaskIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAnalysisTagDetailByTaskIdResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnalysisTagDetailByTaskId"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/listAnalysisTagDetailByTaskId"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-新闻播报-获取热点话题摘要列表
//
// @param request - ListHotTopicSummariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotTopicSummariesResponse
func (client *Client) ListHotTopicSummariesWithContext(ctx context.Context, workspaceId *string, request *ListHotTopicSummariesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHotTopicSummariesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.HotTopic) {
		body["hotTopic"] = request.HotTopic
	}

	if !dara.IsNil(request.HotTopicVersion) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotTopicSummaries"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/listHotTopicSummaries"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotTopicSummariesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业VOC分析
//
// @param tmpReq - RunEnterpriseVocAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunEnterpriseVocAnalysisResponse
func (client *Client) RunEnterpriseVocAnalysisWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunEnterpriseVocAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunEnterpriseVocAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runEnterpriseVocAnalysisWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 企业VOC分析
//
// @param tmpReq - RunEnterpriseVocAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunEnterpriseVocAnalysisResponse
func (client *Client) RunEnterpriseVocAnalysisWithContext(ctx context.Context, workspaceId *string, tmpReq *RunEnterpriseVocAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunEnterpriseVocAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunEnterpriseVocAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterTags) {
		request.FilterTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTags, dara.String("filterTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AkProxy) {
		body["akProxy"] = request.AkProxy
	}

	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.FilterTagsShrink) {
		body["filterTags"] = request.FilterTagsShrink
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.SourceTrace) {
		body["sourceTrace"] = request.SourceTrace
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunEnterpriseVocAnalysis"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runEnterpriseVocAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunEnterpriseVocAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作业批改
//
// @param request - RunEssayCorrectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunEssayCorrectionResponse
func (client *Client) RunEssayCorrectionWithSSECtx(ctx context.Context, workspaceId *string, request *RunEssayCorrectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunEssayCorrectionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runEssayCorrectionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 作业批改
//
// @param request - RunEssayCorrectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunEssayCorrectionResponse
func (client *Client) RunEssayCorrectionWithContext(ctx context.Context, workspaceId *string, request *RunEssayCorrectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunEssayCorrectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Answer) {
		body["answer"] = request.Answer
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OtherReviewPoints) {
		body["otherReviewPoints"] = request.OtherReviewPoints
	}

	if !dara.IsNil(request.Question) {
		body["question"] = request.Question
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	if !dara.IsNil(request.TotalScore) {
		body["totalScore"] = request.TotalScore
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunEssayCorrection"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runEssayCorrection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunEssayCorrectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-热点播报-问答
//
// @param tmpReq - RunHotTopicChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotTopicChatResponse
func (client *Client) RunHotTopicChatWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunHotTopicChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunHotTopicChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runHotTopicChatWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 轻应用-热点播报-问答
//
// @param tmpReq - RunHotTopicChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotTopicChatResponse
func (client *Client) RunHotTopicChatWithContext(ctx context.Context, workspaceId *string, tmpReq *RunHotTopicChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunHotTopicChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunHotTopicChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GenerateOptions) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, dara.String("generateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HotTopics) {
		request.HotTopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotTopics, dara.String("hotTopics"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("messages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StepForBroadcastContentConfig) {
		request.StepForBroadcastContentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForBroadcastContentConfig, dara.String("stepForBroadcastContentConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.GenerateOptionsShrink) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !dara.IsNil(request.HotTopicVersion) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !dara.IsNil(request.HotTopicsShrink) {
		body["hotTopics"] = request.HotTopicsShrink
	}

	if !dara.IsNil(request.ImageCount) {
		body["imageCount"] = request.ImageCount
	}

	if !dara.IsNil(request.MessagesShrink) {
		body["messages"] = request.MessagesShrink
	}

	if !dara.IsNil(request.ModelCustomPromptTemplate) {
		body["modelCustomPromptTemplate"] = request.ModelCustomPromptTemplate
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.StepForBroadcastContentConfigShrink) {
		body["stepForBroadcastContentConfig"] = request.StepForBroadcastContentConfigShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunHotTopicChat"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runHotTopicChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunHotTopicChatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-热点播报-热点摘要生成
//
// @param tmpReq - RunHotTopicSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotTopicSummaryResponse
func (client *Client) RunHotTopicSummaryWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunHotTopicSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunHotTopicSummaryResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runHotTopicSummaryWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 轻应用-热点播报-热点摘要生成
//
// @param tmpReq - RunHotTopicSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunHotTopicSummaryResponse
func (client *Client) RunHotTopicSummaryWithContext(ctx context.Context, workspaceId *string, tmpReq *RunHotTopicSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunHotTopicSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunHotTopicSummaryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StepForCustomSummaryStyleConfig) {
		request.StepForCustomSummaryStyleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForCustomSummaryStyleConfig, dara.String("stepForCustomSummaryStyleConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TopicIds) {
		request.TopicIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicIds, dara.String("topicIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotTopicVersion) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !dara.IsNil(request.StepForCustomSummaryStyleConfigShrink) {
		body["stepForCustomSummaryStyleConfig"] = request.StepForCustomSummaryStyleConfigShrink
	}

	if !dara.IsNil(request.TopicIdsShrink) {
		body["topicIds"] = request.TopicIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunHotTopicSummary"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runHotTopicSummary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunHotTopicSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营销信息抽取服务
//
// @param tmpReq - RunMarketingInformationExtractRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMarketingInformationExtractResponse
func (client *Client) RunMarketingInformationExtractWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunMarketingInformationExtractRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunMarketingInformationExtractResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runMarketingInformationExtractWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 营销信息抽取服务
//
// @param tmpReq - RunMarketingInformationExtractRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMarketingInformationExtractResponse
func (client *Client) RunMarketingInformationExtractWithContext(ctx context.Context, workspaceId *string, tmpReq *RunMarketingInformationExtractRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunMarketingInformationExtractResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunMarketingInformationExtractShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceMaterials) {
		request.SourceMaterialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMaterials, dara.String("sourceMaterials"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.ExtractType) {
		body["extractType"] = request.ExtractType
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.SourceMaterialsShrink) {
		body["sourceMaterials"] = request.SourceMaterialsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunMarketingInformationExtract"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runMarketingInformationExtract"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunMarketingInformationExtractResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营销文案写作服务
//
// @param request - RunMarketingInformationWritingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMarketingInformationWritingResponse
func (client *Client) RunMarketingInformationWritingWithSSECtx(ctx context.Context, workspaceId *string, request *RunMarketingInformationWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunMarketingInformationWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runMarketingInformationWritingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 营销文案写作服务
//
// @param request - RunMarketingInformationWritingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunMarketingInformationWritingResponse
func (client *Client) RunMarketingInformationWritingWithContext(ctx context.Context, workspaceId *string, request *RunMarketingInformationWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunMarketingInformationWritingResponse, _err error) {
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

	if !dara.IsNil(request.CustomLimitation) {
		body["customLimitation"] = request.CustomLimitation
	}

	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.InputExample) {
		body["inputExample"] = request.InputExample
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputExample) {
		body["outputExample"] = request.OutputExample
	}

	if !dara.IsNil(request.SourceMaterial) {
		body["sourceMaterial"] = request.SourceMaterial
	}

	if !dara.IsNil(request.WritingType) {
		body["writingType"] = request.WritingType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunMarketingInformationWriting"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runMarketingInformationWriting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunMarketingInformationWritingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-网络内容审核
//
// @param tmpReq - RunNetworkContentAuditRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNetworkContentAuditResponse
func (client *Client) RunNetworkContentAuditWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunNetworkContentAuditRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunNetworkContentAuditResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runNetworkContentAuditWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 轻应用-网络内容审核
//
// @param tmpReq - RunNetworkContentAuditRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNetworkContentAuditResponse
func (client *Client) RunNetworkContentAuditWithContext(ctx context.Context, workspaceId *string, tmpReq *RunNetworkContentAuditRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunNetworkContentAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunNetworkContentAuditShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BusinessType) {
		body["businessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunNetworkContentAudit"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runNetworkContentAudit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunNetworkContentAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作业批改
//
// @param request - RunOcrParseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunOcrParseResponse
func (client *Client) RunOcrParseWithSSECtx(ctx context.Context, workspaceId *string, request *RunOcrParseRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunOcrParseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runOcrParseWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 作业批改
//
// @param request - RunOcrParseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunOcrParseResponse
func (client *Client) RunOcrParseWithContext(ctx context.Context, workspaceId *string, request *RunOcrParseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunOcrParseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["fileKey"] = request.FileKey
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunOcrParse"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runOcrParse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunOcrParseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 长剧本创作
//
// @param request - RunScriptChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptChatResponse
func (client *Client) RunScriptChatWithSSECtx(ctx context.Context, workspaceId *string, request *RunScriptChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptChatWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 长剧本创作
//
// @param request - RunScriptChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptChatResponse
func (client *Client) RunScriptChatWithContext(ctx context.Context, workspaceId *string, request *RunScriptChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptChat"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunScriptChatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 剧本续写
//
// @param request - RunScriptContinueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptContinueResponse
func (client *Client) RunScriptContinueWithSSECtx(ctx context.Context, workspaceId *string, request *RunScriptContinueRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptContinueResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptContinueWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 剧本续写
//
// @param request - RunScriptContinueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptContinueResponse
func (client *Client) RunScriptContinueWithContext(ctx context.Context, workspaceId *string, request *RunScriptContinueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptContinueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ScriptSummary) {
		body["scriptSummary"] = request.ScriptSummary
	}

	if !dara.IsNil(request.ScriptTypeKeyword) {
		body["scriptTypeKeyword"] = request.ScriptTypeKeyword
	}

	if !dara.IsNil(request.UserProvidedContent) {
		body["userProvidedContent"] = request.UserProvidedContent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptContinue"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptContinue"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunScriptContinueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 剧本策划
//
// @param request - RunScriptPlanningRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptPlanningResponse
func (client *Client) RunScriptPlanningWithSSECtx(ctx context.Context, workspaceId *string, request *RunScriptPlanningRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptPlanningResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptPlanningWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 剧本策划
//
// @param request - RunScriptPlanningRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptPlanningResponse
func (client *Client) RunScriptPlanningWithContext(ctx context.Context, workspaceId *string, request *RunScriptPlanningRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptPlanningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalNote) {
		body["additionalNote"] = request.AdditionalNote
	}

	if !dara.IsNil(request.DialogueInScene) {
		body["dialogueInScene"] = request.DialogueInScene
	}

	if !dara.IsNil(request.PlotConflict) {
		body["plotConflict"] = request.PlotConflict
	}

	if !dara.IsNil(request.ScriptName) {
		body["scriptName"] = request.ScriptName
	}

	if !dara.IsNil(request.ScriptShotCount) {
		body["scriptShotCount"] = request.ScriptShotCount
	}

	if !dara.IsNil(request.ScriptSummary) {
		body["scriptSummary"] = request.ScriptSummary
	}

	if !dara.IsNil(request.ScriptTypeKeyword) {
		body["scriptTypeKeyword"] = request.ScriptTypeKeyword
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptPlanning"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptPlanning"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunScriptPlanningResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 剧本对话内容的整理
//
// @param request - RunScriptRefineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptRefineResponse
func (client *Client) RunScriptRefineWithSSECtx(ctx context.Context, workspaceId *string, request *RunScriptRefineRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptRefineResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptRefineWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 剧本对话内容的整理
//
// @param request - RunScriptRefineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunScriptRefineResponse
func (client *Client) RunScriptRefineWithContext(ctx context.Context, workspaceId *string, request *RunScriptRefineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptRefineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptRefine"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptRefine"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunScriptRefineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文体学习和写作推理服务
//
// @param tmpReq - RunStyleWritingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStyleWritingResponse
func (client *Client) RunStyleWritingWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunStyleWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunStyleWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runStyleWritingWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 文体学习和写作推理服务
//
// @param tmpReq - RunStyleWritingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStyleWritingResponse
func (client *Client) RunStyleWritingWithContext(ctx context.Context, workspaceId *string, tmpReq *RunStyleWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunStyleWritingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunStyleWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LearningSamples) {
		request.LearningSamplesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LearningSamples, dara.String("learningSamples"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReferenceMaterials) {
		request.ReferenceMaterialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceMaterials, dara.String("referenceMaterials"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LearningSamplesShrink) {
		body["learningSamples"] = request.LearningSamplesShrink
	}

	if !dara.IsNil(request.ProcessStage) {
		body["processStage"] = request.ProcessStage
	}

	if !dara.IsNil(request.ReferenceMaterialsShrink) {
		body["referenceMaterials"] = request.ReferenceMaterialsShrink
	}

	if !dara.IsNil(request.StyleFeature) {
		body["styleFeature"] = request.StyleFeature
	}

	if !dara.IsNil(request.UseSearch) {
		body["useSearch"] = request.UseSearch
	}

	if !dara.IsNil(request.WritingTheme) {
		body["writingTheme"] = request.WritingTheme
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunStyleWriting"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runStyleWriting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunStyleWritingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-标签挖掘
//
// @param tmpReq - RunTagMiningAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTagMiningAnalysisResponse
func (client *Client) RunTagMiningAnalysisWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunTagMiningAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunTagMiningAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTagMiningAnalysisWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 轻应用-标签挖掘
//
// @param tmpReq - RunTagMiningAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTagMiningAnalysisResponse
func (client *Client) RunTagMiningAnalysisWithContext(ctx context.Context, workspaceId *string, tmpReq *RunTagMiningAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunTagMiningAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunTagMiningAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BusinessType) {
		body["businessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTagMiningAnalysis"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runTagMiningAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunTagMiningAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-视频理解
//
// @param tmpReq - RunVideoAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunVideoAnalysisResponse
func (client *Client) RunVideoAnalysisWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunVideoAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunVideoAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runVideoAnalysisWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 轻应用-视频理解
//
// @param tmpReq - RunVideoAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunVideoAnalysisResponse
func (client *Client) RunVideoAnalysisWithContext(ctx context.Context, workspaceId *string, tmpReq *RunVideoAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunVideoAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunVideoAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddDocumentParam) {
		request.AddDocumentParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddDocumentParam, dara.String("addDocumentParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeGenerateOptions) {
		request.ExcludeGenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeGenerateOptions, dara.String("excludeGenerateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FrameSampleMethod) {
		request.FrameSampleMethodShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FrameSampleMethod, dara.String("frameSampleMethod"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GenerateOptions) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, dara.String("generateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TextProcessTasks) {
		request.TextProcessTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextProcessTasks, dara.String("textProcessTasks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoCaptionInfo) {
		request.VideoCaptionInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoCaptionInfo, dara.String("videoCaptionInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoRoles) {
		request.VideoRolesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoRoles, dara.String("videoRoles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoUrls) {
		request.VideoUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoUrls, dara.String("videoUrls"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddDocumentParamShrink) {
		body["addDocumentParam"] = request.AddDocumentParamShrink
	}

	if !dara.IsNil(request.AutoRoleRecognitionVideoUrl) {
		body["autoRoleRecognitionVideoUrl"] = request.AutoRoleRecognitionVideoUrl
	}

	if !dara.IsNil(request.ExcludeGenerateOptionsShrink) {
		body["excludeGenerateOptions"] = request.ExcludeGenerateOptionsShrink
	}

	if !dara.IsNil(request.FaceIdentitySimilarityMinScore) {
		body["faceIdentitySimilarityMinScore"] = request.FaceIdentitySimilarityMinScore
	}

	if !dara.IsNil(request.FrameSampleMethodShrink) {
		body["frameSampleMethod"] = request.FrameSampleMethodShrink
	}

	if !dara.IsNil(request.GenerateOptionsShrink) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.ModelCustomPromptTemplate) {
		body["modelCustomPromptTemplate"] = request.ModelCustomPromptTemplate
	}

	if !dara.IsNil(request.ModelCustomPromptTemplateId) {
		body["modelCustomPromptTemplateId"] = request.ModelCustomPromptTemplateId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.SnapshotInterval) {
		body["snapshotInterval"] = request.SnapshotInterval
	}

	if !dara.IsNil(request.SplitInterval) {
		body["splitInterval"] = request.SplitInterval
	}

	if !dara.IsNil(request.SplitType) {
		body["splitType"] = request.SplitType
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.TextProcessTasksShrink) {
		body["textProcessTasks"] = request.TextProcessTasksShrink
	}

	if !dara.IsNil(request.VideoCaptionInfoShrink) {
		body["videoCaptionInfo"] = request.VideoCaptionInfoShrink
	}

	if !dara.IsNil(request.VideoExtraInfo) {
		body["videoExtraInfo"] = request.VideoExtraInfo
	}

	if !dara.IsNil(request.VideoModelCustomPromptTemplate) {
		body["videoModelCustomPromptTemplate"] = request.VideoModelCustomPromptTemplate
	}

	if !dara.IsNil(request.VideoModelId) {
		body["videoModelId"] = request.VideoModelId
	}

	if !dara.IsNil(request.VideoRolesShrink) {
		body["videoRoles"] = request.VideoRolesShrink
	}

	if !dara.IsNil(request.VideoShotFaceIdentityCount) {
		body["videoShotFaceIdentityCount"] = request.VideoShotFaceIdentityCount
	}

	if !dara.IsNil(request.VideoUrl) {
		body["videoUrl"] = request.VideoUrl
	}

	if !dara.IsNil(request.VideoUrlsShrink) {
		body["videoUrls"] = request.VideoUrlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunVideoAnalysis"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runVideoAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunVideoAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-视频拆条
//
// @param tmpReq - RunVideoDetectShotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunVideoDetectShotResponse
func (client *Client) RunVideoDetectShotWithSSECtx(ctx context.Context, workspaceId *string, tmpReq *RunVideoDetectShotRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunVideoDetectShotResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runVideoDetectShotWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 轻应用-视频拆条
//
// @param tmpReq - RunVideoDetectShotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunVideoDetectShotResponse
func (client *Client) RunVideoDetectShotWithContext(ctx context.Context, workspaceId *string, tmpReq *RunVideoDetectShotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunVideoDetectShotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunVideoDetectShotShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecognitionOptions) {
		request.RecognitionOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecognitionOptions, dara.String("recognitionOptions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IntelliSimpPrompt) {
		body["intelliSimpPrompt"] = request.IntelliSimpPrompt
	}

	if !dara.IsNil(request.IntelliSimpPromptTemplateId) {
		body["intelliSimpPromptTemplateId"] = request.IntelliSimpPromptTemplateId
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.ModelCustomPromptTemplateId) {
		body["modelCustomPromptTemplateId"] = request.ModelCustomPromptTemplateId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelVlCustomPromptTemplateId) {
		body["modelVlCustomPromptTemplateId"] = request.ModelVlCustomPromptTemplateId
	}

	if !dara.IsNil(request.OptionsShrink) {
		body["options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.PreModelId) {
		body["preModelId"] = request.PreModelId
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.RecognitionOptionsShrink) {
		body["recognitionOptions"] = request.RecognitionOptionsShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.VideoUrl) {
		body["videoUrl"] = request.VideoUrl
	}

	if !dara.IsNil(request.VlPrompt) {
		body["vlPrompt"] = request.VlPrompt
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunVideoDetectShot"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runVideoDetectShot"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunVideoDetectShotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交企业VOC异步任务
//
// @param tmpReq - SubmitEnterpriseVocAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitEnterpriseVocAnalysisTaskResponse
func (client *Client) SubmitEnterpriseVocAnalysisTaskWithContext(ctx context.Context, workspaceId *string, tmpReq *SubmitEnterpriseVocAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitEnterpriseVocAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitEnterpriseVocAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contents) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, dara.String("contents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterTags) {
		request.FilterTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTags, dara.String("filterTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ContentsShrink) {
		body["contents"] = request.ContentsShrink
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.FileKey) {
		body["fileKey"] = request.FileKey
	}

	if !dara.IsNil(request.FilterTagsShrink) {
		body["filterTags"] = request.FilterTagsShrink
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.SourceTrace) {
		body["sourceTrace"] = request.SourceTrace
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitEnterpriseVocAnalysisTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/submitEnterpriseVocAnalysisTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交批改任务
//
// @param tmpReq - SubmitEssayCorrectionTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitEssayCorrectionTaskResponse
func (client *Client) SubmitEssayCorrectionTaskWithContext(ctx context.Context, workspaceId *string, tmpReq *SubmitEssayCorrectionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitEssayCorrectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitEssayCorrectionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("tasks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OtherReviewPoints) {
		body["otherReviewPoints"] = request.OtherReviewPoints
	}

	if !dara.IsNil(request.Question) {
		body["question"] = request.Question
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	if !dara.IsNil(request.TasksShrink) {
		body["tasks"] = request.TasksShrink
	}

	if !dara.IsNil(request.TotalScore) {
		body["totalScore"] = request.TotalScore
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitEssayCorrectionTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/submitEssayCorrectionTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitEssayCorrectionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-标签挖掘
//
// @param tmpReq - SubmitTagMiningAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTagMiningAnalysisTaskResponse
func (client *Client) SubmitTagMiningAnalysisTaskWithContext(ctx context.Context, workspaceId *string, tmpReq *SubmitTagMiningAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitTagMiningAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitTagMiningAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contents) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, dara.String("contents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BusinessType) {
		body["businessType"] = request.BusinessType
	}

	if !dara.IsNil(request.ContentsShrink) {
		body["contents"] = request.ContentsShrink
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTagMiningAnalysisTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/submitTagMiningAnalysisTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTagMiningAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-提交视频理解任务
//
// @param tmpReq - SubmitVideoAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoAnalysisTaskResponse
func (client *Client) SubmitVideoAnalysisTaskWithContext(ctx context.Context, workspaceId *string, tmpReq *SubmitVideoAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitVideoAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitVideoAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddDocumentParam) {
		request.AddDocumentParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddDocumentParam, dara.String("addDocumentParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeGenerateOptions) {
		request.ExcludeGenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeGenerateOptions, dara.String("excludeGenerateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FrameSampleMethod) {
		request.FrameSampleMethodShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FrameSampleMethod, dara.String("frameSampleMethod"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GenerateOptions) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, dara.String("generateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TextProcessTasks) {
		request.TextProcessTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextProcessTasks, dara.String("textProcessTasks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoCaptionInfo) {
		request.VideoCaptionInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoCaptionInfo, dara.String("videoCaptionInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoRoles) {
		request.VideoRolesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoRoles, dara.String("videoRoles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoUrls) {
		request.VideoUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoUrls, dara.String("videoUrls"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddDocumentParamShrink) {
		body["addDocumentParam"] = request.AddDocumentParamShrink
	}

	if !dara.IsNil(request.AutoRoleRecognitionVideoUrl) {
		body["autoRoleRecognitionVideoUrl"] = request.AutoRoleRecognitionVideoUrl
	}

	if !dara.IsNil(request.DeduplicationId) {
		body["deduplicationId"] = request.DeduplicationId
	}

	if !dara.IsNil(request.ExcludeGenerateOptionsShrink) {
		body["excludeGenerateOptions"] = request.ExcludeGenerateOptionsShrink
	}

	if !dara.IsNil(request.FaceIdentitySimilarityMinScore) {
		body["faceIdentitySimilarityMinScore"] = request.FaceIdentitySimilarityMinScore
	}

	if !dara.IsNil(request.FrameSampleMethodShrink) {
		body["frameSampleMethod"] = request.FrameSampleMethodShrink
	}

	if !dara.IsNil(request.GenerateOptionsShrink) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.ModelCustomPromptTemplate) {
		body["modelCustomPromptTemplate"] = request.ModelCustomPromptTemplate
	}

	if !dara.IsNil(request.ModelCustomPromptTemplateId) {
		body["modelCustomPromptTemplateId"] = request.ModelCustomPromptTemplateId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.SnapshotInterval) {
		body["snapshotInterval"] = request.SnapshotInterval
	}

	if !dara.IsNil(request.SplitInterval) {
		body["splitInterval"] = request.SplitInterval
	}

	if !dara.IsNil(request.SplitType) {
		body["splitType"] = request.SplitType
	}

	if !dara.IsNil(request.TextProcessTasksShrink) {
		body["textProcessTasks"] = request.TextProcessTasksShrink
	}

	if !dara.IsNil(request.VideoCaptionInfoShrink) {
		body["videoCaptionInfo"] = request.VideoCaptionInfoShrink
	}

	if !dara.IsNil(request.VideoExtraInfo) {
		body["videoExtraInfo"] = request.VideoExtraInfo
	}

	if !dara.IsNil(request.VideoModelCustomPromptTemplate) {
		body["videoModelCustomPromptTemplate"] = request.VideoModelCustomPromptTemplate
	}

	if !dara.IsNil(request.VideoModelId) {
		body["videoModelId"] = request.VideoModelId
	}

	if !dara.IsNil(request.VideoRolesShrink) {
		body["videoRoles"] = request.VideoRolesShrink
	}

	if !dara.IsNil(request.VideoShotFaceIdentityCount) {
		body["videoShotFaceIdentityCount"] = request.VideoShotFaceIdentityCount
	}

	if !dara.IsNil(request.VideoUrl) {
		body["videoUrl"] = request.VideoUrl
	}

	if !dara.IsNil(request.VideoUrlsShrink) {
		body["videoUrls"] = request.VideoUrlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoAnalysisTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/submitVideoAnalysisTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 轻应用-提交视频拆条任务
//
// @param tmpReq - SubmitVideoDetectShotTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoDetectShotTaskResponse
func (client *Client) SubmitVideoDetectShotTaskWithContext(ctx context.Context, workspaceId *string, tmpReq *SubmitVideoDetectShotTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitVideoDetectShotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitVideoDetectShotTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecognitionOptions) {
		request.RecognitionOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecognitionOptions, dara.String("recognitionOptions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeduplicationId) {
		body["deduplicationId"] = request.DeduplicationId
	}

	if !dara.IsNil(request.IntelliSimpPrompt) {
		body["intelliSimpPrompt"] = request.IntelliSimpPrompt
	}

	if !dara.IsNil(request.IntelliSimpPromptTemplateId) {
		body["intelliSimpPromptTemplateId"] = request.IntelliSimpPromptTemplateId
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.ModelCustomPromptTemplateId) {
		body["modelCustomPromptTemplateId"] = request.ModelCustomPromptTemplateId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelVlCustomPromptTemplateId) {
		body["modelVlCustomPromptTemplateId"] = request.ModelVlCustomPromptTemplateId
	}

	if !dara.IsNil(request.OptionsShrink) {
		body["options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.PreModelId) {
		body["preModelId"] = request.PreModelId
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.RecognitionOptionsShrink) {
		body["recognitionOptions"] = request.RecognitionOptionsShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.VideoUrl) {
		body["videoUrl"] = request.VideoUrl
	}

	if !dara.IsNil(request.VlPrompt) {
		body["vlPrompt"] = request.VlPrompt
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoDetectShotTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/submitVideoDetectShotTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoDetectShotTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频理解-更新配置
//
// @param request - UpdateVideoAnalysisConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoAnalysisConfigResponse
func (client *Client) UpdateVideoAnalysisConfigWithContext(ctx context.Context, workspaceId *string, request *UpdateVideoAnalysisConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVideoAnalysisConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AsyncConcurrency) {
		body["asyncConcurrency"] = request.AsyncConcurrency
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoAnalysisConfig"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/updateVideoAnalysisConfig"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoAnalysisConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频理解-修改任务状态
//
// @param request - UpdateVideoAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoAnalysisTaskResponse
func (client *Client) UpdateVideoAnalysisTaskWithContext(ctx context.Context, workspaceId *string, request *UpdateVideoAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVideoAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStatus) {
		body["taskStatus"] = request.TaskStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoAnalysisTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/updateVideoAnalysisTask"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频理解-批量修改任务状态
//
// @param tmpReq - UpdateVideoAnalysisTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoAnalysisTasksResponse
func (client *Client) UpdateVideoAnalysisTasksWithContext(ctx context.Context, workspaceId *string, tmpReq *UpdateVideoAnalysisTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVideoAnalysisTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateVideoAnalysisTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("taskIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskIdsShrink) {
		body["taskIds"] = request.TaskIdsShrink
	}

	if !dara.IsNil(request.TaskStatus) {
		body["taskStatus"] = request.TaskStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoAnalysisTasks"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/updateVideoAnalysisTasks"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoAnalysisTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能拆条-更新配置
//
// @param request - UpdateVideoDetectShotConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoDetectShotConfigResponse
func (client *Client) UpdateVideoDetectShotConfigWithContext(ctx context.Context, workspaceId *string, request *UpdateVideoDetectShotConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVideoDetectShotConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AsyncConcurrency) {
		body["asyncConcurrency"] = request.AsyncConcurrency
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoDetectShotConfig"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/videoAnalysis/updateVideoDetectShotConfig"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoDetectShotConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频拆条-修改任务状态
//
// @param request - UpdateVideoDetectShotTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoDetectShotTaskResponse
func (client *Client) UpdateVideoDetectShotTaskWithContext(ctx context.Context, workspaceId *string, request *UpdateVideoDetectShotTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVideoDetectShotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStatus) {
		body["taskStatus"] = request.TaskStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoDetectShotTask"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/updateVideoDetectShotTask"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoDetectShotTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) runEnterpriseVocAnalysisWithSSECtx_opYieldFunc(_yield chan *RunEnterpriseVocAnalysisResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunEnterpriseVocAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunEnterpriseVocAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterTags) {
		request.FilterTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTags, dara.String("filterTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AkProxy) {
		body["akProxy"] = request.AkProxy
	}

	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.FilterTagsShrink) {
		body["filterTags"] = request.FilterTagsShrink
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.SourceTrace) {
		body["sourceTrace"] = request.SourceTrace
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunEnterpriseVocAnalysis"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runEnterpriseVocAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runEssayCorrectionWithSSECtx_opYieldFunc(_yield chan *RunEssayCorrectionResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunEssayCorrectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Answer) {
		body["answer"] = request.Answer
	}

	if !dara.IsNil(request.Grade) {
		body["grade"] = request.Grade
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OtherReviewPoints) {
		body["otherReviewPoints"] = request.OtherReviewPoints
	}

	if !dara.IsNil(request.Question) {
		body["question"] = request.Question
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	if !dara.IsNil(request.TotalScore) {
		body["totalScore"] = request.TotalScore
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunEssayCorrection"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runEssayCorrection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runHotTopicChatWithSSECtx_opYieldFunc(_yield chan *RunHotTopicChatResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunHotTopicChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunHotTopicChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GenerateOptions) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, dara.String("generateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HotTopics) {
		request.HotTopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HotTopics, dara.String("hotTopics"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("messages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StepForBroadcastContentConfig) {
		request.StepForBroadcastContentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForBroadcastContentConfig, dara.String("stepForBroadcastContentConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.GenerateOptionsShrink) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !dara.IsNil(request.HotTopicVersion) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !dara.IsNil(request.HotTopicsShrink) {
		body["hotTopics"] = request.HotTopicsShrink
	}

	if !dara.IsNil(request.ImageCount) {
		body["imageCount"] = request.ImageCount
	}

	if !dara.IsNil(request.MessagesShrink) {
		body["messages"] = request.MessagesShrink
	}

	if !dara.IsNil(request.ModelCustomPromptTemplate) {
		body["modelCustomPromptTemplate"] = request.ModelCustomPromptTemplate
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.StepForBroadcastContentConfigShrink) {
		body["stepForBroadcastContentConfig"] = request.StepForBroadcastContentConfigShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunHotTopicChat"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runHotTopicChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runHotTopicSummaryWithSSECtx_opYieldFunc(_yield chan *RunHotTopicSummaryResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunHotTopicSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunHotTopicSummaryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StepForCustomSummaryStyleConfig) {
		request.StepForCustomSummaryStyleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StepForCustomSummaryStyleConfig, dara.String("stepForCustomSummaryStyleConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TopicIds) {
		request.TopicIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicIds, dara.String("topicIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HotTopicVersion) {
		body["hotTopicVersion"] = request.HotTopicVersion
	}

	if !dara.IsNil(request.StepForCustomSummaryStyleConfigShrink) {
		body["stepForCustomSummaryStyleConfig"] = request.StepForCustomSummaryStyleConfigShrink
	}

	if !dara.IsNil(request.TopicIdsShrink) {
		body["topicIds"] = request.TopicIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunHotTopicSummary"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runHotTopicSummary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runMarketingInformationExtractWithSSECtx_opYieldFunc(_yield chan *RunMarketingInformationExtractResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunMarketingInformationExtractRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunMarketingInformationExtractShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceMaterials) {
		request.SourceMaterialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMaterials, dara.String("sourceMaterials"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.ExtractType) {
		body["extractType"] = request.ExtractType
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.SourceMaterialsShrink) {
		body["sourceMaterials"] = request.SourceMaterialsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunMarketingInformationExtract"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runMarketingInformationExtract"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runMarketingInformationWritingWithSSECtx_opYieldFunc(_yield chan *RunMarketingInformationWritingResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunMarketingInformationWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.CustomLimitation) {
		body["customLimitation"] = request.CustomLimitation
	}

	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.InputExample) {
		body["inputExample"] = request.InputExample
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputExample) {
		body["outputExample"] = request.OutputExample
	}

	if !dara.IsNil(request.SourceMaterial) {
		body["sourceMaterial"] = request.SourceMaterial
	}

	if !dara.IsNil(request.WritingType) {
		body["writingType"] = request.WritingType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunMarketingInformationWriting"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runMarketingInformationWriting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runNetworkContentAuditWithSSECtx_opYieldFunc(_yield chan *RunNetworkContentAuditResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunNetworkContentAuditRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunNetworkContentAuditShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BusinessType) {
		body["businessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunNetworkContentAudit"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runNetworkContentAudit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runOcrParseWithSSECtx_opYieldFunc(_yield chan *RunOcrParseResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunOcrParseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["fileKey"] = request.FileKey
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunOcrParse"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runOcrParse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runScriptChatWithSSECtx_opYieldFunc(_yield chan *RunScriptChatResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunScriptChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptChat"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptChat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runScriptContinueWithSSECtx_opYieldFunc(_yield chan *RunScriptContinueResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunScriptContinueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ScriptSummary) {
		body["scriptSummary"] = request.ScriptSummary
	}

	if !dara.IsNil(request.ScriptTypeKeyword) {
		body["scriptTypeKeyword"] = request.ScriptTypeKeyword
	}

	if !dara.IsNil(request.UserProvidedContent) {
		body["userProvidedContent"] = request.UserProvidedContent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptContinue"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptContinue"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runScriptPlanningWithSSECtx_opYieldFunc(_yield chan *RunScriptPlanningResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunScriptPlanningRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalNote) {
		body["additionalNote"] = request.AdditionalNote
	}

	if !dara.IsNil(request.DialogueInScene) {
		body["dialogueInScene"] = request.DialogueInScene
	}

	if !dara.IsNil(request.PlotConflict) {
		body["plotConflict"] = request.PlotConflict
	}

	if !dara.IsNil(request.ScriptName) {
		body["scriptName"] = request.ScriptName
	}

	if !dara.IsNil(request.ScriptShotCount) {
		body["scriptShotCount"] = request.ScriptShotCount
	}

	if !dara.IsNil(request.ScriptSummary) {
		body["scriptSummary"] = request.ScriptSummary
	}

	if !dara.IsNil(request.ScriptTypeKeyword) {
		body["scriptTypeKeyword"] = request.ScriptTypeKeyword
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptPlanning"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptPlanning"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runScriptRefineWithSSECtx_opYieldFunc(_yield chan *RunScriptRefineResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunScriptRefineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunScriptRefine"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runScriptRefine"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runStyleWritingWithSSECtx_opYieldFunc(_yield chan *RunStyleWritingResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunStyleWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunStyleWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LearningSamples) {
		request.LearningSamplesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LearningSamples, dara.String("learningSamples"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReferenceMaterials) {
		request.ReferenceMaterialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceMaterials, dara.String("referenceMaterials"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LearningSamplesShrink) {
		body["learningSamples"] = request.LearningSamplesShrink
	}

	if !dara.IsNil(request.ProcessStage) {
		body["processStage"] = request.ProcessStage
	}

	if !dara.IsNil(request.ReferenceMaterialsShrink) {
		body["referenceMaterials"] = request.ReferenceMaterialsShrink
	}

	if !dara.IsNil(request.StyleFeature) {
		body["styleFeature"] = request.StyleFeature
	}

	if !dara.IsNil(request.UseSearch) {
		body["useSearch"] = request.UseSearch
	}

	if !dara.IsNil(request.WritingTheme) {
		body["writingTheme"] = request.WritingTheme
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunStyleWriting"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runStyleWriting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runTagMiningAnalysisWithSSECtx_opYieldFunc(_yield chan *RunTagMiningAnalysisResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunTagMiningAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunTagMiningAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BusinessType) {
		body["businessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["extraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OutputFormat) {
		body["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.TagsShrink) {
		body["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDescription) {
		body["taskDescription"] = request.TaskDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTagMiningAnalysis"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runTagMiningAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runVideoAnalysisWithSSECtx_opYieldFunc(_yield chan *RunVideoAnalysisResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunVideoAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunVideoAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddDocumentParam) {
		request.AddDocumentParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddDocumentParam, dara.String("addDocumentParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeGenerateOptions) {
		request.ExcludeGenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeGenerateOptions, dara.String("excludeGenerateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FrameSampleMethod) {
		request.FrameSampleMethodShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FrameSampleMethod, dara.String("frameSampleMethod"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GenerateOptions) {
		request.GenerateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenerateOptions, dara.String("generateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TextProcessTasks) {
		request.TextProcessTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextProcessTasks, dara.String("textProcessTasks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoCaptionInfo) {
		request.VideoCaptionInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoCaptionInfo, dara.String("videoCaptionInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoRoles) {
		request.VideoRolesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoRoles, dara.String("videoRoles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoUrls) {
		request.VideoUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoUrls, dara.String("videoUrls"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddDocumentParamShrink) {
		body["addDocumentParam"] = request.AddDocumentParamShrink
	}

	if !dara.IsNil(request.AutoRoleRecognitionVideoUrl) {
		body["autoRoleRecognitionVideoUrl"] = request.AutoRoleRecognitionVideoUrl
	}

	if !dara.IsNil(request.ExcludeGenerateOptionsShrink) {
		body["excludeGenerateOptions"] = request.ExcludeGenerateOptionsShrink
	}

	if !dara.IsNil(request.FaceIdentitySimilarityMinScore) {
		body["faceIdentitySimilarityMinScore"] = request.FaceIdentitySimilarityMinScore
	}

	if !dara.IsNil(request.FrameSampleMethodShrink) {
		body["frameSampleMethod"] = request.FrameSampleMethodShrink
	}

	if !dara.IsNil(request.GenerateOptionsShrink) {
		body["generateOptions"] = request.GenerateOptionsShrink
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.ModelCustomPromptTemplate) {
		body["modelCustomPromptTemplate"] = request.ModelCustomPromptTemplate
	}

	if !dara.IsNil(request.ModelCustomPromptTemplateId) {
		body["modelCustomPromptTemplateId"] = request.ModelCustomPromptTemplateId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.SnapshotInterval) {
		body["snapshotInterval"] = request.SnapshotInterval
	}

	if !dara.IsNil(request.SplitInterval) {
		body["splitInterval"] = request.SplitInterval
	}

	if !dara.IsNil(request.SplitType) {
		body["splitType"] = request.SplitType
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.TextProcessTasksShrink) {
		body["textProcessTasks"] = request.TextProcessTasksShrink
	}

	if !dara.IsNil(request.VideoCaptionInfoShrink) {
		body["videoCaptionInfo"] = request.VideoCaptionInfoShrink
	}

	if !dara.IsNil(request.VideoExtraInfo) {
		body["videoExtraInfo"] = request.VideoExtraInfo
	}

	if !dara.IsNil(request.VideoModelCustomPromptTemplate) {
		body["videoModelCustomPromptTemplate"] = request.VideoModelCustomPromptTemplate
	}

	if !dara.IsNil(request.VideoModelId) {
		body["videoModelId"] = request.VideoModelId
	}

	if !dara.IsNil(request.VideoRolesShrink) {
		body["videoRoles"] = request.VideoRolesShrink
	}

	if !dara.IsNil(request.VideoShotFaceIdentityCount) {
		body["videoShotFaceIdentityCount"] = request.VideoShotFaceIdentityCount
	}

	if !dara.IsNil(request.VideoUrl) {
		body["videoUrl"] = request.VideoUrl
	}

	if !dara.IsNil(request.VideoUrlsShrink) {
		body["videoUrls"] = request.VideoUrlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunVideoAnalysis"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runVideoAnalysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runVideoDetectShotWithSSECtx_opYieldFunc(_yield chan *RunVideoDetectShotResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, tmpReq *RunVideoDetectShotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunVideoDetectShotShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecognitionOptions) {
		request.RecognitionOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecognitionOptions, dara.String("recognitionOptions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IntelliSimpPrompt) {
		body["intelliSimpPrompt"] = request.IntelliSimpPrompt
	}

	if !dara.IsNil(request.IntelliSimpPromptTemplateId) {
		body["intelliSimpPromptTemplateId"] = request.IntelliSimpPromptTemplateId
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.ModelCustomPromptTemplateId) {
		body["modelCustomPromptTemplateId"] = request.ModelCustomPromptTemplateId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelVlCustomPromptTemplateId) {
		body["modelVlCustomPromptTemplateId"] = request.ModelVlCustomPromptTemplateId
	}

	if !dara.IsNil(request.OptionsShrink) {
		body["options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.OriginalSessionId) {
		body["originalSessionId"] = request.OriginalSessionId
	}

	if !dara.IsNil(request.PreModelId) {
		body["preModelId"] = request.PreModelId
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.RecognitionOptionsShrink) {
		body["recognitionOptions"] = request.RecognitionOptionsShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.VideoUrl) {
		body["videoUrl"] = request.VideoUrl
	}

	if !dara.IsNil(request.VlPrompt) {
		body["vlPrompt"] = request.VlPrompt
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunVideoDetectShot"),
		Version:     dara.String("2024-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/quanmiao/lightapp/runVideoDetectShot"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
