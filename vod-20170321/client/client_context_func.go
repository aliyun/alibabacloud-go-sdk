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
// - Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - <props="china">Before adding an AI template for automated review and smart thumbnail tasks, make sure that you have activated [automated review](https://ai.aliyun.com/vi/censor) or [smart thumbnail](https://ai.aliyun.com/vi/cover).
//
// - <props="intl">Before adding an AI template for automated review and smart thumbnail tasks, make sure that you have activated automated review or smart thumbnail.
//
// @param request - AddAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAITemplateResponse
func (client *Client) AddAITemplateWithContext(ctx context.Context, request *AddAITemplateRequest, runtime *dara.RuntimeOptions) (_result *AddAITemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a category to classify audio, video, image, and short video materials stored in ApsaraVideo VOD, making resource discovery and management more efficient.
//
// Description:
//
// - Audio/video/image categories (`Type` set to `default`) support up to three levels, with a maximum of 100 subcategories per level. Short video material categories (`Type` set to `material`) support up to two levels, with a maximum of 100 subcategories per level.
//
// - After creating a category, you can assign it to media assets during upload or to already uploaded media assets. For more information, see [Media asset categories](https://help.aliyun.com/document_detail/86070.html).
//
// @param request - AddCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCategoryResponse
func (client *Client) AddCategoryWithContext(ctx context.Context, request *AddCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates an online editing project (video editing task).
//
// Description:
//
// - For more information about online editing, see [Online editing](https://help.aliyun.com/document_detail/95482.html).
//
// @param request - AddEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEditingProjectResponse
func (client *Client) AddEditingProjectWithContext(ctx context.Context, request *AddEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *AddEditingProjectResponse, _err error) {
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
// Adds one or more materials to an online editing project.
//
// @param request - AddEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEditingProjectMaterialsResponse
func (client *Client) AddEditingProjectMaterialsWithContext(ctx context.Context, request *AddEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *AddEditingProjectMaterialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Adds transcoding configurations. You can create a transcoding template group or add transcoding templates to a specified template group.
//
// Description:
//
// - Transcoding template groups that are **locked*	- by the ApsaraVideo VOD backend do not support custom operations. You can call the [GetTranscodeTemplateGroup](~~GetTranscodeTemplateGroup~~) operation to query the template configuration and check whether the template group is locked based on the Locked response parameter. You can call the [UpdateTranscodeTemplateGroup](~~UpdateTranscodeTemplateGroup~~) operation to unlock the template before you modify it.
//
// - Because transcoding involves storage addresses of files, you cannot add a transcoding template group if no storage address is available. You can activate a **VOD system bucket*	- in the **ApsaraVideo VOD console > Configuration Management > Media Asset Management Configuration > Storage Management*	- to obtain an available storage address.
//
// - You cannot add transcoding template configurations to a **No Transcoding*	- template group.
//
// - You must specify either **TranscodeTemplateGroupId*	- or **Name**.
//
// - You can create a maximum of 20 transcoding template groups.
//
// - You can add a maximum of 20 transcoding template configurations to a transcoding template group.
//
// - To generate adaptive bitrate streaming addresses through transcoding, you can add a maximum of 10 video packaging templates to a transcoding template group. If more than 10 templates are added, only individual stream addresses are generated instead of adaptive bitrate streaming addresses.
//
// ### QPS limit
//
// The maximum number of queries per second (QPS) per user for this operation is 5. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Manage your calls appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - AddTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTranscodeTemplateGroupResponse
func (client *Client) AddTranscodeTemplateGroupWithContext(ctx context.Context, request *AddTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *AddTranscodeTemplateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Adds a domain name for CDN acceleration in ApsaraVideo VOD.
//
// Description:
//
// - Currently, the only supported service address is **China (Shanghai)**.
//
// - Before creating an accelerated domain name, you must activate [ApsaraVideo VOD](https://help.aliyun.com/document_detail/51512.html), and the accelerated domain name must have a completed ICP filing.
//
// - Origin content that is not hosted on Alibaba Cloud requires review, which will be completed before the next business day.
//
// - You can submit only one accelerated domain name at a time. Each user can add up to 20 domain names.
//
// @param request - AddVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVodDomainResponse
func (client *Client) AddVodDomainWithContext(ctx context.Context, request *AddVodDomainRequest, runtime *dara.RuntimeOptions) (_result *AddVodDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Binds a storage bucket to a sub-application of ApsaraVideo VOD.
//
// Description:
//
// Calls AddVodStorageForApp to enable a VOD system bucket for an ApsaraVideo VOD sub-application.
//
//	<notice>Each sub-application can have at most one VOD system bucket enabled. If you specify an AppId that does not exist or an AppId that already has a VOD system bucket enabled, an error is returned.</notice>
//
// <notice>To call this operation, the caller must have application administrator permissions (VODAppAdministratorAccess). The Alibaba Cloud account has application administrator permissions by default. An application administrator can call AttachAppPolicyToIdentity to grant application permissions to a RAM user or role.</notice>
//
// @param request - AddVodStorageForAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVodStorageForAppResponse
func (client *Client) AddVodStorageForAppWithContext(ctx context.Context, request *AddVodStorageForAppRequest, runtime *dara.RuntimeOptions) (_result *AddVodStorageForAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Adds a snapshot or animated image template.
//
// Description:
//
// - After adding a template, you can pass the snapshot or animated image template ID through the [SubmitSnapshotJob](~~SubmitSnapshotJob~~) or [SubmitDynamicImageJob](~~SubmitDynamicImageJob~~) operation to initiate a snapshot or animated image job.
//
// - You can receive [video snapshot completed](https://help.aliyun.com/document_detail/57337.html) and [video animated image completed](https://help.aliyun.com/document_detail/143490.html) callback messages through HTTP callbacks (compatible with HTTPS) or MNS callbacks. For more information, see [Callback methods](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - AddVodTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVodTemplateResponse
func (client *Client) AddVodTemplateWithContext(ctx context.Context, request *AddVodTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddVodTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// To simplify watermark task processing, ApsaraVideo VOD consolidates complex watermark parameters such as position, size, font, and color into templates, each identified by a unique watermark template ID. Calls this operation to add an image or text watermark template.
//
// Description:
//
// - Call this operation to add an image watermark template (`Image`) or a text watermark template (`Text`). Image watermark templates support the following formats: static images (PNG) and animated images (GIF, APNG, MOV).
//
// - After adding a watermark template by calling this operation, call [AddTranscodeTemplateGroup](~~AddTranscodeTemplateGroup~~) or [UpdateTranscodeTemplateGroup](~~UpdateTranscodeTemplateGroup~~) to associate the watermark template with a transcoding template group for subsequent watermark transcoding.
//
// - For more information about adding image and text watermarks to videos, see [Video watermarks](https://help.aliyun.com/document_detail/99369.html).
//
// @param request - AddWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWatermarkResponse
func (client *Client) AddWatermarkWithContext(ctx context.Context, request *AddWatermarkRequest, runtime *dara.RuntimeOptions) (_result *AddWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to attach access permissions of an ApsaraVideo VOD application to a specified identity (Resource Access Management (RAM) user or RAM role).
//
// Description:
//
//	Notice:  Each Resource Access Management (RAM) user or RAM role can be granted permissions on up to 10 applications.
//
//	Notice: You must have application administrator permissions to invoke this operation. For the first invocation, use your Alibaba Cloud account.
//
// - If the policy name is VODAppAdministratorAccess, AppId is optional. For other policies, AppId is required.
//
// @param request - AttachAppPolicyToIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachAppPolicyToIdentityResponse
func (client *Client) AttachAppPolicyToIdentityWithContext(ctx context.Context, request *AttachAppPolicyToIdentityRequest, runtime *dara.RuntimeOptions) (_result *AttachAppPolicyToIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the basic information and source file information of multiple media assets in a batch.
//
// Description:
//
// - You can retrieve information about up to 20 audio or video files at a time.
//
// - After an audio or video file is uploaded, ApsaraVideo VOD analyzes the uploaded source file. Therefore, media asset information is generated asynchronously. You can configure an [event notification](https://help.aliyun.com/document_detail/55627.html) for the [AudioVideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event. After you receive the [AudioVideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event notification, call this operation to retrieve the audio or video information.
//
// @param request - BatchGetMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetMediaInfosResponse
func (client *Client) BatchGetMediaInfosWithContext(ctx context.Context, request *BatchGetMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *BatchGetMediaInfosResponse, _err error) {
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

	if !dara.IsNil(request.ReferenceIds) {
		query["ReferenceIds"] = request.ReferenceIds
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
// Configures multiple accelerated domain names in a batch.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - You can configure up to 50 domain names at a time.
//
// - After you call this operation to configure certain features for domain names, a unique ConfigId is generated. You can use the ConfigId to update or delete domain name configurations. This operation does not return the ConfigId. To obtain the ConfigId, call the [DescribeVodDomainConfigs](~~DescribeVodDomainConfigs~~) operation.
//
// @param request - BatchSetVodDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetVodDomainConfigsResponse
func (client *Client) BatchSetVodDomainConfigsWithContext(ctx context.Context, request *BatchSetVodDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetVodDomainConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Enables an accelerated domain name that is in the Disabled state.
//
// Description:
//
// - Currently, the only supported endpoint is **China (Shanghai)**.
//
// - If the account associated with the domain name has an overdue payment or the domain name is in an illegal state, you cannot call this operation to enable the ApsaraVideo VOD domain name.
//
// @param request - BatchStartVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartVodDomainResponse
func (client *Client) BatchStartVodDomainWithContext(ctx context.Context, request *BatchStartVodDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStartVodDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Pauses domain name acceleration.
//
// Description:
//
// - Currently, the only supported service address is **China (Shanghai)**.
//
// - After you pause the accelerated domain name, the domain name information is retained. Requests to the accelerated domain name are automatically redirected to the origin server.
//
// @param request - BatchStopVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopVodDomainResponse
func (client *Client) BatchStopVodDomainWithContext(ctx context.Context, request *BatchStopVodDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStopVodDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Cancels URL upload tasks that are in the queue.
//
// Description:
//
// - You can cancel only URL upload nodes whose status is **Pending**. You can call the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation to query the node status.
//
// - Upload nodes that have already started to execute cannot be canceled.
//
// - The request parameters JobIds and UploadUrls must have one specified. If both are specified, only JobIds is processed.
//
// @param request - CancelUrlUploadJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelUrlUploadJobsResponse
func (client *Client) CancelUrlUploadJobsWithContext(ctx context.Context, request *CancelUrlUploadJobsRequest, runtime *dara.RuntimeOptions) (_result *CancelUrlUploadJobsResponse, _err error) {
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
// Transfers a resource to a different resource group.
//
// Description:
//
// Transfers a resource to another resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Each account can create up to 10 applications. For more information, see [Multi-application development guide](https://help.aliyun.com/document_detail/113600.html).
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - CreateAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppInfoResponse
func (client *Client) CreateAppInfoWithContext(ctx context.Context, request *CreateAppInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateAppInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Submits a manual review request to review media information such as videos and audio files.
//
// @param request - CreateAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuditResponse
func (client *Client) CreateAuditWithContext(ctx context.Context, request *CreateAuditRequest, runtime *dara.RuntimeOptions) (_result *CreateAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// The upload URL and credential are provided by ApsaraVideo VOD to address authorization and security concerns, prevent malicious uploads, and enable automatic creation of a media asset ID (MediaId) for management and processing. For auxiliary media assets such as watermarks and subtitles, invoke this operation to obtain the upload credential and create the corresponding media asset information.
//
// Description:
//
// - **Before using this operation, make sure that you understand the billing methods and pricing of ApsaraVideo VOD. Uploading media files to ApsaraVideo VOD incurs storage fees. For more information, see [Media asset storage billing](~~188308#section_e97_xrp_mzz~~). If you have enabled storage transfer acceleration, uploading media files to ApsaraVideo VOD also incurs upload acceleration fees. For more information, see [Storage transfer acceleration billing](~~188310#section_sta_zm2_tsv~~).**
//
// - This operation only obtains the upload URL and credential and creates basic media asset information. It does not upload files. For a complete example of uploading files by using the API, see [Upload media files by using the ApsaraVideo VOD API](https://help.aliyun.com/document_detail/476208.html).
//
// - If the upload credential expires (valid for 3000 seconds), call this operation again to obtain a new upload URL and credential.
//
// - You can configure callbacks to receive event notifications for [auxiliary media asset upload complete](https://help.aliyun.com/document_detail/103250.html) to determine whether the upload is successful.
//
// - Obtaining the upload URL and credential is a core fundamental of ApsaraVideo VOD and a required step for every upload operation. ApsaraVideo VOD provides various upload methods, each with different requirements for obtaining the upload URL and credential. For more information, see the instructions in [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// @param request - CreateUploadAttachedMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadAttachedMediaResponse
func (client *Client) CreateUploadAttachedMediaWithContext(ctx context.Context, request *CreateUploadAttachedMediaRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadAttachedMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the upload URL and upload credential for uploading an image to ApsaraVideo VOD, and creates image information. ApsaraVideo VOD issues upload URLs and credentials to ensure authorization and security, prevent malicious uploads, and supports automatic creation of an image ID (ImageId) for management. You can invoke this operation to obtain the upload URL and credential and create image information.
//
// Description:
//
// - **Before using this operation, make sure that you understand the billing methods and pricing of ApsaraVideo VOD. Uploading media files to ApsaraVideo VOD incurs storage fees. For more information, see [Media asset storage billing](~~188308#section_e97_xrp_mzz~~). If you have enabled storage and transfer acceleration, uploading media files to ApsaraVideo VOD also incurs upload acceleration fees. For more information, see [Storage and transfer acceleration billing](~~188310#section_sta_zm2_tsv~~).**
//
// - This operation only retrieves the upload URL and credential and creates basic media asset information. It does not upload files. For a complete example of uploading files by calling API operations, see [Upload media files by using the ApsaraVideo VOD API](https://help.aliyun.com/document_detail/476208.html).
//
// - Refreshing the upload URL and credential is not supported for image uploads. If the image upload credential expires (the default validity period is 3000 seconds), call this operation again to obtain a new upload URL and credential.
//
// - You can configure callbacks to receive event notifications for [image upload completion](https://help.aliyun.com/document_detail/91968.html) to determine whether the upload is successful.
//
// - Retrieving the upload URL and credential is a core operation of ApsaraVideo VOD and is required for every upload. ApsaraVideo VOD provides multiple upload methods, each with different requirements for retrieving the upload URL and credential. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// @param request - CreateUploadImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadImageResponse
func (client *Client) CreateUploadImageWithContext(ctx context.Context, request *CreateUploadImageRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// ApsaraVideo VOD issues the upload URL and upload credential to ensure authorization and security and prevent malicious uploads. During issuance, a media ID (MediaId), also called a video ID (VideoId), undergoes automatic creation for management. Invoke this operation to obtain the upload URL and upload credential, and create audio or video information.
//
// Description:
//
// - **Before you use this operation, make sure that you are familiar with the billing methods and pricing of ApsaraVideo VOD. Uploading media files to ApsaraVideo VOD incurs storage fees. For more information, see [Media asset storage billing](~~188308#section_e97_xrp_mzz~~). If you have enabled storage and transfer acceleration, uploading media files to ApsaraVideo VOD also incurs upload acceleration fees. For more information, see [Storage and transfer acceleration billing](~~188310#section_sta_zm2_tsv~~). Storage fees are calculated from the time when the file is uploaded. Acceleration fees are calculated when you perform upload operations after the feature is enabled. Simply calling this operation does not incur fees.**
//
// - Obtaining the upload URL and credential is the core foundation of ApsaraVideo VOD and is a required step for every upload operation. ApsaraVideo VOD provides multiple upload methods, each with different requirements for obtaining the upload URL and credential. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// - This operation is used only to obtain the upload URL and credential and create basic media asset information. It does not upload files. For a complete example of uploading files by using API operations, see [Upload media files by using the ApsaraVideo VOD API](https://help.aliyun.com/document_detail/476208.html).
//
// - This operation supports obtaining the upload URL and credential for both video and audio files. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// - If the upload credential expires (the default validity period is 3000 seconds), call the [RefreshUploadVideo](~~RefreshUploadVideo~~) operation to obtain a new upload credential.
//
// - After the upload is complete, you can configure callbacks to receive [upload event notifications](https://help.aliyun.com/document_detail/55396.html) or call the [GetMezzanineInfo](https://help.aliyun.com/document_detail/59624.html) operation to check the file status and determine whether the upload is successful.
//
// - The VideoId parameter returned by this operation can be used for media asset lifecycle management or media processing.
//
// @param request - CreateUploadVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadVideoResponse
func (client *Client) CreateUploadVideoWithContext(ctx context.Context, request *CreateUploadVideoRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.EnableFirstFrameCover) {
		query["EnableFirstFrameCover"] = request.EnableFirstFrameCover
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileSize) {
		query["FileSize"] = request.FileSize
	}

	if !dara.IsNil(request.GenerateThumbnail) {
		query["GenerateThumbnail"] = request.GenerateThumbnail
	}

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Invokes this operation to perform decryption on the CiphertextBlob in a KMS data key (DK).
//
// @param request - DecryptKMSDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptKMSDataKeyResponse
func (client *Client) DecryptKMSDataKeyWithContext(ctx context.Context, request *DecryptKMSDataKeyRequest, runtime *dara.RuntimeOptions) (_result *DecryptKMSDataKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes AI image information.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)*	- and **China (Shanghai)**.
//
// - **This operation only deletes AI image information and does not actually delete image files**.
//
// - A maximum of 10 IDs can be deleted at a time.
//
// @param request - DeleteAIImageInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAIImageInfosResponse
func (client *Client) DeleteAIImageInfosWithContext(ctx context.Context, request *DeleteAIImageInfosRequest, runtime *dara.RuntimeOptions) (_result *DeleteAIImageInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// - Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - An AI template that is set as the default template cannot be deleted.
//
// @param request - DeleteAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAITemplateResponse
func (client *Client) DeleteAITemplateWithContext(ctx context.Context, request *DeleteAITemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAITemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes application information.
//
// Description:
//
// An application cannot be deleted if it contains resources.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this API appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - DeleteAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInfoResponse
func (client *Client) DeleteAppInfoWithContext(ctx context.Context, request *DeleteAppInfoRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes one or more auxiliary media assets that have been uploaded to ApsaraVideo VOD, such as watermark images, subtitle files, and materials.
//
// Description:
//
// - **This operation physically deletes auxiliary media assets. Once deleted, they cannot be recovered. Proceed with caution.**
//
// - You can delete up to 20 auxiliary media assets at a time.
//
// @param request - DeleteAttachedMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAttachedMediaResponse
func (client *Client) DeleteAttachedMediaWithContext(ctx context.Context, request *DeleteAttachedMediaRequest, runtime *dara.RuntimeOptions) (_result *DeleteAttachedMediaResponse, _err error) {
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
// - **This operation deletes a category and all its subcategories (including second-level and third-level categories). Proceed with caution.**
//
// - If a category has been assigned to media assets, deleting the category also removes the category assignment from those media assets.
//
// @param request - DeleteCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithContext(ctx context.Context, request *DeleteCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
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
// Deletes animated sticker information.
//
// Description:
//
// - This operation only deletes the association between animated stickers and videos. It does not delete the actual animated sticker files.
//
// - After the association is deleted, the deleted animated sticker information can no longer be queried by calling the [ListDynamicImage](https://help.aliyun.com/document_detail/180958.html) operation.
//
// - If you do not specify **DynamicImageIds**, all animated stickers associated with the specified VideoId are deleted. However, if the video has more than 10 animated stickers, the deletion request is rejected.
//
// ### QPS limit
//
// The maximum queries per second (QPS) per user for this operation is 10. If the number of calls exceeds the limit, throttling is triggered. This may affect your business. Call this operation as needed. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - DeleteDynamicImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDynamicImageResponse
func (client *Client) DeleteDynamicImageWithContext(ctx context.Context, request *DeleteDynamicImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteDynamicImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes an online editing project.
//
// Description:
//
// - Supports batch deletion.
//
// @param request - DeleteEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEditingProjectResponse
func (client *Client) DeleteEditingProjectWithContext(ctx context.Context, request *DeleteEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteEditingProjectResponse, _err error) {
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
// Deletes materials from an online editing project.
//
// @param request - DeleteEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEditingProjectMaterialsResponse
func (client *Client) DeleteEditingProjectMaterialsWithContext(ctx context.Context, request *DeleteEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEditingProjectMaterialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes images uploaded by users or images generated from video snapshots.
//
// Description:
//
// - **When you call this operation to delete images, the source files are permanently deleted. This action is irreversible. Once deleted, the images cannot be recovered. Proceed with caution.**
//
// - When **DeleteImageType*	- is set to **VideoId**, **VideoId*	- and **ImageType*	- are available and required.
//
// - When **DeleteImageType*	- is set to **ImageURL**, **ImageIds*	- and **ImageURLs*	- are available and required.
//
// - After you call this operation to delete images, CDN caches may still exist in some cases, which means the image URLs may not become invalid immediately.
//
// @param request - DeleteImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageResponse
func (client *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes the callback method, callback URL, and event types of an event notification.
//
// Description:
//
// > For more information, see [Event notification development guide](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - DeleteMessageCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageCallbackResponse
func (client *Client) DeleteMessageCallbackWithContext(ctx context.Context, request *DeleteMessageCallbackRequest, runtime *dara.RuntimeOptions) (_result *DeleteMessageCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes the source files of multiple audio and video files at a time.
//
// Description:
//
// Media processing operations in ApsaraVideo VOD (transcoding, snapshots, automated review, etc.) are performed on source files. Once a source file is deleted, subsequent media processing operations cannot be performed. Proceed with caution.
//
// @param request - DeleteMezzaninesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMezzaninesResponse
func (client *Client) DeleteMezzaninesWithContext(ctx context.Context, request *DeleteMezzaninesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMezzaninesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.ReferenceIds) {
		query["ReferenceIds"] = request.ReferenceIds
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
// Deletes fragment files generated during upload.
//
// Description:
//
// - When you use multipart upload, fragment files may be generated if the upload fails. These fragment files are automatically cleared after 7 days. After the upload is complete or fails, you can call this operation to manually clear the fragment files.
//
// - Calling this operation does not delete the original file or transcoded files. It only deletes fragment files generated during the upload process.
//
// - Calling the [DeleteVideo](https://help.aliyun.com/document_detail/52837.html) operation deletes the complete video file, including fragment files.
//
// @param request - DeleteMultipartUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultipartUploadResponse
func (client *Client) DeleteMultipartUploadWithContext(ctx context.Context, request *DeleteMultipartUploadRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultipartUploadResponse, _err error) {
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
// Deletes media stream (video stream or audio stream) information and storage files.
//
// Description:
//
// ### Usage notes
//
// Batch deletion is supported.
//
// ### QPS limit
//
// A single user can perform a maximum of 50 queries per second (QPS). Throttling is triggered when the QPS limit is exceeded, which may affect your business. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - DeleteStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStreamResponse
func (client *Client) DeleteStreamWithContext(ctx context.Context, request *DeleteStreamRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamResponse, _err error) {
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

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Deletes transcoding configurations. You can delete specific transcoding templates from a transcoding template group or force delete an entire transcoding template group.
//
// Description:
//
// - Default transcoding templates cannot be deleted. Remove the default designation before deleting them.
//
//   - For security protection purposes, a locked transcoding template group cannot be added to, modified, or deleted. Call the [GetTranscodeTemplateGroup](~~GetTranscodeTemplateGroup~~) operation to query the template configuration and check the Locked response parameter to determine whether the template group is locked. Call the [UpdateTranscodeTemplateGroup](~~UpdateTranscodeTemplateGroup~~) operation to unlock the template before making changes.
//
//   - If the ForceDelGroup parameter is empty or set to false, the TranscodeTemplateIds parameter is required.
//
// @param request - DeleteTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTranscodeTemplateGroupResponse
func (client *Client) DeleteTranscodeTemplateGroupWithContext(ctx context.Context, request *DeleteTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteTranscodeTemplateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes complete videos, including video source files, transcoded stream files, and thumbnails.
//
// Description:
//
// ### Usage notes
//
// - **This operation physically deletes videos. Deleted videos cannot be recovered. Proceed with caution.**
//
// - Batch deletion is supported.
//
// - When you delete a video, the source files are deleted, including the video source file, transcoded stream files, and thumbnails. However, the CDN cache is not refreshed through synchronization. If your business requires it, use the purge feature in the ApsaraVideo VOD console to clear stale data from the point of presence. Related operations: see [Purge and prefetch](https://help.aliyun.com/document_detail/86098.html).
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, the API invoke is throttled, which may affect your business. Invoke this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - DeleteVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVideoResponse
func (client *Client) DeleteVideoWithContext(ctx context.Context, request *DeleteVideoRequest, runtime *dara.RuntimeOptions) (_result *DeleteVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReferenceIds) {
		query["ReferenceIds"] = request.ReferenceIds
	}

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
// Deletes an accelerated domain name that has been added for ApsaraVideo VOD.
//
// Description:
//
// >- Currently, the only supported service region is **China (Shanghai)**.
//
// >- This operation makes the domain name inaccessible. Proceed with caution. Before deleting the domain name, restore the A record of the domain name at your DNS service provider.
//
// >- After the domain name is successfully deleted, all related records of the ApsaraVideo VOD domain name are deleted. If you only want to temporarily disable the domain name, use the [DisableVodDomainOffline](https://help.aliyun.com/document_detail/120208.html) operation.
//
// @param request - DeleteVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodDomainResponse
func (client *Client) DeleteVodDomainWithContext(ctx context.Context, request *DeleteVodDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes the configurations of a domain name for CDN acceleration in ApsaraVideo VOD.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - This operation causes the domain name to become inaccessible. Before deleting the domain name, restore the A record of the domain name at your DNS service provider.
//
// - After the domain name is successfully deleted, all related records of the ApsaraVideo VOD domain name are removed. If you only want to temporarily disable the domain name, use the [DisableVodRealtimeLogDelivery](https://help.aliyun.com/document_detail/120208.html) operation.
//
// @param request - DeleteVodSpecificConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodSpecificConfigResponse
func (client *Client) DeleteVodSpecificConfigWithContext(ctx context.Context, request *DeleteVodSpecificConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodSpecificConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes an image watermark template or text watermark template by watermark template ID.
//
// Description:
//
// - **When you delete an image watermark template, the watermark source file is physically deleted and cannot be recovered. Proceed with caution.**
//
// - A watermark template that has been set as the default watermark template cannot be deleted. To delete it, call [SetDefaultWatermark](~~SetDefaultWatermark~~) to set another watermark template as the default template to remove the default status, and then delete it.
//
// @param request - DeleteWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWatermarkResponse
func (client *Client) DeleteWatermarkWithContext(ctx context.Context, request *DeleteWatermarkRequest, runtime *dara.RuntimeOptions) (_result *DeleteWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the time-based distribution of audio and video media assets. The maximum time span between the start time and end time is six months.
//
// Description:
//
// - Currently, this operation is supported only in the China (Shanghai) region.
//
// - If you do not specify StartTime and EndTime, this operation returns data for the past 7 days by default. If you specify StartTime and EndTime, this operation returns data for the specified time range.
//
// @param request - DescribeMediaDistributionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMediaDistributionResponse
func (client *Client) DescribeMediaDistributionWithContext(ctx context.Context, request *DescribeMediaDistributionRequest, runtime *dara.RuntimeOptions) (_result *DescribeMediaDistributionResponse, _err error) {
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
// Retrieves daily playback statistics for top videos, including the number of views, unique viewers, and total playback duration.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - You can query playback statistics for up to the top 1000 videos per day. The top video list is sorted in descending order by the number of views by default.
//
// - Only playback data collected through ApsaraVideo Player SDK is supported.
//
// - Based on UTC+8, playback statistics for the previous day are generated at 9:00 AM each day.
//
// - You can query data generated after January 1, 2018. The maximum time range for a query is 180 days.
//
// @param request - DescribePlayTopVideosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayTopVideosResponse
func (client *Client) DescribePlayTopVideosWithContext(ctx context.Context, request *DescribePlayTopVideosRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayTopVideosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves daily average playback statistics for a specified time range, including average playback duration and average playback count.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - Only playback data collected through ApsaraVideo Player SDK is supported.
//
// - Playback statistics for the previous day are generated at 9:00 AM (UTC+8) each day.
//
// - You can query data generated after 2018-01-01. The maximum time range between the start time and end time is 180 days.
//
// @param request - DescribePlayUserAvgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayUserAvgResponse
func (client *Client) DescribePlayUserAvgWithContext(ctx context.Context, request *DescribePlayUserAvgRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayUserAvgResponse, _err error) {
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
// Retrieves the total daily playback statistics within a specified time range, including total play count, total unique viewers, total playback duration, and playback duration distribution.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - Only playback data from ApsaraVideo Player SDK is supported.
//
// - Based on UTC+8, playback statistics for the previous day are generated at 9:00 AM each day.
//
// - Data after 2018-01-01 can be queried. The maximum time span between the start time and end time is 180 days.
//
// @param request - DescribePlayUserTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayUserTotalResponse
func (client *Client) DescribePlayUserTotalWithContext(ctx context.Context, request *DescribePlayUserTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayUserTotalResponse, _err error) {
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
// Queries the daily playback statistics of a specified video within a specified time range. The statistics include total playback duration, number of playbacks, number of unique viewers, and playback duration distribution.
//
// Description:
//
// - Currently, the service address supports only **China (Shanghai)**.
//
// - Only playback data of videos that use ApsaraVideo Player SDK and rank in the top 1,000 by daily playback count is supported.
//
// - Based on UTC+8, playback statistics for the previous day are generated at 9:00 AM each day.
//
//   - Only data within the last 2 years (730 days) can be queried, and the maximum time span between the start time and end time is 180 days.
//
// @param request - DescribePlayVideoStatisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayVideoStatisResponse
func (client *Client) DescribePlayVideoStatisWithContext(ctx context.Context, request *DescribePlayVideoStatisRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayVideoStatisResponse, _err error) {
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
// Queries usage data of AI processing services such as automated review and media fingerprint.
//
// Description:
//
// - Currently, the service address supports only **China (Shanghai)**.
//
// - If the interval between the start time and end time is within 7 days, hourly data is returned. If the interval is greater than 7 days, daily data is returned. The maximum interval is 31 days.
//
// @param request - DescribeVodAIDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodAIDataResponse
func (client *Client) DescribeVodAIDataWithContext(ctx context.Context, request *DescribeVodAIDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodAIDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the certificate list information.
//
// Description:
//
// - Currently, the service address is supported only in the **China (Shanghai)*	- region.
//
// @param request - DescribeVodCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodCertificateListResponse
func (client *Client) DescribeVodCertificateListWithContext(ctx context.Context, request *DescribeVodCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodCertificateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to query the bandwidth data of an accelerated domain name. Compared with the DescribeVodDomainRealTimeBpsData operation, this operation supports a longer time range for historical data queries (up to 366 days) but provides a larger data time granularity (minimum of 5 minutes) and higher data latency.
//
// Description:
//
// - Currently, this operation is available only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data for the past 24 hours by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// If you specify `StartTime` and `EndTime` without Settings for `Interval`, the default time granularity of returned data, the queryable historical data time range, and the data latency are as follows:
//
// |Time granularity  |Time span per query   |  Queryable historical data time range  |  Data latency   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes  | Time span per query < 3 days   |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time span per query < 31 days  |186 days  |Generally 3-4 hours  |
//
// | 1 day | 31 days ≤ Time span per query ≤ 90 days |366 days  |Generally 4 hours, no more than 24 hours  |
//
// @param request - DescribeVodDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainBpsDataResponse
func (client *Client) DescribeVodDomainBpsDataWithContext(ctx context.Context, request *DescribeVodDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainBpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries bandwidth data of accelerated domain names by protocol type.
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data of up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data of the last 24 hours by default. If you specify `StartTime` and `EndTime`, this operation returns data of the specified time range.
//
// **Time granularity of returned data**
//
// If you specify `StartTime` and `EndTime` without configuring `Interval`, the default time granularity, the maximum time range for historical data queries, and the data delay are as follows:
//
// |Time granularity  |Time range per query   |  Maximum time range for historical data queries  |  Data delay   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes  | Time range per query < 3 days   |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time range per query < 31 days  |186 days  |Typically 3 to 4 hours  |
//
// | 1 day | 31 days ≤ Time range per query ≤ 366 days |366 days  |Typically 4 hours, up to 24 hours  |
//
// @param request - DescribeVodDomainBpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainBpsDataByLayerResponse
func (client *Client) DescribeVodDomainBpsDataByLayerWithContext(ctx context.Context, request *DescribeVodDomainBpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainBpsDataByLayerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the certificate information of a specified accelerated domain name.
//
// Description:
//
// Currently, the only supported service region is **China (Shanghai)**.
//
// @param request - DescribeVodDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainCertificateInfoResponse
func (client *Client) DescribeVodDomainCertificateInfoWithContext(ctx context.Context, request *DescribeVodDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainCertificateInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.HeraApiAutoVersion) {
		query["HeraApiAutoVersion"] = request.HeraApiAutoVersion
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
// Queries domain name configurations. You can query multiple feature configurations in a single request.
//
// Description:
//
// - Currently, the service address supports only **China (Shanghai)**.
//
// @param request - DescribeVodDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainConfigsResponse
func (client *Client) DescribeVodDomainConfigsWithContext(ctx context.Context, request *DescribeVodDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the basic configuration information of a specified China domain name for video-on-demand (VOD) acceleration.
//
// Description:
//
// Currently, the China service address supports only **China (Shanghai)**.
//
// @param request - DescribeVodDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainDetailResponse
func (client *Client) DescribeVodDomainDetailWithContext(ctx context.Context, request *DescribeVodDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the byte hit ratio (percentage of hit bytes) of an accelerated domain name.
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, data from the past 24 hours is returned by default. If you specify `StartTime` and `EndTime`, data for the specified time range is returned.
//
// **Time granularity of returned data**
//
// If you specify `StartTime` and `EndTime` without setting `Interval`, the default time granularity, the maximum time range for historical data queries, and the data delay are as follows:
//
// |Time granularity  |Time span per query   |  Maximum time range for historical data queries  |  Data delay   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes  | Time span per query < 3 days   |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time span per query < 31 days  |186 days  |Generally 3 to 4 hours  |
//
// | 1 day | 31 days ≤ Time span per query ≤ 366 days |366 days  |Generally 4 hours, up to 24 hours  |
//
// @param request - DescribeVodDomainHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainHitRateDataResponse
func (client *Client) DescribeVodDomainHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainHitRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the download URLs of raw CDN access logs for a specified domain name.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - For details about log formats and latency, refer to [Log Management](https://help.aliyun.com/document_detail/86099.html).
//
// - If you do not specify StartTime and EndTime, log data from the past 24 hours is returned by default.
//
// - StartTime and EndTime must be specified together to query logs within the specified time range.
//
// @param request - DescribeVodDomainLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainLogResponse
func (client *Client) DescribeVodDomainLogWithContext(ctx context.Context, request *DescribeVodDomainLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the 95th percentile bandwidth monitoring data for accelerated domain names.
//
// Description:
//
// *Three query methods are available:**
//
// - When both StartTime and EndTime are specified: if the difference between EndTime and StartTime is within 24 hours, the 95th percentile bandwidth of the day that contains StartTime is returned. Otherwise, the 95th percentile bandwidth of the month that contains StartTime is returned.
//
// - When both TimePoint and Cycle are specified, the 95th percentile bandwidth of the cycle that contains TimePoint is returned.
//
// - When StartTime and EndTime are specified with an additional Cycle parameter, the 95th percentile bandwidth for all specified cycles within the query range is returned.
//
// If none of these three methods are specified, the 95th percentile bandwidth of the past 24 hours is returned by default.
//
// - Maximum query span: 90 days.
//
// - Minimum query granularity: 1 day.
//
// - Maximum query range: 90 days.
//
// - Maximum number of calls per user per second: 100.
//
// - Data unit: bit/s.
//
// @param request - DescribeVodDomainMax95BpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainMax95BpsDataResponse
func (client *Client) DescribeVodDomainMax95BpsDataWithContext(ctx context.Context, request *DescribeVodDomainMax95BpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainMax95BpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the queries per second (QPS) of accelerated domain names at a 5-minute granularity. Data from the last 90 days is supported.
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// - Maximum call frequency per user: 100 calls per second.
//
// - If you do not specify StartTime and EndTime, this operation returns data from the last 24 hours. If you specify StartTime and EndTime, this operation returns data for the specified time range.
//
// **Supported time granularities**
//
// The Interval request parameter supports different data time granularities based on the maximum time range per query. The following table describes the queryable historical data time range and data latency for each time granularity:
//
// |Time granularity	|Maximum time range per query	|Queryable historical data time range	|Data latency
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes	|3 days	|93 days	|15 minutes
//
// |1 hour	|31 days	|186 days	|Typically 3-4 hours
//
// |1 day	|366 days	|366 days	|Typically 4 hours, no more than 24 hours
//
// *********
//
// @param request - DescribeVodDomainQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainQpsDataResponse
func (client *Client) DescribeVodDomainQpsDataWithContext(ctx context.Context, request *DescribeVodDomainQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainQpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to query the bandwidth data of an accelerated domain name. Compared with the DescribeVodDomainBpsData operation, this operation supports a smaller time granularity (minimum of 1 minute), lower data latency (minimum of 5 minutes), but a shorter historical data time range (up to 186 days).
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data for the last hour by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// The time granularity of returned data varies based on the time range specified by `StartTime` and `EndTime`. The following table describes the queryable historical data time range and data latency for each time granularity:
//
// |Time granularity  |Time range per query   |  Queryable historical data time range  |  Data latency   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |1 minute  | Time range per query ≤ 1 hour   |7 days  |5 minutes  |
//
// |5 minutes  | 1 hour < Time range per query < 3 days  |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time range per query < 31 days |186 days  |Generally 3 to 4 hours  |
//
// @param request - DescribeVodDomainRealTimeBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeBpsDataResponse
func (client *Client) DescribeVodDomainRealTimeBpsDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeBpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the byte hit rate data of accelerated domain names. This operation supports a minimum time granularity of 1 minute, with a data delay of at least 5 minutes, and allows you to query data from the last 186 days.
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 100 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data from the last hour by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// The time granularity of returned data varies based on the time range specified by `StartTime` and `EndTime`. The following table describes the queryable historical data range and data delay for each time granularity:
//
// |Time granularity  |Time range per query   |  Queryable historical data range  |  Data delay   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |1 minute  | Time range per query ≤ 1 hour   |7 days  |5 minutes  |
//
// |5 minutes  | 1 hour < Time range per query < 3 days  |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time range per query < 31 days |186 days  |Typically 3 to 4 hours  |
//
// @param request - DescribeVodDomainRealTimeByteHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeByteHitRateDataResponse
func (client *Client) DescribeVodDomainRealTimeByteHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeByteHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeByteHitRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries real-time access data for accelerated domain names, including QPS, bandwidth, and HTTP status code data within the last 7 days.
//
// Description:
//
// - Currently, this operation is available only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 20 domain names at a time.
//
// - The maximum time range per query (the time range between StartTime and EndTime) is 10 minutes, and data is returned at a time granularity of 1 minute.
//
// - Only data within the last 7 days can be queried.
//
// @param request - DescribeVodDomainRealTimeDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeDetailDataResponse
func (client *Client) DescribeVodDomainRealTimeDetailDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeDetailDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeDetailDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the total number and proportion of HTTP status codes for an accelerated domain name. This operation supports a minimum data time granularity of 1 minute, with a data delay of at least 5 minutes, and allows you to query data from the last 186 days.
//
// Description:
//
// - Currently, this operation is available only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 100 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data from the last hour by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// The time granularity of returned data varies based on the time range specified by `StartTime` and `EndTime`. The following table describes the queryable historical data range and data delay for each time granularity:
//
// |Time granularity  |Time range per query   |  Queryable historical data range  |  Data delay   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |1 minute  | Time range per query ≤ 1 hour   |7 days  |5 minutes  |
//
// |5 minutes  | 1 hour < Time range per query < 3 days  |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time range per query < 31 days |186 days  |Typically 3 to 4 hours  |
//
// @param request - DescribeVodDomainRealTimeHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeVodDomainRealTimeHttpCodeDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeHttpCodeDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invoke this operation to query the queries per second (QPS) data for access to an accelerated domain name. This operation supports a minimum data time granularity of 1 minute, with a data delay of at least 5 minutes, and allows you to query data from the last 186 days.
//
// Description:
//
// - Currently, this operation is available only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data from the last hour by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// The time granularity of returned data varies based on the time range specified by `StartTime` and `EndTime`. The following table describes the queryable historical data range and data delay for each time granularity:
//
// |Time granularity  |Time range per query   |  Queryable historical data range  |  Data delay   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |1 minute  | Time range per query ≤ 1 hour   |7 days  |5 minutes  |
//
// |5 minutes  | 1 hour < Time range per query < 3 days  |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time range per query < 31 days |186 days  |Typically 3 to 4 hours  |
//
// @param request - DescribeVodDomainRealTimeQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeQpsDataResponse
func (client *Client) DescribeVodDomainRealTimeQpsDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeQpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the request hit ratio data for an accelerated domain name. The minimum time granularity for data queried by this operation is 1 minute. The data latency is at least 5 minutes. You can query data for up to the last 186 days.
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// - Batch queries are supported. You can query data for up to 100 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data for the last 1 hour by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// - The Go language uses the POST method by default. Manually change the request method to GET by declaring `request.Method="GET"`.
//
// **Time granularity of returned data**
//
// The time granularity of returned data varies based on the time range specified by `StartTime` and `EndTime`. The following table describes the queryable historical data range and data latency for each time granularity:
//
// |Time granularity  |Time range per query   |  Queryable historical data range  |  Data latency   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |1 minute  | Time range per query ≤ 1 hour   |7 days  |5 minutes  |
//
// |5 minutes  | 1 hour < Time range per query < 3 days  |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time range per query < 31 days |186 days  |Typically 3 to 4 hours  |
//
// @param request - DescribeVodDomainRealTimeReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeReqHitRateDataResponse
func (client *Client) DescribeVodDomainRealTimeReqHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeReqHitRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to query the traffic data of an accelerated domain name. Compared with the DescribeVodDomainTrafficData operation, this operation supports a smaller time granularity (minimum of 1 minute), lower data latency (minimum of 5 minutes), but a shorter historical data range (up to 186 days).
//
// Description:
//
// - The service address of this operation supports only **China (Shanghai)**.
//
// - Batch queries are supported. You can query data for up to 100 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data for the last hour by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// The time granularity of returned data varies based on the time range specified by `StartTime` and `EndTime`. The following table describes the queryable historical data range and data latency for each time granularity:
//
// |Time granularity  |Time range per query   |  Queryable historical data range  |  Data latency   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |1 minute  | Time range per query ≤ 1 hour   |7 days  |5 minutes  |
//
// |5 minutes  | 1 hour < Time range per query < 3 days  |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time range per query < 31 days |186 days  |Generally 3 to 4 hours  |
//
// @param request - DescribeVodDomainRealTimeTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainRealTimeTrafficDataResponse
func (client *Client) DescribeVodDomainRealTimeTrafficDataWithContext(ctx context.Context, request *DescribeVodDomainRealTimeTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainRealTimeTrafficDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the request hit ratio (percentage of hit requests) of an accelerated domain name.
//
// Description:
//
// - Currently, this operation supports only the following service address: **China (Shanghai)**.
//
// - Batch query is supported. You can query data of up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data of the last 24 hours by default. If you specify `StartTime` and `EndTime`, this operation returns data of the specified time range.
//
// **Time granularity of returned data**
//
// Based on the time span specified by `StartTime` and `EndTime`, and when `Interval` is not set, the default time granularity of returned data, the queryable historical data time range, and data latency are as follows:
//
// |Time granularity  |Time span per query   |  Queryable historical data time range  |  Data latency   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes  | Time span per query < 3 days   |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time span per query < 31 days  |186 days  |Generally 3 to 4 hours  |
//
// | 1 day | 31 days ≤ Time span per query ≤ 90 days |366 days  |Generally 4 hours, no more than 24 hours  |
//
// @param request - DescribeVodDomainReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainReqHitRateDataResponse
func (client *Client) DescribeVodDomainReqHitRateDataWithContext(ctx context.Context, request *DescribeVodDomainReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainReqHitRateDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the back-to-origin bandwidth data of accelerated domain names.
//
// Description:
//
// - Currently, this operation supports only the following service address: **China (Shanghai)**.
//
// - Batch queries are supported. You can query data for up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data from the past 24 hours by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// Based on the time span specified by `StartTime` and `EndTime`, and when `Interval` is not set, the default time granularity of returned data, the queryable historical data range, and data latency are as follows:
//
// |Time granularity  |Time span per query   |  Queryable historical data range  |  Data latency   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes  | Time span per query < 3 days   |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time span per query < 31 days  |186 days  |Generally 3 to 4 hours  |
//
// | 1 day | 31 days ≤ Time span per query ≤ 366 days |366 days  |Generally 4 hours, up to 24 hours  |
//
// @param request - DescribeVodDomainSrcBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainSrcBpsDataResponse
func (client *Client) DescribeVodDomainSrcBpsDataWithContext(ctx context.Context, request *DescribeVodDomainSrcBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainSrcBpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the back-to-origin traffic data of accelerated domain names.
//
// Description:
//
// - Currently, this operation supports only the following service address: **China (Shanghai)**.
//
// - Batch queries are supported. You can query data for up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, this operation returns data from the past 24 hours by default. If you specify `StartTime` and `EndTime`, this operation returns data for the specified time range.
//
// **Time granularity of returned data**
//
// Based on the time span specified by `StartTime` and `EndTime`, and when `Interval` is not set at the same time, the default time granularity, the queryable historical data range, and data delay are as follows:
//
// |Time granularity  |Time span per query   |  Queryable historical data range  |  Data delay   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes  | Time span per query < 3 days   |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time span per query < 31 days  |186 days  |Generally 3 to 4 hours  |
//
// | 1 day | 31 days ≤ Time span per query ≤ 366 days |366 days  |Generally 4 hours, up to 24 hours  |
//
// @param request - DescribeVodDomainSrcTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainSrcTrafficDataResponse
func (client *Client) DescribeVodDomainSrcTrafficDataWithContext(ctx context.Context, request *DescribeVodDomainSrcTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainSrcTrafficDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to query the traffic data of an accelerated domain name. Compared with the DescribeVodDomainRealTimeTrafficData operation, this operation supports a longer historical data query range (up to 366 days) but provides a larger data time granularity (minimum of 5 minutes) and higher data latency.
//
// Description:
//
// - The service address of this operation supports only **China (Shanghai)**.
//
// - Batch queries are supported. You can query data for up to 500 domain names at a time.
//
// - If you do not specify `StartTime` and `EndTime`, data of the last 24 hours is returned by default. If you specify `StartTime` and `EndTime`, data of the specified time range is returned.
//
// **Data time granularity of returned data**
//
// The following table describes the default data time granularity of returned data, the queryable historical data range, and the data latency based on the time span specified by `StartTime` and `EndTime` when `Interval` is not set:
//
// |Time granularity  |Time span per query   |  Queryable historical data range  |  Data latency   |
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes  | Time span per query < 3 days   |93 days  |15 minutes  |
//
// |1 hour  | 3 days ≤ Time span per query < 31 days  |186 days  |Typically 3 to 4 hours  |
//
// | 1 day | 31 days ≤ Time span per query ≤ 366 days |366 days  |Typically 4 hours, up to 24 hours  |
//
// @param request - DescribeVodDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainTrafficDataResponse
func (client *Client) DescribeVodDomainTrafficDataWithContext(ctx context.Context, request *DescribeVodDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainTrafficDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries acceleration traffic or bandwidth usage data.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - Batch domain name queries are supported. Separate multiple domain names with commas (,). You can query up to 100 domain names at a time. If this parameter is left empty, data for all domain names under the account is returned.
//
// - You can query data for up to the last year. The maximum time span for a single query is 3 months. If the query time range is 1 to 3 days, data is returned at hourly granularity. If the query time range is 4 days or more, data is returned at daily granularity.
//
// @param request - DescribeVodDomainUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodDomainUsageDataResponse
func (client *Client) DescribeVodDomainUsageDataWithContext(ctx context.Context, request *DescribeVodDomainUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodDomainUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the usage data of online editing in ApsaraVideo VOD.
//
// Description:
//
// - Single user call frequency: 10 calls per second.
//
// **Supported time granularities**:
//
// The adaptive time granularity and the maximum time range for historical data queries vary based on the maximum time span per query.
//
// | Time granularity          | Maximum time span per query            | Maximum time range for historical data queries    |
//
// | -------------- | -------------- | ------ |
//
// | 1 hour       | 7 days      |   31 days  |
//
// | 1 day  | 31 days     |    366 days  |
//
// @param request - DescribeVodEditingUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodEditingUsageDataResponse
func (client *Client) DescribeVodEditingUsageDataWithContext(ctx context.Context, request *DescribeVodEditingUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodEditingUsageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves playback data of an audio or video file on a specified date by media ID (audio or video ID), including the number of unique visitors, average plays per user, total plays, average play duration per user, and total play duration.
//
// Description:
//
// - Currently, this operation is available only in the **China (Shanghai)*	- region.
//
// - Only playback data collected by ApsaraVideo Player SDK is supported. Traffic statistics for audio-only streams are not supported.
//
// - Only data within the last 30 days can be queried.
//
//		Notice: - Before calling this operation, make sure that ApsaraVideo Player SDK meets the following conditions:
//
//	  - Android Player SDK or iOS Player SDK
//
//	    - The Player SDK version is 5.4.9.2 or later.
//
//	    - A License for the Player SDK has been obtained and integrated. For more information, see [License management](https://help.aliyun.com/document_detail/469166.html).
//
//	    - The event tracking log reporting feature of the Player SDK is enabled. By default, this feature is enabled in ApsaraVideo Player SDK. For more information, see [Create a player for Android](~~311525#section-dc4-gp6-xk2~~) and [Create a player for iOS](~~313855#section-cmf-k7d-jg5~~).
//
//	  - Web Player SDK
//
//	     - The Player SDK version is 2.16.0 or later.
//
//	    - A License for **Playback Quality Monitoring*	- has been obtained and integrated. Submit the [Web Player SDK value-added service application form](https://yida.alibaba-inc.com/o/webplayer#/) to apply. For the License integration method, see the `license` property in [Web SDK API reference](~~125572#section-3ty-gwp-6pa~~).
//
//	    - The event tracking log reporting feature of the Player SDK is enabled. By default, this feature is enabled in ApsaraVideo Player SDK.
//
// @param request - DescribeVodMediaPlayDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodMediaPlayDataResponse
func (client *Client) DescribeVodMediaPlayDataWithContext(ctx context.Context, request *DescribeVodMediaPlayDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodMediaPlayDataResponse, _err error) {
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
// Queries aggregated playback data of the player.
//
// @param request - DescribeVodPlayerCollectDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodPlayerCollectDataResponse
func (client *Client) DescribeVodPlayerCollectDataWithContext(ctx context.Context, request *DescribeVodPlayerCollectDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodPlayerCollectDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the dimension metadata of the player.
//
// @param request - DescribeVodPlayerDimensionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodPlayerDimensionDataResponse
func (client *Client) DescribeVodPlayerDimensionDataWithContext(ctx context.Context, request *DescribeVodPlayerDimensionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodPlayerDimensionDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries player metric data.
//
// @param request - DescribeVodPlayerMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodPlayerMetricDataResponse
func (client *Client) DescribeVodPlayerMetricDataWithContext(ctx context.Context, request *DescribeVodPlayerMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodPlayerMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries access data for an accelerated domain name by ISP or region, including bandwidth, average response rate, page views, cache hit ratio, and request hit ratio.
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// - The maximum time range for a single query (the time range between StartTime and EndTime) is 1 hour.
//
// **Supported time granularities**
//
// Based on the time range specified by `StartTime` and `EndTime`, the default data timestamp granularity, queryable historical data range, and data latency are as follows:
//
// |Time granularity	|Time range per query|Queryable historical data range|Data latency|
//
// | ------------- |------------   | ----------- | ----------- |
//
// |5 minutes	|≤ 1 hour	|93 days	|15 minutes|
//
// @param request - DescribeVodRangeDataByLocateAndIspServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodRangeDataByLocateAndIspServiceResponse
func (client *Client) DescribeVodRangeDataByLocateAndIspServiceWithContext(ctx context.Context, request *DescribeVodRangeDataByLocateAndIspServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodRangeDataByLocateAndIspServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the maximum number and remaining daily quota of URLs and directories for purge and prefetch operations.
//
// Description:
//
// - Currently, the service address supports only **China (Shanghai)**.
//
// - Purge and prefetch operations include the [RefreshVodObjectCaches](https://help.aliyun.com/document_detail/69215.html) purge operation and the [PreloadVodObjectCaches](https://help.aliyun.com/document_detail/69211.html) prefetch operation.
//
// @param request - DescribeVodRefreshQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodRefreshQuotaResponse
func (client *Client) DescribeVodRefreshQuotaWithContext(ctx context.Context, request *DescribeVodRefreshQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodRefreshQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries whether purge and prefetch tasks have taken effect.
//
// Description:
//
// - Currently, the service address supports only **China (Shanghai)**.
//
// - If neither Taskid nor Objectpath is specified, the first page of data (20 entries) within the last 3 days is returned by default. Taskid and Objectpath can be specified at the same time.
//
// - When DomainName or Status is specified, ObjectType is required.
//
// @param request - DescribeVodRefreshTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodRefreshTasksResponse
func (client *Client) DescribeVodRefreshTasksWithContext(ctx context.Context, request *DescribeVodRefreshTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodRefreshTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries certificate list information by domain name.
//
// Description:
//
// This operation currently supports only the **China (Shanghai)*	- region.
//
// @param request - DescribeVodSSLCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodSSLCertificateListResponse
func (client *Client) DescribeVodSSLCertificateListWithContext(ctx context.Context, request *DescribeVodSSLCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodSSLCertificateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the usage of media asset management, including storage space and outbound storage traffic.
//
// Description:
//
// - Currently, the service address supports only **China (Shanghai)**.
//
// - If the interval between the start time and end time is within 7 days, hourly data is returned. If the interval is greater than 7 days, daily data is returned. The maximum interval is 31 days.
//
// @param request - DescribeVodStorageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodStorageDataResponse
func (client *Client) DescribeVodStorageDataWithContext(ctx context.Context, request *DescribeVodStorageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodStorageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the usage of tiered storage for media asset management.
//
// Description:
//
// - Currently, the service is available only in the **China (Shanghai)*	- region.
//
// - If the query time range is within 7 days, hourly data is returned. If the query time range is greater than 7 days, daily data is returned. The maximum time range is 31 days.
//
// @param request - DescribeVodTieringStorageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodTieringStorageDataResponse
func (client *Client) DescribeVodTieringStorageDataWithContext(ctx context.Context, request *DescribeVodTieringStorageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodTieringStorageDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the retrieval data usage of tiered storage in media asset management.
//
// Description:
//
// > - Currently, the service address supports only **China (Shanghai)**.
//
// > - If the query time interval is within 7 days, data at the hour granularity is returned. If the query time interval is greater than 7 days, data at the day granularity is returned. The maximum interval is 31 days.
//
// @param request - DescribeVodTieringStorageRetrievalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodTieringStorageRetrievalDataResponse
func (client *Client) DescribeVodTieringStorageRetrievalDataWithContext(ctx context.Context, request *DescribeVodTieringStorageRetrievalDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodTieringStorageRetrievalDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries transcoding usage data.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - If the interval between the start time and end time is within 7 days, hourly data is returned. If the interval is greater than 7 days, daily data is returned. The maximum interval is 31 days.
//
// @param request - DescribeVodTranscodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodTranscodeDataResponse
func (client *Client) DescribeVodTranscodeDataWithContext(ctx context.Context, request *DescribeVodTranscodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodTranscodeDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of all acceleration domain names under your account for ApsaraVideo VOD.
//
// Description:
//
// - Supports fuzzy match filtering by domain name and filtering by domain name status.
//
// - This operation currently supports only the following region: **China (Shanghai)**.
//
// @param request - DescribeVodUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodUserDomainsResponse
func (client *Client) DescribeVodUserDomainsWithContext(ctx context.Context, request *DescribeVodUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodUserDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the IP list of a domain name.
//
// Description:
//
// This operation is supported only in the China (Shanghai) region.
//
// @param request - DescribeVodUserVipsByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodUserVipsByDomainResponse
func (client *Client) DescribeVodUserVipsByDomainWithContext(ctx context.Context, request *DescribeVodUserVipsByDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodUserVipsByDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the domain ownership verification content.
//
// Description:
//
// - Currently, this operation is supported only in the **China (Shanghai)*	- region.
//
// @param request - DescribeVodVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodVerifyContentResponse
func (client *Client) DescribeVodVerifyContentWithContext(ctx context.Context, request *DescribeVodVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodVerifyContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to revoke application authorization from a specified account identity (Resource Access Management (RAM) user or RAM role).
//
// Description:
//
//	Notice: Each Resource Access Management (RAM) user or RAM role can be granted permissions for up to 10 applications.
//
// -  If the policy name is **VODAppAdministratorAccess**, **AppId*	- is optional. For other policies, **AppId*	- is required.
//
// @param request - DetachAppPolicyFromIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachAppPolicyFromIdentityResponse
func (client *Client) DetachAppPolicyFromIdentityWithContext(ctx context.Context, request *DetachAppPolicyFromIdentityRequest, runtime *dara.RuntimeOptions) (_result *DetachAppPolicyFromIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Generates a key for secure download. The secure download feature of ApsaraVideo Player SDK encrypts videos downloaded to local devices by using a key file. The encrypted videos can only be decrypted and played by using the key file generated by the unique app that is bindable in advance. This effectively protects video content and prevents downloaded videos from being maliciously played or distributed.
//
// Description:
//
// - To use the secure download feature, first enable the download feature in the ApsaraVideo VOD console and set the download method to secure download. For more information, see [Download settings](https://help.aliyun.com/document_detail/86107.html).
//
// - After generating a key for secure download, configure the key in ApsaraVideo Player SDK. For more information, see [Secure download](https://help.aliyun.com/document_detail/124735.html).
//
// @param request - GenerateDownloadSecretKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDownloadSecretKeyResponse
func (client *Client) GenerateDownloadSecretKeyWithContext(ctx context.Context, request *GenerateDownloadSecretKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateDownloadSecretKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes the operation to generate a random KMS data key (DK) for ApsaraVideo VOD HLS encryption.
//
// @param request - GenerateKMSDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateKMSDataKeyResponse
func (client *Client) GenerateKMSDataKeyWithContext(ctx context.Context, request *GenerateKMSDataKeyRequest, runtime *dara.RuntimeOptions) (_result *GenerateKMSDataKeyResponse, _err error) {
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
// Queries the list of AI image processing tasks.
//
// Description:
//
// -  Currently, this operation is supported in the following regions: **China (Beijing)*	- and **China (Shanghai)**.
//
// - Call the [SubmitAIImageJob](~~SubmitAIImageJob~~) operation to submit an AI image processing task before you call this operation to query the list of AI image tasks.
//
// - You can query up to 10 AI image processing tasks at a time.
//
// @param request - GetAIImageJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIImageJobsResponse
func (client *Client) GetAIImageJobsWithContext(ctx context.Context, request *GetAIImageJobsRequest, runtime *dara.RuntimeOptions) (_result *GetAIImageJobsResponse, _err error) {
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
// Queries an automated review job. After you submit an AI job, the job is processed asynchronously. You can call this operation to query job information in real time.
//
// Description:
//
// <props="intl">
//
// - This operation is supported only in the Singapore region.
//
// - Image resources in automated review job results are retained in the free storage provided by ApsaraVideo VOD for only two weeks. After two weeks, the images are automatically deleted.
//
// @param request - GetAIMediaAuditJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIMediaAuditJobResponse
func (client *Client) GetAIMediaAuditJobWithContext(ctx context.Context, request *GetAIMediaAuditJobRequest, runtime *dara.RuntimeOptions) (_result *GetAIMediaAuditJobResponse, _err error) {
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
// Queries an AI template.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - Obtain the AI template ID first, and then call this operation to query the configuration information of the AI template.
//
// @param request - GetAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITemplateResponse
func (client *Client) GetAITemplateWithContext(ctx context.Context, request *GetAITemplateRequest, runtime *dara.RuntimeOptions) (_result *GetAITemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the results of smart tagging for a video.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)*	- and **China (Shanghai)**.
//
// - Retrieves smart tagging results by video ID.
//
// @param request - GetAIVideoTagResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIVideoTagResultResponse
func (client *Client) GetAIVideoTagResultWithContext(ctx context.Context, request *GetAIVideoTagResultRequest, runtime *dara.RuntimeOptions) (_result *GetAIVideoTagResultResponse, _err error) {
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
// Queries application information by application ID.
//
// Description:
//
// - Batch queries are supported.
//
// - AppIds supports a maximum of 10 IDs.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetAppInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInfosResponse
func (client *Client) GetAppInfosWithContext(ctx context.Context, request *GetAppInfosRequest, runtime *dara.RuntimeOptions) (_result *GetAppInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the playback key of an application.
//
// @param request - GetAppPlayKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppPlayKeyResponse
func (client *Client) GetAppPlayKeyWithContext(ctx context.Context, request *GetAppPlayKeyRequest, runtime *dara.RuntimeOptions) (_result *GetAppPlayKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
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
		Action:      dara.String("GetAppPlayKey"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppPlayKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the basic information and access URLs of multiple auxiliary media assets in a batch by specifying their IDs after the assets such as watermark images, subtitle files, and materials are uploaded to ApsaraVideo VOD.
//
// Description:
//
// You can retrieve information about up to 20 auxiliary media assets at a time.
//
// @param request - GetAttachedMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttachedMediaInfoResponse
func (client *Client) GetAttachedMediaInfoWithContext(ctx context.Context, request *GetAttachedMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAttachedMediaInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the history of manual review records.
//
// @param request - GetAuditHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuditHistoryResponse
func (client *Client) GetAuditHistoryWithContext(ctx context.Context, request *GetAuditHistoryRequest, runtime *dara.RuntimeOptions) (_result *GetAuditHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to perform a filtered query for specified categorization information by ID or type, and retrieves the list of its subcategories (next-level categories).
//
// @param request - GetCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCategoriesResponse
func (client *Client) GetCategoriesWithContext(ctx context.Context, request *GetCategoriesRequest, runtime *dara.RuntimeOptions) (_result *GetCategoriesResponse, _err error) {
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
// Retrieves the full traffic data of media assets for a specified date and region. The data is generated based on CDN traffic logs and primarily reflects the traffic consumption of videos. The generated CSV file contains the following information: date, video ID, domain name, traffic, application ID, and category ID. You can download the file to your local machine for scenarios such as operational data analytics.
//
// Description:
//
// - Currently, the service address of this operation only supports: **China (Shanghai)**.
//
// - Only data within the past 90 days can be queried (data starts from April 29, 2025).
//
// - The traffic data provided by this operation is raw traffic data. To align with billing traffic, multiply the data by a TCP coefficient of 1.1.
//
// @param request - GetDailyPlayRegionStatisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDailyPlayRegionStatisResponse
func (client *Client) GetDailyPlayRegionStatisWithContext(ctx context.Context, request *GetDailyPlayRegionStatisRequest, runtime *dara.RuntimeOptions) (_result *GetDailyPlayRegionStatisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the default AI template.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - Currently, only the default AI template for automated review can be queried.
//
// @param request - GetDefaultAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultAITemplateResponse
func (client *Client) GetDefaultAITemplateWithContext(ctx context.Context, request *GetDefaultAITemplateRequest, runtime *dara.RuntimeOptions) (_result *GetDefaultAITemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the results of a digital watermarking (copyright watermark or tracing watermark) extraction job, including the job status and the successfully extracted watermark text.
//
// Description:
//
// - Currently, this operation is available only in the China (Shanghai) and China (Beijing) regions.
//
// - After you call the [SubmitDigitalWatermarkExtractJob](~~SubmitDigitalWatermarkExtractJob~~) operation to extract the copyright watermark or tracing watermark from a video, call this operation to query the extracted watermark text.
//
// - Only watermark extraction jobs from the last 2 years can be queried.
//
// @param request - GetDigitalWatermarkExtractResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDigitalWatermarkExtractResultResponse
func (client *Client) GetDigitalWatermarkExtractResultWithContext(ctx context.Context, request *GetDigitalWatermarkExtractResultRequest, runtime *dara.RuntimeOptions) (_result *GetDigitalWatermarkExtractResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of an online editing project (video editing task).
//
// @param request - GetEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEditingProjectResponse
func (client *Client) GetEditingProjectWithContext(ctx context.Context, request *GetEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *GetEditingProjectResponse, _err error) {
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
// Retrieves the list of materials to be edited in an online editing project.
//
// Description:
//
// During the editing process, materials can be added to the timeline but are not necessarily fully used.
//
// @param request - GetEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEditingProjectMaterialsResponse
func (client *Client) GetEditingProjectMaterialsWithContext(ctx context.Context, request *GetEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *GetEditingProjectMaterialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the basic information and access URL of an image by image ID after the image is uploaded to ApsaraVideo VOD.
//
// Description:
//
// This operation only supports querying information about images uploaded to ApsaraVideo VOD. To query information about snapshots generated from video snapshots, call the [ListSnapshots](~~ListSnapshots~~) operation.
//
// @param request - GetImageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageInfoResponse
func (client *Client) GetImageInfoWithContext(ctx context.Context, request *GetImageInfoRequest, runtime *dara.RuntimeOptions) (_result *GetImageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the basic information and access URLs of multiple images by image ID after the images are uploaded to ApsaraVideo VOD.
//
// Description:
//
// - This operation only supports querying information about images uploaded to ApsaraVideo VOD. To query information about snapshots generated from video snapshots, call the [ListSnapshots](~~ListSnapshots~~) operation.
//
// - You can query information about up to 20 images at a time.
//
// @param request - GetImageInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageInfosResponse
func (client *Client) GetImageInfosWithContext(ctx context.Context, request *GetImageInfosRequest, runtime *dara.RuntimeOptions) (_result *GetImageInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the details of an asynchronous task by job ID.
//
// Description:
//
// *Usage notes**
//
// This operation supports querying asynchronous task data from the last 6 months. Supported task types: transcoding tasks, snapshot tasks, AI tasks, and workflow tasks.
//
// **QPS limit**
//
// The single-user QPS limit for this operation is 15 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDetailResponse
func (client *Client) GetJobDetailWithContext(ctx context.Context, request *GetJobDetailRequest, runtime *dara.RuntimeOptions) (_result *GetJobDetailResponse, _err error) {
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
// Retrieves the results of video AI analysis.
//
// @param request - GetMediaAiAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAiAnalysisResponse
func (client *Client) GetMediaAiAnalysisWithContext(ctx context.Context, request *GetMediaAiAnalysisRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAiAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
	}

	if !dara.IsNil(request.ResultTypes) {
		query["ResultTypes"] = request.ResultTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaAiAnalysis"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaAiAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of audio automated review results.
//
// Description:
//
// ### Usage notes
//
// <props="china">After automated review is complete, if you have configured the [Automated review complete](https://help.aliyun.com/document_detail/89576.html) event notification, the callback URL is notified through a message callback. You can call this operation to query the details of audio review results.
//
// <props="intl">
//
// - This operation is supported only in the Singapore region.
//
// - After automated review is complete, if you have configured the [Automated review complete](https://help.aliyun.com/document_detail/89576.html) event notification, the callback URL is notified through a message callback. You can call this operation to query the details of audio review results.
//
// @param request - GetMediaAuditAudioResultDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditAudioResultDetailResponse
func (client *Client) GetMediaAuditAudioResultDetailWithContext(ctx context.Context, request *GetMediaAuditAudioResultDetailRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditAudioResultDetailResponse, _err error) {
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
// Retrieves the summary of automated review results.
//
// Description:
//
// <props="intl">
//
// ### Usage notes
//
// This operation is supported only in the Singapore region.
//
// ### QPS limit
//
// You can call this operation up to 20 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetMediaAuditResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditResultResponse
func (client *Client) GetMediaAuditResultWithContext(ctx context.Context, request *GetMediaAuditResultRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditResultResponse, _err error) {
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
// Retrieves the details of automated review results. You can call this operation to query the details of review results in real time.
//
// Description:
//
// - By default, only the review screenshot details of violating and suspected violating content are returned. No results are returned for compliant videos and images.
//
// - The image resources of review results are retained in the free storage provided by ApsaraVideo VOD for only 2 weeks. After 2 weeks, the images are automatically deleted.
//
// <props="intl">
//
// - This operation is supported only in the Singapore region.
//
// @param request - GetMediaAuditResultDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditResultDetailResponse
func (client *Client) GetMediaAuditResultDetailWithContext(ctx context.Context, request *GetMediaAuditResultDetailRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditResultDetailResponse, _err error) {
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
// Retrieves the timestamps of all screenshots that contain violations.
//
// Description:
//
// > By default, only screenshot details for violations and suspected violations are returned. No results are returned for compliant videos and images.
//
// <props="intl">
//
// This operation is supported only in the Singapore region.
//
// @param request - GetMediaAuditResultTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaAuditResultTimelineResponse
func (client *Client) GetMediaAuditResultTimelineWithContext(ctx context.Context, request *GetMediaAuditResultTimelineRequest, runtime *dara.RuntimeOptions) (_result *GetMediaAuditResultTimelineResponse, _err error) {
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
// Retrieves media fingerprint results. After a media fingerprint job is complete, you can call this operation to query the results in real time.
//
// Description:
//
// This operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// @param request - GetMediaDNAResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaDNAResultResponse
func (client *Client) GetMediaDNAResultWithContext(ctx context.Context, request *GetMediaDNAResultRequest, runtime *dara.RuntimeOptions) (_result *GetMediaDNAResultResponse, _err error) {
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
// Queries task information such as task status and filtering policies for a video purge or prefetch task.
//
// Description:
//
// ### Usage notes
//
// You can query task information for all audio or video files under a purge or prefetch task, or query task information for a specific audio or video file.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 50 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as needed. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetMediaRefreshJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaRefreshJobsResponse
func (client *Client) GetMediaRefreshJobsWithContext(ctx context.Context, request *GetMediaRefreshJobsRequest, runtime *dara.RuntimeOptions) (_result *GetMediaRefreshJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the callback method, callback URL, and event types of event notifications.
//
// Description:
//
// > For more information, see [Event notification development guide](https://help.aliyun.com/document_detail/55627.html).
//
// @param request - GetMessageCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageCallbackResponse
func (client *Client) GetMessageCallbackWithContext(ctx context.Context, request *GetMessageCallbackRequest, runtime *dara.RuntimeOptions) (_result *GetMessageCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the source file information of an audio or video file, including the file URL, resolution, and bitrate.
//
// Description:
//
// You can retrieve the complete source file information only after a video or audio stream is transcoded.
//
// @param request - GetMezzanineInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMezzanineInfoResponse
func (client *Client) GetMezzanineInfoWithContext(ctx context.Context, request *GetMezzanineInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMezzanineInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Retrieves the playback URL of an audio or video file by providing the audio or video ID, which can then be played using ApsaraVideo Player or a third-party player such as a system-native, open-source, or custom-built player.
//
// Description:
//
// - **Before using this operation, make sure you fully understand the billing methods and pricing of ApsaraVideo VOD. Directly downloading or playing videos from ApsaraVideo VOD playback URLs incurs outbound traffic fees. If no accelerated domain name is configured, refer to [Storage outbound traffic billing](~~188308#section-rwh-e88-f7j~~). If an accelerated domain name is configured, refer to [Acceleration service billing](~~188308#section-c5t-oq9-15e~~). If you have enabled storage transfer acceleration, directly downloading or playing videos from ApsaraVideo VOD playback URLs also incurs download acceleration fees. For billing details, refer to [Storage transfer acceleration billing](~~188310#section_sta_zm2_tsv~~).**
//
// - Only videos in the Normal state (the Status field value is Normal) can be played. For more information about playback URL descriptions and usage limits, refer to [Audio and video playback](https://help.aliyun.com/document_detail/57290.html).
//
// - When the [media storage](https://help.aliyun.com/document_detail/2392368.html) type is non-standard storage, set the StorageClass field of the PlayConfig parameter accordingly. For details, refer to [PlayConfig](~~86952#section-9g7-s9b-v7z~~).
//
// - If video playback is abnormal, call the [GetMezzanineInfo](~~GetMezzanineInfo~~) operation to check whether the video source file information is correct.
//
// <props="china">
//
// - To generate m3u8 tracing watermark video streams by calling this operation, submit a ticket to apply for activation. For information about how to submit a ticket, refer to [Contact us](https://help.aliyun.com/document_detail/464625.html). For more information about tracing watermarks, refer to [Digital watermarking](https://help.aliyun.com/document_detail/2527021.html).
//
// @param request - GetPlayInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPlayInfoResponse
func (client *Client) GetPlayInfoWithContext(ctx context.Context, request *GetPlayInfoRequest, runtime *dara.RuntimeOptions) (_result *GetPlayInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionType) {
		query["AdditionType"] = request.AdditionType
	}

	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.CodecName) {
		query["CodecName"] = request.CodecName
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

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Queries video transcoding summary of one or more audio or video files by their IDs, including video transcoding status and transcoding progress.
//
// Description:
//
// - Because an audio or video file may be transcoded multiple times, this operation returns only the most recent transcoding summary.
//
// - Batch queries are supported. You can query the transcoding summaries of up to 10 audio or video files at a time.
//
// - To query historical transcoding task information, call the [ListTranscodeTask](https://help.aliyun.com/document_detail/109120.html) operation.
//
// - **This operation supports querying transcoding task data only within the last year.**
//
// @param request - GetTranscodeSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeSummaryResponse
func (client *Client) GetTranscodeSummaryWithContext(ctx context.Context, request *GetTranscodeSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the details of transcoding jobs based on a transcoding task ID.
//
// Description:
//
// ### Usage notes
//
// **This operation only supports querying transcoding task data from the last year.**
//
// ### QPS limit
//
// A single user can perform a maximum of 15 queries per second (QPS). Throttling is triggered if this limit is exceeded, which may affect your business. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetTranscodeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeTaskResponse
func (client *Client) GetTranscodeTaskWithContext(ctx context.Context, request *GetTranscodeTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeTaskResponse, _err error) {
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
// Queries the details of a transcoding configuration by transcoding template group ID.
//
// Description:
//
// Retrieves information about a single template group, including the configuration information of all transcoding templates in the group.
//
// @param request - GetTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeTemplateGroupResponse
func (client *Client) GetTranscodeTemplateGroupWithContext(ctx context.Context, request *GetTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeTemplateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves URL upload information.
//
// Description:
//
// - Retrieves URL upload information by using the JobId returned during URL-based upload or the URL used for upload. The information includes the URL upload status, UserData, creation time, and completion time.
//
// - **This operation only supports querying upload task data within the last year.**
//
// - This operation currently supports only the following service regions: **China (Shanghai)*	- and **Singapore**.
//
// - After you call the [UploadMediaByURL](~~UploadMediaByURL~~) operation to upload a media file to ApsaraVideo VOD, you can call this operation to query the upload information of a specified media file by using the upload task IDs (`JobIds`) or the source file URLs (`UploadURLs`).
//
// - When calling this operation, you must specify either `JobIds` or `UploadURLs`. If both are specified, only `JobIds` is processed.
//
// - If the media upload fails, you can call this operation to view the error code and error message. If the media upload succeeds, you can call this operation to view the corresponding media ID.
//
// @param request - GetURLUploadInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetURLUploadInfosResponse
func (client *Client) GetURLUploadInfosWithContext(ctx context.Context, request *GetURLUploadInfosRequest, runtime *dara.RuntimeOptions) (_result *GetURLUploadInfosResponse, _err error) {
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
// Retrieves media upload details by media ID, such as upload time, upload ratio, and upload source. Batch retrieval is supported.
//
// Description:
//
// - This operation only supports retrieving upload details of audio and video files.
//
// - If audio or video files are uploaded through the ApsaraVideo VOD console, you can use this operation to retrieve information such as the upload ratio. If audio or video files are uploaded by using the upload SDK, only the following versions of the [upload SDK](https://help.aliyun.com/document_detail/52200.html) support this operation.
//
// > Only the server upload SDK supports this operation. The client upload SDK does not support this operation. The server upload SDK version requirements are as follows:
//
// > - Java upload SDK: version ≥ 1.4.4
//
// > - C++ upload SDK: version ≥ 1.0.0
//
// > - PHP upload SDK: version ≥ 1.0.2
//
// > - Python upload SDK: version ≥ 1.3.0
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 100 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - GetUploadDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadDetailsResponse
func (client *Client) GetUploadDetailsWithContext(ctx context.Context, request *GetUploadDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetUploadDetailsResponse, _err error) {
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
// Retrieves the information about a single audio or video file by audio or video ID, including the title, description, duration, thumbnail URL, status, creation time, size, snapshots, category, and tags.
//
// Description:
//
// After an audio or video file is uploaded, ApsaraVideo VOD analyzes the uploaded source file. Therefore, media asset information is generated asynchronously. You can configure an [event notification](https://help.aliyun.com/document_detail/55627.html) for [AudioVideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html). After you receive the [AudioVideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event notification, call this operation to retrieve the audio or video information.
//
// @param request - GetVideoInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoInfoResponse
func (client *Client) GetVideoInfoWithContext(ctx context.Context, request *GetVideoInfoRequest, runtime *dara.RuntimeOptions) (_result *GetVideoInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
	}

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
// Retrieves the information about multiple audio and video files at a time by audio or video ID, including the title, description, duration, thumbnail URL, status, creation time, size, snapshots, category, and tags.
//
// Description:
//
// - You can retrieve information about up to 20 audio and video files at a time.
//
// - After an audio or video file is uploaded, ApsaraVideo VOD analyzes the uploaded source file. Therefore, media asset information is generated asynchronously. You can configure the [AudioVideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) [event notification](https://help.aliyun.com/document_detail/55627.html). After you receive the [AudioVideoAnalysisComplete](https://help.aliyun.com/document_detail/99935.html) event notification, call this operation to retrieve the audio and video information.
//
// @param request - GetVideoInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoInfosResponse
func (client *Client) GetVideoInfosWithContext(ctx context.Context, request *GetVideoInfosRequest, runtime *dara.RuntimeOptions) (_result *GetVideoInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReferenceIds) {
		query["ReferenceIds"] = request.ReferenceIds
	}

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
// Retrieves a list of audio and video information.
//
// Description:
//
// This operation retrieves up to **5000*	- audio and video files that match the specified filter conditions (such as video status and category ID). Specify StartTime and EndTime to retrieve data in batches. To query more audio and video files or traverse all audio and video information, see [Search for media information](https://help.aliyun.com/document_detail/86044.html).
//
// @param request - GetVideoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoListResponse
func (client *Client) GetVideoListWithContext(ctx context.Context, request *GetVideoListRequest, runtime *dara.RuntimeOptions) (_result *GetVideoListResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReferenceIds) {
		query["ReferenceIds"] = request.ReferenceIds
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
// Retrieves the playback credential (PlayAuth) for an audio or video file. ApsaraVideo Player SDK uses this credential to automatically obtain the playback URL. Because the playback credential has a validity period and is bound to a specific audio or video file, it cannot be shared or reused. An expired or invalid credential will cause playback failure. This playback method is suitable for audio and video playback scenarios that require high security.
//
// Description:
//
// - When using ApsaraVideo Player SDK (applicable to the PlayAuth playback method), call this operation to obtain the playback credential. ApsaraVideo Player SDK uses the playback credential to automatically obtain the playback URL for playback. For more information, see [ApsaraVideo Player SDK](https://help.aliyun.com/document_detail/125579.html).
//
// - If the playback credential expires, the playback URL cannot be obtained. You must obtain a new playback credential.
//
// @param request - GetVideoPlayAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoPlayAuthResponse
func (client *Client) GetVideoPlayAuthWithContext(ctx context.Context, request *GetVideoPlayAuthRequest, runtime *dara.RuntimeOptions) (_result *GetVideoPlayAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiVersion) {
		query["ApiVersion"] = request.ApiVersion
	}

	if !dara.IsNil(request.AuthInfoTimeout) {
		query["AuthInfoTimeout"] = request.AuthInfoTimeout
	}

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to query the parameter settings of a single image watermark template or text watermark template by watermark template ID, including the position, size, and display time of image watermarks, and the content, font, color, and position of text watermarks.
//
// @param request - GetWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWatermarkResponse
func (client *Client) GetWatermarkWithContext(ctx context.Context, request *GetWatermarkRequest, runtime *dara.RuntimeOptions) (_result *GetWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the execution information of a workflow task.
//
// Description:
//
// - **Before you use this operation, make sure that you are familiar with the billing methods and pricing of ApsaraVideo VOD. Using workflows may incur fees for transcoding, encryption, automated review, and other services. For billing details, see [Billing overview](https://help.aliyun.com/document_detail/188307.html).**
//
// - You can call this operation to query workflow processing tasks. This operation currently supports only video understanding workflow task queries. Workflow tasks of other versions are not supported.
//
// @param request - GetWorkflowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowTaskResponse
func (client *Client) GetWorkflowTaskWithContext(ctx context.Context, request *GetWorkflowTaskRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkflowTask"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AI image information of a specified video.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)*	- and **China (Shanghai)**.
//
// - This operation can query AI image information of only a single video. **Batch queries are not supported**.
//
// @param request - ListAIImageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIImageInfoResponse
func (client *Client) ListAIImageInfoWithContext(ctx context.Context, request *ListAIImageInfoRequest, runtime *dara.RuntimeOptions) (_result *ListAIImageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries intelligent tagging or media fingerprint jobs. After you submit an intelligent tagging or media fingerprint job, the job is processed asynchronously. You can call this operation to query job information in real time.
//
// Description:
//
// - Regions that support media fingerprint: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - Regions that support intelligent tagging: **China (Beijing)*	- and **China (Shanghai)**.
//
// @param request - ListAIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIJobResponse
func (client *Client) ListAIJobWithContext(ctx context.Context, request *ListAIJobRequest, runtime *dara.RuntimeOptions) (_result *ListAIJobResponse, _err error) {
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
// Queries the list of AI templates.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - You can call this operation to query the list of AI templates of a specified type.
//
// @param request - ListAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAITemplateResponse
func (client *Client) ListAITemplateWithContext(ctx context.Context, request *ListAITemplateRequest, runtime *dara.RuntimeOptions) (_result *ListAITemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries a list of applications that you are authorized to access based on specified filter conditions.
//
// Description:
//
// ### Usage notes
//
// You can filter applications by application status.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation properly. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - ListAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInfoResponse
func (client *Client) ListAppInfoWithContext(ctx context.Context, request *ListAppInfoRequest, runtime *dara.RuntimeOptions) (_result *ListAppInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to list the application permissions granted to a specified account identity (Resource Access Management (RAM) user or RAM role).
//
// Description:
//
// - The **IdentityType*	- and **IdentityName*	- parameters take effect only when the caller invokes this operation with administrator permissions. Otherwise, only the application access policies granted to the current account identity are returned.
//
// @param request - ListAppPoliciesForIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppPoliciesForIdentityResponse
func (client *Client) ListAppPoliciesForIdentityWithContext(ctx context.Context, request *ListAppPoliciesForIdentityRequest, runtime *dara.RuntimeOptions) (_result *ListAppPoliciesForIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the list of China Chinese review security IPs.
//
// @param request - ListAuditSecurityIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuditSecurityIpResponse
func (client *Client) ListAuditSecurityIpWithContext(ctx context.Context, request *ListAuditSecurityIpRequest, runtime *dara.RuntimeOptions) (_result *ListAuditSecurityIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of animated images for a video by video ID.
//
// Description:
//
// ### Usage notes
//
// - After animated image capturing for a video is complete, call this operation to obtain the animated image information of the video.
//
// - Animated image tasks can be initiated by calling an API operation ([SubmitDynamicImageJob](https://help.aliyun.com/document_detail/186842.html)) or by using the console. For more information, see [Animated images](https://help.aliyun.com/document_detail/177484.html).
//
// ### QPS limit
//
// The QPS limit for a single user for this operation is 100 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limits](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - ListDynamicImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicImageResponse
func (client *Client) ListDynamicImageWithContext(ctx context.Context, request *ListDynamicImageRequest, runtime *dara.RuntimeOptions) (_result *ListDynamicImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the historical task list based on a media asset ID.
//
// Description:
//
// *Usage notes**
//
// - To query detailed task information, call the [GetJobDetail](https://help.aliyun.com/document_detail/2861326.html) operation.
//
// - This operation only supports querying asynchronous task data within the last 6 months. Supported task types: transcoding tasks, snapshot tasks, and AI tasks.
//
// **QPS limit**
//
// The single-user QPS limit for this operation is 15 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - ListJobInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobInfoResponse
func (client *Client) ListJobInfoWithContext(ctx context.Context, request *ListJobInfoRequest, runtime *dara.RuntimeOptions) (_result *ListJobInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the list of Live to VOD videos.
//
// Description:
//
// A maximum of 5,000 records that match the specified filter conditions can be retrieved.
//
// @param request - ListLiveRecordVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRecordVideoResponse
func (client *Client) ListLiveRecordVideoWithContext(ctx context.Context, request *ListLiveRecordVideoRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRecordVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries snapshots generated by video snapshot jobs and thumbnail snapshots automatically generated by the system during video upload.
//
// Description:
//
// If multiple snapshot jobs have been initiated for a video, this operation returns only the data of the most recent successful snapshot job.
//
// @param request - ListSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshotsWithContext(ctx context.Context, request *ListSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the historical transcoding task information of an audio or video file by its ID. This operation does not return specific job details.
//
// Description:
//
// ### Usage notes
//
// - To query detailed transcoding job information, call the [GetTranscodeTask](https://help.aliyun.com/document_detail/109121.html) operation.
//
// - **This operation supports only querying transcoding task data within the last year.**
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 15 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - ListTranscodeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranscodeTaskResponse
func (client *Client) ListTranscodeTaskWithContext(ctx context.Context, request *ListTranscodeTaskRequest, runtime *dara.RuntimeOptions) (_result *ListTranscodeTaskResponse, _err error) {
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
// Queries the list of transcoding template configurations.
//
// Description:
//
// > This operation does not return the transcoding template configuration information under each transcoding template group. You can call the [GetTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102670.html) operation to obtain the information.
//
// @param request - ListTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranscodeTemplateGroupResponse
func (client *Client) ListTranscodeTemplateGroupWithContext(ctx context.Context, request *ListTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *ListTranscodeTemplateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of snapshot templates.
//
// @param request - ListVodTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVodTemplateResponse
func (client *Client) ListVodTemplateWithContext(ctx context.Context, request *ListVodTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListVodTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to query the parameter settings of all image watermark templates and text watermark templates that have been added in the current service region, including the position, size, and display time of image watermarks, and the content, font, color, position, and other parameter settings of text watermarks.
//
// @param request - ListWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWatermarkResponse
func (client *Client) ListWatermarkWithContext(ctx context.Context, request *ListWatermarkRequest, runtime *dara.RuntimeOptions) (_result *ListWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invokes this operation to migrate resources such as media assets from one application to another. Application administrators can directly transfer resources. Resource Access Management (RAM) users or RAM roles must have write permissions on both the source and destination applications. Batch migration is supported.
//
// @param request - MoveAppResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveAppResourceResponse
func (client *Client) MoveAppResourceWithContext(ctx context.Context, request *MoveAppResourceRequest, runtime *dara.RuntimeOptions) (_result *MoveAppResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Prefetches content from the origin server to L2 Cache nodes so that the first access directly hits the cache, reducing the load on the origin server.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - A maximum of 500 prefetch URL requests can be submitted per account per day. Directory-level prefetch is not supported.
//
// - The purge and prefetch operations include the [RefreshVodObjectCaches](~~RefreshVodObjectCaches~~) purge operation and the [PreloadVodObjectCaches](~~PreloadVodObjectCaches~~) prefetch operation.
//
// @param request - PreloadVodObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadVodObjectCachesResponse
func (client *Client) PreloadVodObjectCachesWithContext(ctx context.Context, request *PreloadVodObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *PreloadVodObjectCachesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Produces one or more videos into a finished video. You can submit source videos directly through the timeline parameter, or create an online editing project first and then submit it for production.
//
// Description:
//
// - **Before you use this operation, make sure that you are familiar with the billing methods and pricing of ApsaraVideo VOD. Online editing is a paid feature. For more information about billing, see [Video editing and production billing](~~188310#section-pyv-b8h-bo7~~).**
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, the online editing project ID is returned (the video has not been produced yet, and the task enters a queue for asynchronous execution). The final result is sent through a callback notification. You can also call [GetEditingProject](https://help.aliyun.com/document_detail/69052.html) to query the task status.
//
// - The video resources used in the online editing timeline can be materials in the material library or videos in the media library. If you use videos from the media library, make sure that their status is Normal.
//
// - Videos are produced based on ProjectId and Timeline. The logic is as follows:
//
//   - ProjectId and Timeline cannot both be empty. Otherwise, no basis exists to produce videos.
//
//   - If ProjectId is empty and Timeline is not empty, an online editing project is automatically created with the specified Timeline. The materials referenced in the Timeline are extracted and set as the project materials. Then, video production begins.
//
//   - If ProjectId is not empty and Timeline is empty, the most recently saved Timeline is retrieved based on ProjectId and used to produce videos.
//
//   - If both ProjectId and Timeline are not empty, the specified Timeline is used to produce videos, and the corresponding online editing project is updated (Timeline and project materials). If other fields are specified, the corresponding project fields are also updated.
//
// - The maximum number of tracks for video tracks, image tracks, and subtitle tracks is 100 each.
//
// - The total number of materials cannot exceed 200, and the total file size of materials cannot exceed 1 TB.
//
// - The region of the input or output bucket must be the same as the region where the ApsaraVideo VOD service is used.
//
// - When the output is a video, the following resolution limits apply to the finished video:
//
//   - Both the width and height must be at least 128 px.
//
//   - Both the width and height must be at most 4096 px.
//
//   - The short side must be at most 2160 px.
//
// - After video production is complete, the video is automatically uploaded to ApsaraVideo VOD. Therefore, after video production is complete, ApsaraVideo VOD sends the **ProduceMediaComplete*	- and **FileUploadComplete*	- event notifications. After the produced video transcoding is complete, the **single definition video transcoding complete*	- and **all definition video transcoding complete*	- event notifications are sent.
//
// - You can also add effects to the produced video. For more details, see [Effects](https://help.aliyun.com/document_detail/69082.html).
//
// @param request - ProduceEditingProjectVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProduceEditingProjectVideoResponse
func (client *Client) ProduceEditingProjectVideoWithContext(ctx context.Context, request *ProduceEditingProjectVideoRequest, runtime *dara.RuntimeOptions) (_result *ProduceEditingProjectVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Submits a refresh or prefetch task for audio or video files by audio or video ID.
//
// Description:
//
// - ApsaraVideo VOD provides resource purge and prefetch features. The purge feature deletes cached resources on points of presence and forces the points of presence to retrieve the latest resources from the origin server through back-to-origin requests. The prefetch feature allows you to download and cache popular resources to points of presence before peak hours to improve access efficiency.
//
// - This operation directly submits a refresh or prefetch node by audio or video ID and supports filtering by streaming format and definition, which allows you to refresh or prefetch specific streams as needed.
//
// - You can submit a refresh or prefetch node for up to 20 audio or video files at a time.
//
// ### QPS limit
//
// The QPS limit for a single user for this operation is 50 calls per second. If the limit is exceeded, the API invocation is throttled, which may affect your business. Invoke this operation properly. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - RefreshMediaPlayUrlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshMediaPlayUrlsResponse
func (client *Client) RefreshMediaPlayUrlsWithContext(ctx context.Context, request *RefreshMediaPlayUrlsRequest, runtime *dara.RuntimeOptions) (_result *RefreshMediaPlayUrlsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Refreshes the upload credential for a video file after the upload times out.
//
// Description:
//
// This operation can also be used to overwrite the source file of a video or audio file. This means that after you obtain the upload URL of the source file, you can upload the file again while keeping the audio or video ID unchanged. However, this may automatically trigger transcoding and snapshot capture if you have configured transcoding or snapshot capture upon upload. For more information, see [Upload URLs and credentials](https://help.aliyun.com/document_detail/55397.html).
//
// @param request - RefreshUploadVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshUploadVideoResponse
func (client *Client) RefreshUploadVideoWithContext(ctx context.Context, request *RefreshUploadVideoRequest, runtime *dara.RuntimeOptions) (_result *RefreshUploadVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Purges file content on nodes. Specifies URL content to purge on cache nodes, and supports batch URL purging.
//
// Description:
//
// - Currently, the only supported service region is **China (Shanghai)**.
//
// - Each account can submit up to 2,000 URL purge requests and 100 directory purge requests per day.
//
// - Purge and prefetch operations include the [RefreshVodObjectCaches](https://help.aliyun.com/document_detail/69215.html) purge operation and the [PreloadVodObjectCaches](https://help.aliyun.com/document_detail/69211.html) prefetch operation.
//
// @param request - RefreshVodObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshVodObjectCachesResponse
func (client *Client) RefreshVodObjectCachesWithContext(ctx context.Context, request *RefreshVodObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *RefreshVodObjectCachesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Registers media assets. Existing media files stored in your own OSS bucket that is connected to ApsaraVideo VOD must be registered to generate the associated data required by VOD before you can use VOD features such as transcoding and snapshotting.
//
// Description:
//
// - For audio and video files already stored in an OSS bucket connected to ApsaraVideo VOD, you must call this operation to generate the associated data required by VOD before you can initiate transcoding, snapshotting, AI processing, and other operations on these files by media ID.
//
// - You can register up to **10 OSS media files*	- at a time, and all media files submitted in a single request must correspond to the same storage address.
//
// - For media files uploaded through VOD, if no transcoding template group ID is specified, the default template group is used for transcoding. In contrast, after media asset registration, transcoding is not automatically triggered if no transcoding template group ID is specified. If a transcoding template group ID is specified, transcoding is performed based on the specified template group.
//
// - If a media file is registered repeatedly, only the **unique media ID associated with it*	- is returned, and no other processing is performed.
//
// - Make sure that the media file you want to register has a valid file name extension. Otherwise, the registration fails.
//
// @param request - RegisterMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterMediaResponse
func (client *Client) RegisterMediaWithContext(ctx context.Context, request *RegisterMediaRequest, runtime *dara.RuntimeOptions) (_result *RegisterMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableFirstFrameCover) {
		query["EnableFirstFrameCover"] = request.EnableFirstFrameCover
	}

	if !dara.IsNil(request.GenerateThumbnail) {
		query["GenerateThumbnail"] = request.GenerateThumbnail
	}

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
// Restores media assets from frozen storage.
//
// Description:
//
// - Make sure that you are familiar with the billing methods and pricing of ApsaraVideo VOD before you call this operation. Restoring media assets incurs storage fees. For more information, see [Media asset storage billing](~~188308#section-e97-xrp-mzz~~).
//
// - This operation applies only to Archive and Cold Archive audio and video files. After a file is restored, it can be accessed. The storage class of an audio or video file that is being restored cannot be changed.
//
// Restoration generates retrieval traffic. After a Cold Archive audio or video file is restored, a Standard storage copy of the file is generated for access. The file copy incurs Standard storage fees until the restoration period ends.
//
// @param request - RestoreMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreMediaResponse
func (client *Client) RestoreMediaWithContext(ctx context.Context, request *RestoreMediaRequest, runtime *dara.RuntimeOptions) (_result *RestoreMediaResponse, _err error) {
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
// Searches for online editing projects (video editing lists).
//
// @param request - SearchEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchEditingProjectResponse
func (client *Client) SearchEditingProjectWithContext(ctx context.Context, request *SearchEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *SearchEditingProjectResponse, _err error) {
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
// Searches for media asset information such as videos, audio files, and images produced by ApsaraVideo VOD. You can use this operation with the media asset search protocol to perform multi-dimensional searches in ApsaraVideo VOD, including specifying return fields, exact matching, fuzzy matching, multi-value queries, range queries, and sort fields.
//
// Description:
//
// For fields that support exact matching and fuzzy matching, when other query methods are used, the returned results follow the query method supported by the field. For example, if a field supports only fuzzy matching, results obtained through multi-value queries are also based on fuzzy matching.
//
// The following describes the limits on the number of data records that can be retrieved:
//
// - Method 1: Paged traversal
//
//	For matched search results, you can set the pagination parameters PageNo (page number) and PageSize (number of records per page) to traverse up to 5,000 records. If the search results exceed 5,000 records, adjust the search conditions to narrow the result range. This method cannot traverse the complete dataset. To traverse more data, refer to Method 2.
//
// - Method 2: Full traversal (for audio and video searches only)
//
//	This method applies to video and audio content searches and supports traversing up to 2 million search results. If the number of search results exceeds 2 million, add more filter conditions to reduce the result count. When using this method, in addition to PageNo and PageSize, you must use the ScrollToken parameter for pagination. Each request supports traversing up to 100 records forward.
//
// Using a PageSize of 20 as an example, the pagination logic is as follows:
//
//   - If PageNo is 1, you can query up to the next 5 pages of data.
//
//   - If PageNo is 2, you can query up to the next 6 pages of data.
//
// Set pagination parameters properly and choose the appropriate traversal method based on the result set size. If you need to page through more than 1,000 records, use Method 2 for faster and more convenient data processing.
//
// @param request - SearchMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaResponse
func (client *Client) SearchMediaWithContext(ctx context.Context, request *SearchMediaRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Sets the playback key for an application.
//
// @param request - SetAppPlayKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAppPlayKeyResponse
func (client *Client) SetAppPlayKeyWithContext(ctx context.Context, request *SetAppPlayKeyRequest, runtime *dara.RuntimeOptions) (_result *SetAppPlayKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayKey) {
		query["PlayKey"] = request.PlayKey
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
		Action:      dara.String("SetAppPlayKey"),
		Version:     dara.String("2017-03-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAppPlayKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the review security IP addresses.
//
// Description:
//
// When a video is in the Checking or Blocked state, only requests from review security IP addresses can play the video.
//
// @param request - SetAuditSecurityIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAuditSecurityIpResponse
func (client *Client) SetAuditSecurityIpWithContext(ctx context.Context, request *SetAuditSecurityIpRequest, runtime *dara.RuntimeOptions) (_result *SetAuditSecurityIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates the content of the cross-domain file crossdomain.xml for ApsaraVideo VOD.
//
// Description:
//
//	Notice: If you access the cross-domain file through a domain name, purge the CDN cache for the update to take effect immediately. You can logon to the console to [purge files](https://help.aliyun.com/document_detail/86098.html) or invoke the [Refresh Cache operation](https://help.aliyun.com/document_detail/69215.html).
//
// @param request - SetCrossdomainContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCrossdomainContentResponse
func (client *Client) SetCrossdomainContentWithContext(ctx context.Context, request *SetCrossdomainContentRequest, runtime *dara.RuntimeOptions) (_result *SetCrossdomainContentResponse, _err error) {
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
// Sets a default AI template.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - Obtain the AI template ID first, and then call this operation to set the template as the default AI template. A default AI template cannot be deleted.
//
// @param request - SetDefaultAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultAITemplateResponse
func (client *Client) SetDefaultAITemplateWithContext(ctx context.Context, request *SetDefaultAITemplateRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultAITemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Sets the default transcoding template group configuration.
//
// @param request - SetDefaultTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultTranscodeTemplateGroupResponse
func (client *Client) SetDefaultTranscodeTemplateGroupWithContext(ctx context.Context, request *SetDefaultTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultTranscodeTemplateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Sets a specified watermark template as the default watermark template.
//
// @param request - SetDefaultWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultWatermarkResponse
func (client *Client) SetDefaultWatermarkWithContext(ctx context.Context, request *SetDefaultWatermarkRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultWatermarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Sets the materials to be edited for an online editing project.
//
// @param request - SetEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetEditingProjectMaterialsResponse
func (client *Client) SetEditingProjectMaterialsWithContext(ctx context.Context, request *SetEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *SetEditingProjectMaterialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Sets the callback method, callback URL, and event types for event notifications.
//
// Description:
//
// HTTP callbacks and Simple Message Queue (formerly MNS) callbacks are supported. For more information, see [Event notifications](https://help.aliyun.com/document_detail/55627.html).
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 15 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - SetMessageCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMessageCallbackResponse
func (client *Client) SetMessageCallbackWithContext(ctx context.Context, request *SetMessageCallbackRequest, runtime *dara.RuntimeOptions) (_result *SetMessageCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Configures whether the certificate feature is enabled for a specified domain name and modifies certificate information.
//
// Description:
//
// - Currently, the service address supports only **China (Shanghai)**.
//
// @param request - SetVodDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVodDomainCertificateResponse
func (client *Client) SetVodDomainCertificateWithContext(ctx context.Context, request *SetVodDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetVodDomainCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Sets whether the certificate feature is enabled for a specified domain name and updates the certificate information.
//
// Description:
//
// - This operation currently supports only the **China East 2 (Shanghai)*	- region.
//
// - Maximum calls per user: 30 calls per second.
//
// - Request method: POST.
//
// @param request - SetVodDomainSSLCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVodDomainSSLCertificateResponse
func (client *Client) SetVodDomainSSLCertificateWithContext(ctx context.Context, request *SetVodDomainSSLCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetVodDomainSSLCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Submits an automated review task for an image. The task is asynchronously executed after it is submitted. The task may not be complete when the response is returned.
//
// Description:
//
// - **Before you use this operation, make sure that you are familiar with the billing methods and pricing of ApsaraVideo VOD. Automated review is a paid feature. For billing details, <props="china">see [Automated review billing](~~188310#section-g7l-s3o-9ng~~).<props="intl">submit a ticket or contact your Alibaba Cloud account manager.**
//
// - <props="china">This operation is supported only in the **China (Shanghai), China (Beijing), and Singapore*	- regions.<props="intl">This operation is supported only in the Singapore region.
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, the task ID is returned. At this point, the task is not complete and enters a queue for asynchronous execution. The final result is sent through a callback notification. You can also call [Query automated review job](https://help.aliyun.com/document_detail/454959.html) to query the task status.
//
// - The size of a single image cannot exceed 20 MB. The height or width cannot exceed 30,000 px. The total number of pixels cannot exceed 250 million px.
//
// - (Recommended) The image resolution is at least 256 × 256 px. A lower resolution may affect the review results.
//
// @param request - SubmitAIImageAuditJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIImageAuditJobResponse
func (client *Client) SubmitAIImageAuditJobWithContext(ctx context.Context, request *SubmitAIImageAuditJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIImageAuditJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CensorProvider) {
		query["CensorProvider"] = request.CensorProvider
	}

	if !dara.IsNil(request.ImageService) {
		query["ImageService"] = request.ImageService
	}

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

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
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
// Submits an AI image processing task.
//
// Description:
//
// - Currently, this operation is supported in the following regions: **China (Beijing)*	- and **China (Shanghai)**.
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, the task ID is returned. The task is not yet complete at this point and enters a background queue for asynchronous execution. The final result is sent through a callback notification. You can also call [GetAIImageJobs](https://help.aliyun.com/document_detail/186923.html) to query the task execution result.
//
// @param request - SubmitAIImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIImageJobResponse
func (client *Client) SubmitAIImageJobWithContext(ctx context.Context, request *SubmitAIImageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIImageJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Submits an intelligent tagging or media fingerprint job.
//
// Description:
//
// - **Before using this operation, make sure that you fully understand the billing methods and pricing of ApsaraVideo VOD. Intelligent tagging and media fingerprint are paid features. For billing details, see [Video AI billing](~~188310#section-g7l-s3o-9ng~~).**
//
// - Regions supported by media fingerprint: **China (Beijing)**, **China (Shanghai)**, and **Singapore**. Regions supported by intelligent tagging: **China (Beijing)*	- and **China (Shanghai)**.
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit an AI job, the job ID is returned. The job is not yet complete at this point and enters a queue for asynchronous execution. We recommend that you configure the [event notification](https://help.aliyun.com/document_detail/55627.html) feature and set the callback event to **Video AI Processing Complete*	- to obtain the final processing result. You can also call [GetTaskDetail](https://help.aliyun.com/document_detail/2861326.html) to query the job status.
//
// - You must activate the media fingerprint or intelligent tagging service before you can call this operation to submit AI jobs. For more information, see [Video AI](https://help.aliyun.com/document_detail/101148.html).
//
// - When you use media fingerprint for the first time, provide your UID and region information and submit a ticket to apply for free activation of the fingerprint library. Otherwise, the media fingerprint feature will not work properly. For more information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
//
// @param request - SubmitAIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIJobResponse
func (client *Client) SubmitAIJobWithContext(ctx context.Context, request *SubmitAIJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIJobResponse, _err error) {
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
// Submits an automated review job for audio and video files. The job is executed asynchronously after submission, and the job may not be complete when the response is returned.
//
// Description:
//
// - **Before using this operation, make sure that you are familiar with the billing methods and pricing of ApsaraVideo VOD. Automated review is a paid feature. For billing details, <props="china">refer to [Automated review billing](~~188310#section-g7l-s3o-9ng~~).<props="intl">submit a ticket or contact your Alibaba Cloud account manager.**
//
// - This operation currently supports only the **Shanghai**, **Beijing**, and **Singapore*	- regions.
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, a task ID is returned. The task is not yet complete at this point and enters a queue for asynchronous execution. The final result is sent through a callback notification. You can also call [Query automated review job](https://help.aliyun.com/document_detail/454959.html) to query the task status.
//
// - For the development guide on submitting automated review jobs, refer to [Automated review](https://help.aliyun.com/document_detail/101148.html).
//
// - After an automated review job is complete, the image resources generated during the job are retained for free for only two weeks in the VOD system bucket allocated by ApsaraVideo VOD. The images are automatically deleted after two weeks.
//
// @param request - SubmitAIMediaAuditJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIMediaAuditJobResponse
func (client *Client) SubmitAIMediaAuditJobWithContext(ctx context.Context, request *SubmitAIMediaAuditJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIMediaAuditJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CensorProvider) {
		query["CensorProvider"] = request.CensorProvider
	}

	if !dara.IsNil(request.MediaAuditConfiguration) {
		query["MediaAuditConfiguration"] = request.MediaAuditConfiguration
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VideoService) {
		query["VideoService"] = request.VideoService
	}

	if !dara.IsNil(request.VoiceService) {
		query["VoiceService"] = request.VoiceService
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
// Submits a digital watermarking (copyright watermark or tracing watermark) extraction job to asynchronously extract a copyright watermark or tracing watermark.
//
// Description:
//
// - **Before using this operation, make sure that you fully understand the billing methods and pricing of ApsaraVideo VOD. Digital watermarking is a paid feature. Both the generation and extraction of digital watermarks incur fees. For billing details, see [Digital watermarking billing](~~188310#62b9c940403se~~).**
//
// - Currently, this operation is available only in the following regions: **China (Shanghai)*	- and **China (Beijing)**.
//
// - <props="china">For more information about the generation and extraction of digital watermarks, see [Digital watermarking](https://help.aliyun.com/document_detail/2527021.html).Before you submit a digital watermark extraction job, make sure that the following conditions are met:
//
//   - The video from which you want to extract the watermark has been uploaded to ApsaraVideo VOD.
//
//   - The duration of the video from which you want to extract the watermark exceeds 6 minutes.
//
// - After you submit a digital watermark extraction job, call the [GetDigitalWatermarkExtractResult](https://help.aliyun.com/document_detail/2587769.html) operation to query the job result.
//
// @param request - SubmitDigitalWatermarkExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDigitalWatermarkExtractJobResponse
func (client *Client) SubmitDigitalWatermarkExtractJobWithContext(ctx context.Context, request *SubmitDigitalWatermarkExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDigitalWatermarkExtractJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Submits a media animated image job for asynchronous processing.
//
// Description:
//
// ### Usage notes
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, the task ID is returned. The task is queued for asynchronous execution in the background. You can receive the final result through a callback notification or proactively query the task status by calling [Get task details](https://help.aliyun.com/document_detail/2861326.html).
//
// - You can submit an animated image job only for videos in the **UploadSucc**, **Transcoding**, **Normal**, **Checking**, or **Blocked*	- state.
//
// - Animated image production is billed as video transcoding at the same rate, based on resolution and duration. For more information, see [Media transcoding billing](https://help.aliyun.com/document_detail/188308.html).
//
// ### QPS limit
//
// The QPS limit for a single user on this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - SubmitDynamicImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDynamicImageJobResponse
func (client *Client) SubmitDynamicImageJobWithContext(ctx context.Context, request *SubmitDynamicImageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDynamicImageJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a media fingerprint job.
//
// Description:
//
// Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// @param request - SubmitMediaDNADeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaDNADeleteJobResponse
func (client *Client) SubmitMediaDNADeleteJobWithContext(ctx context.Context, request *SubmitMediaDNADeleteJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaDNADeleteJobResponse, _err error) {
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
// Transcodes a video by using a China Production Studio (China Production Studio) for preprocessing.
//
// Description:
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, the task ID is returned. The task is queued for asynchronous execution in the background. The final result is sent through a callback notification. You can also call [GetTaskDetail](https://help.aliyun.com/document_detail/2861326.html) to query the task status.
//
// - Video preprocessing is essentially a transcoding process that generates videos that meet the playback requirements of the China Production Studio. Therefore, **metering and billing*	- information is generated. For billing details, see [China Production Studio fees](https://help.aliyun.com/document_detail/64531.html).
//
// - To meet the quality requirements of the China Production Studio for materials, videos in MP4 format with a short side of 360 must meet at least one of the following conditions before preprocessing can be initiated: the resolution exceeds 1920, the bit rate exceeds 6000 kbps, or the frame rate exceeds 25.
//
// - You can receive the [TranscodeComplete](https://help.aliyun.com/document_detail/55638.html) callback message to obtain the processing result. When the callback message contains **Preprocess=true**, it indicates that the preprocessing is complete.
//
// @param request - SubmitPreprocessJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPreprocessJobsResponse
func (client *Client) SubmitPreprocessJobsWithContext(ctx context.Context, request *SubmitPreprocessJobsRequest, runtime *dara.RuntimeOptions) (_result *SubmitPreprocessJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Submits a video snapshot job to start asynchronous snapshot processing.
//
// Description:
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, the task ID is returned (the task is not yet complete and enters a queue for asynchronous execution). The final result is sent through a callback notification. You can also proactively query the task status by calling [Get task details](https://help.aliyun.com/document_detail/2861326.html).
//
// - Only JPG images are supported.
//
// - When the snapshot is complete, an event notification of [Video snapshot complete](https://help.aliyun.com/document_detail/57337.html) with EventType=SnapshotComplete and SubType=SpecifiedTime is sent.
//
// ### QPS limit
//
// The QPS limit for a single user on this operation is 30 calls per second. If this limit is exceeded, the API call is throttled, which may affect your business. Call this operation appropriately. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param tmpReq - SubmitSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSnapshotJobResponse
func (client *Client) SubmitSnapshotJobWithContext(ctx context.Context, tmpReq *SubmitSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSnapshotJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Submits a media transcoding job to start asynchronous transcoding.
//
// Description:
//
// ### Usage notes
//
// - **Before you use this operation, make sure that you fully understand the billing methods and pricing of ApsaraVideo VOD. Transcoding is a paid feature. For more information about billing, see [Transcoding billing](~~188308#section-ejb-nii-nqa~~).**
//
// - This is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, the task ID is returned. The task is not yet complete at this point and enters a queue for asynchronous execution. The final result is sent through a callback notification. You can also call [GetTranscodeTask](https://help.aliyun.com/document_detail/454946.html) to query the task status.
//
// - Only videos in the **UploadSucc**, **Normal**, or **Checking*	- state can be transcoded.
//
// - To obtain transcoding results, configure callback messages: [SingleCompleteEvent](https://help.aliyun.com/document_detail/55636.html) and [AllCompleteEvent](https://help.aliyun.com/document_detail/55638.html).
//
// - This operation supports dynamic replacement of subtitle URLs in HLS adaptive bitrate streaming packaging tasks. If the packaging task does not involve subtitle packaging, do not use this operation to initiate the task. Instead, specify the corresponding transcoding template group ID during video upload to automatically trigger the packaging process.
//
// @param request - SubmitTranscodeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTranscodeJobsResponse
func (client *Client) SubmitTranscodeJobsWithContext(ctx context.Context, request *SubmitTranscodeJobsRequest, runtime *dara.RuntimeOptions) (_result *SubmitTranscodeJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Initiates a VOD workflow for a video.
//
// Description:
//
// - **Before using this operation, make sure that you fully understand the billing methods and pricing of ApsaraVideo VOD. Using workflows may incur fees for transcoding, encryption, automated review, and other services. For more information, see [Billing overview](https://help.aliyun.com/document_detail/188307.html).**
//
// - This operation is an [asynchronous operation](https://help.aliyun.com/document_detail/3027551.html). After you submit a task, a task ID is returned (the task is not yet complete and enters a background queue for asynchronous execution). The final result is sent through a callback notification. You can also call [GetTask](https://help.aliyun.com/document_detail/2861326.html) to query the task status.
//
// - Call this operation to initiate a workflow processing task for a video. For more information about workflows, see [Workflow](https://help.aliyun.com/document_detail/115347.html).
//
// @param request - SubmitWorkflowJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitWorkflowJobResponse
func (client *Client) SubmitWorkflowJobWithContext(ctx context.Context, request *SubmitWorkflowJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitWorkflowJobResponse, _err error) {
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
// - Currently, this operation is supported in the following regions: **China (Beijing)**, **China (Shanghai)**, and **Singapore**.
//
// - After you call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template, you can call this operation to modify the AI template.
//
// @param request - UpdateAITemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAITemplateResponse
func (client *Client) UpdateAITemplateWithContext(ctx context.Context, request *UpdateAITemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateAITemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates application information.
//
// Description:
//
// After an application is created, you can call this operation to locate an application by its application ID and modify the name, description, and status of the application.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation properly. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - UpdateAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppInfoResponse
func (client *Client) UpdateAppInfoWithContext(ctx context.Context, request *UpdateAppInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Batch updates the information of auxiliary media assets, such as title, description, tags, and category, by specifying the unique identifiers (IDs) of the auxiliary media assets that have been uploaded to ApsaraVideo VOD, including watermarked images, subtitle files, and other materials.
//
// Description:
//
// You can update the information of up to 20 auxiliary media assets at a time.
//
// @param request - UpdateAttachedMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAttachedMediaInfosResponse
func (client *Client) UpdateAttachedMediaInfosWithContext(ctx context.Context, request *UpdateAttachedMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *UpdateAttachedMediaInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates the name of a category.
//
// Description:
//
// After a category is created, you can call this operation to modify the name of the category. If the category has been annotated to some media assets, the category name annotated to those media assets is updated synchronously after the category name is modified.
//
// @param request - UpdateCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCategoryResponse
func (client *Client) UpdateCategoryWithContext(ctx context.Context, request *UpdateCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
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
// Modifies an online editing project (video editing task).
//
// @param request - UpdateEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEditingProjectResponse
func (client *Client) UpdateEditingProjectWithContext(ctx context.Context, request *UpdateEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateEditingProjectResponse, _err error) {
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
// Batch modifies the title, description, tags, and category information of images by image ID after the images are uploaded to ApsaraVideo VOD.
//
// Description:
//
// - This operation only supports modifying uploaded images. Modifying images generated from video snapshots is not supported.
//
// - You can modify the information of up to 20 images at a time.
//
// @param request - UpdateImageInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageInfosResponse
func (client *Client) UpdateImageInfosWithContext(ctx context.Context, request *UpdateImageInfosRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies the storage class of a media asset.
//
// Description:
//
// - Before using this operation, make sure that you fully understand the billing methods and pricing of ApsaraVideo VOD. Modifying the storage class of a media asset incurs storage fees. For billing details, see [Media asset storage billing](~~188308#section-e97-xrp-mzz~~).
//
// - Modifying the storage class is an **asynchronous operation**. A callback is sent to you after the entire operation is complete.
//
// - If the current storage class of a media asset is Archive or ColdArchive, calling this operation automatically triggers a restore. After the restore is complete, the storage class is modified. You do not need to manually call the RestoreMedia operation to restore the media asset. For ColdArchive media assets, you need to specify the restore priority. The default value is RestoreTier=Standard.
//
// - A media asset that is being modified cannot be modified again or be used for production or processing.
//
// - Media assets in non-Standard storage classes have minimum storage duration requirements: Infrequent Access/source file Infrequent Access requires at least 30 days. Archive/source file Archive requires at least 60 days. ColdArchive/source file ColdArchive requires at least 180 days. If the storage duration is insufficient, modifying the storage class incurs storage fees for the remaining days. For example, if you modify the storage class from Infrequent Access to Standard after 10 days of storage, you are charged for the remaining 20 days of Infrequent Access storage, totaling 30 days of Infrequent Access storage fees.
//
// - **Modifying the storage class of a self-managed bucket is not supported**.
//
// @param request - UpdateMediaStorageClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaStorageClassResponse
func (client *Client) UpdateMediaStorageClassWithContext(ctx context.Context, request *UpdateMediaStorageClassRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaStorageClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies transcoding configurations. You can modify the configuration of a specified transcoding template in a transcoding template group.
//
// Description:
//
// For security purposes, you cannot add, modify, or delete transcoding template groups that are in the locked state. You can call the [GetTranscodeTemplateGroup](~~GetTranscodeTemplateGroup~~) operation to query the template configuration and check whether the template group is locked based on the Locked response parameter. Alternatively, you can call this operation to unlock the template group before you add, modify, or delete templates.
//
// @param request - UpdateTranscodeTemplateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTranscodeTemplateGroupResponse
func (client *Client) UpdateTranscodeTemplateGroupWithContext(ctx context.Context, request *UpdateTranscodeTemplateGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateTranscodeTemplateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies the information about a single audio or video file.
//
// Description:
//
// ### Operation description
//
// This operation locates an audio or video file by video ID and supports modifying the title, tags, description, and other information of the file. If a parameter is specified, the corresponding field is updated. Otherwise, the field is not overwritten or updated.
//
// ### QPS limit
//
// A single user can perform a maximum of 100 queries per second (QPS). Throttling is triggered when the QPS limit is exceeded, which may affect your business. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - UpdateVideoInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoInfoResponse
func (client *Client) UpdateVideoInfoWithContext(ctx context.Context, request *UpdateVideoInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateVideoInfoResponse, _err error) {
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

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
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
// Modifies the information about multiple audio and video files at a time.
//
// Description:
//
// ### Usage notes
//
// - Audio and video files are identified by their IDs. You can modify the title, tags, and description of audio and video files. If a parameter is specified, the corresponding field is updated. Otherwise, the field is not overwritten or updated.
//
// - You can modify the information about up to 20 audio and video files at a time.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 30 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation properly. For more information, see [QPS limit](https://help.aliyun.com/document_detail/342790.html).
//
// @param request - UpdateVideoInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoInfosResponse
func (client *Client) UpdateVideoInfosWithContext(ctx context.Context, request *UpdateVideoInfosRequest, runtime *dara.RuntimeOptions) (_result *UpdateVideoInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies an accelerated domain name.
//
// Description:
//
// > This operation is available only in the **China (Shanghai)*	- region.
//
// @param request - UpdateVodDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVodDomainResponse
func (client *Client) UpdateVodDomainWithContext(ctx context.Context, request *UpdateVodDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateVodDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies the name and watermark configuration (WatermarkConfig) of an image and text watermark template that was added by calling the AddWatermark operation.
//
// Description:
//
// - After you invoke [AddWatermark](~~AddWatermark~~) to add an image and text watermark template, you can invoke this operation to modify the name and watermark configuration of the template.
//
// - This operation does not support replacing the image in an image watermark template or modifying the template across templatetypes (such as changing an image watermark template to a text watermark template).
//
// @param request - UpdateWatermarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWatermarkResponse
func (client *Client) UpdateWatermarkWithContext(ctx context.Context, request *UpdateWatermarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateWatermarkResponse, _err error) {
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
// Pulls audio and video media files for upload based on source file URLs. Batch upload is supported.
//
// Description:
//
// - **Before using this operation, make sure that you fully understand the billing methods and pricing of ApsaraVideo VOD. Uploading media files to ApsaraVideo VOD incurs storage fees. For billing details, see [Media asset storage billing](~~188308#section_e97_xrp_mzz~~). If you have enabled storage transfer acceleration, uploading media files to ApsaraVideo VOD also incurs upload acceleration fees. For billing details, see [Storage transfer acceleration billing](~~188310#section_sta_zm2_tsv~~).**
//
// - For the media file formats supported by this operation, see [Media formats](~~55396#section-e27-2rj-mde~~).
//
// - This operation is mainly applicable to scenarios where files are not stored on a local server or terminal and need to be uploaded through a URL with public network access.
//
// - This operation is an [asynchronous upload operation](https://help.aliyun.com/document_detail/3027551.html). It is not real-time and does not guarantee timeliness. Generally, the migration upload is completed within hours or even days after the node is submitted. If you have high timeliness requirements, use the upload SDK instead.
//
// - If a callback is configured, you will receive the [URL upload video complete](https://help.aliyun.com/document_detail/86326.html) event notification after the upload is completed. You can call the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation to query the upload status.
//
// - After an upload node is submitted, an asynchronous node is generated in the cloud for execute. All URL upload nodes committed by users in the corresponding service region are queued for execute. The completion time is affected by the number of existing nodes. After the upload is completed, you can associate the URL with the video ID based on the information returned in the event notification (message callback).
//
// - This operation currently supports only the **China (Shanghai)**, **China (Beijing)**, **China (Shenzhen)**, **Singapore**, and **US (Silicon Valley)*	- regions.
//
// - Each time you commit an upload node for the same media file URL, a new media resource is generated in ApsaraVideo VOD (that is, a new media ID is generated).
//
// - If a single file exceeds 20 GB, the upload is failed. If you need to upload a single file larger than 20 GB, use the upload SDK. For more information, see [Overview of the upload SDK](https://help.aliyun.com/document_detail/52200.html).
//
// @param request - UploadMediaByURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMediaByURLResponse
func (client *Client) UploadMediaByURLWithContext(ctx context.Context, request *UploadMediaByURLRequest, runtime *dara.RuntimeOptions) (_result *UploadMediaByURLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EnableFirstFrameCover) {
		query["EnableFirstFrameCover"] = request.EnableFirstFrameCover
	}

	if !dara.IsNil(request.GenerateThumbnail) {
		query["GenerateThumbnail"] = request.GenerateThumbnail
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
// Uploads a transcoded stream file from external storage and mounts it to the corresponding ApsaraVideo VOD media asset.
//
// Description:
//
// - **Before using this operation, make sure that you fully understand the billing methods and pricing of ApsaraVideo VOD. Uploading media files to ApsaraVideo VOD incurs storage fees. For more information, see [Media asset storage billing](~~188308#section_e97_xrp_mzz~~). If you have enabled storage transmission acceleration, upload acceleration fees also apply. For more information, see [Storage transmission acceleration billing](~~188310#section_sta_zm2_tsv~~).**
//
// - This operation is currently supported only in the **Shanghai*	- and **Singapore*	- regions.
//
// - Call this operation to upload a transcoded stream file from external storage and mount it to the corresponding ApsaraVideo VOD media asset. The supported HDR types for transcoded streams are HDR, HDR10, HLG, DolbyVision, HDRVivid, and SDR+.
//
// - You can call the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation to query the upload status. After the upload is complete, you will receive the [URL upload transcoded stream complete](https://help.aliyun.com/document_detail/376427.html) event notification.
//
// @param request - UploadStreamByURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadStreamByURLResponse
func (client *Client) UploadStreamByURLWithContext(ctx context.Context, request *UploadStreamByURLRequest, runtime *dara.RuntimeOptions) (_result *UploadStreamByURLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Calls this operation to verify domain name ownership.
//
// Description:
//
// Currently, the service is supported only in the **China (Shanghai)*	- region.
//
// @param request - VerifyVodDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyVodDomainOwnerResponse
func (client *Client) VerifyVodDomainOwnerWithContext(ctx context.Context, request *VerifyVodDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyVodDomainOwnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
