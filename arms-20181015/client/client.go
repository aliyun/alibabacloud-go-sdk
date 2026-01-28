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
		"ap-northeast-2-pop":          dara.String("arms.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("arms.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("arms.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("arms.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("arms.aliyuncs.com"),
		"cn-edge-1":                   dara.String("arms.aliyuncs.com"),
		"cn-fujian":                   dara.String("arms.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("arms.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("arms.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("arms.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("arms.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("arms.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("arms.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("arms.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("arms.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("arms.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("arms.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("arms.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("arms.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("arms.aliyuncs.com"),
		"cn-wuhan":                    dara.String("arms.aliyuncs.com"),
		"cn-yushanfang":               dara.String("arms.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("arms.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("arms.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("arms.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("arms.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("arms.aliyuncs.com"),
		"me-east-1":                   dara.String("arms.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("arms.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("arms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - ARMSQueryDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ARMSQueryDataSetResponse
func (client *Client) ARMSQueryDataSetWithOptions(request *ARMSQueryDataSetRequest, runtime *dara.RuntimeOptions) (_result *ARMSQueryDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DateStr) {
		query["DateStr"] = request.DateStr
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.HackerUserId) {
		query["HackerUserId"] = request.HackerUserId
	}

	if !dara.IsNil(request.HungryMode) {
		query["HungryMode"] = request.HungryMode
	}

	if !dara.IsNil(request.IntervalInSec) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !dara.IsNil(request.IsDrillDown) {
		query["IsDrillDown"] = request.IsDrillDown
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MaxTime) {
		query["MaxTime"] = request.MaxTime
	}

	if !dara.IsNil(request.Measures) {
		query["Measures"] = request.Measures
	}

	if !dara.IsNil(request.MinTime) {
		query["MinTime"] = request.MinTime
	}

	if !dara.IsNil(request.OptionalDims) {
		query["OptionalDims"] = request.OptionalDims
	}

	if !dara.IsNil(request.OrderByKey) {
		query["OrderByKey"] = request.OrderByKey
	}

	if !dara.IsNil(request.ReduceTail) {
		query["ReduceTail"] = request.ReduceTail
	}

	if !dara.IsNil(request.RequiredDims) {
		query["RequiredDims"] = request.RequiredDims
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ARMSQueryDataSet"),
		Version:     dara.String("2018-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ARMSQueryDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ARMSQueryDataSetRequest
//
// @return ARMSQueryDataSetResponse
func (client *Client) ARMSQueryDataSet(request *ARMSQueryDataSetRequest) (_result *ARMSQueryDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ARMSQueryDataSetResponse{}
	_body, _err := client.ARMSQueryDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MetricQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MetricQueryResponse
func (client *Client) MetricQueryWithOptions(request *MetricQueryRequest, runtime *dara.RuntimeOptions) (_result *MetricQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.HackerUserId) {
		query["HackerUserId"] = request.HackerUserId
	}

	if !dara.IsNil(request.IintervalInSec) {
		query["IintervalInSec"] = request.IintervalInSec
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Measures) {
		query["Measures"] = request.Measures
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MetricQuery"),
		Version:     dara.String("2018-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MetricQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MetricQueryRequest
//
// @return MetricQueryResponse
func (client *Client) MetricQuery(request *MetricQueryRequest) (_result *MetricQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MetricQueryResponse{}
	_body, _err := client.MetricQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
