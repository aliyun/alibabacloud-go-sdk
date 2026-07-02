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
func (client *Client) AddImageMosaicWithContext(ctx context.Context, tmpReq *AddImageMosaicRequest, runtime *dara.RuntimeOptions) (_result *AddImageMosaicResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - AddStoryFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddStoryFilesResponse
func (client *Client) AddStoryFilesWithContext(ctx context.Context, tmpReq *AddStoryFilesRequest, runtime *dara.RuntimeOptions) (_result *AddStoryFilesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachOSSBucketResponse
func (client *Client) AttachOSSBucketWithContext(ctx context.Context, request *AttachOSSBucketRequest, runtime *dara.RuntimeOptions) (_result *AttachOSSBucketResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - BatchDeleteFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteFileMetaResponse
func (client *Client) BatchDeleteFileMetaWithContext(ctx context.Context, tmpReq *BatchDeleteFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - BatchGetFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetFigureClusterResponse
func (client *Client) BatchGetFigureClusterWithContext(ctx context.Context, tmpReq *BatchGetFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *BatchGetFigureClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
// - The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, feel free to join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// @param tmpReq - BatchGetFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetFileMetaResponse
func (client *Client) BatchGetFileMetaWithContext(ctx context.Context, tmpReq *BatchGetFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchGetFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation performs a batch index of object metadata by processing input files for tasks such as label detection, face detection, and location detection. The object metadata is then indexed into a dataset to support various data retrieval methods.
//
// Description:
//
// - **Before you use this API, review the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management.**
//
// - For a list of supported data processing tasks, see [Define a workflow](https://help.aliyun.com/document_detail/466304.html).
//
// - The files to be indexed are subject to limits on their total number and size. For more information about dataset limits, see [Limits](https://help.aliyun.com/document_detail/475569.html). For information about how to create a dataset, see the parameter descriptions.
//
// - For information about the regions that support file indexing, see the dataset and index information in [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - BatchIndexFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchIndexFileMetaResponse
func (client *Client) BatchIndexFileMetaWithContext(ctx context.Context, tmpReq *BatchIndexFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchIndexFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - BatchUpdateFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateFileMetaResponse
func (client *Client) BatchUpdateFileMetaWithContext(ctx context.Context, tmpReq *BatchUpdateFileMetaRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CompareImageFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareImageFacesResponse
func (client *Client) CompareImageFacesWithContext(ctx context.Context, tmpReq *CompareImageFacesRequest, runtime *dara.RuntimeOptions) (_result *CompareImageFacesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ContextualAnswerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContextualAnswerResponse
func (client *Client) ContextualAnswerWithSSECtx(ctx context.Context, tmpReq *ContextualAnswerRequest, runtime *dara.RuntimeOptions, _yield chan *ContextualAnswerResponse, _yieldErr chan error) {
	defer close(_yield)
	client.contextualAnswerWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
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
func (client *Client) ContextualAnswerWithContext(ctx context.Context, tmpReq *ContextualAnswerRequest, runtime *dara.RuntimeOptions) (_result *ContextualAnswerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// ### 注意事项
//
// - 请确保在使用该接口前，已充分了解智能媒体管理产品的收费方式和[价格](https://help.aliyun.com/zh/imm/product-overview/billable-items?spm=openapi-amp.newDocPublishment.0.0.1ecd281fi27Zgk)。
//
// - 调用该接口前，请确保您已通过绑定方式（ [CreateBinding](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-createbinding?spm=a2c4g.11186623.0.0.a3d76f44xJrOnF) ）或者主动索引（ [IndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-indexfilemeta?spm=a2c4g.11186623.help-menu-search-62354.d_0) 或者 [BatchIndexFileMeta](https://help.aliyun.com/zh/imm/developer-reference/api-imm-2020-09-30-batchindexfilemeta?spm=a2c4g.11186623.help-menu-62354.d_5_2_4_2_1_1.f1d86f44iBs3QZ) ）方式将文件索引到数据集（Dataset）中。
//
// - 返回结果仅为示例，根据[工作流模板配置](https://help.aliyun.com/zh/imm/user-guide/workflow-templates-and-operators?spm=a2c4g.11186623.0.0.a3d775abr3hDFp)不同，获取到的文件元数据信息的类别和包含的内容均有可能与示例不同。如果有疑问，请使用钉钉搜索钉钉群号 21714099 加入钉钉群进行反馈。
//
// ### 使用限制
//
// - 历史对话长度最长限制为 100，包括用户消息和助手消息。
//
// - 每条消息长度不超过 1000 个汉字。
//
// @param tmpReq - ContextualRetrievalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContextualRetrievalResponse
func (client *Client) ContextualRetrievalWithContext(ctx context.Context, tmpReq *ContextualRetrievalRequest, runtime *dara.RuntimeOptions) (_result *ContextualRetrievalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to inspect a compressed file and retrieve a list of its contents without decompressing the file.
//
// Description:
//
// > This API is in public preview. If you have any questions, join the DingTalk group to provide feedback. For the DingTalk group number, see [Contact us](https://help.aliyun.com/document_detail/84454.html).
//
// - **Before using this API, make sure you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM)**.
//
//	Notice: The completion time of asynchronous tasks is not guaranteed.
//
// - File count limit: A compressed file can contain a maximum of 80,000 files.
//
// - File size limit: The maximum size is 200 GB for ZIP and RAR files, and 50 GB for 7z files.
//
// - This is an asynchronous API. Task information is saved for 7 days after a task starts and is then deleted. To view the task information, call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation with the returned `TaskId`. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through notification messages.
//
// @param tmpReq - CreateArchiveFileInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArchiveFileInspectionTaskResponse
func (client *Client) CreateArchiveFileInspectionTaskWithContext(ctx context.Context, tmpReq *CreateArchiveFileInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateArchiveFileInspectionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a batch processing task that performs specified operations, such as transcoding and format conversion, on multiple existing files.
//
// Description:
//
// If you want to process data using [Object Storage Service (OSS) data processing](https://help.aliyun.com/document_detail/99372.html), make sure you [bind an OSS bucket](https://help.aliyun.com/document_detail/478206.html) before you create a batch processing task.
//
// @param tmpReq - CreateBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchResponse
func (client *Client) CreateBatchWithContext(ctx context.Context, tmpReq *CreateBatchRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/2743997.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// > Asynchronous processing does not guarantee timely task completion.
//
// Before you create a binding, make sure that the project and the dataset that you want to use exist.
//
// - For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
// - For information about how to create a dataset, see [CreateDataset](https://help.aliyun.com/document_detail/478160.html).
//
// > The CreateBinding operation works by using the [workflow template](https://help.aliyun.com/document_detail/466304.html) that is specified when you created the project or dataset.
//
// After you create a binding between a dataset and an OSS bucket, IMM scans the existing objects in the bucket and extracts metadata based on the scanning result. Then, IMM creates an index from the extracted metadata. If new objects are uploaded to the OSS bucket, IMM tracks and scans the objects and updates the index. For objects whose metadata index is created by calling this operation, you can call query operations, such as [SimpleQuery](https://help.aliyun.com/document_detail/478175.html), to query objects, manage objects, and collect statistics on objects.
//
// @param request - CreateBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBindingResponse
func (client *Client) CreateBindingWithContext(ctx context.Context, request *CreateBindingRequest, runtime *dara.RuntimeOptions) (_result *CreateBindingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The point cloud compression feature compresses point cloud data in Object Storage Service (OSS). This helps reduce data transmission over the network.
//
// Description:
//
// - **Before you use this API, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management.**
//
//	Notice: The completion time of asynchronous tasks is not guaranteed.
//
// - File format: Only point cloud files in the PCD format are supported.
//
// - This is an asynchronous API. After a task starts, its information is saved for only 7 days. After this period, the information cannot be retrieved. To view task information, call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation and use the returned `TaskId`. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through notification messages.
//
// @param tmpReq - CreateCompressPointCloudTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCompressPointCloudTaskResponse
func (client *Client) CreateCompressPointCloudTaskWithContext(ctx context.Context, tmpReq *CreateCompressPointCloudTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCompressPointCloudTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateCustomizedStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomizedStoryResponse
func (client *Client) CreateCustomizedStoryWithContext(ctx context.Context, tmpReq *CreateCustomizedStoryRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomizedStoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset.
//
// Description:
//
// - **Before you use this operation, make sure that you fully understand the billing of Intelligent Media Management (IMM) and its [pricing](https://help.aliyun.com/document_detail/477042.html)**.
//
// - Dataset names must be unique within the same project.
//
// - The number of datasets that can be created is limited. You can call [GetProjcet](https://help.aliyun.com/document_detail/478155.html) to query this limit.
//
// - After you create a dataset, you can call [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) to create file metadata indexes for diversified [data retrieval, statistics](https://help.aliyun.com/document_detail/478175.html), and intelligent management.
//
// @param tmpReq - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, tmpReq *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extracts a blind watermark.
//
// Description:
//
// - Before you use this API, make sure that you understand the billing methods and pricing of Intelligent Media Management (IMM).
//
//	Notice: Asynchronous tasks are not guaranteed to be completed within a specific time frame.
//
// - Make sure that a project is created in IMM. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
//
// - Make sure the service region and project are the same as those used to add the blind watermark using the [EncodeBlindWatermark](https://help.aliyun.com/document_detail/2743655.html) operation. Otherwise, the watermark cannot be extracted.
//
// - The watermark can be extracted even after the image undergoes attacks such as compression, scaling, clipping, and color changes.
//
// - This API is compatible with the previous version of the blind watermarking feature. Some parameters are from the previous DecodeBlindWatermark API.
//
// - This is an asynchronous API. After a task starts, its information is saved for only 7 days. After this period, the information can no longer be retrieved. Call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) API to retrieve the TaskId and view task information. Alternatively, set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through asynchronous notification messages.
//
// @param tmpReq - CreateDecodeBlindWatermarkTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDecodeBlindWatermarkTaskResponse
func (client *Client) CreateDecodeBlindWatermarkTaskWithContext(ctx context.Context, tmpReq *CreateDecodeBlindWatermarkTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDecodeBlindWatermarkTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches a media set for the top N images most similar to a specified image or face ID. The operation returns the corresponding face IDs and bounding boxes, sorted by similarity in descending order.
//
// Description:
//
// - **Before you use this operation, review the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html)**
//
//		Notice:
//
//	The execution time of asynchronous tasks is not guaranteed.
//
// - For each input image, only the face with the largest bounding box is used for the face search.
//
// - This is an asynchronous operation. After a task starts, the task information is retained for 7 days and cannot be retrieved after this period. To retrieve task information, you can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation with the returned `TaskId`. Alternatively, you can configure the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive asynchronous notifications that contain task information.
//
// @param tmpReq - CreateFacesSearchingTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFacesSearchingTaskResponse
func (client *Client) CreateFacesSearchingTaskWithContext(ctx context.Context, tmpReq *CreateFacesSearchingTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFacesSearchingTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a figure clustering task. This task uses an intelligent algorithm to group the faces of different people in images that are indexed in a dataset.
//
// Description:
//
// - **Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management.**
//
//	Notice: The completion time of asynchronous tasks is not guaranteed.
//
// - Before you call this operation, make sure that you have indexed files to a dataset ([CreateDataset](~~CreateDataset~~)) by attaching them ([CreateBinding](~~CreateBinding~~)) or by indexing them ([IndexFileMeta](~~IndexFileMeta~~) or [BatchIndexFileMeta](~~BatchIndexFileMeta~~)).
//
// - Each time you call this operation, files in the dataset ([CreateDataset](~~CreateDataset~~)) are incrementally processed. You can periodically call this operation to process new files.
//
// - After the clustering is complete, you can call the [GetFigureCluster](~~GetFigureCluster~~) or [BatchGetFigureCluster](~~BatchGetFigureCluster~~) operation to retrieve information about specific groups. You can also call [QueryFigureClusters](~~QueryFigureClusters~~) to query and list the groups in the dataset.
//
// - Deleting files from a dataset changes the face clustering results. When all images that contain the faces in a cluster are deleted, the cluster is also deleted.
//
// - This is an asynchronous operation. After a task starts, its information is saved for only 7 days. You cannot retrieve the task information after this period. You can call the [GetTask](~~GetTask~~) or [ListTasks](~~ListTasks~~) operation to view the task information. Alternatively, you can set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information from asynchronous notification messages.
//
// @param tmpReq - CreateFigureClusteringTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFigureClusteringTaskResponse
func (client *Client) CreateFigureClusteringTaskWithContext(ctx context.Context, tmpReq *CreateFigureClusteringTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFigureClusteringTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Merges two or more figure clustering groups into a single figure clustering group.
//
// Description:
//
// - **Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management.**
//
// - Before you call this operation, make sure that you have clustered all faces in the dataset by calling the [CreateFigureClusteringTask](https://help.aliyun.com/document_detail/478180.html) operation.
//
// - Merging unrelated groups affects the feature values of the destination group. This may cause inaccurate grouping of incremental data when you create a figure clustering task.
//
// - This operation is asynchronous. Task information is retained for only 7 days. During this period, you can query task information by calling the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation with the returned `TaskId`. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive asynchronous notification messages about the task.
//
// @param tmpReq - CreateFigureClustersMergingTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFigureClustersMergingTaskResponse
func (client *Client) CreateFigureClustersMergingTaskWithContext(ctx context.Context, tmpReq *CreateFigureClustersMergingTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFigureClustersMergingTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Package Download API
//
// Description:
//
// > This API is in public preview. If you have any questions, join our DingTalk group to provide feedback. For the group number, see [Contact us](https://help.aliyun.com/document_detail/84454.html).
//
// > This API currently supports packaging but not compression. The compression feature will be added later.
//
// - **Before using this API, make sure you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM)**.
//
//	Notice: The completion time of asynchronous tasks is not guaranteed.
//
// - File count limit: You can package up to 80,000 files.
//
// - File size limit: The total size of all files before packaging must not exceed 200 GB.
//
// - This feature supports files of the Standard storage class on OSS. To package files of other storage classes, first [convert their storage class](https://help.aliyun.com/document_detail/90090.html).
//
// - This is an asynchronous API. After a task starts, its information is stored for 7 days. After 7 days, the information can no longer be retrieved. To view task information, call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation with the returned `TaskId`. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through asynchronous notification messages.
//
// @param tmpReq - CreateFileCompressionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileCompressionTaskResponse
func (client *Client) CreateFileCompressionTaskWithContext(ctx context.Context, tmpReq *CreateFileCompressionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFileCompressionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A file decompression task lets you decompress specific files or an entire compressed package to a specified location. Supported formats include Zip, RAR, and 7z.
//
// Description:
//
// > This API is in public preview. If you have any questions, join our DingTalk group to provide feedback. For the group number, see [Contact us](https://help.aliyun.com/document_detail/84454.html).
//
// - **Before you use this API, review the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) for Intelligent Media Management.**
//
//	Notice: Timeliness is not guaranteed for asynchronous tasks.
//
// - File count limit: A compressed package can contain a maximum of 80,000 files.
//
// - File size limit: 200 GB for Zip and RAR formats, and 50 GB for 7z format.
//
// - File decompression tasks use stream decompression, which outputs files as they are decompressed. If an operation is aborted due to file corruption, the files that have already been decompressed are not deleted.
//
// - This is an asynchronous API. Task information is stored for only 7 days and cannot be retrieved after this period. To view the task information, you can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation with the returned `TaskId`. Alternatively, you can set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through an asynchronous notification message.
//
// @param tmpReq - CreateFileUncompressionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileUncompressionTaskResponse
func (client *Client) CreateFileUncompressionTaskWithContext(ctx context.Context, tmpReq *CreateFileUncompressionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateFileUncompressionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a video highlight task. This feature is in invitational preview.
//
// Description:
//
// - **Before you use this operation, make sure that you fully understand the billing of Intelligent Media Management (IMM).*	- For more information, see [Billing overview](https://www.alibabacloud.com/help/en/imm/product-overview/billing-overview). This operation incurs fees for highlight extraction and media processing.
//
// - Before you call this operation, make sure that a project already exists in the current region. For more information, see [Project management](https://www.alibabacloud.com/help/en/imm/developer-reference/api-imm-2020-09-30-createproject).
//
//	Notice: Asynchronous tasks do not guarantee timeliness..
//
// @param tmpReq - CreateHighlightTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHighlightTaskResponse
func (client *Client) CreateHighlightTaskWithContext(ctx context.Context, tmpReq *CreateHighlightTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateHighlightTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateHighlightTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CredentialConfig) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, dara.String("CredentialConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Edit) {
		request.EditShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Edit, dara.String("Edit"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Highlight) {
		request.HighlightShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Highlight, dara.String("Highlight"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CredentialConfigShrink) {
		body["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.EditShrink) {
		body["Edit"] = request.EditShrink
	}

	if !dara.IsNil(request.HighlightShrink) {
		body["Highlight"] = request.HighlightShrink
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.NotificationShrink) {
		body["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.OutputShrink) {
		body["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.SourcesShrink) {
		body["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHighlightTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHighlightTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects non-compliant content in images, such as pornography, terrorism, undesirable scenes, logos, and text-in-image violations.
//
// Description:
//
// - **Before you use this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management.**
//
//		Notice:
//
//	The execution time of asynchronous tasks is not guaranteed.
//
// - Image requirements:
//
//   - Image URLs support the HTTP and HTTPS protocols.
//
//   - The following image formats are supported: PNG, JPG, JPEG, BMP, GIF, and WEBP.
//
//   - The image size cannot exceed 20 MB for both synchronous and asynchronous invocations. The height or width cannot exceed 30,000 pixels, and the total number of pixels cannot exceed 250 million. For GIF images, the total number of pixels cannot exceed 4,194,304, and the height or width cannot exceed 30,000 pixels.
//
//   - The image download timeout period is 3 seconds. If the download takes longer than 3 seconds, a timeout error is returned.
//
//   - For best results, the image resolution should be at least 256 × 256 pixels. Low resolution may affect detection accuracy.
//
//   - The response time for image detection depends on the image download time. Ensure the storage service where the image is stored is stable and reliable. Use Alibaba Cloud Object Storage Service (OSS) or CDN.
//
// - This is an asynchronous operation. After a task starts, its information is saved for only 7 days. You cannot query the information after this period. To view task information, you can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation and use the returned `TaskId`. Alternatively, you can set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through asynchronous notification messages.
//
// > The detection results are returned in an asynchronous notification message. The Suggestion field in the message has one of the following values:
//
// >
//
// > - pass: The image passed the review. No non-compliant content was detected.
//
// >
//
// > - block: The image failed the review. Non-compliant content was detected. The Categories field indicates the non-compliant category. For more information about the categories, see Content Moderation detection results.
//
// >
//
// > - review: The image requires manual review. After the manual review is complete, another asynchronous notification message is sent to inform you of the result.
//
// @param tmpReq - CreateImageModerationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageModerationTaskResponse
func (client *Client) CreateImageModerationTaskWithContext(ctx context.Context, tmpReq *CreateImageModerationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageModerationTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stitches multiple images into a single image based on specified rules and saves the output to a specified OSS object.
//
// Description:
//
// - **Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management.**
//
// - Before you call this operation, ensure that an active project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
// - You can stitch a maximum of 10 images in this operation. The length of a single edge of each image cannot exceed 32,876 pixels. The total number of pixels cannot exceed 1 billion.
//
// - This is an asynchronous operation. After a task starts, its information is saved for 7 days. After this period, you can no longer query the task information. To query task information, call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation and use the returned `TaskId`. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive asynchronous notifications about the task.
//
// @param tmpReq - CreateImageSplicingTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageSplicingTaskResponse
func (client *Client) CreateImageSplicingTaskWithContext(ctx context.Context, tmpReq *CreateImageSplicingTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageSplicingTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts multiple images into a single PDF file and saves the file as a specified OSS object.
//
// Description:
//
// - **Before using this API, make sure you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management.**
//
// - Before calling this API, make sure that an active project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
// - This API supports up to 100 input images.
//
// - This is an asynchronous API. After a task starts, its information is stored for only 7 days and cannot be retrieved after this period. To view task information, call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) API with the returned `TaskId`. You can also receive task information through asynchronous notification messages by setting the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter.
//
// @param tmpReq - CreateImageToPDFTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageToPDFTaskResponse
func (client *Client) CreateImageToPDFTaskWithContext(ctx context.Context, tmpReq *CreateImageToPDFTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImageToPDFTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The spatio-temporal clustering feature classifies files in a dataset based on their time and location. This feature works on indexed files, such as images and videos, that contain shooting time and location data. These classifications can represent content from a user\\"s trip, where files have similar timestamps and locations. The classifications can also represent content shot at different places where a user lives or works. Analyzing the locations and time ranges of these classifications lets you categorize media files, create highlight reels, and generate photo and video stories.
//
// Description:
//
// - **Before you use this operation, you must understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM)**.
//
//	Notice: Asynchronous tasks do not have a guaranteed processing time.
//
// - Before you call this operation, you must index files into a dataset. You can index files by binding data sources using [CreateBinding](https://help.aliyun.com/document_detail/478202.html) or by indexing files using [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html).
//
// - Each call to this operation processes the files in the specified `Dataset` **incrementally**. You can call this operation periodically to process new files.
//
// - After clustering is complete, you can call the [QueryLocationDateClusters](https://help.aliyun.com/document_detail/478189.html) operation to retrieve the clustering results.
//
// - Deleting a file from a dataset does not change the spatio-temporal clusters. To delete existing spatio-temporal clusters, you can call the [DeleteLocationDateCluster](https://help.aliyun.com/document_detail/478191.html) operation.
//
// - This is an asynchronous operation. After a task starts, its information is saved for only 7 days. You cannot retrieve task information after 7 days. You can call the [GetTask](~~GetTask~~) or [ListTasks](~~ListTasks~~) operation to view task information using the returned `TaskId`. You can also configure the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through message notifications.
//
// @param tmpReq - CreateLocationDateClusteringTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLocationDateClusteringTaskResponse
func (client *Client) CreateLocationDateClusteringTaskWithContext(ctx context.Context, tmpReq *CreateLocationDateClusteringTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateLocationDateClusteringTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an asynchronous media transcoding task. This task processes audio and video files for media transcoding, media concatenation, video frame capture, and animated GIF generation.
//
// Description:
//
// - **Before you call this operation, ensure you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/88317.html) for Intelligent Media Management.**
//
// - Before calling this operation, ensure a project is available in the current region. For more information, see [Project Management](https://help.aliyun.com/document_detail/478152.html).
//
//		Notice:
//
//	The completion time of an asynchronous task is not guaranteed.
//
// - When you use this operation for media transcoding, it processes only one video, audio, or subtitle stream by default. You can also configure the number of streams to process.
//
// - When you use this operation for media concatenation, you can specify a maximum of 11 media files. Parameters for operations such as media transcoding and frame capture apply to the final concatenated output.
//
// - This operation is asynchronous. After a task starts, its information is retained for only 7 days. After this period, you cannot retrieve it. To view task information, call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation with the returned `TaskId`. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information via message notifications.
//
// @param tmpReq - CreateMediaConvertTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMediaConvertTaskResponse
func (client *Client) CreateMediaConvertTaskWithContext(ctx context.Context, tmpReq *CreateMediaConvertTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMediaConvertTaskResponse, _err error) {
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

	if !dara.IsNil(tmpReq.TargetGroups) {
		request.TargetGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetGroups, dara.String("TargetGroups"), dara.String("json"))
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

	if !dara.IsNil(request.TargetGroupsShrink) {
		query["TargetGroups"] = request.TargetGroupsShrink
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a document conversion task that converts documents, such as Word, PowerPoint, Excel, and PDF files, stored in Object Storage Service (OSS) into images, text files, or PDF files.
//
// Description:
//
// - **Before you use this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//	Notice: The execution time of asynchronous tasks is not guaranteed.
//
// - Supported input file formats:
//
//   - Word processor documents (Word): doc, docx, wps, wpss, docm, dotm, dot, and dotx.
//
//   - Presentation documents (PowerPoint): pptx, ppt, pot, potx, pps, ppsx, dps, dpt, pptm, potm, ppsm, and dpss.
//
//   - Spreadsheet documents (Excel): xls, xlt, et, ett, xlsx, xltx, csv, xlsb, xlsm, xltm, and ets.
//
//   - PDF documents: pdf.
//
// - Supported output file formats:
//
//   - Images: png and jpg.
//
//   - Text: txt.
//
//   - PDF: pdf.
//
// - The maximum size of a single file is 200 MB. This limit cannot be changed.
//
// - If a file is large or its content is complex, the conversion may time out.
//
// - The number of requests per second is limited to 50 for a single user.
//
// - Task information is stored for only 7 days after a task starts. After this period, the information cannot be retrieved. You can promptly obtain task information using one of the following methods:
//
//   - You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to obtain the returned `TaskId` and view the task information.
//
//   - You can activate Message Service (MNS) in the same region as IMM and configure a subscription to promptly receive task information notifications. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html). For more information about the MNS software development kit (SDK), see [Receive and delete messages](https://help.aliyun.com/document_detail/32449.html).
//
//   - You can activate RocketMQ in the same region as IMM, and create a RocketMQ 4.0 instance, a topic, and a group to promptly receive task information notifications. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html). For more information about how to use RocketMQ, see [Use an SDK for HTTP to send and receive normal messages](https://help.aliyun.com/document_detail/169009.html).
//
//   - You can activate and connect to [EventBridge](https://www.aliyun.com/product/aliware/eventbridge) in the same region as IMM to promptly receive task information notifications. For more information, see [Intelligent Media Management IMM events](https://help.aliyun.com/document_detail/205730.html).
//
// @param tmpReq - CreateOfficeConversionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOfficeConversionTaskResponse
func (client *Client) CreateOfficeConversionTaskWithContext(ctx context.Context, tmpReq *CreateOfficeConversionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateOfficeConversionTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Project names must be unique within the same region.
//
//   - The number of projects that can be created is limited. By default, you can create up to 100 projects. To increase the quota, submit a ticket or search for DingTalk group 88490020073 to join the group and submit a request.
//
//   - After you create a project, you can create other Intelligent Media Management (IMM) resources:
//
//   - [Create a dataset](https://help.aliyun.com/document_detail/478160.html)
//
//   - [Create a trigger](https://help.aliyun.com/document_detail/479912.html)
//
//   - [Create a batch task](https://help.aliyun.com/document_detail/606694.html)
//
//   - [Create a binding task](https://help.aliyun.com/document_detail/478202.html)
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, tmpReq *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The similar image clustering feature groups images that you have indexed in a dataset into clusters based on visual similarity. This feature is useful for scenarios such as deduplicating images or selecting the best shots. For example, you can use it to filter burst photos in an album.
//
// Description:
//
// - **Before calling this operation, review the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
//	Notice: The execution time of asynchronous tasks is not guaranteed.
//
// - Before calling this operation, index files to a dataset. You can index files by attaching a data source using [CreateBinding](https://help.aliyun.com/document_detail/478202.html), or by actively indexing files using [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html).
//
// - Each call to this operation **incrementally*	- processes the files in the specified `Dataset`. You can call this operation periodically to process new files.
//
// - After clustering completes, call the [QuerySimilarImageClusters](https://help.aliyun.com/document_detail/611304.html) operation to retrieve the clustering results.
//
// - Each similar image cluster must contain at least two images. Deleting a file from a dataset changes the similar image clusters. If deleting an image reduces a cluster to fewer than two images, the cluster is automatically deleted.
//
// - This operation is asynchronous. After a task starts, its information is retained for only seven days. You cannot query the information after this period. Call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation and use the returned `TaskId` to view task information. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive asynchronous notification messages about the task.
//
// @param tmpReq - CreateSimilarImageClusteringTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimilarImageClusteringTaskResponse
func (client *Client) CreateSimilarImageClusteringTaskWithContext(ctx context.Context, tmpReq *CreateSimilarImageClusteringTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSimilarImageClusteringTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - **Before calling this operation, understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
// - Before calling this operation, index files to a dataset by calling [CreateBinding](https://help.aliyun.com/document_detail/478202.html), [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html), or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html).
//
// - This is an asynchronous operation. After a task starts, its information is saved for only 7 days. The information cannot be retrieved after this period. Call [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) with the returned TaskId to view task information. Alternatively, set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to obtain task information from notification messages.
//
// @param tmpReq - CreateStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoryResponse
func (client *Client) CreateStoryWithContext(ctx context.Context, tmpReq *CreateStoryRequest, runtime *dara.RuntimeOptions) (_result *CreateStoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trigger to start data processing in Intelligent Media Management (IMM). The trigger is activated by event sources, such as Object Storage Service (OSS), and uses data processing templates to process media files, such as images, videos, and documents.
//
// Description:
//
// To process data from [Object Storage Service](https://help.aliyun.com/document_detail/99372.html), ensure that you have [attached an OSS bucket](https://help.aliyun.com/document_detail/478206.html).
//
// @param tmpReq - CreateTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTriggerResponse
func (client *Client) CreateTriggerWithContext(ctx context.Context, tmpReq *CreateTriggerRequest, runtime *dara.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects labels for scenarios, objects, and events in video content. This feature supports more than 30 categories and thousands of labels. Scenario labels include natural landscapes, life scenes, and disaster scenes. Event labels include talent shows, office work, performances, and production. Object labels include tableware, electronic products, furniture, and vehicles.
//
// Description:
//
// - **Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/2747104.html) of Intelligent Media Management.**
//
// - Before you call this operation, make sure that you have created a project in Intelligent Media Management. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
//
//		Notice:
//
//	The completion time of asynchronous tasks is not guaranteed.
//
// - For more information about the features of this operation, see [Video label detection](https://help.aliyun.com/document_detail/477189.html).
//
// - This operation supports multiple video formats, such as MP4, MPEG-TS, MKV, MOV, AVI, FLV, and M3U8.
//
// - This is an asynchronous operation. After a task starts, its information is stored for seven days. You cannot retrieve the information after this period. Call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation with the returned `TaskId` to view task information. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through message notifications.
//
// @param tmpReq - CreateVideoLabelClassificationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoLabelClassificationTaskResponse
func (client *Client) CreateVideoLabelClassificationTaskWithContext(ctx context.Context, tmpReq *CreateVideoLabelClassificationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVideoLabelClassificationTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects threats or non-compliant content in videos. This operation can be used in scenarios such as pornography detection, terrorism and politically sensitive content detection, text and image violation detection, undesirable scene detection, and logo detection.
//
// Description:
//
// - **Before you use this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management.**
//
//		Notice:
//
//	The completion time of asynchronous tasks is not guaranteed.
//
// - The detection results are returned in an asynchronous notification message. The Suggestion field in the asynchronous notification message can have the following values:
//
//   - pass: The video passed the review. No non-compliant content was detected.
//
//   - block: The video must be blocked. This value is returned when non-compliant content is detected. The Categories field indicates the category of the non-compliant content. For more information about the categories, see [Content Moderation detection results](https://help.aliyun.com/document_detail/2743995.html).
//
//   - review: The video requires manual review. After the manual review is complete, another asynchronous notification message is sent with the result.
//
// - Video snapshot requirements:
//
//   - Video frame URLs support the HTTP and HTTPS protocols.
//
//   - Supported video frame formats: PNG, JPG, JPEG, BMP, GIF, and WEBP.
//
//   - The size of a video frame cannot exceed 10 MB.
//
//   - The recommended resolution for video frames is at least 256 × 256 pixels. A lower resolution may affect detection accuracy.
//
//   - The response time for the video detection operation depends on the download time of the video frames. Make sure that the storage service for your video frames is stable and reliable. We recommend that you use Alibaba Cloud Object Storage Service (OSS) or cache frames with Alibaba Cloud CDN.
//
// - This is an asynchronous operation. After a task is created, the task information is saved for only 7 days. After this period, the information cannot be retrieved. You can call the [GetTask](https://help.aliyun.com/document_detail/478241.html) or [ListTasks](https://help.aliyun.com/document_detail/478242.html) operation to query the task information using the returned `TaskId`. You can also set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive task information through asynchronous notification messages.
//
// @param tmpReq - CreateVideoModerationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoModerationTaskResponse
func (client *Client) CreateVideoModerationTaskWithContext(ctx context.Context, tmpReq *CreateVideoModerationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVideoModerationTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBatchResponse
func (client *Client) DeleteBatchWithContext(ctx context.Context, request *DeleteBatchRequest, runtime *dara.RuntimeOptions) (_result *DeleteBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - If you delete a binding, new changes in the OSS bucket are not synchronized to the dataset. Exercise caution when you perform this operation.
//
// @param request - DeleteBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBindingResponse
func (client *Client) DeleteBindingWithContext(ctx context.Context, request *DeleteBindingRequest, runtime *dara.RuntimeOptions) (_result *DeleteBindingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithContext(ctx context.Context, request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileMetaResponse
func (client *Client) DeleteFileMetaWithContext(ctx context.Context, request *DeleteFileMetaRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLocationDateClusterResponse
func (client *Client) DeleteLocationDateClusterWithContext(ctx context.Context, request *DeleteLocationDateClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteLocationDateClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you delete a project, make sure that all resources in the project, such as datasets, bindings, batch processing tasks, and triggers, are deleted. For more information, see [DeleteDataset](https://help.aliyun.com/document_detail/478164.html), [DeleteBatch](https://help.aliyun.com/document_detail/479918.html), and [DeleteTrigger](https://help.aliyun.com/document_detail/479915.html).
//
// - After a project is deleted, all resources used by the project are recycled, and all related data is lost and cannot be recovered.
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStoryResponse
func (client *Client) DeleteStoryWithContext(ctx context.Context, request *DeleteStoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteStoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTriggerResponse
func (client *Client) DeleteTriggerWithContext(ctx context.Context, request *DeleteTriggerRequest, runtime *dara.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachOSSBucketResponse
func (client *Client) DetachOSSBucketWithContext(ctx context.Context, request *DetachOSSBucketRequest, runtime *dara.RuntimeOptions) (_result *DetachOSSBucketResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DetectImageBodiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageBodiesResponse
func (client *Client) DetectImageBodiesWithContext(ctx context.Context, tmpReq *DetectImageBodiesRequest, runtime *dara.RuntimeOptions) (_result *DetectImageBodiesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DetectImageCarsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageCarsResponse
func (client *Client) DetectImageCarsWithContext(ctx context.Context, tmpReq *DetectImageCarsRequest, runtime *dara.RuntimeOptions) (_result *DetectImageCarsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DetectImageCodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageCodesResponse
func (client *Client) DetectImageCodesWithContext(ctx context.Context, tmpReq *DetectImageCodesRequest, runtime *dara.RuntimeOptions) (_result *DetectImageCodesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects visually optimal cropping regions in an image at a specified aspect ratio by using AI model capabilities.
//
// @param tmpReq - DetectImageCroppingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageCroppingResponse
func (client *Client) DetectImageCroppingWithContext(ctx context.Context, tmpReq *DetectImageCroppingRequest, runtime *dara.RuntimeOptions) (_result *DetectImageCroppingResponse, _err error) {
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

	if !dara.IsNil(tmpReq.InclusionHints) {
		request.InclusionHintsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InclusionHints, dara.String("InclusionHints"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AspectRatios) {
		query["AspectRatios"] = request.AspectRatios
	}

	if !dara.IsNil(request.CredentialConfigShrink) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !dara.IsNil(request.InclusionHintsShrink) {
		query["InclusionHints"] = request.InclusionHintsShrink
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DetectImageFacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageFacesResponse
func (client *Client) DetectImageFacesWithContext(ctx context.Context, tmpReq *DetectImageFacesRequest, runtime *dara.RuntimeOptions) (_result *DetectImageFacesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - Make sure that an IMM [project](https://help.aliyun.com/document_detail/478273.html) is created. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
// - For more information about the features of this operation, see [Image label detection](https://help.aliyun.com/document_detail/477179.html).
//
// - For more information about the input images supported by this operation, see [Limits on images](https://help.aliyun.com/document_detail/475569.html).
//
// @param tmpReq - DetectImageLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageLabelsResponse
func (client *Client) DetectImageLabelsWithContext(ctx context.Context, tmpReq *DetectImageLabelsRequest, runtime *dara.RuntimeOptions) (_result *DetectImageLabelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DetectImageScoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageScoreResponse
func (client *Client) DetectImageScoreWithContext(ctx context.Context, tmpReq *DetectImageScoreRequest, runtime *dara.RuntimeOptions) (_result *DetectImageScoreResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DetectImageTextsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectImageTextsResponse
func (client *Client) DetectImageTextsWithContext(ctx context.Context, tmpReq *DetectImageTextsRequest, runtime *dara.RuntimeOptions) (_result *DetectImageTextsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DetectMediaMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectMediaMetaResponse
func (client *Client) DetectMediaMetaWithContext(ctx context.Context, tmpReq *DetectMediaMetaRequest, runtime *dara.RuntimeOptions) (_result *DetectMediaMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectTextAnomalyResponse
func (client *Client) DetectTextAnomalyWithContext(ctx context.Context, request *DetectTextAnomalyRequest, runtime *dara.RuntimeOptions) (_result *DetectTextAnomalyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EncodeBlindWatermarkResponse
func (client *Client) EncodeBlindWatermarkWithContext(ctx context.Context, request *EncodeBlindWatermarkRequest, runtime *dara.RuntimeOptions) (_result *EncodeBlindWatermarkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ExtractDocumentTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtractDocumentTextResponse
func (client *Client) ExtractDocumentTextWithContext(ctx context.Context, tmpReq *ExtractDocumentTextRequest, runtime *dara.RuntimeOptions) (_result *ExtractDocumentTextResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the extracted file metadata, including the file name, labels, path, custom tags, and other fields. If the value of a metadata field of a file matches the specified string, the metadata of the file is returned.
//
// Description:
//
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of IMM.\\*\\*\\*\\*
//
// - Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
// - The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 88490020073) and share your questions with us.
//
// - For information about the fields that you can use as query conditions, see [Supported fields and operators](https://help.aliyun.com/document_detail/2743991.html).
//
// @param tmpReq - FuzzyQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FuzzyQueryResponse
func (client *Client) FuzzyQueryWithContext(ctx context.Context, tmpReq *FuzzyQueryRequest, runtime *dara.RuntimeOptions) (_result *FuzzyQueryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a playlist from a video file for live transcoding. The output is an M3U8 file that enables immediate playback and on-demand transcoding based on playback progress. Compared with offline transcoding, this method significantly reduces transcoding wait times and lowers transcoding and storage overhead.
//
// Description:
//
// - **Before you use this API, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/88317.html) of Intelligent Media Management.**
//
// - Before you call this API, ensure that an active project exists in the current region. For more information, see [Project management](https://help.aliyun.com/document_detail/478152.html).
//
// - By default, this API processes only one video, audio, or subtitle stream. You can also configure the number of video, audio, and subtitle streams to process.
//
//		Notice:
//
//	The Video, Audio, and Subtitle parameters within Targets cannot all be empty. If a parameter is left empty, the corresponding processing is disabled. For example, if the Video parameter is empty, video processing is disabled, and the output TS file does not contain a video stream.
//
// - The source video must be at least 0.x seconds long. The minimum duration varies slightly based on the output frame rate.
//
// - This API supports generating both Media playlists and Master playlists. For more information, see the parameter descriptions in this document.
//
// - This is a synchronous API. Transcoding is triggered only during playback or pre-transcoding. You can set the [Notification](https://help.aliyun.com/document_detail/2743997.html) parameter to receive the transcoding task result through an asynchronous notification message.
//
// - For more information about the features of this API, see [Live transcoding](https://help.aliyun.com/document_detail/477192.html).
//
// - The data processing feature of OSS can also be used to generate playlists. However, this feature generates only Media Playlists and uses simplified parameters. For more information, see the OSS data processing topic [Generate a playlist](https://help.aliyun.com/document_detail/2709281.html).
//
// @param tmpReq - GenerateVideoPlaylistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateVideoPlaylistResponse
func (client *Client) GenerateVideoPlaylistWithContext(ctx context.Context, tmpReq *GenerateVideoPlaylistRequest, runtime *dara.RuntimeOptions) (_result *GenerateVideoPlaylistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the preview and editing credentials for a document.
//
// Description:
//
// - **Before you use this operation, make sure that you are familiar with the billing of Intelligent Media Management. For more information, see [Pricing](https://help.aliyun.com/document_detail/477042.html)**.
//
// - Do not perform cross-border access on OSS files. For example, if a file is stored in a bucket in the Singapore region, do not initiate preview, read, or download requests from the Chinese mainland. In such scenarios, the network link quality is significantly affected by the cross-border network environment, which may cause increased access latency, preview failures, download interruptions, or unstable connections. Network stability and access experience cannot be guaranteed. Make sure that the access point and the bucket are in the same region to avoid uncertainties caused by cross-border access.
//
// - The access credential expires in 30 minutes, and the refresh credential expires in 1 day.
//
// - The returned expiration time is in UTC, which is 8 hours behind UTC+8.
//
// - Supported input file formats:
//
//   - Word documents: doc, docx, txt, dot, wps, wpt, dotx, docm, dotm, and rtf.
//
//   - Presentation documents (PPT): ppt, pptx, pptm, ppsx, ppsm, pps, potx, potm, dpt, and dps.
//
//   - Excel documents: et, xls, xlt, xlsx, xlsm, xltx, xltm, and csv.
//
//   - PDF documents: pdf.
//
// - The maximum supported file size is 200 MB.
//
// - The maximum supported number of document pages is 5,000.
//
// - For projects created before December 1, 2023, billing is based on the number of document opens. Currently, billing is based on the number of API calls. To switch to the new billing mode, create a new project. Note that each API call can be used by only one user. If the call is reused, only the last user can access the document normally, and the access permissions of other users are revoked.
//
// - Activate Message Service (MNS) in the same region as Intelligent Media Management, create a topic and a queue, and configure a subscription. You can pass in the MNS topic name by using the NotifyTopicName parameter to receive message notifications about file saves. For more information about the MNS SDK, see [Receive and delete messages](https://help.aliyun.com/document_detail/32449.html).
//
// For an example of the JSON format of the Message field in file save message notifications, see [WebOffice message notification format](https://help.aliyun.com/document_detail/2743999.html).
//
// > To use the versioning feature, you must first enable versioning in OSS and then set the History parameter to true.
//
// >.
//
// @param tmpReq - GenerateWebofficeTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateWebofficeTokenResponse
func (client *Client) GenerateWebofficeTokenWithContext(ctx context.Context, tmpReq *GenerateWebofficeTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateWebofficeTokenResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchResponse
func (client *Client) GetBatchWithContext(ctx context.Context, request *GetBatchRequest, runtime *dara.RuntimeOptions) (_result *GetBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
// - Make sure that the binding relationship that you want to query exists. For information about how to create a binding relationship, see [CreateBinding](https://help.aliyun.com/document_detail/478202.html).
//
// @param request - GetBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBindingResponse
func (client *Client) GetBindingWithContext(ctx context.Context, request *GetBindingRequest, runtime *dara.RuntimeOptions) (_result *GetBindingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDRMLicenseResponse
func (client *Client) GetDRMLicenseWithContext(ctx context.Context, request *GetDRMLicenseRequest, runtime *dara.RuntimeOptions) (_result *GetDRMLicenseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - The GetDataset operation supports real-time retrieval of file statistics. You can specify WithStatistics to enable real-time retrieval of file statistics.
//
// @param request - GetDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithContext(ctx context.Context, request *GetDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDecodeBlindWatermarkResultResponse
func (client *Client) GetDecodeBlindWatermarkResultWithContext(ctx context.Context, request *GetDecodeBlindWatermarkResultRequest, runtime *dara.RuntimeOptions) (_result *GetDecodeBlindWatermarkResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - **Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
// - Before you call this operation, make sure that a face clustering task is created to group all faces in a dataset. For information about how to create a face clustering task, see [CreateFigureClusteringTask](~~CreateFigureClusteringTask~~). For information about how to create a dataset, see [CreateDataset](~~CreateDataset~~).
//
// @param request - GetFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFigureClusterResponse
func (client *Client) GetFigureClusterWithContext(ctx context.Context, request *GetFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *GetFigureClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
// - The sample response is provided for reference only. The metadata type and content in your response may differ based on factors such as the [workflow template configurations](https://help.aliyun.com/document_detail/466304.html). For any inquiries, join the DingTalk chat group (ID: 31690030817) and share your questions with us.
//
// @param tmpReq - GetFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileMetaResponse
func (client *Client) GetFileMetaWithContext(ctx context.Context, tmpReq *GetFileMetaRequest, runtime *dara.RuntimeOptions) (_result *GetFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an image content moderation task.
//
// @param request - GetImageModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageModerationResultResponse
func (client *Client) GetImageModerationResultWithContext(ctx context.Context, request *GetImageModerationResultRequest, runtime *dara.RuntimeOptions) (_result *GetImageModerationResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the project name of a project that is bound to a specified OSS bucket in the same region.
//
// Description:
//
// - **Before you use this operation, make sure that you are familiar with the billing of Intelligent Media Management (IMM) and its [pricing](https://help.aliyun.com/document_detail/477042.html).**
//
// - Before you call this operation, make sure that you have called the [AttachOSSBucket](https://help.aliyun.com/document_detail/478206.html) operation to bind a project to an OSS bucket.
//
// @param request - GetOSSBucketAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSBucketAttachmentResponse
func (client *Client) GetOSSBucketAttachmentWithContext(ctx context.Context, request *GetOSSBucketAttachmentRequest, runtime *dara.RuntimeOptions) (_result *GetOSSBucketAttachmentResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specified project, including basic information and statistics information related to datasets and files.
//
// Description:
//
// Querying project information supports real-time retrieval of file statistics information. You can enable this feature through parameter settings. For details, see the request parameters section.
//
//	Notice: Only files in datasets created before December 20, 2025 can be counted.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
// - Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param request - GetStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStoryResponse
func (client *Client) GetStoryWithContext(ctx context.Context, request *GetStoryRequest, runtime *dara.RuntimeOptions) (_result *GetStoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specified asynchronous task. Intelligent Media Management (IMM) supports various asynchronous data processing capabilities, each with its own task creation operation, such as CreateFigureClusteringTask for creating figure clustering tasks and CreateFileCompressionTask for creating file compression tasks. This operation is a general-purpose operation that allows you to query the details of an asynchronous task by task ID and type.
//
// Description:
//
// *Before you use this operation, make sure that you are familiar with the billing of Intelligent Media Management and its [pricing](https://help.aliyun.com/document_detail/477042.html).**.
//
// @param request - GetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, request *GetTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTriggerResponse
func (client *Client) GetTriggerWithContext(ctx context.Context, request *GetTriggerRequest, runtime *dara.RuntimeOptions) (_result *GetTriggerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a video label detection task.
//
// Description:
//
// - Before you call this operation, make sure that you have created a project ([Project](https://help.aliyun.com/document_detail/478273.html)) in Intelligent Media Management. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
//
// - Before you call this operation, make sure that you have created a [video label detection task](https://help.aliyun.com/document_detail/478223.html) and obtained the `TaskId` of the task.
//
// @param request - GetVideoLabelClassificationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoLabelClassificationResultResponse
func (client *Client) GetVideoLabelClassificationResultWithContext(ctx context.Context, request *GetVideoLabelClassificationResultRequest, runtime *dara.RuntimeOptions) (_result *GetVideoLabelClassificationResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a video content moderation task.
//
// @param request - GetVideoModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoModerationResultResponse
func (client *Client) GetVideoModerationResultWithContext(ctx context.Context, request *GetVideoModerationResultRequest, runtime *dara.RuntimeOptions) (_result *GetVideoModerationResultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs data processing on input files for tasks such as label detection, face detection, and location detection. This operation extracts object metadata and creates an index, which lets you retrieve data from a dataset.
//
// Description:
//
// - **Make sure you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management before you use this API.**
//
// - For a list of supported data processing operations for indexing object metadata, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
//
// - The total number and size of files that can be indexed are limited. For more information, see the Dataset limits section in [Limits](https://help.aliyun.com/document_detail/475569.html). For information about how to create a dataset, see the parameter descriptions.
//
// - For a list of regions where you can index object metadata, see the \\"Features supported by region, Datasets and indexes\\" section in [Limits](https://help.aliyun.com/document_detail/475569.html).
//
// - After you index object metadata, you can retrieve data using [Simple query](https://help.aliyun.com/document_detail/478175.html). For information about other retrieval features, see [Query and statistics](https://help.aliyun.com/document_detail/2402363.html). You can also create face groups using [Create a face clustering task](https://help.aliyun.com/document_detail/478180.html). For information about other clustering features, see [Intelligent management](https://help.aliyun.com/document_detail/2402365.html).
//
// > 	- This is an asynchronous operation. After you submit a request, the file is processed. The processing time can range from several seconds to several minutes or longer, depending on the [workflow template and operators](https://help.aliyun.com/document_detail/466304.html) and file content. After the processing is complete, the metadata is stored in the dataset. You can use the [message subscription](https://help.aliyun.com/document_detail/603317.html) feature to receive a notification when the task is complete.
//
// @param tmpReq - IndexFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IndexFileMetaResponse
func (client *Client) IndexFileMetaWithContext(ctx context.Context, tmpReq *IndexFileMetaRequest, runtime *dara.RuntimeOptions) (_result *IndexFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAttachedOSSBucketsResponse
func (client *Client) ListAttachedOSSBucketsWithContext(ctx context.Context, request *ListAttachedOSSBucketsRequest, runtime *dara.RuntimeOptions) (_result *ListAttachedOSSBucketsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBatchesResponse
func (client *Client) ListBatchesWithContext(ctx context.Context, request *ListBatchesRequest, runtime *dara.RuntimeOptions) (_result *ListBatchesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of bindings between datasets and Object Storage Service (OSS) buckets.
//
// Description:
//
// *Before you use this operation, make sure that you are familiar with the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).**
//
// @param request - ListBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindingsResponse
func (client *Client) ListBindingsWithContext(ctx context.Context, request *ListBindingsRequest, runtime *dara.RuntimeOptions) (_result *ListBindingsResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithContext(ctx context.Context, request *ListDatasetsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all projects, including basic information and statistics information related to datasets and files.
//
// Description:
//
// Paging is supported for viewing returned data. When performing a paged query for the first page, set only MaxResults to limit the number of returned entries. The NextToken value in the response serves as the credential for querying subsequent pages. When querying subsequent pages, set the NextToken parameter to the NextToken value obtained from the previous response as the query credential, and set MaxResults to limit the number of returned entries.
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithContext(ctx context.Context, request *ListRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, tmpReq *ListTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTriggersResponse
func (client *Client) ListTriggersWithContext(ctx context.Context, request *ListTriggersRequest, runtime *dara.RuntimeOptions) (_result *ListTriggersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - Before you call this operation, make sure that a face clustering task is created to group all faces in a dataset. For information about how to create a face clustering task, see [CreateFigureClusteringTask](~~CreateFigureClusteringTask~~). For information about how to create a dataset, see [CreateDataset](~~CreateDataset~~).
//
// @param tmpReq - QueryFigureClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFigureClustersResponse
func (client *Client) QueryFigureClustersWithContext(ctx context.Context, tmpReq *QueryFigureClustersRequest, runtime *dara.RuntimeOptions) (_result *QueryFigureClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - QueryLocationDateClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLocationDateClustersResponse
func (client *Client) QueryLocationDateClustersWithContext(ctx context.Context, tmpReq *QueryLocationDateClustersRequest, runtime *dara.RuntimeOptions) (_result *QueryLocationDateClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySimilarImageClustersResponse
func (client *Client) QuerySimilarImageClustersWithContext(ctx context.Context, request *QuerySimilarImageClustersRequest, runtime *dara.RuntimeOptions) (_result *QuerySimilarImageClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM).\\*\\*\\*\\*
//
// - Before you call this operation, make sure that you have indexed file metadata into the dataset automatically by calling the [CreateBinding](https://help.aliyun.com/document_detail/478202.html) operation or manually by calling the [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) or [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) operation.
//
// - Before you call this operation, make sure that you have called the [CreateStory](https://help.aliyun.com/document_detail/478193.html) or [CreateCustomizedStory](https://help.aliyun.com/document_detail/478196.html) operation to create a story.
//
// @param tmpReq - QueryStoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStoriesResponse
func (client *Client) QueryStoriesWithContext(ctx context.Context, tmpReq *QueryStoriesRequest, runtime *dara.RuntimeOptions) (_result *QueryStoriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes a Weboffice access token. A Weboffice access token is valid for 30 minutes. After it expires, you can no longer access Weboffice. To continue accessing Weboffice, call this operation to refresh the Weboffice access token and obtain a new token that is also valid for 30 minutes.
//
// Description:
//
// *Make sure that you are familiar with the billing method and [pricing](https://help.aliyun.com/document_detail/477042.html) of Intelligent Media Management (IMM) before you invoke this operation.**
//
// - For billing details, refer to [WebOffice billing](https://help.aliyun.com/document_detail/2639703.html).
//
// - The access token expires in 30 minutes. Open the preview before the access token expires. After the token expires, previewing is no longer available.
//
// - The refresh token expires in 1 day. Invoke the refresh operation before the refresh token expires. After the token expires, it becomes invalid.
//
// - The returned expiration time is in UTC, which is 8 hours behind UTC+8.
//
// > The access token is used for actual preview session access. The refresh token simplifies the parameter settings required for refreshing tokens. You can use the refresh token to directly obtain a new token with the previously configured settings.
//
// >
//
// @param tmpReq - RefreshWebofficeTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshWebofficeTokenResponse
func (client *Client) RefreshWebofficeTokenWithContext(ctx context.Context, tmpReq *RefreshWebofficeTokenRequest, runtime *dara.RuntimeOptions) (_result *RefreshWebofficeTokenResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - RemoveStoryFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveStoryFilesResponse
func (client *Client) RemoveStoryFilesWithContext(ctx context.Context, tmpReq *RemoveStoryFilesRequest, runtime *dara.RuntimeOptions) (_result *RemoveStoryFilesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeBatchResponse
func (client *Client) ResumeBatchWithContext(ctx context.Context, request *ResumeBatchRequest, runtime *dara.RuntimeOptions) (_result *ResumeBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTriggerResponse
func (client *Client) ResumeTriggerWithContext(ctx context.Context, request *ResumeTriggerRequest, runtime *dara.RuntimeOptions) (_result *ResumeTriggerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - SearchImageFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageFigureClusterResponse
func (client *Client) SearchImageFigureClusterWithContext(ctx context.Context, tmpReq *SearchImageFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *SearchImageFigureClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// ### 注意事项
//
// - **请确保在使用该接口前，已充分了解智能媒体管理产品的收费方式和[价格](https://help.aliyun.com/document_detail/477042.html)**。该接口每次请求，会产生语义理解费用和查询费用两种计费项各一次。
//
// - 调用该接口前，请确保您已通过绑定方式（ [CreateBinding](https://help.aliyun.com/document_detail/478202.html) ）或者主动索引（ [IndexFileMeta](https://help.aliyun.com/document_detail/478166.html) 或者 [BatchIndexFileMeta](https://help.aliyun.com/document_detail/478167.html) ）方式将文件索引到数据集（Dataset）中。
//
// - 返回结果仅为示例，根据[工作流模板配置](https://help.aliyun.com/document_detail/466304.html)不同，获取到的文件元数据信息的类别和包含的内容均有可能与示例不同。如果有疑问，请加入钉钉群进行反馈，钉钉群号请参见[联系我们](https://help.aliyun.com/document_detail/84454.html)。
//
// ### 使用限制
//
// - 每次查询最多返回 100 个文件信息。
//
// - 不支持翻页查询。
//
// - 自然语言理解不保证完全准确。
//
// - 该功能在美国（硅谷），美国（弗吉尼亚）地域下不支持。
//
// ### 使用方式
//
// 使用自然语言关键词对数据集内的文件进行搜索查询。目前支持理解的关键信息包括标签（Labels.LabelName）、时间（ProduceTime）和地点（Address.AddressLine）等。例如，以`2023 年杭州的风景`为条件进行查询，会被智能拆分为如下三个条件，并查找出同时满足这些条件的文件：
//
// - ProduceTime：2023 年 1 月 1 日零点起到 2023 年 12 月 31 日结束止
//
// - Address.AddressLine：包含`杭州`关键词
//
// - Labels.LabelName：包含`风景`标签
//
// 配合[工作流模板配置](https://help.aliyun.com/document_detail/466304.html)，当模板中包含`ImageEmbeddingExtraction`算子时，该搜索请求会提供基于图片内容的搜索，即您输入的`Query`内容会同时被理解为图片内包含的内容，从而实现对图片的智能检索。
//
// @param tmpReq - SemanticQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SemanticQueryResponse
func (client *Client) SemanticQueryWithContext(ctx context.Context, tmpReq *SemanticQueryRequest, runtime *dara.RuntimeOptions) (_result *SemanticQueryResponse, _err error) {
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

	if !dara.IsNil(request.SourceURI) {
		query["SourceURI"] = request.SourceURI
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - SimpleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SimpleQueryResponse
func (client *Client) SimpleQueryWithContext(ctx context.Context, tmpReq *SimpleQueryRequest, runtime *dara.RuntimeOptions) (_result *SimpleQueryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendBatchResponse
func (client *Client) SuspendBatchWithContext(ctx context.Context, request *SuspendBatchRequest, runtime *dara.RuntimeOptions) (_result *SuspendBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendTriggerResponse
func (client *Client) SuspendTriggerWithContext(ctx context.Context, request *SuspendTriggerRequest, runtime *dara.RuntimeOptions) (_result *SuspendTriggerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information for a batch processing task, such as the data source configuration, data processing configuration, and tags.
//
// Description:
//
// - You can update a batch processing task only when its status is Ready or Failed. The update does not change the current status of the task.
//
// - After the update, an incomplete batch processing task does not automatically resume. To resume the task, call the [ResumeBatch](https://help.aliyun.com/document_detail/479914.html) operation.
//
// @param tmpReq - UpdateBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchResponse
func (client *Client) UpdateBatchWithContext(ctx context.Context, tmpReq *UpdateBatchRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a dataset.
//
// Description:
//
// - **Before you use this operation, make sure that you fully understand the billing of Intelligent Media Management (IMM) and its [pricing](https://help.aliyun.com/document_detail/477042.html)**.
//
// - Before you update dataset information, make sure that the dataset has been created. To create a dataset, refer to the request parameter descriptions.
//
// - When you update dataset information, specify only the fields that you want to update. Fields that are not specified remain unchanged.
//
// - After a dataset is updated, the changes may take up to 5 minutes to take effect.
//
// @param tmpReq - UpdateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithContext(ctx context.Context, tmpReq *UpdateDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateFigureClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFigureClusterResponse
func (client *Client) UpdateFigureClusterWithContext(ctx context.Context, tmpReq *UpdateFigureClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateFigureClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateFileMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileMetaResponse
func (client *Client) UpdateFileMetaWithContext(ctx context.Context, tmpReq *UpdateFileMetaRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileMetaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateLocationDateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocationDateClusterResponse
func (client *Client) UpdateLocationDateClusterWithContext(ctx context.Context, tmpReq *UpdateLocationDateClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateLocationDateClusterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a created project.
//
// Description:
//
// - Before updating project information, make sure the project has been created. To create a project, refer to the request parameter descriptions.
//
// - When updating project information, specify only the fields that you want to update. Fields that are not specified remain unchanged.
//
// - Project updates do not take effect immediately. Wait up to 5 minutes for the updates to take effect.
//
// @param tmpReq - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithContext(ctx context.Context, tmpReq *UpdateProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateStoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStoryResponse
func (client *Client) UpdateStoryWithContext(ctx context.Context, tmpReq *UpdateStoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateStoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - UpdateTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTriggerResponse
func (client *Client) UpdateTriggerWithContext(ctx context.Context, tmpReq *UpdateTriggerRequest, runtime *dara.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) contextualAnswerWithSSECtx_opYieldFunc(_yield chan *ContextualAnswerResponse, _yieldErr chan error, ctx context.Context, tmpReq *ContextualAnswerRequest, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
