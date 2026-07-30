// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Retrieves information about multiple media assets in a batch.
//
// Description:
//
// ## Request description.
//
// @param request - BatchGetMediasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetMediasResponse
func (client *Client) BatchGetMediasWithContext(ctx context.Context, request *BatchGetMediasRequest, runtime *dara.RuntimeOptions) (_result *BatchGetMediasResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetMedias"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetMediasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a media asset category.
//
// Description:
//
// Categories support up to three levels, and each level supports up to 100 subcategories.
//
// @param request - CreateAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssetCategoryResponse
func (client *Client) CreateAssetCategoryWithContext(ctx context.Context, request *CreateAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *CreateAssetCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAssetCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a media asset category.
//
// Description:
//
// This operation also deletes all subcategories (including second-level and third-level categories). Proceed with caution.
//
// @param request - DeleteAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAssetCategoryResponse
func (client *Client) DeleteAssetCategoryWithContext(ctx context.Context, request *DeleteAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteAssetCategoryResponse, _err error) {
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
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAssetCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes media asset information.
//
// @param request - DeleteMediasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediasResponse
func (client *Client) DeleteMediasWithContext(ctx context.Context, request *DeleteMediasRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletePhysicalFiles) {
		query["DeletePhysicalFiles"] = request.DeletePhysicalFiles
	}

	if !dara.IsNil(request.InputURLs) {
		query["InputURLs"] = request.InputURLs
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMedias"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified category and the list of its subcategories (immediate child categories).
//
// @param request - GetAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssetCategoryResponse
func (client *Client) GetAssetCategoryWithContext(ctx context.Context, request *GetAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetAssetCategoryResponse, _err error) {
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
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAssetCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an image generation task.
//
// @param request - GetImageGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageGenerationJobResponse
func (client *Client) GetImageGenerationJobWithContext(ctx context.Context, request *GetImageGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *GetImageGenerationJobResponse, _err error) {
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
		Action:      dara.String("GetImageGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询媒资
//
// Description:
//
// ## 请求说明
//
// 该API用于查询媒资内容理解作业。
//
// @param request - GetMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaResponse
func (client *Client) GetMediaWithContext(ctx context.Context, request *GetMediaRequest, runtime *dara.RuntimeOptions) (_result *GetMediaResponse, _err error) {
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

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMedia"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a media asset content understanding job.
//
// Description:
//
// ## Description
//
// This API is used to query a media asset content understanding job.
//
// @param request - GetMediaComprehensionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaComprehensionJobResponse
func (client *Client) GetMediaComprehensionJobWithContext(ctx context.Context, request *GetMediaComprehensionJobRequest, runtime *dara.RuntimeOptions) (_result *GetMediaComprehensionJobResponse, _err error) {
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
		Action:      dara.String("GetMediaComprehensionJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaComprehensionJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a creative script generation task.
//
// @param request - GetRemakeScriptJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRemakeScriptJobResponse
func (client *Client) GetRemakeScriptJobWithContext(ctx context.Context, request *GetRemakeScriptJobRequest, runtime *dara.RuntimeOptions) (_result *GetRemakeScriptJobResponse, _err error) {
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
		Action:      dara.String("GetRemakeScriptJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRemakeScriptJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a video generation task.
//
// @param request - GetVideoGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoGenerationJobResponse
func (client *Client) GetVideoGenerationJobWithContext(ctx context.Context, request *GetVideoGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *GetVideoGenerationJobResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询爆款新视频渲染任务
//
// @param request - GetVideoRenderJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoRenderJobResponse
func (client *Client) GetVideoRenderJobWithContext(ctx context.Context, request *GetVideoRenderJobRequest, runtime *dara.RuntimeOptions) (_result *GetVideoRenderJobResponse, _err error) {
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
		Action:      dara.String("GetVideoRenderJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoRenderJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the membership plan and credit information for a Yike primary account.
//
// @param request - GetYikeAccountCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeAccountCreditResponse
func (client *Client) GetYikeAccountCreditWithContext(ctx context.Context, request *GetYikeAccountCreditRequest, runtime *dara.RuntimeOptions) (_result *GetYikeAccountCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeAccountCredit"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeAccountCreditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the actual credit consumption of a task.
//
// @param request - GetYikeJobCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetYikeJobCreditResponse
func (client *Client) GetYikeJobCreditWithContext(ctx context.Context, request *GetYikeJobCreditRequest, runtime *dara.RuntimeOptions) (_result *GetYikeJobCreditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetYikeJobCredit"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetYikeJobCreditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a media asset.
//
// Description:
//
// ## Operation description
//
// This API is used to query media content understanding jobs.
//
// @param request - ImportMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportMediaResponse
func (client *Client) ImportMediaWithContext(ctx context.Context, request *ImportMediaRequest, runtime *dara.RuntimeOptions) (_result *ImportMediaResponse, _err error) {
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

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicMetaData) {
		query["DynamicMetaData"] = request.DynamicMetaData
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.ImportSource) {
		query["ImportSource"] = request.ImportSource
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaTags) {
		query["MediaTags"] = request.MediaTags
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.RegisterConfig) {
		query["RegisterConfig"] = request.RegisterConfig
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
		Action:      dara.String("ImportMedia"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of categories.
//
// @param request - ListAssetCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssetCategoriesResponse
func (client *Client) ListAssetCategoriesWithContext(ctx context.Context, request *ListAssetCategoriesRequest, runtime *dara.RuntimeOptions) (_result *ListAssetCategoriesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAssetCategories"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAssetCategoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns media asset information that matches the specified filter conditions.
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
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMedia"),
		Version:     dara.String("2026-07-07"),
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
// Submits an image generation task.
//
// @param request - SubmitImageGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageGenerationJobResponse
func (client *Client) SubmitImageGenerationJobWithContext(ctx context.Context, request *SubmitImageGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitImageGenerationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AspectRatio) {
		query["AspectRatio"] = request.AspectRatio
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParameters) {
		query["JobParameters"] = request.JobParameters
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		query["N"] = request.N
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImageGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImageGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交媒资内容理解作业
//
// Description:
//
// ## 请求说明
//
// 该API用于根据提供的媒资文件（比如视频链接）进行内容理解。此外，支持通过`UserData`字段传递自定义参数，在回调时原样返回。
//
// @param request - SubmitMediaComprehensionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaComprehensionJobResponse
func (client *Client) SubmitMediaComprehensionJobWithContext(ctx context.Context, request *SubmitMediaComprehensionJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaComprehensionJobResponse, _err error) {
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

	if !dara.IsNil(request.JobParams) {
		query["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaComprehensionJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaComprehensionJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交爆款复刻新脚本生成任务
//
// Description:
//
// 该 API 用于根据内容理解的结果与新商品/模特信息，仿写生成新的口播脚本。此外，支持通过UserData字段传递自定义参数，在回调时原样返回。
//
// @param request - SubmitRemakeScriptJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitRemakeScriptJobResponse
func (client *Client) SubmitRemakeScriptJobWithContext(ctx context.Context, request *SubmitRemakeScriptJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitRemakeScriptJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RemakeParams) {
		query["RemakeParams"] = request.RemakeParams
	}

	if !dara.IsNil(request.RemakeType) {
		query["RemakeType"] = request.RemakeType
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitRemakeScriptJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitRemakeScriptJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a video generation task.
//
// @param request - SubmitVideoGenerationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoGenerationJobResponse
func (client *Client) SubmitVideoGenerationJobWithContext(ctx context.Context, request *SubmitVideoGenerationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoGenerationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AspectRatio) {
		query["AspectRatio"] = request.AspectRatio
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParameters) {
		query["JobParameters"] = request.JobParameters
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		query["N"] = request.N
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoGenerationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoGenerationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交爆款新视频渲染任务
//
// @param request - SubmitVideoRenderJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoRenderJobResponse
func (client *Client) SubmitVideoRenderJobWithContext(ctx context.Context, request *SubmitVideoRenderJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoRenderJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Script) {
		query["Script"] = request.Script
	}

	if !dara.IsNil(request.Settings) {
		query["Settings"] = request.Settings
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoRenderJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoRenderJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a video translation task that supports subtitle translation, voice translation, and on-screen text translation.
//
// Description:
//
// ## Request description
//
// - This API supports multiple video translation features, including subtitle translation and voice translation.
//
// - The `JobType` parameter defines the task type, such as `SubtitleTranslate` and `VoiceTranslate`.
//
// - The `Input` and `Output` parameters specify the input resource and output path, respectively.
//
// - `JobParameters` contains language configuration and other feature switches, such as `SourceLanguage`, `TargetLanguage`, `NeedDetext`, and `NeedVisualTranslate`.
//
// - `EditingConfig` can be used to specify the style configuration for the final editing and compositing.
//
// - `ClientToken` is an optional parameter used to ensure the idempotence of the request.
//
// - Ensure that all required fields are correctly filled in. Otherwise, the request may fail.
//
// @param request - SubmitVideoTranslationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoTranslationJobResponse
func (client *Client) SubmitVideoTranslationJobWithContext(ctx context.Context, request *SubmitVideoTranslationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoTranslationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Input) {
		body["Input"] = request.Input
	}

	if !dara.IsNil(request.JobParameters) {
		body["JobParameters"] = request.JobParameters
	}

	if !dara.IsNil(request.JobType) {
		body["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Output) {
		body["Output"] = request.Output
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoTranslationJob"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoTranslationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a media asset category.
//
// Description:
//
// After you create a media asset category, you can call this operation to locate and update the name of the media asset category by category ID.
//
// @param request - UpdateAssetCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAssetCategoryResponse
func (client *Client) UpdateAssetCategoryWithContext(ctx context.Context, request *UpdateAssetCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateAssetCategoryResponse, _err error) {
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

	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAssetCategory"),
		Version:     dara.String("2026-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAssetCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates media asset information.
//
// Description:
//
// ## Request description
//
// This API is used to query media content understanding jobs.
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
	if !dara.IsNil(request.AppendTags) {
		query["AppendTags"] = request.AppendTags
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicMetaData) {
		query["DynamicMetaData"] = request.DynamicMetaData
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaTags) {
		query["MediaTags"] = request.MediaTags
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
		Action:      dara.String("UpdateMedia"),
		Version:     dara.String("2026-07-07"),
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
