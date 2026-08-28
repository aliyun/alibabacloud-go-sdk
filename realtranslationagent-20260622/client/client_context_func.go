// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Cancels a translation task that is currently running.
//
// Description:
//
// *Billing description**
//
// After the task is successfully canceled, the Credits frozen for this translation task will be fully refunded to your account.
//
// **Before you begin**
//
// - This operation only supports canceling translation tasks that are in the processing state. Tasks that are completed or failed cannot be canceled.
//
// @param request - CancelTranslationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTranslationTaskResponse
func (client *Client) CancelTranslationTaskWithContext(ctx context.Context, request *CancelTranslationTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelTranslationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTranslationTask"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTranslationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of the original file for a translation task.
//
// @param request - GetOriginalFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOriginalFileUrlResponse
func (client *Client) GetOriginalFileUrlWithContext(ctx context.Context, request *GetOriginalFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetOriginalFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOriginalFileUrl"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOriginalFileUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download URL of the translated file for a translation task.
//
// @param request - GetTranslatedFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslatedFileUrlResponse
func (client *Client) GetTranslatedFileUrlWithContext(ctx context.Context, request *GetTranslatedFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetTranslatedFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslatedFileUrl"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslatedFileUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a translation task.
//
// @param request - GetTranslationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslationTaskResponse
func (client *Client) GetTranslationTaskWithContext(ctx context.Context, request *GetTranslationTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTranslationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranslationTask"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranslationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries translation tasks by paging.
//
// @param request - ListTranslationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranslationTasksResponse
func (client *Client) ListTranslationTasksWithContext(ctx context.Context, request *ListTranslationTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTranslationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OriginalFileName) {
		query["OriginalFileName"] = request.OriginalFileName
	}

	if !dara.IsNil(request.SourceLanguage) {
		query["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetLanguage) {
		query["TargetLanguage"] = request.TargetLanguage
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTranslationTasks"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTranslationTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a translation task. You can submit a new translation task by passing in a TaskId, or resubmit a historical task for translation by passing in a BaseTaskId. After successful submission, the translation task ID and current task status are returned. You can use the task ID to call subsequent operations to query translation progress and results.
//
// Description:
//
// *Billing description**
//
// This operation involves Credits consumption. Before submitting a translation task, ensure that your account has sufficient Credits balance. After calling `UploadTranslationFile`, you can check the `CreditsAvailable` field in the response to confirm whether your current balance meets the requirements of this translation task. For detailed billing information, refer to the `CreditBreakdown` field.
//
// **Task submission description**
//
// - To submit a new translation task, pass in the `TaskId` returned by the `UploadTranslationFile` operation.
//
// - To resubmit a historical task for translation, pass in the task ID of a previously submitted translation task, which is the `BaseTaskId`.
//
// - You must pass in either `TaskId` or `BaseTaskId`. You cannot pass in both at the same time.
//
// **Precautions**
//
// - The `Style` parameter takes effect only when the translation file is a PPT file. Passing in this parameter for files in other formats has no effect.
//
// - For new tasks, you can obtain the list of available fonts from the `Fonts` field in the response of `UploadTranslationFile`. For retranslation of historical tasks, you can obtain the list of available fonts by calling the `GetTranslationTask` operation.
//
// @param tmpReq - SubmitTranslationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTranslationTaskResponse
func (client *Client) SubmitTranslationTaskWithContext(ctx context.Context, tmpReq *SubmitTranslationTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitTranslationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitTranslationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomTerms) {
		request.CustomTermsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomTerms, dara.String("CustomTerms"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.CustomTermsShrink) {
		query["CustomTerms"] = request.CustomTermsShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseTaskId) {
		body["BaseTaskId"] = request.BaseTaskId
	}

	if !dara.IsNil(request.ConfigShrink) {
		body["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTranslationTask"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTranslationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a document, parses document-related information, and generates a translation task. After a successful upload, the task ID and document parsing results are returned, including word count, page count, estimated Credits consumption, estimated translation time, detected language type, and font list. The system also performs sensitive information detection on the uploaded document, and you can decide whether to proceed with submitting the translation task based on the detection results.
//
// Description:
//
// > - This operation only involves document upload and information estimation. **No fees are incurred.*	- Credits consumption starts only after you **officially submit the translation*	- task.
//
// **Language detection**
//
// The system automatically detects the language type of the uploaded document. Currently, Chinese is supported.
//
// **Sensitive information detection**
//
// The system performs sensitive information detection on the uploaded document. If sensitive information is detected, the `SensitiveDetected` field in the response is set to `true`, and the `SensitiveTags` field returns the list of matched keywords.
//
// >  - You can decide whether to proceed with submitting the translation task based on your actual needs.
//
// >  - If the translation quality setting is set to ultimate mode when you submit the task, the system automatically switches the **portions containing sensitive information*	- to auto mode.
//
// **Notes**
//
// - Make sure the uploaded document format is supported by the system. Otherwise, parsing may fail.
//
// - The `EstimatedCostCredits` value in the response is the estimated Credits consumption. The actual consumption is based on the settlement after the translation task is officially submitted.
//
// - The `EstimatedTime` value in the response is the estimated translation duration in milliseconds. The actual translation duration may vary depending on document complexity.
//
// - The `Fonts` field in the response contains the languages that support font modification and the corresponding font lists. You can select an appropriate font based on the target language.
//
// @param request - UploadTranslationFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadTranslationFileResponse
func (client *Client) UploadTranslationFileWithContext(ctx context.Context, request *UploadTranslationFileRequest, runtime *dara.RuntimeOptions) (_result *UploadTranslationFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		query["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.File) {
		query["File"] = request.File
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadTranslationFile"),
		Version:     dara.String("2026-06-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadTranslationFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
