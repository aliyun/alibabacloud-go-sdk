// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("marketplaceIntl")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("marketplaceintl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 获取推送计量数据结果
//
// @param request - DescribePushMeteringDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePushMeteringDataResponse
func (client *Client) DescribePushMeteringDataWithOptions(request *DescribePushMeteringDataRequest, runtime *dara.RuntimeOptions) (_result *DescribePushMeteringDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PushOrderBizId) {
		query["PushOrderBizId"] = request.PushOrderBizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePushMeteringData"),
		Version:     dara.String("2022-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePushMeteringDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推送计量数据结果
//
// @param request - DescribePushMeteringDataRequest
//
// @return DescribePushMeteringDataResponse
func (client *Client) DescribePushMeteringData(request *DescribePushMeteringDataRequest) (_result *DescribePushMeteringDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePushMeteringDataResponse{}
	_body, _err := client.DescribePushMeteringDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卖家查询实例列表
//
// @param request - DescribeSellerInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSellerInstancesResponse
func (client *Client) DescribeSellerInstancesWithOptions(request *DescribeSellerInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSellerInstancesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
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
		Action:      dara.String("DescribeSellerInstances"),
		Version:     dara.String("2022-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSellerInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卖家查询实例列表
//
// @param request - DescribeSellerInstancesRequest
//
// @return DescribeSellerInstancesResponse
func (client *Client) DescribeSellerInstances(request *DescribeSellerInstancesRequest) (_result *DescribeSellerInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSellerInstancesResponse{}
	_body, _err := client.DescribeSellerInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// isv推送实例消息给用户
//
// @param request - NoticeInstanceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NoticeInstanceUserResponse
func (client *Client) NoticeInstanceUserWithOptions(request *NoticeInstanceUserRequest, runtime *dara.RuntimeOptions) (_result *NoticeInstanceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NoticeParam) {
		body["NoticeParam"] = request.NoticeParam
	}

	if !dara.IsNil(request.NoticeType) {
		body["NoticeType"] = request.NoticeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("NoticeInstanceUser"),
		Version:     dara.String("2022-12-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NoticeInstanceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// isv推送实例消息给用户
//
// @param request - NoticeInstanceUserRequest
//
// @return NoticeInstanceUserResponse
func (client *Client) NoticeInstanceUser(request *NoticeInstanceUserRequest) (_result *NoticeInstanceUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &NoticeInstanceUserResponse{}
	_body, _err := client.NoticeInstanceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际云市场推送计量数据
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.GmtCreate) {
		body["GmtCreate"] = request.GmtCreate
	}

	if !dara.IsNil(request.MeteringData) {
		body["MeteringData"] = request.MeteringData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMeteringData"),
		Version:     dara.String("2022-12-30"),
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
// 国际云市场推送计量数据
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
