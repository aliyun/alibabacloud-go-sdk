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
	client.Endpoint, _err = client.GetEndpoint(dara.String("quanmiaolightapp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 取消异步任务
//
// @param request - CancelAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAsyncTaskResponse
func (client *Client) CancelAsyncTaskWithOptions(workspaceId *string, request *CancelAsyncTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelAsyncTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消异步任务
//
// @param request - CancelAsyncTaskRequest
//
// @return CancelAsyncTaskResponse
func (client *Client) CancelAsyncTask(workspaceId *string, request *CancelAsyncTaskRequest) (_result *CancelAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CancelAsyncTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportAnalysisTagDetailByTaskIdWithOptions(workspaceId *string, tmpReq *ExportAnalysisTagDetailByTaskIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportAnalysisTagDetailByTaskIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ExportAnalysisTagDetailByTaskIdRequest
//
// @return ExportAnalysisTagDetailByTaskIdResponse
func (client *Client) ExportAnalysisTagDetailByTaskId(workspaceId *string, request *ExportAnalysisTagDetailByTaskIdRequest) (_result *ExportAnalysisTagDetailByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExportAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.ExportAnalysisTagDetailByTaskIdWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateBroadcastNewsWithOptions(workspaceId *string, request *GenerateBroadcastNewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateBroadcastNewsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GenerateBroadcastNewsResponse
func (client *Client) GenerateBroadcastNews(workspaceId *string, request *GenerateBroadcastNewsRequest) (_result *GenerateBroadcastNewsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateBroadcastNewsResponse{}
	_body, _err := client.GenerateBroadcastNewsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateOutputFormatWithOptions(workspaceId *string, tmpReq *GenerateOutputFormatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateOutputFormatResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GenerateOutputFormatRequest
//
// @return GenerateOutputFormatResponse
func (client *Client) GenerateOutputFormat(workspaceId *string, request *GenerateOutputFormatRequest) (_result *GenerateOutputFormatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateOutputFormatResponse{}
	_body, _err := client.GenerateOutputFormatWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEnterpriseVocAnalysisTaskWithOptions(workspaceId *string, request *GetEnterpriseVocAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnterpriseVocAnalysisTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEnterpriseVocAnalysisTaskResponse
func (client *Client) GetEnterpriseVocAnalysisTask(workspaceId *string, request *GetEnterpriseVocAnalysisTaskRequest) (_result *GetEnterpriseVocAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.GetEnterpriseVocAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEssayCorrectionTaskWithOptions(workspaceId *string, request *GetEssayCorrectionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEssayCorrectionTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEssayCorrectionTaskResponse
func (client *Client) GetEssayCorrectionTask(workspaceId *string, request *GetEssayCorrectionTaskRequest) (_result *GetEssayCorrectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEssayCorrectionTaskResponse{}
	_body, _err := client.GetEssayCorrectionTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFileContentWithOptions(workspaceId *string, request *GetFileContentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFileContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFileContentResponse
func (client *Client) GetFileContent(workspaceId *string, request *GetFileContentRequest) (_result *GetFileContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileContentResponse{}
	_body, _err := client.GetFileContentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTagMiningAnalysisTaskWithOptions(workspaceId *string, request *GetTagMiningAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTagMiningAnalysisTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTagMiningAnalysisTaskResponse
func (client *Client) GetTagMiningAnalysisTask(workspaceId *string, request *GetTagMiningAnalysisTaskRequest) (_result *GetTagMiningAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTagMiningAnalysisTaskResponse{}
	_body, _err := client.GetTagMiningAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetVideoAnalysisConfigWithOptions(workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoAnalysisConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetVideoAnalysisConfigResponse
func (client *Client) GetVideoAnalysisConfig(workspaceId *string) (_result *GetVideoAnalysisConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoAnalysisConfigResponse{}
	_body, _err := client.GetVideoAnalysisConfigWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetVideoAnalysisTaskWithOptions(workspaceId *string, request *GetVideoAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoAnalysisTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetVideoAnalysisTaskResponse
func (client *Client) GetVideoAnalysisTask(workspaceId *string, request *GetVideoAnalysisTaskRequest) (_result *GetVideoAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoAnalysisTaskResponse{}
	_body, _err := client.GetVideoAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) HotNewsRecommendWithOptions(workspaceId *string, request *HotNewsRecommendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *HotNewsRecommendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return HotNewsRecommendResponse
func (client *Client) HotNewsRecommend(workspaceId *string, request *HotNewsRecommendRequest) (_result *HotNewsRecommendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HotNewsRecommendResponse{}
	_body, _err := client.HotNewsRecommendWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAnalysisTagDetailByTaskIdWithOptions(workspaceId *string, request *ListAnalysisTagDetailByTaskIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAnalysisTagDetailByTaskIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAnalysisTagDetailByTaskIdResponse
func (client *Client) ListAnalysisTagDetailByTaskId(workspaceId *string, request *ListAnalysisTagDetailByTaskIdRequest) (_result *ListAnalysisTagDetailByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAnalysisTagDetailByTaskIdResponse{}
	_body, _err := client.ListAnalysisTagDetailByTaskIdWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListHotTopicSummariesWithOptions(workspaceId *string, request *ListHotTopicSummariesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHotTopicSummariesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListHotTopicSummariesResponse
func (client *Client) ListHotTopicSummaries(workspaceId *string, request *ListHotTopicSummariesRequest) (_result *ListHotTopicSummariesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHotTopicSummariesResponse{}
	_body, _err := client.ListHotTopicSummariesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunEnterpriseVocAnalysisWithSSE(workspaceId *string, tmpReq *RunEnterpriseVocAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunEnterpriseVocAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runEnterpriseVocAnalysisWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunEnterpriseVocAnalysisWithOptions(workspaceId *string, tmpReq *RunEnterpriseVocAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunEnterpriseVocAnalysisResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunEnterpriseVocAnalysisRequest
//
// @return RunEnterpriseVocAnalysisResponse
func (client *Client) RunEnterpriseVocAnalysis(workspaceId *string, request *RunEnterpriseVocAnalysisRequest) (_result *RunEnterpriseVocAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunEnterpriseVocAnalysisResponse{}
	_body, _err := client.RunEnterpriseVocAnalysisWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunEssayCorrectionWithSSE(workspaceId *string, request *RunEssayCorrectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunEssayCorrectionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runEssayCorrectionWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
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
func (client *Client) RunEssayCorrectionWithOptions(workspaceId *string, request *RunEssayCorrectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunEssayCorrectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunEssayCorrectionResponse
func (client *Client) RunEssayCorrection(workspaceId *string, request *RunEssayCorrectionRequest) (_result *RunEssayCorrectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunEssayCorrectionResponse{}
	_body, _err := client.RunEssayCorrectionWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunHotTopicChatWithSSE(workspaceId *string, tmpReq *RunHotTopicChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunHotTopicChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runHotTopicChatWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunHotTopicChatWithOptions(workspaceId *string, tmpReq *RunHotTopicChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunHotTopicChatResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunHotTopicChatRequest
//
// @return RunHotTopicChatResponse
func (client *Client) RunHotTopicChat(workspaceId *string, request *RunHotTopicChatRequest) (_result *RunHotTopicChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunHotTopicChatResponse{}
	_body, _err := client.RunHotTopicChatWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunHotTopicSummaryWithSSE(workspaceId *string, tmpReq *RunHotTopicSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunHotTopicSummaryResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runHotTopicSummaryWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunHotTopicSummaryWithOptions(workspaceId *string, tmpReq *RunHotTopicSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunHotTopicSummaryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunHotTopicSummaryRequest
//
// @return RunHotTopicSummaryResponse
func (client *Client) RunHotTopicSummary(workspaceId *string, request *RunHotTopicSummaryRequest) (_result *RunHotTopicSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunHotTopicSummaryResponse{}
	_body, _err := client.RunHotTopicSummaryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunMarketingInformationExtractWithSSE(workspaceId *string, tmpReq *RunMarketingInformationExtractRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunMarketingInformationExtractResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runMarketingInformationExtractWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunMarketingInformationExtractWithOptions(workspaceId *string, tmpReq *RunMarketingInformationExtractRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunMarketingInformationExtractResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunMarketingInformationExtractRequest
//
// @return RunMarketingInformationExtractResponse
func (client *Client) RunMarketingInformationExtract(workspaceId *string, request *RunMarketingInformationExtractRequest) (_result *RunMarketingInformationExtractResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunMarketingInformationExtractResponse{}
	_body, _err := client.RunMarketingInformationExtractWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunMarketingInformationWritingWithSSE(workspaceId *string, request *RunMarketingInformationWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunMarketingInformationWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runMarketingInformationWritingWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
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
func (client *Client) RunMarketingInformationWritingWithOptions(workspaceId *string, request *RunMarketingInformationWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunMarketingInformationWritingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunMarketingInformationWritingResponse
func (client *Client) RunMarketingInformationWriting(workspaceId *string, request *RunMarketingInformationWritingRequest) (_result *RunMarketingInformationWritingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunMarketingInformationWritingResponse{}
	_body, _err := client.RunMarketingInformationWritingWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunNetworkContentAuditWithSSE(workspaceId *string, tmpReq *RunNetworkContentAuditRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunNetworkContentAuditResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runNetworkContentAuditWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunNetworkContentAuditWithOptions(workspaceId *string, tmpReq *RunNetworkContentAuditRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunNetworkContentAuditResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunNetworkContentAuditRequest
//
// @return RunNetworkContentAuditResponse
func (client *Client) RunNetworkContentAudit(workspaceId *string, request *RunNetworkContentAuditRequest) (_result *RunNetworkContentAuditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunNetworkContentAuditResponse{}
	_body, _err := client.RunNetworkContentAuditWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunOcrParseWithSSE(workspaceId *string, request *RunOcrParseRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunOcrParseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runOcrParseWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
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
func (client *Client) RunOcrParseWithOptions(workspaceId *string, request *RunOcrParseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunOcrParseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunOcrParseResponse
func (client *Client) RunOcrParse(workspaceId *string, request *RunOcrParseRequest) (_result *RunOcrParseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunOcrParseResponse{}
	_body, _err := client.RunOcrParseWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunScriptChatWithSSE(workspaceId *string, request *RunScriptChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptChatWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
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
func (client *Client) RunScriptChatWithOptions(workspaceId *string, request *RunScriptChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptChatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunScriptChatResponse
func (client *Client) RunScriptChat(workspaceId *string, request *RunScriptChatRequest) (_result *RunScriptChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptChatResponse{}
	_body, _err := client.RunScriptChatWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunScriptContinueWithSSE(workspaceId *string, request *RunScriptContinueRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptContinueResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptContinueWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
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
func (client *Client) RunScriptContinueWithOptions(workspaceId *string, request *RunScriptContinueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptContinueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunScriptContinueResponse
func (client *Client) RunScriptContinue(workspaceId *string, request *RunScriptContinueRequest) (_result *RunScriptContinueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptContinueResponse{}
	_body, _err := client.RunScriptContinueWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunScriptPlanningWithSSE(workspaceId *string, request *RunScriptPlanningRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptPlanningResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptPlanningWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
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
func (client *Client) RunScriptPlanningWithOptions(workspaceId *string, request *RunScriptPlanningRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptPlanningResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunScriptPlanningResponse
func (client *Client) RunScriptPlanning(workspaceId *string, request *RunScriptPlanningRequest) (_result *RunScriptPlanningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptPlanningResponse{}
	_body, _err := client.RunScriptPlanningWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunScriptRefineWithSSE(workspaceId *string, request *RunScriptRefineRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunScriptRefineResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runScriptRefineWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, request, headers, runtime)
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
func (client *Client) RunScriptRefineWithOptions(workspaceId *string, request *RunScriptRefineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunScriptRefineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunScriptRefineResponse
func (client *Client) RunScriptRefine(workspaceId *string, request *RunScriptRefineRequest) (_result *RunScriptRefineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunScriptRefineResponse{}
	_body, _err := client.RunScriptRefineWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunStyleWritingWithSSE(workspaceId *string, tmpReq *RunStyleWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunStyleWritingResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runStyleWritingWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunStyleWritingWithOptions(workspaceId *string, tmpReq *RunStyleWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunStyleWritingResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunStyleWritingRequest
//
// @return RunStyleWritingResponse
func (client *Client) RunStyleWriting(workspaceId *string, request *RunStyleWritingRequest) (_result *RunStyleWritingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunStyleWritingResponse{}
	_body, _err := client.RunStyleWritingWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunTagMiningAnalysisWithSSE(workspaceId *string, tmpReq *RunTagMiningAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunTagMiningAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runTagMiningAnalysisWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunTagMiningAnalysisWithOptions(workspaceId *string, tmpReq *RunTagMiningAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunTagMiningAnalysisResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunTagMiningAnalysisRequest
//
// @return RunTagMiningAnalysisResponse
func (client *Client) RunTagMiningAnalysis(workspaceId *string, request *RunTagMiningAnalysisRequest) (_result *RunTagMiningAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunTagMiningAnalysisResponse{}
	_body, _err := client.RunTagMiningAnalysisWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunVideoAnalysisWithSSE(workspaceId *string, tmpReq *RunVideoAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunVideoAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runVideoAnalysisWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, tmpReq, headers, runtime)
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
func (client *Client) RunVideoAnalysisWithOptions(workspaceId *string, tmpReq *RunVideoAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunVideoAnalysisResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RunVideoAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
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

	body := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RunVideoAnalysisRequest
//
// @return RunVideoAnalysisResponse
func (client *Client) RunVideoAnalysis(workspaceId *string, request *RunVideoAnalysisRequest) (_result *RunVideoAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunVideoAnalysisResponse{}
	_body, _err := client.RunVideoAnalysisWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitEnterpriseVocAnalysisTaskWithOptions(workspaceId *string, tmpReq *SubmitEnterpriseVocAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitEnterpriseVocAnalysisTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitEnterpriseVocAnalysisTaskRequest
//
// @return SubmitEnterpriseVocAnalysisTaskResponse
func (client *Client) SubmitEnterpriseVocAnalysisTask(workspaceId *string, request *SubmitEnterpriseVocAnalysisTaskRequest) (_result *SubmitEnterpriseVocAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitEnterpriseVocAnalysisTaskResponse{}
	_body, _err := client.SubmitEnterpriseVocAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitEssayCorrectionTaskWithOptions(workspaceId *string, tmpReq *SubmitEssayCorrectionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitEssayCorrectionTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitEssayCorrectionTaskRequest
//
// @return SubmitEssayCorrectionTaskResponse
func (client *Client) SubmitEssayCorrectionTask(workspaceId *string, request *SubmitEssayCorrectionTaskRequest) (_result *SubmitEssayCorrectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitEssayCorrectionTaskResponse{}
	_body, _err := client.SubmitEssayCorrectionTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitTagMiningAnalysisTaskWithOptions(workspaceId *string, tmpReq *SubmitTagMiningAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitTagMiningAnalysisTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitTagMiningAnalysisTaskRequest
//
// @return SubmitTagMiningAnalysisTaskResponse
func (client *Client) SubmitTagMiningAnalysisTask(workspaceId *string, request *SubmitTagMiningAnalysisTaskRequest) (_result *SubmitTagMiningAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitTagMiningAnalysisTaskResponse{}
	_body, _err := client.SubmitTagMiningAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitVideoAnalysisTaskWithOptions(workspaceId *string, tmpReq *SubmitVideoAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitVideoAnalysisTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitVideoAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
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

	body := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubmitVideoAnalysisTaskRequest
//
// @return SubmitVideoAnalysisTaskResponse
func (client *Client) SubmitVideoAnalysisTask(workspaceId *string, request *SubmitVideoAnalysisTaskRequest) (_result *SubmitVideoAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitVideoAnalysisTaskResponse{}
	_body, _err := client.SubmitVideoAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateVideoAnalysisConfigWithOptions(workspaceId *string, request *UpdateVideoAnalysisConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVideoAnalysisConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateVideoAnalysisConfigResponse
func (client *Client) UpdateVideoAnalysisConfig(workspaceId *string, request *UpdateVideoAnalysisConfigRequest) (_result *UpdateVideoAnalysisConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateVideoAnalysisConfigResponse{}
	_body, _err := client.UpdateVideoAnalysisConfigWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateVideoAnalysisTaskWithOptions(workspaceId *string, request *UpdateVideoAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVideoAnalysisTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateVideoAnalysisTaskResponse
func (client *Client) UpdateVideoAnalysisTask(workspaceId *string, request *UpdateVideoAnalysisTaskRequest) (_result *UpdateVideoAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateVideoAnalysisTaskResponse{}
	_body, _err := client.UpdateVideoAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) runEnterpriseVocAnalysisWithSSE_opYieldFunc(_yield chan *RunEnterpriseVocAnalysisResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunEnterpriseVocAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runEssayCorrectionWithSSE_opYieldFunc(_yield chan *RunEssayCorrectionResponse, _yieldErr chan error, workspaceId *string, request *RunEssayCorrectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runHotTopicChatWithSSE_opYieldFunc(_yield chan *RunHotTopicChatResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunHotTopicChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runHotTopicSummaryWithSSE_opYieldFunc(_yield chan *RunHotTopicSummaryResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunHotTopicSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runMarketingInformationExtractWithSSE_opYieldFunc(_yield chan *RunMarketingInformationExtractResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunMarketingInformationExtractRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runMarketingInformationWritingWithSSE_opYieldFunc(_yield chan *RunMarketingInformationWritingResponse, _yieldErr chan error, workspaceId *string, request *RunMarketingInformationWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runNetworkContentAuditWithSSE_opYieldFunc(_yield chan *RunNetworkContentAuditResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunNetworkContentAuditRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runOcrParseWithSSE_opYieldFunc(_yield chan *RunOcrParseResponse, _yieldErr chan error, workspaceId *string, request *RunOcrParseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runScriptChatWithSSE_opYieldFunc(_yield chan *RunScriptChatResponse, _yieldErr chan error, workspaceId *string, request *RunScriptChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runScriptContinueWithSSE_opYieldFunc(_yield chan *RunScriptContinueResponse, _yieldErr chan error, workspaceId *string, request *RunScriptContinueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runScriptPlanningWithSSE_opYieldFunc(_yield chan *RunScriptPlanningResponse, _yieldErr chan error, workspaceId *string, request *RunScriptPlanningRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runScriptRefineWithSSE_opYieldFunc(_yield chan *RunScriptRefineResponse, _yieldErr chan error, workspaceId *string, request *RunScriptRefineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runStyleWritingWithSSE_opYieldFunc(_yield chan *RunStyleWritingResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunStyleWritingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runTagMiningAnalysisWithSSE_opYieldFunc(_yield chan *RunTagMiningAnalysisResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunTagMiningAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runVideoAnalysisWithSSE_opYieldFunc(_yield chan *RunVideoAnalysisResponse, _yieldErr chan error, workspaceId *string, tmpReq *RunVideoAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := tmpReq.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
	}
	request := &RunVideoAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
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

	body := map[string]interface{}{}
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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
