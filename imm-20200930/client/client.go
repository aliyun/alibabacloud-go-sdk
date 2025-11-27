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
		"cn-beijing-gov-1": dara.String("imm-vpc.cn-beijing-gov-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("imm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds mosaics, Gaussian blurs, or solid color shapes to blur one or more areas of an image for privacy protection and saves the output image to the specified path in Object Storage Service (OSS).
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
//		- The operation accepts JPG and PNG images with a maximum side length of 30,000 pixels and a total of up to 250 million pixels.
//
// @param tmpReq - AddImageMosaicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageMosaicResponse
func (client *Client) AddImageMosaicWithOptions(tmpReq *AddImageMosaicRequest, runtime *dara.RuntimeOptions) (_result *AddImageMosaicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddImageMosaicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Targets) {
		request.TargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Targets, dara.String("Targets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ImageFormat) {
		query["ImageFormat"] = request.ImageFormat
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Quality) {
		query["Quality"] = request.Quality
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.TargetsShrink) {
		query["Targets"] = request.TargetsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImageMosaic"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImageMosaicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds mosaics, Gaussian blurs, or solid color shapes to blur one or more areas of an image for privacy protection and saves the output image to the specified path in Object Storage Service (OSS).
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
//		- The operation accepts JPG and PNG images with a maximum side length of 30,000 pixels and a total of up to 250 million pixels.
//
// @param request - AddImageMosaicRequest
//
// @return AddImageMosaicResponse
func (client *Client) AddImageMosaic(request *AddImageMosaicRequest) (_result *AddImageMosaicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddImageMosaicResponse{}
	_body, _err := client.AddImageMosaicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds objects to a story.
//
// @param tmpReq - AddStoryFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddStoryFilesResponse
func (client *Client) AddStoryFilesWithOptions(tmpReq *AddStoryFilesRequest, runtime *dara.RuntimeOptions) (_result *AddStoryFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddStoryFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FilesShrink) {
		body["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddStoryFiles"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddStoryFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds objects to a story.
//
// @param request - AddStoryFilesRequest
//
// @return AddStoryFilesResponse
func (client *Client) AddStoryFiles(request *AddStoryFilesRequest) (_result *AddStoryFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddStoryFilesResponse{}
	_body, _err := client.AddStoryFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds an Object Storage Service (OSS) bucket to the specified project. The binding enables you to use IMM features by using the x-oss-process parameter.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- To use data processing capabilities of IMM based on the x-oss-process parameter, you must bind an OSS bucket to an IMM project. For more information, see [x-oss-process](https://help.aliyun.com/document_detail/2391270.html).
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
// @param request - AttachOSSBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachOSSBucketResponse
func (client *Client) AttachOSSBucketWithOptions(request *AttachOSSBucketRequest, runtime *dara.RuntimeOptions) (_result *AttachOSSBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.OSSBucket) {
		query["OSSBucket"] = request.OSSBucket
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachOSSBucket"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachOSSBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds an Object Storage Service (OSS) bucket to the specified project. The binding enables you to use IMM features by using the x-oss-process parameter.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- To use data processing capabilities of IMM based on the x-oss-process parameter, you must bind an OSS bucket to an IMM project. For more information, see [x-oss-process](https://help.aliyun.com/document_detail/2391270.html).
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
// @param request - AttachOSSBucketRequest
//
// @return AttachOSSBucketResponse
func (client *Client) AttachOSSBucket(request *AttachOSSBucketRequest) (_result *AttachOSSBucketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachOSSBucketResponse{}
	_body, _err := client.AttachOSSBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the metadata of multiple files from a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- A successful deletion message is returned regardless of whether the metadata of the file exists in the dataset.
//
// >
//
//   - If you delete the metadata of a file from a dataset, the file stored in Object Storage Service (OSS) or Photo and Drive Service is **not*	- deleted. If you want to delete the file, use the operations provided by OSS or Photo and Drive Service.
//
//   - Metadata deletion affects existing face groups and stories but does not affect existing spatiotemporal groups.
//
// @param tmpReq - BatchDeleteFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteFileMetaResponse
func (client *Client) BatchDeleteFileMetaWithOptions(tmpReq *BatchDeleteFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeleteFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.URIs) {
		request.URIsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.URIs, dara.String("URIs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.URIsShrink) {
		query["URIs"] = request.URIsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the metadata of multiple files from a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- A successful deletion message is returned regardless of whether the metadata of the file exists in the dataset.
//
// >
//
//   - If you delete the metadata of a file from a dataset, the file stored in Object Storage Service (OSS) or Photo and Drive Service is **not*	- deleted. If you want to delete the file, use the operations provided by OSS or Photo and Drive Service.
//
//   - Metadata deletion affects existing face groups and stories but does not affect existing spatiotemporal groups.
//
// @param request - BatchDeleteFileMetaRequest
//
// @return BatchDeleteFileMetaResponse
func (client *Client) BatchDeleteFileMeta(request *BatchDeleteFileMetaRequest) (_result *BatchDeleteFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteFileMetaResponse{}
	_body, _err := client.BatchDeleteFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries face clusters.
//
// @param tmpReq - BatchGetFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetFigureClusterResponse
func (client *Client) BatchGetFigureClusterWithOptions(tmpReq *BatchGetFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *BatchGetFigureClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetFigureClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectIds) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, dara.String("ObjectIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ObjectIdsShrink) {
		query["ObjectIds"] = request.ObjectIdsShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetFigureCluster"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries face clusters.
//
// @param request - BatchGetFigureClusterRequest
//
// @return BatchGetFigureClusterResponse
func (client *Client) BatchGetFigureCluster(request *BatchGetFigureClusterRequest) (_result *BatchGetFigureClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetFigureClusterResponse{}
	_body, _err := client.BatchGetFigureClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metadata of multiple objects or files in the specified dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, feel free to join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// @param tmpReq - BatchGetFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetFileMetaResponse
func (client *Client) BatchGetFileMetaWithOptions(tmpReq *BatchGetFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchGetFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.URIs) {
		request.URIsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.URIs, dara.String("URIs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WithFields) {
		request.WithFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WithFields, dara.String("WithFields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.URIsShrink) {
		query["URIs"] = request.URIsShrink
	}

	if !dara.IsNil(request.WithFieldsShrink) {
		query["WithFields"] = request.WithFieldsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metadata of multiple objects or files in the specified dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, feel free to join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// @param request - BatchGetFileMetaRequest
//
// @return BatchGetFileMetaResponse
func (client *Client) BatchGetFileMeta(request *BatchGetFileMetaRequest) (_result *BatchGetFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetFileMetaResponse{}
	_body, _err := client.BatchGetFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Indexes metadata of multiple objects into the specified dataset. The process involves data processing operations such as label detection, face detection, and location detection. Metadata indexing helps meet diverse data retrieval requirements.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Data processing operations supported for metadata processing vary with workflow templates. For more information, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
//
//		- Metadata indexing poses limits on the total number and size of objects. For more information about these limits, see [Limits](https://help.aliyun.com/document_detail/475569.html). For more information about how to create
//
//		- Metadata indexing is available in specific regions. For information about regions that support metadata indexing, see the "Data management and indexing" section of the [Limits](https://help.aliyun.com/document_detail/475569.html) topic.
//
// @param tmpReq - BatchIndexFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchIndexFileMetaResponse
func (client *Client) BatchIndexFileMetaWithOptions(tmpReq *BatchIndexFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchIndexFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchIndexFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FilesShrink) {
		query["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchIndexFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchIndexFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Indexes metadata of multiple objects into the specified dataset. The process involves data processing operations such as label detection, face detection, and location detection. Metadata indexing helps meet diverse data retrieval requirements.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Data processing operations supported for metadata processing vary with workflow templates. For more information, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
//
//		- Metadata indexing poses limits on the total number and size of objects. For more information about these limits, see [Limits](https://help.aliyun.com/document_detail/475569.html). For more information about how to create
//
//		- Metadata indexing is available in specific regions. For information about regions that support metadata indexing, see the "Data management and indexing" section of the [Limits](https://help.aliyun.com/document_detail/475569.html) topic.
//
// @param request - BatchIndexFileMetaRequest
//
// @return BatchIndexFileMetaResponse
func (client *Client) BatchIndexFileMeta(request *BatchIndexFileMetaRequest) (_result *BatchIndexFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchIndexFileMetaResponse{}
	_body, _err := client.BatchIndexFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates some metadata items of files indexed into a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- You cannot call this operation to update all metadata. You can update only metadata fields such as CustomLabels, CustomId, and Figures. For more information, see the "Request parameters" section of this topic.
//
// @param tmpReq - BatchUpdateFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateFileMetaResponse
func (client *Client) BatchUpdateFileMetaWithOptions(tmpReq *BatchUpdateFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FilesShrink) {
		query["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates some metadata items of files indexed into a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- You cannot call this operation to update all metadata. You can update only metadata fields such as CustomLabels, CustomId, and Figures. For more information, see the "Request parameters" section of this topic.
//
// @param request - BatchUpdateFileMetaRequest
//
// @return BatchUpdateFileMetaResponse
func (client *Client) BatchUpdateFileMeta(request *BatchUpdateFileMetaRequest) (_result *BatchUpdateFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUpdateFileMetaResponse{}
	_body, _err := client.BatchUpdateFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Compares the similarity of the largest faces in two images. The largest face refers to the largest face frame in an image after face detection.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- For the input image, only the face with the largest face frame in the image is used for face comparison. The face frame detection result is consistent with the responses of the [DetectImageFaces](https://help.aliyun.com/document_detail/478213.html) operation.
//
// @param tmpReq - CompareImageFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareImageFacesResponse
func (client *Client) CompareImageFacesWithOptions(tmpReq *CompareImageFacesRequest, runtime *dara.RuntimeOptions) (_result *CompareImageFacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CompareImageFacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Source) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, dara.String("Source"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceShrink) {
		query["Source"] = request.SourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareImageFaces"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Compares the similarity of the largest faces in two images. The largest face refers to the largest face frame in an image after face detection.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- For the input image, only the face with the largest face frame in the image is used for face comparison. The face frame detection result is consistent with the responses of the [DetectImageFaces](https://help.aliyun.com/document_detail/478213.html) operation.
//
// @param request - CompareImageFacesRequest
//
// @return CompareImageFacesResponse
func (client *Client) CompareImageFaces(request *CompareImageFacesRequest) (_result *CompareImageFacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.CompareImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Phase II of AI Assistant, Q\\&A API
//
// Description:
//
// ### Precautions
//
// - Before using this interface, please make sure you fully understand the billing method and [pricing](https://help.aliyun.com/zh/imm/product-overview/billable-items?spm=openapi-amp.newDocPublishment.0.0.1ecd281fi27Zgk) of the Intelligent Media Management product.
//
// - Before calling this interface, ensure that you have indexed the files into the dataset (Dataset) through binding ([CreateBinding](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-createbinding?spm=a2c4g.11186623.0.0.a3d76f44xJrOnF)) or active indexing ([IndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-indexfilemeta?spm=a2c4g.11186623.help-menu-search-62354.d_0) or [BatchIndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-batchindexfilemeta?spm=a2c4g.11186623.help-menu-62354.d_5_2_4_2_1_1.f1d86f44iBs3QZ)).
//
// - The returned result is only an example. Depending on the [workflow template configuration](https://help.aliyun.com/zh/imm/user-guide/workflow-templates-and-operators?spm=a2c4g.11186623.0.0.a3d775abr3hDFp), the categories and content of the file metadata information obtained may differ from the example. If you have any questions, please join the DingTalk group by searching for the group number 21714099 in DingTalk.
//
// ### Usage Restrictions
//
// - The maximum length of the historical conversation is 100, including both user and assistant messages.
//
// - Each message should not exceed 1000 Chinese characters.
//
// @param tmpReq - ContextualAnswerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContextualAnswerResponse
func (client *Client) ContextualAnswerWithSSE(tmpReq *ContextualAnswerRequest, runtime *dara.RuntimeOptions, _yield chan *ContextualAnswerResponse, _yieldErr chan error) {
	defer close(_yield)
	client.contextualAnswerWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
	return
}

// Summary:
//
// Phase II of AI Assistant, Q\\&A API
//
// Description:
//
// ### Precautions
//
// - Before using this interface, please make sure you fully understand the billing method and [pricing](https://help.aliyun.com/zh/imm/product-overview/billable-items?spm=openapi-amp.newDocPublishment.0.0.1ecd281fi27Zgk) of the Intelligent Media Management product.
//
// - Before calling this interface, ensure that you have indexed the files into the dataset (Dataset) through binding ([CreateBinding](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-createbinding?spm=a2c4g.11186623.0.0.a3d76f44xJrOnF)) or active indexing ([IndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-indexfilemeta?spm=a2c4g.11186623.help-menu-search-62354.d_0) or [BatchIndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-batchindexfilemeta?spm=a2c4g.11186623.help-menu-62354.d_5_2_4_2_1_1.f1d86f44iBs3QZ)).
//
// - The returned result is only an example. Depending on the [workflow template configuration](https://help.aliyun.com/zh/imm/user-guide/workflow-templates-and-operators?spm=a2c4g.11186623.0.0.a3d775abr3hDFp), the categories and content of the file metadata information obtained may differ from the example. If you have any questions, please join the DingTalk group by searching for the group number 21714099 in DingTalk.
//
// ### Usage Restrictions
//
// - The maximum length of the historical conversation is 100, including both user and assistant messages.
//
// - Each message should not exceed 1000 Chinese characters.
//
// @param tmpReq - ContextualAnswerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContextualAnswerResponse
func (client *Client) ContextualAnswerWithOptions(tmpReq *ContextualAnswerRequest, runtime *dara.RuntimeOptions) (_result *ContextualAnswerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ContextualAnswerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("Messages"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilesShrink) {
		body["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.MessagesShrink) {
		body["Messages"] = request.MessagesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContextualAnswer"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContextualAnswerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Phase II of AI Assistant, Q\\&A API
//
// Description:
//
// ### Precautions
//
// - Before using this interface, please make sure you fully understand the billing method and [pricing](https://help.aliyun.com/zh/imm/product-overview/billable-items?spm=openapi-amp.newDocPublishment.0.0.1ecd281fi27Zgk) of the Intelligent Media Management product.
//
// - Before calling this interface, ensure that you have indexed the files into the dataset (Dataset) through binding ([CreateBinding](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-createbinding?spm=a2c4g.11186623.0.0.a3d76f44xJrOnF)) or active indexing ([IndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-indexfilemeta?spm=a2c4g.11186623.help-menu-search-62354.d_0) or [BatchIndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-batchindexfilemeta?spm=a2c4g.11186623.help-menu-62354.d_5_2_4_2_1_1.f1d86f44iBs3QZ)).
//
// - The returned result is only an example. Depending on the [workflow template configuration](https://help.aliyun.com/zh/imm/user-guide/workflow-templates-and-operators?spm=a2c4g.11186623.0.0.a3d775abr3hDFp), the categories and content of the file metadata information obtained may differ from the example. If you have any questions, please join the DingTalk group by searching for the group number 21714099 in DingTalk.
//
// ### Usage Restrictions
//
// - The maximum length of the historical conversation is 100, including both user and assistant messages.
//
// - Each message should not exceed 1000 Chinese characters.
//
// @param request - ContextualAnswerRequest
//
// @return ContextualAnswerResponse
func (client *Client) ContextualAnswer(request *ContextualAnswerRequest) (_result *ContextualAnswerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContextualAnswerResponse{}
	_body, _err := client.ContextualAnswerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves semantically similar documents. The operation is designed for multi-turn conversations and can process message input in historical conversations. The operation returns results that are highly related to the current conversation based on an in-depth understanding of contextual content. It provides consistent and efficient information retrieval in multi-turn conversations.
//
// Description:
//
// ### [](#)Usage notes
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/zh/imm/product-overview/billable-items?spm=openapi-amp.newDocPublishment.0.0.1ecd281fi27Zgk) of Intelligent Media Management (IMM).
//
//   - Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-createbinding?spm=a2c4g.11186623.0.0.a3d76f44xJrOnF) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-indexfilemeta?spm=a2c4g.11186623.help-menu-search-62354.d_0) or [BatchIndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-batchindexfilemeta?spm=a2c4g.11186623.help-menu-62354.d_5_2_4_2_1_1.f1d86f44iBs3QZ) operation.
//
//   - The response provided in this example is for reference only. The categories and content of metadata vary based on configurations of [workflow templates](https://help.aliyun.com/zh/imm/user-guide/workflow-templates-and-operators?spm=a2c4g.11186623.0.0.a3d775abr3hDFp). For any inquiries, join the DingTalk chat group (ID: 21714099) for feedback.
//
// ### [](#)Limitations
//
//   - The conversation history can hold up to 100 messages, including user-sent messages and assistant-generated messages.
//
//   - Each message cannot exceed 1,000 characters in length.
//
// @param tmpReq - ContextualRetrievalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContextualRetrievalResponse
func (client *Client) ContextualRetrievalWithOptions(tmpReq *ContextualRetrievalRequest, runtime *dara.RuntimeOptions) (_result *ContextualRetrievalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ContextualRetrievalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("Messages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SmartClusterIds) {
		request.SmartClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SmartClusterIds, dara.String("SmartClusterIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RecallOnly) {
		query["RecallOnly"] = request.RecallOnly
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MessagesShrink) {
		body["Messages"] = request.MessagesShrink
	}

	if !dara.IsNil(request.SmartClusterIdsShrink) {
		body["SmartClusterIds"] = request.SmartClusterIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContextualRetrieval"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContextualRetrievalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves semantically similar documents. The operation is designed for multi-turn conversations and can process message input in historical conversations. The operation returns results that are highly related to the current conversation based on an in-depth understanding of contextual content. It provides consistent and efficient information retrieval in multi-turn conversations.
//
// Description:
//
// ### [](#)Usage notes
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/zh/imm/product-overview/billable-items?spm=openapi-amp.newDocPublishment.0.0.1ecd281fi27Zgk) of Intelligent Media Management (IMM).
//
//   - Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-createbinding?spm=a2c4g.11186623.0.0.a3d76f44xJrOnF) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-indexfilemeta?spm=a2c4g.11186623.help-menu-search-62354.d_0) or [BatchIndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-batchindexfilemeta?spm=a2c4g.11186623.help-menu-62354.d_5_2_4_2_1_1.f1d86f44iBs3QZ) operation.
//
//   - The response provided in this example is for reference only. The categories and content of metadata vary based on configurations of [workflow templates](https://help.aliyun.com/zh/imm/user-guide/workflow-templates-and-operators?spm=a2c4g.11186623.0.0.a3d775abr3hDFp). For any inquiries, join the DingTalk chat group (ID: 21714099) for feedback.
//
// ### [](#)Limitations
//
//   - The conversation history can hold up to 100 messages, including user-sent messages and assistant-generated messages.
//
//   - Each message cannot exceed 1,000 characters in length.
//
// @param request - ContextualRetrievalRequest
//
// @return ContextualRetrievalResponse
func (client *Client) ContextualRetrieval(request *ContextualRetrievalRequest) (_result *ContextualRetrievalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContextualRetrievalResponse{}
	_body, _err := client.ContextualRetrievalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an archive file inspection task to preview the files in a package without decompressing the package.
//
// Description:
//
// >  The operation is in public preview. For any inquires, join our DingTalk chat group (ID: 31690030817) and share your questions with us.
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//     **
//
//     **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//   - The operation supports a package that contains up to 80,000 files.
//
//   - The operation supports ZIP or RAR packages up to 200 GB in size, or 7z packages up to 50 GB in size.
//
//   - This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.“ If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateArchiveFileInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArchiveFileInspectionTaskResponse
func (client *Client) CreateArchiveFileInspectionTaskWithOptions(tmpReq *CreateArchiveFileInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateArchiveFileInspectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateArchiveFileInspectionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateArchiveFileInspectionTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateArchiveFileInspectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an archive file inspection task to preview the files in a package without decompressing the package.
//
// Description:
//
// >  The operation is in public preview. For any inquires, join our DingTalk chat group (ID: 31690030817) and share your questions with us.
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//     **
//
//     **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//   - The operation supports a package that contains up to 80,000 files.
//
//   - The operation supports ZIP or RAR packages up to 200 GB in size, or 7z packages up to 50 GB in size.
//
//   - This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.“ If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateArchiveFileInspectionTaskRequest
//
// @return CreateArchiveFileInspectionTaskResponse
func (client *Client) CreateArchiveFileInspectionTask(request *CreateArchiveFileInspectionTaskRequest) (_result *CreateArchiveFileInspectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateArchiveFileInspectionTaskResponse{}
	_body, _err := client.CreateArchiveFileInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a batch processing task to perform a data processing operation, such as transcoding or format conversion, on multiple existing files at a time.
//
// Description:
//
// If you want to create a batch processing task to process data in [OSS](https://help.aliyun.com/document_detail/99372.html), make sure that you have bound the dataset to the OSS bucket where the data is stored. For more information about how to bind a dataset to a bucket, see [AttachOSSBucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param tmpReq - CreateBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchResponse
func (client *Client) CreateBatchWithOptions(tmpReq *CreateBatchRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Actions) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, dara.String("Actions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionsShrink) {
		body["Actions"] = request.ActionsShrink
	}

	if !dara.IsNil(request.InputShrink) {
		body["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		body["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ServiceRole) {
		body["ServiceRole"] = request.ServiceRole
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBatch"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a batch processing task to perform a data processing operation, such as transcoding or format conversion, on multiple existing files at a time.
//
// Description:
//
// If you want to create a batch processing task to process data in [OSS](https://help.aliyun.com/document_detail/99372.html), make sure that you have bound the dataset to the OSS bucket where the data is stored. For more information about how to bind a dataset to a bucket, see [AttachOSSBucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param request - CreateBatchRequest
//
// @return CreateBatchResponse
func (client *Client) CreateBatch(request *CreateBatchRequest) (_result *CreateBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBatchResponse{}
	_body, _err := client.CreateBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a binding relationship between a dataset and an Object Storage Service (OSS) bucket. This allows for the automatic synchronization of incremental and full data and indexing.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/2743997.html) of Intelligent Media Management (IMM).****
//
// >  Asynchronous processing does not guarantee timely task completion.
//
// Before you create a binding, make sure that the project and the dataset that you want to use exist.
//
//   - For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//   - For information about how to create a dataset, see [CreateDataset](https://help.aliyun.com/document_detail/478160.html).
//
// >  The CreateBinding operation works by using the [workflow template](https://help.aliyun.com/document_detail/466304.html) that is specified when you created the project or dataset.
//
// After you create a binding between a dataset and an OSS bucket, IMM scans the existing objects in the bucket and extracts metadata based on the scanning result. Then, IMM creates an index from the extracted metadata. If new objects are uploaded to the OSS bucket, IMM tracks and scans the objects and updates the index. For objects whose metadata index is created by calling this operation, you can call query operations, such as [SimpleQuery](https://help.aliyun.com/document_detail/478175.html), to query objects, manage objects, and collect statistics on objects.
//
// @param request - CreateBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBindingResponse
func (client *Client) CreateBindingWithOptions(request *CreateBindingRequest, runtime *dara.RuntimeOptions) (_result *CreateBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBinding"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a binding relationship between a dataset and an Object Storage Service (OSS) bucket. This allows for the automatic synchronization of incremental and full data and indexing.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/2743997.html) of Intelligent Media Management (IMM).****
//
// >  Asynchronous processing does not guarantee timely task completion.
//
// Before you create a binding, make sure that the project and the dataset that you want to use exist.
//
//   - For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//   - For information about how to create a dataset, see [CreateDataset](https://help.aliyun.com/document_detail/478160.html).
//
// >  The CreateBinding operation works by using the [workflow template](https://help.aliyun.com/document_detail/466304.html) that is specified when you created the project or dataset.
//
// After you create a binding between a dataset and an OSS bucket, IMM scans the existing objects in the bucket and extracts metadata based on the scanning result. Then, IMM creates an index from the extracted metadata. If new objects are uploaded to the OSS bucket, IMM tracks and scans the objects and updates the index. For objects whose metadata index is created by calling this operation, you can call query operations, such as [SimpleQuery](https://help.aliyun.com/document_detail/478175.html), to query objects, manage objects, and collect statistics on objects.
//
// @param request - CreateBindingRequest
//
// @return CreateBindingResponse
func (client *Client) CreateBinding(request *CreateBindingRequest) (_result *CreateBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBindingResponse{}
	_body, _err := client.CreateBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Compresses point cloud data (PCD) in Object Storage Service (OSS) to reduce the amount of data transferred over networks.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- This operation supports only Point Cloud Data (PCD) files.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications. >
//
// @param tmpReq - CreateCompressPointCloudTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCompressPointCloudTaskResponse
func (client *Client) CreateCompressPointCloudTaskWithOptions(tmpReq *CreateCompressPointCloudTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCompressPointCloudTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCompressPointCloudTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KdtreeOption) {
		request.KdtreeOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KdtreeOption, dara.String("KdtreeOption"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OctreeOption) {
		request.OctreeOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OctreeOption, dara.String("OctreeOption"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PointCloudFields) {
		request.PointCloudFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PointCloudFields, dara.String("PointCloudFields"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CompressMethod) {
		query["CompressMethod"] = request.CompressMethod
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.KdtreeOptionShrink) {
		query["KdtreeOption"] = request.KdtreeOptionShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.OctreeOptionShrink) {
		query["OctreeOption"] = request.OctreeOptionShrink
	}

	if !dara.IsNil(request.PointCloudFieldsShrink) {
		query["PointCloudFields"] = request.PointCloudFieldsShrink
	}

	if !dara.IsNil(request.PointCloudFileFormat) {
		query["PointCloudFileFormat"] = request.PointCloudFileFormat
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCompressPointCloudTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCompressPointCloudTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Compresses point cloud data (PCD) in Object Storage Service (OSS) to reduce the amount of data transferred over networks.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- This operation supports only Point Cloud Data (PCD) files.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications. >
//
// @param request - CreateCompressPointCloudTaskRequest
//
// @return CreateCompressPointCloudTaskResponse
func (client *Client) CreateCompressPointCloudTask(request *CreateCompressPointCloudTaskRequest) (_result *CreateCompressPointCloudTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCompressPointCloudTaskResponse{}
	_body, _err := client.CreateCompressPointCloudTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a story based on the specified images and videos.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
// @param tmpReq - CreateCustomizedStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomizedStoryResponse
func (client *Client) CreateCustomizedStoryWithOptions(tmpReq *CreateCustomizedStoryRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomizedStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCustomizedStoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Cover) {
		request.CoverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cover, dara.String("Cover"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomLabels) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, dara.String("CustomLabels"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CoverShrink) {
		body["Cover"] = request.CoverShrink
	}

	if !dara.IsNil(request.CustomLabelsShrink) {
		body["CustomLabels"] = request.CustomLabelsShrink
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FilesShrink) {
		body["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.StoryName) {
		body["StoryName"] = request.StoryName
	}

	if !dara.IsNil(request.StorySubType) {
		body["StorySubType"] = request.StorySubType
	}

	if !dara.IsNil(request.StoryType) {
		body["StoryType"] = request.StoryType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomizedStory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomizedStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a story based on the specified images and videos.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
// @param request - CreateCustomizedStoryRequest
//
// @return CreateCustomizedStoryResponse
func (client *Client) CreateCustomizedStory(request *CreateCustomizedStoryRequest) (_result *CreateCustomizedStoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomizedStoryResponse{}
	_body, _err := client.CreateCustomizedStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Dataset
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - The dataset name must be unique within the same project.
//
// - There is a limit to the number of datasets that can be created, which can be queried through [GetProject](https://help.aliyun.com/document_detail/478155.html).
//
// - After creating a dataset, you can use [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) to build file metadata indexes, enabling diversified [data retrieval and statistics](https://help.aliyun.com/document_detail/478175.html) and intelligent management.
//
// @param tmpReq - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithOptions(tmpReq *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkflowParameters) {
		request.WorkflowParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkflowParameters, dara.String("WorkflowParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetMaxBindCount) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !dara.IsNil(request.DatasetMaxEntityCount) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !dara.IsNil(request.DatasetMaxFileCount) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !dara.IsNil(request.DatasetMaxRelationCount) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !dara.IsNil(request.DatasetMaxTotalFileSize) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.WorkflowParametersShrink) {
		query["WorkflowParameters"] = request.WorkflowParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Dataset
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - The dataset name must be unique within the same project.
//
// - There is a limit to the number of datasets that can be created, which can be queried through [GetProject](https://help.aliyun.com/document_detail/478155.html).
//
// - After creating a dataset, you can use [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) to build file metadata indexes, enabling diversified [data retrieval and statistics](https://help.aliyun.com/document_detail/478175.html) and intelligent management.
//
// @param request - CreateDatasetRequest
//
// @return CreateDatasetResponse
func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Decodes the blind watermark in an image.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of Intelligent Media Management (IMM).
//
//	    **
//
//	    **Note that*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Make sure that an IMM project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- The region and project specified in the request to decode a blind watermark must match those in the [EncodeBlindWatermark](https://help.aliyun.com/document_detail/2743655.html) request to encode the blind watermark.
//
//		- A blind watermark can still be extracted even if attacks, such as compression, scaling, cropping, and color transformation, are performed on the image.
//
//		- This operation is compatible with its earlier version DecodeBlindWatermark.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task. If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateDecodeBlindWatermarkTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDecodeBlindWatermarkTaskResponse
func (client *Client) CreateDecodeBlindWatermarkTaskWithOptions(tmpReq *CreateDecodeBlindWatermarkTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDecodeBlindWatermarkTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDecodeBlindWatermarkTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageQuality) {
		query["ImageQuality"] = request.ImageQuality
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.OriginalImageURI) {
		query["OriginalImageURI"] = request.OriginalImageURI
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.StrengthLevel) {
		query["StrengthLevel"] = request.StrengthLevel
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.WatermarkType) {
		query["WatermarkType"] = request.WatermarkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDecodeBlindWatermarkTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDecodeBlindWatermarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decodes the blind watermark in an image.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of Intelligent Media Management (IMM).
//
//	    **
//
//	    **Note that*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Make sure that an IMM project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- The region and project specified in the request to decode a blind watermark must match those in the [EncodeBlindWatermark](https://help.aliyun.com/document_detail/2743655.html) request to encode the blind watermark.
//
//		- A blind watermark can still be extracted even if attacks, such as compression, scaling, cropping, and color transformation, are performed on the image.
//
//		- This operation is compatible with its earlier version DecodeBlindWatermark.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task. If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateDecodeBlindWatermarkTaskRequest
//
// @return CreateDecodeBlindWatermarkTaskResponse
func (client *Client) CreateDecodeBlindWatermarkTask(request *CreateDecodeBlindWatermarkTaskRequest) (_result *CreateDecodeBlindWatermarkTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDecodeBlindWatermarkTaskResponse{}
	_body, _err := client.CreateDecodeBlindWatermarkTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches the dataset for the specified number of images most similar to the specified image or face and returns face IDs and boundaries in descending order of similarity.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- The operation searches for faces that are similar to the face within the largest bounding box in each input image.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateFacesSearchingTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFacesSearchingTaskResponse
func (client *Client) CreateFacesSearchingTaskWithOptions(tmpReq *CreateFacesSearchingTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFacesSearchingTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFacesSearchingTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxResult) {
		query["MaxResult"] = request.MaxResult
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFacesSearchingTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFacesSearchingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches the dataset for the specified number of images most similar to the specified image or face and returns face IDs and boundaries in descending order of similarity.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- The operation searches for faces that are similar to the face within the largest bounding box in each input image.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateFacesSearchingTaskRequest
//
// @return CreateFacesSearchingTaskResponse
func (client *Client) CreateFacesSearchingTask(request *CreateFacesSearchingTaskRequest) (_result *CreateFacesSearchingTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFacesSearchingTaskResponse{}
	_body, _err := client.CreateFacesSearchingTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a face clustering task to cluster faces of different persons in images by person based on the intelligent algorithms.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Before you call this operation, make sure that you have indexed file metadata into the [dataset](~~CreateDataset~~) automatically by calling the [CreateBinding](~~CreateBinding~~) operation or manually by calling the [IndexFileMeta](~~IndexFileMeta~~) or [BatchIndexFileMeta](~~BatchIndexFileMeta~~) operation.
//
//		- Each call to the operation incrementally processes metadata in the [dataset](~~CreateDataset~~). You can regularly call this operation to process incremental files.
//
//		- After the clustering task is complete, you can call the [GetFigureCluster](~~GetFigureCluster~~) or [BatchGetFigureCluster](~~BatchGetFigureCluster~~) operation to query information about a specific cluster. You can also call the [QueryFigureClusters](~~QueryFigureClusters~~) operation to query all face clusters of the specified dataset.
//
//		- Removing image information from the dataset causes changes to face clusters. When images that contain all faces in a cluster are removed, the cluster is deleted.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can call the [GetTask](~~GetTask~~) or [ListTasks](~~ListTasks~~) operation to query information about the task. If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateFigureClusteringTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFigureClusteringTaskResponse
func (client *Client) CreateFigureClusteringTaskWithOptions(tmpReq *CreateFigureClusteringTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFigureClusteringTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFigureClusteringTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFigureClusteringTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFigureClusteringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a face clustering task to cluster faces of different persons in images by person based on the intelligent algorithms.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Before you call this operation, make sure that you have indexed file metadata into the [dataset](~~CreateDataset~~) automatically by calling the [CreateBinding](~~CreateBinding~~) operation or manually by calling the [IndexFileMeta](~~IndexFileMeta~~) or [BatchIndexFileMeta](~~BatchIndexFileMeta~~) operation.
//
//		- Each call to the operation incrementally processes metadata in the [dataset](~~CreateDataset~~). You can regularly call this operation to process incremental files.
//
//		- After the clustering task is complete, you can call the [GetFigureCluster](~~GetFigureCluster~~) or [BatchGetFigureCluster](~~BatchGetFigureCluster~~) operation to query information about a specific cluster. You can also call the [QueryFigureClusters](~~QueryFigureClusters~~) operation to query all face clusters of the specified dataset.
//
//		- Removing image information from the dataset causes changes to face clusters. When images that contain all faces in a cluster are removed, the cluster is deleted.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can call the [GetTask](~~GetTask~~) or [ListTasks](~~ListTasks~~) operation to query information about the task. If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateFigureClusteringTaskRequest
//
// @return CreateFigureClusteringTaskResponse
func (client *Client) CreateFigureClusteringTask(request *CreateFigureClusteringTaskRequest) (_result *CreateFigureClusteringTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFigureClusteringTaskResponse{}
	_body, _err := client.CreateFigureClusteringTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Merges two or more face clustering groups into one face clustering group.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have called the [CreateFigureClusteringTask](https://help.aliyun.com/document_detail/478180.html) operation to cluster all faces in the dataset.
//
//		- If you merge unrelated groups, the feature values of the target groups are affected. As a result, the incremental data may be inaccurately grouped when you create a face clustering task.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateFigureClustersMergingTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFigureClustersMergingTaskResponse
func (client *Client) CreateFigureClustersMergingTaskWithOptions(tmpReq *CreateFigureClustersMergingTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFigureClustersMergingTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFigureClustersMergingTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Froms) {
		request.FromsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Froms, dara.String("Froms"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.FromsShrink) {
		query["Froms"] = request.FromsShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFigureClustersMergingTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFigureClustersMergingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Merges two or more face clustering groups into one face clustering group.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have called the [CreateFigureClusteringTask](https://help.aliyun.com/document_detail/478180.html) operation to cluster all faces in the dataset.
//
//		- If you merge unrelated groups, the feature values of the target groups are affected. As a result, the incremental data may be inaccurately grouped when you create a face clustering task.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateFigureClustersMergingTaskRequest
//
// @return CreateFigureClustersMergingTaskResponse
func (client *Client) CreateFigureClustersMergingTask(request *CreateFigureClustersMergingTaskRequest) (_result *CreateFigureClustersMergingTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFigureClustersMergingTaskResponse{}
	_body, _err := client.CreateFigureClustersMergingTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a file packing task.
//
// Description:
//
// >  The operation is in public preview. For any inquires, join our DingTalk group (ID: 88490020073) and share your questions with us.
//
// >  The operation supports file packing only. Compression support will be added later.
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//     **
//
//     **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//   - A call to the operation can pack up to 80,000 objects into a package.
//
//   - The total size of all objects to be packed into a package cannot exceed 200 GB.
//
//   - The operation can pack only Standard objects in Object Storage Service (OSS). To pack an object in another storage class, you must first [convert the storage class of the object](https://help.aliyun.com/document_detail/90090.html).
//
//   - This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.“ If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateFileCompressionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileCompressionTaskResponse
func (client *Client) CreateFileCompressionTaskWithOptions(tmpReq *CreateFileCompressionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFileCompressionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFileCompressionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CompressedFormat) {
		query["CompressedFormat"] = request.CompressedFormat
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceManifestURI) {
		query["SourceManifestURI"] = request.SourceManifestURI
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileCompressionTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileCompressionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file packing task.
//
// Description:
//
// >  The operation is in public preview. For any inquires, join our DingTalk group (ID: 88490020073) and share your questions with us.
//
// >  The operation supports file packing only. Compression support will be added later.
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//     **
//
//     **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//   - A call to the operation can pack up to 80,000 objects into a package.
//
//   - The total size of all objects to be packed into a package cannot exceed 200 GB.
//
//   - The operation can pack only Standard objects in Object Storage Service (OSS). To pack an object in another storage class, you must first [convert the storage class of the object](https://help.aliyun.com/document_detail/90090.html).
//
//   - This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.“ If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateFileCompressionTaskRequest
//
// @return CreateFileCompressionTaskResponse
func (client *Client) CreateFileCompressionTask(request *CreateFileCompressionTaskRequest) (_result *CreateFileCompressionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileCompressionTaskResponse{}
	_body, _err := client.CreateFileCompressionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Extracts the specified files from a ZIP, RAR, or 7z package to the specified directory or decompresses the entire package.
//
// Description:
//
// >  The operation is in public preview. For any inquires, join our DingTalk group (ID: 88490020073) and share your questions with us.
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//     **
//
//     **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//   - The operation supports a package that contains up to 80,000 files.
//
//   - The operation supports ZIP or RAR packages up to 200 GB in size, or 7z packages up to 50 GB in size.
//
//   - The operation extracts files in streams to the specified directory. If the file extraction task is interrupted by a corrupt file, files that have been extracted are not deleted.
//
//   - This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.“ If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateFileUncompressionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileUncompressionTaskResponse
func (client *Client) CreateFileUncompressionTaskWithOptions(tmpReq *CreateFileUncompressionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFileUncompressionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFileUncompressionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SelectedFiles) {
		request.SelectedFilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedFiles, dara.String("SelectedFiles"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SelectedFilesShrink) {
		query["SelectedFiles"] = request.SelectedFilesShrink
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileUncompressionTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileUncompressionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extracts the specified files from a ZIP, RAR, or 7z package to the specified directory or decompresses the entire package.
//
// Description:
//
// >  The operation is in public preview. For any inquires, join our DingTalk group (ID: 88490020073) and share your questions with us.
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//     **
//
//     **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//   - The operation supports a package that contains up to 80,000 files.
//
//   - The operation supports ZIP or RAR packages up to 200 GB in size, or 7z packages up to 50 GB in size.
//
//   - The operation extracts files in streams to the specified directory. If the file extraction task is interrupted by a corrupt file, files that have been extracted are not deleted.
//
//   - This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.“ If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateFileUncompressionTaskRequest
//
// @return CreateFileUncompressionTaskResponse
func (client *Client) CreateFileUncompressionTask(request *CreateFileUncompressionTaskRequest) (_result *CreateFileUncompressionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileUncompressionTaskResponse{}
	_body, _err := client.CreateFileUncompressionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image moderation task to ensure image content compliance. You can call this operation to identify inappropriate content, such as pornography, violence, terrorism, politically sensitive content, undesirable scenes, unauthorized logos, and non-compliant ads.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- The image for which you want to create a content moderation task must meet the following requirements:
//
//	    	- The image URL uses the HTTP or HTTPS protocol.
//
//	    	- The image is in one of the following formats: PNG, JPG, JPEG, BMP, GIF, and WebP
//
//	    	- The image size is limited to 20 MB for synchronous and asynchronous calls, with a maximum height or width of 30,000 pixels. The total number of pixels in the image cannot exceed 250 million. GIF images are limited to 4,194,304 pixels, with a maximum height or width of 30,000 pixels.
//
//	    	- The image download time is limited to 3 seconds. If the download takes longer, a timeout error occurs.
//
//	    	- To ensure effective moderation, we recommend that you submit an image with dimensions of at least 256 × 256 pixels.
//
//	    	- The response time of the CreateImageModerationTask operation varies based on the duration of the image download. Make sure that the image is stored in a stable and reliable service. We recommend that you store images on Alibaba Cloud Object Storage Service (OSS) or cache them on Alibaba Cloud CDN.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can also obtain information about the task based on notifications.
//
// >  The detection result is sent as an asynchronous notification. The Suggestion field of the notification can have one of the following values:
//
//   - pass: No non-compliant content is found.
//
//   - block: Non-compliant content is detected. The Categories field value indicates the non-compliance categories. For more information, see Content moderation results.
//
//   - review: A manual review is needed. After the manual review is completed, an asynchronous notification is sent to inform you of the result. >
//
// @param tmpReq - CreateImageModerationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageModerationTaskResponse
func (client *Client) CreateImageModerationTaskWithOptions(tmpReq *CreateImageModerationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageModerationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateImageModerationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Scenes) {
		request.ScenesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scenes, dara.String("Scenes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.MaxFrames) {
		query["MaxFrames"] = request.MaxFrames
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ScenesShrink) {
		query["Scenes"] = request.ScenesShrink
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageModerationTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageModerationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image moderation task to ensure image content compliance. You can call this operation to identify inappropriate content, such as pornography, violence, terrorism, politically sensitive content, undesirable scenes, unauthorized logos, and non-compliant ads.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- The image for which you want to create a content moderation task must meet the following requirements:
//
//	    	- The image URL uses the HTTP or HTTPS protocol.
//
//	    	- The image is in one of the following formats: PNG, JPG, JPEG, BMP, GIF, and WebP
//
//	    	- The image size is limited to 20 MB for synchronous and asynchronous calls, with a maximum height or width of 30,000 pixels. The total number of pixels in the image cannot exceed 250 million. GIF images are limited to 4,194,304 pixels, with a maximum height or width of 30,000 pixels.
//
//	    	- The image download time is limited to 3 seconds. If the download takes longer, a timeout error occurs.
//
//	    	- To ensure effective moderation, we recommend that you submit an image with dimensions of at least 256 × 256 pixels.
//
//	    	- The response time of the CreateImageModerationTask operation varies based on the duration of the image download. Make sure that the image is stored in a stable and reliable service. We recommend that you store images on Alibaba Cloud Object Storage Service (OSS) or cache them on Alibaba Cloud CDN.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can also obtain information about the task based on notifications.
//
// >  The detection result is sent as an asynchronous notification. The Suggestion field of the notification can have one of the following values:
//
//   - pass: No non-compliant content is found.
//
//   - block: Non-compliant content is detected. The Categories field value indicates the non-compliance categories. For more information, see Content moderation results.
//
//   - review: A manual review is needed. After the manual review is completed, an asynchronous notification is sent to inform you of the result. >
//
// @param request - CreateImageModerationTaskRequest
//
// @return CreateImageModerationTaskResponse
func (client *Client) CreateImageModerationTask(request *CreateImageModerationTaskRequest) (_result *CreateImageModerationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageModerationTaskResponse{}
	_body, _err := client.CreateImageModerationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image splicing task. You can call this operation to splice multiple images into one based on a given rule and save the final image into an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
//		- You can call this operation to merge up to 10 images. Each side of an image cannot exceed 32,876 pixels, and the total number of pixels of the image cannot exceed 1 billion.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateImageSplicingTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageSplicingTaskResponse
func (client *Client) CreateImageSplicingTaskWithOptions(tmpReq *CreateImageSplicingTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageSplicingTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateImageSplicingTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Align) {
		query["Align"] = request.Align
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.ImageFormat) {
		query["ImageFormat"] = request.ImageFormat
	}

	if !dara.IsNil(request.Margin) {
		query["Margin"] = request.Margin
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Padding) {
		query["Padding"] = request.Padding
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Quality) {
		query["Quality"] = request.Quality
	}

	if !dara.IsNil(request.ScaleType) {
		query["ScaleType"] = request.ScaleType
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageSplicingTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageSplicingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image splicing task. You can call this operation to splice multiple images into one based on a given rule and save the final image into an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
//		- You can call this operation to merge up to 10 images. Each side of an image cannot exceed 32,876 pixels, and the total number of pixels of the image cannot exceed 1 billion.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateImageSplicingTaskRequest
//
// @return CreateImageSplicingTaskResponse
func (client *Client) CreateImageSplicingTask(request *CreateImageSplicingTaskRequest) (_result *CreateImageSplicingTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageSplicingTaskResponse{}
	_body, _err := client.CreateImageSplicingTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts multiple images into one single PDF file and stores the PDF file to the specified path in Object Storage Service (OSS).
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
//		- You can specify up to 100 images in a call to the operation.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateImageToPDFTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageToPDFTaskResponse
func (client *Client) CreateImageToPDFTaskWithOptions(tmpReq *CreateImageToPDFTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageToPDFTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateImageToPDFTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageToPDFTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageToPDFTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts multiple images into one single PDF file and stores the PDF file to the specified path in Object Storage Service (OSS).
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
//		- You can specify up to 100 images in a call to the operation.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateImageToPDFTaskRequest
//
// @return CreateImageToPDFTaskResponse
func (client *Client) CreateImageToPDFTask(request *CreateImageToPDFTaskRequest) (_result *CreateImageToPDFTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageToPDFTaskResponse{}
	_body, _err := client.CreateImageToPDFTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a spatiotemporal clustering task to cluster photos and videos based on geolocation and time information. Spatiotemporal clustering allows you to group photos and videos taken during a travel or at different places by their spatial and temporal similarity. Based on spatiotemporal clustering, you can develop media capabilities such as media file categorization, photo collections, and image and video-based stories.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Each call to the operation incrementally processes metadata in the dataset.****`` You can regularly call this operation to process incremental files.
//
//		- After a spatiotemporal clustering task is complete, you can call the [QueryLocationDateClusters](https://help.aliyun.com/document_detail/478189.html) operation to query the spatiotemporal clustering result.
//
//		- Removing metadata from a dataset does not affect existing spatiotemporal clusters for the dataset. To delete a spatiotemporal cluster, call the [DeleteLocationDateCluster](https://help.aliyun.com/document_detail/478191.html) operation.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateLocationDateClusteringTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLocationDateClusteringTaskResponse
func (client *Client) CreateLocationDateClusteringTaskWithOptions(tmpReq *CreateLocationDateClusteringTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateLocationDateClusteringTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLocationDateClusteringTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DateOptions) {
		request.DateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DateOptions, dara.String("DateOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LocationOptions) {
		request.LocationOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationOptions, dara.String("LocationOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DateOptionsShrink) {
		query["DateOptions"] = request.DateOptionsShrink
	}

	if !dara.IsNil(request.LocationOptionsShrink) {
		query["LocationOptions"] = request.LocationOptionsShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLocationDateClusteringTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLocationDateClusteringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a spatiotemporal clustering task to cluster photos and videos based on geolocation and time information. Spatiotemporal clustering allows you to group photos and videos taken during a travel or at different places by their spatial and temporal similarity. Based on spatiotemporal clustering, you can develop media capabilities such as media file categorization, photo collections, and image and video-based stories.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Each call to the operation incrementally processes metadata in the dataset.****`` You can regularly call this operation to process incremental files.
//
//		- After a spatiotemporal clustering task is complete, you can call the [QueryLocationDateClusters](https://help.aliyun.com/document_detail/478189.html) operation to query the spatiotemporal clustering result.
//
//		- Removing metadata from a dataset does not affect existing spatiotemporal clusters for the dataset. To delete a spatiotemporal cluster, call the [DeleteLocationDateCluster](https://help.aliyun.com/document_detail/478191.html) operation.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateLocationDateClusteringTaskRequest
//
// @return CreateLocationDateClusteringTaskResponse
func (client *Client) CreateLocationDateClusteringTask(request *CreateLocationDateClusteringTaskRequest) (_result *CreateLocationDateClusteringTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLocationDateClusteringTaskResponse{}
	_body, _err := client.CreateLocationDateClusteringTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Transcoding Service
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/88317.html) of the Intelligent Media Management product before using this interface.**
//
// - Before calling this interface, make sure that there is an available project (Project) in the current Region. For more details, see [Project Management](https://help.aliyun.com/document_detail/478152.html).
//
//	Notice: Asynchronous tasks do not guarantee timeliness.
//
// - When using this interface for media transcoding, by default, only one video/audio/subtitle stream is processed, but you can also configure the number of video/audio/subtitle streams to be processed.
//
// - When using this interface for media concatenation, a maximum of 11 media files are supported. In this case, the configured transcoding, frame extraction, and other parameters will apply to the concatenated media data.
//
// - This is an asynchronous interface. After the task starts, the task information is retained for 7 days. If it exceeds 7 days, the information cannot be retrieved. Call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) interface to get the returned `TaskId` and view the task information. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) message notification parameter to obtain task information through message notifications.
//
// @param tmpReq - CreateMediaConvertTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMediaConvertTaskResponse
func (client *Client) CreateMediaConvertTaskWithOptions(tmpReq *CreateMediaConvertTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMediaConvertTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMediaConvertTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Targets) {
		request.TargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Targets, dara.String("Targets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlignmentIndex) {
		query["AlignmentIndex"] = request.AlignmentIndex
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TargetsShrink) {
		query["Targets"] = request.TargetsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMediaConvertTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMediaConvertTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Transcoding Service
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/88317.html) of the Intelligent Media Management product before using this interface.**
//
// - Before calling this interface, make sure that there is an available project (Project) in the current Region. For more details, see [Project Management](https://help.aliyun.com/document_detail/478152.html).
//
//	Notice: Asynchronous tasks do not guarantee timeliness.
//
// - When using this interface for media transcoding, by default, only one video/audio/subtitle stream is processed, but you can also configure the number of video/audio/subtitle streams to be processed.
//
// - When using this interface for media concatenation, a maximum of 11 media files are supported. In this case, the configured transcoding, frame extraction, and other parameters will apply to the concatenated media data.
//
// - This is an asynchronous interface. After the task starts, the task information is retained for 7 days. If it exceeds 7 days, the information cannot be retrieved. Call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) interface to get the returned `TaskId` and view the task information. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) message notification parameter to obtain task information through message notifications.
//
// @param request - CreateMediaConvertTaskRequest
//
// @return CreateMediaConvertTaskResponse
func (client *Client) CreateMediaConvertTask(request *CreateMediaConvertTaskRequest) (_result *CreateMediaConvertTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMediaConvertTaskResponse{}
	_body, _err := client.CreateMediaConvertTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a document format conversion task to convert the format of a document stored in an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Supported input formats:
//
//	    	- Text documents: doc, docx, wps, wpss, docm, dotm, dot, dotx, and html
//
//	    	- Presentation documents: pptx, ppt, pot, potx, pps, ppsx, dps, dpt, pptm, potm, ppsm, and dpss
//
//	    	- Spreadsheet documents: xls, xlt, et, ett, xlsx, xltx, csv, xlsb, xlsm, xltm, and ets
//
//	    	- PDF documents: pdf
//
//		- Supported output formats:
//
//	    	- Image files: png and jpg
//
//	    	- Text files: txt
//
//	    	- PDF files: pdf
//
//		- Each input document can be up to 200 MB in size. The upper limit cannot be adjusted.
//
//		- If the document size is large or the content is complex, the conversion task may time out.
//
//		- The limit on the number of requests per second for a single user is 50.
//
//		- The operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can use one of the following methods to query the task information in a timely manner:
//
//	    	- Call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.``
//
//	    	- In the region in which the IMM project is located, configure a Simple Message Queue (SMQ) subscription to receive task information notifications. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html). For information about SMQ SDKs, see [Use queues](https://help.aliyun.com/document_detail/32449.html).
//
//	    	- In the region in which the IMM project is located, create an ApsaraMQ for RocketMQ 4.0 instance, a topic, and a group to receive task notifications. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html). For more information about how to use ApsaraMQ for RocketMQ, see [Call HTTP SDKs to send and subscribe to messages](https://help.aliyun.com/document_detail/169009.html).
//
//	    	- In the region in which the IMM project is located, use [EventBridge](https://www.aliyun.com/product/aliware/eventbridge) to receive task information notifications. For more information, see [IMM events](https://help.aliyun.com/document_detail/205730.html).
//
// @param tmpReq - CreateOfficeConversionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOfficeConversionTaskResponse
func (client *Client) CreateOfficeConversionTaskWithOptions(tmpReq *CreateOfficeConversionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateOfficeConversionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOfficeConversionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TrimPolicy) {
		request.TrimPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrimPolicy, dara.String("TrimPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.EndPage) {
		query["EndPage"] = request.EndPage
	}

	if !dara.IsNil(request.FirstPage) {
		query["FirstPage"] = request.FirstPage
	}

	if !dara.IsNil(request.FitToHeight) {
		query["FitToHeight"] = request.FitToHeight
	}

	if !dara.IsNil(request.FitToWidth) {
		query["FitToWidth"] = request.FitToWidth
	}

	if !dara.IsNil(request.HoldLineFeed) {
		query["HoldLineFeed"] = request.HoldLineFeed
	}

	if !dara.IsNil(request.ImageDPI) {
		query["ImageDPI"] = request.ImageDPI
	}

	if !dara.IsNil(request.LongPicture) {
		query["LongPicture"] = request.LongPicture
	}

	if !dara.IsNil(request.LongText) {
		query["LongText"] = request.LongText
	}

	if !dara.IsNil(request.MaxSheetColumn) {
		query["MaxSheetColumn"] = request.MaxSheetColumn
	}

	if !dara.IsNil(request.MaxSheetRow) {
		query["MaxSheetRow"] = request.MaxSheetRow
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Pages) {
		query["Pages"] = request.Pages
	}

	if !dara.IsNil(request.PaperHorizontal) {
		query["PaperHorizontal"] = request.PaperHorizontal
	}

	if !dara.IsNil(request.PaperSize) {
		query["PaperSize"] = request.PaperSize
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Quality) {
		query["Quality"] = request.Quality
	}

	if !dara.IsNil(request.ScalePercentage) {
		query["ScalePercentage"] = request.ScalePercentage
	}

	if !dara.IsNil(request.SheetCount) {
		query["SheetCount"] = request.SheetCount
	}

	if !dara.IsNil(request.SheetIndex) {
		query["SheetIndex"] = request.SheetIndex
	}

	if !dara.IsNil(request.ShowComments) {
		query["ShowComments"] = request.ShowComments
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.StartPage) {
		query["StartPage"] = request.StartPage
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	if !dara.IsNil(request.TargetURIPrefix) {
		query["TargetURIPrefix"] = request.TargetURIPrefix
	}

	if !dara.IsNil(request.TrimPolicyShrink) {
		query["TrimPolicy"] = request.TrimPolicyShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourcesShrink) {
		body["Sources"] = request.SourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOfficeConversionTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a document format conversion task to convert the format of a document stored in an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Supported input formats:
//
//	    	- Text documents: doc, docx, wps, wpss, docm, dotm, dot, dotx, and html
//
//	    	- Presentation documents: pptx, ppt, pot, potx, pps, ppsx, dps, dpt, pptm, potm, ppsm, and dpss
//
//	    	- Spreadsheet documents: xls, xlt, et, ett, xlsx, xltx, csv, xlsb, xlsm, xltm, and ets
//
//	    	- PDF documents: pdf
//
//		- Supported output formats:
//
//	    	- Image files: png and jpg
//
//	    	- Text files: txt
//
//	    	- PDF files: pdf
//
//		- Each input document can be up to 200 MB in size. The upper limit cannot be adjusted.
//
//		- If the document size is large or the content is complex, the conversion task may time out.
//
//		- The limit on the number of requests per second for a single user is 50.
//
//		- The operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can use one of the following methods to query the task information in a timely manner:
//
//	    	- Call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.``
//
//	    	- In the region in which the IMM project is located, configure a Simple Message Queue (SMQ) subscription to receive task information notifications. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html). For information about SMQ SDKs, see [Use queues](https://help.aliyun.com/document_detail/32449.html).
//
//	    	- In the region in which the IMM project is located, create an ApsaraMQ for RocketMQ 4.0 instance, a topic, and a group to receive task notifications. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html). For more information about how to use ApsaraMQ for RocketMQ, see [Call HTTP SDKs to send and subscribe to messages](https://help.aliyun.com/document_detail/169009.html).
//
//	    	- In the region in which the IMM project is located, use [EventBridge](https://www.aliyun.com/product/aliware/eventbridge) to receive task information notifications. For more information, see [IMM events](https://help.aliyun.com/document_detail/205730.html).
//
// @param request - CreateOfficeConversionTaskRequest
//
// @return CreateOfficeConversionTaskResponse
func (client *Client) CreateOfficeConversionTask(request *CreateOfficeConversionTaskRequest) (_result *CreateOfficeConversionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.CreateOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a project.
//
// Description:
//
//	  The name of a project must be unique in a region.
//
//		- By default, you can create up to 100 projects in a region. If you want to request a quota increase to create more projects, submit a ticket or join the DingTalk chat group (ID: 88490020073).
//
//		- After you create a project, you can create other Intelligent Media Management (IMM) resources in the project. For more information, see the following links:
//
//	    	- [CreateDataset](https://help.aliyun.com/document_detail/478160.html)
//
//	    	- [CreateTrigger](https://help.aliyun.com/document_detail/479912.html)
//
//	    	- [CreateBatch](https://help.aliyun.com/document_detail/606694.html)
//
//	    	- [CreateBinding](https://help.aliyun.com/document_detail/478202.html)
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetMaxBindCount) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !dara.IsNil(request.DatasetMaxEntityCount) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !dara.IsNil(request.DatasetMaxFileCount) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !dara.IsNil(request.DatasetMaxRelationCount) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !dara.IsNil(request.DatasetMaxTotalFileSize) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ProjectMaxDatasetCount) {
		query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ServiceRole) {
		query["ServiceRole"] = request.ServiceRole
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a project.
//
// Description:
//
//	  The name of a project must be unique in a region.
//
//		- By default, you can create up to 100 projects in a region. If you want to request a quota increase to create more projects, submit a ticket or join the DingTalk chat group (ID: 88490020073).
//
//		- After you create a project, you can create other Intelligent Media Management (IMM) resources in the project. For more information, see the following links:
//
//	    	- [CreateDataset](https://help.aliyun.com/document_detail/478160.html)
//
//	    	- [CreateTrigger](https://help.aliyun.com/document_detail/479912.html)
//
//	    	- [CreateBatch](https://help.aliyun.com/document_detail/606694.html)
//
//	    	- [CreateBinding](https://help.aliyun.com/document_detail/478202.html)
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clusters images indexed into a dataset by similarity. Image clustering is suitable for image deduplication and selection. For example, you can use image clustering to filter photos in your album that are taken in continuous shooting mode.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note that*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Each call to the operation incrementally processes metadata in the dataset.****`` You can regularly call this operation to process incremental files.
//
//		- After clustering is completed, you can call the [QuerySimilarImageClusters](https://help.aliyun.com/document_detail/611304.html) operation to query image clustering results.
//
//		- An image cluster contains at lest two images. Removing similar images from the dataset affects existing image clusters. If image deletion reduces the number of images in a cluster to less than 2, the cluster is automatically deleted.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateSimilarImageClusteringTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimilarImageClusteringTaskResponse
func (client *Client) CreateSimilarImageClusteringTaskWithOptions(tmpReq *CreateSimilarImageClusteringTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSimilarImageClusteringTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSimilarImageClusteringTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSimilarImageClusteringTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSimilarImageClusteringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clusters images indexed into a dataset by similarity. Image clustering is suitable for image deduplication and selection. For example, you can use image clustering to filter photos in your album that are taken in continuous shooting mode.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note that*	- Asynchronous processing does not guarantee timely task completion.
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Each call to the operation incrementally processes metadata in the dataset.****`` You can regularly call this operation to process incremental files.
//
//		- After clustering is completed, you can call the [QuerySimilarImageClusters](https://help.aliyun.com/document_detail/611304.html) operation to query image clustering results.
//
//		- An image cluster contains at lest two images. Removing similar images from the dataset affects existing image clusters. If image deletion reduces the number of images in a cluster to less than 2, the cluster is automatically deleted.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateSimilarImageClusteringTaskRequest
//
// @return CreateSimilarImageClusteringTaskResponse
func (client *Client) CreateSimilarImageClusteringTask(request *CreateSimilarImageClusteringTaskRequest) (_result *CreateSimilarImageClusteringTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSimilarImageClusteringTaskResponse{}
	_body, _err := client.CreateSimilarImageClusteringTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a story.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) to query information about the task. If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoryResponse
func (client *Client) CreateStoryWithOptions(tmpReq *CreateStoryRequest, runtime *dara.RuntimeOptions) (_result *CreateStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Address) {
		request.AddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Address, dara.String("Address"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomLabels) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, dara.String("CustomLabels"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddressShrink) {
		body["Address"] = request.AddressShrink
	}

	if !dara.IsNil(request.CustomId) {
		body["CustomId"] = request.CustomId
	}

	if !dara.IsNil(request.CustomLabelsShrink) {
		body["CustomLabels"] = request.CustomLabelsShrink
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxFileCount) {
		body["MaxFileCount"] = request.MaxFileCount
	}

	if !dara.IsNil(request.MinFileCount) {
		body["MinFileCount"] = request.MinFileCount
	}

	if !dara.IsNil(request.NotifyTopicName) {
		body["NotifyTopicName"] = request.NotifyTopicName
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.StoryEndTime) {
		body["StoryEndTime"] = request.StoryEndTime
	}

	if !dara.IsNil(request.StoryName) {
		body["StoryName"] = request.StoryName
	}

	if !dara.IsNil(request.StoryStartTime) {
		body["StoryStartTime"] = request.StoryStartTime
	}

	if !dara.IsNil(request.StorySubType) {
		body["StorySubType"] = request.StorySubType
	}

	if !dara.IsNil(request.StoryType) {
		body["StoryType"] = request.StoryType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a story.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The operation is an asynchronous operation. After a task is executed, the task information is saved only for seven days. When the retention period ends, the task information can no longer be retrieved. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) to query information about the task. If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateStoryRequest
//
// @return CreateStoryResponse
func (client *Client) CreateStory(request *CreateStoryRequest) (_result *CreateStoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStoryResponse{}
	_body, _err := client.CreateStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a trigger. A trigger can trigger Intelligent Media Management (IMM) based on events such as events in Object Storage Service (OSS) to process files, such as images, videos, and documents based on data processing templates.
//
// Description:
//
// If you want to create a trigger to process data in [OSS](https://help.aliyun.com/document_detail/99372.html), make sure that you have bound the dataset to the OSS bucket where the data is stored. For more information about how to bind a dataset to a bucket, see [AttachOSSBucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param tmpReq - CreateTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTriggerResponse
func (client *Client) CreateTriggerWithOptions(tmpReq *CreateTriggerRequest, runtime *dara.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTriggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Actions) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, dara.String("Actions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionsShrink) {
		body["Actions"] = request.ActionsShrink
	}

	if !dara.IsNil(request.InputShrink) {
		body["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		body["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ServiceRole) {
		body["ServiceRole"] = request.ServiceRole
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrigger"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trigger. A trigger can trigger Intelligent Media Management (IMM) based on events such as events in Object Storage Service (OSS) to process files, such as images, videos, and documents based on data processing templates.
//
// Description:
//
// If you want to create a trigger to process data in [OSS](https://help.aliyun.com/document_detail/99372.html), make sure that you have bound the dataset to the OSS bucket where the data is stored. For more information about how to bind a dataset to a bucket, see [AttachOSSBucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param request - CreateTriggerRequest
//
// @return CreateTriggerResponse
func (client *Client) CreateTrigger(request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects the scene, object, and event tag information of video content. Scene information includes categories such as natural landscapes, life scenes, and disaster scenes. Event information includes categories such as talent shows, office events, performances, and production events. Object information includes categories such as tableware, electronic products, furniture, and transportation. Video tag detection supports more than 30 tag categories and thousands of tags.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/2747104.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that an IMM project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- For more information about video label detection, see [Video label detection](https://help.aliyun.com/document_detail/477189.html).
//
//		- This operation supports multiple video formats, such as MP4, MPEG-TS, MKV, MOV, AVI, FLV, and M3U8.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param tmpReq - CreateVideoLabelClassificationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoLabelClassificationTaskResponse
func (client *Client) CreateVideoLabelClassificationTaskWithOptions(tmpReq *CreateVideoLabelClassificationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVideoLabelClassificationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVideoLabelClassificationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoLabelClassificationTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoLabelClassificationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects the scene, object, and event tag information of video content. Scene information includes categories such as natural landscapes, life scenes, and disaster scenes. Event information includes categories such as talent shows, office events, performances, and production events. Object information includes categories such as tableware, electronic products, furniture, and transportation. Video tag detection supports more than 30 tag categories and thousands of tags.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/2747104.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that an IMM project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- For more information about video label detection, see [Video label detection](https://help.aliyun.com/document_detail/477189.html).
//
//		- This operation supports multiple video formats, such as MP4, MPEG-TS, MKV, MOV, AVI, FLV, and M3U8.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications.
//
// @param request - CreateVideoLabelClassificationTaskRequest
//
// @return CreateVideoLabelClassificationTaskResponse
func (client *Client) CreateVideoLabelClassificationTask(request *CreateVideoLabelClassificationTaskRequest) (_result *CreateVideoLabelClassificationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVideoLabelClassificationTaskResponse{}
	_body, _err := client.CreateVideoLabelClassificationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects risky or non-compliant content from videos. You can use this operation in scenarios such as intelligent pornography detection, terrorist content and political bias detection, ad violation detection, and logo detection.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- The detection result is sent as an asynchronous notification. The Suggestion parameter in asynchronous notifications supports the following values:
//
//	    	- pass: No non-compliant content is found.
//
//	    	- block: Non-compliant content is detected. The Categories field value indicates the non-compliance category. For more information, see [Content moderation results](https://help.aliyun.com/document_detail/2743995.html).
//
//	    	- review: A manual review is needed. After the manual review is completed, an asynchronous notification is sent to inform you about the result.
//
//		- The following video frame requirements apply:
//
//	    	- The URLs for video frames must use HTTP or HTTPS.
//
//	    	- Video frames must be in PNG, JPG, JPEG, BMP, GIF, or WebP format.
//
//	    	- The size of a video frame cannot exceed 10 MB.
//
//	    	- The resolution for video frames is not lower than 256 × 256 pixels. A frame resolution lower than this recommended resolution may affect detection accuracy.
//
//	    	- The response time of the operation varies based on the amount of time required to download frames. Make sure that video frames to be detected are stored in a reliable and stable service. We recommend that you store video frames in OSS or cache video frames on Alibaba Cloud CDN.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications. >
//
// @param tmpReq - CreateVideoModerationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoModerationTaskResponse
func (client *Client) CreateVideoModerationTaskWithOptions(tmpReq *CreateVideoModerationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVideoModerationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVideoModerationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Scenes) {
		request.ScenesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scenes, dara.String("Scenes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.MaxFrames) {
		query["MaxFrames"] = request.MaxFrames
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ScenesShrink) {
		query["Scenes"] = request.ScenesShrink
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoModerationTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoModerationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects risky or non-compliant content from videos. You can use this operation in scenarios such as intelligent pornography detection, terrorist content and political bias detection, ad violation detection, and logo detection.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//	    **
//
//	    **Note*	- Asynchronous processing does not guarantee timely task completion.
//
//		- The detection result is sent as an asynchronous notification. The Suggestion parameter in asynchronous notifications supports the following values:
//
//	    	- pass: No non-compliant content is found.
//
//	    	- block: Non-compliant content is detected. The Categories field value indicates the non-compliance category. For more information, see [Content moderation results](https://help.aliyun.com/document_detail/2743995.html).
//
//	    	- review: A manual review is needed. After the manual review is completed, an asynchronous notification is sent to inform you about the result.
//
//		- The following video frame requirements apply:
//
//	    	- The URLs for video frames must use HTTP or HTTPS.
//
//	    	- Video frames must be in PNG, JPG, JPEG, BMP, GIF, or WebP format.
//
//	    	- The size of a video frame cannot exceed 10 MB.
//
//	    	- The resolution for video frames is not lower than 256 × 256 pixels. A frame resolution lower than this recommended resolution may affect detection accuracy.
//
//	    	- The response time of the operation varies based on the amount of time required to download frames. Make sure that video frames to be detected are stored in a reliable and stable service. We recommend that you store video frames in OSS or cache video frames on Alibaba Cloud CDN.
//
//		- This operation is an asynchronous operation. After a task is executed, the task information is retained only for seven days and cannot be retrieved when the retention period elapses. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query information about the task.`` If you specify [Notification](https://help.aliyun.com/document_detail/2743997.html), you can obtain information about the task based on notifications. >
//
// @param request - CreateVideoModerationTaskRequest
//
// @return CreateVideoModerationTaskResponse
func (client *Client) CreateVideoModerationTask(request *CreateVideoModerationTaskRequest) (_result *CreateVideoModerationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVideoModerationTaskResponse{}
	_body, _err := client.CreateVideoModerationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a batch processing task.
//
// Description:
//
//	  You can delete only a batch processing task that is in one of the following states: Ready, Failed, Suspended, and Succeeded.
//
//		- Before you delete a batch processing task, you can call the [GetBatch](https://help.aliyun.com/document_detail/479922.html) operation to query the task status. This ensures a successful deletion.
//
// @param request - DeleteBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBatchResponse
func (client *Client) DeleteBatchWithOptions(request *DeleteBatchRequest, runtime *dara.RuntimeOptions) (_result *DeleteBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBatch"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a batch processing task.
//
// Description:
//
//	  You can delete only a batch processing task that is in one of the following states: Ready, Failed, Suspended, and Succeeded.
//
//		- Before you delete a batch processing task, you can call the [GetBatch](https://help.aliyun.com/document_detail/479922.html) operation to query the task status. This ensures a successful deletion.
//
// @param request - DeleteBatchRequest
//
// @return DeleteBatchResponse
func (client *Client) DeleteBatch(request *DeleteBatchRequest) (_result *DeleteBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBatchResponse{}
	_body, _err := client.DeleteBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the binding between a dataset and an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- If you delete a binding, new changes in the OSS bucket are not synchronized to the dataset. Exercise caution when you perform this operation.
//
// @param request - DeleteBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBindingResponse
func (client *Client) DeleteBindingWithOptions(request *DeleteBindingRequest, runtime *dara.RuntimeOptions) (_result *DeleteBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBinding"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the binding between a dataset and an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- If you delete a binding, new changes in the OSS bucket are not synchronized to the dataset. Exercise caution when you perform this operation.
//
// @param request - DeleteBindingRequest
//
// @return DeleteBindingResponse
func (client *Client) DeleteBinding(request *DeleteBindingRequest) (_result *DeleteBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBindingResponse{}
	_body, _err := client.DeleteBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dataset.
//
// Description:
//
//	  Before you delete a dataset, make sure that you have deleted all indexes in the dataset. For more information about how to delete indexes, see [DeleteFileMeta](https://help.aliyun.com/document_detail/478172.html) and [BatchDeleteFileMeta](https://help.aliyun.com/document_detail/478173.html).
//
//		- Before you [delete a dataset](https://help.aliyun.com/document_detail/478160.html), make sure that you have deleted all bindings between the dataset and Object Storage Service (OSS) buckets. For more information about how to delete a binding, see [DeleteBinding](https://help.aliyun.com/document_detail/478205.html). The [DeleteBinding](https://help.aliyun.com/document_detail/478205.html) operation does not delete an index that is manually created, even if you set the `Cleanup` parameter to `true`. To delete indexes that are manually created, you must call the [DeleteFileMeta](https://help.aliyun.com/document_detail/478172.html) or [BatchDeleteFileMeta](https://help.aliyun.com/document_detail/478173.html) operation. For more information about the differences between automatically and manually created indexes, see [Create a metadata index](https://help.aliyun.com/document_detail/478166.html).
//
// @param request - DeleteDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithOptions(request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset.
//
// Description:
//
//	  Before you delete a dataset, make sure that you have deleted all indexes in the dataset. For more information about how to delete indexes, see [DeleteFileMeta](https://help.aliyun.com/document_detail/478172.html) and [BatchDeleteFileMeta](https://help.aliyun.com/document_detail/478173.html).
//
//		- Before you [delete a dataset](https://help.aliyun.com/document_detail/478160.html), make sure that you have deleted all bindings between the dataset and Object Storage Service (OSS) buckets. For more information about how to delete a binding, see [DeleteBinding](https://help.aliyun.com/document_detail/478205.html). The [DeleteBinding](https://help.aliyun.com/document_detail/478205.html) operation does not delete an index that is manually created, even if you set the `Cleanup` parameter to `true`. To delete indexes that are manually created, you must call the [DeleteFileMeta](https://help.aliyun.com/document_detail/478172.html) or [BatchDeleteFileMeta](https://help.aliyun.com/document_detail/478173.html) operation. For more information about the differences between automatically and manually created indexes, see [Create a metadata index](https://help.aliyun.com/document_detail/478166.html).
//
// @param request - DeleteDatasetRequest
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDataset(request *DeleteDatasetRequest) (_result *DeleteDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the metadata of a file from a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- A successful deletion message is returned regardless of whether the metadata of the file exists in the dataset.
//
// >
//
//   - The objects stored in Object Storage Service (OSS) or Photo and Drive Service are **not*	- deleted if you delete metadata from a dataset. If you want to delete the file, call the corresponding operations of OSS and Photo and Drive Service.
//
//   - When you delete file metadata, the corresponding face clustering group information and story (if any) are changed, but the spatiotemporal clustering is not changed.
//
// @param request - DeleteFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileMetaResponse
func (client *Client) DeleteFileMetaWithOptions(request *DeleteFileMetaRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the metadata of a file from a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- A successful deletion message is returned regardless of whether the metadata of the file exists in the dataset.
//
// >
//
//   - The objects stored in Object Storage Service (OSS) or Photo and Drive Service are **not*	- deleted if you delete metadata from a dataset. If you want to delete the file, call the corresponding operations of OSS and Photo and Drive Service.
//
//   - When you delete file metadata, the corresponding face clustering group information and story (if any) are changed, but the spatiotemporal clustering is not changed.
//
// @param request - DeleteFileMetaRequest
//
// @return DeleteFileMetaResponse
func (client *Client) DeleteFileMeta(request *DeleteFileMetaRequest) (_result *DeleteFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFileMetaResponse{}
	_body, _err := client.DeleteFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a spatiotemporal cluster.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.****
//
//		- Before you call this operation, you must call the [CreateLocationDateClusteringTask](https://help.aliyun.com/document_detail/478188.html) operation to perform spatiotemporal clustering.
//
//		- A successful deletion is returned regardless of whether a spatiotemporal clustering group ID exists.
//
// @param request - DeleteLocationDateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLocationDateClusterResponse
func (client *Client) DeleteLocationDateClusterWithOptions(request *DeleteLocationDateClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteLocationDateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLocationDateCluster"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLocationDateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a spatiotemporal cluster.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.****
//
//		- Before you call this operation, you must call the [CreateLocationDateClusteringTask](https://help.aliyun.com/document_detail/478188.html) operation to perform spatiotemporal clustering.
//
//		- A successful deletion is returned regardless of whether a spatiotemporal clustering group ID exists.
//
// @param request - DeleteLocationDateClusterRequest
//
// @return DeleteLocationDateClusterResponse
func (client *Client) DeleteLocationDateCluster(request *DeleteLocationDateClusterRequest) (_result *DeleteLocationDateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLocationDateClusterResponse{}
	_body, _err := client.DeleteLocationDateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a project.
//
// Description:
//
//	  Before you delete a project, make sure that all resources in the project, such as datasets, bindings, batch processing tasks, and triggers, are deleted. For more information, see [DeleteDataset](https://help.aliyun.com/document_detail/478164.html), [DeleteBatch](https://help.aliyun.com/document_detail/479918.html), and [DeleteTrigger](https://help.aliyun.com/document_detail/479915.html).
//
//		- After a project is deleted, all resources used by the project are recycled, and all related data is lost and cannot be recovered.
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a project.
//
// Description:
//
//	  Before you delete a project, make sure that all resources in the project, such as datasets, bindings, batch processing tasks, and triggers, are deleted. For more information, see [DeleteDataset](https://help.aliyun.com/document_detail/478164.html), [DeleteBatch](https://help.aliyun.com/document_detail/479918.html), and [DeleteTrigger](https://help.aliyun.com/document_detail/479915.html).
//
//		- After a project is deleted, all resources used by the project are recycled, and all related data is lost and cannot be recovered.
//
// @param request - DeleteProjectRequest
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a story.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param request - DeleteStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStoryResponse
func (client *Client) DeleteStoryWithOptions(request *DeleteStoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a story.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param request - DeleteStoryRequest
//
// @return DeleteStoryResponse
func (client *Client) DeleteStory(request *DeleteStoryRequest) (_result *DeleteStoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStoryResponse{}
	_body, _err := client.DeleteStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a trigger.
//
// Description:
//
// You can delete a trigger only if the trigger is in one of the following states: Ready, Failed, Suspended, and Succeeded. You cannot delete a trigger that is in the Running state.
//
// @param request - DeleteTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTriggerResponse
func (client *Client) DeleteTriggerWithOptions(request *DeleteTriggerRequest, runtime *dara.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrigger"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a trigger.
//
// Description:
//
// You can delete a trigger only if the trigger is in one of the following states: Ready, Failed, Suspended, and Succeeded. You cannot delete a trigger that is in the Running state.
//
// @param request - DeleteTriggerRequest
//
// @return DeleteTriggerResponse
func (client *Client) DeleteTrigger(request *DeleteTriggerRequest) (_result *DeleteTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.DeleteTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds an Object Storage Service (OSS) bucket from the corresponding project.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that the project is bound to a bucket. For more information, see [AttachOSSBucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param request - DetachOSSBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachOSSBucketResponse
func (client *Client) DetachOSSBucketWithOptions(request *DetachOSSBucketRequest, runtime *dara.RuntimeOptions) (_result *DetachOSSBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OSSBucket) {
		query["OSSBucket"] = request.OSSBucket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachOSSBucket"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachOSSBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an Object Storage Service (OSS) bucket from the corresponding project.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that the project is bound to a bucket. For more information, see [AttachOSSBucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param request - DetachOSSBucketRequest
//
// @return DetachOSSBucketResponse
func (client *Client) DetachOSSBucket(request *DetachOSSBucketRequest) (_result *DetachOSSBucketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachOSSBucketResponse{}
	_body, _err := client.DetachOSSBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects human body information, such as the confidence level and body bounding box, in an image.
//
// Description:
//
//	  Before you call this operation, make sure that an Intelligent Media Management (IMM) project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- For information about the image encoding formats supported by this operation, see [Limits on images](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - DetectImageBodiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageBodiesResponse
func (client *Client) DetectImageBodiesWithOptions(tmpReq *DetectImageBodiesRequest, runtime *dara.RuntimeOptions) (_result *DetectImageBodiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageBodiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Sensitivity) {
		query["Sensitivity"] = request.Sensitivity
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageBodies"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects human body information, such as the confidence level and body bounding box, in an image.
//
// Description:
//
//	  Before you call this operation, make sure that an Intelligent Media Management (IMM) project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- For information about the image encoding formats supported by this operation, see [Limits on images](https://help.aliyun.com/document_detail/475569.html).
//
// @param request - DetectImageBodiesRequest
//
// @return DetectImageBodiesResponse
func (client *Client) DetectImageBodies(request *DetectImageBodiesRequest) (_result *DetectImageBodiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.DetectImageBodiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects the outline data, attributes, and license plate information of vehicles in an image. The vehicle attributes include the vehicle color (CarColor) and vehicle type (CarType). The license plate information includes the recognition content (Content) and plate frame (Boundary).
//
// Description:
//
//	For information about the image encoding formats supported by this operation, see [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - DetectImageCarsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageCarsResponse
func (client *Client) DetectImageCarsWithOptions(tmpReq *DetectImageCarsRequest, runtime *dara.RuntimeOptions) (_result *DetectImageCarsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageCarsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageCars"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageCarsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects the outline data, attributes, and license plate information of vehicles in an image. The vehicle attributes include the vehicle color (CarColor) and vehicle type (CarType). The license plate information includes the recognition content (Content) and plate frame (Boundary).
//
// Description:
//
//	For information about the image encoding formats supported by this operation, see [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// @param request - DetectImageCarsRequest
//
// @return DetectImageCarsResponse
func (client *Client) DetectImageCars(request *DetectImageCarsRequest) (_result *DetectImageCarsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageCarsResponse{}
	_body, _err := client.DetectImageCarsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects barcodes and QR codes in an image.
//
// Description:
//
//	For information about the image encoding formats supported by this operation, see [Limits on images](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - DetectImageCodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageCodesResponse
func (client *Client) DetectImageCodesWithOptions(tmpReq *DetectImageCodesRequest, runtime *dara.RuntimeOptions) (_result *DetectImageCodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageCodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageCodes"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageCodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects barcodes and QR codes in an image.
//
// Description:
//
//	For information about the image encoding formats supported by this operation, see [Limits on images](https://help.aliyun.com/document_detail/475569.html).
//
// @param request - DetectImageCodesRequest
//
// @return DetectImageCodesResponse
func (client *Client) DetectImageCodes(request *DetectImageCodesRequest) (_result *DetectImageCodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageCodesResponse{}
	_body, _err := client.DetectImageCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects the cropping area that produces the optimal visual effect based on a given image ratio by using AI model capabilities.
//
// @param tmpReq - DetectImageCroppingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageCroppingResponse
func (client *Client) DetectImageCroppingWithOptions(tmpReq *DetectImageCroppingRequest, runtime *dara.RuntimeOptions) (_result *DetectImageCroppingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageCroppingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AspectRatios) {
		query["AspectRatios"] = request.AspectRatios
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageCropping"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageCroppingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects the cropping area that produces the optimal visual effect based on a given image ratio by using AI model capabilities.
//
// @param request - DetectImageCroppingRequest
//
// @return DetectImageCroppingResponse
func (client *Client) DetectImageCropping(request *DetectImageCroppingRequest) (_result *DetectImageCroppingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageCroppingResponse{}
	_body, _err := client.DetectImageCroppingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects faces from an image, including face boundary information, attributes, and quality. The boundary information includes the distance from the y-coordinate of the vertex to the top edge (Top), distance from the x-coordinate of the vertex to the left edge (Left), height (Height), and width (Width). Face attributes include the age (Age), age standard deviation (AgeSD), gender (Gender), emotion (Emotion), mouth opening (Mouth), beard (Beard), hat wearing (Hat), mask wearing (Mask), glasses wearing (Glasses), head orientation (HeadPose), attractiveness (Attractive), and confidence levels for preceding attributes. Quality information includes the face quality score (FaceQuality) and face resolution (Sharpness).
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- For information about the image encoding formats supported by this operation, see [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - DetectImageFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageFacesResponse
func (client *Client) DetectImageFacesWithOptions(tmpReq *DetectImageFacesRequest, runtime *dara.RuntimeOptions) (_result *DetectImageFacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageFacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageFaces"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects faces from an image, including face boundary information, attributes, and quality. The boundary information includes the distance from the y-coordinate of the vertex to the top edge (Top), distance from the x-coordinate of the vertex to the left edge (Left), height (Height), and width (Width). Face attributes include the age (Age), age standard deviation (AgeSD), gender (Gender), emotion (Emotion), mouth opening (Mouth), beard (Beard), hat wearing (Hat), mask wearing (Mask), glasses wearing (Glasses), head orientation (HeadPose), attractiveness (Attractive), and confidence levels for preceding attributes. Quality information includes the face quality score (FaceQuality) and face resolution (Sharpness).
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- For information about the image encoding formats supported by this operation, see [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// @param request - DetectImageFacesRequest
//
// @return DetectImageFacesResponse
func (client *Client) DetectImageFaces(request *DetectImageFacesRequest) (_result *DetectImageFacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.DetectImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects scene, object, and event information in an image. Scene information includes natural landscapes, daily life, and disasters. Event information includes talent shows, office events, performances, and production events. Object information includes tableware, electronics, furniture, and transportation. The DetectImageLabels operation supports more than 30 different categories and thousands of labels.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that an IMM [project](https://help.aliyun.com/document_detail/478273.html) is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- For more information about the features of this operation, see [Image label detection](https://help.aliyun.com/document_detail/477179.html).
//
//		- For more information about the input images supported by this operation, see [Limits on images](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - DetectImageLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageLabelsResponse
func (client *Client) DetectImageLabelsWithOptions(tmpReq *DetectImageLabelsRequest, runtime *dara.RuntimeOptions) (_result *DetectImageLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageLabelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageLabels"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects scene, object, and event information in an image. Scene information includes natural landscapes, daily life, and disasters. Event information includes talent shows, office events, performances, and production events. Object information includes tableware, electronics, furniture, and transportation. The DetectImageLabels operation supports more than 30 different categories and thousands of labels.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that an IMM [project](https://help.aliyun.com/document_detail/478273.html) is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- For more information about the features of this operation, see [Image label detection](https://help.aliyun.com/document_detail/477179.html).
//
//		- For more information about the input images supported by this operation, see [Limits on images](https://help.aliyun.com/document_detail/475569.html).
//
// @param request - DetectImageLabelsRequest
//
// @return DetectImageLabelsResponse
func (client *Client) DetectImageLabels(request *DetectImageLabelsRequest) (_result *DetectImageLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageLabelsResponse{}
	_body, _err := client.DetectImageLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calculates the aesthetics quality score of an image based on metrics such as the composition, brightness, contrast, color, and resolution. The operation returns a score within the range from 0 to 1. A higher score indicates better image quality.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478273.html).[](~~478152~~)
//
//		- For information about the image encoding formats supported by this operation, see [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - DetectImageScoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageScoreResponse
func (client *Client) DetectImageScoreWithOptions(tmpReq *DetectImageScoreRequest, runtime *dara.RuntimeOptions) (_result *DetectImageScoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageScoreShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageScore"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageScoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calculates the aesthetics quality score of an image based on metrics such as the composition, brightness, contrast, color, and resolution. The operation returns a score within the range from 0 to 1. A higher score indicates better image quality.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478273.html).[](~~478152~~)
//
//		- For information about the image encoding formats supported by this operation, see [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// @param request - DetectImageScoreRequest
//
// @return DetectImageScoreResponse
func (client *Client) DetectImageScore(request *DetectImageScoreRequest) (_result *DetectImageScoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageScoreResponse{}
	_body, _err := client.DetectImageScoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Recognizes and extracts text content from an image.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- The size of the image cannot exceed 20 MB.
//
//		- The shortest side of the image is not less than 20 px, and the longest side is not more than 30,000 px.
//
//		- The aspect ratio of the image is less than 1:2.
//
//		- We recommend that you do not use an image that is smaller than 15 px × 15 px in size. Otherwise, the recognition rate is low.
//
// @param tmpReq - DetectImageTextsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageTextsResponse
func (client *Client) DetectImageTextsWithOptions(tmpReq *DetectImageTextsRequest, runtime *dara.RuntimeOptions) (_result *DetectImageTextsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectImageTextsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectImageTexts"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectImageTextsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Recognizes and extracts text content from an image.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- The size of the image cannot exceed 20 MB.
//
//		- The shortest side of the image is not less than 20 px, and the longest side is not more than 30,000 px.
//
//		- The aspect ratio of the image is less than 1:2.
//
//		- We recommend that you do not use an image that is smaller than 15 px × 15 px in size. Otherwise, the recognition rate is low.
//
// @param request - DetectImageTextsRequest
//
// @return DetectImageTextsResponse
func (client *Client) DetectImageTexts(request *DetectImageTextsRequest) (_result *DetectImageTextsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectImageTextsResponse{}
	_body, _err := client.DetectImageTextsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries media metadata, including the media format and stream information.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
// @param tmpReq - DetectMediaMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectMediaMetaResponse
func (client *Client) DetectMediaMetaWithOptions(tmpReq *DetectMediaMetaRequest, runtime *dara.RuntimeOptions) (_result *DetectMediaMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectMediaMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectMediaMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectMediaMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media metadata, including the media format and stream information.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Make sure that the specified project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
// @param request - DetectMediaMetaRequest
//
// @return DetectMediaMetaResponse
func (client *Client) DetectMediaMeta(request *DetectMediaMetaRequest) (_result *DetectMediaMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectMediaMetaResponse{}
	_body, _err := client.DetectMediaMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects whether specified text contains anomalies, such as pornography, advertisements, excessive junk content, politically sensitive content, and abuse.
//
// Description:
//
//	Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
// >  The text compliance detection feature only supports Chinese characters.
//
// @param request - DetectTextAnomalyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectTextAnomalyResponse
func (client *Client) DetectTextAnomalyWithOptions(request *DetectTextAnomalyRequest, runtime *dara.RuntimeOptions) (_result *DetectTextAnomalyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectTextAnomaly"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectTextAnomalyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects whether specified text contains anomalies, such as pornography, advertisements, excessive junk content, politically sensitive content, and abuse.
//
// Description:
//
//	Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
// >  The text compliance detection feature only supports Chinese characters.
//
// @param request - DetectTextAnomalyRequest
//
// @return DetectTextAnomalyResponse
func (client *Client) DetectTextAnomaly(request *DetectTextAnomalyRequest) (_result *DetectTextAnomalyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectTextAnomalyResponse{}
	_body, _err := client.DetectTextAnomalyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Embeds specific textual information into an image as watermarks. These watermarks are visually imperceptible and do not affect the aesthetics of the image or the integrity of the original data. The watermarks can be extracted by using the CreateDecodeBlindWatermarkTask operation.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of Intelligent Media Management (IMM).
//
//		- Make sure that an IMM project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- You can embed only text as blind watermarks to an image.
//
//		- The format of the output image is the same as that of the input image.
//
//		- A blind watermark can still be extracted even if attacks, such as compression, scaling, cropping, and color transformation, are performed on the image.
//
//		- Pure black and white images and images with low resolution (roughly less than 200 px × 200 px,) are not supported.
//
// @param request - EncodeBlindWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncodeBlindWatermarkResponse
func (client *Client) EncodeBlindWatermarkWithOptions(request *EncodeBlindWatermarkRequest, runtime *dara.RuntimeOptions) (_result *EncodeBlindWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ImageQuality) {
		query["ImageQuality"] = request.ImageQuality
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.StrengthLevel) {
		query["StrengthLevel"] = request.StrengthLevel
	}

	if !dara.IsNil(request.TargetURI) {
		query["TargetURI"] = request.TargetURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EncodeBlindWatermark"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EncodeBlindWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Embeds specific textual information into an image as watermarks. These watermarks are visually imperceptible and do not affect the aesthetics of the image or the integrity of the original data. The watermarks can be extracted by using the CreateDecodeBlindWatermarkTask operation.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of Intelligent Media Management (IMM).
//
//		- Make sure that an IMM project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- You can embed only text as blind watermarks to an image.
//
//		- The format of the output image is the same as that of the input image.
//
//		- A blind watermark can still be extracted even if attacks, such as compression, scaling, cropping, and color transformation, are performed on the image.
//
//		- Pure black and white images and images with low resolution (roughly less than 200 px × 200 px,) are not supported.
//
// @param request - EncodeBlindWatermarkRequest
//
// @return EncodeBlindWatermarkResponse
func (client *Client) EncodeBlindWatermark(request *EncodeBlindWatermarkRequest) (_result *EncodeBlindWatermarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EncodeBlindWatermarkResponse{}
	_body, _err := client.EncodeBlindWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Extract text from the document
//
// Description:
//
// - **Before using this interface, please make sure you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/88317.html) of the Intelligent Media Management product.**
//
// - Before calling this interface, ensure that there is an available project ([Project](https://help.aliyun.com/document_detail/478273.html)) in the current Region. For more details, see [Project Management](https://help.aliyun.com/document_detail/478152.html).
//
// - Supports common Word, Excel, PPT, PDF, and TXT documents.
//
// - The file size must not exceed 200 MB. The extracted plain text file size should not exceed 2 MB (approximately 600,000 Chinese characters).
//
//	Notice: If the document format is complex or the text volume is too large, a timeout error may occur. In such scenarios, it is recommended to use the [CreateOfficeConversionTask](478228) interface and specify the output format as txt to achieve similar functionality.
//
// @param tmpReq - ExtractDocumentTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtractDocumentTextResponse
func (client *Client) ExtractDocumentTextWithOptions(tmpReq *ExtractDocumentTextRequest, runtime *dara.RuntimeOptions) (_result *ExtractDocumentTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExtractDocumentTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExtractDocumentText"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExtractDocumentTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Extract text from the document
//
// Description:
//
// - **Before using this interface, please make sure you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/88317.html) of the Intelligent Media Management product.**
//
// - Before calling this interface, ensure that there is an available project ([Project](https://help.aliyun.com/document_detail/478273.html)) in the current Region. For more details, see [Project Management](https://help.aliyun.com/document_detail/478152.html).
//
// - Supports common Word, Excel, PPT, PDF, and TXT documents.
//
// - The file size must not exceed 200 MB. The extracted plain text file size should not exceed 2 MB (approximately 600,000 Chinese characters).
//
//	Notice: If the document format is complex or the text volume is too large, a timeout error may occur. In such scenarios, it is recommended to use the [CreateOfficeConversionTask](478228) interface and specify the output format as txt to achieve similar functionality.
//
// @param request - ExtractDocumentTextRequest
//
// @return ExtractDocumentTextResponse
func (client *Client) ExtractDocumentText(request *ExtractDocumentTextRequest) (_result *ExtractDocumentTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExtractDocumentTextResponse{}
	_body, _err := client.ExtractDocumentTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the extracted file metadata, including the file name, labels, path, custom tags, text, and other fields. If the value of a metadata field of a file matches the specified string, the metadata of the file is returned.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 88490020073) and share your questions with us.
//
//		- For information about the fields that you can use as query conditions, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
//
// @param tmpReq - FuzzyQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FuzzyQueryResponse
func (client *Client) FuzzyQueryWithOptions(tmpReq *FuzzyQueryRequest, runtime *dara.RuntimeOptions) (_result *FuzzyQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FuzzyQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WithFields) {
		request.WithFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WithFields, dara.String("WithFields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.WithFieldsShrink) {
		query["WithFields"] = request.WithFieldsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FuzzyQuery"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FuzzyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the extracted file metadata, including the file name, labels, path, custom tags, text, and other fields. If the value of a metadata field of a file matches the specified string, the metadata of the file is returned.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 88490020073) and share your questions with us.
//
//		- For information about the fields that you can use as query conditions, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
//
// @param request - FuzzyQueryRequest
//
// @return FuzzyQueryResponse
func (client *Client) FuzzyQuery(request *FuzzyQueryRequest) (_result *FuzzyQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FuzzyQueryResponse{}
	_body, _err := client.FuzzyQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a live transcoding playlist and converts video files into M3U8 files. After a playlist is generated, the videos in the playlist are immediately played and the video files are transcoded based on the playback progress. Compared with offline transcoding, online transcoding significantly reduces the time spent in waiting for the videos to be transcoded and reduces transcoding and storage costs.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).**
//
//		- Make sure that the project that you want to use is available in the current region. For more information, see [Project Management](https://help.aliyun.com/document_detail/478152.html).
//
//		- By default, you can call this operation to process only one video, audio, or subtitle track. You can specify the number of the video, audio, or subtitle tracks that you want to process.
//
//		- You can call this operation to generate a media playlist and a master playlist. For more information, see the parameter description.
//
//		- This operation is a synchronous operation. Synchronous or asynchronous transcoding is triggered only during playback or pre-transcoding. You can configure the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to obtain the transcoding task result.
//
//		- For information about the feature description of this operation, see [Live transcoding](https://help.aliyun.com/document_detail/477192.html).
//
//		- The data processing capability of Object Storage Service (OSS) also provides the playlist generation feature. However, this feature can generate only a media playlist, and related parameters are simplified.
//
// @param tmpReq - GenerateVideoPlaylistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateVideoPlaylistResponse
func (client *Client) GenerateVideoPlaylistWithOptions(tmpReq *GenerateVideoPlaylistRequest, runtime *dara.RuntimeOptions) (_result *GenerateVideoPlaylistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateVideoPlaylistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceSubtitles) {
		request.SourceSubtitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceSubtitles, dara.String("SourceSubtitles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Targets) {
		request.TargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Targets, dara.String("Targets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.MasterURI) {
		query["MasterURI"] = request.MasterURI
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.OverwritePolicy) {
		query["OverwritePolicy"] = request.OverwritePolicy
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceDuration) {
		query["SourceDuration"] = request.SourceDuration
	}

	if !dara.IsNil(request.SourceStartTime) {
		query["SourceStartTime"] = request.SourceStartTime
	}

	if !dara.IsNil(request.SourceSubtitlesShrink) {
		query["SourceSubtitles"] = request.SourceSubtitlesShrink
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TargetsShrink) {
		query["Targets"] = request.TargetsShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateVideoPlaylist"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateVideoPlaylistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a live transcoding playlist and converts video files into M3U8 files. After a playlist is generated, the videos in the playlist are immediately played and the video files are transcoded based on the playback progress. Compared with offline transcoding, online transcoding significantly reduces the time spent in waiting for the videos to be transcoded and reduces transcoding and storage costs.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).**
//
//		- Make sure that the project that you want to use is available in the current region. For more information, see [Project Management](https://help.aliyun.com/document_detail/478152.html).
//
//		- By default, you can call this operation to process only one video, audio, or subtitle track. You can specify the number of the video, audio, or subtitle tracks that you want to process.
//
//		- You can call this operation to generate a media playlist and a master playlist. For more information, see the parameter description.
//
//		- This operation is a synchronous operation. Synchronous or asynchronous transcoding is triggered only during playback or pre-transcoding. You can configure the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to obtain the transcoding task result.
//
//		- For information about the feature description of this operation, see [Live transcoding](https://help.aliyun.com/document_detail/477192.html).
//
//		- The data processing capability of Object Storage Service (OSS) also provides the playlist generation feature. However, this feature can generate only a media playlist, and related parameters are simplified.
//
// @param request - GenerateVideoPlaylistRequest
//
// @return GenerateVideoPlaylistResponse
func (client *Client) GenerateVideoPlaylist(request *GenerateVideoPlaylistRequest) (_result *GenerateVideoPlaylistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateVideoPlaylistResponse{}
	_body, _err := client.GenerateVideoPlaylistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an access token for document preview or editing.
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - The access token expires in 30 minutes, and the refresh token expires in 1 day.
//
// - The returned expiration time is in UTC, which has an 8-hour difference from Beijing Time.
//
// - Supported input file formats:
//
//   - Word documents: doc, docx, txt, dot, wps, wpt, dotx, docm, dotm, rtf.
//
//   - Presentation documents (PPT): ppt, pptx, pptm, ppsx, ppsm, pps, potx, potm, dpt, dps.
//
//   - Spreadsheet documents (Excel): et, xls, xlt, xlsx, xlsm, xltx, xltm, csv
//
//   - PDF documents: pdf.
//
// - Supports files up to 200MB.
//
// - Supports documents with a maximum of 5000 pages.
//
// - Projects created before 2023-12-01 are billed based on the number of document openings. Currently, billing is based on the number of API calls. To switch to the new billing model, simply create a new project. Note that one API call can only be used by one user; if reused, only the last user will have normal access, and the access rights of other users will be revoked.
//
// - In the same region as the Intelligent Media Management, activate MNS service, create topics and queues, and configure subscription relationships. You can pass the MNS topic name through the NotifyTopicName parameter to receive message notifications for file saves. For more information about the MNS SDK, see [Receiving and Deleting Messages](https://help.aliyun.com/document_detail/32449.html).
//
// For an example of the JSON format of the Message field in file save message notifications, refer to [WebOffice Message Notification Format](https://help.aliyun.com/document_detail/2743999.html).
//
// > To use the multi-version feature, you must first enable the multi-version feature in OSS, then set the \\"History\\" parameter to true.
//
// >
//
// @param tmpReq - GenerateWebofficeTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateWebofficeTokenResponse
func (client *Client) GenerateWebofficeTokenWithOptions(tmpReq *GenerateWebofficeTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateWebofficeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateWebofficeTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Permission) {
		request.PermissionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permission, dara.String("Permission"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.User) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, dara.String("User"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Watermark) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Watermark, dara.String("Watermark"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CachePreview) {
		query["CachePreview"] = request.CachePreview
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ExternalUploaded) {
		query["ExternalUploaded"] = request.ExternalUploaded
	}

	if !dara.IsNil(request.Filename) {
		query["Filename"] = request.Filename
	}

	if !dara.IsNil(request.Hidecmb) {
		query["Hidecmb"] = request.Hidecmb
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.NotifyTopicName) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PermissionShrink) {
		query["Permission"] = request.PermissionShrink
	}

	if !dara.IsNil(request.PreviewPages) {
		query["PreviewPages"] = request.PreviewPages
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Referer) {
		query["Referer"] = request.Referer
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	if !dara.IsNil(request.UserShrink) {
		query["User"] = request.UserShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.WatermarkShrink) {
		query["Watermark"] = request.WatermarkShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateWebofficeToken"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateWebofficeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an access token for document preview or editing.
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - The access token expires in 30 minutes, and the refresh token expires in 1 day.
//
// - The returned expiration time is in UTC, which has an 8-hour difference from Beijing Time.
//
// - Supported input file formats:
//
//   - Word documents: doc, docx, txt, dot, wps, wpt, dotx, docm, dotm, rtf.
//
//   - Presentation documents (PPT): ppt, pptx, pptm, ppsx, ppsm, pps, potx, potm, dpt, dps.
//
//   - Spreadsheet documents (Excel): et, xls, xlt, xlsx, xlsm, xltx, xltm, csv
//
//   - PDF documents: pdf.
//
// - Supports files up to 200MB.
//
// - Supports documents with a maximum of 5000 pages.
//
// - Projects created before 2023-12-01 are billed based on the number of document openings. Currently, billing is based on the number of API calls. To switch to the new billing model, simply create a new project. Note that one API call can only be used by one user; if reused, only the last user will have normal access, and the access rights of other users will be revoked.
//
// - In the same region as the Intelligent Media Management, activate MNS service, create topics and queues, and configure subscription relationships. You can pass the MNS topic name through the NotifyTopicName parameter to receive message notifications for file saves. For more information about the MNS SDK, see [Receiving and Deleting Messages](https://help.aliyun.com/document_detail/32449.html).
//
// For an example of the JSON format of the Message field in file save message notifications, refer to [WebOffice Message Notification Format](https://help.aliyun.com/document_detail/2743999.html).
//
// > To use the multi-version feature, you must first enable the multi-version feature in OSS, then set the \\"History\\" parameter to true.
//
// >
//
// @param request - GenerateWebofficeTokenRequest
//
// @return GenerateWebofficeTokenResponse
func (client *Client) GenerateWebofficeToken(request *GenerateWebofficeTokenRequest) (_result *GenerateWebofficeTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateWebofficeTokenResponse{}
	_body, _err := client.GenerateWebofficeTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a batch processing task.
//
// @param request - GetBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchResponse
func (client *Client) GetBatchWithOptions(request *GetBatchRequest, runtime *dara.RuntimeOptions) (_result *GetBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatch"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a batch processing task.
//
// @param request - GetBatchRequest
//
// @return GetBatchResponse
func (client *Client) GetBatch(request *GetBatchRequest) (_result *GetBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchResponse{}
	_body, _err := client.GetBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the binding relationship between a specific dataset and an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//		- Make sure that the binding relationship that you want to query exists. For information about how to create a binding relationship, see [CreateBinding](https://help.aliyun.com/document_detail/478202.html).
//
// @param request - GetBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBindingResponse
func (client *Client) GetBindingWithOptions(request *GetBindingRequest, runtime *dara.RuntimeOptions) (_result *GetBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBinding"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding relationship between a specific dataset and an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//		- Make sure that the binding relationship that you want to query exists. For information about how to create a binding relationship, see [CreateBinding](https://help.aliyun.com/document_detail/478202.html).
//
// @param request - GetBindingRequest
//
// @return GetBindingResponse
func (client *Client) GetBinding(request *GetBindingRequest) (_result *GetBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBindingResponse{}
	_body, _err := client.GetBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetDRMLicense is deprecated
//
// Summary:
//
// drmlicense获取
//
// @param request - GetDRMLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDRMLicenseResponse
func (client *Client) GetDRMLicenseWithOptions(request *GetDRMLicenseRequest, runtime *dara.RuntimeOptions) (_result *GetDRMLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.NotifyEndpoint) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !dara.IsNil(request.NotifyTopicName) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ProtectionSystem) {
		query["ProtectionSystem"] = request.ProtectionSystem
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDRMLicense"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDRMLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetDRMLicense is deprecated
//
// Summary:
//
// drmlicense获取
//
// @param request - GetDRMLicenseRequest
//
// @return GetDRMLicenseResponse
// Deprecated
func (client *Client) GetDRMLicense(request *GetDRMLicenseRequest) (_result *GetDRMLicenseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDRMLicenseResponse{}
	_body, _err := client.GetDRMLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- The GetDataset operation supports real-time retrieval of file statistics. You can specify WithStatistics to enable real-time retrieval of file statistics.
//
// @param request - GetDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithOptions(request *GetDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.WithStatistics) {
		query["WithStatistics"] = request.WithStatistics
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataset"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- The GetDataset operation supports real-time retrieval of file statistics. You can specify WithStatistics to enable real-time retrieval of file statistics.
//
// @param request - GetDatasetRequest
//
// @return GetDatasetResponse
func (client *Client) GetDataset(request *GetDatasetRequest) (_result *GetDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of an invisible watermark parsing task.
//
// Description:
//
//	  Before you call this operation, make sure that an Intelligent Media Management (IMM) project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- Before you call this operation, make sure that an invisible watermark task is created and the task ID is obtained.``
//
// @param request - GetDecodeBlindWatermarkResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDecodeBlindWatermarkResultResponse
func (client *Client) GetDecodeBlindWatermarkResultWithOptions(request *GetDecodeBlindWatermarkResultRequest, runtime *dara.RuntimeOptions) (_result *GetDecodeBlindWatermarkResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDecodeBlindWatermarkResult"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDecodeBlindWatermarkResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of an invisible watermark parsing task.
//
// Description:
//
//	  Before you call this operation, make sure that an Intelligent Media Management (IMM) project is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- Before you call this operation, make sure that an invisible watermark task is created and the task ID is obtained.``
//
// @param request - GetDecodeBlindWatermarkResultRequest
//
// @return GetDecodeBlindWatermarkResultResponse
func (client *Client) GetDecodeBlindWatermarkResult(request *GetDecodeBlindWatermarkResultRequest) (_result *GetDecodeBlindWatermarkResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDecodeBlindWatermarkResultResponse{}
	_body, _err := client.GetDecodeBlindWatermarkResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains basic information about face clustering, including the creation time, number of images, and cover.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//		- Before you call this operation, make sure that a face clustering task is created to group all faces in a dataset. For information about how to create a face clustering task, see [CreateFigureClusteringTask](~~CreateFigureClusteringTask~~). For information about how to create a dataset, see [CreateDataset](~~CreateDataset~~).
//
// @param request - GetFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFigureClusterResponse
func (client *Client) GetFigureClusterWithOptions(request *GetFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *GetFigureClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFigureCluster"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains basic information about face clustering, including the creation time, number of images, and cover.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//		- Before you call this operation, make sure that a face clustering task is created to group all faces in a dataset. For information about how to create a face clustering task, see [CreateFigureClusteringTask](~~CreateFigureClusteringTask~~). For information about how to create a dataset, see [CreateDataset](~~CreateDataset~~).
//
// @param request - GetFigureClusterRequest
//
// @return GetFigureClusterResponse
func (client *Client) GetFigureCluster(request *GetFigureClusterRequest) (_result *GetFigureClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFigureClusterResponse{}
	_body, _err := client.GetFigureClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metadata of a file whose metadata is indexed into the dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// @param tmpReq - GetFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileMetaResponse
func (client *Client) GetFileMetaWithOptions(tmpReq *GetFileMetaRequest, runtime *dara.RuntimeOptions) (_result *GetFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WithFields) {
		request.WithFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WithFields, dara.String("WithFields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	if !dara.IsNil(request.WithFieldsShrink) {
		query["WithFields"] = request.WithFieldsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metadata of a file whose metadata is indexed into the dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// @param request - GetFileMetaRequest
//
// @return GetFileMetaResponse
func (client *Client) GetFileMeta(request *GetFileMetaRequest) (_result *GetFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileMetaResponse{}
	_body, _err := client.GetFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an image compliance detection task.
//
// @param request - GetImageModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageModerationResultResponse
func (client *Client) GetImageModerationResultWithOptions(request *GetImageModerationResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageModerationResult"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an image compliance detection task.
//
// @param request - GetImageModerationResultRequest
//
// @return GetImageModerationResultResponse
func (client *Client) GetImageModerationResult(request *GetImageModerationResultRequest) (_result *GetImageModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageModerationResultResponse{}
	_body, _err := client.GetImageModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the name of the project bound to an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//		- Before you call this operation, make sure that [the project whose name you want to query is bound to the specified OSS bucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param request - GetOSSBucketAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSBucketAttachmentResponse
func (client *Client) GetOSSBucketAttachmentWithOptions(request *GetOSSBucketAttachmentRequest, runtime *dara.RuntimeOptions) (_result *GetOSSBucketAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OSSBucket) {
		query["OSSBucket"] = request.OSSBucket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOSSBucketAttachment"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOSSBucketAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the name of the project bound to an Object Storage Service (OSS) bucket.
//
// Description:
//
//	  **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//		- Before you call this operation, make sure that [the project whose name you want to query is bound to the specified OSS bucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param request - GetOSSBucketAttachmentRequest
//
// @return GetOSSBucketAttachmentResponse
func (client *Client) GetOSSBucketAttachment(request *GetOSSBucketAttachmentRequest) (_result *GetOSSBucketAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOSSBucketAttachmentResponse{}
	_body, _err := client.GetOSSBucketAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information, datasets, and file statistics of a project.
//
// Description:
//
// When you call this operation, you can enable the real-time retrieval of file statistics based on your business requirements. For more information, see the "Request parameters" section of this topic.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.WithStatistics) {
		query["WithStatistics"] = request.WithStatistics
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information, datasets, and file statistics of a project.
//
// Description:
//
// When you call this operation, you can enable the real-time retrieval of file statistics based on your business requirements. For more information, see the "Request parameters" section of this topic.
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a story.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param request - GetStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStoryResponse
func (client *Client) GetStoryWithOptions(request *GetStoryRequest, runtime *dara.RuntimeOptions) (_result *GetStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a story.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param request - GetStoryRequest
//
// @return GetStoryResponse
func (client *Client) GetStory(request *GetStoryRequest) (_result *GetStoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStoryResponse{}
	_body, _err := client.GetStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an asynchronous task. Intelligent Media Management (IMM) has multiple asynchronous data processing capabilities, each of which has its own operation for creating tasks. For example, you can call the CreateFigureClusteringTask operation to create a face clustering task and the CreateFileCompressionTask operation to create a file compression task. The GetTask operation is a general operation. You can call this operation to query information about asynchronous tasks by task ID or type.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.
//
// @param request - GetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RequestDefinition) {
		query["RequestDefinition"] = request.RequestDefinition
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an asynchronous task. Intelligent Media Management (IMM) has multiple asynchronous data processing capabilities, each of which has its own operation for creating tasks. For example, you can call the CreateFigureClusteringTask operation to create a face clustering task and the CreateFileCompressionTask operation to create a file compression task. The GetTask operation is a general operation. You can call this operation to query information about asynchronous tasks by task ID or type.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.
//
// @param request - GetTaskRequest
//
// @return GetTaskResponse
func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a trigger.
//
// @param request - GetTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTriggerResponse
func (client *Client) GetTriggerWithOptions(request *GetTriggerRequest, runtime *dara.RuntimeOptions) (_result *GetTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrigger"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a trigger.
//
// @param request - GetTriggerRequest
//
// @return GetTriggerResponse
func (client *Client) GetTrigger(request *GetTriggerRequest) (_result *GetTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTriggerResponse{}
	_body, _err := client.GetTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of a video label detection task.
//
// Description:
//
//	  Before you call this operation, make sure that a [project](https://help.aliyun.com/document_detail/478273.html) is created on Intelligent Media Management (IMM). For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- Before you call this operation, make sure that a video label detection task is created and the `TaskId` of the task is obtained. For more information, see [CreateVideoLabelClassificationTask](https://help.aliyun.com/document_detail/478223.html).
//
// @param request - GetVideoLabelClassificationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoLabelClassificationResultResponse
func (client *Client) GetVideoLabelClassificationResultWithOptions(request *GetVideoLabelClassificationResultRequest, runtime *dara.RuntimeOptions) (_result *GetVideoLabelClassificationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoLabelClassificationResult"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoLabelClassificationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a video label detection task.
//
// Description:
//
//	  Before you call this operation, make sure that a [project](https://help.aliyun.com/document_detail/478273.html) is created on Intelligent Media Management (IMM). For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
//		- Before you call this operation, make sure that a video label detection task is created and the `TaskId` of the task is obtained. For more information, see [CreateVideoLabelClassificationTask](https://help.aliyun.com/document_detail/478223.html).
//
// @param request - GetVideoLabelClassificationResultRequest
//
// @return GetVideoLabelClassificationResultResponse
func (client *Client) GetVideoLabelClassificationResult(request *GetVideoLabelClassificationResultRequest) (_result *GetVideoLabelClassificationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoLabelClassificationResultResponse{}
	_body, _err := client.GetVideoLabelClassificationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of a video moderation task.
//
// @param request - GetVideoModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoModerationResultResponse
func (client *Client) GetVideoModerationResultWithOptions(request *GetVideoModerationResultRequest, runtime *dara.RuntimeOptions) (_result *GetVideoModerationResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoModerationResult"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a video moderation task.
//
// @param request - GetVideoModerationResultRequest
//
// @return GetVideoModerationResultResponse
func (client *Client) GetVideoModerationResult(request *GetVideoModerationResultRequest) (_result *GetVideoModerationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoModerationResultResponse{}
	_body, _err := client.GetVideoModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an index from metadata extracted by using techniques such as label recognition, face detection, and location detection from input files. You can retrieve data from the same dataset by using multiple methods.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- For information about how to create indexes from metadata, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
//
//		- For information about the limits on the maximum number and size of index files that you can create, see the "Limits on datasets" section of the [Limits](https://help.aliyun.com/document_detail/475569.html) topic. For information about how to create a dataset, see the "CreateDataset" topic.
//
//		- For information about the regions in which you can create index files from metadata, see the "Datasets and indexes" section of the [Limits](https://help.aliyun.com/document_detail/475569.html) topic.
//
//		- After you create an index from metadata, you can try [simple query](https://help.aliyun.com/document_detail/478175.html) to retrieve data. For information about other query capabilities, see [Query and statistics](https://help.aliyun.com/document_detail/2402363.html). You can also [create a face clustering task](https://help.aliyun.com/document_detail/478180.html) to group faces. For information about other clustering capabilities, see [Intelligent management](https://help.aliyun.com/document_detail/2402365.html).
//
// **
//
// **Usage notes**
//
//   - The IndexFileMeta operation is asynchronous, indicating that it takes some time to process the data after a request is submitted. After the processing is complete, the metadata is stored in your dataset. The amount of time it takes for this process varies based on [the workflow template, the operator](https://help.aliyun.com/document_detail/466304.html), and the content of the file, ranging from several seconds to several minutes or even longer. You can subscribe to [Simple Message Service](https://help.aliyun.com/document_detail/2743997.html) for task completion notifications.
//
// @param tmpReq - IndexFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IndexFileMetaResponse
func (client *Client) IndexFileMetaWithOptions(tmpReq *IndexFileMetaRequest, runtime *dara.RuntimeOptions) (_result *IndexFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &IndexFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.File) {
		request.FileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.File, dara.String("File"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FileShrink) {
		query["File"] = request.FileShrink
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IndexFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IndexFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an index from metadata extracted by using techniques such as label recognition, face detection, and location detection from input files. You can retrieve data from the same dataset by using multiple methods.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- For information about how to create indexes from metadata, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
//
//		- For information about the limits on the maximum number and size of index files that you can create, see the "Limits on datasets" section of the [Limits](https://help.aliyun.com/document_detail/475569.html) topic. For information about how to create a dataset, see the "CreateDataset" topic.
//
//		- For information about the regions in which you can create index files from metadata, see the "Datasets and indexes" section of the [Limits](https://help.aliyun.com/document_detail/475569.html) topic.
//
//		- After you create an index from metadata, you can try [simple query](https://help.aliyun.com/document_detail/478175.html) to retrieve data. For information about other query capabilities, see [Query and statistics](https://help.aliyun.com/document_detail/2402363.html). You can also [create a face clustering task](https://help.aliyun.com/document_detail/478180.html) to group faces. For information about other clustering capabilities, see [Intelligent management](https://help.aliyun.com/document_detail/2402365.html).
//
// **
//
// **Usage notes**
//
//   - The IndexFileMeta operation is asynchronous, indicating that it takes some time to process the data after a request is submitted. After the processing is complete, the metadata is stored in your dataset. The amount of time it takes for this process varies based on [the workflow template, the operator](https://help.aliyun.com/document_detail/466304.html), and the content of the file, ranging from several seconds to several minutes or even longer. You can subscribe to [Simple Message Service](https://help.aliyun.com/document_detail/2743997.html) for task completion notifications.
//
// @param request - IndexFileMetaRequest
//
// @return IndexFileMetaResponse
func (client *Client) IndexFileMeta(request *IndexFileMetaRequest) (_result *IndexFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IndexFileMetaResponse{}
	_body, _err := client.IndexFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List bound attachments
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/88317.html) of the Intelligent Media Management product before using this interface.**
//
// - Ensure that you have called [Bind Object Storage Bucket](～～478206～～) to bind the OSS Bucket to the project.
//
// @param request - ListAttachedOSSBucketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAttachedOSSBucketsResponse
func (client *Client) ListAttachedOSSBucketsWithOptions(request *ListAttachedOSSBucketsRequest, runtime *dara.RuntimeOptions) (_result *ListAttachedOSSBucketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAttachedOSSBuckets"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAttachedOSSBucketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List bound attachments
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/88317.html) of the Intelligent Media Management product before using this interface.**
//
// - Ensure that you have called [Bind Object Storage Bucket](～～478206～～) to bind the OSS Bucket to the project.
//
// @param request - ListAttachedOSSBucketsRequest
//
// @return ListAttachedOSSBucketsResponse
func (client *Client) ListAttachedOSSBuckets(request *ListAttachedOSSBucketsRequest) (_result *ListAttachedOSSBucketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAttachedOSSBucketsResponse{}
	_body, _err := client.ListAttachedOSSBucketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries batch processing tasks. You can query batch processing tasks based on conditions such task tags and status. The results can be sorted.
//
// @param request - ListBatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBatchesResponse
func (client *Client) ListBatchesWithOptions(request *ListBatchesRequest, runtime *dara.RuntimeOptions) (_result *ListBatchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.TagSelector) {
		query["TagSelector"] = request.TagSelector
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBatches"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries batch processing tasks. You can query batch processing tasks based on conditions such task tags and status. The results can be sorted.
//
// @param request - ListBatchesRequest
//
// @return ListBatchesResponse
func (client *Client) ListBatches(request *ListBatchesRequest) (_result *ListBatchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBatchesResponse{}
	_body, _err := client.ListBatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bindings between a dataset and Object Storage Service (OSS) buckets.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).
//
// @param request - ListBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindingsResponse
func (client *Client) ListBindingsWithOptions(request *ListBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBindings"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bindings between a dataset and Object Storage Service (OSS) buckets.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).
//
// @param request - ListBindingsRequest
//
// @return ListBindingsResponse
func (client *Client) ListBindings(request *ListBindingsRequest) (_result *ListBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBindingsResponse{}
	_body, _err := client.ListBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of datasets. You can query the list by dataset prefix.
//
// @param request - ListDatasetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithOptions(request *ListDatasetsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasets"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of datasets. You can query the list by dataset prefix.
//
// @param request - ListDatasetsRequest
//
// @return ListDatasetsResponse
func (client *Client) ListDatasets(request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries projects. You can call this operation to query the basic information, datasets, and file statistics of multiple projects at the same time.
//
// Description:
//
// The ListProjects operation supports pagination. When you call this operation, you must specify the token that is obtained from the previous query as the value of NextToken. You must also specify MaxResults to limit the number of entries to return.
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries projects. You can call this operation to query the basic information, datasets, and file statistics of multiple projects at the same time.
//
// Description:
//
// The ListProjects operation supports pagination. When you call this operation, you must specify the token that is obtained from the previous query as the value of NextToken. You must also specify MaxResults to limit the number of entries to return.
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the list of regions
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the list of regions
//
// @param request - ListRegionsRequest
//
// @return ListRegionsResponse
func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists tasks based on specific conditions, such as by time range and by tag.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).
//
// @param tmpReq - ListTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(tmpReq *ListTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EndTimeRange) {
		request.EndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EndTimeRange, dara.String("EndTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StartTimeRange) {
		request.StartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartTimeRange, dara.String("StartTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskTypes) {
		request.TaskTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskTypes, dara.String("TaskTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimeRangeShrink) {
		query["EndTimeRange"] = request.EndTimeRangeShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RequestDefinition) {
		query["RequestDefinition"] = request.RequestDefinition
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.StartTimeRangeShrink) {
		query["StartTimeRange"] = request.StartTimeRangeShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagSelector) {
		query["TagSelector"] = request.TagSelector
	}

	if !dara.IsNil(request.TaskTypesShrink) {
		query["TaskTypes"] = request.TaskTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists tasks based on specific conditions, such as by time range and by tag.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries triggers by tag or status.
//
// @param request - ListTriggersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTriggersResponse
func (client *Client) ListTriggersWithOptions(request *ListTriggersRequest, runtime *dara.RuntimeOptions) (_result *ListTriggersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.TagSelector) {
		query["TagSelector"] = request.TagSelector
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTriggers"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTriggersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries triggers by tag or status.
//
// @param request - ListTriggersRequest
//
// @return ListTriggersResponse
func (client *Client) ListTriggers(request *ListTriggersRequest) (_result *ListTriggersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTriggersResponse{}
	_body, _err := client.ListTriggersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries face groups based on given conditions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that a face clustering task is created to group all faces in a dataset. For information about how to create a face clustering task, see [CreateFigureClusteringTask](~~CreateFigureClusteringTask~~). For information about how to create a dataset, see [CreateDataset](~~CreateDataset~~).
//
// @param tmpReq - QueryFigureClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFigureClustersResponse
func (client *Client) QueryFigureClustersWithOptions(tmpReq *QueryFigureClustersRequest, runtime *dara.RuntimeOptions) (_result *QueryFigureClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryFigureClustersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateTimeRange) {
		request.CreateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTimeRange, dara.String("CreateTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateTimeRange) {
		request.UpdateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateTimeRange, dara.String("UpdateTimeRange"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeRangeShrink) {
		query["CreateTimeRange"] = request.CreateTimeRangeShrink
	}

	if !dara.IsNil(request.CustomLabels) {
		query["CustomLabels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.UpdateTimeRangeShrink) {
		query["UpdateTimeRange"] = request.UpdateTimeRangeShrink
	}

	if !dara.IsNil(request.WithTotalCount) {
		query["WithTotalCount"] = request.WithTotalCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFigureClusters"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFigureClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries face groups based on given conditions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that a face clustering task is created to group all faces in a dataset. For information about how to create a face clustering task, see [CreateFigureClusteringTask](~~CreateFigureClusteringTask~~). For information about how to create a dataset, see [CreateDataset](~~CreateDataset~~).
//
// @param request - QueryFigureClustersRequest
//
// @return QueryFigureClustersResponse
func (client *Client) QueryFigureClusters(request *QueryFigureClustersRequest) (_result *QueryFigureClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFigureClustersResponse{}
	_body, _err := client.QueryFigureClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of spatiotemporal clusters based on the specified conditions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.****
//
//		- Before you call this operation, make sure that you have called the [CreateLocationDateClusteringTask](https://help.aliyun.com/document_detail/478188.html) operation to create spatiotemporal clusters in the project.
//
// @param tmpReq - QueryLocationDateClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLocationDateClustersResponse
func (client *Client) QueryLocationDateClustersWithOptions(tmpReq *QueryLocationDateClustersRequest, runtime *dara.RuntimeOptions) (_result *QueryLocationDateClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryLocationDateClustersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Address) {
		request.AddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Address, dara.String("Address"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CreateTimeRange) {
		request.CreateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTimeRange, dara.String("CreateTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LocationDateClusterEndTimeRange) {
		request.LocationDateClusterEndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDateClusterEndTimeRange, dara.String("LocationDateClusterEndTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LocationDateClusterLevels) {
		request.LocationDateClusterLevelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDateClusterLevels, dara.String("LocationDateClusterLevels"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LocationDateClusterStartTimeRange) {
		request.LocationDateClusterStartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDateClusterStartTimeRange, dara.String("LocationDateClusterStartTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateTimeRange) {
		request.UpdateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateTimeRange, dara.String("UpdateTimeRange"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressShrink) {
		query["Address"] = request.AddressShrink
	}

	if !dara.IsNil(request.CreateTimeRangeShrink) {
		query["CreateTimeRange"] = request.CreateTimeRangeShrink
	}

	if !dara.IsNil(request.CustomLabels) {
		query["CustomLabels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.LocationDateClusterEndTimeRangeShrink) {
		query["LocationDateClusterEndTimeRange"] = request.LocationDateClusterEndTimeRangeShrink
	}

	if !dara.IsNil(request.LocationDateClusterLevelsShrink) {
		query["LocationDateClusterLevels"] = request.LocationDateClusterLevelsShrink
	}

	if !dara.IsNil(request.LocationDateClusterStartTimeRangeShrink) {
		query["LocationDateClusterStartTimeRange"] = request.LocationDateClusterStartTimeRangeShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UpdateTimeRangeShrink) {
		query["UpdateTimeRange"] = request.UpdateTimeRangeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLocationDateClusters"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLocationDateClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of spatiotemporal clusters based on the specified conditions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.****
//
//		- Before you call this operation, make sure that you have called the [CreateLocationDateClusteringTask](https://help.aliyun.com/document_detail/478188.html) operation to create spatiotemporal clusters in the project.
//
// @param request - QueryLocationDateClustersRequest
//
// @return QueryLocationDateClustersResponse
func (client *Client) QueryLocationDateClusters(request *QueryLocationDateClustersRequest) (_result *QueryLocationDateClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryLocationDateClustersResponse{}
	_body, _err := client.QueryLocationDateClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to query the list of similar image clusters.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, you must call the [CreateSimilarImageClusteringTask](https://help.aliyun.com/document_detail/611302.html) operation to cluster similar images in the dataset.
//
// @param request - QuerySimilarImageClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySimilarImageClustersResponse
func (client *Client) QuerySimilarImageClustersWithOptions(request *QuerySimilarImageClustersRequest, runtime *dara.RuntimeOptions) (_result *QuerySimilarImageClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomLabels) {
		query["CustomLabels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySimilarImageClusters"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySimilarImageClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the list of similar image clusters.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, you must call the [CreateSimilarImageClusteringTask](https://help.aliyun.com/document_detail/611302.html) operation to cluster similar images in the dataset.
//
// @param request - QuerySimilarImageClustersRequest
//
// @return QuerySimilarImageClustersResponse
func (client *Client) QuerySimilarImageClusters(request *QuerySimilarImageClustersRequest) (_result *QuerySimilarImageClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySimilarImageClustersResponse{}
	_body, _err := client.QuerySimilarImageClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries stories based on the specified conditions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param tmpReq - QueryStoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStoriesResponse
func (client *Client) QueryStoriesWithOptions(tmpReq *QueryStoriesRequest, runtime *dara.RuntimeOptions) (_result *QueryStoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryStoriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateTimeRange) {
		request.CreateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTimeRange, dara.String("CreateTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FigureClusterIds) {
		request.FigureClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FigureClusterIds, dara.String("FigureClusterIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StoryEndTimeRange) {
		request.StoryEndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoryEndTimeRange, dara.String("StoryEndTimeRange"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StoryStartTimeRange) {
		request.StoryStartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoryStartTimeRange, dara.String("StoryStartTimeRange"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeRangeShrink) {
		query["CreateTimeRange"] = request.CreateTimeRangeShrink
	}

	if !dara.IsNil(request.CustomLabels) {
		query["CustomLabels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FigureClusterIdsShrink) {
		query["FigureClusterIds"] = request.FigureClusterIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.StoryEndTimeRangeShrink) {
		query["StoryEndTimeRange"] = request.StoryEndTimeRangeShrink
	}

	if !dara.IsNil(request.StoryName) {
		query["StoryName"] = request.StoryName
	}

	if !dara.IsNil(request.StoryStartTimeRangeShrink) {
		query["StoryStartTimeRange"] = request.StoryStartTimeRangeShrink
	}

	if !dara.IsNil(request.StorySubType) {
		query["StorySubType"] = request.StorySubType
	}

	if !dara.IsNil(request.StoryType) {
		query["StoryType"] = request.StoryType
	}

	if !dara.IsNil(request.WithEmptyStories) {
		query["WithEmptyStories"] = request.WithEmptyStories
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryStories"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryStoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stories based on the specified conditions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param request - QueryStoriesRequest
//
// @return QueryStoriesResponse
func (client *Client) QueryStories(request *QueryStoriesRequest) (_result *QueryStoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryStoriesResponse{}
	_body, _err := client.QueryStoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Refresh Document Preview and Editing Token
//
// Description:
//
// *Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - For detailed billing information, refer to the [WebOffice Billing Instructions](https://help.aliyun.com/document_detail/2639703.html).
//
// - The access token expires after 30 minutes. You must open the preview before the access token expires; otherwise, you will not be able to preview.
//
// - The refresh token expires after 1 day. You need to call the refresh interface before the refresh token expires; otherwise, the token will become invalid.
//
// - The expiration time returned is in UTC, which has an 8-hour difference from Beijing Time.
//
// > The access token is used for actual preview session access, while the refresh token is used to reduce the parameters required for users to refresh tokens. You can use the refresh token to directly obtain a new token based on previous configurations.
//
// >
//
// @param tmpReq - RefreshWebofficeTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshWebofficeTokenResponse
func (client *Client) RefreshWebofficeTokenWithOptions(tmpReq *RefreshWebofficeTokenRequest, runtime *dara.RuntimeOptions) (_result *RefreshWebofficeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RefreshWebofficeTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RefreshToken) {
		query["RefreshToken"] = request.RefreshToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshWebofficeToken"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Refresh Document Preview and Editing Token
//
// Description:
//
// *Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - For detailed billing information, refer to the [WebOffice Billing Instructions](https://help.aliyun.com/document_detail/2639703.html).
//
// - The access token expires after 30 minutes. You must open the preview before the access token expires; otherwise, you will not be able to preview.
//
// - The refresh token expires after 1 day. You need to call the refresh interface before the refresh token expires; otherwise, the token will become invalid.
//
// - The expiration time returned is in UTC, which has an 8-hour difference from Beijing Time.
//
// > The access token is used for actual preview session access, while the refresh token is used to reduce the parameters required for users to refresh tokens. You can use the refresh token to directly obtain a new token based on previous configurations.
//
// >
//
// @param request - RefreshWebofficeTokenRequest
//
// @return RefreshWebofficeTokenResponse
func (client *Client) RefreshWebofficeToken(request *RefreshWebofficeTokenRequest) (_result *RefreshWebofficeTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.RefreshWebofficeTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes files from a story.
//
// @param tmpReq - RemoveStoryFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveStoryFilesResponse
func (client *Client) RemoveStoryFilesWithOptions(tmpReq *RemoveStoryFilesRequest, runtime *dara.RuntimeOptions) (_result *RemoveStoryFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveStoryFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FilesShrink) {
		body["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveStoryFiles"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveStoryFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes files from a story.
//
// @param request - RemoveStoryFilesRequest
//
// @return RemoveStoryFilesResponse
func (client *Client) RemoveStoryFiles(request *RemoveStoryFilesRequest) (_result *RemoveStoryFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveStoryFilesResponse{}
	_body, _err := client.RemoveStoryFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a batch processing task that is in the Suspended or Failed state.
//
// Description:
//
// You can resume a batch processing task only when the task is in the Suspended or Failed state. A batch processing task continues to provide services after you resume the task.
//
// @param request - ResumeBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeBatchResponse
func (client *Client) ResumeBatchWithOptions(request *ResumeBatchRequest, runtime *dara.RuntimeOptions) (_result *ResumeBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeBatch"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a batch processing task that is in the Suspended or Failed state.
//
// Description:
//
// You can resume a batch processing task only when the task is in the Suspended or Failed state. A batch processing task continues to provide services after you resume the task.
//
// @param request - ResumeBatchRequest
//
// @return ResumeBatchResponse
func (client *Client) ResumeBatch(request *ResumeBatchRequest) (_result *ResumeBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeBatchResponse{}
	_body, _err := client.ResumeBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a trigger that is in the Suspended or Failed state.
//
// Description:
//
// You can resume only a trigger that is in the Suspended or Failed state. After you resume a trigger, the trigger continues to provide services as expected.
//
// @param request - ResumeTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTriggerResponse
func (client *Client) ResumeTriggerWithOptions(request *ResumeTriggerRequest, runtime *dara.RuntimeOptions) (_result *ResumeTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeTrigger"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a trigger that is in the Suspended or Failed state.
//
// Description:
//
// You can resume only a trigger that is in the Suspended or Failed state. After you resume a trigger, the trigger continues to provide services as expected.
//
// @param request - ResumeTriggerRequest
//
// @return ResumeTriggerResponse
func (client *Client) ResumeTrigger(request *ResumeTriggerRequest) (_result *ResumeTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeTriggerResponse{}
	_body, _err := client.ResumeTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries face clusters that contain a specific face in an image. Each face cluster contains information such as bounding boxes and similarity.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have created a face clustering task by calling the [CreateFigureClusteringTask](https://help.aliyun.com/document_detail/478180.html) operation to cluster all faces in the dataset.
//
// @param tmpReq - SearchImageFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageFigureClusterResponse
func (client *Client) SearchImageFigureClusterWithOptions(tmpReq *SearchImageFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *SearchImageFigureClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchImageFigureClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImageFigureCluster"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries face clusters that contain a specific face in an image. Each face cluster contains information such as bounding boxes and similarity.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have created a face clustering task by calling the [CreateFigureClusteringTask](https://help.aliyun.com/document_detail/478180.html) operation to cluster all faces in the dataset.
//
// @param request - SearchImageFigureClusterRequest
//
// @return SearchImageFigureClusterResponse
func (client *Client) SearchImageFigureCluster(request *SearchImageFigureClusterRequest) (_result *SearchImageFigureClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchImageFigureClusterResponse{}
	_body, _err := client.SearchImageFigureClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metadata in a dataset by inputting natural language.
//
// Description:
//
// ### [](#)Precautions
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).***	- Each time you call this operation, you are charged for semantic understanding and query fees.
//
//   - Before you call this operation, make sure that the file that you want to use is indexed into the dataset that you use. To index a file into a dataset, you can call one of the following operations: [CreateBinding](https://help.aliyun.com/document_detail/478202.html), [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html), and [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html).
//
//   - The response provided in this example is for reference only. The categories and content of metadata vary based on configurations of [workflow templates](https://help.aliyun.com/document_detail/466304.html). If you have questions, search for and join the DingTalk group numbered 21714099.
//
// ### [](#)Usage limits
//
//   - Each time you call this operation, up to 1,000 metadata files are returned.
//
//   - Pagination is not supported.
//
//   - The natural language processing capability may not always produce completely accurate results.
//
// ### [](#)Usage methods
//
// You can query files within a dataset by using natural language keywords. Key information supported for understanding includes labels (Labels.LabelName), time (ProduceTime), and location (Address.AddressLine). For example, if you use `2023 Hangzhou scenery` as the query criterion, the operation intelligently breaks the query criterion down into the following sub-criteria, and returns the files that meet all the sub-criteria:
//
//   - ProduceTime: 00:00 on January 1, 2023 to 00:00 on December 31, 2023.
//
//   - Address.AddressLine: `Hangzhou`
//
//   - Labels.LabelName: `scenery`.
//
// When you call this operation, you can configure a [workflow template](https://help.aliyun.com/document_detail/466304.html) that includes the `ImageEmbeddingExtraction` operator. This allows the operation to return image content when the query you input matches the image content, thereby achieving intelligent image retrieval.“
//
// @param tmpReq - SemanticQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SemanticQueryResponse
func (client *Client) SemanticQueryWithOptions(tmpReq *SemanticQueryRequest, runtime *dara.RuntimeOptions) (_result *SemanticQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SemanticQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MediaTypes) {
		request.MediaTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MediaTypes, dara.String("MediaTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WithFields) {
		request.WithFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WithFields, dara.String("WithFields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MediaTypesShrink) {
		query["MediaTypes"] = request.MediaTypesShrink
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.WithFieldsShrink) {
		query["WithFields"] = request.WithFieldsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SemanticQuery"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SemanticQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metadata in a dataset by inputting natural language.
//
// Description:
//
// ### [](#)Precautions
//
//   - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).***	- Each time you call this operation, you are charged for semantic understanding and query fees.
//
//   - Before you call this operation, make sure that the file that you want to use is indexed into the dataset that you use. To index a file into a dataset, you can call one of the following operations: [CreateBinding](https://help.aliyun.com/document_detail/478202.html), [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html), and [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html).
//
//   - The response provided in this example is for reference only. The categories and content of metadata vary based on configurations of [workflow templates](https://help.aliyun.com/document_detail/466304.html). If you have questions, search for and join the DingTalk group numbered 21714099.
//
// ### [](#)Usage limits
//
//   - Each time you call this operation, up to 1,000 metadata files are returned.
//
//   - Pagination is not supported.
//
//   - The natural language processing capability may not always produce completely accurate results.
//
// ### [](#)Usage methods
//
// You can query files within a dataset by using natural language keywords. Key information supported for understanding includes labels (Labels.LabelName), time (ProduceTime), and location (Address.AddressLine). For example, if you use `2023 Hangzhou scenery` as the query criterion, the operation intelligently breaks the query criterion down into the following sub-criteria, and returns the files that meet all the sub-criteria:
//
//   - ProduceTime: 00:00 on January 1, 2023 to 00:00 on December 31, 2023.
//
//   - Address.AddressLine: `Hangzhou`
//
//   - Labels.LabelName: `scenery`.
//
// When you call this operation, you can configure a [workflow template](https://help.aliyun.com/document_detail/466304.html) that includes the `ImageEmbeddingExtraction` operator. This allows the operation to return image content when the query you input matches the image content, thereby achieving intelligent image retrieval.“
//
// @param request - SemanticQueryRequest
//
// @return SemanticQueryResponse
func (client *Client) SemanticQuery(request *SemanticQueryRequest) (_result *SemanticQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SemanticQueryResponse{}
	_body, _err := client.SemanticQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries files in a dataset by performing a simple query operation. The operation supports logical expressions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// **Limits**
//
//   - Each query returns information about up to 100 files.
//
//   - Each query returns up to 2,000 aggregations.
//
//   - A subquery supports up to 100 conditions.
//
//   - A subquery can have a maximum nesting depth of 5 levels.
//
// **Example query conditions**
//
//   - Retrieve JPEG images larger than 1,000 pixels:
//
// <!---->
//
//	    {
//
//	      "SubQueries":[
//
//	        {
//
//	          "Field":"ContentType",
//
//	          "Value": "image/jpeg",
//
//	          "Operation":"eq"
//
//	        },
//
//	        {
//
//	          "Field":"ImageWidth",
//
//	          "Value":"1000",
//
//	          "Operation":"gt"
//
//	        }
//
//	      ],
//
//	      "Operation":"and"
//
//	    }
//
//		- Search `oss://examplebucket/path/` for objects that have the `TV` or `Stereo` label and are larger than 10 MB in size:
//
// >  This query requires matching files to have the `TV` or `Stereo` label. The two labels are specified as separate objects in the `Labels` fields.
//
// ```
//
// {
//
//	"SubQueries": [
//
//	  {
//
//	    "Field": "URI",
//
//	    "Value": "oss://examplebucket/path/",
//
//	    "Operation": "prefix"
//
//	  },
//
//	  {
//
//	    "Field": "Size",
//
//	    "Value": "1048576",
//
//	    "Operation": "gt"
//
//	  },
//
//	  {
//
//	    "SubQueries": [
//
//	      {
//
//	        "Field": "Labels.LabelName",
//
//	        "Value": "TV",
//
//	        "Operation": "eq"
//
//	      },
//
//	      {
//
//	        "Field": "Labels.LabelName",
//
//	        "Value": "Stereo",
//
//	        "Operation": "eq"
//
//	      }
//
//	    ],
//
//	    "Operation": "or"
//
//	  }
//
//	],
//
//	"Operation": "and"
//
// }
//
// ```
//
//   - Exclude images that contain a face of a male over the age of 36:
//
// >  In this example query, an image will be excluded from the query results if it contains a face of a male over the age of 36. This query is different from excluding an image that contains a male face or a face of a person over the age of 36. In this query, you need to use the `nested` operator to specify that the conditions are met on the same element.
//
//	    {
//
//	    	"Operation": "not",
//
//	    	"SubQueries": [{
//
//	    		"Operation": "nested",
//
//	    		"SubQueries": [{
//
//	    			"Operation": "and",
//
//	    			"SubQueries": [{
//
//	    				"Field": "Figures.Age",
//
//	    				"Operation": "gt",
//
//	    				"Value": "36"
//
//	    			}, {
//
//	    				"Field": "Figures.Gender",
//
//	    				"Operation": "eq",
//
//	    				"Value": "male"
//
//	    			}]
//
//	    		}]
//
//	    	}]
//
//	    }
//
//		- Query JPEG images that have both custom labels and system labels:
//
// <!---->
//
//	{
//
//	  "SubQueries":[
//
//	    {
//
//	      "Field":"ContentType",
//
//	      "Value": "image/jpeg",
//
//	      "Operation":"eq"
//
//	    },
//
//	    {
//
//	      "Field":"CustomLabels.test",
//
//	      "Operation":"exist"
//
//	    },
//
//	    {
//
//	      "Field":"Labels.LabelName",
//
//	      "Operation":"exist"
//
//	    }
//
//	  ],
//
//	  "Operation":"and"
//
//	}
//
// You can also perform aggregate operations to collect and analyze different data based on the specified conditions. For example, you can calculate the sum, count, average value, or maximum value of all files that meet the query conditions. You can also calculate the size distribution of images that meet the query conditions.
//
// @param tmpReq - SimpleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SimpleQueryResponse
func (client *Client) SimpleQueryWithOptions(tmpReq *SimpleQueryRequest, runtime *dara.RuntimeOptions) (_result *SimpleQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SimpleQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Aggregations) {
		request.AggregationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Aggregations, dara.String("Aggregations"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WithFields) {
		request.WithFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WithFields, dara.String("WithFields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregationsShrink) {
		query["Aggregations"] = request.AggregationsShrink
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.WithFieldsShrink) {
		query["WithFields"] = request.WithFieldsShrink
	}

	if !dara.IsNil(request.WithoutTotalHits) {
		query["WithoutTotalHits"] = request.WithoutTotalHits
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SimpleQuery"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SimpleQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries files in a dataset by performing a simple query operation. The operation supports logical expressions.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
//		- The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// **Limits**
//
//   - Each query returns information about up to 100 files.
//
//   - Each query returns up to 2,000 aggregations.
//
//   - A subquery supports up to 100 conditions.
//
//   - A subquery can have a maximum nesting depth of 5 levels.
//
// **Example query conditions**
//
//   - Retrieve JPEG images larger than 1,000 pixels:
//
// <!---->
//
//	    {
//
//	      "SubQueries":[
//
//	        {
//
//	          "Field":"ContentType",
//
//	          "Value": "image/jpeg",
//
//	          "Operation":"eq"
//
//	        },
//
//	        {
//
//	          "Field":"ImageWidth",
//
//	          "Value":"1000",
//
//	          "Operation":"gt"
//
//	        }
//
//	      ],
//
//	      "Operation":"and"
//
//	    }
//
//		- Search `oss://examplebucket/path/` for objects that have the `TV` or `Stereo` label and are larger than 10 MB in size:
//
// >  This query requires matching files to have the `TV` or `Stereo` label. The two labels are specified as separate objects in the `Labels` fields.
//
// ```
//
// {
//
//	"SubQueries": [
//
//	  {
//
//	    "Field": "URI",
//
//	    "Value": "oss://examplebucket/path/",
//
//	    "Operation": "prefix"
//
//	  },
//
//	  {
//
//	    "Field": "Size",
//
//	    "Value": "1048576",
//
//	    "Operation": "gt"
//
//	  },
//
//	  {
//
//	    "SubQueries": [
//
//	      {
//
//	        "Field": "Labels.LabelName",
//
//	        "Value": "TV",
//
//	        "Operation": "eq"
//
//	      },
//
//	      {
//
//	        "Field": "Labels.LabelName",
//
//	        "Value": "Stereo",
//
//	        "Operation": "eq"
//
//	      }
//
//	    ],
//
//	    "Operation": "or"
//
//	  }
//
//	],
//
//	"Operation": "and"
//
// }
//
// ```
//
//   - Exclude images that contain a face of a male over the age of 36:
//
// >  In this example query, an image will be excluded from the query results if it contains a face of a male over the age of 36. This query is different from excluding an image that contains a male face or a face of a person over the age of 36. In this query, you need to use the `nested` operator to specify that the conditions are met on the same element.
//
//	    {
//
//	    	"Operation": "not",
//
//	    	"SubQueries": [{
//
//	    		"Operation": "nested",
//
//	    		"SubQueries": [{
//
//	    			"Operation": "and",
//
//	    			"SubQueries": [{
//
//	    				"Field": "Figures.Age",
//
//	    				"Operation": "gt",
//
//	    				"Value": "36"
//
//	    			}, {
//
//	    				"Field": "Figures.Gender",
//
//	    				"Operation": "eq",
//
//	    				"Value": "male"
//
//	    			}]
//
//	    		}]
//
//	    	}]
//
//	    }
//
//		- Query JPEG images that have both custom labels and system labels:
//
// <!---->
//
//	{
//
//	  "SubQueries":[
//
//	    {
//
//	      "Field":"ContentType",
//
//	      "Value": "image/jpeg",
//
//	      "Operation":"eq"
//
//	    },
//
//	    {
//
//	      "Field":"CustomLabels.test",
//
//	      "Operation":"exist"
//
//	    },
//
//	    {
//
//	      "Field":"Labels.LabelName",
//
//	      "Operation":"exist"
//
//	    }
//
//	  ],
//
//	  "Operation":"and"
//
//	}
//
// You can also perform aggregate operations to collect and analyze different data based on the specified conditions. For example, you can calculate the sum, count, average value, or maximum value of all files that meet the query conditions. You can also calculate the size distribution of images that meet the query conditions.
//
// @param request - SimpleQueryRequest
//
// @return SimpleQueryResponse
func (client *Client) SimpleQuery(request *SimpleQueryRequest) (_result *SimpleQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SimpleQueryResponse{}
	_body, _err := client.SimpleQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends a batch processing task.
//
// Description:
//
// You can suspend a batch processing task that is in the Running state. You can call the [ResumeBatch](https://help.aliyun.com/document_detail/479914.html) operation to resume a batch processing task that is suspended.
//
// @param request - SuspendBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendBatchResponse
func (client *Client) SuspendBatchWithOptions(request *SuspendBatchRequest, runtime *dara.RuntimeOptions) (_result *SuspendBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendBatch"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends a batch processing task.
//
// Description:
//
// You can suspend a batch processing task that is in the Running state. You can call the [ResumeBatch](https://help.aliyun.com/document_detail/479914.html) operation to resume a batch processing task that is suspended.
//
// @param request - SuspendBatchRequest
//
// @return SuspendBatchResponse
func (client *Client) SuspendBatch(request *SuspendBatchRequest) (_result *SuspendBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendBatchResponse{}
	_body, _err := client.SuspendBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends a running trigger.
//
// Description:
//
// The operation can be used to suspend a trigger only in the Running state. If you want to resume a suspended trigger, call the [ResumeTrigger](https://help.aliyun.com/document_detail/479919.html) operation.
//
// @param request - SuspendTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendTriggerResponse
func (client *Client) SuspendTriggerWithOptions(request *SuspendTriggerRequest, runtime *dara.RuntimeOptions) (_result *SuspendTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendTrigger"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends a running trigger.
//
// Description:
//
// The operation can be used to suspend a trigger only in the Running state. If you want to resume a suspended trigger, call the [ResumeTrigger](https://help.aliyun.com/document_detail/479919.html) operation.
//
// @param request - SuspendTriggerRequest
//
// @return SuspendTriggerResponse
func (client *Client) SuspendTrigger(request *SuspendTriggerRequest) (_result *SuspendTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendTriggerResponse{}
	_body, _err := client.SuspendTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information about a batch processing task, including the input data source, data processing settings, and tags.
//
// Description:
//
//	  You can update only a batch processing task that is in the Ready or Failed state. The update operation does not change the status of the batch processing task.
//
//		- If you update a batch processing task that is in progress, the task is not automatically resumed after the update is complete. You must call the [ResumeBatch](https://help.aliyun.com/document_detail/479914.html) operation to resume the task.
//
// @param tmpReq - UpdateBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchResponse
func (client *Client) UpdateBatchWithOptions(tmpReq *UpdateBatchRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Actions) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, dara.String("Actions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionsShrink) {
		body["Actions"] = request.ActionsShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InputShrink) {
		body["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBatch"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a batch processing task, including the input data source, data processing settings, and tags.
//
// Description:
//
//	  You can update only a batch processing task that is in the Ready or Failed state. The update operation does not change the status of the batch processing task.
//
//		- If you update a batch processing task that is in progress, the task is not automatically resumed after the update is complete. You must call the [ResumeBatch](https://help.aliyun.com/document_detail/479914.html) operation to resume the task.
//
// @param request - UpdateBatchRequest
//
// @return UpdateBatchResponse
func (client *Client) UpdateBatch(request *UpdateBatchRequest) (_result *UpdateBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBatchResponse{}
	_body, _err := client.UpdateBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Media Set
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - When updating dataset information, make sure the dataset has been successfully created. For creating a dataset, please refer to the request parameter description.
//
// - When updating dataset information, only fill in the fields that need to be updated; unfilled fields will not change.
//
// - The update of the dataset will not take effect immediately and may require up to 5 minutes to become effective.
//
// @param tmpReq - UpdateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithOptions(tmpReq *UpdateDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkflowParameters) {
		request.WorkflowParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkflowParameters, dara.String("WorkflowParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetMaxBindCount) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !dara.IsNil(request.DatasetMaxEntityCount) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !dara.IsNil(request.DatasetMaxFileCount) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !dara.IsNil(request.DatasetMaxRelationCount) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !dara.IsNil(request.DatasetMaxTotalFileSize) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.WorkflowParametersShrink) {
		query["WorkflowParameters"] = request.WorkflowParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataset"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Media Set
//
// Description:
//
// - **Please ensure that you fully understand the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of the Intelligent Media Management product before using this interface.**
//
// - When updating dataset information, make sure the dataset has been successfully created. For creating a dataset, please refer to the request parameter description.
//
// - When updating dataset information, only fill in the fields that need to be updated; unfilled fields will not change.
//
// - The update of the dataset will not take effect immediately and may require up to 5 minutes to become effective.
//
// @param request - UpdateDatasetRequest
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDataset(request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information about a face cluster, such as the cluster name and labels.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have called the [CreateFigureClusteringTask](https://help.aliyun.com/document_detail/478180.html) operation to cluster all faces in the dataset.
//
//		- The operation updates only the cover image, cluster name, and tags.
//
//		- After the operation is successful, you can call the [GetFigureCluster](https://help.aliyun.com/document_detail/478182.html) or [BatchGetFigureCluster](https://help.aliyun.com/document_detail/2248450.html) operation to query the updated cluster.
//
// @param tmpReq - UpdateFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFigureClusterResponse
func (client *Client) UpdateFigureClusterWithOptions(tmpReq *UpdateFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateFigureClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFigureClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FigureCluster) {
		request.FigureClusterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FigureCluster, dara.String("FigureCluster"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FigureClusterShrink) {
		query["FigureCluster"] = request.FigureClusterShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFigureCluster"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a face cluster, such as the cluster name and labels.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have called the [CreateFigureClusteringTask](https://help.aliyun.com/document_detail/478180.html) operation to cluster all faces in the dataset.
//
//		- The operation updates only the cover image, cluster name, and tags.
//
//		- After the operation is successful, you can call the [GetFigureCluster](https://help.aliyun.com/document_detail/478182.html) or [BatchGetFigureCluster](https://help.aliyun.com/document_detail/2248450.html) operation to query the updated cluster.
//
// @param request - UpdateFigureClusterRequest
//
// @return UpdateFigureClusterResponse
func (client *Client) UpdateFigureCluster(request *UpdateFigureClusterRequest) (_result *UpdateFigureClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFigureClusterResponse{}
	_body, _err := client.UpdateFigureClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the partial metadata of the indexed files in a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- You cannot call this operation to update all metadata. You can update only metadata specified by CustomLabels, CustomId, and Figures. For more information, see the "Request parameters" section of this topic.
//
// @param tmpReq - UpdateFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileMetaResponse
func (client *Client) UpdateFileMetaWithOptions(tmpReq *UpdateFileMetaRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.File) {
		request.FileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.File, dara.String("File"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.FileShrink) {
		query["File"] = request.FileShrink
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFileMeta"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the partial metadata of the indexed files in a dataset.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- You cannot call this operation to update all metadata. You can update only metadata specified by CustomLabels, CustomId, and Figures. For more information, see the "Request parameters" section of this topic.
//
// @param request - UpdateFileMetaRequest
//
// @return UpdateFileMetaResponse
func (client *Client) UpdateFileMeta(request *UpdateFileMetaRequest) (_result *UpdateFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFileMetaResponse{}
	_body, _err := client.UpdateFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a spatiotemporal cluster.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have called the [CreateLocationDateClusteringTask](https://help.aliyun.com/document_detail/478188.html) operation to create spatiotemporal clusters in the project.
//
// @param tmpReq - UpdateLocationDateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocationDateClusterResponse
func (client *Client) UpdateLocationDateClusterWithOptions(tmpReq *UpdateLocationDateClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateLocationDateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLocationDateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomLabels) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, dara.String("CustomLabels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomId) {
		query["CustomId"] = request.CustomId
	}

	if !dara.IsNil(request.CustomLabelsShrink) {
		query["CustomLabels"] = request.CustomLabelsShrink
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLocationDateCluster"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLocationDateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a spatiotemporal cluster.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).****
//
//		- Before you call this operation, make sure that you have called the [CreateLocationDateClusteringTask](https://help.aliyun.com/document_detail/478188.html) operation to create spatiotemporal clusters in the project.
//
// @param request - UpdateLocationDateClusterRequest
//
// @return UpdateLocationDateClusterResponse
func (client *Client) UpdateLocationDateCluster(request *UpdateLocationDateClusterRequest) (_result *UpdateLocationDateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLocationDateClusterResponse{}
	_body, _err := client.UpdateLocationDateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information about a project.
//
// Description:
//
//	  Before you call this operation, make sure that the project exists. For information about how to create a project, see "CreateProject".
//
//		- When you call this operation, you need to specify only the parameters that you want to update. The parameters that you do not specify remain unchanged after you call this operation.
//
//		- Wait for up to 5 minutes for the update to take effect.
//
// @param tmpReq - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(tmpReq *UpdateProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetMaxBindCount) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !dara.IsNil(request.DatasetMaxEntityCount) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !dara.IsNil(request.DatasetMaxFileCount) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !dara.IsNil(request.DatasetMaxRelationCount) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !dara.IsNil(request.DatasetMaxTotalFileSize) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ProjectMaxDatasetCount) {
		query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ServiceRole) {
		query["ServiceRole"] = request.ServiceRole
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a project.
//
// Description:
//
//	  Before you call this operation, make sure that the project exists. For information about how to create a project, see "CreateProject".
//
//		- When you call this operation, you need to specify only the parameters that you want to update. The parameters that you do not specify remain unchanged after you call this operation.
//
//		- Wait for up to 5 minutes for the update to take effect.
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a story, such as the story name and cover image.
//
// @param tmpReq - UpdateStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStoryResponse
func (client *Client) UpdateStoryWithOptions(tmpReq *UpdateStoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Cover) {
		request.CoverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cover, dara.String("Cover"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomLabels) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, dara.String("CustomLabels"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CoverShrink) {
		body["Cover"] = request.CoverShrink
	}

	if !dara.IsNil(request.CustomId) {
		body["CustomId"] = request.CustomId
	}

	if !dara.IsNil(request.CustomLabelsShrink) {
		body["CustomLabels"] = request.CustomLabelsShrink
	}

	if !dara.IsNil(request.DatasetName) {
		body["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.StoryName) {
		body["StoryName"] = request.StoryName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a story, such as the story name and cover image.
//
// @param request - UpdateStoryRequest
//
// @return UpdateStoryResponse
func (client *Client) UpdateStory(request *UpdateStoryRequest) (_result *UpdateStoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStoryResponse{}
	_body, _err := client.UpdateStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information about a trigger, such as the input data source, data processing settings, and tags.
//
// Description:
//
//	  You can update only a trigger that is in the Ready or Failed state. The update operation does not change the trigger status.
//
//		- After you update a trigger, the uncompleted tasks under the original trigger are no longer executed. You can call the [ResumeTrigger](https://help.aliyun.com/document_detail/479916.html) operation to resume the execution of the trigger.
//
// @param tmpReq - UpdateTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTriggerResponse
func (client *Client) UpdateTriggerWithOptions(tmpReq *UpdateTriggerRequest, runtime *dara.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTriggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Actions) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, dara.String("Actions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionsShrink) {
		body["Actions"] = request.ActionsShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InputShrink) {
		body["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrigger"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a trigger, such as the input data source, data processing settings, and tags.
//
// Description:
//
//	  You can update only a trigger that is in the Ready or Failed state. The update operation does not change the trigger status.
//
//		- After you update a trigger, the uncompleted tasks under the original trigger are no longer executed. You can call the [ResumeTrigger](https://help.aliyun.com/document_detail/479916.html) operation to resume the execution of the trigger.
//
// @param request - UpdateTriggerRequest
//
// @return UpdateTriggerResponse
func (client *Client) UpdateTrigger(request *UpdateTriggerRequest) (_result *UpdateTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.UpdateTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) contextualAnswerWithSSE_opYieldFunc(_yield chan *ContextualAnswerResponse, _yieldErr chan error, tmpReq *ContextualAnswerRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &ContextualAnswerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Files) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, dara.String("Files"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("Messages"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilesShrink) {
		body["Files"] = request.FilesShrink
	}

	if !dara.IsNil(request.MessagesShrink) {
		body["Messages"] = request.MessagesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContextualAnswer"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
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
