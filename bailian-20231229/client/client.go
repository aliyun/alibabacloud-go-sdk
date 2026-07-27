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
	client.EndpointMap = map[string]*string{
		"cn-beijing":     dara.String("bailian.cn-beijing.aliyuncs.com"),
		"ap-southeast-1": dara.String("bailian.ap-southeast-1.aliyuncs.com"),
	}
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
// Creates a category in a specified workspace to classify and manage files. Each workspace supports a maximum of 500 categories.
//
// Description:
//
// - You cannot use an API to add data tables. To add data tables, go to the [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) page in the console.
//
// - A RAM user must obtain the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. The `AliyunBailianDataFullAccess` permission, which includes the `sfm:AddCategory` permission, is required. An Alibaba Cloud account can call this operation directly without requiring authorization. To call this operation, use the latest version of the [Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Rate limiting:*	- Frequent calls to this operation are subject to rate limiting. Do not exceed a frequency of 5 calls per second. If rate limiting is triggered, try again later.
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

	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
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
// Creates a category in a specified workspace to classify and manage files. Each workspace supports a maximum of 500 categories.
//
// Description:
//
// - You cannot use an API to add data tables. To add data tables, go to the [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) page in the console.
//
// - A RAM user must obtain the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. The `AliyunBailianDataFullAccess` permission, which includes the `sfm:AddCategory` permission, is required. An Alibaba Cloud account can call this operation directly without requiring authorization. To call this operation, use the latest version of the [Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Rate limiting:*	- Frequent calls to this operation are subject to rate limiting. Do not exceed a frequency of 5 calls per second. If rate limiting is triggered, try again later.
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
// Adds chunks to a document search (document), data query (table), or image Q&A (image) knowledge base.
//
// Description:
//
// - This operation adds chunk content to a specified knowledge base of the document search (document), data query (table), or image Q&A (image) type. Related operations on multimedia search (multimedia) knowledge bases are not supported. For data query and image Q&A knowledge bases, this operation takes effect only when the data source is a spreadsheet connector (excel).
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:ChunkList permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest [Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation has idempotence.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
//
// @param tmpReq - AddChunkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddChunkResponse
func (client *Client) AddChunkWithOptions(WorkspaceId *string, tmpReq *AddChunkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddChunkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddChunkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Field) {
		request.FieldShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Field, dara.String("field"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.DataId) {
		query["dataId"] = request.DataId
	}

	if !dara.IsNil(request.FieldShrink) {
		query["field"] = request.FieldShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddChunk"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/chunk/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddChunkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds chunks to a document search (document), data query (table), or image Q&A (image) knowledge base.
//
// Description:
//
// - This operation adds chunk content to a specified knowledge base of the document search (document), data query (table), or image Q&A (image) type. Related operations on multimedia search (multimedia) knowledge bases are not supported. For data query and image Q&A knowledge bases, this operation takes effect only when the data source is a spreadsheet connector (excel).
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:ChunkList permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest [Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation has idempotence.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
//
// @param request - AddChunkRequest
//
// @return AddChunkResponse
func (client *Client) AddChunk(WorkspaceId *string, request *AddChunkRequest) (_result *AddChunkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddChunkResponse{}
	_body, _err := client.AddChunkWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a connector. This API currently supports only file connectors.
//
// Description:
//
// - To call this operation, a RAM user (sub-account) must first have the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio and be a member of a [business space](https://help.aliyun.com/document_detail/2851098.html). This requires the `AliyunBailianDataFullAccess` policy, which includes the sfm:AddCategory permission. A primary account can call this operation directly without authorization. We recommend using the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is not idempotent.
//
// **Throttling:**
//
// Do not call this operation more than 5 times per second. If a request is throttled, try again later.
//
// @param tmpReq - AddConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddConnectorResponse
func (client *Client) AddConnectorWithOptions(WorkspaceId *string, tmpReq *AddConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileConnectorConfig) {
		request.FileConnectorConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileConnectorConfig, dara.String("FileConnectorConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorName) {
		body["ConnectorName"] = request.ConnectorName
	}

	if !dara.IsNil(request.ConnectorType) {
		body["ConnectorType"] = request.ConnectorType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileConnectorConfigShrink) {
		body["FileConnectorConfig"] = request.FileConnectorConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddConnector"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/connector"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a connector. This API currently supports only file connectors.
//
// Description:
//
// - To call this operation, a RAM user (sub-account) must first have the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio and be a member of a [business space](https://help.aliyun.com/document_detail/2851098.html). This requires the `AliyunBailianDataFullAccess` policy, which includes the sfm:AddCategory permission. A primary account can call this operation directly without authorization. We recommend using the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is not idempotent.
//
// **Throttling:**
//
// Do not call this operation more than 5 times per second. If a request is throttled, try again later.
//
// @param request - AddConnectorRequest
//
// @return AddConnectorResponse
func (client *Client) AddConnector(WorkspaceId *string, request *AddConnectorRequest) (_result *AddConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddConnectorResponse{}
	_body, _err := client.AddConnectorWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports a file from the temporary storage of Alibaba Cloud Model Studio into an Alibaba Cloud Model Studio data connection (formerly known as application data).
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:AddFile permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Invoke this operation by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Throttling:**
//
// This operation is subject to throttling. Do not exceed 10 calls per second. If you are throttled, retry later.
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
	if !dara.IsNil(tmpReq.ParserConfig) {
		request.ParserConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParserConfig, dara.String("ParserConfig"), dara.String("json"))
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

	if !dara.IsNil(request.LeaseId) {
		body["LeaseId"] = request.LeaseId
	}

	if !dara.IsNil(request.OriginalFileUrl) {
		body["OriginalFileUrl"] = request.OriginalFileUrl
	}

	if !dara.IsNil(request.Parser) {
		body["Parser"] = request.Parser
	}

	if !dara.IsNil(request.ParserConfigShrink) {
		body["ParserConfig"] = request.ParserConfigShrink
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
// Imports a file from the temporary storage of Alibaba Cloud Model Studio into an Alibaba Cloud Model Studio data connection (formerly known as application data).
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:AddFile permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Invoke this operation by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Throttling:**
//
// This operation is subject to throttling. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Imports files from an authorized OSS bucket into Alibaba Cloud Model Studio application data.
//
// Description:
//
// - Make sure the OSS bucket belongs to the same Alibaba Cloud account as Model Studio and that authorization has been completed as described in [Importing data from OSS configuration instructions](https://help.aliyun.com/document_detail/2782155.html).
//
//   - Supported bucket storage types do not include Archive, Cold Archive, or Deep Cold Archive. Buckets with content encryption enabled are supported. Buckets with public read/write, public read, or private access are supported.
//
//   - To use a bucket with [Referer-based hotlink protection](https://help.aliyun.com/document_detail/2636937.html) enabled, add the domain name `*.console.aliyun.com` to the Referer whitelist as described in [Allow access only from trusted websites](https://help.aliyun.com/document_detail/2636937.html).
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (the `AliyunBailianDataFullAccess` permission is required, which includes the sfm:AddFilesFromAuthorizedOss permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Call this operation by using the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation does not have idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 5 calls per second. If throttled, retry later.
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
// Imports files from an authorized OSS bucket into Alibaba Cloud Model Studio application data.
//
// Description:
//
// - Make sure the OSS bucket belongs to the same Alibaba Cloud account as Model Studio and that authorization has been completed as described in [Importing data from OSS configuration instructions](https://help.aliyun.com/document_detail/2782155.html).
//
//   - Supported bucket storage types do not include Archive, Cold Archive, or Deep Cold Archive. Buckets with content encryption enabled are supported. Buckets with public read/write, public read, or private access are supported.
//
//   - To use a bucket with [Referer-based hotlink protection](https://help.aliyun.com/document_detail/2636937.html) enabled, add the domain name `*.console.aliyun.com` to the Referer whitelist as described in [Allow access only from trusted websites](https://help.aliyun.com/document_detail/2636937.html).
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (the `AliyunBailianDataFullAccess` permission is required, which includes the sfm:AddFilesFromAuthorizedOss permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Call this operation by using the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation does not have idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 5 calls per second. If throttled, retry later.
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
// Add a table for the table data connector.
//
// Description:
//
// - RAM users (sub-accounts) must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Chinese China Bailian (requires `AliyunBailianDataFullAccess`, which includes the sfm:AddTable permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this API. Alibaba Cloud accounts (primary accounts) can call this API directly without authorization. We recommend that you call this API by using the latest <props="china">[Chinese China Bailian SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Chinese China Bailian SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This API is not idempotent.
//
// **Throttling:**
//
// This API is subject to throttling when called frequently. Do not exceed 10 calls per second. If throttling is triggered, retry later.
//
// @param tmpReq - AddTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTableResponse
func (client *Client) AddTableWithOptions(WorkspaceId *string, tmpReq *AddTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableColumns) {
		request.TableColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableColumns, dara.String("TableColumns"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.TableColumnsShrink) {
		body["TableColumns"] = request.TableColumnsShrink
	}

	if !dara.IsNil(request.TableDesc) {
		body["TableDesc"] = request.TableDesc
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTable"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/table"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a table for the table data connector.
//
// Description:
//
// - RAM users (sub-accounts) must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Chinese China Bailian (requires `AliyunBailianDataFullAccess`, which includes the sfm:AddTable permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this API. Alibaba Cloud accounts (primary accounts) can call this API directly without authorization. We recommend that you call this API by using the latest <props="china">[Chinese China Bailian SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Chinese China Bailian SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This API is not idempotent.
//
// **Throttling:**
//
// This API is subject to throttling when called frequently. Do not exceed 10 calls per second. If throttling is triggered, retry later.
//
// @param request - AddTableRequest
//
// @return AddTableResponse
func (client *Client) AddTable(WorkspaceId *string, request *AddTableRequest) (_result *AddTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddTableResponse{}
	_body, _err := client.AddTableWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Request an upload lease for uploading knowledge base files or files for agent application conversational interactions.
//
// Description:
//
// - RAM users (sub-accounts) must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ApplyFileUploadLease permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this API. Alibaba Cloud accounts (primary accounts) can directly call this API without authorization. We recommend that you call this API by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This API is not idempotent.
//
// **Throttling:**
//
// This API is subject to throttling if called too frequently. The frequency must not exceed 10 calls per second. If throttled, please retry later.
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
// Request an upload lease for uploading knowledge base files or files for agent application conversational interactions.
//
// Description:
//
// - RAM users (sub-accounts) must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ApplyFileUploadLease permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this API. Alibaba Cloud accounts (primary accounts) can directly call this API without authorization. We recommend that you call this API by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This API is not idempotent.
//
// **Throttling:**
//
// This API is subject to throttling if called too frequently. The frequency must not exceed 10 calls per second. If throttled, please retry later.
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
// This interface is intended for pro-code deployment only; other scenarios are currently not supported. It is used to apply for a temporary file upload lease. After obtaining the lease, you must upload the file manually.
//
// Description:
//
// 1\\. This interface is intended for pro-code deployment only; other scenarios are currently not supported. 2. After obtaining the temporary file upload lease via this interface, upload the file manually.
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
// This interface is intended for pro-code deployment only; other scenarios are currently not supported. It is used to apply for a temporary file upload lease. After obtaining the lease, you must upload the file manually.
//
// Description:
//
// 1\\. This interface is intended for pro-code deployment only; other scenarios are currently not supported. 2. After obtaining the temporary file upload lease via this interface, upload the file manually.
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
// This operation updates document tags in a data connection in batches.
//
// @param tmpReq - BatchUpdateFileTagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateFileTagResponse
func (client *Client) BatchUpdateFileTagWithOptions(WorkspaceId *string, tmpReq *BatchUpdateFileTagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchUpdateFileTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateFileTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileInfos) {
		request.FileInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileInfos, dara.String("FileInfos"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileInfosShrink) {
		body["FileInfos"] = request.FileInfosShrink
	}

	if !dara.IsNil(request.UpdateMode) {
		body["UpdateMode"] = request.UpdateMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateFileTag"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/batchupdatetag"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateFileTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation updates document tags in a data connection in batches.
//
// @param request - BatchUpdateFileTagRequest
//
// @return BatchUpdateFileTagResponse
func (client *Client) BatchUpdateFileTag(WorkspaceId *string, request *BatchUpdateFileTagRequest) (_result *BatchUpdateFileTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUpdateFileTagResponse{}
	_body, _err := client.BatchUpdateFileTagWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the parsing method for specific file types. For example, specifies large model document parsing for .pdf files or Qwen VL parsing for .jpg files.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ChangeParseSetting permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Invoke this operation by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// Configures the parsing method for specific file types. For example, specifies large model document parsing for .pdf files or Qwen VL parsing for .jpg files.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ChangeParseSetting permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Invoke this operation by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// Creates a knowledge base, either an unstructured knowledge base based on documents or audio/video, or a structured knowledge base for data queries or image-based Q&A.
//
// Description:
//
// - **Permission requirements**:
//
//   - **Resource Access Management (RAM) user**: Obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio first (you can use the `AliyunBailianDataFullAccess` policy, which includes the sfm:CreateIndex permission required by this operation), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation.
//
//   - **Alibaba Cloud account**: Has permissions by default and can invoke this operation directly.
//
// - **Calling method**: Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK has encapsulated the complex signature calculation logic and simplifies the invocation procedure.
//
// - **What to do next**: This operation only performs initialization of the knowledge base creation job. After invoking this operation, you must invoke the **SubmitIndexJob*	- operation to complete the creation (otherwise, you will get an empty knowledge base). For code examples, refer to [Knowledge Base API Guide](https://help.aliyun.com/document_detail/2852772.html).
//
// - **Idempotence**: This operation does not have idempotence. Repeated invocations may create multiple knowledge bases with the same name. Implement idempotent invocations by querying first and then creating.
//
// **Rate limit:**
//
// This operation is subject to rate limiting. Do not exceed 10 calls per second. If you are throttled, retry later.
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

	if !dara.IsNil(request.RerankInstruct) {
		query["RerankInstruct"] = request.RerankInstruct
	}

	if !dara.IsNil(request.RerankMinScore) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !dara.IsNil(request.RerankMode) {
		query["RerankMode"] = request.RerankMode
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

	if !dara.IsNil(request.ChannelType) {
		query["channelType"] = request.ChannelType
	}

	if !dara.IsNil(request.ChunkMode) {
		query["chunkMode"] = request.ChunkMode
	}

	if !dara.IsNil(request.ConnectId) {
		query["connectId"] = request.ConnectId
	}

	if !dara.IsNil(request.Database) {
		query["database"] = request.Database
	}

	if !dara.IsNil(request.DatasourceCode) {
		query["datasourceCode"] = request.DatasourceCode
	}

	if !dara.IsNil(request.EnableHeaders) {
		query["enableHeaders"] = request.EnableHeaders
	}

	if !dara.IsNil(request.KnowledgeScene) {
		query["knowledgeScene"] = request.KnowledgeScene
	}

	if !dara.IsNil(request.KnowledgeType) {
		query["knowledgeType"] = request.KnowledgeType
	}

	if !dara.IsNil(request.MetaExtractColumnsShrink) {
		query["metaExtractColumns"] = request.MetaExtractColumnsShrink
	}

	if !dara.IsNil(request.PipelineCommercialCu) {
		query["pipelineCommercialCu"] = request.PipelineCommercialCu
	}

	if !dara.IsNil(request.PipelineCommercialType) {
		query["pipelineCommercialType"] = request.PipelineCommercialType
	}

	if !dara.IsNil(request.PipelineRetrieveRateLimitStrategy) {
		query["pipelineRetrieveRateLimitStrategy"] = request.PipelineRetrieveRateLimitStrategy
	}

	if !dara.IsNil(request.Table) {
		query["table"] = request.Table
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
// Creates a knowledge base, either an unstructured knowledge base based on documents or audio/video, or a structured knowledge base for data queries or image-based Q&A.
//
// Description:
//
// - **Permission requirements**:
//
//   - **Resource Access Management (RAM) user**: Obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio first (you can use the `AliyunBailianDataFullAccess` policy, which includes the sfm:CreateIndex permission required by this operation), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation.
//
//   - **Alibaba Cloud account**: Has permissions by default and can invoke this operation directly.
//
// - **Calling method**: Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK has encapsulated the complex signature calculation logic and simplifies the invocation procedure.
//
// - **What to do next**: This operation only performs initialization of the knowledge base creation job. After invoking this operation, you must invoke the **SubmitIndexJob*	- operation to complete the creation (otherwise, you will get an empty knowledge base). For code examples, refer to [Knowledge Base API Guide](https://help.aliyun.com/document_detail/2852772.html).
//
// - **Idempotence**: This operation does not have idempotence. Repeated invocations may create multiple knowledge bases with the same name. Implement idempotent invocations by querying first and then creating.
//
// **Rate limit:**
//
// This operation is subject to rate limiting. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Creates a long-term memory.
//
// Description:
//
// - You can store specific information from conversations (memory nodes. For more information, see [Long-term memory](https://www.alibabacloud.com/help/en/model-studio/user-guide/long-term-memory)) in a long-term memory. Agent applications can then reference this information in subsequent conversations. This is not an automatic creation procedure. You must first invoke the [CreateMemory](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-creatememory) operation to create a long-term memory and obtain the `memoryId`. Then pass the `memoryId` when you [invoke the agent application through the API](https://www.alibabacloud.com/help/en/model-studio/user-guide/application-calling).
//
//	> Long-term memory does not support storing and managing user profiles through the API. Perform related operations in the console. For more information, see [Long-term memory](https://www.alibabacloud.com/help/en/model-studio/user-guide/long-term-memory#578ebae524m6l).
//
// - If you pass a `memoryId`, the system uses automatic creation to generate memory nodes (MemoryNode) under the specified long-term memory based on conversation records. You can also invoke the [CreateMemoryNode](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-creatememorynode) operation to manually create memory nodes.
//
// - This operation does not support idempotence.
//
// **Throttling:*	- Ensure that the interval between two consecutive requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
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
// Creates a long-term memory.
//
// Description:
//
// - You can store specific information from conversations (memory nodes. For more information, see [Long-term memory](https://www.alibabacloud.com/help/en/model-studio/user-guide/long-term-memory)) in a long-term memory. Agent applications can then reference this information in subsequent conversations. This is not an automatic creation procedure. You must first invoke the [CreateMemory](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-creatememory) operation to create a long-term memory and obtain the `memoryId`. Then pass the `memoryId` when you [invoke the agent application through the API](https://www.alibabacloud.com/help/en/model-studio/user-guide/application-calling).
//
//	> Long-term memory does not support storing and managing user profiles through the API. Perform related operations in the console. For more information, see [Long-term memory](https://www.alibabacloud.com/help/en/model-studio/user-guide/long-term-memory#578ebae524m6l).
//
// - If you pass a `memoryId`, the system uses automatic creation to generate memory nodes (MemoryNode) under the specified long-term memory based on conversation records. You can also invoke the [CreateMemoryNode](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-creatememorynode) operation to manually create memory nodes.
//
// - This operation does not support idempotence.
//
// **Throttling:*	- Ensure that the interval between two consecutive requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
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
// Creates a memory node.
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
// Creates a memory node.
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
// Create a prompt template.
//
// Description:
//
// This API does not currently support the creation of text-to-image prompt templates.
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
// Create a prompt template.
//
// Description:
//
// This API does not currently support the creation of text-to-image prompt templates.
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
// @param request - DeleteAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithOptions(workspaceId *string, appCode *string, request *DeleteAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
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
// @param request - DeleteAgentRequest
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgent(workspaceId *string, appCode *string, request *DeleteAgentRequest) (_result *DeleteAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(workspaceId, appCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Permanently deletes a specified category.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:DeleteCategory permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is idempotent.
//
// **Rate limiting:**
//
// This operation is subject to rate limiting. Do not exceed 5 calls per second. If you are throttled, retry later.
//
// @param request - DeleteCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithOptions(CategoryId *string, WorkspaceId *string, request *DeleteCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
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
// Permanently deletes a specified category.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:DeleteCategory permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is idempotent.
//
// **Rate limiting:**
//
// This operation is subject to rate limiting. Do not exceed 5 calls per second. If you are throttled, retry later.
//
// @param request - DeleteCategoryRequest
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategory(CategoryId *string, WorkspaceId *string, request *DeleteCategoryRequest) (_result *DeleteCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(CategoryId, WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes specified text chunks from a knowledge base. Deleted text chunks cannot be retrieved or recalled.
//
// Description:
//
// <warning>  Deleted text chunks cannot be recovered (hard delete). Proceed with caution.
//
// - **Permission requirements**:
//
//   - **Resource Access Management (RAM) user**: Obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio first (you can use the `AliyunBailianDataFullAccess` policy, which includes the sfm:DeleteChunk permission required by this operation), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation.
//
//   - **Alibaba Cloud account**: Has permissions by default and can invoke this operation directly.
//
// - **Invocation method**: Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK provides encapsulation of complex signature calculation logic and simplifies the invocation procedure.
//
// - **Effective latency**: Changes typically take effect immediately. During peak hours, there may be a slight delay (seconds).
//
// - **Idempotence**: This operation is idempotent. Repeated calls to delete an already deleted text chunk return a success response.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// Deletes specified text chunks from a knowledge base. Deleted text chunks cannot be retrieved or recalled.
//
// Description:
//
// <warning>  Deleted text chunks cannot be recovered (hard delete). Proceed with caution.
//
// - **Permission requirements**:
//
//   - **Resource Access Management (RAM) user**: Obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio first (you can use the `AliyunBailianDataFullAccess` policy, which includes the sfm:DeleteChunk permission required by this operation), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation.
//
//   - **Alibaba Cloud account**: Has permissions by default and can invoke this operation directly.
//
// - **Invocation method**: Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK provides encapsulation of complex signature calculation logic and simplifies the invocation procedure.
//
// - **Effective latency**: Changes typically take effect immediately. During peak hours, there may be a slight delay (seconds).
//
// - **Idempotence**: This operation is idempotent. Repeated calls to delete an already deleted text chunk return a success response.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// 删除连接器
//
// Description:
//
// - RAM用户（子账号）需要首先获取阿里云百炼的[API权限](https://help.aliyun.com/document_detail/2848578.html)（需要`AliyunBailianDataFullAccess`，已包括sfm:DeleteConnector权限点），并[加入一个业务空间](https://help.aliyun.com/document_detail/2851098.html)后，方可调用本接口。阿里云账号（主账号）可直接调用无须授权。建议您通过最新版[阿里云百炼SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)[阿里云百炼SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29)来调用本接口。
//
// - 本接口不具备幂等性。
//
// **限流说明：**
//
// 本接口频繁调用会被限流，频率请勿超过5次/秒。如遇限流，请稍后重试。
//
// @param request - DeleteConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectorResponse
func (client *Client) DeleteConnectorWithOptions(ConnectorId *string, WorkspaceId *string, request *DeleteConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConnectorResponse, _err error) {
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
		Action:      dara.String("DeleteConnector"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/connector/" + dara.PercentEncode(dara.StringValue(ConnectorId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除连接器
//
// Description:
//
// - RAM用户（子账号）需要首先获取阿里云百炼的[API权限](https://help.aliyun.com/document_detail/2848578.html)（需要`AliyunBailianDataFullAccess`，已包括sfm:DeleteConnector权限点），并[加入一个业务空间](https://help.aliyun.com/document_detail/2851098.html)后，方可调用本接口。阿里云账号（主账号）可直接调用无须授权。建议您通过最新版[阿里云百炼SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)[阿里云百炼SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29)来调用本接口。
//
// - 本接口不具备幂等性。
//
// **限流说明：**
//
// 本接口频繁调用会被限流，频率请勿超过5次/秒。如遇限流，请稍后重试。
//
// @param request - DeleteConnectorRequest
//
// @return DeleteConnectorResponse
func (client *Client) DeleteConnector(ConnectorId *string, WorkspaceId *string, request *DeleteConnectorRequest) (_result *DeleteConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConnectorResponse{}
	_body, _err := client.DeleteConnectorWithOptions(ConnectorId, WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Permanently delete a specified file from application data. Deleting data tables via API is not supported. For details, see the API Guide below.
//
// Description:
//
// - Deleting data tables via API is not supported. To delete a data table or specific data within a table, go to <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) to perform the operation.
//
// - This API is used to delete files in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) and does not affect any knowledge bases that have already been built. To delete a file from a knowledge base, invoke the **DeleteIndexDocument*	- API.
//
// - A RAM user must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the `sfm:DeleteFile` permission point) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this API. An Alibaba Cloud account can invoke this API directly without authorization. We recommend that you use the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this API.
//
// - This API can only delete files whose status is either Failed to Parse (`PARSE_FAILED`) or Parse Succeeded (`PARSE_SUCCESS`).
//
// - This API is idempotent.
//
// **Rate Limiting Notice:**
//
// Frequent invocation of this API will trigger rate limiting. Do not exceed 10 requests per second. If rate limited, retry after a short wait.
//
// @param request - DeleteFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(FileId *string, WorkspaceId *string, request *DeleteFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
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
// Permanently delete a specified file from application data. Deleting data tables via API is not supported. For details, see the API Guide below.
//
// Description:
//
// - Deleting data tables via API is not supported. To delete a data table or specific data within a table, go to <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) to perform the operation.
//
// - This API is used to delete files in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) and does not affect any knowledge bases that have already been built. To delete a file from a knowledge base, invoke the **DeleteIndexDocument*	- API.
//
// - A RAM user must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the `sfm:DeleteFile` permission point) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this API. An Alibaba Cloud account can invoke this API directly without authorization. We recommend that you use the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this API.
//
// - This API can only delete files whose status is either Failed to Parse (`PARSE_FAILED`) or Parse Succeeded (`PARSE_SUCCESS`).
//
// - This API is idempotent.
//
// **Rate Limiting Notice:**
//
// Frequent invocation of this API will trigger rate limiting. Do not exceed 10 requests per second. If rate limited, retry after a short wait.
//
// @param request - DeleteFileRequest
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(FileId *string, WorkspaceId *string, request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(FileId, WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete files in batch
//
// Description:
//
// - Deleting data tables through the API is not supported. To delete a data table or specific data in a table, go to <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
//
// - This API is used to delete files in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). It does not affect knowledge bases that have already been built. To delete files in a knowledge base, call the **DeleteIndexDocument*	- operation.
//
// - A RAM user (sub-account) must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which already includes the sfm:DeleteFiles permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. The Alibaba Cloud account (main account) can call this operation directly without authorization. We recommend that you use the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation can only delete files whose status is parsing failed (PARSE_FAILED) or parsing succeeded (PARSE_SUCCESS).
//
// - This operation is idempotent.
//
// **Throttling:**
//
// Frequent calls to this operation are throttled. Do not exceed 10 queries per second (QPS). If you are throttled, try again later.
//
// @param tmpReq - DeleteFilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFilesResponse
func (client *Client) DeleteFilesWithOptions(WorkspaceId *string, tmpReq *DeleteFilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileIdsShrink) {
		body["FileIds"] = request.FileIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFiles"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/file/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete files in batch
//
// Description:
//
// - Deleting data tables through the API is not supported. To delete a data table or specific data in a table, go to <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
//
// - This API is used to delete files in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). It does not affect knowledge bases that have already been built. To delete files in a knowledge base, call the **DeleteIndexDocument*	- operation.
//
// - A RAM user (sub-account) must first obtain the [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which already includes the sfm:DeleteFiles permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. The Alibaba Cloud account (main account) can call this operation directly without authorization. We recommend that you use the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation can only delete files whose status is parsing failed (PARSE_FAILED) or parsing succeeded (PARSE_SUCCESS).
//
// - This operation is idempotent.
//
// **Throttling:**
//
// Frequent calls to this operation are throttled. Do not exceed 10 queries per second (QPS). If you are throttled, try again later.
//
// @param request - DeleteFilesRequest
//
// @return DeleteFilesResponse
func (client *Client) DeleteFiles(WorkspaceId *string, request *DeleteFilesRequest) (_result *DeleteFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFilesResponse{}
	_body, _err := client.DeleteFilesWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Permanently deletes a specified knowledge base.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:DeleteIndex permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - If the knowledge base is associated with an application, you must first dissociate it from the application before deleting it. This can currently only be done through the console. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
//
// - Deletion is irreversible. A deleted knowledge base cannot be recovered. Proceed with caution.
//
// - Invoking this operation does not delete files that have been imported into <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
//
// - This operation has idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Permanently deletes a specified knowledge base.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:DeleteIndex permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - If the knowledge base is associated with an application, you must first dissociate it from the application before deleting it. This can currently only be done through the console. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
//
// - Deletion is irreversible. A deleted knowledge base cannot be recovered. Proceed with caution.
//
// - Invoking this operation does not delete files that have been imported into <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
//
// - This operation has idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Permanently deletes files from a specified knowledge base.
//
// Description:
//
// - This operation does not support deleting data from data query or image Q&A knowledge bases. Use the Model Studio console instead.
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:DeleteIndexDocument permission), before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - Before calling this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - You can only delete files whose status is import failed (INSERT_ERROR) or import succeeded (FINISH) in the knowledge base. To query the file status in a specified knowledge base, call the **ListIndexDocuments*	- operation.
//
// - Deletion is irreversible. The content of deleted files cannot be recovered, and the **Retrieve*	- operation can no longer retrieve related information. Proceed with caution.
//
// - Calling this operation does not delete documents that have been imported into <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
//
// - This operation is idempotent.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Permanently deletes files from a specified knowledge base.
//
// Description:
//
// - This operation does not support deleting data from data query or image Q&A knowledge bases. Use the Model Studio console instead.
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requiring `AliyunBailianDataFullAccess`, which includes the sfm:DeleteIndexDocument permission), before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - Before calling this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - You can only delete files whose status is import failed (INSERT_ERROR) or import succeeded (FINISH) in the knowledge base. To query the file status in a specified knowledge base, call the **ListIndexDocuments*	- operation.
//
// - Deletion is irreversible. The content of deleted files cannot be recovered, and the **Retrieve*	- operation can no longer retrieve related information. Proceed with caution.
//
// - Calling this operation does not delete documents that have been imported into <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
//
// - This operation is idempotent.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Permanently deletes a specified long-term memory.
//
// Description:
//
// - Before calling this operation, make sure that your long-term memory has been created and has not been deleted (that is, the memoryId is valid).
//
// - The delete operation is irreversible. The deleted long-term memory, including all of its long-term memory nodes, cannot be recovered. The [GetMemory](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-getmemory) operation will no longer be able to retrieve its information. Proceed with caution.
//
// - This operation is idempotent.
//
// **Rate limit:*	- Make sure that the interval between two requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
//
// @param request - DeleteMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(workspaceId *string, memoryId *string, request *DeleteMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
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
// Permanently deletes a specified long-term memory.
//
// Description:
//
// - Before calling this operation, make sure that your long-term memory has been created and has not been deleted (that is, the memoryId is valid).
//
// - The delete operation is irreversible. The deleted long-term memory, including all of its long-term memory nodes, cannot be recovered. The [GetMemory](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-getmemory) operation will no longer be able to retrieve its information. Proceed with caution.
//
// - This operation is idempotent.
//
// **Rate limit:*	- Make sure that the interval between two requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
//
// @param request - DeleteMemoryRequest
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(workspaceId *string, memoryId *string, request *DeleteMemoryRequest) (_result *DeleteMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a memory fragment.
//
// @param request - DeleteMemoryNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryNodeResponse
func (client *Client) DeleteMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, request *DeleteMemoryNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryNodeResponse, _err error) {
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
// Deletes a memory fragment.
//
// @param request - DeleteMemoryNodeRequest
//
// @return DeleteMemoryNodeResponse
func (client *Client) DeleteMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string, request *DeleteMemoryNodeRequest) (_result *DeleteMemoryNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryNodeResponse{}
	_body, _err := client.DeleteMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, request, headers, runtime)
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
// @param request - DeletePromptTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePromptTemplateResponse
func (client *Client) DeletePromptTemplateWithOptions(workspaceId *string, promptTemplateId *string, request *DeletePromptTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePromptTemplateResponse, _err error) {
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
// @param request - DeletePromptTemplateRequest
//
// @return DeletePromptTemplateResponse
func (client *Client) DeletePromptTemplate(workspaceId *string, promptTemplateId *string, request *DeletePromptTemplateRequest) (_result *DeletePromptTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePromptTemplateResponse{}
	_body, _err := client.DeletePromptTemplateWithOptions(workspaceId, promptTemplateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information of a file in application data, including the file name, type, and status.
//
// Description:
//
// - A Resource Access Management (RAM) user must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (`AliyunBailianDataFullAccess` or `AliyunBailianDataReadOnlyAccess`, both of which include the sfm:DescribeFile permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. An Alibaba Cloud account can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - This operation is idempotent.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
//
// @param request - DescribeFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileResponse
func (client *Client) DescribeFileWithOptions(WorkspaceId *string, FileId *string, request *DescribeFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeFileResponse, _err error) {
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
// Queries the basic information of a file in application data, including the file name, type, and status.
//
// Description:
//
// - A Resource Access Management (RAM) user must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (`AliyunBailianDataFullAccess` or `AliyunBailianDataReadOnlyAccess`, both of which include the sfm:DescribeFile permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. An Alibaba Cloud account can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - This operation is idempotent.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
//
// @param request - DescribeFileRequest
//
// @return DescribeFileResponse
func (client *Client) DescribeFile(WorkspaceId *string, FileId *string, request *DescribeFileRequest) (_result *DescribeFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFileResponse{}
	_body, _err := client.DescribeFileWithOptions(WorkspaceId, FileId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tipping status of the Alipay wallet bound to an application.
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
// Queries the tipping status of the Alipay wallet bound to an application.
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
// Retrieves the Alipay tipping URL for an application.
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
// Retrieves the Alipay tipping URL for an application.
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
// Description:
//
// - RAM用户（子账号）需要首先获取阿里云百炼的[API权限](https://help.aliyun.com/document_detail/2848578.html)（`AliyunBailianDataFullAccess`或`AliyunBailianDataReadOnlyAccess`均可，已包括sfm:GetAvailableParserTypes权限点），并[加入一个业务空间](https://help.aliyun.com/document_detail/2851098.html)后，方可调用本接口。阿里云账号（主账号）可直接调用无须授权。建议您通过最新版<props="china">[阿里云百炼SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[阿里云百炼SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29)来调用本接口。
//
// - 本接口具有幂等性。
//
// **限流说明：**
//
// 本接口频繁调用会被限流，频率请勿超过10次/秒。如遇限流，请稍后重试。
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
// Description:
//
// - RAM用户（子账号）需要首先获取阿里云百炼的[API权限](https://help.aliyun.com/document_detail/2848578.html)（`AliyunBailianDataFullAccess`或`AliyunBailianDataReadOnlyAccess`均可，已包括sfm:GetAvailableParserTypes权限点），并[加入一个业务空间](https://help.aliyun.com/document_detail/2851098.html)后，方可调用本接口。阿里云账号（主账号）可直接调用无须授权。建议您通过最新版<props="china">[阿里云百炼SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[阿里云百炼SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29)来调用本接口。
//
// - 本接口具有幂等性。
//
// **限流说明：**
//
// 本接口频繁调用会被限流，频率请勿超过10次/秒。如遇限流，请稍后重试。
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
// Retrieves details about a connector. This operation currently supports only file connectors.
//
// Description:
//
// - To call this operation, a RAM user (sub-account) must have the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and must [join a business space](https://help.aliyun.com/document_detail/2851098.html). This requires the `AliyunBailianDataFullAccess` policy, which includes the sfm:GetConnector permission. An Alibaba Cloud account (primary account) can call this operation directly. We recommend using the latest [Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is idempotent.
//
// **Throttling:**
//
// This operation is subject to throttling. Do not exceed a frequency of 5 calls per second. If a request is throttled, try again later.
//
// @param request - GetConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectorResponse
func (client *Client) GetConnectorWithOptions(WorkspaceId *string, request *GetConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorName) {
		query["ConnectorName"] = request.ConnectorName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnector"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/connector"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details about a connector. This operation currently supports only file connectors.
//
// Description:
//
// - To call this operation, a RAM user (sub-account) must have the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and must [join a business space](https://help.aliyun.com/document_detail/2851098.html). This requires the `AliyunBailianDataFullAccess` policy, which includes the sfm:GetConnector permission. An Alibaba Cloud account (primary account) can call this operation directly. We recommend using the latest [Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is idempotent.
//
// **Throttling:**
//
// This operation is subject to throttling. Do not exceed a frequency of 5 calls per second. If a request is throttled, try again later.
//
// @param request - GetConnectorRequest
//
// @return GetConnectorResponse
func (client *Client) GetConnector(WorkspaceId *string, request *GetConnectorRequest) (_result *GetConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConnectorResponse{}
	_body, _err := client.GetConnectorWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current status of a specified knowledge base creation task or knowledge base document append task.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (`AliyunBailianDataFullAccess` or `AliyunBailianDataReadOnlyAccess`, both of which include the sfm:GetIndexJobStatus permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - You must have a knowledge base job that is in progress. You can create a knowledge base creation task by invoking the **SubmitIndexJob*	- operation, or create a knowledge base document append task by invoking the **SubmitIndexAddDocumentsJob*	- operation. Obtain the corresponding `JobId`.
//
// - Invoke this operation at intervals of 5 seconds or more.
//
// - This operation supports idempotence.
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
// Queries the current status of a specified knowledge base creation task or knowledge base document append task.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (`AliyunBailianDataFullAccess` or `AliyunBailianDataReadOnlyAccess`, both of which include the sfm:GetIndexJobStatus permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - You must have a knowledge base job that is in progress. You can create a knowledge base creation task by invoking the **SubmitIndexJob*	- operation, or create a knowledge base document append task by invoking the **SubmitIndexAddDocumentsJob*	- operation. Obtain the corresponding `JobId`.
//
// - Invoke this operation at intervals of 5 seconds or more.
//
// - This operation supports idempotence.
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
// You can call the GetIndexMonitor operation to query monitoring data for a specified knowledge base within a specific time range. This data is crucial for App Performance Analytics, capacity planning, and cost management. The monitoring data includes two main dimensions: storage and retrieval. Storage monitoring retrieves the index storage limit and current usage of the knowledge base. Retrieval monitoring retrieves performance metrics for the query period, such as peak queries per second (QPS), total requests, and average QPS. The metrics are provided as totals and are also broken down by time window. The requests are categorized as successful, failed, and rate-limited.
//
// Description:
//
// <props="intl">
//
// This operation is not available on the Alibaba Cloud International Website (www\\.alibabacloud.com).
//
// <props="china">
//
// - Before you call this operation, a RAM user must obtain the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (which requires the `AliyunBailianDataFullAccess` permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). Alibaba Cloud accounts can call this operation directly without authorization. You can call this operation using the latest version of the [Alibaba Cloud Model Studio software development kit (SDK)](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29). Before you call this operation, make sure that the specified knowledge base has been created and has not been deleted. This means that the knowledge base ID (`IndexId`) must be valid. This operation is idempotent. The maximum query time range (EndTimestamp - StartTimestamp) is 30 days. The granularity of the time window in the returned data is dynamically adjusted based on the query time range.
//
// @param request - GetIndexMonitorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexMonitorResponse
func (client *Client) GetIndexMonitorWithOptions(WorkspaceId *string, request *GetIndexMonitorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIndexMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.IndexId) {
		query["IndexId"] = request.IndexId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIndexMonitor"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/rag/index/monitor"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIndexMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the GetIndexMonitor operation to query monitoring data for a specified knowledge base within a specific time range. This data is crucial for App Performance Analytics, capacity planning, and cost management. The monitoring data includes two main dimensions: storage and retrieval. Storage monitoring retrieves the index storage limit and current usage of the knowledge base. Retrieval monitoring retrieves performance metrics for the query period, such as peak queries per second (QPS), total requests, and average QPS. The metrics are provided as totals and are also broken down by time window. The requests are categorized as successful, failed, and rate-limited.
//
// Description:
//
// <props="intl">
//
// This operation is not available on the Alibaba Cloud International Website (www\\.alibabacloud.com).
//
// <props="china">
//
// - Before you call this operation, a RAM user must obtain the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (which requires the `AliyunBailianDataFullAccess` permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). Alibaba Cloud accounts can call this operation directly without authorization. You can call this operation using the latest version of the [Alibaba Cloud Model Studio software development kit (SDK)](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29). Before you call this operation, make sure that the specified knowledge base has been created and has not been deleted. This means that the knowledge base ID (`IndexId`) must be valid. This operation is idempotent. The maximum query time range (EndTimestamp - StartTimestamp) is 30 days. The granularity of the time window in the returned data is dynamically adjusted based on the query time range.
//
// @param request - GetIndexMonitorRequest
//
// @return GetIndexMonitorResponse
func (client *Client) GetIndexMonitor(WorkspaceId *string, request *GetIndexMonitorRequest) (_result *GetIndexMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexMonitorResponse{}
	_body, _err := client.GetIndexMonitorWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the description of a specified long-term memory.
//
// Description:
//
// - This operation is idempotent.
//
// **Rate limit:*	- Ensure that the interval between two consecutive requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
//
// @param request - GetMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithOptions(workspaceId *string, memoryId *string, request *GetMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
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
// Retrieves the description of a specified long-term memory.
//
// Description:
//
// - This operation is idempotent.
//
// **Rate limit:*	- Ensure that the interval between two consecutive requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
//
// @param request - GetMemoryRequest
//
// @return GetMemoryResponse
func (client *Client) GetMemory(workspaceId *string, memoryId *string, request *GetMemoryRequest) (_result *GetMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryResponse{}
	_body, _err := client.GetMemoryWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a memory fragment.
//
// @param request - GetMemoryNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryNodeResponse
func (client *Client) GetMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, request *GetMemoryNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryNodeResponse, _err error) {
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
// Retrieves a memory fragment.
//
// @param request - GetMemoryNodeRequest
//
// @return GetMemoryNodeResponse
func (client *Client) GetMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string, request *GetMemoryNodeRequest) (_result *GetMemoryNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryNodeResponse{}
	_body, _err := client.GetMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, request, headers, runtime)
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
// Description:
//
// - RAM用户（子账号）需要首先获取阿里云百炼的[API权限](https://help.aliyun.com/document_detail/2848578.html)（`AliyunBailianDataFullAccess`或`AliyunBailianDataReadOnlyAccess`均可，已包括sfm:GetParseSettings权限点），并[加入一个业务空间](https://help.aliyun.com/document_detail/2851098.html)后，方可调用本接口。阿里云账号（主账号）可直接调用无须授权。建议您通过最新版<props="china">[阿里云百炼SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[阿里云百炼SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29)来调用本接口。
//
// - 本接口具有幂等性。
//
// **限流说明：**
//
// 本接口频繁调用会被限流，频率请勿超过10次/秒。如遇限流，请稍后重试。
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
// Description:
//
// - RAM用户（子账号）需要首先获取阿里云百炼的[API权限](https://help.aliyun.com/document_detail/2848578.html)（`AliyunBailianDataFullAccess`或`AliyunBailianDataReadOnlyAccess`均可，已包括sfm:GetParseSettings权限点），并[加入一个业务空间](https://help.aliyun.com/document_detail/2851098.html)后，方可调用本接口。阿里云账号（主账号）可直接调用无须授权。建议您通过最新版<props="china">[阿里云百炼SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[阿里云百炼SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29)来调用本接口。
//
// - 本接口具有幂等性。
//
// **限流说明：**
//
// 本接口频繁调用会被限流，频率请勿超过10次/秒。如遇限流，请稍后重试。
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
// @param request - GetPromptTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPromptTemplateResponse
func (client *Client) GetPromptTemplateWithOptions(workspaceId *string, promptTemplateId *string, request *GetPromptTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPromptTemplateResponse, _err error) {
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
// @param request - GetPromptTemplateRequest
//
// @return GetPromptTemplateResponse
func (client *Client) GetPromptTemplate(workspaceId *string, promptTemplateId *string, request *GetPromptTemplateRequest) (_result *GetPromptTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPromptTemplateResponse{}
	_body, _err := client.GetPromptTemplateWithOptions(workspaceId, promptTemplateId, request, headers, runtime)
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
// @param request - GetPublishedAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublishedAgentResponse
func (client *Client) GetPublishedAgentWithOptions(workspaceId *string, appCode *string, request *GetPublishedAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPublishedAgentResponse, _err error) {
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
// @param request - GetPublishedAgentRequest
//
// @return GetPublishedAgentResponse
func (client *Client) GetPublishedAgent(workspaceId *string, appCode *string, request *GetPublishedAgentRequest) (_result *GetPublishedAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPublishedAgentResponse{}
	_body, _err := client.GetPublishedAgentWithOptions(workspaceId, appCode, request, headers, runtime)
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
// Retrieves the details of one or more categories in a specified workspace.
//
// Description:
//
// - This API does not support querying data tables.
//
// - To call this API, a RAM user must first obtain the required [API permission](https://help.aliyun.com/document_detail/2848578.html) for Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). This requires the `AliyunBailianDataFullAccess` permission, which includes the `sfm:ListCategory` permission. Alibaba Cloud accounts can call this API directly. Use the latest version of the <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation supports pagination. When making your first request, set the `MaxResults` parameter to specify the maximum number of items to return. If more items are available, the response includes a `NextToken`. To retrieve the next page of results, set the `NextToken` parameter to the value from the previous response and specify `MaxResults` again. An empty `NextToken` indicates that no more results are available.
//
// - This operation is idempotent.
//
// **Rate limiting:*	- This API is subject to rate limiting. Do not exceed 5 requests per second. If the system throttles a request, retry it after a short interval.
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

	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
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
// Retrieves the details of one or more categories in a specified workspace.
//
// Description:
//
// - This API does not support querying data tables.
//
// - To call this API, a RAM user must first obtain the required [API permission](https://help.aliyun.com/document_detail/2848578.html) for Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). This requires the `AliyunBailianDataFullAccess` permission, which includes the `sfm:ListCategory` permission. Alibaba Cloud accounts can call this API directly. Use the latest version of the <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation supports pagination. When making your first request, set the `MaxResults` parameter to specify the maximum number of items to return. If more items are available, the response includes a `NextToken`. To retrieve the next page of results, set the `NextToken` parameter to the value from the previous response and specify `MaxResults` again. An empty `NextToken` indicates that no more results are available.
//
// - This operation is idempotent.
//
// **Rate limiting:*	- This API is subject to rate limiting. Do not exceed 5 requests per second. If the system throttles a request, retry it after a short interval.
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
// Queries the list and information of text chunks.
//
// Description:
//
// - For document search<props="china"> or audio/video search knowledge bases, this operation queries all chunks of a specified file. For data query or image Q&A knowledge bases, this operation retrieves information about all text chunks.
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ChunkList permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation has idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Queries the list and information of text chunks.
//
// Description:
//
// - For document search<props="china"> or audio/video search knowledge bases, this operation queries all chunks of a specified file. For data query or image Q&A knowledge bases, this operation retrieves information about all text chunks.
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ChunkList permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation has idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Retrieves the details of one or more documents in a specified category.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListFile permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Invoke this operation by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - For paging on the first page, only set `MaxResults` to limit the number of entries returned. The `NextToken` in the response serves as the credential for querying subsequent pages. When querying subsequent pages, set the `NextToken` parameter to the `NextToken` value obtained from the previous response (if `NextToken` is empty, all results have been returned and no further requests are needed), and set `MaxResults` to limit the number of entries returned.
//
// - This operation is idempotent.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 5 calls per second. If throttled, retry later.
//
// @param tmpReq - ListFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileResponse
func (client *Client) ListFileWithOptions(WorkspaceId *string, tmpReq *ListFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.FileIdsShrink) {
		query["FileIds"] = request.FileIdsShrink
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
// Retrieves the details of one or more documents in a specified category.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListFile permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Invoke this operation by using the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - For paging on the first page, only set `MaxResults` to limit the number of entries returned. The `NextToken` in the response serves as the credential for querying subsequent pages. When querying subsequent pages, set the `NextToken` parameter to the `NextToken` value obtained from the previous response (if `NextToken` is empty, all results have been returned and no further requests are needed), and set `MaxResults` to limit the number of entries returned.
//
// - This operation is idempotent.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 5 calls per second. If throttled, retry later.
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
// Retrieves files and their summary information from a specified knowledge base.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListIndexFiles permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation has idempotence.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 15 calls per second. If you are throttled, retry later.
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
// Retrieves files and their summary information from a specified knowledge base.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListIndexFiles permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation has idempotence.
//
// **Throttling:**
//
// This operation is throttled if called too frequently. Do not exceed 15 calls per second. If you are throttled, retry later.
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
// Retrieves the files in a specified knowledge base along with their details.
//
// Description:
//
// - RAM users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListIndexFiles permission), before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - Before calling this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation is idempotent.
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
// Retrieves the files in a specified knowledge base along with their details.
//
// Description:
//
// - RAM users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListIndexFiles permission), before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - Before calling this operation, make sure that your knowledge base has been created and has not been deleted (that is, the knowledge base ID `IndexId` is valid).
//
// - This operation is idempotent.
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
// Retrieves the list of knowledge bases in a specified workspace.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListIndex permission) before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is idempotent.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// Retrieves the list of knowledge bases in a specified workspace.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:ListIndex permission) before calling this operation. Alibaba Cloud accounts can call this operation directly without authorization. Use the latest <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is idempotent.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// Retrieves the details of one or more long-term memory entities in a specified workspace.
//
// Description:
//
// - When querying the first page of a paging query, set only `MaxResults` to limit the number of entries returned. The `NextToken` value in the response serves as the credential for querying subsequent pages. When querying subsequent pages, set the `NextToken` parameter to the `NextToken` value obtained from the previous response as the query credential (if `NextToken` is empty, all results have been returned and no further requests are needed), and settings `MaxResults` to limit the number of entries returned.
//
// - This operation supports idempotence.
//
// **Rate limit:*	- Ensure that the interval between two consecutive requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
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
// Retrieves the details of one or more long-term memory entities in a specified workspace.
//
// Description:
//
// - When querying the first page of a paging query, set only `MaxResults` to limit the number of entries returned. The `NextToken` value in the response serves as the credential for querying subsequent pages. When querying subsequent pages, set the `NextToken` parameter to the `NextToken` value obtained from the previous response as the query credential (if `NextToken` is empty, all results have been returned and no further requests are needed), and settings `MaxResults` to limit the number of entries returned.
//
// - This operation supports idempotence.
//
// **Rate limit:*	- Ensure that the interval between two consecutive requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
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
// Retrieves a list of memory nodes.
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
// Retrieves a list of memory nodes.
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
// Retrieves information from a specified knowledge base.
//
// Description:
//
// <props="china">
//
// - **How to call**: To retrieve information from a knowledge base, use the latest [Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29) with an [AccessKey](https://help.aliyun.com/document_detail/116401.html) or [Spring AI Alibaba](https://help.aliyun.com/document_detail/2990886.html) with an Alibaba Cloud Model Studio [API key](https://help.aliyun.com/document_detail/2712195.html). Both tools simplify your API calls by handling the complex signature calculation.
//
// - **Required permissions**:
//
//   - **RAM user (sub-account)**: To call this API, a RAM user must be granted [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). You can use the `AliyunBailianDataFullAccess` policy, which includes the required `sfm:Retrieve` permission.
//
//   - **Alibaba Cloud account (main account)**: This account has the required permissions by default and can call the API directly.
//
// - **Response latency**: This API call involves complex retrieval and matching operations, which can cause longer response times. We recommend configuring appropriate request timeouts and retry strategies.
//
// - **Idempotency**: This API is idempotent.
//
// <props="intl">
//
// - **How to call**: We recommend using the latest [Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this API. The SDK simplifies API calls by handling the complex signature calculation.
//
// - **Required permissions**:
//
//   - **RAM user (sub-account)**: To call this API, a RAM user must be granted [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). You can use the `AliyunBailianDataFullAccess` policy, which includes the required `sfm:Retrieve` permission.
//
//   - **Alibaba Cloud account (main account)**: This account has the required permissions by default and can call the API directly.
//
// - **Response latency**: This API call involves complex retrieval and matching operations, which can cause longer response times. We recommend configuring appropriate request timeouts and retry strategies.
//
// - **Idempotency**: This API is idempotent.
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
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

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

	if !dara.IsNil(request.ExtraShrink) {
		query["Extra"] = request.ExtraShrink
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
// Retrieves information from a specified knowledge base.
//
// Description:
//
// <props="china">
//
// - **How to call**: To retrieve information from a knowledge base, use the latest [Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29) with an [AccessKey](https://help.aliyun.com/document_detail/116401.html) or [Spring AI Alibaba](https://help.aliyun.com/document_detail/2990886.html) with an Alibaba Cloud Model Studio [API key](https://help.aliyun.com/document_detail/2712195.html). Both tools simplify your API calls by handling the complex signature calculation.
//
// - **Required permissions**:
//
//   - **RAM user (sub-account)**: To call this API, a RAM user must be granted [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). You can use the `AliyunBailianDataFullAccess` policy, which includes the required `sfm:Retrieve` permission.
//
//   - **Alibaba Cloud account (main account)**: This account has the required permissions by default and can call the API directly.
//
// - **Response latency**: This API call involves complex retrieval and matching operations, which can cause longer response times. We recommend configuring appropriate request timeouts and retry strategies.
//
// - **Idempotency**: This API is idempotent.
//
// <props="intl">
//
// - **How to call**: We recommend using the latest [Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this API. The SDK simplifies API calls by handling the complex signature calculation.
//
// - **Required permissions**:
//
//   - **RAM user (sub-account)**: To call this API, a RAM user must be granted [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). You can use the `AliyunBailianDataFullAccess` policy, which includes the required `sfm:Retrieve` permission.
//
//   - **Alibaba Cloud account (main account)**: This account has the required permissions by default and can call the API directly.
//
// - **Response latency**: This API call involves complex retrieval and matching operations, which can cause longer response times. We recommend configuring appropriate request timeouts and retry strategies.
//
// - **Idempotency**: This API is idempotent.
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
// Adds parsed files to the specified knowledge base.
//
// Description:
//
// <props="china">
//
// - This API does not support knowledge bases for data queries or image Q\\&A. To update these knowledge bases, see the [knowledge base](https://help.aliyun.com/document_detail/2807740.html) documentation.
//
// <props="intl">
//
// - This API does not support knowledge bases for data queries or image Q\\&A. To update these knowledge bases, see the [knowledge base](https://help.aliyun.com/document_detail/2807740.html) documentation.
//
// - A RAM user (sub-account) can call this API only after being granted the required [api permission](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (specifically, the `AliyunBailianDataFullAccess` policy, which includes the `sfm:SubmitIndexAddDocumentsJob` permission) and joining a [workspace](https://help.aliyun.com/document_detail/2851098.html). An Alibaba Cloud account can call this API directly without authorization. We recommend using the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this API.
//
// - Before calling this API, ensure your knowledge base exists and has a valid knowledge base ID (`IndexId`).
//
// - Before calling this API, you must first upload files to Alibaba Cloud Model Studio using the **AddFile*	- API.
//
// - After calling this API, the job runs in the background and may take several hours to complete, especially during peak times. Do not submit duplicate requests until the job is complete. To check the job status, call the **GetIndexJobStatus*	- API. The `Documents` file list returned by the GetIndexJobStatus API contains all files for the job, which is uniquely identified by the `job_id` you provided. You can check this list to verify whether each file was imported (parsed) successfully. Note that frequent calls to the GetIndexJobStatus API are subject to rate limiting. Do not exceed 20 calls per minute.
//
// - A successful API call indicates the job has been submitted for processing, which takes time. This API is not idempotent, so do not send duplicate requests; doing so will create multiple jobs.
//
// **Rate limiting:*	- This API is limited to 10 calls per second. If you exceed this limit, wait before retrying.
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

	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
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

	if !dara.IsNil(request.ExtraShrink) {
		query["Extra"] = request.ExtraShrink
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
// Adds parsed files to the specified knowledge base.
//
// Description:
//
// <props="china">
//
// - This API does not support knowledge bases for data queries or image Q\\&A. To update these knowledge bases, see the [knowledge base](https://help.aliyun.com/document_detail/2807740.html) documentation.
//
// <props="intl">
//
// - This API does not support knowledge bases for data queries or image Q\\&A. To update these knowledge bases, see the [knowledge base](https://help.aliyun.com/document_detail/2807740.html) documentation.
//
// - A RAM user (sub-account) can call this API only after being granted the required [api permission](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (specifically, the `AliyunBailianDataFullAccess` policy, which includes the `sfm:SubmitIndexAddDocumentsJob` permission) and joining a [workspace](https://help.aliyun.com/document_detail/2851098.html). An Alibaba Cloud account can call this API directly without authorization. We recommend using the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this API.
//
// - Before calling this API, ensure your knowledge base exists and has a valid knowledge base ID (`IndexId`).
//
// - Before calling this API, you must first upload files to Alibaba Cloud Model Studio using the **AddFile*	- API.
//
// - After calling this API, the job runs in the background and may take several hours to complete, especially during peak times. Do not submit duplicate requests until the job is complete. To check the job status, call the **GetIndexJobStatus*	- API. The `Documents` file list returned by the GetIndexJobStatus API contains all files for the job, which is uniquely identified by the `job_id` you provided. You can check this list to verify whether each file was imported (parsed) successfully. Note that frequent calls to the GetIndexJobStatus API are subject to rate limiting. Do not exceed 20 calls per minute.
//
// - A successful API call indicates the job has been submitted for processing, which takes time. This API is not idempotent, so do not send duplicate requests; doing so will create multiple jobs.
//
// **Rate limiting:*	- This API is limited to 10 calls per second. If you exceed this limit, wait before retrying.
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
// Submits a specified CreateIndex task to complete knowledge base creation.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:SubmitIndexJob permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, you must invoke the **CreateIndex*	- operation and obtain the corresponding `IndexId`.
//
// - After invoking this operation, the node requires time to execute and may take several hours during peak periods. Do not submit duplicate requests before the node completes. To query the node execution status, invoke the **GetIndexJobStatus*	- operation.
//
// - After the knowledge base is created, you can associate it with an agent application<props="china"> or workflow application in the same workspace through <props="china">[Application Management](https://bailian.console.aliyun.com/?tab=app#/app-center)<props="intl">[Application Management](https://modelstudio.console.alibabacloud.com/?tab=app#/app-center) (or pass the `IndexID` through `rag_options` in [Application Calls](https://help.aliyun.com/document_detail/2846132.html)) to supplement your Model Studio application with private knowledge and up-to-date information. You can also choose not to use a Model Studio application and directly query the knowledge base by invoking the **Retrieve*	- operation.
//
// - This operation does not support idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// Submits a specified CreateIndex task to complete knowledge base creation.
//
// Description:
//
// - Resource Access Management (RAM) users must first obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (requires `AliyunBailianDataFullAccess`, which includes the sfm:SubmitIndexJob permission), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before invoking this operation. Alibaba Cloud accounts can invoke this operation directly without authorization. Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to invoke this operation.
//
// - Before invoking this operation, you must invoke the **CreateIndex*	- operation and obtain the corresponding `IndexId`.
//
// - After invoking this operation, the node requires time to execute and may take several hours during peak periods. Do not submit duplicate requests before the node completes. To query the node execution status, invoke the **GetIndexJobStatus*	- operation.
//
// - After the knowledge base is created, you can associate it with an agent application<props="china"> or workflow application in the same workspace through <props="china">[Application Management](https://bailian.console.aliyun.com/?tab=app#/app-center)<props="intl">[Application Management](https://modelstudio.console.alibabacloud.com/?tab=app#/app-center) (or pass the `IndexID` through `rag_options` in [Application Calls](https://help.aliyun.com/document_detail/2846132.html)) to supplement your Model Studio application with private knowledge and up-to-date information. You can also choose not to use a Model Studio application and directly query the knowledge base by invoking the **Retrieve*	- operation.
//
// - This operation does not support idempotence.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If throttled, retry later.
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
// Modifies the content and title of a specified text chunk in a knowledge base, and specifies whether the chunk participates in knowledge base retrieval.
//
// Description:
//
// - **Key limits**: This operation supports only document search knowledge bases. Data query and image Q&A knowledge bases are not supported.
//
// - **Permission requirements**:
//
//   - **Resource Access Management (RAM) user**: Before invoking this operation, obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (you can use the `AliyunBailianDataFullAccess` policy, which includes the sfm:UpdateChunk permission required by this operation), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//   - **Alibaba Cloud account**: Has permissions by default and can invoke this operation directly.
//
// - **Calling method**: Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK provides encapsulation of complex signature calculation logic and simplifies the invocation procedure.
//
// - **Effective latency**: Updates typically take effect immediately. During peak hours, a slight delay (seconds) may occur.
//
// - **Idempotence**: This operation is idempotent. If you repeat the operation on a text chunk that has already been updated, the operation returns a success response.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Modifies the content and title of a specified text chunk in a knowledge base, and specifies whether the chunk participates in knowledge base retrieval.
//
// Description:
//
// - **Key limits**: This operation supports only document search knowledge bases. Data query and image Q&A knowledge bases are not supported.
//
// - **Permission requirements**:
//
//   - **Resource Access Management (RAM) user**: Before invoking this operation, obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Model Studio (you can use the `AliyunBailianDataFullAccess` policy, which includes the sfm:UpdateChunk permission required by this operation), and [join a workspace](https://help.aliyun.com/document_detail/2851098.html).
//
//   - **Alibaba Cloud account**: Has permissions by default and can invoke this operation directly.
//
// - **Calling method**: Use the latest <props="china">[Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29). The SDK provides encapsulation of complex signature calculation logic and simplifies the invocation procedure.
//
// - **Effective latency**: Updates typically take effect immediately. During peak hours, a slight delay (seconds) may occur.
//
// - **Idempotence**: This operation is idempotent. If you repeat the operation on a text chunk that has already been updated, the operation returns a success response.
//
// **Rate limit:**
//
// This operation is throttled if called too frequently. Do not exceed 10 calls per second. If you are throttled, retry later.
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
// Updates a connector.
//
// Description:
//
// - A RAM user can call this operation only after they join a workspace and are granted the required [API permission](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (Bailian). The `AliyunBailianDataFullAccess` policy, which includes the `sfm:UpdateConnector` permission, is required. An Alibaba Cloud account can call this operation directly. Use the latest version of the <props="china">[Alibaba Cloud Model Studio (Bailian) SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio (Bailian) SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is not idempotent.
//
// **Throttling:*	- If you call this operation too frequently, the system may throttle your requests. Do not exceed a frequency of 5 calls per second. If a request is throttled, try again later.
//
// @param request - UpdateConnectorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnectorResponse
func (client *Client) UpdateConnectorWithOptions(WorkspaceId *string, ConnectorId *string, request *UpdateConnectorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorName) {
		body["ConnectorName"] = request.ConnectorName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnector"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/connector/" + dara.PercentEncode(dara.StringValue(ConnectorId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a connector.
//
// Description:
//
// - A RAM user can call this operation only after they join a workspace and are granted the required [API permission](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (Bailian). The `AliyunBailianDataFullAccess` policy, which includes the `sfm:UpdateConnector` permission, is required. An Alibaba Cloud account can call this operation directly. Use the latest version of the <props="china">[Alibaba Cloud Model Studio (Bailian) SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio (Bailian) SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// - This operation is not idempotent.
//
// **Throttling:*	- If you call this operation too frequently, the system may throttle your requests. Do not exceed a frequency of 5 calls per second. If a request is throttled, try again later.
//
// @param request - UpdateConnectorRequest
//
// @return UpdateConnectorResponse
func (client *Client) UpdateConnector(WorkspaceId *string, ConnectorId *string, request *UpdateConnectorRequest) (_result *UpdateConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConnectorResponse{}
	_body, _err := client.UpdateConnectorWithOptions(WorkspaceId, ConnectorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the tags for a specified file.
//
// Description:
//
// - A RAM User (sub-account) must be granted the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (the `AliyunBailianDataFullAccess` policy, which includes the `sfm:UpdateFileTag` permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. An Alibaba Cloud account (main account) can call this operation directly without authorization. We recommend using the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// **Throttling:*	- Do not call this operation more than 5 times per second. If a request is throttled, try again later.
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
// Updates the tags for a specified file.
//
// Description:
//
// - A RAM User (sub-account) must be granted the required [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio (the `AliyunBailianDataFullAccess` policy, which includes the `sfm:UpdateFileTag` permission) and [join a workspace](https://help.aliyun.com/document_detail/2851098.html) before calling this operation. An Alibaba Cloud account (main account) can call this operation directly without authorization. We recommend using the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29) to call this operation.
//
// **Throttling:*	- Do not call this operation more than 5 times per second. If a request is throttled, try again later.
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
// Updates the configuration of a specified knowledge base.
//
// Description:
//
// <props="intl">This operation is not available on the Alibaba Cloud International Website (www\\.alibabacloud.com).<props="china">
//
// Before a RAM user can call this operation, the RAM user must have the `AliyunBailianDataFullAccess` permission for Alibaba Cloud Model Studio. For more information, see [Grant permissions](https://help.aliyun.com/document_detail/2848578.html). The RAM user must also be added to a workspace. For more information, see [Add a member to a workspace](https://help.aliyun.com/document_detail/2851098.html). An Alibaba Cloud account can call this operation without authorization. Use the latest version of the Alibaba Cloud Model Studio SDK to call this operation. For more information, see [Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29).
//
// Before you call this operation, ensure that the knowledge base is created and has not been deleted. The knowledge base ID (`Id`) must be valid.
//
// This operation is idempotent.
//
// @param request - UpdateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIndexResponse
func (client *Client) UpdateIndexWithOptions(WorkspaceId *string, request *UpdateIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DenseSimilarityTopK) {
		query["DenseSimilarityTopK"] = request.DenseSimilarityTopK
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PipelineCommercialCu) {
		query["PipelineCommercialCu"] = request.PipelineCommercialCu
	}

	if !dara.IsNil(request.PipelineCommercialType) {
		query["PipelineCommercialType"] = request.PipelineCommercialType
	}

	if !dara.IsNil(request.RerankMinScore) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !dara.IsNil(request.SparseSimilarityTopK) {
		query["SparseSimilarityTopK"] = request.SparseSimilarityTopK
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIndex"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/index/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specified knowledge base.
//
// Description:
//
// <props="intl">This operation is not available on the Alibaba Cloud International Website (www\\.alibabacloud.com).<props="china">
//
// Before a RAM user can call this operation, the RAM user must have the `AliyunBailianDataFullAccess` permission for Alibaba Cloud Model Studio. For more information, see [Grant permissions](https://help.aliyun.com/document_detail/2848578.html). The RAM user must also be added to a workspace. For more information, see [Add a member to a workspace](https://help.aliyun.com/document_detail/2851098.html). An Alibaba Cloud account can call this operation without authorization. Use the latest version of the Alibaba Cloud Model Studio SDK to call this operation. For more information, see [Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29).
//
// Before you call this operation, ensure that the knowledge base is created and has not been deleted. The knowledge base ID (`Id`) must be valid.
//
// This operation is idempotent.
//
// @param request - UpdateIndexRequest
//
// @return UpdateIndexResponse
func (client *Client) UpdateIndex(WorkspaceId *string, request *UpdateIndexRequest) (_result *UpdateIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIndexResponse{}
	_body, _err := client.UpdateIndexWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the description of a specified long-term memory.
//
// Description:
//
// - This operation is idempotent.
//
// **Rate limit:*	- Ensure that the interval between two requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
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
// Updates the description of a specified long-term memory.
//
// Description:
//
// - This operation is idempotent.
//
// **Rate limit:*	- Ensure that the interval between two requests is at least 1 second. Otherwise, throttling may be triggered. If throttling occurs, retry later.
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
// Updates a memory fragment.
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
// Updates a memory fragment.
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

// Summary:
//
// Update a table in an Alibaba Cloud Model Studio data connector using a file from an authorized OSS bucket.
//
// Description:
//
// - Ensure that the OSS bucket belongs to the same Alibaba Cloud account as your Alibaba Cloud Model Studio instance. You must also complete the authorization steps described in [Configure data import from OSS](https://help.aliyun.com/document_detail/2782155.html).
//
//   - The bucket storage class must not be Archive, Cold Archive, or Deep Cold Archive. Buckets with server-side encryption are supported. public-read-write, public-read, and private buckets are also supported.
//
//   - If you use a bucket with [Referer-based hotlink protection](https://help.aliyun.com/document_detail/2636937.html), you must add `*.console.aliyun.com` to the Referer whitelist. For more information, see [Allow access only from trusted websites](https://help.aliyun.com/document_detail/2636937.html).
//
// - Before a RAM user (sub-account) can call this operation, they must obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). The AliyunBailianDataFullAccess policy includes the required `sfm:UpdateTableFromAuthorizedOss` permission. An Alibaba Cloud account (root account) can call this operation directly without additional permissions. We recommend that you use the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29) or <props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Rate limiting:*	- This operation is subject to rate limiting. Do not call it more than five times per second. If you reach the limit, wait before you try again.
//
// @param request - UpdateTableFromAuthorizedOssRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableFromAuthorizedOssResponse
func (client *Client) UpdateTableFromAuthorizedOssWithOptions(WorkspaceId *string, TableId *string, request *UpdateTableFromAuthorizedOssRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTableFromAuthorizedOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OssBucket) {
		body["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssKey) {
		body["OssKey"] = request.OssKey
	}

	if !dara.IsNil(request.OssRegionId) {
		body["OssRegionId"] = request.OssRegionId
	}

	if !dara.IsNil(request.UpdateMode) {
		body["UpdateMode"] = request.UpdateMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableFromAuthorizedOss"),
		Version:     dara.String("2023-12-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/datacenter/table/fromoss/" + dara.PercentEncode(dara.StringValue(TableId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableFromAuthorizedOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a table in an Alibaba Cloud Model Studio data connector using a file from an authorized OSS bucket.
//
// Description:
//
// - Ensure that the OSS bucket belongs to the same Alibaba Cloud account as your Alibaba Cloud Model Studio instance. You must also complete the authorization steps described in [Configure data import from OSS](https://help.aliyun.com/document_detail/2782155.html).
//
//   - The bucket storage class must not be Archive, Cold Archive, or Deep Cold Archive. Buckets with server-side encryption are supported. public-read-write, public-read, and private buckets are also supported.
//
//   - If you use a bucket with [Referer-based hotlink protection](https://help.aliyun.com/document_detail/2636937.html), you must add `*.console.aliyun.com` to the Referer whitelist. For more information, see [Allow access only from trusted websites](https://help.aliyun.com/document_detail/2636937.html).
//
// - Before a RAM user (sub-account) can call this operation, they must obtain [API permissions](https://help.aliyun.com/document_detail/2848578.html) for Alibaba Cloud Model Studio and [join a workspace](https://help.aliyun.com/document_detail/2851098.html). The AliyunBailianDataFullAccess policy includes the required `sfm:UpdateTableFromAuthorizedOss` permission. An Alibaba Cloud account (root account) can call this operation directly without additional permissions. We recommend that you use the latest version of the <props="china">[Alibaba Cloud Model Studio SDK](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29) or <props="intl">[Alibaba Cloud Model Studio SDK](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29).
//
// - This operation is not idempotent.
//
// **Rate limiting:*	- This operation is subject to rate limiting. Do not call it more than five times per second. If you reach the limit, wait before you try again.
//
// @param request - UpdateTableFromAuthorizedOssRequest
//
// @return UpdateTableFromAuthorizedOssResponse
func (client *Client) UpdateTableFromAuthorizedOss(WorkspaceId *string, TableId *string, request *UpdateTableFromAuthorizedOssRequest) (_result *UpdateTableFromAuthorizedOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTableFromAuthorizedOssResponse{}
	_body, _err := client.UpdateTableFromAuthorizedOssWithOptions(WorkspaceId, TableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
