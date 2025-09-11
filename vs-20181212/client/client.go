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

// @param request - AddVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVsPullStreamInfoConfigResponse
func (client *Client) AddVsPullStreamInfoConfigWithOptions(request *AddVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *AddVsPullStreamInfoConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 云应用服务实例与项目进行关联。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于将满足特定条件的实例与指定项目进行关联。
//
// @param tmpReq - AssociateRenderingProjectInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateRenderingProjectInstancesResponse
func (client *Client) AssociateRenderingProjectInstancesWithOptions(tmpReq *AssociateRenderingProjectInstancesRequest, runtime *dara.RuntimeOptions) (_result *AssociateRenderingProjectInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 云应用服务实例与项目进行关联。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于将满足特定条件的实例与指定项目进行关联。
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

// @param request - BatchBindDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindDirectoriesResponse
func (client *Client) BatchBindDirectoriesWithOptions(request *BatchBindDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindDirectoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchBindParentPlatformDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindParentPlatformDevicesResponse
func (client *Client) BatchBindParentPlatformDevicesWithOptions(request *BatchBindParentPlatformDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindParentPlatformDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchBindPurchasedDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindPurchasedDevicesResponse
func (client *Client) BatchBindPurchasedDevicesWithOptions(request *BatchBindPurchasedDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindPurchasedDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchBindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindTemplateResponse
func (client *Client) BatchBindTemplateWithOptions(request *BatchBindTemplateRequest, runtime *dara.RuntimeOptions) (_result *BatchBindTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchBindTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchBindTemplatesResponse
func (client *Client) BatchBindTemplatesWithOptions(request *BatchBindTemplatesRequest, runtime *dara.RuntimeOptions) (_result *BatchBindTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchDeleteDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDevicesResponse
func (client *Client) BatchDeleteDevicesWithOptions(request *BatchDeleteDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchDeleteVsDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteVsDomainConfigsResponse
func (client *Client) BatchDeleteVsDomainConfigsWithOptions(request *BatchDeleteVsDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteVsDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchForbidVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchForbidVsStreamResponse
func (client *Client) BatchForbidVsStreamWithOptions(request *BatchForbidVsStreamRequest, runtime *dara.RuntimeOptions) (_result *BatchForbidVsStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchResumeVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchResumeVsStreamResponse
func (client *Client) BatchResumeVsStreamWithOptions(request *BatchResumeVsStreamRequest, runtime *dara.RuntimeOptions) (_result *BatchResumeVsStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchSetVsDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetVsDomainConfigsResponse
func (client *Client) BatchSetVsDomainConfigsWithOptions(request *BatchSetVsDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetVsDomainConfigsResponse, _err error) {
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

// @param request - BatchStartDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartDevicesResponse
func (client *Client) BatchStartDevicesWithOptions(request *BatchStartDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchStartDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchStartStreamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartStreamsResponse
func (client *Client) BatchStartStreamsWithOptions(request *BatchStartStreamsRequest, runtime *dara.RuntimeOptions) (_result *BatchStartStreamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchStopDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopDevicesResponse
func (client *Client) BatchStopDevicesWithOptions(request *BatchStopDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchStopDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchStopStreamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopStreamsResponse
func (client *Client) BatchStopStreamsWithOptions(request *BatchStopStreamsRequest, runtime *dara.RuntimeOptions) (_result *BatchStopStreamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchUnbindDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindDirectoriesResponse
func (client *Client) BatchUnbindDirectoriesWithOptions(request *BatchUnbindDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindDirectoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchUnbindParentPlatformDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindParentPlatformDevicesResponse
func (client *Client) BatchUnbindParentPlatformDevicesWithOptions(request *BatchUnbindParentPlatformDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindParentPlatformDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchUnbindPurchasedDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindPurchasedDevicesResponse
func (client *Client) BatchUnbindPurchasedDevicesWithOptions(request *BatchUnbindPurchasedDevicesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindPurchasedDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchUnbindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindTemplateResponse
func (client *Client) BatchUnbindTemplateWithOptions(request *BatchUnbindTemplateRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BatchUnbindTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUnbindTemplatesResponse
func (client *Client) BatchUnbindTemplatesWithOptions(request *BatchUnbindTemplatesRequest, runtime *dara.RuntimeOptions) (_result *BatchUnbindTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BindDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindDirectoryResponse
func (client *Client) BindDirectoryWithOptions(request *BindDirectoryRequest, runtime *dara.RuntimeOptions) (_result *BindDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BindParentPlatformDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindParentPlatformDeviceResponse
func (client *Client) BindParentPlatformDeviceWithOptions(request *BindParentPlatformDeviceRequest, runtime *dara.RuntimeOptions) (_result *BindParentPlatformDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BindPurchasedDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPurchasedDeviceResponse
func (client *Client) BindPurchasedDeviceWithOptions(request *BindPurchasedDeviceRequest, runtime *dara.RuntimeOptions) (_result *BindPurchasedDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - BindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindTemplateResponse
func (client *Client) BindTemplateWithOptions(request *BindTemplateRequest, runtime *dara.RuntimeOptions) (_result *BindTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ContinuousAdjustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinuousAdjustResponse
func (client *Client) ContinuousAdjustWithOptions(request *ContinuousAdjustRequest, runtime *dara.RuntimeOptions) (_result *ContinuousAdjustResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ContinuousMoveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinuousMoveResponse
func (client *Client) ContinuousMoveWithOptions(request *ContinuousMoveRequest, runtime *dara.RuntimeOptions) (_result *ContinuousMoveResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CreateDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeviceResponse
func (client *Client) CreateDeviceWithOptions(request *CreateDeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CreateDeviceAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeviceAlarmResponse
func (client *Client) CreateDeviceAlarmWithOptions(request *CreateDeviceAlarmRequest, runtime *dara.RuntimeOptions) (_result *CreateDeviceAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithOptions(request *CreateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CreateParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParentPlatformResponse
func (client *Client) CreateParentPlatformWithOptions(request *CreateParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *CreateParentPlatformResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建云渲染数据包
//
// @param request - CreateRenderingDataPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingDataPackageResponse
func (client *Client) CreateRenderingDataPackageWithOptions(request *CreateRenderingDataPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingDataPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建云渲染数据包
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
// 申请云渲染资源实例
//
// @param tmpReq - CreateRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingInstanceResponse
func (client *Client) CreateRenderingInstanceWithOptions(tmpReq *CreateRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 申请云渲染资源实例
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
// 创建自定义网关
//
// @param request - CreateRenderingInstanceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingInstanceGatewayResponse
func (client *Client) CreateRenderingInstanceGatewayWithOptions(request *CreateRenderingInstanceGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingInstanceGatewayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 创建自定义网关
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
// 创建一个新的云应用服务项目，并设置相关属性。
//
// @param tmpReq - CreateRenderingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRenderingProjectResponse
func (client *Client) CreateRenderingProjectWithOptions(tmpReq *CreateRenderingProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateRenderingProjectResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 创建一个新的云应用服务项目，并设置相关属性。
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

// @param request - CreateStreamSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamSnapshotResponse
func (client *Client) CreateStreamSnapshotWithOptions(request *CreateStreamSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamSnapshotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除云应用
//
// @param request - DeleteCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudAppResponse
func (client *Client) DeleteCloudAppWithOptions(request *DeleteCloudAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudAppResponse, _err error) {
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
// 删除云应用
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

// @param request - DeleteDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeviceResponse
func (client *Client) DeleteDeviceWithOptions(request *DeleteDeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除文件对象。
//
// @param request - DeleteFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除文件对象。
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

// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DeleteParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParentPlatformResponse
func (client *Client) DeleteParentPlatformWithOptions(request *DeleteParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *DeleteParentPlatformResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DeletePresetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePresetResponse
func (client *Client) DeletePresetWithOptions(request *DeletePresetRequest, runtime *dara.RuntimeOptions) (_result *DeletePresetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除公钥信息
//
// @param request - DeletePublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePublicKeyResponse
func (client *Client) DeletePublicKeyWithOptions(request *DeletePublicKeyRequest, runtime *dara.RuntimeOptions) (_result *DeletePublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除公钥信息
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
// 删除云渲染实例配置参数
//
// @param tmpReq - DeleteRenderingInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingInstanceConfigurationResponse
func (client *Client) DeleteRenderingInstanceConfigurationWithOptions(tmpReq *DeleteRenderingInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingInstanceConfigurationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 删除云渲染实例配置参数
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
// 删除自定义网关
//
// @param request - DeleteRenderingInstanceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingInstanceGatewayResponse
func (client *Client) DeleteRenderingInstanceGatewayWithOptions(request *DeleteRenderingInstanceGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingInstanceGatewayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除自定义网关
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
// 清除实例设置
//
// @param tmpReq - DeleteRenderingInstanceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingInstanceSettingsResponse
func (client *Client) DeleteRenderingInstanceSettingsWithOptions(tmpReq *DeleteRenderingInstanceSettingsRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingInstanceSettingsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 清除实例设置
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
// 删除一个云应用服务项目，有在线会话等业务调度数据的项目不允许删除。
//
// @param request - DeleteRenderingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRenderingProjectResponse
func (client *Client) DeleteRenderingProjectWithOptions(request *DeleteRenderingProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteRenderingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除一个云应用服务项目，有在线会话等业务调度数据的项目不允许删除。
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

// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DeleteVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVsPullStreamInfoConfigResponse
func (client *Client) DeleteVsPullStreamInfoConfigWithOptions(request *DeleteVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteVsPullStreamInfoConfigResponse, _err error) {
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

// @param request - DeleteVsStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVsStreamsNotifyUrlConfigResponse
func (client *Client) DeleteVsStreamsNotifyUrlConfigWithOptions(request *DeleteVsStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteVsStreamsNotifyUrlConfigResponse, _err error) {
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

// @param request - DescribeAccountStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountStatResponse
func (client *Client) DescribeAccountStatWithOptions(request *DescribeAccountStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountStatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceResponse
func (client *Client) DescribeDeviceWithOptions(request *DescribeDeviceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDeviceChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceChannelsResponse
func (client *Client) DescribeDeviceChannelsWithOptions(request *DescribeDeviceChannelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceChannelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDeviceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceGatewayResponse
func (client *Client) DescribeDeviceGatewayWithOptions(request *DescribeDeviceGatewayRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceGatewayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDeviceURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceURLResponse
func (client *Client) DescribeDeviceURLWithOptions(request *DescribeDeviceURLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceURLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDevicesResponse
func (client *Client) DescribeDevicesWithOptions(request *DescribeDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoryResponse
func (client *Client) DescribeDirectoryWithOptions(request *DescribeDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupResponse
func (client *Client) DescribeGroupWithOptions(request *DescribeGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroupsWithOptions(request *DescribeGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentPlatformResponse
func (client *Client) DescribeParentPlatformWithOptions(request *DescribeParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentPlatformResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeParentPlatformDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentPlatformDevicesResponse
func (client *Client) DescribeParentPlatformDevicesWithOptions(request *DescribeParentPlatformDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentPlatformDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeParentPlatformsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentPlatformsResponse
func (client *Client) DescribeParentPlatformsWithOptions(request *DescribeParentPlatformsRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentPlatformsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribePresetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePresetsResponse
func (client *Client) DescribePresetsWithOptions(request *DescribePresetsRequest, runtime *dara.RuntimeOptions) (_result *DescribePresetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribePurchasedDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedDeviceResponse
func (client *Client) DescribePurchasedDeviceWithOptions(request *DescribePurchasedDeviceRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribePurchasedDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedDevicesResponse
func (client *Client) DescribePurchasedDevicesWithOptions(request *DescribePurchasedDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordsResponse
func (client *Client) DescribeRecordsWithOptions(request *DescribeRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordsResponse, _err error) {
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
// 查询云渲染实例详细信息。
//
// @param request - DescribeRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingInstanceResponse
func (client *Client) DescribeRenderingInstanceWithOptions(request *DescribeRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询云渲染实例详细信息。
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
// 查询云渲染实例模块配置参数
//
// @param tmpReq - DescribeRenderingInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingInstanceConfigurationResponse
func (client *Client) DescribeRenderingInstanceConfigurationWithOptions(tmpReq *DescribeRenderingInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingInstanceConfigurationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询云渲染实例模块配置参数
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
// 查询实例配置
//
// @param tmpReq - DescribeRenderingInstanceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingInstanceSettingsResponse
func (client *Client) DescribeRenderingInstanceSettingsWithOptions(tmpReq *DescribeRenderingInstanceSettingsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingInstanceSettingsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 查询实例配置
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
// 输出会话的详情信息，包含关联的实例、网络出口等信息。
//
// @param request - DescribeRenderingSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenderingSessionResponse
func (client *Client) DescribeRenderingSessionWithOptions(request *DescribeRenderingSessionRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenderingSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 输出会话的详情信息，包含关联的实例、网络出口等信息。
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

// @param request - DescribeStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamResponse
func (client *Client) DescribeStreamWithOptions(request *DescribeStreamRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeStreamURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamURLResponse
func (client *Client) DescribeStreamURLWithOptions(request *DescribeStreamURLRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamURLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeStreamVodListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamVodListResponse
func (client *Client) DescribeStreamVodListWithOptions(request *DescribeStreamVodListRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamVodListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeStreamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamsResponse
func (client *Client) DescribeStreamsWithOptions(request *DescribeStreamsRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateResponse
func (client *Client) DescribeTemplateWithOptions(request *DescribeTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeVodStreamURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVodStreamURLResponse
func (client *Client) DescribeVodStreamURLWithOptions(request *DescribeVodStreamURLRequest, runtime *dara.RuntimeOptions) (_result *DescribeVodStreamURLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeVsCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsCertificateDetailResponse
func (client *Client) DescribeVsCertificateDetailWithOptions(request *DescribeVsCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeVsCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsCertificateListResponse
func (client *Client) DescribeVsCertificateListWithOptions(request *DescribeVsCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsCertificateListResponse, _err error) {
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

// @param request - DescribeVsDevicesDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDevicesDataResponse
func (client *Client) DescribeVsDevicesDataWithOptions(request *DescribeVsDevicesDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDevicesDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeVsDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainBpsDataResponse
func (client *Client) DescribeVsDomainBpsDataWithOptions(request *DescribeVsDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainBpsDataResponse, _err error) {
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

// @param request - DescribeVsDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainCertificateInfoResponse
func (client *Client) DescribeVsDomainCertificateInfoWithOptions(request *DescribeVsDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainCertificateInfoResponse, _err error) {
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

// @param request - DescribeVsDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainConfigsResponse
func (client *Client) DescribeVsDomainConfigsWithOptions(request *DescribeVsDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainConfigsResponse, _err error) {
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

// @param request - DescribeVsDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainDetailResponse
func (client *Client) DescribeVsDomainDetailWithOptions(request *DescribeVsDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainDetailResponse, _err error) {
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

// @param request - DescribeVsDomainPvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainPvDataResponse
func (client *Client) DescribeVsDomainPvDataWithOptions(request *DescribeVsDomainPvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainPvDataResponse, _err error) {
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

// @param request - DescribeVsDomainPvUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainPvUvDataResponse
func (client *Client) DescribeVsDomainPvUvDataWithOptions(request *DescribeVsDomainPvUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainPvUvDataResponse, _err error) {
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

// @param request - DescribeVsDomainRecordDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainRecordDataResponse
func (client *Client) DescribeVsDomainRecordDataWithOptions(request *DescribeVsDomainRecordDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainRecordDataResponse, _err error) {
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

// @param request - DescribeVsDomainRegionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainRegionDataResponse
func (client *Client) DescribeVsDomainRegionDataWithOptions(request *DescribeVsDomainRegionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainRegionDataResponse, _err error) {
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

// @param request - DescribeVsDomainReqBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainReqBpsDataResponse
func (client *Client) DescribeVsDomainReqBpsDataWithOptions(request *DescribeVsDomainReqBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainReqBpsDataResponse, _err error) {
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

// @param request - DescribeVsDomainReqTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainReqTrafficDataResponse
func (client *Client) DescribeVsDomainReqTrafficDataWithOptions(request *DescribeVsDomainReqTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainReqTrafficDataResponse, _err error) {
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

// @param request - DescribeVsDomainSnapshotDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainSnapshotDataResponse
func (client *Client) DescribeVsDomainSnapshotDataWithOptions(request *DescribeVsDomainSnapshotDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainSnapshotDataResponse, _err error) {
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

// @param request - DescribeVsDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainTrafficDataResponse
func (client *Client) DescribeVsDomainTrafficDataWithOptions(request *DescribeVsDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainTrafficDataResponse, _err error) {
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

// @param request - DescribeVsDomainUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsDomainUvDataResponse
func (client *Client) DescribeVsDomainUvDataWithOptions(request *DescribeVsDomainUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsDomainUvDataResponse, _err error) {
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

// @param request - DescribeVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsPullStreamInfoConfigResponse
func (client *Client) DescribeVsPullStreamInfoConfigWithOptions(request *DescribeVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsPullStreamInfoConfigResponse, _err error) {
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

// @param request - DescribeVsStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsStreamsNotifyUrlConfigResponse
func (client *Client) DescribeVsStreamsNotifyUrlConfigWithOptions(request *DescribeVsStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsStreamsNotifyUrlConfigResponse, _err error) {
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

// @param request - DescribeVsStreamsOnlineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsStreamsOnlineListResponse
func (client *Client) DescribeVsStreamsOnlineListWithOptions(request *DescribeVsStreamsOnlineListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsStreamsOnlineListResponse, _err error) {
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

// @param request - DescribeVsStreamsPublishListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsStreamsPublishListResponse
func (client *Client) DescribeVsStreamsPublishListWithOptions(request *DescribeVsStreamsPublishListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsStreamsPublishListResponse, _err error) {
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

// @param request - DescribeVsTopDomainsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsTopDomainsByFlowResponse
func (client *Client) DescribeVsTopDomainsByFlowWithOptions(request *DescribeVsTopDomainsByFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsTopDomainsByFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeVsUpPeakPublishStreamDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsUpPeakPublishStreamDataResponse
func (client *Client) DescribeVsUpPeakPublishStreamDataWithOptions(request *DescribeVsUpPeakPublishStreamDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsUpPeakPublishStreamDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DescribeVsVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVsVerifyContentResponse
func (client *Client) DescribeVsVerifyContentWithOptions(request *DescribeVsVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeVsVerifyContentResponse, _err error) {
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
// 云应用服务实例与项目解除关联
//
// @param tmpReq - DisassociateRenderingProjectInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateRenderingProjectInstancesResponse
func (client *Client) DisassociateRenderingProjectInstancesWithOptions(tmpReq *DisassociateRenderingProjectInstancesRequest, runtime *dara.RuntimeOptions) (_result *DisassociateRenderingProjectInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 云应用服务实例与项目解除关联
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

// @param request - ForbidVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForbidVsStreamResponse
func (client *Client) ForbidVsStreamWithOptions(request *ForbidVsStreamRequest, runtime *dara.RuntimeOptions) (_result *ForbidVsStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询命令的执行状态与结果。
//
// @param request - GetRenderingInstanceCommandsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenderingInstanceCommandsStatusResponse
func (client *Client) GetRenderingInstanceCommandsStatusWithOptions(request *GetRenderingInstanceCommandsStatusRequest, runtime *dara.RuntimeOptions) (_result *GetRenderingInstanceCommandsStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询命令的执行状态与结果。
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
// 获取云渲染实例流连接信息，每次流化建联前都需要调用此接口获取最新连接信息
//
// @param request - GetRenderingInstanceStreamingInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenderingInstanceStreamingInfoResponse
func (client *Client) GetRenderingInstanceStreamingInfoWithOptions(request *GetRenderingInstanceStreamingInfoRequest, runtime *dara.RuntimeOptions) (_result *GetRenderingInstanceStreamingInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 获取云渲染实例流连接信息，每次流化建联前都需要调用此接口获取最新连接信息
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
// 输出满足特定条件的资源各状态数据量统计值。
//
// @param request - GetRenderingProjectInstanceStateMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenderingProjectInstanceStateMetricsResponse
func (client *Client) GetRenderingProjectInstanceStateMetricsWithOptions(request *GetRenderingProjectInstanceStateMetricsRequest, runtime *dara.RuntimeOptions) (_result *GetRenderingProjectInstanceStateMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 输出满足特定条件的资源各状态数据量统计值。
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

// @param request - GotoPresetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GotoPresetResponse
func (client *Client) GotoPresetWithOptions(request *GotoPresetRequest, runtime *dara.RuntimeOptions) (_result *GotoPresetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 安装云应用
//
// @param tmpReq - InstallCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCloudAppResponse
func (client *Client) InstallCloudAppWithOptions(tmpReq *InstallCloudAppRequest, runtime *dara.RuntimeOptions) (_result *InstallCloudAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 安装云应用
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
// 查询云应用安装信息列表
//
// @param request - ListCloudAppInstallationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAppInstallationsResponse
func (client *Client) ListCloudAppInstallationsWithOptions(request *ListCloudAppInstallationsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAppInstallationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询云应用安装信息列表
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
// 查询一个云应用的Patch列表。
//
// @param request - ListCloudAppPatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAppPatchesResponse
func (client *Client) ListCloudAppPatchesWithOptions(request *ListCloudAppPatchesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAppPatchesResponse, _err error) {
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
// 查询一个云应用的Patch列表。
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
// 查询云应用列表
//
// @param request - ListCloudAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAppsResponse
func (client *Client) ListCloudAppsWithOptions(request *ListCloudAppsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询云应用列表
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
// 查询文件的实例推送状态信息列表。
//
// @param request - ListFilePushStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilePushStatusesResponse
func (client *Client) ListFilePushStatusesWithOptions(request *ListFilePushStatusesRequest, runtime *dara.RuntimeOptions) (_result *ListFilePushStatusesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询文件的实例推送状态信息列表。
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
// 查询可用文件列表。
//
// @param request - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithOptions(request *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询可用文件列表。
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
// 查询公钥信息
//
// @param request - ListPublicKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicKeysResponse
func (client *Client) ListPublicKeysWithOptions(request *ListPublicKeysRequest, runtime *dara.RuntimeOptions) (_result *ListPublicKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询公钥信息
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
// 查询所有云应用数据包信息，支持分页查询。
//
// @param request - ListRenderingDataPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingDataPackagesResponse
func (client *Client) ListRenderingDataPackagesWithOptions(request *ListRenderingDataPackagesRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingDataPackagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询所有云应用数据包信息，支持分页查询。
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
// 查询自定义网关
//
// @param request - ListRenderingInstanceGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingInstanceGatewayResponse
func (client *Client) ListRenderingInstanceGatewayWithOptions(request *ListRenderingInstanceGatewayRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingInstanceGatewayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询自定义网关
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
// 查询所有云渲染实例信息，支持分页查询。
//
// @param request - ListRenderingInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingInstancesResponse
func (client *Client) ListRenderingInstancesWithOptions(request *ListRenderingInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询所有云渲染实例信息，支持分页查询。
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
// 分页查询项目关联的云应用服务实例列表。
//
// Description:
//
// ## 请求说明
//
// - 该接口支持通过多种筛选条件（如状态、实例ID等）来查询指定项目下的云应用服务实例。
//
// @param request - ListRenderingProjectInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingProjectInstancesResponse
func (client *Client) ListRenderingProjectInstancesWithOptions(request *ListRenderingProjectInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingProjectInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询项目关联的云应用服务实例列表。
//
// Description:
//
// ## 请求说明
//
// - 该接口支持通过多种筛选条件（如状态、实例ID等）来查询指定项目下的云应用服务实例。
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
// 分页查询用户下的云应用服务项目基本信息列表。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于分页查询指定用户下的渲染项目基本信息列表。
//
// - 可通过 `ProjectId` 和 `ProjectName` 进行过滤查询。
//
// @param request - ListRenderingProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingProjectsResponse
func (client *Client) ListRenderingProjectsWithOptions(request *ListRenderingProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询用户下的云应用服务项目基本信息列表。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于分页查询指定用户下的渲染项目基本信息列表。
//
// - 可通过 `ProjectId` 和 `ProjectName` 进行过滤查询。
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
// 分页查询指定条件下的渲染会话列表。
//
// Description:
//
// ## 请求说明
//
// - 该接口支持通过多种参数组合来过滤和分页查询用户的渲染会话列表。
//
// - `SessionId` 和 `ClientId` 参数至少需要提供一个，但两者都不是必选的。如果同时提供了两个参数，则将根据这两个参数进行更精确的匹配。
//
// @param request - ListRenderingSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRenderingSessionsResponse
func (client *Client) ListRenderingSessionsWithOptions(request *ListRenderingSessionsRequest, runtime *dara.RuntimeOptions) (_result *ListRenderingSessionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 分页查询指定条件下的渲染会话列表。
//
// Description:
//
// ## 请求说明
//
// - 该接口支持通过多种参数组合来过滤和分页查询用户的渲染会话列表。
//
// - `SessionId` 和 `ClientId` 参数至少需要提供一个，但两者都不是必选的。如果同时提供了两个参数，则将根据这两个参数进行更精确的匹配。
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
// 安全登陆管理
//
// @param request - ManageLoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageLoginResponse
func (client *Client) ManageLoginWithOptions(request *ManageLoginRequest, runtime *dara.RuntimeOptions) (_result *ManageLoginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 安全登陆管理
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

// @param request - ModifyDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceResponse
func (client *Client) ModifyDeviceWithOptions(request *ModifyDeviceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ModifyDeviceAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceAlarmResponse
func (client *Client) ModifyDeviceAlarmWithOptions(request *ModifyDeviceAlarmRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ModifyDeviceCaptureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceCaptureResponse
func (client *Client) ModifyDeviceCaptureWithOptions(request *ModifyDeviceCaptureRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceCaptureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ModifyDeviceChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceChannelsResponse
func (client *Client) ModifyDeviceChannelsWithOptions(request *ModifyDeviceChannelsRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceChannelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ModifyDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDirectoryResponse
func (client *Client) ModifyDirectoryWithOptions(request *ModifyDirectoryRequest, runtime *dara.RuntimeOptions) (_result *ModifyDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ModifyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGroupResponse
func (client *Client) ModifyGroupWithOptions(request *ModifyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ModifyParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParentPlatformResponse
func (client *Client) ModifyParentPlatformWithOptions(request *ModifyParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *ModifyParentPlatformResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 变配云渲染资源实例
//
// @param request - ModifyRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRenderingInstanceResponse
func (client *Client) ModifyRenderingInstanceWithOptions(request *ModifyRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyRenderingInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 变配云渲染资源实例
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
// 修改云应用服务实例密码
//
// @param request - ModifyRenderingInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRenderingInstanceAttributeResponse
func (client *Client) ModifyRenderingInstanceAttributeWithOptions(request *ModifyRenderingInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRenderingInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改云应用服务实例密码
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
// 修改云渲染实例限速带宽
//
// @param request - ModifyRenderingInstanceBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRenderingInstanceBandwidthResponse
func (client *Client) ModifyRenderingInstanceBandwidthWithOptions(request *ModifyRenderingInstanceBandwidthRequest, runtime *dara.RuntimeOptions) (_result *ModifyRenderingInstanceBandwidthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 修改云渲染实例限速带宽
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

// @param request - ModifyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplateWithOptions(request *ModifyTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - OpenVsServiceRequest
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
// 预推文件到云渲染实例。
//
// @param request - PushFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushFileResponse
func (client *Client) PushFileWithOptions(request *PushFileRequest, runtime *dara.RuntimeOptions) (_result *PushFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 预推文件到云渲染实例。
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
// 重启云渲染实例
//
// @param request - RebootRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootRenderingInstanceResponse
func (client *Client) RebootRenderingInstanceWithOptions(request *RebootRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebootRenderingInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 重启云渲染实例
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
// 恢复数据到云渲染实例
//
// @param request - RecoverRenderingDataPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverRenderingDataPackageResponse
func (client *Client) RecoverRenderingDataPackageWithOptions(request *RecoverRenderingDataPackageRequest, runtime *dara.RuntimeOptions) (_result *RecoverRenderingDataPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 恢复数据到云渲染实例
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
// 更新实例流连接信息
//
// @param tmpReq - RefreshRenderingInstanceStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshRenderingInstanceStreamingResponse
func (client *Client) RefreshRenderingInstanceStreamingWithOptions(tmpReq *RefreshRenderingInstanceStreamingRequest, runtime *dara.RuntimeOptions) (_result *RefreshRenderingInstanceStreamingResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新实例流连接信息
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
// 释放云渲染数据包
//
// @param request - ReleaseRenderingDataPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseRenderingDataPackageResponse
func (client *Client) ReleaseRenderingDataPackageWithOptions(request *ReleaseRenderingDataPackageRequest, runtime *dara.RuntimeOptions) (_result *ReleaseRenderingDataPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 释放云渲染数据包
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
// 释放云渲染实例
//
// @param request - ReleaseRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseRenderingInstanceResponse
func (client *Client) ReleaseRenderingInstanceWithOptions(request *ReleaseRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseRenderingInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 释放云渲染实例
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
// 续费云渲染资源实例
//
// @param request - RenewRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewRenderingInstanceResponse
func (client *Client) RenewRenderingInstanceWithOptions(request *RenewRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewRenderingInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 续费云渲染资源实例
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
// 重置云渲染实例
//
// @param request - ResetRenderingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetRenderingInstanceResponse
func (client *Client) ResetRenderingInstanceWithOptions(request *ResetRenderingInstanceRequest, runtime *dara.RuntimeOptions) (_result *ResetRenderingInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 重置云渲染实例
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

// @param request - ResumeVsStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeVsStreamResponse
func (client *Client) ResumeVsStreamWithOptions(request *ResumeVsStreamRequest, runtime *dara.RuntimeOptions) (_result *ResumeVsStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 下发shell命令，支持同步/异步响应命令。
//
// @param request - SendRenderingInstanceCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRenderingInstanceCommandsResponse
func (client *Client) SendRenderingInstanceCommandsWithOptions(request *SendRenderingInstanceCommandsRequest, runtime *dara.RuntimeOptions) (_result *SendRenderingInstanceCommandsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 下发shell命令，支持同步/异步响应命令。
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

// @param request - SetPresetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPresetResponse
func (client *Client) SetPresetWithOptions(request *SetPresetRequest, runtime *dara.RuntimeOptions) (_result *SetPresetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - SetVsDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVsDomainCertificateResponse
func (client *Client) SetVsDomainCertificateWithOptions(request *SetVsDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetVsDomainCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - SetVsStreamsNotifyUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVsStreamsNotifyUrlConfigResponse
func (client *Client) SetVsStreamsNotifyUrlConfigWithOptions(request *SetVsStreamsNotifyUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *SetVsStreamsNotifyUrlConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StartDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDeviceResponse
func (client *Client) StartDeviceWithOptions(request *StartDeviceRequest, runtime *dara.RuntimeOptions) (_result *StartDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StartParentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartParentPlatformResponse
func (client *Client) StartParentPlatformWithOptions(request *StartParentPlatformRequest, runtime *dara.RuntimeOptions) (_result *StartParentPlatformResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StartRecordStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordStreamResponse
func (client *Client) StartRecordStreamWithOptions(request *StartRecordStreamRequest, runtime *dara.RuntimeOptions) (_result *StartRecordStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 调度一个空闲云应用服务实例，并完成服务启动。
//
// @param tmpReq - StartRenderingSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRenderingSessionResponse
func (client *Client) StartRenderingSessionWithOptions(tmpReq *StartRenderingSessionRequest, runtime *dara.RuntimeOptions) (_result *StartRenderingSessionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 调度一个空闲云应用服务实例，并完成服务启动。
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

// @param request - StartStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartStreamResponse
func (client *Client) StartStreamWithOptions(request *StartStreamRequest, runtime *dara.RuntimeOptions) (_result *StartStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StartTransferStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTransferStreamResponse
func (client *Client) StartTransferStreamWithOptions(request *StartTransferStreamRequest, runtime *dara.RuntimeOptions) (_result *StartTransferStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StopAdjustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAdjustResponse
func (client *Client) StopAdjustWithOptions(request *StopAdjustRequest, runtime *dara.RuntimeOptions) (_result *StopAdjustResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StopDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDeviceResponse
func (client *Client) StopDeviceWithOptions(request *StopDeviceRequest, runtime *dara.RuntimeOptions) (_result *StopDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StopMoveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMoveResponse
func (client *Client) StopMoveWithOptions(request *StopMoveRequest, runtime *dara.RuntimeOptions) (_result *StopMoveResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StopRecordStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordStreamResponse
func (client *Client) StopRecordStreamWithOptions(request *StopRecordStreamRequest, runtime *dara.RuntimeOptions) (_result *StopRecordStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 关闭指定的云应用服务会话并回收相关实例资源。
//
// Description:
//
// ## 请求说明
//
// @param request - StopRenderingSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRenderingSessionResponse
func (client *Client) StopRenderingSessionWithOptions(request *StopRenderingSessionRequest, runtime *dara.RuntimeOptions) (_result *StopRenderingSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 关闭指定的云应用服务会话并回收相关实例资源。
//
// Description:
//
// ## 请求说明
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

// @param request - StopStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStreamResponse
func (client *Client) StopStreamWithOptions(request *StopStreamRequest, runtime *dara.RuntimeOptions) (_result *StopStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - StopTransferStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTransferStreamResponse
func (client *Client) StopTransferStreamWithOptions(request *StopTransferStreamRequest, runtime *dara.RuntimeOptions) (_result *StopTransferStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - SyncCatalogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncCatalogsResponse
func (client *Client) SyncCatalogsWithOptions(request *SyncCatalogsRequest, runtime *dara.RuntimeOptions) (_result *SyncCatalogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - UnbindDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDirectoryResponse
func (client *Client) UnbindDirectoryWithOptions(request *UnbindDirectoryRequest, runtime *dara.RuntimeOptions) (_result *UnbindDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - UnbindParentPlatformDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindParentPlatformDeviceResponse
func (client *Client) UnbindParentPlatformDeviceWithOptions(request *UnbindParentPlatformDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnbindParentPlatformDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - UnbindPurchasedDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPurchasedDeviceResponse
func (client *Client) UnbindPurchasedDeviceWithOptions(request *UnbindPurchasedDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnbindPurchasedDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - UnbindTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindTemplateResponse
func (client *Client) UnbindTemplateWithOptions(request *UnbindTemplateRequest, runtime *dara.RuntimeOptions) (_result *UnbindTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 卸载云应用
//
// @param tmpReq - UninstallCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallCloudAppResponse
func (client *Client) UninstallCloudAppWithOptions(tmpReq *UninstallCloudAppRequest, runtime *dara.RuntimeOptions) (_result *UninstallCloudAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 卸载云应用
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

// @param request - UnlockDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockDeviceResponse
func (client *Client) UnlockDeviceWithOptions(request *UnlockDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnlockDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新云应用信息
//
// @param tmpReq - UpdateCloudAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudAppInfoResponse
func (client *Client) UpdateCloudAppInfoWithOptions(tmpReq *UpdateCloudAppInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudAppInfoResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateCloudAppInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Patch) {
		request.PatchShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Patch, dara.String("Patch"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.PatchShrink) {
		query["Patch"] = request.PatchShrink
	}

	if !dara.IsNil(request.StablePatchId) {
		query["StablePatchId"] = request.StablePatchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
// 更新云应用信息
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
// 更新文件信息。
//
// @param request - UpdateFileInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileInfoResponse
func (client *Client) UpdateFileInfoWithOptions(request *UpdateFileInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 更新文件信息。
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
// 更新云渲染实例配置参数
//
// @param tmpReq - UpdateRenderingInstanceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRenderingInstanceConfigurationResponse
func (client *Client) UpdateRenderingInstanceConfigurationWithOptions(tmpReq *UpdateRenderingInstanceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateRenderingInstanceConfigurationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新云渲染实例配置参数
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
// 更新实例设置
//
// @param tmpReq - UpdateRenderingInstanceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRenderingInstanceSettingsResponse
func (client *Client) UpdateRenderingInstanceSettingsWithOptions(tmpReq *UpdateRenderingInstanceSettingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateRenderingInstanceSettingsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新实例设置
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
// 更新一个项目的属性信息
//
// @param tmpReq - UpdateRenderingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRenderingProjectResponse
func (client *Client) UpdateRenderingProjectWithOptions(tmpReq *UpdateRenderingProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateRenderingProjectResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 更新一个项目的属性信息
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

// @param request - UpdateVsPullStreamInfoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVsPullStreamInfoConfigResponse
func (client *Client) UpdateVsPullStreamInfoConfigWithOptions(request *UpdateVsPullStreamInfoConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateVsPullStreamInfoConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 应用上架
//
// @param request - UploadCloudAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCloudAppResponse
func (client *Client) UploadCloudAppWithOptions(request *UploadCloudAppRequest, runtime *dara.RuntimeOptions) (_result *UploadCloudAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 应用上架
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
// 文件上传
//
// @param request - UploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadFileResponse
func (client *Client) UploadFileWithOptions(request *UploadFileRequest, runtime *dara.RuntimeOptions) (_result *UploadFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 文件上传
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
// 上传公钥，用于安全登陆鉴权。
//
// @param request - UploadPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadPublicKeyResponse
func (client *Client) UploadPublicKeyWithOptions(request *UploadPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *UploadPublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 上传公钥，用于安全登陆鉴权。
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

// @param request - VerifyVsDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyVsDomainOwnerResponse
func (client *Client) VerifyVsDomainOwnerWithOptions(request *VerifyVsDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyVsDomainOwnerResponse, _err error) {
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
