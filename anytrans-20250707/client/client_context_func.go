// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 通义多模态翻译批量翻译
//
// @param tmpReq - BatchTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchTranslateResponse
func (client *Client) BatchTranslateWithContext(ctx context.Context, tmpReq *BatchTranslateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchTranslateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BatchTranslateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("ext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Text) {
		request.TextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Text, dara.String("text"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtShrink) {
		body["ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Format) {
		body["format"] = request.Format
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["targetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.TextShrink) {
		body["text"] = request.TextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchTranslate"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/batch"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译获文档翻译任务
//
// @param request - GetDocTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocTranslateTaskResponse
func (client *Client) GetDocTranslateTaskWithContext(ctx context.Context, request *GetDocTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocTranslateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/doc/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
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

// Summary:
//
// 通义多模态翻译获取html翻译结果
//
// @param request - GetHtmlTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHtmlTranslateTaskResponse
func (client *Client) GetHtmlTranslateTaskWithContext(ctx context.Context, request *GetHtmlTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHtmlTranslateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHtmlTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/html/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHtmlTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译获取图片翻译任务
//
// @param request - GetImageTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTranslateTaskResponse
func (client *Client) GetImageTranslateTaskWithContext(ctx context.Context, request *GetImageTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetImageTranslateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/image/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
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
// 通义多模态翻译获取长文翻译结果
//
// @param request - GetLongTextTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLongTextTranslateTaskResponse
func (client *Client) GetLongTextTranslateTaskWithContext(ctx context.Context, request *GetLongTextTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLongTextTranslateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLongTextTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/longText/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLongTextTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译提交文档翻译任务
//
// @param tmpReq - SubmitDocTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocTranslateTaskResponse
func (client *Client) SubmitDocTranslateTaskWithContext(ctx context.Context, tmpReq *SubmitDocTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocTranslateTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitDocTranslateTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtShrink) {
		body["ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Format) {
		body["format"] = request.Format
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["targetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/doc/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译提交html翻译任务
//
// @param tmpReq - SubmitHtmlTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitHtmlTranslateTaskResponse
func (client *Client) SubmitHtmlTranslateTaskWithContext(ctx context.Context, tmpReq *SubmitHtmlTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitHtmlTranslateTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitHtmlTranslateTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtShrink) {
		body["ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Format) {
		body["format"] = request.Format
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["targetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitHtmlTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/html/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitHtmlTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译提交图片翻译任务
//
// @param tmpReq - SubmitImageTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageTranslateTaskResponse
func (client *Client) SubmitImageTranslateTaskWithContext(ctx context.Context, tmpReq *SubmitImageTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitImageTranslateTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitImageTranslateTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("ext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TargetLanguage) {
		request.TargetLanguageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetLanguage, dara.String("targetLanguage"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtShrink) {
		body["ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Format) {
		body["format"] = request.Format
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguageShrink) {
		body["targetLanguage"] = request.TargetLanguageShrink
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImageTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/image/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImageTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译提交长文翻译任务
//
// @param tmpReq - SubmitLongTextTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitLongTextTranslateTaskResponse
func (client *Client) SubmitLongTextTranslateTaskWithContext(ctx context.Context, tmpReq *SubmitLongTextTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitLongTextTranslateTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitLongTextTranslateTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtShrink) {
		body["ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Format) {
		body["format"] = request.Format
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["targetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitLongTextTranslateTask"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/longText/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitLongTextTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译术语编辑
//
// @param tmpReq - TermEditRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TermEditResponse
func (client *Client) TermEditWithContext(ctx context.Context, tmpReq *TermEditRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TermEditResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TermEditShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.ExtShrink) {
		body["ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["targetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TermEdit"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/intervene/edit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TermEditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译术语查询
//
// @param request - TermQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TermQueryResponse
func (client *Client) TermQueryWithContext(ctx context.Context, request *TermQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TermQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["targetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TermQuery"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/intervene/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TermQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通义多模态翻译文本翻译
//
// @param tmpReq - TextTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextTranslateResponse
func (client *Client) TextTranslateWithContext(ctx context.Context, tmpReq *TextTranslateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TextTranslateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &TextTranslateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtShrink) {
		body["ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Format) {
		body["format"] = request.Format
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.SourceLanguage) {
		body["sourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.TargetLanguage) {
		body["targetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TextTranslate"),
		Version:     dara.String("2025-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/anytrans/translate/text"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TextTranslateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
