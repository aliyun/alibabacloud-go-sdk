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
		"cn-qingdao":                  dara.String("antiddos.aliyuncs.com"),
		"cn-beijing":                  dara.String("antiddos.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("antiddos-openapi.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":                dara.String("antiddos-openapi.cn-huhehaote.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("antiddos-openapi.cn-wulanchabu.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("antiddos.aliyuncs.com"),
		"cn-shanghai":                 dara.String("antiddos.aliyuncs.com"),
		"cn-nanjing":                  dara.String("antiddos-openapi.cn-hangzhou-cloudstone.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("antiddos.aliyuncs.com"),
		"cn-heyuan":                   dara.String("antiddos-openapi.cn-heyuan.aliyuncs.com"),
		"cn-guangzhou":                dara.String("antiddos-openapi.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":                  dara.String("antiddos-openapi.cn-chengdu.aliyuncs.com"),
		"cn-hongkong":                 dara.String("antiddos.aliyuncs.com"),
		"ap-northeast-1":              dara.String("antiddos-openapi.ap-northeast-1.aliyuncs.com"),
		"ap-northeast-2":              dara.String("antiddos-openapi.ap-northeast-2.aliyuncs.com"),
		"ap-southeast-1":              dara.String("antiddos.aliyuncs.com"),
		"ap-southeast-2":              dara.String("antiddos-openapi.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              dara.String("antiddos-openapi.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":              dara.String("antiddos-openapi.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-6":              dara.String("antiddos-openapi.ap-southeast-6.aliyuncs.com"),
		"us-east-1":                   dara.String("antiddos.aliyuncs.com"),
		"us-west-1":                   dara.String("antiddos.aliyuncs.com"),
		"eu-west-1":                   dara.String("antiddos-openapi.eu-west-1.aliyuncs.com"),
		"eu-central-1":                dara.String("antiddos-openapi.eu-central-1.aliyuncs.com"),
		"ap-south-1":                  dara.String("antiddos-openapi.ap-south-1.aliyuncs.com"),
		"me-east-1":                   dara.String("antiddos-openapi.me-east-1.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("antiddos.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("antiddos.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("antiddos.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("antiddos.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("antiddos.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("antiddos.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("antiddos.aliyuncs.com"),
		"cn-edge-1":                   dara.String("antiddos.aliyuncs.com"),
		"cn-fujian":                   dara.String("antiddos.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("antiddos.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("antiddos.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("antiddos.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("antiddos.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("antiddos.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("antiddos.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("antiddos.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("antiddos.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("antiddos.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("antiddos.aliyuncs.com"),
		"cn-wuhan":                    dara.String("antiddos.aliyuncs.com"),
		"cn-yushanfang":               dara.String("antiddos.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("antiddos.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("antiddos.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("antiddos-openapi.cn-zhangjiakou.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("antiddos.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("antiddos.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("antiddos.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("antiddos-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Queries the configurations of the Anti-DDoS Origin instance that is associated with an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeBgpPackByIp operation to query the configurations of the Anti-DDoS Origin instance that is associated with an asset. The configurations include the basic protection threshold, burstable protection threshold, and expiration time.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeBgpPackByIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBgpPackByIpResponse
func (client *Client) DescribeBgpPackByIpWithOptions(request *DescribeBgpPackByIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeBgpPackByIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBgpPackByIp"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBgpPackByIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the Anti-DDoS Origin instance that is associated with an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeBgpPackByIp operation to query the configurations of the Anti-DDoS Origin instance that is associated with an asset. The configurations include the basic protection threshold, burstable protection threshold, and expiration time.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeBgpPackByIpRequest
//
// @return DescribeBgpPackByIpResponse
func (client *Client) DescribeBgpPackByIp(request *DescribeBgpPackByIpRequest) (_result *DescribeBgpPackByIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBgpPackByIpResponse{}
	_body, _err := client.DescribeBgpPackByIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download link to the traffic data that is captured when a DDoS attack event occurs.
//
// Description:
//
// You can call the DescribeCap operation to query the download link to the traffic data that is captured when a DDoS attack event occurs. You can download the traffic data from the download link and use the data as evidence.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCapResponse
func (client *Client) DescribeCapWithOptions(request *DescribeCapRequest, runtime *dara.RuntimeOptions) (_result *DescribeCapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BegTime) {
		query["BegTime"] = request.BegTime
	}

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetIp) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCap"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download link to the traffic data that is captured when a DDoS attack event occurs.
//
// Description:
//
// You can call the DescribeCap operation to query the download link to the traffic data that is captured when a DDoS attack event occurs. You can download the traffic data from the download link and use the data as evidence.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCapRequest
//
// @return DescribeCapResponse
func (client *Client) DescribeCap(request *DescribeCapRequest) (_result *DescribeCapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCapResponse{}
	_body, _err := client.DescribeCapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are under DDoS attacks in a specific region. The assets are assigned public IP addresses.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosCountResponse
func (client *Client) DescribeDdosCountWithOptions(request *DescribeDdosCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDdosCount"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are under DDoS attacks in a specific region. The assets are assigned public IP addresses.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCountRequest
//
// @return DescribeDdosCountResponse
func (client *Client) DescribeDdosCount(request *DescribeDdosCountRequest) (_result *DescribeDdosCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDdosCountResponse{}
	_body, _err := client.DescribeDdosCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the security credit score of the current Alibaba Cloud account in a specific region.
//
// Description:
//
// You can call the DescribeDdosCredit operation to query the details of the security credit score of the current Alibaba Cloud account in a specific region. The details include the security credit score, security credit level, and the time period after which blackhole filtering is automatically deactivated.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCreditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosCreditResponse
func (client *Client) DescribeDdosCreditWithOptions(request *DescribeDdosCreditRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosCreditResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDdosCredit"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosCreditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the security credit score of the current Alibaba Cloud account in a specific region.
//
// Description:
//
// You can call the DescribeDdosCredit operation to query the details of the security credit score of the current Alibaba Cloud account in a specific region. The details include the security credit score, security credit level, and the time period after which blackhole filtering is automatically deactivated.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosCreditRequest
//
// @return DescribeDdosCreditResponse
func (client *Client) DescribeDdosCredit(request *DescribeDdosCreditRequest) (_result *DescribeDdosCreditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDdosCreditResponse{}
	_body, _err := client.DescribeDdosCreditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS attack events that occur on an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeDdosEventList operation to query the details of the DDoS attack events that occur on an asset by page. The details include the start time, end time, and status of each DDoS attack event.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosEventListResponse
func (client *Client) DescribeDdosEventListWithOptions(request *DescribeDdosEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosEventListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetIp) {
		query["InternetIp"] = request.InternetIp
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryDays) {
		query["QueryDays"] = request.QueryDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDdosEventList"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS attack events that occur on an asset. The asset is assigned a public IP address.
//
// Description:
//
// You can call the DescribeDdosEventList operation to query the details of the DDoS attack events that occur on an asset by page. The details include the start time, end time, and status of each DDoS attack event.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosEventListRequest
//
// @return DescribeDdosEventListResponse
func (client *Client) DescribeDdosEventList(request *DescribeDdosEventListRequest) (_result *DescribeDdosEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDdosEventListResponse{}
	_body, _err := client.DescribeDdosEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// You can call the DescribeDdosThreshold operation to query the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The details include the current traffic scrubbing threshold, maximum traffic scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosThresholdResponse
func (client *Client) DescribeDdosThresholdWithOptions(request *DescribeDdosThresholdRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosThresholdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.DdosType) {
		query["DdosType"] = request.DdosType
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDdosThreshold"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// You can call the DescribeDdosThreshold operation to query the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The details include the current traffic scrubbing threshold, maximum traffic scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosThresholdRequest
//
// @return DescribeDdosThresholdResponse
func (client *Client) DescribeDdosThreshold(request *DescribeDdosThresholdRequest) (_result *DescribeDdosThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDdosThresholdResponse{}
	_body, _err := client.DescribeDdosThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the assets within the current Alibaba Cloud account. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses. This operation is phased out. We recommend that you use the DescribeInstanceIpAddress operation.
//
// Description:
//
// You can call the DescribeInstance operation to query the details of the assets that are within the current Alibaba Cloud account by page. The details include the IDs and IP addresses of the assets, the basic protection thresholds and traffic scrubbing thresholds that are configured for the assets in Anti-DDoS Origin, and whether the assets are associated with Anti-DDoS Origin instances.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.DdosStatus) {
		query["DdosStatus"] = request.DdosStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIp) {
		query["InstanceIp"] = request.InstanceIp
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2017-05-18"),
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
// Queries the details of the assets within the current Alibaba Cloud account. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses. This operation is phased out. We recommend that you use the DescribeInstanceIpAddress operation.
//
// Description:
//
// You can call the DescribeInstance operation to query the details of the assets that are within the current Alibaba Cloud account by page. The details include the IDs and IP addresses of the assets, the basic protection thresholds and traffic scrubbing thresholds that are configured for the assets in Anti-DDoS Origin, and whether the assets are associated with Anti-DDoS Origin instances.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Queries the details of the assets within the current Alibaba Cloud account and the details of the Anti-DDoS Origin instance to which the assets belong. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeInstanceIpAddress operation to query the DDoS mitigation information and the details of the Anti-DDoS Origin instance. The information and the details include the basic protection threshold and traffic scrubbing threshold for the assets, DDoS mitigation status of the assets, ID of the instance, and the mitigation status of the instance.
//
// ## Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceIpAddressResponse
func (client *Client) DescribeInstanceIpAddressWithOptions(request *DescribeInstanceIpAddressRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceIpAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.DdosStatus) {
		query["DdosStatus"] = request.DdosStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIp) {
		query["InstanceIp"] = request.InstanceIp
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceIpAddress"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the assets within the current Alibaba Cloud account and the details of the Anti-DDoS Origin instance to which the assets belong. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeInstanceIpAddress operation to query the DDoS mitigation information and the details of the Anti-DDoS Origin instance. The information and the details include the basic protection threshold and traffic scrubbing threshold for the assets, DDoS mitigation status of the assets, ID of the instance, and the mitigation status of the instance.
//
// ## Limits
//
// You can call this operation up to 200 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceIpAddressRequest
//
// @return DescribeInstanceIpAddressResponse
func (client *Client) DescribeInstanceIpAddress(request *DescribeInstanceIpAddressRequest) (_result *DescribeInstanceIpAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceIpAddressResponse{}
	_body, _err := client.DescribeInstanceIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeIpDdosThreshold operation to query the details of the DDoS mitigation threshold or traffic scrubbing threshold for a specific asset. The details include the current traffic scrubbing threshold, maximum scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpDdosThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpDdosThresholdResponse
func (client *Client) DescribeIpDdosThresholdWithOptions(request *DescribeIpDdosThresholdRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpDdosThresholdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.DdosType) {
		query["DdosType"] = request.DdosType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetIp) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpDdosThreshold"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpDdosThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the DDoS mitigation thresholds or traffic scrubbing thresholds for specified assets. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// If one or more assets of the current Alibaba Cloud account are added to an Anti-DDoS Origin instance, you can call the DescribeIpDdosThreshold operation to query the details of the DDoS mitigation threshold or traffic scrubbing threshold for a specific asset. The details include the current traffic scrubbing threshold, maximum scrubbing threshold, current DDoS mitigation threshold, and maximum DDoS mitigation threshold.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpDdosThresholdRequest
//
// @return DescribeIpDdosThresholdResponse
func (client *Client) DescribeIpDdosThreshold(request *DescribeIpDdosThresholdRequest) (_result *DescribeIpDdosThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpDdosThresholdResponse{}
	_body, _err := client.DescribeIpDdosThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the region to which the public IP address of the asset within the current Alibaba Cloud account belongs. The asset can be an elastic IP address (EIP). The asset can also be an Elastic Compute Service (ECS) instance or Server Load Balancer (SLB) instance that is assigned a public IP address.
//
// Description:
//
// You can call the DescribeIpLocationService operation to query the region of the public IP address for a specified asset that is within the current Alibaba Cloud account. You can also query the details of the Anti-DDoS Origin instance to which the asset is added. The details include the ID and name.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpLocationServiceResponse
func (client *Client) DescribeIpLocationServiceWithOptions(request *DescribeIpLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpLocationServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InternetIp) {
		query["InternetIp"] = request.InternetIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpLocationService"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the region to which the public IP address of the asset within the current Alibaba Cloud account belongs. The asset can be an elastic IP address (EIP). The asset can also be an Elastic Compute Service (ECS) instance or Server Load Balancer (SLB) instance that is assigned a public IP address.
//
// Description:
//
// You can call the DescribeIpLocationService operation to query the region of the public IP address for a specified asset that is within the current Alibaba Cloud account. You can also query the details of the Anti-DDoS Origin instance to which the asset is added. The details include the ID and name.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeIpLocationServiceRequest
//
// @return DescribeIpLocationServiceResponse
func (client *Client) DescribeIpLocationService(request *DescribeIpLocationServiceRequest) (_result *DescribeIpLocationServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpLocationServiceResponse{}
	_body, _err := client.DescribeIpLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions in which Anti-DDoS Origin Basic is available.
//
// Description:
//
// You can call this operation to query information about the regions in which Anti-DDoS Origin Basic is available. The information includes the region ID, region name, and code.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions in which Anti-DDoS Origin Basic is available.
//
// Description:
//
// You can call this operation to query information about the regions in which Anti-DDoS Origin Basic is available. The information includes the region ID, region name, and code.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the scrubbing thresholds for an asset that is assigned a public IP address.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDefenseThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseThresholdResponse
func (client *Client) ModifyDefenseThresholdWithOptions(request *ModifyDefenseThresholdRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseThresholdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bps) {
		query["Bps"] = request.Bps
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetIp) {
		query["InternetIp"] = request.InternetIp
	}

	if !dara.IsNil(request.IsAuto) {
		query["IsAuto"] = request.IsAuto
	}

	if !dara.IsNil(request.Pps) {
		query["Pps"] = request.Pps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseThreshold"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the scrubbing thresholds for an asset that is assigned a public IP address.
//
// Description:
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDefenseThresholdRequest
//
// @return ModifyDefenseThresholdResponse
func (client *Client) ModifyDefenseThreshold(request *ModifyDefenseThresholdRequest) (_result *ModifyDefenseThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseThresholdResponse{}
	_body, _err := client.ModifyDefenseThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the scrubbing thresholds for an asset that is assigned a public IP address. This operation is a synchronous operation that supports Terraform.
//
// Description:
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyIpDefenseThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpDefenseThresholdResponse
func (client *Client) ModifyIpDefenseThresholdWithOptions(request *ModifyIpDefenseThresholdRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpDefenseThresholdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bps) {
		query["Bps"] = request.Bps
	}

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetIp) {
		query["InternetIp"] = request.InternetIp
	}

	if !dara.IsNil(request.IsAuto) {
		query["IsAuto"] = request.IsAuto
	}

	if !dara.IsNil(request.Pps) {
		query["Pps"] = request.Pps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIpDefenseThreshold"),
		Version:     dara.String("2017-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpDefenseThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the scrubbing thresholds for an asset that is assigned a public IP address. This operation is a synchronous operation that supports Terraform.
//
// Description:
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyIpDefenseThresholdRequest
//
// @return ModifyIpDefenseThresholdResponse
func (client *Client) ModifyIpDefenseThreshold(request *ModifyIpDefenseThresholdRequest) (_result *ModifyIpDefenseThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyIpDefenseThresholdResponse{}
	_body, _err := client.ModifyIpDefenseThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
