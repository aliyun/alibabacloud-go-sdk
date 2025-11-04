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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("bailian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加类目
//
// @param request - AddCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCategoryResponse
func (client *Client) AddCategoryWithOptions(WorkspaceId *string, request *AddCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		body["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.CategoryType) {
		body["CategoryType"] = request.CategoryType
	}

	if !dara.IsNil(request.ParentCategoryId) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCategory"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/category/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加类目
//
// @param request - AddCategoryRequest
//
// @return AddCategoryResponse
func (client *Client) AddCategory(WorkspaceId *string, request *AddCategoryRequest) (_result *AddCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddCategoryResponse{}
	_body, _err := client.AddCategoryWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports an unstructured document stored in the temporary storage space to Data Management. You cannot use the API to import structured documents. Use the console instead.
//
// Description:
//
//	Before you call this operation, make sure that you have obtained the lease and uploaded the document to the temporary storage space by using the [ApplyFileUploadLease](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-applyfileuploadlease) operation. For more information, see [Upload files by calling API](https://www.alibabacloud.com/help/en/model-studio/developer-reference/upload-files-by-calling-api).
//
// >  After you call this operation, the used lease ID expires immediately. Do not use the same lease ID to submit new requests.
//
//   - You must call this operation within 12 hours after you call the [ApplyFileUploadLease](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-applyfileuploadlease) operation. Otherwise, the lease expires and the request fails.
//
//   - After you call this operation, the system parses and imports your document. The process takes some time.
//
//   - This interface is not idempotent.
//
// @param tmpReq - AddFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFileResponse
func (client *Client) AddFileWithOptions(WorkspaceId *string, tmpReq *AddFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CategoryType) {
		body["CategoryType"] = request.CategoryType
	}

	if !dara.IsNil(request.LeaseId) {
		body["LeaseId"] = request.LeaseId
	}

	if !dara.IsNil(request.OriginalFileUrl) {
		body["OriginalFileUrl"] = request.OriginalFileUrl
	}

	if !dara.IsNil(request.Parser) {
		body["Parser"] = request.Parser
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFile"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/file"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports an unstructured document stored in the temporary storage space to Data Management. You cannot use the API to import structured documents. Use the console instead.
//
// Description:
//
//	Before you call this operation, make sure that you have obtained the lease and uploaded the document to the temporary storage space by using the [ApplyFileUploadLease](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-applyfileuploadlease) operation. For more information, see [Upload files by calling API](https://www.alibabacloud.com/help/en/model-studio/developer-reference/upload-files-by-calling-api).
//
// >  After you call this operation, the used lease ID expires immediately. Do not use the same lease ID to submit new requests.
//
//   - You must call this operation within 12 hours after you call the [ApplyFileUploadLease](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-applyfileuploadlease) operation. Otherwise, the lease expires and the request fails.
//
//   - After you call this operation, the system parses and imports your document. The process takes some time.
//
//   - This interface is not idempotent.
//
// @param request - AddFileRequest
//
// @return AddFileResponse
func (client *Client) AddFile(WorkspaceId *string, request *AddFileRequest) (_result *AddFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddFileResponse{}
	_body, _err := client.AddFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将已授权OSS Bucket中的文件添加到百炼应用数据
//
// @param tmpReq - AddFilesFromAuthorizedOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFilesFromAuthorizedOssResponse
func (client *Client) AddFilesFromAuthorizedOssWithOptions(WorkspaceId *string, tmpReq *AddFilesFromAuthorizedOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddFilesFromAuthorizedOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddFilesFromAuthorizedOssShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileDetails) {
		request.FileDetailsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileDetails, dara.String("FileDetails"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CategoryType) {
		body["CategoryType"] = request.CategoryType
	}

	if !dara.IsNil(request.FileDetailsShrink) {
		body["FileDetails"] = request.FileDetailsShrink
	}

	if !dara.IsNil(request.OssBucketName) {
		body["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssRegionId) {
		body["OssRegionId"] = request.OssRegionId
	}

	if !dara.IsNil(request.OverWriteFileByOssKey) {
		body["OverWriteFileByOssKey"] = request.OverWriteFileByOssKey
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFilesFromAuthorizedOss"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/file/fromoss"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFilesFromAuthorizedOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将已授权OSS Bucket中的文件添加到百炼应用数据
//
// @param request - AddFilesFromAuthorizedOssRequest
//
// @return AddFilesFromAuthorizedOssResponse
func (client *Client) AddFilesFromAuthorizedOss(WorkspaceId *string, request *AddFilesFromAuthorizedOssRequest) (_result *AddFilesFromAuthorizedOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddFilesFromAuthorizedOssResponse{}
	_body, _err := client.AddFilesFromAuthorizedOssWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a document upload lease to upload a document.
//
// Description:
//
//	  This operation returns an HTTP URL that can be used to upload an unstructured document (the lease) and parameters required for the upload. Structured documents are not supported.
//
//		- The HTTP URL returned by this operation is valid only for minutes. Upload the document before the URL expires.
//
//		- After you apply for a lease and upload a document, the document is stored in a temporary storage space for 12 hours.
//
//		- This interface is not idempotent.
//
// @param request - ApplyFileUploadLeaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLeaseWithOptions(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApplyFileUploadLeaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryType) {
		body["CategoryType"] = request.CategoryType
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Md5) {
		body["Md5"] = request.Md5
	}

	if !dara.IsNil(request.SizeInBytes) {
		body["SizeInBytes"] = request.SizeInBytes
	}

	if !dara.IsNil(request.UseInternalEndpoint) {
		body["UseInternalEndpoint"] = request.UseInternalEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyFileUploadLease"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/category/" + dara.PercentEncode(dara.StringValue(CategoryId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a document upload lease to upload a document.
//
// Description:
//
//	  This operation returns an HTTP URL that can be used to upload an unstructured document (the lease) and parameters required for the upload. Structured documents are not supported.
//
//		- The HTTP URL returned by this operation is valid only for minutes. Upload the document before the URL expires.
//
//		- After you apply for a lease and upload a document, the document is stored in a temporary storage space for 12 hours.
//
//		- This interface is not idempotent.
//
// @param request - ApplyFileUploadLeaseRequest
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLease(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest) (_result *ApplyFileUploadLeaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.ApplyFileUploadLeaseWithOptions(CategoryId, WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请临时文件存储上传许可
//
// @param request - ApplyTempStorageLeaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyTempStorageLeaseResponse
func (client *Client) ApplyTempStorageLeaseWithOptions(WorkspaceId *string, request *ApplyTempStorageLeaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApplyTempStorageLeaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.SizeInBytes) {
		body["SizeInBytes"] = request.SizeInBytes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyTempStorageLease"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyTempStorageLeaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请临时文件存储上传许可
//
// @param request - ApplyTempStorageLeaseRequest
//
// @return ApplyTempStorageLeaseResponse
func (client *Client) ApplyTempStorageLease(WorkspaceId *string, request *ApplyTempStorageLeaseRequest) (_result *ApplyTempStorageLeaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyTempStorageLeaseResponse{}
	_body, _err := client.ApplyTempStorageLeaseWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configure the parsing method for a specific file type. For example, use LLM parsing for .pdf files, or use Qwen VL parsing for .jpg files.
//
// @param tmpReq - ChangeParseSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeParseSettingResponse
func (client *Client) ChangeParseSettingWithOptions(WorkspaceId *string, tmpReq *ChangeParseSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeParseSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChangeParseSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ParserConfig) {
		request.ParserConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParserConfig, dara.String("ParserConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.Parser) {
		body["Parser"] = request.Parser
	}

	if !dara.IsNil(request.ParserConfigShrink) {
		body["ParserConfig"] = request.ParserConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeParseSetting"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/parser/settings"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeParseSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configure the parsing method for a specific file type. For example, use LLM parsing for .pdf files, or use Qwen VL parsing for .jpg files.
//
// @param request - ChangeParseSettingRequest
//
// @return ChangeParseSettingResponse
func (client *Client) ChangeParseSetting(WorkspaceId *string, request *ChangeParseSettingRequest) (_result *ChangeParseSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeParseSettingResponse{}
	_body, _err := client.ChangeParseSettingWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建并发布智能体应用
//
// @param tmpReq - CreateAndPulishAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndPulishAgentResponse
func (client *Client) CreateAndPulishAgentWithOptions(workspaceId *string, tmpReq *CreateAndPulishAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAndPulishAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAndPulishAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationConfig) {
		request.ApplicationConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationConfig, dara.String("applicationConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SampleLibrary) {
		request.SampleLibraryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SampleLibrary, dara.String("sampleLibrary"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigShrink) {
		body["applicationConfig"] = request.ApplicationConfigShrink
	}

	if !dara.IsNil(request.Instructions) {
		body["instructions"] = request.Instructions
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SampleLibraryShrink) {
		body["sampleLibrary"] = request.SampleLibraryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndPulishAgent"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/application/agents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndPulishAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建并发布智能体应用
//
// @param request - CreateAndPulishAgentRequest
//
// @return CreateAndPulishAgentResponse
func (client *Client) CreateAndPulishAgent(workspaceId *string, request *CreateAndPulishAgentRequest) (_result *CreateAndPulishAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAndPulishAgentResponse{}
	_body, _err := client.CreateAndPulishAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a knowledge base of the document search type.
//
// Description:
//
//	  **Limits**: This operation can create only knowledge base of the document search type. Data query and image Q\\&A types are not supported. Use the console instead.
//
//		- **Required permissions**
//
//	    	- **RAM users**: Must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) of Model Studio (such as the `AliyunBailianDataFullAccess` policy, which includes the sfm:CreateIndex permission required), and [become member of a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//	    	- **Alibaba Cloud account**: Has the permission by default, and can call the operation directly.
//
//		- **Call method**: We recommend using the latest version of the [GenAI Service Platform SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK encapsulates complex signature computational logic to simplify the call process.
//
//		- **What to do next**: This operation only initializes knowledge base creation job. After that, call **SubmitIndexJob*	- to complete the creation. Otherwise, you will get an empty knowledge base. For more information about the sample code, see [Knowledge base API guide](https://help.aliyun.com/document_detail/2852772.html).
//
//		- **Idempotence**: This operation is not idempotent. If you call the operation for multiple times, you may create several knowledge bases with the same name. We recommend following a "query first, then create" logic.
//
// **Rate limit:*	- Rate limiting will be triggered if you call this operation frequently. Do not exceed 10 times per second. If limiting is triggered, try again later.
//
// @param tmpReq - CreateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexResponse
func (client *Client) CreateIndexWithOptions(WorkspaceId *string, tmpReq *CreateIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateIndexShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryIds) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, dara.String("CategoryIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Columns) {
		request.ColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Columns, dara.String("Columns"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DataSource) {
		request.DataSourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSource, dara.String("DataSource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocumentIds) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, dara.String("DocumentIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TableIds) {
		request.TableIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableIds, dara.String("TableIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MetaExtractColumns) {
		request.MetaExtractColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetaExtractColumns, dara.String("metaExtractColumns"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryIdsShrink) {
		query["CategoryIds"] = request.CategoryIdsShrink
	}

	if !dara.IsNil(request.ChunkSize) {
		query["ChunkSize"] = request.ChunkSize
	}

	if !dara.IsNil(request.ColumnsShrink) {
		query["Columns"] = request.ColumnsShrink
	}

	if !dara.IsNil(request.CreateIndexType) {
		query["CreateIndexType"] = request.CreateIndexType
	}

	if !dara.IsNil(request.DataSourceShrink) {
		query["DataSource"] = request.DataSourceShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DocumentIdsShrink) {
		query["DocumentIds"] = request.DocumentIdsShrink
	}

	if !dara.IsNil(request.EmbeddingModelName) {
		query["EmbeddingModelName"] = request.EmbeddingModelName
	}

	if !dara.IsNil(request.EnableRewrite) {
		query["EnableRewrite"] = request.EnableRewrite
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OverlapSize) {
		query["OverlapSize"] = request.OverlapSize
	}

	if !dara.IsNil(request.RerankMinScore) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !dara.IsNil(request.RerankModelName) {
		query["RerankModelName"] = request.RerankModelName
	}

	if !dara.IsNil(request.Separator) {
		query["Separator"] = request.Separator
	}

	if !dara.IsNil(request.SinkInstanceId) {
		query["SinkInstanceId"] = request.SinkInstanceId
	}

	if !dara.IsNil(request.SinkRegion) {
		query["SinkRegion"] = request.SinkRegion
	}

	if !dara.IsNil(request.SinkType) {
		query["SinkType"] = request.SinkType
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StructureType) {
		query["StructureType"] = request.StructureType
	}

	if !dara.IsNil(request.TableIdsShrink) {
		query["TableIds"] = request.TableIdsShrink
	}

	if !dara.IsNil(request.ChunkMode) {
		query["chunkMode"] = request.ChunkMode
	}

	if !dara.IsNil(request.EnableHeaders) {
		query["enableHeaders"] = request.EnableHeaders
	}

	if !dara.IsNil(request.MetaExtractColumnsShrink) {
		query["metaExtractColumns"] = request.MetaExtractColumnsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIndex"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a knowledge base of the document search type.
//
// Description:
//
//	  **Limits**: This operation can create only knowledge base of the document search type. Data query and image Q\\&A types are not supported. Use the console instead.
//
//		- **Required permissions**
//
//	    	- **RAM users**: Must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) of Model Studio (such as the `AliyunBailianDataFullAccess` policy, which includes the sfm:CreateIndex permission required), and [become member of a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//	    	- **Alibaba Cloud account**: Has the permission by default, and can call the operation directly.
//
//		- **Call method**: We recommend using the latest version of the [GenAI Service Platform SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK encapsulates complex signature computational logic to simplify the call process.
//
//		- **What to do next**: This operation only initializes knowledge base creation job. After that, call **SubmitIndexJob*	- to complete the creation. Otherwise, you will get an empty knowledge base. For more information about the sample code, see [Knowledge base API guide](https://help.aliyun.com/document_detail/2852772.html).
//
//		- **Idempotence**: This operation is not idempotent. If you call the operation for multiple times, you may create several knowledge bases with the same name. We recommend following a "query first, then create" logic.
//
// **Rate limit:*	- Rate limiting will be triggered if you call this operation frequently. Do not exceed 10 times per second. If limiting is triggered, try again later.
//
// @param request - CreateIndexRequest
//
// @return CreateIndexResponse
func (client *Client) CreateIndex(WorkspaceId *string, request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Memory
//
// @param request - CreateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryResponse
func (client *Client) CreateMemoryWithOptions(workspaceId *string, request *CreateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemory"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Memory
//
// @param request - CreateMemoryRequest
//
// @return CreateMemoryResponse
func (client *Client) CreateMemory(workspaceId *string, request *CreateMemoryRequest) (_result *CreateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryResponse{}
	_body, _err := client.CreateMemoryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建记忆Node
//
// @param request - CreateMemoryNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryNodeResponse
func (client *Client) CreateMemoryNodeWithOptions(workspaceId *string, memoryId *string, request *CreateMemoryNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemoryNode"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId)) + "/memoryNodes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建记忆Node
//
// @param request - CreateMemoryNodeRequest
//
// @return CreateMemoryNodeResponse
func (client *Client) CreateMemoryNode(workspaceId *string, memoryId *string, request *CreateMemoryNodeRequest) (_result *CreateMemoryNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryNodeResponse{}
	_body, _err := client.CreateMemoryNodeWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a prompt template.
//
// @param request - CreatePromptTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePromptTemplateResponse
func (client *Client) CreatePromptTemplateWithOptions(workspaceId *string, request *CreatePromptTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePromptTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePromptTemplate"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/promptTemplates"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePromptTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a prompt template.
//
// @param request - CreatePromptTemplateRequest
//
// @return CreatePromptTemplateResponse
func (client *Client) CreatePromptTemplate(workspaceId *string, request *CreatePromptTemplateRequest) (_result *CreatePromptTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePromptTemplateResponse{}
	_body, _err := client.CreatePromptTemplateWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除智能体
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithOptions(workspaceId *string, appCode *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgent"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/application/agents/" + dara.PercentEncode(dara.StringValue(appCode))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除智能体
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgent(workspaceId *string, appCode *string) (_result *DeleteAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(workspaceId, appCode, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除类目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithOptions(CategoryId *string, WorkspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCategory"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/category/" + dara.PercentEncode(dara.StringValue(CategoryId)) + "/"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除类目
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategory(CategoryId *string, WorkspaceId *string) (_result *DeleteCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(CategoryId, WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified text chunk from a knowledge base. The deleted chunk cannot be retrieved or recalled.
//
// Description:
//
// *
//
// **Warning*	- After a text chunk is deleted, it cannot be restored. Proceed with caution.
//
//   - **Required permissions**:
//
//   - **RAM users**: Must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) of Model Studio (such as the `AliyunBailianDataFullAccess` policy, which includes the sfm:DeleteChunk permission required), and [become member of a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//   - **Alibaba Cloud account**: Has the permission by default, and can call the operation directly.
//
//   - **Call method**: We recommend using the latest version of the [GenAI Service Platform SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK encapsulates complex signature computational logic to simplify the call process.
//
//   - **Delay**: The update takes effect immediately. During peak hours, the update may take place in seconds.
//
//   - **Idempotence**: This operation is idempotent. If you perform a repeated operation on a chunk that has already been deleted, the interface returns a success.
//
// **Rate limit:*	- Rate limiting will be triggered if you call this operation frequently. Do not exceed 10 times per second. If limiting is triggered, try again later.
//
// @param tmpReq - DeleteChunkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChunkResponse
func (client *Client) DeleteChunkWithOptions(WorkspaceId *string, tmpReq *DeleteChunkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteChunkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteChunkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChunkIds) {
		request.ChunkIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChunkIds, dara.String("ChunkIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChunkIdsShrink) {
		query["ChunkIds"] = request.ChunkIdsShrink
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChunk"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/chunk/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChunkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified text chunk from a knowledge base. The deleted chunk cannot be retrieved or recalled.
//
// Description:
//
// *
//
// **Warning*	- After a text chunk is deleted, it cannot be restored. Proceed with caution.
//
//   - **Required permissions**:
//
//   - **RAM users**: Must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) of Model Studio (such as the `AliyunBailianDataFullAccess` policy, which includes the sfm:DeleteChunk permission required), and [become member of a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//   - **Alibaba Cloud account**: Has the permission by default, and can call the operation directly.
//
//   - **Call method**: We recommend using the latest version of the [GenAI Service Platform SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK encapsulates complex signature computational logic to simplify the call process.
//
//   - **Delay**: The update takes effect immediately. During peak hours, the update may take place in seconds.
//
//   - **Idempotence**: This operation is idempotent. If you perform a repeated operation on a chunk that has already been deleted, the interface returns a success.
//
// **Rate limit:*	- Rate limiting will be triggered if you call this operation frequently. Do not exceed 10 times per second. If limiting is triggered, try again later.
//
// @param request - DeleteChunkRequest
//
// @return DeleteChunkResponse
func (client *Client) DeleteChunk(WorkspaceId *string, request *DeleteChunkRequest) (_result *DeleteChunkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteChunkResponse{}
	_body, _err := client.DeleteChunkWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(FileId *string, WorkspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFile"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/file/" + dara.PercentEncode(dara.StringValue(FileId)) + "/"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(FileId *string, WorkspaceId *string) (_result *DeleteFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(FileId, WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified knowledge base permanently.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- If a knowledge base is being called by an application, disassociate the knowledge base before you can delete it. To disassociate the knowledge base, you must use the console. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base).
//
//		- After you delete a knowledge base, it cannot be recovered. We recommend that you proceed with caution.
//
//		- Imported documents are not deleted from the [Data Management](https://bailian.console.aliyun.com/#/data-center) if you call this operation.
//
//		- This interface is idempotent.
//
// @param request - DeleteIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndexWithOptions(WorkspaceId *string, request *DeleteIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndex"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified knowledge base permanently.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- If a knowledge base is being called by an application, disassociate the knowledge base before you can delete it. To disassociate the knowledge base, you must use the console. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base).
//
//		- After you delete a knowledge base, it cannot be recovered. We recommend that you proceed with caution.
//
//		- Imported documents are not deleted from the [Data Management](https://bailian.console.aliyun.com/#/data-center) if you call this operation.
//
//		- This interface is idempotent.
//
// @param request - DeleteIndexRequest
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndex(WorkspaceId *string, request *DeleteIndexRequest) (_result *DeleteIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexResponse{}
	_body, _err := client.DeleteIndexWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes one or more documents from a specified unstructured knowledge base permanently.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- Only documents with the INSERT_ERROR and FINISH states can be deleted. To query the status of documents in a specified knowledge base, call the [ListIndexDocuments](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-listindexdocuments) operation.
//
//		- After you delete a document, it cannot be recovered and the [Retrieve](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-retrieve) operation cannot query information about the document. We recommend that you proceed with caution.
//
//		- Imported documents are not deleted from the [Data Management](https://bailian.console.aliyun.com/#/data-center) if you call this operation.
//
//		- This interface is idempotent.
//
// @param tmpReq - DeleteIndexDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexDocumentResponse
func (client *Client) DeleteIndexDocumentWithOptions(WorkspaceId *string, tmpReq *DeleteIndexDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndexDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteIndexDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocumentIds) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, dara.String("DocumentIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIdsShrink) {
		query["DocumentIds"] = request.DocumentIdsShrink
	}

	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndexDocument"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/delete_index_document"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndexDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more documents from a specified unstructured knowledge base permanently.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- Only documents with the INSERT_ERROR and FINISH states can be deleted. To query the status of documents in a specified knowledge base, call the [ListIndexDocuments](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-listindexdocuments) operation.
//
//		- After you delete a document, it cannot be recovered and the [Retrieve](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-retrieve) operation cannot query information about the document. We recommend that you proceed with caution.
//
//		- Imported documents are not deleted from the [Data Management](https://bailian.console.aliyun.com/#/data-center) if you call this operation.
//
//		- This interface is idempotent.
//
// @param request - DeleteIndexDocumentRequest
//
// @return DeleteIndexDocumentResponse
func (client *Client) DeleteIndexDocument(WorkspaceId *string, request *DeleteIndexDocumentRequest) (_result *DeleteIndexDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexDocumentResponse{}
	_body, _err := client.DeleteIndexDocumentWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除memory
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(workspaceId *string, memoryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemory"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除memory
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(workspaceId *string, memoryId *string) (_result *DeleteMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(workspaceId, memoryId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除记忆Node
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryNodeResponse
func (client *Client) DeleteMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryNodeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemoryNode"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId)) + "/memoryNodes/" + dara.PercentEncode(dara.StringValue(memoryNodeId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除记忆Node
//
// @return DeleteMemoryNodeResponse
func (client *Client) DeleteMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string) (_result *DeleteMemoryNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryNodeResponse{}
	_body, _err := client.DeleteMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a prompt template based on the template ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePromptTemplateResponse
func (client *Client) DeletePromptTemplateWithOptions(workspaceId *string, promptTemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePromptTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePromptTemplate"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/promptTemplates/" + dara.PercentEncode(dara.StringValue(promptTemplateId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePromptTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a prompt template based on the template ID.
//
// @return DeletePromptTemplateResponse
func (client *Client) DeletePromptTemplate(workspaceId *string, promptTemplateId *string) (_result *DeletePromptTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePromptTemplateResponse{}
	_body, _err := client.DeletePromptTemplateWithOptions(workspaceId, promptTemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an unstructured document.
//
// Description:
//
// Before you call this API, make sure that your document is uploaded to the [Data Management](https://bailian.console.aliyun.com/knowledge-base#/data-center) page of Alibaba Cloud Model Studio.
//
//   - You can also call this operation to query unstructured documents that you upload on the [Data Management](https://bailian.console.aliyun.com/knowledge-base#/data-center) page.
//
//   - This operation is idempotent.
//
// **Throttling:*	- Make sure that the interval between the two queries is at least 15 seconds. Otherwise, you may trigger system throttling. If throttling is triggered, try again later.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileResponse
func (client *Client) DescribeFileWithOptions(WorkspaceId *string, FileId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeFileResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFile"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/file/" + dara.PercentEncode(dara.StringValue(FileId)) + "/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an unstructured document.
//
// Description:
//
// Before you call this API, make sure that your document is uploaded to the [Data Management](https://bailian.console.aliyun.com/knowledge-base#/data-center) page of Alibaba Cloud Model Studio.
//
//   - You can also call this operation to query unstructured documents that you upload on the [Data Management](https://bailian.console.aliyun.com/knowledge-base#/data-center) page.
//
//   - This operation is idempotent.
//
// **Throttling:*	- Make sure that the interval between the two queries is at least 15 seconds. Otherwise, you may trigger system throttling. If throttling is triggered, try again later.
//
// @return DescribeFileResponse
func (client *Client) DescribeFile(WorkspaceId *string, FileId *string) (_result *DescribeFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFileResponse{}
	_body, _err := client.DescribeFileWithOptions(WorkspaceId, FileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支付宝打赏状态
//
// @param request - GetAlipayTransferStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlipayTransferStatusResponse
func (client *Client) GetAlipayTransferStatusWithOptions(request *GetAlipayTransferStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlipayTransferStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["code"] = request.Code
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspace_id"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlipayTransferStatus"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/alipay/transfer/status"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlipayTransferStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支付宝打赏状态
//
// @param request - GetAlipayTransferStatusRequest
//
// @return GetAlipayTransferStatusResponse
func (client *Client) GetAlipayTransferStatus(request *GetAlipayTransferStatusRequest) (_result *GetAlipayTransferStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlipayTransferStatusResponse{}
	_body, _err := client.GetAlipayTransferStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 支付宝打赏链接
//
// @param request - GetAlipayUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlipayUrlResponse
func (client *Client) GetAlipayUrlWithOptions(request *GetAlipayUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlipayUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["app_id"] = request.AppId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspace_id"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlipayUrl"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/alipay/transfer/url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 支付宝打赏链接
//
// @param request - GetAlipayUrlRequest
//
// @return GetAlipayUrlResponse
func (client *Client) GetAlipayUrl(request *GetAlipayUrlRequest) (_result *GetAlipayUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.GetAlipayUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all supported parser types based on the input file type (file extension).
//
// @param request - GetAvailableParserTypesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAvailableParserTypesResponse
func (client *Client) GetAvailableParserTypesWithOptions(WorkspaceId *string, request *GetAvailableParserTypesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAvailableParserTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAvailableParserTypes"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/parser/parsertype"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAvailableParserTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all supported parser types based on the input file type (file extension).
//
// @param request - GetAvailableParserTypesRequest
//
// @return GetAvailableParserTypesResponse
func (client *Client) GetAvailableParserTypes(WorkspaceId *string, request *GetAvailableParserTypesRequest) (_result *GetAvailableParserTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAvailableParserTypesResponse{}
	_body, _err := client.GetAvailableParserTypesWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current status of a specified knowledge base creation or add document job.
//
// Description:
//
// 1.  A knowledge base job is running. You can call the [SubmitIndexJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexjob) operation to create a creation job or the [SubmitIndexAddDocumentsJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexadddocumentsjob) operation to create a add document job. Then, obtain the `JobId` returned by the operations.
//
// 2.  We recommend that you call this operation at intervals of more than 5 seconds.
//
// 3.  This interface is idempotent.
//
// @param request - GetIndexJobStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexJobStatusResponse
func (client *Client) GetIndexJobStatusWithOptions(WorkspaceId *string, request *GetIndexJobStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIndexJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIndexJobStatus"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/job/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIndexJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current status of a specified knowledge base creation or add document job.
//
// Description:
//
// 1.  A knowledge base job is running. You can call the [SubmitIndexJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexjob) operation to create a creation job or the [SubmitIndexAddDocumentsJob](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-submitindexadddocumentsjob) operation to create a add document job. Then, obtain the `JobId` returned by the operations.
//
// 2.  We recommend that you call this operation at intervals of more than 5 seconds.
//
// 3.  This interface is idempotent.
//
// @param request - GetIndexJobStatusRequest
//
// @return GetIndexJobStatusResponse
func (client *Client) GetIndexJobStatus(WorkspaceId *string, request *GetIndexJobStatusRequest) (_result *GetIndexJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexJobStatusResponse{}
	_body, _err := client.GetIndexJobStatusWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithOptions(workspaceId *string, memoryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemory"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @return GetMemoryResponse
func (client *Client) GetMemory(workspaceId *string, memoryId *string) (_result *GetMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryResponse{}
	_body, _err := client.GetMemoryWithOptions(workspaceId, memoryId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取记忆Node
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryNodeResponse
func (client *Client) GetMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryNodeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemoryNode"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId)) + "/memoryNodes/" + dara.PercentEncode(dara.StringValue(memoryNodeId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取记忆Node
//
// @return GetMemoryNodeResponse
func (client *Client) GetMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string) (_result *GetMemoryNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryNodeResponse{}
	_body, _err := client.GetMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data parsing settings in a specified category.
//
// @param request - GetParseSettingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParseSettingsResponse
func (client *Client) GetParseSettingsWithOptions(WorkspaceId *string, request *GetParseSettingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetParseSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParseSettings"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/parser/settings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParseSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data parsing settings in a specified category.
//
// @param request - GetParseSettingsRequest
//
// @return GetParseSettingsResponse
func (client *Client) GetParseSettings(WorkspaceId *string, request *GetParseSettingsRequest) (_result *GetParseSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetParseSettingsResponse{}
	_body, _err := client.GetParseSettingsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a prompt template based on the template ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPromptTemplateResponse
func (client *Client) GetPromptTemplateWithOptions(workspaceId *string, promptTemplateId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPromptTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPromptTemplate"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/promptTemplates/" + dara.PercentEncode(dara.StringValue(promptTemplateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPromptTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a prompt template based on the template ID.
//
// @return GetPromptTemplateResponse
func (client *Client) GetPromptTemplate(workspaceId *string, promptTemplateId *string) (_result *GetPromptTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPromptTemplateResponse{}
	_body, _err := client.GetPromptTemplateWithOptions(workspaceId, promptTemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取发布态智能体应用
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublishedAgentResponse
func (client *Client) GetPublishedAgentWithOptions(workspaceId *string, appCode *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPublishedAgentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPublishedAgent"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/application/agents/" + dara.PercentEncode(dara.StringValue(appCode))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPublishedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取发布态智能体应用
//
// @return GetPublishedAgentResponse
func (client *Client) GetPublishedAgent(workspaceId *string, appCode *string) (_result *GetPublishedAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPublishedAgentResponse{}
	_body, _err := client.GetPublishedAgentWithOptions(workspaceId, appCode, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 高代码部署服务
//
// @param request - HighCodeDeployRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HighCodeDeployResponse
func (client *Client) HighCodeDeployWithOptions(workspaceId *string, request *HighCodeDeployRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *HighCodeDeployResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentDesc) {
		body["agentDesc"] = request.AgentDesc
	}

	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentName) {
		body["agentName"] = request.AgentName
	}

	if !dara.IsNil(request.SourceCodeName) {
		body["sourceCodeName"] = request.SourceCodeName
	}

	if !dara.IsNil(request.SourceCodeOssUrl) {
		body["sourceCodeOssUrl"] = request.SourceCodeOssUrl
	}

	if !dara.IsNil(request.TelemetryEnabled) {
		body["telemetryEnabled"] = request.TelemetryEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HighCodeDeploy"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/highCode/publish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HighCodeDeployResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 高代码部署服务
//
// @param request - HighCodeDeployRequest
//
// @return HighCodeDeployResponse
func (client *Client) HighCodeDeploy(workspaceId *string, request *HighCodeDeployRequest) (_result *HighCodeDeployResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HighCodeDeployResponse{}
	_body, _err := client.HighCodeDeployWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListCategory
//
// @param request - ListCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoryResponse
func (client *Client) ListCategoryWithOptions(WorkspaceId *string, request *ListCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		body["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.CategoryType) {
		body["CategoryType"] = request.CategoryType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentCategoryId) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCategory"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/categories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListCategory
//
// @param request - ListCategoryRequest
//
// @return ListCategoryResponse
func (client *Client) ListCategory(WorkspaceId *string, request *ListCategoryRequest) (_result *ListCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCategoryResponse{}
	_body, _err := client.ListCategoryWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// For unstructured knowledge base, obtains the details of all chunks of a specified document; for structured knowledge base, obtains the details of all chunks.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- This interface is idempotent.
//
// @param request - ListChunksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChunksResponse
func (client *Client) ListChunksWithOptions(WorkspaceId *string, request *ListChunksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListChunksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.Filed) {
		body["Filed"] = request.Filed
	}

	if !dara.IsNil(request.IndexId) {
		body["IndexId"] = request.IndexId
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChunks"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/list_chunks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChunksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For unstructured knowledge base, obtains the details of all chunks of a specified document; for structured knowledge base, obtains the details of all chunks.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- This interface is idempotent.
//
// @param request - ListChunksRequest
//
// @return ListChunksResponse
func (client *Client) ListChunks(WorkspaceId *string, request *ListChunksRequest) (_result *ListChunksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChunksResponse{}
	_body, _err := client.ListChunksWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of one or more documents in a specified category.
//
// Description:
//
//	  If you are using a RAM user, you must first obtain the OpenAPI management permissions (namely sfm:ListFile) of Model Studio. For more information, see [Grant OpenAPI permissions to a RAM user](https://help.aliyun.com/document_detail/2848578.html). If you are using the Alibaba Cloud account, you do not need permissions. We recommend that you use [the latest version of the SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
//		- During a paged query, set `MaxResults` to specify the maximum number of entries to return. The return value of `NextToken` is a pagination token that can be used in the next call to retrieve a new page of results. When you query subsequent pages, set the `NextToken` parameter to the `NextToken` obtained in the last returned result. You can also set the `MaxResults` parameter to limit the number of entries to be returned. If no `NextToken` is returned, the result is completely returned and no more requests are required.
//
//		- This operation is idempotent.
//
// **Throttling:*	- Throttling will be triggered if you call this operation frequently. Do not exceed 5 times per second. If throttling is triggered, try again later.
//
// @param request - ListFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileResponse
func (client *Client) ListFileWithOptions(WorkspaceId *string, request *ListFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFile"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/files"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more documents in a specified category.
//
// Description:
//
//	  If you are using a RAM user, you must first obtain the OpenAPI management permissions (namely sfm:ListFile) of Model Studio. For more information, see [Grant OpenAPI permissions to a RAM user](https://help.aliyun.com/document_detail/2848578.html). If you are using the Alibaba Cloud account, you do not need permissions. We recommend that you use [the latest version of the SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
//		- During a paged query, set `MaxResults` to specify the maximum number of entries to return. The return value of `NextToken` is a pagination token that can be used in the next call to retrieve a new page of results. When you query subsequent pages, set the `NextToken` parameter to the `NextToken` obtained in the last returned result. You can also set the `MaxResults` parameter to limit the number of entries to be returned. If no `NextToken` is returned, the result is completely returned and no more requests are required.
//
//		- This operation is idempotent.
//
// **Throttling:*	- Throttling will be triggered if you call this operation frequently. Do not exceed 5 times per second. If throttling is triggered, try again later.
//
// @param request - ListFileRequest
//
// @return ListFileResponse
func (client *Client) ListFile(WorkspaceId *string, request *ListFileRequest) (_result *ListFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFileResponse{}
	_body, _err := client.ListFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of one or more documents in a specified knowledge base.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- This interface is idempotent.
//
// @param request - ListIndexDocumentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexDocumentsResponse
func (client *Client) ListIndexDocumentsWithOptions(WorkspaceId *string, request *ListIndexDocumentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndexDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentName) {
		query["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.DocumentStatus) {
		query["DocumentStatus"] = request.DocumentStatus
	}

	if !dara.IsNil(request.EnableNameLike) {
		query["EnableNameLike"] = request.EnableNameLike
	}

	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
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
		Action:      dara.String("ListIndexDocuments"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/list_index_documents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndexDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more documents in a specified knowledge base.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- This interface is idempotent.
//
// @param request - ListIndexDocumentsRequest
//
// @return ListIndexDocumentsResponse
func (client *Client) ListIndexDocuments(WorkspaceId *string, request *ListIndexDocumentsRequest) (_result *ListIndexDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexDocumentsResponse{}
	_body, _err := client.ListIndexDocumentsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Index文件详情
//
// @param request - ListIndexFileDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexFileDetailsResponse
func (client *Client) ListIndexFileDetailsWithOptions(WorkspaceId *string, request *ListIndexFileDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndexFileDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentName) {
		query["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.DocumentStatus) {
		query["DocumentStatus"] = request.DocumentStatus
	}

	if !dara.IsNil(request.EnableNameLike) {
		query["EnableNameLike"] = request.EnableNameLike
	}

	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
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
		Action:      dara.String("ListIndexFileDetails"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/list_index_file_detail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndexFileDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Index文件详情
//
// @param request - ListIndexFileDetailsRequest
//
// @return ListIndexFileDetailsResponse
func (client *Client) ListIndexFileDetails(WorkspaceId *string, request *ListIndexFileDetailsRequest) (_result *ListIndexFileDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexFileDetailsResponse{}
	_body, _err := client.ListIndexFileDetailsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists knowledge bases in a specified workspace.
//
// Description:
//
// This interface is idempotent.
//
// @param request - ListIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndicesResponse
func (client *Client) ListIndicesWithOptions(WorkspaceId *string, request *ListIndicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexName) {
		query["IndexName"] = request.IndexName
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
		Action:      dara.String("ListIndices"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/list_indices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists knowledge bases in a specified workspace.
//
// Description:
//
// This interface is idempotent.
//
// @param request - ListIndicesRequest
//
// @return ListIndicesResponse
func (client *Client) ListIndices(WorkspaceId *string, request *ListIndicesRequest) (_result *ListIndicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndicesResponse{}
	_body, _err := client.ListIndicesWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @param request - ListMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoriesResponse
func (client *Client) ListMemoriesWithOptions(workspaceId *string, request *ListMemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoriesResponse, _err error) {
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
		Action:      dara.String("ListMemories"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @param request - ListMemoriesRequest
//
// @return ListMemoriesResponse
func (client *Client) ListMemories(workspaceId *string, request *ListMemoriesRequest) (_result *ListMemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoriesResponse{}
	_body, _err := client.ListMemoriesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取记忆Node列表
//
// @param request - ListMemoryNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryNodesResponse
func (client *Client) ListMemoryNodesWithOptions(workspaceId *string, memoryId *string, request *ListMemoryNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryNodesResponse, _err error) {
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
		Action:      dara.String("ListMemoryNodes"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId)) + "/memoryNodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoryNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取记忆Node列表
//
// @param request - ListMemoryNodesRequest
//
// @return ListMemoryNodesResponse
func (client *Client) ListMemoryNodes(workspaceId *string, memoryId *string, request *ListMemoryNodesRequest) (_result *ListMemoryNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoryNodesResponse{}
	_body, _err := client.ListMemoryNodesWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of prompt templates.
//
// @param request - ListPromptTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPromptTemplatesResponse
func (client *Client) ListPromptTemplatesWithOptions(workspaceId *string, request *ListPromptTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPromptTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPromptTemplates"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/promptTemplates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPromptTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of prompt templates.
//
// @param request - ListPromptTemplatesRequest
//
// @return ListPromptTemplatesResponse
func (client *Client) ListPromptTemplates(workspaceId *string, request *ListPromptTemplatesRequest) (_result *ListPromptTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPromptTemplatesResponse{}
	_body, _err := client.ListPromptTemplatesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询已发布的智能体应用列表
//
// @param request - ListPublishedAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedAgentResponse
func (client *Client) ListPublishedAgentWithOptions(workspaceId *string, request *ListPublishedAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPublishedAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["pageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishedAgent"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/application/agents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询已发布的智能体应用列表
//
// @param request - ListPublishedAgentRequest
//
// @return ListPublishedAgentResponse
func (client *Client) ListPublishedAgent(workspaceId *string, request *ListPublishedAgentRequest) (_result *ListPublishedAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPublishedAgentResponse{}
	_body, _err := client.ListPublishedAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information from a specified knowledge base.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- The response time may be long because this operation involves complex retrieval and matching. We recommend that you set appropriate timeout and retry policy for requests.
//
//		- This interface is idempotent.
//
// @param tmpReq - RetrieveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveResponse
func (client *Client) RetrieveWithOptions(WorkspaceId *string, tmpReq *RetrieveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetrieveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RetrieveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Images) {
		request.ImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Images, dara.String("Images"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.QueryHistory) {
		request.QueryHistoryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryHistory, dara.String("QueryHistory"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rerank) {
		request.RerankShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rerank, dara.String("Rerank"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rewrite) {
		request.RewriteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rewrite, dara.String("Rewrite"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchFilters) {
		request.SearchFiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchFilters, dara.String("SearchFilters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DenseSimilarityTopK) {
		query["DenseSimilarityTopK"] = request.DenseSimilarityTopK
	}

	if !dara.IsNil(request.EnableReranking) {
		query["EnableReranking"] = request.EnableReranking
	}

	if !dara.IsNil(request.EnableRewrite) {
		query["EnableRewrite"] = request.EnableRewrite
	}

	if !dara.IsNil(request.ImagesShrink) {
		query["Images"] = request.ImagesShrink
	}

	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.QueryHistoryShrink) {
		query["QueryHistory"] = request.QueryHistoryShrink
	}

	if !dara.IsNil(request.RerankShrink) {
		query["Rerank"] = request.RerankShrink
	}

	if !dara.IsNil(request.RerankMinScore) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !dara.IsNil(request.RerankTopN) {
		query["RerankTopN"] = request.RerankTopN
	}

	if !dara.IsNil(request.RewriteShrink) {
		query["Rewrite"] = request.RewriteShrink
	}

	if !dara.IsNil(request.SaveRetrieverHistory) {
		query["SaveRetrieverHistory"] = request.SaveRetrieverHistory
	}

	if !dara.IsNil(request.SearchFiltersShrink) {
		query["SearchFilters"] = request.SearchFiltersShrink
	}

	if !dara.IsNil(request.SparseSimilarityTopK) {
		query["SparseSimilarityTopK"] = request.SparseSimilarityTopK
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Retrieve"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/retrieve"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information from a specified knowledge base.
//
// Description:
//
//	  Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- The response time may be long because this operation involves complex retrieval and matching. We recommend that you set appropriate timeout and retry policy for requests.
//
//		- This interface is idempotent.
//
// @param request - RetrieveRequest
//
// @return RetrieveResponse
func (client *Client) Retrieve(WorkspaceId *string, request *RetrieveRequest) (_result *RetrieveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetrieveResponse{}
	_body, _err := client.RetrieveWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds parsed documents to an unstructured knowledge base.
//
// Description:
//
//	  You must first upload documents to [Data Management](https://bailian.console.aliyun.com/#/data-center) and obtain the `FileId`. The documents are the knowledge source of the knowledge base. For more information, see [Import Data](https://www.alibabacloud.com/help/en/model-studio/user-guide/data-import-instructions).
//
//		- Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- After you call this operation, you can call the [GetIndexJobStatus](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-getindexjobstatus) operation to query the status of the job. More than 20 calls to the GetIndexJobStatus operation per minute may trigger throttling.
//
//		- Execution takes a period of time after this operation is called. Do not make new request before the request is returned. This interface is not idempotent.
//
// @param tmpReq - SubmitIndexAddDocumentsJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIndexAddDocumentsJobResponse
func (client *Client) SubmitIndexAddDocumentsJobWithOptions(WorkspaceId *string, tmpReq *SubmitIndexAddDocumentsJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitIndexAddDocumentsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitIndexAddDocumentsJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryIds) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, dara.String("CategoryIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocumentIds) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, dara.String("DocumentIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryIdsShrink) {
		query["CategoryIds"] = request.CategoryIdsShrink
	}

	if !dara.IsNil(request.ChunkMode) {
		query["ChunkMode"] = request.ChunkMode
	}

	if !dara.IsNil(request.ChunkSize) {
		query["ChunkSize"] = request.ChunkSize
	}

	if !dara.IsNil(request.DocumentIdsShrink) {
		query["DocumentIds"] = request.DocumentIdsShrink
	}

	if !dara.IsNil(request.EnableHeaders) {
		query["EnableHeaders"] = request.EnableHeaders
	}

	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
	}

	if !dara.IsNil(request.OverlapSize) {
		query["OverlapSize"] = request.OverlapSize
	}

	if !dara.IsNil(request.Separator) {
		query["Separator"] = request.Separator
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIndexAddDocumentsJob"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/add_documents_to_index"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIndexAddDocumentsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds parsed documents to an unstructured knowledge base.
//
// Description:
//
//	  You must first upload documents to [Data Management](https://bailian.console.aliyun.com/#/data-center) and obtain the `FileId`. The documents are the knowledge source of the knowledge base. For more information, see [Import Data](https://www.alibabacloud.com/help/en/model-studio/user-guide/data-import-instructions).
//
//		- Before you call this operation, make sure that your knowledge base is created and is not deleted. That is, the primary key ID of the knowledge base `IndexId` is valid.
//
//		- After you call this operation, you can call the [GetIndexJobStatus](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-getindexjobstatus) operation to query the status of the job. More than 20 calls to the GetIndexJobStatus operation per minute may trigger throttling.
//
//		- Execution takes a period of time after this operation is called. Do not make new request before the request is returned. This interface is not idempotent.
//
// @param request - SubmitIndexAddDocumentsJobRequest
//
// @return SubmitIndexAddDocumentsJobResponse
func (client *Client) SubmitIndexAddDocumentsJob(WorkspaceId *string, request *SubmitIndexAddDocumentsJobRequest) (_result *SubmitIndexAddDocumentsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitIndexAddDocumentsJobResponse{}
	_body, _err := client.SubmitIndexAddDocumentsJobWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a specified CreateIndex job to complete knowledge base creation.
//
// Description:
//
// 1.  Before you call this operation, you must call the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation and obtain the `IndexId`.
//
// 2.  Execution takes a period of time after this operation is called. Do not make new request before the request is returned.
//
// 3.  If you want to query the execution status of the job after you call this operation, call the [GetIndexJobStatus](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-getindexjobstatus) operation.
//
// 4.  This interface is not idempotent.
//
// @param request - SubmitIndexJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIndexJobResponse
func (client *Client) SubmitIndexJobWithOptions(WorkspaceId *string, request *SubmitIndexJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitIndexJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIndexJob"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/submit_index_job"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIndexJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a specified CreateIndex job to complete knowledge base creation.
//
// Description:
//
// 1.  Before you call this operation, you must call the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation and obtain the `IndexId`.
//
// 2.  Execution takes a period of time after this operation is called. Do not make new request before the request is returned.
//
// 3.  If you want to query the execution status of the job after you call this operation, call the [GetIndexJobStatus](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-getindexjobstatus) operation.
//
// 4.  This interface is not idempotent.
//
// @param request - SubmitIndexJobRequest
//
// @return SubmitIndexJobResponse
func (client *Client) SubmitIndexJob(WorkspaceId *string, request *SubmitIndexJobRequest) (_result *SubmitIndexJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitIndexJobResponse{}
	_body, _err := client.SubmitIndexJobWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新并发布智能体应用
//
// @param tmpReq - UpdateAndPublishAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAndPublishAgentResponse
func (client *Client) UpdateAndPublishAgentWithOptions(workspaceId *string, appCode *string, tmpReq *UpdateAndPublishAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAndPublishAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAndPublishAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationConfig) {
		request.ApplicationConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationConfig, dara.String("applicationConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SampleLibrary) {
		request.SampleLibraryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SampleLibrary, dara.String("sampleLibrary"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigShrink) {
		body["applicationConfig"] = request.ApplicationConfigShrink
	}

	if !dara.IsNil(request.Instructions) {
		body["instructions"] = request.Instructions
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SampleLibraryShrink) {
		body["sampleLibrary"] = request.SampleLibraryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAndPublishAgent"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/application/agents/" + dara.PercentEncode(dara.StringValue(appCode))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAndPublishAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新并发布智能体应用
//
// @param request - UpdateAndPublishAgentRequest
//
// @return UpdateAndPublishAgentResponse
func (client *Client) UpdateAndPublishAgent(workspaceId *string, appCode *string, request *UpdateAndPublishAgentRequest) (_result *UpdateAndPublishAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAndPublishAgentResponse{}
	_body, _err := client.UpdateAndPublishAgentWithOptions(workspaceId, appCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 选择更新并发布智能体应用
//
// @param tmpReq - UpdateAndPublishAgentSelectiveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAndPublishAgentSelectiveResponse
func (client *Client) UpdateAndPublishAgentSelectiveWithOptions(workspaceId *string, appCode *string, tmpReq *UpdateAndPublishAgentSelectiveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAndPublishAgentSelectiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAndPublishAgentSelectiveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationConfig) {
		request.ApplicationConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationConfig, dara.String("applicationConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SampleLibrary) {
		request.SampleLibraryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SampleLibrary, dara.String("sampleLibrary"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigShrink) {
		body["applicationConfig"] = request.ApplicationConfigShrink
	}

	if !dara.IsNil(request.Instructions) {
		body["instructions"] = request.Instructions
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SampleLibraryShrink) {
		body["sampleLibrary"] = request.SampleLibraryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAndPublishAgentSelective"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/application/agents/" + dara.PercentEncode(dara.StringValue(appCode)) + "/updateAndPublishAgentSelective"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAndPublishAgentSelectiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 选择更新并发布智能体应用
//
// @param request - UpdateAndPublishAgentSelectiveRequest
//
// @return UpdateAndPublishAgentSelectiveResponse
func (client *Client) UpdateAndPublishAgentSelective(workspaceId *string, appCode *string, request *UpdateAndPublishAgentSelectiveRequest) (_result *UpdateAndPublishAgentSelectiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAndPublishAgentSelectiveResponse{}
	_body, _err := client.UpdateAndPublishAgentSelectiveWithOptions(workspaceId, appCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the content and title of a specified text chunk in the knowledge base, and sets whether the chunk participates in knowledge base retrieval.
//
// Description:
//
//	  **Limits**: This operation supports only knowledge base of the document search type. Data query and image Q\\&A types are not supported.
//
//		- **Required permissions**:
//
//	    	- **RAM users**: Must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) of Model Studio (such as the `AliyunBailianDataFullAccess` policy, which includes the sfm:UpdateChunk permission required), and [become member of a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//	    	- **Alibaba Cloud account**: Has the permission by default, and can call the operation directly.
//
//		- **Call method**: We recommend using the latest version of the [GenAI Service Platform SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK encapsulates complex signature computational logic to simplify the call process.
//
//		- **Delay**: The update takes effect immediately. During peak hours, the update may take place in seconds.
//
//		- **Idempotence**: This operation is idempotent. If you perform a repeated operation on a chunk that has already been updated, the interface returns a success.
//
// **Rate limit:*	- Rate limiting will be triggered if you call this operation frequently. Do not exceed 10 times per second. If limiting is triggered, try again later.
//
// @param request - UpdateChunkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChunkResponse
func (client *Client) UpdateChunkWithOptions(WorkspaceId *string, request *UpdateChunkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateChunkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChunkId) {
		query["ChunkId"] = request.ChunkId
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.IsDisplayedChunkContent) {
		query["IsDisplayedChunkContent"] = request.IsDisplayedChunkContent
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChunk"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/chunk/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChunkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the content and title of a specified text chunk in the knowledge base, and sets whether the chunk participates in knowledge base retrieval.
//
// Description:
//
//	  **Limits**: This operation supports only knowledge base of the document search type. Data query and image Q\\&A types are not supported.
//
//		- **Required permissions**:
//
//	    	- **RAM users**: Must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) of Model Studio (such as the `AliyunBailianDataFullAccess` policy, which includes the sfm:UpdateChunk permission required), and [become member of a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//	    	- **Alibaba Cloud account**: Has the permission by default, and can call the operation directly.
//
//		- **Call method**: We recommend using the latest version of the [GenAI Service Platform SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK encapsulates complex signature computational logic to simplify the call process.
//
//		- **Delay**: The update takes effect immediately. During peak hours, the update may take place in seconds.
//
//		- **Idempotence**: This operation is idempotent. If you perform a repeated operation on a chunk that has already been updated, the interface returns a success.
//
// **Rate limit:*	- Rate limiting will be triggered if you call this operation frequently. Do not exceed 10 times per second. If limiting is triggered, try again later.
//
// @param request - UpdateChunkRequest
//
// @return UpdateChunkResponse
func (client *Client) UpdateChunk(WorkspaceId *string, request *UpdateChunkRequest) (_result *UpdateChunkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateChunkResponse{}
	_body, _err := client.UpdateChunkWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档Tag
//
// @param tmpReq - UpdateFileTagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileTagResponse
func (client *Client) UpdateFileTagWithOptions(WorkspaceId *string, FileId *string, tmpReq *UpdateFileTagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFileTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFileTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFileTag"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/file/" + dara.PercentEncode(dara.StringValue(FileId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档Tag
//
// @param request - UpdateFileTagRequest
//
// @return UpdateFileTagResponse
func (client *Client) UpdateFileTag(WorkspaceId *string, FileId *string, request *UpdateFileTagRequest) (_result *UpdateFileTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFileTagResponse{}
	_body, _err := client.UpdateFileTagWithOptions(WorkspaceId, FileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新memory
//
// @param request - UpdateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithOptions(workspaceId *string, memoryId *string, request *UpdateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemory"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新memory
//
// @param request - UpdateMemoryRequest
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemory(workspaceId *string, memoryId *string, request *UpdateMemoryRequest) (_result *UpdateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryResponse{}
	_body, _err := client.UpdateMemoryWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新记忆Node
//
// @param request - UpdateMemoryNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryNodeResponse
func (client *Client) UpdateMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, request *UpdateMemoryNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemoryNode"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/memories/" + dara.PercentEncode(dara.StringValue(memoryId)) + "/memoryNodes/" + dara.PercentEncode(dara.StringValue(memoryNodeId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新记忆Node
//
// @param request - UpdateMemoryNodeRequest
//
// @return UpdateMemoryNodeResponse
func (client *Client) UpdateMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string, request *UpdateMemoryNodeRequest) (_result *UpdateMemoryNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryNodeResponse{}
	_body, _err := client.UpdateMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a prompt template based on the template ID.
//
// @param request - UpdatePromptTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePromptTemplateResponse
func (client *Client) UpdatePromptTemplateWithOptions(workspaceId *string, promptTemplateId *string, request *UpdatePromptTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePromptTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePromptTemplate"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/promptTemplates/" + dara.PercentEncode(dara.StringValue(promptTemplateId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePromptTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a prompt template based on the template ID.
//
// @param request - UpdatePromptTemplateRequest
//
// @return UpdatePromptTemplateResponse
func (client *Client) UpdatePromptTemplate(workspaceId *string, promptTemplateId *string, request *UpdatePromptTemplateRequest) (_result *UpdatePromptTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePromptTemplateResponse{}
	_body, _err := client.UpdatePromptTemplateWithOptions(workspaceId, promptTemplateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
