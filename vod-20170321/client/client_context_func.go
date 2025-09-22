// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds an AI template for automated review and smart thumbnail tasks.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- Before you add an AI template for automated review and smart thumbnail tasks, make sure that [automated review](https://ai.aliyun.com/vi/censor) and [smart thumbnail](https://ai.aliyun.com/vi/cover) are enabled.
//
// @param request - AddAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAITemplateResponse
func (client *Client) AddAITemplateWithContext(ctx context.Context, request *AddAITemplateRequest, runtime *dara.RuntimeOptions) (_result *AddAITemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAITemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAITemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a video category. You can call this operation to categorize media assets including audio or video files, images, and short video materials in ApsaraVideo VOD. This simplifies the query and management of media assets.
//
// Description:
//
//	  You can create a maximum of 3 levels of categories for audio, video, and image files and 2 levels of categories for short video materials. Each category level can contain a maximum of 100 subcategories. To create categories for audio and video files, set `Type` to `default`. To create categories for short video materials, set `Type` to `material`.
//
//		- After you create a category, you can categorize media resources during upload or categorize the uploaded media resources. For more information, see [Manage video categories](https://help.aliyun.com/document_detail/86070.html).
//
// @param request - AddCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCategoryResponse
func (client *Client) AddCategoryWithContext(ctx context.Context, request *AddCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateName) {
		query["CateName"] = request.CateName
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCategory"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an online editing project.
//
// Description:
//
//	For more information about the online editing feature, see [Overview](https://help.aliyun.com/document_detail/95482.html).
//
// @param request - AddEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEditingProjectResponse
func (client *Client) AddEditingProjectWithContext(ctx context.Context, request *AddEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *AddEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Division) {
		query["Division"] = request.Division
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

	if !dara.IsNil(request.Timeline) {
		query["Timeline"] = request.Timeline
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEditingProject"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEditingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more materials to an editing project.
//
// @param request - AddEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEditingProjectMaterialsResponse
func (client *Client) AddEditingProjectMaterialsWithContext(ctx context.Context, request *AddEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *AddEditingProjectMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaterialIds) {
		query["MaterialIds"] = request.MaterialIds
	}

	if !dara.IsNil(request.MaterialType) {
		query["MaterialType"] = request.MaterialType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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
		Action:      dara.String("AddEditingProjectMaterials"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEditingProjectMaterialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a transcoding template group or adds transcoding templates to a transcoding template group.
//
// Description:
//
//	  You cannot perform custom operations on transcoding template groups that are **locked*	- in the ApsaraVideo VOD console. You can call the [GetTranscodeTemplateGroup](~~GetTranscodeTemplateGroup~~) operation to query the information about a transcoding template group and check whether the transcoding template group is locked based on the value of the Locked parameter. You can call the [UpdateTranscodeTemplateGroup](~~UpdateTranscodeTemplateGroup~~) operation to unlock a transcoding template group if it is locked. Then, you can perform custom operations on the transcoding template group.
//
//		- An Object Storage Service (OSS) bucket is required to store files that are used for transcoding. You cannot create a transcoding template group if no bucket is available. To activate a bucket, perform the following operations: Log on to the ApsaraVideo VOD console. In the left-side navigation pane, choose **Configuration Management > Media Management > Storage**. On the **Storage*	- page, activate the bucket that is allocated by ApsaraVideo VOD.
//
//		- You cannot add transcoding templates to the **No Transcoding*	- template group.
//
//		- You can create a maximum of 20 transcoding template groups.
//
//		- You can add a maximum of 20 transcoding templates to a transcoding template group.
//
//		- If you want to generate a URL for adaptive bitrate streaming, you can add video packaging templates to a transcoding template group. You can add a maximum of 10 video packaging templates to a transcoding template group. If you add more than 10 video packaging templates, URLs of the video transcoded based on the video packaging templates are generated but the URL for adaptive bitrate streaming is not generated.
//
// ### QPS limits
//
// You can call this operation up to five times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - AddTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTranscodeTemplateGroupResponse
func (client *Client) AddTranscodeTemplateGroupWithContext(ctx context.Context, request *AddTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *AddTranscodeTemplateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TranscodeTemplateGroupId) {
		query["TranscodeTemplateGroupId"] = request.TranscodeTemplateGroupId
	}

	if !dara.IsNil(request.TranscodeTemplateList) {
		query["TranscodeTemplateList"] = request.TranscodeTemplateList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTranscodeTemplateGroup"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTranscodeTemplateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to accelerate in ApsaraVideo VOD.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- Before you add a domain name to accelerate, you must activate ApsaraVideo VOD and apply for an Internet content provider (ICP) filing for the domain name. For more information about how to activate ApsaraVideo VOD, see [Activate ApsaraVideo VOD](https://help.aliyun.com/document_detail/51512.html).
//
//		- If the content on the origin server is not stored on Alibaba Cloud, the content must be reviewed by Alibaba Cloud. The review will be complete by the end of the next business day after you submit an application.
//
//		- You can add only one domain name to accelerate in a request. You can add a maximum of 20 accelerated domain names within an Alibaba Cloud account.
//
// @param request - AddVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVodDomainResponse
func (client *Client) AddVodDomainWithContext(ctx context.Context, request *AddVodDomainRequest, runtime *dara.RuntimeOptions) (_result *AddVodDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddVodDomain"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVodDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a storage bucket to one or more applications in ApsaraVideo VOD.
//
// Description:
//
// You can call this operation to add a buckets to an ApsaraVideo VOD applications.
//
// > You can add only one ApsaraVideo VOD bucket for each application. If you specify an AppId that does not exist or the ID of an application for which an VOD bucket is enabled, an error is returned.
//
// @param request - AddVodStorageForAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVodStorageForAppResponse
func (client *Client) AddVodStorageForAppWithContext(ctx context.Context, request *AddVodStorageForAppRequest, runtime *dara.RuntimeOptions) (_result *AddVodStorageForAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddVodStorageForApp"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVodStorageForAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a snapshot template or frame animation template.
//
// Description:
//
//	  After you add a snapshot template, you can call the [SubmitSnapshotJob](https://help.aliyun.com/document_detail/72213.html) operation and specify the template ID to submit a snapshot job.
//
//		- You can use the HTTP (HTTPS compatible) callback or MNS callback method to receive the [SnapshotComplete](https://help.aliyun.com/document_detail/57337.html) callback. For more information, see [Overview](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - AddVodTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVodTemplateResponse
func (client *Client) AddVodTemplateWithContext(ctx context.Context, request *AddVodTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddVodTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddVodTemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVodTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image or text watermark. ApsaraVideo VOD allows you to create watermark templates to reuse your parameter configurations such as watermark position, size, font, and color. Each watermark template is assigned a unique ID. This simplifies the progress of creating watermark tasks.
//
// Description:
//
//	  You can call this operation to create an `Image` watermark template or a `Text` watermark template. You can use static images in the PNG format or dynamic images in the GIF, APNG, and MOV formats as image watermarks.
//
//		- After you call this operation to create a watermark template, you must call the [AddTranscodeTemplateGroup](~~AddTranscodeTemplateGroup~~) or [UpdateTranscodeTemplateGroup](~~UpdateTranscodeTemplateGroup~~) operation to associate the watermark template with a transcoding template group. This way, you can add watermarks to videos during transcoding.
//
//		- For more information, see [Video watermarks](https://help.aliyun.com/document_detail/99369.html).
//
// @param request - AddWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWatermarkResponse
func (client *Client) AddWatermarkWithContext(ctx context.Context, request *AddWatermarkRequest, runtime *dara.RuntimeOptions) (_result *AddWatermarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.WatermarkConfig) {
		query["WatermarkConfig"] = request.WatermarkConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWatermark"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a RAM user or RAM role permissions to access ApsaraVideo VOD applications.
//
// Description:
//
// > You can grant a RAM user or RAM role permissions to access up to 10 applications.
//
// @param request - AttachAppPolicyToIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachAppPolicyToIdentityResponse
func (client *Client) AttachAppPolicyToIdentityWithContext(ctx context.Context, request *AttachAppPolicyToIdentityRequest, runtime *dara.RuntimeOptions) (_result *AttachAppPolicyToIdentityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.IdentityName) {
		query["IdentityName"] = request.IdentityName
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.PolicyNames) {
		query["PolicyNames"] = request.PolicyNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachAppPolicyToIdentity"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachAppPolicyToIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the basic information and source file information of multiple media assets.
//
// Description:
//
//	  You can specify up to 20 audio or video file IDs in each request.
//
//		- After a media file is uploaded, ApsaraVideo VOD processes the source file. Then, information about the media file is asynchronously generated. You can configure notifications for the [VideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event and call this operation to query information about a media file after you receive notifications for the [VideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event. For more information, see [Overview](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - BatchGetMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetMediaInfosResponse
func (client *Client) BatchGetMediaInfosWithContext(ctx context.Context, request *BatchGetMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *BatchGetMediaInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetMediaInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetMediaInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures one or more domain names for CDN.
//
// Description:
//
// > This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - BatchSetVodDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetVodDomainConfigsResponse
func (client *Client) BatchSetVodDomainConfigsWithContext(ctx context.Context, request *BatchSetVodDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetVodDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetVodDomainConfigs"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetVodDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables accelerated domain names that are in the disabled state.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- If the domain name that you want to enable is invalid or your Alibaba Cloud account has overdue payments, you cannot call this operation to enable the domain name.
//
// @param request - BatchStartVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartVodDomainResponse
func (client *Client) BatchStartVodDomainWithContext(ctx context.Context, request *BatchStartVodDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStartVodDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStartVodDomain"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStartVodDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables accelerated domain names.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- After you disable an accelerated domain name, the information about the domain name is retained. The system automatically reroutes all the requests that are destined for the domain name to the origin server.
//
// @param request - BatchStopVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopVodDomainResponse
func (client *Client) BatchStopVodDomainWithContext(ctx context.Context, request *BatchStopVodDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStopVodDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStopVodDomain"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStopVodDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels URL-based upload jobs in the queue.
//
// Description:
//
//	  You can cancel only URL-based upload jobs in the **Pending*	- state. You can query the status of a URL-based upload job by calling the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation.
//
//		- You cannot cancel an upload job that already starts.
//
//		- You must specify either JobIds or UploadUrls. If you specify both parameters, only JobIds takes effect.
//
// @param request - CancelUrlUploadJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelUrlUploadJobsResponse
func (client *Client) CancelUrlUploadJobsWithContext(ctx context.Context, request *CancelUrlUploadJobsRequest, runtime *dara.RuntimeOptions) (_result *CancelUrlUploadJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.UploadUrls) {
		query["UploadUrls"] = request.UploadUrls
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelUrlUploadJobs"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelUrlUploadJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers a resource to a specified resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// Description:
//
// You can create up to 10 applications within an Alibaba Cloud account. For more information, see [Multi-application service](https://help.aliyun.com/document_detail/113600.html).
//
// ### QPS limits
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits on API operations in ApsaraVideo VOD](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - CreateAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppInfoResponse
func (client *Client) CreateAppInfoWithContext(ctx context.Context, request *CreateAppInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateAppInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs manual review on media files, such as audio and video files.
//
// @param request - CreateAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuditResponse
func (client *Client) CreateAuditWithContext(ctx context.Context, request *CreateAuditRequest, runtime *dara.RuntimeOptions) (_result *CreateAuditResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditContent) {
		query["AuditContent"] = request.AuditContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAudit"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an upload URL and an upload credential for an auxiliary media asset such as a watermark image, subtitle file, or material and generates the media ID. ApsaraVideo VOD issues upload URLs and credentials to perform authorization and ensure security. This prevents unauthorized users from uploading media files. ApsaraVideo VOD generates media IDs together with upload URLs and credentials. Media IDs are used in lifecycle management and media processing.
//
// Description:
//
//	  **Make sure that you understand the billing method and prices of ApsaraVideo VOD before you call this operation. You are charged storage fees after you upload media files to ApsaraVideo VOD. For more information, see [Billing of media asset storage](~~188308#section_e97_xrp_mzz~~). If you have activated the acceleration service, you are charged acceleration fees when you upload media files to ApsaraVideo VOD. For more information, see [Billing of acceleration traffic](~~188310#section_sta_zm2_tsv~~).**
//
//		- You can call this operation only to obtain the upload URLs and credentials for media files and create media assets in ApsaraVideo VOD. You cannot call this operation to upload media files. For more information about how to upload media files by calling API operations, see [Upload media files by calling API operations](https://help.aliyun.com/document_detail/476208.html).
//
//		- If the upload credential expires after 3,000 seconds, you can call the CreateUploadAttachedMedia operation again to obtain a new upload URL and a new upload credential.
//
//		- You can configure a callback to receive an [AttachedMediaUploadComplete](https://help.aliyun.com/document_detail/103250.html) event notification to determine whether the upload is successful.
//
//		- You must obtain a URL and a credential before you upload a media file to ApsaraVideo VOD. ApsaraVideo VOD supports multiple upload methods. Each method has different requirements on upload URLs and credentials. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// @param request - CreateUploadAttachedMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadAttachedMediaResponse
func (client *Client) CreateUploadAttachedMediaWithContext(ctx context.Context, request *CreateUploadAttachedMediaRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadAttachedMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.CateIds) {
		query["CateIds"] = request.CateIds
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileSize) {
		query["FileSize"] = request.FileSize
	}

	if !dara.IsNil(request.MediaExt) {
		query["MediaExt"] = request.MediaExt
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
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
		Action:      dara.String("CreateUploadAttachedMedia"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUploadAttachedMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a URL and a credential for uploading an image.
//
// Description:
//
//	  **Make sure that you understand the billing method and price of ApsaraVideo VOD before you call this operation. You are charged storage fees after you upload media files to ApsaraVideo VOD. For more information, see [Billing of media asset storage](~~188308#section_e97_xrp_mzz~~). If you have activated the acceleration service, you are charged acceleration fees when you upload media files to ApsaraVideo VOD. For more information, see [Billing of acceleration traffic](~~188310#section_sta_zm2_tsv~~).**
//
//		- You must obtain a URL and a credential before you upload an image to ApsaraVideo VOD. ApsaraVideo VOD provides multiple upload methods. You can upload files by using server upload SDKs, client upload SDKs, URLs, Object Storage Service (OSS) API, or OSS SDKs. Each upload method has different requirements for obtaining upload URLs and credentials. For more information, see the "Usage notes" section of the [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html) topic.
//
//		- You cannot refresh the upload URL or credential when you upload images. If the image upload credential expires, you can call this operation to obtain a new upload URL and credential. By default, the validity period of an image upload credential is 3,000 seconds.
//
//		- You can call the [CreateUploadAttachedMedia](https://help.aliyun.com/document_detail/98467.html) operation to upload image watermarks.
//
//		- You can configure a callback for [ImageUploadComplete](https://help.aliyun.com/document_detail/91968.html) to receive notifications about the image upload status.
//
// @param request - CreateUploadImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadImageResponse
func (client *Client) CreateUploadImageWithContext(ctx context.Context, request *CreateUploadImageRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageExt) {
		query["ImageExt"] = request.ImageExt
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.OriginalFileName) {
		query["OriginalFileName"] = request.OriginalFileName
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
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
		Action:      dara.String("CreateUploadImage"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUploadImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an upload URL and an upload credential for uploading an audio or video file and generates the audio or video ID. ApsaraVideo VOD issues upload URLs and credentials to perform authorization and ensure security. This prevents unauthorized users from uploading media files. ApsaraVideo VOD generates media IDs, video IDs, and image IDs together with upload URLs and credentials. Media IDs are used in lifecycle management and media processing.
//
// Description:
//
//	  **Make sure that you understand the billing method and prices of ApsaraVideo VOD before you call this operation. You are charged storage fees after you upload media files to ApsaraVideo VOD. For more information, see [Billing of media asset storage](~~188308#section_e97_xrp_mzz~~). If you have activated the acceleration service, you are charged acceleration fees when you upload media files to ApsaraVideo VOD. For more information, see [Billing of acceleration traffic](~~188310#section_sta_zm2_tsv~~).**
//
//		- You can call this operation to obtain upload URLs and credentials for video and audio files. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
//		- You can call this operation only to obtain the upload URLs and credentials for media files and create media assets in ApsaraVideo VOD. You cannot call this operation to upload media files. For more information about how to upload media files by calling API operations, see [Upload media files by calling API operations](https://help.aliyun.com/document_detail/476208.html).
//
//		- If the upload credential expires, call the [RefreshUploadVideo](~~RefreshUploadVideo~~) operation to obtain a new upload credential. The default validity period of an upload credential is 3,000 seconds.
//
//		- You can configure a callback to receive an event notification when an audio or video file is uploaded. Alternatively, after you upload an audio or video file, you can call the [GetMezzanineInfo](https://help.aliyun.com/document_detail/59624.html) operation to determine whether the upload is successful. For more information, see [Overview](https://help.aliyun.com/document_detail/55396.html).
//
//		- The value of the VideoId parameter that is returned after you call this operation can be used for media processing or the lifecycle management of media assets.
//
//		- You must obtain a URL and a credential before you upload a media file to ApsaraVideo VOD. ApsaraVideo VOD supports multiple upload methods. Each method has different requirements on upload URLs and credentials. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// @param request - CreateUploadVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadVideoResponse
func (client *Client) CreateUploadVideoWithContext(ctx context.Context, request *CreateUploadVideoRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadVideoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileSize) {
		query["FileSize"] = request.FileSize
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TemplateGroupId) {
		query["TemplateGroupId"] = request.TemplateGroupId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUploadVideo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUploadVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decrypts the ciphertext specified by CiphertextBlob in the Key Management Service (KMS) data key.
//
// @param request - DecryptKMSDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptKMSDataKeyResponse
func (client *Client) DecryptKMSDataKeyWithContext(ctx context.Context, request *DecryptKMSDataKeyRequest, runtime *dara.RuntimeOptions) (_result *DecryptKMSDataKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CipherText) {
		query["CipherText"] = request.CipherText
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
		Action:      dara.String("DecryptKMSDataKey"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DecryptKMSDataKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the information about one or more images that are submitted for AI processing.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)*	- and **China (Shanghai)**.
//
//		- This operation deletes only information about images that are submitted for AI processing. The image files are not deleted.
//
// @param request - DeleteAIImageInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAIImageInfosResponse
func (client *Client) DeleteAIImageInfosWithContext(ctx context.Context, request *DeleteAIImageInfosRequest, runtime *dara.RuntimeOptions) (_result *DeleteAIImageInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIImageInfoIds) {
		query["AIImageInfoIds"] = request.AIImageInfoIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAIImageInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAIImageInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an AI template.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- You cannot delete an AI template that is set as the default template.
//
// @param request - DeleteAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAITemplateResponse
func (client *Client) DeleteAITemplateWithContext(ctx context.Context, request *DeleteAITemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAITemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAITemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAITemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// Description:
//
// Application with resources can not be deleted.
//
// @param request - DeleteAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInfoResponse
func (client *Client) DeleteAppInfoWithContext(ctx context.Context, request *DeleteAppInfoRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes auxiliary media assets from ApsaraVideo VOD. You can delete multiple auxiliary media assets such as watermark images, subtitle files, and materials in a batch.
//
// Description:
//
//	  **This operation physically deletes auxiliary media assets. You cannot recover the auxiliary media assets that you deleted. Exercise caution when you call this operation.**
//
//		- You can delete a maximum of 20 auxiliary media assets in one request.
//
// @param request - DeleteAttachedMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAttachedMediaResponse
func (client *Client) DeleteAttachedMediaWithContext(ctx context.Context, request *DeleteAttachedMediaRequest, runtime *dara.RuntimeOptions) (_result *DeleteAttachedMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAttachedMedia"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAttachedMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a category and its subcategories.
//
// Description:
//
//	  **After you call this operation to delete a category, all subcategories including level 2 and level 3 categories are deleted at the same time. Exercise caution when you call this operation.**
//
//		- If you have classified specific media resources to a category, the category names labeled on these media resources are automatically deleted when you delete the category.
//
// @param request - DeleteCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithContext(ctx context.Context, request *DeleteCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCategory"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the information about animated stickers.
//
// Description:
//
// > This operation deletes only the information about animated stickers, but not the animated stickers themselves.
//
// @param request - DeleteDynamicImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDynamicImageResponse
func (client *Client) DeleteDynamicImageWithContext(ctx context.Context, request *DeleteDynamicImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteDynamicImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicImageIds) {
		query["DynamicImageIds"] = request.DynamicImageIds
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDynamicImage"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDynamicImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes online editing projects.
//
// Description:
//
//	You can call this operation to delete multiple online editing projects at a time.
//
// @param request - DeleteEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEditingProjectResponse
func (client *Client) DeleteEditingProjectWithContext(ctx context.Context, request *DeleteEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
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
		Action:      dara.String("DeleteEditingProject"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEditingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除剪辑资源
//
// @param request - DeleteEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEditingProjectMaterialsResponse
func (client *Client) DeleteEditingProjectMaterialsWithContext(ctx context.Context, request *DeleteEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEditingProjectMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaterialIds) {
		query["MaterialIds"] = request.MaterialIds
	}

	if !dara.IsNil(request.MaterialType) {
		query["MaterialType"] = request.MaterialType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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
		Action:      dara.String("DeleteEditingProjectMaterials"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEditingProjectMaterialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes uploaded images and video snapshots that are automatically captured.
//
// Description:
//
//	  **After you call this operation to delete an image, the source file is permanently deleted and cannot be recovered. Exercise caution when you call this operation.**
//
//		- If some images are cached on Alibaba Cloud CDN points of presence (POPs), the image URLs do not immediately become invalid.
//
//		- You can call this operation to delete uploaded images and video snapshots.
//
// @param request - DeleteImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageResponse
func (client *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteImageType) {
		query["DeleteImageType"] = request.DeleteImageType
	}

	if !dara.IsNil(request.ImageIds) {
		query["ImageIds"] = request.ImageIds
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.ImageURLs) {
		query["ImageURLs"] = request.ImageURLs
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImage"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the callback method, callback URL, and event type of an event notification.
//
// Description:
//
// > For more information, see [Overview](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - DeleteMessageCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageCallbackResponse
func (client *Client) DeleteMessageCallbackWithContext(ctx context.Context, request *DeleteMessageCallbackRequest, runtime *dara.RuntimeOptions) (_result *DeleteMessageCallbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMessageCallback"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMessageCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more source files at a time.
//
// Description:
//
// All media processing operations in ApsaraVideo VOD, such as transcoding, snapshot capture, and content moderation, are performed based on source files. If you delete the source files, you cannot perform media processing operations. Exercise caution when you call this operation.
//
// @param request - DeleteMezzaninesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMezzaninesResponse
func (client *Client) DeleteMezzaninesWithContext(ctx context.Context, request *DeleteMezzaninesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMezzaninesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.VideoIds) {
		query["VideoIds"] = request.VideoIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMezzanines"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMezzaninesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the parts generated during an upload.
//
// Description:
//
//	  During multipart upload, useless parts may be retained if the upload fails. These useless parts are automatically deleted after 7 days. You can call this operation to delete the generated parts after the upload is successful or fails.
//
//		- This operation does not delete the source file or transcoded file, but deletes only the parts generated during the upload.
//
//		- If you call the [DeleteVideo](https://help.aliyun.com/document_detail/52837.html) operation, the entire video file is deleted, including the generated parts.
//
// @param request - DeleteMultipartUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultipartUploadResponse
func (client *Client) DeleteMultipartUploadWithContext(ctx context.Context, request *DeleteMultipartUploadRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultipartUploadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultipartUpload"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultipartUploadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more video or audio streams and their storage files at a time.
//
// @param request - DeleteStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStreamResponse
func (client *Client) DeleteStreamWithContext(ctx context.Context, request *DeleteStreamRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStream"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more transcoding templates from a transcoding template group or forcibly deletes a transcoding template group.
//
// Description:
//
//	  You cannot call this operation to delete the default transcoding template. You can delete the transcoding template when it is no longer specified as the default one.
//
//		- For security purposes, you cannot add, modify, or delete transcoding templates in a transcoding template group that is locked. To check whether a transcoding template group is locked, call the [GetTranscodeTemplateGroup](~~GetTranscodeTemplateGroup~~) operation and obtain the Locked parameter from the response. To modify transcoding templates within a locked transcoding template group, you must call the [UpdateTranscodeTemplateGroup](~~UpdateTranscodeTemplateGroup~~) operation to unlock the transcoding template group first.
//
// @param request - DeleteTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTranscodeTemplateGroupResponse
func (client *Client) DeleteTranscodeTemplateGroupWithContext(ctx context.Context, request *DeleteTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteTranscodeTemplateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelGroup) {
		query["ForceDelGroup"] = request.ForceDelGroup
	}

	if !dara.IsNil(request.TranscodeTemplateGroupId) {
		query["TranscodeTemplateGroupId"] = request.TranscodeTemplateGroupId
	}

	if !dara.IsNil(request.TranscodeTemplateIds) {
		query["TranscodeTemplateIds"] = request.TranscodeTemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTranscodeTemplateGroup"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTranscodeTemplateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more videos at a time, including their mezzanine files, transcoded stream files, and thumbnail snapshots.
//
// Description:
//
//	  This operation physically deletes videos. Deleted videos cannot be recovered. Exercise caution when you call this operation.
//
//		- You can call this operation to delete multiple videos at a time.
//
//		- When you delete a video, its source file, transcoded stream file, and thumbnail screenshot are also deleted. However, the Alibaba Cloud Content Delivery Network (CDN) cache is not refreshed simultaneously. You can use the refresh feature in the ApsaraVideo VOD console to clear garbage data on CDN nodes. For more information, see [Refresh and prefetch](https://help.aliyun.com/document_detail/86098.html).
//
// @param request - DeleteVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVideoResponse
func (client *Client) DeleteVideoWithContext(ctx context.Context, request *DeleteVideoRequest, runtime *dara.RuntimeOptions) (_result *DeleteVideoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VideoIds) {
		query["VideoIds"] = request.VideoIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVideo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a domain name for CDN from ApsaraVideo VOD.
//
// Description:
//
// > 	- This operation is available only in the **China (Shanghai)*	- region.
//
// > 	- After a domain name for CDN is removed from ApsaraVideo VOD, the domain name becomes unavailable. Proceed with caution. We recommend that you restore the A record at your DNS service provider before you remove the domain name for CDN.
//
// > 	- After you call this operation to remove a domain name for CDN from ApsaraVideo VOD, all records that are related to the domain name are deleted. If you only want to disable a domain name for CDN, call the [BatchStopVodDomain](https://help.aliyun.com/document_detail/120208.html) operation.
//
// @param request - DeleteVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodDomainResponse
func (client *Client) DeleteVodDomainWithContext(ctx context.Context, request *DeleteVodDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVodDomain"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVodDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configurations of a domain name for CDN.
//
// Description:
//
// >
//
//   - This operation is available only in the **China (Shanghai)*	- region.
//
//   - After the configurations of a domain name for CDN are deleted, the domain name becomes unavailable. We recommend that you restore the A record at your DNS service provider before you delete the configurations of the domain name for CDN.
//
//   - After you call this operation to remove a domain name for CDN from ApsaraVideo VOD, all records that are related to the domain name are deleted. If you only want to disable a domain name for CDN, call the [BatchStopVodDomain](https://help.aliyun.com/document_detail/120208.html) operation.
//
// @param request - DeleteVodSpecificConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodSpecificConfigResponse
func (client *Client) DeleteVodSpecificConfigWithContext(ctx context.Context, request *DeleteVodSpecificConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodSpecificConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVodSpecificConfig"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVodSpecificConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a snapshot template.
//
// @param request - DeleteVodTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodTemplateResponse
func (client *Client) DeleteVodTemplateWithContext(ctx context.Context, request *DeleteVodTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VodTemplateId) {
		query["VodTemplateId"] = request.VodTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVodTemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVodTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image watermark or text watermark template.
//
// Description:
//
//	  **After you delete an image watermark template, the source watermark file is physically deleted and cannot be restored. Exercise caution when you call this operation.**
//
//		- You cannot delete the default watermark template. To delete a default watermark template, call the [SetDefaultWatermark](~~SetDefaultWatermark~~) operation to set another watermark template as the default one.
//
// @param request - DeleteWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWatermarkResponse
func (client *Client) DeleteWatermarkWithContext(ctx context.Context, request *DeleteWatermarkRequest, runtime *dara.RuntimeOptions) (_result *DeleteWatermarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WatermarkId) {
		query["WatermarkId"] = request.WatermarkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWatermark"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution of media asset data by time. The maximum time range to query is 6 months.
//
// Description:
//
//	  This operation is available only in the China (Shanghai) region.
//
//		- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the previous 7 days. If you set both the parameters, the request returns the data collected within the specified time range.
//
// @param request - DescribeMediaDistributionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMediaDistributionResponse
func (client *Client) DescribeMediaDistributionWithContext(ctx context.Context, request *DescribeMediaDistributionRequest, runtime *dara.RuntimeOptions) (_result *DescribeMediaDistributionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StorageClass) {
		query["StorageClass"] = request.StorageClass
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMediaDistribution"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMediaDistributionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries daily playback statistics on top videos, including video views, unique visitors, and total playback duration.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- You can query playback statistics on top 1,000 videos at most on a specified day. By default, top videos are sorted in descending order based on video views.
//
//		- You can call this operation to query only playback statistics collected on videos that are played by using ApsaraVideo Player SDKs.
//
//		- Playback statistics for the previous day are generated at 09:00 on the current day, in UTC+8.
//
//		- You can query data that is generated since January 1, 2018. The maximum time range to query is 180 days.
//
// @param request - DescribePlayTopVideosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayTopVideosResponse
func (client *Client) DescribePlayTopVideosWithContext(ctx context.Context, request *DescribePlayTopVideosRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayTopVideosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		query["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePlayTopVideos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlayTopVideosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on average playback each day in a specified time range.
//
// Description:
//
// > 	- This operation is available only in the **China (Shanghai)*	- region.
//
// > 	- You can call this operation to query only playback statistics collected on videos that are played by using ApsaraVideo Player SDKs.
//
// > 	- Playback statistics for the previous day are generated at 09:00 on the current day, in UTC+8.
//
// > 	- You can query data that is generated since January 1, 2018. The maximum time range to query is 180 days.
//
// @param request - DescribePlayUserAvgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayUserAvgResponse
func (client *Client) DescribePlayUserAvgWithContext(ctx context.Context, request *DescribePlayUserAvgRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayUserAvgResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePlayUserAvg"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlayUserAvgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the daily playback statistics in a specified time range. The playback statistics include the total number of views, total number of viewers, total playback duration, and playback duration distribution.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- You can call this operation to query only playback statistics collected on videos that are played by using ApsaraVideo Player SDKs.
//
//		- Playback statistics for the current day are generated at 09:00 (UTC+8) on the next day.
//
//		- You can query data that is generated since January 1, 2018. The maximum time range to query is 180 days.
//
// @param request - DescribePlayUserTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayUserTotalResponse
func (client *Client) DescribePlayUserTotalWithContext(ctx context.Context, request *DescribePlayUserTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayUserTotalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePlayUserTotal"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlayUserTotalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries daily playback statistics on a video in the specified time range.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- You can call this operation to query only playback statistics collected on videos that are played by using ApsaraVideo Player SDKs.
//
//		- Playback statistics for the current day are generated at 09:00 (UTC+8) on the next day.
//
//		- You can query only data in the last 730 days. The maximum time range to query is 180 days.
//
// @param request - DescribePlayVideoStatisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayVideoStatisResponse
func (client *Client) DescribePlayVideoStatisWithContext(ctx context.Context, request *DescribePlayVideoStatisRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayVideoStatisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePlayVideoStatis"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlayVideoStatisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on video AI of different types, such as automated review and media fingerprinting.
//
// Description:
//
// > 	- This operation is available only in the **China (Shanghai)*	- region.
//
// >	- If the time range to query is less than or equal to seven days, the system returns the statistics collected on an hourly basis. If the time range to query is greater than seven days, the system returns the statistics collected on a daily basis. The maximum time range that you can specify to query is 31 days.
//
// @param request - DescribeVodAIDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodAIDataResponse
func (client *Client) DescribeVodAIDataWithContext(ctx context.Context, request *DescribeVodAIDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodAIDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIType) {
		query["AIType"] = request.AIType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodAIData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodAIDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates of a specified domain name for CDN or all the domain names for CDN within your Alibaba Cloud account.
//
// Description:
//
// >  This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - DescribeVodCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodCertificateListResponse
func (client *Client) DescribeVodCertificateListWithContext(ctx context.Context, request *DescribeVodCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodCertificateList"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodCertificateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth for one or more specified domain names for CDN.
//
// Description:
//
// If you specify neither the StartTime parameter nor the EndTime parameter, the data in the last 24 hours is queried. Alternatively, you can specify both the StartTime and EndTime parameters to query data that is generated in the specified duration. You can query data for the last 90 days at most.
//
// @param request - DescribeVodDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainBpsDataResponse
func (client *Client) DescribeVodDomainBpsDataWithContext(ctx context.Context, request *DescribeVodDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainBpsData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data by protocol.
//
// Description:
//
// You can call this API operation up to 20 times per second per account. If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range. Time granularity
//
// The time granularity supported by Interval, the maximum time period within which historical data is available, and the data delay vary based on the time range to query, as described in the following table.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |15 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|3 to 4 hours|
//
// |1 day|90 days|366 days|4 hours in most cases, not more than 24 hours|
//
// @param request - DescribeVodDomainBpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainBpsDataByLayerResponse
func (client *Client) DescribeVodDomainBpsDataByLayerWithContext(ctx context.Context, request *DescribeVodDomainBpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainBpsDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainBpsDataByLayer"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainBpsDataByLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate information about an accelerated domain name.
//
// Description:
//
// This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - DescribeVodDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainCertificateInfoResponse
func (client *Client) DescribeVodDomainCertificateInfoWithContext(ctx context.Context, request *DescribeVodDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainCertificateInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainCertificateInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainCertificateInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a domain name for CDN. You can query the configurations of multiple features at a time.
//
// Description:
//
// > This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - DescribeVodDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainConfigsResponse
func (client *Client) DescribeVodDomainConfigsWithContext(ctx context.Context, request *DescribeVodDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainConfigs"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about a specified domain name for CDN.
//
// Description:
//
// > This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - DescribeVodDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainDetailResponse
func (client *Client) DescribeVodDomainDetailWithContext(ctx context.Context, request *DescribeVodDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainDetail"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the byte hit ratios of accelerated domain names. Byte hit ratios are measured in percentage.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 500 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 24 hours is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay when you do not set `Interval`.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|Time range per query &#x3C; 3 days|93 days|15 minutes|
//
// |1 hour|3 days ≤ Time range per query &#x3C; 31 days|186 days|3 to 4 hours|
//
// |1 day|31 days ≤ Time span of a single query ≤ 366 days|366 days|4 hours in most cases, not more than 24 hours|
//
// @param request - DescribeVodDomainHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainHitRateDataResponse
func (client *Client) DescribeVodDomainHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainHitRateData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainHitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the CDN access logs for a domain name, including the log path.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- For more information about the log format and latency, see [Download logs](https://help.aliyun.com/document_detail/86099.html).
//
//		- If you specify neither the StartTime parameter nor the EndTime parameter, the log data in the last 24 hours is queried.
//
//		- You can specify both the StartTime and EndTime parameters to query the log data that is generated in the specified time range.
//
// @param request - DescribeVodDomainLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainLogResponse
func (client *Client) DescribeVodDomainLogWithContext(ctx context.Context, request *DescribeVodDomainLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainLog"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the 95th percentile bandwidth data of an accelerated domain name.
//
// @param request - DescribeVodDomainMax95BpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainMax95BpsDataResponse
func (client *Client) DescribeVodDomainMax95BpsDataWithContext(ctx context.Context, request *DescribeVodDomainMax95BpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainMax95BpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainMax95BpsData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainMax95BpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) for one or more accelerated domain names. Data is collected every 5 minutes. You can query data collected in the last 90 days.
//
// Description:
//
// This operation is available only in the China (Shanghai) region.
//
//   - You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|3 to 4 hours|
//
// |1 day|366 days|366 days|4 to 24 hours|
//
// ---
//
// @param request - DescribeVodDomainQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainQpsDataResponse
func (client *Client) DescribeVodDomainQpsDataWithContext(ctx context.Context, request *DescribeVodDomainQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainQpsData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainQpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data for one or more accelerated domains. The minimum time granularity is 1 minute. The minimum data latency is 5 minutes. You can query data in the last 186 days. Compared with the DescribeVodDomainBpsData operation, this operation provides a smaller time granularity, lower data latency, and allows you to query historical data within a shorter time period.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 500 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 1 hour is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|Time range per query ≤ 1 hour|7 days|5 minutes|
//
// |5 minutes|1 Hour &#x3C; Time range per query ≤ 3 days|93 days|15 minutes|
//
// |1 hour|3 days &#x3C; Time range per query ≤ 31 days|186 days|3 to 4 hours|
//
// @param request - DescribeVodDomainRealTimeBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeBpsDataResponse
func (client *Client) DescribeVodDomainRealTimeBpsDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainRealTimeBpsData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainRealTimeBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the byte hit ratio for one or more accelerated domains. The minimum time granularity is 1 minute. The minimum data latency is 5 minutes. You can query data in the last 186 days.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 100 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 1 hour is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|Time range per query ≤ 1 hour|7 days|5 minutes|
//
// |5 minutes|1 Hour &#x3C; Time range per query ≤ 3 days|93 days|15 minutes|
//
// |1 hour|3 days &#x3C; Time range per query ≤ 31 days|186 days|3 to 4 hours|
//
// @param request - DescribeVodDomainRealTimeByteHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeByteHitRateDataResponse
func (client *Client) DescribeVodDomainRealTimeByteHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeByteHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeByteHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainRealTimeByteHitRateData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries real-time monitoring data of one or more accelerated domain names.
//
// Description:
//
// You can query data within the last seven days. Data is collected every minute. You can call this API operation up to 10 times per second per account.
//
// @param request - DescribeVodDomainRealTimeDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeDetailDataResponse
func (client *Client) DescribeVodDomainRealTimeDetailDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeDetailDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeDetailDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainRealTimeDetailData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainRealTimeDetailDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number of HTTP status codes and proportion of each HTTP status code for one or more accelerated domains. The minimum time granularity is 1 minute. The minimum data latency is 5 minutes. You can query data in the last 186 days.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 100 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 1 hour is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Time range per query|Historical data available (days)|Data latency|
//
// |---|---|---|---|
//
// |1 minute|Time range per query ≤ 1 hour|7 days|5 minutes|
//
// |5 minutes|1 hour &#x3C; Time range per query &#x3C; 3 days|93 days|15 minutes|
//
// |1 hour|3 days ≤ Time range per query &#x3C; 31 days|186 days|3 to 4 hours|
//
// @param request - DescribeVodDomainRealTimeHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeVodDomainRealTimeHttpCodeDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainRealTimeHttpCodeData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) for one or more accelerated domains. The minimum time granularity is 1 minute. The minimum data latency is 5 minutes. You can query data in the last 186 days.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 500 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 1 hour is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|Time range per query ≤ 1 hour|7 days|5 minutes|
//
// |5 minutes|1 Hour &#x3C; Time range per query ≤ 3 days|93 days|15 minutes|
//
// |1 hour|3 days &#x3C; Time range per query ≤ 31 days|186 days|3 to 4 hours|
//
// @param request - DescribeVodDomainRealTimeQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeQpsDataResponse
func (client *Client) DescribeVodDomainRealTimeQpsDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainRealTimeQpsData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainRealTimeQpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request hit ratio data for one or more accelerated domain names. The minimum time granularity is 1 minute. The minimum data latency is 5 minutes. You can query data in the last 186 days.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 100 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 1 hour is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
//   - By default, the POST method is used for Go. To use the FET method, you must declare `request.Method="GET"`.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|Time range per query ≤ 1 hour|7 days|5 minutes|
//
// |5 minutes|1 hour &#x3C; Time range per query &#x3C; 3 days|93 days|15 minutes|
//
// |1 hour|3 days ≤ Time range per query &#x3C; 31 days|186 days|3 to 4 hours|
//
// @param request - DescribeVodDomainRealTimeReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeReqHitRateDataResponse
func (client *Client) DescribeVodDomainRealTimeReqHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeReqHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainRealTimeReqHitRateData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic data for one or more accelerated domains. The minimum time granularity is 1 minute. The minimum data latency is 5 minutes. You can query data in the last 186 days. Compared with the DescribeVodDomainTrafficData operation, this operation provides a smaller time granularity, lower data latency, and allows you to query historical data within a shorter time period.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 100 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 1 hour is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|Time range per query ≤ 1 hour|7 days|5 minutes|
//
// |5 minutes|1 Hour &#x3C; Time range per query ≤ 3 days|93 days|15 minutes|
//
// |1 hour|3 days &#x3C; Time range per query ≤ 31 days|186 days|3 to 4 hours|
//
// @param request - DescribeVodDomainRealTimeTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeTrafficDataResponse
func (client *Client) DescribeVodDomainRealTimeTrafficDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainRealTimeTrafficData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainRealTimeTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the byte hit ratio for one or more accelerated domains. Request hit ratios are measured in percentage.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 500 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 24 hours is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay when you do not set `Interval`.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|Time range per query &#x3C; 3 days|93 days|15 minutes|
//
// |1 hour|3 days ≤ Time range per query &#x3C; 31 days|186 days|3 to 4 hours|
//
// |1 day|31 days ≤ Time range per query ≤ 90 days|366 days|4 hours in most cases, not more than 24 hours|
//
// @param request - DescribeVodDomainReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainReqHitRateDataResponse
func (client *Client) DescribeVodDomainReqHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainReqHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainReqHitRateData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainReqHitRateDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data during back-to-origin routing for one or more accelerated domain names.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 500 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 24 hours is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay when you do not set `Interval`.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|Time range per query &#x3C; 3 days|93 days|15 minutes|
//
// |1 hour|3 days ≤ Time range per query &#x3C; 31 days|186 days|3 to 4 hours|
//
// |1 day|31 days ≤ Time span of a single query ≤ 366 days|366 days|4 hours in most cases, not more than 24 hours|
//
// @param request - DescribeVodDomainSrcBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainSrcBpsDataResponse
func (client *Client) DescribeVodDomainSrcBpsDataWithContext(ctx context.Context, request *DescribeVodDomainSrcBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainSrcBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainSrcBpsData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainSrcBpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries origin traffic data for accelerated domain names in ApsaraVideo VOD. The traffic is measured in bytes.
//
// Description:
//
// This operation is available only in the **China (Shanghai)*	- region.
//
//   - ApsaraVideo VOD stores the origin traffic data for 90 days before the data is deleted.
//
//   - If you do not set the `StartTime` or `EndTime` parameter, the request returns the data collected in the last 24 hours. If you set both the `StartTime` and `EndTime` parameters, the request returns the data collected within the specified time range.
//
//   - You can specify a maximum of 500 domain names in a request. Separate multiple domain names with commas (,). If you specify multiple domain names in a request, aggregation results are returned.
//
// ### Time granularity
//
// The time granularity supported by the Interval parameter varies based on the time range per query specified by using `StartTime` and `EndTime`. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Time range per query (days)|Historical data available (days)|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|(0, 3\\]|93|15 minutes|
//
// |1 hour|(3, 31\\]|186|4 hours|
//
// |1 day|(31, 366\\]|366|04:00 on the next day|
//
// @param request - DescribeVodDomainSrcTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainSrcTrafficDataResponse
func (client *Client) DescribeVodDomainSrcTrafficDataWithContext(ctx context.Context, request *DescribeVodDomainSrcTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainSrcTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainSrcTrafficData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainSrcTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic data for one or more accelerated domains. The minimum time granularity is 5 minutes. You can query data in the last 366 days. Compared with the DescribeVodDomainRealTimeTrafficData operation, this operation provides a greater time granularity, higher data latency, but allows you to query historical data within a longer time period.
//
// Description:
//
// This operation is supported only in the **China (Shanghai)*	- region.
//
//   - You can specify a maximum of 500 accelerated domain names.
//
//   - If you specify neither `StartTime` nor `EndTime`, the data of the last 24 hour is queried. You can specify both `StartTime` and `EndTime` parameters to query data of a specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the `StartTime` and `EndTime` parameters. The following table describes the time period within which historical data is available and the data delay when you do not set `Interval`.
//
// |Time granularity|Time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|Time range per query &#x3C; 3 days|93 days|15 minutes|
//
// |1 hour|3 days ≤ Time range per query &#x3C; 31 days|186 days|3 to 4 hours|
//
// |1 day|31 days ≤ Time range per query ≤ 366 days|366 days|4 hours in most cases, not more than 24 hours|
//
// @param request - DescribeVodDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainTrafficDataResponse
func (client *Client) DescribeVodDomainTrafficDataWithContext(ctx context.Context, request *DescribeVodDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainTrafficData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic or bandwidth data of one or more accelerated domain names.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- You can specify up to 100 accelerated domain names in a request. Separate multiple domain names with commas (,). If you do not specify an accelerated domain name, the data of all accelerated domain names within your Alibaba Cloud account is returned.
//
//		- You can query data in the last year. The maximum time range that can be queried is three months. If you specify a time range of one to three days, the system returns data on an hourly basis. If you specify a time range of four days or more, the system returns data on a daily basis.
//
// @param request - DescribeVodDomainUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainUsageDataResponse
func (client *Client) DescribeVodDomainUsageDataWithContext(ctx context.Context, request *DescribeVodDomainUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodDomainUsageData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodDomainUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 点播云剪辑用量查询
//
// @param request - DescribeVodEditingUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodEditingUsageDataResponse
func (client *Client) DescribeVodEditingUsageDataWithContext(ctx context.Context, request *DescribeVodEditingUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodEditingUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodEditingUsageData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodEditingUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the playback statistics based on the media ID. You can call this operation to query information such as the number of visits, average video views per viewer, total number of views, average playback duration per viewer, and total playback duration.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- Only playback data in ApsaraVideo Player SDK is collected.
//
//		- You can query only data within the last 30 days.
//
//		- Before you call this operation, make sure that the following requirements are met:
//
//	    	- ApsaraVideo Player SDK for Android or iOS
//
//	        	- ApsaraVideo Player SDK for Android or iOS V5.4.9.2 or later is used.
//
//	        	- A license for ApsaraVideo Player SDK is obtained. For more information, see [Manage licenses](https://help.aliyun.com/document_detail/469166.html).
//
//	        	- The log reporting feature is enabled. By default, the feature is enabled for ApsaraVideo Player SDKs. For more information, see [Integrate ApsaraVideo Player SDK for Android](~~311525#section-dc4-gp6-xk2~~) and [Integrate ApsaraVideo Player SDK for iOS](~~313855#section-cmf-k7d-jg5~~).
//
//	    	- ApsaraVideo Player SDK for Web
//
//	        	- ApsaraVideo Player SDK for Web V2.16.0 or later is used.
//
//	        	- A license for **playback quality monitoring*	- is obtained. To apply for the license, [submit a request on Yida to enable value-added features for ApsaraVideo Player SDK for Web](https://yida.alibaba-inc.com/o/webplayer#/). For more information, see the description of the `license` parameter in the [API operations](~~125572#section-3ty-gwp-6pa~~) topic.
//
//	        	- The log reporting feature is enabled. By default, the feature is enabled for ApsaraVideo Player SDKs.
//
// @param request - DescribeVodMediaPlayDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodMediaPlayDataResponse
func (client *Client) DescribeVodMediaPlayDataWithContext(ctx context.Context, request *DescribeVodMediaPlayDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodMediaPlayDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OrderName) {
		query["OrderName"] = request.OrderName
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.Os) {
		query["Os"] = request.Os
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlayDate) {
		query["PlayDate"] = request.PlayDate
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.TerminalType) {
		query["TerminalType"] = request.TerminalType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodMediaPlayData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodMediaPlayDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询播放器指标数据
//
// @param request - DescribeVodPlayerCollectDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodPlayerCollectDataResponse
func (client *Client) DescribeVodPlayerCollectDataWithContext(ctx context.Context, request *DescribeVodPlayerCollectDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodPlayerCollectDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Os) {
		query["Os"] = request.Os
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TerminalType) {
		query["TerminalType"] = request.TerminalType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodPlayerCollectData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodPlayerCollectDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询播放器指标数据
//
// @param request - DescribeVodPlayerDimensionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodPlayerDimensionDataResponse
func (client *Client) DescribeVodPlayerDimensionDataWithContext(ctx context.Context, request *DescribeVodPlayerDimensionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodPlayerDimensionDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodPlayerDimensionData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodPlayerDimensionDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询播放器指标数据
//
// @param request - DescribeVodPlayerMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodPlayerMetricDataResponse
func (client *Client) DescribeVodPlayerMetricDataWithContext(ctx context.Context, request *DescribeVodPlayerMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodPlayerMetricDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Os) {
		query["Os"] = request.Os
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TerminalType) {
		query["TerminalType"] = request.TerminalType
	}

	if !dara.IsNil(request.Top) {
		query["Top"] = request.Top
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodPlayerMetricData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodPlayerMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data by Internet service provider (ISP) and region.
//
// Description:
//
// The data is collected every 5 minutes. You can call this API operation up to 20 times per second per account. Time granularity
//
// The time granularity supported by Interval, the maximum time period within which historical data is available, and the data delay vary based on the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|1 hour|93 days|15 minutes|
//
// @param request - DescribeVodRangeDataByLocateAndIspServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodRangeDataByLocateAndIspServiceResponse
func (client *Client) DescribeVodRangeDataByLocateAndIspServiceWithContext(ctx context.Context, request *DescribeVodRangeDataByLocateAndIspServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodRangeDataByLocateAndIspServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodRangeDataByLocateAndIspService"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodRangeDataByLocateAndIspServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum number and remaining number of requests to refresh or prefetch files on the current day. You can prefetch files based on URLs and refresh files based on URLs or directories.
//
// Description:
//
// > 	- This operation is available only in the **China (Shanghai)*	- region.
//
// > 	- You can call the [RefreshVodObjectCaches](https://help.aliyun.com/document_detail/69215.html) operation to refresh content and the [PreloadVodObjectCaches](https://help.aliyun.com/document_detail/69211.html) operation to prefetch content.
//
// @param request - DescribeVodRefreshQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodRefreshQuotaResponse
func (client *Client) DescribeVodRefreshQuotaWithContext(ctx context.Context, request *DescribeVodRefreshQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodRefreshQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodRefreshQuota"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodRefreshQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more refresh or prefetch tasks.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- If you do not specify the TaskId or ObjectPath parameter, the data in the last three days is returned on the first page. By default, one page displays a maximum of 20 entries. You can specify the TaskId and ObjectPath parameters at the same time.
//
// @param request - DescribeVodRefreshTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodRefreshTasksResponse
func (client *Client) DescribeVodRefreshTasksWithContext(ctx context.Context, request *DescribeVodRefreshTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodRefreshTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodRefreshTasks"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodRefreshTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates by domain name.
//
// @param request - DescribeVodSSLCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodSSLCertificateListResponse
func (client *Client) DescribeVodSSLCertificateListWithContext(ctx context.Context, request *DescribeVodSSLCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodSSLCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
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

	if !dara.IsNil(request.SearchKeyword) {
		query["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodSSLCertificateList"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodSSLCertificateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage of storage-related resources, including the storage volume and outbound traffic.
//
// Description:
//
// > 	- This operation is available only in the **China (Shanghai)*	- region.
//
// >	- If the time range to query is less than or equal to seven days, the system returns the statistics collected on an hourly basis. If the time range to query is greater than seven days, the system returns the statistics collected on a daily basis. The maximum time range that you can specify to query is 31 days.
//
// @param request - DescribeVodStorageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodStorageDataResponse
func (client *Client) DescribeVodStorageDataWithContext(ctx context.Context, request *DescribeVodStorageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodStorageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodStorageData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodStorageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage of tiered storage for media assets.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- If you specify a time range within 7 days, the request returns the data based on hours. If you specify a time range longer than 7 days, the request returns the data based on days. The maximum time range is 31 days.
//
// @param request - DescribeVodTieringStorageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodTieringStorageDataResponse
func (client *Client) DescribeVodTieringStorageDataWithContext(ctx context.Context, request *DescribeVodTieringStorageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodTieringStorageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StorageClass) {
		query["StorageClass"] = request.StorageClass
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodTieringStorageData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodTieringStorageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data retrieval from tiered storage.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- If you specify a time range within 7 days, the request returns the data based on hours. If you specify a time range longer than 7 days, the request returns the data based on days. The maximum time range is 31 days.
//
// @param request - DescribeVodTieringStorageRetrievalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodTieringStorageRetrievalDataResponse
func (client *Client) DescribeVodTieringStorageRetrievalDataWithContext(ctx context.Context, request *DescribeVodTieringStorageRetrievalDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodTieringStorageRetrievalDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StorageClass) {
		query["StorageClass"] = request.StorageClass
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodTieringStorageRetrievalData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodTieringStorageRetrievalDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the transcoding statistics.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- If the time range to query is less than or equal to seven days, the system returns the statistics collected on an hourly basis. If the time range to query is greater than seven days, the system returns the statistics collected on a daily basis. The maximum time range that you can specify to query is 31 days.
//
// @param request - DescribeVodTranscodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodTranscodeDataResponse
func (client *Client) DescribeVodTranscodeDataWithContext(ctx context.Context, request *DescribeVodTranscodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodTranscodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodTranscodeData"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodTranscodeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names for CDN within your Alibaba Cloud account.
//
// Description:
//
//	  You can filter domain names by name and status. Fuzzy match is supported for domain name-based query.
//
//		- This operation is available only in the China (Shanghai) region.
//
// @param request - DescribeVodUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodUserDomainsResponse
func (client *Client) DescribeVodUserDomainsWithContext(ctx context.Context, request *DescribeVodUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodUserDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodUserDomains"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodUserDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取域名Vip
//
// @param request - DescribeVodUserVipsByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodUserVipsByDomainResponse
func (client *Client) DescribeVodUserVipsByDomainWithContext(ctx context.Context, request *DescribeVodUserVipsByDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodUserVipsByDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Available) {
		query["Available"] = request.Available
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodUserVipsByDomain"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodUserVipsByDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ownership verification content.
//
// Description:
//
//	This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - DescribeVodVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodVerifyContentResponse
func (client *Client) DescribeVodVerifyContentWithContext(ctx context.Context, request *DescribeVodVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodVerifyContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodVerifyContent"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodVerifyContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes application permissions from the specified identity. The identity may a RAM user or RAM role.
//
// Description:
//
// >  You can grant a maximum of 10 application permissions to a RAM user or RAM role.
//
// @param request - DetachAppPolicyFromIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachAppPolicyFromIdentityResponse
func (client *Client) DetachAppPolicyFromIdentityWithContext(ctx context.Context, request *DetachAppPolicyFromIdentityRequest, runtime *dara.RuntimeOptions) (_result *DetachAppPolicyFromIdentityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.IdentityName) {
		query["IdentityName"] = request.IdentityName
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.PolicyNames) {
		query["PolicyNames"] = request.PolicyNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachAppPolicyFromIdentity"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachAppPolicyFromIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a key for secure download. ApsaraVideo Player SDK provides the secure download feature. Videos that are downloaded to your local device in this mode are encrypted. You can play the encrypted videos only by using the key file generated from the app that you specified. Secure download protects your videos from malicious playback or distribution.
//
// Description:
//
//	  To use the secure download feature, you must enable the download feature in the ApsaraVideo VOD console and set the download method to secure download. For more information, see [Configure download settings](https://help.aliyun.com/document_detail/86107.html).
//
//		- After you generate a key for secure download, you must configure the key in ApsaraVideo Player SDK. For more information, see [Secure download](https://help.aliyun.com/document_detail/124735.html).
//
// @param request - GenerateDownloadSecretKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDownloadSecretKeyResponse
func (client *Client) GenerateDownloadSecretKeyWithContext(ctx context.Context, request *GenerateDownloadSecretKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateDownloadSecretKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppDecryptKey) {
		query["AppDecryptKey"] = request.AppDecryptKey
	}

	if !dara.IsNil(request.AppIdentification) {
		query["AppIdentification"] = request.AppIdentification
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
		Action:      dara.String("GenerateDownloadSecretKey"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDownloadSecretKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a random Key Management Service (KMS) data key used for HLS encryption in ApsaraVideo VOD.
//
// @param request - GenerateKMSDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateKMSDataKeyResponse
func (client *Client) GenerateKMSDataKeyWithContext(ctx context.Context, request *GenerateKMSDataKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateKMSDataKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateKMSDataKey"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateKMSDataKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries jobs of image AI processing.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)*	- and **China (Shanghai)**.
//
//		- Call the [SubmitAIImageJob](~~SubmitAIImageJob~~) operation to submit image AI processing jobs before you call this operation to query image AI processing jobs.
//
//		- You can query a maximum of 10 jobs of image AI processing in one request.
//
// @param request - GetAIImageJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIImageJobsResponse
func (client *Client) GetAIImageJobsWithContext(ctx context.Context, request *GetAIImageJobsRequest, runtime *dara.RuntimeOptions) (_result *GetAIImageJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("GetAIImageJobs"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAIImageJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an intelligent review job. After the job is submitted, it is processed asynchronously. You can call this operation to query the job information in real time.
//
// Description:
//
// ApsaraVideo VOD stores the snapshots of the intelligent review results free of charge for two weeks. After this period, the snapshots are automatically deleted.
//
// @param request - GetAIMediaAuditJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIMediaAuditJobResponse
func (client *Client) GetAIMediaAuditJobWithContext(ctx context.Context, request *GetAIMediaAuditJobRequest, runtime *dara.RuntimeOptions) (_result *GetAIMediaAuditJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAIMediaAuditJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAIMediaAuditJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an AI template.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- Before you call this operation to query details of an AI template, you must obtain the ID of the AI template.
//
// @param request - GetAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITemplateResponse
func (client *Client) GetAITemplateWithContext(ctx context.Context, request *GetAITemplateRequest, runtime *dara.RuntimeOptions) (_result *GetAITemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAITemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAITemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of smart tagging jobs.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- You can obtain the smart tagging results by using the video ID.
//
// @param request - GetAIVideoTagResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIVideoTagResultResponse
func (client *Client) GetAIVideoTagResultWithContext(ctx context.Context, request *GetAIVideoTagResultRequest, runtime *dara.RuntimeOptions) (_result *GetAIVideoTagResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAIVideoTagResult"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAIVideoTagResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more applications based on application IDs.
//
// Description:
//
// You can specify multiple accelerated domain names in a request.
//
// @param request - GetAppInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInfosResponse
func (client *Client) GetAppInfosWithContext(ctx context.Context, request *GetAppInfosRequest, runtime *dara.RuntimeOptions) (_result *GetAppInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the URL and basic information about one or more auxiliary media assets such as watermark images, subtitle files, and materials based on IDs.
//
// Description:
//
// You can query information about up to 20 auxiliary media assets in a request.
//
// @param request - GetAttachedMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttachedMediaInfoResponse
func (client *Client) GetAttachedMediaInfoWithContext(ctx context.Context, request *GetAttachedMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAttachedMediaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAttachedMediaInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAttachedMediaInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the manual review history.
//
// @param request - GetAuditHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuditHistoryResponse
func (client *Client) GetAuditHistoryWithContext(ctx context.Context, request *GetAuditHistoryRequest, runtime *dara.RuntimeOptions) (_result *GetAuditHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuditHistory"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuditHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specific category and its subcategories based on the ID or type of the category.
//
// @param request - GetCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCategoriesResponse
func (client *Client) GetCategoriesWithContext(ctx context.Context, request *GetCategoriesRequest, runtime *dara.RuntimeOptions) (_result *GetCategoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCategories"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCategoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 支持区域化媒资ID级别播放数据查询
//
// @param request - GetDailyPlayRegionStatisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDailyPlayRegionStatisResponse
func (client *Client) GetDailyPlayRegionStatisWithContext(ctx context.Context, request *GetDailyPlayRegionStatisRequest, runtime *dara.RuntimeOptions) (_result *GetDailyPlayRegionStatisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.MediaRegion) {
		query["MediaRegion"] = request.MediaRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDailyPlayRegionStatis"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDailyPlayRegionStatisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the default AI template.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- You can query information only about the default AI template for automated review.
//
// @param request - GetDefaultAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultAITemplateResponse
func (client *Client) GetDefaultAITemplateWithContext(ctx context.Context, request *GetDefaultAITemplateRequest, runtime *dara.RuntimeOptions) (_result *GetDefaultAITemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDefaultAITemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDefaultAITemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a digital watermark extraction job. You can call this operation to obtain information such as the job status and the content of the copyright or user-tracing watermark.
//
// Description:
//
//	  This operation is supported only in the China (Shanghai) and China (Beijing) regions.
//
//		- You can call this operation to query the watermark content after you call the [SubmitDigitalWatermarkExtractJob](~~SubmitDigitalWatermarkExtractJob~~) operation to extract the copyright or user-tracing watermark in a video.
//
//		- You can query watermark content extracted only from watermark extraction jobs that are submitted in the last 2 years.
//
// @param request - GetDigitalWatermarkExtractResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDigitalWatermarkExtractResultResponse
func (client *Client) GetDigitalWatermarkExtractResultWithContext(ctx context.Context, request *GetDigitalWatermarkExtractResultRequest, runtime *dara.RuntimeOptions) (_result *GetDigitalWatermarkExtractResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtractType) {
		query["ExtractType"] = request.ExtractType
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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
		Action:      dara.String("GetDigitalWatermarkExtractResult"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDigitalWatermarkExtractResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取剪辑工程
//
// @param request - GetEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEditingProjectResponse
func (client *Client) GetEditingProjectWithContext(ctx context.Context, request *GetEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *GetEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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
		Action:      dara.String("GetEditingProject"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries materials to be edited for an online editing project.
//
// Description:
//
// During editing, you can add materials to the timeline, but some of them may not be used.
//
// @param request - GetEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEditingProjectMaterialsResponse
func (client *Client) GetEditingProjectMaterialsWithContext(ctx context.Context, request *GetEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *GetEditingProjectMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaterialType) {
		query["MaterialType"] = request.MaterialType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEditingProjectMaterials"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEditingProjectMaterialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information and access URL of an image based on the image ID.
//
// @param request - GetImageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageInfoResponse
func (client *Client) GetImageInfoWithContext(ctx context.Context, request *GetImageInfoRequest, runtime *dara.RuntimeOptions) (_result *GetImageInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about multiple images at a time.
//
// Description:
//
//	  You can call the [CreateUploadImage](~~CreateUploadImage~~) operation to upload images to ApsaraVideo VOD and call this operation to query the basic information about multiple images at a time.
//
//		- To query information about video snapshots, call the [ListSnapshots](~~ListSnapshots~~) operation.
//
//		- You can specify up to 20 image IDs in one call.
//
// @param request - GetImageInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageInfosResponse
func (client *Client) GetImageInfosWithContext(ctx context.Context, request *GetImageInfosRequest, runtime *dara.RuntimeOptions) (_result *GetImageInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.ImageIds) {
		query["ImageIds"] = request.ImageIds
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an asynchronous task based on jobId.
//
// Description:
//
// ***
//
// You can call this operation to query only asynchronous tasks of the last six months. The types of tasks that you can query include transcoding tasks, snapshot tasks, and AI tasks.
//
// **QPS limit**
//
// You can call this operation up to 15 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDetailResponse
func (client *Client) GetJobDetailWithContext(ctx context.Context, request *GetJobDetailRequest, runtime *dara.RuntimeOptions) (_result *GetJobDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobDetail"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of audio review results.
//
// Description:
//
// If notifications for the [CreateAuditComplete](https://help.aliyun.com/document_detail/89576.html) event are configured, event notifications are sent to the callback URL after automated review is complete. You can call this operation to query the details of audio review results.
//
// @param request - GetMediaAuditAudioResultDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditAudioResultDetailResponse
func (client *Client) GetMediaAuditAudioResultDetailWithContext(ctx context.Context, request *GetMediaAuditAudioResultDetailRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditAudioResultDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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
		Action:      dara.String("GetMediaAuditAudioResultDetail"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaAuditAudioResultDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the summary of automated review results.
//
// @param request - GetMediaAuditResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditResultResponse
func (client *Client) GetMediaAuditResultWithContext(ctx context.Context, request *GetMediaAuditResultRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaAuditResult"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaAuditResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of automated review results. You can call this operation to query the details of review results in real time.
//
// Description:
//
//	  By default, only details of snapshots that violate content regulations and potentially violate content regulations are returned.
//
//		- ApsaraVideo VOD stores the snapshots in the automated review results free of charge for two weeks. After this period, the snapshots are automatically deleted.
//
//		- This operation is available only in the Singapore region.
//
// @param request - GetMediaAuditResultDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditResultDetailResponse
func (client *Client) GetMediaAuditResultDetailWithContext(ctx context.Context, request *GetMediaAuditResultDetailRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditResultDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaAuditResultDetail"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaAuditResultDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the timelines of all snapshots that violate content regulations.
//
// Description:
//
// >  By default, only details of snapshots that violate content regulations and potentially violate content regulations are returned.
//
// This operation is available only in the Singapore region.
//
// @param request - GetMediaAuditResultTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditResultTimelineResponse
func (client *Client) GetMediaAuditResultTimelineWithContext(ctx context.Context, request *GetMediaAuditResultTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditResultTimelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaAuditResultTimeline"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaAuditResultTimelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a media fingerprinting result. After a media fingerprinting job is complete, you can call this operation to query the media fingerprinting result.
//
// Description:
//
// Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// @param request - GetMediaDNAResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaDNAResultResponse
func (client *Client) GetMediaDNAResultWithContext(ctx context.Context, request *GetMediaDNAResultRequest, runtime *dara.RuntimeOptions) (_result *GetMediaDNAResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaDNAResult"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaDNAResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about media refresh or prefetch jobs, such as the job status and filtering conditions.
//
// Description:
//
// You can query the information about all media files or a specific media file in a refresh or prefetch job.
//
// ### QPS limits
//
// You can call this operation up to 50 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits on API operations in ApsaraVideo VoD](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetMediaRefreshJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaRefreshJobsResponse
func (client *Client) GetMediaRefreshJobsWithContext(ctx context.Context, request *GetMediaRefreshJobsRequest, runtime *dara.RuntimeOptions) (_result *GetMediaRefreshJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaRefreshJobs"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaRefreshJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the callback method, callback URL, and event type for event notifications.
//
// Description:
//
// > For more information, see [Event notification](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - GetMessageCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageCallbackResponse
func (client *Client) GetMessageCallbackWithContext(ctx context.Context, request *GetMessageCallbackRequest, runtime *dara.RuntimeOptions) (_result *GetMessageCallbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageCallback"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the mezzanine file of an audio or video. The information includes the mezzanine file URL, resolution, and bitrate of the audio or video.
//
// Description:
//
// You can obtain complete information about the source file only after a stream is transcoded.
//
// @param request - GetMezzanineInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMezzanineInfoResponse
func (client *Client) GetMezzanineInfoWithContext(ctx context.Context, request *GetMezzanineInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMezzanineInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionType) {
		query["AdditionType"] = request.AdditionType
	}

	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMezzanineInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMezzanineInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the playback URL by the audio or video ID. Then, you can use ApsaraVideo Player or a third-party player, such as a system player, open source player, orself-developed player, to play the audio or video.
//
// Description:
//
//	  **Make sure that you understand the billing method and price of ApsaraVideo VOD before you call this operation. You are charged for outbound traffic when you download or play videos based on URLs in ApsaraVideo VOD. For more information about billing of outbound traffic, see [Billing of outbound traffic](~~188308#section-rwh-e88-f7j~~). If you have configured an accelerated domain name, see [Billing of the acceleration service](~~188308#section-c5t-oq9-15e~~). If you have activated the acceleration service, you are charged acceleration fees when you upload media files to ApsaraVideo VOD. For more information, see [Billing of acceleration traffic](~~188310#section_sta_zm2_tsv~~).**
//
//		- Only videos whose Status is Normal can be played. For more information, see [Overview](https://help.aliyun.com/document_detail/57290.html).
//
//		- If video playback fails, you can call the [GetMezzanineInfo](~~GetMezzanineInfo~~) operation to check whether the video source information is correct.
//
// @param request - GetPlayInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPlayInfoResponse
func (client *Client) GetPlayInfoWithContext(ctx context.Context, request *GetPlayInfoRequest, runtime *dara.RuntimeOptions) (_result *GetPlayInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionType) {
		query["AdditionType"] = request.AdditionType
	}

	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.Definition) {
		query["Definition"] = request.Definition
	}

	if !dara.IsNil(request.DigitalWatermarkType) {
		query["DigitalWatermarkType"] = request.DigitalWatermarkType
	}

	if !dara.IsNil(request.Formats) {
		query["Formats"] = request.Formats
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
	}

	if !dara.IsNil(request.PlayConfig) {
		query["PlayConfig"] = request.PlayConfig
	}

	if !dara.IsNil(request.ReAuthInfo) {
		query["ReAuthInfo"] = request.ReAuthInfo
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.Trace) {
		query["Trace"] = request.Trace
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPlayInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPlayInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries transcoding summaries of audio and video files based on the file ID. A transcoding summary includes the status and progress of transcoding.
//
// Description:
//
//	  An audio or video file may be transcoded multiple times. This operation returns only the latest transcoding summary.
//
//		- You can query transcoding summaries for a maximum of 10 audio and video files in one request.
//
//		- You can call the [ListTranscodeTask](https://help.aliyun.com/document_detail/109120.html) operation to query historical transcoding tasks.
//
//		- **You can call this operation to query information only about transcoding tasks created within the past year.
//
// @param request - GetTranscodeSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeSummaryResponse
func (client *Client) GetTranscodeSummaryWithContext(ctx context.Context, request *GetTranscodeSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VideoIds) {
		query["VideoIds"] = request.VideoIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranscodeSummary"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranscodeSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about transcoding jobs based on the transcoding task ID.
//
// Description:
//
// You can call this operation to query only transcoding tasks created within the past year.
//
// @param request - GetTranscodeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeTaskResponse
func (client *Client) GetTranscodeTaskWithContext(ctx context.Context, request *GetTranscodeTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.TranscodeTaskId) {
		query["TranscodeTaskId"] = request.TranscodeTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranscodeTask"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranscodeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a transcoding template group based on the template group ID.
//
// Description:
//
// This operation returns information about the specified transcoding template group and the configurations of all the transcoding templates in the group.
//
// @param request - GetTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeTemplateGroupResponse
func (client *Client) GetTranscodeTemplateGroupWithContext(ctx context.Context, request *GetTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeTemplateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TranscodeTemplateGroupId) {
		query["TranscodeTemplateGroupId"] = request.TranscodeTemplateGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranscodeTemplateGroup"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranscodeTemplateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about URL-based upload jobs.
//
// Description:
//
// You can query the information about a URL-based upload job by specifying the upload URL or using the job ID returned when you upload media files. The information includes the status of the upload job, custom configurations, the time when the job was created, and the time when the job was complete.
//
// If the upload fails, you can view the error code and error message. If the upload is successful, you can obtain the video ID.
//
// @param request - GetURLUploadInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetURLUploadInfosResponse
func (client *Client) GetURLUploadInfosWithContext(ctx context.Context, request *GetURLUploadInfosRequest, runtime *dara.RuntimeOptions) (_result *GetURLUploadInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.UploadURLs) {
		query["UploadURLs"] = request.UploadURLs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetURLUploadInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetURLUploadInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the upload details, such as the upload time, upload ratio, and upload source, about one or more media files based on the media IDs.
//
// Description:
//
//	  You can call this operation to obtain the upload details only about audio and video files.
//
//		- If you use the ApsaraVideo VOD console to upload audio and video files, you can call this operation to query information such as the upload ratio. If you use an upload SDK to upload audio and video files, make sure that the version of the [upload SDK](https://help.aliyun.com/document_detail/52200.html) meets one of the following requirements:
//
//	    	- The version of the upload SDK for Java is 1.4.4 or later.
//
//	    	- The version of the upload SDK for C++ is 1.0.0 or later.
//
//	    	- The version of the upload SDK for PHP is 1.0.2 or later.
//
//	    	- The version of the upload SDK for Python is 1.3.0 or later.
//
//	    	- The version of the upload SDK for JavaScript is 1.4.0 or later.
//
//	    	- The version of the upload SDK for Android is 1.5.0 or later.
//
//	    	- The version of the upload SDK for iOS is 1.5.0 or later.
//
// @param request - GetUploadDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadDetailsResponse
func (client *Client) GetUploadDetailsWithContext(ctx context.Context, request *GetUploadDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetUploadDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadDetails"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the title, description, duration, thumbnail URL, status, creation time, size, snapshots, category, and tags of a media file based on the file ID.
//
// Description:
//
// After a media file is uploaded, ApsaraVideo VOD processes the source file. Then, information about the media file is asynchronously generated. You can configure notifications for the [VideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event and call this operation to query information about a media file after you receive notifications for the [VideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event. For more information, see [Overview](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - GetVideoInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoInfoResponse
func (client *Client) GetVideoInfoWithContext(ctx context.Context, request *GetVideoInfoRequest, runtime *dara.RuntimeOptions) (_result *GetVideoInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information such as the title, description, duration, thumbnail URL, status, creation time, size, snapshots, category, and tags about multiple audio or video files based on IDs.
//
// Description:
//
//	  You can specify up to 20 audio or video file IDs in each request.
//
//		- After a media file is uploaded, ApsaraVideo VOD processes the source file. Then, information about the media file is asynchronously generated. You can configure notifications for the [VideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event and call this operation to query information about a media file after you receive notifications for the [VideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event. For more information, see [Overview](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - GetVideoInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoInfosResponse
func (client *Client) GetVideoInfosWithContext(ctx context.Context, request *GetVideoInfosRequest, runtime *dara.RuntimeOptions) (_result *GetVideoInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VideoIds) {
		query["VideoIds"] = request.VideoIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about media files.
//
// Description:
//
// You can call this operation to query information about media files based on the filter conditions that you specify, such as video status and category ID. Information about a maximum of **5,000*	- media files can be returned for each request. We recommend that you set the StartTime and EndTime parameters to specify a time range for each request. For more information about how to query information about more media files or even all media files, see [SearchMedia](https://help.aliyun.com/document_detail/86044.html).
//
// @param request - GetVideoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoListResponse
func (client *Client) GetVideoListWithContext(ctx context.Context, request *GetVideoListRequest, runtime *dara.RuntimeOptions) (_result *GetVideoListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoList"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the credential required for media playback. ApsaraVideo Player SDK automatically obtains the playback URL based on the playback credential. Each playback credential can be used to obtain the playback URL only for a specific video in a specific period of time. You cannot obtain the playback URL if the credential expires or is incorrect. You can use PlayAuth-based playback when you require high security for audio and video playback.
//
// Description:
//
//	  You can call this operation to obtain a playback credential when you use ApsaraVideo Player SDK to play a media file based on PlayAuth. The credential is used to obtain the playback URL. For more information, see [ApsaraVideo Player SDK](https://help.aliyun.com/document_detail/125579.html).
//
//		- You cannot obtain the playback URL of a video by using a credential that has expired. A new credential is required.
//
// @param request - GetVideoPlayAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoPlayAuthResponse
func (client *Client) GetVideoPlayAuthWithContext(ctx context.Context, request *GetVideoPlayAuthRequest, runtime *dara.RuntimeOptions) (_result *GetVideoPlayAuthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiVersion) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !dara.IsNil(request.AuthInfoTimeout) {
		query["AuthInfoTimeout"] = request.AuthInfoTimeout
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoPlayAuth"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoPlayAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a single snapshot template.
//
// @param request - GetVodTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVodTemplateResponse
func (client *Client) GetVodTemplateWithContext(ctx context.Context, request *GetVodTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetVodTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VodTemplateId) {
		query["VodTemplateId"] = request.VodTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVodTemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVodTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an image or text watermark based on the watermark template ID. You can call this operation to obtain information such as the position, size, and display time of an image watermark or the content, position, font, and font color of a text watermark.
//
// @param request - GetWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWatermarkResponse
func (client *Client) GetWatermarkWithContext(ctx context.Context, request *GetWatermarkRequest, runtime *dara.RuntimeOptions) (_result *GetWatermarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WatermarkId) {
		query["WatermarkId"] = request.WatermarkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWatermark"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AI processing results about the images of a specified video.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)*	- and **China (Shanghai)**.
//
//		- You can call this operation to query AI processing results about images of a specified video. Images of different videos cannot be queried in one request.
//
// @param request - ListAIImageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIImageInfoResponse
func (client *Client) ListAIImageInfoWithContext(ctx context.Context, request *ListAIImageInfoRequest, runtime *dara.RuntimeOptions) (_result *ListAIImageInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIImageInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIImageInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries AI jobs. After a job is submitted, ApsaraVideo VOD asynchronously processes the job. You can call this operation to query the job information in real time.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- You can call this operation to query video fingerprinting jobs and smart tagging jobs.
//
// @param request - ListAIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIJobResponse
func (client *Client) ListAIJobWithContext(ctx context.Context, request *ListAIJobRequest, runtime *dara.RuntimeOptions) (_result *ListAIJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("ListAIJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries AI templates.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- You can call this operation to query AI templates of a specified type.
//
// @param request - ListAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAITemplateResponse
func (client *Client) ListAITemplateWithContext(ctx context.Context, request *ListAITemplateRequest, runtime *dara.RuntimeOptions) (_result *ListAITemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAITemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAITemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that you are authorized to manage based on query conditions.
//
// Description:
//
// ### [](#)Usage notes
//
// You can query applications based on states.
//
// ### [](#qps-)QPS limit
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits on API operations](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - ListAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInfoResponse
func (client *Client) ListAppInfoWithContext(ctx context.Context, request *ListAppInfoRequest, runtime *dara.RuntimeOptions) (_result *ListAppInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the application policies that are attached to the specified identity. The identity may be a RAM user or RAM role.
//
// Description:
//
// > The IdentityType and IdentityName parameters take effect only when an identity assumes the application administrator role to call this operation. Otherwise, only application policies that are attached to the current identity are returned.
//
// @param request - ListAppPoliciesForIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppPoliciesForIdentityResponse
func (client *Client) ListAppPoliciesForIdentityWithContext(ctx context.Context, request *ListAppPoliciesForIdentityRequest, runtime *dara.RuntimeOptions) (_result *ListAppPoliciesForIdentityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.IdentityName) {
		query["IdentityName"] = request.IdentityName
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppPoliciesForIdentity"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppPoliciesForIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses in a review security group.
//
// @param request - ListAuditSecurityIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuditSecurityIpResponse
func (client *Client) ListAuditSecurityIpWithContext(ctx context.Context, request *ListAuditSecurityIpRequest, runtime *dara.RuntimeOptions) (_result *ListAuditSecurityIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityGroupName) {
		query["SecurityGroupName"] = request.SecurityGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuditSecurityIp"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuditSecurityIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about animated stickers of a video based on the video ID.
//
// @param request - ListDynamicImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicImageResponse
func (client *Client) ListDynamicImageWithContext(ctx context.Context, request *ListDynamicImageRequest, runtime *dara.RuntimeOptions) (_result *ListDynamicImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDynamicImage"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDynamicImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries historical tasks based on the media asset ID.
//
// Description:
//
// ***
//
//   - You can call the [GetJobDetail](https://apiworkbench.aliyun-inc.com/document/vod/2017-03-21/GetJobDetail?spm=openapi-amp.newDocPublishment.0.0.616a281fSegn0e) operation to query detailed information about the tasks.
//
//   - You can call this operation to query only asynchronous tasks of the last six months. The types of tasks that you can query include transcoding tasks, snapshot tasks, and AI tasks.
//
// **QPS limits**
//
// You can call this operation up to 15 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - ListJobInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobInfoResponse
func (client *Client) ListJobInfoWithContext(ctx context.Context, request *ListJobInfoRequest, runtime *dara.RuntimeOptions) (_result *ListJobInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries live-to-VOD videos.
//
// Description:
//
// You can query up to 5,000 videos based on the specified filter condition.
//
// @param request - ListLiveRecordVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRecordVideoResponse
func (client *Client) ListLiveRecordVideoWithContext(ctx context.Context, request *ListLiveRecordVideoRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRecordVideoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveRecordVideo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveRecordVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the snapshots that are captured by submitting snapshot jobs or snapshots that are generated by the system when you upload the video.
//
// Description:
//
// If multiple snapshots exist for a video, you can call this operation to query information about the latest snapshot.
//
// @param request - ListSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshotsWithContext(ctx context.Context, request *ListSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SnapshotType) {
		query["SnapshotType"] = request.SnapshotType
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSnapshots"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries transcoding tasks based on the media ID. This operation does not return specific job information.
//
// Description:
//
//	  You can call the [GetTranscodeTask](https://help.aliyun.com/document_detail/109121.html) operation to query details about transcoding jobs.
//
//		- **You can call this operation to query only transcoding tasks created within the past year.**
//
// @param request - ListTranscodeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranscodeTaskResponse
func (client *Client) ListTranscodeTaskWithContext(ctx context.Context, request *ListTranscodeTaskRequest, runtime *dara.RuntimeOptions) (_result *ListTranscodeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTranscodeTask"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTranscodeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries transcoding template groups.
//
// Description:
//
// > This operation does not return the configurations of transcoding templates in each transcoding template group. To query the configurations of transcoding templates in a specific transcoding template group, call the [GetTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102670.html) operation.
//
// @param request - ListTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranscodeTemplateGroupResponse
func (client *Client) ListTranscodeTemplateGroupWithContext(ctx context.Context, request *ListTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *ListTranscodeTemplateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTranscodeTemplateGroup"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTranscodeTemplateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries snapshot templates.
//
// @param request - ListVodTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVodTemplateResponse
func (client *Client) ListVodTemplateWithContext(ctx context.Context, request *ListVodTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListVodTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVodTemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVodTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information about all image and text watermark templates in a region. You can call this operation to obtain information such as the position, size, and display time of image watermarks or the content, position, font, and font color of text watermarks.
//
// @param request - ListWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWatermarkResponse
func (client *Client) ListWatermarkWithContext(ctx context.Context, request *ListWatermarkRequest, runtime *dara.RuntimeOptions) (_result *ListWatermarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWatermark"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates resources between applications. The application administrator can directly migrate resources between applications. Resource Access Management (RAM) users or RAM roles must obtain the write permissions on the source and destination applications before they migrate resources between applications. Multiple resources can be migrated at a time.
//
// @param request - MoveAppResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveAppResourceResponse
func (client *Client) MoveAppResourceWithContext(ctx context.Context, request *MoveAppResourceRequest, runtime *dara.RuntimeOptions) (_result *MoveAppResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TargetAppId) {
		query["TargetAppId"] = request.TargetAppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveAppResource"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveAppResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prefetches resources from an origin server to L2 nodes. Users can directly hit the cache upon their first visits. This way, workloads on the origin server can be reduced.
//
// Description:
//
// > 	- This operation is available only in the **China (Shanghai)*	- region.
//
// > 	- You can submit a maximum of 500 requests to prefetch resources based on URLs each day by using an Alibaba Cloud account. You cannot prefetch resources based on directories.
//
// > 	- You can call the [RefreshVodObjectCaches](https://help.aliyun.com/document_detail/69215.html) operation to refresh content and the [PreloadVodObjectCaches](https://help.aliyun.com/document_detail/69211.htmll) operation to prefetch content.
//
// @param request - PreloadVodObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadVodObjectCachesResponse
func (client *Client) PreloadVodObjectCachesWithContext(ctx context.Context, request *PreloadVodObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *PreloadVodObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.L2Preload) {
		query["L2Preload"] = request.L2Preload
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WithHeader) {
		query["WithHeader"] = request.WithHeader
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreloadVodObjectCaches"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreloadVodObjectCachesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Produces a video from one or more source files. You can directly specify source files by configuring the Timeline parameter. Alternatively, you can specify source files after you create an online editing project.
//
// Description:
//
//	  **Make sure that you understand the billing method and price of ApsaraVideo VOD before you call this operation. You are charged for using the online editing feature. For more information, see [Billing](~~188310#section-pyv-b8h-bo7~~).**
//
//		- This operation returns only the submission result of a video production task. When the submission result is returned, video production may still be in progress. After a video production task is submitted, the task is queued in the background for asynchronous processing.
//
//		- The source files that are used in the timeline of an online editing project can be materials directly uploaded to the online project or selected from the media asset library. Only media assets that are in the Normal state can be used in the project.
//
//		- Videos are produced based on ProjectId and Timeline. The following content describes the parameter configurations:
//
//	    	- You must specify ProjectId or Timeline. If you leave both parameters empty, the video cannot be produced.
//
//	    	- If you specify Timeline and leave ProjectId empty, the system automatically creates an online editing project based on Timeline and adds the materials specified in the Timeline to the project to produce videos.
//
//	    	- If you specify ProjectId and leave Timeline empty, the system automatically uses the latest timeline information of the project to produce videos.
//
//	    	- If you specify both ProjectId and Timeline, the system automatically uses the timeline information that you specified to produce videos and updates the project timeline and materials. You can also specify other parameters to update the corresponding information about the online editing project.
//
//		- You can create up to 100 video tracks, 100 image tracks, and 100 subtitle tracks in a project.
//
//		- The total size of material files cannot exceed 1 TB.
//
//		- The buckets in which the materials reside and where the exported videos are stored must be in the same region as the region where ApsaraVideo VOD is activated.
//
//		- The exported videos must meet the following requirements:
//
//	    	- The width and height of the video image cannot be less than 128 pixels.
//
//	    	- The width and height of the video image cannot exceed 4,096 pixels.
//
//	    	- The width cannot exceed 2,160 pixels.
//
//		- After a video is produced, the video is automatically uploaded to ApsaraVideo VOD. Then, the **ProduceMediaComplete*	- and **FileUploadComplete*	- event notifications are sent to you. After the produced video is transcoded, the **StreamTranscodeComplete*	- and **TranscodeComplete*	- event notifications are sent to you.
//
//		- You can add special effects to the video. For more information, see [Special effects](https://help.aliyun.com/document_detail/69082.html).
//
// @param request - ProduceEditingProjectVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProduceEditingProjectVideoResponse
func (client *Client) ProduceEditingProjectVideoWithContext(ctx context.Context, request *ProduceEditingProjectVideoRequest, runtime *dara.RuntimeOptions) (_result *ProduceEditingProjectVideoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MediaMetadata) {
		query["MediaMetadata"] = request.MediaMetadata
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProduceConfig) {
		query["ProduceConfig"] = request.ProduceConfig
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Timeline) {
		query["Timeline"] = request.Timeline
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
		Action:      dara.String("ProduceEditingProjectVideo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProduceEditingProjectVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits media refresh or prefetch tasks based on the media IDs.
//
// Description:
//
//	  ApsaraVideo VOD allows you to purge and prefetch resources. The purge feature forces the point of presence (POP) to clear cached resources and retrieve the latest resources from origin servers. The prefetch feature allows the POP to retrieve frequently accessed resources from origin servers during off-peak hours. This increases the cache hit ratio.
//
//		- You can call this operation to submit purge or prefetch tasks based on the media ID. You can also specify the format and resolution of the media streams to purge or prefetch based on your business requirements.
//
//		- You can submit a maximum of 20 purge or prefetch tasks at a time.
//
// @param request - RefreshMediaPlayUrlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshMediaPlayUrlsResponse
func (client *Client) RefreshMediaPlayUrlsWithContext(ctx context.Context, request *RefreshMediaPlayUrlsRequest, runtime *dara.RuntimeOptions) (_result *RefreshMediaPlayUrlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Definitions) {
		query["Definitions"] = request.Definitions
	}

	if !dara.IsNil(request.Formats) {
		query["Formats"] = request.Formats
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SliceCount) {
		query["SliceCount"] = request.SliceCount
	}

	if !dara.IsNil(request.SliceFlag) {
		query["SliceFlag"] = request.SliceFlag
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshMediaPlayUrls"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshMediaPlayUrlsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a new upload credential after a file failed to be uploaded.
//
// Description:
//
// You can also call this operation to overwrite the source file of an audio or video file. After you call this operation, the system obtains the upload URL and uploads a new source file without changing the ID of the audio or video file. If you have configured transcoding or snapshot capture for the upload, the transcoding or snapshot capture job is automatically triggered. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// @param request - RefreshUploadVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshUploadVideoResponse
func (client *Client) RefreshUploadVideoWithContext(ctx context.Context, request *RefreshUploadVideoRequest, runtime *dara.RuntimeOptions) (_result *RefreshUploadVideoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshUploadVideo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshUploadVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes files on Alibaba Cloud CDN nodes. You can refresh multiple files at a time based on URLs.
//
// Description:
//
//	  This operation is available only in the **China (Shanghai)*	- region.
//
//		- You can submit a maximum of 2,000 requests to refresh resources based on URLs and 100 requests to refresh resources based on directories each day by using an Alibaba Cloud account.
//
//		- You can call the [RefreshVodObjectCaches](https://help.aliyun.com/document_detail/69215.html) operation to refresh content and the [PreloadVodObjectCaches](https://help.aliyun.com/document_detail/69211.html) operation to prefetch content.
//
// @param request - RefreshVodObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshVodObjectCachesResponse
func (client *Client) RefreshVodObjectCachesWithContext(ctx context.Context, request *RefreshVodObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *RefreshVodObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshVodObjectCaches"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshVodObjectCachesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers media files. After you add an Object Storage Service (OSS) bucket to ApsaraVideo VOD, you must register the media files in the bucket to generate the required information before you use features such as transcoding and snapshot capture on the media files.
//
// Description:
//
//	  After you add an OSS bucket to ApsaraVideo VOD, you must register media files in the OSS bucket to generate the required information. Then, you can use media IDs for features such as transcoding, snapshot capture, and AI processing.use features such as xxx on media files by specifying their IDs?
//
//		- You can register up to 10 media files in an OSS bucket in a request. The media files must be stored in the same bucket.
//
//		- If you do not specify a transcoding template group ID when you upload a media file to ApsaraVideo VOD, the media file is automatically transcoded based on the default template group. If you do not specify a transcoding template group ID after you register a media file, the media file is not automatically transcoded. The registered media files are automatically transcoded only if you specify a transcoding template group ID.
//
//		- If the media file that you want to register has been registered, this operation returns only the unique media ID that is associated with the media file. No further operation is performed.
//
//		- Make sure that the media file that you want to register has a valid suffix. Otherwise, the registration fails.
//
// @param request - RegisterMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterMediaResponse
func (client *Client) RegisterMediaWithContext(ctx context.Context, request *RegisterMediaRequest, runtime *dara.RuntimeOptions) (_result *RegisterMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegisterMetadatas) {
		query["RegisterMetadatas"] = request.RegisterMetadatas
	}

	if !dara.IsNil(request.TemplateGroupId) {
		query["TemplateGroupId"] = request.TemplateGroupId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterMedia"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores media assets.
//
// Description:
//
// You can call this operation to restore only Archive and Cold Archive audio and video files. You can access the audio and video files after the files are restored. You cannot change the storage class of an audio or video file that is being restored. You are charged for the retrieval traffic generated during restoration. After a Cold Archive audio or video file is restored, a Standard replica of the file is generated for access. You are charged for the storage of the replica before the file returns to the frozen state.
//
// @param request - RestoreMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreMediaResponse
func (client *Client) RestoreMediaWithContext(ctx context.Context, request *RestoreMediaRequest, runtime *dara.RuntimeOptions) (_result *RestoreMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.RestoreDays) {
		query["RestoreDays"] = request.RestoreDays
	}

	if !dara.IsNil(request.RestoreTier) {
		query["RestoreTier"] = request.RestoreTier
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreMedia"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries online editing projects.
//
// @param request - SearchEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchEditingProjectResponse
func (client *Client) SearchEditingProjectWithContext(ctx context.Context, request *SearchEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *SearchEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchEditingProject"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about videos, audio, images, and auxiliary media assets. You can call this operation and specify the search protocol to query media assets based on the return fields, fields used for exact match, fields used for fuzzy match, fields used for a multi-value query, fields used for a range query, and sort fields.
//
// Description:
//
// The maximum number of data records that you can query varies based on the method used to query the data. You can use the following methods to query data:
//
//   - Method 1: Traverse data by page
//
//     You can use the PageNo and PageSize parameters to traverse up to 5,000 data records that meet the specified filter condition. PageNo specifies the page number and PageSize specifies the number of data records displayed on a page. If the number of data records that meet the specified filter condition exceeds 5,000, change the filter conditions to narrow down the results. You cannot use this method to traverse all data records. If you want to traverse more data records, use Method 2.
//
//   - Method 2: Traverse all data (available only for audio and video files)
//
//     You can use this method to traverse up to 2 million data records related to audio and video files. If the number of data records that meet the specified filter condition exceeds 2 million, change the filter conditions to narrow down the results. To traverse data page by page, you must set the PageNo, PageSize, and ScrollToken parameters. The total number of data records from the current page to the target page cannot exceed 100. For example, you set PageSize to 20. The following content describes the traverse logic:
//
//   - When the PageNo parameter is set to 1, you can traverse data records from page 1 to page 5.
//
//   - When the PageNo parameter is set to 2, you can traverse data records from page 2 to page 6.
//
// Make sure that you set the appropriate page number and page size, and use a traverse method based on the number of results that meet your filter condition.
//
// @param request - SearchMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaResponse
func (client *Client) SearchMediaWithContext(ctx context.Context, request *SearchMediaRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Match) {
		query["Match"] = request.Match
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScrollToken) {
		query["ScrollToken"] = request.ScrollToken
	}

	if !dara.IsNil(request.SearchType) {
		query["SearchType"] = request.SearchType
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMedia"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages the IP addresses in review security groups.
//
// Description:
//
// You can play videos in the Checking or Blocked state only from the IP addresses that are added to review security groups.
//
// @param request - SetAuditSecurityIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAuditSecurityIpResponse
func (client *Client) SetAuditSecurityIpWithContext(ctx context.Context, request *SetAuditSecurityIpRequest, runtime *dara.RuntimeOptions) (_result *SetAuditSecurityIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ips) {
		query["Ips"] = request.Ips
	}

	if !dara.IsNil(request.OperateMode) {
		query["OperateMode"] = request.OperateMode
	}

	if !dara.IsNil(request.SecurityGroupName) {
		query["SecurityGroupName"] = request.SecurityGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAuditSecurityIp"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAuditSecurityIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the cross-domain policy file crossdomain.xml.
//
// Description:
//
// > After you use the cross-domain policy file to update the resources on the origin server, you must refresh the resources that are cached on Alibaba Cloud CDN nodes. You can use the ApsaraVideo VOD console to refresh resources. For more information, see [Refresh and prefetch](https://help.aliyun.com/document_detail/86098.html). Alternatively, you can call the [RefreshVodObjectCaches](https://help.aliyun.com/document_detail/69215.html) operation to refresh resources.
//
// @param request - SetCrossdomainContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCrossdomainContentResponse
func (client *Client) SetCrossdomainContentWithContext(ctx context.Context, request *SetCrossdomainContentRequest, runtime *dara.RuntimeOptions) (_result *SetCrossdomainContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
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

	if !dara.IsNil(request.ResourceRealOwnerId) {
		query["ResourceRealOwnerId"] = request.ResourceRealOwnerId
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCrossdomainContent"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCrossdomainContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies an AI template as the default template.
//
// Description:
//
// Specifies an AI template as the default template.
//
// @param request - SetDefaultAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultAITemplateResponse
func (client *Client) SetDefaultAITemplateWithContext(ctx context.Context, request *SetDefaultAITemplateRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultAITemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultAITemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultAITemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a transcoding template group as the default one.
//
// @param request - SetDefaultTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultTranscodeTemplateGroupResponse
func (client *Client) SetDefaultTranscodeTemplateGroupWithContext(ctx context.Context, request *SetDefaultTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultTranscodeTemplateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TranscodeTemplateGroupId) {
		query["TranscodeTemplateGroupId"] = request.TranscodeTemplateGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultTranscodeTemplateGroup"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultTranscodeTemplateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a watermark template as the default one.
//
// @param request - SetDefaultWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultWatermarkResponse
func (client *Client) SetDefaultWatermarkWithContext(ctx context.Context, request *SetDefaultWatermarkRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultWatermarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WatermarkId) {
		query["WatermarkId"] = request.WatermarkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultWatermark"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the media assets that you want to edit in an online editing project.
//
// @param request - SetEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetEditingProjectMaterialsResponse
func (client *Client) SetEditingProjectMaterialsWithContext(ctx context.Context, request *SetEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *SetEditingProjectMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaterialIds) {
		query["MaterialIds"] = request.MaterialIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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
		Action:      dara.String("SetEditingProjectMaterials"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetEditingProjectMaterialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the callback method, callback URL, and event type of an event notification.
//
// Description:
//
// HTTP callbacks and MNS callbacks are supported. For more information, see [Overview](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - SetMessageCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMessageCallbackResponse
func (client *Client) SetMessageCallbackWithContext(ctx context.Context, request *SetMessageCallbackRequest, runtime *dara.RuntimeOptions) (_result *SetMessageCallbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.AuthSwitch) {
		query["AuthSwitch"] = request.AuthSwitch
	}

	if !dara.IsNil(request.CallbackType) {
		query["CallbackType"] = request.CallbackType
	}

	if !dara.IsNil(request.CallbackURL) {
		query["CallbackURL"] = request.CallbackURL
	}

	if !dara.IsNil(request.EventTypeList) {
		query["EventTypeList"] = request.EventTypeList
	}

	if !dara.IsNil(request.MnsEndpoint) {
		query["MnsEndpoint"] = request.MnsEndpoint
	}

	if !dara.IsNil(request.MnsQueueName) {
		query["MnsQueueName"] = request.MnsQueueName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetMessageCallback"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetMessageCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the certificate of a domain name and modifies the certificate information.
//
// Description:
//
// > This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - SetVodDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVodDomainCertificateResponse
func (client *Client) SetVodDomainCertificateWithContext(ctx context.Context, request *SetVodDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetVodDomainCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLPri) {
		query["SSLPri"] = request.SSLPri
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVodDomainCertificate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVodDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the SSL certificate of a domain name and updates the certificate information.
//
// @param request - SetVodDomainSSLCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVodDomainSSLCertificateResponse
func (client *Client) SetVodDomainSSLCertificateWithContext(ctx context.Context, request *SetVodDomainSSLCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetVodDomainSSLCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertRegion) {
		query["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLPri) {
		query["SSLPri"] = request.SSLPri
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVodDomainSSLCertificate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVodDomainSSLCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an automated review job for an image. After the job is submitted, the job is processed in an asynchronous manner. The operation may return a response before the job is complete.
//
// Description:
//
// This operation is available only in the Singapore region.
//
// @param request - SubmitAIImageAuditJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIImageAuditJobResponse
func (client *Client) SubmitAIImageAuditJobWithContext(ctx context.Context, request *SubmitAIImageAuditJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIImageAuditJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaAuditConfiguration) {
		query["MediaAuditConfiguration"] = request.MediaAuditConfiguration
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAIImageAuditJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAIImageAuditJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits jobs of image AI processing.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)*	- and **China (Shanghai)**.
//
//		- After you call this operation, you can call the [GetAIImageJobs](https://help.aliyun.com/document_detail/186923.html) operation to query the job execution result.
//
// @param request - SubmitAIImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIImageJobResponse
func (client *Client) SubmitAIImageJobWithContext(ctx context.Context, request *SubmitAIImageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIImageJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIPipelineId) {
		query["AIPipelineId"] = request.AIPipelineId
	}

	if !dara.IsNil(request.AITemplateId) {
		query["AITemplateId"] = request.AITemplateId
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

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAIImageJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAIImageJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a smart tagging or video fingerprinting job.
//
// Description:
//
//	  **Make sure that you understand the billing method and price of ApsaraVideo VOD before you call this operation. You are charged for using the smart tagging and video fingerprinting features. For more information, see [Billing of video AI](~~188310#section-g7l-s3o-9ng~~).**
//
//		- Regions that support the video fingerprinting feature: **China (Beijing)**, **China (Shanghai)**, and **Singapore**. Regions that support the smart tagging feature: **China (Beijing)*	- and **China (Shanghai)**.
//
//		- You need to enable the video fingerprinting feature or the smart tagging feature before you can call this operation to submit jobs. For more information, see [Overview](https://help.aliyun.com/document_detail/101148.html).
//
//		- If this is the first time you use the video fingerprinting feature, you must submit a ticket to apply for using the media fingerprint library for free. Otherwise, the video fingerprinting feature will be affected. For more information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
//
//		- After you submit an AI job, ApsaraVideo VOD asynchronously processes the job. The operation may return a response before the job is complete. You can configure the [Event Notification](https://help.aliyun.com/document_detail/55627.html) feature and set the callback event to **AI Processing Completed**. After you receive the event notification, you can query the execution result of the AI job.
//
// @param request - SubmitAIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIJobResponse
func (client *Client) SubmitAIJobWithContext(ctx context.Context, request *SubmitAIJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
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

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAIJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an automated review job for a media file. After the job is submitted, ApsaraVideo VOD asynchronously processes the job. Therefore, the operation may return a response before the job is complete.
//
// Description:
//
//	  **Make sure that you understand the billing methods and price of ApsaraVideo VOD before you call this operation. You are charged for using the automated review feature. For more information about billing, submit a ticket or contact your account manager.**
//
//		- You can call this operation only in the **China (Shanghai)**, **China (Beijing)**, and **Singapore*	- regions.
//
//		- For more information, see [Automated review](https://help.aliyun.com/document_detail/101148.html).
//
//		- After an automated review job is complete, the images generated during the review are stored in the VOD bucket for two weeks free of charge. The images are automatically deleted after two weeks.
//
// @param request - SubmitAIMediaAuditJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIMediaAuditJobResponse
func (client *Client) SubmitAIMediaAuditJobWithContext(ctx context.Context, request *SubmitAIMediaAuditJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIMediaAuditJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaAuditConfiguration) {
		query["MediaAuditConfiguration"] = request.MediaAuditConfiguration
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAIMediaAuditJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAIMediaAuditJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a digital watermark extraction job. You can call this operation to asynchronously extract a copyright watermark or user-tracing watermark.
//
// Description:
//
//	  **Make sure that you understand the billing methods and price of ApsaraVideo VOD before you call this operation. You are charged for generating and extracting digital watermarks. For more information, see [Billing](~~188310#62b9c940403se~~).**
//
//		- This operation is supported only in the **China (Shanghai)*	- and **China (Beijing)*	- regions.
//
//		- Before you submit a digital watermark extraction job, make sure that the following conditions are met:
//
//	    	- The video from which you want to extract the watermark is uploaded to the ApsaraVideo VOD.
//
//	    	- The video from which you want to extract the watermark is longer than 6 minutes.
//
// @param request - SubmitDigitalWatermarkExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDigitalWatermarkExtractJobResponse
func (client *Client) SubmitDigitalWatermarkExtractJobWithContext(ctx context.Context, request *SubmitDigitalWatermarkExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDigitalWatermarkExtractJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtractType) {
		query["ExtractType"] = request.ExtractType
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
		Action:      dara.String("SubmitDigitalWatermarkExtractJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDigitalWatermarkExtractJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a frame animation job and starts asynchronous processing.
//
// Description:
//
//	  You can capture a part of a video and generate animated images only when the video is in the **Uploaded**, **Transcoding**, **Normal**, **Reviewing**, or **Flagged*	- state.
//
//		- The fees for frame animation are included in your video transcoding bill. You are charged based on the output resolution and the duration. For more information, see [Billing of basic services](https://help.aliyun.com/document_detail/188308.html).
//
// ### QPS limits
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit on API operations](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - SubmitDynamicImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDynamicImageJobResponse
func (client *Client) SubmitDynamicImageJobWithContext(ctx context.Context, request *SubmitDynamicImageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDynamicImageJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicImageTemplateId) {
		query["DynamicImageTemplateId"] = request.DynamicImageTemplateId
	}

	if !dara.IsNil(request.OverrideParams) {
		query["OverrideParams"] = request.OverrideParams
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDynamicImageJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDynamicImageJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a video fingerprinting job.
//
// Description:
//
// Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// @param request - SubmitMediaDNADeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaDNADeleteJobResponse
func (client *Client) SubmitMediaDNADeleteJobWithContext(ctx context.Context, request *SubmitMediaDNADeleteJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaDNADeleteJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaDNADeleteJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaDNADeleteJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transcodes a video by using the production studio.
//
// Description:
//
//	  During video preprocessing, videos are transcoded to meet the playback requirements of the production studio. Therefore, **you are charged for video preprocessing**. For more information about billing, see [Billing of production studios](https://help.aliyun.com/document_detail/64531.html).
//
//		- You can obtain the preprocessing result in the [TranscodeComplete](https://help.aliyun.com/document_detail/55638.html) event notification. If **Preprocess=true*	- is returned in the event notification, the video is transcoded.
//
// @param request - SubmitPreprocessJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPreprocessJobsResponse
func (client *Client) SubmitPreprocessJobsWithContext(ctx context.Context, request *SubmitPreprocessJobsRequest, runtime *dara.RuntimeOptions) (_result *SubmitPreprocessJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PreprocessType) {
		query["PreprocessType"] = request.PreprocessType
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitPreprocessJobs"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitPreprocessJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a snapshot job for a video and starts asynchronous snapshot processing.
//
// Description:
//
//	  Only snapshots in the JPG format are generated.
//
//		- After a snapshot is captured, the [SnapshotComplete](https://help.aliyun.com/document_detail/57337.html) callback is fired and EventType=SnapshotComplete, SubType=SpecifiedTime is returned.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 30 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342790.html).
//
// @param tmpReq - SubmitSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSnapshotJobResponse
func (client *Client) SubmitSnapshotJobWithContext(ctx context.Context, tmpReq *SubmitSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSnapshotJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitSnapshotJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SpecifiedOffsetTimes) {
		request.SpecifiedOffsetTimesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpecifiedOffsetTimes, dara.String("SpecifiedOffsetTimes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.Height) {
		query["Height"] = request.Height
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.SnapshotTemplateId) {
		query["SnapshotTemplateId"] = request.SnapshotTemplateId
	}

	if !dara.IsNil(request.SpecifiedOffsetTime) {
		query["SpecifiedOffsetTime"] = request.SpecifiedOffsetTime
	}

	if !dara.IsNil(request.SpecifiedOffsetTimesShrink) {
		query["SpecifiedOffsetTimes"] = request.SpecifiedOffsetTimesShrink
	}

	if !dara.IsNil(request.SpriteSnapshotConfig) {
		query["SpriteSnapshotConfig"] = request.SpriteSnapshotConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	if !dara.IsNil(request.Width) {
		query["Width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSnapshotJob"),
		Version:     dara.String("2017-03-21"),
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
// Submits a transcoding job to start transcoding in an asynchronous manner.
//
// Description:
//
// ### [](#)Usage notes
//
//   - **Make sure that you understand the billing methods and prices of ApsaraVideo VOD before you call this operation. For more information about billing of the transcoding feature, see [Billing of basic services](~~188308#section-ejb-nii-nqa~~).**
//
//   - You can transcode a video only in the Uploaded, Normal, or Reviewing state.
//
//   - You can obtain the transcoding results from the [StreamTranscodeComplete](https://help.aliyun.com/document_detail/55636.html) or [TranscodeComplete](https://help.aliyun.com/document_detail/55638.html) callback.
//
//   - You can call this operation to dynamically override the subtitle URL in an HTTP Live Streaming (HLS) packaging task. If the packaging task does not contain subtitles, we recommend that you specify the ID of the specific packaging template group when you upload the video instead of calling this operation.
//
// @param request - SubmitTranscodeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTranscodeJobsResponse
func (client *Client) SubmitTranscodeJobsWithContext(ctx context.Context, request *SubmitTranscodeJobsRequest, runtime *dara.RuntimeOptions) (_result *SubmitTranscodeJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptConfig) {
		query["EncryptConfig"] = request.EncryptConfig
	}

	if !dara.IsNil(request.OverrideParams) {
		query["OverrideParams"] = request.OverrideParams
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TemplateGroupId) {
		query["TemplateGroupId"] = request.TemplateGroupId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTranscodeJobs"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTranscodeJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a workflow to process media files.
//
// Description:
//
// **Make sure that you understand the billing method and price of ApsaraVideo VOD before you call this operation. When you use workflows to process videos, you may be charged for transcoding, encryption, and automated review. For more information, see [Billing overview](https://help.aliyun.com/document_detail/188307.html).**
//
//   - You can call this operation to initiate a VOD workflow to process media files. For more information, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
//
// @param request - SubmitWorkflowJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitWorkflowJobResponse
func (client *Client) SubmitWorkflowJobWithContext(ctx context.Context, request *SubmitWorkflowJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitWorkflowJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitWorkflowJob"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitWorkflowJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an AI template.
//
// Description:
//
//	  Regions that support this operation: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
//		- After you call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template, you can call this operation to modify the AI template.
//
// @param request - UpdateAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAITemplateResponse
func (client *Client) UpdateAITemplateWithContext(ctx context.Context, request *UpdateAITemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateAITemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("UpdateAITemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAITemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about an application.
//
// Description:
//
// ## QPS limit
//
// A single user can perform a maximum of 30 queries per second (QPS). Throttling is triggered when the number of calls per second exceeds the QPS limit. The throttling may affect your business. Thus, we recommend that you observe the QPS limit on this operation.
//
// @param request - UpdateAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppInfoResponse
func (client *Client) UpdateAppInfoWithContext(ctx context.Context, request *UpdateAppInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about multiple auxiliary media assets such as watermark images, subtitle files, and materials in a batch based on IDs. You can modify information such as the title, description, tags, and category.
//
// Description:
//
// You can modify the information about up to 20 auxiliary media assets at a time.
//
// @param request - UpdateAttachedMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAttachedMediaInfosResponse
func (client *Client) UpdateAttachedMediaInfosWithContext(ctx context.Context, request *UpdateAttachedMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *UpdateAttachedMediaInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UpdateContent) {
		query["UpdateContent"] = request.UpdateContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAttachedMediaInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAttachedMediaInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a video category.
//
// Description:
//
// After you create a category, you can call this operation to modify the name of the category. If you have classified specific media resources to this category, the category names that are labeled on the media resources are automatically updated.
//
// @param request - UpdateCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCategoryResponse
func (client *Client) UpdateCategoryWithContext(ctx context.Context, request *UpdateCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.CateName) {
		query["CateName"] = request.CateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCategory"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an online editing project.
//
// @param request - UpdateEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEditingProjectResponse
func (client *Client) UpdateEditingProjectWithContext(ctx context.Context, request *UpdateEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Timeline) {
		query["Timeline"] = request.Timeline
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEditingProject"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about one or more images at a time.
//
// Description:
//
//	  You can call this operation to modify information such as the title, tags, description, and category about images based on image IDs. You must pass in the parameters that you want to modify. Otherwise, parameter configurations are not overwritten.
//
//		- You can modify the information about up to 20 images at a time.
//
// @param request - UpdateImageInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageInfosResponse
func (client *Client) UpdateImageInfosWithContext(ctx context.Context, request *UpdateImageInfosRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UpdateContent) {
		query["UpdateContent"] = request.UpdateContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImageInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the storage classes of media assets.
//
// Description:
//
//	  This operation is an asynchronous operation. You can call this operation to modify the storage classes of media assets. After the storage class is modified, a callback notification is sent.
//
//		- If the storage class of the media asset is Archive or Cold Archive and you call this operation to modify the storage class of the media asset, the media asset is automatically restored before the storage class is modified. You do not need to call the RestoreMedia operation to restore the media asset. You must specify the restoration priority for Cold Archive objects. Default configuration: RestoreTier=Standard.
//
//		- Media assets whose storage classes are being modified cannot be used or processed.
//
//		- Non-Standard objects have minimum storage durations. If an object is stored for less than the minimum storage duration, the storage class of the object cannot be changed. The following content describes the minimum storage durations for objects in different storage classes: IA or IA storage for source files: 30 days, Archive or Archive storage for source files: 60 days, Cold Archive or Cold Archive for source files: 180 days.
//
// @param request - UpdateMediaStorageClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaStorageClassResponse
func (client *Client) UpdateMediaStorageClassWithContext(ctx context.Context, request *UpdateMediaStorageClassRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaStorageClassResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowUpdateWithoutTimeLimit) {
		query["AllowUpdateWithoutTimeLimit"] = request.AllowUpdateWithoutTimeLimit
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.RestoreTier) {
		query["RestoreTier"] = request.RestoreTier
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.StorageClass) {
		query["StorageClass"] = request.StorageClass
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaStorageClass"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaStorageClassResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a transcoding template group or configurations of transcoding templates in the transcoding template group.
//
// Description:
//
// For security purposes, you cannot add, modify, or delete transcoding templates in a transcoding template group that is locked. You can call the [GetTranscodeTemplateGroup](~~GetTranscodeTemplateGroup~~) operation to query the configurations of a transcoding template group, check whether the transcoding template group is locked by using the response parameter Locked, and unlock the transcoding template group before you perform operations such as add, modify, and delete transcoding templates.
//
// @param request - UpdateTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTranscodeTemplateGroupResponse
func (client *Client) UpdateTranscodeTemplateGroupWithContext(ctx context.Context, request *UpdateTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateTranscodeTemplateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Locked) {
		query["Locked"] = request.Locked
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TranscodeTemplateGroupId) {
		query["TranscodeTemplateGroupId"] = request.TranscodeTemplateGroupId
	}

	if !dara.IsNil(request.TranscodeTemplateList) {
		query["TranscodeTemplateList"] = request.TranscodeTemplateList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTranscodeTemplateGroup"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTranscodeTemplateGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about an audio or video file.
//
// Description:
//
// ### [](#)
//
// You can call this operation to modify information such as the title, tags, and description about audio and video files based on audio or video IDs. You must pass in the parameters that you want to modify. Otherwise, parameter configurations are not overwritten.
//
// ### [](#qps-)Queries per second (QPS) limit
//
// You can call this operation up to 100 times per second per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits on API operations](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - UpdateVideoInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoInfoResponse
func (client *Client) UpdateVideoInfoWithContext(ctx context.Context, request *UpdateVideoInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateVideoInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoInfo"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about multiple videos at a time.
//
// Description:
//
// The specific parameter of a video is updated only when a new value is passed in the parameter.
//
// @param request - UpdateVideoInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoInfosResponse
func (client *Client) UpdateVideoInfosWithContext(ctx context.Context, request *UpdateVideoInfosRequest, runtime *dara.RuntimeOptions) (_result *UpdateVideoInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UpdateContent) {
		query["UpdateContent"] = request.UpdateContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoInfos"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a specific accelerated domain name.
//
// Description:
//
// # UpdateVodDomain
//
// @param request - UpdateVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVodDomainResponse
func (client *Client) UpdateVodDomainWithContext(ctx context.Context, request *UpdateVodDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateVodDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVodDomain"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVodDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a snapshot template.
//
// @param request - UpdateVodTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVodTemplateResponse
func (client *Client) UpdateVodTemplateWithContext(ctx context.Context, request *UpdateVodTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateVodTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.VodTemplateId) {
		query["VodTemplateId"] = request.VodTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVodTemplate"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVodTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and configurations of a watermark template after you create a watermark template.
//
// Description:
//
//	  You can modify the name and configurations of the watermark template after you call the [AddWatermark](~~AddWatermark~~) operation to create a watermark template.
//
//		- You cannot call this operation to change the image in an image watermark template.
//
// @param request - UpdateWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWatermarkResponse
func (client *Client) UpdateWatermarkWithContext(ctx context.Context, request *UpdateWatermarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateWatermarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.WatermarkConfig) {
		query["WatermarkConfig"] = request.WatermarkConfig
	}

	if !dara.IsNil(request.WatermarkId) {
		query["WatermarkId"] = request.WatermarkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWatermark"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWatermarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads media files based on URLs.
//
// Description:
//
//	  You can call this operation to upload media files that are not stored on a local server or device and must be uploaded based on URLs over the Internet.
//
//		- The URL-based upload jobs are asynchronous. After you submit a URL-based upload job by calling this operation, it may take hours, even days to complete. If you require high timeliness, we recommend that you use the upload SDK.
//
//		- If you configure callbacks, you can receive an [UploadByURLComplete](https://help.aliyun.com/document_detail/86326.html) event notification after the media file is uploaded. You can query the upload status by calling the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation.
//
//		- After you submit an upload job, the job is asynchronously processed on the cloud. All URL-based upload jobs that are submitted in each region are queued. The waiting time for the upload job depends on the number of queued jobs. After the upload job is complete, you can associate the playback URL included in the callback with the media ID.
//
//		- You can call this operation only in the **China (Shanghai)*	- and **Singapore*	- regions.
//
//		- Every time you submit a URL-based upload job, a new media ID is generated in ApsaraVideo VOD.
//
// @param request - UploadMediaByURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMediaByURLResponse
func (client *Client) UploadMediaByURLWithContext(ctx context.Context, request *UploadMediaByURLRequest, runtime *dara.RuntimeOptions) (_result *UploadMediaByURLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StorageLocation) {
		query["StorageLocation"] = request.StorageLocation
	}

	if !dara.IsNil(request.TemplateGroupId) {
		query["TemplateGroupId"] = request.TemplateGroupId
	}

	if !dara.IsNil(request.UploadMetadatas) {
		query["UploadMetadatas"] = request.UploadMetadatas
	}

	if !dara.IsNil(request.UploadURLs) {
		query["UploadURLs"] = request.UploadURLs
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadMediaByURL"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadMediaByURLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads transcoded streams to ApsaraVideo VOD from external storage.
//
// Description:
//
//	  **Make sure that you understand the billing method and price of ApsaraVideo VOD before you call this operation. You are charged storage fees after you upload media files to ApsaraVideo VOD. For more information, see [Billing of media asset storage](~~188308#section_e97_xrp_mzz~~). If you have activated the acceleration service, you are charged acceleration fees when you upload media files to ApsaraVideo VOD. For more information, see [Billing of acceleration traffic](~~188310#section_sta_zm2_tsv~~).**
//
//		- This operation is available only in the **China (Shanghai)*	- and **Singapore*	- regions.
//
//		- You can call this operation to upload transcoded streams to ApsaraVideo VOD from external storage. The following HDR types of transcoded streams are supported: HDR, HDR 10, HLG, Dolby Vision, HDR Vivid, and SDR+.
//
//		- You can call the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation to query the upload status. After the upload is complete, the callback of the [UploadByURLComplete](https://help.aliyun.com/document_detail/376427.html) event is returned.
//
// @param request - UploadStreamByURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadStreamByURLResponse
func (client *Client) UploadStreamByURLWithContext(ctx context.Context, request *UploadStreamByURLRequest, runtime *dara.RuntimeOptions) (_result *UploadStreamByURLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		query["Definition"] = request.Definition
	}

	if !dara.IsNil(request.FileExtension) {
		query["FileExtension"] = request.FileExtension
	}

	if !dara.IsNil(request.HDRType) {
		query["HDRType"] = request.HDRType
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.StreamURL) {
		query["StreamURL"] = request.StreamURL
	}

	if !dara.IsNil(request.UploadMetadata) {
		query["UploadMetadata"] = request.UploadMetadata
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadStreamByURL"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadStreamByURLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a specified domain name.
//
// Description:
//
// This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - VerifyVodDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyVodDomainOwnerResponse
func (client *Client) VerifyVodDomainOwnerWithContext(ctx context.Context, request *VerifyVodDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyVodDomainOwnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyVodDomainOwner"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyVodDomainOwnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
