// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Compares configurations between two versions of a published playbook.
//
// @param request - ComparePlaybooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ComparePlaybooksResponse
func (client *Client) ComparePlaybooksWithContext(ctx context.Context, request *ComparePlaybooksRequest, runtime *dara.RuntimeOptions) (_result *ComparePlaybooksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewPlaybookReleaseId) {
		query["NewPlaybookReleaseId"] = request.NewPlaybookReleaseId
	}

	if !dara.IsNil(request.OldPlaybookReleaseId) {
		query["OldPlaybookReleaseId"] = request.OldPlaybookReleaseId
	}

	if !dara.IsNil(request.PlaybookUuid) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ComparePlaybooks"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ComparePlaybooksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Convert XML configuration.
//
// Description:
//
// Please ensure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/sas/detail/sas) of the orchestration product before using this interface.
//
// @param request - ConvertPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertPlaybookResponse
func (client *Client) ConvertPlaybookWithContext(ctx context.Context, request *ConvertPlaybookRequest, runtime *dara.RuntimeOptions) (_result *ConvertPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RoleFor) {
		query["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Taskflow) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a playbook.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR) or pricing for the log data added to the Cloud Threat Detection and Response (CTDR) feature. For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - CopyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyPlaybookResponse
func (client *Client) CopyPlaybookWithContext(ctx context.Context, request *CopyPlaybookRequest, runtime *dara.RuntimeOptions) (_result *CopyPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RoleFor) {
		query["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["ReleaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.SourcePlaybookUuid) {
		body["SourcePlaybookUuid"] = request.SourcePlaybookUuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// New Playbook.
//
// Description:
//
// Create Playbook.
//
// @param request - CreatePlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePlaybookResponse
func (client *Client) CreatePlaybookWithContext(ctx context.Context, request *CreatePlaybookRequest, runtime *dara.RuntimeOptions) (_result *CreatePlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InputParams) {
		body["InputParams"] = request.InputParams
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OutputParams) {
		body["OutputParams"] = request.OutputParams
	}

	if !dara.IsNil(request.TaskflowType) {
		body["TaskflowType"] = request.TaskflowType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Debugs a playbook.
//
// @param request - DebugPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugPlaybookResponse
func (client *Client) DebugPlaybookWithContext(ctx context.Context, request *DebugPlaybookRequest, runtime *dara.RuntimeOptions) (_result *DebugPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.Record) {
		body["Record"] = request.Record
	}

	if !dara.IsNil(request.Taskflow) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the assets in a component.
//
// @param request - DeleteComponentAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentAssetResponse
func (client *Client) DeleteComponentAssetWithContext(ctx context.Context, request *DeleteComponentAssetRequest, runtime *dara.RuntimeOptions) (_result *DeleteComponentAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComponentAsset"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComponentAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom playbook.
//
// @param request - DeletePlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePlaybookResponse
func (client *Client) DeletePlaybookWithContext(ctx context.Context, request *DeletePlaybookRequest, runtime *dara.RuntimeOptions) (_result *DeletePlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of assets in a component. The metadata of an asset refers to the fields that describe the asset.
//
// @param request - DescribeComponentAssetFormRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentAssetFormResponse
func (client *Client) DescribeComponentAssetFormWithContext(ctx context.Context, request *DescribeComponentAssetFormRequest, runtime *dara.RuntimeOptions) (_result *DescribeComponentAssetFormResponse, _err error) {
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
		Action:      dara.String("DescribeComponentAssetForm"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentAssetFormResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of assets in a component.
//
// @param request - DescribeComponentAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentAssetsResponse
func (client *Client) DescribeComponentAssetsWithContext(ctx context.Context, request *DescribeComponentAssetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeComponentAssetsResponse, _err error) {
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
		Action:      dara.String("DescribeComponentAssets"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of common components that are available.
//
// @param request - DescribeComponentListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentListResponse
func (client *Client) DescribeComponentListWithContext(ctx context.Context, request *DescribeComponentListRequest, runtime *dara.RuntimeOptions) (_result *DescribeComponentListResponse, _err error) {
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
		Action:      dara.String("DescribeComponentList"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of predefined components that are available.
//
// @param request - DescribeComponentPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentPlaybookResponse
func (client *Client) DescribeComponentPlaybookWithContext(ctx context.Context, request *DescribeComponentPlaybookRequest, runtime *dara.RuntimeOptions) (_result *DescribeComponentPlaybookResponse, _err error) {
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
		Action:      dara.String("DescribeComponentPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the JavaScript file of a component. The component uses the returned JavaScript file for page rendering.
//
// @param request - DescribeComponentsJsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentsJsResponse
func (client *Client) DescribeComponentsJsWithContext(ctx context.Context, request *DescribeComponentsJsRequest, runtime *dara.RuntimeOptions) (_result *DescribeComponentsJsResponse, _err error) {
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
		Action:      dara.String("DescribeComponentsJs"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentsJsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the published versions of a playbook after deduplication.
//
// @param request - DescribeDistinctReleasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistinctReleasesResponse
func (client *Client) DescribeDistinctReleasesWithContext(ctx context.Context, request *DescribeDistinctReleasesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDistinctReleasesResponse, _err error) {
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
		Action:      dara.String("DescribeDistinctReleases"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDistinctReleasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries enumeration items that are required by a cloud service.
//
// @param request - DescribeEnumItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnumItemsResponse
func (client *Client) DescribeEnumItemsWithContext(ctx context.Context, request *DescribeEnumItemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnumItemsResponse, _err error) {
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
		Action:      dara.String("DescribeEnumItems"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnumItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the playbooks that are available for an automatic response plan.
//
// @param request - DescribeExecutePlaybooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExecutePlaybooksResponse
func (client *Client) DescribeExecutePlaybooksWithContext(ctx context.Context, request *DescribeExecutePlaybooksRequest, runtime *dara.RuntimeOptions) (_result *DescribeExecutePlaybooksResponse, _err error) {
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
		Action:      dara.String("DescribeExecutePlaybooks"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExecutePlaybooksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global configuration information about a cloud service.
//
// @param request - DescribeFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFieldResponse
func (client *Client) DescribeFieldWithContext(ctx context.Context, request *DescribeFieldRequest, runtime *dara.RuntimeOptions) (_result *DescribeFieldResponse, _err error) {
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
		Action:      dara.String("DescribeField"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries groups of Alibaba Cloud services.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - DescribeGroupProductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupProductionsResponse
func (client *Client) DescribeGroupProductionsWithContext(ctx context.Context, request *DescribeGroupProductionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupProductionsResponse, _err error) {
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
		Action:      dara.String("DescribeGroupProductions"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupProductionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the output structure information of each node in a playbook based on the most recent running record of the playbook.
//
// @param request - DescribeLatestRecordSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLatestRecordSchemaResponse
func (client *Client) DescribeLatestRecordSchemaWithContext(ctx context.Context, request *DescribeLatestRecordSchemaRequest, runtime *dara.RuntimeOptions) (_result *DescribeLatestRecordSchemaResponse, _err error) {
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
		Action:      dara.String("DescribeLatestRecordSchema"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLatestRecordSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries recommended dynamic input parameters of a component for playbook orchestration.
//
// @param request - DescribeNodeParamTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeParamTagsResponse
func (client *Client) DescribeNodeParamTagsWithContext(ctx context.Context, request *DescribeNodeParamTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodeParamTagsResponse, _err error) {
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
		Action:      dara.String("DescribeNodeParamTags"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodeParamTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries notification templates.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - DescribeNotifyTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNotifyTemplateListResponse
func (client *Client) DescribeNotifyTemplateListWithContext(ctx context.Context, request *DescribeNotifyTemplateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNotifyTemplateListResponse, _err error) {
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
		Action:      dara.String("DescribeNotifyTemplateList"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNotifyTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an API operation.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR) or pricing for the log data added to the Cloud Threat Detection and Response (CTDR) feature. For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - DescribeOpenApiInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenApiInfoResponse
func (client *Client) DescribeOpenApiInfoWithContext(ctx context.Context, request *DescribeOpenApiInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpenApiInfoResponse, _err error) {
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
		Action:      dara.String("DescribeOpenApiInfo"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOpenApiInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the API operations of an Alibaba Cloud service.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR) or the pricing for log data added to the Cloud Threat Detection and Response (CTDR) feature. For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - DescribeOpenApiListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenApiListResponse
func (client *Client) DescribeOpenApiListWithContext(ctx context.Context, request *DescribeOpenApiListRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpenApiListResponse, _err error) {
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
		Action:      dara.String("DescribeOpenApiList"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOpenApiListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the XML configuration of a playbook.
//
// @param request - DescribePlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookResponse
func (client *Client) DescribePlaybookWithContext(ctx context.Context, request *DescribePlaybookRequest, runtime *dara.RuntimeOptions) (_result *DescribePlaybookResponse, _err error) {
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
		Action:      dara.String("DescribePlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the input and output parameter configurations of a playbook.
//
// @param request - DescribePlaybookInputOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookInputOutputResponse
func (client *Client) DescribePlaybookInputOutputWithContext(ctx context.Context, request *DescribePlaybookInputOutputRequest, runtime *dara.RuntimeOptions) (_result *DescribePlaybookInputOutputResponse, _err error) {
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
		Action:      dara.String("DescribePlaybookInputOutput"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlaybookInputOutputResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics of a playbook. The metrics include the playbook name, playbook description, the number of times that the playbook is run, and the failure rate of the playbook.
//
// @param request - DescribePlaybookMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookMetricsResponse
func (client *Client) DescribePlaybookMetricsWithContext(ctx context.Context, request *DescribePlaybookMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribePlaybookMetricsResponse, _err error) {
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
		Action:      dara.String("DescribePlaybookMetrics"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlaybookMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical output data of a component node.
//
// @param request - DescribePlaybookNodesOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookNodesOutputResponse
func (client *Client) DescribePlaybookNodesOutputWithContext(ctx context.Context, request *DescribePlaybookNodesOutputRequest, runtime *dara.RuntimeOptions) (_result *DescribePlaybookNodesOutputResponse, _err error) {
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
		Action:      dara.String("DescribePlaybookNodesOutput"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlaybookNodesOutputResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of Security Orchestration Automation Response (SOAR), such as the numbers of created and enabled playbooks.
//
// @param request - DescribePlaybookNumberMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookNumberMetricsResponse
func (client *Client) DescribePlaybookNumberMetricsWithContext(ctx context.Context, request *DescribePlaybookNumberMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribePlaybookNumberMetricsResponse, _err error) {
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
		Action:      dara.String("DescribePlaybookNumberMetrics"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlaybookNumberMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the published versions of a playbook.
//
// @param request - DescribePlaybookReleasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookReleasesResponse
func (client *Client) DescribePlaybookReleasesWithContext(ctx context.Context, request *DescribePlaybookReleasesRequest, runtime *dara.RuntimeOptions) (_result *DescribePlaybookReleasesResponse, _err error) {
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
		Action:      dara.String("DescribePlaybookReleases"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlaybookReleasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of playbooks.
//
// @param request - DescribePlaybooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybooksResponse
func (client *Client) DescribePlaybooksWithContext(ctx context.Context, request *DescribePlaybooksRequest, runtime *dara.RuntimeOptions) (_result *DescribePlaybooksResponse, _err error) {
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
		Action:      dara.String("DescribePlaybooks"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlaybooksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an API operation.
//
// @param request - DescribePopApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePopApiResponse
func (client *Client) DescribePopApiWithContext(ctx context.Context, request *DescribePopApiRequest, runtime *dara.RuntimeOptions) (_result *DescribePopApiResponse, _err error) {
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
		Action:      dara.String("DescribePopApi"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePopApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR) or pricing for the log data added to the Cloud Threat Detection and Response (CTDR) feature. For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - DescribeProcessStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessStatisticsResponse
func (client *Client) DescribeProcessStatisticsWithContext(ctx context.Context, request *DescribeProcessStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeProcessStatisticsResponse, _err error) {
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
		Action:      dara.String("DescribeProcessStatistics"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProcessStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the number of associated disposal tasks based on the entity UUID.
//
// @param request - DescribeProcessTaskCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessTaskCountResponse
func (client *Client) DescribeProcessTaskCountWithContext(ctx context.Context, request *DescribeProcessTaskCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeProcessTaskCountResponse, _err error) {
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
		Action:      dara.String("DescribeProcessTaskCount"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProcessTaskCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about handling tasks. When you use Security Orchestration Automation Response (SOAR) to handle events, handling tasks are generated in the handling center.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR) or pricing for the log data added to the Cloud Threat Detection and Response (CTDR) feature. For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - DescribeProcessTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessTasksResponse
func (client *Client) DescribeProcessTasksWithContext(ctx context.Context, request *DescribeProcessTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeProcessTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EntityName) {
		query["EntityName"] = request.EntityName
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.EntityUuid) {
		query["EntityUuid"] = request.EntityUuid
	}

	if !dara.IsNil(request.EventUuid) {
		query["EventUuid"] = request.EventUuid
	}

	if !dara.IsNil(request.OrderField) {
		query["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParamContent) {
		query["ParamContent"] = request.ParamContent
	}

	if !dara.IsNil(request.ProcessActionEnd) {
		query["ProcessActionEnd"] = request.ProcessActionEnd
	}

	if !dara.IsNil(request.ProcessActionStart) {
		query["ProcessActionStart"] = request.ProcessActionStart
	}

	if !dara.IsNil(request.ProcessRemoveEnd) {
		query["ProcessRemoveEnd"] = request.ProcessRemoveEnd
	}

	if !dara.IsNil(request.ProcessRemoveStart) {
		query["ProcessRemoveStart"] = request.ProcessRemoveStart
	}

	if !dara.IsNil(request.ProcessStrategyUuid) {
		query["ProcessStrategyUuid"] = request.ProcessStrategyUuid
	}

	if !dara.IsNil(request.ReqUuid) {
		query["ReqUuid"] = request.ReqUuid
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TriggerSource) {
		query["TriggerSource"] = request.TriggerSource
	}

	if !dara.IsNil(request.YunCode) {
		query["YunCode"] = request.YunCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProcessTasks"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProcessTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data that is returned when a component initiates an action in a playbook task.
//
// @param request - DescribeSoarRecordActionOutputListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarRecordActionOutputListResponse
func (client *Client) DescribeSoarRecordActionOutputListWithContext(ctx context.Context, request *DescribeSoarRecordActionOutputListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSoarRecordActionOutputListResponse, _err error) {
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
		Action:      dara.String("DescribeSoarRecordActionOutputList"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSoarRecordActionOutputListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the input and output data of a component action. You can call this operation after a playbook is run.
//
// @param request - DescribeSoarRecordInOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarRecordInOutputResponse
func (client *Client) DescribeSoarRecordInOutputWithContext(ctx context.Context, request *DescribeSoarRecordInOutputRequest, runtime *dara.RuntimeOptions) (_result *DescribeSoarRecordInOutputResponse, _err error) {
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
		Action:      dara.String("DescribeSoarRecordInOutput"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSoarRecordInOutputResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the execution records of a playbook.
//
// @param request - DescribeSoarRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarRecordsResponse
func (client *Client) DescribeSoarRecordsWithContext(ctx context.Context, request *DescribeSoarRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSoarRecordsResponse, _err error) {
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
		Action:      dara.String("DescribeSoarRecords"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSoarRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution records of a component during the running of a playbook.
//
// @param request - DescribeSoarTaskAndActionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarTaskAndActionsResponse
func (client *Client) DescribeSoarTaskAndActionsWithContext(ctx context.Context, request *DescribeSoarTaskAndActionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSoarTaskAndActionsResponse, _err error) {
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
		Action:      dara.String("DescribeSoarTaskAndActions"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSoarTaskAndActionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the commands that can be run to obtain objects.
//
// @param request - DescribeSophonCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSophonCommandsResponse
func (client *Client) DescribeSophonCommandsWithContext(ctx context.Context, request *DescribeSophonCommandsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSophonCommandsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSophonCommands"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSophonCommandsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query OpenApi List of Cloud Vendors.
//
// Description:
//
// Please ensure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/sas/detail/sas) of the response orchestration product (i.e., threat analysis and response log access traffic) before using this interface.
//
// @param request - DescribeVendorApiListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVendorApiListResponse
func (client *Client) DescribeVendorApiListWithContext(ctx context.Context, request *DescribeVendorApiListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVendorApiListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.VendorCode) {
		query["VendorCode"] = request.VendorCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVendorApiList"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVendorApiListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operational logs of a Python3 script by using the UUID that is returned when the script is run. The UUID is specified by requestUuid.
//
// @param request - DescriberPython3ScriptLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescriberPython3ScriptLogsResponse
func (client *Client) DescriberPython3ScriptLogsWithContext(ctx context.Context, request *DescriberPython3ScriptLogsRequest, runtime *dara.RuntimeOptions) (_result *DescriberPython3ScriptLogsResponse, _err error) {
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
		Action:      dara.String("DescriberPython3ScriptLogs"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescriberPython3ScriptLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about the asset that is configured for a component.
//
// @param request - ModifyComponentAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyComponentAssetResponse
func (client *Client) ModifyComponentAssetWithContext(ctx context.Context, request *ModifyComponentAssetRequest, runtime *dara.RuntimeOptions) (_result *ModifyComponentAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetConfig) {
		query["AssetConfig"] = request.AssetConfig
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyComponentAsset"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyComponentAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a playbook.
//
// @param request - ModifyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPlaybookResponse
func (client *Client) ModifyPlaybookWithContext(ctx context.Context, request *ModifyPlaybookRequest, runtime *dara.RuntimeOptions) (_result *ModifyPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.Taskflow) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the input and output parameters of a playbook.
//
// @param request - ModifyPlaybookInputOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPlaybookInputOutputResponse
func (client *Client) ModifyPlaybookInputOutputWithContext(ctx context.Context, request *ModifyPlaybookInputOutputRequest, runtime *dara.RuntimeOptions) (_result *ModifyPlaybookInputOutputResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExeConfig) {
		body["ExeConfig"] = request.ExeConfig
	}

	if !dara.IsNil(request.InputParams) {
		body["InputParams"] = request.InputParams
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OutputParams) {
		body["OutputParams"] = request.OutputParams
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPlaybookInputOutput"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPlaybookInputOutputResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes the playbook. After the playbook is published, the playbook runs based on the new logic.
//
// @param request - PublishPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishPlaybookResponse
func (client *Client) PublishPlaybookWithContext(ctx context.Context, request *PublishPlaybookRequest, runtime *dara.RuntimeOptions) (_result *PublishPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all playbooks at a time.
//
// @param request - QueryTreeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTreeDataResponse
func (client *Client) QueryTreeDataWithContext(ctx context.Context, request *QueryTreeDataRequest, runtime *dara.RuntimeOptions) (_result *QueryTreeDataResponse, _err error) {
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
		Action:      dara.String("QueryTreeData"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTreeDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a playbook to a specific version. You can determine whether to publish the new playbook version during the rollback.
//
// @param request - RevertPlaybookReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertPlaybookReleaseResponse
func (client *Client) RevertPlaybookReleaseWithContext(ctx context.Context, request *RevertPlaybookReleaseRequest, runtime *dara.RuntimeOptions) (_result *RevertPlaybookReleaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsPublish) {
		body["IsPublish"] = request.IsPublish
	}

	if !dara.IsNil(request.PlayReleaseId) {
		body["PlayReleaseId"] = request.PlayReleaseId
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertPlaybookRelease"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertPlaybookReleaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs the email notification component to send messages.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR) or pricing for the log data added to the Cloud Threat Detection and Response (CTDR) feature. For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - RunNotifyComponentWithEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNotifyComponentWithEmailResponse
func (client *Client) RunNotifyComponentWithEmailWithContext(ctx context.Context, request *RunNotifyComponentWithEmailRequest, runtime *dara.RuntimeOptions) (_result *RunNotifyComponentWithEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionName) {
		query["ActionName"] = request.ActionName
	}

	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.PlaybookUuid) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.Receivers) {
		query["Receivers"] = request.Receivers
	}

	if !dara.IsNil(request.RoleFor) {
		query["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.Subject) {
		query["Subject"] = request.Subject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunNotifyComponentWithEmail"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunNotifyComponentWithEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Execute Notification Component - Send Message via Message Center.
//
// Description:
//
// Please ensure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/sas/detail/sas) of the response orchestration product (i.e., Threat Analysis and Response Log Access Traffic) before using this interface.
//
// @param request - RunNotifyComponentWithMessageCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNotifyComponentWithMessageCenterResponse
func (client *Client) RunNotifyComponentWithMessageCenterWithContext(ctx context.Context, request *RunNotifyComponentWithMessageCenterRequest, runtime *dara.RuntimeOptions) (_result *RunNotifyComponentWithMessageCenterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionName) {
		query["ActionName"] = request.ActionName
	}

	if !dara.IsNil(request.Aliuid) {
		query["Aliuid"] = request.Aliuid
	}

	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.ChannelTypeList) {
		query["ChannelTypeList"] = request.ChannelTypeList
	}

	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.PlaybookUuid) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.RoleFor) {
		query["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunNotifyComponentWithMessageCenter"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunNotifyComponentWithMessageCenterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs the webhook notification component to send messages.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR) or pricing for the log data added to the Cloud Threat Detection and Response (CTDR) feature. For more information, see [Pricing](https://www.aliyun.com/price/product#/sas/detail/sas).
//
// @param request - RunNotifyComponentWithWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNotifyComponentWithWebhookResponse
func (client *Client) RunNotifyComponentWithWebhookWithContext(ctx context.Context, request *RunNotifyComponentWithWebhookRequest, runtime *dara.RuntimeOptions) (_result *RunNotifyComponentWithWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionName) {
		query["ActionName"] = request.ActionName
	}

	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MsgType) {
		query["MsgType"] = request.MsgType
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.PlaybookUuid) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.RoleFor) {
		query["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.Secret) {
		query["Secret"] = request.Secret
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunNotifyComponentWithWebhook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunNotifyComponentWithWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits and runs a Python3 script. You can call this operation only for data processing.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.4c41281fWhbdPa#/commodity/vm_intl).
//
// @param request - RunPython3ScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPython3ScriptResponse
func (client *Client) RunPython3ScriptWithContext(ctx context.Context, request *RunPython3ScriptRequest, runtime *dara.RuntimeOptions) (_result *RunPython3ScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PythonVersion) {
		query["PythonVersion"] = request.PythonVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeName) {
		body["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.PythonScript) {
		body["PythonScript"] = request.PythonScript
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunPython3Script"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunPython3ScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers an enabled custom playbook or a predefined playbook.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=a2796.7960336.3034855210.1.7adab91arMeIx2#/commodity/vm_intl).
//
// @param request - TriggerPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerPlaybookResponse
func (client *Client) TriggerPlaybookWithContext(ctx context.Context, request *TriggerPlaybookRequest, runtime *dara.RuntimeOptions) (_result *TriggerPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InputParam) {
		body["InputParam"] = request.InputParam
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs an action on a handling task that is generated by the handling center when an event is handled by using Security Orchestration Automation Response (SOAR). For example, you can call this operation to cancel blocking or isolation, or retry blocking.
//
// @param request - TriggerProcessTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerProcessTaskResponse
func (client *Client) TriggerProcessTaskWithContext(ctx context.Context, request *TriggerProcessTaskRequest, runtime *dara.RuntimeOptions) (_result *TriggerProcessTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerProcessTask"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerProcessTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a playbook or a command.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=a2796.7960336.3034855210.1.7adab91arMeIx2#/commodity/vm_intl).
//
// @param request - TriggerSophonPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerSophonPlaybookResponse
func (client *Client) TriggerSophonPlaybookWithContext(ctx context.Context, request *TriggerSophonPlaybookRequest, runtime *dara.RuntimeOptions) (_result *TriggerSophonPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandName) {
		query["CommandName"] = request.CommandName
	}

	if !dara.IsNil(request.InputParams) {
		query["InputParams"] = request.InputParams
	}

	if !dara.IsNil(request.SophonTaskId) {
		query["SophonTaskId"] = request.SophonTaskId
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerSophonPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerSophonPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the configuration of the playbook is correct and whether the logic of the orchestration is reasonable.
//
// @param request - VerifyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyPlaybookResponse
func (client *Client) VerifyPlaybookWithContext(ctx context.Context, request *VerifyPlaybookRequest, runtime *dara.RuntimeOptions) (_result *VerifyPlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.TaskFlow) {
		body["TaskFlow"] = request.TaskFlow
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyPlaybook"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the syntax of a Python code snippet is correct.
//
// @param request - VerifyPythonFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyPythonFileResponse
func (client *Client) VerifyPythonFileWithContext(ctx context.Context, request *VerifyPythonFileRequest, runtime *dara.RuntimeOptions) (_result *VerifyPythonFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyPythonFile"),
		Version:     dara.String("2022-07-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyPythonFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
