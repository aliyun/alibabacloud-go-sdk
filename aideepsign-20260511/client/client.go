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
	EnableValidate  *bool
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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("aideepsign"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Stops a batch C2PA trusted signing task for a specified OSS bucket.
//
// Description:
//
// ## Operation description
//
// - This operation stops a batch trusted signing task for a bucket owned by the currently authenticated account. TaskId is the parent task ID of the bucket returned by CreateC2paBucketSignTask.
//
// - When the returned Status is `CANCELING`, the task is completing the stop process. Repeated requests may return the completed `CANCELED` status.
//
// - Call GetC2paSignResult to query the task status and final result.
//
// @param request - CancelC2paBucketSignTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelC2paBucketSignTaskResponse
func (client *Client) CancelC2paBucketSignTaskWithOptions(request *CancelC2paBucketSignTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelC2paBucketSignTaskResponse, _err error) {
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
		Action:      dara.String("CancelC2paBucketSignTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelC2paBucketSignTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a batch C2PA trusted signing task for a specified OSS bucket.
//
// Description:
//
// ## Operation description
//
// - This operation stops a batch trusted signing task for a bucket owned by the currently authenticated account. TaskId is the parent task ID of the bucket returned by CreateC2paBucketSignTask.
//
// - When the returned Status is `CANCELING`, the task is completing the stop process. Repeated requests may return the completed `CANCELED` status.
//
// - Call GetC2paSignResult to query the task status and final result.
//
// @param request - CancelC2paBucketSignTaskRequest
//
// @return CancelC2paBucketSignTaskResponse
func (client *Client) CancelC2paBucketSignTask(request *CancelC2paBucketSignTaskRequest) (_result *CancelC2paBucketSignTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelC2paBucketSignTaskResponse{}
	_body, _err := client.CancelC2paBucketSignTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a batch C2PA content credential signing task for objects in a specified OSS bucket.
//
// Description:
//
// ## Operation description
//
// - This operation creates a batch C2PA content credential signing task for files under a specified source prefix in an OSS bucket. It supports JPEG, JPG, PNG, MP4, FLAC, M4A, MP3, and WAV formats, and allows mixed processing under the same prefix.
//
// - `SourcePrefix` must be `ai-deepsign/` or its subdirectory, and must not overlap with the output prefix `ai-deepsign-signed/`.
//
// - By default, a single task processes up to 20,000 non-directory objects. The size of a single file must not exceed 800 MiB.
//
// - The RAM identity of the caller must have permissions to call this operation. The service-linked role `AliyunServiceRoleForAIDeepSign` must have access permissions to the specified bucket and associated paths.
//
// - File name extensions are case-insensitive. Files in unsupported formats are recorded as failed, but this does not abort the processing of other files.
//
// - The signing results are written to the `ai-deepsign-signed/<TaskId>/` directory in the same bucket. The subdirectories relative to `SourcePrefix` and the source file name extensions are preserved, and `_signed` is added before the file name extension. The processing report is `c2pa_report.json` in the task directory.
//
// - For example, if `SourcePrefix=ai-deepsign/images/`, the signing result for the source file `ai-deepsign/images/audio/demo.m4a` is `ai-deepsign-signed/<TaskId>/audio/demo_signed.m4a`.
//
// - A successful creation only indicates that the task has been accepted. Call `GetC2paSignResult` to query the results. A `SUCCEEDED` status for a batch task does not mean that every file was signed successfully. Check the final statistics and processing report for file-by-file results.
//
// @param request - CreateC2paBucketSignTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateC2paBucketSignTaskResponse
func (client *Client) CreateC2paBucketSignTaskWithOptions(request *CreateC2paBucketSignTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateC2paBucketSignTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.BucketRegionId) {
		query["BucketRegionId"] = request.BucketRegionId
	}

	if !dara.IsNil(request.SoftwareAgent) {
		query["SoftwareAgent"] = request.SoftwareAgent
	}

	if !dara.IsNil(request.SourcePrefix) {
		query["SourcePrefix"] = request.SourcePrefix
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateC2paBucketSignTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateC2paBucketSignTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a batch C2PA content credential signing task for objects in a specified OSS bucket.
//
// Description:
//
// ## Operation description
//
// - This operation creates a batch C2PA content credential signing task for files under a specified source prefix in an OSS bucket. It supports JPEG, JPG, PNG, MP4, FLAC, M4A, MP3, and WAV formats, and allows mixed processing under the same prefix.
//
// - `SourcePrefix` must be `ai-deepsign/` or its subdirectory, and must not overlap with the output prefix `ai-deepsign-signed/`.
//
// - By default, a single task processes up to 20,000 non-directory objects. The size of a single file must not exceed 800 MiB.
//
// - The RAM identity of the caller must have permissions to call this operation. The service-linked role `AliyunServiceRoleForAIDeepSign` must have access permissions to the specified bucket and associated paths.
//
// - File name extensions are case-insensitive. Files in unsupported formats are recorded as failed, but this does not abort the processing of other files.
//
// - The signing results are written to the `ai-deepsign-signed/<TaskId>/` directory in the same bucket. The subdirectories relative to `SourcePrefix` and the source file name extensions are preserved, and `_signed` is added before the file name extension. The processing report is `c2pa_report.json` in the task directory.
//
// - For example, if `SourcePrefix=ai-deepsign/images/`, the signing result for the source file `ai-deepsign/images/audio/demo.m4a` is `ai-deepsign-signed/<TaskId>/audio/demo_signed.m4a`.
//
// - A successful creation only indicates that the task has been accepted. Call `GetC2paSignResult` to query the results. A `SUCCEEDED` status for a batch task does not mean that every file was signed successfully. Check the final statistics and processing report for file-by-file results.
//
// @param request - CreateC2paBucketSignTaskRequest
//
// @return CreateC2paBucketSignTaskResponse
func (client *Client) CreateC2paBucketSignTask(request *CreateC2paBucketSignTaskRequest) (_result *CreateC2paBucketSignTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateC2paBucketSignTaskResponse{}
	_body, _err := client.CreateC2paBucketSignTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a C2PA content credential signing task for a specified file and returns the task ID.
//
// Description:
//
// ## Request description
//
// - This operation adds standard C2PA content credentials to a single JPEG/JPG, PNG, MP4, FLAC, M4A, MP3, or WAV file.
//
// - `FileUrl` must be an absolute HTTP or HTTPS URL accessible by the service and cannot contain a username or password.
//
// - File name extensions are case-insensitive, and URL query parameters are not used for file name identification. For example, the file name for `https://example.com/media/demo.m4a?token=xxx` is `demo.m4a`.
//
// - The source URL must directly return an HTTP 200 status code. Redirect URLs are not supported.
//
// - The size of a single file cannot exceed 800 MiB. Provide a valid file in a supported format. Changing only the file name extension does not transform the file format.
//
// - The signing result retains the source file name extension and adds `_signed` before the extension. For example, `demo.m4a` becomes `demo_signed.m4a`.
//
// - Unsupported file name extensions return `InvalidParam`, and the error message is `FileUrl only supports files with .jpg, .jpeg, .png, .mp4, .flac, .m4a, .mp3, or .wav extensions`.
//
// - A successful creation only indicates that the task has been accepted. Save the returned `TaskId` and invoke `GetC2paSignResult` to query the task status and final result.
//
// @param request - CreateC2paSignTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateC2paSignTaskResponse
func (client *Client) CreateC2paSignTaskWithOptions(request *CreateC2paSignTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateC2paSignTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.SoftwareAgent) {
		query["SoftwareAgent"] = request.SoftwareAgent
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateC2paSignTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateC2paSignTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a C2PA content credential signing task for a specified file and returns the task ID.
//
// Description:
//
// ## Request description
//
// - This operation adds standard C2PA content credentials to a single JPEG/JPG, PNG, MP4, FLAC, M4A, MP3, or WAV file.
//
// - `FileUrl` must be an absolute HTTP or HTTPS URL accessible by the service and cannot contain a username or password.
//
// - File name extensions are case-insensitive, and URL query parameters are not used for file name identification. For example, the file name for `https://example.com/media/demo.m4a?token=xxx` is `demo.m4a`.
//
// - The source URL must directly return an HTTP 200 status code. Redirect URLs are not supported.
//
// - The size of a single file cannot exceed 800 MiB. Provide a valid file in a supported format. Changing only the file name extension does not transform the file format.
//
// - The signing result retains the source file name extension and adds `_signed` before the extension. For example, `demo.m4a` becomes `demo_signed.m4a`.
//
// - Unsupported file name extensions return `InvalidParam`, and the error message is `FileUrl only supports files with .jpg, .jpeg, .png, .mp4, .flac, .m4a, .mp3, or .wav extensions`.
//
// - A successful creation only indicates that the task has been accepted. Save the returned `TaskId` and invoke `GetC2paSignResult` to query the task status and final result.
//
// @param request - CreateC2paSignTaskRequest
//
// @return CreateC2paSignTaskResponse
func (client *Client) CreateC2paSignTask(request *CreateC2paSignTaskRequest) (_result *CreateC2paSignTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateC2paSignTaskResponse{}
	_body, _err := client.CreateC2paSignTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an image for defect detection and returns a task ID for querying results.
//
// Description:
//
// ## Request description
//
// - This operation commits an image to the server for defect detection, including but not limited to blurriness, distortion, glare, and other anomalies in the image subject.
//
// - **ImageUrl is required*	- and must be a publicly accessible HTTP/HTTPS URL. Currently, only image URLs accessible from the Chinese mainland public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL of the image as the input parameter for this operation. References: [How to obtain the URL of a single object or the URLs of multiple objects](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - After successful commit, a node ID is returned immediately. Regardless of whether the node has been completed or is still being processed, call the GetImageDefectDetectionTaskResult operation to perform polling for the final detection result.
//
// - If a parameter fault or other issue occurs during the commit procedure, the corresponding error message is returned directly without creating a node record.
//
// - The image file size cannot exceed 10 MB. Detection is failed for images that exceed this limit.
//
// @param request - CreateImageDefectDetectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageDefectDetectionTaskResponse
func (client *Client) CreateImageDefectDetectionTaskWithOptions(request *CreateImageDefectDetectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageDefectDetectionTaskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageDefectDetectionTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageDefectDetectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an image for defect detection and returns a task ID for querying results.
//
// Description:
//
// ## Request description
//
// - This operation commits an image to the server for defect detection, including but not limited to blurriness, distortion, glare, and other anomalies in the image subject.
//
// - **ImageUrl is required*	- and must be a publicly accessible HTTP/HTTPS URL. Currently, only image URLs accessible from the Chinese mainland public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL of the image as the input parameter for this operation. References: [How to obtain the URL of a single object or the URLs of multiple objects](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - After successful commit, a node ID is returned immediately. Regardless of whether the node has been completed or is still being processed, call the GetImageDefectDetectionTaskResult operation to perform polling for the final detection result.
//
// - If a parameter fault or other issue occurs during the commit procedure, the corresponding error message is returned directly without creating a node record.
//
// - The image file size cannot exceed 10 MB. Detection is failed for images that exceed this limit.
//
// @param request - CreateImageDefectDetectionTaskRequest
//
// @return CreateImageDefectDetectionTaskResponse
func (client *Client) CreateImageDefectDetectionTask(request *CreateImageDefectDetectionTaskRequest) (_result *CreateImageDefectDetectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageDefectDetectionTaskResponse{}
	_body, _err := client.CreateImageDefectDetectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an asynchronous image tampering detection task and returns the task ID.
//
// Description:
//
// ## Request description
//
// - **ImageUrl*	- is required and must be a publicly accessible HTTP/HTTPS URL.
//
// - External images are not currently supported. Only image URLs accessible from the Chinese mainland public network are supported. Upload images to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL as the input parameter for this operation. References: [How to obtain the URL of a single object or the URLs of multiple objects](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects).
//
// - The **DetectType*	- parameter is deprecated. This parameter is still accepted (`tamper`/`aigc`) but does not affect the actual detection pattern. Regardless of the value passed, tampering detection is executed and `DetectMode=tamper` is returned. New callers should not pass this parameter.
//
// - The **UseTamperSmallModel*	- parameter specifies whether to use the small model. If set to `true`, the small model is used and no large model report is generated. If set to `false` or not passed, the large model is used by default and a comprehensive large model analysis report is returned.
//
// ## Settings
//
// - When **UseTamperSmallModel*	- is not specified, the large model is used by default.
//
// @param request - CreateImageDetectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageDetectionTaskResponse
func (client *Client) CreateImageDetectionTaskWithOptions(request *CreateImageDetectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageDetectionTaskResponse, _err error) {
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

	if !dara.IsNil(request.DetectType) {
		query["DetectType"] = request.DetectType
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.UseTamperSmallModel) {
		query["UseTamperSmallModel"] = request.UseTamperSmallModel
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an asynchronous image tampering detection task and returns the task ID.
//
// Description:
//
// ## Request description
//
// - **ImageUrl*	- is required and must be a publicly accessible HTTP/HTTPS URL.
//
// - External images are not currently supported. Only image URLs accessible from the Chinese mainland public network are supported. Upload images to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL as the input parameter for this operation. References: [How to obtain the URL of a single object or the URLs of multiple objects](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects).
//
// - The **DetectType*	- parameter is deprecated. This parameter is still accepted (`tamper`/`aigc`) but does not affect the actual detection pattern. Regardless of the value passed, tampering detection is executed and `DetectMode=tamper` is returned. New callers should not pass this parameter.
//
// - The **UseTamperSmallModel*	- parameter specifies whether to use the small model. If set to `true`, the small model is used and no large model report is generated. If set to `false` or not passed, the large model is used by default and a comprehensive large model analysis report is returned.
//
// ## Settings
//
// - When **UseTamperSmallModel*	- is not specified, the large model is used by default.
//
// @param request - CreateImageDetectionTaskRequest
//
// @return CreateImageDetectionTaskResponse
func (client *Client) CreateImageDetectionTask(request *CreateImageDetectionTaskRequest) (_result *CreateImageDetectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageDetectionTaskResponse{}
	_body, _err := client.CreateImageDetectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a personal income tax record screenshot for forgery detection and returns a task ID.
//
// Description:
//
// ## Request description
//
// - **ImageUrl*	- is required and must be a publicly accessible HTTP/HTTPS URL (the server needs to download the image).
//
// - External images are not supported. Only image URLs accessible from the Chinese mainland public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL of the image as the input parameter of this operation. References: [How do I obtain the URL of one or more files?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - The image file size cannot exceed 10 MB. Detection of images that exceed this limit will be failed.
//
// @param request - CreateImageIncomeTaxDetectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageIncomeTaxDetectionTaskResponse
func (client *Client) CreateImageIncomeTaxDetectionTaskWithOptions(request *CreateImageIncomeTaxDetectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageIncomeTaxDetectionTaskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageIncomeTaxDetectionTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageIncomeTaxDetectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a personal income tax record screenshot for forgery detection and returns a task ID.
//
// Description:
//
// ## Request description
//
// - **ImageUrl*	- is required and must be a publicly accessible HTTP/HTTPS URL (the server needs to download the image).
//
// - External images are not supported. Only image URLs accessible from the Chinese mainland public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL of the image as the input parameter of this operation. References: [How do I obtain the URL of one or more files?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - The image file size cannot exceed 10 MB. Detection of images that exceed this limit will be failed.
//
// @param request - CreateImageIncomeTaxDetectionTaskRequest
//
// @return CreateImageIncomeTaxDetectionTaskResponse
func (client *Client) CreateImageIncomeTaxDetectionTask(request *CreateImageIncomeTaxDetectionTaskRequest) (_result *CreateImageIncomeTaxDetectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageIncomeTaxDetectionTaskResponse{}
	_body, _err := client.CreateImageIncomeTaxDetectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateImageTaskWithOptions(request *CreateImageTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateImageTaskResponse
func (client *Client) CreateImageTask(request *CreateImageTaskRequest) (_result *CreateImageTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageTaskResponse{}
	_body, _err := client.CreateImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - **ImageUrl*	- is required and must be a publicly accessible HTTP/HTTPS URL.
//
// - External images are not supported. Only image URLs accessible from the Chinese mainland public network are supported. Upload images to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the pre-signed URL of the image as the input parameter of this operation. References: [How to obtain the URL of a single file or the URLs of multiple files](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects).
//
// - The size of the image uploaded by using ImageUrl cannot exceed 10 MB.
//
// - The optional parameter **ClientToken*	- ensures the idempotence of the request. The client generates this parameter value, which must be unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters.
//
// @param request - CreateSensitiveScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSensitiveScanTaskResponse
func (client *Client) CreateSensitiveScanTaskWithOptions(request *CreateSensitiveScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSensitiveScanTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - **ImageUrl*	- is required and must be a publicly accessible HTTP/HTTPS URL.
//
// - External images are not supported. Only image URLs accessible from the Chinese mainland public network are supported. Upload images to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the pre-signed URL of the image as the input parameter of this operation. References: [How to obtain the URL of a single file or the URLs of multiple files](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects).
//
// - The size of the image uploaded by using ImageUrl cannot exceed 10 MB.
//
// - The optional parameter **ClientToken*	- ensures the idempotence of the request. The client generates this parameter value, which must be unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters.
//
// @param request - CreateSensitiveScanTaskRequest
//
// @return CreateSensitiveScanTaskResponse
func (client *Client) CreateSensitiveScanTask(request *CreateSensitiveScanTaskRequest) (_result *CreateSensitiveScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSensitiveScanTaskResponse{}
	_body, _err := client.CreateSensitiveScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a video generation task.
//
// Description:
//
// ## Request description
//
// - This operation creates an AI video generation task. The system generates a video based on the positive prompt provided by the user.
//
// - By default, the generated video is automatically embedded with a C2PA digital signature.
//
// - After creating the task, use the `TaskId` to call the `GetVideoTaskResult` operation to query the generation result.
//
// @param request - CreateVideoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoTaskResponse
func (client *Client) CreateVideoTaskWithOptions(request *CreateVideoTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVideoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Audio) {
		query["Audio"] = request.Audio
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptExtend) {
		query["PromptExtend"] = request.PromptExtend
	}

	if !dara.IsNil(request.Ratio) {
		query["Ratio"] = request.Ratio
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.Seed) {
		query["Seed"] = request.Seed
	}

	if !dara.IsNil(request.Watermark) {
		query["Watermark"] = request.Watermark
	}

	if !dara.IsNil(request.Media) {
		query["Media"] = request.Media
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a video generation task.
//
// Description:
//
// ## Request description
//
// - This operation creates an AI video generation task. The system generates a video based on the positive prompt provided by the user.
//
// - By default, the generated video is automatically embedded with a C2PA digital signature.
//
// - After creating the task, use the `TaskId` to call the `GetVideoTaskResult` operation to query the generation result.
//
// @param request - CreateVideoTaskRequest
//
// @return CreateVideoTaskResponse
func (client *Client) CreateVideoTask(request *CreateVideoTaskRequest) (_result *CreateVideoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVideoTaskResponse{}
	_body, _err := client.CreateVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an image or video file for invisible watermark embedding and returns a task ID for querying the result.
//
// Description:
//
// ## Operation description
//
// - **FileUrl*	- must be a publicly accessible HTTP/HTTPS URL. External files are not supported. Only file URLs accessible over the Chinese mainland public network are supported. Upload the file to an Alibaba Cloud Object Storage Service (OSS) bucket in a Chinese mainland region first, and then use the presigned URL of the file as the input parameter for this operation. References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - **FileType*	- must match the actual file type. The same value must be specified during both embedding and fetch.
//
// - Supported file specifications: Video files must be no longer than 10 minutes and no larger than 100 MB. Image files must be no larger than 100 MB.
//
// - The task completion time is approximately equal to the file duration. For example, embedding a 10-second video takes about 10 seconds.
//
// - If the submission fails, no task record is created and a synchronous error is returned.
//
// @param request - CreateWatermarkEmbedTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWatermarkEmbedTaskResponse
func (client *Client) CreateWatermarkEmbedTaskWithOptions(request *CreateWatermarkEmbedTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateWatermarkEmbedTaskResponse, _err error) {
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

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Reserved) {
		query["Reserved"] = request.Reserved
	}

	if !dara.IsNil(request.WatermarkContent) {
		query["WatermarkContent"] = request.WatermarkContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWatermarkEmbedTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWatermarkEmbedTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an image or video file for invisible watermark embedding and returns a task ID for querying the result.
//
// Description:
//
// ## Operation description
//
// - **FileUrl*	- must be a publicly accessible HTTP/HTTPS URL. External files are not supported. Only file URLs accessible over the Chinese mainland public network are supported. Upload the file to an Alibaba Cloud Object Storage Service (OSS) bucket in a Chinese mainland region first, and then use the presigned URL of the file as the input parameter for this operation. References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - **FileType*	- must match the actual file type. The same value must be specified during both embedding and fetch.
//
// - Supported file specifications: Video files must be no longer than 10 minutes and no larger than 100 MB. Image files must be no larger than 100 MB.
//
// - The task completion time is approximately equal to the file duration. For example, embedding a 10-second video takes about 10 seconds.
//
// - If the submission fails, no task record is created and a synchronous error is returned.
//
// @param request - CreateWatermarkEmbedTaskRequest
//
// @return CreateWatermarkEmbedTaskResponse
func (client *Client) CreateWatermarkEmbedTask(request *CreateWatermarkEmbedTaskRequest) (_result *CreateWatermarkEmbedTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWatermarkEmbedTaskResponse{}
	_body, _err := client.CreateWatermarkEmbedTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an image or video file to extract a dark watermark and returns a task ID.
//
// Description:
//
// ## Request description
//
// - **FileUrl*	- must be a publicly accessible HTTP/HTTPS URL. External files are not supported. Only file URLs accessible from the Chinese mainland public network are supported. Upload the file to an Alibaba Cloud OSS bucket in a region within the Chinese mainland first, and then use the presigned URL of the file as the input parameter of this operation. References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - **FileType must be the same as the one used during embedding**: Make sure that the file type (image or video) for extraction is the same as the file type used when the watermark was embedded.
//
// - **File specification limits**: Video files must be no longer than 10 minutes and no larger than 100 MB. Image files must also be no larger than 100 MB.
//
// - **Processing time**: Approximately equal to the actual playback duration of the file. For example, a 10-second video takes approximately 10 seconds to complete the watermark fetch procedure.
//
// @param request - CreateWatermarkExtractTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWatermarkExtractTaskResponse
func (client *Client) CreateWatermarkExtractTaskWithOptions(request *CreateWatermarkExtractTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateWatermarkExtractTaskResponse, _err error) {
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

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Reserved) {
		query["Reserved"] = request.Reserved
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWatermarkExtractTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWatermarkExtractTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an image or video file to extract a dark watermark and returns a task ID.
//
// Description:
//
// ## Request description
//
// - **FileUrl*	- must be a publicly accessible HTTP/HTTPS URL. External files are not supported. Only file URLs accessible from the Chinese mainland public network are supported. Upload the file to an Alibaba Cloud OSS bucket in a region within the Chinese mainland first, and then use the presigned URL of the file as the input parameter of this operation. References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// - **FileType must be the same as the one used during embedding**: Make sure that the file type (image or video) for extraction is the same as the file type used when the watermark was embedded.
//
// - **File specification limits**: Video files must be no longer than 10 minutes and no larger than 100 MB. Image files must also be no larger than 100 MB.
//
// - **Processing time**: Approximately equal to the actual playback duration of the file. For example, a 10-second video takes approximately 10 seconds to complete the watermark fetch procedure.
//
// @param request - CreateWatermarkExtractTaskRequest
//
// @return CreateWatermarkExtractTaskResponse
func (client *Client) CreateWatermarkExtractTask(request *CreateWatermarkExtractTaskRequest) (_result *CreateWatermarkExtractTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWatermarkExtractTaskResponse{}
	_body, _err := client.CreateWatermarkExtractTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Determines whether an image is AI-generated and returns detection labels with confidence scores.
//
// Description:
//
// ## Request description
//
// - This operation detects whether a specified image is AI-generated content (AIGC). You can specify the image to detect by using an image URL or an OSS ObjectKey.
//
// - You must provide at least one of ImageUrl and ObjectKey. If both are provided, ObjectKey takes precedence.
//
// - When using the ObjectKey method, the system verifies whether the ObjectKey belongs to the current caller. Cross-tenant access is not allowed.
//
// - This is a synchronous operation suitable for real-time detection of a single image. If you need asynchronous detection or credential detection at the same time, use the CreateImageDetectionTask operation.
//
// @param request - DetectAigcImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectAigcImageResponse
func (client *Client) DetectAigcImageWithOptions(request *DetectAigcImageRequest, runtime *dara.RuntimeOptions) (_result *DetectAigcImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Determines whether an image is AI-generated and returns detection labels with confidence scores.
//
// Description:
//
// ## Request description
//
// - This operation detects whether a specified image is AI-generated content (AIGC). You can specify the image to detect by using an image URL or an OSS ObjectKey.
//
// - You must provide at least one of ImageUrl and ObjectKey. If both are provided, ObjectKey takes precedence.
//
// - When using the ObjectKey method, the system verifies whether the ObjectKey belongs to the current caller. Cross-tenant access is not allowed.
//
// - This is a synchronous operation suitable for real-time detection of a single image. If you need asynchronous detection or credential detection at the same time, use the CreateImageDetectionTask operation.
//
// @param request - DetectAigcImageRequest
//
// @return DetectAigcImageResponse
func (client *Client) DetectAigcImage(request *DetectAigcImageRequest) (_result *DetectAigcImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectAigcImageResponse{}
	_body, _err := client.DetectAigcImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - This operation detects and returns basic information about an image, including but not limited to the file name, image format (such as JPEG/PNG), file size, and resolution.
//
// - `ImageUrl` is required and must be a publicly accessible HTTP/HTTPS URL. Currently, only image URLs accessible from the Chinese mainland public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL of the image as the input parameter for this operation. References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// @param request - DetectImageBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageBasicInfoResponse
func (client *Client) DetectImageBasicInfoWithOptions(request *DetectImageBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *DetectImageBasicInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - This operation detects and returns basic information about an image, including but not limited to the file name, image format (such as JPEG/PNG), file size, and resolution.
//
// - `ImageUrl` is required and must be a publicly accessible HTTP/HTTPS URL. Currently, only image URLs accessible from the Chinese mainland public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a Chinese mainland region first, and then use the presigned URL of the image as the input parameter for this operation. References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// @param request - DetectImageBasicInfoRequest
//
// @return DetectImageBasicInfoResponse
func (client *Client) DetectImageBasicInfo(request *DetectImageBasicInfoRequest) (_result *DetectImageBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageBasicInfoResponse{}
	_body, _err := client.DetectImageBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and results of a single-file or OSS Bucket batch C2PA signing task.
//
// Description:
//
// ## Request description
//
// - This operation queries the status and final results of nodes created by using CreateC2paSignTask or CreateC2paBucketSignTask.
//
// - It supports signing node queries for images, videos, and various audio formats (FLAC, M4A, MP3, and WAV).
//
// - For BUCKET nodes, intermediate statistics information is not provided when the node is in the PROCESSING, PAUSING, PAUSED, or CANCELING state. Complete statistics information and the report location are returned only after the node enters the SUCCEEDED or CANCELED state.
//
// - A valid TaskId must be provided for the query. This ID must be the value returned when the node was created. It supports the standard UUID format and is compatible with legacy decimal node IDs.
//
// @param request - GetC2paSignResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetC2paSignResultResponse
func (client *Client) GetC2paSignResultWithOptions(request *GetC2paSignResultRequest, runtime *dara.RuntimeOptions) (_result *GetC2paSignResultResponse, _err error) {
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
		Action:      dara.String("GetC2paSignResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetC2paSignResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and results of a single-file or OSS Bucket batch C2PA signing task.
//
// Description:
//
// ## Request description
//
// - This operation queries the status and final results of nodes created by using CreateC2paSignTask or CreateC2paBucketSignTask.
//
// - It supports signing node queries for images, videos, and various audio formats (FLAC, M4A, MP3, and WAV).
//
// - For BUCKET nodes, intermediate statistics information is not provided when the node is in the PROCESSING, PAUSING, PAUSED, or CANCELING state. Complete statistics information and the report location are returned only after the node enters the SUCCEEDED or CANCELED state.
//
// - A valid TaskId must be provided for the query. This ID must be the value returned when the node was created. It supports the standard UUID format and is compatible with legacy decimal node IDs.
//
// @param request - GetC2paSignResultRequest
//
// @return GetC2paSignResultResponse
func (client *Client) GetC2paSignResult(request *GetC2paSignResultRequest) (_result *GetC2paSignResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetC2paSignResultResponse{}
	_body, _err := client.GetC2paSignResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and result of an image defect detection task, including the abnormality score and reason.
//
// Description:
//
// ## Request description
//
// This operation queries the execution status and result of an image defect detection task created by `CreateImageDefectDetectionTask`. Follow these polling conventions:
//
//   - **Set the polling interval to at least 5 seconds**: When the task is in the `running` state, each call to this operation triggers a real downstream query. Excessively frequent polling only increases system load and does not speed up result retrieval.
//
//   - **Total timeout is 30 minutes**: If no result is received within 30 minutes after the task is created, the server marks the task as `failed` with the message `poll timeout: no result within 30m0s`. After that, the task no longer initiates new downstream query requests.
//
//   - After the task reaches the `succeeded` or `failed` desired state, subsequent calls to this operation directly return the stored result without triggering additional downstream processing.
//
// Note: The content in the `Result` field may be HTML-escaped JSON data. Unescape it before parsing.
//
// @param request - GetImageDefectDetectionTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDefectDetectionTaskResultResponse
func (client *Client) GetImageDefectDetectionTaskResultWithOptions(request *GetImageDefectDetectionTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageDefectDetectionTaskResultResponse, _err error) {
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
		Action:      dara.String("GetImageDefectDetectionTaskResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageDefectDetectionTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and result of an image defect detection task, including the abnormality score and reason.
//
// Description:
//
// ## Request description
//
// This operation queries the execution status and result of an image defect detection task created by `CreateImageDefectDetectionTask`. Follow these polling conventions:
//
//   - **Set the polling interval to at least 5 seconds**: When the task is in the `running` state, each call to this operation triggers a real downstream query. Excessively frequent polling only increases system load and does not speed up result retrieval.
//
//   - **Total timeout is 30 minutes**: If no result is received within 30 minutes after the task is created, the server marks the task as `failed` with the message `poll timeout: no result within 30m0s`. After that, the task no longer initiates new downstream query requests.
//
//   - After the task reaches the `succeeded` or `failed` desired state, subsequent calls to this operation directly return the stored result without triggering additional downstream processing.
//
// Note: The content in the `Result` field may be HTML-escaped JSON data. Unescape it before parsing.
//
// @param request - GetImageDefectDetectionTaskResultRequest
//
// @return GetImageDefectDetectionTaskResultResponse
func (client *Client) GetImageDefectDetectionTaskResult(request *GetImageDefectDetectionTaskResultRequest) (_result *GetImageDefectDetectionTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageDefectDetectionTaskResultResponse{}
	_body, _err := client.GetImageDefectDetectionTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and result of an image detection task created by CreateImageDetectionTask.
//
// Description:
//
// ## Description
//
// Call this operation to query the execution status and result of a specified image detection task, including AIGC detection labels and tamper detection results. Poll at intervals of 2 to 5 seconds until the task status changes to succeeded or failed.
//
// ## Precautions
//
//   - Ensure that you use a valid TaskId for the query.
//
//   - If the task is not complete, increase the polling interval to avoid wasting resources due to frequent requests.
//
// @param request - GetImageDetectionTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDetectionTaskResultResponse
func (client *Client) GetImageDetectionTaskResultWithOptions(request *GetImageDetectionTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageDetectionTaskResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and result of an image detection task created by CreateImageDetectionTask.
//
// Description:
//
// ## Description
//
// Call this operation to query the execution status and result of a specified image detection task, including AIGC detection labels and tamper detection results. Poll at intervals of 2 to 5 seconds until the task status changes to succeeded or failed.
//
// ## Precautions
//
//   - Ensure that you use a valid TaskId for the query.
//
//   - If the task is not complete, increase the polling interval to avoid wasting resources due to frequent requests.
//
// @param request - GetImageDetectionTaskResultRequest
//
// @return GetImageDetectionTaskResultResponse
func (client *Client) GetImageDetectionTaskResult(request *GetImageDetectionTaskResultRequest) (_result *GetImageDetectionTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageDetectionTaskResultResponse{}
	_body, _err := client.GetImageDetectionTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and list of abnormal items for an individual income tax screenshot detection task.
//
// Description:
//
// ## Operation description
//
// - **Polling recommendation**: Set the polling interval to at least 5 seconds. Polling more frequently while the task is in `running` status does not return results faster.
//
// - **Total timeout**: If no result is obtained within 30 minutes after the task is created, the task is marked as `failed` with the Message `poll timeout: no result within 30m0s`.
//
// - **Desired state handling**: `succeeded` and `failed` are the final states of the task. The result does not change after the task reaches either state.
//
// - **Items is an empty array**: This indicates that no abnormal items are detected. It is not an error.
//
// - **Coordinates may be an empty array**: Some abnormalities are image-wide in nature (for example, quality issues affecting the entire image) and cannot be located within a specific bounding box. In this case, `Coordinates` is an empty array, but the abnormal item itself is still valid and should still be presented to the user.
//
// - **Stage is an open value set**: Future versions may introduce new `Stage` values. Display unknown values as-is. Do not perform whitelist validation on `Stage` or filter abnormal items based on it.
//
// - **quality stage**: A `Stage` value of `quality` indicates that the image quality is too low and the detection result has low confidence. Prompt the user to resubmit a clear original screenshot and run the detection again.
//
// @param request - GetImageIncomeTaxDetectionTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageIncomeTaxDetectionTaskResultResponse
func (client *Client) GetImageIncomeTaxDetectionTaskResultWithOptions(request *GetImageIncomeTaxDetectionTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageIncomeTaxDetectionTaskResultResponse, _err error) {
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
		Action:      dara.String("GetImageIncomeTaxDetectionTaskResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageIncomeTaxDetectionTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and list of abnormal items for an individual income tax screenshot detection task.
//
// Description:
//
// ## Operation description
//
// - **Polling recommendation**: Set the polling interval to at least 5 seconds. Polling more frequently while the task is in `running` status does not return results faster.
//
// - **Total timeout**: If no result is obtained within 30 minutes after the task is created, the task is marked as `failed` with the Message `poll timeout: no result within 30m0s`.
//
// - **Desired state handling**: `succeeded` and `failed` are the final states of the task. The result does not change after the task reaches either state.
//
// - **Items is an empty array**: This indicates that no abnormal items are detected. It is not an error.
//
// - **Coordinates may be an empty array**: Some abnormalities are image-wide in nature (for example, quality issues affecting the entire image) and cannot be located within a specific bounding box. In this case, `Coordinates` is an empty array, but the abnormal item itself is still valid and should still be presented to the user.
//
// - **Stage is an open value set**: Future versions may introduce new `Stage` values. Display unknown values as-is. Do not perform whitelist validation on `Stage` or filter abnormal items based on it.
//
// - **quality stage**: A `Stage` value of `quality` indicates that the image quality is too low and the detection result has low confidence. Prompt the user to resubmit a clear original screenshot and run the detection again.
//
// @param request - GetImageIncomeTaxDetectionTaskResultRequest
//
// @return GetImageIncomeTaxDetectionTaskResultResponse
func (client *Client) GetImageIncomeTaxDetectionTaskResult(request *GetImageIncomeTaxDetectionTaskResultRequest) (_result *GetImageIncomeTaxDetectionTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageIncomeTaxDetectionTaskResultResponse{}
	_body, _err := client.GetImageIncomeTaxDetectionTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetImageTaskResultWithOptions(request *GetImageTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageTaskResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetImageTaskResultResponse
func (client *Client) GetImageTaskResult(request *GetImageTaskResultRequest) (_result *GetImageTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageTaskResultResponse{}
	_body, _err := client.GetImageTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSensitiveScanResultWithOptions(request *GetSensitiveScanResultRequest, runtime *dara.RuntimeOptions) (_result *GetSensitiveScanResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSensitiveScanResultResponse
func (client *Client) GetSensitiveScanResult(request *GetSensitiveScanResultRequest) (_result *GetSensitiveScanResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSensitiveScanResultResponse{}
	_body, _err := client.GetSensitiveScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the task result of a video generation node.
//
// Description:
//
// ## Request description
//
// - Call this operation to query the execution status and result of an asynchronous video generation task created by calling `CreateVideoTask`.
//
// - Poll the task status at an interval of 2 to 5 seconds until the task status changes to `succeeded` or `failed`.
//
// @param request - GetVideoTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoTaskResultResponse
func (client *Client) GetVideoTaskResultWithOptions(request *GetVideoTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetVideoTaskResultResponse, _err error) {
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
		Action:      dara.String("GetVideoTaskResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task result of a video generation node.
//
// Description:
//
// ## Request description
//
// - Call this operation to query the execution status and result of an asynchronous video generation task created by calling `CreateVideoTask`.
//
// - Poll the task status at an interval of 2 to 5 seconds until the task status changes to `succeeded` or `failed`.
//
// @param request - GetVideoTaskResultRequest
//
// @return GetVideoTaskResultResponse
func (client *Client) GetVideoTaskResult(request *GetVideoTaskResultRequest) (_result *GetVideoTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoTaskResultResponse{}
	_body, _err := client.GetVideoTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and download URL of an invisible watermark embedding task.
//
// Description:
//
// ## Operation description
//
// - This operation queries the execution status and result of an invisible watermark embedding task created by `CreateWatermarkEmbedTask`.
//
// - Tasks are processed asynchronously on the server side. Call this operation to retrieve the latest status.
//
// - **Polling recommendations**: Ensure that the polling interval is at least 5 seconds. For video tasks, the expected completion time is approximately equal to the media duration (about 1x speed). Image tasks are completed within 5 minutes.
//
// - **Timeout mechanism**: The total timeout window for video tasks is 10 minutes, and for image tasks is 5 minutes. Tasks that are not completed within this period are marked as failed, and the corresponding error message is returned.
//
// - When the task status is `succeeded` or `failed`, the task has reached a desired state, and the result no longer changes.
//
// @param request - GetWatermarkEmbedTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWatermarkEmbedTaskResultResponse
func (client *Client) GetWatermarkEmbedTaskResultWithOptions(request *GetWatermarkEmbedTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetWatermarkEmbedTaskResultResponse, _err error) {
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
		Action:      dara.String("GetWatermarkEmbedTaskResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWatermarkEmbedTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and download URL of an invisible watermark embedding task.
//
// Description:
//
// ## Operation description
//
// - This operation queries the execution status and result of an invisible watermark embedding task created by `CreateWatermarkEmbedTask`.
//
// - Tasks are processed asynchronously on the server side. Call this operation to retrieve the latest status.
//
// - **Polling recommendations**: Ensure that the polling interval is at least 5 seconds. For video tasks, the expected completion time is approximately equal to the media duration (about 1x speed). Image tasks are completed within 5 minutes.
//
// - **Timeout mechanism**: The total timeout window for video tasks is 10 minutes, and for image tasks is 5 minutes. Tasks that are not completed within this period are marked as failed, and the corresponding error message is returned.
//
// - When the task status is `succeeded` or `failed`, the task has reached a desired state, and the result no longer changes.
//
// @param request - GetWatermarkEmbedTaskResultRequest
//
// @return GetWatermarkEmbedTaskResultResponse
func (client *Client) GetWatermarkEmbedTaskResult(request *GetWatermarkEmbedTaskResultRequest) (_result *GetWatermarkEmbedTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWatermarkEmbedTaskResultResponse{}
	_body, _err := client.GetWatermarkEmbedTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and extraction result of a dark watermark extraction task.
//
// Description:
//
// ## Operation description
//
// - **Polling recommendation**: Set the polling interval to at least 5 seconds. The task completion time is approximately equal to the duration of the carrier media (approximately 1x speed).
//
// - **Timeout handling**: The total timeout window for video tasks is 10 minutes, and for image tasks is 5 minutes. If a task is not completed within this time, the task status changes to `failed`.
//
// - **Desired state confirmation**: `succeeded` or `failed` is the final status. The result does not change after the task reaches either status.
//
// @param request - GetWatermarkExtractTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWatermarkExtractTaskResultResponse
func (client *Client) GetWatermarkExtractTaskResultWithOptions(request *GetWatermarkExtractTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetWatermarkExtractTaskResultResponse, _err error) {
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
		Action:      dara.String("GetWatermarkExtractTaskResult"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWatermarkExtractTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and extraction result of a dark watermark extraction task.
//
// Description:
//
// ## Operation description
//
// - **Polling recommendation**: Set the polling interval to at least 5 seconds. The task completion time is approximately equal to the duration of the carrier media (approximately 1x speed).
//
// - **Timeout handling**: The total timeout window for video tasks is 10 minutes, and for image tasks is 5 minutes. If a task is not completed within this time, the task status changes to `failed`.
//
// - **Desired state confirmation**: `succeeded` or `failed` is the final status. The result does not change after the task reaches either status.
//
// @param request - GetWatermarkExtractTaskResultRequest
//
// @return GetWatermarkExtractTaskResultResponse
func (client *Client) GetWatermarkExtractTaskResult(request *GetWatermarkExtractTaskResultRequest) (_result *GetWatermarkExtractTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWatermarkExtractTaskResultResponse{}
	_body, _err := client.GetWatermarkExtractTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses a batch C2PA trusted signing task for a specified OSS bucket.
//
// Description:
//
// ## Operation description
//
// - This operation pauses a batch trusted signing parent task for a bucket that is owned and still controllable by the currently authenticated account.
//
// - The paused state is persisted to the database and remains effective after a service restart.
//
// - If no files have been submitted to the trusted signing service, the task directly enters the `PAUSED` state. If submitted files exist, the task first enters the `PAUSING` state, stops scanning and new submissions, and continues to query the actual results of those files. After all results converge, the task enters the `PAUSED` state.
//
// - Pausing does not interrupt an in-progress individual OSS or trusted signing service invoke, nor does it revoke signatures or metering that have already been completed.
//
// - Invoking this operation repeatedly on a task that is already in the `PAUSING` or `PAUSED` state is idempotent and succeeds.
//
// - Tasks in the `SUCCEEDED`, `FAILED`, `CANCELING`, or `CANCELED` state cannot be paused.
//
// @param request - PauseC2paBucketSignTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseC2paBucketSignTaskResponse
func (client *Client) PauseC2paBucketSignTaskWithOptions(request *PauseC2paBucketSignTaskRequest, runtime *dara.RuntimeOptions) (_result *PauseC2paBucketSignTaskResponse, _err error) {
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
		Action:      dara.String("PauseC2paBucketSignTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseC2paBucketSignTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a batch C2PA trusted signing task for a specified OSS bucket.
//
// Description:
//
// ## Operation description
//
// - This operation pauses a batch trusted signing parent task for a bucket that is owned and still controllable by the currently authenticated account.
//
// - The paused state is persisted to the database and remains effective after a service restart.
//
// - If no files have been submitted to the trusted signing service, the task directly enters the `PAUSED` state. If submitted files exist, the task first enters the `PAUSING` state, stops scanning and new submissions, and continues to query the actual results of those files. After all results converge, the task enters the `PAUSED` state.
//
// - Pausing does not interrupt an in-progress individual OSS or trusted signing service invoke, nor does it revoke signatures or metering that have already been completed.
//
// - Invoking this operation repeatedly on a task that is already in the `PAUSING` or `PAUSED` state is idempotent and succeeds.
//
// - Tasks in the `SUCCEEDED`, `FAILED`, `CANCELING`, or `CANCELED` state cannot be paused.
//
// @param request - PauseC2paBucketSignTaskRequest
//
// @return PauseC2paBucketSignTaskResponse
func (client *Client) PauseC2paBucketSignTask(request *PauseC2paBucketSignTaskRequest) (_result *PauseC2paBucketSignTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseC2paBucketSignTaskResponse{}
	_body, _err := client.PauseC2paBucketSignTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a paused or pausing OSS Bucket batch C2PA trusted signing task.
//
// Description:
//
// ## Operation description
//
// - This operation resumes a Bucket batch trusted signing parent task owned by the currently authenticated account.
//
// - After resumption, the task continues execution from the saved scan cursor and per-file stage. It does not rebuild the object list or resubmit trusted signing service tasks that have already been persisted.
//
// - This operation supports resuming from the `PAUSED` or `PAUSING` state. For tasks already in the `PROCESSING` state, repeated calls are idempotent and succeed.
//
// - If the task was previously auto-paused due to failures in shared dependencies such as user OSS, internal OSS, or the trusted signing service, the pause error recorded on the parent task is cleared upon resumption. If the dependencies have not yet recovered, the task may enter the `PAUSING` or `PAUSED` state again with a new reason recorded.
//
// - Before resuming a task that was paused due to permission issues, restore the OSS permissions of `AliyunServiceRoleForAIDeepSign` for the source prefix and output prefix. Otherwise, the task will be auto-paused again.
//
// - Tasks in the `SUCCEEDED`, `FAILED`, `CANCELING`, or `CANCELED` state cannot be resumed.
//
// @param request - ResumeC2paBucketSignTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeC2paBucketSignTaskResponse
func (client *Client) ResumeC2paBucketSignTaskWithOptions(request *ResumeC2paBucketSignTaskRequest, runtime *dara.RuntimeOptions) (_result *ResumeC2paBucketSignTaskResponse, _err error) {
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
		Action:      dara.String("ResumeC2paBucketSignTask"),
		Version:     dara.String("2026-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeC2paBucketSignTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a paused or pausing OSS Bucket batch C2PA trusted signing task.
//
// Description:
//
// ## Operation description
//
// - This operation resumes a Bucket batch trusted signing parent task owned by the currently authenticated account.
//
// - After resumption, the task continues execution from the saved scan cursor and per-file stage. It does not rebuild the object list or resubmit trusted signing service tasks that have already been persisted.
//
// - This operation supports resuming from the `PAUSED` or `PAUSING` state. For tasks already in the `PROCESSING` state, repeated calls are idempotent and succeed.
//
// - If the task was previously auto-paused due to failures in shared dependencies such as user OSS, internal OSS, or the trusted signing service, the pause error recorded on the parent task is cleared upon resumption. If the dependencies have not yet recovered, the task may enter the `PAUSING` or `PAUSED` state again with a new reason recorded.
//
// - Before resuming a task that was paused due to permission issues, restore the OSS permissions of `AliyunServiceRoleForAIDeepSign` for the source prefix and output prefix. Otherwise, the task will be auto-paused again.
//
// - Tasks in the `SUCCEEDED`, `FAILED`, `CANCELING`, or `CANCELED` state cannot be resumed.
//
// @param request - ResumeC2paBucketSignTaskRequest
//
// @return ResumeC2paBucketSignTaskResponse
func (client *Client) ResumeC2paBucketSignTask(request *ResumeC2paBucketSignTaskRequest) (_result *ResumeC2paBucketSignTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeC2paBucketSignTaskResponse{}
	_body, _err := client.ResumeC2paBucketSignTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Signs an image with a basic C2PA signature by using a platform-owned certificate, supports content integrity verification, and returns the download URL of the signed image.
//
// Description:
//
// ## Signature types
//
// - **Basic signature**: Signs with a platform-owned certificate and supports content integrity verification.
//
// - **Trusted signature**: In addition to integrity verification, the signing certificate chain can be verified through the C2PA official trust list.
//
// This operation provides basic signing. For trusted signing, call `CreateC2paSignTask` to create a task and call `GetC2paSignResult` to query the result.
//
// ## Request description
//
// - `ImageUrl` is required and must be a publicly accessible HTTP/HTTPS URL.
//
// - External image URLs are not currently supported. Only image URLs that are publicly accessible within the Chinese mainland are supported. Upload the image to an Alibaba Cloud OSS bucket in a region within the Chinese mainland first, and then use the pre-signed URL of the image as the input parameter of this operation. References: [How to obtain the URL of a single file or the URLs of multiple files](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects).
//
// - Supported image formats: JPEG and PNG. Unsupported formats return the C2PA_FORMAT_UNSUPPORTED fault.
//
// @param request - SignUserImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignUserImageResponse
func (client *Client) SignUserImageWithOptions(request *SignUserImageRequest, runtime *dara.RuntimeOptions) (_result *SignUserImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Signs an image with a basic C2PA signature by using a platform-owned certificate, supports content integrity verification, and returns the download URL of the signed image.
//
// Description:
//
// ## Signature types
//
// - **Basic signature**: Signs with a platform-owned certificate and supports content integrity verification.
//
// - **Trusted signature**: In addition to integrity verification, the signing certificate chain can be verified through the C2PA official trust list.
//
// This operation provides basic signing. For trusted signing, call `CreateC2paSignTask` to create a task and call `GetC2paSignResult` to query the result.
//
// ## Request description
//
// - `ImageUrl` is required and must be a publicly accessible HTTP/HTTPS URL.
//
// - External image URLs are not currently supported. Only image URLs that are publicly accessible within the Chinese mainland are supported. Upload the image to an Alibaba Cloud OSS bucket in a region within the Chinese mainland first, and then use the pre-signed URL of the image as the input parameter of this operation. References: [How to obtain the URL of a single file or the URLs of multiple files](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects).
//
// - Supported image formats: JPEG and PNG. Unsupported formats return the C2PA_FORMAT_UNSUPPORTED fault.
//
// @param request - SignUserImageRequest
//
// @return SignUserImageResponse
func (client *Client) SignUserImage(request *SignUserImageRequest) (_result *SignUserImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SignUserImageResponse{}
	_body, _err := client.SignUserImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the C2PA signature in an image and returns the content integrity verification status and issuer certificate trust information.
//
// Description:
//
// ## Signature types
//
// - **Basic signature**: Uses the platform\\"s own certificate for signing. Supports content integrity verification.
//
// - **Trusted signature**: In addition to integrity verification, the signing certificate chain can be verified through the C2PA official trust list.
//
// ## Verification result description
//
// - Content integrity verification checks whether the image and its signature-protected provenance information remain consistent with the state at the time of signing. It does not indicate that the visual content itself is authentic.
//
// - `VerificationState=Valid` indicates that the signature and data hash are valid and content integrity verification has passed. This alone cannot be used to determine that the signature is a trusted signature.
//
// - `VerificationState=Trusted` and `IssuerTrusted` present the certificate trust verification results. Whether a signature qualifies as a trusted signature as defined in this topic depends on whether the signing certificate chain can be verified through the C2PA official trust list, not solely on the certificate issuer name.
//
// ## Request description
//
// - This operation verifies the C2PA (Coalition for Content Provenance and Authenticity) digital signature embedded in a specified image.
//
// - `ImageUrl` **does not currently support images from external networks. Only image URLs accessible over the China public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a China region first, and then use the presigned URL of the image as the input parameter for this operation.*	- References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// @param request - VerifyImageSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyImageSignatureResponse
func (client *Client) VerifyImageSignatureWithOptions(request *VerifyImageSignatureRequest, runtime *dara.RuntimeOptions) (_result *VerifyImageSignatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the C2PA signature in an image and returns the content integrity verification status and issuer certificate trust information.
//
// Description:
//
// ## Signature types
//
// - **Basic signature**: Uses the platform\\"s own certificate for signing. Supports content integrity verification.
//
// - **Trusted signature**: In addition to integrity verification, the signing certificate chain can be verified through the C2PA official trust list.
//
// ## Verification result description
//
// - Content integrity verification checks whether the image and its signature-protected provenance information remain consistent with the state at the time of signing. It does not indicate that the visual content itself is authentic.
//
// - `VerificationState=Valid` indicates that the signature and data hash are valid and content integrity verification has passed. This alone cannot be used to determine that the signature is a trusted signature.
//
// - `VerificationState=Trusted` and `IssuerTrusted` present the certificate trust verification results. Whether a signature qualifies as a trusted signature as defined in this topic depends on whether the signing certificate chain can be verified through the C2PA official trust list, not solely on the certificate issuer name.
//
// ## Request description
//
// - This operation verifies the C2PA (Coalition for Content Provenance and Authenticity) digital signature embedded in a specified image.
//
// - `ImageUrl` **does not currently support images from external networks. Only image URLs accessible over the China public network are supported. Upload the image to an Alibaba Cloud OSS bucket in a China region first, and then use the presigned URL of the image as the input parameter for this operation.*	- References: [How do I obtain the URL of one or more objects?](https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects)
//
// @param request - VerifyImageSignatureRequest
//
// @return VerifyImageSignatureResponse
func (client *Client) VerifyImageSignature(request *VerifyImageSignatureRequest) (_result *VerifyImageSignatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyImageSignatureResponse{}
	_body, _err := client.VerifyImageSignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
