// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - CreateAsyncTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsyncTranslateResponse
func (client *Client) CreateAsyncTranslateWithContext(ctx context.Context, request *CreateAsyncTranslateRequest, runtime *dara.RuntimeOptions) (_result *CreateAsyncTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAsyncTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAsyncTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDocTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocTranslateTaskResponse
func (client *Client) CreateDocTranslateTaskWithContext(ctx context.Context, request *CreateDocTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDocTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackUrl) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图片翻译任务
//
// @param request - CreateImageTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageTranslateTaskResponse
func (client *Client) CreateImageTranslateTaskWithContext(ctx context.Context, request *CreateImageTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.UrlList) {
		body["UrlList"] = request.UrlList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - GetAsyncTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTranslateResponse
func (client *Client) GetAsyncTranslateWithContext(ctx context.Context, request *GetAsyncTranslateRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量文本翻译
//
// @param request - GetBatchTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTranslateResponse
func (client *Client) GetBatchTranslateWithContext(ctx context.Context, request *GetBatchTranslateRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetBatchTranslateByVPC
//
// @param request - GetBatchTranslateByVPCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTranslateByVPCResponse
func (client *Client) GetBatchTranslateByVPCWithContext(ctx context.Context, request *GetBatchTranslateByVPCRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTranslateByVPCResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTranslateByVPC"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTranslateByVPCResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectLanguageResponse
func (client *Client) GetDetectLanguageWithContext(ctx context.Context, request *GetDetectLanguageRequest, runtime *dara.RuntimeOptions) (_result *GetDetectLanguageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetectLanguage"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetectLanguageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectLanguageVpcResponse
func (client *Client) GetDetectLanguageVpcWithContext(ctx context.Context, request *GetDetectLanguageVpcRequest, runtime *dara.RuntimeOptions) (_result *GetDetectLanguageVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetectLanguageVpc"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetectLanguageVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档翻译任务
//
// @param request - GetDocTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocTranslateTaskResponse
func (client *Client) GetDocTranslateTaskWithContext(ctx context.Context, request *GetDocTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDocTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetImageDiagnoseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDiagnoseResponse
func (client *Client) GetImageDiagnoseWithContext(ctx context.Context, request *GetImageDiagnoseRequest, runtime *dara.RuntimeOptions) (_result *GetImageDiagnoseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageDiagnose"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageDiagnoseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片翻译结果
//
// @param request - GetImageTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTranslateResponse
func (client *Client) GetImageTranslateWithContext(ctx context.Context, request *GetImageTranslateRequest, runtime *dara.RuntimeOptions) (_result *GetImageTranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageTranslate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片翻译任务
//
// @param request - GetImageTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTranslateTaskResponse
func (client *Client) GetImageTranslateTaskWithContext(ctx context.Context, request *GetImageTranslateTaskRequest, runtime *dara.RuntimeOptions) (_result *GetImageTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageTranslateTask"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTitleDiagnose
//
// @param request - GetTitleDiagnoseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleDiagnoseResponse
func (client *Client) GetTitleDiagnoseWithContext(ctx context.Context, request *GetTitleDiagnoseRequest, runtime *dara.RuntimeOptions) (_result *GetTitleDiagnoseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTitleDiagnose"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTitleDiagnoseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTitleGenerate
//
// @param request - GetTitleGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleGenerateResponse
func (client *Client) GetTitleGenerateWithContext(ctx context.Context, request *GetTitleGenerateRequest, runtime *dara.RuntimeOptions) (_result *GetTitleGenerateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		body["Attributes"] = request.Attributes
	}

	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.HotWords) {
		body["HotWords"] = request.HotWords
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTitleGenerate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTitleGenerateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTitleIntelligence
//
// @param request - GetTitleIntelligenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleIntelligenceResponse
func (client *Client) GetTitleIntelligenceWithContext(ctx context.Context, request *GetTitleIntelligenceRequest, runtime *dara.RuntimeOptions) (_result *GetTitleIntelligenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CatLevelThreeId) {
		body["CatLevelThreeId"] = request.CatLevelThreeId
	}

	if !dara.IsNil(request.CatLevelTwoId) {
		body["CatLevelTwoId"] = request.CatLevelTwoId
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Keywords) {
		body["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTitleIntelligence"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTitleIntelligenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片批量翻译结果
//
// @param request - GetTranslateImageBatchResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslateImageBatchResultResponse
func (client *Client) GetTranslateImageBatchResultWithContext(ctx context.Context, request *GetTranslateImageBatchResultRequest, runtime *dara.RuntimeOptions) (_result *GetTranslateImageBatchResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslateImageBatchResult"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslateImageBatchResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTranslateReport
//
// @param request - GetTranslateReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslateReportResponse
func (client *Client) GetTranslateReportWithContext(ctx context.Context, request *GetTranslateReportRequest, runtime *dara.RuntimeOptions) (_result *GetTranslateReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslateReport"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslateReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通服务
//
// @param request - OpenAlimtServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAlimtServiceResponse
func (client *Client) OpenAlimtServiceWithContext(ctx context.Context, request *OpenAlimtServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenAlimtServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenAlimtService"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenAlimtServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 专业文本翻译
//
// @param request - TranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateResponse
func (client *Client) TranslateWithContext(ctx context.Context, request *TranslateRequest, runtime *dara.RuntimeOptions) (_result *TranslateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Translate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # TranslateCertificate
//
// @param request - TranslateCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateCertificateResponse
func (client *Client) TranslateCertificateWithContext(ctx context.Context, request *TranslateCertificateRequest, runtime *dara.RuntimeOptions) (_result *TranslateCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertificateType) {
		body["CertificateType"] = request.CertificateType
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ResultType) {
		body["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateCertificate"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI TranslateECommerce is deprecated, please use alimt::2018-10-12::Translate instead.
//
// Summary:
//
// # TranslateECommerce
//
// @param request - TranslateECommerceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateECommerceResponse
func (client *Client) TranslateECommerceWithContext(ctx context.Context, request *TranslateECommerceRequest, runtime *dara.RuntimeOptions) (_result *TranslateECommerceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateECommerce"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateECommerceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本通用翻译
//
// @param request - TranslateGeneralRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateGeneralResponse
func (client *Client) TranslateGeneralWithContext(ctx context.Context, request *TranslateGeneralRequest, runtime *dara.RuntimeOptions) (_result *TranslateGeneralResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateGeneral"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateGeneralResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # TranslateGeneralVpc
//
// @param request - TranslateGeneralVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateGeneralVpcResponse
func (client *Client) TranslateGeneralVpcWithContext(ctx context.Context, request *TranslateGeneralVpcRequest, runtime *dara.RuntimeOptions) (_result *TranslateGeneralVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateGeneralVpc"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateGeneralVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 公有云图片翻译产品API
//
// @param request - TranslateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateImageResponse
func (client *Client) TranslateImageWithContext(ctx context.Context, request *TranslateImageRequest, runtime *dara.RuntimeOptions) (_result *TranslateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ext) {
		body["Ext"] = request.Ext
	}

	if !dara.IsNil(request.Field) {
		body["Field"] = request.Field
	}

	if !dara.IsNil(request.ImageBase64) {
		body["ImageBase64"] = request.ImageBase64
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateImage"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量图片翻译接口
//
// @param request - TranslateImageBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateImageBatchResponse
func (client *Client) TranslateImageBatchWithContext(ctx context.Context, request *TranslateImageBatchRequest, runtime *dara.RuntimeOptions) (_result *TranslateImageBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomTaskId) {
		body["CustomTaskId"] = request.CustomTaskId
	}

	if !dara.IsNil(request.Ext) {
		body["Ext"] = request.Ext
	}

	if !dara.IsNil(request.Field) {
		body["Field"] = request.Field
	}

	if !dara.IsNil(request.ImageUrls) {
		body["ImageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateImageBatch"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateImageBatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索翻译
//
// @param request - TranslateSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateSearchResponse
func (client *Client) TranslateSearchWithContext(ctx context.Context, request *TranslateSearchRequest, runtime *dara.RuntimeOptions) (_result *TranslateSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.SourceText) {
		body["SourceText"] = request.SourceText
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TranslateSearch"),
		Version:     dara.String("2018-10-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TranslateSearchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
