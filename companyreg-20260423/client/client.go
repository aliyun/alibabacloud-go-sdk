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
		"ap-northeast-1":              dara.String("companyreg.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("companyreg.aliyuncs.com"),
		"ap-south-1":                  dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-1":              dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-2":              dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-3":              dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-5":              dara.String("companyreg.aliyuncs.com"),
		"cn-beijing":                  dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("companyreg.aliyuncs.com"),
		"cn-chengdu":                  dara.String("companyreg.aliyuncs.com"),
		"cn-edge-1":                   dara.String("companyreg.aliyuncs.com"),
		"cn-fujian":                   dara.String("companyreg.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("companyreg.aliyuncs.com"),
		"cn-hongkong":                 dara.String("companyreg.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("companyreg.aliyuncs.com"),
		"cn-huhehaote":                dara.String("companyreg.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("companyreg.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("companyreg.aliyuncs.com"),
		"cn-qingdao":                  dara.String("companyreg.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai":                 dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("companyreg.aliyuncs.com"),
		"cn-wuhan":                    dara.String("companyreg.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("companyreg.aliyuncs.com"),
		"cn-yushanfang":               dara.String("companyreg.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("companyreg.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("companyreg.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("companyreg.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("companyreg.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("companyreg.aliyuncs.com"),
		"eu-central-1":                dara.String("companyreg.aliyuncs.com"),
		"eu-west-1":                   dara.String("companyreg.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("companyreg.aliyuncs.com"),
		"me-east-1":                   dara.String("companyreg.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("companyreg.aliyuncs.com"),
		"us-east-1":                   dara.String("companyreg.aliyuncs.com"),
		"us-west-1":                   dara.String("companyreg.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("companyreg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 查询已备案信息
//
// @param request - QuerySuccessIcpDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySuccessIcpDataResponse
func (client *Client) QuerySuccessIcpDataWithOptions(request *QuerySuccessIcpDataRequest, runtime *dara.RuntimeOptions) (_result *QuerySuccessIcpDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySuccessIcpData"),
		Version:     dara.String("2026-04-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySuccessIcpDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询已备案信息
//
// @param request - QuerySuccessIcpDataRequest
//
// @return QuerySuccessIcpDataResponse
func (client *Client) QuerySuccessIcpData(request *QuerySuccessIcpDataRequest) (_result *QuerySuccessIcpDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySuccessIcpDataResponse{}
	_body, _err := client.QuerySuccessIcpDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
