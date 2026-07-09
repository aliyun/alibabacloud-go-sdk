// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates an asynchronous image detection task that supports AIGC and tampering detection.
//
// Description:
//
// ## Operation description
//
// - This operation creates an asynchronous image detection task that supports automatic classification, AIGC detection, and tampering detection.
//
// - You must specify at least one of `ImageUrl` and `ObjectKey`. If both are specified, `ObjectKey` takes precedence.
//
// - If you set `DetectType` to `auto`, the system automatically determines whether to perform AIGC detection or tampering detection based on the image content.
//
// @param request - CreateImageDetectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageDetectionTaskResponse
func (client *Client) CreateImageDetectionTaskWithContext(ctx context.Context, request *CreateImageDetectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageDetectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.DetectType) {
		query["DetectType"] = request.DetectType
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageDetectionTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageDetectionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates images based on a provided text description and returns a task ID.
//
// Description:
//
// ## Operation description
//
// - This operation creates an AI image generation task. The system generates images based on the positive prompt provided by the user.
//
// - You can configure parameters such as negative prompt, model, and image size to optimize the generation results.
//
// - By default, generated images are automatically embedded with a C2PA digital signature. You can optionally add a watermark in the lower-right corner.
//
// - Set the ClientToken parameter to ensure the idempotence of the request and guarantee uniqueness across different requests.
//
// - After the task is created, call the GetImageTaskResult operation with the TaskId to query the generation results.
//
// @param request - CreateImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageTaskResponse
func (client *Client) CreateImageTaskWithContext(ctx context.Context, request *CreateImageTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		query["N"] = request.N
	}

	if !dara.IsNil(request.NegativePrompt) {
		query["NegativePrompt"] = request.NegativePrompt
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptExtend) {
		query["PromptExtend"] = request.PromptExtend
	}

	if !dara.IsNil(request.Seed) {
		query["Seed"] = request.Seed
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Watermark) {
		query["Watermark"] = request.Watermark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image sensitive information scan task and returns the task ID.
//
// Description:
//
// ## Operation description
//
// - This operation creates an image sensitive information scan task. The system performs sensitive data identification on the specified image.
//
// - You can specify the image to scan by using an image URL or an OSS ObjectKey.
//
// - The image size cannot exceed 10 MB.
//
// - You must specify at least one of ImageUrl and ObjectKey. If both are specified, ObjectKey takes precedence.
//
// - When you use ObjectKey, make sure that the key belongs to the namespace of the current caller. Cross-tenant access is not allowed.
//
// - You can use the ClientToken parameter to ensure the idempotence of the request.
//
// @param request - CreateSensitiveScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSensitiveScanTaskResponse
func (client *Client) CreateSensitiveScanTaskWithContext(ctx context.Context, request *CreateSensitiveScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSensitiveScanTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSensitiveScanTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSensitiveScanTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Determines whether an image is AI-generated and returns detection labels and confidence levels.
//
// Description:
//
// ## Operation description
//
// - This operation detects whether a specified image is AI-generated content (AIGC). You can specify the image to detect by using an image URL or an OSS ObjectKey.
//
// - You must provide at least one of ImageUrl and ObjectKey. If both are provided, ObjectKey takes precedence.
//
// - When you use the ObjectKey method, the system verifies whether the ObjectKey belongs to the current caller. Cross-tenant access is not allowed.
//
// - This is a synchronous operation suitable for real-time detection of a single image. To perform asynchronous detection or credential detection at the same time, use the CreateImageDetectionTask operation.
//
// @param request - DetectAigcImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectAigcImageResponse
func (client *Client) DetectAigcImageWithContext(ctx context.Context, request *DetectAigcImageRequest, runtime *dara.RuntimeOptions) (_result *DetectAigcImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectAigcImage"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectAigcImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves basic information about a specified image, such as the file name, format, size, and resolution.
//
// Description:
//
// ## Description
//
// - This operation detects and returns basic information about an image, including but not limited to the file name, image format (such as JPEG or PNG), file size, and resolution.
//
// - You can specify the image to detect by providing an image URL or an OSS ObjectKey. If both ImageUrl and ObjectKey are provided, ObjectKey takes precedence.
//
// - When using ObjectKey, ensure that the object belongs to the namespace of the current caller. Cross-tenant access is not allowed.
//
// - The optional parameter ClientToken ensures the idempotence of the request. Generate a new unique value for each request.
//
// @param request - DetectImageBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageBasicInfoResponse
func (client *Client) DetectImageBasicInfoWithContext(ctx context.Context, request *DetectImageBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *DetectImageBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageBasicInfo"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageBasicInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and results of an image detection task, including AIGC detection labels and tamper detection results.
//
// Description:
//
// ## Operation description
//
// Call this operation to query the execution status and results of an asynchronous detection task created by `CreateImageDetectionTask`. Poll at intervals of 2 to 5 seconds until the task status changes to `succeeded` or `failed`.
//
// ### Before you begin
//
// - Use a valid `TaskId` for the query.
//
// - If the task is not complete, increase the polling interval to avoid unnecessary resource consumption caused by frequent requests.
//
// @param request - GetImageDetectionTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDetectionTaskResultResponse
func (client *Client) GetImageDetectionTaskResultWithContext(ctx context.Context, request *GetImageDetectionTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageDetectionTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageDetectionTaskResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageDetectionTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and result of an asynchronous image generation task and retrieves the image download URL.
//
// Description:
//
// ## Operation description
//
// - Call this operation to query the execution status and result of an asynchronous image generation task created by `CreateImageTask`.
//
// - Poll at intervals of 2 to 5 seconds until the task status changes to `succeeded` or `failed`.
//
// - The image download URL (Url) returned after the task succeeds is a pre-signed URL that is valid for 1 hour.
//
// - To ensure idempotence of the request, set the `ClientToken` parameter.
//
// @param request - GetImageTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTaskResultResponse
func (client *Client) GetImageTaskResultWithContext(ctx context.Context, request *GetImageTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageTaskResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and result of a sensitive information scan task and returns the details of the sensitive data discovered during the scan.
//
// Description:
//
// ## Operation description
//
// - Call this operation to query the execution status and result of a sensitive information scan task created by `CreateSensitiveScanTask`.
//
// - Poll at intervals of 3 to 5 seconds until the task status changes to `completed` or `terminated`.
//
// - The `ClientToken` parameter ensures the idempotence of the request. It is generated by the client, must be unique across different requests, supports ASCII characters, and cannot exceed 64 characters in length.
//
// @param request - GetSensitiveScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSensitiveScanResultResponse
func (client *Client) GetSensitiveScanResultWithContext(ctx context.Context, request *GetSensitiveScanResultRequest, runtime *dara.RuntimeOptions) (_result *GetSensitiveScanResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSensitiveScanResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSensitiveScanResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Embeds a C2PA digital signature into a user-uploaded image and returns the download URL.
//
// Description:
//
// ## Operation description
//
// - Specify at least one of `ImageUrl` and `ObjectKey`. If both are specified, `ObjectKey` takes precedence.
//
// - When you use `ObjectKey`, the system verifies that the `ObjectKey` belongs to the current caller. Cross-tenant access is not allowed.
//
// - Supported image formats are JPEG and PNG. Unsupported formats return the `C2PA_FORMAT_UNSUPPORTED` error.
//
// - If the original image already contains a C2PA signature, the system retains the original signature as an ingredient and appends a new signature.
//
// - Use the `ClientToken` parameter to ensure idempotence. Make sure the value is unique across different requests and does not exceed 64 characters.
//
// @param request - SignUserImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignUserImageResponse
func (client *Client) SignUserImageWithContext(ctx context.Context, request *SignUserImageRequest, runtime *dara.RuntimeOptions) (_result *SignUserImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SignUserImage"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SignUserImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the C2PA digital signature in an image and returns the signature status and issuer information.
//
// Description:
//
// ## Operation description
//
// - This operation verifies the C2PA digital signature embedded in an image and returns the signature verification status, issuer trustworthiness, issuer information, and the complete content credentials manifest.
//
// - You can specify the image to verify by using an image URL or an OSS ObjectKey. If both ImageUrl and ObjectKey are provided, ObjectKey takes precedence.
//
// - When you use the ObjectKey method, the system verifies whether the ObjectKey belongs to the current caller. Cross-tenant access is not allowed.
//
// - To ensure request idempotency, provide the ClientToken parameter. Ensure that the value is unique across different requests and does not exceed 64 characters.
//
// @param request - VerifyImageSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyImageSignatureResponse
func (client *Client) VerifyImageSignatureWithContext(ctx context.Context, request *VerifyImageSignatureRequest, runtime *dara.RuntimeOptions) (_result *VerifyImageSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyImageSignature"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyImageSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
