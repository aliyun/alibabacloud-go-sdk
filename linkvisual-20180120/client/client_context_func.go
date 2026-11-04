// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates an intelligent alert event recording plan for an IP camera.
//
// @param request - AddEventRecordPlanDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEventRecordPlanDeviceResponse
func (client *Client) AddEventRecordPlanDeviceWithContext(ctx context.Context, request *AddEventRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *AddEventRecordPlanDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEventRecordPlanDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEventRecordPlanDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a device group for the face library.
//
// Description:
//
// ## Usage notes
//
// - The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// - After you invoke this API operation, the system returns the ID of the device group (**DeviceGroupId**). Save this value locally for use when invoking other API operations.
//
// @param request - AddFaceDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceDeviceGroupResponse
func (client *Client) AddFaceDeviceGroupWithContext(ctx context.Context, request *AddFaceDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupName) {
		query["DeviceGroupName"] = request.DeviceGroupName
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceDeviceGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceDeviceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a device to a device group.
//
// @param request - AddFaceDeviceToDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceDeviceToDeviceGroupResponse
func (client *Client) AddFaceDeviceToDeviceGroupWithContext(ctx context.Context, request *AddFaceDeviceToDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceDeviceToDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceDeviceToDeviceGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceDeviceToDeviceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to the face database.
//
// Description:
//
// ## Usage notes
//
// - The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// - Configure a URL for downloading the user\\"s face image in advance.
//
//	> Only JPEG and JPG image formats are supported, and the image URL cannot contain Chinese characters.
//
// - After you invoke this API operation, the system returns a user ID (**UserId**). Save this value locally for use when you invoke other API operations.
//
// @param request - AddFaceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceUserResponse
func (client *Client) AddFaceUserWithContext(ctx context.Context, request *AddFaceUserRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomUserId) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !dara.IsNil(request.FacePicUrl) {
		query["FacePicUrl"] = request.FacePicUrl
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceUser"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user group.
//
// Description:
//
// ## Usage notes
//
// - The maximum queries per second (QPS) limit for a single Alibaba Cloud account to invoke this API operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// - After you invoke this API operation, the system returns the ID of the user group (**UserGroupId**). Save this value locally for use when you invoke other API operations.
//
// @param request - AddFaceUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceUserGroupResponse
func (client *Client) AddFaceUserGroupWithContext(ctx context.Context, request *AddFaceUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceUserGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a face synchronization relationship between a user group and a device group.
//
// @param request - AddFaceUserGroupAndDeviceGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) AddFaceUserGroupAndDeviceGroupRelationWithContext(ctx context.Context, request *AddFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.Relation) {
		query["Relation"] = request.Relation
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceUserGroupAndDeviceGroupRelation"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a face image for a user in the face database.
//
// Description:
//
// ## Usage notes
//
// - The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// - Configure a download URL for the user\\"s face image in advance.
//
//	> Only images in JPEG and JPG format are supported, and the image URL cannot contain Chinese characters.
//
// - After you invoke this API operation, the system returns the unique identity (**Data**) of the user\\"s face image. Save this value locally for use when you invoke other API operations.
//
// @param request - AddFaceUserPictureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceUserPictureResponse
func (client *Client) AddFaceUserPictureWithContext(ctx context.Context, request *AddFaceUserPictureRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserPictureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FacePicUrl) {
		query["FacePicUrl"] = request.FacePicUrl
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceUserPicture"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceUserPictureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a user group.
//
// Description:
//
// ## Limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - AddFaceUserToUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceUserToUserGroupResponse
func (client *Client) AddFaceUserToUserGroupWithContext(ctx context.Context, request *AddFaceUserToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserToUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFaceUserToUserGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFaceUserToUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a continuous recording plan for an IPC device.
//
// @param request - AddRecordPlanDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordPlanDeviceResponse
func (client *Client) AddRecordPlanDeviceWithContext(ctx context.Context, request *AddRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *AddRecordPlanDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecordPlanDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecordPlanDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchQueryVisionDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryVisionDeviceInfoResponse
func (client *Client) BatchQueryVisionDeviceInfoWithContext(ctx context.Context, request *BatchQueryVisionDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *BatchQueryVisionDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceNameList) {
		query["DeviceNameList"] = request.DeviceNameList
	}

	if !dara.IsNil(request.IotIdList) {
		query["IotIdList"] = request.IotIdList
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchQueryVisionDeviceInfo"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchQueryVisionDeviceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds devices to an image search instance in batches.
//
// Description:
//
// ## Limits
//
// If you invoke this operation on an Enterprise instance, you must specify the **IotInstanceId*	- request parameter. Otherwise, the invocation failed.
//
// @param request - BindPictureSearchAppWithDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPictureSearchAppWithDevicesResponse
func (client *Client) BindPictureSearchAppWithDevicesWithContext(ctx context.Context, request *BindPictureSearchAppWithDevicesRequest, runtime *dara.RuntimeOptions) (_result *BindPictureSearchAppWithDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.DeviceIotIds) {
		query["DeviceIotIds"] = request.DeviceIotIds
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPictureSearchAppWithDevices"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the face of a specified user is successfully synchronized on the device.
//
// @param request - CheckFaceUserDoExistOnDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckFaceUserDoExistOnDeviceResponse
func (client *Client) CheckFaceUserDoExistOnDeviceWithContext(ctx context.Context, request *CheckFaceUserDoExistOnDeviceRequest, runtime *dara.RuntimeOptions) (_result *CheckFaceUserDoExistOnDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckFaceUserDoExistOnDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckFaceUserDoExistOnDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears all face databases under a specified business isolation ID for a V-series edge all-in-one machine.
//
// Description:
//
// ## Limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - ClearFaceDeviceDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearFaceDeviceDBResponse
func (client *Client) ClearFaceDeviceDBWithContext(ctx context.Context, request *ClearFaceDeviceDBRequest, runtime *dara.RuntimeOptions) (_result *ClearFaceDeviceDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearFaceDeviceDB"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearFaceDeviceDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a recording plan for intelligent alert events.
//
// @param request - CreateEventRecordPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventRecordPlanResponse
func (client *Client) CreateEventRecordPlanWithContext(ctx context.Context, request *CreateEventRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateEventRecordPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventTypes) {
		query["EventTypes"] = request.EventTypes
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PreRecordDuration) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !dara.IsNil(request.RecordDuration) {
		query["RecordDuration"] = request.RecordDuration
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventRecordPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventRecordPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateGbDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGbDeviceResponse
func (client *Client) CreateGbDeviceWithContext(ctx context.Context, request *CreateGbDeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateGbDeviceResponse, _err error) {
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

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.MediaNetProtocol) {
		query["MediaNetProtocol"] = request.MediaNetProtocol
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.SubProductKey) {
		query["SubProductKey"] = request.SubProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGbDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGbDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a local video recording upload job.
//
// Description:
//
// A single local video recording upload job supports uploading up to 200 files. The maximum number of items in the TimeSlot list is 200.
//
// @param request - CreateLocalFileUploadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLocalFileUploadJobResponse
func (client *Client) CreateLocalFileUploadJobWithContext(ctx context.Context, request *CreateLocalFileUploadJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLocalFileUploadJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.TimeSlot) {
		query["TimeSlot"] = request.TimeSlot
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLocalFileUploadJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLocalFileUploadJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to download local video recordings by time.
//
// Description:
//
// ## Limits
//
// - The maximum duration for a single time-based download is 30 minutes.
//
// - Download tasks and files are retained for up to 1 day. You can re-download them after they expire.
//
// - A single instance supports up to 10 concurrent downloads. Completed downloads are not counted.
//
// - If the audio and video encoding parameters change, the downloaded recording defaults to the first half of the recording.
//
// - Due to fluctuations in the device\\"s upstream network, the link may be abnormally disconnected or the device may actively drop data. As a result, the uploaded and downloaded recordings may be incomplete.
//
// - The downloaded recording data is determined by the device\\"s stream pushing logic. If the pushed data is abnormal, the duration of the downloaded recording may not match the time period specified in the task settings.
//
// ## QPS limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this operation is 50.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - CreateLocalRecordDownloadByTimeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLocalRecordDownloadByTimeJobResponse
func (client *Client) CreateLocalRecordDownloadByTimeJobWithContext(ctx context.Context, request *CreateLocalRecordDownloadByTimeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLocalRecordDownloadByTimeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLocalRecordDownloadByTimeJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLocalRecordDownloadByTimeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image search application instance.
//
// Description:
//
// ## Limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
// > A Resource Access Management (RAM) user shares the quota of the Alibaba Cloud account.
//
// @param request - CreatePictureSearchAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePictureSearchAppResponse
func (client *Client) CreatePictureSearchAppWithContext(ctx context.Context, request *CreatePictureSearchAppRequest, runtime *dara.RuntimeOptions) (_result *CreatePictureSearchAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePictureSearchApp"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePictureSearchAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an integrated image search task.
//
// Description:
//
// Integrated image search retrieves both face and human body features simultaneously during an image search.
//
// > This feature requires the deployment of face and human body detection algorithms on V-series edge all-in-one machines.
//
// @param request - CreatePictureSearchJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePictureSearchJobResponse
func (client *Client) CreatePictureSearchJobWithContext(ctx context.Context, request *CreatePictureSearchJobRequest, runtime *dara.RuntimeOptions) (_result *CreatePictureSearchJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.BodyThreshold) {
		query["BodyThreshold"] = request.BodyThreshold
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PictureSearchType) {
		query["PictureSearchType"] = request.PictureSearchType
	}

	if !dara.IsNil(request.SearchPicUrl) {
		query["SearchPicUrl"] = request.SearchPicUrl
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePictureSearchJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePictureSearchJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to download cloud recordings by time.
//
// Description:
//
// ## Limits
//
// - Maximum duration for a single time-based download: 30 minutes.
//
// - Download time precision: The precision of the download time range is within 10 seconds. For example, if you select a 20-second data segment to download, up to 30 seconds of data may be downloaded.
//
// - Maximum number of recording files for a single time-based download: 100.
//
// - Download tasks and files are retained for up to 1 day. You can re-download them after they expire.
//
// - Maximum concurrent downloads per instance: 10. Completed downloads are not counted.
//
// - If the audio and video encoding parameters change, the first half of the recording is downloaded by default.
//
// - If the download time ranges of recordings overlap, the cloud deduplicates the overlapping data.
//
// ## QPS limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to call this API operation is 50.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the account.
//
// @param request - CreateRecordDownloadByTimeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecordDownloadByTimeJobResponse
func (client *Client) CreateRecordDownloadByTimeJobWithContext(ctx context.Context, request *CreateRecordDownloadByTimeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecordDownloadByTimeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.RecordType) {
		query["RecordType"] = request.RecordType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecordDownloadByTimeJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecordDownloadByTimeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a recording plan based on a time template.
//
// @param request - CreateRecordPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecordPlanResponse
func (client *Client) CreateRecordPlanWithContext(ctx context.Context, request *CreateRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateRecordPlanResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecordPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecordPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateRtmpDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRtmpDeviceResponse
func (client *Client) CreateRtmpDeviceWithContext(ctx context.Context, request *CreateRtmpDeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateRtmpDeviceResponse, _err error) {
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

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.PullAuthKey) {
		query["PullAuthKey"] = request.PullAuthKey
	}

	if !dara.IsNil(request.PullKeyExpireTime) {
		query["PullKeyExpireTime"] = request.PullKeyExpireTime
	}

	if !dara.IsNil(request.PushAuthKey) {
		query["PushAuthKey"] = request.PushAuthKey
	}

	if !dara.IsNil(request.PushKeyExpireTime) {
		query["PushKeyExpireTime"] = request.PushKeyExpireTime
	}

	if !dara.IsNil(request.SubStreamName) {
		query["SubStreamName"] = request.SubStreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRtmpDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRtmpDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a video stream pushing job.
//
// Description:
//
// ## QPS limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 50.
//
// > All Resource Access Management (RAM) users under an Alibaba Cloud account share the quota of the account.
//
// @param request - CreateStreamPushJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamPushJobResponse
func (client *Client) CreateStreamPushJobWithContext(ctx context.Context, request *CreateStreamPushJobRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamPushJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.PushUrl) {
		query["PushUrl"] = request.PushUrl
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamPushJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamPushJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud snapshot task.
//
// Description:
//
// ## QPS limit
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this operation is 50.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - CreateStreamSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamSnapshotJobResponse
func (client *Client) CreateStreamSnapshotJobWithContext(ctx context.Context, request *CreateStreamSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamSnapshotJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.SnapshotInterval) {
		query["SnapshotInterval"] = request.SnapshotInterval
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamSnapshotJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamSnapshotJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a time template.
//
// @param request - CreateTimeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTimeTemplateResponse
func (client *Client) CreateTimeTemplateWithContext(ctx context.Context, request *CreateTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTimeTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllDay) {
		query["AllDay"] = request.AllDay
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TimeSections) {
		query["TimeSections"] = request.TimeSections
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTimeTemplate"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTimeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an intelligent event recording plan.
//
// @param request - DeleteEventRecordPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRecordPlanResponse
func (client *Client) DeleteEventRecordPlanWithContext(ctx context.Context, request *DeleteEventRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRecordPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventRecordPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventRecordPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the IP cameras associated with an intelligent alert event recording plan.
//
// @param request - DeleteEventRecordPlanDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRecordPlanDeviceResponse
func (client *Client) DeleteEventRecordPlanDeviceWithContext(ctx context.Context, request *DeleteEventRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRecordPlanDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventRecordPlanDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventRecordPlanDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a device group.
//
// @param request - DeleteFaceDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceDeviceGroupResponse
func (client *Client) DeleteFaceDeviceGroupWithContext(ctx context.Context, request *DeleteFaceDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceDeviceGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceDeviceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user from the face management service.
//
// @param request - DeleteFaceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceUserResponse
func (client *Client) DeleteFaceUserWithContext(ctx context.Context, request *DeleteFaceUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceUser"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified user group.
//
// Description:
//
// ## Limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this operation is 100.
//
// > Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - DeleteFaceUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceUserGroupResponse
func (client *Client) DeleteFaceUserGroupWithContext(ctx context.Context, request *DeleteFaceUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceUserGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the face synchronization relationship between a user group and a device group.
//
// Description:
//
// ## Before you begin
//
// The maximum number of queries per second (QPS) per Alibaba Cloud account for invoking this operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - DeleteFaceUserGroupAndDeviceGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelationWithContext(ctx context.Context, request *DeleteFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlId) {
		query["ControlId"] = request.ControlId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceUserGroupAndDeviceGroupRelation"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified face image from a user in the face library.
//
// Description:
//
// ## Limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this operation is 100.
//
// > Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - DeleteFaceUserPictureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceUserPictureResponse
func (client *Client) DeleteFaceUserPictureWithContext(ctx context.Context, request *DeleteFaceUserPictureRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserPictureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FacePicMd5) {
		query["FacePicMd5"] = request.FacePicMd5
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceUserPicture"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceUserPictureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteGbDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGbDeviceResponse
func (client *Client) DeleteGbDeviceWithContext(ctx context.Context, request *DeleteGbDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGbDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGbDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGbDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a local recording upload task.
//
// Description:
//
// After you stop a local recording upload task, the recording files that have been created are not deleted.
//
// @param request - DeleteLocalFileUploadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLocalFileUploadJobResponse
func (client *Client) DeleteLocalFileUploadJobWithContext(ctx context.Context, request *DeleteLocalFileUploadJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteLocalFileUploadJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLocalFileUploadJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLocalFileUploadJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple cloud pictures with specified picture IDs for a specified IPC device.
//
// @param request - DeletePictureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePictureResponse
func (client *Client) DeletePictureWithContext(ctx context.Context, request *DeletePictureRequest, runtime *dara.RuntimeOptions) (_result *DeletePictureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PictureIdList) {
		query["PictureIdList"] = request.PictureIdList
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePicture"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePictureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes video files stored in the cloud.
//
// @param request - DeleteRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordResponse
func (client *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.FileNameList) {
		query["FileNameList"] = request.FileNameList
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecord"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a recording plan.
//
// @param request - DeleteRecordPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordPlanResponse
func (client *Client) DeleteRecordPlanWithContext(ctx context.Context, request *DeleteRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecordPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecordPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the recording plan associated with an IPC device.
//
// @param request - DeleteRecordPlanDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordPlanDeviceResponse
func (client *Client) DeleteRecordPlanDeviceWithContext(ctx context.Context, request *DeleteRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordPlanDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecordPlanDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecordPlanDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRtmpDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRtmpDeviceResponse
func (client *Client) DeleteRtmpDeviceWithContext(ctx context.Context, request *DeleteRtmpDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteRtmpDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRtmpDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRtmpDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRtmpKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRtmpKeyResponse
func (client *Client) DeleteRtmpKeyWithContext(ctx context.Context, request *DeleteRtmpKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteRtmpKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRtmpKey"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRtmpKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a video stream pushing job.
//
// Description:
//
// ## QPS limit
//
// A single Alibaba Cloud account can invoke this operation up to 50 times per second (QPS).
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of that account.
//
// @param request - DeleteStreamPushJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStreamPushJobResponse
func (client *Client) DeleteStreamPushJobWithContext(ctx context.Context, request *DeleteStreamPushJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamPushJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStreamPushJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStreamPushJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a video snapshot task.
//
// Description:
//
// ## QPS limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 50.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - DeleteStreamSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStreamSnapshotJobResponse
func (client *Client) DeleteStreamSnapshotJobWithContext(ctx context.Context, request *DeleteStreamSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamSnapshotJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStreamSnapshotJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStreamSnapshotJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a time template.
//
// @param request - DeleteTimeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTimeTemplateResponse
func (client *Client) DeleteTimeTemplateWithContext(ctx context.Context, request *DeleteTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTimeTemplateResponse, _err error) {
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
		Action:      dara.String("DeleteTimeTemplate"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTimeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableGbSubDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableGbSubDeviceResponse
func (client *Client) EnableGbSubDeviceWithContext(ctx context.Context, request *EnableGbSubDeviceRequest, runtime *dara.RuntimeOptions) (_result *EnableGbSubDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.SubDeviceId) {
		query["SubDeviceId"] = request.SubDeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableGbSubDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableGbSubDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an integrated image search task.
//
// Description:
//
// Integrated image search retrieves both facial and body features during an image search.
//
// > This feature requires the deployment of face and body detection algorithms on V-series edge all-in-one machines.
//
// @param request - GetPictureSearchJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPictureSearchJobStatusResponse
func (client *Client) GetPictureSearchJobStatusWithContext(ctx context.Context, request *GetPictureSearchJobStatusRequest, runtime *dara.RuntimeOptions) (_result *GetPictureSearchJobStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPictureSearchJobStatus"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPictureSearchJobStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an image search application instance to search for images.
//
// @param request - PictureSearchPictureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PictureSearchPictureResponse
func (client *Client) PictureSearchPictureWithContext(ctx context.Context, request *PictureSearchPictureRequest, runtime *dara.RuntimeOptions) (_result *PictureSearchPictureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.ContainPicUrl) {
		query["ContainPicUrl"] = request.ContainPicUrl
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PictureSearchType) {
		query["PictureSearchType"] = request.PictureSearchType
	}

	if !dara.IsNil(request.SearchPicUrl) {
		query["SearchPicUrl"] = request.SearchPicUrl
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PictureSearchPicture"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PictureSearchPictureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the intelligent alert events for vehicle operations.
//
// Description:
//
// ## Usage notes
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to call this API operation is 10.
//
// > Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - QueryCarProcessEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCarProcessEventsResponse
func (client *Client) QueryCarProcessEventsWithContext(ctx context.Context, request *QueryCarProcessEventsRequest, runtime *dara.RuntimeOptions) (_result *QueryCarProcessEventsResponse, _err error) {
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

	if !dara.IsNil(request.AreaIndex) {
		query["AreaIndex"] = request.AreaIndex
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlateNo) {
		query["PlateNo"] = request.PlateNo
	}

	if !dara.IsNil(request.SubDeviceName) {
		query["SubDeviceName"] = request.SubDeviceName
	}

	if !dara.IsNil(request.SubIotId) {
		query["SubIotId"] = request.SubIotId
	}

	if !dara.IsNil(request.SubProductKey) {
		query["SubProductKey"] = request.SubProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCarProcessEvents"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCarProcessEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of intelligent alert events for an IPC device.
//
// @param request - QueryDeviceEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceEventResponse
func (client *Client) QueryDeviceEventWithContext(ctx context.Context, request *QueryDeviceEventRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceEvent"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the image captured and uploaded by an IPC device when an intelligent alert event is triggered.
//
// @param request - QueryDeviceEventPictureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceEventPictureResponse
func (client *Client) QueryDeviceEventPictureWithContext(ctx context.Context, request *QueryDeviceEventPictureRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceEventPictureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceEventPicture"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceEventPictureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the URL of the recording file for a specified intelligent alert event of an IP camera (IPC) device.
//
// @param request - QueryDeviceEventRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceEventRecordResponse
func (client *Client) QueryDeviceEventRecordWithContext(ctx context.Context, request *QueryDeviceEventRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceEventRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceEventRecord"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceEventRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves image URLs based on a list of image IDs.
//
// @param request - QueryDevicePictureByListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicePictureByListResponse
func (client *Client) QueryDevicePictureByListWithContext(ctx context.Context, request *QueryDevicePictureByListRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicePictureByListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PictureIdList) {
		query["PictureIdList"] = request.PictureIdList
	}

	if !dara.IsNil(request.PictureType) {
		query["PictureType"] = request.PictureType
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.ThumbWidth) {
		query["ThumbWidth"] = request.ThumbWidth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDevicePictureByList"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDevicePictureByListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about images captured by an IPC device.
//
// @param request - QueryDevicePictureFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicePictureFileResponse
func (client *Client) QueryDevicePictureFileWithContext(ctx context.Context, request *QueryDevicePictureFileRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicePictureFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaptureId) {
		query["CaptureId"] = request.CaptureId
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PictureType) {
		query["PictureType"] = request.PictureType
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.ThumbWidth) {
		query["ThumbWidth"] = request.ThumbWidth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDevicePictureFile"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDevicePictureFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud storage image lifecycle for an IPC device.
//
// @param request - QueryDevicePictureLifeCycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicePictureLifeCycleResponse
func (client *Client) QueryDevicePictureLifeCycleWithContext(ctx context.Context, request *QueryDevicePictureLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicePictureLifeCycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDevicePictureLifeCycle"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDevicePictureLifeCycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the cloud storage recording lifecycle of an IPC device.
//
// @param request - QueryDeviceRecordLifeCycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceRecordLifeCycleResponse
func (client *Client) QueryDeviceRecordLifeCycleWithContext(ctx context.Context, request *QueryDeviceRecordLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceRecordLifeCycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceList) {
		query["DeviceList"] = request.DeviceList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceRecordLifeCycle"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceRecordLifeCycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the URL of a recording file with a specified file name from the storage card of a local network video recorder or IP camera.
//
// @param request - QueryDeviceVodUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceVodUrlResponse
func (client *Client) QueryDeviceVodUrlWithContext(ctx context.Context, request *QueryDeviceVodUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceVodUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EnableStun) {
		query["EnableStun"] = request.EnableStun
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PlayUnLimited) {
		query["PlayUnLimited"] = request.PlayUnLimited
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.Scheme) {
		query["Scheme"] = request.Scheme
	}

	if !dara.IsNil(request.SeekTime) {
		query["SeekTime"] = request.SeekTime
	}

	if !dara.IsNil(request.ShouldEncrypt) {
		query["ShouldEncrypt"] = request.ShouldEncrypt
	}

	if !dara.IsNil(request.UrlValidDuration) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceVodUrl"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceVodUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the playback URL of a recording file within a specified time range from the storage card of a local network video recorder (NVR) or IP camera (IPC) device.
//
// Description:
//
// ## Limits
//
// This operation is supported only in scenarios where IPC devices are directly connected to IoT Platform.
//
// @param request - QueryDeviceVodUrlByTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceVodUrlByTimeResponse
func (client *Client) QueryDeviceVodUrlByTimeWithContext(ctx context.Context, request *QueryDeviceVodUrlByTimeRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceVodUrlByTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EnableStun) {
		query["EnableStun"] = request.EnableStun
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PlayUnLimited) {
		query["PlayUnLimited"] = request.PlayUnLimited
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.Scheme) {
		query["Scheme"] = request.Scheme
	}

	if !dara.IsNil(request.SeekTime) {
		query["SeekTime"] = request.SeekTime
	}

	if !dara.IsNil(request.ShouldEncrypt) {
		query["ShouldEncrypt"] = request.ShouldEncrypt
	}

	if !dara.IsNil(request.UrlValidDuration) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceVodUrlByTime"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceVodUrlByTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an intelligent alert event recording plan.
//
// @param request - QueryEventRecordPlanDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEventRecordPlanDetailResponse
func (client *Client) QueryEventRecordPlanDetailWithContext(ctx context.Context, request *QueryEventRecordPlanDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlanDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEventRecordPlanDetail"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEventRecordPlanDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the intelligent alert event recording plan of an IPC device by device ID.
//
// @param request - QueryEventRecordPlanDeviceByDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEventRecordPlanDeviceByDeviceResponse
func (client *Client) QueryEventRecordPlanDeviceByDeviceWithContext(ctx context.Context, request *QueryEventRecordPlanDeviceByDeviceRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlanDeviceByDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEventRecordPlanDeviceByDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEventRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cameras associated with an intelligent alert event recording plan.
//
// @param request - QueryEventRecordPlanDeviceByPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEventRecordPlanDeviceByPlanResponse
func (client *Client) QueryEventRecordPlanDeviceByPlanWithContext(ctx context.Context, request *QueryEventRecordPlanDeviceByPlanRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlanDeviceByPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEventRecordPlanDeviceByPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEventRecordPlanDeviceByPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of intelligent alert event recording plans.
//
// @param request - QueryEventRecordPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEventRecordPlansResponse
func (client *Client) QueryEventRecordPlansWithContext(ctx context.Context, request *QueryEventRecordPlansRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEventRecordPlans"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEventRecordPlansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all device groups.
//
// @param request - QueryFaceAllDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceAllDeviceGroupResponse
func (client *Client) QueryFaceAllDeviceGroupWithContext(ctx context.Context, request *QueryFaceAllDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
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
		Action:      dara.String("QueryFaceAllDeviceGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceAllDeviceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all user groups.
//
// @param request - QueryFaceAllUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceAllUserGroupResponse
func (client *Client) QueryFaceAllUserGroupWithContext(ctx context.Context, request *QueryFaceAllUserGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
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
		Action:      dara.String("QueryFaceAllUserGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceAllUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user groups and device groups with face synchronization relationships.
//
// @param request - QueryFaceAllUserGroupAndDeviceGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceAllUserGroupAndDeviceGroupRelationResponse
func (client *Client) QueryFaceAllUserGroupAndDeviceGroupRelationWithContext(ctx context.Context, request *QueryFaceAllUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllUserGroupAndDeviceGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
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
		Action:      dara.String("QueryFaceAllUserGroupAndDeviceGroupRelation"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceAllUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all users in a user group.
//
// @param request - QueryFaceAllUserIdsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceAllUserIdsByGroupIdResponse
func (client *Client) QueryFaceAllUserIdsByGroupIdWithContext(ctx context.Context, request *QueryFaceAllUserIdsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllUserIdsByGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceAllUserIdsByGroupId"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceAllUserIdsByGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user ID in the caller\\"s service based on the user ID from the face image library synchronization service.
//
// @param request - QueryFaceCustomUserIdByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceCustomUserIdByUserIdResponse
func (client *Client) QueryFaceCustomUserIdByUserIdWithContext(ctx context.Context, request *QueryFaceCustomUserIdByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceCustomUserIdByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceCustomUserIdByUserId"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceCustomUserIdByUserIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the device groups to which a device belongs.
//
// @param request - QueryFaceDeviceGroupsByDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceDeviceGroupsByDeviceResponse
func (client *Client) QueryFaceDeviceGroupsByDeviceWithContext(ctx context.Context, request *QueryFaceDeviceGroupsByDeviceRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceDeviceGroupsByDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceDeviceGroupsByDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceDeviceGroupsByDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a user.
//
// @param request - QueryFaceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceUserResponse
func (client *Client) QueryFaceUserWithContext(ctx context.Context, request *QueryFaceUserRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceUser"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user information in batches based on a list of user IDs.
//
// Description:
//
// ## Usage notes
//
// - The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - QueryFaceUserBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceUserBatchResponse
func (client *Client) QueryFaceUserBatchWithContext(ctx context.Context, request *QueryFaceUserBatchRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserIdList) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceUserBatch"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceUserBatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries face data by face name.
//
// Description:
//
// ## Usage notes
//
// - The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - QueryFaceUserByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceUserByNameResponse
func (client *Client) QueryFaceUserByNameWithContext(ctx context.Context, request *QueryFaceUserByNameRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceUserByName"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceUserByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user groups to which a user belongs.
//
// @param request - QueryFaceUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceUserGroupResponse
func (client *Client) QueryFaceUserGroupWithContext(ctx context.Context, request *QueryFaceUserGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceUserGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the face synchronization relationships between user groups and device groups.
//
// @param request - QueryFaceUserGroupAndDeviceGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) QueryFaceUserGroupAndDeviceGroupRelationWithContext(ctx context.Context, request *QueryFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlId) {
		query["ControlId"] = request.ControlId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceUserGroupAndDeviceGroupRelation"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user ID in the face library synchronization service based on the user ID in the caller system.
//
// @param request - QueryFaceUserIdByCustomUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceUserIdByCustomUserIdResponse
func (client *Client) QueryFaceUserIdByCustomUserIdWithContext(ctx context.Context, request *QueryFaceUserIdByCustomUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserIdByCustomUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomUserId) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceUserIdByCustomUserId"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceUserIdByCustomUserIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryGbSubDeviceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGbSubDeviceListResponse
func (client *Client) QueryGbSubDeviceListWithContext(ctx context.Context, request *QueryGbSubDeviceListRequest, runtime *dara.RuntimeOptions) (_result *QueryGbSubDeviceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PageStart) {
		query["PageStart"] = request.PageStart
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGbSubDeviceList"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGbSubDeviceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the live streaming URL of an IPC device.
//
// @param request - QueryLiveStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveStreamingResponse
func (client *Client) QueryLiveStreamingWithContext(ctx context.Context, request *QueryLiveStreamingRequest, runtime *dara.RuntimeOptions) (_result *QueryLiveStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheDuration) {
		query["CacheDuration"] = request.CacheDuration
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EnableStun) {
		query["EnableStun"] = request.EnableStun
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.ForceIFrame) {
		query["ForceIFrame"] = request.ForceIFrame
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PlayUnLimited) {
		query["PlayUnLimited"] = request.PlayUnLimited
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.Scheme) {
		query["Scheme"] = request.Scheme
	}

	if !dara.IsNil(request.ShouldEncrypt) {
		query["ShouldEncrypt"] = request.ShouldEncrypt
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.UrlValidDuration) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLiveStreaming"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLiveStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task result of a local video recording upload node.
//
// @param request - QueryLocalFileUploadJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLocalFileUploadJobResponse
func (client *Client) QueryLocalFileUploadJobWithContext(ctx context.Context, request *QueryLocalFileUploadJobRequest, runtime *dara.RuntimeOptions) (_result *QueryLocalFileUploadJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLocalFileUploadJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLocalFileUploadJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether an IP camera has video recordings for each day of a specified month.
//
// @param request - QueryMonthRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMonthRecordResponse
func (client *Client) QueryMonthRecordWithContext(ctx context.Context, request *QueryMonthRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryMonthRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.Month) {
		query["Month"] = request.Month
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMonthRecord"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMonthRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of images.
//
// Description:
//
// ## QPS limit
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 500.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - QueryPictureFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPictureFilesResponse
func (client *Client) QueryPictureFilesWithContext(ctx context.Context, request *QueryPictureFilesRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PictureSource) {
		query["PictureSource"] = request.PictureSource
	}

	if !dara.IsNil(request.PictureType) {
		query["PictureType"] = request.PictureType
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPictureFiles"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPictureFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of V-series edge all-in-one machines bound to an image search application instance.
//
// @param request - QueryPictureSearchAiboxesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPictureSearchAiboxesResponse
func (client *Client) QueryPictureSearchAiboxesWithContext(ctx context.Context, request *QueryPictureSearchAiboxesRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchAiboxesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPictureSearchAiboxes"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPictureSearchAiboxesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of created image search application instances.
//
// Description:
//
// ## Limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
// > Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - QueryPictureSearchAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPictureSearchAppsResponse
func (client *Client) QueryPictureSearchAppsWithContext(ctx context.Context, request *QueryPictureSearchAppsRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPictureSearchApps"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPictureSearchAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of devices bound to an image search application instance.
//
// @param request - QueryPictureSearchDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPictureSearchDevicesResponse
func (client *Client) QueryPictureSearchDevicesWithContext(ctx context.Context, request *QueryPictureSearchDevicesRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPictureSearchDevices"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPictureSearchDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of integrated image search jobs.
//
// Description:
//
// Integrated image search means that both faces and human bodies are retrieved simultaneously during an image search.
//
// > This feature requires the deployment of face and human body detection algorithms on V-series edge all-in-one machines.
//
// @param request - QueryPictureSearchJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPictureSearchJobResponse
func (client *Client) QueryPictureSearchJobWithContext(ctx context.Context, request *QueryPictureSearchJobRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.JobStatus) {
		query["JobStatus"] = request.JobStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPictureSearchJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPictureSearchJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the image retrieval results of an integrated image search job.
//
// @param request - QueryPictureSearchJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPictureSearchJobResultResponse
func (client *Client) QueryPictureSearchJobResultWithContext(ctx context.Context, request *QueryPictureSearchJobResultRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchJobResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPictureSearchJobResult"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPictureSearchJobResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cloud video recordings.
//
// Description:
//
// ## QPS limit
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 250.
//
// > All Resource Access Management (RAM) users within a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - QueryRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordResponse
func (client *Client) QueryRecordWithContext(ctx context.Context, request *QueryRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.NeedSnapshot) {
		query["NeedSnapshot"] = request.NeedSnapshot
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.RecordType) {
		query["RecordType"] = request.RecordType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecord"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the URL of a recording file by recording file ID.
//
// @param request - QueryRecordByRecordIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordByRecordIdResponse
func (client *Client) QueryRecordByRecordIdWithContext(ctx context.Context, request *QueryRecordByRecordIdRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordByRecordIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordByRecordId"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordByRecordIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cloud record download task by task ID, including the start time, recording progress, recording type, and download URL.
//
// Description:
//
// ## Limits
//
// The obtained recording download URL is valid for 1 hour.
//
// ## QPS limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 50.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - QueryRecordDownloadJobByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordDownloadJobByIdResponse
func (client *Client) QueryRecordDownloadJobByIdWithContext(ctx context.Context, request *QueryRecordDownloadJobByIdRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordDownloadJobByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordDownloadJobById"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordDownloadJobByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cloud recording download tasks.
//
// Description:
//
// ## Limits
//
// The obtained recording download URL is valid for 1 hour.
//
// ## QPS limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 50.
//
// > All Resource Access Management (RAM) users within a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - QueryRecordDownloadJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordDownloadJobListResponse
func (client *Client) QueryRecordDownloadJobListWithContext(ctx context.Context, request *QueryRecordDownloadJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordDownloadJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordDownloadJobList"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordDownloadJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Requests the download URL of a cloud recording file.
//
// Description:
//
// ## Usage notes
//
// Generating a cloud recording file takes a long time. You must poll this operation to obtain the progress. The download URL is generated only after the progress is complete.
//
// @param request - QueryRecordDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordDownloadUrlResponse
func (client *Client) QueryRecordDownloadUrlWithContext(ctx context.Context, request *QueryRecordDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordDownloadUrl"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordDownloadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a recording plan.
//
// @param request - QueryRecordPlanDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordPlanDetailResponse
func (client *Client) QueryRecordPlanDetailWithContext(ctx context.Context, request *QueryRecordPlanDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlanDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordPlanDetail"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordPlanDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recording plan of a video device based on the device ID of the IP camera (IPC).
//
// @param request - QueryRecordPlanDeviceByDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordPlanDeviceByDeviceResponse
func (client *Client) QueryRecordPlanDeviceByDeviceWithContext(ctx context.Context, request *QueryRecordPlanDeviceByDeviceRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlanDeviceByDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordPlanDeviceByDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the devices associated with a recording plan.
//
// @param request - QueryRecordPlanDeviceByPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordPlanDeviceByPlanResponse
func (client *Client) QueryRecordPlanDeviceByPlanWithContext(ctx context.Context, request *QueryRecordPlanDeviceByPlanRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlanDeviceByPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordPlanDeviceByPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordPlanDeviceByPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of recording plans.
//
// @param request - QueryRecordPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordPlansResponse
func (client *Client) QueryRecordPlansWithContext(ctx context.Context, request *QueryRecordPlansRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordPlans"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordPlansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the playback URL of a cloud recording file.
//
// @param request - QueryRecordUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordUrlResponse
func (client *Client) QueryRecordUrlWithContext(ctx context.Context, request *QueryRecordUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.SeekTime) {
		query["SeekTime"] = request.SeekTime
	}

	if !dara.IsNil(request.UrlValidDuration) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordUrl"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the playback URL for recording files within a specified time range.
//
// @param request - QueryRecordUrlByTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordUrlByTimeResponse
func (client *Client) QueryRecordUrlByTimeWithContext(ctx context.Context, request *QueryRecordUrlByTimeRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordUrlByTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.UrlValidDuration) {
		query["UrlValidDuration"] = request.UrlValidDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordUrlByTime"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordUrlByTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryRtmpKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRtmpKeyResponse
func (client *Client) QueryRtmpKeyWithContext(ctx context.Context, request *QueryRtmpKeyRequest, runtime *dara.RuntimeOptions) (_result *QueryRtmpKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRtmpKey"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRtmpKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a video stream push job, such as the creation time, current push status, job type, and push URL.
//
// Description:
//
// ## QPS limit
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to call this API operation is 50.
//
// > All Resource Access Management (RAM) users within a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - QueryStreamPushJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStreamPushJobResponse
func (client *Client) QueryStreamPushJobWithContext(ctx context.Context, request *QueryStreamPushJobRequest, runtime *dara.RuntimeOptions) (_result *QueryStreamPushJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryStreamPushJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryStreamPushJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list and details of video stream push jobs, including the creation time, current push status, job type, and push URL.
//
// Description:
//
// ## QPS limit
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this operation is 50.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - QueryStreamPushJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStreamPushJobListResponse
func (client *Client) QueryStreamPushJobListWithContext(ctx context.Context, request *QueryStreamPushJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryStreamPushJobListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryStreamPushJobList"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryStreamPushJobListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list and details of cloud snapshot tasks, including the snapshot interval and stream type.
//
// Description:
//
// ## QPS limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 50.
//
// > All Resource Access Management (RAM) users under a single Alibaba Cloud account share the quota of the Alibaba Cloud account.
//
// @param request - QueryStreamSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStreamSnapshotJobResponse
func (client *Client) QueryStreamSnapshotJobWithContext(ctx context.Context, request *QueryStreamSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *QueryStreamSnapshotJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryStreamSnapshotJob"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryStreamSnapshotJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time template information of a device.
//
// @param request - QueryTimeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTimeTemplateResponse
func (client *Client) QueryTimeTemplateWithContext(ctx context.Context, request *QueryTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryTimeTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTimeTemplate"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTimeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified time template.
//
// @param request - QueryTimeTemplateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTimeTemplateDetailResponse
func (client *Client) QueryTimeTemplateDetailWithContext(ctx context.Context, request *QueryTimeTemplateDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryTimeTemplateDetailResponse, _err error) {
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
		Action:      dara.String("QueryTimeTemplateDetail"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTimeTemplateDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryVisionDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVisionDeviceInfoResponse
func (client *Client) QueryVisionDeviceInfoWithContext(ctx context.Context, request *QueryVisionDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryVisionDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVisionDeviceInfo"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVisionDeviceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the voice intercom address of an IPC device.
//
// @param request - QueryVoiceIntercomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVoiceIntercomResponse
func (client *Client) QueryVoiceIntercomWithContext(ctx context.Context, request *QueryVoiceIntercomRequest, runtime *dara.RuntimeOptions) (_result *QueryVoiceIntercomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.Scheme) {
		query["Scheme"] = request.Scheme
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVoiceIntercom"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVoiceIntercomResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshGbSubDeviceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshGbSubDeviceListResponse
func (client *Client) RefreshGbSubDeviceListWithContext(ctx context.Context, request *RefreshGbSubDeviceListRequest, runtime *dara.RuntimeOptions) (_result *RefreshGbSubDeviceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshGbSubDeviceList"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshGbSubDeviceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a device from a device group.
//
// @param request - RemoveFaceDeviceFromDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFaceDeviceFromDeviceGroupResponse
func (client *Client) RemoveFaceDeviceFromDeviceGroupWithContext(ctx context.Context, request *RemoveFaceDeviceFromDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveFaceDeviceFromDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveFaceDeviceFromDeviceGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveFaceDeviceFromDeviceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a user from a user group.
//
// @param request - RemoveFaceUserFromUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFaceUserFromUserGroupResponse
func (client *Client) RemoveFaceUserFromUserGroupWithContext(ctx context.Context, request *RemoveFaceUserFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveFaceUserFromUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveFaceUserFromUserGroup"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveFaceUserFromUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the cloud storage image lifecycle for an IPC device.
//
// Description:
//
// > By default, images are permanently stored in the cloud before you configure the cloud storage image lifecycle.
//
// @param request - SetDevicePictureLifeCycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDevicePictureLifeCycleResponse
func (client *Client) SetDevicePictureLifeCycleWithContext(ctx context.Context, request *SetDevicePictureLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *SetDevicePictureLifeCycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Day) {
		query["Day"] = request.Day
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDevicePictureLifeCycle"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDevicePictureLifeCycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the recording lifecycle of an IPC device, which specifies the number of days that recording files are stored in the cloud.
//
// @param request - SetDeviceRecordLifeCycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeviceRecordLifeCycleResponse
func (client *Client) SetDeviceRecordLifeCycleWithContext(ctx context.Context, request *SetDeviceRecordLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *SetDeviceRecordLifeCycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Day) {
		query["Day"] = request.Day
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDeviceRecordLifeCycle"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDeviceRecordLifeCycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the live streaming of an IPC device.
//
// Description:
//
// ## Usage notes
//
// - Before you invoke this operation, you must implement the Thing Specification Language (TSL) model service for stopping live streaming with the identity **StopPushStreaming*	- on the IPC device. If this service is not implemented on the IPC device, or the IPC device does not return a normal response after receiving the service invocation, the IPC device will not stop pushing the stream. For more information about TSL models, see [What is a TSL model](https://help.aliyun.com/document_detail/73727.html).
//
// - If no one is watching the stream on the IPC device, you do not need to invoke this operation. IoT Platform automatically triggers the device to stop pushing the stream.
//
// - During the procedure of triggering a video recording event, invoking this operation causes interruptions to the video file. Use this operation with caution.
//
// @param request - StopLiveStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLiveStreamingResponse
func (client *Client) StopLiveStreamingWithContext(ctx context.Context, request *StopLiveStreamingRequest, runtime *dara.RuntimeOptions) (_result *StopLiveStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLiveStreaming"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLiveStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an IP camera from actively triggering video recording.
//
// @param request - StopTriggeredRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTriggeredRecordResponse
func (client *Client) StopTriggeredRecordWithContext(ctx context.Context, request *StopTriggeredRecordRequest, runtime *dara.RuntimeOptions) (_result *StopTriggeredRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTriggeredRecord"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTriggeredRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers devices to another instance.
//
// Description:
//
// Before you call this operation to transfer devices to another instance, contact technical support to evaluate the migration plan and prevent features from becoming unavailable after the migration.
//
// **Before you begin**
//
// - The devices to be transferred have the video service feature enabled.
//
//   - Devices in the Enterprise instance: Log on to the IoT Platform console and click **Enterprise instance**. In the left-side navigation pane, choose **Video Service*	- > **Access Service**. On the **Access Service*	- page, click the **Product Management*	- tab. Check the product to which the device to be transferred belongs, and make sure that the **Enabled/Disabled Status*	- column displays **Enabled**.
//
//   - Devices in an Enterprise Edition instance: The video service is enabled.
//
// - The device distribution task is complete. For more information, see [CreateDeviceDistributeJob](https://help.aliyun.com/document_detail/199390.html) or [Device distribution](https://help.aliyun.com/document_detail/143450.html).
//
// - Make sure that a picture lifecycle is configured for camera devices. If no picture lifecycle is configured for a camera device, calling this operation automatically sets the picture lifecycle to 180 days. For more information, see [SetDevicePictureLifeCycle](https://help.aliyun.com/document_detail/107880.html).
//
// **Background information**
//
// - If the source instance and target instance are of different types, review the Billable methods of the instances before the migration. For more information about Billable methods, see [Billing overview](https://help.aliyun.com/document_detail/124250.html).
//
// - During the migration, devices go offline. The more devices you migrate, the longer the migration takes. Data generated by the devices during this period cannot be retained.
//
// @param request - TransferDeviceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferDeviceInstanceResponse
func (client *Client) TransferDeviceInstanceWithContext(ctx context.Context, request *TransferDeviceInstanceRequest, runtime *dara.RuntimeOptions) (_result *TransferDeviceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceNameList) {
		query["DeviceNameList"] = request.DeviceNameList
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferDeviceInstance"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferDeviceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers an IPC device to capture an image and upload it to the cloud.
//
// @param request - TriggerCapturePictureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerCapturePictureResponse
func (client *Client) TriggerCapturePictureWithContext(ctx context.Context, request *TriggerCapturePictureRequest, runtime *dara.RuntimeOptions) (_result *TriggerCapturePictureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerCapturePicture"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerCapturePictureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers the specified IPC device to start recording and uploads the recording file to the cloud.
//
// @param request - TriggerRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerRecordResponse
func (client *Client) TriggerRecordWithContext(ctx context.Context, request *TriggerRecordRequest, runtime *dara.RuntimeOptions) (_result *TriggerRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.PreRecordDuration) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.RecordDuration) {
		query["RecordDuration"] = request.RecordDuration
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerRecord"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates multiple devices from an image search application instance.
//
// Description:
//
// ## Limits
//
// If you invoke this operation on an Enterprise instance, you must specify the **IotInstanceId*	- request parameter. Otherwise, the call is failed.
//
// @param request - UnbindPictureSearchAppWithDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPictureSearchAppWithDevicesResponse
func (client *Client) UnbindPictureSearchAppWithDevicesWithContext(ctx context.Context, request *UnbindPictureSearchAppWithDevicesRequest, runtime *dara.RuntimeOptions) (_result *UnbindPictureSearchAppWithDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.DeviceIotIds) {
		query["DeviceIotIds"] = request.DeviceIotIds
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindPictureSearchAppWithDevices"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the video recording plan for intelligent alert events.
//
// @param request - UpdateEventRecordPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventRecordPlanResponse
func (client *Client) UpdateEventRecordPlanWithContext(ctx context.Context, request *UpdateEventRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventRecordPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventTypes) {
		query["EventTypes"] = request.EventTypes
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.PreRecordDuration) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !dara.IsNil(request.RecordDuration) {
		query["RecordDuration"] = request.RecordDuration
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventRecordPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventRecordPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates user information.
//
// @param request - UpdateFaceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFaceUserResponse
func (client *Client) UpdateFaceUserWithContext(ctx context.Context, request *UpdateFaceUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateFaceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomUserId) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !dara.IsNil(request.FacePicUrl) {
		query["FacePicUrl"] = request.FacePicUrl
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFaceUser"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFaceUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the face synchronization relationship between a user group and a device group.
//
// Description:
//
// ## Limits
//
// The maximum queries per second (QPS) for a single Alibaba Cloud account to invoke this API operation is 100.
//
//	> Resource Access Management (RAM) users share the quota of the Alibaba Cloud account.
//
// @param request - UpdateFaceUserGroupAndDeviceGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) UpdateFaceUserGroupAndDeviceGroupRelationWithContext(ctx context.Context, request *UpdateFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *UpdateFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlId) {
		query["ControlId"] = request.ControlId
	}

	if !dara.IsNil(request.IsolationId) {
		query["IsolationId"] = request.IsolationId
	}

	if !dara.IsNil(request.Relation) {
		query["Relation"] = request.Relation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFaceUserGroupAndDeviceGroupRelation"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateGbDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGbDeviceResponse
func (client *Client) UpdateGbDeviceWithContext(ctx context.Context, request *UpdateGbDeviceRequest, runtime *dara.RuntimeOptions) (_result *UpdateGbDeviceResponse, _err error) {
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

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGbDevice"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGbDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the IP protocol version supported by an instance.
//
// @param request - UpdateInstanceInternetProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceInternetProtocolResponse
func (client *Client) UpdateInstanceInternetProtocolWithContext(ctx context.Context, request *UpdateInstanceInternetProtocolRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceInternetProtocolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceInternetProtocol"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceInternetProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of an image search application instance.
//
// @param request - UpdatePictureSearchAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePictureSearchAppResponse
func (client *Client) UpdatePictureSearchAppWithContext(ctx context.Context, request *UpdatePictureSearchAppRequest, runtime *dara.RuntimeOptions) (_result *UpdatePictureSearchAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePictureSearchApp"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePictureSearchAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a recording plan.
//
// @param request - UpdateRecordPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordPlanResponse
func (client *Client) UpdateRecordPlanWithContext(ctx context.Context, request *UpdateRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordPlanResponse, _err error) {
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

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecordPlan"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecordPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRtmpKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRtmpKeyResponse
func (client *Client) UpdateRtmpKeyWithContext(ctx context.Context, request *UpdateRtmpKeyRequest, runtime *dara.RuntimeOptions) (_result *UpdateRtmpKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotInstanceId) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.PullAuthKey) {
		query["PullAuthKey"] = request.PullAuthKey
	}

	if !dara.IsNil(request.PullKeyExpireTime) {
		query["PullKeyExpireTime"] = request.PullKeyExpireTime
	}

	if !dara.IsNil(request.PushAuthKey) {
		query["PushAuthKey"] = request.PushAuthKey
	}

	if !dara.IsNil(request.PushKeyExpireTime) {
		query["PushKeyExpireTime"] = request.PushKeyExpireTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRtmpKey"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRtmpKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a time template.
//
// @param request - UpdateTimeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTimeTemplateResponse
func (client *Client) UpdateTimeTemplateWithContext(ctx context.Context, request *UpdateTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateTimeTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllDay) {
		query["AllDay"] = request.AllDay
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TimeSections) {
		query["TimeSections"] = request.TimeSections
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTimeTemplate"),
		Version:     dara.String("2018-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTimeTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
