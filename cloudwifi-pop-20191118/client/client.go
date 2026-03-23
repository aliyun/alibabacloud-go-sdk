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
		"cn-hangzhou": dara.String("cloudwf.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloudwifi-pop"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param tmpReq - AddApListToApgroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddApListToApgroupResponse
func (client *Client) AddApListToApgroupWithOptions(tmpReq *AddApListToApgroupRequest, runtime *dara.RuntimeOptions) (_result *AddApListToApgroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddApListToApgroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApMacList) {
		request.ApMacListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApMacList, dara.String("ApMacList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApGroupId) {
		query["ApGroupId"] = request.ApGroupId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApMacListShrink) {
		body["ApMacList"] = request.ApMacListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddApListToApgroup"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddApListToApgroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddApListToApgroupRequest
//
// @return AddApListToApgroupResponse
func (client *Client) AddApListToApgroup(request *AddApListToApgroupRequest) (_result *AddApListToApgroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddApListToApgroupResponse{}
	_body, _err := client.AddApListToApgroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除设备的三方app
//
// @param request - DelApThirdAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelApThirdAppResponse
func (client *Client) DelApThirdAppWithOptions(request *DelApThirdAppRequest, runtime *dara.RuntimeOptions) (_result *DelApThirdAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApAssetId) {
		query["ApAssetId"] = request.ApAssetId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelApThirdApp"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelApThirdAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除设备的三方app
//
// @param request - DelApThirdAppRequest
//
// @return DelApThirdAppResponse
func (client *Client) DelApThirdApp(request *DelApThirdAppRequest) (_result *DelApThirdAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelApThirdAppResponse{}
	_body, _err := client.DelApThirdAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteApSsidConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApSsidConfigResponse
func (client *Client) DeleteApSsidConfigWithOptions(request *DeleteApSsidConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteApSsidConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.RadioIndex) {
		query["RadioIndex"] = request.RadioIndex
	}

	if !dara.IsNil(request.Ssid) {
		query["Ssid"] = request.Ssid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApSsidConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApSsidConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteApSsidConfigRequest
//
// @return DeleteApSsidConfigResponse
func (client *Client) DeleteApSsidConfig(request *DeleteApSsidConfigRequest) (_result *DeleteApSsidConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApSsidConfigResponse{}
	_body, _err := client.DeleteApSsidConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteApgroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApgroupConfigResponse
func (client *Client) DeleteApgroupConfigWithOptions(request *DeleteApgroupConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteApgroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApGroupUUId) {
		query["ApGroupUUId"] = request.ApGroupUUId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApgroupConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApgroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteApgroupConfigRequest
//
// @return DeleteApgroupConfigResponse
func (client *Client) DeleteApgroupConfig(request *DeleteApgroupConfigRequest) (_result *DeleteApgroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApgroupConfigResponse{}
	_body, _err := client.DeleteApgroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AP组上的SSID
//
// @param request - DeleteApgroupSsidConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApgroupSsidConfigResponse
func (client *Client) DeleteApgroupSsidConfigWithOptions(request *DeleteApgroupSsidConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteApgroupSsidConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApgroupSsidConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApgroupSsidConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AP组上的SSID
//
// @param request - DeleteApgroupSsidConfigRequest
//
// @return DeleteApgroupSsidConfigResponse
func (client *Client) DeleteApgroupSsidConfig(request *DeleteApgroupSsidConfigRequest) (_result *DeleteApgroupSsidConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApgroupSsidConfigResponse{}
	_body, _err := client.DeleteApgroupSsidConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除设备组的三方app
//
// @param request - DeleteApgroupThirdAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApgroupThirdAppResponse
func (client *Client) DeleteApgroupThirdAppWithOptions(request *DeleteApgroupThirdAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteApgroupThirdAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApgroupThirdApp"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApgroupThirdAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除设备组的三方app
//
// @param request - DeleteApgroupThirdAppRequest
//
// @return DeleteApgroupThirdAppResponse
func (client *Client) DeleteApgroupThirdApp(request *DeleteApgroupThirdAppRequest) (_result *DeleteApgroupThirdAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApgroupThirdAppResponse{}
	_body, _err := client.DeleteApgroupThirdAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteNetDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetDeviceInfoResponse
func (client *Client) DeleteNetDeviceInfoWithOptions(request *DeleteNetDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetDeviceInfo"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteNetDeviceInfoRequest
//
// @return DeleteNetDeviceInfoResponse
func (client *Client) DeleteNetDeviceInfo(request *DeleteNetDeviceInfoRequest) (_result *DeleteNetDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetDeviceInfoResponse{}
	_body, _err := client.DeleteNetDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置设备组的三方app
//
// @param request - EditApgroupThirdAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditApgroupThirdAppResponse
func (client *Client) EditApgroupThirdAppWithOptions(request *EditApgroupThirdAppRequest, runtime *dara.RuntimeOptions) (_result *EditApgroupThirdAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppData) {
		query["AppData"] = request.AppData
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ApplyToSubGroup) {
		query["ApplyToSubGroup"] = request.ApplyToSubGroup
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InheritParentGroup) {
		query["InheritParentGroup"] = request.InheritParentGroup
	}

	if !dara.IsNil(request.ThirdAppName) {
		query["ThirdAppName"] = request.ThirdAppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditApgroupThirdApp"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditApgroupThirdAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置设备组的三方app
//
// @param request - EditApgroupThirdAppRequest
//
// @return EditApgroupThirdAppResponse
func (client *Client) EditApgroupThirdApp(request *EditApgroupThirdAppRequest) (_result *EditApgroupThirdAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditApgroupThirdAppResponse{}
	_body, _err := client.EditApgroupThirdAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EffectApConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EffectApConfigResponse
func (client *Client) EffectApConfigWithOptions(request *EffectApConfigRequest, runtime *dara.RuntimeOptions) (_result *EffectApConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EffectApConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EffectApConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EffectApConfigRequest
//
// @return EffectApConfigResponse
func (client *Client) EffectApConfig(request *EffectApConfigRequest) (_result *EffectApConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EffectApConfigResponse{}
	_body, _err := client.EffectApConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EffectApgroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EffectApgroupConfigResponse
func (client *Client) EffectApgroupConfigWithOptions(request *EffectApgroupConfigRequest, runtime *dara.RuntimeOptions) (_result *EffectApgroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApGroupUUId) {
		query["ApGroupUUId"] = request.ApGroupUUId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EffectApgroupConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EffectApgroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EffectApgroupConfigRequest
//
// @return EffectApgroupConfigResponse
func (client *Client) EffectApgroupConfig(request *EffectApgroupConfigRequest) (_result *EffectApgroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EffectApgroupConfigResponse{}
	_body, _err := client.EffectApgroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询蚂蚁环境终端状态
//
// @param request - GetAntStaStatusByMacRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAntStaStatusByMacResponse
func (client *Client) GetAntStaStatusByMacWithOptions(request *GetAntStaStatusByMacRequest, runtime *dara.RuntimeOptions) (_result *GetAntStaStatusByMacResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.StaMac) {
		query["StaMac"] = request.StaMac
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAntStaStatusByMac"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAntStaStatusByMacResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询蚂蚁环境终端状态
//
// @param request - GetAntStaStatusByMacRequest
//
// @return GetAntStaStatusByMacResponse
func (client *Client) GetAntStaStatusByMac(request *GetAntStaStatusByMacRequest) (_result *GetAntStaStatusByMacResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAntStaStatusByMacResponse{}
	_body, _err := client.GetAntStaStatusByMacWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApAddressByMacRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApAddressByMacResponse
func (client *Client) GetApAddressByMacWithOptions(request *GetApAddressByMacRequest, runtime *dara.RuntimeOptions) (_result *GetApAddressByMacResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApAddressByMac"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApAddressByMacResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApAddressByMacRequest
//
// @return GetApAddressByMacResponse
func (client *Client) GetApAddressByMac(request *GetApAddressByMacRequest) (_result *GetApAddressByMacResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApAddressByMacResponse{}
	_body, _err := client.GetApAddressByMacWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApAssetResponse
func (client *Client) GetApAssetWithOptions(request *GetApAssetRequest, runtime *dara.RuntimeOptions) (_result *GetApAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApAsset"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApAssetRequest
//
// @return GetApAssetResponse
func (client *Client) GetApAsset(request *GetApAssetRequest) (_result *GetApAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApAssetResponse{}
	_body, _err := client.GetApAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApDetailStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApDetailStatusResponse
func (client *Client) GetApDetailStatusWithOptions(request *GetApDetailStatusRequest, runtime *dara.RuntimeOptions) (_result *GetApDetailStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	if !dara.IsNil(request.NeedApgroupInfo) {
		query["NeedApgroupInfo"] = request.NeedApgroupInfo
	}

	if !dara.IsNil(request.NeedRadioStatus) {
		query["NeedRadioStatus"] = request.NeedRadioStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApDetailStatus"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApDetailStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApDetailStatusRequest
//
// @return GetApDetailStatusResponse
func (client *Client) GetApDetailStatus(request *GetApDetailStatusRequest) (_result *GetApDetailStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApDetailStatusResponse{}
	_body, _err := client.GetApDetailStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApDetailedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApDetailedConfigResponse
func (client *Client) GetApDetailedConfigWithOptions(request *GetApDetailedConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApDetailedConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApDetailedConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApDetailedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApDetailedConfigRequest
//
// @return GetApDetailedConfigResponse
func (client *Client) GetApDetailedConfig(request *GetApDetailedConfigRequest) (_result *GetApDetailedConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApDetailedConfigResponse{}
	_body, _err := client.GetApDetailedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApInfoFromPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApInfoFromPoolResponse
func (client *Client) GetApInfoFromPoolWithOptions(request *GetApInfoFromPoolRequest, runtime *dara.RuntimeOptions) (_result *GetApInfoFromPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApInfoFromPool"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApInfoFromPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApInfoFromPoolRequest
//
// @return GetApInfoFromPoolResponse
func (client *Client) GetApInfoFromPool(request *GetApInfoFromPoolRequest) (_result *GetApInfoFromPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApInfoFromPoolResponse{}
	_body, _err := client.GetApInfoFromPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApRunHistoryTimeSerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApRunHistoryTimeSerResponse
func (client *Client) GetApRunHistoryTimeSerWithOptions(request *GetApRunHistoryTimeSerRequest, runtime *dara.RuntimeOptions) (_result *GetApRunHistoryTimeSerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApRunHistoryTimeSer"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApRunHistoryTimeSerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApRunHistoryTimeSerRequest
//
// @return GetApRunHistoryTimeSerResponse
func (client *Client) GetApRunHistoryTimeSer(request *GetApRunHistoryTimeSerRequest) (_result *GetApRunHistoryTimeSerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApRunHistoryTimeSerResponse{}
	_body, _err := client.GetApRunHistoryTimeSerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApStatusByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApStatusByGroupIdResponse
func (client *Client) GetApStatusByGroupIdWithOptions(request *GetApStatusByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetApStatusByGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApStatusByGroupId"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApStatusByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApStatusByGroupIdRequest
//
// @return GetApStatusByGroupIdResponse
func (client *Client) GetApStatusByGroupId(request *GetApStatusByGroupIdRequest) (_result *GetApStatusByGroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApStatusByGroupIdResponse{}
	_body, _err := client.GetApStatusByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据组id查组的基本信息（名称之类）
//
// @param request - GetApgroupConfigByIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApgroupConfigByIdentityResponse
func (client *Client) GetApgroupConfigByIdentityWithOptions(request *GetApgroupConfigByIdentityRequest, runtime *dara.RuntimeOptions) (_result *GetApgroupConfigByIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.ApgroupUuid) {
		query["ApgroupUuid"] = request.ApgroupUuid
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApgroupConfigByIdentity"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApgroupConfigByIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据组id查组的基本信息（名称之类）
//
// @param request - GetApgroupConfigByIdentityRequest
//
// @return GetApgroupConfigByIdentityResponse
func (client *Client) GetApgroupConfigByIdentity(request *GetApgroupConfigByIdentityRequest) (_result *GetApgroupConfigByIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApgroupConfigByIdentityResponse{}
	_body, _err := client.GetApgroupConfigByIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备组的完整配置(含ssid、portal等）
//
// @param request - GetApgroupDetailedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApgroupDetailedConfigResponse
func (client *Client) GetApgroupDetailedConfigWithOptions(request *GetApgroupDetailedConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApgroupDetailedConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.ApgroupUuid) {
		query["ApgroupUuid"] = request.ApgroupUuid
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApgroupDetailedConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApgroupDetailedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备组的完整配置(含ssid、portal等）
//
// @param request - GetApgroupDetailedConfigRequest
//
// @return GetApgroupDetailedConfigResponse
func (client *Client) GetApgroupDetailedConfig(request *GetApgroupDetailedConfigRequest) (_result *GetApgroupDetailedConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApgroupDetailedConfigResponse{}
	_body, _err := client.GetApgroupDetailedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApgroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApgroupIdResponse
func (client *Client) GetApgroupIdWithOptions(request *GetApgroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetApgroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApgroupId"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApgroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApgroupIdRequest
//
// @return GetApgroupIdResponse
func (client *Client) GetApgroupId(request *GetApgroupIdRequest) (_result *GetApgroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApgroupIdResponse{}
	_body, _err := client.GetApgroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetApgroupSsidConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApgroupSsidConfigResponse
func (client *Client) GetApgroupSsidConfigWithOptions(request *GetApgroupSsidConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApgroupSsidConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApGroupUUId) {
		query["ApGroupUUId"] = request.ApGroupUUId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApgroupSsidConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApgroupSsidConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetApgroupSsidConfigRequest
//
// @return GetApgroupSsidConfigResponse
func (client *Client) GetApgroupSsidConfig(request *GetApgroupSsidConfigRequest) (_result *GetApgroupSsidConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApgroupSsidConfigResponse{}
	_body, _err := client.GetApgroupSsidConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询批量任务的执行进度
//
// @param request - GetBatchTaskProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskProgressResponse
func (client *Client) GetBatchTaskProgressWithOptions(request *GetBatchTaskProgressRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskProgress"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询批量任务的执行进度
//
// @param request - GetBatchTaskProgressRequest
//
// @return GetBatchTaskProgressResponse
func (client *Client) GetBatchTaskProgress(request *GetBatchTaskProgressRequest) (_result *GetBatchTaskProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchTaskProgressResponse{}
	_body, _err := client.GetBatchTaskProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetGroupMiscAggTimeSerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupMiscAggTimeSerResponse
func (client *Client) GetGroupMiscAggTimeSerWithOptions(request *GetGroupMiscAggTimeSerRequest, runtime *dara.RuntimeOptions) (_result *GetGroupMiscAggTimeSerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupUuid) {
		query["ApgroupUuid"] = request.ApgroupUuid
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroupMiscAggTimeSer"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupMiscAggTimeSerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetGroupMiscAggTimeSerRequest
//
// @return GetGroupMiscAggTimeSerResponse
func (client *Client) GetGroupMiscAggTimeSer(request *GetGroupMiscAggTimeSerRequest) (_result *GetGroupMiscAggTimeSerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGroupMiscAggTimeSerResponse{}
	_body, _err := client.GetGroupMiscAggTimeSerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNetDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetDeviceInfoResponse
func (client *Client) GetNetDeviceInfoWithOptions(request *GetNetDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *GetNetDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Idc) {
		query["Idc"] = request.Idc
	}

	if !dara.IsNil(request.LogicNetPod) {
		query["LogicNetPod"] = request.LogicNetPod
	}

	if !dara.IsNil(request.Manufacturer) {
		query["Manufacturer"] = request.Manufacturer
	}

	if !dara.IsNil(request.MgnIp) {
		query["MgnIp"] = request.MgnIp
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.NetPod) {
		query["NetPod"] = request.NetPod
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.ServiceTag) {
		query["ServiceTag"] = request.ServiceTag
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetDeviceInfo"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNetDeviceInfoRequest
//
// @return GetNetDeviceInfoResponse
func (client *Client) GetNetDeviceInfo(request *GetNetDeviceInfoRequest) (_result *GetNetDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetDeviceInfoResponse{}
	_body, _err := client.GetNetDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNetDeviceInfoWithSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetDeviceInfoWithSizeResponse
func (client *Client) GetNetDeviceInfoWithSizeWithOptions(request *GetNetDeviceInfoWithSizeRequest, runtime *dara.RuntimeOptions) (_result *GetNetDeviceInfoWithSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Idc) {
		query["Idc"] = request.Idc
	}

	if !dara.IsNil(request.LogicNetPod) {
		query["LogicNetPod"] = request.LogicNetPod
	}

	if !dara.IsNil(request.Manufacturer) {
		query["Manufacturer"] = request.Manufacturer
	}

	if !dara.IsNil(request.MgnIp) {
		query["MgnIp"] = request.MgnIp
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.NetPod) {
		query["NetPod"] = request.NetPod
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.ServiceTag) {
		query["ServiceTag"] = request.ServiceTag
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetDeviceInfoWithSize"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetDeviceInfoWithSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNetDeviceInfoWithSizeRequest
//
// @return GetNetDeviceInfoWithSizeResponse
func (client *Client) GetNetDeviceInfoWithSize(request *GetNetDeviceInfoWithSizeRequest) (_result *GetNetDeviceInfoWithSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetDeviceInfoWithSizeResponse{}
	_body, _err := client.GetNetDeviceInfoWithSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取体验加PV/UV信息
//
// @param request - GetPageVisitDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPageVisitDataResponse
func (client *Client) GetPageVisitDataWithOptions(request *GetPageVisitDataRequest, runtime *dara.RuntimeOptions) (_result *GetPageVisitDataResponse, _err error) {
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
		Action:      dara.String("GetPageVisitData"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPageVisitDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取体验加PV/UV信息
//
// @param request - GetPageVisitDataRequest
//
// @return GetPageVisitDataResponse
func (client *Client) GetPageVisitData(request *GetPageVisitDataRequest) (_result *GetPageVisitDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPageVisitDataResponse{}
	_body, _err := client.GetPageVisitDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRadioRunHistoryTimeSerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRadioRunHistoryTimeSerResponse
func (client *Client) GetRadioRunHistoryTimeSerWithOptions(request *GetRadioRunHistoryTimeSerRequest, runtime *dara.RuntimeOptions) (_result *GetRadioRunHistoryTimeSerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRadioRunHistoryTimeSer"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRadioRunHistoryTimeSerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRadioRunHistoryTimeSerRequest
//
// @return GetRadioRunHistoryTimeSerResponse
func (client *Client) GetRadioRunHistoryTimeSer(request *GetRadioRunHistoryTimeSerRequest) (_result *GetRadioRunHistoryTimeSerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRadioRunHistoryTimeSerResponse{}
	_body, _err := client.GetRadioRunHistoryTimeSerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询终端设备的详细状态
//
// @param request - GetStaDetailedStatusByMacRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStaDetailedStatusByMacResponse
func (client *Client) GetStaDetailedStatusByMacWithOptions(request *GetStaDetailedStatusByMacRequest, runtime *dara.RuntimeOptions) (_result *GetStaDetailedStatusByMacResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.StaMac) {
		query["StaMac"] = request.StaMac
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStaDetailedStatusByMac"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStaDetailedStatusByMacResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询终端设备的详细状态
//
// @param request - GetStaDetailedStatusByMacRequest
//
// @return GetStaDetailedStatusByMacResponse
func (client *Client) GetStaDetailedStatusByMac(request *GetStaDetailedStatusByMacRequest) (_result *GetStaDetailedStatusByMacResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStaDetailedStatusByMacResponse{}
	_body, _err := client.GetStaDetailedStatusByMacWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetStaStatusListByApRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStaStatusListByApResponse
func (client *Client) GetStaStatusListByApWithOptions(request *GetStaStatusListByApRequest, runtime *dara.RuntimeOptions) (_result *GetStaStatusListByApResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStaStatusListByAp"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStaStatusListByApResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetStaStatusListByApRequest
//
// @return GetStaStatusListByApResponse
func (client *Client) GetStaStatusListByAp(request *GetStaStatusListByApRequest) (_result *GetStaStatusListByApResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStaStatusListByApResponse{}
	_body, _err := client.GetStaStatusListByApWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 判断是否刑天的业务
//
// @param request - JudgeXingTianBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JudgeXingTianBusinessResponse
func (client *Client) JudgeXingTianBusinessWithOptions(request *JudgeXingTianBusinessRequest, runtime *dara.RuntimeOptions) (_result *JudgeXingTianBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.RealmId) {
		query["RealmId"] = request.RealmId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JudgeXingTianBusiness"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JudgeXingTianBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 判断是否刑天的业务
//
// @param request - JudgeXingTianBusinessRequest
//
// @return JudgeXingTianBusinessResponse
func (client *Client) JudgeXingTianBusiness(request *JudgeXingTianBusinessRequest) (_result *JudgeXingTianBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &JudgeXingTianBusinessResponse{}
	_body, _err := client.JudgeXingTianBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 踢除蚂蚁环境的终端
//
// @param request - KickAntStaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KickAntStaResponse
func (client *Client) KickAntStaWithOptions(request *KickAntStaRequest, runtime *dara.RuntimeOptions) (_result *KickAntStaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.StaMac) {
		query["StaMac"] = request.StaMac
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KickAntSta"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KickAntStaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 踢除蚂蚁环境的终端
//
// @param request - KickAntStaRequest
//
// @return KickAntStaResponse
func (client *Client) KickAntSta(request *KickAntStaRequest) (_result *KickAntStaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KickAntStaResponse{}
	_body, _err := client.KickAntStaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - KickStaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KickStaResponse
func (client *Client) KickStaWithOptions(request *KickStaRequest, runtime *dara.RuntimeOptions) (_result *KickStaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.StaMac) {
		query["StaMac"] = request.StaMac
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KickSta"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KickStaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - KickStaRequest
//
// @return KickStaResponse
func (client *Client) KickSta(request *KickStaRequest) (_result *KickStaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KickStaResponse{}
	_body, _err := client.KickStaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据组id查组内第一级子组
//
// @param request - ListApgroupDescendantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApgroupDescendantResponse
func (client *Client) ListApgroupDescendantWithOptions(request *ListApgroupDescendantRequest, runtime *dara.RuntimeOptions) (_result *ListApgroupDescendantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.ApgroupUuid) {
		query["ApgroupUuid"] = request.ApgroupUuid
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApgroupDescendant"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApgroupDescendantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据组id查组内第一级子组
//
// @param request - ListApgroupDescendantRequest
//
// @return ListApgroupDescendantResponse
func (client *Client) ListApgroupDescendant(request *ListApgroupDescendantRequest) (_result *ListApgroupDescendantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApgroupDescendantResponse{}
	_body, _err := client.ListApgroupDescendantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListJobOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobOrdersResponse
func (client *Client) ListJobOrdersWithOptions(request *ListJobOrdersRequest, runtime *dara.RuntimeOptions) (_result *ListJobOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ChangingType) {
		query["ChangingType"] = request.ChangingType
	}

	if !dara.IsNil(request.ClientSystem) {
		query["ClientSystem"] = request.ClientSystem
	}

	if !dara.IsNil(request.ClientUniqueId) {
		query["ClientUniqueId"] = request.ClientUniqueId
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Handler) {
		query["Handler"] = request.Handler
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
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
		Action:      dara.String("ListJobOrders"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListJobOrdersRequest
//
// @return ListJobOrdersResponse
func (client *Client) ListJobOrders(request *ListJobOrdersRequest) (_result *ListJobOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobOrdersResponse{}
	_body, _err := client.ListJobOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListJobOrdersWithSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobOrdersWithSizeResponse
func (client *Client) ListJobOrdersWithSizeWithOptions(request *ListJobOrdersWithSizeRequest, runtime *dara.RuntimeOptions) (_result *ListJobOrdersWithSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ChangingType) {
		query["ChangingType"] = request.ChangingType
	}

	if !dara.IsNil(request.ClientSystem) {
		query["ClientSystem"] = request.ClientSystem
	}

	if !dara.IsNil(request.ClientUniqueId) {
		query["ClientUniqueId"] = request.ClientUniqueId
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Handler) {
		query["Handler"] = request.Handler
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
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
		Action:      dara.String("ListJobOrdersWithSize"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobOrdersWithSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListJobOrdersWithSizeRequest
//
// @return ListJobOrdersWithSizeResponse
func (client *Client) ListJobOrdersWithSize(request *ListJobOrdersWithSizeRequest) (_result *ListJobOrdersWithSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobOrdersWithSizeResponse{}
	_body, _err := client.ListJobOrdersWithSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - NewApgroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NewApgroupConfigResponse
func (client *Client) NewApgroupConfigWithOptions(request *NewApgroupConfigRequest, runtime *dara.RuntimeOptions) (_result *NewApgroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentApgroupId) {
		query["ParentApgroupId"] = request.ParentApgroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("NewApgroupConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NewApgroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - NewApgroupConfigRequest
//
// @return NewApgroupConfigResponse
func (client *Client) NewApgroupConfig(request *NewApgroupConfigRequest) (_result *NewApgroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &NewApgroupConfigResponse{}
	_body, _err := client.NewApgroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - NewJobOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NewJobOrderResponse
func (client *Client) NewJobOrderWithOptions(tmpReq *NewJobOrderRequest, runtime *dara.RuntimeOptions) (_result *NewJobOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &NewJobOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChangeType) {
		query["ChangeType"] = request.ChangeType
	}

	if !dara.IsNil(request.ClientSystem) {
		query["ClientSystem"] = request.ClientSystem
	}

	if !dara.IsNil(request.ClientUniqueId) {
		query["ClientUniqueId"] = request.ClientUniqueId
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.ReferUrl) {
		query["ReferUrl"] = request.ReferUrl
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("NewJobOrder"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NewJobOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - NewJobOrderRequest
//
// @return NewJobOrderResponse
func (client *Client) NewJobOrder(request *NewJobOrderRequest) (_result *NewJobOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &NewJobOrderResponse{}
	_body, _err := client.NewJobOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - NewNetDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NewNetDeviceInfoResponse
func (client *Client) NewNetDeviceInfoWithOptions(request *NewNetDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *NewNetDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Devices) {
		body["Devices"] = request.Devices
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("NewNetDeviceInfo"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NewNetDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - NewNetDeviceInfoRequest
//
// @return NewNetDeviceInfoResponse
func (client *Client) NewNetDeviceInfo(request *NewNetDeviceInfoRequest) (_result *NewNetDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &NewNetDeviceInfoResponse{}
	_body, _err := client.NewNetDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PutAppConfigAndSaveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutAppConfigAndSaveResponse
func (client *Client) PutAppConfigAndSaveWithOptions(request *PutAppConfigAndSaveRequest, runtime *dara.RuntimeOptions) (_result *PutAppConfigAndSaveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ConfigureType) {
		query["ConfigureType"] = request.ConfigureType
	}

	if !dara.IsNil(request.CurrentTime) {
		query["CurrentTime"] = request.CurrentTime
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutAppConfigAndSave"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutAppConfigAndSaveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutAppConfigAndSaveRequest
//
// @return PutAppConfigAndSaveResponse
func (client *Client) PutAppConfigAndSave(request *PutAppConfigAndSaveRequest) (_result *PutAppConfigAndSaveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutAppConfigAndSaveResponse{}
	_body, _err := client.PutAppConfigAndSaveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryJobOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryJobOrderResponse
func (client *Client) QueryJobOrderWithOptions(request *QueryJobOrderRequest, runtime *dara.RuntimeOptions) (_result *QueryJobOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryJobOrder"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryJobOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryJobOrderRequest
//
// @return QueryJobOrderResponse
func (client *Client) QueryJobOrder(request *QueryJobOrderRequest) (_result *QueryJobOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryJobOrderResponse{}
	_body, _err := client.QueryJobOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RebootApRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootApResponse
func (client *Client) RebootApWithOptions(request *RebootApRequest, runtime *dara.RuntimeOptions) (_result *RebootApResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootAp"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootApResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RebootApRequest
//
// @return RebootApResponse
func (client *Client) RebootAp(request *RebootApRequest) (_result *RebootApResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebootApResponse{}
	_body, _err := client.RebootApWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RegisterApAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterApAssetResponse
func (client *Client) RegisterApAssetWithOptions(request *RegisterApAssetRequest, runtime *dara.RuntimeOptions) (_result *RegisterApAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApGroupUUId) {
		query["ApGroupUUId"] = request.ApGroupUUId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	if !dara.IsNil(request.SerialNo) {
		query["SerialNo"] = request.SerialNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterApAsset"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterApAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RegisterApAssetRequest
//
// @return RegisterApAssetResponse
func (client *Client) RegisterApAsset(request *RegisterApAssetRequest) (_result *RegisterApAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterApAssetResponse{}
	_body, _err := client.RegisterApAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RepairApRadioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RepairApRadioResponse
func (client *Client) RepairApRadioWithOptions(request *RepairApRadioRequest, runtime *dara.RuntimeOptions) (_result *RepairApRadioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.RadioIndex) {
		query["RadioIndex"] = request.RadioIndex
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RepairApRadio"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RepairApRadioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RepairApRadioRequest
//
// @return RepairApRadioResponse
func (client *Client) RepairApRadio(request *RepairApRadioRequest) (_result *RepairApRadioResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RepairApRadioResponse{}
	_body, _err := client.RepairApRadioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存设备的基本配置
//
// @param request - SaveApBasicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApBasicConfigResponse
func (client *Client) SaveApBasicConfigWithOptions(request *SaveApBasicConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveApBasicConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Dai) {
		query["Dai"] = request.Dai
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EchoInt) {
		query["EchoInt"] = request.EchoInt
	}

	if !dara.IsNil(request.Failover) {
		query["Failover"] = request.Failover
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InsecureAllowed) {
		query["InsecureAllowed"] = request.InsecureAllowed
	}

	if !dara.IsNil(request.LanIp) {
		query["LanIp"] = request.LanIp
	}

	if !dara.IsNil(request.LanMask) {
		query["LanMask"] = request.LanMask
	}

	if !dara.IsNil(request.LanStatus) {
		query["LanStatus"] = request.LanStatus
	}

	if !dara.IsNil(request.LogIp) {
		query["LogIp"] = request.LogIp
	}

	if !dara.IsNil(request.LogLevel) {
		query["LogLevel"] = request.LogLevel
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Passwd) {
		query["Passwd"] = request.Passwd
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Route) {
		query["Route"] = request.Route
	}

	if !dara.IsNil(request.Scan) {
		query["Scan"] = request.Scan
	}

	if !dara.IsNil(request.TokenServer) {
		query["TokenServer"] = request.TokenServer
	}

	if !dara.IsNil(request.Vpn) {
		query["Vpn"] = request.Vpn
	}

	if !dara.IsNil(request.WorkMode) {
		query["WorkMode"] = request.WorkMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApBasicConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApBasicConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存设备的基本配置
//
// @param request - SaveApBasicConfigRequest
//
// @return SaveApBasicConfigResponse
func (client *Client) SaveApBasicConfig(request *SaveApBasicConfigRequest) (_result *SaveApBasicConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApBasicConfigResponse{}
	_body, _err := client.SaveApBasicConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置设备的portal
//
// @param tmpReq - SaveApPortalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApPortalConfigResponse
func (client *Client) SaveApPortalConfigWithOptions(tmpReq *SaveApPortalConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveApPortalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveApPortalConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PortalTypes) {
		request.PortalTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PortalTypes, dara.String("PortalTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApConfigId) {
		query["ApConfigId"] = request.ApConfigId
	}

	if !dara.IsNil(request.AppAuthUrl) {
		query["AppAuthUrl"] = request.AppAuthUrl
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.AuthSecret) {
		query["AuthSecret"] = request.AuthSecret
	}

	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.ClientDownload) {
		query["ClientDownload"] = request.ClientDownload
	}

	if !dara.IsNil(request.ClientUpload) {
		query["ClientUpload"] = request.ClientUpload
	}

	if !dara.IsNil(request.Countdown) {
		query["Countdown"] = request.Countdown
	}

	if !dara.IsNil(request.Network) {
		query["Network"] = request.Network
	}

	if !dara.IsNil(request.PortalTypesShrink) {
		query["PortalTypes"] = request.PortalTypesShrink
	}

	if !dara.IsNil(request.PortalUrl) {
		query["PortalUrl"] = request.PortalUrl
	}

	if !dara.IsNil(request.TimeStamp) {
		query["TimeStamp"] = request.TimeStamp
	}

	if !dara.IsNil(request.TotalDownload) {
		query["TotalDownload"] = request.TotalDownload
	}

	if !dara.IsNil(request.TotalUpload) {
		query["TotalUpload"] = request.TotalUpload
	}

	if !dara.IsNil(request.WebAuthUrl) {
		query["WebAuthUrl"] = request.WebAuthUrl
	}

	if !dara.IsNil(request.Whitelist) {
		query["Whitelist"] = request.Whitelist
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApPortalConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApPortalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置设备的portal
//
// @param request - SaveApPortalConfigRequest
//
// @return SaveApPortalConfigResponse
func (client *Client) SaveApPortalConfig(request *SaveApPortalConfigRequest) (_result *SaveApPortalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApPortalConfigResponse{}
	_body, _err := client.SaveApPortalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveApRadioConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApRadioConfigResponse
func (client *Client) SaveApRadioConfigWithOptions(request *SaveApRadioConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveApRadioConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BcastRate) {
		query["BcastRate"] = request.BcastRate
	}

	if !dara.IsNil(request.BeaconInt) {
		query["BeaconInt"] = request.BeaconInt
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.Frag) {
		query["Frag"] = request.Frag
	}

	if !dara.IsNil(request.Htmode) {
		query["Htmode"] = request.Htmode
	}

	if !dara.IsNil(request.Hwmode) {
		query["Hwmode"] = request.Hwmode
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.McastRate) {
		query["McastRate"] = request.McastRate
	}

	if !dara.IsNil(request.MgmtRate) {
		query["MgmtRate"] = request.MgmtRate
	}

	if !dara.IsNil(request.Minrate) {
		query["Minrate"] = request.Minrate
	}

	if !dara.IsNil(request.Noscan) {
		query["Noscan"] = request.Noscan
	}

	if !dara.IsNil(request.Probereq) {
		query["Probereq"] = request.Probereq
	}

	if !dara.IsNil(request.RequireMode) {
		query["RequireMode"] = request.RequireMode
	}

	if !dara.IsNil(request.Rts) {
		query["Rts"] = request.Rts
	}

	if !dara.IsNil(request.Shortgi) {
		query["Shortgi"] = request.Shortgi
	}

	if !dara.IsNil(request.Txpower) {
		query["Txpower"] = request.Txpower
	}

	if !dara.IsNil(request.Uapsd) {
		query["Uapsd"] = request.Uapsd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApRadioConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApRadioConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveApRadioConfigRequest
//
// @return SaveApRadioConfigResponse
func (client *Client) SaveApRadioConfig(request *SaveApRadioConfigRequest) (_result *SaveApRadioConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApRadioConfigResponse{}
	_body, _err := client.SaveApRadioConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveApSsidConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApSsidConfigResponse
func (client *Client) SaveApSsidConfigWithOptions(request *SaveApSsidConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveApSsidConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcctPort) {
		query["AcctPort"] = request.AcctPort
	}

	if !dara.IsNil(request.AcctSecret) {
		query["AcctSecret"] = request.AcctSecret
	}

	if !dara.IsNil(request.AcctServer) {
		query["AcctServer"] = request.AcctServer
	}

	if !dara.IsNil(request.AcctStatusServerWork) {
		query["AcctStatusServerWork"] = request.AcctStatusServerWork
	}

	if !dara.IsNil(request.ApAssetId) {
		query["ApAssetId"] = request.ApAssetId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ArpProxyEnable) {
		query["ArpProxyEnable"] = request.ArpProxyEnable
	}

	if !dara.IsNil(request.AuthCache) {
		query["AuthCache"] = request.AuthCache
	}

	if !dara.IsNil(request.AuthPort) {
		query["AuthPort"] = request.AuthPort
	}

	if !dara.IsNil(request.AuthSecret) {
		query["AuthSecret"] = request.AuthSecret
	}

	if !dara.IsNil(request.AuthServer) {
		query["AuthServer"] = request.AuthServer
	}

	if !dara.IsNil(request.AuthStatusServerWork) {
		query["AuthStatusServerWork"] = request.AuthStatusServerWork
	}

	if !dara.IsNil(request.Cir) {
		query["Cir"] = request.Cir
	}

	if !dara.IsNil(request.CirStep) {
		query["CirStep"] = request.CirStep
	}

	if !dara.IsNil(request.CirType) {
		query["CirType"] = request.CirType
	}

	if !dara.IsNil(request.CirUl) {
		query["CirUl"] = request.CirUl
	}

	if !dara.IsNil(request.DaeClient) {
		query["DaeClient"] = request.DaeClient
	}

	if !dara.IsNil(request.DaePort) {
		query["DaePort"] = request.DaePort
	}

	if !dara.IsNil(request.DaeSecret) {
		query["DaeSecret"] = request.DaeSecret
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.DisassocLowAck) {
		query["DisassocLowAck"] = request.DisassocLowAck
	}

	if !dara.IsNil(request.DisassocWeakRssi) {
		query["DisassocWeakRssi"] = request.DisassocWeakRssi
	}

	if !dara.IsNil(request.DynamicVlan) {
		query["DynamicVlan"] = request.DynamicVlan
	}

	if !dara.IsNil(request.EncKey) {
		query["EncKey"] = request.EncKey
	}

	if !dara.IsNil(request.Encryption) {
		query["Encryption"] = request.Encryption
	}

	if !dara.IsNil(request.FourthAuthPort) {
		query["FourthAuthPort"] = request.FourthAuthPort
	}

	if !dara.IsNil(request.FourthAuthSecret) {
		query["FourthAuthSecret"] = request.FourthAuthSecret
	}

	if !dara.IsNil(request.FourthAuthServer) {
		query["FourthAuthServer"] = request.FourthAuthServer
	}

	if !dara.IsNil(request.FtOverDs) {
		query["FtOverDs"] = request.FtOverDs
	}

	if !dara.IsNil(request.Hidden) {
		query["Hidden"] = request.Hidden
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Ieee80211r) {
		query["Ieee80211r"] = request.Ieee80211r
	}

	if !dara.IsNil(request.Ieee80211w) {
		query["Ieee80211w"] = request.Ieee80211w
	}

	if !dara.IsNil(request.IgnoreWeakProbe) {
		query["IgnoreWeakProbe"] = request.IgnoreWeakProbe
	}

	if !dara.IsNil(request.Isolate) {
		query["Isolate"] = request.Isolate
	}

	if !dara.IsNil(request.LiteEffect) {
		query["LiteEffect"] = request.LiteEffect
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	if !dara.IsNil(request.MaxInactivity) {
		query["MaxInactivity"] = request.MaxInactivity
	}

	if !dara.IsNil(request.Maxassoc) {
		query["Maxassoc"] = request.Maxassoc
	}

	if !dara.IsNil(request.MobilityDomain) {
		query["MobilityDomain"] = request.MobilityDomain
	}

	if !dara.IsNil(request.MulticastForward) {
		query["MulticastForward"] = request.MulticastForward
	}

	if !dara.IsNil(request.Nasid) {
		query["Nasid"] = request.Nasid
	}

	if !dara.IsNil(request.NdProxyWork) {
		query["NdProxyWork"] = request.NdProxyWork
	}

	if !dara.IsNil(request.Network) {
		query["Network"] = request.Network
	}

	if !dara.IsNil(request.Ownip) {
		query["Ownip"] = request.Ownip
	}

	if !dara.IsNil(request.RadioIndex) {
		query["RadioIndex"] = request.RadioIndex
	}

	if !dara.IsNil(request.SecondaryAcctPort) {
		query["SecondaryAcctPort"] = request.SecondaryAcctPort
	}

	if !dara.IsNil(request.SecondaryAcctSecret) {
		query["SecondaryAcctSecret"] = request.SecondaryAcctSecret
	}

	if !dara.IsNil(request.SecondaryAcctServer) {
		query["SecondaryAcctServer"] = request.SecondaryAcctServer
	}

	if !dara.IsNil(request.SecondaryAuthPort) {
		query["SecondaryAuthPort"] = request.SecondaryAuthPort
	}

	if !dara.IsNil(request.SecondaryAuthSecret) {
		query["SecondaryAuthSecret"] = request.SecondaryAuthSecret
	}

	if !dara.IsNil(request.SecondaryAuthServer) {
		query["SecondaryAuthServer"] = request.SecondaryAuthServer
	}

	if !dara.IsNil(request.SendConfigToAp) {
		query["SendConfigToAp"] = request.SendConfigToAp
	}

	if !dara.IsNil(request.ShortPreamble) {
		query["ShortPreamble"] = request.ShortPreamble
	}

	if !dara.IsNil(request.Ssid) {
		query["Ssid"] = request.Ssid
	}

	if !dara.IsNil(request.SsidLb) {
		query["SsidLb"] = request.SsidLb
	}

	if !dara.IsNil(request.ThirdAuthPort) {
		query["ThirdAuthPort"] = request.ThirdAuthPort
	}

	if !dara.IsNil(request.ThirdAuthSecret) {
		query["ThirdAuthSecret"] = request.ThirdAuthSecret
	}

	if !dara.IsNil(request.ThirdAuthServer) {
		query["ThirdAuthServer"] = request.ThirdAuthServer
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VlanDhcp) {
		query["VlanDhcp"] = request.VlanDhcp
	}

	if !dara.IsNil(request.Wmm) {
		query["Wmm"] = request.Wmm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApSsidConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApSsidConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveApSsidConfigRequest
//
// @return SaveApSsidConfigResponse
func (client *Client) SaveApSsidConfig(request *SaveApSsidConfigRequest) (_result *SaveApSsidConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApSsidConfigResponse{}
	_body, _err := client.SaveApSsidConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置设备的三方app
//
// @param request - SaveApThirdAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApThirdAppResponse
func (client *Client) SaveApThirdAppWithOptions(request *SaveApThirdAppRequest, runtime *dara.RuntimeOptions) (_result *SaveApThirdAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddStyle) {
		query["AddStyle"] = request.AddStyle
	}

	if !dara.IsNil(request.ApAssetId) {
		query["ApAssetId"] = request.ApAssetId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppData) {
		query["AppData"] = request.AppData
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	if !dara.IsNil(request.ThirdAppName) {
		query["ThirdAppName"] = request.ThirdAppName
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApThirdApp"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApThirdAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置设备的三方app
//
// @param request - SaveApThirdAppRequest
//
// @return SaveApThirdAppResponse
func (client *Client) SaveApThirdApp(request *SaveApThirdAppRequest) (_result *SaveApThirdAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApThirdAppResponse{}
	_body, _err := client.SaveApThirdAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存设备组的基本配置
//
// @param request - SaveApgroupBasicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApgroupBasicConfigResponse
func (client *Client) SaveApgroupBasicConfigWithOptions(request *SaveApgroupBasicConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveApgroupBasicConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Dai) {
		query["Dai"] = request.Dai
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EchoInt) {
		query["EchoInt"] = request.EchoInt
	}

	if !dara.IsNil(request.Failover) {
		query["Failover"] = request.Failover
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InsecureAllowed) {
		query["InsecureAllowed"] = request.InsecureAllowed
	}

	if !dara.IsNil(request.IsConfigStrongConsistency) {
		query["IsConfigStrongConsistency"] = request.IsConfigStrongConsistency
	}

	if !dara.IsNil(request.LanIp) {
		query["LanIp"] = request.LanIp
	}

	if !dara.IsNil(request.LanMask) {
		query["LanMask"] = request.LanMask
	}

	if !dara.IsNil(request.LanStatus) {
		query["LanStatus"] = request.LanStatus
	}

	if !dara.IsNil(request.LogIp) {
		query["LogIp"] = request.LogIp
	}

	if !dara.IsNil(request.LogLevel) {
		query["LogLevel"] = request.LogLevel
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentApgroupId) {
		query["ParentApgroupId"] = request.ParentApgroupId
	}

	if !dara.IsNil(request.Passwd) {
		query["Passwd"] = request.Passwd
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Route) {
		query["Route"] = request.Route
	}

	if !dara.IsNil(request.Scan) {
		query["Scan"] = request.Scan
	}

	if !dara.IsNil(request.TokenServer) {
		query["TokenServer"] = request.TokenServer
	}

	if !dara.IsNil(request.Vpn) {
		query["Vpn"] = request.Vpn
	}

	if !dara.IsNil(request.WorkMode) {
		query["WorkMode"] = request.WorkMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApgroupBasicConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApgroupBasicConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存设备组的基本配置
//
// @param request - SaveApgroupBasicConfigRequest
//
// @return SaveApgroupBasicConfigResponse
func (client *Client) SaveApgroupBasicConfig(request *SaveApgroupBasicConfigRequest) (_result *SaveApgroupBasicConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApgroupBasicConfigResponse{}
	_body, _err := client.SaveApgroupBasicConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置设备组的portal
//
// @param tmpReq - SaveApgroupPortalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApgroupPortalConfigResponse
func (client *Client) SaveApgroupPortalConfigWithOptions(tmpReq *SaveApgroupPortalConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveApgroupPortalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveApgroupPortalConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PortalTypes) {
		request.PortalTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PortalTypes, dara.String("PortalTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.AppAuthUrl) {
		query["AppAuthUrl"] = request.AppAuthUrl
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.AuthSecret) {
		query["AuthSecret"] = request.AuthSecret
	}

	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.ClientDownload) {
		query["ClientDownload"] = request.ClientDownload
	}

	if !dara.IsNil(request.ClientUpload) {
		query["ClientUpload"] = request.ClientUpload
	}

	if !dara.IsNil(request.Countdown) {
		query["Countdown"] = request.Countdown
	}

	if !dara.IsNil(request.Network) {
		query["Network"] = request.Network
	}

	if !dara.IsNil(request.PortalTypesShrink) {
		query["PortalTypes"] = request.PortalTypesShrink
	}

	if !dara.IsNil(request.PortalUrl) {
		query["PortalUrl"] = request.PortalUrl
	}

	if !dara.IsNil(request.TimeStamp) {
		query["TimeStamp"] = request.TimeStamp
	}

	if !dara.IsNil(request.TotalDownload) {
		query["TotalDownload"] = request.TotalDownload
	}

	if !dara.IsNil(request.TotalUpload) {
		query["TotalUpload"] = request.TotalUpload
	}

	if !dara.IsNil(request.WebAuthUrl) {
		query["WebAuthUrl"] = request.WebAuthUrl
	}

	if !dara.IsNil(request.Whitelist) {
		query["Whitelist"] = request.Whitelist
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApgroupPortalConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApgroupPortalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置设备组的portal
//
// @param request - SaveApgroupPortalConfigRequest
//
// @return SaveApgroupPortalConfigResponse
func (client *Client) SaveApgroupPortalConfig(request *SaveApgroupPortalConfigRequest) (_result *SaveApgroupPortalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApgroupPortalConfigResponse{}
	_body, _err := client.SaveApgroupPortalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveApgroupSsidConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveApgroupSsidConfigResponse
func (client *Client) SaveApgroupSsidConfigWithOptions(request *SaveApgroupSsidConfigRequest, runtime *dara.RuntimeOptions) (_result *SaveApgroupSsidConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcctPort) {
		query["AcctPort"] = request.AcctPort
	}

	if !dara.IsNil(request.AcctSecret) {
		query["AcctSecret"] = request.AcctSecret
	}

	if !dara.IsNil(request.AcctServer) {
		query["AcctServer"] = request.AcctServer
	}

	if !dara.IsNil(request.ApgroupId) {
		query["ApgroupId"] = request.ApgroupId
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AuthCache) {
		query["AuthCache"] = request.AuthCache
	}

	if !dara.IsNil(request.AuthPort) {
		query["AuthPort"] = request.AuthPort
	}

	if !dara.IsNil(request.AuthSecret) {
		query["AuthSecret"] = request.AuthSecret
	}

	if !dara.IsNil(request.AuthServer) {
		query["AuthServer"] = request.AuthServer
	}

	if !dara.IsNil(request.Binding) {
		query["Binding"] = request.Binding
	}

	if !dara.IsNil(request.Cir) {
		query["Cir"] = request.Cir
	}

	if !dara.IsNil(request.DaeClient) {
		query["DaeClient"] = request.DaeClient
	}

	if !dara.IsNil(request.DaePort) {
		query["DaePort"] = request.DaePort
	}

	if !dara.IsNil(request.DaeSecret) {
		query["DaeSecret"] = request.DaeSecret
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.DisassocLowAck) {
		query["DisassocLowAck"] = request.DisassocLowAck
	}

	if !dara.IsNil(request.DisassocWeakRssi) {
		query["DisassocWeakRssi"] = request.DisassocWeakRssi
	}

	if !dara.IsNil(request.DynamicVlan) {
		query["DynamicVlan"] = request.DynamicVlan
	}

	if !dara.IsNil(request.Effect) {
		query["Effect"] = request.Effect
	}

	if !dara.IsNil(request.EncKey) {
		query["EncKey"] = request.EncKey
	}

	if !dara.IsNil(request.Encryption) {
		query["Encryption"] = request.Encryption
	}

	if !dara.IsNil(request.Hidden) {
		query["Hidden"] = request.Hidden
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Ieee80211w) {
		query["Ieee80211w"] = request.Ieee80211w
	}

	if !dara.IsNil(request.IgnoreWeakProbe) {
		query["IgnoreWeakProbe"] = request.IgnoreWeakProbe
	}

	if !dara.IsNil(request.Isolate) {
		query["Isolate"] = request.Isolate
	}

	if !dara.IsNil(request.LiteEffect) {
		query["LiteEffect"] = request.LiteEffect
	}

	if !dara.IsNil(request.MaxInactivity) {
		query["MaxInactivity"] = request.MaxInactivity
	}

	if !dara.IsNil(request.Maxassoc) {
		query["Maxassoc"] = request.Maxassoc
	}

	if !dara.IsNil(request.MulticastForward) {
		query["MulticastForward"] = request.MulticastForward
	}

	if !dara.IsNil(request.Nasid) {
		query["Nasid"] = request.Nasid
	}

	if !dara.IsNil(request.Network) {
		query["Network"] = request.Network
	}

	if !dara.IsNil(request.NewSsid) {
		query["NewSsid"] = request.NewSsid
	}

	if !dara.IsNil(request.Ownip) {
		query["Ownip"] = request.Ownip
	}

	if !dara.IsNil(request.SecondaryAcctPort) {
		query["SecondaryAcctPort"] = request.SecondaryAcctPort
	}

	if !dara.IsNil(request.SecondaryAcctSecret) {
		query["SecondaryAcctSecret"] = request.SecondaryAcctSecret
	}

	if !dara.IsNil(request.SecondaryAcctServer) {
		query["SecondaryAcctServer"] = request.SecondaryAcctServer
	}

	if !dara.IsNil(request.SecondaryAuthPort) {
		query["SecondaryAuthPort"] = request.SecondaryAuthPort
	}

	if !dara.IsNil(request.SecondaryAuthSecret) {
		query["SecondaryAuthSecret"] = request.SecondaryAuthSecret
	}

	if !dara.IsNil(request.SecondaryAuthServer) {
		query["SecondaryAuthServer"] = request.SecondaryAuthServer
	}

	if !dara.IsNil(request.ShortPreamble) {
		query["ShortPreamble"] = request.ShortPreamble
	}

	if !dara.IsNil(request.Ssid) {
		query["Ssid"] = request.Ssid
	}

	if !dara.IsNil(request.SsidLb) {
		query["SsidLb"] = request.SsidLb
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VlanDhcp) {
		query["VlanDhcp"] = request.VlanDhcp
	}

	if !dara.IsNil(request.Wmm) {
		query["Wmm"] = request.Wmm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveApgroupSsidConfig"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveApgroupSsidConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveApgroupSsidConfigRequest
//
// @return SaveApgroupSsidConfigResponse
func (client *Client) SaveApgroupSsidConfig(request *SaveApgroupSsidConfigRequest) (_result *SaveApgroupSsidConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveApgroupSsidConfigResponse{}
	_body, _err := client.SaveApgroupSsidConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetApAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApAddressResponse
func (client *Client) SetApAddressWithOptions(request *SetApAddressRequest, runtime *dara.RuntimeOptions) (_result *SetApAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApAreaName) {
		query["ApAreaName"] = request.ApAreaName
	}

	if !dara.IsNil(request.ApBuildingName) {
		query["ApBuildingName"] = request.ApBuildingName
	}

	if !dara.IsNil(request.ApCampusName) {
		query["ApCampusName"] = request.ApCampusName
	}

	if !dara.IsNil(request.ApCityName) {
		query["ApCityName"] = request.ApCityName
	}

	if !dara.IsNil(request.ApFloor) {
		query["ApFloor"] = request.ApFloor
	}

	if !dara.IsNil(request.ApGroup) {
		query["ApGroup"] = request.ApGroup
	}

	if !dara.IsNil(request.ApName) {
		query["ApName"] = request.ApName
	}

	if !dara.IsNil(request.ApNationName) {
		query["ApNationName"] = request.ApNationName
	}

	if !dara.IsNil(request.ApProvinceName) {
		query["ApProvinceName"] = request.ApProvinceName
	}

	if !dara.IsNil(request.ApUnitId) {
		query["ApUnitId"] = request.ApUnitId
	}

	if !dara.IsNil(request.ApUnitName) {
		query["ApUnitName"] = request.ApUnitName
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Lat) {
		query["Lat"] = request.Lat
	}

	if !dara.IsNil(request.Lng) {
		query["Lng"] = request.Lng
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApAddress"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetApAddressRequest
//
// @return SetApAddressResponse
func (client *Client) SetApAddress(request *SetApAddressRequest) (_result *SetApAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApAddressResponse{}
	_body, _err := client.SetApAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetApNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApNameResponse
func (client *Client) SetApNameWithOptions(request *SetApNameRequest, runtime *dara.RuntimeOptions) (_result *SetApNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApMac) {
		query["ApMac"] = request.ApMac
	}

	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApName"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetApNameRequest
//
// @return SetApNameResponse
func (client *Client) SetApName(request *SetApNameRequest) (_result *SetApNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApNameResponse{}
	_body, _err := client.SetApNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注销AP设备
//
// @param request - UnRegisterApAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnRegisterApAssetResponse
func (client *Client) UnRegisterApAssetWithOptions(request *UnRegisterApAssetRequest, runtime *dara.RuntimeOptions) (_result *UnRegisterApAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AssetApgroupId) {
		query["AssetApgroupId"] = request.AssetApgroupId
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	if !dara.IsNil(request.SerialNo) {
		query["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.UseFor) {
		query["UseFor"] = request.UseFor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnRegisterApAsset"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnRegisterApAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注销AP设备
//
// @param request - UnRegisterApAssetRequest
//
// @return UnRegisterApAssetResponse
func (client *Client) UnRegisterApAsset(request *UnRegisterApAssetRequest) (_result *UnRegisterApAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnRegisterApAssetResponse{}
	_body, _err := client.UnRegisterApAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateNetDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetDeviceInfoResponse
func (client *Client) UpdateNetDeviceInfoWithOptions(request *UpdateNetDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Devices) {
		body["Devices"] = request.Devices
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNetDeviceInfo"),
		Version:     dara.String("2019-11-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNetDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateNetDeviceInfoRequest
//
// @return UpdateNetDeviceInfoResponse
func (client *Client) UpdateNetDeviceInfo(request *UpdateNetDeviceInfoRequest) (_result *UpdateNetDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNetDeviceInfoResponse{}
	_body, _err := client.UpdateNetDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
