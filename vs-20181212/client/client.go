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
		"cn-shenzhen": dara.String("vs.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai": dara.String("vs.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":  dara.String("vs.cn-qingdao.aliyuncs.com"),
		"cn-beijing":  dara.String("vs.cn-beijing.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("vs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds one or more instances to a specified cluster.
//
// Description:
//
// ## Usage notes
//
// - **HiveId*	- is a required parameter that specifies the ID of the target cluster.
//
// - **InstanceIds*	- is a required parameter that specifies a list of instance IDs to add.
//
// - Adding an instance that already exists in the target cluster returns an error message.
//
// - The response includes lists of successful and failed instances. This allows you to verify which instances were added and review the reasons for any failures.
//
// @param tmpReq - AddHiveEdgeWorkersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddHiveEdgeWorkersResponse
func (client *Client) AddHiveEdgeWorkersWithOptions(tmpReq *AddHiveEdgeWorkersRequest, runtime *dara.RuntimeOptions) (_result *AddHiveEdgeWorkersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddHiveEdgeWorkersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HiveId) {
		query["HiveId"] = request.HiveId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddHiveEdgeWorkers"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddHiveEdgeWorkersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more instances to a specified cluster.
//
// Description:
//
// ## Usage notes
//
// - **HiveId*	- is a required parameter that specifies the ID of the target cluster.
//
// - **InstanceIds*	- is a required parameter that specifies a list of instance IDs to add.
//
// - Adding an instance that already exists in the target cluster returns an error message.
//
// - The response includes lists of successful and failed instances. This allows you to verify which instances were added and review the reasons for any failures.
//
// @param request - AddHiveEdgeWorkersRequest
//
// @return AddHiveEdgeWorkersResponse
func (client *Client) AddHiveEdgeWorkers(request *AddHiveEdgeWorkersRequest) (_result *AddHiveEdgeWorkersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddHiveEdgeWorkersResponse{}
	_body, _err := client.AddHiveEdgeWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a stream pulling configuration.
//
// @param request - AddVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVsPullStreamInfoConfigResponse
func (client *Client) AddVsPullStreamInfoConfigWithOptions(request *AddVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *AddVsPullStreamInfoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Always) {
		query["Always"] = request.Always
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceUrl) {
		query["SourceUrl"] = request.SourceUrl
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
		Action:      dara.String("AddVsPullStreamInfoConfig"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a stream pulling configuration.
//
// @param request - AddVsPullStreamInfoConfigRequest
//
// @return AddVsPullStreamInfoConfigResponse
func (client *Client) AddVsPullStreamInfoConfig(request *AddVsPullStreamInfoConfigRequest) (_result *AddVsPullStreamInfoConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddVsPullStreamInfoConfigResponse{}
	_body, _err := client.AddVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates cloud application service instances with a project.
//
// Description:
//
// ## Request description
//
// - This operation associates instances that meet specific conditions with a specified project.
//
// @param tmpReq - AssociateRenderingProjectInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateRenderingProjectInstancesResponse
func (client *Client) AssociateRenderingProjectInstancesWithOptions(tmpReq *AssociateRenderingProjectInstancesRequest, runtime *dara.RuntimeOptions) (_result *AssociateRenderingProjectInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AssociateRenderingProjectInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RenderingInstanceIds) {
		request.RenderingInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RenderingInstanceIds, dara.String("RenderingInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RenderingInstanceIdsShrink) {
		query["RenderingInstanceIds"] = request.RenderingInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateRenderingProjectInstances"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateRenderingProjectInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates cloud application service instances with a project.
//
// Description:
//
// ## Request description
//
// - This operation associates instances that meet specific conditions with a specified project.
//
// @param request - AssociateRenderingProjectInstancesRequest
//
// @return AssociateRenderingProjectInstancesResponse
func (client *Client) AssociateRenderingProjectInstances(request *AssociateRenderingProjectInstancesRequest) (_result *AssociateRenderingProjectInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateRenderingProjectInstancesResponse{}
	_body, _err := client.AssociateRenderingProjectInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds multiple devices to directories in a single operation.
//
// @param request - BatchBindDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindDirectoriesResponse
func (client *Client) BatchBindDirectoriesWithOptions(request *BatchBindDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchBindDirectories"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchBindDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds multiple devices to directories in a single operation.
//
// @param request - BatchBindDirectoriesRequest
//
// @return BatchBindDirectoriesResponse
func (client *Client) BatchBindDirectories(request *BatchBindDirectoriesRequest) (_result *BatchBindDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchBindDirectoriesResponse{}
	_body, _err := client.BatchBindDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds multiple devices to a parent platform for push in batches.
//
// @param request - BatchBindParentPlatformDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindParentPlatformDevicesResponse
func (client *Client) BatchBindParentPlatformDevicesWithOptions(request *BatchBindParentPlatformDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindParentPlatformDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParentPlatformId) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchBindParentPlatformDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchBindParentPlatformDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds multiple devices to a parent platform for push in batches.
//
// @param request - BatchBindParentPlatformDevicesRequest
//
// @return BatchBindParentPlatformDevicesResponse
func (client *Client) BatchBindParentPlatformDevices(request *BatchBindParentPlatformDevicesRequest) (_result *BatchBindParentPlatformDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchBindParentPlatformDevicesResponse{}
	_body, _err := client.BatchBindParentPlatformDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds multiple purchased devices.
//
// @param request - BatchBindPurchasedDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindPurchasedDevicesResponse
func (client *Client) BatchBindPurchasedDevicesWithOptions(request *BatchBindPurchasedDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindPurchasedDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchBindPurchasedDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchBindPurchasedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds multiple purchased devices.
//
// @param request - BatchBindPurchasedDevicesRequest
//
// @return BatchBindPurchasedDevicesResponse
func (client *Client) BatchBindPurchasedDevices(request *BatchBindPurchasedDevicesRequest) (_result *BatchBindPurchasedDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchBindPurchasedDevicesResponse{}
	_body, _err := client.BatchBindPurchasedDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Bind templates to multiple specified instances, such as instances bound to spaces and streams.
//
// @param request - BatchBindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindTemplateResponse
func (client *Client) BatchBindTemplateWithOptions(request *BatchBindTemplateRequest, runtime *dara.RuntimeOptions) (_result *BatchBindTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyAll) {
		query["ApplyAll"] = request.ApplyAll
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Replace) {
		query["Replace"] = request.Replace
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchBindTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchBindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Bind templates to multiple specified instances, such as instances bound to spaces and streams.
//
// @param request - BatchBindTemplateRequest
//
// @return BatchBindTemplateResponse
func (client *Client) BatchBindTemplate(request *BatchBindTemplateRequest) (_result *BatchBindTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchBindTemplateResponse{}
	_body, _err := client.BatchBindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds multiple templates in a single operation.
//
// @param request - BatchBindTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindTemplatesResponse
func (client *Client) BatchBindTemplatesWithOptions(request *BatchBindTemplatesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyAll) {
		query["ApplyAll"] = request.ApplyAll
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Replace) {
		query["Replace"] = request.Replace
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchBindTemplates"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchBindTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds multiple templates in a single operation.
//
// @param request - BatchBindTemplatesRequest
//
// @return BatchBindTemplatesResponse
func (client *Client) BatchBindTemplates(request *BatchBindTemplatesRequest) (_result *BatchBindTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchBindTemplatesResponse{}
	_body, _err := client.BatchBindTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves screenshots of cloud application service instances.
//
// Description:
//
// ## Request description
//
// - **Authentication**: Requests must include the `AliUid` parameter for identity verification.
//
// - **Instance specification**: Use `RenderingInstanceIds` to specify the instances to capture screenshots from.
//
// - **Screenshot quality**: Use the `Quality` parameter to set the image quality of screenshots. The default value is 75 (if not configured). Valid values: 1 to 100.
//
// - **Response handling**: The response contains lists of successful and failed instances with related information, including download URLs and screenshot completion times.
//
// @param tmpReq - BatchCaptureRenderingInstanceScreenshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCaptureRenderingInstanceScreenshotResponse
func (client *Client) BatchCaptureRenderingInstanceScreenshotWithOptions(tmpReq *BatchCaptureRenderingInstanceScreenshotRequest, runtime *dara.RuntimeOptions) (_result *BatchCaptureRenderingInstanceScreenshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCaptureRenderingInstanceScreenshotShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RenderingInstanceIds) {
		request.RenderingInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RenderingInstanceIds, dara.String("RenderingInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Quality) {
		query["Quality"] = request.Quality
	}

	if !dara.IsNil(request.RenderingInstanceIdsShrink) {
		query["RenderingInstanceIds"] = request.RenderingInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCaptureRenderingInstanceScreenshot"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCaptureRenderingInstanceScreenshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves screenshots of cloud application service instances.
//
// Description:
//
// ## Request description
//
// - **Authentication**: Requests must include the `AliUid` parameter for identity verification.
//
// - **Instance specification**: Use `RenderingInstanceIds` to specify the instances to capture screenshots from.
//
// - **Screenshot quality**: Use the `Quality` parameter to set the image quality of screenshots. The default value is 75 (if not configured). Valid values: 1 to 100.
//
// - **Response handling**: The response contains lists of successful and failed instances with related information, including download URLs and screenshot completion times.
//
// @param request - BatchCaptureRenderingInstanceScreenshotRequest
//
// @return BatchCaptureRenderingInstanceScreenshotResponse
func (client *Client) BatchCaptureRenderingInstanceScreenshot(request *BatchCaptureRenderingInstanceScreenshotRequest) (_result *BatchCaptureRenderingInstanceScreenshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCaptureRenderingInstanceScreenshotResponse{}
	_body, _err := client.BatchCaptureRenderingInstanceScreenshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple devices in a single operation.
//
// @param request - BatchDeleteDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDevicesResponse
func (client *Client) BatchDeleteDevicesWithOptions(request *BatchDeleteDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDevicesResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple devices in a single operation.
//
// @param request - BatchDeleteDevicesRequest
//
// @return BatchDeleteDevicesResponse
func (client *Client) BatchDeleteDevices(request *BatchDeleteDevicesRequest) (_result *BatchDeleteDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDevicesResponse{}
	_body, _err := client.BatchDeleteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes domain name configurations in a batch.
//
// @param request - BatchDeleteVsDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteVsDomainConfigsResponse
func (client *Client) BatchDeleteVsDomainConfigsWithOptions(request *BatchDeleteVsDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteVsDomainConfigsResponse, _err error) {
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

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteVsDomainConfigs"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteVsDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes domain name configurations in a batch.
//
// @param request - BatchDeleteVsDomainConfigsRequest
//
// @return BatchDeleteVsDomainConfigsResponse
func (client *Client) BatchDeleteVsDomainConfigs(request *BatchDeleteVsDomainConfigsRequest) (_result *BatchDeleteVsDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteVsDomainConfigsResponse{}
	_body, _err := client.BatchDeleteVsDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stop stream ingest for one or more streams. You can schedule when to resume ingest.
//
// @param request - BatchForbidVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchForbidVsStreamResponse
func (client *Client) BatchForbidVsStreamWithOptions(request *BatchForbidVsStreamRequest, runtime *dara.RuntimeOptions) (_result *BatchForbidVsStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ControlStreamAction) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LiveStreamType) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !dara.IsNil(request.Oneshot) {
		query["Oneshot"] = request.Oneshot
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResumeTime) {
		query["ResumeTime"] = request.ResumeTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchForbidVsStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchForbidVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop stream ingest for one or more streams. You can schedule when to resume ingest.
//
// @param request - BatchForbidVsStreamRequest
//
// @return BatchForbidVsStreamResponse
func (client *Client) BatchForbidVsStream(request *BatchForbidVsStreamRequest) (_result *BatchForbidVsStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchForbidVsStreamResponse{}
	_body, _err := client.BatchForbidVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes stream ingest for one or more streams.
//
// @param request - BatchResumeVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchResumeVsStreamResponse
func (client *Client) BatchResumeVsStreamWithOptions(request *BatchResumeVsStreamRequest, runtime *dara.RuntimeOptions) (_result *BatchResumeVsStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ControlStreamAction) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LiveStreamType) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchResumeVsStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchResumeVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes stream ingest for one or more streams.
//
// @param request - BatchResumeVsStreamRequest
//
// @return BatchResumeVsStreamResponse
func (client *Client) BatchResumeVsStream(request *BatchResumeVsStreamRequest) (_result *BatchResumeVsStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchResumeVsStreamResponse{}
	_body, _err := client.BatchResumeVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configure multiple domain names in batch.
//
// @param request - BatchSetVsDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetVsDomainConfigsResponse
func (client *Client) BatchSetVsDomainConfigsWithOptions(request *BatchSetVsDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetVsDomainConfigsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetVsDomainConfigs"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetVsDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configure multiple domain names in batch.
//
// @param request - BatchSetVsDomainConfigsRequest
//
// @return BatchSetVsDomainConfigsResponse
func (client *Client) BatchSetVsDomainConfigs(request *BatchSetVsDomainConfigsRequest) (_result *BatchSetVsDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSetVsDomainConfigsResponse{}
	_body, _err := client.BatchSetVsDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start stream pulling for multiple devices at once.
//
// @param request - BatchStartDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartDevicesResponse
func (client *Client) BatchStartDevicesWithOptions(request *BatchStartDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchStartDevicesResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStartDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStartDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start stream pulling for multiple devices at once.
//
// @param request - BatchStartDevicesRequest
//
// @return BatchStartDevicesResponse
func (client *Client) BatchStartDevices(request *BatchStartDevicesRequest) (_result *BatchStartDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStartDevicesResponse{}
	_body, _err := client.BatchStartDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts multiple streams.
//
// @param request - BatchStartStreamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartStreamsResponse
func (client *Client) BatchStartStreamsWithOptions(request *BatchStartStreamsRequest, runtime *dara.RuntimeOptions) (_result *BatchStartStreamsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStartStreams"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStartStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts multiple streams.
//
// @param request - BatchStartStreamsRequest
//
// @return BatchStartStreamsResponse
func (client *Client) BatchStartStreams(request *BatchStartStreamsRequest) (_result *BatchStartStreamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStartStreamsResponse{}
	_body, _err := client.BatchStartStreamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops stream pulling for multiple devices.
//
// @param request - BatchStopDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopDevicesResponse
func (client *Client) BatchStopDevicesWithOptions(request *BatchStopDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchStopDevicesResponse, _err error) {
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
		Action:      dara.String("BatchStopDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStopDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops stream pulling for multiple devices.
//
// @param request - BatchStopDevicesRequest
//
// @return BatchStopDevicesResponse
func (client *Client) BatchStopDevices(request *BatchStopDevicesRequest) (_result *BatchStopDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStopDevicesResponse{}
	_body, _err := client.BatchStopDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops multiple streams in a batch.
//
// @param request - BatchStopStreamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopStreamsResponse
func (client *Client) BatchStopStreamsWithOptions(request *BatchStopStreamsRequest, runtime *dara.RuntimeOptions) (_result *BatchStopStreamsResponse, _err error) {
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
		Action:      dara.String("BatchStopStreams"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStopStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple streams in a batch.
//
// @param request - BatchStopStreamsRequest
//
// @return BatchStopStreamsResponse
func (client *Client) BatchStopStreams(request *BatchStopStreamsRequest) (_result *BatchStopStreamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStopStreamsResponse{}
	_body, _err := client.BatchStopStreamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches multiple devices from a folder in bulk.
//
// @param request - BatchUnbindDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindDirectoriesResponse
func (client *Client) BatchUnbindDirectoriesWithOptions(request *BatchUnbindDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUnbindDirectories"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUnbindDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches multiple devices from a folder in bulk.
//
// @param request - BatchUnbindDirectoriesRequest
//
// @return BatchUnbindDirectoriesResponse
func (client *Client) BatchUnbindDirectories(request *BatchUnbindDirectoriesRequest) (_result *BatchUnbindDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUnbindDirectoriesResponse{}
	_body, _err := client.BatchUnbindDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch unbind multiple devices from parent platform push.
//
// @param request - BatchUnbindParentPlatformDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindParentPlatformDevicesResponse
func (client *Client) BatchUnbindParentPlatformDevicesWithOptions(request *BatchUnbindParentPlatformDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindParentPlatformDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParentPlatformId) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUnbindParentPlatformDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUnbindParentPlatformDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch unbind multiple devices from parent platform push.
//
// @param request - BatchUnbindParentPlatformDevicesRequest
//
// @return BatchUnbindParentPlatformDevicesResponse
func (client *Client) BatchUnbindParentPlatformDevices(request *BatchUnbindParentPlatformDevicesRequest) (_result *BatchUnbindParentPlatformDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUnbindParentPlatformDevicesResponse{}
	_body, _err := client.BatchUnbindParentPlatformDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detach multiple purchased devices from a space in a single operation.
//
// @param request - BatchUnbindPurchasedDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindPurchasedDevicesResponse
func (client *Client) BatchUnbindPurchasedDevicesWithOptions(request *BatchUnbindPurchasedDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindPurchasedDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUnbindPurchasedDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUnbindPurchasedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach multiple purchased devices from a space in a single operation.
//
// @param request - BatchUnbindPurchasedDevicesRequest
//
// @return BatchUnbindPurchasedDevicesResponse
func (client *Client) BatchUnbindPurchasedDevices(request *BatchUnbindPurchasedDevicesRequest) (_result *BatchUnbindPurchasedDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUnbindPurchasedDevicesResponse{}
	_body, _err := client.BatchUnbindPurchasedDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detach a template from multiple specified instances, such as space instances or stream instances.
//
// Description:
//
// > Specify at least one of TemplateId or TemplateType.
//
// @param request - BatchUnbindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindTemplateResponse
func (client *Client) BatchUnbindTemplateWithOptions(request *BatchUnbindTemplateRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUnbindTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUnbindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach a template from multiple specified instances, such as space instances or stream instances.
//
// Description:
//
// > Specify at least one of TemplateId or TemplateType.
//
// @param request - BatchUnbindTemplateRequest
//
// @return BatchUnbindTemplateResponse
func (client *Client) BatchUnbindTemplate(request *BatchUnbindTemplateRequest) (_result *BatchUnbindTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUnbindTemplateResponse{}
	_body, _err := client.BatchUnbindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbind multiple templates simultaneously.
//
// @param request - BatchUnbindTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindTemplatesResponse
func (client *Client) BatchUnbindTemplatesWithOptions(request *BatchUnbindTemplatesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUnbindTemplates"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUnbindTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbind multiple templates simultaneously.
//
// @param request - BatchUnbindTemplatesRequest
//
// @return BatchUnbindTemplatesResponse
func (client *Client) BatchUnbindTemplates(request *BatchUnbindTemplatesRequest) (_result *BatchUnbindTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUnbindTemplatesResponse{}
	_body, _err := client.BatchUnbindTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attach a device to a folder.
//
// @param request - BindDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindDirectoryResponse
func (client *Client) BindDirectoryWithOptions(request *BindDirectoryRequest, runtime *dara.RuntimeOptions) (_result *BindDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindDirectory"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attach a device to a folder.
//
// @param request - BindDirectoryRequest
//
// @return BindDirectoryResponse
func (client *Client) BindDirectory(request *BindDirectoryRequest) (_result *BindDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindDirectoryResponse{}
	_body, _err := client.BindDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a device to push streams to a parent platform.
//
// @param request - BindParentPlatformDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindParentPlatformDeviceResponse
func (client *Client) BindParentPlatformDeviceWithOptions(request *BindParentPlatformDeviceRequest, runtime *dara.RuntimeOptions) (_result *BindParentPlatformDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParentPlatformId) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindParentPlatformDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindParentPlatformDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a device to push streams to a parent platform.
//
// @param request - BindParentPlatformDeviceRequest
//
// @return BindParentPlatformDeviceResponse
func (client *Client) BindParentPlatformDevice(request *BindParentPlatformDeviceRequest) (_result *BindParentPlatformDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindParentPlatformDeviceResponse{}
	_body, _err := client.BindParentPlatformDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attach purchased devices to a space.
//
// @param request - BindPurchasedDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPurchasedDeviceResponse
func (client *Client) BindPurchasedDeviceWithOptions(request *BindPurchasedDeviceRequest, runtime *dara.RuntimeOptions) (_result *BindPurchasedDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPurchasedDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attach purchased devices to a space.
//
// @param request - BindPurchasedDeviceRequest
//
// @return BindPurchasedDeviceResponse
func (client *Client) BindPurchasedDevice(request *BindPurchasedDeviceRequest) (_result *BindPurchasedDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindPurchasedDeviceResponse{}
	_body, _err := client.BindPurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a template to a specified instance, such as a group or stream.
//
// @param request - BindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindTemplateResponse
func (client *Client) BindTemplateWithOptions(request *BindTemplateRequest, runtime *dara.RuntimeOptions) (_result *BindTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyAll) {
		query["ApplyAll"] = request.ApplyAll
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Replace) {
		query["Replace"] = request.Replace
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a template to a specified instance, such as a group or stream.
//
// @param request - BindTemplateRequest
//
// @return BindTemplateResponse
func (client *Client) BindTemplate(request *BindTemplateRequest) (_result *BindTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindTemplateResponse{}
	_body, _err := client.BindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a Comfy task.
//
// Description:
//
// > Stop the parent platform before canceling the task.
//
// @param request - CancelComfyTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelComfyTaskResponse
func (client *Client) CancelComfyTaskWithOptions(request *CancelComfyTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelComfyTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelComfyTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelComfyTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a Comfy task.
//
// Description:
//
// > Stop the parent platform before canceling the task.
//
// @param request - CancelComfyTaskRequest
//
// @return CancelComfyTaskResponse
func (client *Client) CancelComfyTask(request *CancelComfyTaskRequest) (_result *CancelComfyTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelComfyTaskResponse{}
	_body, _err := client.CancelComfyTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Continuously adjust lens parameters such as aperture and zoom.
//
// @param request - ContinuousAdjustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinuousAdjustResponse
func (client *Client) ContinuousAdjustWithOptions(request *ContinuousAdjustRequest, runtime *dara.RuntimeOptions) (_result *ContinuousAdjustResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Focus) {
		query["Focus"] = request.Focus
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Iris) {
		query["Iris"] = request.Iris
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinuousAdjust"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinuousAdjustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Continuously adjust lens parameters such as aperture and zoom.
//
// @param request - ContinuousAdjustRequest
//
// @return ContinuousAdjustResponse
func (client *Client) ContinuousAdjust(request *ContinuousAdjustRequest) (_result *ContinuousAdjustResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinuousAdjustResponse{}
	_body, _err := client.ContinuousAdjustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rotate the camera continuously by panning, tilting, or zooming.
//
// @param request - ContinuousMoveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinuousMoveResponse
func (client *Client) ContinuousMoveWithOptions(request *ContinuousMoveRequest, runtime *dara.RuntimeOptions) (_result *ContinuousMoveResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Pan) {
		query["Pan"] = request.Pan
	}

	if !dara.IsNil(request.Tilt) {
		query["Tilt"] = request.Tilt
	}

	if !dara.IsNil(request.Zoom) {
		query["Zoom"] = request.Zoom
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinuousMove"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinuousMoveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rotate the camera continuously by panning, tilting, or zooming.
//
// @param request - ContinuousMoveRequest
//
// @return ContinuousMoveResponse
func (client *Client) ContinuousMove(request *ContinuousMoveRequest) (_result *ContinuousMoveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinuousMoveResponse{}
	_body, _err := client.ContinuousMoveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a Comfy task.
//
// Description:
//
// > You must first enable the on-demand screenshot feature in the associated screenshot template.
//
// @param request - CreateComfyTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComfyTaskResponse
func (client *Client) CreateComfyTaskWithOptions(request *CreateComfyTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateComfyTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HiveId) {
		query["HiveId"] = request.HiveId
	}

	if !dara.IsNil(request.UserParameters) {
		query["UserParameters"] = request.UserParameters
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComfyTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComfyTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a Comfy task.
//
// Description:
//
// > You must first enable the on-demand screenshot feature in the associated screenshot template.
//
// @param request - CreateComfyTaskRequest
//
// @return CreateComfyTaskResponse
func (client *Client) CreateComfyTask(request *CreateComfyTaskRequest) (_result *CreateComfyTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComfyTaskResponse{}
	_body, _err := client.CreateComfyTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a directory for user data.
//
// Description:
//
// > You must specify either a template ID or a template type.
//
// @param request - CreateComfyUserDataDirRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComfyUserDataDirResponse
func (client *Client) CreateComfyUserDataDirWithOptions(request *CreateComfyUserDataDirRequest, runtime *dara.RuntimeOptions) (_result *CreateComfyUserDataDirResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComfyUserDataDir"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComfyUserDataDirResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory for user data.
//
// Description:
//
// > You must specify either a template ID or a template type.
//
// @param request - CreateComfyUserDataDirRequest
//
// @return CreateComfyUserDataDirResponse
func (client *Client) CreateComfyUserDataDir(request *CreateComfyUserDataDirRequest) (_result *CreateComfyUserDataDirResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComfyUserDataDirResponse{}
	_body, _err := client.CreateComfyUserDataDirWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Comfy workflow.
//
// Description:
//
// > You must enable the on-demand screenshot feature in the associated screenshot template before calling this operation.
//
// @param request - CreateComfyWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComfyWorkflowResponse
func (client *Client) CreateComfyWorkflowWithOptions(request *CreateComfyWorkflowRequest, runtime *dara.RuntimeOptions) (_result *CreateComfyWorkflowResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Workflow) {
		query["Workflow"] = request.Workflow
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComfyWorkflow"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComfyWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Comfy workflow.
//
// Description:
//
// > You must enable the on-demand screenshot feature in the associated screenshot template before calling this operation.
//
// @param request - CreateComfyWorkflowRequest
//
// @return CreateComfyWorkflowResponse
func (client *Client) CreateComfyWorkflow(request *CreateComfyWorkflowRequest) (_result *CreateComfyWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComfyWorkflowResponse{}
	_body, _err := client.CreateComfyWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a new device.
//
// @param request - CreateDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeviceResponse
func (client *Client) CreateDeviceWithOptions(request *CreateDeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmMethod) {
		query["AlarmMethod"] = request.AlarmMethod
	}

	if !dara.IsNil(request.AutoDirectory) {
		query["AutoDirectory"] = request.AutoDirectory
	}

	if !dara.IsNil(request.AutoPos) {
		query["AutoPos"] = request.AutoPos
	}

	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Dsn) {
		query["Dsn"] = request.Dsn
	}

	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Latitude) {
		query["Latitude"] = request.Latitude
	}

	if !dara.IsNil(request.Longitude) {
		query["Longitude"] = request.Longitude
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PosInterval) {
		query["PosInterval"] = request.PosInterval
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.Vendor) {
		query["Vendor"] = request.Vendor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a new device.
//
// @param request - CreateDeviceRequest
//
// @return CreateDeviceResponse
func (client *Client) CreateDevice(request *CreateDeviceRequest) (_result *CreateDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDeviceResponse{}
	_body, _err := client.CreateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reports a device alert.
//
// @param request - CreateDeviceAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeviceAlarmResponse
func (client *Client) CreateDeviceAlarmWithOptions(request *CreateDeviceAlarmRequest, runtime *dara.RuntimeOptions) (_result *CreateDeviceAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Alarm) {
		query["Alarm"] = request.Alarm
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubAlarm) {
		query["SubAlarm"] = request.SubAlarm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeviceAlarm"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeviceAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports a device alert.
//
// @param request - CreateDeviceAlarmRequest
//
// @return CreateDeviceAlarmResponse
func (client *Client) CreateDeviceAlarm(request *CreateDeviceAlarmRequest) (_result *CreateDeviceAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDeviceAlarmResponse{}
	_body, _err := client.CreateDeviceAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new folder.
//
// @param request - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithOptions(request *CreateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDirectory"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new folder.
//
// @param request - CreateDirectoryRequest
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectory(request *CreateDirectoryRequest) (_result *CreateDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CreateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a new workspace.
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InProtocol) {
		query["InProtocol"] = request.InProtocol
	}

	if !dara.IsNil(request.LazyPull) {
		query["LazyPull"] = request.LazyPull
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutProtocol) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.PushDomain) {
		query["PushDomain"] = request.PushDomain
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a new workspace.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cluster.
//
// Description:
//
// ## Description
//
// - This operation creates an empty cluster to manage workloads.
//
// @param request - CreateHiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHiveResponse
func (client *Client) CreateHiveWithOptions(request *CreateHiveRequest, runtime *dara.RuntimeOptions) (_result *CreateHiveResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHive"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cluster.
//
// Description:
//
// ## Description
//
// - This operation creates an empty cluster to manage workloads.
//
// @param request - CreateHiveRequest
//
// @return CreateHiveResponse
func (client *Client) CreateHive(request *CreateHiveRequest) (_result *CreateHiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHiveResponse{}
	_body, _err := client.CreateHiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a new parent platform.
//
// @param request - CreateParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParentPlatformResponse
func (client *Client) CreateParentPlatformWithOptions(request *CreateParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *CreateParentPlatformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.ClientAuth) {
		query["ClientAuth"] = request.ClientAuth
	}

	if !dara.IsNil(request.ClientPassword) {
		query["ClientPassword"] = request.ClientPassword
	}

	if !dara.IsNil(request.ClientUsername) {
		query["ClientUsername"] = request.ClientUsername
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParentPlatform"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a new parent platform.
//
// @param request - CreateParentPlatformRequest
//
// @return CreateParentPlatformResponse
func (client *Client) CreateParentPlatform(request *CreateParentPlatformRequest) (_result *CreateParentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateParentPlatformResponse{}
	_body, _err := client.CreateParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data pack for a cloud application service.
//
// @param request - CreateRenderingDataPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingDataPackageResponse
func (client *Client) CreateRenderingDataPackageWithOptions(request *CreateRenderingDataPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingDataPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRenderingDataPackage"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRenderingDataPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data pack for a cloud application service.
//
// @param request - CreateRenderingDataPackageRequest
//
// @return CreateRenderingDataPackageResponse
func (client *Client) CreateRenderingDataPackage(request *CreateRenderingDataPackageRequest) (_result *CreateRenderingDataPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRenderingDataPackageResponse{}
	_body, _err := client.CreateRenderingDataPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call CreateRenderingInstance to create a cloud application service instance.
//
// @param tmpReq - CreateRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingInstanceResponse
func (client *Client) CreateRenderingInstanceWithOptions(tmpReq *CreateRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRenderingInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Attributes) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, dara.String("Attributes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClientInfo) {
		request.ClientInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientInfo, dara.String("ClientInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttributesShrink) {
		query["Attributes"] = request.AttributesShrink
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ClientInfoShrink) {
		query["ClientInfo"] = request.ClientInfoShrink
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.InternetMaxBandwidth) {
		query["InternetMaxBandwidth"] = request.InternetMaxBandwidth
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RenderingSpec) {
		query["RenderingSpec"] = request.RenderingSpec
	}

	if !dara.IsNil(request.StorageSize) {
		query["StorageSize"] = request.StorageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRenderingInstance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRenderingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call CreateRenderingInstance to create a cloud application service instance.
//
// @param request - CreateRenderingInstanceRequest
//
// @return CreateRenderingInstanceResponse
func (client *Client) CreateRenderingInstance(request *CreateRenderingInstanceRequest) (_result *CreateRenderingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRenderingInstanceResponse{}
	_body, _err := client.CreateRenderingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom gateway.
//
// Description:
//
// > You can specify a template ID or a template type.
//
// @param request - CreateRenderingInstanceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingInstanceGatewayResponse
func (client *Client) CreateRenderingInstanceGatewayWithOptions(request *CreateRenderingInstanceGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingInstanceGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayInstanceId) {
		query["GatewayInstanceId"] = request.GatewayInstanceId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRenderingInstanceGateway"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRenderingInstanceGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom gateway.
//
// Description:
//
// > You can specify a template ID or a template type.
//
// @param request - CreateRenderingInstanceGatewayRequest
//
// @return CreateRenderingInstanceGatewayResponse
func (client *Client) CreateRenderingInstanceGateway(request *CreateRenderingInstanceGatewayRequest) (_result *CreateRenderingInstanceGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRenderingInstanceGatewayResponse{}
	_body, _err := client.CreateRenderingInstanceGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cloud application service project and configures its properties, such as session attributes.
//
// @param tmpReq - CreateRenderingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingProjectResponse
func (client *Client) CreateRenderingProjectWithOptions(tmpReq *CreateRenderingProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRenderingProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionAttribs) {
		request.SessionAttribsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionAttribs, dara.String("SessionAttribs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SessionAttribsShrink) {
		query["SessionAttribs"] = request.SessionAttribsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRenderingProject"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRenderingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud application service project and configures its properties, such as session attributes.
//
// @param request - CreateRenderingProjectRequest
//
// @return CreateRenderingProjectResponse
func (client *Client) CreateRenderingProject(request *CreateRenderingProjectRequest) (_result *CreateRenderingProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRenderingProjectResponse{}
	_body, _err := client.CreateRenderingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an on-demand snapshot for the specified stream.
//
// Description:
//
// > You must first enable the on-demand snapshot feature in the attached snapshot template.
//
// @param request - CreateStreamSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamSnapshotResponse
func (client *Client) CreateStreamSnapshotWithOptions(request *CreateStreamSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamSnapshotResponse, _err error) {
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

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamSnapshot"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an on-demand snapshot for the specified stream.
//
// Description:
//
// > You must first enable the on-demand snapshot feature in the attached snapshot template.
//
// @param request - CreateStreamSnapshotRequest
//
// @return CreateStreamSnapshotResponse
func (client *Client) CreateStreamSnapshot(request *CreateStreamSnapshotRequest) (_result *CreateStreamSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStreamSnapshotResponse{}
	_body, _err := client.CreateStreamSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a new template.
//
// @param request - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileFormat) {
		query["FileFormat"] = request.FileFormat
	}

	if !dara.IsNil(request.Flv) {
		query["Flv"] = request.Flv
	}

	if !dara.IsNil(request.HlsM3u8) {
		query["HlsM3u8"] = request.HlsM3u8
	}

	if !dara.IsNil(request.HlsTs) {
		query["HlsTs"] = request.HlsTs
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.JpgOnDemand) {
		query["JpgOnDemand"] = request.JpgOnDemand
	}

	if !dara.IsNil(request.JpgOverwrite) {
		query["JpgOverwrite"] = request.JpgOverwrite
	}

	if !dara.IsNil(request.JpgSequence) {
		query["JpgSequence"] = request.JpgSequence
	}

	if !dara.IsNil(request.Mp4) {
		query["Mp4"] = request.Mp4
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssFilePrefix) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	if !dara.IsNil(request.TransConfigsJSON) {
		query["TransConfigsJSON"] = request.TransConfigsJSON
	}

	if !dara.IsNil(request.Trigger) {
		query["Trigger"] = request.Trigger
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a new template.
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从集群删除负载
//
// Description:
//
// ## 请求说明
//
// - **HiveId*	- 是必填参数，表示要操作的集群ID。
//
// - **InstanceIds*	- 是必填参数，需要提供一个负载ID列表，用于指定要从集群中解绑的负载实例。
//
// - 解绑操作成功后，会返回成功和失败的负载实例列表及其相关信息。
//
// @param tmpReq - DelHiveEdgeWorkersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelHiveEdgeWorkersResponse
func (client *Client) DelHiveEdgeWorkersWithOptions(tmpReq *DelHiveEdgeWorkersRequest, runtime *dara.RuntimeOptions) (_result *DelHiveEdgeWorkersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DelHiveEdgeWorkersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HiveId) {
		query["HiveId"] = request.HiveId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelHiveEdgeWorkers"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelHiveEdgeWorkersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从集群删除负载
//
// Description:
//
// ## 请求说明
//
// - **HiveId*	- 是必填参数，表示要操作的集群ID。
//
// - **InstanceIds*	- 是必填参数，需要提供一个负载ID列表，用于指定要从集群中解绑的负载实例。
//
// - 解绑操作成功后，会返回成功和失败的负载实例列表及其相关信息。
//
// @param request - DelHiveEdgeWorkersRequest
//
// @return DelHiveEdgeWorkersResponse
func (client *Client) DelHiveEdgeWorkers(request *DelHiveEdgeWorkersRequest) (_result *DelHiveEdgeWorkersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelHiveEdgeWorkersResponse{}
	_body, _err := client.DelHiveEdgeWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a cloud application. You cannot delete a cloud application that is in use.
//
// @param request - DeleteCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudAppResponse
func (client *Client) DeleteCloudAppWithOptions(request *DeleteCloudAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudAppResponse, _err error) {
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
		Action:      dara.String("DeleteCloudApp"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cloud application. You cannot delete a cloud application that is in use.
//
// @param request - DeleteCloudAppRequest
//
// @return DeleteCloudAppResponse
func (client *Client) DeleteCloudApp(request *DeleteCloudAppRequest) (_result *DeleteCloudAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudAppResponse{}
	_body, _err := client.DeleteCloudAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Deleting artifacts
//
// Description:
//
// > Stop the parent platform before you delete a production.
//
// @param request - DeleteComfyProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComfyProductionResponse
func (client *Client) DeleteComfyProductionWithOptions(request *DeleteComfyProductionRequest, runtime *dara.RuntimeOptions) (_result *DeleteComfyProductionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductionId) {
		query["ProductionId"] = request.ProductionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComfyProduction"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComfyProductionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Deleting artifacts
//
// Description:
//
// > Stop the parent platform before you delete a production.
//
// @param request - DeleteComfyProductionRequest
//
// @return DeleteComfyProductionResponse
func (client *Client) DeleteComfyProduction(request *DeleteComfyProductionRequest) (_result *DeleteComfyProductionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComfyProductionResponse{}
	_body, _err := client.DeleteComfyProductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file or directory from user data.
//
// Description:
//
// > You must stop the upper-level platform before performing this operation.
//
// @param request - DeleteComfyUserDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComfyUserDataResponse
func (client *Client) DeleteComfyUserDataWithOptions(request *DeleteComfyUserDataRequest, runtime *dara.RuntimeOptions) (_result *DeleteComfyUserDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComfyUserData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComfyUserDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file or directory from user data.
//
// Description:
//
// > You must stop the upper-level platform before performing this operation.
//
// @param request - DeleteComfyUserDataRequest
//
// @return DeleteComfyUserDataResponse
func (client *Client) DeleteComfyUserData(request *DeleteComfyUserDataRequest) (_result *DeleteComfyUserDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComfyUserDataResponse{}
	_body, _err := client.DeleteComfyUserDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Comfy workflow.
//
// Description:
//
// > You must stop the parent platform before you can delete the workflow.
//
// @param request - DeleteComfyWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComfyWorkflowResponse
func (client *Client) DeleteComfyWorkflowWithOptions(request *DeleteComfyWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteComfyWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComfyWorkflow"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComfyWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Comfy workflow.
//
// Description:
//
// > You must stop the parent platform before you can delete the workflow.
//
// @param request - DeleteComfyWorkflowRequest
//
// @return DeleteComfyWorkflowResponse
func (client *Client) DeleteComfyWorkflow(request *DeleteComfyWorkflowRequest) (_result *DeleteComfyWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComfyWorkflowResponse{}
	_body, _err := client.DeleteComfyWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a device from a space.
//
// @param request - DeleteDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeviceResponse
func (client *Client) DeleteDeviceWithOptions(request *DeleteDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeviceResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a device from a space.
//
// @param request - DeleteDeviceRequest
//
// @return DeleteDeviceResponse
func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (_result *DeleteDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DeleteDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a folder.
//
// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDirectory"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a folder.
//
// @param request - DeleteDirectoryRequest
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectory(request *DeleteDirectoryRequest) (_result *DeleteDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.DeleteDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You cannot delete a file while it is uploading or pre-pushing. After deletion, all related push records become invalid. You can push a file with the same name again.
//
// @param request - DeleteFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFile"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You cannot delete a file while it is uploading or pre-pushing. After deletion, all related push records become invalid. You can push a file with the same name again.
//
// @param request - DeleteFileRequest
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a workspace.
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroup"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a workspace.
//
// @param request - DeleteGroupRequest
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除集群
//
// Description:
//
// ## 请求说明
//
// - 需要确保该集群内所有应用服务已清空，否则无法执行删除操作。
//
// - `HiveId` 是必填参数，用于标识待删除的集群。
//
// @param request - DeleteHiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHiveResponse
func (client *Client) DeleteHiveWithOptions(request *DeleteHiveRequest, runtime *dara.RuntimeOptions) (_result *DeleteHiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HiveId) {
		query["HiveId"] = request.HiveId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHive"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除集群
//
// Description:
//
// ## 请求说明
//
// - 需要确保该集群内所有应用服务已清空，否则无法执行删除操作。
//
// - `HiveId` 是必填参数，用于标识待删除的集群。
//
// @param request - DeleteHiveRequest
//
// @return DeleteHiveResponse
func (client *Client) DeleteHive(request *DeleteHiveRequest) (_result *DeleteHiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHiveResponse{}
	_body, _err := client.DeleteHiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a parent platform.
//
// Description:
//
// > You must stop the parent platform before you delete it.
//
// @param request - DeleteParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParentPlatformResponse
func (client *Client) DeleteParentPlatformWithOptions(request *DeleteParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *DeleteParentPlatformResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParentPlatform"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a parent platform.
//
// Description:
//
// > You must stop the parent platform before you delete it.
//
// @param request - DeleteParentPlatformRequest
//
// @return DeleteParentPlatformResponse
func (client *Client) DeleteParentPlatform(request *DeleteParentPlatformRequest) (_result *DeleteParentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteParentPlatformResponse{}
	_body, _err := client.DeleteParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a preset.
//
// @param request - DeletePresetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePresetResponse
func (client *Client) DeletePresetWithOptions(request *DeletePresetRequest, runtime *dara.RuntimeOptions) (_result *DeletePresetResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PresetId) {
		query["PresetId"] = request.PresetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePreset"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePresetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a preset.
//
// @param request - DeletePresetRequest
//
// @return DeletePresetResponse
func (client *Client) DeletePreset(request *DeletePresetRequest) (_result *DeletePresetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePresetResponse{}
	_body, _err := client.DeletePresetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified public key. This action automatically revokes logon authorization for all associated cloud application service instances.
//
// @param request - DeletePublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePublicKeyResponse
func (client *Client) DeletePublicKeyWithOptions(request *DeletePublicKeyRequest, runtime *dara.RuntimeOptions) (_result *DeletePublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePublicKey"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified public key. This action automatically revokes logon authorization for all associated cloud application service instances.
//
// @param request - DeletePublicKeyRequest
//
// @return DeletePublicKeyResponse
func (client *Client) DeletePublicKey(request *DeletePublicKeyRequest) (_result *DeletePublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePublicKeyResponse{}
	_body, _err := client.DeletePublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the configuration of a cloud application service instance. This operation deletes only module properties that are configured using the UpdateRenderingInstanceConfiguration operation.
//
// @param tmpReq - DeleteRenderingInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingInstanceConfigurationResponse
func (client *Client) DeleteRenderingInstanceConfigurationWithOptions(tmpReq *DeleteRenderingInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingInstanceConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRenderingInstanceConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configuration) {
		request.ConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configuration, dara.String("Configuration"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationShrink) {
		body["Configuration"] = request.ConfigurationShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRenderingInstanceConfiguration"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRenderingInstanceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configuration of a cloud application service instance. This operation deletes only module properties that are configured using the UpdateRenderingInstanceConfiguration operation.
//
// @param request - DeleteRenderingInstanceConfigurationRequest
//
// @return DeleteRenderingInstanceConfigurationResponse
func (client *Client) DeleteRenderingInstanceConfiguration(request *DeleteRenderingInstanceConfigurationRequest) (_result *DeleteRenderingInstanceConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRenderingInstanceConfigurationResponse{}
	_body, _err := client.DeleteRenderingInstanceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom gateway.
//
// Description:
//
// > Stop the parent platform before you delete the gateway.
//
// @param request - DeleteRenderingInstanceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingInstanceGatewayResponse
func (client *Client) DeleteRenderingInstanceGatewayWithOptions(request *DeleteRenderingInstanceGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingInstanceGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRenderingInstanceGateway"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRenderingInstanceGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom gateway.
//
// Description:
//
// > Stop the parent platform before you delete the gateway.
//
// @param request - DeleteRenderingInstanceGatewayRequest
//
// @return DeleteRenderingInstanceGatewayResponse
func (client *Client) DeleteRenderingInstanceGateway(request *DeleteRenderingInstanceGatewayRequest) (_result *DeleteRenderingInstanceGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRenderingInstanceGatewayResponse{}
	_body, _err := client.DeleteRenderingInstanceGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call DeleteRenderingInstanceSettings to delete the settings of a cloud application service instance.
//
// @param tmpReq - DeleteRenderingInstanceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingInstanceSettingsResponse
func (client *Client) DeleteRenderingInstanceSettingsWithOptions(tmpReq *DeleteRenderingInstanceSettingsRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingInstanceSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRenderingInstanceSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AttributeNames) {
		request.AttributeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttributeNames, dara.String("AttributeNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttributeNamesShrink) {
		query["AttributeNames"] = request.AttributeNamesShrink
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRenderingInstanceSettings"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRenderingInstanceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call DeleteRenderingInstanceSettings to delete the settings of a cloud application service instance.
//
// @param request - DeleteRenderingInstanceSettingsRequest
//
// @return DeleteRenderingInstanceSettingsResponse
func (client *Client) DeleteRenderingInstanceSettings(request *DeleteRenderingInstanceSettingsRequest) (_result *DeleteRenderingInstanceSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRenderingInstanceSettingsResponse{}
	_body, _err := client.DeleteRenderingInstanceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a Data Service Project. Projects that have business scheduling data, such as active sessions, cannot be deleted.
//
// @param request - DeleteRenderingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingProjectResponse
func (client *Client) DeleteRenderingProjectWithOptions(request *DeleteRenderingProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRenderingProject"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRenderingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a Data Service Project. Projects that have business scheduling data, such as active sessions, cannot be deleted.
//
// @param request - DeleteRenderingProjectRequest
//
// @return DeleteRenderingProjectResponse
func (client *Client) DeleteRenderingProject(request *DeleteRenderingProjectRequest) (_result *DeleteRenderingProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRenderingProjectResponse{}
	_body, _err := client.DeleteRenderingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete stream pulling information.
//
// @param request - DeleteVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVsPullStreamInfoConfigResponse
func (client *Client) DeleteVsPullStreamInfoConfigWithOptions(request *DeleteVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteVsPullStreamInfoConfigResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVsPullStreamInfoConfig"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete stream pulling information.
//
// @param request - DeleteVsPullStreamInfoConfigRequest
//
// @return DeleteVsPullStreamInfoConfigResponse
func (client *Client) DeleteVsPullStreamInfoConfig(request *DeleteVsPullStreamInfoConfigRequest) (_result *DeleteVsPullStreamInfoConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVsPullStreamInfoConfigResponse{}
	_body, _err := client.DeleteVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the callback configuration for stream ingest.
//
// @param request - DeleteVsStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVsStreamsNotifyUrlConfigResponse
func (client *Client) DeleteVsStreamsNotifyUrlConfigWithOptions(request *DeleteVsStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteVsStreamsNotifyUrlConfigResponse, _err error) {
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
		Action:      dara.String("DeleteVsStreamsNotifyUrlConfig"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the callback configuration for stream ingest.
//
// @param request - DeleteVsStreamsNotifyUrlConfigRequest
//
// @return DeleteVsStreamsNotifyUrlConfigResponse
func (client *Client) DeleteVsStreamsNotifyUrlConfig(request *DeleteVsStreamsNotifyUrlConfigRequest) (_result *DeleteVsStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DeleteVsStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query all resource information for an account in a specified region.
//
// @param request - DescribeAccountStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountStatResponse
func (client *Client) DescribeAccountStatWithOptions(request *DescribeAccountStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountStatResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountStat"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query all resource information for an account in a specified region.
//
// @param request - DescribeAccountStatRequest
//
// @return DescribeAccountStatResponse
func (client *Client) DescribeAccountStat(request *DescribeAccountStatRequest) (_result *DescribeAccountStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountStatResponse{}
	_body, _err := client.DescribeAccountStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a download link for a production.
//
// Description:
//
// > Screenshot queries do not support pagination and must be performed iteratively. To fetch the next page, use the extStartTime value from the response as the StartTime for your subsequent request.
//
// @param request - DescribeComfyProductionDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComfyProductionDownloadUrlResponse
func (client *Client) DescribeComfyProductionDownloadUrlWithOptions(request *DescribeComfyProductionDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeComfyProductionDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductionId) {
		query["ProductionId"] = request.ProductionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComfyProductionDownloadUrl"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComfyProductionDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a download link for a production.
//
// Description:
//
// > Screenshot queries do not support pagination and must be performed iteratively. To fetch the next page, use the extStartTime value from the response as the StartTime for your subsequent request.
//
// @param request - DescribeComfyProductionDownloadUrlRequest
//
// @return DescribeComfyProductionDownloadUrlResponse
func (client *Client) DescribeComfyProductionDownloadUrl(request *DescribeComfyProductionDownloadUrlRequest) (_result *DescribeComfyProductionDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComfyProductionDownloadUrlResponse{}
	_body, _err := client.DescribeComfyProductionDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists a user\\"s productions.
//
// Description:
//
// > This API uses pagination. Use the PageNumber and PageSize parameters to navigate through the results.
//
// @param request - DescribeComfyProductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComfyProductionsResponse
func (client *Client) DescribeComfyProductionsWithOptions(request *DescribeComfyProductionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeComfyProductionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComfyProductions"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComfyProductionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists a user\\"s productions.
//
// Description:
//
// > This API uses pagination. Use the PageNumber and PageSize parameters to navigate through the results.
//
// @param request - DescribeComfyProductionsRequest
//
// @return DescribeComfyProductionsResponse
func (client *Client) DescribeComfyProductions(request *DescribeComfyProductionsRequest) (_result *DescribeComfyProductionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComfyProductionsResponse{}
	_body, _err := client.DescribeComfyProductionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Comfy tasks.
//
// Description:
//
// > Querying by screenshot does not support pagination and only supports iteration. To request the next page, use the extStartTime parameter value from the response as the StartTime for the new request.
//
// @param request - DescribeComfyTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComfyTasksResponse
func (client *Client) DescribeComfyTasksWithOptions(request *DescribeComfyTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeComfyTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskState) {
		query["TaskState"] = request.TaskState
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComfyTasks"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComfyTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Comfy tasks.
//
// Description:
//
// > Querying by screenshot does not support pagination and only supports iteration. To request the next page, use the extStartTime parameter value from the response as the StartTime for the new request.
//
// @param request - DescribeComfyTasksRequest
//
// @return DescribeComfyTasksResponse
func (client *Client) DescribeComfyTasks(request *DescribeComfyTasksRequest) (_result *DescribeComfyTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComfyTasksResponse{}
	_body, _err := client.DescribeComfyTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets a download URL for user data.
//
// @param request - DescribeComfyUserDataDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComfyUserDataDownloadUrlResponse
func (client *Client) DescribeComfyUserDataDownloadUrlWithOptions(request *DescribeComfyUserDataDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeComfyUserDataDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComfyUserDataDownloadUrl"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComfyUserDataDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a download URL for user data.
//
// @param request - DescribeComfyUserDataDownloadUrlRequest
//
// @return DescribeComfyUserDataDownloadUrlResponse
func (client *Client) DescribeComfyUserDataDownloadUrl(request *DescribeComfyUserDataDownloadUrlRequest) (_result *DescribeComfyUserDataDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComfyUserDataDownloadUrlResponse{}
	_body, _err := client.DescribeComfyUserDataDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a URL to upload a user file.
//
// Description:
//
// You can upload files using the retrieved URL and the Alibaba Cloud OSS software development kit (SDK).
//
// @param request - DescribeComfyUserDataUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComfyUserDataUploadUrlResponse
func (client *Client) DescribeComfyUserDataUploadUrlWithOptions(request *DescribeComfyUserDataUploadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeComfyUserDataUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.FileMd5) {
		query["FileMd5"] = request.FileMd5
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileSizeBytes) {
		query["FileSizeBytes"] = request.FileSizeBytes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComfyUserDataUploadUrl"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComfyUserDataUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a URL to upload a user file.
//
// Description:
//
// You can upload files using the retrieved URL and the Alibaba Cloud OSS software development kit (SDK).
//
// @param request - DescribeComfyUserDataUploadUrlRequest
//
// @return DescribeComfyUserDataUploadUrlResponse
func (client *Client) DescribeComfyUserDataUploadUrl(request *DescribeComfyUserDataUploadUrlRequest) (_result *DescribeComfyUserDataUploadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComfyUserDataUploadUrlResponse{}
	_body, _err := client.DescribeComfyUserDataUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all files and folders in the user data.
//
// Description:
//
// If StartTime and EndTime are not specified, data from the last 24 hours is read by default. To query a specific time range, you must specify both StartTime and EndTime. The maximum time range for a query is 31 days.
//
// - You can query multiple domain names in a batch. Separate the domain names with a comma (,).
//
// - You can retrieve data from the last 90 days.
//
// - The time granularity is one hour.
//
// @param request - DescribeComfyUserDatasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComfyUserDatasResponse
func (client *Client) DescribeComfyUserDatasWithOptions(request *DescribeComfyUserDatasRequest, runtime *dara.RuntimeOptions) (_result *DescribeComfyUserDatasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
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
		Action:      dara.String("DescribeComfyUserDatas"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComfyUserDatasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all files and folders in the user data.
//
// Description:
//
// If StartTime and EndTime are not specified, data from the last 24 hours is read by default. To query a specific time range, you must specify both StartTime and EndTime. The maximum time range for a query is 31 days.
//
// - You can query multiple domain names in a batch. Separate the domain names with a comma (,).
//
// - You can retrieve data from the last 90 days.
//
// - The time granularity is one hour.
//
// @param request - DescribeComfyUserDatasRequest
//
// @return DescribeComfyUserDatasResponse
func (client *Client) DescribeComfyUserDatas(request *DescribeComfyUserDatasRequest) (_result *DescribeComfyUserDatasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComfyUserDatasResponse{}
	_body, _err := client.DescribeComfyUserDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation retrieves a paginated list of your Comfy workflows.
//
// Description:
//
// \\> 截图查询目前不支持分页，仅支持按迭代方式。使用返回结果里的extStartTime参数值，作为新请求的StartTime可请求下一页。
//
// @param request - DescribeComfyWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComfyWorkflowsResponse
func (client *Client) DescribeComfyWorkflowsWithOptions(request *DescribeComfyWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *DescribeComfyWorkflowsResponse, _err error) {
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
		Action:      dara.String("DescribeComfyWorkflows"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComfyWorkflowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation retrieves a paginated list of your Comfy workflows.
//
// Description:
//
// \\> 截图查询目前不支持分页，仅支持按迭代方式。使用返回结果里的extStartTime参数值，作为新请求的StartTime可请求下一页。
//
// @param request - DescribeComfyWorkflowsRequest
//
// @return DescribeComfyWorkflowsResponse
func (client *Client) DescribeComfyWorkflows(request *DescribeComfyWorkflowsRequest) (_result *DescribeComfyWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeComfyWorkflowsResponse{}
	_body, _err := client.DescribeComfyWorkflowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query information about a device.
//
// @param request - DescribeDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceResponse
func (client *Client) DescribeDeviceWithOptions(request *DescribeDeviceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceResponse, _err error) {
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

	if !dara.IsNil(request.IncludeDirectory) {
		query["IncludeDirectory"] = request.IncludeDirectory
	}

	if !dara.IsNil(request.IncludeStats) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query information about a device.
//
// @param request - DescribeDeviceRequest
//
// @return DescribeDeviceResponse
func (client *Client) DescribeDevice(request *DescribeDeviceRequest) (_result *DescribeDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeviceResponse{}
	_body, _err := client.DescribeDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of device channels.
//
// @param request - DescribeDeviceChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceChannelsResponse
func (client *Client) DescribeDeviceChannelsWithOptions(request *DescribeDeviceChannelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceChannelsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceChannels"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of device channels.
//
// @param request - DescribeDeviceChannelsRequest
//
// @return DescribeDeviceChannelsResponse
func (client *Client) DescribeDeviceChannels(request *DescribeDeviceChannelsRequest) (_result *DescribeDeviceChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeviceChannelsResponse{}
	_body, _err := client.DescribeDeviceChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a device gateway.
//
// @param request - DescribeDeviceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceGatewayResponse
func (client *Client) DescribeDeviceGatewayWithOptions(request *DescribeDeviceGatewayRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceGateway"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a device gateway.
//
// @param request - DescribeDeviceGatewayRequest
//
// @return DescribeDeviceGatewayResponse
func (client *Client) DescribeDeviceGateway(request *DescribeDeviceGatewayRequest) (_result *DescribeDeviceGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeviceGatewayResponse{}
	_body, _err := client.DescribeDeviceGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the URL information for a device stream.
//
// @param request - DescribeDeviceURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceURLResponse
func (client *Client) DescribeDeviceURLWithOptions(request *DescribeDeviceURLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceURLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		query["Auth"] = request.Auth
	}

	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OutProtocol) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceURL"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the URL information for a device stream.
//
// @param request - DescribeDeviceURLRequest
//
// @return DescribeDeviceURLResponse
func (client *Client) DescribeDeviceURL(request *DescribeDeviceURLRequest) (_result *DescribeDeviceURLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeviceURLResponse{}
	_body, _err := client.DescribeDeviceURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of devices.
//
// @param request - DescribeDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDevicesResponse
func (client *Client) DescribeDevicesWithOptions(request *DescribeDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Dsn) {
		query["Dsn"] = request.Dsn
	}

	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.IncludeDirectory) {
		query["IncludeDirectory"] = request.IncludeDirectory
	}

	if !dara.IsNil(request.IncludeStats) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Vendor) {
		query["Vendor"] = request.Vendor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of devices.
//
// @param request - DescribeDevicesRequest
//
// @return DescribeDevicesResponse
func (client *Client) DescribeDevices(request *DescribeDevicesRequest) (_result *DescribeDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.DescribeDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of directories.
//
// @param request - DescribeDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NoPagination) {
		query["NoPagination"] = request.NoPagination
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDirectories"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of directories.
//
// @param request - DescribeDirectoriesRequest
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a directory.
//
// @param request - DescribeDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoryResponse
func (client *Client) DescribeDirectoryWithOptions(request *DescribeDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirectoryResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDirectory"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a directory.
//
// @param request - DescribeDirectoryRequest
//
// @return DescribeDirectoryResponse
func (client *Client) DescribeDirectory(request *DescribeDirectoryRequest) (_result *DescribeDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDirectoryResponse{}
	_body, _err := client.DescribeDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about a space.
//
// @param request - DescribeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupResponse
func (client *Client) DescribeGroupWithOptions(request *DescribeGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupResponse, _err error) {
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

	if !dara.IsNil(request.IncludeStats) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroup"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a space.
//
// @param request - DescribeGroupRequest
//
// @return DescribeGroupResponse
func (client *Client) DescribeGroup(request *DescribeGroupRequest) (_result *DescribeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupResponse{}
	_body, _err := client.DescribeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can query the list of spaces.
//
// @param request - DescribeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroupsWithOptions(request *DescribeGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupsResponse, _err error) {
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

	if !dara.IsNil(request.InProtocol) {
		query["InProtocol"] = request.InProtocol
	}

	if !dara.IsNil(request.IncludeStats) {
		query["IncludeStats"] = request.IncludeStats
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroups"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can query the list of spaces.
//
// @param request - DescribeGroupsRequest
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroups(request *DescribeGroupsRequest) (_result *DescribeGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.DescribeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a parent platform.
//
// @param request - DescribeParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentPlatformResponse
func (client *Client) DescribeParentPlatformWithOptions(request *DescribeParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentPlatformResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParentPlatform"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a parent platform.
//
// @param request - DescribeParentPlatformRequest
//
// @return DescribeParentPlatformResponse
func (client *Client) DescribeParentPlatform(request *DescribeParentPlatformRequest) (_result *DescribeParentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParentPlatformResponse{}
	_body, _err := client.DescribeParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of devices under a parent platform.
//
// @param request - DescribeParentPlatformDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentPlatformDevicesResponse
func (client *Client) DescribeParentPlatformDevicesWithOptions(request *DescribeParentPlatformDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentPlatformDevicesResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParentPlatformDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParentPlatformDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of devices under a parent platform.
//
// @param request - DescribeParentPlatformDevicesRequest
//
// @return DescribeParentPlatformDevicesResponse
func (client *Client) DescribeParentPlatformDevices(request *DescribeParentPlatformDevicesRequest) (_result *DescribeParentPlatformDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParentPlatformDevicesResponse{}
	_body, _err := client.DescribeParentPlatformDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of parent platforms.
//
// @param request - DescribeParentPlatformsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentPlatformsResponse
func (client *Client) DescribeParentPlatformsWithOptions(request *DescribeParentPlatformsRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentPlatformsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParentPlatforms"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParentPlatformsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of parent platforms.
//
// @param request - DescribeParentPlatformsRequest
//
// @return DescribeParentPlatformsResponse
func (client *Client) DescribeParentPlatforms(request *DescribeParentPlatformsRequest) (_result *DescribeParentPlatformsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParentPlatformsResponse{}
	_body, _err := client.DescribeParentPlatformsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the list of presets.
//
// @param request - DescribePresetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePresetsResponse
func (client *Client) DescribePresetsWithOptions(request *DescribePresetsRequest, runtime *dara.RuntimeOptions) (_result *DescribePresetsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePresets"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePresetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of presets.
//
// @param request - DescribePresetsRequest
//
// @return DescribePresetsResponse
func (client *Client) DescribePresets(request *DescribePresetsRequest) (_result *DescribePresetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePresetsResponse{}
	_body, _err := client.DescribePresetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribePublishStreamStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePublishStreamStatusResponse
func (client *Client) DescribePublishStreamStatusWithOptions(request *DescribePublishStreamStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePublishStreamStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePublishStreamStatus"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePublishStreamStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePublishStreamStatusRequest
//
// @return DescribePublishStreamStatusResponse
func (client *Client) DescribePublishStreamStatus(request *DescribePublishStreamStatusRequest) (_result *DescribePublishStreamStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePublishStreamStatusResponse{}
	_body, _err := client.DescribePublishStreamStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about purchased devices.
//
// @param request - DescribePurchasedDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedDeviceResponse
func (client *Client) DescribePurchasedDeviceWithOptions(request *DescribePurchasedDeviceRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedDeviceResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about purchased devices.
//
// @param request - DescribePurchasedDeviceRequest
//
// @return DescribePurchasedDeviceResponse
func (client *Client) DescribePurchasedDevice(request *DescribePurchasedDeviceRequest) (_result *DescribePurchasedDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePurchasedDeviceResponse{}
	_body, _err := client.DescribePurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of purchased devices.
//
// @param request - DescribePurchasedDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedDevicesResponse
func (client *Client) DescribePurchasedDevicesWithOptions(request *DescribePurchasedDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.SubType) {
		query["SubType"] = request.SubType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Vendor) {
		query["Vendor"] = request.Vendor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedDevices"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of purchased devices.
//
// @param request - DescribePurchasedDevicesRequest
//
// @return DescribePurchasedDevicesResponse
func (client *Client) DescribePurchasedDevices(request *DescribePurchasedDevicesRequest) (_result *DescribePurchasedDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePurchasedDevicesResponse{}
	_body, _err := client.DescribePurchasedDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of recordings.
//
// Description:
//
// > Paging is not supported for snapshot queries. Only iteration is supported. To request the next page, use the NextStartTime value from the response as the StartTime for the new request.
//
// @param request - DescribeRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordsResponse
func (client *Client) DescribeRecordsWithOptions(request *DescribeRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordsResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateBucket) {
		query["PrivateBucket"] = request.PrivateBucket
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamId) {
		query["StreamId"] = request.StreamId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecords"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of recordings.
//
// Description:
//
// > Paging is not supported for snapshot queries. Only iteration is supported. To request the next page, use the NextStartTime value from the response as the StartTime for the new request.
//
// @param request - DescribeRecordsRequest
//
// @return DescribeRecordsResponse
func (client *Client) DescribeRecords(request *DescribeRecordsRequest) (_result *DescribeRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordsResponse{}
	_body, _err := client.DescribeRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a cloud application service instance.
//
// @param request - DescribeRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingInstanceResponse
func (client *Client) DescribeRenderingInstanceWithOptions(request *DescribeRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRenderingInstance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRenderingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cloud application service instance.
//
// @param request - DescribeRenderingInstanceRequest
//
// @return DescribeRenderingInstanceResponse
func (client *Client) DescribeRenderingInstance(request *DescribeRenderingInstanceRequest) (_result *DescribeRenderingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRenderingInstanceResponse{}
	_body, _err := client.DescribeRenderingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time configuration of a cloud application service instance.
//
// @param tmpReq - DescribeRenderingInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingInstanceConfigurationResponse
func (client *Client) DescribeRenderingInstanceConfigurationWithOptions(tmpReq *DescribeRenderingInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingInstanceConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRenderingInstanceConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configuration) {
		request.ConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configuration, dara.String("Configuration"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRenderingInstanceConfiguration"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRenderingInstanceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time configuration of a cloud application service instance.
//
// @param request - DescribeRenderingInstanceConfigurationRequest
//
// @return DescribeRenderingInstanceConfigurationResponse
func (client *Client) DescribeRenderingInstanceConfiguration(request *DescribeRenderingInstanceConfigurationRequest) (_result *DescribeRenderingInstanceConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRenderingInstanceConfigurationResponse{}
	_body, _err := client.DescribeRenderingInstanceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// DescribeRenderingInstanceSettings queries the configuration of a Cloud Application service instance.
//
// @param tmpReq - DescribeRenderingInstanceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingInstanceSettingsResponse
func (client *Client) DescribeRenderingInstanceSettingsWithOptions(tmpReq *DescribeRenderingInstanceSettingsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingInstanceSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRenderingInstanceSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AttributeNames) {
		request.AttributeNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttributeNames, dara.String("AttributeNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttributeNamesShrink) {
		query["AttributeNames"] = request.AttributeNamesShrink
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRenderingInstanceSettings"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRenderingInstanceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DescribeRenderingInstanceSettings queries the configuration of a Cloud Application service instance.
//
// @param request - DescribeRenderingInstanceSettingsRequest
//
// @return DescribeRenderingInstanceSettingsResponse
func (client *Client) DescribeRenderingInstanceSettings(request *DescribeRenderingInstanceSettingsRequest) (_result *DescribeRenderingInstanceSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRenderingInstanceSettingsResponse{}
	_body, _err := client.DescribeRenderingInstanceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve details about a rendering session, including the current session state, network access IP address and port, and the location of the cloud application service instance.
//
// @param request - DescribeRenderingSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingSessionResponse
func (client *Client) DescribeRenderingSessionWithOptions(request *DescribeRenderingSessionRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRenderingSession"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRenderingSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve details about a rendering session, including the current session state, network access IP address and port, and the location of the cloud application service instance.
//
// @param request - DescribeRenderingSessionRequest
//
// @return DescribeRenderingSessionResponse
func (client *Client) DescribeRenderingSession(request *DescribeRenderingSessionRequest) (_result *DescribeRenderingSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRenderingSessionResponse{}
	_body, _err := client.DescribeRenderingSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a stream.
//
// @param request - DescribeStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamResponse
func (client *Client) DescribeStreamWithOptions(request *DescribeStreamRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a stream.
//
// @param request - DescribeStreamRequest
//
// @return DescribeStreamResponse
func (client *Client) DescribeStream(request *DescribeStreamRequest) (_result *DescribeStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamResponse{}
	_body, _err := client.DescribeStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the URL of a stream.
//
// @param request - DescribeStreamURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamURLResponse
func (client *Client) DescribeStreamURLWithOptions(request *DescribeStreamURLRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamURLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		query["Auth"] = request.Auth
	}

	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OutProtocol) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Transcode) {
		query["Transcode"] = request.Transcode
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreamURL"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the URL of a stream.
//
// @param request - DescribeStreamURLRequest
//
// @return DescribeStreamURLResponse
func (client *Client) DescribeStreamURL(request *DescribeStreamURLRequest) (_result *DescribeStreamURLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamURLResponse{}
	_body, _err := client.DescribeStreamURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the stream VOD record list, such as historical stream list from NVR.
//
// @param request - DescribeStreamVodListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamVodListResponse
func (client *Client) DescribeStreamVodListWithOptions(request *DescribeStreamVodListRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamVodListResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
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
		Action:      dara.String("DescribeStreamVodList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamVodListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the stream VOD record list, such as historical stream list from NVR.
//
// @param request - DescribeStreamVodListRequest
//
// @return DescribeStreamVodListResponse
func (client *Client) DescribeStreamVodList(request *DescribeStreamVodListRequest) (_result *DescribeStreamVodListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamVodListResponse{}
	_body, _err := client.DescribeStreamVodListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists video streams. You can filter the results by stream ID, name, group ID, device ID, or other criteria.
//
// @param request - DescribeStreamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamsResponse
func (client *Client) DescribeStreamsWithOptions(request *DescribeStreamsRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreams"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists video streams. You can filter the results by stream ID, name, group ID, device ID, or other criteria.
//
// @param request - DescribeStreamsRequest
//
// @return DescribeStreamsResponse
func (client *Client) DescribeStreams(request *DescribeStreamsRequest) (_result *DescribeStreamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamsResponse{}
	_body, _err := client.DescribeStreamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query information about a template.
//
// @param request - DescribeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateResponse
func (client *Client) DescribeTemplateWithOptions(request *DescribeTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query information about a template.
//
// @param request - DescribeTemplateRequest
//
// @return DescribeTemplateResponse
func (client *Client) DescribeTemplate(request *DescribeTemplateRequest) (_result *DescribeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplateResponse{}
	_body, _err := client.DescribeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List templates.
//
// @param request - DescribeTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplates"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List templates.
//
// @param request - DescribeTemplatesRequest
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the URL information of a video-on-demand (VOD) stream.
//
// @param request - DescribeVodStreamURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodStreamURLResponse
func (client *Client) DescribeVodStreamURLWithOptions(request *DescribeVodStreamURLRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodStreamURLResponse, _err error) {
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

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVodStreamURL"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVodStreamURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the URL information of a video-on-demand (VOD) stream.
//
// @param request - DescribeVodStreamURLRequest
//
// @return DescribeVodStreamURLResponse
func (client *Client) DescribeVodStreamURL(request *DescribeVodStreamURLRequest) (_result *DescribeVodStreamURLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVodStreamURLResponse{}
	_body, _err := client.DescribeVodStreamURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve certificate details.
//
// @param request - DescribeVsCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsCertificateDetailResponse
func (client *Client) DescribeVsCertificateDetailWithOptions(request *DescribeVsCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsCertificateDetailResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsCertificateDetail"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve certificate details.
//
// @param request - DescribeVsCertificateDetailRequest
//
// @return DescribeVsCertificateDetailResponse
func (client *Client) DescribeVsCertificateDetail(request *DescribeVsCertificateDetailRequest) (_result *DescribeVsCertificateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsCertificateDetailResponse{}
	_body, _err := client.DescribeVsCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of domain name certificates.
//
// @param request - DescribeVsCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsCertificateListResponse
func (client *Client) DescribeVsCertificateListWithOptions(request *DescribeVsCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsCertificateListResponse, _err error) {
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
		Action:      dara.String("DescribeVsCertificateList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of domain name certificates.
//
// @param request - DescribeVsCertificateListRequest
//
// @return DescribeVsCertificateListResponse
func (client *Client) DescribeVsCertificateList(request *DescribeVsCertificateListRequest) (_result *DescribeVsCertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsCertificateListResponse{}
	_body, _err := client.DescribeVsCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves usage data for connected devices.
//
// Description:
//
// If you do not specify \\`StartTime\\` and \\`EndTime\\`, the API retrieves data from the last 24 hours by default. To query data for a specific time range, you must specify both \\`StartTime\\` and \\`EndTime\\`. The maximum time range for a single query is 31 days.
//
// - You can query multiple domain names at once. Separate the domain names with commas.
//
// - You can retrieve data from the last 90 days.
//
// - The time granularity is one hour.
//
// @param request - DescribeVsDevicesDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDevicesDataResponse
func (client *Client) DescribeVsDevicesDataWithOptions(request *DescribeVsDevicesDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDevicesDataResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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
		Action:      dara.String("DescribeVsDevicesData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDevicesDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves usage data for connected devices.
//
// Description:
//
// If you do not specify \\`StartTime\\` and \\`EndTime\\`, the API retrieves data from the last 24 hours by default. To query data for a specific time range, you must specify both \\`StartTime\\` and \\`EndTime\\`. The maximum time range for a single query is 31 days.
//
// - You can query multiple domain names at once. Separate the domain names with commas.
//
// - You can retrieve data from the last 90 days.
//
// - The time granularity is one hour.
//
// @param request - DescribeVsDevicesDataRequest
//
// @return DescribeVsDevicesDataResponse
func (client *Client) DescribeVsDevicesData(request *DescribeVsDevicesDataRequest) (_result *DescribeVsDevicesDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDevicesDataResponse{}
	_body, _err := client.DescribeVsDevicesDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries network bandwidth monitoring data for Domain Names.
//
// @param request - DescribeVsDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainBpsDataResponse
func (client *Client) DescribeVsDomainBpsDataWithOptions(request *DescribeVsDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainBpsDataResponse, _err error) {
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
		Action:      dara.String("DescribeVsDomainBpsData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network bandwidth monitoring data for Domain Names.
//
// @param request - DescribeVsDomainBpsDataRequest
//
// @return DescribeVsDomainBpsDataResponse
func (client *Client) DescribeVsDomainBpsData(request *DescribeVsDomainBpsDataRequest) (_result *DescribeVsDomainBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainBpsDataResponse{}
	_body, _err := client.DescribeVsDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the certificate information for a specified accelerated domain name.
//
// @param request - DescribeVsDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainCertificateInfoResponse
func (client *Client) DescribeVsDomainCertificateInfoWithOptions(request *DescribeVsDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainCertificateInfoResponse, _err error) {
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
		Action:      dara.String("DescribeVsDomainCertificateInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainCertificateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the certificate information for a specified accelerated domain name.
//
// @param request - DescribeVsDomainCertificateInfoRequest
//
// @return DescribeVsDomainCertificateInfoResponse
func (client *Client) DescribeVsDomainCertificateInfo(request *DescribeVsDomainCertificateInfoRequest) (_result *DescribeVsDomainCertificateInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainCertificateInfoResponse{}
	_body, _err := client.DescribeVsDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain name configurations. You can query the configurations of multiple features in a single request.
//
// @param request - DescribeVsDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainConfigsResponse
func (client *Client) DescribeVsDomainConfigsWithOptions(request *DescribeVsDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainConfigsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsDomainConfigs"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain name configurations. You can query the configurations of multiple features in a single request.
//
// @param request - DescribeVsDomainConfigsRequest
//
// @return DescribeVsDomainConfigsResponse
func (client *Client) DescribeVsDomainConfigs(request *DescribeVsDomainConfigsRequest) (_result *DescribeVsDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainConfigsResponse{}
	_body, _err := client.DescribeVsDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the basic configuration information for a specified Visual Edge Computing Service domain name.
//
// @param request - DescribeVsDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainDetailResponse
func (client *Client) DescribeVsDomainDetailWithOptions(request *DescribeVsDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainDetailResponse, _err error) {
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
		Action:      dara.String("DescribeVsDomainDetail"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the basic configuration information for a specified Visual Edge Computing Service domain name.
//
// @param request - DescribeVsDomainDetailRequest
//
// @return DescribeVsDomainDetailResponse
func (client *Client) DescribeVsDomainDetail(request *DescribeVsDomainDetailRequest) (_result *DescribeVsDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainDetailResponse{}
	_body, _err := client.DescribeVsDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the page view (PV) data for a domain name.
//
// @param request - DescribeVsDomainPvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainPvDataResponse
func (client *Client) DescribeVsDomainPvDataWithOptions(request *DescribeVsDomainPvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainPvDataResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsDomainPvData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainPvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the page view (PV) data for a domain name.
//
// @param request - DescribeVsDomainPvDataRequest
//
// @return DescribeVsDomainPvDataResponse
func (client *Client) DescribeVsDomainPvData(request *DescribeVsDomainPvDataRequest) (_result *DescribeVsDomainPvDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainPvDataResponse{}
	_body, _err := client.DescribeVsDomainPvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve page view (PV) and unique visitor (UV) data for a Visual Edge Computing Service domain.
//
// @param request - DescribeVsDomainPvUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainPvUvDataResponse
func (client *Client) DescribeVsDomainPvUvDataWithOptions(request *DescribeVsDomainPvUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainPvUvDataResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsDomainPvUvData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainPvUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve page view (PV) and unique visitor (UV) data for a Visual Edge Computing Service domain.
//
// @param request - DescribeVsDomainPvUvDataRequest
//
// @return DescribeVsDomainPvUvDataResponse
func (client *Client) DescribeVsDomainPvUvData(request *DescribeVsDomainPvUvDataRequest) (_result *DescribeVsDomainPvUvDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainPvUvDataResponse{}
	_body, _err := client.DescribeVsDomainPvUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve domain name record data.
//
// @param request - DescribeVsDomainRecordDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainRecordDataResponse
func (client *Client) DescribeVsDomainRecordDataWithOptions(request *DescribeVsDomainRecordDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainRecordDataResponse, _err error) {
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
		Action:      dara.String("DescribeVsDomainRecordData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainRecordDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve domain name record data.
//
// @param request - DescribeVsDomainRecordDataRequest
//
// @return DescribeVsDomainRecordDataResponse
func (client *Client) DescribeVsDomainRecordData(request *DescribeVsDomainRecordDataRequest) (_result *DescribeVsDomainRecordDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainRecordDataResponse{}
	_body, _err := client.DescribeVsDomainRecordDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves domain region data.
//
// @param request - DescribeVsDomainRegionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainRegionDataResponse
func (client *Client) DescribeVsDomainRegionDataWithOptions(request *DescribeVsDomainRegionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainRegionDataResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsDomainRegionData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainRegionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves domain region data.
//
// @param request - DescribeVsDomainRegionDataRequest
//
// @return DescribeVsDomainRegionDataResponse
func (client *Client) DescribeVsDomainRegionData(request *DescribeVsDomainRegionDataRequest) (_result *DescribeVsDomainRegionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainRegionDataResponse{}
	_body, _err := client.DescribeVsDomainRegionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query network request monitoring data for a domain name.
//
// @param request - DescribeVsDomainReqBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainReqBpsDataResponse
func (client *Client) DescribeVsDomainReqBpsDataWithOptions(request *DescribeVsDomainReqBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainReqBpsDataResponse, _err error) {
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
		Action:      dara.String("DescribeVsDomainReqBpsData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainReqBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query network request monitoring data for a domain name.
//
// @param request - DescribeVsDomainReqBpsDataRequest
//
// @return DescribeVsDomainReqBpsDataResponse
func (client *Client) DescribeVsDomainReqBpsData(request *DescribeVsDomainReqBpsDataRequest) (_result *DescribeVsDomainReqBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainReqBpsDataResponse{}
	_body, _err := client.DescribeVsDomainReqBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain traffic data for domain name requests.
//
// @param request - DescribeVsDomainReqTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainReqTrafficDataResponse
func (client *Client) DescribeVsDomainReqTrafficDataWithOptions(request *DescribeVsDomainReqTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainReqTrafficDataResponse, _err error) {
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
		Action:      dara.String("DescribeVsDomainReqTrafficData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainReqTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain traffic data for domain name requests.
//
// @param request - DescribeVsDomainReqTrafficDataRequest
//
// @return DescribeVsDomainReqTrafficDataResponse
func (client *Client) DescribeVsDomainReqTrafficData(request *DescribeVsDomainReqTrafficDataRequest) (_result *DescribeVsDomainReqTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainReqTrafficDataResponse{}
	_body, _err := client.DescribeVsDomainReqTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves snapshot data for a domain name.
//
// @param request - DescribeVsDomainSnapshotDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainSnapshotDataResponse
func (client *Client) DescribeVsDomainSnapshotDataWithOptions(request *DescribeVsDomainSnapshotDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainSnapshotDataResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsDomainSnapshotData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainSnapshotDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves snapshot data for a domain name.
//
// @param request - DescribeVsDomainSnapshotDataRequest
//
// @return DescribeVsDomainSnapshotDataResponse
func (client *Client) DescribeVsDomainSnapshotData(request *DescribeVsDomainSnapshotDataRequest) (_result *DescribeVsDomainSnapshotDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainSnapshotDataResponse{}
	_body, _err := client.DescribeVsDomainSnapshotDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve traffic data for a domain name.
//
// @param request - DescribeVsDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainTrafficDataResponse
func (client *Client) DescribeVsDomainTrafficDataWithOptions(request *DescribeVsDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainTrafficDataResponse, _err error) {
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
		Action:      dara.String("DescribeVsDomainTrafficData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve traffic data for a domain name.
//
// @param request - DescribeVsDomainTrafficDataRequest
//
// @return DescribeVsDomainTrafficDataResponse
func (client *Client) DescribeVsDomainTrafficData(request *DescribeVsDomainTrafficDataRequest) (_result *DescribeVsDomainTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainTrafficDataResponse{}
	_body, _err := client.DescribeVsDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve UV data by domain name.
//
// @param request - DescribeVsDomainUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainUvDataResponse
func (client *Client) DescribeVsDomainUvDataWithOptions(request *DescribeVsDomainUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainUvDataResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsDomainUvData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsDomainUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve UV data by domain name.
//
// @param request - DescribeVsDomainUvDataRequest
//
// @return DescribeVsDomainUvDataResponse
func (client *Client) DescribeVsDomainUvData(request *DescribeVsDomainUvDataRequest) (_result *DescribeVsDomainUvDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsDomainUvDataResponse{}
	_body, _err := client.DescribeVsDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the pull stream configurations for a domain name.
//
// @param request - DescribeVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsPullStreamInfoConfigResponse
func (client *Client) DescribeVsPullStreamInfoConfigWithOptions(request *DescribeVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsPullStreamInfoConfigResponse, _err error) {
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
		Action:      dara.String("DescribeVsPullStreamInfoConfig"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pull stream configurations for a domain name.
//
// @param request - DescribeVsPullStreamInfoConfigRequest
//
// @return DescribeVsPullStreamInfoConfigResponse
func (client *Client) DescribeVsPullStreamInfoConfig(request *DescribeVsPullStreamInfoConfigRequest) (_result *DescribeVsPullStreamInfoConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsPullStreamInfoConfigResponse{}
	_body, _err := client.DescribeVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the stream ingest callback configuration.
//
// @param request - DescribeVsStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsStreamsNotifyUrlConfigResponse
func (client *Client) DescribeVsStreamsNotifyUrlConfigWithOptions(request *DescribeVsStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsStreamsNotifyUrlConfigResponse, _err error) {
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
		Action:      dara.String("DescribeVsStreamsNotifyUrlConfig"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the stream ingest callback configuration.
//
// @param request - DescribeVsStreamsNotifyUrlConfigRequest
//
// @return DescribeVsStreamsNotifyUrlConfigResponse
func (client *Client) DescribeVsStreamsNotifyUrlConfig(request *DescribeVsStreamsNotifyUrlConfigRequest) (_result *DescribeVsStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DescribeVsStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about all active streams for a specified domain name or application.
//
// @param request - DescribeVsStreamsOnlineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsStreamsOnlineListResponse
func (client *Client) DescribeVsStreamsOnlineListWithOptions(request *DescribeVsStreamsOnlineListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsStreamsOnlineListResponse, _err error) {
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

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsStreamsOnlineList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsStreamsOnlineListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about all active streams for a specified domain name or application.
//
// @param request - DescribeVsStreamsOnlineListRequest
//
// @return DescribeVsStreamsOnlineListResponse
func (client *Client) DescribeVsStreamsOnlineList(request *DescribeVsStreamsOnlineListRequest) (_result *DescribeVsStreamsOnlineListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsStreamsOnlineListResponse{}
	_body, _err := client.DescribeVsStreamsOnlineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve stream ingest records for a domain, an application under that domain, or a specific stream within a specified time range.
//
// @param request - DescribeVsStreamsPublishListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsStreamsPublishListResponse
func (client *Client) DescribeVsStreamsPublishListWithOptions(request *DescribeVsStreamsPublishListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsStreamsPublishListResponse, _err error) {
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

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
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

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsStreamsPublishList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsStreamsPublishListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve stream ingest records for a domain, an application under that domain, or a specific stream within a specified time range.
//
// @param request - DescribeVsStreamsPublishListRequest
//
// @return DescribeVsStreamsPublishListResponse
func (client *Client) DescribeVsStreamsPublishList(request *DescribeVsStreamsPublishListRequest) (_result *DescribeVsStreamsPublishListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsStreamsPublishListResponse{}
	_body, _err := client.DescribeVsStreamsPublishListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a user\\"s domain names ranked by traffic.
//
// If you do not specify StartTime and EndTime, data for the current month is retrieved by default. To query data over a specific time range, you must specify both StartTime and EndTime.
//
// \\	- You can retrieve data for a maximum of 90 days.
//
// @param request - DescribeVsTopDomainsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsTopDomainsByFlowResponse
func (client *Client) DescribeVsTopDomainsByFlowWithOptions(request *DescribeVsTopDomainsByFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsTopDomainsByFlowResponse, _err error) {
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

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
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
		Action:      dara.String("DescribeVsTopDomainsByFlow"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsTopDomainsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a user\\"s domain names ranked by traffic.
//
// If you do not specify StartTime and EndTime, data for the current month is retrieved by default. To query data over a specific time range, you must specify both StartTime and EndTime.
//
// \\	- You can retrieve data for a maximum of 90 days.
//
// @param request - DescribeVsTopDomainsByFlowRequest
//
// @return DescribeVsTopDomainsByFlowResponse
func (client *Client) DescribeVsTopDomainsByFlow(request *DescribeVsTopDomainsByFlowRequest) (_result *DescribeVsTopDomainsByFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsTopDomainsByFlowResponse{}
	_body, _err := client.DescribeVsTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the daily peak number of concurrent stream ingest operations.
//
// @param request - DescribeVsUpPeakPublishStreamDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsUpPeakPublishStreamDataResponse
func (client *Client) DescribeVsUpPeakPublishStreamDataWithOptions(request *DescribeVsUpPeakPublishStreamDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsUpPeakPublishStreamDataResponse, _err error) {
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

	if !dara.IsNil(request.DomainSwitch) {
		query["DomainSwitch"] = request.DomainSwitch
	}

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
		Action:      dara.String("DescribeVsUpPeakPublishStreamData"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsUpPeakPublishStreamDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the daily peak number of concurrent stream ingest operations.
//
// @param request - DescribeVsUpPeakPublishStreamDataRequest
//
// @return DescribeVsUpPeakPublishStreamDataResponse
func (client *Client) DescribeVsUpPeakPublishStreamData(request *DescribeVsUpPeakPublishStreamDataRequest) (_result *DescribeVsUpPeakPublishStreamDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsUpPeakPublishStreamDataResponse{}
	_body, _err := client.DescribeVsUpPeakPublishStreamDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeVsUserResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsUserResourcePackageResponse
func (client *Client) DescribeVsUserResourcePackageWithOptions(request *DescribeVsUserResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsUserResourcePackageResponse, _err error) {
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
		Action:      dara.String("DescribeVsUserResourcePackage"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsUserResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeVsUserResourcePackageRequest
//
// @return DescribeVsUserResourcePackageResponse
func (client *Client) DescribeVsUserResourcePackage(request *DescribeVsUserResourcePackageRequest) (_result *DescribeVsUserResourcePackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsUserResourcePackageResponse{}
	_body, _err := client.DescribeVsUserResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeVsVerifyContent
//
// @param request - DescribeVsVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsVerifyContentResponse
func (client *Client) DescribeVsVerifyContentWithOptions(request *DescribeVsVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsVerifyContentResponse, _err error) {
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
		Action:      dara.String("DescribeVsVerifyContent"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVsVerifyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeVsVerifyContent
//
// @param request - DescribeVsVerifyContentRequest
//
// @return DescribeVsVerifyContentResponse
func (client *Client) DescribeVsVerifyContent(request *DescribeVsVerifyContentRequest) (_result *DescribeVsVerifyContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVsVerifyContentResponse{}
	_body, _err := client.DescribeVsVerifyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociate cloud application service instances from a project.
//
// @param tmpReq - DisassociateRenderingProjectInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateRenderingProjectInstancesResponse
func (client *Client) DisassociateRenderingProjectInstancesWithOptions(tmpReq *DisassociateRenderingProjectInstancesRequest, runtime *dara.RuntimeOptions) (_result *DisassociateRenderingProjectInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DisassociateRenderingProjectInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RenderingInstanceIds) {
		request.RenderingInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RenderingInstanceIds, dara.String("RenderingInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RenderingInstanceIdsShrink) {
		query["RenderingInstanceIds"] = request.RenderingInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateRenderingProjectInstances"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateRenderingProjectInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociate cloud application service instances from a project.
//
// @param request - DisassociateRenderingProjectInstancesRequest
//
// @return DisassociateRenderingProjectInstancesResponse
func (client *Client) DisassociateRenderingProjectInstances(request *DisassociateRenderingProjectInstancesRequest) (_result *DisassociateRenderingProjectInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisassociateRenderingProjectInstancesResponse{}
	_body, _err := client.DisassociateRenderingProjectInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Forbids pushing a specific stream. You can schedule a time to resume the stream.
//
// @param request - ForbidVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForbidVsStreamResponse
func (client *Client) ForbidVsStreamWithOptions(request *ForbidVsStreamRequest, runtime *dara.RuntimeOptions) (_result *ForbidVsStreamResponse, _err error) {
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

	if !dara.IsNil(request.ControlStreamAction) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LiveStreamType) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !dara.IsNil(request.Oneshot) {
		query["Oneshot"] = request.Oneshot
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResumeTime) {
		query["ResumeTime"] = request.ResumeTime
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForbidVsStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForbidVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Forbids pushing a specific stream. You can schedule a time to resume the stream.
//
// @param request - ForbidVsStreamRequest
//
// @return ForbidVsStreamResponse
func (client *Client) ForbidVsStream(request *ForbidVsStreamRequest) (_result *ForbidVsStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ForbidVsStreamResponse{}
	_body, _err := client.ForbidVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution status of a control command to determine whether the command was successful and to retrieve the result string.
//
// @param request - GetRenderingInstanceCommandsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenderingInstanceCommandsStatusResponse
func (client *Client) GetRenderingInstanceCommandsStatusWithOptions(request *GetRenderingInstanceCommandsStatusRequest, runtime *dara.RuntimeOptions) (_result *GetRenderingInstanceCommandsStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CmdId) {
		query["CmdId"] = request.CmdId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRenderingInstanceCommandsStatus"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRenderingInstanceCommandsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of a control command to determine whether the command was successful and to retrieve the result string.
//
// @param request - GetRenderingInstanceCommandsStatusRequest
//
// @return GetRenderingInstanceCommandsStatusResponse
func (client *Client) GetRenderingInstanceCommandsStatus(request *GetRenderingInstanceCommandsStatusRequest) (_result *GetRenderingInstanceCommandsStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRenderingInstanceCommandsStatusResponse{}
	_body, _err := client.GetRenderingInstanceCommandsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the streaming connection information for a cloud application service instance. Call this operation before establishing each streaming connection to obtain the latest connection details.
//
// @param request - GetRenderingInstanceStreamingInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenderingInstanceStreamingInfoResponse
func (client *Client) GetRenderingInstanceStreamingInfoWithOptions(request *GetRenderingInstanceStreamingInfoRequest, runtime *dara.RuntimeOptions) (_result *GetRenderingInstanceStreamingInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRenderingInstanceStreamingInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRenderingInstanceStreamingInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the streaming connection information for a cloud application service instance. Call this operation before establishing each streaming connection to obtain the latest connection details.
//
// @param request - GetRenderingInstanceStreamingInfoRequest
//
// @return GetRenderingInstanceStreamingInfoResponse
func (client *Client) GetRenderingInstanceStreamingInfo(request *GetRenderingInstanceStreamingInfoRequest) (_result *GetRenderingInstanceStreamingInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRenderingInstanceStreamingInfoResponse{}
	_body, _err := client.GetRenderingInstanceStreamingInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data volume statistics for the states of project instances that meet specified conditions.
//
// @param request - GetRenderingProjectInstanceStateMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenderingProjectInstanceStateMetricsResponse
func (client *Client) GetRenderingProjectInstanceStateMetricsWithOptions(request *GetRenderingProjectInstanceStateMetricsRequest, runtime *dara.RuntimeOptions) (_result *GetRenderingProjectInstanceStateMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRenderingProjectInstanceStateMetrics"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRenderingProjectInstanceStateMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data volume statistics for the states of project instances that meet specified conditions.
//
// @param request - GetRenderingProjectInstanceStateMetricsRequest
//
// @return GetRenderingProjectInstanceStateMetricsResponse
func (client *Client) GetRenderingProjectInstanceStateMetrics(request *GetRenderingProjectInstanceStateMetricsRequest) (_result *GetRenderingProjectInstanceStateMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRenderingProjectInstanceStateMetricsResponse{}
	_body, _err := client.GetRenderingProjectInstanceStateMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves to a specified preset.
//
// @param request - GotoPresetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GotoPresetResponse
func (client *Client) GotoPresetWithOptions(request *GotoPresetRequest, runtime *dara.RuntimeOptions) (_result *GotoPresetResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PresetId) {
		query["PresetId"] = request.PresetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GotoPreset"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GotoPresetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves to a specified preset.
//
// @param request - GotoPresetRequest
//
// @return GotoPresetResponse
func (client *Client) GotoPreset(request *GotoPresetRequest) (_result *GotoPresetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GotoPresetResponse{}
	_body, _err := client.GotoPresetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs a cloud application to a specified cloud application instance. This is an asynchronous interface. To monitor the installation progress, use the ListCloudAppInstallations interface.
//
// @param tmpReq - InstallCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCloudAppResponse
func (client *Client) InstallCloudAppWithOptions(tmpReq *InstallCloudAppRequest, runtime *dara.RuntimeOptions) (_result *InstallCloudAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InstallCloudAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RenderingInstanceIds) {
		request.RenderingInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RenderingInstanceIds, dara.String("RenderingInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PatchId) {
		query["PatchId"] = request.PatchId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.RenderingInstanceIdsShrink) {
		query["RenderingInstanceIds"] = request.RenderingInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallCloudApp"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallCloudAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs a cloud application to a specified cloud application instance. This is an asynchronous interface. To monitor the installation progress, use the ListCloudAppInstallations interface.
//
// @param request - InstallCloudAppRequest
//
// @return InstallCloudAppResponse
func (client *Client) InstallCloudApp(request *InstallCloudAppRequest) (_result *InstallCloudAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallCloudAppResponse{}
	_body, _err := client.InstallCloudAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists cloud application installations. The response includes the installation status of cloud application service instances and supports paged queries.
//
// @param request - ListCloudAppInstallationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAppInstallationsResponse
func (client *Client) ListCloudAppInstallationsWithOptions(request *ListCloudAppInstallationsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAppInstallationsResponse, _err error) {
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
		Action:      dara.String("ListCloudAppInstallations"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudAppInstallationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists cloud application installations. The response includes the installation status of cloud application service instances and supports paged queries.
//
// @param request - ListCloudAppInstallationsRequest
//
// @return ListCloudAppInstallationsResponse
func (client *Client) ListCloudAppInstallations(request *ListCloudAppInstallationsRequest) (_result *ListCloudAppInstallationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudAppInstallationsResponse{}
	_body, _err := client.ListCloudAppInstallationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of patches for a cloud application.
//
// Description:
//
// > Specify at least one of the template ID or the template type.
//
// @param request - ListCloudAppPatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAppPatchesResponse
func (client *Client) ListCloudAppPatchesWithOptions(request *ListCloudAppPatchesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAppPatchesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PatchId) {
		query["PatchId"] = request.PatchId
	}

	if !dara.IsNil(request.PatchName) {
		query["PatchName"] = request.PatchName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudAppPatches"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudAppPatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of patches for a cloud application.
//
// Description:
//
// > Specify at least one of the template ID or the template type.
//
// @param request - ListCloudAppPatchesRequest
//
// @return ListCloudAppPatchesResponse
func (client *Client) ListCloudAppPatches(request *ListCloudAppPatchesRequest) (_result *ListCloudAppPatchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudAppPatchesResponse{}
	_body, _err := client.ListCloudAppPatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of cloud applications. This operation supports paged queries.
//
// @param request - ListCloudAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAppsResponse
func (client *Client) ListCloudAppsWithOptions(request *ListCloudAppsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAppsResponse, _err error) {
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
		Action:      dara.String("ListCloudApps"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of cloud applications. This operation supports paged queries.
//
// @param request - ListCloudAppsRequest
//
// @return ListCloudAppsResponse
func (client *Client) ListCloudApps(request *ListCloudAppsRequest) (_result *ListCloudAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudAppsResponse{}
	_body, _err := client.ListCloudAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries payload information for cloud application services. This operation supports paged queries.
//
// Description:
//
// ## Request description
//
// - This API queries payload information for cloud application services and supports filtering and paged queries using various parameters.
//
// - Optional parameters include `Spec`, `Statuses`, `InstanceIds`, `PlanIds`, and `HiveIds`.
//
// - For paged queries, you can use the `PageNumber` and `PageSize` parameters to control the amount of data returned. The default page size is 10 records, and the maximum is 100 records.
//
// - You can specify a time range for the query using the `StartTime` and `EndTime` parameters.
//
// @param tmpReq - ListEdgeWorkersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeWorkersResponse
func (client *Client) ListEdgeWorkersWithOptions(tmpReq *ListEdgeWorkersRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeWorkersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListEdgeWorkersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HiveIds) {
		request.HiveIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HiveIds, dara.String("HiveIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PlanIds) {
		request.PlanIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlanIds, dara.String("PlanIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HiveIdsShrink) {
		query["HiveIds"] = request.HiveIdsShrink
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlanIdsShrink) {
		query["PlanIds"] = request.PlanIdsShrink
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatusesShrink) {
		query["Statuses"] = request.StatusesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEdgeWorkers"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeWorkersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries payload information for cloud application services. This operation supports paged queries.
//
// Description:
//
// ## Request description
//
// - This API queries payload information for cloud application services and supports filtering and paged queries using various parameters.
//
// - Optional parameters include `Spec`, `Statuses`, `InstanceIds`, `PlanIds`, and `HiveIds`.
//
// - For paged queries, you can use the `PageNumber` and `PageSize` parameters to control the amount of data returned. The default page size is 10 records, and the maximum is 100 records.
//
// - You can specify a time range for the query using the `StartTime` and `EndTime` parameters.
//
// @param request - ListEdgeWorkersRequest
//
// @return ListEdgeWorkersResponse
func (client *Client) ListEdgeWorkers(request *ListEdgeWorkersRequest) (_result *ListEdgeWorkersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeWorkersResponse{}
	_body, _err := client.ListEdgeWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the push status records for a file pushed to cloud application service instances. It supports paged query.
//
// @param request - ListFilePushStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilePushStatusesResponse
func (client *Client) ListFilePushStatusesWithOptions(request *ListFilePushStatusesRequest, runtime *dara.RuntimeOptions) (_result *ListFilePushStatusesResponse, _err error) {
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
		Action:      dara.String("ListFilePushStatuses"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilePushStatusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the push status records for a file pushed to cloud application service instances. It supports paged query.
//
// @param request - ListFilePushStatusesRequest
//
// @return ListFilePushStatusesResponse
func (client *Client) ListFilePushStatuses(request *ListFilePushStatusesRequest) (_result *ListFilePushStatusesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFilePushStatusesResponse{}
	_body, _err := client.ListFilePushStatusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists uploaded files. The response includes the upload status for each file and supports paged queries.
//
// @param request - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithOptions(request *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
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
		Action:      dara.String("ListFiles"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists uploaded files. The response includes the upload status for each file and supports paged queries.
//
// @param request - ListFilesRequest
//
// @return ListFilesResponse
func (client *Client) ListFiles(request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询所有集群信息，支持分页查询。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于查询用户创建的所有集群信息。
//
// - 支持通过 `HiveId` 和 `Name` 参数进行过滤查询。
//
// - 分页参数 `PageNumber` 和 `PageSize` 可以控制返回结果的数量和页码，默认每页显示10条记录，最大支持100条。
//
// - `StartTime` 和 `EndTime` 参数可用于指定时间范围内的集群信息查询，但非必填项。
//
// @param request - ListHivesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHivesResponse
func (client *Client) ListHivesWithOptions(request *ListHivesRequest, runtime *dara.RuntimeOptions) (_result *ListHivesResponse, _err error) {
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

	if !dara.IsNil(request.HiveId) {
		query["HiveId"] = request.HiveId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
		Action:      dara.String("ListHives"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHivesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有集群信息，支持分页查询。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于查询用户创建的所有集群信息。
//
// - 支持通过 `HiveId` 和 `Name` 参数进行过滤查询。
//
// - 分页参数 `PageNumber` 和 `PageSize` 可以控制返回结果的数量和页码，默认每页显示10条记录，最大支持100条。
//
// - `StartTime` 和 `EndTime` 参数可用于指定时间范围内的集群信息查询，但非必填项。
//
// @param request - ListHivesRequest
//
// @return ListHivesResponse
func (client *Client) ListHives(request *ListHivesRequest) (_result *ListHivesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHivesResponse{}
	_body, _err := client.ListHivesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of public keys that match the specified criteria. This operation supports pagination.
//
// @param request - ListPublicKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicKeysResponse
func (client *Client) ListPublicKeysWithOptions(request *ListPublicKeysRequest, runtime *dara.RuntimeOptions) (_result *ListPublicKeysResponse, _err error) {
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
		Action:      dara.String("ListPublicKeys"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublicKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of public keys that match the specified criteria. This operation supports pagination.
//
// @param request - ListPublicKeysRequest
//
// @return ListPublicKeysResponse
func (client *Client) ListPublicKeys(request *ListPublicKeysRequest) (_result *ListPublicKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPublicKeysResponse{}
	_body, _err := client.ListPublicKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the data packets of cloud applications. Paged queries are supported.
//
// @param request - ListRenderingDataPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingDataPackagesResponse
func (client *Client) ListRenderingDataPackagesWithOptions(request *ListRenderingDataPackagesRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingDataPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.DataPackageId) {
		query["DataPackageId"] = request.DataPackageId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRenderingDataPackages"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRenderingDataPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the data packets of cloud applications. Paged queries are supported.
//
// @param request - ListRenderingDataPackagesRequest
//
// @return ListRenderingDataPackagesResponse
func (client *Client) ListRenderingDataPackages(request *ListRenderingDataPackagesRequest) (_result *ListRenderingDataPackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRenderingDataPackagesResponse{}
	_body, _err := client.ListRenderingDataPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom gateways.
//
// Description:
//
// > Specify at least the template ID or the template type.
//
// @param request - ListRenderingInstanceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingInstanceGatewayResponse
func (client *Client) ListRenderingInstanceGatewayWithOptions(request *ListRenderingInstanceGatewayRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingInstanceGatewayResponse, _err error) {
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

	if !dara.IsNil(request.GatewayInstanceId) {
		query["GatewayInstanceId"] = request.GatewayInstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRenderingInstanceGateway"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRenderingInstanceGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom gateways.
//
// Description:
//
// > Specify at least the template ID or the template type.
//
// @param request - ListRenderingInstanceGatewayRequest
//
// @return ListRenderingInstanceGatewayResponse
func (client *Client) ListRenderingInstanceGateway(request *ListRenderingInstanceGatewayRequest) (_result *ListRenderingInstanceGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRenderingInstanceGatewayResponse{}
	_body, _err := client.ListRenderingInstanceGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists basic information about cloud application service instances and supports paged queries.
//
// @param request - ListRenderingInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingInstancesResponse
func (client *Client) ListRenderingInstancesWithOptions(request *ListRenderingInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingInstancesResponse, _err error) {
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
		Action:      dara.String("ListRenderingInstances"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRenderingInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists basic information about cloud application service instances and supports paged queries.
//
// @param request - ListRenderingInstancesRequest
//
// @return ListRenderingInstancesResponse
func (client *Client) ListRenderingInstances(request *ListRenderingInstancesRequest) (_result *ListRenderingInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRenderingInstancesResponse{}
	_body, _err := client.ListRenderingInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve a paginated list of cloud application service instances associated with a project.
//
// Description:
//
// ## Request description
//
// - This operation enables you to query cloud application service instances in a project using multiple filter conditions, such as status and instance ID.
//
// @param request - ListRenderingProjectInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingProjectInstancesResponse
func (client *Client) ListRenderingProjectInstancesWithOptions(request *ListRenderingProjectInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingProjectInstancesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRenderingProjectInstances"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRenderingProjectInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve a paginated list of cloud application service instances associated with a project.
//
// Description:
//
// ## Request description
//
// - This operation enables you to query cloud application service instances in a project using multiple filter conditions, such as status and instance ID.
//
// @param request - ListRenderingProjectInstancesRequest
//
// @return ListRenderingProjectInstancesResponse
func (client *Client) ListRenderingProjectInstances(request *ListRenderingProjectInstancesRequest) (_result *ListRenderingProjectInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRenderingProjectInstancesResponse{}
	_body, _err := client.ListRenderingProjectInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain a paged list of basic information about cloud application projects for the current user.
//
// Description:
//
// ## Request details
//
// - This operation returns a paged list of basic information about rendering projects for a specified user.
//
// - Filter results by `ProjectId` or `ProjectName`.
//
// @param request - ListRenderingProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingProjectsResponse
func (client *Client) ListRenderingProjectsWithOptions(request *ListRenderingProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingProjectsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRenderingProjects"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRenderingProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain a paged list of basic information about cloud application projects for the current user.
//
// Description:
//
// ## Request details
//
// - This operation returns a paged list of basic information about rendering projects for a specified user.
//
// - Filter results by `ProjectId` or `ProjectName`.
//
// @param request - ListRenderingProjectsRequest
//
// @return ListRenderingProjectsResponse
func (client *Client) ListRenderingProjects(request *ListRenderingProjectsRequest) (_result *ListRenderingProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRenderingProjectsResponse{}
	_body, _err := client.ListRenderingProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query for the list of cloud application service sessions based on specified conditions.
//
// Description:
//
// ## Request Description
//
// - This API supports filtering and paged query of user rendering session lists with various parameter combinations.
//
// - You must provide at least one of the `SessionId` or `ClientId` parameters. Neither parameter is mandatory independently. If both parameters are provided, a more precise match is performed based on these two parameters.
//
// @param request - ListRenderingSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingSessionsResponse
func (client *Client) ListRenderingSessionsWithOptions(request *ListRenderingSessionsRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingSessionsResponse, _err error) {
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

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PatchId) {
		query["PatchId"] = request.PatchId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRenderingSessions"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRenderingSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query for the list of cloud application service sessions based on specified conditions.
//
// Description:
//
// ## Request Description
//
// - This API supports filtering and paged query of user rendering session lists with various parameter combinations.
//
// - You must provide at least one of the `SessionId` or `ClientId` parameters. Neither parameter is mandatory independently. If both parameters are provided, a more precise match is performed based on these two parameters.
//
// @param request - ListRenderingSessionsRequest
//
// @return ListRenderingSessionsResponse
func (client *Client) ListRenderingSessions(request *ListRenderingSessionsRequest) (_result *ListRenderingSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRenderingSessionsResponse{}
	_body, _err := client.ListRenderingSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询规格信息，支持分页查询。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于查询所有可用的云应用服务规格信息。
//
// - 支持通过 `Specification` 参数过滤特定规格。
//
// - 分页查询时，可以通过 `PageNumber` 和 `PageSize` 参数控制返回的数据量。
//
// @param request - ListSpecificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSpecificationsResponse
func (client *Client) ListSpecificationsWithOptions(request *ListSpecificationsRequest, runtime *dara.RuntimeOptions) (_result *ListSpecificationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSpecifications"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSpecificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询规格信息，支持分页查询。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于查询所有可用的云应用服务规格信息。
//
// - 支持通过 `Specification` 参数过滤特定规格。
//
// - 分页查询时，可以通过 `PageNumber` 和 `PageSize` 参数控制返回的数据量。
//
// @param request - ListSpecificationsRequest
//
// @return ListSpecificationsResponse
func (client *Client) ListSpecifications(request *ListSpecificationsRequest) (_result *ListSpecificationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSpecificationsResponse{}
	_body, _err := client.ListSpecificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manages secure logons.
//
// @param request - ManageLoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageLoginResponse
func (client *Client) ManageLoginWithOptions(request *ManageLoginRequest, runtime *dara.RuntimeOptions) (_result *ManageLoginResponse, _err error) {
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

	if !dara.IsNil(request.KeyGroup) {
		query["KeyGroup"] = request.KeyGroup
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManageLogin"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManageLoginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages secure logons.
//
// @param request - ManageLoginRequest
//
// @return ManageLoginResponse
func (client *Client) ManageLogin(request *ManageLoginRequest) (_result *ManageLoginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManageLoginResponse{}
	_body, _err := client.ManageLoginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the metadata of a specified workflow.
//
// Description:
//
// \\> 截图查询目前不支持分页，仅支持按迭代方式。使用返回结果里的extStartTime参数值，作为新请求的StartTime可请求下一页。
//
// @param request - ModifyComfyWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyComfyWorkflowResponse
func (client *Client) ModifyComfyWorkflowWithOptions(request *ModifyComfyWorkflowRequest, runtime *dara.RuntimeOptions) (_result *ModifyComfyWorkflowResponse, _err error) {
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

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowName) {
		query["WorkflowName"] = request.WorkflowName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyComfyWorkflow"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyComfyWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the metadata of a specified workflow.
//
// Description:
//
// \\> 截图查询目前不支持分页，仅支持按迭代方式。使用返回结果里的extStartTime参数值，作为新请求的StartTime可请求下一页。
//
// @param request - ModifyComfyWorkflowRequest
//
// @return ModifyComfyWorkflowResponse
func (client *Client) ModifyComfyWorkflow(request *ModifyComfyWorkflowRequest) (_result *ModifyComfyWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyComfyWorkflowResponse{}
	_body, _err := client.ModifyComfyWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update device information.
//
// @param request - ModifyDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceResponse
func (client *Client) ModifyDeviceWithOptions(request *ModifyDeviceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmMethod) {
		query["AlarmMethod"] = request.AlarmMethod
	}

	if !dara.IsNil(request.AutoDirectory) {
		query["AutoDirectory"] = request.AutoDirectory
	}

	if !dara.IsNil(request.AutoPos) {
		query["AutoPos"] = request.AutoPos
	}

	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Latitude) {
		query["Latitude"] = request.Latitude
	}

	if !dara.IsNil(request.Longitude) {
		query["Longitude"] = request.Longitude
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PosInterval) {
		query["PosInterval"] = request.PosInterval
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.Vendor) {
		query["Vendor"] = request.Vendor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update device information.
//
// @param request - ModifyDeviceRequest
//
// @return ModifyDeviceResponse
func (client *Client) ModifyDevice(request *ModifyDeviceRequest) (_result *ModifyDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDeviceResponse{}
	_body, _err := client.ModifyDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the alarm status of a device.
//
// @param request - ModifyDeviceAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceAlarmResponse
func (client *Client) ModifyDeviceAlarmWithOptions(request *ModifyDeviceAlarmRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		query["AlarmId"] = request.AlarmId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDeviceAlarm"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeviceAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the alarm status of a device.
//
// @param request - ModifyDeviceAlarmRequest
//
// @return ModifyDeviceAlarmResponse
func (client *Client) ModifyDeviceAlarm(request *ModifyDeviceAlarmRequest) (_result *ModifyDeviceAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDeviceAlarmResponse{}
	_body, _err := client.ModifyDeviceAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the device image capture configuration.
//
// @param request - ModifyDeviceCaptureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceCaptureResponse
func (client *Client) ModifyDeviceCaptureWithOptions(request *ModifyDeviceCaptureRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceCaptureResponse, _err error) {
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

	if !dara.IsNil(request.Image) {
		query["Image"] = request.Image
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Video) {
		query["Video"] = request.Video
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDeviceCapture"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeviceCaptureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the device image capture configuration.
//
// @param request - ModifyDeviceCaptureRequest
//
// @return ModifyDeviceCaptureResponse
func (client *Client) ModifyDeviceCapture(request *ModifyDeviceCaptureRequest) (_result *ModifyDeviceCaptureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDeviceCaptureResponse{}
	_body, _err := client.ModifyDeviceCaptureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the list of channels for a device.
//
// @param request - ModifyDeviceChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceChannelsResponse
func (client *Client) ModifyDeviceChannelsWithOptions(request *ModifyDeviceChannelsRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceChannelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Channels) {
		query["Channels"] = request.Channels
	}

	if !dara.IsNil(request.DeviceStatus) {
		query["DeviceStatus"] = request.DeviceStatus
	}

	if !dara.IsNil(request.Dsn) {
		query["Dsn"] = request.Dsn
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDeviceChannels"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeviceChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the list of channels for a device.
//
// @param request - ModifyDeviceChannelsRequest
//
// @return ModifyDeviceChannelsResponse
func (client *Client) ModifyDeviceChannels(request *ModifyDeviceChannelsRequest) (_result *ModifyDeviceChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDeviceChannelsResponse{}
	_body, _err := client.ModifyDeviceChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information of a directory.
//
// @param request - ModifyDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDirectoryResponse
func (client *Client) ModifyDirectoryWithOptions(request *ModifyDirectoryRequest, runtime *dara.RuntimeOptions) (_result *ModifyDirectoryResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDirectory"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a directory.
//
// @param request - ModifyDirectoryRequest
//
// @return ModifyDirectoryResponse
func (client *Client) ModifyDirectory(request *ModifyDirectoryRequest) (_result *ModifyDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDirectoryResponse{}
	_body, _err := client.ModifyDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can modify the details of a space.
//
// @param request - ModifyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGroupResponse
func (client *Client) ModifyGroupWithOptions(request *ModifyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InProtocol) {
		query["InProtocol"] = request.InProtocol
	}

	if !dara.IsNil(request.LazyPull) {
		query["LazyPull"] = request.LazyPull
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutProtocol) {
		query["OutProtocol"] = request.OutProtocol
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.PushDomain) {
		query["PushDomain"] = request.PushDomain
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGroup"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can modify the details of a space.
//
// @param request - ModifyGroupRequest
//
// @return ModifyGroupResponse
func (client *Client) ModifyGroup(request *ModifyGroupRequest) (_result *ModifyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGroupResponse{}
	_body, _err := client.ModifyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name or description of a specified cluster.
//
// Description:
//
// ## Request
//
// - This API modifies the name and/or description of an existing cluster.
//
// - `HiveId` is a required parameter that identifies the cluster to modify.
//
// - The `Name` and `Description` parameters are optional. You can specify either or both to update the corresponding attributes of the cluster.
//
// @param request - ModifyHiveAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHiveAttributeResponse
func (client *Client) ModifyHiveAttributeWithOptions(request *ModifyHiveAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyHiveAttributeResponse, _err error) {
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

	if !dara.IsNil(request.HiveId) {
		query["HiveId"] = request.HiveId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHiveAttribute"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHiveAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name or description of a specified cluster.
//
// Description:
//
// ## Request
//
// - This API modifies the name and/or description of an existing cluster.
//
// - `HiveId` is a required parameter that identifies the cluster to modify.
//
// - The `Name` and `Description` parameters are optional. You can specify either or both to update the corresponding attributes of the cluster.
//
// @param request - ModifyHiveAttributeRequest
//
// @return ModifyHiveAttributeResponse
func (client *Client) ModifyHiveAttribute(request *ModifyHiveAttributeRequest) (_result *ModifyHiveAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHiveAttributeResponse{}
	_body, _err := client.ModifyHiveAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information of a parent platform.
//
// @param request - ModifyParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParentPlatformResponse
func (client *Client) ModifyParentPlatformWithOptions(request *ModifyParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *ModifyParentPlatformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.ClientAuth) {
		query["ClientAuth"] = request.ClientAuth
	}

	if !dara.IsNil(request.ClientPassword) {
		query["ClientPassword"] = request.ClientPassword
	}

	if !dara.IsNil(request.ClientUsername) {
		query["ClientUsername"] = request.ClientUsername
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GbId) {
		query["GbId"] = request.GbId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyParentPlatform"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a parent platform.
//
// @param request - ModifyParentPlatformRequest
//
// @return ModifyParentPlatformResponse
func (client *Client) ModifyParentPlatform(request *ModifyParentPlatformRequest) (_result *ModifyParentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyParentPlatformResponse{}
	_body, _err := client.ModifyParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Change the billing method for a Graphic Computing Service instance.
//
// @param request - ModifyRenderingChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRenderingChargeTypeResponse
func (client *Client) ModifyRenderingChargeTypeWithOptions(request *ModifyRenderingChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRenderingChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRenderingChargeType"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRenderingChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Change the billing method for a Graphic Computing Service instance.
//
// @param request - ModifyRenderingChargeTypeRequest
//
// @return ModifyRenderingChargeTypeResponse
func (client *Client) ModifyRenderingChargeType(request *ModifyRenderingChargeTypeRequest) (_result *ModifyRenderingChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRenderingChargeTypeResponse{}
	_body, _err := client.ModifyRenderingChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades a cloud application service instance.
//
// @param request - ModifyRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRenderingInstanceResponse
func (client *Client) ModifyRenderingInstanceWithOptions(request *ModifyRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyRenderingInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.RenderingSpec) {
		query["RenderingSpec"] = request.RenderingSpec
	}

	if !dara.IsNil(request.StorageSize) {
		query["StorageSize"] = request.StorageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRenderingInstance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRenderingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades a cloud application service instance.
//
// @param request - ModifyRenderingInstanceRequest
//
// @return ModifyRenderingInstanceResponse
func (client *Client) ModifyRenderingInstance(request *ModifyRenderingInstanceRequest) (_result *ModifyRenderingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRenderingInstanceResponse{}
	_body, _err := client.ModifyRenderingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a cloud application service instance.
//
// @param request - ModifyRenderingInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRenderingInstanceAttributeResponse
func (client *Client) ModifyRenderingInstanceAttributeWithOptions(request *ModifyRenderingInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRenderingInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRenderingInstanceAttribute"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRenderingInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a cloud application service instance.
//
// @param request - ModifyRenderingInstanceAttributeRequest
//
// @return ModifyRenderingInstanceAttributeResponse
func (client *Client) ModifyRenderingInstanceAttribute(request *ModifyRenderingInstanceAttributeRequest) (_result *ModifyRenderingInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRenderingInstanceAttributeResponse{}
	_body, _err := client.ModifyRenderingInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the rate limiting bandwidth for a cloud application service instance. You can call the DescribeRenderingInstance operation to retrieve the current rate limiting value and check the status of the rate limiting update.
//
// @param request - ModifyRenderingInstanceBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRenderingInstanceBandwidthResponse
func (client *Client) ModifyRenderingInstanceBandwidthWithOptions(request *ModifyRenderingInstanceBandwidthRequest, runtime *dara.RuntimeOptions) (_result *ModifyRenderingInstanceBandwidthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxEgressBandwidth) {
		query["MaxEgressBandwidth"] = request.MaxEgressBandwidth
	}

	if !dara.IsNil(request.MaxIngressBandwidth) {
		query["MaxIngressBandwidth"] = request.MaxIngressBandwidth
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRenderingInstanceBandwidth"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRenderingInstanceBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the rate limiting bandwidth for a cloud application service instance. You can call the DescribeRenderingInstance operation to retrieve the current rate limiting value and check the status of the rate limiting update.
//
// @param request - ModifyRenderingInstanceBandwidthRequest
//
// @return ModifyRenderingInstanceBandwidthResponse
func (client *Client) ModifyRenderingInstanceBandwidth(request *ModifyRenderingInstanceBandwidthRequest) (_result *ModifyRenderingInstanceBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRenderingInstanceBandwidthResponse{}
	_body, _err := client.ModifyRenderingInstanceBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies template information.
//
// @param request - ModifyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplateWithOptions(request *ModifyTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileFormat) {
		query["FileFormat"] = request.FileFormat
	}

	if !dara.IsNil(request.Flv) {
		query["Flv"] = request.Flv
	}

	if !dara.IsNil(request.HlsM3u8) {
		query["HlsM3u8"] = request.HlsM3u8
	}

	if !dara.IsNil(request.HlsTs) {
		query["HlsTs"] = request.HlsTs
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.JpgOnDemand) {
		query["JpgOnDemand"] = request.JpgOnDemand
	}

	if !dara.IsNil(request.JpgOverwrite) {
		query["JpgOverwrite"] = request.JpgOverwrite
	}

	if !dara.IsNil(request.JpgSequence) {
		query["JpgSequence"] = request.JpgSequence
	}

	if !dara.IsNil(request.Mp4) {
		query["Mp4"] = request.Mp4
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssFilePrefix) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	if !dara.IsNil(request.TransConfigsJSON) {
		query["TransConfigsJSON"] = request.TransConfigsJSON
	}

	if !dara.IsNil(request.Trigger) {
		query["Trigger"] = request.Trigger
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies template information.
//
// @param request - ModifyTemplateRequest
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplate(request *ModifyTemplateRequest) (_result *ModifyTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.ModifyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves the specified cloud application service instances from their current cluster to the target Hive.
//
// Description:
//
// ## Request
//
// - Ensure the target Hive has sufficient resources to accommodate the instances.
//
// @param tmpReq - MoveHiveEdgeWorkersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveHiveEdgeWorkersResponse
func (client *Client) MoveHiveEdgeWorkersWithOptions(tmpReq *MoveHiveEdgeWorkersRequest, runtime *dara.RuntimeOptions) (_result *MoveHiveEdgeWorkersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MoveHiveEdgeWorkersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HiveId) {
		query["HiveId"] = request.HiveId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveHiveEdgeWorkers"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveHiveEdgeWorkersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves the specified cloud application service instances from their current cluster to the target Hive.
//
// Description:
//
// ## Request
//
// - Ensure the target Hive has sufficient resources to accommodate the instances.
//
// @param request - MoveHiveEdgeWorkersRequest
//
// @return MoveHiveEdgeWorkersResponse
func (client *Client) MoveHiveEdgeWorkers(request *MoveHiveEdgeWorkersRequest) (_result *MoveHiveEdgeWorkersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveHiveEdgeWorkersResponse{}
	_body, _err := client.MoveHiveEdgeWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates the service.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenVsServiceResponse
func (client *Client) OpenVsServiceWithOptions(runtime *dara.RuntimeOptions) (_result *OpenVsServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("OpenVsService"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenVsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the service.
//
// @return OpenVsServiceResponse
func (client *Client) OpenVsService() (_result *OpenVsServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenVsServiceResponse{}
	_body, _err := client.OpenVsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Push a file to a specified cloud application service instance. This is an asynchronous operation. You can query the push progress using the ListFilePushStatuses operation.
//
// @param request - PushFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushFileResponse
func (client *Client) PushFileWithOptions(request *PushFileRequest, runtime *dara.RuntimeOptions) (_result *PushFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushFile"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Push a file to a specified cloud application service instance. This is an asynchronous operation. You can query the push progress using the ListFilePushStatuses operation.
//
// @param request - PushFileRequest
//
// @return PushFileResponse
func (client *Client) PushFile(request *PushFileRequest) (_result *PushFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushFileResponse{}
	_body, _err := client.PushFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a cloud application service instance. You can call the DescribeRenderingInstance API to monitor the restart progress.
//
// @param request - RebootRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootRenderingInstanceResponse
func (client *Client) RebootRenderingInstanceWithOptions(request *RebootRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebootRenderingInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootRenderingInstance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootRenderingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a cloud application service instance. You can call the DescribeRenderingInstance API to monitor the restart progress.
//
// @param request - RebootRenderingInstanceRequest
//
// @return RebootRenderingInstanceResponse
func (client *Client) RebootRenderingInstance(request *RebootRenderingInstanceRequest) (_result *RebootRenderingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebootRenderingInstanceResponse{}
	_body, _err := client.RebootRenderingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts the host of a cloud application service instance.
//
// @param tmpReq - RebootRenderingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootRenderingServerResponse
func (client *Client) RebootRenderingServerWithOptions(tmpReq *RebootRenderingServerRequest, runtime *dara.RuntimeOptions) (_result *RebootRenderingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RebootRenderingServerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RenderingInstanceIds) {
		request.RenderingInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RenderingInstanceIds, dara.String("RenderingInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceIdsShrink) {
		query["RenderingInstanceIds"] = request.RenderingInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootRenderingServer"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootRenderingServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts the host of a cloud application service instance.
//
// @param request - RebootRenderingServerRequest
//
// @return RebootRenderingServerResponse
func (client *Client) RebootRenderingServer(request *RebootRenderingServerRequest) (_result *RebootRenderingServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebootRenderingServerResponse{}
	_body, _err := client.RebootRenderingServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Recover data to a Graphic Computing Service instance
//
// @param request - RecoverRenderingDataPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverRenderingDataPackageResponse
func (client *Client) RecoverRenderingDataPackageWithOptions(request *RecoverRenderingDataPackageRequest, runtime *dara.RuntimeOptions) (_result *RecoverRenderingDataPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataPackageId) {
		query["DataPackageId"] = request.DataPackageId
	}

	if !dara.IsNil(request.LoadMode) {
		query["LoadMode"] = request.LoadMode
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecoverRenderingDataPackage"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoverRenderingDataPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Recover data to a Graphic Computing Service instance
//
// @param request - RecoverRenderingDataPackageRequest
//
// @return RecoverRenderingDataPackageResponse
func (client *Client) RecoverRenderingDataPackage(request *RecoverRenderingDataPackageRequest) (_result *RecoverRenderingDataPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecoverRenderingDataPackageResponse{}
	_body, _err := client.RecoverRenderingDataPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call RefreshRenderingInstanceStreaming to refresh the stream connection for a cloud application service instance.
//
// Description:
//
// > Specify at least one of the template ID or template type.
//
// @param tmpReq - RefreshRenderingInstanceStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshRenderingInstanceStreamingResponse
func (client *Client) RefreshRenderingInstanceStreamingWithOptions(tmpReq *RefreshRenderingInstanceStreamingRequest, runtime *dara.RuntimeOptions) (_result *RefreshRenderingInstanceStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RefreshRenderingInstanceStreamingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClientInfo) {
		request.ClientInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientInfo, dara.String("ClientInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientInfoShrink) {
		query["ClientInfo"] = request.ClientInfoShrink
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshRenderingInstanceStreaming"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshRenderingInstanceStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call RefreshRenderingInstanceStreaming to refresh the stream connection for a cloud application service instance.
//
// Description:
//
// > Specify at least one of the template ID or template type.
//
// @param request - RefreshRenderingInstanceStreamingRequest
//
// @return RefreshRenderingInstanceStreamingResponse
func (client *Client) RefreshRenderingInstanceStreaming(request *RefreshRenderingInstanceStreamingRequest) (_result *RefreshRenderingInstanceStreamingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshRenderingInstanceStreamingResponse{}
	_body, _err := client.RefreshRenderingInstanceStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Release a cloud application service data pack
//
// @param request - ReleaseRenderingDataPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseRenderingDataPackageResponse
func (client *Client) ReleaseRenderingDataPackageWithOptions(request *ReleaseRenderingDataPackageRequest, runtime *dara.RuntimeOptions) (_result *ReleaseRenderingDataPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataPackageId) {
		query["DataPackageId"] = request.DataPackageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseRenderingDataPackage"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseRenderingDataPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Release a cloud application service data pack
//
// @param request - ReleaseRenderingDataPackageRequest
//
// @return ReleaseRenderingDataPackageResponse
func (client *Client) ReleaseRenderingDataPackage(request *ReleaseRenderingDataPackageRequest) (_result *ReleaseRenderingDataPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseRenderingDataPackageResponse{}
	_body, _err := client.ReleaseRenderingDataPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invoke ReleaseRenderingInstance to release a Graphic Computing Service application instance.
//
// @param request - ReleaseRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseRenderingInstanceResponse
func (client *Client) ReleaseRenderingInstanceWithOptions(request *ReleaseRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseRenderingInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseRenderingInstance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseRenderingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke ReleaseRenderingInstance to release a Graphic Computing Service application instance.
//
// @param request - ReleaseRenderingInstanceRequest
//
// @return ReleaseRenderingInstanceResponse
func (client *Client) ReleaseRenderingInstance(request *ReleaseRenderingInstanceRequest) (_result *ReleaseRenderingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseRenderingInstanceResponse{}
	_body, _err := client.ReleaseRenderingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invoke RenewRenderingInstance to renew a cloud application service instance.
//
// @param request - RenewRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewRenderingInstanceResponse
func (client *Client) RenewRenderingInstanceWithOptions(request *RenewRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewRenderingInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewRenderingInstance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewRenderingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke RenewRenderingInstance to renew a cloud application service instance.
//
// @param request - RenewRenderingInstanceRequest
//
// @return RenewRenderingInstanceResponse
func (client *Client) RenewRenderingInstance(request *RenewRenderingInstanceRequest) (_result *RenewRenderingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewRenderingInstanceResponse{}
	_body, _err := client.RenewRenderingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a cloud application service instance. You can query the DescribeRenderingInstance interface to obtain the reset progress.
//
// @param request - ResetRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetRenderingInstanceResponse
func (client *Client) ResetRenderingInstanceWithOptions(request *ResetRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *ResetRenderingInstanceResponse, _err error) {
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

	if !dara.IsNil(request.DataPackageId) {
		query["DataPackageId"] = request.DataPackageId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetRenderingInstance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetRenderingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a cloud application service instance. You can query the DescribeRenderingInstance interface to obtain the reset progress.
//
// @param request - ResetRenderingInstanceRequest
//
// @return ResetRenderingInstanceResponse
func (client *Client) ResetRenderingInstance(request *ResetRenderingInstanceRequest) (_result *ResetRenderingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetRenderingInstanceResponse{}
	_body, _err := client.ResetRenderingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes pushing for a stream.
//
// @param request - ResumeVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeVsStreamResponse
func (client *Client) ResumeVsStreamWithOptions(request *ResumeVsStreamRequest, runtime *dara.RuntimeOptions) (_result *ResumeVsStreamResponse, _err error) {
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

	if !dara.IsNil(request.ControlStreamAction) {
		query["ControlStreamAction"] = request.ControlStreamAction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.LiveStreamType) {
		query["LiveStreamType"] = request.LiveStreamType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeVsStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeVsStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes pushing for a stream.
//
// @param request - ResumeVsStreamRequest
//
// @return ResumeVsStreamResponse
func (client *Client) ResumeVsStream(request *ResumeVsStreamRequest) (_result *ResumeVsStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeVsStreamResponse{}
	_body, _err := client.ResumeVsStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends shell control instructions to a cloud application service instance. This operation supports both sync and asynchronous command responses. The sync scenario is not suitable for time-consuming commands. The maximum execution time cannot exceed 30 s. In an asynchronous scenario, you can call the GetRenderingInstanceCommandsStatus operation to query the execution status and result of a command.
//
// @param request - SendRenderingInstanceCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRenderingInstanceCommandsResponse
func (client *Client) SendRenderingInstanceCommandsWithOptions(request *SendRenderingInstanceCommandsRequest, runtime *dara.RuntimeOptions) (_result *SendRenderingInstanceCommandsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Commands) {
		body["Commands"] = request.Commands
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendRenderingInstanceCommands"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendRenderingInstanceCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends shell control instructions to a cloud application service instance. This operation supports both sync and asynchronous command responses. The sync scenario is not suitable for time-consuming commands. The maximum execution time cannot exceed 30 s. In an asynchronous scenario, you can call the GetRenderingInstanceCommandsStatus operation to query the execution status and result of a command.
//
// @param request - SendRenderingInstanceCommandsRequest
//
// @return SendRenderingInstanceCommandsResponse
func (client *Client) SendRenderingInstanceCommands(request *SendRenderingInstanceCommandsRequest) (_result *SendRenderingInstanceCommandsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendRenderingInstanceCommandsResponse{}
	_body, _err := client.SendRenderingInstanceCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set a preset position.
//
// @param request - SetPresetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPresetResponse
func (client *Client) SetPresetWithOptions(request *SetPresetRequest, runtime *dara.RuntimeOptions) (_result *SetPresetResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PresetId) {
		query["PresetId"] = request.PresetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPreset"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPresetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set a preset position.
//
// @param request - SetPresetRequest
//
// @return SetPresetResponse
func (client *Client) SetPreset(request *SetPresetRequest) (_result *SetPresetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPresetResponse{}
	_body, _err := client.SetPresetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enable or disable the certificate feature for a domain name.
//
// @param request - SetVsDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVsDomainCertificateResponse
func (client *Client) SetVsDomainCertificateWithOptions(request *SetVsDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetVsDomainCertificateResponse, _err error) {
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

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ForceSet) {
		query["ForceSet"] = request.ForceSet
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVsDomainCertificate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVsDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enable or disable the certificate feature for a domain name.
//
// @param request - SetVsDomainCertificateRequest
//
// @return SetVsDomainCertificateResponse
func (client *Client) SetVsDomainCertificate(request *SetVsDomainCertificateRequest) (_result *SetVsDomainCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetVsDomainCertificateResponse{}
	_body, _err := client.SetVsDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configure stream ingest callbacks.
//
// @param request - SetVsStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVsStreamsNotifyUrlConfigResponse
func (client *Client) SetVsStreamsNotifyUrlConfigWithOptions(request *SetVsStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *SetVsStreamsNotifyUrlConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVsStreamsNotifyUrlConfig"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configure stream ingest callbacks.
//
// @param request - SetVsStreamsNotifyUrlConfigRequest
//
// @return SetVsStreamsNotifyUrlConfigResponse
func (client *Client) SetVsStreamsNotifyUrlConfig(request *SetVsStreamsNotifyUrlConfigRequest) (_result *SetVsStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetVsStreamsNotifyUrlConfigResponse{}
	_body, _err := client.SetVsStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start stream pulling from a device. This action starts all streams on the device.
//
// Description:
//
// Each device currently supports only one ingest endpoint. The effect is the same as StartStream.
//
// @param request - StartDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDeviceResponse
func (client *Client) StartDeviceWithOptions(request *StartDeviceRequest, runtime *dara.RuntimeOptions) (_result *StartDeviceResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start stream pulling from a device. This action starts all streams on the device.
//
// Description:
//
// Each device currently supports only one ingest endpoint. The effect is the same as StartStream.
//
// @param request - StartDeviceRequest
//
// @return StartDeviceResponse
func (client *Client) StartDevice(request *StartDeviceRequest) (_result *StartDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDeviceResponse{}
	_body, _err := client.StartDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts interactions with the parent platform, such as registration and keep-alive.
//
// @param request - StartParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartParentPlatformResponse
func (client *Client) StartParentPlatformWithOptions(request *StartParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *StartParentPlatformResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartParentPlatform"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartParentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts interactions with the parent platform, such as registration and keep-alive.
//
// @param request - StartParentPlatformRequest
//
// @return StartParentPlatformResponse
func (client *Client) StartParentPlatform(request *StartParentPlatformRequest) (_result *StartParentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartParentPlatformResponse{}
	_body, _err := client.StartParentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartPublishStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPublishStreamResponse
func (client *Client) StartPublishStreamWithOptions(request *StartPublishStreamRequest, runtime *dara.RuntimeOptions) (_result *StartPublishStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PublishUrl) {
		query["PublishUrl"] = request.PublishUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPublishStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPublishStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartPublishStreamRequest
//
// @return StartPublishStreamResponse
func (client *Client) StartPublishStream(request *StartPublishStreamRequest) (_result *StartPublishStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartPublishStreamResponse{}
	_body, _err := client.StartPublishStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts on-demand recording for the specified stream.
//
// Description:
//
// > - An on-demand record template is required. You must first attach one to the space or stream.
//
// >
//
// > - You can specify a stream in two ways: using its ID or its PlayDomain/App/Name.
//
// @param request - StartRecordStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordStreamResponse
func (client *Client) StartRecordStreamWithOptions(request *StartRecordStreamRequest, runtime *dara.RuntimeOptions) (_result *StartRecordStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRecordStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRecordStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts on-demand recording for the specified stream.
//
// Description:
//
// > - An on-demand record template is required. You must first attach one to the space or stream.
//
// >
//
// > - You can specify a stream in two ways: using its ID or its PlayDomain/App/Name.
//
// @param request - StartRecordStreamRequest
//
// @return StartRecordStreamResponse
func (client *Client) StartRecordStream(request *StartRecordStreamRequest) (_result *StartRecordStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRecordStreamResponse{}
	_body, _err := client.StartRecordStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Schedules an idle cloud application service instance for the requesting client (ClientId) and starts the service. If the requesting client (ClientId) sends another start request after a successful start and the associated session is in the SessionStartSuspended state, the session is restarted. If the session is in any other state, the session information is returned directly.
//
// @param tmpReq - StartRenderingSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRenderingSessionResponse
func (client *Client) StartRenderingSessionWithOptions(tmpReq *StartRenderingSessionRequest, runtime *dara.RuntimeOptions) (_result *StartRenderingSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartRenderingSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClientParams) {
		request.ClientParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientParams, dara.String("ClientParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientParamsShrink) {
		query["ClientParams"] = request.ClientParamsShrink
	}

	if !dara.IsNil(request.PatchId) {
		query["PatchId"] = request.PatchId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRenderingSession"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRenderingSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Schedules an idle cloud application service instance for the requesting client (ClientId) and starts the service. If the requesting client (ClientId) sends another start request after a successful start and the associated session is in the SessionStartSuspended state, the session is restarted. If the session is in any other state, the session information is returned directly.
//
// @param request - StartRenderingSessionRequest
//
// @return StartRenderingSessionResponse
func (client *Client) StartRenderingSession(request *StartRenderingSessionRequest) (_result *StartRenderingSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRenderingSessionResponse{}
	_body, _err := client.StartRenderingSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start a stream.
//
// @param request - StartStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartStreamResponse
func (client *Client) StartStreamWithOptions(request *StartStreamRequest, runtime *dara.RuntimeOptions) (_result *StartStreamResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
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
		Action:      dara.String("StartStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start a stream.
//
// @param request - StartStreamRequest
//
// @return StartStreamResponse
func (client *Client) StartStream(request *StartStreamRequest) (_result *StartStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartStreamResponse{}
	_body, _err := client.StartStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts forwarding a stream to an external address.
//
// @param request - StartTransferStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTransferStreamResponse
func (client *Client) StartTransferStreamWithOptions(request *StartTransferStreamRequest, runtime *dara.RuntimeOptions) (_result *StartTransferStreamResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Transcode) {
		query["Transcode"] = request.Transcode
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTransferStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTransferStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts forwarding a stream to an external address.
//
// @param request - StartTransferStreamRequest
//
// @return StartTransferStreamResponse
func (client *Client) StartTransferStream(request *StartTransferStreamRequest) (_result *StartTransferStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartTransferStreamResponse{}
	_body, _err := client.StartTransferStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops lens adjustments, such as aperture or zoom changes.
//
// @param request - StopAdjustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAdjustResponse
func (client *Client) StopAdjustWithOptions(request *StopAdjustRequest, runtime *dara.RuntimeOptions) (_result *StopAdjustResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Focus) {
		query["Focus"] = request.Focus
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Iris) {
		query["Iris"] = request.Iris
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAdjust"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAdjustResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops lens adjustments, such as aperture or zoom changes.
//
// @param request - StopAdjustRequest
//
// @return StopAdjustResponse
func (client *Client) StopAdjust(request *StopAdjustRequest) (_result *StopAdjustResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopAdjustResponse{}
	_body, _err := client.StopAdjustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops stream pulling for a device. This operation terminates all streams on that device.
//
// Description:
//
// Stops stream pulling for a device. This operation terminates all streams on that device.
//
// @param request - StopDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDeviceResponse
func (client *Client) StopDeviceWithOptions(request *StopDeviceRequest, runtime *dara.RuntimeOptions) (_result *StopDeviceResponse, _err error) {
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
		Action:      dara.String("StopDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops stream pulling for a device. This operation terminates all streams on that device.
//
// Description:
//
// Stops stream pulling for a device. This operation terminates all streams on that device.
//
// @param request - StopDeviceRequest
//
// @return StopDeviceResponse
func (client *Client) StopDevice(request *StopDeviceRequest) (_result *StopDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDeviceResponse{}
	_body, _err := client.StopDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops camera movement, such as panning, tilting, and zooming.
//
// @param request - StopMoveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMoveResponse
func (client *Client) StopMoveWithOptions(request *StopMoveRequest, runtime *dara.RuntimeOptions) (_result *StopMoveResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Pan) {
		query["Pan"] = request.Pan
	}

	if !dara.IsNil(request.Tilt) {
		query["Tilt"] = request.Tilt
	}

	if !dara.IsNil(request.Zoom) {
		query["Zoom"] = request.Zoom
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMove"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMoveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops camera movement, such as panning, tilting, and zooming.
//
// @param request - StopMoveRequest
//
// @return StopMoveResponse
func (client *Client) StopMove(request *StopMoveRequest) (_result *StopMoveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopMoveResponse{}
	_body, _err := client.StopMoveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StopPublishStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPublishStreamResponse
func (client *Client) StopPublishStreamWithOptions(request *StopPublishStreamRequest, runtime *dara.RuntimeOptions) (_result *StopPublishStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopPublishStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopPublishStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopPublishStreamRequest
//
// @return StopPublishStreamResponse
func (client *Client) StopPublishStream(request *StopPublishStreamRequest) (_result *StopPublishStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopPublishStreamResponse{}
	_body, _err := client.StopPublishStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops on-demand recording for a specified stream.
//
// Description:
//
// > You can specify a stream by ID or by PlayDomain/App/Name.
//
// @param request - StopRecordStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordStreamResponse
func (client *Client) StopRecordStreamWithOptions(request *StopRecordStreamRequest, runtime *dara.RuntimeOptions) (_result *StopRecordStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.App) {
		query["App"] = request.App
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRecordStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRecordStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops on-demand recording for a specified stream.
//
// Description:
//
// > You can specify a stream by ID or by PlayDomain/App/Name.
//
// @param request - StopRecordStreamRequest
//
// @return StopRecordStreamResponse
func (client *Client) StopRecordStream(request *StopRecordStreamRequest) (_result *StopRecordStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopRecordStreamResponse{}
	_body, _err := client.StopRecordStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Shut down the specified cloud application service session and revoke the associated instance resources.
//
// Description:
//
// ## Request information
//
// @param request - StopRenderingSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRenderingSessionResponse
func (client *Client) StopRenderingSessionWithOptions(request *StopRenderingSessionRequest, runtime *dara.RuntimeOptions) (_result *StopRenderingSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRenderingSession"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRenderingSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shut down the specified cloud application service session and revoke the associated instance resources.
//
// Description:
//
// ## Request information
//
// @param request - StopRenderingSessionRequest
//
// @return StopRenderingSessionResponse
func (client *Client) StopRenderingSession(request *StopRenderingSessionRequest) (_result *StopRenderingSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopRenderingSessionResponse{}
	_body, _err := client.StopRenderingSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a stream.
//
// @param request - StopStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStreamResponse
func (client *Client) StopStreamWithOptions(request *StopStreamRequest, runtime *dara.RuntimeOptions) (_result *StopStreamResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
		Action:      dara.String("StopStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a stream.
//
// @param request - StopStreamRequest
//
// @return StopStreamResponse
func (client *Client) StopStream(request *StopStreamRequest) (_result *StopStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopStreamResponse{}
	_body, _err := client.StopStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a stream.
//
// @param request - StopTransferStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTransferStreamResponse
func (client *Client) StopTransferStreamWithOptions(request *StopTransferStreamRequest, runtime *dara.RuntimeOptions) (_result *StopTransferStreamResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Transcode) {
		query["Transcode"] = request.Transcode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTransferStream"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTransferStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a stream.
//
// @param request - StopTransferStreamRequest
//
// @return StopTransferStreamResponse
func (client *Client) StopTransferStream(request *StopTransferStreamRequest) (_result *StopTransferStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopTransferStreamResponse{}
	_body, _err := client.StopTransferStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes platform channel information.
//
// @param request - SyncCatalogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncCatalogsResponse
func (client *Client) SyncCatalogsWithOptions(request *SyncCatalogsRequest, runtime *dara.RuntimeOptions) (_result *SyncCatalogsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncCatalogs"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncCatalogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes platform channel information.
//
// @param request - SyncCatalogsRequest
//
// @return SyncCatalogsResponse
func (client *Client) SyncCatalogs(request *SyncCatalogsRequest) (_result *SyncCatalogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncCatalogsResponse{}
	_body, _err := client.SyncCatalogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detach a device from a folder.
//
// @param request - UnbindDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDirectoryResponse
func (client *Client) UnbindDirectoryWithOptions(request *UnbindDirectoryRequest, runtime *dara.RuntimeOptions) (_result *UnbindDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindDirectory"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach a device from a folder.
//
// @param request - UnbindDirectoryRequest
//
// @return UnbindDirectoryResponse
func (client *Client) UnbindDirectory(request *UnbindDirectoryRequest) (_result *UnbindDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindDirectoryResponse{}
	_body, _err := client.UnbindDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dissociates a device from a parent platform push configuration so that the device is no longer pushed.
//
// @param request - UnbindParentPlatformDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindParentPlatformDeviceResponse
func (client *Client) UnbindParentPlatformDeviceWithOptions(request *UnbindParentPlatformDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnbindParentPlatformDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParentPlatformId) {
		query["ParentPlatformId"] = request.ParentPlatformId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindParentPlatformDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindParentPlatformDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a device from a parent platform push configuration so that the device is no longer pushed.
//
// @param request - UnbindParentPlatformDeviceRequest
//
// @return UnbindParentPlatformDeviceResponse
func (client *Client) UnbindParentPlatformDevice(request *UnbindParentPlatformDeviceRequest) (_result *UnbindParentPlatformDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindParentPlatformDeviceResponse{}
	_body, _err := client.UnbindParentPlatformDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a purchased device from a space.
//
// @param request - UnbindPurchasedDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPurchasedDeviceResponse
func (client *Client) UnbindPurchasedDeviceWithOptions(request *UnbindPurchasedDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnbindPurchasedDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindPurchasedDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindPurchasedDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a purchased device from a space.
//
// @param request - UnbindPurchasedDeviceRequest
//
// @return UnbindPurchasedDeviceResponse
func (client *Client) UnbindPurchasedDevice(request *UnbindPurchasedDeviceRequest) (_result *UnbindPurchasedDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindPurchasedDeviceResponse{}
	_body, _err := client.UnbindPurchasedDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbind a template from a specified instance, such as a group instance or a stream.
//
// Description:
//
// > Specify at least one of TemplateId or TemplateType.
//
// @param request - UnbindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindTemplateResponse
func (client *Client) UnbindTemplateWithOptions(request *UnbindTemplateRequest, runtime *dara.RuntimeOptions) (_result *UnbindTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindTemplate"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbind a template from a specified instance, such as a group instance or a stream.
//
// Description:
//
// > Specify at least one of TemplateId or TemplateType.
//
// @param request - UnbindTemplateRequest
//
// @return UnbindTemplateResponse
func (client *Client) UnbindTemplate(request *UnbindTemplateRequest) (_result *UnbindTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindTemplateResponse{}
	_body, _err := client.UnbindTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can uninstall a specified cloud application from a specified cloud application instance. This operation is asynchronous. You can use the ListCloudAppInstallations operation to check the uninstallation progress. After successful uninstallation, the query operation no longer returns related information.
//
// @param tmpReq - UninstallCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallCloudAppResponse
func (client *Client) UninstallCloudAppWithOptions(tmpReq *UninstallCloudAppRequest, runtime *dara.RuntimeOptions) (_result *UninstallCloudAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UninstallCloudAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RenderingInstanceIds) {
		request.RenderingInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RenderingInstanceIds, dara.String("RenderingInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PatchId) {
		query["PatchId"] = request.PatchId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.RenderingInstanceIdsShrink) {
		query["RenderingInstanceIds"] = request.RenderingInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallCloudApp"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallCloudAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can uninstall a specified cloud application from a specified cloud application instance. This operation is asynchronous. You can use the ListCloudAppInstallations operation to check the uninstallation progress. After successful uninstallation, the query operation no longer returns related information.
//
// @param request - UninstallCloudAppRequest
//
// @return UninstallCloudAppResponse
func (client *Client) UninstallCloudApp(request *UninstallCloudAppRequest) (_result *UninstallCloudAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallCloudAppResponse{}
	_body, _err := client.UninstallCloudAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlock a device.
//
// @param request - UnlockDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockDeviceResponse
func (client *Client) UnlockDeviceWithOptions(request *UnlockDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnlockDeviceResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockDevice"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlock a device.
//
// @param request - UnlockDeviceRequest
//
// @return UnlockDeviceResponse
func (client *Client) UnlockDevice(request *UnlockDeviceRequest) (_result *UnlockDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockDeviceResponse{}
	_body, _err := client.UnlockDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information for a cloud application, such as its description and tags. You can upload patch or hotfix packages and create hotfix packages for the Android cloud application marketplace. A cloud application supports up to 20 patch packages, but only one package can be in the uploading state at a time.
//
// @param tmpReq - UpdateCloudAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudAppInfoResponse
func (client *Client) UpdateCloudAppInfoWithOptions(tmpReq *UpdateCloudAppInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudAppInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudAppInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Patch) {
		request.PatchShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Patch, dara.String("Patch"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PkgLabels) {
		request.PkgLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PkgLabels, dara.String("PkgLabels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.PkgLabelsShrink) {
		query["PkgLabels"] = request.PkgLabelsShrink
	}

	if !dara.IsNil(request.StablePatchId) {
		query["StablePatchId"] = request.StablePatchId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PatchShrink) {
		body["Patch"] = request.PatchShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudAppInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudAppInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information for a cloud application, such as its description and tags. You can upload patch or hotfix packages and create hotfix packages for the Android cloud application marketplace. A cloud application supports up to 20 patch packages, but only one package can be in the uploading state at a time.
//
// @param request - UpdateCloudAppInfoRequest
//
// @return UpdateCloudAppInfoResponse
func (client *Client) UpdateCloudAppInfo(request *UpdateCloudAppInfoRequest) (_result *UpdateCloudAppInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudAppInfoResponse{}
	_body, _err := client.UpdateCloudAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update basic information for a file, such as its description.
//
// @param request - UpdateFileInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileInfoResponse
func (client *Client) UpdateFileInfoWithOptions(request *UpdateFileInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileInfoResponse, _err error) {
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

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFileInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update basic information for a file, such as its description.
//
// @param request - UpdateFileInfoRequest
//
// @return UpdateFileInfoResponse
func (client *Client) UpdateFileInfo(request *UpdateFileInfoRequest) (_result *UpdateFileInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFileInfoResponse{}
	_body, _err := client.UpdateFileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation updates the configuration parameters of a cloud application service instance. It lets you modify various configurations of the Cloud Android system, such as prop, location, and network, to create a real device simulation.
//
// You can retrieve the configured values for the real device simulation by calling the DescribeRenderingInstance API.
//
// To query the configuration parameters of the real-time environment, see the DescribeRenderingInstanceConfiguration API.
//
// @param tmpReq - UpdateRenderingInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRenderingInstanceConfigurationResponse
func (client *Client) UpdateRenderingInstanceConfigurationWithOptions(tmpReq *UpdateRenderingInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateRenderingInstanceConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRenderingInstanceConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configuration) {
		request.ConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configuration, dara.String("Configuration"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationShrink) {
		body["Configuration"] = request.ConfigurationShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRenderingInstanceConfiguration"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRenderingInstanceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation updates the configuration parameters of a cloud application service instance. It lets you modify various configurations of the Cloud Android system, such as prop, location, and network, to create a real device simulation.
//
// You can retrieve the configured values for the real device simulation by calling the DescribeRenderingInstance API.
//
// To query the configuration parameters of the real-time environment, see the DescribeRenderingInstanceConfiguration API.
//
// @param request - UpdateRenderingInstanceConfigurationRequest
//
// @return UpdateRenderingInstanceConfigurationResponse
func (client *Client) UpdateRenderingInstanceConfiguration(request *UpdateRenderingInstanceConfigurationRequest) (_result *UpdateRenderingInstanceConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRenderingInstanceConfigurationResponse{}
	_body, _err := client.UpdateRenderingInstanceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the settings of a cloud application service instance.
//
// @param tmpReq - UpdateRenderingInstanceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRenderingInstanceSettingsResponse
func (client *Client) UpdateRenderingInstanceSettingsWithOptions(tmpReq *UpdateRenderingInstanceSettingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateRenderingInstanceSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRenderingInstanceSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Settings) {
		request.SettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Settings, dara.String("Settings"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderingInstanceId) {
		query["RenderingInstanceId"] = request.RenderingInstanceId
	}

	if !dara.IsNil(request.SettingsShrink) {
		query["Settings"] = request.SettingsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRenderingInstanceSettings"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRenderingInstanceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the settings of a cloud application service instance.
//
// @param request - UpdateRenderingInstanceSettingsRequest
//
// @return UpdateRenderingInstanceSettingsResponse
func (client *Client) UpdateRenderingInstanceSettings(request *UpdateRenderingInstanceSettingsRequest) (_result *UpdateRenderingInstanceSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRenderingInstanceSettingsResponse{}
	_body, _err := client.UpdateRenderingInstanceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a project’s properties.
//
// @param tmpReq - UpdateRenderingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRenderingProjectResponse
func (client *Client) UpdateRenderingProjectWithOptions(tmpReq *UpdateRenderingProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateRenderingProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRenderingProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionAttribs) {
		request.SessionAttribsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionAttribs, dara.String("SessionAttribs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SessionAttribsShrink) {
		query["SessionAttribs"] = request.SessionAttribsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRenderingProject"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRenderingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a project’s properties.
//
// @param request - UpdateRenderingProjectRequest
//
// @return UpdateRenderingProjectResponse
func (client *Client) UpdateRenderingProject(request *UpdateRenderingProjectRequest) (_result *UpdateRenderingProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRenderingProjectResponse{}
	_body, _err := client.UpdateRenderingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration for stream pulling. You can modify the start and end times of origin server addresses in an existing stream pulling task.
//
// @param request - UpdateVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVsPullStreamInfoConfigResponse
func (client *Client) UpdateVsPullStreamInfoConfigWithOptions(request *UpdateVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateVsPullStreamInfoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Always) {
		query["Always"] = request.Always
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceUrl) {
		query["SourceUrl"] = request.SourceUrl
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
		Action:      dara.String("UpdateVsPullStreamInfoConfig"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVsPullStreamInfoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration for stream pulling. You can modify the start and end times of origin server addresses in an existing stream pulling task.
//
// @param request - UpdateVsPullStreamInfoConfigRequest
//
// @return UpdateVsPullStreamInfoConfigResponse
func (client *Client) UpdateVsPullStreamInfoConfig(request *UpdateVsPullStreamInfoConfigRequest) (_result *UpdateVsPullStreamInfoConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVsPullStreamInfoConfigResponse{}
	_body, _err := client.UpdateVsPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upload or list a cloud application package. This is an asynchronous API. Use the ListCloudApps API to check upload progress.
//
// @param tmpReq - UploadCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCloudAppResponse
func (client *Client) UploadCloudAppWithOptions(tmpReq *UploadCloudAppRequest, runtime *dara.RuntimeOptions) (_result *UploadCloudAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UploadCloudAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PkgLabels) {
		request.PkgLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PkgLabels, dara.String("PkgLabels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DownloadUrl) {
		query["DownloadUrl"] = request.DownloadUrl
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.PkgFormat) {
		query["PkgFormat"] = request.PkgFormat
	}

	if !dara.IsNil(request.PkgLabelsShrink) {
		query["PkgLabels"] = request.PkgLabelsShrink
	}

	if !dara.IsNil(request.PkgType) {
		query["PkgType"] = request.PkgType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadCloudApp"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadCloudAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upload or list a cloud application package. This is an asynchronous API. Use the ListCloudApps API to check upload progress.
//
// @param request - UploadCloudAppRequest
//
// @return UploadCloudAppResponse
func (client *Client) UploadCloudApp(request *UploadCloudAppRequest) (_result *UploadCloudAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadCloudAppResponse{}
	_body, _err := client.UploadCloudAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a file from a public URL to local or cloud storage. This is an asynchronous operation. You can call the ListFiles operation to monitor the upload progress.
//
// @param request - UploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadFileResponse
func (client *Client) UploadFileWithOptions(request *UploadFileRequest, runtime *dara.RuntimeOptions) (_result *UploadFileResponse, _err error) {
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

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.OriginUrl) {
		query["OriginUrl"] = request.OriginUrl
	}

	if !dara.IsNil(request.TargetPath) {
		query["TargetPath"] = request.TargetPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadFile"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a file from a public URL to local or cloud storage. This is an asynchronous operation. You can call the ListFiles operation to monitor the upload progress.
//
// @param request - UploadFileRequest
//
// @return UploadFileResponse
func (client *Client) UploadFile(request *UploadFileRequest) (_result *UploadFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadFileResponse{}
	_body, _err := client.UploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upload a new public key.
//
// @param request - UploadPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadPublicKeyResponse
func (client *Client) UploadPublicKeyWithOptions(request *UploadPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *UploadPublicKeyResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.KeyGroup) {
		query["KeyGroup"] = request.KeyGroup
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadPublicKey"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upload a new public key.
//
// @param request - UploadPublicKeyRequest
//
// @return UploadPublicKeyResponse
func (client *Client) UploadPublicKey(request *UploadPublicKeyRequest) (_result *UploadPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadPublicKeyResponse{}
	_body, _err := client.UploadPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # VerifyVsDomainOwner
//
// @param request - VerifyVsDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyVsDomainOwnerResponse
func (client *Client) VerifyVsDomainOwnerWithOptions(request *VerifyVsDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyVsDomainOwnerResponse, _err error) {
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
		Action:      dara.String("VerifyVsDomainOwner"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyVsDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # VerifyVsDomainOwner
//
// @param request - VerifyVsDomainOwnerRequest
//
// @return VerifyVsDomainOwnerResponse
func (client *Client) VerifyVsDomainOwner(request *VerifyVsDomainOwnerRequest) (_result *VerifyVsDomainOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyVsDomainOwnerResponse{}
	_body, _err := client.VerifyVsDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
