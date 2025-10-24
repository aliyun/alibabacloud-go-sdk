// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Activates a media workflow.
//
// Description:
//
// You can call this operation to activate a media workflow that has been deactivated. After you activate a media workflow, you cannot modify the workflow information, such as the name, topology, or trigger mode. A media workflow is activated by default after it is created.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ActivateMediaWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateMediaWorkflowResponse
func (client *Client) ActivateMediaWorkflowWithContext(ctx context.Context, request *ActivateMediaWorkflowRequest, runtime *dara.RuntimeOptions) (_result *ActivateMediaWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaWorkflowId) {
		query["MediaWorkflowId"] = request.MediaWorkflowId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateMediaWorkflow"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateMediaWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a media file.
//
// Description:
//
//	  You can call this operation to process videos that are uploaded to Object Storage Service (OSS) but not processed. This way, you do not need to upload the videos to OSS again. If you have configured media workflows, OSS automatically notifies ApsaraVideo Media Processing (MPS) when a media file is uploaded to OSS. MPS automatically finds the corresponding workflow in the Active state based on the specified OSS bucket and object. Therefore, in most cases, you do not need to manually call the AddMedia operation to process the media file.
//
//		- Media information is automatically obtained only when the specified media workflow is in the Active state. If no media workflow is specified or the specified media workflow is not in the Active state, media information is not obtained.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - AddMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMediaResponse
func (client *Client) AddMediaWithContext(ctx context.Context, request *AddMediaRequest, runtime *dara.RuntimeOptions) (_result *AddMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileURL) {
		query["FileURL"] = request.FileURL
	}

	if !dara.IsNil(request.InputUnbind) {
		query["InputUnbind"] = request.InputUnbind
	}

	if !dara.IsNil(request.MediaWorkflowId) {
		query["MediaWorkflowId"] = request.MediaWorkflowId
	}

	if !dara.IsNil(request.MediaWorkflowUserData) {
		query["MediaWorkflowUserData"] = request.MediaWorkflowUserData
	}

	if !dara.IsNil(request.OverrideParams) {
		query["OverrideParams"] = request.OverrideParams
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMedia"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a tag to a media file.
//
// Description:
//
// You can call this operation to add only one tag. To add multiple tags at a time, you can call the [UpdateMedia](https://help.aliyun.com/document_detail/44464.html) operation.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - AddMediaTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMediaTagResponse
func (client *Client) AddMediaTagWithContext(ctx context.Context, request *AddMediaTagRequest, runtime *dara.RuntimeOptions) (_result *AddMediaTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMediaTag"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMediaTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a media workflow.
//
// Description:
//
//	  You can call this operation to define the topology, activities, and dependencies of a media workflow. The topology is represented by a directed acyclic graph (DAG) in the console. For more information, see [Workflow activities](https://help.aliyun.com/document_detail/68494.html). You can view and run the workflows that are created by calling this operation in the ApsaraVideo Media Processing (MPS) console.
//
//		- MPS media workflows can be automatically triggered only by using the prefix of the file path. Automatic triggering by using the suffix is not supported. For more information about the trigger rules, see [Workflow triggering rules for files](https://help.aliyun.com/document_detail/68574.html).
//
// ### [](#qps)QPS limits
//
// You can call this API operation up to 100 times per second per account. Requests that exceed this limit are dropped, and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - AddMediaWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMediaWorkflowResponse
func (client *Client) AddMediaWorkflowWithContext(ctx context.Context, request *AddMediaWorkflowRequest, runtime *dara.RuntimeOptions) (_result *AddMediaWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Topology) {
		query["Topology"] = request.Topology
	}

	if !dara.IsNil(request.TriggerMode) {
		query["TriggerMode"] = request.TriggerMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMediaWorkflow"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMediaWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - AddPipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPipelineResponse
func (client *Client) AddPipelineWithContext(ctx context.Context, request *AddPipelineRequest, runtime *dara.RuntimeOptions) (_result *AddPipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NotifyConfig) {
		query["NotifyConfig"] = request.NotifyConfig
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	if !dara.IsNil(request.SpeedLevel) {
		query["SpeedLevel"] = request.SpeedLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPipeline"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加labelVersion、knowledgeConfig配置
//
// @param request - AddSmarttagTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmarttagTemplateResponse
func (client *Client) AddSmarttagTemplateWithContext(ctx context.Context, request *AddSmarttagTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddSmarttagTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalyseTypes) {
		query["AnalyseTypes"] = request.AnalyseTypes
	}

	if !dara.IsNil(request.FaceCategoryIds) {
		query["FaceCategoryIds"] = request.FaceCategoryIds
	}

	if !dara.IsNil(request.FaceCustomParamsConfig) {
		query["FaceCustomParamsConfig"] = request.FaceCustomParamsConfig
	}

	if !dara.IsNil(request.Industry) {
		query["Industry"] = request.Industry
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.KeywordConfig) {
		query["KeywordConfig"] = request.KeywordConfig
	}

	if !dara.IsNil(request.KnowledgeConfig) {
		query["KnowledgeConfig"] = request.KnowledgeConfig
	}

	if !dara.IsNil(request.LabelCustomCategoryIds) {
		query["LabelCustomCategoryIds"] = request.LabelCustomCategoryIds
	}

	if !dara.IsNil(request.LabelCustomParamsConfig) {
		query["LabelCustomParamsConfig"] = request.LabelCustomParamsConfig
	}

	if !dara.IsNil(request.LabelType) {
		query["LabelType"] = request.LabelType
	}

	if !dara.IsNil(request.LabelVersion) {
		query["LabelVersion"] = request.LabelVersion
	}

	if !dara.IsNil(request.LandmarkGroupIds) {
		query["LandmarkGroupIds"] = request.LandmarkGroupIds
	}

	if !dara.IsNil(request.ObjectGroupIds) {
		query["ObjectGroupIds"] = request.ObjectGroupIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSmarttagTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSmarttagTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom transcoding template. You need to configure the information such as the container format, video stream settings, and audio stream settings.
//
// Description:
//
// When you call this operation, you need to set transcoding parameters such as those related to the container format, video stream, and audio stream. If you do not specify some parameters, streams that are generated by using the template do not contain the information specified by those parameters.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - AddTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTemplateResponse
func (client *Client) AddTemplateWithContext(ctx context.Context, request *AddTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddTemplateResponse, _err error) {
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

	if !dara.IsNil(request.Container) {
		query["Container"] = request.Container
	}

	if !dara.IsNil(request.MuxConfig) {
		query["MuxConfig"] = request.MuxConfig
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransConfig) {
		query["TransConfig"] = request.TransConfig
	}

	if !dara.IsNil(request.Video) {
		query["Video"] = request.Video
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a watermark template.
//
// Description:
//
// After you create a watermark template by calling this operation, you can specify the watermark template and watermark asset when you [submit a transcoding job](https://help.aliyun.com/document_detail/29226.html). This allows you to add watermark information to the output video.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - AddWaterMarkTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWaterMarkTemplateResponse
func (client *Client) AddWaterMarkTemplateWithContext(ctx context.Context, request *AddWaterMarkTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddWaterMarkTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWaterMarkTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWaterMarkTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds an input media bucket.
//
// Description:
//
// Before you call this operation to bind an input media bucket, you must create a media bucket. For more information, see [Add media buckets](https://help.aliyun.com/document_detail/42430.html).
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - BindInputBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindInputBucketResponse
func (client *Client) BindInputBucketWithContext(ctx context.Context, request *BindInputBucketRequest, runtime *dara.RuntimeOptions) (_result *BindInputBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Referer) {
		query["Referer"] = request.Referer
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindInputBucket"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindInputBucketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds an output media bucket to the media library.
//
// Description:
//
// Before you call this operation to bind an output media bucket to the media library, you must create a media bucket. For more information, see [Add media buckets](https://help.aliyun.com/document_detail/42430.html).
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - BindOutputBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindOutputBucketResponse
func (client *Client) BindOutputBucketWithContext(ctx context.Context, request *BindOutputBucketRequest, runtime *dara.RuntimeOptions) (_result *BindOutputBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindOutputBucket"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindOutputBucketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a transcoding job.
//
// Description:
//
//	  You can cancel a transcoding job only if the job is in the Submitted state.
//
//		- We recommend that you call the **UpdatePipeline*	- operation to set the status of the ApsaraVideo Media Processing (MPS) queue to Paused before you cancel a job. This suspends job scheduling in the MPS queue. After the job is canceled, you must set the status of the MPS queue back to Active so that the other jobs in the MPS queue can be scheduled.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - CancelJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelJobResponse
func (client *Client) CancelJobWithContext(ctx context.Context, request *CancelJobRequest, runtime *dara.RuntimeOptions) (_result *CancelJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateCustomEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomEntityResponse
func (client *Client) CreateCustomEntityWithContext(ctx context.Context, request *CreateCustomEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomEntityInfo) {
		query["CustomEntityInfo"] = request.CustomEntityInfo
	}

	if !dara.IsNil(request.CustomEntityName) {
		query["CustomEntityName"] = request.CustomEntityName
	}

	if !dara.IsNil(request.CustomGroupId) {
		query["CustomGroupId"] = request.CustomGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomEntity"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateCustomGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomGroupResponse
func (client *Client) CreateCustomGroupWithContext(ctx context.Context, request *CreateCustomGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomGroupDescription) {
		query["CustomGroupDescription"] = request.CustomGroupDescription
	}

	if !dara.IsNil(request.CustomGroupName) {
		query["CustomGroupName"] = request.CustomGroupName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomGroup"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a job of creating a media fingerprint library.
//
// Description:
//
//	  You can call this operation to submit a job to create a video or text fingerprint library. You can use a text fingerprint library to store fingerprints for text.
//
//		- You can submit a job of creating a text fingerprint library only in the China (Shanghai) region.
//
//		- By default, you can submit up to 10 jobs of creating a video fingerprint library to an ApsaraVideo Media Processing (MPS) queue at a time. If you submit more than 10 jobs at a time, the call may fail.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - CreateFpShotDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFpShotDBResponse
func (client *Client) CreateFpShotDBWithContext(ctx context.Context, request *CreateFpShotDBRequest, runtime *dara.RuntimeOptions) (_result *CreateFpShotDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFpShotDB"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFpShotDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a media workflow.
//
// Description:
//
// The time when the media workflow was created.
//
// @param request - DeactivateMediaWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateMediaWorkflowResponse
func (client *Client) DeactivateMediaWorkflowWithContext(ctx context.Context, request *DeactivateMediaWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DeactivateMediaWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaWorkflowId) {
		query["MediaWorkflowId"] = request.MediaWorkflowId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactivateMediaWorkflow"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactivateMediaWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom entity from a custom library.
//
// @param request - DeleteCustomEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomEntityResponse
func (client *Client) DeleteCustomEntityWithContext(ctx context.Context, request *DeleteCustomEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomEntityId) {
		query["CustomEntityId"] = request.CustomEntityId
	}

	if !dara.IsNil(request.CustomGroupId) {
		query["CustomGroupId"] = request.CustomGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomEntity"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a custom image library.
//
// Description:
//
// You can call this operation only in the China (Beijing), China (Shanghai), and China (Hangzhou) regions.
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeleteCustomGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomGroupResponse
func (client *Client) DeleteCustomGroupWithContext(ctx context.Context, request *DeleteCustomGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomGroupId) {
		query["CustomGroupId"] = request.CustomGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomGroup"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomViewResponse
func (client *Client) DeleteCustomViewWithContext(ctx context.Context, request *DeleteCustomViewRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomEntityId) {
		query["CustomEntityId"] = request.CustomEntityId
	}

	if !dara.IsNil(request.CustomGroupId) {
		query["CustomGroupId"] = request.CustomGroupId
	}

	if !dara.IsNil(request.CustomViewId) {
		query["CustomViewId"] = request.CustomViewId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomView"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes media files from ApsaraVideo Media Processing (MPS).
//
// Description:
//
// This operation allows you to logically delete a media file. The media file can no longer be processed, but the corresponding objects in the input and output Object Storage Service (OSS) buckets are retained.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeleteMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaResponse
func (client *Client) DeleteMediaWithContext(ctx context.Context, request *DeleteMediaRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMedia"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a tag for a media file.
//
// Description:
//
// You can call this operation to remove only one tag at a time.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeleteMediaTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaTagResponse
func (client *Client) DeleteMediaTagWithContext(ctx context.Context, request *DeleteMediaTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaTag"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a media workflow. This does not affect workflow instances that are running.
//
// Description:
//
// After you delete or disable a workflow, the workflow cannot be used. In this case, the workflow is not automatically triggered when you upload a file to the bucket specified by the workflow.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeleteMediaWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaWorkflowResponse
func (client *Client) DeleteMediaWorkflowWithContext(ctx context.Context, request *DeleteMediaWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaWorkflowId) {
		query["MediaWorkflowId"] = request.MediaWorkflowId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaWorkflow"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ApsaraVideo Media Processing (MPS) queue.
//
// Description:
//
// You can call this operation to delete only one MPS queue at a time.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeletePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineResponse
func (client *Client) DeletePipelineWithContext(ctx context.Context, request *DeletePipelineRequest, runtime *dara.RuntimeOptions) (_result *DeletePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePipeline"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// Description:
//
// You can call this operation to delete only one template at a time.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped, and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeleteSmarttagTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmarttagTemplateResponse
func (client *Client) DeleteSmarttagTemplateWithContext(ctx context.Context, request *DeleteSmarttagTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmarttagTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSmarttagTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmarttagTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom transcoding template.
//
// Description:
//
// A custom transcoding template cannot be deleted if it is being used by a job that has been submitted.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithContext(ctx context.Context, request *DeleteTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a watermark template.
//
// Description:
//
// A watermark template cannot be deleted if it is being used by a submitted job.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - DeleteWaterMarkTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWaterMarkTemplateResponse
func (client *Client) DeleteWaterMarkTemplateWithContext(ctx context.Context, request *DeleteWaterMarkTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteWaterMarkTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WaterMarkTemplateId) {
		query["WaterMarkTemplateId"] = request.WaterMarkTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWaterMarkTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWaterMarkTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reviews images and text and returns the review results.
//
// Description:
//
//	  The moderation results are synchronously returned after the moderation is complete.
//
//		- You can use the image and text moderation feature only in the China (Beijing), China (Shanghai), and Singapore regions.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ImAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImAuditResponse
func (client *Client) ImAuditWithContext(ctx context.Context, request *ImAuditRequest, runtime *dara.RuntimeOptions) (_result *ImAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Contents) {
		query["Contents"] = request.Contents
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scenes) {
		query["Scenes"] = request.Scenes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImAudit"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a job of importing text files to a text fingerprint library.
//
// Description:
//
//	  You can call this operation to import multiple text files to a text fingerprint library at a time. The system extracts fingerprints from the text files and saves the fingerprints to the text fingerprint library.
//
//		- You can call this operation only in the China (Shanghai) region.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ImportFpShotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportFpShotJobResponse
func (client *Client) ImportFpShotJobWithContext(ctx context.Context, request *ImportFpShotJobRequest, runtime *dara.RuntimeOptions) (_result *ImportFpShotJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FpDBId) {
		query["FpDBId"] = request.FpDBId
	}

	if !dara.IsNil(request.FpImportConfig) {
		query["FpImportConfig"] = request.FpImportConfig
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportFpShotJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportFpShotJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all media buckets bound to the media library.
//
// Description:
//
// A maximum of 100 media buckets can be returned.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ListAllMediaBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllMediaBucketResponse
func (client *Client) ListAllMediaBucketWithContext(ctx context.Context, request *ListAllMediaBucketRequest, runtime *dara.RuntimeOptions) (_result *ListAllMediaBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllMediaBucket"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllMediaBucketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of entities in a custom library.
//
// @param request - ListCustomEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomEntitiesResponse
func (client *Client) ListCustomEntitiesWithContext(ctx context.Context, request *ListCustomEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListCustomEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomGroupId) {
		query["CustomGroupId"] = request.CustomGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomEntities"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCustomGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomGroupsResponse
func (client *Client) ListCustomGroupsWithContext(ctx context.Context, request *ListCustomGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomGroups"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all figures and faces in a specific figure library.
//
// Description:
//
// You can specify the ID of a figure or a figure library to query the corresponding information. If neither the figure ID nor figure library ID is specified, the operation returns the information about all figures and faces in all figure libraries within the current RAM user.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ListCustomPersonsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomPersonsResponse
func (client *Client) ListCustomPersonsWithContext(ctx context.Context, request *ListCustomPersonsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomPersonsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomPersons"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomPersonsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCustomViewsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomViewsResponse
func (client *Client) ListCustomViewsWithContext(ctx context.Context, request *ListCustomViewsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomViewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomEntityId) {
		query["CustomEntityId"] = request.CustomEntityId
	}

	if !dara.IsNil(request.CustomGroupId) {
		query["CustomGroupId"] = request.CustomGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomViews"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomViewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media fingerprint libraries.
//
// Description:
//
//	  You can call this operation to query the status and information about the media fingerprint libraries based on the specified IDs.
//
//		- You can query text fingerprint libraries only in the China (Shanghai) region.
//
//		- You can call this operation to query up to 10 media fingerprint libraries.
//
// ### QPS limit
//
// You can call this operation up to 500 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ListFpShotDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFpShotDBResponse
func (client *Client) ListFpShotDBWithContext(ctx context.Context, request *ListFpShotDBRequest, runtime *dara.RuntimeOptions) (_result *ListFpShotDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FpDBIds) {
		query["FpDBIds"] = request.FpDBIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFpShotDB"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFpShotDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media files in a media fingerprint library.
//
// Description:
//
//	  You can call this operation to query media files in a specific media fingerprint library based on the library ID. This operation supports paged queries.
//
//		- You can call this operation only in the China (Beijing), China (Hangzhou), China (Shanghai), and Singapore regions.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ListFpShotFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFpShotFilesResponse
func (client *Client) ListFpShotFilesWithContext(ctx context.Context, request *ListFpShotFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFpShotFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FpDBId) {
		query["FpDBId"] = request.FpDBId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFpShotFiles"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFpShotFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries jobs of importing text files to a text fingerprint library.
//
// Description:
//
// You can call this operation only in the China (Shanghai) region.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ListFpShotImportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFpShotImportJobResponse
func (client *Client) ListFpShotImportJobWithContext(ctx context.Context, request *ListFpShotImportJobRequest, runtime *dara.RuntimeOptions) (_result *ListFpShotImportJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFpShotImportJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFpShotImportJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Traverses transcoding jobs
//
// Description:
//
//	  By default, the returned transcoding jobs are sorted by CreationTime in descending order.
//
//		- You can call this operation to return transcoding jobs of the last 90 days. The jobs are returned based on the specified time range.
//
//		- You can filter query results by configuring request parameters such as job status, creation time interval, and ApsaraVideo Media Processing (MPS) queue for transcoding.
//
//		- By default, MPS does not allow you to access data across regions within the same account. Before you call this operation, make sure that the region that you specify is the same as the region of the transcoding jobs to be queried. Otherwise, this operation may fail to be called, or invalid information may be returned.
//
// ### [](#qps)QPS limits
//
// You can call this API operation up to 100 times per second per account. Requests that exceed this limit are dropped, and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ListJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobResponse
func (client *Client) ListJobWithContext(ctx context.Context, request *ListJobRequest, runtime *dara.RuntimeOptions) (_result *ListJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfJobCreatedTimeRange) {
		query["EndOfJobCreatedTimeRange"] = request.EndOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartOfJobCreatedTimeRange) {
		query["StartOfJobCreatedTimeRange"] = request.StartOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution instances of a media workflow.
//
// Description:
//
// This operation returns execution instances only in the recent 90 days.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - ListMediaWorkflowExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaWorkflowExecutionsResponse
func (client *Client) ListMediaWorkflowExecutionsWithContext(ctx context.Context, request *ListMediaWorkflowExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListMediaWorkflowExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputFileURL) {
		query["InputFileURL"] = request.InputFileURL
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.MediaWorkflowId) {
		query["MediaWorkflowId"] = request.MediaWorkflowId
	}

	if !dara.IsNil(request.MediaWorkflowName) {
		query["MediaWorkflowName"] = request.MediaWorkflowName
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaWorkflowExecutions"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaWorkflowExecutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the template analysis job and returns a list of available preset templates when the template analysis job is complete.
//
// Description:
//
// The time when the job was created.
//
// @param request - QueryAnalysisJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAnalysisJobListResponse
func (client *Client) QueryAnalysisJobListWithContext(ctx context.Context, request *QueryAnalysisJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryAnalysisJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisJobIds) {
		query["AnalysisJobIds"] = request.AnalysisJobIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAnalysisJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAnalysisJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询版权水印提取任务
//
// @param request - QueryCopyrightExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCopyrightExtractJobResponse
func (client *Client) QueryCopyrightExtractJobWithContext(ctx context.Context, request *QueryCopyrightExtractJobRequest, runtime *dara.RuntimeOptions) (_result *QueryCopyrightExtractJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCopyrightExtractJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCopyrightExtractJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频版权水印任务
//
// @param request - QueryCopyrightJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCopyrightJobResponse
func (client *Client) QueryCopyrightJobWithContext(ctx context.Context, request *QueryCopyrightJobRequest, runtime *dara.RuntimeOptions) (_result *QueryCopyrightJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCopyrightJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCopyrightJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the jobs of clearing or deleting a media fingerprint library.
//
// Description:
//
// You can call this operation to query the specified jobs of clearing or deleting a media fingerprint library based on the job IDs. If you do not specify job IDs, the system returns the latest 20 jobs that are submitted.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryFpDBDeleteJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFpDBDeleteJobListResponse
func (client *Client) QueryFpDBDeleteJobListWithContext(ctx context.Context, request *QueryFpDBDeleteJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryFpDBDeleteJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFpDBDeleteJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFpDBDeleteJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the jobs of deleting media files from a media fingerprint library.
//
// Description:
//
// You can call this operation to query the specified jobs of deleting media files from a media fingerprint library based on the job IDs. If you do not specify job IDs, the system returns the latest 20 jobs that are submitted.
//
// ### QPS limit
//
// You can call this operation up to 500 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryFpFileDeleteJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFpFileDeleteJobListResponse
func (client *Client) QueryFpFileDeleteJobListWithContext(ctx context.Context, request *QueryFpFileDeleteJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryFpFileDeleteJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFpFileDeleteJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFpFileDeleteJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media fingerprint analysis jobs. You can call this operation to query video fingerprint analysis jobs and text fingerprint analysis jobs.
//
// Description:
//
//	  After a media fingerprint analysis job is submitted, the media fingerprinting service compares the fingerprints of the job input with those of the media files in the media fingerprint library. You can call this operation to query the job results.
//
//		- You can query the results of a text fingerprint analysis job only in the China (Shanghai) region.
//
// ### [](#qps)QPS limits
//
// You can call this API operation up to 100 times per second per account. Requests that exceed this limit are dropped, and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryFpShotJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFpShotJobListResponse
func (client *Client) QueryFpShotJobListWithContext(ctx context.Context, request *QueryFpShotJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryFpShotJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfJobCreatedTimeRange) {
		query["EndOfJobCreatedTimeRange"] = request.EndOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartOfJobCreatedTimeRange) {
		query["StartOfJobCreatedTimeRange"] = request.StartOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFpShotJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFpShotJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryIProductionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIProductionJobResponse
func (client *Client) QueryIProductionJobWithContext(ctx context.Context, request *QueryIProductionJobRequest, runtime *dara.RuntimeOptions) (_result *QueryIProductionJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryIProductionJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryIProductionJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries transcoding jobs at a time by job ID.
//
// Description:
//
//	  By default, returned jobs are sorted in descending order by CreationTime.
//
//		- You can call this operation to query up to 10 transcoding jobs at a time.
//
//		- If you do not set the JobIds parameter, the `InvalidParameter` error code is returned.
//
// ## QPS limit
//
// You can call this API operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/qps-limit).
//
// @param request - QueryJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryJobListResponse
func (client *Client) QueryJobListWithContext(ctx context.Context, request *QueryJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a content moderation job.
//
// Description:
//
// In the content moderation results, the moderation results of the video are sorted in ascending order by time into a timeline. If the video is long, the content moderation results are paginated, and the first page is returned. You can call this operation again to query the remaining moderation results of the video.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryMediaCensorJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaCensorJobDetailResponse
func (client *Client) QueryMediaCensorJobDetailWithContext(ctx context.Context, request *QueryMediaCensorJobDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaCensorJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaCensorJobDetail"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaCensorJobDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries content moderation jobs.
//
// Description:
//
// You can call this operation to query only the content moderation jobs within three months.
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryMediaCensorJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaCensorJobListResponse
func (client *Client) QueryMediaCensorJobListWithContext(ctx context.Context, request *QueryMediaCensorJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaCensorJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfJobCreatedTimeRange) {
		query["EndOfJobCreatedTimeRange"] = request.EndOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartOfJobCreatedTimeRange) {
		query["StartOfJobCreatedTimeRange"] = request.StartOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaCensorJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaCensorJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of media information analysis jobs.
//
// Description:
//
//	  In asynchronous mode, the media information can be retrieved only after the Message Service (MNS) callback of **submitting a media information job*	- is returned. If you have not retrieved the media information for a long period, the job may have failed.
//
//		- You can call this operation to query up to 10 media information analysis jobs at a time.
//
//		- By default, returned jobs are sorted in descending order by CreationTime.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryMediaInfoJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaInfoJobListResponse
func (client *Client) QueryMediaInfoJobListWithContext(ctx context.Context, request *QueryMediaInfoJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaInfoJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaInfoJobIds) {
		query["MediaInfoJobIds"] = request.MediaInfoJobIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaInfoJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaInfoJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media files based on media file IDs.
//
// Description:
//
// You can call this operation to query up to 10 media files at a time.
//
// ## QPS limit
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limits](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/qps-limit).
//
// @param request - QueryMediaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaListResponse
func (client *Client) QueryMediaListWithContext(ctx context.Context, request *QueryMediaListRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeMediaInfo) {
		query["IncludeMediaInfo"] = request.IncludeMediaInfo
	}

	if !dara.IsNil(request.IncludePlayList) {
		query["IncludePlayList"] = request.IncludePlayList
	}

	if !dara.IsNil(request.IncludeSnapshotList) {
		query["IncludeSnapshotList"] = request.IncludeSnapshotList
	}

	if !dara.IsNil(request.IncludeSummaryList) {
		query["IncludeSummaryList"] = request.IncludeSummaryList
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media files based on their Object Storage Service (OSS) URLs.
//
// Description:
//
//	  You can call this operation to query up to 10 media files at a time.
//
//		- Before you call this operation, you must call the [AddMedia](https://help.aliyun.com/document_detail/44458.html) operation to add media files.
//
//		- You can call this operation to query only media files that are processed in a workflow. To obtain comprehensive information about a media file that is newly uploaded to OSS, you can call this operation after the corresponding workflow is complete. To query media files that are not processed in a workflow, you must call the [SubmitMediaInfoJob](https://help.aliyun.com/document_detail/29220.html) operation to submit a media information analysis job. After the job is complete, you can query the information about the media files.
//
// ## QPS limit
//
// You can call this API operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limit](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/qps-limit).
//
// @param request - QueryMediaListByURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaListByURLResponse
func (client *Client) QueryMediaListByURLWithContext(ctx context.Context, request *QueryMediaListByURLRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaListByURLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileURLs) {
		query["FileURLs"] = request.FileURLs
	}

	if !dara.IsNil(request.IncludeMediaInfo) {
		query["IncludeMediaInfo"] = request.IncludeMediaInfo
	}

	if !dara.IsNil(request.IncludePlayList) {
		query["IncludePlayList"] = request.IncludePlayList
	}

	if !dara.IsNil(request.IncludeSnapshotList) {
		query["IncludeSnapshotList"] = request.IncludeSnapshotList
	}

	if !dara.IsNil(request.IncludeSummaryList) {
		query["IncludeSummaryList"] = request.IncludeSummaryList
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaListByURL"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaListByURLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media workflow execution instances.
//
// Description:
//
//	  You can call this operation to query a maximum of 10 media workflow execution instances at a time.
//
//		- Before you call this operation, make sure that the workflow pipeline is enabled. Otherwise, the workflow may not run as expected. For example, the following exceptions may occur: the workflow node is invalid and jobs created in the workflow cannot be executed.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryMediaWorkflowExecutionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaWorkflowExecutionListResponse
func (client *Client) QueryMediaWorkflowExecutionListWithContext(ctx context.Context, request *QueryMediaWorkflowExecutionListRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaWorkflowExecutionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RunIds) {
		query["RunIds"] = request.RunIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaWorkflowExecutionList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaWorkflowExecutionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries registered media workflows.
//
// Description:
//
// You can call this operation to query up to 10 media workflows at a time.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryMediaWorkflowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaWorkflowListResponse
func (client *Client) QueryMediaWorkflowListWithContext(ctx context.Context, request *QueryMediaWorkflowListRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaWorkflowListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaWorkflowIds) {
		query["MediaWorkflowIds"] = request.MediaWorkflowIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaWorkflowList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaWorkflowListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an ApsaraVideo Media Processing (MPS) queue by using the ID of the queue.
//
// Description:
//
//	  You can call this operation to query up to 10 MPS queues at a time.
//
//		- If `"Code": "InvalidIdentity.ServiceDisabled","Message": "The request identity was not allowed operated.","Recommend"` is returned after you call this operation, check whether the RAM user that you use is assigned the AliyunMTSDefaultRole role to obtain the permissions on MPS and whether your Alibaba Cloud account has overdue payments.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryPipelineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPipelineListResponse
func (client *Client) QueryPipelineListWithContext(ctx context.Context, request *QueryPipelineListRequest, runtime *dara.RuntimeOptions) (_result *QueryPipelineListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineIds) {
		query["PipelineIds"] = request.PipelineIds
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPipelineList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPipelineListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询智能标签任务
//
// @param request - QuerySmarttagJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmarttagJobResponse
func (client *Client) QuerySmarttagJobWithContext(ctx context.Context, request *QuerySmarttagJobRequest, runtime *dara.RuntimeOptions) (_result *QuerySmarttagJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmarttagJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmarttagJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the analysis template of a smart tagging job.
//
// Description:
//
// If you call this operation to query the information about a smart tagging template, you must specify the template ID. Otherwise, the operation returns the information about all the templates that are created by the current RAM user.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QuerySmarttagTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmarttagTemplateListResponse
func (client *Client) QuerySmarttagTemplateListWithContext(ctx context.Context, request *QuerySmarttagTemplateListRequest, runtime *dara.RuntimeOptions) (_result *QuerySmarttagTemplateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmarttagTemplateList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmarttagTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries snapshot jobs.
//
// Description:
//
// The status of the snapshot jobs that you want to query.
//
//   - **Submitted**: The job was submitted.
//
//   - **Snapshoting**: The job is being processed.
//
//   - **Success**: The job was successfully processed.
//
//   - **Fail**: The job failed.
//
// @param request - QuerySnapshotJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySnapshotJobListResponse
func (client *Client) QuerySnapshotJobListWithContext(ctx context.Context, request *QuerySnapshotJobListRequest, runtime *dara.RuntimeOptions) (_result *QuerySnapshotJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfJobCreatedTimeRange) {
		query["EndOfJobCreatedTimeRange"] = request.EndOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SnapshotJobIds) {
		query["SnapshotJobIds"] = request.SnapshotJobIds
	}

	if !dara.IsNil(request.StartOfJobCreatedTimeRange) {
		query["StartOfJobCreatedTimeRange"] = request.StartOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySnapshotJobList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySnapshotJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query up to 10 transcoding templates at a time.
//
// Description:
//
// The IDs of the transcoding templates that you want to query. You can query up to 10 transcoding templates at a time. Separate multiple IDs of custom transcoding templates with commas (,).
//
// @param request - QueryTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTemplateListResponse
func (client *Client) QueryTemplateListWithContext(ctx context.Context, request *QueryTemplateListRequest, runtime *dara.RuntimeOptions) (_result *QueryTemplateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTemplateList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频溯源水印ab流任务
//
// @param request - QueryTraceAbJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTraceAbJobResponse
func (client *Client) QueryTraceAbJobWithContext(ctx context.Context, request *QueryTraceAbJobRequest, runtime *dara.RuntimeOptions) (_result *QueryTraceAbJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTraceAbJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTraceAbJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询溯源水印提取任务
//
// @param request - QueryTraceExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTraceExtractJobResponse
func (client *Client) QueryTraceExtractJobWithContext(ctx context.Context, request *QueryTraceExtractJobRequest, runtime *dara.RuntimeOptions) (_result *QueryTraceExtractJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTraceExtractJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTraceExtractJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频溯源水印m3u8任务
//
// @param request - QueryTraceM3u8JobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTraceM3u8JobResponse
func (client *Client) QueryTraceM3u8JobWithContext(ctx context.Context, request *QueryTraceM3u8JobRequest, runtime *dara.RuntimeOptions) (_result *QueryTraceM3u8JobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTraceM3u8Job"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTraceM3u8JobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries watermark templates.
//
// Description:
//
// You can call this operation to query up to 10 watermark templates at a time.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - QueryWaterMarkTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWaterMarkTemplateListResponse
func (client *Client) QueryWaterMarkTemplateListWithContext(ctx context.Context, request *QueryWaterMarkTemplateListRequest, runtime *dara.RuntimeOptions) (_result *QueryWaterMarkTemplateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WaterMarkTemplateIds) {
		query["WaterMarkTemplateIds"] = request.WaterMarkTemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWaterMarkTemplateList"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWaterMarkTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a custom face.
//
// Description:
//
//	  You can call this operation to register only one custom face at a time.
//
//		- A maximum of 10 images can be registered for a custom face.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - RegisterCustomFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCustomFaceResponse
func (client *Client) RegisterCustomFaceWithContext(ctx context.Context, request *RegisterCustomFaceRequest, runtime *dara.RuntimeOptions) (_result *RegisterCustomFaceResponse, _err error) {
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

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.PersonName) {
		query["PersonName"] = request.PersonName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterCustomFace"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterCustomFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册自定义实体项
//
// @param request - RegisterCustomViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCustomViewResponse
func (client *Client) RegisterCustomViewWithContext(ctx context.Context, request *RegisterCustomViewRequest, runtime *dara.RuntimeOptions) (_result *RegisterCustomViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.CustomEntityId) {
		query["CustomEntityId"] = request.CustomEntityId
	}

	if !dara.IsNil(request.CustomGroupId) {
		query["CustomGroupId"] = request.CustomGroupId
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.LabelPrompt) {
		query["LabelPrompt"] = request.LabelPrompt
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterCustomView"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterCustomViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media workflows in the specified state.
//
// Description:
//
// You can call this operation to query media workflows in the specified state. If you do not specify the state, all media workflows are queried by default.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SearchMediaWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaWorkflowResponse
func (client *Client) SearchMediaWorkflowWithContext(ctx context.Context, request *SearchMediaWorkflowRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StateList) {
		query["StateList"] = request.StateList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMediaWorkflow"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for ApsaraVideo Media Processing (MPS) queues in the specified state.
//
// Description:
//
// You can call this operation to query MPS queues in the specified state. If you do not specify the state, all MPS queues are queried by default.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SearchPipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchPipelineResponse
func (client *Client) SearchPipelineWithContext(ctx context.Context, request *SearchPipelineRequest, runtime *dara.RuntimeOptions) (_result *SearchPipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchPipeline"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchPipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query custom transcoding templates in the specified state.
//
// Description:
//
// You can call this operation up to 100 times per second. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SearchTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTemplateResponse
func (client *Client) SearchTemplateWithContext(ctx context.Context, request *SearchTemplateRequest, runtime *dara.RuntimeOptions) (_result *SearchTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamePrefix) {
		query["NamePrefix"] = request.NamePrefix
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for watermark templates.
//
// Description:
//
// The total number of returned entries.
//
// @param request - SearchWaterMarkTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchWaterMarkTemplateResponse
func (client *Client) SearchWaterMarkTemplateWithContext(ctx context.Context, request *SearchWaterMarkTemplateRequest, runtime *dara.RuntimeOptions) (_result *SearchWaterMarkTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchWaterMarkTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchWaterMarkTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a preset template analysis job.
//
// Description:
//
//	  After you call the SubmitAnalysisJob operation to submit a preset template analysis job, ApsaraVideo Media Processing (MPS) intelligently analyzes the input file of the job and recommends a suitable preset template. You can call the [QueryAnalysisJobList](https://help.aliyun.com/document_detail/29224.html) operation to query the analysis result or enable asynchronous notifications to receive the analysis result.
//
//		- The analysis result is retained only for two weeks after it is generated. The analysis result is deleted after two weeks. If you use the recommended preset template in a transcoding job after two weeks, the job fails, and the `AnalysisResultNotFound` error code is returned.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitAnalysisJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAnalysisJobResponse
func (client *Client) SubmitAnalysisJobWithContext(ctx context.Context, request *SubmitAnalysisJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAnalysisJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisConfig) {
		query["AnalysisConfig"] = request.AnalysisConfig
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAnalysisJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAnalysisJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交版权水印提取任务
//
// @param request - SubmitCopyrightExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCopyrightExtractJobResponse
func (client *Client) SubmitCopyrightExtractJobWithContext(ctx context.Context, request *SubmitCopyrightExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitCopyrightExtractJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCopyrightExtractJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCopyrightExtractJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频版权水印
//
// @param request - SubmitCopyrightJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCopyrightJobResponse
func (client *Client) SubmitCopyrightJobWithContext(ctx context.Context, request *SubmitCopyrightJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitCopyrightJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.Output) {
		query["Output"] = request.Output
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TotalTime) {
		query["TotalTime"] = request.TotalTime
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCopyrightJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCopyrightJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a job of clearing or deleting a media fingerprint library.
//
// Description:
//
// You can call this operation to clear or delete the specified media fingerprint library based on the library ID. If you clear a media fingerprint library, the content in the library is deleted, but the library is not deleted. If you delete a media fingerprint library, both the library and the content in the library are deleted. If you do not specify the operation type, the system clears the media fingerprint library by default.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitFpDBDeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFpDBDeleteJobResponse
func (client *Client) SubmitFpDBDeleteJobWithContext(ctx context.Context, request *SubmitFpDBDeleteJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitFpDBDeleteJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DelType) {
		query["DelType"] = request.DelType
	}

	if !dara.IsNil(request.FpDBId) {
		query["FpDBId"] = request.FpDBId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitFpDBDeleteJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitFpDBDeleteJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a job of deleting media files from a media fingerprint library.
//
// Description:
//
// ## [](#)Limits
//
//   - You can call this operation to delete up to 200 media files from a media fingerprint library at a time.
//
//   - This operation is available in the following regions: China (Beijing), China (Hangzhou), China (Shanghai), and Singapore.
//
// ## [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitFpFileDeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFpFileDeleteJobResponse
func (client *Client) SubmitFpFileDeleteJobWithContext(ctx context.Context, request *SubmitFpFileDeleteJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitFpFileDeleteJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileIds) {
		query["FileIds"] = request.FileIds
	}

	if !dara.IsNil(request.FpDBId) {
		query["FpDBId"] = request.FpDBId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.PrimaryKeys) {
		query["PrimaryKeys"] = request.PrimaryKeys
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitFpFileDeleteJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitFpFileDeleteJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a media fingerprint analysis job to query the media fingerprint repository for a media file with the identical or similar fingerprint as the input file.
//
// Description:
//
//	  You can call this operation to submit a video, audio, image, or text fingerprint analysis job.
//
//		- This operation asynchronously submits a job. The query results may not have been generated when the response is returned. After the results are generated, an asynchronous message is returned.
//
//		- You can submit a text fingerprint analysis job only in the China (Shanghai) region.
//
//		- The input file of the job must be in one of the following formats:
//
//	    	- Image formats: JPEG, PNG, and BMP.
//
//	    	- Video formats: MP4, AVI, MKV, MPG, TS, MOV, FLV, MXF.
//
//	    	- Video encoding formats: MPEG2, MPEG4, H264, HEVC, and WMV.
//
// ### QPS limit
//
// You can call this operation up to 150 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitFpShotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFpShotJobResponse
func (client *Client) SubmitFpShotJobWithContext(ctx context.Context, request *SubmitFpShotJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitFpShotJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FpShotConfig) {
		query["FpShotConfig"] = request.FpShotConfig
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitFpShotJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitFpShotJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
//	  Jobs that are submitted by calling this operation run in an asynchronous manner. After a job is added to the ApsaraVideo Media Processing (MPS) queue, the job is scheduled to run. You can call the [QueryIProductionJob](https://help.aliyun.com/document_detail/170217.html) operation or configure a callback to query the job result.
//
//		- Capabilities provided by the intelligent production feature vary based on the region. Before you call this operation to submit an intelligent production job, check whether the job is supported in the region in which your service is activated. For more information, see [Regions and endpoints](https://help.aliyun.com/document_detail/43248.html).
//
// ### [](#qps)QPS limit
//
// You can call this API operation up to 100 times per second per account. Requests that exceed this limit are dropped, and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitIProductionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIProductionJobResponse
func (client *Client) SubmitIProductionJobWithContext(ctx context.Context, request *SubmitIProductionJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitIProductionJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParams) {
		query["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.Output) {
		query["Output"] = request.Output
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScheduleParams) {
		query["ScheduleParams"] = request.ScheduleParams
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIProductionJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIProductionJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交图片版权水印任务
//
// @param request - SubmitImageCopyrightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageCopyrightResponse
func (client *Client) SubmitImageCopyrightWithContext(ctx context.Context, request *SubmitImageCopyrightRequest, runtime *dara.RuntimeOptions) (_result *SubmitImageCopyrightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.Output) {
		query["Output"] = request.Output
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImageCopyright"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImageCopyrightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits transcoding jobs. If the transcoding jobs and workflows created in the ApsaraVideo Media Processing (MPS) console cannot meet your business requirements, you can call the SubmitJobs operation to submit transcoding jobs. Specify transcoding parameters as required when you call the SubmitJobs operation.
//
// Description:
//
//	  SubmitJobs is an asynchronous operation. After you submit transcoding jobs, the jobs are added to an MPS queue to be scheduled and run. The transcoding jobs may not have been complete when the response is returned. After you call this operation, you can call the [QueryJobList](https://help.aliyun.com/document_detail/602836.html) operation to query the job results. You can also associate a Message Service (MNS) queue or topic with the MPS queue to receive notifications on the jobs. For more information, see [Receive notifications](https://help.aliyun.com/document_detail/42618.html).
//
//		- An input file can be up to 100 GB in size. If the size of the input file exceeds this limit, the job may fail.
//
//		- If you use an **intelligent preset template*	- to transcode an input file, you must first call the [SubmitAnalysisJob](https://help.aliyun.com/document_detail/29223.html) operation to submit a preset template analysis job. After the analysis job is complete, you can call the [QueryAnalysisJobList](https://help.aliyun.com/document_detail/29224.html)operation to obtain the available preset templates for the input file. When you submit a transcoding job, set TemplateId to the ID of an available preset template. If you specify a preset template that is not in the available preset templates, the transcoding job fails.
//
//		- If you use a **static preset template*	- to transcode an input file, you do not need to submit a preset template analysis job.
//
//		- If you want to use multiple accounts in MPS, you can create Resource Access Management (RAM) users by using your Alibaba Cloud account. For more information, see [Create a RAM user and grant permissions to the RAM user](https://help.aliyun.com/document_detail/42841.html). If the Alibaba Cloud account that is used to query transcoding jobs is not the one that is used to submit the transcoding jobs, no data is returned.
//
//		- For information about transcoding FAQ, see [FAQ about MPS](https://help.aliyun.com/document_detail/38986.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped, and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitJobsResponse
func (client *Client) SubmitJobsWithContext(ctx context.Context, request *SubmitJobsRequest, runtime *dara.RuntimeOptions) (_result *SubmitJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OutputBucket) {
		query["OutputBucket"] = request.OutputBucket
	}

	if !dara.IsNil(request.OutputLocation) {
		query["OutputLocation"] = request.OutputLocation
	}

	if !dara.IsNil(request.Outputs) {
		query["Outputs"] = request.Outputs
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitJobs"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a content moderation job.
//
// Description:
//
// The job that you submit by calling this operation is run in asynchronous mode. The job is added to an ApsaraVideo Media Processing (MPS) queue and then scheduled, queued, and run. You can call the [QueryMediaCensorJobDetail](https://help.aliyun.com/document_detail/91779.html) operation or configure an asynchronous notification to obtain the job result.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitMediaCensorJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaCensorJobResponse
func (client *Client) SubmitMediaCensorJobWithContext(ctx context.Context, request *SubmitMediaCensorJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaCensorJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Barrages) {
		query["Barrages"] = request.Barrages
	}

	if !dara.IsNil(request.CoverImages) {
		query["CoverImages"] = request.CoverImages
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExternalUrl) {
		query["ExternalUrl"] = request.ExternalUrl
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VideoCensorConfig) {
		query["VideoCensorConfig"] = request.VideoCensorConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaCensorJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaCensorJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a media information analysis job.
//
// Description:
//
//	  Before you submit a media information analysis job, make sure that the input file is uploaded to an Object Storage Service (OSS) bucket. Otherwise, the job fails. You can configure upload callbacks to be notified of the upload status of files.****
//
//		- A media information analysis job can be run in synchronous or asynchronous mode.
//
//		- In asynchronous mode, the media information analysis job is submitted to and scheduled in an ApsaraVideo Media Processing (MPS) queue. In this case, the media information analysis job may be queued. The media information analysis job may not be generated when the response to the SubmitMediaInfoJob operation is returned. After the execution is complete, you can call the [QueryMediaInfoJobList](https://help.aliyun.com/document_detail/602828.html) operation to poll the execution results, or associate a Message Service (MNS) queue or topic with the MPS queue to receive the execution results. For more information, see [Receive message notifications](https://www.alibabacloud.com/help/en/mps/receive-message-notifications/?spm=a2c63.p38356.0.0.b48576d2jxNSca).
//
// ### QPS limits
//
// You can call this operation up to 50 times per second per account. If the number of requests that you send to call this operation within one second reaches the request limit of this operation, new requests fail and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitMediaInfoJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaInfoJobResponse
func (client *Client) SubmitMediaInfoJobWithContext(ctx context.Context, request *SubmitMediaInfoJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaInfoJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Async) {
		query["Async"] = request.Async
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaInfoJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaInfoJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitSmarttagJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmarttagJobResponse
func (client *Client) SubmitSmarttagJobWithContext(ctx context.Context, request *SubmitSmarttagJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmarttagJobResponse, _err error) {
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

	if !dara.IsNil(request.ContentAddr) {
		query["ContentAddr"] = request.ContentAddr
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSmarttagJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSmarttagJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a snapshot job.
//
// Description:
//
//	  Only JPG images can be generated by calling this operation.
//
//		- Asynchronous mode: This operation may return a response before snapshots are captured. Snapshot jobs are queued in the background and asynchronously processed by ApsaraVideo Media Processing (MPS). If the **Interval*	- or **Num*	- parameter is set, the snapshot job is processed in asynchronous mode. For more information about FAQ about capturing snapshots, see [FAQ about taking snapshots](https://help.aliyun.com/document_detail/60805.html).
//
//		- Notifications: When you submit a snapshot job, the **PipelineId*	- parameter is required. An asynchronous message is sent only after the notification feature is enabled for the MPS queue.
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - SubmitSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSnapshotJobResponse
func (client *Client) SubmitSnapshotJobWithContext(ctx context.Context, request *SubmitSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSnapshotJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SnapshotConfig) {
		query["SnapshotConfig"] = request.SnapshotConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSnapshotJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSnapshotJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频溯源水印ab流任务
//
// @param request - SubmitTraceAbJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTraceAbJobResponse
func (client *Client) SubmitTraceAbJobWithContext(ctx context.Context, request *SubmitTraceAbJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTraceAbJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.CipherBase64ed) {
		query["CipherBase64ed"] = request.CipherBase64ed
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Output) {
		query["Output"] = request.Output
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TotalTime) {
		query["TotalTime"] = request.TotalTime
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTraceAbJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTraceAbJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交溯源水印提取任务
//
// @param request - SubmitTraceExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTraceExtractJobResponse
func (client *Client) SubmitTraceExtractJobWithContext(ctx context.Context, request *SubmitTraceExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTraceExtractJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTraceExtractJob"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTraceExtractJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频溯源水印m3u8文件任务
//
// @param request - SubmitTraceM3u8JobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTraceM3u8JobResponse
func (client *Client) SubmitTraceM3u8JobWithContext(ctx context.Context, request *SubmitTraceM3u8JobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTraceM3u8JobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyUri) {
		query["KeyUri"] = request.KeyUri
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.Output) {
		query["Output"] = request.Output
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.Trace) {
		query["Trace"] = request.Trace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTraceM3u8Job"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTraceM3u8JobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The description of the figure. The description can be up to 512 characters in length.
//
// Description:
//
// The response parameters.
//
// @param request - TagCustomPersonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagCustomPersonResponse
func (client *Client) TagCustomPersonWithContext(ctx context.Context, request *TagCustomPersonRequest, runtime *dara.RuntimeOptions) (_result *TagCustomPersonResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryDescription) {
		query["CategoryDescription"] = request.CategoryDescription
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PersonDescription) {
		query["PersonDescription"] = request.PersonDescription
	}

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.PersonName) {
		query["PersonName"] = request.PersonName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagCustomPerson"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagCustomPersonResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an input media bucket from the media library.
//
// Description:
//
// # Usage notes
//
// You can call this operation to unbind an input media bucket from the media library based on the name of the output media bucket.
//
// # QPS limits
//
// You can call this API operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UnbindInputBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindInputBucketResponse
func (client *Client) UnbindInputBucketWithContext(ctx context.Context, request *UnbindInputBucketRequest, runtime *dara.RuntimeOptions) (_result *UnbindInputBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindInputBucket"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindInputBucketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to unbind an output media bucket from the media library based on the name of the output media bucket.
//
// ## QPS limit
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limits](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/qps-limit).
//
// Description:
//
// The name of the output media bucket to be unbound. To obtain the media bucket name, you can log on to the **ApsaraVideo Media Processing (MPS)*	- console and choose **Workflows*	- > **Media Buckets*	- in the left-side navigation pane. Alternatively, you can log on to the **Object Storage Service (OSS) console*	- and click **My OSS Paths**.
//
// @param request - UnbindOutputBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindOutputBucketResponse
func (client *Client) UnbindOutputBucketWithContext(ctx context.Context, request *UnbindOutputBucketRequest, runtime *dara.RuntimeOptions) (_result *UnbindOutputBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindOutputBucket"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindOutputBucketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom face, all the custom faces that are registered in a custom figure library, or a custom figure library.
//
// Description:
//
// You can call this operation to delete a specific custom face, all the custom faces that are registered in a custom figure library, or a custom figure library.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped, and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UnregisterCustomFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnregisterCustomFaceResponse
func (client *Client) UnregisterCustomFaceWithContext(ctx context.Context, request *UnregisterCustomFaceRequest, runtime *dara.RuntimeOptions) (_result *UnregisterCustomFaceResponse, _err error) {
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

	if !dara.IsNil(request.FaceId) {
		query["FaceId"] = request.FaceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnregisterCustomFace"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnregisterCustomFaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a media file.
//
// Description:
//
// The basic information that you can update by calling this operation includes the title, description, and category of a media file. This operation applies to a full update. You must set all the parameters unless you want to replace the value of a specific parameter with a NULL value.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UpdateMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaResponse
func (client *Client) UpdateMediaWithContext(ctx context.Context, request *UpdateMediaRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMedia"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the category to which a media file belongs.
//
// Description:
//
// You can call this operation to update only the category of a media file. For more information about how to update all the information about a media file, see [UpdateMedia](https://help.aliyun.com/document_detail/44464.html).
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UpdateMediaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaCategoryResponse
func (client *Client) UpdateMediaCategoryWithContext(ctx context.Context, request *UpdateMediaCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaCategory"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the thumbnail of a media file.
//
// Description:
//
// You can call this operation to update only the thumbnail of a media file. For more information about how to update all the information about a media file, see [UpdateMedia](https://help.aliyun.com/document_detail/44464.html).
//
// ## Limits on QPS
//
// You can call this operation up to 100 times per second. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limits](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/qps-limit).
//
// @param request - UpdateMediaCoverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaCoverResponse
func (client *Client) UpdateMediaCoverWithContext(ctx context.Context, request *UpdateMediaCoverRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaCoverResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaCover"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaCoverResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the publishing status of a media file.
//
// Description:
//
// The published state indicates that the access control list (ACL) of media playback resources and snapshot files is set to inherit the ACL of the bucket to which they belong. The unpublished state indicates that the ACL of media playback resources and snapshot files is set to private.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UpdateMediaPublishStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaPublishStateResponse
func (client *Client) UpdateMediaPublishStateWithContext(ctx context.Context, request *UpdateMediaPublishStateRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaPublishStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Publish) {
		query["Publish"] = request.Publish
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaPublishState"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaPublishStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the topology of a media workflow.
//
// Description:
//
//	  You can call this operation to update the topology of a media workflow. To update the trigger mode of a media workflow, call the [UpdateMediaWorkflowTriggerMode](https://help.aliyun.com/document_detail/70372.html) operation.
//
//		- After you delete or deactivate a media workflow, the workflow cannot be used. In this case, the workflow is not automatically triggered when you upload a file to the bucket specified by the workflow.
//
// <warning>Deleting or deactivating a workflow will not affect tasks that have already been submitted. If a workflow is deleted or deactivated after a task has been submitted, tasks that are already in the processing queue will not be canceled and will be executed normally and charged the corresponding fees.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).>
//
// @param request - UpdateMediaWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaWorkflowResponse
func (client *Client) UpdateMediaWorkflowWithContext(ctx context.Context, request *UpdateMediaWorkflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaWorkflowId) {
		query["MediaWorkflowId"] = request.MediaWorkflowId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Topology) {
		query["Topology"] = request.Topology
	}

	if !dara.IsNil(request.TriggerMode) {
		query["TriggerMode"] = request.TriggerMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaWorkflow"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the trigger mode of a media workflow.
//
// Description:
//
// You can call this operation only to modify the trigger mode of a media workflow. To modify other information about the workflow, call the [UpdateMediaWorkflow](https://help.aliyun.com/document_detail/44438.html) operation.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UpdateMediaWorkflowTriggerModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaWorkflowTriggerModeResponse
func (client *Client) UpdateMediaWorkflowTriggerModeWithContext(ctx context.Context, request *UpdateMediaWorkflowTriggerModeRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaWorkflowTriggerModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaWorkflowId) {
		query["MediaWorkflowId"] = request.MediaWorkflowId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TriggerMode) {
		query["TriggerMode"] = request.TriggerMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaWorkflowTriggerMode"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaWorkflowTriggerModeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an ApsaraVideo Media Processing (MPS) queue.
//
// Description:
//
//	  You can call this operation to modify the name, status, and notification settings of the specified MPS queue.
//
//		- If a paused MPS queue is selected in a workflow or a job, such as a video review or media fingerprint job, the workflow or job fails.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UpdatePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipelineWithContext(ctx context.Context, request *UpdatePipelineRequest, runtime *dara.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtendConfig) {
		query["ExtendConfig"] = request.ExtendConfig
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NotifyConfig) {
		query["NotifyConfig"] = request.NotifyConfig
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipeline"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新智能标签模板接口支持cpv
//
// @param request - UpdateSmarttagTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmarttagTemplateResponse
func (client *Client) UpdateSmarttagTemplateWithContext(ctx context.Context, request *UpdateSmarttagTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmarttagTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalyseTypes) {
		query["AnalyseTypes"] = request.AnalyseTypes
	}

	if !dara.IsNil(request.FaceCategoryIds) {
		query["FaceCategoryIds"] = request.FaceCategoryIds
	}

	if !dara.IsNil(request.FaceCustomParamsConfig) {
		query["FaceCustomParamsConfig"] = request.FaceCustomParamsConfig
	}

	if !dara.IsNil(request.Industry) {
		query["Industry"] = request.Industry
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.KeywordConfig) {
		query["KeywordConfig"] = request.KeywordConfig
	}

	if !dara.IsNil(request.KnowledgeConfig) {
		query["KnowledgeConfig"] = request.KnowledgeConfig
	}

	if !dara.IsNil(request.LabelType) {
		query["LabelType"] = request.LabelType
	}

	if !dara.IsNil(request.LabelVersion) {
		query["LabelVersion"] = request.LabelVersion
	}

	if !dara.IsNil(request.LandmarkGroupIds) {
		query["LandmarkGroupIds"] = request.LandmarkGroupIds
	}

	if !dara.IsNil(request.ObjectGroupIds) {
		query["ObjectGroupIds"] = request.ObjectGroupIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmarttagTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmarttagTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// A custom transcoding template cannot be updated if it is being used by a job that has been submitted.The ID of the template. You can obtain the template ID from the response of the [AddTemplate](https://help.aliyun.com/document_detail/213306.html) operation.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UpdateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithContext(ctx context.Context, request *UpdateTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
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

	if !dara.IsNil(request.Container) {
		query["Container"] = request.Container
	}

	if !dara.IsNil(request.MuxConfig) {
		query["MuxConfig"] = request.MuxConfig
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TransConfig) {
		query["TransConfig"] = request.TransConfig
	}

	if !dara.IsNil(request.Video) {
		query["Video"] = request.Video
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name and configurations of the specified watermark template.
//
// Description:
//
//	  You can call this operation to update the information about a watermark template based on the ID of the watermark template. For example, you can update the name and configurations of a watermark template.
//
//		- A watermark template cannot be updated if it is being used by a job that has been submitted.
//
// ### QPS limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342832.html).
//
// @param request - UpdateWaterMarkTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWaterMarkTemplateResponse
func (client *Client) UpdateWaterMarkTemplateWithContext(ctx context.Context, request *UpdateWaterMarkTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateWaterMarkTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WaterMarkTemplateId) {
		query["WaterMarkTemplateId"] = request.WaterMarkTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWaterMarkTemplate"),
		Version:     dara.String("2014-06-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWaterMarkTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
