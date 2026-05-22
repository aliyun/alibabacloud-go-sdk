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
		"cn-hangzhou":           dara.String("market.aliyuncs.com"),
		"ap-northeast-1":        dara.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        dara.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        dara.String("market.ap-southeast-1.aliyuncs.com"),
		"cn-beijing":            dara.String("market.aliyuncs.com"),
		"cn-chengdu":            dara.String("market.aliyuncs.com"),
		"cn-hongkong":           dara.String("market.aliyuncs.com"),
		"cn-huhehaote":          dara.String("market.aliyuncs.com"),
		"cn-qingdao":            dara.String("market.aliyuncs.com"),
		"cn-shanghai":           dara.String("market.aliyuncs.com"),
		"cn-shenzhen":           dara.String("market.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("market.aliyuncs.com"),
		"eu-central-1":          dara.String("market.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             dara.String("market.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             dara.String("market.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             dara.String("market.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             dara.String("market.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("market.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("market.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("market.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("market.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("market"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 增加STS支持
//
// @param request - ActivateLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicenseWithOptions(request *ActivateLicenseRequest, runtime *dara.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identification) {
		query["Identification"] = request.Identification
	}

	if !dara.IsNil(request.LicenseCode) {
		query["LicenseCode"] = request.LicenseCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateLicense"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加STS支持
//
// @param request - ActivateLicenseRequest
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicense(request *ActivateLicenseRequest) (_result *ActivateLicenseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.ActivateLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置自动续费
//
// @param request - AutoRenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AutoRenewInstanceResponse
func (client *Client) AutoRenewInstanceWithOptions(request *AutoRenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *AutoRenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewCycle) {
		body["AutoRenewCycle"] = request.AutoRenewCycle
	}

	if !dara.IsNil(request.AutoRenewDuration) {
		body["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.OrderBizId) {
		body["OrderBizId"] = request.OrderBizId
	}

	if !dara.IsNil(request.OwnerId) {
		body["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AutoRenewInstance"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AutoRenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置自动续费
//
// @param request - AutoRenewInstanceRequest
//
// @return AutoRenewInstanceResponse
func (client *Client) AutoRenewInstance(request *AutoRenewInstanceRequest) (_result *AutoRenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AutoRenewInstanceResponse{}
	_body, _err := client.AutoRenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 确认查收订阅通知
//
// @param request - ConfirmNotificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmNotificationResponse
func (client *Client) ConfirmNotificationWithOptions(request *ConfirmNotificationRequest, runtime *dara.RuntimeOptions) (_result *ConfirmNotificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NotificationRequestId) {
		query["NotificationRequestId"] = request.NotificationRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmNotification"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmNotificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 确认查收订阅通知
//
// @param request - ConfirmNotificationRequest
//
// @return ConfirmNotificationResponse
func (client *Client) ConfirmNotification(request *ConfirmNotificationRequest) (_result *ConfirmNotificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfirmNotificationResponse{}
	_body, _err := client.ConfirmNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建订单
//
// @param request - CreateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
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

	if !dara.IsNil(request.Commodity) {
		query["Commodity"] = request.Commodity
	}

	if !dara.IsNil(request.OrderSouce) {
		query["OrderSouce"] = request.OrderSouce
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrder"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建订单
//
// @param request - CreateOrderRequest
//
// @return CreateOrderResponse
func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 跨账号角色授权，根据token获取用户信息
//
// @param request - CrossAccountVerifyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CrossAccountVerifyTokenResponse
func (client *Client) CrossAccountVerifyTokenWithOptions(request *CrossAccountVerifyTokenRequest, runtime *dara.RuntimeOptions) (_result *CrossAccountVerifyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CrossAccountVerifyToken"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CrossAccountVerifyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 跨账号角色授权，根据token获取用户信息
//
// @param request - CrossAccountVerifyTokenRequest
//
// @return CrossAccountVerifyTokenResponse
func (client *Client) CrossAccountVerifyToken(request *CrossAccountVerifyTokenRequest) (_result *CrossAccountVerifyTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CrossAccountVerifyTokenResponse{}
	_body, _err := client.CrossAccountVerifyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询API用量
//
// @param request - DescribeApiMeteringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiMeteringResponse
func (client *Client) DescribeApiMeteringWithOptions(request *DescribeApiMeteringRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiMeteringResponse, _err error) {
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
		Action:      dara.String("DescribeApiMetering"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiMeteringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询API用量
//
// @param request - DescribeApiMeteringRequest
//
// @return DescribeApiMeteringResponse
func (client *Client) DescribeApiMetering(request *DescribeApiMeteringRequest) (_result *DescribeApiMeteringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiMeteringResponse{}
	_body, _err := client.DescribeApiMeteringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工作流当前节点信息
//
// @param request - DescribeCurrentNodeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCurrentNodeInfoResponse
func (client *Client) DescribeCurrentNodeInfoWithOptions(request *DescribeCurrentNodeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCurrentNodeInfoResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCurrentNodeInfo"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCurrentNodeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工作流当前节点信息
//
// @param request - DescribeCurrentNodeInfoRequest
//
// @return DescribeCurrentNodeInfoResponse
func (client *Client) DescribeCurrentNodeInfo(request *DescribeCurrentNodeInfoRequest) (_result *DescribeCurrentNodeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCurrentNodeInfoResponse{}
	_body, _err := client.DescribeCurrentNodeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取推广商品
//
// @param request - DescribeDistributionProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistributionProductsResponse
func (client *Client) DescribeDistributionProductsWithOptions(request *DescribeDistributionProductsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDistributionProductsResponse, _err error) {
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
		Action:      dara.String("DescribeDistributionProducts"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDistributionProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取推广商品
//
// @param request - DescribeDistributionProductsRequest
//
// @return DescribeDistributionProductsResponse
func (client *Client) DescribeDistributionProducts(request *DescribeDistributionProductsRequest) (_result *DescribeDistributionProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDistributionProductsResponse{}
	_body, _err := client.DescribeDistributionProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取并生成推广商品-链接
//
// @param tmpReq - DescribeDistributionProductsLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistributionProductsLinkResponse
func (client *Client) DescribeDistributionProductsLinkWithOptions(tmpReq *DescribeDistributionProductsLinkRequest, runtime *dara.RuntimeOptions) (_result *DescribeDistributionProductsLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeDistributionProductsLinkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Codes) {
		request.CodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Codes, dara.String("Codes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CodesShrink) {
		query["Codes"] = request.CodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDistributionProductsLink"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDistributionProductsLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取并生成推广商品-链接
//
// @param request - DescribeDistributionProductsLinkRequest
//
// @return DescribeDistributionProductsLinkResponse
func (client *Client) DescribeDistributionProductsLink(request *DescribeDistributionProductsLinkRequest) (_result *DescribeDistributionProductsLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDistributionProductsLinkResponse{}
	_body, _err := client.DescribeDistributionProductsLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询订阅通知失败列表
//
// @param request - DescribeFailedNotificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFailedNotificationsResponse
func (client *Client) DescribeFailedNotificationsWithOptions(request *DescribeFailedNotificationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFailedNotificationsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFailedNotifications"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFailedNotificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订阅通知失败列表
//
// @param request - DescribeFailedNotificationsRequest
//
// @return DescribeFailedNotificationsResponse
func (client *Client) DescribeFailedNotifications(request *DescribeFailedNotificationsRequest) (_result *DescribeFailedNotificationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFailedNotificationsResponse{}
	_body, _err := client.DescribeFailedNotificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务商侧查询镜像实例信息
//
// @param request - DescribeImageInstanceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageInstanceForIsvResponse
func (client *Client) DescribeImageInstanceForIsvWithOptions(request *DescribeImageInstanceForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageInstanceForIsvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerPk) {
		query["CustomerPk"] = request.CustomerPk
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageInstanceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageInstanceForIsvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商侧查询镜像实例信息
//
// @param request - DescribeImageInstanceForIsvRequest
//
// @return DescribeImageInstanceForIsvResponse
func (client *Client) DescribeImageInstanceForIsv(request *DescribeImageInstanceForIsvRequest) (_result *DescribeImageInstanceForIsvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImageInstanceForIsvResponse{}
	_body, _err := client.DescribeImageInstanceForIsvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
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

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例
//
// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务商侧查询实例信息
//
// @param request - DescribeInstanceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceForIsvResponse
func (client *Client) DescribeInstanceForIsvWithOptions(request *DescribeInstanceForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceForIsvResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceForIsvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商侧查询实例信息
//
// @param request - DescribeInstanceForIsvRequest
//
// @return DescribeInstanceForIsvResponse
func (client *Client) DescribeInstanceForIsv(request *DescribeInstanceForIsvRequest) (_result *DescribeInstanceForIsvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceForIsvResponse{}
	_body, _err := client.DescribeInstanceForIsvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例列表
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Codes) {
		query["Codes"] = request.Codes
	}

	if !dara.IsNil(request.ExceptCodes) {
		query["ExceptCodes"] = request.ExceptCodes
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例列表
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询供应商下的发票信息
//
// @param request - DescribeInvoiceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvoiceForIsvResponse
func (client *Client) DescribeInvoiceForIsvWithOptions(request *DescribeInvoiceForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvoiceForIsvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.InvoiceId) {
		query["InvoiceId"] = request.InvoiceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvoiceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvoiceForIsvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询供应商下的发票信息
//
// @param request - DescribeInvoiceForIsvRequest
//
// @return DescribeInvoiceForIsvResponse
func (client *Client) DescribeInvoiceForIsv(request *DescribeInvoiceForIsvRequest) (_result *DescribeInvoiceForIsvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInvoiceForIsvResponse{}
	_body, _err := client.DescribeInvoiceForIsvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取License
//
// @param request - DescribeLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLicenseResponse
func (client *Client) DescribeLicenseWithOptions(request *DescribeLicenseRequest, runtime *dara.RuntimeOptions) (_result *DescribeLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LicenseCode) {
		query["LicenseCode"] = request.LicenseCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLicense"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取License
//
// @param request - DescribeLicenseRequest
//
// @return DescribeLicenseResponse
func (client *Client) DescribeLicense(request *DescribeLicenseRequest) (_result *DescribeLicenseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLicenseResponse{}
	_body, _err := client.DescribeLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询订单
//
// @param request - DescribeOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrderResponse
func (client *Client) DescribeOrderWithOptions(request *DescribeOrderRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOrder"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订单
//
// @param request - DescribeOrderRequest
//
// @return DescribeOrderResponse
func (client *Client) DescribeOrder(request *DescribeOrderRequest) (_result *DescribeOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOrderResponse{}
	_body, _err := client.DescribeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务商侧查询订单详情
//
// @param request - DescribeOrderForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrderForIsvResponse
func (client *Client) DescribeOrderForIsvWithOptions(request *DescribeOrderForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrderForIsvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOrderForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOrderForIsvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商侧查询订单详情
//
// @param request - DescribeOrderForIsvRequest
//
// @return DescribeOrderForIsvResponse
func (client *Client) DescribeOrderForIsv(request *DescribeOrderForIsvRequest) (_result *DescribeOrderForIsvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOrderForIsvResponse{}
	_body, _err := client.DescribeOrderForIsvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询价格
//
// @param request - DescribePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceResponse
func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Commodity) {
		query["Commodity"] = request.Commodity
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrice"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询价格
//
// @param request - DescribePriceRequest
//
// @return DescribePriceResponse
func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductResponse
func (client *Client) DescribeProductWithOptions(request *DescribeProductRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.QueryDraft) {
		query["QueryDraft"] = request.QueryDraft
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProduct"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeProductRequest
//
// @return DescribeProductResponse
func (client *Client) DescribeProduct(request *DescribeProductRequest) (_result *DescribeProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductResponse{}
	_body, _err := client.DescribeProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductsResponse
func (client *Client) DescribeProductsWithOptions(request *DescribeProductsRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchTerm) {
		query["SearchTerm"] = request.SearchTerm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProducts"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeProductsRequest
//
// @return DescribeProductsResponse
func (client *Client) DescribeProducts(request *DescribeProductsRequest) (_result *DescribeProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductsResponse{}
	_body, _err := client.DescribeProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 附件信息
//
// @param request - DescribeProjectAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectAttachmentsResponse
func (client *Client) DescribeProjectAttachmentsWithOptions(request *DescribeProjectAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectAttachmentsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectAttachments"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 附件信息
//
// @param request - DescribeProjectAttachmentsRequest
//
// @return DescribeProjectAttachmentsResponse
func (client *Client) DescribeProjectAttachments(request *DescribeProjectAttachmentsRequest) (_result *DescribeProjectAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProjectAttachmentsResponse{}
	_body, _err := client.DescribeProjectAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 项目信息
//
// @param request - DescribeProjectInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectInfoResponse
func (client *Client) DescribeProjectInfoWithOptions(request *DescribeProjectInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectInfoResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectInfo"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目信息
//
// @param request - DescribeProjectInfoRequest
//
// @return DescribeProjectInfoResponse
func (client *Client) DescribeProjectInfo(request *DescribeProjectInfoRequest) (_result *DescribeProjectInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProjectInfoResponse{}
	_body, _err := client.DescribeProjectInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目留言信息
//
// @param request - DescribeProjectMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectMessagesResponse
func (client *Client) DescribeProjectMessagesWithOptions(request *DescribeProjectMessagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectMessagesResponse, _err error) {
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

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectMessages"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目留言信息
//
// @param request - DescribeProjectMessagesRequest
//
// @return DescribeProjectMessagesResponse
func (client *Client) DescribeProjectMessages(request *DescribeProjectMessagesRequest) (_result *DescribeProjectMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProjectMessagesResponse{}
	_body, _err := client.DescribeProjectMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目流程节点list
//
// Description:
//
// *
//
// **
//
// @param request - DescribeProjectNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectNodesResponse
func (client *Client) DescribeProjectNodesWithOptions(request *DescribeProjectNodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectNodesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectNodes"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目流程节点list
//
// Description:
//
// *
//
// **
//
// @param request - DescribeProjectNodesRequest
//
// @return DescribeProjectNodesResponse
func (client *Client) DescribeProjectNodes(request *DescribeProjectNodesRequest) (_result *DescribeProjectNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProjectNodesResponse{}
	_body, _err := client.DescribeProjectNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目操作记录列表
//
// @param request - DescribeProjectOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectOperateLogsResponse
func (client *Client) DescribeProjectOperateLogsWithOptions(request *DescribeProjectOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectOperateLogsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectOperateLogs"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectOperateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目操作记录列表
//
// @param request - DescribeProjectOperateLogsRequest
//
// @return DescribeProjectOperateLogsResponse
func (client *Client) DescribeProjectOperateLogs(request *DescribeProjectOperateLogsRequest) (_result *DescribeProjectOperateLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProjectOperateLogsResponse{}
	_body, _err := client.DescribeProjectOperateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 完成当前流程节点
//
// @param request - FinishCurrentProjectNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishCurrentProjectNodeResponse
func (client *Client) FinishCurrentProjectNodeWithOptions(request *FinishCurrentProjectNodeRequest, runtime *dara.RuntimeOptions) (_result *FinishCurrentProjectNodeResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.TemplateForm) {
		query["TemplateForm"] = request.TemplateForm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishCurrentProjectNode"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishCurrentProjectNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完成当前流程节点
//
// @param request - FinishCurrentProjectNodeRequest
//
// @return FinishCurrentProjectNodeResponse
func (client *Client) FinishCurrentProjectNode(request *FinishCurrentProjectNodeRequest) (_result *FinishCurrentProjectNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FinishCurrentProjectNodeResponse{}
	_body, _err := client.FinishCurrentProjectNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发票受理接口
//
// @param request - ModifyInvoiceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInvoiceForIsvResponse
func (client *Client) ModifyInvoiceForIsvWithOptions(request *ModifyInvoiceForIsvRequest, runtime *dara.RuntimeOptions) (_result *ModifyInvoiceForIsvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckNotice) {
		query["CheckNotice"] = request.CheckNotice
	}

	if !dara.IsNil(request.ElectronUrl) {
		query["ElectronUrl"] = request.ElectronUrl
	}

	if !dara.IsNil(request.InvoiceId) {
		query["InvoiceId"] = request.InvoiceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInvoiceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInvoiceForIsvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发票受理接口
//
// @param request - ModifyInvoiceForIsvRequest
//
// @return ModifyInvoiceForIsvResponse
func (client *Client) ModifyInvoiceForIsv(request *ModifyInvoiceForIsvRequest) (_result *ModifyInvoiceForIsvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInvoiceForIsvResponse{}
	_body, _err := client.ModifyInvoiceForIsvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停项目
//
// @param request - PauseProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseProjectResponse
func (client *Client) PauseProjectWithOptions(request *PauseProjectRequest, runtime *dara.RuntimeOptions) (_result *PauseProjectResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseProject"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停项目
//
// @param request - PauseProjectRequest
//
// @return PauseProjectResponse
func (client *Client) PauseProject(request *PauseProjectRequest) (_result *PauseProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseProjectResponse{}
	_body, _err := client.PauseProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 云市场计量推送接口
//
// @param request - PushMeteringDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringDataWithOptions(request *PushMeteringDataRequest, runtime *dara.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Metering) {
		query["Metering"] = request.Metering
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMeteringData"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 云市场计量推送接口
//
// @param request - PushMeteringDataRequest
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringData(request *PushMeteringDataRequest) (_result *PushMeteringDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.PushMeteringDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 按次售卖按量计费
//
// @param request - PushTimesUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushTimesUsageResponse
func (client *Client) PushTimesUsageWithOptions(request *PushTimesUsageRequest, runtime *dara.RuntimeOptions) (_result *PushTimesUsageResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Times) {
		query["Times"] = request.Times
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushTimesUsage"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushTimesUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按次售卖按量计费
//
// @param request - PushTimesUsageRequest
//
// @return PushTimesUsageResponse
func (client *Client) PushTimesUsage(request *PushTimesUsageRequest) (_result *PushTimesUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushTimesUsageResponse{}
	_body, _err := client.PushTimesUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复已暂停的项目
//
// @param request - ResumeProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeProjectResponse
func (client *Client) ResumeProjectWithOptions(request *ResumeProjectRequest, runtime *dara.RuntimeOptions) (_result *ResumeProjectResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeProject"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复已暂停的项目
//
// @param request - ResumeProjectRequest
//
// @return ResumeProjectResponse
func (client *Client) ResumeProject(request *ResumeProjectRequest) (_result *ResumeProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeProjectResponse{}
	_body, _err := client.ResumeProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 当前流程节点回滚到上一步
//
// @param request - RollbackCurrentProjectNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackCurrentProjectNodeResponse
func (client *Client) RollbackCurrentProjectNodeWithOptions(request *RollbackCurrentProjectNodeRequest, runtime *dara.RuntimeOptions) (_result *RollbackCurrentProjectNodeResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackCurrentProjectNode"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackCurrentProjectNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 当前流程节点回滚到上一步
//
// @param request - RollbackCurrentProjectNodeRequest
//
// @return RollbackCurrentProjectNodeResponse
func (client *Client) RollbackCurrentProjectNode(request *RollbackCurrentProjectNodeRequest) (_result *RollbackCurrentProjectNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackCurrentProjectNodeResponse{}
	_body, _err := client.RollbackCurrentProjectNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
