// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("linkvisual.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("linkvisual.aliyuncs.com"),
		"ap-south-1":                  tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-1":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-2":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-3":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-5":              tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-chengdu":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-edge-1":                   tea.String("linkvisual.aliyuncs.com"),
		"cn-fujian":                   tea.String("linkvisual.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("linkvisual.aliyuncs.com"),
		"cn-hongkong":                 tea.String("linkvisual.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("linkvisual.aliyuncs.com"),
		"cn-huhehaote":                tea.String("linkvisual.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("linkvisual.aliyuncs.com"),
		"cn-qingdao":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-wuhan":                    tea.String("linkvisual.aliyuncs.com"),
		"cn-yushanfang":               tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("linkvisual.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("linkvisual.aliyuncs.com"),
		"eu-central-1":                tea.String("linkvisual.aliyuncs.com"),
		"eu-west-1":                   tea.String("linkvisual.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("linkvisual.aliyuncs.com"),
		"me-east-1":                   tea.String("linkvisual.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("linkvisual.aliyuncs.com"),
		"us-east-1":                   tea.String("linkvisual.aliyuncs.com"),
		"us-west-1":                   tea.String("linkvisual.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkvisual"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
