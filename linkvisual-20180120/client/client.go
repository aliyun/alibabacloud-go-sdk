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
		"ap-northeast-1":              dara.String("linkvisual.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("linkvisual.aliyuncs.com"),
		"ap-south-1":                  dara.String("linkvisual.aliyuncs.com"),
		"ap-southeast-1":              dara.String("linkvisual.aliyuncs.com"),
		"ap-southeast-2":              dara.String("linkvisual.aliyuncs.com"),
		"ap-southeast-3":              dara.String("linkvisual.aliyuncs.com"),
		"ap-southeast-5":              dara.String("linkvisual.aliyuncs.com"),
		"cn-beijing":                  dara.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("linkvisual.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("linkvisual.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("linkvisual.aliyuncs.com"),
		"cn-chengdu":                  dara.String("linkvisual.aliyuncs.com"),
		"cn-edge-1":                   dara.String("linkvisual.aliyuncs.com"),
		"cn-fujian":                   dara.String("linkvisual.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("linkvisual.aliyuncs.com"),
		"cn-hongkong":                 dara.String("linkvisual.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("linkvisual.aliyuncs.com"),
		"cn-huhehaote":                dara.String("linkvisual.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("linkvisual.aliyuncs.com"),
		"cn-qingdao":                  dara.String("linkvisual.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("linkvisual.aliyuncs.com"),
		"cn-wuhan":                    dara.String("linkvisual.aliyuncs.com"),
		"cn-yushanfang":               dara.String("linkvisual.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("linkvisual.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("linkvisual.aliyuncs.com"),
		"eu-central-1":                dara.String("linkvisual.aliyuncs.com"),
		"eu-west-1":                   dara.String("linkvisual.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("linkvisual.aliyuncs.com"),
		"me-east-1":                   dara.String("linkvisual.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("linkvisual.aliyuncs.com"),
		"us-east-1":                   dara.String("linkvisual.aliyuncs.com"),
		"us-west-1":                   dara.String("linkvisual.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("linkvisual"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates an intelligent alert event recording plan for an IP camera.
//
// @param request - AddEventRecordPlanDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEventRecordPlanDeviceResponse
func (client *Client) AddEventRecordPlanDeviceWithOptions(request *AddEventRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *AddEventRecordPlanDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an intelligent alert event recording plan for an IP camera.
//
// @param request - AddEventRecordPlanDeviceRequest
//
// @return AddEventRecordPlanDeviceResponse
func (client *Client) AddEventRecordPlanDevice(request *AddEventRecordPlanDeviceRequest) (_result *AddEventRecordPlanDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddEventRecordPlanDeviceResponse{}
	_body, _err := client.AddEventRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFaceDeviceGroupWithOptions(request *AddFaceDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceDeviceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddFaceDeviceGroupResponse
func (client *Client) AddFaceDeviceGroup(request *AddFaceDeviceGroupRequest) (_result *AddFaceDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFaceDeviceGroupResponse{}
	_body, _err := client.AddFaceDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFaceDeviceToDeviceGroupWithOptions(request *AddFaceDeviceToDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceDeviceToDeviceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddFaceDeviceToDeviceGroupResponse
func (client *Client) AddFaceDeviceToDeviceGroup(request *AddFaceDeviceToDeviceGroupRequest) (_result *AddFaceDeviceToDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFaceDeviceToDeviceGroupResponse{}
	_body, _err := client.AddFaceDeviceToDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFaceUserWithOptions(request *AddFaceUserRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddFaceUserResponse
func (client *Client) AddFaceUser(request *AddFaceUserRequest) (_result *AddFaceUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFaceUserResponse{}
	_body, _err := client.AddFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFaceUserGroupWithOptions(request *AddFaceUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddFaceUserGroupResponse
func (client *Client) AddFaceUserGroup(request *AddFaceUserGroupRequest) (_result *AddFaceUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFaceUserGroupResponse{}
	_body, _err := client.AddFaceUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFaceUserGroupAndDeviceGroupRelationWithOptions(request *AddFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) AddFaceUserGroupAndDeviceGroupRelation(request *AddFaceUserGroupAndDeviceGroupRelationRequest) (_result *AddFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.AddFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFaceUserPictureWithOptions(request *AddFaceUserPictureRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserPictureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddFaceUserPictureResponse
func (client *Client) AddFaceUserPicture(request *AddFaceUserPictureRequest) (_result *AddFaceUserPictureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFaceUserPictureResponse{}
	_body, _err := client.AddFaceUserPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFaceUserToUserGroupWithOptions(request *AddFaceUserToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddFaceUserToUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddFaceUserToUserGroupResponse
func (client *Client) AddFaceUserToUserGroup(request *AddFaceUserToUserGroupRequest) (_result *AddFaceUserToUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFaceUserToUserGroupResponse{}
	_body, _err := client.AddFaceUserToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddRecordPlanDeviceWithOptions(request *AddRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *AddRecordPlanDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddRecordPlanDeviceResponse
func (client *Client) AddRecordPlanDevice(request *AddRecordPlanDeviceRequest) (_result *AddRecordPlanDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRecordPlanDeviceResponse{}
	_body, _err := client.AddRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BatchQueryVisionDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryVisionDeviceInfoResponse
func (client *Client) BatchQueryVisionDeviceInfoWithOptions(request *BatchQueryVisionDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *BatchQueryVisionDeviceInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchQueryVisionDeviceInfoRequest
//
// @return BatchQueryVisionDeviceInfoResponse
func (client *Client) BatchQueryVisionDeviceInfo(request *BatchQueryVisionDeviceInfoRequest) (_result *BatchQueryVisionDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchQueryVisionDeviceInfoResponse{}
	_body, _err := client.BatchQueryVisionDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BindPictureSearchAppWithDevicesWithOptions(request *BindPictureSearchAppWithDevicesRequest, runtime *dara.RuntimeOptions) (_result *BindPictureSearchAppWithDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BindPictureSearchAppWithDevicesResponse
func (client *Client) BindPictureSearchAppWithDevices(request *BindPictureSearchAppWithDevicesRequest) (_result *BindPictureSearchAppWithDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.BindPictureSearchAppWithDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckFaceUserDoExistOnDeviceWithOptions(request *CheckFaceUserDoExistOnDeviceRequest, runtime *dara.RuntimeOptions) (_result *CheckFaceUserDoExistOnDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckFaceUserDoExistOnDeviceResponse
func (client *Client) CheckFaceUserDoExistOnDevice(request *CheckFaceUserDoExistOnDeviceRequest) (_result *CheckFaceUserDoExistOnDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckFaceUserDoExistOnDeviceResponse{}
	_body, _err := client.CheckFaceUserDoExistOnDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ClearFaceDeviceDBWithOptions(request *ClearFaceDeviceDBRequest, runtime *dara.RuntimeOptions) (_result *ClearFaceDeviceDBResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ClearFaceDeviceDBResponse
func (client *Client) ClearFaceDeviceDB(request *ClearFaceDeviceDBRequest) (_result *ClearFaceDeviceDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClearFaceDeviceDBResponse{}
	_body, _err := client.ClearFaceDeviceDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateEventRecordPlanWithOptions(request *CreateEventRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateEventRecordPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateEventRecordPlanResponse
func (client *Client) CreateEventRecordPlan(request *CreateEventRecordPlanRequest) (_result *CreateEventRecordPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEventRecordPlanResponse{}
	_body, _err := client.CreateEventRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateGbDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGbDeviceResponse
func (client *Client) CreateGbDeviceWithOptions(request *CreateGbDeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateGbDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateGbDeviceRequest
//
// @return CreateGbDeviceResponse
func (client *Client) CreateGbDevice(request *CreateGbDeviceRequest) (_result *CreateGbDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGbDeviceResponse{}
	_body, _err := client.CreateGbDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateLocalFileUploadJobWithOptions(request *CreateLocalFileUploadJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLocalFileUploadJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateLocalFileUploadJobResponse
func (client *Client) CreateLocalFileUploadJob(request *CreateLocalFileUploadJobRequest) (_result *CreateLocalFileUploadJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLocalFileUploadJobResponse{}
	_body, _err := client.CreateLocalFileUploadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateLocalRecordDownloadByTimeJobWithOptions(request *CreateLocalRecordDownloadByTimeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLocalRecordDownloadByTimeJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateLocalRecordDownloadByTimeJobResponse
func (client *Client) CreateLocalRecordDownloadByTimeJob(request *CreateLocalRecordDownloadByTimeJobRequest) (_result *CreateLocalRecordDownloadByTimeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLocalRecordDownloadByTimeJobResponse{}
	_body, _err := client.CreateLocalRecordDownloadByTimeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePictureSearchAppWithOptions(request *CreatePictureSearchAppRequest, runtime *dara.RuntimeOptions) (_result *CreatePictureSearchAppResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreatePictureSearchAppResponse
func (client *Client) CreatePictureSearchApp(request *CreatePictureSearchAppRequest) (_result *CreatePictureSearchAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePictureSearchAppResponse{}
	_body, _err := client.CreatePictureSearchAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePictureSearchJobWithOptions(request *CreatePictureSearchJobRequest, runtime *dara.RuntimeOptions) (_result *CreatePictureSearchJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreatePictureSearchJobResponse
func (client *Client) CreatePictureSearchJob(request *CreatePictureSearchJobRequest) (_result *CreatePictureSearchJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePictureSearchJobResponse{}
	_body, _err := client.CreatePictureSearchJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRecordDownloadByTimeJobWithOptions(request *CreateRecordDownloadByTimeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecordDownloadByTimeJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRecordDownloadByTimeJobResponse
func (client *Client) CreateRecordDownloadByTimeJob(request *CreateRecordDownloadByTimeJobRequest) (_result *CreateRecordDownloadByTimeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecordDownloadByTimeJobResponse{}
	_body, _err := client.CreateRecordDownloadByTimeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRecordPlanWithOptions(request *CreateRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateRecordPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRecordPlanResponse
func (client *Client) CreateRecordPlan(request *CreateRecordPlanRequest) (_result *CreateRecordPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecordPlanResponse{}
	_body, _err := client.CreateRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateRtmpDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRtmpDeviceResponse
func (client *Client) CreateRtmpDeviceWithOptions(request *CreateRtmpDeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateRtmpDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateRtmpDeviceRequest
//
// @return CreateRtmpDeviceResponse
func (client *Client) CreateRtmpDevice(request *CreateRtmpDeviceRequest) (_result *CreateRtmpDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRtmpDeviceResponse{}
	_body, _err := client.CreateRtmpDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateStreamPushJobWithOptions(request *CreateStreamPushJobRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamPushJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateStreamPushJobResponse
func (client *Client) CreateStreamPushJob(request *CreateStreamPushJobRequest) (_result *CreateStreamPushJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStreamPushJobResponse{}
	_body, _err := client.CreateStreamPushJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateStreamSnapshotJobWithOptions(request *CreateStreamSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamSnapshotJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateStreamSnapshotJobResponse
func (client *Client) CreateStreamSnapshotJob(request *CreateStreamSnapshotJobRequest) (_result *CreateStreamSnapshotJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStreamSnapshotJobResponse{}
	_body, _err := client.CreateStreamSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTimeTemplateWithOptions(request *CreateTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTimeTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateTimeTemplateResponse
func (client *Client) CreateTimeTemplate(request *CreateTimeTemplateRequest) (_result *CreateTimeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTimeTemplateResponse{}
	_body, _err := client.CreateTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteEventRecordPlanWithOptions(request *DeleteEventRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRecordPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteEventRecordPlanResponse
func (client *Client) DeleteEventRecordPlan(request *DeleteEventRecordPlanRequest) (_result *DeleteEventRecordPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEventRecordPlanResponse{}
	_body, _err := client.DeleteEventRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteEventRecordPlanDeviceWithOptions(request *DeleteEventRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRecordPlanDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteEventRecordPlanDeviceResponse
func (client *Client) DeleteEventRecordPlanDevice(request *DeleteEventRecordPlanDeviceRequest) (_result *DeleteEventRecordPlanDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEventRecordPlanDeviceResponse{}
	_body, _err := client.DeleteEventRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFaceDeviceGroupWithOptions(request *DeleteFaceDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceDeviceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFaceDeviceGroupResponse
func (client *Client) DeleteFaceDeviceGroup(request *DeleteFaceDeviceGroupRequest) (_result *DeleteFaceDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceDeviceGroupResponse{}
	_body, _err := client.DeleteFaceDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFaceUserWithOptions(request *DeleteFaceUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFaceUserResponse
func (client *Client) DeleteFaceUser(request *DeleteFaceUserRequest) (_result *DeleteFaceUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceUserResponse{}
	_body, _err := client.DeleteFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFaceUserGroupWithOptions(request *DeleteFaceUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFaceUserGroupResponse
func (client *Client) DeleteFaceUserGroup(request *DeleteFaceUserGroupRequest) (_result *DeleteFaceUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceUserGroupResponse{}
	_body, _err := client.DeleteFaceUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelationWithOptions(request *DeleteFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelation(request *DeleteFaceUserGroupAndDeviceGroupRelationRequest) (_result *DeleteFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.DeleteFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFaceUserPictureWithOptions(request *DeleteFaceUserPictureRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceUserPictureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFaceUserPictureResponse
func (client *Client) DeleteFaceUserPicture(request *DeleteFaceUserPictureRequest) (_result *DeleteFaceUserPictureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceUserPictureResponse{}
	_body, _err := client.DeleteFaceUserPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteGbDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGbDeviceResponse
func (client *Client) DeleteGbDeviceWithOptions(request *DeleteGbDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGbDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteGbDeviceRequest
//
// @return DeleteGbDeviceResponse
func (client *Client) DeleteGbDevice(request *DeleteGbDeviceRequest) (_result *DeleteGbDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGbDeviceResponse{}
	_body, _err := client.DeleteGbDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteLocalFileUploadJobWithOptions(request *DeleteLocalFileUploadJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteLocalFileUploadJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteLocalFileUploadJobResponse
func (client *Client) DeleteLocalFileUploadJob(request *DeleteLocalFileUploadJobRequest) (_result *DeleteLocalFileUploadJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLocalFileUploadJobResponse{}
	_body, _err := client.DeleteLocalFileUploadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeletePictureWithOptions(request *DeletePictureRequest, runtime *dara.RuntimeOptions) (_result *DeletePictureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeletePictureResponse
func (client *Client) DeletePicture(request *DeletePictureRequest) (_result *DeletePictureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePictureResponse{}
	_body, _err := client.DeletePictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRecordWithOptions(request *DeleteRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRecordResponse
func (client *Client) DeleteRecord(request *DeleteRecordRequest) (_result *DeleteRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecordResponse{}
	_body, _err := client.DeleteRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRecordPlanWithOptions(request *DeleteRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRecordPlanResponse
func (client *Client) DeleteRecordPlan(request *DeleteRecordPlanRequest) (_result *DeleteRecordPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecordPlanResponse{}
	_body, _err := client.DeleteRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRecordPlanDeviceWithOptions(request *DeleteRecordPlanDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordPlanDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRecordPlanDeviceResponse
func (client *Client) DeleteRecordPlanDevice(request *DeleteRecordPlanDeviceRequest) (_result *DeleteRecordPlanDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecordPlanDeviceResponse{}
	_body, _err := client.DeleteRecordPlanDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteRtmpDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRtmpDeviceResponse
func (client *Client) DeleteRtmpDeviceWithOptions(request *DeleteRtmpDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteRtmpDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRtmpDeviceRequest
//
// @return DeleteRtmpDeviceResponse
func (client *Client) DeleteRtmpDevice(request *DeleteRtmpDeviceRequest) (_result *DeleteRtmpDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRtmpDeviceResponse{}
	_body, _err := client.DeleteRtmpDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteRtmpKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRtmpKeyResponse
func (client *Client) DeleteRtmpKeyWithOptions(request *DeleteRtmpKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteRtmpKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRtmpKeyRequest
//
// @return DeleteRtmpKeyResponse
func (client *Client) DeleteRtmpKey(request *DeleteRtmpKeyRequest) (_result *DeleteRtmpKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRtmpKeyResponse{}
	_body, _err := client.DeleteRtmpKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteStreamPushJobWithOptions(request *DeleteStreamPushJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamPushJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteStreamPushJobResponse
func (client *Client) DeleteStreamPushJob(request *DeleteStreamPushJobRequest) (_result *DeleteStreamPushJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStreamPushJobResponse{}
	_body, _err := client.DeleteStreamPushJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteStreamSnapshotJobWithOptions(request *DeleteStreamSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteStreamSnapshotJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteStreamSnapshotJobResponse
func (client *Client) DeleteStreamSnapshotJob(request *DeleteStreamSnapshotJobRequest) (_result *DeleteStreamSnapshotJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStreamSnapshotJobResponse{}
	_body, _err := client.DeleteStreamSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTimeTemplateWithOptions(request *DeleteTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTimeTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteTimeTemplateResponse
func (client *Client) DeleteTimeTemplate(request *DeleteTimeTemplateRequest) (_result *DeleteTimeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTimeTemplateResponse{}
	_body, _err := client.DeleteTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableGbSubDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableGbSubDeviceResponse
func (client *Client) EnableGbSubDeviceWithOptions(request *EnableGbSubDeviceRequest, runtime *dara.RuntimeOptions) (_result *EnableGbSubDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableGbSubDeviceRequest
//
// @return EnableGbSubDeviceResponse
func (client *Client) EnableGbSubDevice(request *EnableGbSubDeviceRequest) (_result *EnableGbSubDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableGbSubDeviceResponse{}
	_body, _err := client.EnableGbSubDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPictureSearchJobStatusWithOptions(request *GetPictureSearchJobStatusRequest, runtime *dara.RuntimeOptions) (_result *GetPictureSearchJobStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPictureSearchJobStatusResponse
func (client *Client) GetPictureSearchJobStatus(request *GetPictureSearchJobStatusRequest) (_result *GetPictureSearchJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPictureSearchJobStatusResponse{}
	_body, _err := client.GetPictureSearchJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PictureSearchPictureWithOptions(request *PictureSearchPictureRequest, runtime *dara.RuntimeOptions) (_result *PictureSearchPictureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PictureSearchPictureResponse
func (client *Client) PictureSearchPicture(request *PictureSearchPictureRequest) (_result *PictureSearchPictureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PictureSearchPictureResponse{}
	_body, _err := client.PictureSearchPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryCarProcessEventsWithOptions(request *QueryCarProcessEventsRequest, runtime *dara.RuntimeOptions) (_result *QueryCarProcessEventsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryCarProcessEventsResponse
func (client *Client) QueryCarProcessEvents(request *QueryCarProcessEventsRequest) (_result *QueryCarProcessEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCarProcessEventsResponse{}
	_body, _err := client.QueryCarProcessEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDeviceEventWithOptions(request *QueryDeviceEventRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDeviceEventResponse
func (client *Client) QueryDeviceEvent(request *QueryDeviceEventRequest) (_result *QueryDeviceEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceEventResponse{}
	_body, _err := client.QueryDeviceEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDeviceEventPictureWithOptions(request *QueryDeviceEventPictureRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceEventPictureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDeviceEventPictureResponse
func (client *Client) QueryDeviceEventPicture(request *QueryDeviceEventPictureRequest) (_result *QueryDeviceEventPictureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceEventPictureResponse{}
	_body, _err := client.QueryDeviceEventPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDeviceEventRecordWithOptions(request *QueryDeviceEventRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceEventRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDeviceEventRecordResponse
func (client *Client) QueryDeviceEventRecord(request *QueryDeviceEventRecordRequest) (_result *QueryDeviceEventRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceEventRecordResponse{}
	_body, _err := client.QueryDeviceEventRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDevicePictureByListWithOptions(request *QueryDevicePictureByListRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicePictureByListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDevicePictureByListResponse
func (client *Client) QueryDevicePictureByList(request *QueryDevicePictureByListRequest) (_result *QueryDevicePictureByListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDevicePictureByListResponse{}
	_body, _err := client.QueryDevicePictureByListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDevicePictureFileWithOptions(request *QueryDevicePictureFileRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicePictureFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDevicePictureFileResponse
func (client *Client) QueryDevicePictureFile(request *QueryDevicePictureFileRequest) (_result *QueryDevicePictureFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDevicePictureFileResponse{}
	_body, _err := client.QueryDevicePictureFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDevicePictureLifeCycleWithOptions(request *QueryDevicePictureLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicePictureLifeCycleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDevicePictureLifeCycleResponse
func (client *Client) QueryDevicePictureLifeCycle(request *QueryDevicePictureLifeCycleRequest) (_result *QueryDevicePictureLifeCycleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDevicePictureLifeCycleResponse{}
	_body, _err := client.QueryDevicePictureLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDeviceRecordLifeCycleWithOptions(request *QueryDeviceRecordLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceRecordLifeCycleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDeviceRecordLifeCycleResponse
func (client *Client) QueryDeviceRecordLifeCycle(request *QueryDeviceRecordLifeCycleRequest) (_result *QueryDeviceRecordLifeCycleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceRecordLifeCycleResponse{}
	_body, _err := client.QueryDeviceRecordLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDeviceVodUrlWithOptions(request *QueryDeviceVodUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceVodUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDeviceVodUrlResponse
func (client *Client) QueryDeviceVodUrl(request *QueryDeviceVodUrlRequest) (_result *QueryDeviceVodUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceVodUrlResponse{}
	_body, _err := client.QueryDeviceVodUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDeviceVodUrlByTimeWithOptions(request *QueryDeviceVodUrlByTimeRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceVodUrlByTimeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryDeviceVodUrlByTimeResponse
func (client *Client) QueryDeviceVodUrlByTime(request *QueryDeviceVodUrlByTimeRequest) (_result *QueryDeviceVodUrlByTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceVodUrlByTimeResponse{}
	_body, _err := client.QueryDeviceVodUrlByTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryEventRecordPlanDetailWithOptions(request *QueryEventRecordPlanDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlanDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryEventRecordPlanDetailResponse
func (client *Client) QueryEventRecordPlanDetail(request *QueryEventRecordPlanDetailRequest) (_result *QueryEventRecordPlanDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEventRecordPlanDetailResponse{}
	_body, _err := client.QueryEventRecordPlanDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryEventRecordPlanDeviceByDeviceWithOptions(request *QueryEventRecordPlanDeviceByDeviceRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlanDeviceByDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryEventRecordPlanDeviceByDeviceResponse
func (client *Client) QueryEventRecordPlanDeviceByDevice(request *QueryEventRecordPlanDeviceByDeviceRequest) (_result *QueryEventRecordPlanDeviceByDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEventRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.QueryEventRecordPlanDeviceByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryEventRecordPlanDeviceByPlanWithOptions(request *QueryEventRecordPlanDeviceByPlanRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlanDeviceByPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryEventRecordPlanDeviceByPlanResponse
func (client *Client) QueryEventRecordPlanDeviceByPlan(request *QueryEventRecordPlanDeviceByPlanRequest) (_result *QueryEventRecordPlanDeviceByPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEventRecordPlanDeviceByPlanResponse{}
	_body, _err := client.QueryEventRecordPlanDeviceByPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryEventRecordPlansWithOptions(request *QueryEventRecordPlansRequest, runtime *dara.RuntimeOptions) (_result *QueryEventRecordPlansResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryEventRecordPlansResponse
func (client *Client) QueryEventRecordPlans(request *QueryEventRecordPlansRequest) (_result *QueryEventRecordPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEventRecordPlansResponse{}
	_body, _err := client.QueryEventRecordPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceAllDeviceGroupWithOptions(request *QueryFaceAllDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllDeviceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceAllDeviceGroupResponse
func (client *Client) QueryFaceAllDeviceGroup(request *QueryFaceAllDeviceGroupRequest) (_result *QueryFaceAllDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceAllDeviceGroupResponse{}
	_body, _err := client.QueryFaceAllDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceAllUserGroupWithOptions(request *QueryFaceAllUserGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceAllUserGroupResponse
func (client *Client) QueryFaceAllUserGroup(request *QueryFaceAllUserGroupRequest) (_result *QueryFaceAllUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceAllUserGroupResponse{}
	_body, _err := client.QueryFaceAllUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceAllUserGroupAndDeviceGroupRelationWithOptions(request *QueryFaceAllUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllUserGroupAndDeviceGroupRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceAllUserGroupAndDeviceGroupRelationResponse
func (client *Client) QueryFaceAllUserGroupAndDeviceGroupRelation(request *QueryFaceAllUserGroupAndDeviceGroupRelationRequest) (_result *QueryFaceAllUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceAllUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.QueryFaceAllUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceAllUserIdsByGroupIdWithOptions(request *QueryFaceAllUserIdsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceAllUserIdsByGroupIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceAllUserIdsByGroupIdResponse
func (client *Client) QueryFaceAllUserIdsByGroupId(request *QueryFaceAllUserIdsByGroupIdRequest) (_result *QueryFaceAllUserIdsByGroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceAllUserIdsByGroupIdResponse{}
	_body, _err := client.QueryFaceAllUserIdsByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceCustomUserIdByUserIdWithOptions(request *QueryFaceCustomUserIdByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceCustomUserIdByUserIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceCustomUserIdByUserIdResponse
func (client *Client) QueryFaceCustomUserIdByUserId(request *QueryFaceCustomUserIdByUserIdRequest) (_result *QueryFaceCustomUserIdByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceCustomUserIdByUserIdResponse{}
	_body, _err := client.QueryFaceCustomUserIdByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceDeviceGroupsByDeviceWithOptions(request *QueryFaceDeviceGroupsByDeviceRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceDeviceGroupsByDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceDeviceGroupsByDeviceResponse
func (client *Client) QueryFaceDeviceGroupsByDevice(request *QueryFaceDeviceGroupsByDeviceRequest) (_result *QueryFaceDeviceGroupsByDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceDeviceGroupsByDeviceResponse{}
	_body, _err := client.QueryFaceDeviceGroupsByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceUserWithOptions(request *QueryFaceUserRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceUserResponse
func (client *Client) QueryFaceUser(request *QueryFaceUserRequest) (_result *QueryFaceUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceUserResponse{}
	_body, _err := client.QueryFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceUserBatchWithOptions(request *QueryFaceUserBatchRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserBatchResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceUserBatchResponse
func (client *Client) QueryFaceUserBatch(request *QueryFaceUserBatchRequest) (_result *QueryFaceUserBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceUserBatchResponse{}
	_body, _err := client.QueryFaceUserBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceUserByNameWithOptions(request *QueryFaceUserByNameRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserByNameResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceUserByNameResponse
func (client *Client) QueryFaceUserByName(request *QueryFaceUserByNameRequest) (_result *QueryFaceUserByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceUserByNameResponse{}
	_body, _err := client.QueryFaceUserByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceUserGroupWithOptions(request *QueryFaceUserGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceUserGroupResponse
func (client *Client) QueryFaceUserGroup(request *QueryFaceUserGroupRequest) (_result *QueryFaceUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceUserGroupResponse{}
	_body, _err := client.QueryFaceUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceUserGroupAndDeviceGroupRelationWithOptions(request *QueryFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) QueryFaceUserGroupAndDeviceGroupRelation(request *QueryFaceUserGroupAndDeviceGroupRelationRequest) (_result *QueryFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.QueryFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryFaceUserIdByCustomUserIdWithOptions(request *QueryFaceUserIdByCustomUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceUserIdByCustomUserIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryFaceUserIdByCustomUserIdResponse
func (client *Client) QueryFaceUserIdByCustomUserId(request *QueryFaceUserIdByCustomUserIdRequest) (_result *QueryFaceUserIdByCustomUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceUserIdByCustomUserIdResponse{}
	_body, _err := client.QueryFaceUserIdByCustomUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryGbSubDeviceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGbSubDeviceListResponse
func (client *Client) QueryGbSubDeviceListWithOptions(request *QueryGbSubDeviceListRequest, runtime *dara.RuntimeOptions) (_result *QueryGbSubDeviceListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryGbSubDeviceListRequest
//
// @return QueryGbSubDeviceListResponse
func (client *Client) QueryGbSubDeviceList(request *QueryGbSubDeviceListRequest) (_result *QueryGbSubDeviceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGbSubDeviceListResponse{}
	_body, _err := client.QueryGbSubDeviceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryLiveStreamingWithOptions(request *QueryLiveStreamingRequest, runtime *dara.RuntimeOptions) (_result *QueryLiveStreamingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryLiveStreamingResponse
func (client *Client) QueryLiveStreaming(request *QueryLiveStreamingRequest) (_result *QueryLiveStreamingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryLiveStreamingResponse{}
	_body, _err := client.QueryLiveStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryLocalFileUploadJobWithOptions(request *QueryLocalFileUploadJobRequest, runtime *dara.RuntimeOptions) (_result *QueryLocalFileUploadJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryLocalFileUploadJobResponse
func (client *Client) QueryLocalFileUploadJob(request *QueryLocalFileUploadJobRequest) (_result *QueryLocalFileUploadJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryLocalFileUploadJobResponse{}
	_body, _err := client.QueryLocalFileUploadJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMonthRecordWithOptions(request *QueryMonthRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryMonthRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryMonthRecordResponse
func (client *Client) QueryMonthRecord(request *QueryMonthRecordRequest) (_result *QueryMonthRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMonthRecordResponse{}
	_body, _err := client.QueryMonthRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPictureFilesWithOptions(request *QueryPictureFilesRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureFilesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPictureFilesResponse
func (client *Client) QueryPictureFiles(request *QueryPictureFilesRequest) (_result *QueryPictureFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPictureFilesResponse{}
	_body, _err := client.QueryPictureFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPictureSearchAiboxesWithOptions(request *QueryPictureSearchAiboxesRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchAiboxesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPictureSearchAiboxesResponse
func (client *Client) QueryPictureSearchAiboxes(request *QueryPictureSearchAiboxesRequest) (_result *QueryPictureSearchAiboxesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPictureSearchAiboxesResponse{}
	_body, _err := client.QueryPictureSearchAiboxesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPictureSearchAppsWithOptions(request *QueryPictureSearchAppsRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchAppsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPictureSearchAppsResponse
func (client *Client) QueryPictureSearchApps(request *QueryPictureSearchAppsRequest) (_result *QueryPictureSearchAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPictureSearchAppsResponse{}
	_body, _err := client.QueryPictureSearchAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPictureSearchDevicesWithOptions(request *QueryPictureSearchDevicesRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPictureSearchDevicesResponse
func (client *Client) QueryPictureSearchDevices(request *QueryPictureSearchDevicesRequest) (_result *QueryPictureSearchDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPictureSearchDevicesResponse{}
	_body, _err := client.QueryPictureSearchDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPictureSearchJobWithOptions(request *QueryPictureSearchJobRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPictureSearchJobResponse
func (client *Client) QueryPictureSearchJob(request *QueryPictureSearchJobRequest) (_result *QueryPictureSearchJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPictureSearchJobResponse{}
	_body, _err := client.QueryPictureSearchJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPictureSearchJobResultWithOptions(request *QueryPictureSearchJobResultRequest, runtime *dara.RuntimeOptions) (_result *QueryPictureSearchJobResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPictureSearchJobResultResponse
func (client *Client) QueryPictureSearchJobResult(request *QueryPictureSearchJobResultRequest) (_result *QueryPictureSearchJobResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPictureSearchJobResultResponse{}
	_body, _err := client.QueryPictureSearchJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordWithOptions(request *QueryRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordResponse
func (client *Client) QueryRecord(request *QueryRecordRequest) (_result *QueryRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordResponse{}
	_body, _err := client.QueryRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordByRecordIdWithOptions(request *QueryRecordByRecordIdRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordByRecordIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordByRecordIdResponse
func (client *Client) QueryRecordByRecordId(request *QueryRecordByRecordIdRequest) (_result *QueryRecordByRecordIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordByRecordIdResponse{}
	_body, _err := client.QueryRecordByRecordIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordDownloadJobByIdWithOptions(request *QueryRecordDownloadJobByIdRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordDownloadJobByIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordDownloadJobByIdResponse
func (client *Client) QueryRecordDownloadJobById(request *QueryRecordDownloadJobByIdRequest) (_result *QueryRecordDownloadJobByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordDownloadJobByIdResponse{}
	_body, _err := client.QueryRecordDownloadJobByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordDownloadJobListWithOptions(request *QueryRecordDownloadJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordDownloadJobListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordDownloadJobListResponse
func (client *Client) QueryRecordDownloadJobList(request *QueryRecordDownloadJobListRequest) (_result *QueryRecordDownloadJobListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordDownloadJobListResponse{}
	_body, _err := client.QueryRecordDownloadJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordDownloadUrlWithOptions(request *QueryRecordDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordDownloadUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordDownloadUrlResponse
func (client *Client) QueryRecordDownloadUrl(request *QueryRecordDownloadUrlRequest) (_result *QueryRecordDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordDownloadUrlResponse{}
	_body, _err := client.QueryRecordDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordPlanDetailWithOptions(request *QueryRecordPlanDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlanDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordPlanDetailResponse
func (client *Client) QueryRecordPlanDetail(request *QueryRecordPlanDetailRequest) (_result *QueryRecordPlanDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordPlanDetailResponse{}
	_body, _err := client.QueryRecordPlanDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordPlanDeviceByDeviceWithOptions(request *QueryRecordPlanDeviceByDeviceRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlanDeviceByDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordPlanDeviceByDeviceResponse
func (client *Client) QueryRecordPlanDeviceByDevice(request *QueryRecordPlanDeviceByDeviceRequest) (_result *QueryRecordPlanDeviceByDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordPlanDeviceByDeviceResponse{}
	_body, _err := client.QueryRecordPlanDeviceByDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordPlanDeviceByPlanWithOptions(request *QueryRecordPlanDeviceByPlanRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlanDeviceByPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordPlanDeviceByPlanResponse
func (client *Client) QueryRecordPlanDeviceByPlan(request *QueryRecordPlanDeviceByPlanRequest) (_result *QueryRecordPlanDeviceByPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordPlanDeviceByPlanResponse{}
	_body, _err := client.QueryRecordPlanDeviceByPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordPlansWithOptions(request *QueryRecordPlansRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordPlansResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordPlansResponse
func (client *Client) QueryRecordPlans(request *QueryRecordPlansRequest) (_result *QueryRecordPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordPlansResponse{}
	_body, _err := client.QueryRecordPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordUrlWithOptions(request *QueryRecordUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordUrlResponse
func (client *Client) QueryRecordUrl(request *QueryRecordUrlRequest) (_result *QueryRecordUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordUrlResponse{}
	_body, _err := client.QueryRecordUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordUrlByTimeWithOptions(request *QueryRecordUrlByTimeRequest, runtime *dara.RuntimeOptions) (_result *QueryRecordUrlByTimeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryRecordUrlByTimeResponse
func (client *Client) QueryRecordUrlByTime(request *QueryRecordUrlByTimeRequest) (_result *QueryRecordUrlByTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecordUrlByTimeResponse{}
	_body, _err := client.QueryRecordUrlByTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryRtmpKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRtmpKeyResponse
func (client *Client) QueryRtmpKeyWithOptions(request *QueryRtmpKeyRequest, runtime *dara.RuntimeOptions) (_result *QueryRtmpKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryRtmpKeyRequest
//
// @return QueryRtmpKeyResponse
func (client *Client) QueryRtmpKey(request *QueryRtmpKeyRequest) (_result *QueryRtmpKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRtmpKeyResponse{}
	_body, _err := client.QueryRtmpKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryStreamPushJobWithOptions(request *QueryStreamPushJobRequest, runtime *dara.RuntimeOptions) (_result *QueryStreamPushJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryStreamPushJobResponse
func (client *Client) QueryStreamPushJob(request *QueryStreamPushJobRequest) (_result *QueryStreamPushJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryStreamPushJobResponse{}
	_body, _err := client.QueryStreamPushJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryStreamPushJobListWithOptions(request *QueryStreamPushJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryStreamPushJobListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryStreamPushJobListResponse
func (client *Client) QueryStreamPushJobList(request *QueryStreamPushJobListRequest) (_result *QueryStreamPushJobListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryStreamPushJobListResponse{}
	_body, _err := client.QueryStreamPushJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryStreamSnapshotJobWithOptions(request *QueryStreamSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *QueryStreamSnapshotJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryStreamSnapshotJobResponse
func (client *Client) QueryStreamSnapshotJob(request *QueryStreamSnapshotJobRequest) (_result *QueryStreamSnapshotJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryStreamSnapshotJobResponse{}
	_body, _err := client.QueryStreamSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryTimeTemplateWithOptions(request *QueryTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryTimeTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryTimeTemplateResponse
func (client *Client) QueryTimeTemplate(request *QueryTimeTemplateRequest) (_result *QueryTimeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTimeTemplateResponse{}
	_body, _err := client.QueryTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryTimeTemplateDetailWithOptions(request *QueryTimeTemplateDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryTimeTemplateDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryTimeTemplateDetailResponse
func (client *Client) QueryTimeTemplateDetail(request *QueryTimeTemplateDetailRequest) (_result *QueryTimeTemplateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTimeTemplateDetailResponse{}
	_body, _err := client.QueryTimeTemplateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryVisionDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVisionDeviceInfoResponse
func (client *Client) QueryVisionDeviceInfoWithOptions(request *QueryVisionDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryVisionDeviceInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryVisionDeviceInfoRequest
//
// @return QueryVisionDeviceInfoResponse
func (client *Client) QueryVisionDeviceInfo(request *QueryVisionDeviceInfoRequest) (_result *QueryVisionDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVisionDeviceInfoResponse{}
	_body, _err := client.QueryVisionDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryVoiceIntercomWithOptions(request *QueryVoiceIntercomRequest, runtime *dara.RuntimeOptions) (_result *QueryVoiceIntercomResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryVoiceIntercomResponse
func (client *Client) QueryVoiceIntercom(request *QueryVoiceIntercomRequest) (_result *QueryVoiceIntercomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVoiceIntercomResponse{}
	_body, _err := client.QueryVoiceIntercomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RefreshGbSubDeviceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshGbSubDeviceListResponse
func (client *Client) RefreshGbSubDeviceListWithOptions(request *RefreshGbSubDeviceListRequest, runtime *dara.RuntimeOptions) (_result *RefreshGbSubDeviceListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshGbSubDeviceListRequest
//
// @return RefreshGbSubDeviceListResponse
func (client *Client) RefreshGbSubDeviceList(request *RefreshGbSubDeviceListRequest) (_result *RefreshGbSubDeviceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshGbSubDeviceListResponse{}
	_body, _err := client.RefreshGbSubDeviceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveFaceDeviceFromDeviceGroupWithOptions(request *RemoveFaceDeviceFromDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveFaceDeviceFromDeviceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveFaceDeviceFromDeviceGroupResponse
func (client *Client) RemoveFaceDeviceFromDeviceGroup(request *RemoveFaceDeviceFromDeviceGroupRequest) (_result *RemoveFaceDeviceFromDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveFaceDeviceFromDeviceGroupResponse{}
	_body, _err := client.RemoveFaceDeviceFromDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveFaceUserFromUserGroupWithOptions(request *RemoveFaceUserFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveFaceUserFromUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveFaceUserFromUserGroupResponse
func (client *Client) RemoveFaceUserFromUserGroup(request *RemoveFaceUserFromUserGroupRequest) (_result *RemoveFaceUserFromUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveFaceUserFromUserGroupResponse{}
	_body, _err := client.RemoveFaceUserFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetDevicePictureLifeCycleWithOptions(request *SetDevicePictureLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *SetDevicePictureLifeCycleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetDevicePictureLifeCycleResponse
func (client *Client) SetDevicePictureLifeCycle(request *SetDevicePictureLifeCycleRequest) (_result *SetDevicePictureLifeCycleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDevicePictureLifeCycleResponse{}
	_body, _err := client.SetDevicePictureLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetDeviceRecordLifeCycleWithOptions(request *SetDeviceRecordLifeCycleRequest, runtime *dara.RuntimeOptions) (_result *SetDeviceRecordLifeCycleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetDeviceRecordLifeCycleResponse
func (client *Client) SetDeviceRecordLifeCycle(request *SetDeviceRecordLifeCycleRequest) (_result *SetDeviceRecordLifeCycleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDeviceRecordLifeCycleResponse{}
	_body, _err := client.SetDeviceRecordLifeCycleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopLiveStreamingWithOptions(request *StopLiveStreamingRequest, runtime *dara.RuntimeOptions) (_result *StopLiveStreamingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopLiveStreamingResponse
func (client *Client) StopLiveStreaming(request *StopLiveStreamingRequest) (_result *StopLiveStreamingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopLiveStreamingResponse{}
	_body, _err := client.StopLiveStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopTriggeredRecordWithOptions(request *StopTriggeredRecordRequest, runtime *dara.RuntimeOptions) (_result *StopTriggeredRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopTriggeredRecordResponse
func (client *Client) StopTriggeredRecord(request *StopTriggeredRecordRequest) (_result *StopTriggeredRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopTriggeredRecordResponse{}
	_body, _err := client.StopTriggeredRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TransferDeviceInstanceWithOptions(request *TransferDeviceInstanceRequest, runtime *dara.RuntimeOptions) (_result *TransferDeviceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TransferDeviceInstanceResponse
func (client *Client) TransferDeviceInstance(request *TransferDeviceInstanceRequest) (_result *TransferDeviceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferDeviceInstanceResponse{}
	_body, _err := client.TransferDeviceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TriggerCapturePictureWithOptions(request *TriggerCapturePictureRequest, runtime *dara.RuntimeOptions) (_result *TriggerCapturePictureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TriggerCapturePictureResponse
func (client *Client) TriggerCapturePicture(request *TriggerCapturePictureRequest) (_result *TriggerCapturePictureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TriggerCapturePictureResponse{}
	_body, _err := client.TriggerCapturePictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TriggerRecordWithOptions(request *TriggerRecordRequest, runtime *dara.RuntimeOptions) (_result *TriggerRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TriggerRecordResponse
func (client *Client) TriggerRecord(request *TriggerRecordRequest) (_result *TriggerRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TriggerRecordResponse{}
	_body, _err := client.TriggerRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnbindPictureSearchAppWithDevicesWithOptions(request *UnbindPictureSearchAppWithDevicesRequest, runtime *dara.RuntimeOptions) (_result *UnbindPictureSearchAppWithDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnbindPictureSearchAppWithDevicesResponse
func (client *Client) UnbindPictureSearchAppWithDevices(request *UnbindPictureSearchAppWithDevicesRequest) (_result *UnbindPictureSearchAppWithDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindPictureSearchAppWithDevicesResponse{}
	_body, _err := client.UnbindPictureSearchAppWithDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateEventRecordPlanWithOptions(request *UpdateEventRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventRecordPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateEventRecordPlanResponse
func (client *Client) UpdateEventRecordPlan(request *UpdateEventRecordPlanRequest) (_result *UpdateEventRecordPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEventRecordPlanResponse{}
	_body, _err := client.UpdateEventRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateFaceUserWithOptions(request *UpdateFaceUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateFaceUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateFaceUserResponse
func (client *Client) UpdateFaceUser(request *UpdateFaceUserRequest) (_result *UpdateFaceUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFaceUserResponse{}
	_body, _err := client.UpdateFaceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateFaceUserGroupAndDeviceGroupRelationWithOptions(request *UpdateFaceUserGroupAndDeviceGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *UpdateFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateFaceUserGroupAndDeviceGroupRelationResponse
func (client *Client) UpdateFaceUserGroupAndDeviceGroupRelation(request *UpdateFaceUserGroupAndDeviceGroupRelationRequest) (_result *UpdateFaceUserGroupAndDeviceGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFaceUserGroupAndDeviceGroupRelationResponse{}
	_body, _err := client.UpdateFaceUserGroupAndDeviceGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateGbDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGbDeviceResponse
func (client *Client) UpdateGbDeviceWithOptions(request *UpdateGbDeviceRequest, runtime *dara.RuntimeOptions) (_result *UpdateGbDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateGbDeviceRequest
//
// @return UpdateGbDeviceResponse
func (client *Client) UpdateGbDevice(request *UpdateGbDeviceRequest) (_result *UpdateGbDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGbDeviceResponse{}
	_body, _err := client.UpdateGbDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateInstanceInternetProtocolWithOptions(request *UpdateInstanceInternetProtocolRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceInternetProtocolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateInstanceInternetProtocolResponse
func (client *Client) UpdateInstanceInternetProtocol(request *UpdateInstanceInternetProtocolRequest) (_result *UpdateInstanceInternetProtocolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceInternetProtocolResponse{}
	_body, _err := client.UpdateInstanceInternetProtocolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdatePictureSearchAppWithOptions(request *UpdatePictureSearchAppRequest, runtime *dara.RuntimeOptions) (_result *UpdatePictureSearchAppResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdatePictureSearchAppResponse
func (client *Client) UpdatePictureSearchApp(request *UpdatePictureSearchAppRequest) (_result *UpdatePictureSearchAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePictureSearchAppResponse{}
	_body, _err := client.UpdatePictureSearchAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRecordPlanWithOptions(request *UpdateRecordPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRecordPlanResponse
func (client *Client) UpdateRecordPlan(request *UpdateRecordPlanRequest) (_result *UpdateRecordPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecordPlanResponse{}
	_body, _err := client.UpdateRecordPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateRtmpKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRtmpKeyResponse
func (client *Client) UpdateRtmpKeyWithOptions(request *UpdateRtmpKeyRequest, runtime *dara.RuntimeOptions) (_result *UpdateRtmpKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRtmpKeyRequest
//
// @return UpdateRtmpKeyResponse
func (client *Client) UpdateRtmpKey(request *UpdateRtmpKeyRequest) (_result *UpdateRtmpKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRtmpKeyResponse{}
	_body, _err := client.UpdateRtmpKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTimeTemplateWithOptions(request *UpdateTimeTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateTimeTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateTimeTemplateResponse
func (client *Client) UpdateTimeTemplate(request *UpdateTimeTemplateRequest) (_result *UpdateTimeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTimeTemplateResponse{}
	_body, _err := client.UpdateTimeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
