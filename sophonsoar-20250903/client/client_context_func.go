// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Create Component Asset.
//
// Description:
//
// Please ensure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/sas/detail/sas) of the response orchestration product (i.e., Threat Analysis and Response Log Ingress Traffic) before using this interface.
//
// @param request - CreateComponentAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentAssetResponse
func (client *Client) CreateComponentAssetWithContext(ctx context.Context, request *CreateComponentAssetRequest, runtime *dara.RuntimeOptions) (_result *CreateComponentAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentAssetName) {
		body["ComponentAssetName"] = request.ComponentAssetName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ComponentAssetValues) {
		bodyFlat["ComponentAssetValues"] = request.ComponentAssetValues
	}

	if !dara.IsNil(request.ComponentName) {
		body["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComponentAsset"),
		Version:     dara.String("2025-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComponentAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
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
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PlaybookDescription) {
		body["PlaybookDescription"] = request.PlaybookDescription
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.PlaybookInputConfigs) {
		bodyFlat["PlaybookInputConfigs"] = request.PlaybookInputConfigs
	}

	if !dara.IsNil(request.PlaybookName) {
		body["PlaybookName"] = request.PlaybookName
	}

	if !dara.IsNil(request.PlaybookOutputConfigs) {
		bodyFlat["PlaybookOutputConfigs"] = request.PlaybookOutputConfigs
	}

	if !dara.IsNil(request.PlaybookParamType) {
		body["PlaybookParamType"] = request.PlaybookParamType
	}

	if !dara.IsNil(request.PlaybookTaskFlow) {
		body["PlaybookTaskFlow"] = request.PlaybookTaskFlow
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.SrcPlaybookUuid) {
		body["SrcPlaybookUuid"] = request.SrcPlaybookUuid
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePlaybook"),
		Version:     dara.String("2025-09-03"),
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
// Delete Component Asset.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the billing method and [pricing](https://www.aliyun.com/price/product#/sas/detail/sas) of the response orchestration product (i.e., Threat Analysis and Response Log Access Traffic).
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentAssetUuid) {
		body["ComponentAssetUuid"] = request.ComponentAssetUuid
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComponentAsset"),
		Version:     dara.String("2025-09-03"),
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
// Delete Playbook.
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

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePlaybook"),
		Version:     dara.String("2025-09-03"),
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
// Get playbook details.
//
// @param request - GetPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPlaybookResponse
func (client *Client) GetPlaybookWithContext(ctx context.Context, request *GetPlaybookRequest, runtime *dara.RuntimeOptions) (_result *GetPlaybookResponse, _err error) {
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

	if !dara.IsNil(request.PlaybookVersion) {
		body["PlaybookVersion"] = request.PlaybookVersion
	}

	if !dara.IsNil(request.PlaybookVersionType) {
		body["PlaybookVersionType"] = request.PlaybookVersionType
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPlaybook"),
		Version:     dara.String("2025-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the list of assets configured for a component.
//
// @param request - ListComponentAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentAssetsResponse
func (client *Client) ListComponentAssetsWithContext(ctx context.Context, request *ListComponentAssetsRequest, runtime *dara.RuntimeOptions) (_result *ListComponentAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentAssetUuid) {
		body["ComponentAssetUuid"] = request.ComponentAssetUuid
	}

	if !dara.IsNil(request.ComponentName) {
		body["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponentAssets"),
		Version:     dara.String("2025-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get Component List.
//
// @param request - ListComponentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentsResponse
func (client *Client) ListComponentsWithContext(ctx context.Context, request *ListComponentsRequest, runtime *dara.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentName) {
		body["ComponentName"] = request.ComponentName
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponents"),
		Version:     dara.String("2025-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get Playbook List.
//
// @param tmpReq - ListPlaybooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPlaybooksResponse
func (client *Client) ListPlaybooksWithContext(ctx context.Context, tmpReq *ListPlaybooksRequest, runtime *dara.RuntimeOptions) (_result *ListPlaybooksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPlaybooksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PlaybookParamTypes) {
		request.PlaybookParamTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlaybookParamTypes, dara.String("PlaybookParamTypes"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.PlaybookUuids) {
		request.PlaybookUuidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlaybookUuids, dara.String("PlaybookUuids"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlaybookExecutionEndTime) {
		body["PlaybookExecutionEndTime"] = request.PlaybookExecutionEndTime
	}

	if !dara.IsNil(request.PlaybookExecutionStartTime) {
		body["PlaybookExecutionStartTime"] = request.PlaybookExecutionStartTime
	}

	if !dara.IsNil(request.PlaybookName) {
		body["PlaybookName"] = request.PlaybookName
	}

	if !dara.IsNil(request.PlaybookParamTypesShrink) {
		body["PlaybookParamTypes"] = request.PlaybookParamTypesShrink
	}

	if !dara.IsNil(request.PlaybookStatus) {
		body["PlaybookStatus"] = request.PlaybookStatus
	}

	if !dara.IsNil(request.PlaybookType) {
		body["PlaybookType"] = request.PlaybookType
	}

	if !dara.IsNil(request.PlaybookUuidsShrink) {
		body["PlaybookUuids"] = request.PlaybookUuidsShrink
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlaybooks"),
		Version:     dara.String("2025-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPlaybooksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update Component Asset.
//
// @param request - UpdateComponentAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentAssetResponse
func (client *Client) UpdateComponentAssetWithContext(ctx context.Context, request *UpdateComponentAssetRequest, runtime *dara.RuntimeOptions) (_result *UpdateComponentAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentAssetName) {
		body["ComponentAssetName"] = request.ComponentAssetName
	}

	if !dara.IsNil(request.ComponentAssetUuid) {
		body["ComponentAssetUuid"] = request.ComponentAssetUuid
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ComponentAssetValues) {
		bodyFlat["ComponentAssetValues"] = request.ComponentAssetValues
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComponentAsset"),
		Version:     dara.String("2025-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComponentAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update Playbook.
//
// @param tmpReq - UpdatePlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePlaybookResponse
func (client *Client) UpdatePlaybookWithContext(ctx context.Context, tmpReq *UpdatePlaybookRequest, runtime *dara.RuntimeOptions) (_result *UpdatePlaybookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePlaybookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PlaybookInputConfigs) {
		request.PlaybookInputConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlaybookInputConfigs, dara.String("PlaybookInputConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PlaybookOutputConfigs) {
		request.PlaybookOutputConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlaybookOutputConfigs, dara.String("PlaybookOutputConfigs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PlaybookDescription) {
		body["PlaybookDescription"] = request.PlaybookDescription
	}

	if !dara.IsNil(request.PlaybookInputConfigsShrink) {
		body["PlaybookInputConfigs"] = request.PlaybookInputConfigsShrink
	}

	if !dara.IsNil(request.PlaybookName) {
		body["PlaybookName"] = request.PlaybookName
	}

	if !dara.IsNil(request.PlaybookOutputConfigsShrink) {
		body["PlaybookOutputConfigs"] = request.PlaybookOutputConfigsShrink
	}

	if !dara.IsNil(request.PlaybookParamType) {
		body["PlaybookParamType"] = request.PlaybookParamType
	}

	if !dara.IsNil(request.PlaybookTaskFlow) {
		body["PlaybookTaskFlow"] = request.PlaybookTaskFlow
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePlaybook"),
		Version:     dara.String("2025-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePlaybookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
