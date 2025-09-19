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
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              dara.String("cloudpush.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("cloudpush.aliyuncs.com"),
		"ap-south-1":                  dara.String("cloudpush.aliyuncs.com"),
		"ap-southeast-1":              dara.String("cloudpush.aliyuncs.com"),
		"ap-southeast-2":              dara.String("cloudpush.aliyuncs.com"),
		"ap-southeast-3":              dara.String("cloudpush.aliyuncs.com"),
		"ap-southeast-5":              dara.String("cloudpush.aliyuncs.com"),
		"cn-beijing":                  dara.String("cloudpush.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("cloudpush.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("cloudpush.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("cloudpush.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("cloudpush.aliyuncs.com"),
		"cn-chengdu":                  dara.String("cloudpush.aliyuncs.com"),
		"cn-edge-1":                   dara.String("cloudpush.aliyuncs.com"),
		"cn-fujian":                   dara.String("cloudpush.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("cloudpush.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("cloudpush.aliyuncs.com"),
		"cn-hongkong":                 dara.String("cloudpush.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("cloudpush.aliyuncs.com"),
		"cn-huhehaote":                dara.String("cloudpush.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("cloudpush.aliyuncs.com"),
		"cn-qingdao":                  dara.String("cloudpush.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("cloudpush.aliyuncs.com"),
		"cn-shanghai":                 dara.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("cloudpush.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("cloudpush.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("cloudpush.aliyuncs.com"),
		"cn-wuhan":                    dara.String("cloudpush.aliyuncs.com"),
		"cn-yushanfang":               dara.String("cloudpush.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("cloudpush.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("cloudpush.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("cloudpush.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("cloudpush.aliyuncs.com"),
		"eu-central-1":                dara.String("cloudpush.aliyuncs.com"),
		"eu-west-1":                   dara.String("cloudpush.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("cloudpush.aliyuncs.com"),
		"me-east-1":                   dara.String("cloudpush.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("cloudpush.aliyuncs.com"),
		"us-east-1":                   dara.String("cloudpush.aliyuncs.com"),
		"us-west-1":                   dara.String("cloudpush.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("push"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 绑定别名
//
// @param request - BindAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAliasResponse
func (client *Client) BindAliasWithOptions(request *BindAliasRequest, runtime *dara.RuntimeOptions) (_result *BindAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAlias"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定别名
//
// @param request - BindAliasRequest
//
// @return BindAliasResponse
func (client *Client) BindAlias(request *BindAliasRequest) (_result *BindAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAliasResponse{}
	_body, _err := client.BindAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定手机号码
//
// @param request - BindPhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPhoneResponse
func (client *Client) BindPhoneWithOptions(request *BindPhoneRequest, runtime *dara.RuntimeOptions) (_result *BindPhoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPhone"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定手机号码
//
// @param request - BindPhoneRequest
//
// @return BindPhoneResponse
func (client *Client) BindPhone(request *BindPhoneRequest) (_result *BindPhoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindPhoneResponse{}
	_body, _err := client.BindPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定标签
//
// @param request - BindTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindTagResponse
func (client *Client) BindTagWithOptions(request *BindTagRequest, runtime *dara.RuntimeOptions) (_result *BindTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.ClientKey) {
		query["ClientKey"] = request.ClientKey
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindTag"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定标签
//
// @param request - BindTagRequest
//
// @return BindTagResponse
func (client *Client) BindTag(request *BindTagRequest) (_result *BindTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindTagResponse{}
	_body, _err := client.BindTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消定时推送任务
//
// @param request - CancelPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPushResponse
func (client *Client) CancelPushWithOptions(request *CancelPushRequest, runtime *dara.RuntimeOptions) (_result *CancelPushResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消定时推送任务
//
// @param request - CancelPushRequest
//
// @return CancelPushResponse
func (client *Client) CancelPush(request *CancelPushRequest) (_result *CancelPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelPushResponse{}
	_body, _err := client.CancelPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CheckCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCertificateResponse
func (client *Client) CheckCertificateWithOptions(request *CheckCertificateRequest, runtime *dara.RuntimeOptions) (_result *CheckCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCertificate"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckCertificateRequest
//
// @return CheckCertificateResponse
func (client *Client) CheckCertificate(request *CheckCertificateRequest) (_result *CheckCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckCertificateResponse{}
	_body, _err := client.CheckCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CheckDevice is deprecated, please use Push::2016-08-01::CheckDevices instead.
//
// Summary:
//
// 【废弃】验证设备有效性
//
// @param request - CheckDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDeviceResponse
func (client *Client) CheckDeviceWithOptions(request *CheckDeviceRequest, runtime *dara.RuntimeOptions) (_result *CheckDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDevice"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CheckDevice is deprecated, please use Push::2016-08-01::CheckDevices instead.
//
// Summary:
//
// 【废弃】验证设备有效性
//
// @param request - CheckDeviceRequest
//
// @return CheckDeviceResponse
// Deprecated
func (client *Client) CheckDevice(request *CheckDeviceRequest) (_result *CheckDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDeviceResponse{}
	_body, _err := client.CheckDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量检查设备有效性
//
// @param request - CheckDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDevicesResponse
func (client *Client) CheckDevicesWithOptions(request *CheckDevicesRequest, runtime *dara.RuntimeOptions) (_result *CheckDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceIds) {
		query["DeviceIds"] = request.DeviceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDevices"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量检查设备有效性
//
// @param request - CheckDevicesRequest
//
// @return CheckDevicesResponse
func (client *Client) CheckDevices(request *CheckDevicesRequest) (_result *CheckDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDevicesResponse{}
	_body, _err := client.CheckDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 完成持续推送任务
//
// @param request - CompleteContinuouslyPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteContinuouslyPushResponse
func (client *Client) CompleteContinuouslyPushWithOptions(request *CompleteContinuouslyPushRequest, runtime *dara.RuntimeOptions) (_result *CompleteContinuouslyPushResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteContinuouslyPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteContinuouslyPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完成持续推送任务
//
// @param request - CompleteContinuouslyPushRequest
//
// @return CompleteContinuouslyPushResponse
func (client *Client) CompleteContinuouslyPush(request *CompleteContinuouslyPushRequest) (_result *CompleteContinuouslyPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompleteContinuouslyPushResponse{}
	_body, _err := client.CompleteContinuouslyPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 持续推送
//
// @param request - ContinuouslyPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinuouslyPushResponse
func (client *Client) ContinuouslyPushWithOptions(request *ContinuouslyPushRequest, runtime *dara.RuntimeOptions) (_result *ContinuouslyPushResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinuouslyPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinuouslyPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 持续推送
//
// @param request - ContinuouslyPushRequest
//
// @return ContinuouslyPushResponse
func (client *Client) ContinuouslyPush(request *ContinuouslyPushRequest) (_result *ContinuouslyPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinuouslyPushResponse{}
	_body, _err := client.ContinuouslyPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListSummaryApps is deprecated, please use Mhub::2017-08-25::ListApps instead.
//
// Summary:
//
// 【废弃】查询用户已创建的app列表
//
// @param request - ListSummaryAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSummaryAppsResponse
// Deprecated
func (client *Client) ListSummaryAppsWithOptions(runtime *dara.RuntimeOptions) (_result *ListSummaryAppsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListSummaryApps"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSummaryAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListSummaryApps is deprecated, please use Mhub::2017-08-25::ListApps instead.
//
// Summary:
//
// 【废弃】查询用户已创建的app列表
//
// @return ListSummaryAppsResponse
// Deprecated
func (client *Client) ListSummaryApps() (_result *ListSummaryAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSummaryAppsResponse{}
	_body, _err := client.ListSummaryAppsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标签列表
//
// @param request - ListTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTags"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标签列表
//
// @param request - ListTagsRequest
//
// @return ListTagsResponse
func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量推送
//
// @param request - MassPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MassPushResponse
func (client *Client) MassPushWithOptions(request *MassPushRequest, runtime *dara.RuntimeOptions) (_result *MassPushResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.IdempotentToken) {
		query["IdempotentToken"] = request.IdempotentToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PushTask) {
		body["PushTask"] = request.PushTask
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MassPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MassPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量推送
//
// @param request - MassPushRequest
//
// @return MassPushResponse
func (client *Client) MassPush(request *MassPushRequest) (_result *MassPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MassPushResponse{}
	_body, _err := client.MassPushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 高级推送接口
//
// @param tmpReq - PushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResponse
func (client *Client) PushWithOptions(tmpReq *PushRequest, runtime *dara.RuntimeOptions) (_result *PushResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PushShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidOppoPrivateContentParameters) {
		request.AndroidOppoPrivateContentParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidOppoPrivateContentParameters, dara.String("AndroidOppoPrivateContentParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidOppoPrivateTitleParameters) {
		request.AndroidOppoPrivateTitleParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidOppoPrivateTitleParameters, dara.String("AndroidOppoPrivateTitleParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidActivity) {
		query["AndroidActivity"] = request.AndroidActivity
	}

	if !dara.IsNil(request.AndroidBadgeAddNum) {
		query["AndroidBadgeAddNum"] = request.AndroidBadgeAddNum
	}

	if !dara.IsNil(request.AndroidBadgeClass) {
		query["AndroidBadgeClass"] = request.AndroidBadgeClass
	}

	if !dara.IsNil(request.AndroidBadgeSetNum) {
		query["AndroidBadgeSetNum"] = request.AndroidBadgeSetNum
	}

	if !dara.IsNil(request.AndroidBigBody) {
		query["AndroidBigBody"] = request.AndroidBigBody
	}

	if !dara.IsNil(request.AndroidBigPictureUrl) {
		query["AndroidBigPictureUrl"] = request.AndroidBigPictureUrl
	}

	if !dara.IsNil(request.AndroidBigTitle) {
		query["AndroidBigTitle"] = request.AndroidBigTitle
	}

	if !dara.IsNil(request.AndroidExtParameters) {
		query["AndroidExtParameters"] = request.AndroidExtParameters
	}

	if !dara.IsNil(request.AndroidHonorTargetUserType) {
		query["AndroidHonorTargetUserType"] = request.AndroidHonorTargetUserType
	}

	if !dara.IsNil(request.AndroidHuaweiReceiptId) {
		query["AndroidHuaweiReceiptId"] = request.AndroidHuaweiReceiptId
	}

	if !dara.IsNil(request.AndroidHuaweiTargetUserType) {
		query["AndroidHuaweiTargetUserType"] = request.AndroidHuaweiTargetUserType
	}

	if !dara.IsNil(request.AndroidImageUrl) {
		query["AndroidImageUrl"] = request.AndroidImageUrl
	}

	if !dara.IsNil(request.AndroidInboxBody) {
		query["AndroidInboxBody"] = request.AndroidInboxBody
	}

	if !dara.IsNil(request.AndroidMeizuNoticeMsgType) {
		query["AndroidMeizuNoticeMsgType"] = request.AndroidMeizuNoticeMsgType
	}

	if !dara.IsNil(request.AndroidMessageHuaweiCategory) {
		query["AndroidMessageHuaweiCategory"] = request.AndroidMessageHuaweiCategory
	}

	if !dara.IsNil(request.AndroidMessageHuaweiUrgency) {
		query["AndroidMessageHuaweiUrgency"] = request.AndroidMessageHuaweiUrgency
	}

	if !dara.IsNil(request.AndroidMessageOppoCategory) {
		query["AndroidMessageOppoCategory"] = request.AndroidMessageOppoCategory
	}

	if !dara.IsNil(request.AndroidMessageOppoNotifyLevel) {
		query["AndroidMessageOppoNotifyLevel"] = request.AndroidMessageOppoNotifyLevel
	}

	if !dara.IsNil(request.AndroidMessageVivoCategory) {
		query["AndroidMessageVivoCategory"] = request.AndroidMessageVivoCategory
	}

	if !dara.IsNil(request.AndroidMusic) {
		query["AndroidMusic"] = request.AndroidMusic
	}

	if !dara.IsNil(request.AndroidNotificationBarPriority) {
		query["AndroidNotificationBarPriority"] = request.AndroidNotificationBarPriority
	}

	if !dara.IsNil(request.AndroidNotificationBarType) {
		query["AndroidNotificationBarType"] = request.AndroidNotificationBarType
	}

	if !dara.IsNil(request.AndroidNotificationChannel) {
		query["AndroidNotificationChannel"] = request.AndroidNotificationChannel
	}

	if !dara.IsNil(request.AndroidNotificationGroup) {
		query["AndroidNotificationGroup"] = request.AndroidNotificationGroup
	}

	if !dara.IsNil(request.AndroidNotificationHonorChannel) {
		query["AndroidNotificationHonorChannel"] = request.AndroidNotificationHonorChannel
	}

	if !dara.IsNil(request.AndroidNotificationHuaweiChannel) {
		query["AndroidNotificationHuaweiChannel"] = request.AndroidNotificationHuaweiChannel
	}

	if !dara.IsNil(request.AndroidNotificationNotifyId) {
		query["AndroidNotificationNotifyId"] = request.AndroidNotificationNotifyId
	}

	if !dara.IsNil(request.AndroidNotificationThreadId) {
		query["AndroidNotificationThreadId"] = request.AndroidNotificationThreadId
	}

	if !dara.IsNil(request.AndroidNotificationVivoChannel) {
		query["AndroidNotificationVivoChannel"] = request.AndroidNotificationVivoChannel
	}

	if !dara.IsNil(request.AndroidNotificationXiaomiChannel) {
		query["AndroidNotificationXiaomiChannel"] = request.AndroidNotificationXiaomiChannel
	}

	if !dara.IsNil(request.AndroidNotifyType) {
		query["AndroidNotifyType"] = request.AndroidNotifyType
	}

	if !dara.IsNil(request.AndroidOpenType) {
		query["AndroidOpenType"] = request.AndroidOpenType
	}

	if !dara.IsNil(request.AndroidOpenUrl) {
		query["AndroidOpenUrl"] = request.AndroidOpenUrl
	}

	if !dara.IsNil(request.AndroidOppoPrivateContentParametersShrink) {
		query["AndroidOppoPrivateContentParameters"] = request.AndroidOppoPrivateContentParametersShrink
	}

	if !dara.IsNil(request.AndroidOppoPrivateMsgTemplateId) {
		query["AndroidOppoPrivateMsgTemplateId"] = request.AndroidOppoPrivateMsgTemplateId
	}

	if !dara.IsNil(request.AndroidOppoPrivateTitleParametersShrink) {
		query["AndroidOppoPrivateTitleParameters"] = request.AndroidOppoPrivateTitleParametersShrink
	}

	if !dara.IsNil(request.AndroidPopupActivity) {
		query["AndroidPopupActivity"] = request.AndroidPopupActivity
	}

	if !dara.IsNil(request.AndroidPopupBody) {
		query["AndroidPopupBody"] = request.AndroidPopupBody
	}

	if !dara.IsNil(request.AndroidPopupTitle) {
		query["AndroidPopupTitle"] = request.AndroidPopupTitle
	}

	if !dara.IsNil(request.AndroidRemind) {
		query["AndroidRemind"] = request.AndroidRemind
	}

	if !dara.IsNil(request.AndroidRenderStyle) {
		query["AndroidRenderStyle"] = request.AndroidRenderStyle
	}

	if !dara.IsNil(request.AndroidTargetUserType) {
		query["AndroidTargetUserType"] = request.AndroidTargetUserType
	}

	if !dara.IsNil(request.AndroidVivoPushMode) {
		query["AndroidVivoPushMode"] = request.AndroidVivoPushMode
	}

	if !dara.IsNil(request.AndroidVivoReceiptId) {
		query["AndroidVivoReceiptId"] = request.AndroidVivoReceiptId
	}

	if !dara.IsNil(request.AndroidXiaoMiActivity) {
		query["AndroidXiaoMiActivity"] = request.AndroidXiaoMiActivity
	}

	if !dara.IsNil(request.AndroidXiaoMiNotifyBody) {
		query["AndroidXiaoMiNotifyBody"] = request.AndroidXiaoMiNotifyBody
	}

	if !dara.IsNil(request.AndroidXiaoMiNotifyTitle) {
		query["AndroidXiaoMiNotifyTitle"] = request.AndroidXiaoMiNotifyTitle
	}

	if !dara.IsNil(request.AndroidXiaomiBigPictureUrl) {
		query["AndroidXiaomiBigPictureUrl"] = request.AndroidXiaomiBigPictureUrl
	}

	if !dara.IsNil(request.AndroidXiaomiImageUrl) {
		query["AndroidXiaomiImageUrl"] = request.AndroidXiaomiImageUrl
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.HarmonyAction) {
		query["HarmonyAction"] = request.HarmonyAction
	}

	if !dara.IsNil(request.HarmonyActionType) {
		query["HarmonyActionType"] = request.HarmonyActionType
	}

	if !dara.IsNil(request.HarmonyBadgeAddNum) {
		query["HarmonyBadgeAddNum"] = request.HarmonyBadgeAddNum
	}

	if !dara.IsNil(request.HarmonyBadgeSetNum) {
		query["HarmonyBadgeSetNum"] = request.HarmonyBadgeSetNum
	}

	if !dara.IsNil(request.HarmonyCategory) {
		query["HarmonyCategory"] = request.HarmonyCategory
	}

	if !dara.IsNil(request.HarmonyExtParameters) {
		query["HarmonyExtParameters"] = request.HarmonyExtParameters
	}

	if !dara.IsNil(request.HarmonyExtensionExtraData) {
		query["HarmonyExtensionExtraData"] = request.HarmonyExtensionExtraData
	}

	if !dara.IsNil(request.HarmonyExtensionPush) {
		query["HarmonyExtensionPush"] = request.HarmonyExtensionPush
	}

	if !dara.IsNil(request.HarmonyImageUrl) {
		query["HarmonyImageUrl"] = request.HarmonyImageUrl
	}

	if !dara.IsNil(request.HarmonyInboxContent) {
		query["HarmonyInboxContent"] = request.HarmonyInboxContent
	}

	if !dara.IsNil(request.HarmonyNotificationSlotType) {
		query["HarmonyNotificationSlotType"] = request.HarmonyNotificationSlotType
	}

	if !dara.IsNil(request.HarmonyNotifyId) {
		query["HarmonyNotifyId"] = request.HarmonyNotifyId
	}

	if !dara.IsNil(request.HarmonyReceiptId) {
		query["HarmonyReceiptId"] = request.HarmonyReceiptId
	}

	if !dara.IsNil(request.HarmonyRemind) {
		query["HarmonyRemind"] = request.HarmonyRemind
	}

	if !dara.IsNil(request.HarmonyRemindBody) {
		query["HarmonyRemindBody"] = request.HarmonyRemindBody
	}

	if !dara.IsNil(request.HarmonyRemindTitle) {
		query["HarmonyRemindTitle"] = request.HarmonyRemindTitle
	}

	if !dara.IsNil(request.HarmonyRenderStyle) {
		query["HarmonyRenderStyle"] = request.HarmonyRenderStyle
	}

	if !dara.IsNil(request.HarmonyTestMessage) {
		query["HarmonyTestMessage"] = request.HarmonyTestMessage
	}

	if !dara.IsNil(request.HarmonyUri) {
		query["HarmonyUri"] = request.HarmonyUri
	}

	if !dara.IsNil(request.IdempotentToken) {
		query["IdempotentToken"] = request.IdempotentToken
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.PushTime) {
		query["PushTime"] = request.PushTime
	}

	if !dara.IsNil(request.PushType) {
		query["PushType"] = request.PushType
	}

	if !dara.IsNil(request.SendChannels) {
		query["SendChannels"] = request.SendChannels
	}

	if !dara.IsNil(request.SendSpeed) {
		query["SendSpeed"] = request.SendSpeed
	}

	if !dara.IsNil(request.SmsDelaySecs) {
		query["SmsDelaySecs"] = request.SmsDelaySecs
	}

	if !dara.IsNil(request.SmsParams) {
		query["SmsParams"] = request.SmsParams
	}

	if !dara.IsNil(request.SmsSendPolicy) {
		query["SmsSendPolicy"] = request.SmsSendPolicy
	}

	if !dara.IsNil(request.SmsSignName) {
		query["SmsSignName"] = request.SmsSignName
	}

	if !dara.IsNil(request.SmsTemplateName) {
		query["SmsTemplateName"] = request.SmsTemplateName
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.Trim) {
		query["Trim"] = request.Trim
	}

	if !dara.IsNil(request.IOSApnsEnv) {
		query["iOSApnsEnv"] = request.IOSApnsEnv
	}

	if !dara.IsNil(request.IOSBadge) {
		query["iOSBadge"] = request.IOSBadge
	}

	if !dara.IsNil(request.IOSBadgeAutoIncrement) {
		query["iOSBadgeAutoIncrement"] = request.IOSBadgeAutoIncrement
	}

	if !dara.IsNil(request.IOSExtParameters) {
		query["iOSExtParameters"] = request.IOSExtParameters
	}

	if !dara.IsNil(request.IOSInterruptionLevel) {
		query["iOSInterruptionLevel"] = request.IOSInterruptionLevel
	}

	if !dara.IsNil(request.IOSLiveActivityAttributes) {
		query["iOSLiveActivityAttributes"] = request.IOSLiveActivityAttributes
	}

	if !dara.IsNil(request.IOSLiveActivityAttributesType) {
		query["iOSLiveActivityAttributesType"] = request.IOSLiveActivityAttributesType
	}

	if !dara.IsNil(request.IOSLiveActivityContentState) {
		query["iOSLiveActivityContentState"] = request.IOSLiveActivityContentState
	}

	if !dara.IsNil(request.IOSLiveActivityDismissalDate) {
		query["iOSLiveActivityDismissalDate"] = request.IOSLiveActivityDismissalDate
	}

	if !dara.IsNil(request.IOSLiveActivityEvent) {
		query["iOSLiveActivityEvent"] = request.IOSLiveActivityEvent
	}

	if !dara.IsNil(request.IOSLiveActivityId) {
		query["iOSLiveActivityId"] = request.IOSLiveActivityId
	}

	if !dara.IsNil(request.IOSLiveActivityStaleDate) {
		query["iOSLiveActivityStaleDate"] = request.IOSLiveActivityStaleDate
	}

	if !dara.IsNil(request.IOSMusic) {
		query["iOSMusic"] = request.IOSMusic
	}

	if !dara.IsNil(request.IOSMutableContent) {
		query["iOSMutableContent"] = request.IOSMutableContent
	}

	if !dara.IsNil(request.IOSNotificationCategory) {
		query["iOSNotificationCategory"] = request.IOSNotificationCategory
	}

	if !dara.IsNil(request.IOSNotificationCollapseId) {
		query["iOSNotificationCollapseId"] = request.IOSNotificationCollapseId
	}

	if !dara.IsNil(request.IOSNotificationThreadId) {
		query["iOSNotificationThreadId"] = request.IOSNotificationThreadId
	}

	if !dara.IsNil(request.IOSRelevanceScore) {
		query["iOSRelevanceScore"] = request.IOSRelevanceScore
	}

	if !dara.IsNil(request.IOSRemind) {
		query["iOSRemind"] = request.IOSRemind
	}

	if !dara.IsNil(request.IOSRemindBody) {
		query["iOSRemindBody"] = request.IOSRemindBody
	}

	if !dara.IsNil(request.IOSSilentNotification) {
		query["iOSSilentNotification"] = request.IOSSilentNotification
	}

	if !dara.IsNil(request.IOSSubtitle) {
		query["iOSSubtitle"] = request.IOSSubtitle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Push"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 高级推送接口
//
// @param request - PushRequest
//
// @return PushResponse
func (client *Client) Push(request *PushRequest) (_result *PushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushResponse{}
	_body, _err := client.PushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送消息给Android设备
//
// @param request - PushMessageToAndroidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMessageToAndroidResponse
func (client *Client) PushMessageToAndroidWithOptions(request *PushMessageToAndroidRequest, runtime *dara.RuntimeOptions) (_result *PushMessageToAndroidResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMessageToAndroid"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMessageToAndroidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送消息给Android设备
//
// @param request - PushMessageToAndroidRequest
//
// @return PushMessageToAndroidResponse
func (client *Client) PushMessageToAndroid(request *PushMessageToAndroidRequest) (_result *PushMessageToAndroidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushMessageToAndroidResponse{}
	_body, _err := client.PushMessageToAndroidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送消息给iOS设备
//
// @param request - PushMessageToiOSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMessageToiOSResponse
func (client *Client) PushMessageToiOSWithOptions(request *PushMessageToiOSRequest, runtime *dara.RuntimeOptions) (_result *PushMessageToiOSResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMessageToiOS"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMessageToiOSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送消息给iOS设备
//
// @param request - PushMessageToiOSRequest
//
// @return PushMessageToiOSResponse
func (client *Client) PushMessageToiOS(request *PushMessageToiOSRequest) (_result *PushMessageToiOSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushMessageToiOSResponse{}
	_body, _err := client.PushMessageToiOSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送通知给Android设备
//
// @param request - PushNoticeToAndroidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushNoticeToAndroidResponse
func (client *Client) PushNoticeToAndroidWithOptions(request *PushNoticeToAndroidRequest, runtime *dara.RuntimeOptions) (_result *PushNoticeToAndroidResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.ExtParameters) {
		query["ExtParameters"] = request.ExtParameters
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushNoticeToAndroid"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushNoticeToAndroidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送通知给Android设备
//
// @param request - PushNoticeToAndroidRequest
//
// @return PushNoticeToAndroidResponse
func (client *Client) PushNoticeToAndroid(request *PushNoticeToAndroidRequest) (_result *PushNoticeToAndroidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushNoticeToAndroidResponse{}
	_body, _err := client.PushNoticeToAndroidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送通知给iOS设备
//
// @param request - PushNoticeToiOSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushNoticeToiOSResponse
func (client *Client) PushNoticeToiOSWithOptions(request *PushNoticeToiOSRequest, runtime *dara.RuntimeOptions) (_result *PushNoticeToiOSResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApnsEnv) {
		query["ApnsEnv"] = request.ApnsEnv
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.ExtParameters) {
		query["ExtParameters"] = request.ExtParameters
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushNoticeToiOS"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushNoticeToiOSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送通知给iOS设备
//
// @param request - PushNoticeToiOSRequest
//
// @return PushNoticeToiOSResponse
func (client *Client) PushNoticeToiOS(request *PushNoticeToiOSRequest) (_result *PushNoticeToiOSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushNoticeToiOSResponse{}
	_body, _err := client.PushNoticeToiOSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询别名
//
// @param request - QueryAliasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAliasesResponse
func (client *Client) QueryAliasesWithOptions(request *QueryAliasesRequest, runtime *dara.RuntimeOptions) (_result *QueryAliasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAliases"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAliasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询别名
//
// @param request - QueryAliasesRequest
//
// @return QueryAliasesResponse
func (client *Client) QueryAliases(request *QueryAliasesRequest) (_result *QueryAliasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAliasesResponse{}
	_body, _err := client.QueryAliasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备详情
//
// @param request - QueryDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceInfoResponse
func (client *Client) QueryDeviceInfoWithOptions(request *QueryDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceInfo"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备详情
//
// @param request - QueryDeviceInfoRequest
//
// @return QueryDeviceInfoResponse
func (client *Client) QueryDeviceInfo(request *QueryDeviceInfoRequest) (_result *QueryDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceInfoResponse{}
	_body, _err := client.QueryDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备新增与留存
//
// @param request - QueryDeviceStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceStatResponse
func (client *Client) QueryDeviceStatWithOptions(request *QueryDeviceStatRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceStatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceStat"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备新增与留存
//
// @param request - QueryDeviceStatRequest
//
// @return QueryDeviceStatResponse
func (client *Client) QueryDeviceStat(request *QueryDeviceStatRequest) (_result *QueryDeviceStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeviceStatResponse{}
	_body, _err := client.QueryDeviceStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过账户查询设备列表
//
// @param request - QueryDevicesByAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicesByAccountResponse
func (client *Client) QueryDevicesByAccountWithOptions(request *QueryDevicesByAccountRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicesByAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDevicesByAccount"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDevicesByAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过账户查询设备列表
//
// @param request - QueryDevicesByAccountRequest
//
// @return QueryDevicesByAccountResponse
func (client *Client) QueryDevicesByAccount(request *QueryDevicesByAccountRequest) (_result *QueryDevicesByAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDevicesByAccountResponse{}
	_body, _err := client.QueryDevicesByAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过别名查询设备列表
//
// @param request - QueryDevicesByAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicesByAliasResponse
func (client *Client) QueryDevicesByAliasWithOptions(request *QueryDevicesByAliasRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicesByAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["Alias"] = request.Alias
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDevicesByAlias"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDevicesByAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过别名查询设备列表
//
// @param request - QueryDevicesByAliasRequest
//
// @return QueryDevicesByAliasResponse
func (client *Client) QueryDevicesByAlias(request *QueryDevicesByAliasRequest) (_result *QueryDevicesByAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDevicesByAliasResponse{}
	_body, _err := client.QueryDevicesByAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryPushRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushRecordsResponse
func (client *Client) QueryPushRecordsWithOptions(request *QueryPushRecordsRequest, runtime *dara.RuntimeOptions) (_result *QueryPushRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PushType) {
		query["PushType"] = request.PushType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushRecords"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPushRecordsRequest
//
// @return QueryPushRecordsResponse
func (client *Client) QueryPushRecords(request *QueryPushRecordsRequest) (_result *QueryPushRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPushRecordsResponse{}
	_body, _err := client.QueryPushRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # App维度推送统计
//
// @param request - QueryPushStatByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushStatByAppResponse
func (client *Client) QueryPushStatByAppWithOptions(request *QueryPushStatByAppRequest, runtime *dara.RuntimeOptions) (_result *QueryPushStatByAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushStatByApp"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushStatByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # App维度推送统计
//
// @param request - QueryPushStatByAppRequest
//
// @return QueryPushStatByAppResponse
func (client *Client) QueryPushStatByApp(request *QueryPushStatByAppRequest) (_result *QueryPushStatByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPushStatByAppResponse{}
	_body, _err := client.QueryPushStatByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 任务维度推送统计
//
// @param request - QueryPushStatByMsgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushStatByMsgResponse
func (client *Client) QueryPushStatByMsgWithOptions(request *QueryPushStatByMsgRequest, runtime *dara.RuntimeOptions) (_result *QueryPushStatByMsgResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushStatByMsg"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushStatByMsgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务维度推送统计
//
// @param request - QueryPushStatByMsgRequest
//
// @return QueryPushStatByMsgResponse
func (client *Client) QueryPushStatByMsg(request *QueryPushStatByMsgRequest) (_result *QueryPushStatByMsgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPushStatByMsgResponse{}
	_body, _err := client.QueryPushStatByMsgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询标签列表
//
// @param request - QueryTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagsResponse
func (client *Client) QueryTagsWithOptions(request *QueryTagsRequest, runtime *dara.RuntimeOptions) (_result *QueryTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.ClientKey) {
		query["ClientKey"] = request.ClientKey
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTags"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询标签列表
//
// @param request - QueryTagsRequest
//
// @return QueryTagsResponse
func (client *Client) QueryTags(request *QueryTagsRequest) (_result *QueryTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTagsResponse{}
	_body, _err := client.QueryTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 去重设备统计
//
// @param request - QueryUniqueDeviceStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUniqueDeviceStatResponse
func (client *Client) QueryUniqueDeviceStatWithOptions(request *QueryUniqueDeviceStatRequest, runtime *dara.RuntimeOptions) (_result *QueryUniqueDeviceStatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUniqueDeviceStat"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUniqueDeviceStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 去重设备统计
//
// @param request - QueryUniqueDeviceStatRequest
//
// @return QueryUniqueDeviceStatResponse
func (client *Client) QueryUniqueDeviceStat(request *QueryUniqueDeviceStatRequest) (_result *QueryUniqueDeviceStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUniqueDeviceStatResponse{}
	_body, _err := client.QueryUniqueDeviceStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - RemoveTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTagResponse
func (client *Client) RemoveTagWithOptions(request *RemoveTagRequest, runtime *dara.RuntimeOptions) (_result *RemoveTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTag"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - RemoveTagRequest
//
// @return RemoveTagResponse
func (client *Client) RemoveTag(request *RemoveTagRequest) (_result *RemoveTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTagResponse{}
	_body, _err := client.RemoveTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑别名
//
// @param request - UnbindAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAliasResponse
func (client *Client) UnbindAliasWithOptions(request *UnbindAliasRequest, runtime *dara.RuntimeOptions) (_result *UnbindAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.UnbindAll) {
		query["UnbindAll"] = request.UnbindAll
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAlias"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑别名
//
// @param request - UnbindAliasRequest
//
// @return UnbindAliasResponse
func (client *Client) UnbindAlias(request *UnbindAliasRequest) (_result *UnbindAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindAliasResponse{}
	_body, _err := client.UnbindAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑手机号码
//
// @param request - UnbindPhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPhoneResponse
func (client *Client) UnbindPhoneWithOptions(request *UnbindPhoneRequest, runtime *dara.RuntimeOptions) (_result *UnbindPhoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindPhone"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑手机号码
//
// @param request - UnbindPhoneRequest
//
// @return UnbindPhoneResponse
func (client *Client) UnbindPhone(request *UnbindPhoneRequest) (_result *UnbindPhoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindPhoneResponse{}
	_body, _err := client.UnbindPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定标签
//
// @param request - UnbindTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindTagResponse
func (client *Client) UnbindTagWithOptions(request *UnbindTagRequest, runtime *dara.RuntimeOptions) (_result *UnbindTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.ClientKey) {
		query["ClientKey"] = request.ClientKey
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindTag"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定标签
//
// @param request - UnbindTagRequest
//
// @return UnbindTagResponse
func (client *Client) UnbindTag(request *UnbindTagRequest) (_result *UnbindTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindTagResponse{}
	_body, _err := client.UnbindTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
